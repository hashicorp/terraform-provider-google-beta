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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/spanner/InstanceConfig.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package spanner

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func replicasHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(m["location"].(string)))) // ToLower just in case
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(m["type"].(string))))
	var isLeader interface{}
	if m["defaultLeaderLocation"] != nil {
		isLeader = m["defaultLeaderLocation"]
	} else {
		isLeader = false
	}
	buf.WriteString(fmt.Sprintf("%v-", isLeader.(bool)))
	return tpgresource.Hashcode(buf.String())
}

func getBaseInstanceConfigReplicas(d *schema.ResourceData, config *transport_tpg.Config, baseConfigProp interface{}, billingProject, userAgent string) ([]interface{}, error) {
	url, err := tpgresource.ReplaceVars(d, config, "{{SpannerBasePath}}")
	if err != nil {
		return nil, err
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    fmt.Sprintf("%s%s", url, baseConfigProp.(string)),
		UserAgent: userAgent,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return nil, fmt.Errorf("Error fetching base InstanceConfig: %s", err)
	}

	data, ok := res["replicas"]
	if !ok || data == nil {
		log.Print("[DEBUG] No replicas in the base InstanceConfig.")
		return nil, nil
	}

	return data.([]interface{}), nil
}

func ResourceSpannerInstanceConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceSpannerInstanceConfigCreate,
		Read:   resourceSpannerInstanceConfigRead,
		Update: resourceSpannerInstanceConfigUpdate,
		Delete: resourceSpannerInstanceConfigDelete,

		Importer: &schema.ResourceImporter{
			State: resourceSpannerInstanceConfigImport,
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
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The name of this instance configuration as it appears in UIs.`,
			},
			"replicas": {
				Type:        schema.TypeSet,
				Required:    true,
				ForceNew:    true,
				Description: `The geographic placement of nodes in this instance configuration and their replication properties.`,
				Elem:        spannerInstanceConfigReplicasSchema(),
				Set:         replicasHash,
			},
			"base_config": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
				Description: `Base configuration name, e.g. nam3, based on which this configuration is created.
Only set for user managed configurations.
baseConfig must refer to a configuration of type GOOGLE_MANAGED in the same project as this configuration.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `An object containing a list of "key": value pairs.
Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.


**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
				Description: `A unique identifier for the instance configuration. Values are of the
form projects/<project>/instanceConfigs/[a-z][-a-z0-9]*`,
			},
			"config_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. Whether this instance config is a Google or User Managed Configuration.`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
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

func spannerInstanceConfigReplicasSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"default_leader_location": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Description: `If true, this location is designated as the default leader location where
leader replicas are placed.`,
				Default: false,
			},
			"location": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The location of the serving resources, e.g. "us-central1".`,
			},
			"type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"READ_WRITE", "READ_ONLY", "WITNESS", ""}),
				Description: `Indicates the type of replica.  See the [replica types
documentation](https://cloud.google.com/spanner/docs/replication#replica_types)
for more details. Possible values: ["READ_WRITE", "READ_ONLY", "WITNESS"]`,
			},
		},
	}
}

func resourceSpannerInstanceConfigCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandSpannerInstanceConfigName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	displayNameProp, err := expandSpannerInstanceConfigDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	baseConfigProp, err := expandSpannerInstanceConfigBaseConfig(d.Get("base_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("base_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(baseConfigProp)) && (ok || !reflect.DeepEqual(v, baseConfigProp)) {
		obj["baseConfig"] = baseConfigProp
	}
	replicasProp, err := expandSpannerInstanceConfigReplicas(d.Get("replicas"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("replicas"); !tpgresource.IsEmptyValue(reflect.ValueOf(replicasProp)) && (ok || !reflect.DeepEqual(v, replicasProp)) {
		obj["replicas"] = replicasProp
	}
	labelsProp, err := expandSpannerInstanceConfigEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	obj, err = resourceSpannerInstanceConfigEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SpannerBasePath}}projects/{{project}}/instanceConfigs")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new InstanceConfig: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for InstanceConfig: %s", err)
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
		return fmt.Errorf("Error creating InstanceConfig: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{project}}/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = SpannerOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating InstanceConfig", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create InstanceConfig: %s", err)
	}

	opRes, err = resourceSpannerInstanceConfigDecoder(d, meta, opRes)
	if err != nil {
		return fmt.Errorf("Error decoding response from operation: %s", err)
	}
	if opRes == nil {
		return fmt.Errorf("Error decoding response from operation, could not find object")
	}

	if err := d.Set("name", flattenSpannerInstanceConfigName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "{{project}}/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating InstanceConfig %q: %#v", d.Id(), res)

	return resourceSpannerInstanceConfigRead(d, meta)
}

func resourceSpannerInstanceConfigRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SpannerBasePath}}projects/{{project}}/instanceConfigs/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for InstanceConfig: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("SpannerInstanceConfig %q", d.Id()))
	}

	res, err = resourceSpannerInstanceConfigDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing SpannerInstanceConfig because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading InstanceConfig: %s", err)
	}

	if err := d.Set("name", flattenSpannerInstanceConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading InstanceConfig: %s", err)
	}
	if err := d.Set("display_name", flattenSpannerInstanceConfigDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading InstanceConfig: %s", err)
	}
	if err := d.Set("base_config", flattenSpannerInstanceConfigBaseConfig(res["baseConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading InstanceConfig: %s", err)
	}
	if err := d.Set("config_type", flattenSpannerInstanceConfigConfigType(res["configType"], d, config)); err != nil {
		return fmt.Errorf("Error reading InstanceConfig: %s", err)
	}
	if err := d.Set("replicas", flattenSpannerInstanceConfigReplicas(res["replicas"], d, config)); err != nil {
		return fmt.Errorf("Error reading InstanceConfig: %s", err)
	}
	if err := d.Set("labels", flattenSpannerInstanceConfigLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading InstanceConfig: %s", err)
	}
	if err := d.Set("terraform_labels", flattenSpannerInstanceConfigTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading InstanceConfig: %s", err)
	}
	if err := d.Set("effective_labels", flattenSpannerInstanceConfigEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading InstanceConfig: %s", err)
	}

	return nil
}

func resourceSpannerInstanceConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for InstanceConfig: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandSpannerInstanceConfigDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	labelsProp, err := expandSpannerInstanceConfigEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	obj, err = resourceSpannerInstanceConfigUpdateEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SpannerBasePath}}projects/{{project}}/instanceConfigs/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating InstanceConfig %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
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
			return fmt.Errorf("Error updating InstanceConfig %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating InstanceConfig %q: %#v", d.Id(), res)
		}

		err = SpannerOperationWaitTime(
			config, res, project, "Updating InstanceConfig", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceSpannerInstanceConfigRead(d, meta)
}

func resourceSpannerInstanceConfigDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for InstanceConfig: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{SpannerBasePath}}projects/{{project}}/instanceConfigs/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting InstanceConfig %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "InstanceConfig")
	}

	log.Printf("[DEBUG] Finished deleting InstanceConfig %q: %#v", d.Id(), res)
	return nil
}

func resourceSpannerInstanceConfigImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/instanceConfigs/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "{{project}}/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenSpannerInstanceConfigName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSpannerInstanceConfigDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSpannerInstanceConfigBaseConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func flattenSpannerInstanceConfigConfigType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSpannerInstanceConfigReplicas(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(replicasHash, []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"location":                flattenSpannerInstanceConfigReplicasLocation(original["location"], d, config),
			"type":                    flattenSpannerInstanceConfigReplicasType(original["type"], d, config),
			"default_leader_location": flattenSpannerInstanceConfigReplicasDefaultLeaderLocation(original["defaultLeaderLocation"], d, config),
		})
	}
	return transformed
}
func flattenSpannerInstanceConfigReplicasLocation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSpannerInstanceConfigReplicasType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSpannerInstanceConfigReplicasDefaultLeaderLocation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSpannerInstanceConfigLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenSpannerInstanceConfigTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenSpannerInstanceConfigEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandSpannerInstanceConfigName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerInstanceConfigDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerInstanceConfigBaseConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	r := regexp.MustCompile("projects/(.+)/instanceConfigs/(.+)")
	if r.MatchString(v.(string)) {
		return v.(string), nil
	}

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return nil, err
	}

	return fmt.Sprintf("projects/%s/instanceConfigs/%s", project, v.(string)), nil
}

func expandSpannerInstanceConfigReplicas(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedLocation, err := expandSpannerInstanceConfigReplicasLocation(original["location"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedLocation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["location"] = transformedLocation
		}

		transformedType, err := expandSpannerInstanceConfigReplicasType(original["type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["type"] = transformedType
		}

		transformedDefaultLeaderLocation, err := expandSpannerInstanceConfigReplicasDefaultLeaderLocation(original["default_leader_location"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDefaultLeaderLocation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["defaultLeaderLocation"] = transformedDefaultLeaderLocation
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandSpannerInstanceConfigReplicasLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerInstanceConfigReplicasType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerInstanceConfigReplicasDefaultLeaderLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerInstanceConfigEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func resourceSpannerInstanceConfigEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	config := meta.(*transport_tpg.Config)
	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return nil, err
	}
	newObj := make(map[string]interface{})
	if obj["name"] == nil {
		return nil, fmt.Errorf("Error setting instance config name")
	}
	newObj["instanceConfigId"] = obj["name"]
	obj["name"] = fmt.Sprintf("projects/%s/instanceConfigs/%s", project, obj["name"])
	baseReplicas, err := getBaseInstanceConfigReplicas(d, config, obj["baseConfig"], project, meta.(*transport_tpg.Config).UserAgent)
	if err != nil {
		return nil, err
	}
	r := obj["replicas"].([]interface{})
	obj["replicas"] = append(r, baseReplicas...)
	newObj["instanceConfig"] = obj
	return newObj, nil
}

func resourceSpannerInstanceConfigUpdateEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	project, err := tpgresource.GetProject(d, meta.(*transport_tpg.Config))
	if err != nil {
		return nil, err
	}
	obj["name"] = fmt.Sprintf("projects/%s/instanceConfigs/%s", project, obj["name"])
	newObj := make(map[string]interface{})
	newObj["instanceConfig"] = obj
	return newObj, nil
}

func resourceSpannerInstanceConfigDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	config := meta.(*transport_tpg.Config)
	d.SetId(res["name"].(string))
	if err := tpgresource.ParseImportId([]string{"projects/(?P<project>[^/]+)/instanceConfigs/(?P<name>[^/]+)"}, d, config); err != nil {
		return nil, err
	}
	res["project"] = d.Get("project").(string)
	res["name"] = d.Get("name").(string)
	id, err := tpgresource.ReplaceVars(d, config, "{{project}}/{{name}}")
	if err != nil {
		return nil, err
	}
	baseReplicas, err := getBaseInstanceConfigReplicas(d, config, res["baseConfig"], res["project"].(string), config.UserAgent)
	if err != nil {
		return nil, err
	}
	customReplica := make(map[int]interface{})
	for _, b := range baseReplicas {
		customReplica[replicasHash(b)] = b
	}
	var cR []interface{}
	for _, r := range res["replicas"].([]interface{}) {
		if _, ok := customReplica[replicasHash(r)]; !ok {
			cR = append(cR, r)
		}
	}
	res["replicas"] = cR
	d.SetId(id)
	return res, nil
}
