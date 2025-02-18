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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/kms/KeyRingImportJob.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package kms

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceKMSKeyRingImportJob() *schema.Resource {
	return &schema.Resource{
		Create: resourceKMSKeyRingImportJobCreate,
		Read:   resourceKMSKeyRingImportJobRead,
		Delete: resourceKMSKeyRingImportJobDelete,

		Importer: &schema.ResourceImporter{
			State: resourceKMSKeyRingImportJobImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"import_job_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `It must be unique within a KeyRing and match the regular expression [a-zA-Z0-9_-]{1,63}`,
			},
			"import_method": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"RSA_OAEP_3072_SHA1_AES_256", "RSA_OAEP_4096_SHA1_AES_256"}),
				Description:  `The wrapping method to be used for incoming key material. Possible values: ["RSA_OAEP_3072_SHA1_AES_256", "RSA_OAEP_4096_SHA1_AES_256"]`,
			},
			"key_ring": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: kmsCryptoKeyRingsEquivalent,
				Description: `The KeyRing that this import job belongs to.
Format: ''projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}''.`,
			},
			"protection_level": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"SOFTWARE", "HSM", "EXTERNAL"}),
				Description: `The protection level of the ImportJob. This must match the protectionLevel of the
versionTemplate on the CryptoKey you attempt to import into. Possible values: ["SOFTWARE", "HSM", "EXTERNAL"]`,
			},
			"attestation": {
				Type:     schema.TypeList,
				Computed: true,
				Description: `Statement that was generated and signed by the key creator (for example, an HSM) at key creation time.
Use this statement to verify attributes of the key as stored on the HSM, independently of Google.
Only present if the chosen ImportMethod is one with a protection level of HSM.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"content": {
							Type:     schema.TypeString,
							Computed: true,
							Description: `The attestation data provided by the HSM when the key operation was performed.
A base64-encoded string.`,
						},
						"format": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The format of the attestation data.`,
						},
					},
				},
			},
			"expire_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The time at which this resource is scheduled for expiration and can no longer be used.
This is in RFC3339 text format.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The resource name for this ImportJob in the format projects/*/locations/*/keyRings/*/importJobs/*.`,
			},
			"public_key": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `The public key with which to wrap key material prior to import. Only returned if state is 'ACTIVE'.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pem": {
							Type:     schema.TypeString,
							Computed: true,
							Description: `The public key, encoded in PEM format. For more information, see the RFC 7468 sections
for General Considerations and Textual Encoding of Subject Public Key Info.`,
						},
					},
				},
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The current state of the ImportJob, indicating if it can be used.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceKMSKeyRingImportJobCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	importMethodProp, err := expandKMSKeyRingImportJobImportMethod(d.Get("import_method"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("import_method"); !tpgresource.IsEmptyValue(reflect.ValueOf(importMethodProp)) && (ok || !reflect.DeepEqual(v, importMethodProp)) {
		obj["importMethod"] = importMethodProp
	}
	protectionLevelProp, err := expandKMSKeyRingImportJobProtectionLevel(d.Get("protection_level"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("protection_level"); !tpgresource.IsEmptyValue(reflect.ValueOf(protectionLevelProp)) && (ok || !reflect.DeepEqual(v, protectionLevelProp)) {
		obj["protectionLevel"] = protectionLevelProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{KMSBasePath}}{{key_ring}}/importJobs?importJobId={{import_job_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new KeyRingImportJob: %#v", obj)
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
		return fmt.Errorf("Error creating KeyRingImportJob: %s", err)
	}
	if err := d.Set("name", flattenKMSKeyRingImportJobName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating KeyRingImportJob %q: %#v", d.Id(), res)

	return resourceKMSKeyRingImportJobRead(d, meta)
}

func resourceKMSKeyRingImportJobRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{KMSBasePath}}{{name}}")
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("KMSKeyRingImportJob %q", d.Id()))
	}

	if err := d.Set("name", flattenKMSKeyRingImportJobName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading KeyRingImportJob: %s", err)
	}
	if err := d.Set("import_method", flattenKMSKeyRingImportJobImportMethod(res["importMethod"], d, config)); err != nil {
		return fmt.Errorf("Error reading KeyRingImportJob: %s", err)
	}
	if err := d.Set("protection_level", flattenKMSKeyRingImportJobProtectionLevel(res["protectionLevel"], d, config)); err != nil {
		return fmt.Errorf("Error reading KeyRingImportJob: %s", err)
	}
	if err := d.Set("expire_time", flattenKMSKeyRingImportJobExpireTime(res["expireTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading KeyRingImportJob: %s", err)
	}
	if err := d.Set("state", flattenKMSKeyRingImportJobState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading KeyRingImportJob: %s", err)
	}
	if err := d.Set("public_key", flattenKMSKeyRingImportJobPublicKey(res["publicKey"], d, config)); err != nil {
		return fmt.Errorf("Error reading KeyRingImportJob: %s", err)
	}
	if err := d.Set("attestation", flattenKMSKeyRingImportJobAttestation(res["attestation"], d, config)); err != nil {
		return fmt.Errorf("Error reading KeyRingImportJob: %s", err)
	}

	return nil
}

func resourceKMSKeyRingImportJobDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{KMSBasePath}}{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting KeyRingImportJob %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "KeyRingImportJob")
	}

	log.Printf("[DEBUG] Finished deleting KeyRingImportJob %q: %#v", d.Id(), res)
	return nil
}

func resourceKMSKeyRingImportJobImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	config := meta.(*transport_tpg.Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := tpgresource.ParseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}

	stringParts := strings.Split(d.Get("name").(string), "/")
	if len(stringParts) != 8 {
		return nil, fmt.Errorf(
			"Saw %s when the name is expected to have shape %s",
			d.Get("name"),
			"projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/importJobs/{{importJobId}}",
		)
	}

	if err := d.Set("key_ring", stringParts[3]); err != nil {
		return nil, fmt.Errorf("Error setting key_ring: %s", err)
	}
	if err := d.Set("import_job_id", stringParts[5]); err != nil {
		return nil, fmt.Errorf("Error setting import_job_id: %s", err)
	}
	return []*schema.ResourceData{d}, nil
}

func flattenKMSKeyRingImportJobName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenKMSKeyRingImportJobImportMethod(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenKMSKeyRingImportJobProtectionLevel(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenKMSKeyRingImportJobExpireTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenKMSKeyRingImportJobState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenKMSKeyRingImportJobPublicKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["pem"] =
		flattenKMSKeyRingImportJobPublicKeyPem(original["pem"], d, config)
	return []interface{}{transformed}
}
func flattenKMSKeyRingImportJobPublicKeyPem(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenKMSKeyRingImportJobAttestation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["format"] =
		flattenKMSKeyRingImportJobAttestationFormat(original["format"], d, config)
	transformed["content"] =
		flattenKMSKeyRingImportJobAttestationContent(original["content"], d, config)
	return []interface{}{transformed}
}
func flattenKMSKeyRingImportJobAttestationFormat(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenKMSKeyRingImportJobAttestationContent(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandKMSKeyRingImportJobImportMethod(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandKMSKeyRingImportJobProtectionLevel(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
