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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/apphub/Application.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package apphub

import (
	"context"
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
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func apphubApplicationCustomizeDiff(_ context.Context, diff *schema.ResourceDiff, meta interface{}) error {
	if diff.HasChange("location") || diff.HasChange("scope.0.type") {
		location := diff.Get("location")
		scope_type := diff.Get("scope.0.type")

		if scope_type == "GLOBAL" {
			if location != "global" {
				return fmt.Errorf("Error validating location %s with %s scope type", location, scope_type)
			}
		} else {
			if location == "global" {
				return fmt.Errorf("Error validating location %s with %s scope type", location, scope_type)
			}
		}
	}
	return nil
}

func ResourceApphubApplication() *schema.Resource {
	return &schema.Resource{
		Create: resourceApphubApplicationCreate,
		Read:   resourceApphubApplicationRead,
		Update: resourceApphubApplicationUpdate,
		Delete: resourceApphubApplicationDelete,

		Importer: &schema.ResourceImporter{
			State: resourceApphubApplicationImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			apphubApplicationCustomizeDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"application_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Required. The Application identifier.`,
			},
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Part of 'parent'. See documentation of 'projectsId'.`,
			},
			"scope": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `Scope of an application.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: verify.ValidateEnum([]string{"REGIONAL", "GLOBAL"}),
							Description: `Required. Scope Type. 
 Possible values:
REGIONAL
GLOBAL Possible values: ["REGIONAL", "GLOBAL"]`,
						},
					},
				},
			},
			"attributes": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Consumer provided attributes.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"business_owners": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Optional. Business team that ensures user needs are met and value is delivered`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"email": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `Required. Email address of the contacts.`,
									},
									"display_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `Optional. Contact's name.`,
									},
								},
							},
						},
						"criticality": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Criticality of the Application, Service, or Workload`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: verify.ValidateEnum([]string{"MISSION_CRITICAL", "HIGH", "MEDIUM", "LOW"}),
										Description:  `Criticality type. Possible values: ["MISSION_CRITICAL", "HIGH", "MEDIUM", "LOW"]`,
									},
								},
							},
						},
						"developer_owners": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Optional. Developer team that owns development and coding.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"email": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `Required. Email address of the contacts.`,
									},
									"display_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `Optional. Contact's name.`,
									},
								},
							},
						},
						"environment": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Environment of the Application, Service, or Workload`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: verify.ValidateEnum([]string{"PRODUCTION", "STAGING", "TEST", "DEVELOPMENT"}),
										Description:  `Environment type. Possible values: ["PRODUCTION", "STAGING", "TEST", "DEVELOPMENT"]`,
									},
								},
							},
						},
						"operator_owners": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Optional. Operator team that ensures runtime and operations.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"email": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `Required. Email address of the contacts.`,
									},
									"display_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `Optional. Contact's name.`,
									},
								},
							},
						},
					},
				},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Optional. User-defined description of an Application.`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Optional. User-defined name for the Application.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. Create time.`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Identifier. The resource name of an Application. Format:
