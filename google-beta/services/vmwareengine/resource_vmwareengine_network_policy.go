// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package vmwareengine

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceVmwareengineNetworkPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceVmwareengineNetworkPolicyCreate,
		Read:   resourceVmwareengineNetworkPolicyRead,
		Update: resourceVmwareengineNetworkPolicyUpdate,
		Delete: resourceVmwareengineNetworkPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceVmwareengineNetworkPolicyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"edge_services_cidr": {
				Type:     schema.TypeString,
				Required: true,
				Description: `IP address range in CIDR notation used to create internet access and external IP access.
An RFC 1918 CIDR block, with a "/26" prefix, is required. The range cannot overlap with any
prefixes either in the consumer VPC network or in use by the private clouds attached to that VPC network.`,
			},
			"location": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The resource name of the location (region) to create the new network policy in.
Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names.
For example: projects/my-project/locations/us-central1`,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The ID of the Network Policy.`,
			},
			"vmware_engine_network": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The relative resource name of the VMware Engine network. Specify the name in the following form:
projects/{project}/locations/{location}/vmwareEngineNetworks/{vmwareEngineNetworkId} where {project}
can either be a project number or a project ID.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `User-provided description for this network policy.`,
			},
			"external_ip": {
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				Description: `Network service that allows External IP addresses to be assigned to VMware workloads.
This service can only be enabled when internetAccess is also enabled.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: `True if the service is enabled; false otherwise.`,
						},
						"state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `State of the service. New values may be added to this enum when appropriate.`,
						},
					},
				},
			},
			"internet_access": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				Description: `Network service that allows VMware workloads to access the internet.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: `True if the service is enabled; false otherwise.`,
						},
						"state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `State of the service. New values may be added to this enum when appropriate.`,
						},
					},
				},
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Creation time of this resource.
A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and
up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".`,
			},
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `System-generated unique identifier for the resource.`,
			},
			"update_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Last updated time of this resource.
A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".`,
			},
			"vmware_engine_network_canonical": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The canonical name of the VMware Engine network in the form:
