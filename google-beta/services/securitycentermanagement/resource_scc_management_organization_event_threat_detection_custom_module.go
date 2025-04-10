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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/securitycentermanagement/OrganizationEventThreatDetectionCustomModule.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package securitycentermanagement

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/structure"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceSecurityCenterManagementOrganizationEventThreatDetectionCustomModule() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleCreate,
		Read:   resourceSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleRead,
		Update: resourceSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleUpdate,
		Delete: resourceSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleDelete,

		Importer: &schema.ResourceImporter{
			State: resourceSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"organization": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Numerical ID of the parent organization.`,
			},
			"config": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsJSON,
				StateFunc:    func(v interface{}) string { s, _ := structure.NormalizeJsonString(v); return s },
				Description: `Config for the module. For the resident module, its config value is defined at this level.
For the inherited module, its config value is inherited from the ancestor module.`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The human readable name to be displayed for the module.`,
			},
			"enablement_state": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: verify.ValidateEnum([]string{"ENABLED", "DISABLED", ""}),
				Description:  `The state of enablement for the module at the given level of the hierarchy. Possible values: ["ENABLED", "DISABLED"]`,
			},
			"location": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `Location ID of the parent organization. Only global is supported at the moment.`,
				Default:     "global",
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `Immutable. Type for the module. e.g. CONFIGURABLE_BAD_IP.`,
			},
			"last_editor": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The editor that last updated the custom module`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The resource name of the Event Threat Detection custom module.
Its format is "organizations/{organization}/locations/{location}/eventThreatDetectionCustomModules/{eventThreatDetectionCustomModule}".`,
			},
			"update_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The time at which the custom module was last updated.

A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and
up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	configProp, err := expandSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleConfig(d.Get("config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("config"); !tpgresource.IsEmptyValue(reflect.ValueOf(configProp)) && (ok || !reflect.DeepEqual(v, configProp)) {
		obj["config"] = configProp
	}
	enablementStateProp, err := expandSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleEnablementState(d.Get("enablement_state"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enablement_state"); !tpgresource.IsEmptyValue(reflect.ValueOf(enablementStateProp)) && (ok || !reflect.DeepEqual(v, enablementStateProp)) {
		obj["enablementState"] = enablementStateProp
	}
	typeProp, err := expandSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	displayNameProp, err := expandSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}

	lockName, err := tpgresource.ReplaceVars(d, config, "organizations/{{organization}}/locations/{{location}}/eventThreatDetectionCustomModules")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{SecurityCenterManagementBasePath}}organizations/{{organization}}/locations/{{location}}/eventThreatDetectionCustomModules")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new OrganizationEventThreatDetectionCustomModule: %#v", obj)
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
		return fmt.Errorf("Error creating OrganizationEventThreatDetectionCustomModule: %s", err)
	}
	// Set computed resource properties from create API response so that they're available on the subsequent Read
	// call.
	if err := d.Set("name", flattenSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "organizations/{{organization}}/locations/{{location}}/eventThreatDetectionCustomModules/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating OrganizationEventThreatDetectionCustomModule %q: %#v", d.Id(), res)

	return resourceSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleRead(d, meta)
}

func resourceSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecurityCenterManagementBasePath}}organizations/{{organization}}/locations/{{location}}/eventThreatDetectionCustomModules/{{name}}")
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("SecurityCenterManagementOrganizationEventThreatDetectionCustomModule %q", d.Id()))
	}

	if err := d.Set("name", flattenSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationEventThreatDetectionCustomModule: %s", err)
	}
	if err := d.Set("config", flattenSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleConfig(res["config"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationEventThreatDetectionCustomModule: %s", err)
	}
	if err := d.Set("enablement_state", flattenSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleEnablementState(res["enablementState"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationEventThreatDetectionCustomModule: %s", err)
	}
	if err := d.Set("type", flattenSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleType(res["type"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationEventThreatDetectionCustomModule: %s", err)
	}
	if err := d.Set("display_name", flattenSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationEventThreatDetectionCustomModule: %s", err)
	}
	if err := d.Set("update_time", flattenSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationEventThreatDetectionCustomModule: %s", err)
	}
	if err := d.Set("last_editor", flattenSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleLastEditor(res["lastEditor"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationEventThreatDetectionCustomModule: %s", err)
	}

	return nil
}

func resourceSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	configProp, err := expandSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleConfig(d.Get("config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("config"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, configProp)) {
		obj["config"] = configProp
	}
	enablementStateProp, err := expandSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleEnablementState(d.Get("enablement_state"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enablement_state"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, enablementStateProp)) {
		obj["enablementState"] = enablementStateProp
	}
	displayNameProp, err := expandSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}

	lockName, err := tpgresource.ReplaceVars(d, config, "organizations/{{organization}}/locations/{{location}}/eventThreatDetectionCustomModules")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{SecurityCenterManagementBasePath}}organizations/{{organization}}/locations/{{location}}/eventThreatDetectionCustomModules/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating OrganizationEventThreatDetectionCustomModule %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("config") {
		updateMask = append(updateMask, "config")
	}

	if d.HasChange("enablement_state") {
		updateMask = append(updateMask, "enablementState")
	}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
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
			return fmt.Errorf("Error updating OrganizationEventThreatDetectionCustomModule %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating OrganizationEventThreatDetectionCustomModule %q: %#v", d.Id(), res)
		}

	}

	return resourceSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleRead(d, meta)
}

func resourceSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	lockName, err := tpgresource.ReplaceVars(d, config, "organizations/{{organization}}/locations/{{location}}/eventThreatDetectionCustomModules")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{SecurityCenterManagementBasePath}}organizations/{{organization}}/locations/{{location}}/eventThreatDetectionCustomModules/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting OrganizationEventThreatDetectionCustomModule %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "OrganizationEventThreatDetectionCustomModule")
	}

	log.Printf("[DEBUG] Finished deleting OrganizationEventThreatDetectionCustomModule %q: %#v", d.Id(), res)
	return nil
}

func resourceSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^organizations/(?P<organization>[^/]+)/locations/(?P<location>[^/]+)/eventThreatDetectionCustomModules/(?P<name>[^/]+)$",
		"^(?P<organization>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "organizations/{{organization}}/locations/{{location}}/eventThreatDetectionCustomModules/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func flattenSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	b, err := json.Marshal(v)
	if err != nil {
		// TODO: return error once https://github.com/GoogleCloudPlatform/magic-modules/issues/3257 is fixed.
		log.Printf("[ERROR] failed to marshal schema to JSON: %v", err)
	}
	return string(b)
}

func flattenSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleEnablementState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleLastEditor(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	b := []byte(v.(string))
	if len(b) == 0 {
		return nil, nil
	}
	m := make(map[string]interface{})
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func expandSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleEnablementState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterManagementOrganizationEventThreatDetectionCustomModuleDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
