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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/compute/FirewallPolicy.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package compute

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceComputeFirewallPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeFirewallPolicyCreate,
		Read:   resourceComputeFirewallPolicyRead,
		Update: resourceComputeFirewallPolicyUpdate,
		Delete: resourceComputeFirewallPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeFirewallPolicyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"parent": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The parent of the firewall policy.`,
			},
			"short_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `User-provided name of the Organization firewall policy. The name should be unique in the organization in which the firewall policy is created.
This field is not applicable to network firewall policies. This name must be set on creation and cannot be changed. The name must be 1-63 characters long, and comply with RFC1035.
Specifically, the name must be 1-63 characters long and match the regular expression [a-z]([-a-z0-9]*[a-z0-9])? which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `An optional description of this resource. Provide this property when you create the resource.`,
			},
			"creation_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC3339 text format.`,
			},
			"fingerprint": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Fingerprint of the resource. This field is used internally during updates of this resource.`,
			},
			"firewall_policy_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The unique identifier for the resource. This identifier is defined by the server.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Name of the resource. It is a numeric ID allocated by GCP which uniquely identifies the Firewall Policy.`,
			},
			"rule_tuple_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `Total count of all firewall policy rule tuples. A firewall policy can not exceed a set number of tuples.`,
			},
			"self_link": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Server-defined URL for the resource.`,
			},
			"self_link_with_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Server-defined URL for this resource with the resource id.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceComputeFirewallPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	shortNameProp, err := expandComputeFirewallPolicyShortName(d.Get("short_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("short_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(shortNameProp)) && (ok || !reflect.DeepEqual(v, shortNameProp)) {
		obj["shortName"] = shortNameProp
	}
	descriptionProp, err := expandComputeFirewallPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	parentProp, err := expandComputeFirewallPolicyParent(d.Get("parent"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("parent"); !tpgresource.IsEmptyValue(reflect.ValueOf(parentProp)) && (ok || !reflect.DeepEqual(v, parentProp)) {
		obj["parent"] = parentProp
	}
	fingerprintProp, err := expandComputeFirewallPolicyFingerprint(d.Get("fingerprint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("fingerprint"); !tpgresource.IsEmptyValue(reflect.ValueOf(fingerprintProp)) && (ok || !reflect.DeepEqual(v, fingerprintProp)) {
		obj["fingerprint"] = fingerprintProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}locations/global/firewallPolicies?parentId={{parent}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new FirewallPolicy: %#v", obj)
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
		return fmt.Errorf("Error creating FirewallPolicy: %s", err)
	}
	// Set computed resource properties from create API response so that they're available on the subsequent Read
	// call.
	if err := d.Set("name", flattenComputeFirewallPolicyName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "locations/global/firewallPolicies/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	var opRes map[string]interface{}
	err = ComputeOrgOperationWaitTimeWithResponse(
		config, res, &opRes, d.Get("parent").(string), "FirewallPolicy operation", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create FirewallPolicy: %s", err)
	}

	firewallPolicyId, ok := opRes["targetId"]
	if !ok {
		return fmt.Errorf("Create response didn't contain targetId. Create may not have succeeded.")
	}
	if err := d.Set("name", firewallPolicyId.(string)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err = tpgresource.ReplaceVars(d, config, "locations/global/firewallPolicies/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating FirewallPolicy %q: %#v", d.Id(), res)

	return resourceComputeFirewallPolicyRead(d, meta)
}

func resourceComputeFirewallPolicyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}locations/global/firewallPolicies/{{name}}")
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ComputeFirewallPolicy %q", d.Id()))
	}

	if err := d.Set("creation_timestamp", flattenComputeFirewallPolicyCreationTimestamp(res["creationTimestamp"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallPolicy: %s", err)
	}
	if err := d.Set("name", flattenComputeFirewallPolicyName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallPolicy: %s", err)
	}
	if err := d.Set("firewall_policy_id", flattenComputeFirewallPolicyFirewallPolicyId(res["id"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallPolicy: %s", err)
	}
	if err := d.Set("short_name", flattenComputeFirewallPolicyShortName(res["shortName"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallPolicy: %s", err)
	}
	if err := d.Set("description", flattenComputeFirewallPolicyDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallPolicy: %s", err)
	}
	if err := d.Set("parent", flattenComputeFirewallPolicyParent(res["parent"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallPolicy: %s", err)
	}
	if err := d.Set("fingerprint", flattenComputeFirewallPolicyFingerprint(res["fingerprint"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallPolicy: %s", err)
	}
	if err := d.Set("self_link", flattenComputeFirewallPolicySelfLink(res["selfLink"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallPolicy: %s", err)
	}
	if err := d.Set("self_link_with_id", flattenComputeFirewallPolicySelfLinkWithId(res["selfLinkWithId"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallPolicy: %s", err)
	}
	if err := d.Set("rule_tuple_count", flattenComputeFirewallPolicyRuleTupleCount(res["ruleTupleCount"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallPolicy: %s", err)
	}

	return nil
}

func resourceComputeFirewallPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeFirewallPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	fingerprintProp, err := expandComputeFirewallPolicyFingerprint(d.Get("fingerprint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("fingerprint"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, fingerprintProp)) {
		obj["fingerprint"] = fingerprintProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}locations/global/firewallPolicies/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating FirewallPolicy %q: %#v", d.Id(), obj)
	headers := make(http.Header)

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

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
		return fmt.Errorf("Error updating FirewallPolicy %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating FirewallPolicy %q: %#v", d.Id(), res)
	}

	var opRes map[string]interface{}
	err = ComputeOrgOperationWaitTimeWithResponse(
		config, res, &opRes, d.Get("parent").(string), "FirewallPolicy operation", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		return fmt.Errorf("Error waiting for FirewallPolicy operation: %s", err)
	}
	return resourceComputeFirewallPolicyRead(d, meta)
}

func resourceComputeFirewallPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}locations/global/firewallPolicies/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting FirewallPolicy %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "FirewallPolicy")
	}

	var opRes map[string]interface{}
	err = ComputeOrgOperationWaitTimeWithResponse(
		config, res, &opRes, d.Get("parent").(string), "FirewallPolicy operation", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		return fmt.Errorf("Error waiting for FirewallPolicy operation: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting FirewallPolicy %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeFirewallPolicyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^locations/global/firewallPolicies/(?P<name>[^/]+)$",
		"^(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "locations/global/firewallPolicies/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeFirewallPolicyCreationTimestamp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallPolicyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallPolicyFirewallPolicyId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallPolicyShortName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallPolicyDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallPolicyParent(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallPolicyFingerprint(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallPolicySelfLink(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallPolicySelfLinkWithId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallPolicyRuleTupleCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func expandComputeFirewallPolicyShortName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyParent(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyFingerprint(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
