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

package gemini

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

func ResourceGeminiReleaseChannelSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceGeminiReleaseChannelSettingCreate,
		Read:   resourceGeminiReleaseChannelSettingRead,
		Update: resourceGeminiReleaseChannelSettingUpdate,
		Delete: resourceGeminiReleaseChannelSettingDelete,

		Importer: &schema.ResourceImporter{
			State: resourceGeminiReleaseChannelSettingImport,
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
				Description: `Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.`,
			},
			"release_channel_setting_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Required. Id of the requesting object.
If auto-generating Id server-side, remove this field and
release_channel_setting_id from the method_signature of Create RPC`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Optional. Labels as key value pairs.

**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"release_channel": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `Optional. Release channel to be used.
Possible values:
STABLE
EXPERIMENTAL`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. [Output only] Create time stamp.`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Identifier. Name of the resource.
Format:projects/{project}/locations/{location}/releaseChannelSettings/{releaseChannelSetting}`,
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. [Output only] Update time stamp.`,
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

func resourceGeminiReleaseChannelSettingCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	releaseChannelProp, err := expandGeminiReleaseChannelSettingReleaseChannel(d.Get("release_channel"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("release_channel"); !tpgresource.IsEmptyValue(reflect.ValueOf(releaseChannelProp)) && (ok || !reflect.DeepEqual(v, releaseChannelProp)) {
		obj["releaseChannel"] = releaseChannelProp
	}
	labelsProp, err := expandGeminiReleaseChannelSettingEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	lockName, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/releaseChannelSettings/{{release_channel_setting_id}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{GeminiBasePath}}projects/{{project}}/locations/{{location}}/releaseChannelSettings?releaseChannelSettingId={{release_channel_setting_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ReleaseChannelSetting: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ReleaseChannelSetting: %s", err)
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
		return fmt.Errorf("Error creating ReleaseChannelSetting: %s", err)
	}
	if err := d.Set("name", flattenGeminiReleaseChannelSettingName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/releaseChannelSettings/{{release_channel_setting_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating ReleaseChannelSetting %q: %#v", d.Id(), res)

	return resourceGeminiReleaseChannelSettingRead(d, meta)
}

func resourceGeminiReleaseChannelSettingRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{GeminiBasePath}}projects/{{project}}/locations/{{location}}/releaseChannelSettings/{{release_channel_setting_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ReleaseChannelSetting: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("GeminiReleaseChannelSetting %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ReleaseChannelSetting: %s", err)
	}

	if err := d.Set("create_time", flattenGeminiReleaseChannelSettingCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading ReleaseChannelSetting: %s", err)
	}
	if err := d.Set("update_time", flattenGeminiReleaseChannelSettingUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading ReleaseChannelSetting: %s", err)
	}
	if err := d.Set("labels", flattenGeminiReleaseChannelSettingLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading ReleaseChannelSetting: %s", err)
	}
	if err := d.Set("release_channel", flattenGeminiReleaseChannelSettingReleaseChannel(res["releaseChannel"], d, config)); err != nil {
		return fmt.Errorf("Error reading ReleaseChannelSetting: %s", err)
	}
	if err := d.Set("name", flattenGeminiReleaseChannelSettingName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading ReleaseChannelSetting: %s", err)
	}
	if err := d.Set("terraform_labels", flattenGeminiReleaseChannelSettingTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading ReleaseChannelSetting: %s", err)
	}
	if err := d.Set("effective_labels", flattenGeminiReleaseChannelSettingEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading ReleaseChannelSetting: %s", err)
	}

	return nil
}

func resourceGeminiReleaseChannelSettingUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ReleaseChannelSetting: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	releaseChannelProp, err := expandGeminiReleaseChannelSettingReleaseChannel(d.Get("release_channel"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("release_channel"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, releaseChannelProp)) {
		obj["releaseChannel"] = releaseChannelProp
	}
	labelsProp, err := expandGeminiReleaseChannelSettingEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	lockName, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/releaseChannelSettings/{{release_channel_setting_id}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{GeminiBasePath}}projects/{{project}}/locations/{{location}}/releaseChannelSettings/{{release_channel_setting_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ReleaseChannelSetting %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("release_channel") {
		updateMask = append(updateMask, "releaseChannel")
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
			return fmt.Errorf("Error updating ReleaseChannelSetting %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating ReleaseChannelSetting %q: %#v", d.Id(), res)
		}

	}

	return resourceGeminiReleaseChannelSettingRead(d, meta)
}

func resourceGeminiReleaseChannelSettingDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ReleaseChannelSetting: %s", err)
	}
	billingProject = project

	lockName, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/releaseChannelSettings/{{release_channel_setting_id}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{GeminiBasePath}}projects/{{project}}/locations/{{location}}/releaseChannelSettings/{{release_channel_setting_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting ReleaseChannelSetting %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "ReleaseChannelSetting")
	}

	log.Printf("[DEBUG] Finished deleting ReleaseChannelSetting %q: %#v", d.Id(), res)
	return nil
}

func resourceGeminiReleaseChannelSettingImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/releaseChannelSettings/(?P<release_channel_setting_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<release_channel_setting_id>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<release_channel_setting_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/releaseChannelSettings/{{release_channel_setting_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenGeminiReleaseChannelSettingCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGeminiReleaseChannelSettingUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGeminiReleaseChannelSettingLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenGeminiReleaseChannelSettingReleaseChannel(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGeminiReleaseChannelSettingName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGeminiReleaseChannelSettingTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenGeminiReleaseChannelSettingEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandGeminiReleaseChannelSettingReleaseChannel(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGeminiReleaseChannelSettingEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
