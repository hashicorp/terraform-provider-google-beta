// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/networksecurity/MirroringDeploymentGroup.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package networksecurity

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceNetworkSecurityMirroringDeploymentGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkSecurityMirroringDeploymentGroupCreate,
		Read:   resourceNetworkSecurityMirroringDeploymentGroupRead,
		Update: resourceNetworkSecurityMirroringDeploymentGroupUpdate,
		Delete: resourceNetworkSecurityMirroringDeploymentGroupDelete,

		Importer: &schema.ResourceImporter{
			State: resourceNetworkSecurityMirroringDeploymentGroupImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The cloud location of the deployment group, currently restricted to 'global'.`,
			},
			"mirroring_deployment_group_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The ID to use for the new deployment group, which will become the final
component of the deployment group's resource name.`,
			},
			"network": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The network that will be used for all child deployments, for example:
'projects/{project}/global/networks/{network}'.
See https://google.aip.dev/124.`,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `User-provided description of the deployment group.
Used as additional context for the deployment group.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Labels are key/value pairs that help to organize and filter resources.

**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"connected_endpoint_groups": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `The list of endpoint groups that are connected to this resource.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
							Description: `The connected endpoint group's resource name, for example:
'projects/123456789/locations/global/mirroringEndpointGroups/my-eg'.
See https://google.aip.dev/124.`,
						},
					},
				},
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The timestamp when the resource was created.
See https://google.aip.dev/148#timestamps.`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"locations": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: `The list of locations where the deployment group is present.`,
				Elem:        networksecurityMirroringDeploymentGroupLocationsSchema(),
				// Default schema.HashSchema is used.
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The resource name of this deployment group, for example:
'projects/123456789/locations/global/mirroringDeploymentGroups/my-dg'.
See https://google.aip.dev/122 for more details.`,
			},
			"reconciling": {
				Type:     schema.TypeBool,
				Computed: true,
				Description: `The current state of the resource does not match the user's intended state,
and the system is working to reconcile them. This is part of the normal
operation (e.g. adding a new deployment to the group)
See https://google.aip.dev/128.`,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The current state of the deployment group.
See https://google.aip.dev/216.
Possible values:
STATE_UNSPECIFIED
ACTIVE
CREATING
DELETING`,
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"update_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The timestamp when the resource was most recently updated.
See https://google.aip.dev/148#timestamps.`,
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

func networksecurityMirroringDeploymentGroupLocationsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The cloud location, e.g. 'us-central1-a' or 'asia-south1-b'.`,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The current state of the association in this location.
Possible values:
STATE_UNSPECIFIED
ACTIVE
OUT_OF_SYNC`,
			},
		},
	}
}

func resourceNetworkSecurityMirroringDeploymentGroupCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	networkProp, err := expandNetworkSecurityMirroringDeploymentGroupNetwork(d.Get("network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	descriptionProp, err := expandNetworkSecurityMirroringDeploymentGroupDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandNetworkSecurityMirroringDeploymentGroupEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/mirroringDeploymentGroups?mirroringDeploymentGroupId={{mirroring_deployment_group_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new MirroringDeploymentGroup: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for MirroringDeploymentGroup: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating MirroringDeploymentGroup: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/mirroringDeploymentGroups/{{mirroring_deployment_group_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = NetworkSecurityOperationWaitTime(
		config, res, project, "Creating MirroringDeploymentGroup", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create MirroringDeploymentGroup: %s", err)
	}

	log.Printf("[DEBUG] Finished creating MirroringDeploymentGroup %q: %#v", d.Id(), res)

	return resourceNetworkSecurityMirroringDeploymentGroupRead(d, meta)
}

