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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/cloudquotas/QuotaPreference.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package cloudquotas

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceCloudQuotasQuotaPreference() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudQuotasQuotaPreferenceCreate,
		Read:   resourceCloudQuotasQuotaPreferenceRead,
		Update: resourceCloudQuotasQuotaPreferenceUpdate,
		Delete: resourceCloudQuotasQuotaPreferenceDelete,

		Importer: &schema.ResourceImporter{
			State: resourceCloudQuotasQuotaPreferenceImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"parent": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `The parent of the quota preference. Allowed parents are "projects/[project-id / number]" or "folders/[folder-id / number]" or "organizations/[org-id / number]".`,
			},
			"quota_config": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `The preferred quota configuration.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"preferred_value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The preferred value. Must be greater than or equal to -1. If set to -1, it means the value is "unlimited".`,
						},
						"annotations": {
							Type:     schema.TypeMap,
							Optional: true,
							Description: `The annotations map for clients to store small amounts of arbitrary data. Do not put PII or other sensitive information here. See https://google.aip.dev/128#annotations.

An object containing a list of "key: value" pairs. Example: '{ "name": "wrench", "mass": "1.3kg", "count": "3" }'.`,
							Elem: &schema.Schema{Type: schema.TypeString},
						},
						"granted_value": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Granted quota value.`,
						},
						"request_origin": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The origin of the quota preference request.`,
						},
						"state_detail": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Optional details about the state of this quota preference.`,
						},
						"trace_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The trace id that the Google Cloud uses to provision the requested quota. This trace id may be used by the client to contact Cloud support to track the state of a quota preference request. The trace id is only produced for increase requests and is unique for each request. The quota decrease requests do not have a trace id.`,
						},
					},
				},
			},
			"quota_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				Description: `The id of the quota to which the quota preference is applied. A quota id is unique in the service.
Example: 'CPUS-per-project-region'.`,
			},
			"service": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: `The name of the service to which the quota preference is applied.`,
			},
			"contact_email": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `An email address that can be used for quota related communication between the Google Cloud and the user in case the Google Cloud needs further information to make a decision on whether the user preferred quota can be granted.

The Google account for the email address must have quota update permission for the project, folder or organization this quota preference is for.`,
			},
			"dimensions": {
				Type:     schema.TypeMap,
				Computed: true,
				Optional: true,
				Description: `The dimensions that this quota preference applies to. The key of the map entry is the name of a dimension, such as "region", "zone", "network_id", and the value of the map entry is the dimension value. If a dimension is missing from the map of dimensions, the quota preference applies to all the dimension values except for those that have other quota preferences configured for the specific value.

NOTE: QuotaPreferences can only be applied across all values of "user" and "resource" dimension. Do not set values for "user" or "resource" in the dimension map.

Example: '{"provider": "Foo Inc"}' where "provider" is a service specific dimension.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"ignore_safety_checks": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: verify.ValidateEnum([]string{"QUOTA_SAFETY_CHECK_UNSPECIFIED", "QUOTA_DECREASE_BELOW_USAGE", "QUOTA_DECREASE_PERCENTAGE_TOO_HIGH", ""}),
				Description:  `The list of quota safety checks to be ignored. Default value: "QUOTA_SAFETY_CHECK_UNSPECIFIED" Possible values: ["QUOTA_SAFETY_CHECK_UNSPECIFIED", "QUOTA_DECREASE_BELOW_USAGE", "QUOTA_DECREASE_PERCENTAGE_TOO_HIGH"]`,
				Default:      "QUOTA_SAFETY_CHECK_UNSPECIFIED",
			},
			"justification": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The reason / justification for this quota preference.`,
			},
			"name": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The resource name of the quota preference. Required except in the CREATE requests.`,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Create time stamp.
A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: '2014-10-02T15:01:23Z' and '2014-10-02T15:01:23.045123456Z'.`,
			},
			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The current etag of the quota preference. If an etag is provided on update and does not match the current server's etag of the quota preference, the request will be blocked and an ABORTED error will be returned. See https://google.aip.dev/134#etags for more details on etags.`,
			},
			"reconciling": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: `Is the quota preference pending Google Cloud approval and fulfillment.`,
			},
			"update_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Update time stamp.
