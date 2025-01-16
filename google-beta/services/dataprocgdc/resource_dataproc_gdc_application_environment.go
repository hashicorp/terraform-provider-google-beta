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

package dataprocgdc

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

func ResourceDataprocGdcApplicationEnvironment() *schema.Resource {
	return &schema.Resource{
		Create: resourceDataprocGdcApplicationEnvironmentCreate,
		Read:   resourceDataprocGdcApplicationEnvironmentRead,
		Update: resourceDataprocGdcApplicationEnvironmentUpdate,
		Delete: resourceDataprocGdcApplicationEnvironmentDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDataprocGdcApplicationEnvironmentImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.SetAnnotationsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The location of the application environment`,
			},
			"serviceinstance": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The id of the service instance to which this application environment belongs.`,
			},
			"annotations": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `The annotations to associate with this application environment. Annotations may be used to store client information, but are not used by the server.

**Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
Please refer to the field 'effective_annotations' for all of the annotations present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"application_environment_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The id of the application environment`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `User-provided human-readable name to be used in user interfaces.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `The labels to associate with this application environment. Labels may be used for filtering and billing tracking. 

**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"namespace": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The name of the namespace in which to create this ApplicationEnvironment. This namespace must already exist in the cluster`,
			},
			"spark_application_environment_config": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Represents the SparkApplicationEnvironmentConfig.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default_properties": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: `A map of default Spark properties to apply to workloads in this application environment. These defaults may be overridden by per-application properties.`,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"default_version": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `The default Dataproc version to use for applications submitted to this application environment`,
						},
					},
				},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The timestamp when the resource was created.`,
			},
			"effective_annotations": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of annotations (key/value pairs) present on the resource in GCP, including the annotations configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
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
				Description: `Identifier. The name of the application environment. Format: projects/{project}/locations/{location}/serviceInstances/{service_instance}/applicationEnvironments/{application_environment_id}`,
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
				Description: `System generated unique identifier for this application environment, formatted as UUID4.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The timestamp when the resource was most recently updated.`,
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

