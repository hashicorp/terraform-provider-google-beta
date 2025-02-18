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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/networksecurity/GatewaySecurityPolicyRule.yaml
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
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceNetworkSecurityGatewaySecurityPolicyRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkSecurityGatewaySecurityPolicyRuleCreate,
		Read:   resourceNetworkSecurityGatewaySecurityPolicyRuleRead,
		Update: resourceNetworkSecurityGatewaySecurityPolicyRuleUpdate,
		Delete: resourceNetworkSecurityGatewaySecurityPolicyRuleDelete,

		Importer: &schema.ResourceImporter{
			State: resourceNetworkSecurityGatewaySecurityPolicyRuleImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"basic_profile": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: verify.ValidateEnum([]string{"BASIC_PROFILE_UNSPECIFIED", "ALLOW", "DENY"}),
				Description:  `Profile which tells what the primitive action should be. Possible values are: * ALLOW * DENY. Possible values: ["BASIC_PROFILE_UNSPECIFIED", "ALLOW", "DENY"]`,
			},
			"enabled": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: `Whether the rule is enforced.`,
			},
			"gateway_security_policy": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the gatewat security policy this rule belongs to.`,
			},
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The location of the gateway security policy.`,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Name of the resource. ame is the full resource name so projects/{project}/locations/{location}/gatewaySecurityPolicies/{gateway_security_policy}/rules/{rule}
rule should match the pattern: (^a-z?$).`,
			},
			"priority": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: `Priority of the rule. Lower number corresponds to higher precedence.`,
			},
			"session_matcher": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `CEL expression for matching on session criteria.`,
			},
			"application_matcher": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `CEL expression for matching on L7/application level criteria.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Free-text description of the resource.`,
			},
			"tls_inspection_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: `Flag to enable TLS inspection of traffic matching on. Can only be true if the
parent GatewaySecurityPolicy references a TLSInspectionConfig.`,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The timestamp when the resource was created.
A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z"`,
			},
			"self_link": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Server-defined URL of this resource.`,
			},
			"update_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The timestamp when the resource was updated.
A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".`,
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

