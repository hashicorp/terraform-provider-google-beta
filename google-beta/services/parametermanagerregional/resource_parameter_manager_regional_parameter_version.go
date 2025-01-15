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

package parametermanagerregional

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"

	"google.golang.org/api/googleapi"
)

func ResourceParameterManagerRegionalRegionalParameterVersion() *schema.Resource {
	return &schema.Resource{
		Create: resourceParameterManagerRegionalRegionalParameterVersionCreate,
		Read:   resourceParameterManagerRegionalRegionalParameterVersionRead,
		Update: resourceParameterManagerRegionalRegionalParameterVersionUpdate,
		Delete: resourceParameterManagerRegionalRegionalParameterVersionDelete,

		Importer: &schema.ResourceImporter{
			State: resourceParameterManagerRegionalRegionalParameterVersionImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"parameter": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `Parameter Manager Regional Parameter resource.`,
			},
			"parameter_version_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Version ID of the Regional Parameter Version Resource. This must be unique within the Regional Parameter.`,
			},
			"parameter_data": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The Regional Parameter data.`,
				Sensitive:   true,
			},

			"disabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `The current state of Regional Parameter Version. This field is only applicable for updating Regional Parameter Version.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time at which the Regional Parameter Version was created.`,
			},
			"location": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Location of Parameter Manager Regional parameter resource.`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The resource name of the Regional Parameter Version. Format:
'projects/{{project}}/locations/{{location}}/parameters/{{parameter_id}}/versions/{{parameter_version_id}}'`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time at which the Regional Parameter Version was updated.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceParameterManagerRegionalRegionalParameterVersionCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	disabledProp, err := expandParameterManagerRegionalRegionalParameterVersionDisabled(d.Get("disabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(disabledProp)) && (ok || !reflect.DeepEqual(v, disabledProp)) {
		obj["disabled"] = disabledProp
	}
	payloadProp, err := expandParameterManagerRegionalRegionalParameterVersionPayload(nil, d, config)
	if err != nil {
		return err
	} else if !tpgresource.IsEmptyValue(reflect.ValueOf(payloadProp)) {
		obj["payload"] = payloadProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ParameterManagerRegionalBasePath}}{{parameter}}/versions?parameter_version_id={{parameter_version_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new RegionalParameterVersion: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	parameter := d.Get("parameter").(string)
	parameterRegex := regexp.MustCompile("projects/(.+)/locations/(.+)/parameters/(.+)$")

	parts := parameterRegex.FindStringSubmatch(parameter)
	if len(parts) != 4 {
		return fmt.Errorf("parameter does not fit the format `projects/{{project}}/locations/{{location}}/parameters/{{parameter}}`")
	}

	if err := d.Set("location", parts[2]); err != nil {
		return fmt.Errorf("Error setting location: %s", err)
	}

	// Override the url after setting the location
	url, err = tpgresource.ReplaceVars(d, config, "{{ParameterManagerRegionalBasePath}}{{parameter}}/versions?parameter_version_id={{parameter_version_id}}")
	if err != nil {
		return err
	}
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
		return fmt.Errorf("Error creating RegionalParameterVersion: %s", err)
	}
	if err := d.Set("name", flattenParameterManagerRegionalRegionalParameterVersionName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{parameter}}/versions/{{parameter_version_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating RegionalParameterVersion %q: %#v", d.Id(), res)

	return resourceParameterManagerRegionalRegionalParameterVersionRead(d, meta)
}

func resourceParameterManagerRegionalRegionalParameterVersionRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ParameterManagerRegionalBasePath}}{{parameter}}/versions/{{parameter_version_id}}")
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ParameterManagerRegionalRegionalParameterVersion %q", d.Id()))
	}

	if err := d.Set("name", flattenParameterManagerRegionalRegionalParameterVersionName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionalParameterVersion: %s", err)
	}
	if err := d.Set("create_time", flattenParameterManagerRegionalRegionalParameterVersionCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionalParameterVersion: %s", err)
	}
	if err := d.Set("update_time", flattenParameterManagerRegionalRegionalParameterVersionUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionalParameterVersion: %s", err)
	}
	if err := d.Set("disabled", flattenParameterManagerRegionalRegionalParameterVersionDisabled(res["disabled"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionalParameterVersion: %s", err)
	}
	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if flattenedProp := flattenParameterManagerRegionalRegionalParameterVersionPayload(res["payload"], d, config); flattenedProp != nil {
		if gerr, ok := flattenedProp.(*googleapi.Error); ok {
			return fmt.Errorf("Error reading RegionalParameterVersion: %s", gerr)
		}
		casted := flattenedProp.([]interface{})[0]
		if casted != nil {
			for k, v := range casted.(map[string]interface{}) {
				if err := d.Set(k, v); err != nil {
					return fmt.Errorf("Error setting %s: %s", k, err)
				}
			}
		}
	}

	return nil
}

func resourceParameterManagerRegionalRegionalParameterVersionUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	disabledProp, err := expandParameterManagerRegionalRegionalParameterVersionDisabled(d.Get("disabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, disabledProp)) {
		obj["disabled"] = disabledProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ParameterManagerRegionalBasePath}}{{parameter}}/versions/{{parameter_version_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating RegionalParameterVersion %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("disabled") {
		updateMask = append(updateMask, "disabled")
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
			return fmt.Errorf("Error updating RegionalParameterVersion %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating RegionalParameterVersion %q: %#v", d.Id(), res)
		}

	}

	return resourceParameterManagerRegionalRegionalParameterVersionRead(d, meta)
}

func resourceParameterManagerRegionalRegionalParameterVersionDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{ParameterManagerRegionalBasePath}}{{parameter}}/versions/{{parameter_version_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting RegionalParameterVersion %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "RegionalParameterVersion")
	}

	log.Printf("[DEBUG] Finished deleting RegionalParameterVersion %q: %#v", d.Id(), res)
	return nil
}

func resourceParameterManagerRegionalRegionalParameterVersionImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := tpgresource.ParseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}

	name := d.Get("name").(string)
	parameterRegex := regexp.MustCompile("(projects/.+/locations/.+/parameters/.+)/versions/.+$")
	versionRegex := regexp.MustCompile("projects/(.+)/locations/(.+)/parameters/(.+)/versions/(.+)$")

	parts := parameterRegex.FindStringSubmatch(name)
	if len(parts) != 2 {
		return nil, fmt.Errorf("Version name does not fit the format `projects/{{project}}/locations/{{location}}/parameters/{{parameter_id}}/versions/{{parameter_version_id}}`")
	}
	if err := d.Set("parameter", parts[1]); err != nil {
		return nil, fmt.Errorf("Error setting parameter: %s", err)
	}

	parts = versionRegex.FindStringSubmatch(name)

	if err := d.Set("parameter_version_id", parts[4]); err != nil {
		return nil, fmt.Errorf("Error setting parameter_version_id: %s", err)
	}

	if err := d.Set("location", parts[2]); err != nil {
		return nil, fmt.Errorf("Error setting location: %s", err)
	}

	return []*schema.ResourceData{d}, nil
}

func flattenParameterManagerRegionalRegionalParameterVersionName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenParameterManagerRegionalRegionalParameterVersionCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenParameterManagerRegionalRegionalParameterVersionUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenParameterManagerRegionalRegionalParameterVersionDisabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenParameterManagerRegionalRegionalParameterVersionPayload(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	data, err := base64.StdEncoding.DecodeString(original["data"].(string))
	if err != nil {
		return err
	}
	transformed["parameter_data"] = string(data)
	return []interface{}{transformed}
}

func expandParameterManagerRegionalRegionalParameterVersionDisabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandParameterManagerRegionalRegionalParameterVersionPayload(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	transformed := make(map[string]interface{})
	transformedParameterData, err := expandParameterManagerRegionalRegionalParameterVersionPayloadParameterData(d.Get("parameter_data"), d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedParameterData); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["data"] = transformedParameterData
	}

	return transformed, nil
}

func expandParameterManagerRegionalRegionalParameterVersionPayloadParameterData(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	if v == nil {
		return nil, nil
	}

	return base64.StdEncoding.EncodeToString([]byte(v.(string))), nil
}
