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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/networksecurity/InterceptEndpointGroupAssociation.yaml
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

func ResourceNetworkSecurityInterceptEndpointGroupAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkSecurityInterceptEndpointGroupAssociationCreate,
		Read:   resourceNetworkSecurityInterceptEndpointGroupAssociationRead,
		Update: resourceNetworkSecurityInterceptEndpointGroupAssociationUpdate,
		Delete: resourceNetworkSecurityInterceptEndpointGroupAssociationDelete,

		Importer: &schema.ResourceImporter{
			State: resourceNetworkSecurityInterceptEndpointGroupAssociationImport,
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
			"intercept_endpoint_group": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The endpoint group that this association is connected to, for example:
'projects/123456789/locations/global/interceptEndpointGroups/my-eg'.
See https://google.aip.dev/124.`,
			},
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The cloud location of the association, currently restricted to 'global'.`,
			},
			"network": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The VPC network that is associated. for example:
'projects/123456789/global/networks/my-network'.
See https://google.aip.dev/124.`,
			},
			"intercept_endpoint_group_association_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The ID to use for the new association, which will become the final
component of the endpoint group's resource name. If not provided, the
server will generate a unique ID.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Labels are key/value pairs that help to organize and filter resources.

**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
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
				Type:     schema.TypeSet,
				Computed: true,
				Description: `The list of locations where the association is configured. This information
is retrieved from the linked endpoint group.`,
				Elem: networksecurityInterceptEndpointGroupAssociationLocationsSchema(),
				// Default schema.HashSchema is used.
			},
			"locations_details": {
				Type:       schema.TypeList,
				Computed:   true,
				Deprecated: "`locationsDetails` is deprecated and will be removed in a future major release. Use `locations` instead.",
				Description: `The list of locations where the association is present. This information
is retrieved from the linked endpoint group, and not configured as part
of the association itself.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"location": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The cloud location, e.g. 'us-central1-a' or 'asia-south1'.`,
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
				},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The resource name of this endpoint group association, for example:
