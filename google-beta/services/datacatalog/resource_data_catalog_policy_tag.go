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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/datacatalog/PolicyTag.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package datacatalog

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
)

func ResourceDataCatalogPolicyTag() *schema.Resource {
	return &schema.Resource{
		Create: resourceDataCatalogPolicyTagCreate,
		Read:   resourceDataCatalogPolicyTagRead,
		Update: resourceDataCatalogPolicyTagUpdate,
		Delete: resourceDataCatalogPolicyTagDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDataCatalogPolicyTagImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				Description: `User defined name of this policy tag. It must: be unique within the parent
taxonomy; contain only unicode letters, numbers, underscores, dashes and spaces;
not start or end with spaces; and be at most 200 bytes long when encoded in UTF-8.`,
			},
			"taxonomy": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Taxonomy the policy tag is associated with`,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `Description of this policy tag. It must: contain only unicode characters, tabs,
newlines, carriage returns and page breaks; and be at most 2000 bytes long when
encoded in UTF-8. If not set, defaults to an empty description.
If not set, defaults to an empty description.`,
			},
			"parent_policy_tag": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `Resource name of this policy tag's parent policy tag.
If empty, it means this policy tag is a top level policy tag.
If not set, defaults to an empty string.`,
			},
			"child_policy_tags": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `Resource names of child policy tags of this policy tag.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Resource name of this policy tag, whose format is:
"projects/{project}/locations/{region}/taxonomies/{taxonomy}/policyTags/{policytag}"`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceDataCatalogPolicyTagCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandDataCatalogPolicyTagDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandDataCatalogPolicyTagDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	parentPolicyTagProp, err := expandDataCatalogPolicyTagParentPolicyTag(d.Get("parent_policy_tag"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("parent_policy_tag"); !tpgresource.IsEmptyValue(reflect.ValueOf(parentPolicyTagProp)) && (ok || !reflect.DeepEqual(v, parentPolicyTagProp)) {
		obj["parentPolicyTag"] = parentPolicyTagProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DataCatalogBasePath}}{{taxonomy}}/policyTags")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new PolicyTag: %#v", obj)
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
		return fmt.Errorf("Error creating PolicyTag: %s", err)
	}
	// Set computed resource properties from create API response so that they're available on the subsequent Read
	// call.
	if err := d.Set("name", flattenDataCatalogPolicyTagName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating PolicyTag %q: %#v", d.Id(), res)

	return resourceDataCatalogPolicyTagRead(d, meta)
}

func resourceDataCatalogPolicyTagRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DataCatalogBasePath}}{{name}}")
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("DataCatalogPolicyTag %q", d.Id()))
	}

	if err := d.Set("name", flattenDataCatalogPolicyTagName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading PolicyTag: %s", err)
	}
	if err := d.Set("display_name", flattenDataCatalogPolicyTagDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading PolicyTag: %s", err)
	}
	if err := d.Set("description", flattenDataCatalogPolicyTagDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading PolicyTag: %s", err)
	}
	if err := d.Set("parent_policy_tag", flattenDataCatalogPolicyTagParentPolicyTag(res["parentPolicyTag"], d, config)); err != nil {
		return fmt.Errorf("Error reading PolicyTag: %s", err)
	}
	if err := d.Set("child_policy_tags", flattenDataCatalogPolicyTagChildPolicyTags(res["childPolicyTags"], d, config)); err != nil {
		return fmt.Errorf("Error reading PolicyTag: %s", err)
	}

	return nil
}

func resourceDataCatalogPolicyTagUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	displayNameProp, err := expandDataCatalogPolicyTagDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandDataCatalogPolicyTagDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	parentPolicyTagProp, err := expandDataCatalogPolicyTagParentPolicyTag(d.Get("parent_policy_tag"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("parent_policy_tag"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, parentPolicyTagProp)) {
		obj["parentPolicyTag"] = parentPolicyTagProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DataCatalogBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating PolicyTag %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("parent_policy_tag") {
		updateMask = append(updateMask, "parentPolicyTag")
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
			return fmt.Errorf("Error updating PolicyTag %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating PolicyTag %q: %#v", d.Id(), res)
		}

	}

	return resourceDataCatalogPolicyTagRead(d, meta)
}

func resourceDataCatalogPolicyTagDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{DataCatalogBasePath}}{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting PolicyTag %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "PolicyTag")
	}

	log.Printf("[DEBUG] Finished deleting PolicyTag %q: %#v", d.Id(), res)
	return nil
}

func resourceDataCatalogPolicyTagImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)

	if err := tpgresource.ParseImportId([]string{
		"(?P<taxonomy>projects/[^/]+/locations/[^/]+/taxonomies/[^/]+)/policyTags/(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}

	originalName := d.Get("name").(string)
	originalTaxonomy := d.Get("taxonomy").(string)
	name := fmt.Sprintf("%s/policyTags/%s", originalTaxonomy, originalName)

	if err := d.Set("name", name); err != nil {
		return nil, fmt.Errorf("Error setting name: %s", err)
	}
	d.SetId(name)
	return []*schema.ResourceData{d}, nil
}

func flattenDataCatalogPolicyTagName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataCatalogPolicyTagDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataCatalogPolicyTagDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataCatalogPolicyTagParentPolicyTag(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataCatalogPolicyTagChildPolicyTags(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandDataCatalogPolicyTagDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogPolicyTagDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogPolicyTagParentPolicyTag(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
