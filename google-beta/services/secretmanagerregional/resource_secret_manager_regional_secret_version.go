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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/secretmanagerregional/RegionalSecretVersion.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package secretmanagerregional

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"

	"google.golang.org/api/googleapi"
)

func ResourceSecretManagerRegionalRegionalSecretVersion() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecretManagerRegionalRegionalSecretVersionCreate,
		Read:   resourceSecretManagerRegionalRegionalSecretVersionRead,
		Update: resourceSecretManagerRegionalRegionalSecretVersionUpdate,
		Delete: resourceSecretManagerRegionalRegionalSecretVersionDelete,

		Importer: &schema.ResourceImporter{
			State: resourceSecretManagerRegionalRegionalSecretVersionImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"secret_data": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The secret data. Must be no larger than 64KiB.`,
				Sensitive:   true,
			},

			"secret": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `Secret Manager regional secret resource.`,
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `The current state of the regional secret version.`,
				Default:     true,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time at which the regional secret version was created.`,
			},
			"customer_managed_encryption": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `The customer-managed encryption configuration of the regional secret.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kms_key_version_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The resource name of the Cloud KMS CryptoKey used to encrypt secret payloads.`,
						},
					},
				},
			},
			"destroy_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time at which the regional secret version was destroyed. Only present if state is DESTROYED.`,
			},
			"location": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Location of Secret Manager regional secret resource.`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The resource name of the regional secret version. Format:
'projects/{{project}}/locations/{{location}}/secrets/{{secret_id}}/versions/{{version}}'`,
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The version of the Regional Secret.`,
			},
			"deletion_policy": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `The deletion policy for the regional secret version. Setting 'ABANDON' allows the resource
to be abandoned rather than deleted. Setting 'DISABLE' allows the resource to be
disabled rather than deleted. Default is 'DELETE'. Possible values are:
  * DELETE
  * DISABLE
  * ABANDON`,
				Default: "DELETE",
			},
			"is_secret_data_base64": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     false,
				Description: `If set to 'true', the secret data is expected to be base64-encoded string and would be sent as is.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceSecretManagerRegionalRegionalSecretVersionCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	stateProp, err := expandSecretManagerRegionalRegionalSecretVersionEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(stateProp)) && (ok || !reflect.DeepEqual(v, stateProp)) {
		obj["state"] = stateProp
	}
	payloadProp, err := expandSecretManagerRegionalRegionalSecretVersionPayload(nil, d, config)
	if err != nil {
		return err
	} else if !tpgresource.IsEmptyValue(reflect.ValueOf(payloadProp)) {
		obj["payload"] = payloadProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecretManagerRegionalBasePath}}{{secret}}:addVersion")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new RegionalSecretVersion: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	secret := d.Get("secret").(string)
	secretRegex := regexp.MustCompile("projects/(.+)/locations/(.+)/secrets/(.+)$")

	parts := secretRegex.FindStringSubmatch(secret)
	if len(parts) != 4 {
		return fmt.Errorf("secret does not fit the format `projects/{{project}}/locations/{{location}}/secrets/{{secret}}`")
	}

	if err := d.Set("location", parts[2]); err != nil {
		return fmt.Errorf("Error setting location: %s", err)
	}

	// Override the url after setting the location
	url, err = tpgresource.ReplaceVars(d, config, "{{SecretManagerRegionalBasePath}}{{secret}}:addVersion")
	if err != nil {
		return err
	}
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
		return fmt.Errorf("Error creating RegionalSecretVersion: %s", err)
	}
	// Set computed resource properties from create API response so that they're available on the subsequent Read
	// call.
	res, err = resourceSecretManagerRegionalRegionalSecretVersionDecoder(d, meta, res)
	if err != nil {
		return fmt.Errorf("decoding response: %w", err)
	}
	if res == nil {
		return fmt.Errorf("decoding response, could not find object")
	}
	if err := d.Set("name", flattenSecretManagerRegionalRegionalSecretVersionName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// `name` is autogenerated from the api so needs to be set post-create
	name, ok := res["name"]
	if !ok {
		return fmt.Errorf("Create response didn't contain critical fields. Create may not have succeeded.")
	}
	if err := d.Set("name", name.(string)); err != nil {
		return fmt.Errorf("Error setting name: %s", err)
	}
	d.SetId(name.(string))

	_, err = expandSecretManagerRegionalRegionalSecretVersionEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished creating RegionalSecretVersion %q: %#v", d.Id(), res)

	return resourceSecretManagerRegionalRegionalSecretVersionRead(d, meta)
}