A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: '2014-10-02T15:01:23Z' and '2014-10-02T15:01:23.045123456Z'.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceCloudQuotasQuotaPreferenceCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandCloudQuotasQuotaPreferenceName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	serviceProp, err := expandCloudQuotasQuotaPreferenceService(d.Get("service"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("service"); !tpgresource.IsEmptyValue(reflect.ValueOf(serviceProp)) && (ok || !reflect.DeepEqual(v, serviceProp)) {
		obj["service"] = serviceProp
	}
	quotaIdProp, err := expandCloudQuotasQuotaPreferenceQuotaId(d.Get("quota_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("quota_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(quotaIdProp)) && (ok || !reflect.DeepEqual(v, quotaIdProp)) {
		obj["quotaId"] = quotaIdProp
	}
	quotaConfigProp, err := expandCloudQuotasQuotaPreferenceQuotaConfig(d.Get("quota_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("quota_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(quotaConfigProp)) && (ok || !reflect.DeepEqual(v, quotaConfigProp)) {
		obj["quotaConfig"] = quotaConfigProp
	}
	dimensionsProp, err := expandCloudQuotasQuotaPreferenceDimensions(d.Get("dimensions"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("dimensions"); !tpgresource.IsEmptyValue(reflect.ValueOf(dimensionsProp)) && (ok || !reflect.DeepEqual(v, dimensionsProp)) {
		obj["dimensions"] = dimensionsProp
	}
	justificationProp, err := expandCloudQuotasQuotaPreferenceJustification(d.Get("justification"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("justification"); !tpgresource.IsEmptyValue(reflect.ValueOf(justificationProp)) && (ok || !reflect.DeepEqual(v, justificationProp)) {
		obj["justification"] = justificationProp
	}
	contactEmailProp, err := expandCloudQuotasQuotaPreferenceContactEmail(d.Get("contact_email"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("contact_email"); !tpgresource.IsEmptyValue(reflect.ValueOf(contactEmailProp)) && (ok || !reflect.DeepEqual(v, contactEmailProp)) {
		obj["contactEmail"] = contactEmailProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{CloudQuotasBasePath}}{{parent}}/locations/global/quotaPreferences?quotaPreferenceId={{name}}&ignoreSafetyChecks={{ignore_safety_checks}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new QuotaPreference: %#v", obj)
	billingProject := ""

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
		return fmt.Errorf("Error creating QuotaPreference: %s", err)
	}
	// Set computed resource properties from create API response so that they're available on the subsequent Read
	// call.
	err = resourceCloudQuotasQuotaPreferencePostCreateSetComputedFields(d, meta, res)
	if err != nil {
		return fmt.Errorf("setting computed ID format fields: %w", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{parent}}/locations/global/quotaPreferences/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating QuotaPreference %q: %#v", d.Id(), res)

	return resourceCloudQuotasQuotaPreferenceRead(d, meta)
}

func resourceCloudQuotasQuotaPreferenceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{CloudQuotasBasePath}}{{parent}}/locations/global/quotaPreferences/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("CloudQuotasQuotaPreference %q", d.Id()))
	}

	if err := d.Set("name", flattenCloudQuotasQuotaPreferenceName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading QuotaPreference: %s", err)
	}
	if err := d.Set("service", flattenCloudQuotasQuotaPreferenceService(res["service"], d, config)); err != nil {
		return fmt.Errorf("Error reading QuotaPreference: %s", err)
	}
	if err := d.Set("quota_id", flattenCloudQuotasQuotaPreferenceQuotaId(res["quotaId"], d, config)); err != nil {
		return fmt.Errorf("Error reading QuotaPreference: %s", err)
	}
	if err := d.Set("quota_config", flattenCloudQuotasQuotaPreferenceQuotaConfig(res["quotaConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading QuotaPreference: %s", err)
	}
	if err := d.Set("dimensions", flattenCloudQuotasQuotaPreferenceDimensions(res["dimensions"], d, config)); err != nil {
		return fmt.Errorf("Error reading QuotaPreference: %s", err)
	}
	if err := d.Set("etag", flattenCloudQuotasQuotaPreferenceEtag(res["etag"], d, config)); err != nil {
		return fmt.Errorf("Error reading QuotaPreference: %s", err)
	}
	if err := d.Set("create_time", flattenCloudQuotasQuotaPreferenceCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading QuotaPreference: %s", err)
	}
	if err := d.Set("update_time", flattenCloudQuotasQuotaPreferenceUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading QuotaPreference: %s", err)
	}
	if err := d.Set("reconciling", flattenCloudQuotasQuotaPreferenceReconciling(res["reconciling"], d, config)); err != nil {
		return fmt.Errorf("Error reading QuotaPreference: %s", err)
	}

	return nil
}

func resourceCloudQuotasQuotaPreferenceUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	nameProp, err := expandCloudQuotasQuotaPreferenceName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	serviceProp, err := expandCloudQuotasQuotaPreferenceService(d.Get("service"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("service"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, serviceProp)) {
		obj["service"] = serviceProp
	}
	quotaIdProp, err := expandCloudQuotasQuotaPreferenceQuotaId(d.Get("quota_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("quota_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, quotaIdProp)) {
		obj["quotaId"] = quotaIdProp
	}
	quotaConfigProp, err := expandCloudQuotasQuotaPreferenceQuotaConfig(d.Get("quota_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("quota_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, quotaConfigProp)) {
		obj["quotaConfig"] = quotaConfigProp
	}
	dimensionsProp, err := expandCloudQuotasQuotaPreferenceDimensions(d.Get("dimensions"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("dimensions"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, dimensionsProp)) {
		obj["dimensions"] = dimensionsProp
	}
	justificationProp, err := expandCloudQuotasQuotaPreferenceJustification(d.Get("justification"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("justification"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, justificationProp)) {
		obj["justification"] = justificationProp
	}
	contactEmailProp, err := expandCloudQuotasQuotaPreferenceContactEmail(d.Get("contact_email"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("contact_email"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, contactEmailProp)) {
		obj["contactEmail"] = contactEmailProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{CloudQuotasBasePath}}{{parent}}/locations/global/quotaPreferences/{{name}}?ignoreSafetyChecks={{ignore_safety_checks}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating QuotaPreference %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("name") {
		updateMask = append(updateMask, "name")
	}

	if d.HasChange("service") {
		updateMask = append(updateMask, "service")
	}

	if d.HasChange("quota_id") {
		updateMask = append(updateMask, "quotaId")
	}

	if d.HasChange("quota_config") {
		updateMask = append(updateMask, "quotaConfig")
	}

	if d.HasChange("dimensions") {
		updateMask = append(updateMask, "dimensions")
	}

	if d.HasChange("justification") {
		updateMask = append(updateMask, "justification")
	}

	if d.HasChange("contact_email") {
		updateMask = append(updateMask, "contactEmail")
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
			return fmt.Errorf("Error updating QuotaPreference %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating QuotaPreference %q: %#v", d.Id(), res)
		}

	}

	return resourceCloudQuotasQuotaPreferenceRead(d, meta)
}

func resourceCloudQuotasQuotaPreferenceDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARNING] CloudQuotas QuotaPreference resources"+
		" cannot be deleted from Google Cloud. The resource %s will be removed from Terraform"+
		" state, but will still be present on Google Cloud.", d.Id())
	d.SetId("")

	return nil
}

func resourceCloudQuotasQuotaPreferenceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^(?P<parent>.+)/locations/global/quotaPreferences/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "{{parent}}/locations/global/quotaPreferences/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenCloudQuotasQuotaPreferenceName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func flattenCloudQuotasQuotaPreferenceService(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudQuotasQuotaPreferenceQuotaId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudQuotasQuotaPreferenceQuotaConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["preferred_value"] =
		flattenCloudQuotasQuotaPreferenceQuotaConfigPreferredValue(original["preferredValue"], d, config)
	transformed["state_detail"] =
		flattenCloudQuotasQuotaPreferenceQuotaConfigStateDetail(original["stateDetail"], d, config)
	transformed["granted_value"] =
		flattenCloudQuotasQuotaPreferenceQuotaConfigGrantedValue(original["grantedValue"], d, config)
	transformed["trace_id"] =
		flattenCloudQuotasQuotaPreferenceQuotaConfigTraceId(original["traceId"], d, config)
	transformed["annotations"] =
		flattenCloudQuotasQuotaPreferenceQuotaConfigAnnotations(original["annotations"], d, config)
	transformed["request_origin"] =
		flattenCloudQuotasQuotaPreferenceQuotaConfigRequestOrigin(original["requestOrigin"], d, config)
	return []interface{}{transformed}
}
func flattenCloudQuotasQuotaPreferenceQuotaConfigPreferredValue(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudQuotasQuotaPreferenceQuotaConfigStateDetail(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudQuotasQuotaPreferenceQuotaConfigGrantedValue(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudQuotasQuotaPreferenceQuotaConfigTraceId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudQuotasQuotaPreferenceQuotaConfigAnnotations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// ignore read on this field
	return d.Get("quota_config.0.annotations")
}

func flattenCloudQuotasQuotaPreferenceQuotaConfigRequestOrigin(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudQuotasQuotaPreferenceDimensions(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudQuotasQuotaPreferenceEtag(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudQuotasQuotaPreferenceCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudQuotasQuotaPreferenceUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudQuotasQuotaPreferenceReconciling(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandCloudQuotasQuotaPreferenceName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return tpgresource.GetResourceNameFromSelfLink(v.(string)), nil
}

func expandCloudQuotasQuotaPreferenceService(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudQuotasQuotaPreferenceQuotaId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudQuotasQuotaPreferenceQuotaConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPreferredValue, err := expandCloudQuotasQuotaPreferenceQuotaConfigPreferredValue(original["preferred_value"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPreferredValue); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["preferredValue"] = transformedPreferredValue
	}

	transformedStateDetail, err := expandCloudQuotasQuotaPreferenceQuotaConfigStateDetail(original["state_detail"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStateDetail); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["stateDetail"] = transformedStateDetail
	}

	transformedGrantedValue, err := expandCloudQuotasQuotaPreferenceQuotaConfigGrantedValue(original["granted_value"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGrantedValue); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["grantedValue"] = transformedGrantedValue
	}

	transformedTraceId, err := expandCloudQuotasQuotaPreferenceQuotaConfigTraceId(original["trace_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTraceId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["traceId"] = transformedTraceId
	}

	transformedAnnotations, err := expandCloudQuotasQuotaPreferenceQuotaConfigAnnotations(original["annotations"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAnnotations); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["annotations"] = transformedAnnotations
	}

	transformedRequestOrigin, err := expandCloudQuotasQuotaPreferenceQuotaConfigRequestOrigin(original["request_origin"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequestOrigin); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["requestOrigin"] = transformedRequestOrigin
	}

	return transformed, nil
}

func expandCloudQuotasQuotaPreferenceQuotaConfigPreferredValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudQuotasQuotaPreferenceQuotaConfigStateDetail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudQuotasQuotaPreferenceQuotaConfigGrantedValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudQuotasQuotaPreferenceQuotaConfigTraceId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return nil, nil
}

func expandCloudQuotasQuotaPreferenceQuotaConfigAnnotations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCloudQuotasQuotaPreferenceQuotaConfigRequestOrigin(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudQuotasQuotaPreferenceDimensions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCloudQuotasQuotaPreferenceJustification(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudQuotasQuotaPreferenceContactEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceCloudQuotasQuotaPreferencePostCreateSetComputedFields(d *schema.ResourceData, meta interface{}, res map[string]interface{}) error {
	config := meta.(*transport_tpg.Config)
	// name is set by API when unset
	if tpgresource.IsEmptyValue(reflect.ValueOf(d.Get("name"))) {
		if err := d.Set("name", flattenCloudQuotasQuotaPreferenceName(res["name"], d, config)); err != nil {
			return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
		}
	}
	return nil
}