"projects/{host-project-id}/locations/{location}/applications/{application-id}"`,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Output only. Application state. 
 Possible values:
 STATE_UNSPECIFIED
CREATING
ACTIVE
DELETING`,
			},
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. A universally unique identifier (in UUID4 format) for the 'Application'.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. Update time.`,
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

func resourceApphubApplicationCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandApphubApplicationDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandApphubApplicationDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	attributesProp, err := expandApphubApplicationAttributes(d.Get("attributes"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("attributes"); !tpgresource.IsEmptyValue(reflect.ValueOf(attributesProp)) && (ok || !reflect.DeepEqual(v, attributesProp)) {
		obj["attributes"] = attributesProp
	}
	scopeProp, err := expandApphubApplicationScope(d.Get("scope"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("scope"); !tpgresource.IsEmptyValue(reflect.ValueOf(scopeProp)) && (ok || !reflect.DeepEqual(v, scopeProp)) {
		obj["scope"] = scopeProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ApphubBasePath}}projects/{{project}}/locations/{{location}}/applications?applicationId={{application_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Application: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Application: %s", err)
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
		return fmt.Errorf("Error creating Application: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/applications/{{application_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = ApphubOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating Application", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create Application: %s", err)
	}

	if err := d.Set("name", flattenApphubApplicationName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/applications/{{application_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Application %q: %#v", d.Id(), res)

	return resourceApphubApplicationRead(d, meta)
}

func resourceApphubApplicationRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ApphubBasePath}}projects/{{project}}/locations/{{location}}/applications/{{application_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Application: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ApphubApplication %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Application: %s", err)
	}

	if err := d.Set("name", flattenApphubApplicationName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Application: %s", err)
	}
	if err := d.Set("display_name", flattenApphubApplicationDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Application: %s", err)
	}
	if err := d.Set("description", flattenApphubApplicationDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Application: %s", err)
	}
	if err := d.Set("attributes", flattenApphubApplicationAttributes(res["attributes"], d, config)); err != nil {
		return fmt.Errorf("Error reading Application: %s", err)
	}
	if err := d.Set("create_time", flattenApphubApplicationCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Application: %s", err)
	}
	if err := d.Set("update_time", flattenApphubApplicationUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Application: %s", err)
	}
	if err := d.Set("scope", flattenApphubApplicationScope(res["scope"], d, config)); err != nil {
		return fmt.Errorf("Error reading Application: %s", err)
	}
	if err := d.Set("uid", flattenApphubApplicationUid(res["uid"], d, config)); err != nil {
		return fmt.Errorf("Error reading Application: %s", err)
	}
	if err := d.Set("state", flattenApphubApplicationState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading Application: %s", err)
	}

	return nil
}

func resourceApphubApplicationUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Application: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandApphubApplicationDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandApphubApplicationDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	attributesProp, err := expandApphubApplicationAttributes(d.Get("attributes"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("attributes"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, attributesProp)) {
		obj["attributes"] = attributesProp
	}
	scopeProp, err := expandApphubApplicationScope(d.Get("scope"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("scope"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, scopeProp)) {
		obj["scope"] = scopeProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ApphubBasePath}}projects/{{project}}/locations/{{location}}/applications/{{application_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Application %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("attributes") {
		updateMask = append(updateMask, "attributes")
	}

	if d.HasChange("scope") {
		updateMask = append(updateMask, "scope")
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
			return fmt.Errorf("Error updating Application %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating Application %q: %#v", d.Id(), res)
		}

		err = ApphubOperationWaitTime(
			config, res, project, "Updating Application", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceApphubApplicationRead(d, meta)
}

func resourceApphubApplicationDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Application: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{ApphubBasePath}}projects/{{project}}/locations/{{location}}/applications/{{application_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting Application %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "Application")
	}

	err = ApphubOperationWaitTime(
		config, res, project, "Deleting Application", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Application %q: %#v", d.Id(), res)
	return nil
}

func resourceApphubApplicationImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/applications/(?P<application_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<application_id>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<application_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/applications/{{application_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenApphubApplicationName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApphubApplicationDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApphubApplicationDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApphubApplicationAttributes(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["criticality"] =
		flattenApphubApplicationAttributesCriticality(original["criticality"], d, config)
	transformed["environment"] =
		flattenApphubApplicationAttributesEnvironment(original["environment"], d, config)
	transformed["developer_owners"] =
		flattenApphubApplicationAttributesDeveloperOwners(original["developerOwners"], d, config)
	transformed["operator_owners"] =
		flattenApphubApplicationAttributesOperatorOwners(original["operatorOwners"], d, config)
	transformed["business_owners"] =
		flattenApphubApplicationAttributesBusinessOwners(original["businessOwners"], d, config)
	return []interface{}{transformed}
}
func flattenApphubApplicationAttributesCriticality(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["type"] =
		flattenApphubApplicationAttributesCriticalityType(original["type"], d, config)
	return []interface{}{transformed}
}
func flattenApphubApplicationAttributesCriticalityType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApphubApplicationAttributesEnvironment(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["type"] =
		flattenApphubApplicationAttributesEnvironmentType(original["type"], d, config)
	return []interface{}{transformed}
}
func flattenApphubApplicationAttributesEnvironmentType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApphubApplicationAttributesDeveloperOwners(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"display_name": flattenApphubApplicationAttributesDeveloperOwnersDisplayName(original["displayName"], d, config),
			"email":        flattenApphubApplicationAttributesDeveloperOwnersEmail(original["email"], d, config),
		})
	}
	return transformed
}
func flattenApphubApplicationAttributesDeveloperOwnersDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApphubApplicationAttributesDeveloperOwnersEmail(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApphubApplicationAttributesOperatorOwners(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"display_name": flattenApphubApplicationAttributesOperatorOwnersDisplayName(original["displayName"], d, config),
			"email":        flattenApphubApplicationAttributesOperatorOwnersEmail(original["email"], d, config),
		})
	}
	return transformed
}
func flattenApphubApplicationAttributesOperatorOwnersDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApphubApplicationAttributesOperatorOwnersEmail(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApphubApplicationAttributesBusinessOwners(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"display_name": flattenApphubApplicationAttributesBusinessOwnersDisplayName(original["displayName"], d, config),
			"email":        flattenApphubApplicationAttributesBusinessOwnersEmail(original["email"], d, config),
		})
	}
	return transformed
}
func flattenApphubApplicationAttributesBusinessOwnersDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApphubApplicationAttributesBusinessOwnersEmail(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApphubApplicationCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApphubApplicationUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApphubApplicationScope(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["type"] =
		flattenApphubApplicationScopeType(original["type"], d, config)
	return []interface{}{transformed}
}
func flattenApphubApplicationScopeType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApphubApplicationUid(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApphubApplicationState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandApphubApplicationDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApphubApplicationDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApphubApplicationAttributes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCriticality, err := expandApphubApplicationAttributesCriticality(original["criticality"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCriticality); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["criticality"] = transformedCriticality
	}

	transformedEnvironment, err := expandApphubApplicationAttributesEnvironment(original["environment"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnvironment); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["environment"] = transformedEnvironment
	}

	transformedDeveloperOwners, err := expandApphubApplicationAttributesDeveloperOwners(original["developer_owners"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDeveloperOwners); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["developerOwners"] = transformedDeveloperOwners
	}

	transformedOperatorOwners, err := expandApphubApplicationAttributesOperatorOwners(original["operator_owners"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOperatorOwners); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["operatorOwners"] = transformedOperatorOwners
	}

	transformedBusinessOwners, err := expandApphubApplicationAttributesBusinessOwners(original["business_owners"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBusinessOwners); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["businessOwners"] = transformedBusinessOwners
	}

	return transformed, nil
}

func expandApphubApplicationAttributesCriticality(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedType, err := expandApphubApplicationAttributesCriticalityType(original["type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["type"] = transformedType
	}

	return transformed, nil
}

func expandApphubApplicationAttributesCriticalityType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApphubApplicationAttributesEnvironment(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedType, err := expandApphubApplicationAttributesEnvironmentType(original["type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["type"] = transformedType
	}

	return transformed, nil
}

func expandApphubApplicationAttributesEnvironmentType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApphubApplicationAttributesDeveloperOwners(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDisplayName, err := expandApphubApplicationAttributesDeveloperOwnersDisplayName(original["display_name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDisplayName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["displayName"] = transformedDisplayName
		}

		transformedEmail, err := expandApphubApplicationAttributesDeveloperOwnersEmail(original["email"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEmail); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["email"] = transformedEmail
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandApphubApplicationAttributesDeveloperOwnersDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApphubApplicationAttributesDeveloperOwnersEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApphubApplicationAttributesOperatorOwners(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDisplayName, err := expandApphubApplicationAttributesOperatorOwnersDisplayName(original["display_name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDisplayName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["displayName"] = transformedDisplayName
		}

		transformedEmail, err := expandApphubApplicationAttributesOperatorOwnersEmail(original["email"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEmail); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["email"] = transformedEmail
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandApphubApplicationAttributesOperatorOwnersDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApphubApplicationAttributesOperatorOwnersEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApphubApplicationAttributesBusinessOwners(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDisplayName, err := expandApphubApplicationAttributesBusinessOwnersDisplayName(original["display_name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDisplayName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["displayName"] = transformedDisplayName
		}

		transformedEmail, err := expandApphubApplicationAttributesBusinessOwnersEmail(original["email"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEmail); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["email"] = transformedEmail
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandApphubApplicationAttributesBusinessOwnersDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApphubApplicationAttributesBusinessOwnersEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApphubApplicationScope(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedType, err := expandApphubApplicationScopeType(original["type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["type"] = transformedType
	}

	return transformed, nil
}

func expandApphubApplicationScopeType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