func resourceDataprocGdcApplicationEnvironmentCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandDataprocGdcApplicationEnvironmentDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	sparkApplicationEnvironmentConfigProp, err := expandDataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfig(d.Get("spark_application_environment_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("spark_application_environment_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(sparkApplicationEnvironmentConfigProp)) && (ok || !reflect.DeepEqual(v, sparkApplicationEnvironmentConfigProp)) {
		obj["sparkApplicationEnvironmentConfig"] = sparkApplicationEnvironmentConfigProp
	}
	namespaceProp, err := expandDataprocGdcApplicationEnvironmentNamespace(d.Get("namespace"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("namespace"); !tpgresource.IsEmptyValue(reflect.ValueOf(namespaceProp)) && (ok || !reflect.DeepEqual(v, namespaceProp)) {
		obj["namespace"] = namespaceProp
	}
	labelsProp, err := expandDataprocGdcApplicationEnvironmentEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	annotationsProp, err := expandDataprocGdcApplicationEnvironmentEffectiveAnnotations(d.Get("effective_annotations"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_annotations"); !tpgresource.IsEmptyValue(reflect.ValueOf(annotationsProp)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DataprocGdcBasePath}}projects/{{project}}/locations/{{location}}/serviceInstances/{{serviceinstance}}/applicationEnvironments?applicationEnvironmentId={{application_environment_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ApplicationEnvironment: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ApplicationEnvironment: %s", err)
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
		return fmt.Errorf("Error creating ApplicationEnvironment: %s", err)
	}
	if err := d.Set("name", flattenDataprocGdcApplicationEnvironmentName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/serviceInstances/{{serviceinstance}}/applicationEnvironments/{{application_environment_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating ApplicationEnvironment %q: %#v", d.Id(), res)

	return resourceDataprocGdcApplicationEnvironmentRead(d, meta)
}

func resourceDataprocGdcApplicationEnvironmentRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DataprocGdcBasePath}}projects/{{project}}/locations/{{location}}/serviceInstances/{{serviceinstance}}/applicationEnvironments/{{application_environment_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ApplicationEnvironment: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("DataprocGdcApplicationEnvironment %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ApplicationEnvironment: %s", err)
	}

	if err := d.Set("name", flattenDataprocGdcApplicationEnvironmentName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading ApplicationEnvironment: %s", err)
	}
	if err := d.Set("uid", flattenDataprocGdcApplicationEnvironmentUid(res["uid"], d, config)); err != nil {
		return fmt.Errorf("Error reading ApplicationEnvironment: %s", err)
	}
	if err := d.Set("display_name", flattenDataprocGdcApplicationEnvironmentDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading ApplicationEnvironment: %s", err)
	}
	if err := d.Set("create_time", flattenDataprocGdcApplicationEnvironmentCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading ApplicationEnvironment: %s", err)
	}
	if err := d.Set("update_time", flattenDataprocGdcApplicationEnvironmentUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading ApplicationEnvironment: %s", err)
	}
	if err := d.Set("labels", flattenDataprocGdcApplicationEnvironmentLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading ApplicationEnvironment: %s", err)
	}
	if err := d.Set("annotations", flattenDataprocGdcApplicationEnvironmentAnnotations(res["annotations"], d, config)); err != nil {
		return fmt.Errorf("Error reading ApplicationEnvironment: %s", err)
	}
	if err := d.Set("spark_application_environment_config", flattenDataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfig(res["sparkApplicationEnvironmentConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading ApplicationEnvironment: %s", err)
	}
	if err := d.Set("namespace", flattenDataprocGdcApplicationEnvironmentNamespace(res["namespace"], d, config)); err != nil {
		return fmt.Errorf("Error reading ApplicationEnvironment: %s", err)
	}
	if err := d.Set("terraform_labels", flattenDataprocGdcApplicationEnvironmentTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading ApplicationEnvironment: %s", err)
	}
	if err := d.Set("effective_labels", flattenDataprocGdcApplicationEnvironmentEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading ApplicationEnvironment: %s", err)
	}
	if err := d.Set("effective_annotations", flattenDataprocGdcApplicationEnvironmentEffectiveAnnotations(res["annotations"], d, config)); err != nil {
		return fmt.Errorf("Error reading ApplicationEnvironment: %s", err)
	}

	return nil
}

func resourceDataprocGdcApplicationEnvironmentUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ApplicationEnvironment: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandDataprocGdcApplicationEnvironmentDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	sparkApplicationEnvironmentConfigProp, err := expandDataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfig(d.Get("spark_application_environment_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("spark_application_environment_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, sparkApplicationEnvironmentConfigProp)) {
		obj["sparkApplicationEnvironmentConfig"] = sparkApplicationEnvironmentConfigProp
	}
	namespaceProp, err := expandDataprocGdcApplicationEnvironmentNamespace(d.Get("namespace"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("namespace"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, namespaceProp)) {
		obj["namespace"] = namespaceProp
	}
	labelsProp, err := expandDataprocGdcApplicationEnvironmentEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	annotationsProp, err := expandDataprocGdcApplicationEnvironmentEffectiveAnnotations(d.Get("effective_annotations"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_annotations"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DataprocGdcBasePath}}projects/{{project}}/locations/{{location}}/serviceInstances/{{serviceinstance}}/applicationEnvironments/{{application_environment_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ApplicationEnvironment %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("spark_application_environment_config") {
		updateMask = append(updateMask, "sparkApplicationEnvironmentConfig")
	}

	if d.HasChange("namespace") {
		updateMask = append(updateMask, "namespace")
	}

	if d.HasChange("effective_labels") {
		updateMask = append(updateMask, "labels")
	}

	if d.HasChange("effective_annotations") {
		updateMask = append(updateMask, "annotations")
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
			return fmt.Errorf("Error updating ApplicationEnvironment %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating ApplicationEnvironment %q: %#v", d.Id(), res)
		}

	}

	return resourceDataprocGdcApplicationEnvironmentRead(d, meta)
}

func resourceDataprocGdcApplicationEnvironmentDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ApplicationEnvironment: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{DataprocGdcBasePath}}projects/{{project}}/locations/{{location}}/serviceInstances/{{serviceinstance}}/applicationEnvironments/{{application_environment_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting ApplicationEnvironment %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "ApplicationEnvironment")
	}

	log.Printf("[DEBUG] Finished deleting ApplicationEnvironment %q: %#v", d.Id(), res)
	return nil
}

func resourceDataprocGdcApplicationEnvironmentImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/serviceInstances/(?P<serviceinstance>[^/]+)/applicationEnvironments/(?P<application_environment_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<serviceinstance>[^/]+)/(?P<application_environment_id>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<serviceinstance>[^/]+)/(?P<application_environment_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/serviceInstances/{{serviceinstance}}/applicationEnvironments/{{application_environment_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenDataprocGdcApplicationEnvironmentName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocGdcApplicationEnvironmentUid(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocGdcApplicationEnvironmentDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocGdcApplicationEnvironmentCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocGdcApplicationEnvironmentUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocGdcApplicationEnvironmentLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenDataprocGdcApplicationEnvironmentAnnotations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("annotations"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenDataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["default_properties"] =
		flattenDataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigDefaultProperties(original["defaultProperties"], d, config)
	transformed["default_version"] =
		flattenDataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigDefaultVersion(original["defaultVersion"], d, config)
	return []interface{}{transformed}
}
func flattenDataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigDefaultProperties(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigDefaultVersion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocGdcApplicationEnvironmentNamespace(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocGdcApplicationEnvironmentTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenDataprocGdcApplicationEnvironmentEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocGdcApplicationEnvironmentEffectiveAnnotations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandDataprocGdcApplicationEnvironmentDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDefaultProperties, err := expandDataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigDefaultProperties(original["default_properties"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDefaultProperties); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["defaultProperties"] = transformedDefaultProperties
	}

	transformedDefaultVersion, err := expandDataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigDefaultVersion(original["default_version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDefaultVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["defaultVersion"] = transformedDefaultVersion
	}

	return transformed, nil
}

func expandDataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigDefaultProperties(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandDataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigDefaultVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocGdcApplicationEnvironmentNamespace(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocGdcApplicationEnvironmentEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandDataprocGdcApplicationEnvironmentEffectiveAnnotations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