'projects/123456789/locations/global/interceptEndpointGroupAssociations/my-eg-association'.
See https://google.aip.dev/122 for more details.`,
			},
			"reconciling": {
				Type:     schema.TypeBool,
				Computed: true,
				Description: `The current state of the resource does not match the user's intended state,
and the system is working to reconcile them. This part of the normal
operation (e.g. adding a new location to the target deployment group).
See https://google.aip.dev/128.`,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Current state of the endpoint group association.
Possible values:
STATE_UNSPECIFIED
ACTIVE
CREATING
DELETING
CLOSED
OUT_OF_SYNC
DELETE_FAILED`,
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

func networksecurityInterceptEndpointGroupAssociationLocationsSchema() *schema.Resource {
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

func resourceNetworkSecurityInterceptEndpointGroupAssociationCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	interceptEndpointGroupProp, err := expandNetworkSecurityInterceptEndpointGroupAssociationInterceptEndpointGroup(d.Get("intercept_endpoint_group"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("intercept_endpoint_group"); !tpgresource.IsEmptyValue(reflect.ValueOf(interceptEndpointGroupProp)) && (ok || !reflect.DeepEqual(v, interceptEndpointGroupProp)) {
		obj["interceptEndpointGroup"] = interceptEndpointGroupProp
	}
	networkProp, err := expandNetworkSecurityInterceptEndpointGroupAssociationNetwork(d.Get("network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	labelsProp, err := expandNetworkSecurityInterceptEndpointGroupAssociationEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/interceptEndpointGroupAssociations?interceptEndpointGroupAssociationId={{intercept_endpoint_group_association_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new InterceptEndpointGroupAssociation: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for InterceptEndpointGroupAssociation: %s", err)
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
		return fmt.Errorf("Error creating InterceptEndpointGroupAssociation: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/interceptEndpointGroupAssociations/{{intercept_endpoint_group_association_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = NetworkSecurityOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating InterceptEndpointGroupAssociation", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create InterceptEndpointGroupAssociation: %s", err)
	}

	if err := d.Set("name", flattenNetworkSecurityInterceptEndpointGroupAssociationName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/interceptEndpointGroupAssociations/{{intercept_endpoint_group_association_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating InterceptEndpointGroupAssociation %q: %#v", d.Id(), res)

	return resourceNetworkSecurityInterceptEndpointGroupAssociationRead(d, meta)
}

func resourceNetworkSecurityInterceptEndpointGroupAssociationRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/interceptEndpointGroupAssociations/{{intercept_endpoint_group_association_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for InterceptEndpointGroupAssociation: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("NetworkSecurityInterceptEndpointGroupAssociation %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading InterceptEndpointGroupAssociation: %s", err)
	}

	if err := d.Set("name", flattenNetworkSecurityInterceptEndpointGroupAssociationName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterceptEndpointGroupAssociation: %s", err)
	}
	if err := d.Set("create_time", flattenNetworkSecurityInterceptEndpointGroupAssociationCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterceptEndpointGroupAssociation: %s", err)
	}
	if err := d.Set("update_time", flattenNetworkSecurityInterceptEndpointGroupAssociationUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterceptEndpointGroupAssociation: %s", err)
	}
	if err := d.Set("labels", flattenNetworkSecurityInterceptEndpointGroupAssociationLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterceptEndpointGroupAssociation: %s", err)
	}
	if err := d.Set("intercept_endpoint_group", flattenNetworkSecurityInterceptEndpointGroupAssociationInterceptEndpointGroup(res["interceptEndpointGroup"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterceptEndpointGroupAssociation: %s", err)
	}
	if err := d.Set("network", flattenNetworkSecurityInterceptEndpointGroupAssociationNetwork(res["network"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterceptEndpointGroupAssociation: %s", err)
	}
	if err := d.Set("locations_details", flattenNetworkSecurityInterceptEndpointGroupAssociationLocationsDetails(res["locationsDetails"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterceptEndpointGroupAssociation: %s", err)
	}
	if err := d.Set("state", flattenNetworkSecurityInterceptEndpointGroupAssociationState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterceptEndpointGroupAssociation: %s", err)
	}
	if err := d.Set("reconciling", flattenNetworkSecurityInterceptEndpointGroupAssociationReconciling(res["reconciling"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterceptEndpointGroupAssociation: %s", err)
	}
	if err := d.Set("locations", flattenNetworkSecurityInterceptEndpointGroupAssociationLocations(res["locations"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterceptEndpointGroupAssociation: %s", err)
	}
	if err := d.Set("terraform_labels", flattenNetworkSecurityInterceptEndpointGroupAssociationTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterceptEndpointGroupAssociation: %s", err)
	}
	if err := d.Set("effective_labels", flattenNetworkSecurityInterceptEndpointGroupAssociationEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterceptEndpointGroupAssociation: %s", err)
	}

	return nil
}

func resourceNetworkSecurityInterceptEndpointGroupAssociationUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for InterceptEndpointGroupAssociation: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	labelsProp, err := expandNetworkSecurityInterceptEndpointGroupAssociationEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/interceptEndpointGroupAssociations/{{intercept_endpoint_group_association_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating InterceptEndpointGroupAssociation %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

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
			return fmt.Errorf("Error updating InterceptEndpointGroupAssociation %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating InterceptEndpointGroupAssociation %q: %#v", d.Id(), res)
		}

		err = NetworkSecurityOperationWaitTime(
			config, res, project, "Updating InterceptEndpointGroupAssociation", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceNetworkSecurityInterceptEndpointGroupAssociationRead(d, meta)
}

func resourceNetworkSecurityInterceptEndpointGroupAssociationDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for InterceptEndpointGroupAssociation: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/interceptEndpointGroupAssociations/{{intercept_endpoint_group_association_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting InterceptEndpointGroupAssociation %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "InterceptEndpointGroupAssociation")
	}

	err = NetworkSecurityOperationWaitTime(
		config, res, project, "Deleting InterceptEndpointGroupAssociation", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting InterceptEndpointGroupAssociation %q: %#v", d.Id(), res)
	return nil
}

func resourceNetworkSecurityInterceptEndpointGroupAssociationImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/interceptEndpointGroupAssociations/(?P<intercept_endpoint_group_association_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<intercept_endpoint_group_association_id>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<intercept_endpoint_group_association_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/interceptEndpointGroupAssociations/{{intercept_endpoint_group_association_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNetworkSecurityInterceptEndpointGroupAssociationName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityInterceptEndpointGroupAssociationCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityInterceptEndpointGroupAssociationUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityInterceptEndpointGroupAssociationLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenNetworkSecurityInterceptEndpointGroupAssociationInterceptEndpointGroup(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityInterceptEndpointGroupAssociationNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityInterceptEndpointGroupAssociationLocationsDetails(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"location": flattenNetworkSecurityInterceptEndpointGroupAssociationLocationsDetailsLocation(original["location"], d, config),
			"state":    flattenNetworkSecurityInterceptEndpointGroupAssociationLocationsDetailsState(original["state"], d, config),
		})
	}
	return transformed
}
func flattenNetworkSecurityInterceptEndpointGroupAssociationLocationsDetailsLocation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityInterceptEndpointGroupAssociationLocationsDetailsState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityInterceptEndpointGroupAssociationState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityInterceptEndpointGroupAssociationReconciling(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityInterceptEndpointGroupAssociationLocations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(schema.HashResource(networksecurityInterceptEndpointGroupAssociationLocationsSchema()), []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"location": flattenNetworkSecurityInterceptEndpointGroupAssociationLocationsLocation(original["location"], d, config),
			"state":    flattenNetworkSecurityInterceptEndpointGroupAssociationLocationsState(original["state"], d, config),
		})
	}
	return transformed
}
func flattenNetworkSecurityInterceptEndpointGroupAssociationLocationsLocation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityInterceptEndpointGroupAssociationLocationsState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityInterceptEndpointGroupAssociationTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenNetworkSecurityInterceptEndpointGroupAssociationEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNetworkSecurityInterceptEndpointGroupAssociationInterceptEndpointGroup(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityInterceptEndpointGroupAssociationNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityInterceptEndpointGroupAssociationEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
