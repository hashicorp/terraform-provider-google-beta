// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute

import (
	"fmt"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"

	compute "google.golang.org/api/compute/v0.beta"
)

func computeInstanceDeleteAccessConfigs(d *schema.ResourceData, config *transport_tpg.Config, instNetworkInterface *compute.NetworkInterface, project, zone, userAgent, instanceName string) error {
	// Delete any accessConfig that currently exists in instNetworkInterface
	for _, ac := range instNetworkInterface.AccessConfigs {
		op, err := config.NewComputeClient(userAgent).Instances.DeleteAccessConfig(
			project, zone, instanceName, ac.Name, instNetworkInterface.Name).Do()
		if err != nil {
			return fmt.Errorf("Error deleting old access_config: %s", err)
		}
		opErr := ComputeOperationWaitTime(config, op, project, "old access_config to delete", userAgent, d.Timeout(schema.TimeoutUpdate))
		if opErr != nil {
			return opErr
		}
	}
	return nil
}

func computeInstanceAddAccessConfigs(d *schema.ResourceData, config *transport_tpg.Config, instNetworkInterface *compute.NetworkInterface, accessConfigs []*compute.AccessConfig, project, zone, userAgent, instanceName string) error {
	// Create new ones
	for _, ac := range accessConfigs {
		op, err := config.NewComputeClient(userAgent).Instances.AddAccessConfig(project, zone, instanceName, instNetworkInterface.Name, ac).Do()
		if err != nil {
			return fmt.Errorf("Error adding new access_config: %s", err)
		}
		opErr := ComputeOperationWaitTime(config, op, project, "new access_config to add", userAgent, d.Timeout(schema.TimeoutUpdate))
		if opErr != nil {
			return opErr
		}
	}
	return nil
}

func computeInstanceCreateUpdateWhileStoppedCall(d *schema.ResourceData, config *transport_tpg.Config, networkInterfacePatchObj *compute.NetworkInterface, accessConfigs []*compute.AccessConfig, accessConfigsHaveChanged bool, index int, project, zone, userAgent, instanceName string) func(inst *compute.Instance) error {

	// Access configs' ip changes when the instance stops invalidating our fingerprint
	// expect caller to re-validate instance before calling patch this is why we expect
	// instance to be passed in
	return func(instance *compute.Instance) error {

		instNetworkInterface := instance.NetworkInterfaces[index]
		networkInterfacePatchObj.Fingerprint = instNetworkInterface.Fingerprint

		// Access config can run into some issues since we can't tell the difference between
		// the users declared intent (config within their hcl file) and what we have inferred from the
		// server (terraform state). Access configs contain an ip subproperty that can be incompatible
		// with the subnetwork/network we are transitioning to. Due to this we only change access
		// configs if we notice the configuration (user intent) changes.
		if accessConfigsHaveChanged {
			err := computeInstanceDeleteAccessConfigs(d, config, instNetworkInterface, project, zone, userAgent, instanceName)
			if err != nil {
				return err
			}
		}

		op, err := config.NewComputeClient(userAgent).Instances.UpdateNetworkInterface(project, zone, instanceName, instNetworkInterface.Name, networkInterfacePatchObj).Do()
		if err != nil {
			return errwrap.Wrapf("Error updating network interface: {{err}}", err)
		}
		opErr := ComputeOperationWaitTime(config, op, project, "network interface to update", userAgent, d.Timeout(schema.TimeoutUpdate))
		if opErr != nil {
			return opErr
		}

		if accessConfigsHaveChanged {
			err := computeInstanceAddAccessConfigs(d, config, instNetworkInterface, accessConfigs, project, zone, userAgent, instanceName)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func computeInstanceAddSecurityPolicy(d *schema.ResourceData, config *transport_tpg.Config, securityPolicyWithNics map[string][]string, project, zone, userAgent, instanceName string) error {
	for sp, nics := range securityPolicyWithNics {
		req := &compute.InstancesSetSecurityPolicyRequest{
			NetworkInterfaces: nics,
			SecurityPolicy:    sp,
		}
		op, err := config.NewComputeClient(userAgent).Instances.SetSecurityPolicy(project, zone, instanceName, req).Do()
		if err != nil {
			return fmt.Errorf("Error adding security policy: %s", err)
		}
		opErr := ComputeOperationWaitTime(config, op, project, "security_policy to add", userAgent, d.Timeout(schema.TimeoutUpdate))
		if opErr != nil {
			return opErr
		}
	}

	return nil
}

func computeInstanceMapSecurityPoliciesCreate(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string][]string, error) {
	securityPolicies := make(map[string][]string)
	configs := d.Get("network_interface").([]interface{})
	for i, raw := range configs {
		data := raw.(map[string]interface{})
		secPolicy := data["security_policy"].(string)
		err := validateSecurityPolicy(data)
		if err != nil {
			return securityPolicies, err
		}

		if secPolicy != "" {
			// Network interfaces use the nicN naming format and is only know after the instance is created.
			nicName := fmt.Sprintf("nic%d", i)
			securityPolicies[secPolicy] = append(securityPolicies[secPolicy], nicName)
		}
	}

	return securityPolicies, nil
}

func computeInstanceMapSecurityPoliciesUpdate(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string][]string, error) {
	securityPolicies := make(map[string][]string)
	configs := d.Get("network_interface").([]interface{})
	for i, raw := range configs {
		data := raw.(map[string]interface{})
		secPolicy := data["security_policy"].(string)
		err := validateSecurityPolicy(data)
		if err != nil {
			return securityPolicies, err
		}

		// Network interfaces use the nicN naming format and is only know after the instance is created.
		nicName := fmt.Sprintf("nic%d", i)
		// To cleanup the security policy from the interface we should send something like this on the api: {"":[nic0, nic1]}
		securityPolicies[secPolicy] = append(securityPolicies[secPolicy], nicName)
	}

	return securityPolicies, nil
}

func validateSecurityPolicy(rawNetworkInterface map[string]interface{}) error {
	acessConfigs := expandAccessConfigs(rawNetworkInterface["access_config"].([]interface{}))
	ipv6AccessConfigs := expandIpv6AccessConfigs(rawNetworkInterface["ipv6_access_config"].([]interface{}))
	secPolicy := rawNetworkInterface["security_policy"].(string)

	if secPolicy != "" && len(acessConfigs) == 0 && len(ipv6AccessConfigs) == 0 {
		return fmt.Errorf("Error setting security policy to the instance since at least one access config must exist")
	}

	return nil
}
