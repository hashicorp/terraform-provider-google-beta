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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/securesourcemanager/Instance.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package securesourcemanager

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

func ResourceSecureSourceManagerInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecureSourceManagerInstanceCreate,
		Read:   resourceSecureSourceManagerInstanceRead,
		Update: resourceSecureSourceManagerInstanceUpdate,
		Delete: resourceSecureSourceManagerInstanceDelete,

		Importer: &schema.ResourceImporter{
			State: resourceSecureSourceManagerInstanceImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name for the Instance.`,
			},
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The location for the Instance.`,
			},
			"kms_key": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `Customer-managed encryption key name, in the format projects/*/locations/*/keyRings/*/cryptoKeys/*.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Labels as key value pairs.


**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"private_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `Private settings for private instance.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ca_pool": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: `CA pool resource, resource must in the format of 'projects/{project}/locations/{location}/caPools/{ca_pool}'.`,
						},
						"is_private": {
							Type:        schema.TypeBool,
							Required:    true,
							ForceNew:    true,
							Description: `'Indicate if it's private instance.'`,
						},
						"http_service_attachment": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Service Attachment for HTTP, resource is in the format of 'projects/{project}/regions/{region}/serviceAttachments/{service_attachment}'.`,
						},
						"ssh_service_attachment": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Service Attachment for SSH, resource is in the format of 'projects/{project}/regions/{region}/serviceAttachments/{service_attachment}'.`,
						},
					},
				},
			},
			"workforce_identity_federation_config": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Description: `Configuration for Workforce Identity Federation to support third party identity provider.
If unset, defaults to the Google OIDC IdP.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Required:    true,
							ForceNew:    true,
							Description: `'Whether Workforce Identity Federation is enabled.'`,
						},
					},
				},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the Instance was created in UTC.`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				ForceNew:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"host_config": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `A list of hostnames for this instance.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"api": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `API hostname.`,
						},
						"git_http": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Git HTTP hostname.`,
						},
						"git_ssh": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Git SSH hostname.`,
						},
						"html": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `HTML hostname.`,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The resource name for the Instance.`,
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The current state of the Instance.`,
			},
			"state_note": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Provides information about the current instance state.`,
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the Instance was updated in UTC.`,
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

func resourceSecureSourceManagerInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	kmsKeyProp, err := expandSecureSourceManagerInstanceKmsKey(d.Get("kms_key"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("kms_key"); !tpgresource.IsEmptyValue(reflect.ValueOf(kmsKeyProp)) && (ok || !reflect.DeepEqual(v, kmsKeyProp)) {
		obj["kmsKey"] = kmsKeyProp
	}
	privateConfigProp, err := expandSecureSourceManagerInstancePrivateConfig(d.Get("private_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("private_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(privateConfigProp)) && (ok || !reflect.DeepEqual(v, privateConfigProp)) {
		obj["privateConfig"] = privateConfigProp
	}
	workforceIdentityFederationConfigProp, err := expandSecureSourceManagerInstanceWorkforceIdentityFederationConfig(d.Get("workforce_identity_federation_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("workforce_identity_federation_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(workforceIdentityFederationConfigProp)) && (ok || !reflect.DeepEqual(v, workforceIdentityFederationConfigProp)) {
		obj["workforceIdentityFederationConfig"] = workforceIdentityFederationConfigProp
	}
	labelsProp, err := expandSecureSourceManagerInstanceEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecureSourceManagerBasePath}}projects/{{project}}/locations/{{location}}/instances?instance_id={{instance_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Instance: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Instance: %s", err)
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
		return fmt.Errorf("Error creating Instance: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/instances/{{instance_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = SecureSourceManagerOperationWaitTime(
		config, res, project, "Creating Instance", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Instance: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Instance %q: %#v", d.Id(), res)

	return resourceSecureSourceManagerInstanceRead(d, meta)
}

func resourceSecureSourceManagerInstanceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecureSourceManagerBasePath}}projects/{{project}}/locations/{{location}}/instances/{{instance_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Instance: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("SecureSourceManagerInstance %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}

	if err := d.Set("name", flattenSecureSourceManagerInstanceName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("create_time", flattenSecureSourceManagerInstanceCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("update_time", flattenSecureSourceManagerInstanceUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("labels", flattenSecureSourceManagerInstanceLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("state", flattenSecureSourceManagerInstanceState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("host_config", flattenSecureSourceManagerInstanceHostConfig(res["hostConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("state_note", flattenSecureSourceManagerInstanceStateNote(res["stateNote"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("kms_key", flattenSecureSourceManagerInstanceKmsKey(res["kmsKey"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("private_config", flattenSecureSourceManagerInstancePrivateConfig(res["privateConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("workforce_identity_federation_config", flattenSecureSourceManagerInstanceWorkforceIdentityFederationConfig(res["workforceIdentityFederationConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("terraform_labels", flattenSecureSourceManagerInstanceTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("effective_labels", flattenSecureSourceManagerInstanceEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}

	return nil
}

func resourceSecureSourceManagerInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	// Only the root field "labels" and "terraform_labels" are mutable
	return resourceSecureSourceManagerInstanceRead(d, meta)
}

func resourceSecureSourceManagerInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Instance: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{SecureSourceManagerBasePath}}projects/{{project}}/locations/{{location}}/instances/{{instance_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting Instance %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "Instance")
	}

	err = SecureSourceManagerOperationWaitTime(
		config, res, project, "Deleting Instance", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Instance %q: %#v", d.Id(), res)
	return nil
}

func resourceSecureSourceManagerInstanceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/instances/(?P<instance_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<instance_id>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<instance_id>[^/]+)$",
		"^(?P<instance_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/instances/{{instance_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenSecureSourceManagerInstanceName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecureSourceManagerInstanceCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecureSourceManagerInstanceUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecureSourceManagerInstanceLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenSecureSourceManagerInstanceState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecureSourceManagerInstanceHostConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["html"] =
		flattenSecureSourceManagerInstanceHostConfigHtml(original["html"], d, config)
	transformed["api"] =
		flattenSecureSourceManagerInstanceHostConfigApi(original["api"], d, config)
	transformed["git_http"] =
		flattenSecureSourceManagerInstanceHostConfigGitHttp(original["gitHttp"], d, config)
	transformed["git_ssh"] =
		flattenSecureSourceManagerInstanceHostConfigGitSsh(original["gitSsh"], d, config)
	return []interface{}{transformed}
}
func flattenSecureSourceManagerInstanceHostConfigHtml(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecureSourceManagerInstanceHostConfigApi(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecureSourceManagerInstanceHostConfigGitHttp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecureSourceManagerInstanceHostConfigGitSsh(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecureSourceManagerInstanceStateNote(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecureSourceManagerInstanceKmsKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecureSourceManagerInstancePrivateConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["is_private"] =
		flattenSecureSourceManagerInstancePrivateConfigIsPrivate(original["isPrivate"], d, config)
	transformed["ca_pool"] =
		flattenSecureSourceManagerInstancePrivateConfigCaPool(original["caPool"], d, config)
	transformed["http_service_attachment"] =
		flattenSecureSourceManagerInstancePrivateConfigHttpServiceAttachment(original["httpServiceAttachment"], d, config)
	transformed["ssh_service_attachment"] =
		flattenSecureSourceManagerInstancePrivateConfigSshServiceAttachment(original["sshServiceAttachment"], d, config)
	return []interface{}{transformed}
}
func flattenSecureSourceManagerInstancePrivateConfigIsPrivate(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecureSourceManagerInstancePrivateConfigCaPool(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecureSourceManagerInstancePrivateConfigHttpServiceAttachment(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecureSourceManagerInstancePrivateConfigSshServiceAttachment(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecureSourceManagerInstanceWorkforceIdentityFederationConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["enabled"] =
		flattenSecureSourceManagerInstanceWorkforceIdentityFederationConfigEnabled(original["enabled"], d, config)
	return []interface{}{transformed}
}
func flattenSecureSourceManagerInstanceWorkforceIdentityFederationConfigEnabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecureSourceManagerInstanceTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenSecureSourceManagerInstanceEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandSecureSourceManagerInstanceKmsKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecureSourceManagerInstancePrivateConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIsPrivate, err := expandSecureSourceManagerInstancePrivateConfigIsPrivate(original["is_private"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIsPrivate); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["isPrivate"] = transformedIsPrivate
	}

	transformedCaPool, err := expandSecureSourceManagerInstancePrivateConfigCaPool(original["ca_pool"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCaPool); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["caPool"] = transformedCaPool
	}

	transformedHttpServiceAttachment, err := expandSecureSourceManagerInstancePrivateConfigHttpServiceAttachment(original["http_service_attachment"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHttpServiceAttachment); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["httpServiceAttachment"] = transformedHttpServiceAttachment
	}

	transformedSshServiceAttachment, err := expandSecureSourceManagerInstancePrivateConfigSshServiceAttachment(original["ssh_service_attachment"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSshServiceAttachment); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["sshServiceAttachment"] = transformedSshServiceAttachment
	}

	return transformed, nil
}

func expandSecureSourceManagerInstancePrivateConfigIsPrivate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecureSourceManagerInstancePrivateConfigCaPool(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecureSourceManagerInstancePrivateConfigHttpServiceAttachment(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecureSourceManagerInstancePrivateConfigSshServiceAttachment(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecureSourceManagerInstanceWorkforceIdentityFederationConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnabled, err := expandSecureSourceManagerInstanceWorkforceIdentityFederationConfigEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enabled"] = transformedEnabled
	}

	return transformed, nil
}

func expandSecureSourceManagerInstanceWorkforceIdentityFederationConfigEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecureSourceManagerInstanceEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
