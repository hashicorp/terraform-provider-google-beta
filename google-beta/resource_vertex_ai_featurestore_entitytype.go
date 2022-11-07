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

package google

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVertexAIFeaturestoreEntitytype() *schema.Resource {
	return &schema.Resource{
		Create: resourceVertexAIFeaturestoreEntitytypeCreate,
		Read:   resourceVertexAIFeaturestoreEntitytypeRead,
		Update: resourceVertexAIFeaturestoreEntitytypeUpdate,
		Delete: resourceVertexAIFeaturestoreEntitytypeDelete,

		Importer: &schema.ResourceImporter{
			State: resourceVertexAIFeaturestoreEntitytypeImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"featurestore": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}.`,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `A set of key/value label pairs to assign to this EntityType.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"monitoring_config": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `The default monitoring configuration for all Features under this EntityType.

If this is populated with [FeaturestoreMonitoringConfig.monitoring_interval] specified, snapshot analysis monitoring is enabled. Otherwise, snapshot analysis monitoring is disabled.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"snapshot_analysis": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Configuration of how features in Featurestore are monitored.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"disabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: `The monitoring schedule for snapshot analysis. For EntityType-level config: unset / disabled = true indicates disabled by default for Features under it; otherwise by default enable snapshot analysis monitoring with monitoringInterval for Features under it.`,
										Default:     false,
									},
									"monitoring_interval": {
										Type:       schema.TypeString,
										Optional:   true,
										Deprecated: "This field is unavailable in the GA provider and will be removed from the beta provider in a future release.",
										Description: `Configuration of the snapshot analysis based monitoring pipeline running interval. The value is rolled up to full day.

A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".`,
									},
								},
							},
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The name of the EntityType. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The timestamp of when the featurestore was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.`,
			},
			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Used to perform consistent read-modify-write updates.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The timestamp of when the featurestore was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.`,
			},
			"region": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The region of the EntityType.",
			},
		},
		UseJSONNumber: true,
	}
}

func resourceVertexAIFeaturestoreEntitytypeCreate(d *schema.ResourceData, meta interface{}) error {
	var project string
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	labelsProp, err := expandVertexAIFeaturestoreEntitytypeLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	monitoringConfigProp, err := expandVertexAIFeaturestoreEntitytypeMonitoringConfig(d.Get("monitoring_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("monitoring_config"); !isEmptyValue(reflect.ValueOf(monitoringConfigProp)) && (ok || !reflect.DeepEqual(v, monitoringConfigProp)) {
		obj["monitoringConfig"] = monitoringConfigProp
	}

	obj, err = resourceVertexAIFeaturestoreEntitytypeEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{VertexAIBasePath}}{{featurestore}}/entityTypes?entityTypeId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new FeaturestoreEntitytype: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	if v, ok := d.GetOk("featurestore"); ok {
		re := regexp.MustCompile("projects/([a-zA-Z0-9-]*)/(?:locations|regions)/([a-zA-Z0-9-]*)")
		switch {
		case re.MatchString(v.(string)):
			if res := re.FindStringSubmatch(v.(string)); len(res) == 3 && res[1] != "" {
				project = res[1]
			}
		}
	}
	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating FeaturestoreEntitytype: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{featurestore}}/entityTypes/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = vertexAIOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating FeaturestoreEntitytype", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create FeaturestoreEntitytype: %s", err)
	}

	// This may have caused the ID to update - update it if so.
	id, err = replaceVars(d, config, "{{featurestore}}/entityTypes/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating FeaturestoreEntitytype %q: %#v", d.Id(), res)

	return resourceVertexAIFeaturestoreEntitytypeRead(d, meta)
}

func resourceVertexAIFeaturestoreEntitytypeRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{VertexAIBasePath}}{{featurestore}}/entityTypes/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("VertexAIFeaturestoreEntitytype %q", d.Id()))
	}

	if err := d.Set("create_time", flattenVertexAIFeaturestoreEntitytypeCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeaturestoreEntitytype: %s", err)
	}
	if err := d.Set("update_time", flattenVertexAIFeaturestoreEntitytypeUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeaturestoreEntitytype: %s", err)
	}
	if err := d.Set("labels", flattenVertexAIFeaturestoreEntitytypeLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeaturestoreEntitytype: %s", err)
	}
	if err := d.Set("monitoring_config", flattenVertexAIFeaturestoreEntitytypeMonitoringConfig(res["monitoringConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeaturestoreEntitytype: %s", err)
	}

	return nil
}