func resourceNetworkSecurityGatewaySecurityPolicyRuleCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	enabledProp, err := expandNetworkSecurityGatewaySecurityPolicyRuleEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(enabledProp)) && (ok || !reflect.DeepEqual(v, enabledProp)) {
		obj["enabled"] = enabledProp
	}
	priorityProp, err := expandNetworkSecurityGatewaySecurityPolicyRulePriority(d.Get("priority"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("priority"); !tpgresource.IsEmptyValue(reflect.ValueOf(priorityProp)) && (ok || !reflect.DeepEqual(v, priorityProp)) {
		obj["priority"] = priorityProp
	}
	descriptionProp, err := expandNetworkSecurityGatewaySecurityPolicyRuleDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	sessionMatcherProp, err := expandNetworkSecurityGatewaySecurityPolicyRuleSessionMatcher(d.Get("session_matcher"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("session_matcher"); !tpgresource.IsEmptyValue(reflect.ValueOf(sessionMatcherProp)) && (ok || !reflect.DeepEqual(v, sessionMatcherProp)) {
		obj["sessionMatcher"] = sessionMatcherProp
	}
	applicationMatcherProp, err := expandNetworkSecurityGatewaySecurityPolicyRuleApplicationMatcher(d.Get("application_matcher"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("application_matcher"); !tpgresource.IsEmptyValue(reflect.ValueOf(applicationMatcherProp)) && (ok || !reflect.DeepEqual(v, applicationMatcherProp)) {
		obj["applicationMatcher"] = applicationMatcherProp
	}
	tlsInspectionEnabledProp, err := expandNetworkSecurityGatewaySecurityPolicyRuleTlsInspectionEnabled(d.Get("tls_inspection_enabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("tls_inspection_enabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(tlsInspectionEnabledProp)) && (ok || !reflect.DeepEqual(v, tlsInspectionEnabledProp)) {
		obj["tlsInspectionEnabled"] = tlsInspectionEnabledProp
	}
	basicProfileProp, err := expandNetworkSecurityGatewaySecurityPolicyRuleBasicProfile(d.Get("basic_profile"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("basic_profile"); !tpgresource.IsEmptyValue(reflect.ValueOf(basicProfileProp)) && (ok || !reflect.DeepEqual(v, basicProfileProp)) {
		obj["basicProfile"] = basicProfileProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/gatewaySecurityPolicies/{{gateway_security_policy}}/rules?gatewaySecurityPolicyRuleId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new GatewaySecurityPolicyRule: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for GatewaySecurityPolicyRule: %s", err)
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
		return fmt.Errorf("Error creating GatewaySecurityPolicyRule: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/gatewaySecurityPolicies/{{gateway_security_policy}}/rules/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = NetworkSecurityOperationWaitTime(
		config, res, project, "Creating GatewaySecurityPolicyRule", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create GatewaySecurityPolicyRule: %s", err)
	}

	log.Printf("[DEBUG] Finished creating GatewaySecurityPolicyRule %q: %#v", d.Id(), res)

	return resourceNetworkSecurityGatewaySecurityPolicyRuleRead(d, meta)
}

func resourceNetworkSecurityGatewaySecurityPolicyRuleRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/gatewaySecurityPolicies/{{gateway_security_policy}}/rules/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for GatewaySecurityPolicyRule: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("NetworkSecurityGatewaySecurityPolicyRule %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading GatewaySecurityPolicyRule: %s", err)
	}

	if err := d.Set("self_link", flattenNetworkSecurityGatewaySecurityPolicyRuleSelfLink(res["selfLink"], d, config)); err != nil {
		return fmt.Errorf("Error reading GatewaySecurityPolicyRule: %s", err)
	}
	if err := d.Set("create_time", flattenNetworkSecurityGatewaySecurityPolicyRuleCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading GatewaySecurityPolicyRule: %s", err)
	}
	if err := d.Set("update_time", flattenNetworkSecurityGatewaySecurityPolicyRuleUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading GatewaySecurityPolicyRule: %s", err)
	}
	if err := d.Set("enabled", flattenNetworkSecurityGatewaySecurityPolicyRuleEnabled(res["enabled"], d, config)); err != nil {
		return fmt.Errorf("Error reading GatewaySecurityPolicyRule: %s", err)
	}
	if err := d.Set("priority", flattenNetworkSecurityGatewaySecurityPolicyRulePriority(res["priority"], d, config)); err != nil {
		return fmt.Errorf("Error reading GatewaySecurityPolicyRule: %s", err)
	}
	if err := d.Set("description", flattenNetworkSecurityGatewaySecurityPolicyRuleDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading GatewaySecurityPolicyRule: %s", err)
	}
	if err := d.Set("session_matcher", flattenNetworkSecurityGatewaySecurityPolicyRuleSessionMatcher(res["sessionMatcher"], d, config)); err != nil {
		return fmt.Errorf("Error reading GatewaySecurityPolicyRule: %s", err)
	}
	if err := d.Set("application_matcher", flattenNetworkSecurityGatewaySecurityPolicyRuleApplicationMatcher(res["applicationMatcher"], d, config)); err != nil {
		return fmt.Errorf("Error reading GatewaySecurityPolicyRule: %s", err)
	}
	if err := d.Set("tls_inspection_enabled", flattenNetworkSecurityGatewaySecurityPolicyRuleTlsInspectionEnabled(res["tlsInspectionEnabled"], d, config)); err != nil {
		return fmt.Errorf("Error reading GatewaySecurityPolicyRule: %s", err)
	}
	if err := d.Set("basic_profile", flattenNetworkSecurityGatewaySecurityPolicyRuleBasicProfile(res["basicProfile"], d, config)); err != nil {
		return fmt.Errorf("Error reading GatewaySecurityPolicyRule: %s", err)
	}

	return nil
}

func resourceNetworkSecurityGatewaySecurityPolicyRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for GatewaySecurityPolicyRule: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	enabledProp, err := expandNetworkSecurityGatewaySecurityPolicyRuleEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, enabledProp)) {
		obj["enabled"] = enabledProp
	}
	priorityProp, err := expandNetworkSecurityGatewaySecurityPolicyRulePriority(d.Get("priority"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("priority"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, priorityProp)) {
		obj["priority"] = priorityProp
	}
	descriptionProp, err := expandNetworkSecurityGatewaySecurityPolicyRuleDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	sessionMatcherProp, err := expandNetworkSecurityGatewaySecurityPolicyRuleSessionMatcher(d.Get("session_matcher"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("session_matcher"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, sessionMatcherProp)) {
		obj["sessionMatcher"] = sessionMatcherProp
	}
	applicationMatcherProp, err := expandNetworkSecurityGatewaySecurityPolicyRuleApplicationMatcher(d.Get("application_matcher"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("application_matcher"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, applicationMatcherProp)) {
		obj["applicationMatcher"] = applicationMatcherProp
	}
	tlsInspectionEnabledProp, err := expandNetworkSecurityGatewaySecurityPolicyRuleTlsInspectionEnabled(d.Get("tls_inspection_enabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("tls_inspection_enabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, tlsInspectionEnabledProp)) {
		obj["tlsInspectionEnabled"] = tlsInspectionEnabledProp
	}
	basicProfileProp, err := expandNetworkSecurityGatewaySecurityPolicyRuleBasicProfile(d.Get("basic_profile"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("basic_profile"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, basicProfileProp)) {
		obj["basicProfile"] = basicProfileProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/gatewaySecurityPolicies/{{gateway_security_policy}}/rules/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating GatewaySecurityPolicyRule %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("enabled") {
		updateMask = append(updateMask, "enabled")
	}

	if d.HasChange("priority") {
		updateMask = append(updateMask, "priority")
	}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("session_matcher") {
		updateMask = append(updateMask, "sessionMatcher")
	}

	if d.HasChange("application_matcher") {
		updateMask = append(updateMask, "applicationMatcher")
	}

	if d.HasChange("tls_inspection_enabled") {
		updateMask = append(updateMask, "tlsInspectionEnabled")
	}

	if d.HasChange("basic_profile") {
		updateMask = append(updateMask, "basicProfile")
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
			return fmt.Errorf("Error updating GatewaySecurityPolicyRule %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating GatewaySecurityPolicyRule %q: %#v", d.Id(), res)
		}

		err = NetworkSecurityOperationWaitTime(
			config, res, project, "Updating GatewaySecurityPolicyRule", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceNetworkSecurityGatewaySecurityPolicyRuleRead(d, meta)
}

func resourceNetworkSecurityGatewaySecurityPolicyRuleDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for GatewaySecurityPolicyRule: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/gatewaySecurityPolicies/{{gateway_security_policy}}/rules/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting GatewaySecurityPolicyRule %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "GatewaySecurityPolicyRule")
	}

	err = NetworkSecurityOperationWaitTime(
		config, res, project, "Deleting GatewaySecurityPolicyRule", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting GatewaySecurityPolicyRule %q: %#v", d.Id(), res)
	return nil
}

func resourceNetworkSecurityGatewaySecurityPolicyRuleImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/gatewaySecurityPolicies/(?P<gateway_security_policy>[^/]+)/rules/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<gateway_security_policy>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<gateway_security_policy>[^/]+)/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/gatewaySecurityPolicies/{{gateway_security_policy}}/rules/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNetworkSecurityGatewaySecurityPolicyRuleSelfLink(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityGatewaySecurityPolicyRuleCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityGatewaySecurityPolicyRuleUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityGatewaySecurityPolicyRuleEnabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityGatewaySecurityPolicyRulePriority(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenNetworkSecurityGatewaySecurityPolicyRuleDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityGatewaySecurityPolicyRuleSessionMatcher(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityGatewaySecurityPolicyRuleApplicationMatcher(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityGatewaySecurityPolicyRuleTlsInspectionEnabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityGatewaySecurityPolicyRuleBasicProfile(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNetworkSecurityGatewaySecurityPolicyRuleEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityGatewaySecurityPolicyRulePriority(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityGatewaySecurityPolicyRuleDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityGatewaySecurityPolicyRuleSessionMatcher(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityGatewaySecurityPolicyRuleApplicationMatcher(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityGatewaySecurityPolicyRuleTlsInspectionEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityGatewaySecurityPolicyRuleBasicProfile(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
