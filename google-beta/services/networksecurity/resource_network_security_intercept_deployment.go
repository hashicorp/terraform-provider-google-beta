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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/networksecurity/InterceptDeployment.yaml
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

func ResourceNetworkSecurityInterceptDeployment() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkSecurityInterceptDeploymentCreate,
		Read:   resourceNetworkSecurityInterceptDeploymentRead,
		Update: resourceNetworkSecurityInterceptDeploymentUpdate,
		Delete: resourceNetworkSecurityInterceptDeploymentDelete,

		Importer: &schema.ResourceImporter{
			State: resourceNetworkSecurityInterceptDeploymentImport,
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
			"forwarding_rule": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The regional forwarding rule that fronts the interceptors, for example:
'projects/123456789/regions/us-central1/forwardingRules/my-rule'.
See https://google.aip.dev/124.`,
			},
			"intercept_deployment_group": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The deployment group that this deployment is a part of, for example:
'projects/123456789/locations/global/interceptDeploymentGroups/my-dg'.
See https://google.aip.dev/124.`,
			},
			"intercept_deployment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The ID to use for the new deployment, which will become the final
component of the deployment's resource name.`,
			},
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The cloud location of the deployment, e.g. 'us-central1-a' or 'asia-south1-b'.`,
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
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The resource name of this deployment, for example:
'projects/123456789/locations/us-central1-a/interceptDeployments/my-dep'.
See https://google.aip.dev/122 for more details.`,
			},
			"reconciling": {
				Type:     schema.TypeBool,
				Computed: true,
				Description: `The current state of the resource does not match the user's intended state,
and the system is working to reconcile them. This part of the normal
operation (e.g. linking a new association to the parent group).
See https://google.aip.dev/128.`,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The current state of the deployment.
See https://google.aip.dev/216.
Possible values:
STATE_UNSPECIFIED
ACTIVE
CREATING
DELETING
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

func resourceNetworkSecurityInterceptDeploymentCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	forwardingRuleProp, err := expandNetworkSecurityInterceptDeploymentForwardingRule(d.Get("forwarding_rule"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("forwarding_rule"); !tpgresource.IsEmptyValue(reflect.ValueOf(forwardingRuleProp)) && (ok || !reflect.DeepEqual(v, forwardingRuleProp)) {
		obj["forwardingRule"] = forwardingRuleProp
	}
	interceptDeploymentGroupProp, err := expandNetworkSecurityInterceptDeploymentInterceptDeploymentGroup(d.Get("intercept_deployment_group"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("intercept_deployment_group"); !tpgresource.IsEmptyValue(reflect.ValueOf(interceptDeploymentGroupProp)) && (ok || !reflect.DeepEqual(v, interceptDeploymentGroupProp)) {
		obj["interceptDeploymentGroup"] = interceptDeploymentGroupProp
	}
	labelsProp, err := expandNetworkSecurityInterceptDeploymentEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/interceptDeployments?interceptDeploymentId={{intercept_deployment_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new InterceptDeployment: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for InterceptDeployment: %s", err)
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
		return fmt.Errorf("Error creating InterceptDeployment: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/interceptDeployments/{{intercept_deployment_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = NetworkSecurityOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating InterceptDeployment", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create InterceptDeployment: %s", err)
	}

	if err := d.Set("name", flattenNetworkSecurityInterceptDeploymentName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/interceptDeployments/{{intercept_deployment_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating InterceptDeployment %q: %#v", d.Id(), res)

	return resourceNetworkSecurityInterceptDeploymentRead(d, meta)
}

func resourceNetworkSecurityInterceptDeploymentRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/interceptDeployments/{{intercept_deployment_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for InterceptDeployment: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("NetworkSecurityInterceptDeployment %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading InterceptDeployment: %s", err)
	}

	if err := d.Set("name", flattenNetworkSecurityInterceptDeploymentName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterceptDeployment: %s", err)
	}
	if err := d.Set("create_time", flattenNetworkSecurityInterceptDeploymentCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterceptDeployment: %s", err)
	}
	if err := d.Set("update_time", flattenNetworkSecurityInterceptDeploymentUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterceptDeployment: %s", err)
	}
	if err := d.Set("labels", flattenNetworkSecurityInterceptDeploymentLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterceptDeployment: %s", err)
	}
	if err := d.Set("forwarding_rule", flattenNetworkSecurityInterceptDeploymentForwardingRule(res["forwardingRule"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterceptDeployment: %s", err)
	}
	if err := d.Set("intercept_deployment_group", flattenNetworkSecurityInterceptDeploymentInterceptDeploymentGroup(res["interceptDeploymentGroup"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterceptDeployment: %s", err)
	}
	if err := d.Set("state", flattenNetworkSecurityInterceptDeploymentState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterceptDeployment: %s", err)
	}
	if err := d.Set("reconciling", flattenNetworkSecurityInterceptDeploymentReconciling(res["reconciling"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterceptDeployment: %s", err)
	}
	if err := d.Set("terraform_labels", flattenNetworkSecurityInterceptDeploymentTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterceptDeployment: %s", err)
	}
	if err := d.Set("effective_labels", flattenNetworkSecurityInterceptDeploymentEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterceptDeployment: %s", err)
	}

	return nil
}

func resourceNetworkSecurityInterceptDeploymentUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for InterceptDeployment: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	labelsProp, err := expandNetworkSecurityInterceptDeploymentEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/interceptDeployments/{{intercept_deployment_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating InterceptDeployment %q: %#v", d.Id(), obj)
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
			return fmt.Errorf("Error updating InterceptDeployment %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating InterceptDeployment %q: %#v", d.Id(), res)
		}

		err = NetworkSecurityOperationWaitTime(
			config, res, project, "Updating InterceptDeployment", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceNetworkSecurityInterceptDeploymentRead(d, meta)
}

func resourceNetworkSecurityInterceptDeploymentDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for InterceptDeployment: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/interceptDeployments/{{intercept_deployment_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting InterceptDeployment %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "InterceptDeployment")
	}

	err = NetworkSecurityOperationWaitTime(
		config, res, project, "Deleting InterceptDeployment", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting InterceptDeployment %q: %#v", d.Id(), res)
	return nil
}

func resourceNetworkSecurityInterceptDeploymentImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/interceptDeployments/(?P<intercept_deployment_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<intercept_deployment_id>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<intercept_deployment_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/interceptDeployments/{{intercept_deployment_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNetworkSecurityInterceptDeploymentName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityInterceptDeploymentCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityInterceptDeploymentUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityInterceptDeploymentLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenNetworkSecurityInterceptDeploymentForwardingRule(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityInterceptDeploymentInterceptDeploymentGroup(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityInterceptDeploymentState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityInterceptDeploymentReconciling(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityInterceptDeploymentTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenNetworkSecurityInterceptDeploymentEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNetworkSecurityInterceptDeploymentForwardingRule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityInterceptDeploymentInterceptDeploymentGroup(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityInterceptDeploymentEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