func resourceVertexAIFeaturestoreEntitytypeUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	labelsProp, err := expandVertexAIFeaturestoreEntitytypeLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	monitoringConfigProp, err := expandVertexAIFeaturestoreEntitytypeMonitoringConfig(d.Get("monitoring_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("monitoring_config"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, monitoringConfigProp)) {
		obj["monitoringConfig"] = monitoringConfigProp
	}

	obj, err = resourceVertexAIFeaturestoreEntitytypeEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{VertexAIBasePath}}{{featurestore}}/entityTypes/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating FeaturestoreEntitytype %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}

	if d.HasChange("monitoring_config") {
		updateMask = append(updateMask, "monitoringConfig")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating FeaturestoreEntitytype %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating FeaturestoreEntitytype %q: %#v", d.Id(), res)
	}

	return resourceVertexAIFeaturestoreEntitytypeRead(d, meta)
}

func resourceVertexAIFeaturestoreEntitytypeDelete(d *schema.ResourceData, meta interface{}) error {
	var project string
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := replaceVars(d, config, "{{VertexAIBasePath}}{{featurestore}}/entityTypes/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	if v, ok := d.GetOk("featurestore"); ok {
		re := regexp.MustCompile("projects/([a-zA-Z0-9-]*)/(?:locations|regions)/([a-zA-Z0-9-]*)")
		switch {
		case re.MatchString(v.(string)):
			if res := re.FindStringSubmatch(v.(string)); len(res) == 3 && res[1] != "" {
				project = res[1]
			}
		}
	}
	log.Printf("[DEBUG] Deleting FeaturestoreEntitytype %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "FeaturestoreEntitytype")
	}

	err = vertexAIOperationWaitTime(
		config, res, project, "Deleting FeaturestoreEntitytype", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting FeaturestoreEntitytype %q: %#v", d.Id(), res)
	return nil
}

func resourceVertexAIFeaturestoreEntitytypeImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"(?P<featurestore>.+)/entityTypes/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{featurestore}}/entityTypes/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	featurestore := d.Get("featurestore").(string)

	re := regexp.MustCompile("^projects/(.+)/locations/(.+)/featurestores/(.+)$")
	if parts := re.FindStringSubmatch(featurestore); parts != nil {
		d.Set("region", parts[2])
	}

	return []*schema.ResourceData{d}, nil
}

func flattenVertexAIFeaturestoreEntitytypeCreateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenVertexAIFeaturestoreEntitytypeUpdateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenVertexAIFeaturestoreEntitytypeLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenVertexAIFeaturestoreEntitytypeMonitoringConfig(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["snapshot_analysis"] =
		flattenVertexAIFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysis(original["snapshotAnalysis"], d, config)
	return []interface{}{transformed}
}
func flattenVertexAIFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysis(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["disabled"] =
		flattenVertexAIFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisDisabled(original["disabled"], d, config)
	transformed["monitoring_interval"] =
		flattenVertexAIFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisMonitoringInterval(original["monitoringInterval"], d, config)
	return []interface{}{transformed}
}
func flattenVertexAIFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisDisabled(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenVertexAIFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisMonitoringInterval(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandVertexAIFeaturestoreEntitytypeLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandVertexAIFeaturestoreEntitytypeMonitoringConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSnapshotAnalysis, err := expandVertexAIFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysis(original["snapshot_analysis"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSnapshotAnalysis); val.IsValid() && !isEmptyValue(val) {
		transformed["snapshotAnalysis"] = transformedSnapshotAnalysis
	}

	return transformed, nil
}

func expandVertexAIFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysis(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDisabled, err := expandVertexAIFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisDisabled(original["disabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDisabled); val.IsValid() && !isEmptyValue(val) {
		transformed["disabled"] = transformedDisabled
	}

	transformedMonitoringInterval, err := expandVertexAIFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisMonitoringInterval(original["monitoring_interval"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMonitoringInterval); val.IsValid() && !isEmptyValue(val) {
		transformed["monitoringInterval"] = transformedMonitoringInterval
	}

	return transformed, nil
}

func expandVertexAIFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisDisabled(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisMonitoringInterval(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourceVertexAIFeaturestoreEntitytypeEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	if v, ok := d.GetOk("featurestore"); ok {
		re := regexp.MustCompile("projects/(.+)/locations/(.+)/featurestores/(.+)$")
		if parts := re.FindStringSubmatch(v.(string)); parts != nil {
			d.Set("region", parts[2])
		}
	}

	return obj, nil
}