func resourceSecretManagerRegionalRegionalSecretVersionRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecretManagerRegionalBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	secret := d.Get("secret").(string)
	secretRegex := regexp.MustCompile("projects/(.+)/locations/(.+)/secrets/(.+)$")

	parts := secretRegex.FindStringSubmatch(secret)
	if len(parts) != 4 {
		return fmt.Errorf("secret does not fit the format `projects/{{project}}/locations/{{location}}/secrets/{{secret}}`")
	}

	if err := d.Set("location", parts[2]); err != nil {
		return fmt.Errorf("Error setting location: %s", err)
	}

	// Override the url after setting the location
	url, err = tpgresource.ReplaceVars(d, config, "{{SecretManagerRegionalBasePath}}{{name}}")
	if err != nil {
		return err
	}

	// Explicitly set the field to default value if unset
	if _, ok := d.GetOkExists("is_secret_data_base64"); !ok {
		if err := d.Set("is_secret_data_base64", false); err != nil {
			return fmt.Errorf("Error setting is_secret_data_base64: %s", err)
		}
	}
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("SecretManagerRegionalRegionalSecretVersion %q", d.Id()))
	}

	res, err = resourceSecretManagerRegionalRegionalSecretVersionDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing SecretManagerRegionalRegionalSecretVersion because it no longer exists.")
		d.SetId("")
		return nil
	}

	// Explicitly set virtual fields to default values if unset
	if _, ok := d.GetOkExists("deletion_policy"); !ok {
		if err := d.Set("deletion_policy", "DELETE"); err != nil {
			return fmt.Errorf("Error setting deletion_policy: %s", err)
		}
	}

	if err := d.Set("name", flattenSecretManagerRegionalRegionalSecretVersionName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionalSecretVersion: %s", err)
	}
	if err := d.Set("create_time", flattenSecretManagerRegionalRegionalSecretVersionCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionalSecretVersion: %s", err)
	}
	if err := d.Set("destroy_time", flattenSecretManagerRegionalRegionalSecretVersionDestroyTime(res["destroyTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionalSecretVersion: %s", err)
	}
	if err := d.Set("customer_managed_encryption", flattenSecretManagerRegionalRegionalSecretVersionCustomerManagedEncryption(res["customerManagedEncryption"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionalSecretVersion: %s", err)
	}
	if err := d.Set("version", flattenSecretManagerRegionalRegionalSecretVersionVersion(res["version"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionalSecretVersion: %s", err)
	}
	if err := d.Set("enabled", flattenSecretManagerRegionalRegionalSecretVersionEnabled(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionalSecretVersion: %s", err)
	}
	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if flattenedProp := flattenSecretManagerRegionalRegionalSecretVersionPayload(res["payload"], d, config); flattenedProp != nil {
		if gerr, ok := flattenedProp.(*googleapi.Error); ok {
			return fmt.Errorf("Error reading RegionalSecretVersion: %s", gerr)
		}
		casted := flattenedProp.([]interface{})[0]
		if casted != nil {
			for k, v := range casted.(map[string]interface{}) {
				if err := d.Set(k, v); err != nil {
					return fmt.Errorf("Error setting %s: %s", k, err)
				}
			}
		}
	}

	return nil
}

func resourceSecretManagerRegionalRegionalSecretVersionUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	_, err := expandSecretManagerRegionalRegionalSecretVersionEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return err
	}

	return resourceSecretManagerRegionalRegionalSecretVersionRead(d, meta)
}

func resourceSecretManagerRegionalRegionalSecretVersionDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{SecretManagerRegionalBasePath}}{{name}}:destroy")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	deletionPolicy := d.Get("deletion_policy")

	if deletionPolicy == "ABANDON" {
		return nil
	} else if deletionPolicy == "DISABLE" {
		url, err = tpgresource.ReplaceVars(d, config, "{{SecretManagerRegionalBasePath}}{{name}}:disable")
		if err != nil {
			return err
		}
	}

	log.Printf("[DEBUG] Deleting RegionalSecretVersion %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "RegionalSecretVersion")
	}

	log.Printf("[DEBUG] Finished deleting RegionalSecretVersion %q: %#v", d.Id(), res)
	return nil
}

func resourceSecretManagerRegionalRegionalSecretVersionImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := tpgresource.ParseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}

	name := d.Get("name").(string)
	secretRegex := regexp.MustCompile("(projects/.+/locations/.+/secrets/.+)/versions/.+$")
	versionRegex := regexp.MustCompile("projects/(.+)/locations/(.+)/secrets/(.+)/versions/(.+)$")

	parts := secretRegex.FindStringSubmatch(name)
	if len(parts) != 2 {
		return nil, fmt.Errorf("Version name does not fit the format `projects/{{project}}/locations/{{location}}/secrets/{{secret}}/versions/{{version}}`")
	}
	if err := d.Set("secret", parts[1]); err != nil {
		return nil, fmt.Errorf("Error setting secret: %s", err)
	}

	parts = versionRegex.FindStringSubmatch(name)

	if err := d.Set("version", parts[4]); err != nil {
		return nil, fmt.Errorf("Error setting version: %s", err)
	}

	// Explicitly set virtual fields to default values on import
	if err := d.Set("deletion_policy", "DELETE"); err != nil {
		return nil, fmt.Errorf("Error setting deletion policy: %s", err)
	}

	if err := d.Set("location", parts[2]); err != nil {
		return nil, fmt.Errorf("Error setting location: %s", err)
	}

	return []*schema.ResourceData{d}, nil
}

