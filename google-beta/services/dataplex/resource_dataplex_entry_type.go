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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/dataplex/EntryType.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package dataplex

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

func ResourceDataplexEntryType() *schema.Resource {
	return &schema.Resource{
		Create: resourceDataplexEntryTypeCreate,
		Read:   resourceDataplexEntryTypeRead,
		Update: resourceDataplexEntryTypeUpdate,
		Delete: resourceDataplexEntryTypeDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDataplexEntryTypeImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Description of the EntryType.`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `User friendly display name.`,
			},
			"entry_type_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The entry type id of the entry type.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `User-defined labels for the EntryType.


**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"location": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The location where entry type will be created in.`,
			},
			"platform": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The platform that Entries of this type belongs to.`,
			},
			"required_aspects": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `AspectInfo for the entry type.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Required aspect type for the entry type.`,
						},
					},
				},
			},
			"system": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The system that Entries of this type belongs to.`,
			},
			"type_aliases": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Indicates the class this Entry Type belongs to, for example, TABLE, DATABASE, MODEL.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time when the EntryType was created.`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The relative resource name of the EntryType, of the form: projects/{project_number}/locations/{location_id}/entryTypes/{entry_type_id}`,
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `System generated globally unique ID for the EntryType. This ID will be different if the EntryType is deleted and re-created with the same name.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time when the EntryType was last updated.`,
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

func resourceDataplexEntryTypeCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandDataplexEntryTypeDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	displayNameProp, err := expandDataplexEntryTypeDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	typeAliasesProp, err := expandDataplexEntryTypeTypeAliases(d.Get("type_aliases"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type_aliases"); !tpgresource.IsEmptyValue(reflect.ValueOf(typeAliasesProp)) && (ok || !reflect.DeepEqual(v, typeAliasesProp)) {
		obj["typeAliases"] = typeAliasesProp
	}
	platformProp, err := expandDataplexEntryTypePlatform(d.Get("platform"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("platform"); !tpgresource.IsEmptyValue(reflect.ValueOf(platformProp)) && (ok || !reflect.DeepEqual(v, platformProp)) {
		obj["platform"] = platformProp
	}
	systemProp, err := expandDataplexEntryTypeSystem(d.Get("system"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("system"); !tpgresource.IsEmptyValue(reflect.ValueOf(systemProp)) && (ok || !reflect.DeepEqual(v, systemProp)) {
		obj["system"] = systemProp
	}
	requiredAspectsProp, err := expandDataplexEntryTypeRequiredAspects(d.Get("required_aspects"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("required_aspects"); !tpgresource.IsEmptyValue(reflect.ValueOf(requiredAspectsProp)) && (ok || !reflect.DeepEqual(v, requiredAspectsProp)) {
		obj["requiredAspects"] = requiredAspectsProp
	}
	labelsProp, err := expandDataplexEntryTypeEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DataplexBasePath}}projects/{{project}}/locations/{{location}}/entryTypes?entryTypeId={{entry_type_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new EntryType: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for EntryType: %s", err)
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
		return fmt.Errorf("Error creating EntryType: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/entryTypes/{{entry_type_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = DataplexOperationWaitTime(
		config, res, project, "Creating EntryType", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create EntryType: %s", err)
	}

	log.Printf("[DEBUG] Finished creating EntryType %q: %#v", d.Id(), res)

	return resourceDataplexEntryTypeRead(d, meta)
}

func resourceDataplexEntryTypeRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DataplexBasePath}}projects/{{project}}/locations/{{location}}/entryTypes/{{entry_type_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for EntryType: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("DataplexEntryType %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading EntryType: %s", err)
	}

	if err := d.Set("name", flattenDataplexEntryTypeName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading EntryType: %s", err)
	}
	if err := d.Set("uid", flattenDataplexEntryTypeUid(res["uid"], d, config)); err != nil {
		return fmt.Errorf("Error reading EntryType: %s", err)
	}
	if err := d.Set("create_time", flattenDataplexEntryTypeCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading EntryType: %s", err)
	}
	if err := d.Set("update_time", flattenDataplexEntryTypeUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading EntryType: %s", err)
	}
	if err := d.Set("description", flattenDataplexEntryTypeDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading EntryType: %s", err)
	}
	if err := d.Set("display_name", flattenDataplexEntryTypeDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading EntryType: %s", err)
	}
	if err := d.Set("labels", flattenDataplexEntryTypeLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading EntryType: %s", err)
	}
	if err := d.Set("type_aliases", flattenDataplexEntryTypeTypeAliases(res["typeAliases"], d, config)); err != nil {
		return fmt.Errorf("Error reading EntryType: %s", err)
	}
	if err := d.Set("platform", flattenDataplexEntryTypePlatform(res["platform"], d, config)); err != nil {
		return fmt.Errorf("Error reading EntryType: %s", err)
	}
	if err := d.Set("system", flattenDataplexEntryTypeSystem(res["system"], d, config)); err != nil {
		return fmt.Errorf("Error reading EntryType: %s", err)
	}
	if err := d.Set("required_aspects", flattenDataplexEntryTypeRequiredAspects(res["requiredAspects"], d, config)); err != nil {
		return fmt.Errorf("Error reading EntryType: %s", err)
	}
	if err := d.Set("terraform_labels", flattenDataplexEntryTypeTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading EntryType: %s", err)
	}
	if err := d.Set("effective_labels", flattenDataplexEntryTypeEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading EntryType: %s", err)
	}

	return nil
}

func resourceDataplexEntryTypeUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for EntryType: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandDataplexEntryTypeDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	displayNameProp, err := expandDataplexEntryTypeDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	typeAliasesProp, err := expandDataplexEntryTypeTypeAliases(d.Get("type_aliases"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type_aliases"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, typeAliasesProp)) {
		obj["typeAliases"] = typeAliasesProp
	}
	platformProp, err := expandDataplexEntryTypePlatform(d.Get("platform"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("platform"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, platformProp)) {
		obj["platform"] = platformProp
	}
	systemProp, err := expandDataplexEntryTypeSystem(d.Get("system"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("system"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, systemProp)) {
		obj["system"] = systemProp
	}
	requiredAspectsProp, err := expandDataplexEntryTypeRequiredAspects(d.Get("required_aspects"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("required_aspects"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, requiredAspectsProp)) {
		obj["requiredAspects"] = requiredAspectsProp
	}
	labelsProp, err := expandDataplexEntryTypeEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DataplexBasePath}}projects/{{project}}/locations/{{location}}/entryTypes/{{entry_type_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating EntryType %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("type_aliases") {
		updateMask = append(updateMask, "typeAliases")
	}

	if d.HasChange("platform") {
		updateMask = append(updateMask, "platform")
	}

	if d.HasChange("system") {
		updateMask = append(updateMask, "system")
	}

	if d.HasChange("required_aspects") {
		updateMask = append(updateMask, "requiredAspects")
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
			return fmt.Errorf("Error updating EntryType %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating EntryType %q: %#v", d.Id(), res)
		}

		err = DataplexOperationWaitTime(
			config, res, project, "Updating EntryType", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceDataplexEntryTypeRead(d, meta)
}

func resourceDataplexEntryTypeDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for EntryType: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{DataplexBasePath}}projects/{{project}}/locations/{{location}}/entryTypes/{{entry_type_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting EntryType %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "EntryType")
	}

	err = DataplexOperationWaitTime(
		config, res, project, "Deleting EntryType", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting EntryType %q: %#v", d.Id(), res)
	return nil
}

func resourceDataplexEntryTypeImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/entryTypes/(?P<entry_type_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<entry_type_id>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<entry_type_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/entryTypes/{{entry_type_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenDataplexEntryTypeName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexEntryTypeUid(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexEntryTypeCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexEntryTypeUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexEntryTypeDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexEntryTypeDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexEntryTypeLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenDataplexEntryTypeTypeAliases(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexEntryTypePlatform(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexEntryTypeSystem(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexEntryTypeRequiredAspects(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"type": flattenDataplexEntryTypeRequiredAspectsType(original["type"], d, config),
		})
	}
	return transformed
}
func flattenDataplexEntryTypeRequiredAspectsType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexEntryTypeTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenDataplexEntryTypeEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandDataplexEntryTypeDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexEntryTypeDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexEntryTypeTypeAliases(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexEntryTypePlatform(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexEntryTypeSystem(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexEntryTypeRequiredAspects(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedType, err := expandDataplexEntryTypeRequiredAspectsType(original["type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["type"] = transformedType
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDataplexEntryTypeRequiredAspectsType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexEntryTypeEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