projects/{project_number}/locations/{location}/vmwareEngineNetworks/{vmwareEngineNetworkId}`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceVmwareengineNetworkPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	edgeServicesCidrProp, err := expandVmwareengineNetworkPolicyEdgeServicesCidr(d.Get("edge_services_cidr"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("edge_services_cidr"); !tpgresource.IsEmptyValue(reflect.ValueOf(edgeServicesCidrProp)) && (ok || !reflect.DeepEqual(v, edgeServicesCidrProp)) {
		obj["edgeServicesCidr"] = edgeServicesCidrProp
	}
	descriptionProp, err := expandVmwareengineNetworkPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	vmwareEngineNetworkProp, err := expandVmwareengineNetworkPolicyVmwareEngineNetwork(d.Get("vmware_engine_network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("vmware_engine_network"); !tpgresource.IsEmptyValue(reflect.ValueOf(vmwareEngineNetworkProp)) && (ok || !reflect.DeepEqual(v, vmwareEngineNetworkProp)) {
		obj["vmwareEngineNetwork"] = vmwareEngineNetworkProp
	}
	internetAccessProp, err := expandVmwareengineNetworkPolicyInternetAccess(d.Get("internet_access"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("internet_access"); !tpgresource.IsEmptyValue(reflect.ValueOf(internetAccessProp)) && (ok || !reflect.DeepEqual(v, internetAccessProp)) {
		obj["internetAccess"] = internetAccessProp
	}
	externalIpProp, err := expandVmwareengineNetworkPolicyExternalIp(d.Get("external_ip"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("external_ip"); !tpgresource.IsEmptyValue(reflect.ValueOf(externalIpProp)) && (ok || !reflect.DeepEqual(v, externalIpProp)) {
		obj["externalIp"] = externalIpProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{VmwareengineBasePath}}projects/{{project}}/locations/{{location}}/networkPolicies?networkPolicyId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new NetworkPolicy: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for NetworkPolicy: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating NetworkPolicy: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/networkPolicies/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = VmwareengineOperationWaitTime(
		config, res, project, "Creating NetworkPolicy", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create NetworkPolicy: %s", err)
	}

	log.Printf("[DEBUG] Finished creating NetworkPolicy %q: %#v", d.Id(), res)

	return resourceVmwareengineNetworkPolicyRead(d, meta)
}

func resourceVmwareengineNetworkPolicyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{VmwareengineBasePath}}projects/{{project}}/locations/{{location}}/networkPolicies/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for NetworkPolicy: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("VmwareengineNetworkPolicy %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading NetworkPolicy: %s", err)
	}

	if err := d.Set("create_time", flattenVmwareengineNetworkPolicyCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkPolicy: %s", err)
	}
	if err := d.Set("update_time", flattenVmwareengineNetworkPolicyUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkPolicy: %s", err)
	}
	if err := d.Set("uid", flattenVmwareengineNetworkPolicyUid(res["uid"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkPolicy: %s", err)
	}
	if err := d.Set("vmware_engine_network_canonical", flattenVmwareengineNetworkPolicyVmwareEngineNetworkCanonical(res["vmwareEngineNetworkCanonical"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkPolicy: %s", err)
	}
	if err := d.Set("edge_services_cidr", flattenVmwareengineNetworkPolicyEdgeServicesCidr(res["edgeServicesCidr"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkPolicy: %s", err)
	}
	if err := d.Set("description", flattenVmwareengineNetworkPolicyDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkPolicy: %s", err)
	}
	if err := d.Set("vmware_engine_network", flattenVmwareengineNetworkPolicyVmwareEngineNetwork(res["vmwareEngineNetwork"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkPolicy: %s", err)
	}
	if err := d.Set("internet_access", flattenVmwareengineNetworkPolicyInternetAccess(res["internetAccess"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkPolicy: %s", err)
	}
	if err := d.Set("external_ip", flattenVmwareengineNetworkPolicyExternalIp(res["externalIp"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkPolicy: %s", err)
	}

	return nil
}

func resourceVmwareengineNetworkPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for NetworkPolicy: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	edgeServicesCidrProp, err := expandVmwareengineNetworkPolicyEdgeServicesCidr(d.Get("edge_services_cidr"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("edge_services_cidr"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, edgeServicesCidrProp)) {
		obj["edgeServicesCidr"] = edgeServicesCidrProp
	}
	descriptionProp, err := expandVmwareengineNetworkPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	internetAccessProp, err := expandVmwareengineNetworkPolicyInternetAccess(d.Get("internet_access"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("internet_access"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, internetAccessProp)) {
		obj["internetAccess"] = internetAccessProp
	}
	externalIpProp, err := expandVmwareengineNetworkPolicyExternalIp(d.Get("external_ip"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("external_ip"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, externalIpProp)) {
		obj["externalIp"] = externalIpProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{VmwareengineBasePath}}projects/{{project}}/locations/{{location}}/networkPolicies/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating NetworkPolicy %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
	})

	if err != nil {
		return fmt.Errorf("Error updating NetworkPolicy %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating NetworkPolicy %q: %#v", d.Id(), res)
	}

	err = VmwareengineOperationWaitTime(
		config, res, project, "Updating NetworkPolicy", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceVmwareengineNetworkPolicyRead(d, meta)
}

func resourceVmwareengineNetworkPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for NetworkPolicy: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{VmwareengineBasePath}}projects/{{project}}/locations/{{location}}/networkPolicies/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting NetworkPolicy %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "NetworkPolicy")
	}

	err = VmwareengineOperationWaitTime(
		config, res, project, "Deleting NetworkPolicy", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting NetworkPolicy %q: %#v", d.Id(), res)
	return nil
}

func resourceVmwareengineNetworkPolicyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/networkPolicies/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/networkPolicies/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenVmwareengineNetworkPolicyCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareengineNetworkPolicyUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareengineNetworkPolicyUid(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareengineNetworkPolicyVmwareEngineNetworkCanonical(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareengineNetworkPolicyEdgeServicesCidr(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareengineNetworkPolicyDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareengineNetworkPolicyVmwareEngineNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareengineNetworkPolicyInternetAccess(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["enabled"] =
		flattenVmwareengineNetworkPolicyInternetAccessEnabled(original["enabled"], d, config)
	transformed["state"] =
		flattenVmwareengineNetworkPolicyInternetAccessState(original["state"], d, config)
	return []interface{}{transformed}
}
func flattenVmwareengineNetworkPolicyInternetAccessEnabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareengineNetworkPolicyInternetAccessState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareengineNetworkPolicyExternalIp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["enabled"] =
		flattenVmwareengineNetworkPolicyExternalIpEnabled(original["enabled"], d, config)
	transformed["state"] =
		flattenVmwareengineNetworkPolicyExternalIpState(original["state"], d, config)
	return []interface{}{transformed}
}
func flattenVmwareengineNetworkPolicyExternalIpEnabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareengineNetworkPolicyExternalIpState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandVmwareengineNetworkPolicyEdgeServicesCidr(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVmwareengineNetworkPolicyDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVmwareengineNetworkPolicyVmwareEngineNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVmwareengineNetworkPolicyInternetAccess(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnabled, err := expandVmwareengineNetworkPolicyInternetAccessEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["enabled"] = transformedEnabled
	}

	transformedState, err := expandVmwareengineNetworkPolicyInternetAccessState(original["state"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedState); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["state"] = transformedState
	}

	return transformed, nil
}

func expandVmwareengineNetworkPolicyInternetAccessEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVmwareengineNetworkPolicyInternetAccessState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVmwareengineNetworkPolicyExternalIp(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnabled, err := expandVmwareengineNetworkPolicyExternalIpEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["enabled"] = transformedEnabled
	}

	transformedState, err := expandVmwareengineNetworkPolicyExternalIpState(original["state"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedState); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["state"] = transformedState
	}

	return transformed, nil
}

func expandVmwareengineNetworkPolicyExternalIpEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVmwareengineNetworkPolicyExternalIpState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
