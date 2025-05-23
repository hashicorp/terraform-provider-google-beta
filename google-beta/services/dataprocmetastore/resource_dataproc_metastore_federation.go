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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/metastore/Federation.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package dataprocmetastore

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

func ResourceDataprocMetastoreFederation() *schema.Resource {
	return &schema.Resource{
		Create: resourceDataprocMetastoreFederationCreate,
		Read:   resourceDataprocMetastoreFederationRead,
		Update: resourceDataprocMetastoreFederationUpdate,
		Delete: resourceDataprocMetastoreFederationDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDataprocMetastoreFederationImport,
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
			"backend_metastores": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: `A map from BackendMetastore rank to BackendMetastores from which the federation service serves metadata at query time. The map key represents the order in which BackendMetastores should be evaluated to resolve database names at query time and should be greater than or equal to zero. A BackendMetastore with a lower number will be evaluated before a BackendMetastore with a higher number.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rank": {
							Type:     schema.TypeString,
							Required: true,
						},
						"metastore_type": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: verify.ValidateEnum([]string{"METASTORE_TYPE_UNSPECIFIED", "DATAPROC_METASTORE", "BIGQUERY"}),
							Description:  `The type of the backend metastore. Possible values: ["METASTORE_TYPE_UNSPECIFIED", "DATAPROC_METASTORE", "BIGQUERY"]`,
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The relative resource name of the metastore that is being federated. The formats of the relative resource names for the currently supported metastores are listed below: Dataplex: projects/{projectId}/locations/{location}/lakes/{lake_id} BigQuery: projects/{projectId} Dataproc Metastore: projects/{projectId}/locations/{location}/services/{serviceId}`,
						},
					},
				},
			},
			"federation_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The ID of the metastore federation. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
3 and 63 characters.`,
			},
			"version": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The Apache Hive metastore version of the federation. All backend metastore versions must be compatible with the federation version.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `User-defined labels for the metastore federation.

**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"location": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The location where the metastore federation should reside.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The time when the metastore federation was created.`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"endpoint_uri": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The URI of the endpoint used to access the metastore federation.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The relative resource name of the metastore federation.`,
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The current state of the metastore federation.`,
			},
			"state_message": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Additional information about the current state of the metastore federation, if available.`,
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
				Description: `The globally unique resource identifier of the metastore federation.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The time when the metastore federation was last updated.`,
			},
			"deletion_protection": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: `Whether Terraform will be prevented from destroying the federation. Defaults to false.
When the field is set to true in Terraform state, a 'terraform apply'
or 'terraform destroy' that would delete the federation will fail.`,
				Default: false,
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

func resourceDataprocMetastoreFederationCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	versionProp, err := expandDataprocMetastoreFederationVersion(d.Get("version"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("version"); !tpgresource.IsEmptyValue(reflect.ValueOf(versionProp)) && (ok || !reflect.DeepEqual(v, versionProp)) {
		obj["version"] = versionProp
	}
	backendMetastoresProp, err := expandDataprocMetastoreFederationBackendMetastores(d.Get("backend_metastores"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("backend_metastores"); !tpgresource.IsEmptyValue(reflect.ValueOf(backendMetastoresProp)) && (ok || !reflect.DeepEqual(v, backendMetastoresProp)) {
		obj["backendMetastores"] = backendMetastoresProp
	}
	labelsProp, err := expandDataprocMetastoreFederationEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DataprocMetastoreBasePath}}projects/{{project}}/locations/{{location}}/federations?federationId={{federation_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Federation: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Federation: %s", err)
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
		return fmt.Errorf("Error creating Federation: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/federations/{{federation_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = DataprocMetastoreOperationWaitTime(
		config, res, project, "Creating Federation", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Federation: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Federation %q: %#v", d.Id(), res)

	return resourceDataprocMetastoreFederationRead(d, meta)
}

func resourceDataprocMetastoreFederationRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DataprocMetastoreBasePath}}projects/{{project}}/locations/{{location}}/federations/{{federation_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Federation: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("DataprocMetastoreFederation %q", d.Id()))
	}

	// Explicitly set virtual fields to default values if unset
	if _, ok := d.GetOkExists("deletion_protection"); !ok {
		if err := d.Set("deletion_protection", false); err != nil {
			return fmt.Errorf("Error setting deletion_protection: %s", err)
		}
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Federation: %s", err)
	}

	if err := d.Set("name", flattenDataprocMetastoreFederationName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Federation: %s", err)
	}
	if err := d.Set("create_time", flattenDataprocMetastoreFederationCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Federation: %s", err)
	}
	if err := d.Set("update_time", flattenDataprocMetastoreFederationUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Federation: %s", err)
	}
	if err := d.Set("labels", flattenDataprocMetastoreFederationLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Federation: %s", err)
	}
	if err := d.Set("endpoint_uri", flattenDataprocMetastoreFederationEndpointUri(res["endpointUri"], d, config)); err != nil {
		return fmt.Errorf("Error reading Federation: %s", err)
	}
	if err := d.Set("state", flattenDataprocMetastoreFederationState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading Federation: %s", err)
	}
	if err := d.Set("state_message", flattenDataprocMetastoreFederationStateMessage(res["stateMessage"], d, config)); err != nil {
		return fmt.Errorf("Error reading Federation: %s", err)
	}
	if err := d.Set("uid", flattenDataprocMetastoreFederationUid(res["uid"], d, config)); err != nil {
		return fmt.Errorf("Error reading Federation: %s", err)
	}
	if err := d.Set("version", flattenDataprocMetastoreFederationVersion(res["version"], d, config)); err != nil {
		return fmt.Errorf("Error reading Federation: %s", err)
	}
	if err := d.Set("backend_metastores", flattenDataprocMetastoreFederationBackendMetastores(res["backendMetastores"], d, config)); err != nil {
		return fmt.Errorf("Error reading Federation: %s", err)
	}
	if err := d.Set("terraform_labels", flattenDataprocMetastoreFederationTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Federation: %s", err)
	}
	if err := d.Set("effective_labels", flattenDataprocMetastoreFederationEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Federation: %s", err)
	}

	return nil
}

func resourceDataprocMetastoreFederationUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Federation: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	backendMetastoresProp, err := expandDataprocMetastoreFederationBackendMetastores(d.Get("backend_metastores"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("backend_metastores"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, backendMetastoresProp)) {
		obj["backendMetastores"] = backendMetastoresProp
	}
	labelsProp, err := expandDataprocMetastoreFederationEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DataprocMetastoreBasePath}}projects/{{project}}/locations/{{location}}/federations/{{federation_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Federation %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("backend_metastores") {
		updateMask = append(updateMask, "backendMetastores")
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
			return fmt.Errorf("Error updating Federation %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating Federation %q: %#v", d.Id(), res)
		}

		err = DataprocMetastoreOperationWaitTime(
			config, res, project, "Updating Federation", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceDataprocMetastoreFederationRead(d, meta)
}

func resourceDataprocMetastoreFederationDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Federation: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{DataprocMetastoreBasePath}}projects/{{project}}/locations/{{location}}/federations/{{federation_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	if d.Get("deletion_protection").(bool) {
		return fmt.Errorf("cannot destroy metastore federation without setting deletion_protection=false and running `terraform apply`")
	}

	log.Printf("[DEBUG] Deleting Federation %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "Federation")
	}

	err = DataprocMetastoreOperationWaitTime(
		config, res, project, "Deleting Federation", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Federation %q: %#v", d.Id(), res)
	return nil
}

func resourceDataprocMetastoreFederationImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/federations/(?P<federation_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<federation_id>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<federation_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/federations/{{federation_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Explicitly set virtual fields to default values on import
	if err := d.Set("deletion_protection", false); err != nil {
		return nil, fmt.Errorf("Error setting deletion_protection: %s", err)
	}

	return []*schema.ResourceData{d}, nil
}

func flattenDataprocMetastoreFederationName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocMetastoreFederationCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocMetastoreFederationUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocMetastoreFederationLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenDataprocMetastoreFederationEndpointUri(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocMetastoreFederationState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocMetastoreFederationStateMessage(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocMetastoreFederationUid(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocMetastoreFederationVersion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocMetastoreFederationBackendMetastores(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.(map[string]interface{})
	transformed := make([]interface{}, 0, len(l))
	for k, raw := range l {
		original := raw.(map[string]interface{})
		transformed = append(transformed, map[string]interface{}{
			"rank":           k,
			"name":           flattenDataprocMetastoreFederationBackendMetastoresName(original["name"], d, config),
			"metastore_type": flattenDataprocMetastoreFederationBackendMetastoresMetastoreType(original["metastoreType"], d, config),
		})
	}
	return transformed
}
func flattenDataprocMetastoreFederationBackendMetastoresName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocMetastoreFederationBackendMetastoresMetastoreType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocMetastoreFederationTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenDataprocMetastoreFederationEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandDataprocMetastoreFederationVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocMetastoreFederationBackendMetastores(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	if v == nil {
		return map[string]interface{}{}, nil
	}
	m := make(map[string]interface{})
	for _, raw := range v.(*schema.Set).List() {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandDataprocMetastoreFederationBackendMetastoresName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedMetastoreType, err := expandDataprocMetastoreFederationBackendMetastoresMetastoreType(original["metastore_type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMetastoreType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["metastoreType"] = transformedMetastoreType
		}

		transformedRank, err := tpgresource.ExpandString(original["rank"], d, config)
		if err != nil {
			return nil, err
		}
		m[transformedRank] = transformed
	}
	return m, nil
}

func expandDataprocMetastoreFederationBackendMetastoresName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocMetastoreFederationBackendMetastoresMetastoreType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocMetastoreFederationEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