func resourceNetworkSecurityMirroringDeploymentGroupRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/mirroringDeploymentGroups/{{mirroring_deployment_group_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for MirroringDeploymentGroup: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("NetworkSecurityMirroringDeploymentGroup %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading MirroringDeploymentGroup: %s", err)
	}

	if err := d.Set("name", flattenNetworkSecurityMirroringDeploymentGroupName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading MirroringDeploymentGroup: %s", err)
	}
	if err := d.Set("create_time", flattenNetworkSecurityMirroringDeploymentGroupCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading MirroringDeploymentGroup: %s", err)
	}
	if err := d.Set("update_time", flattenNetworkSecurityMirroringDeploymentGroupUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading MirroringDeploymentGroup: %s", err)
	}
	if err := d.Set("labels", flattenNetworkSecurityMirroringDeploymentGroupLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading MirroringDeploymentGroup: %s", err)
	}
	if err := d.Set("network", flattenNetworkSecurityMirroringDeploymentGroupNetwork(res["network"], d, config)); err != nil {
		return fmt.Errorf("Error reading MirroringDeploymentGroup: %s", err)
	}
	if err := d.Set("connected_endpoint_groups", flattenNetworkSecurityMirroringDeploymentGroupConnectedEndpointGroups(res["connectedEndpointGroups"], d, config)); err != nil {
		return fmt.Errorf("Error reading MirroringDeploymentGroup: %s", err)
	}
	if err := d.Set("state", flattenNetworkSecurityMirroringDeploymentGroupState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading MirroringDeploymentGroup: %s", err)
	}
	if err := d.Set("reconciling", flattenNetworkSecurityMirroringDeploymentGroupReconciling(res["reconciling"], d, config)); err != nil {
		return fmt.Errorf("Error reading MirroringDeploymentGroup: %s", err)
	}
	if err := d.Set("description", flattenNetworkSecurityMirroringDeploymentGroupDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading MirroringDeploymentGroup: %s", err)
	}
	if err := d.Set("locations", flattenNetworkSecurityMirroringDeploymentGroupLocations(res["locations"], d, config)); err != nil {
		return fmt.Errorf("Error reading MirroringDeploymentGroup: %s", err)
	}
	if err := d.Set("terraform_labels", flattenNetworkSecurityMirroringDeploymentGroupTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading MirroringDeploymentGroup: %s", err)
	}
	if err := d.Set("effective_labels", flattenNetworkSecurityMirroringDeploymentGroupEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading MirroringDeploymentGroup: %s", err)
	}

	return nil
}

func resourceNetworkSecurityMirroringDeploymentGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for MirroringDeploymentGroup: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkSecurityMirroringDeploymentGroupDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandNetworkSecurityMirroringDeploymentGroupEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/mirroringDeploymentGroups/{{mirroring_deployment_group_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating MirroringDeploymentGroup %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("effective_labels") {
		updateMask = append(updateMask, "labels")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// if updateMask is empty we are not updating anything so skip the post
	if len(updateMask) > 0 {
		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "PATCH",
			Project:   billingProject,
			RawURL:    url,
			UserAgent: userAgent,
			Body:      obj,
			Timeout:   d.Timeout(schema.TimeoutUpdate),
			Headers:   headers,
		})

		if err != nil {
			return fmt.Errorf("Error updating MirroringDeploymentGroup %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating MirroringDeploymentGroup %q: %#v", d.Id(), res)
		}

		err = NetworkSecurityOperationWaitTime(
			config, res, project, "Updating MirroringDeploymentGroup", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceNetworkSecurityMirroringDeploymentGroupRead(d, meta)
}

func resourceNetworkSecurityMirroringDeploymentGroupDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for MirroringDeploymentGroup: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/mirroringDeploymentGroups/{{mirroring_deployment_group_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting MirroringDeploymentGroup %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "MirroringDeploymentGroup")
	}

	err = NetworkSecurityOperationWaitTime(
		config, res, project, "Deleting MirroringDeploymentGroup", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting MirroringDeploymentGroup %q: %#v", d.Id(), res)
	return nil
}

func resourceNetworkSecurityMirroringDeploymentGroupImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/mirroringDeploymentGroups/(?P<mirroring_deployment_group_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<mirroring_deployment_group_id>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<mirroring_deployment_group_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/mirroringDeploymentGroups/{{mirroring_deployment_group_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNetworkSecurityMirroringDeploymentGroupName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityMirroringDeploymentGroupCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityMirroringDeploymentGroupUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityMirroringDeploymentGroupLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenNetworkSecurityMirroringDeploymentGroupNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityMirroringDeploymentGroupConnectedEndpointGroups(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"name": flattenNetworkSecurityMirroringDeploymentGroupConnectedEndpointGroupsName(original["name"], d, config),
		})
	}
	return transformed
}
func flattenNetworkSecurityMirroringDeploymentGroupConnectedEndpointGroupsName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityMirroringDeploymentGroupState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityMirroringDeploymentGroupReconciling(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityMirroringDeploymentGroupDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityMirroringDeploymentGroupLocations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(schema.HashResource(networksecurityMirroringDeploymentGroupLocationsSchema()), []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"location": flattenNetworkSecurityMirroringDeploymentGroupLocationsLocation(original["location"], d, config),
			"state":    flattenNetworkSecurityMirroringDeploymentGroupLocationsState(original["state"], d, config),
		})
	}
	return transformed
}
func flattenNetworkSecurityMirroringDeploymentGroupLocationsLocation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityMirroringDeploymentGroupLocationsState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityMirroringDeploymentGroupTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("terraform_labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenNetworkSecurityMirroringDeploymentGroupEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNetworkSecurityMirroringDeploymentGroupNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityMirroringDeploymentGroupDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityMirroringDeploymentGroupEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