func flattenSecretManagerRegionalRegionalSecretVersionName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecretManagerRegionalRegionalSecretVersionCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecretManagerRegionalRegionalSecretVersionDestroyTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecretManagerRegionalRegionalSecretVersionCustomerManagedEncryption(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["kms_key_version_name"] =
		flattenSecretManagerRegionalRegionalSecretVersionCustomerManagedEncryptionKmsKeyVersionName(original["kmsKeyVersionName"], d, config)
	return []interface{}{transformed}
}
func flattenSecretManagerRegionalRegionalSecretVersionCustomerManagedEncryptionKmsKeyVersionName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecretManagerRegionalRegionalSecretVersionVersion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	name := d.Get("name").(string)
	secretRegex := regexp.MustCompile("projects/(.+)/locations/(.+)/secrets/(.+)/versions/(.+)$")

	parts := secretRegex.FindStringSubmatch(name)
	if len(parts) != 5 {
		return fmt.Errorf("Version name does not fit the format `projects/{{project}}/locations/{{location}}/secrets/{{secret}}/versions/{{version}}`")
	}

	return parts[4]
}

func flattenSecretManagerRegionalRegionalSecretVersionEnabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v.(string) == "ENABLED" {
		return true
	}

	return false
}

func flattenSecretManagerRegionalRegionalSecretVersionPayload(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	transformed := make(map[string]interface{})

	// if this secret version is disabled, the api will return an error, as the value cannot be accessed, return what we have
	if d.Get("enabled").(bool) == false {
		transformed["secret_data"] = d.Get("secret_data")
		return []interface{}{transformed}
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecretManagerRegionalBasePath}}{{name}}:access")
	if err != nil {
		return err
	}

	parts := strings.Split(d.Get("name").(string), "/")
	project := parts[1]

	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	accessRes, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   project,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return err
	}

	if d.Get("is_secret_data_base64").(bool) {
		transformed["secret_data"] = accessRes["payload"].(map[string]interface{})["data"].(string)
	} else {
		data, err := base64.StdEncoding.DecodeString(accessRes["payload"].(map[string]interface{})["data"].(string))
		if err != nil {
			return err
		}
		transformed["secret_data"] = string(data)
	}
	return []interface{}{transformed}
}

func expandSecretManagerRegionalRegionalSecretVersionEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	name := d.Get("name").(string)
	if name == "" {
		return "", nil
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecretManagerRegionalBasePath}}{{name}}")
	if err != nil {
		return nil, err
	}

	if v == true {
		url = fmt.Sprintf("%s:enable", url)
	} else {
		url = fmt.Sprintf("%s:disable", url)
	}

	parts := strings.Split(name, "/")
	project := parts[1]

	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return nil, err
	}

	_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   project,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func expandSecretManagerRegionalRegionalSecretVersionPayload(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	transformed := make(map[string]interface{})
	transformedSecretData, err := expandSecretManagerRegionalRegionalSecretVersionPayloadSecretData(d.Get("secret_data"), d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSecretData); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["data"] = transformedSecretData
	}

	return transformed, nil
}

func expandSecretManagerRegionalRegionalSecretVersionPayloadSecretData(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	if v == nil {
		return nil, nil
	}

	if d.Get("is_secret_data_base64").(bool) {
		return v, nil
	}
	return base64.StdEncoding.EncodeToString([]byte(v.(string))), nil
}

func resourceSecretManagerRegionalRegionalSecretVersionDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	if v := res["state"]; v == "DESTROYED" {
		return nil, nil
	}

	return res, nil
}
