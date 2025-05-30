// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: Handwritten     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/services/apigee/resource_apigee_api.go
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
//
//     This file is partially automatically generated by Magic Modules and with manual
//     changes to resourceApigeeApiCreate
//
// ----------------------------------------------------------------------------

package apigee

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceApigeeApi() *schema.Resource {
	return &schema.Resource{
		Create: resourceApigeeApiCreate,
		Read:   resourceApigeeApiRead,
		Update: resourceApigeeApiUpdate,
		Delete: resourceApigeeApiDelete,

		Importer: &schema.ResourceImporter{
			State: resourceApigeeApiImport,
		},

		CustomizeDiff: customdiff.All(
			/*
				If any of the config_bundle, detect_md5hash or md5hash is changed,
				then an update is expected, so we tell Terraform core to expect update on meta_data,
				latest_revision_id and revision
			*/

			customdiff.ComputedIf("meta_data", apigeeApiDetectBundleUpdate),
			customdiff.ComputedIf("latest_revision_id", apigeeApiDetectBundleUpdate),
			customdiff.ComputedIf("revision", apigeeApiDetectBundleUpdate),
		),

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Name of the API proxy. This field only accepts the following characters: A-Za-z0-9._-.`,
			},
			"org_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The Apigee Organization name associated with the Apigee instance.`,
			},
			"latest_revision_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The id of the most recently created revision for this API proxy.`,
			},
			"meta_data": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `Metadata describing the API proxy.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Time at which the API proxy was created, in milliseconds since epoch.`,
						},
						"last_modified_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Time at which the API proxy was most recently modified, in milliseconds since epoch.`,
						},
						"sub_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The type of entity described`,
						},
					},
				},
			},
			"revision": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `A list of revisions of this API proxy.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"config_bundle": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Path to the config zip bundle`,
			},
			"md5hash": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Base 64 MD5 hash of the uploaded config bundle.`,
			},
			"detect_md5hash": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A hash of local config bundle in string, user needs to use a Terraform Hash function of their choice. A change in hash will trigger an update.`,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					localMd5Hash := ""
					if config_bundle, ok := d.GetOkExists("config_bundle"); ok {
						localMd5Hash = tpgresource.GetFileMd5Hash(config_bundle.(string))
					}
					if localMd5Hash == "" {
						return false
					}

					// `old` is the md5 hash we speculated from server responses,
					// when apply responded with succeed, hash is set to the hash of uploaded bundle
					if old != localMd5Hash {
						return false
					}

					return true
				},
			},
		},
		UseJSONNumber: true,
	}
}

func resourceApigeeApiCreate(d *schema.ResourceData, meta interface{}) error {
	ctx := context.TODO()
	tflog.Info(ctx, "resourceApigeeApiCreate")
	log.Printf("[DEBUG] resourceApigeeApiCreate")

	log.Printf("[DEBUG] resourceApigeeApiCreate, name=			 	%s", d.Get("name").(string))
	log.Printf("[DEBUG] resourceApigeeApiCreate, org_id=, 			%s", d.Get("org_id").(string))
	log.Printf("[DEBUG] resourceApigeeApiCreate, config_bundle=, 	%s", d.Get("config_bundle").(string))

	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	var file *os.File
	var localMd5Hash string
	if configBundlePath, ok := d.GetOk("config_bundle"); ok {
		var err error
		file, err = os.Open(configBundlePath.(string))
		if err != nil {
			return err
		}
		localMd5Hash = tpgresource.GetFileMd5Hash(configBundlePath.(string))
	} else {
		return fmt.Errorf("Error, \"config_bundle\" must be specified")
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ApigeeBasePath}}organizations/{{org_id}}/apis?name={{name}}&action=import")
	if err != nil {
		return err
	}
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	log.Printf("[DEBUG] resourceApigeeApiCreate, url=, 	%s", url)
	res, err := sendRequestRawBodyWithTimeout(config, "POST", billingProject, url, userAgent, file, "application/octet-stream", d.Timeout(schema.TimeoutCreate))

	log.Printf("[DEBUG] sendRequestRawBodyWithTimeout Done")
	if err != nil {
		return fmt.Errorf("Error creating API proxy: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "organizations/{{org_id}}/apis/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)
	log.Printf("[DEBUG] create d.SetId done, id = %s", id)

	log.Printf("[DEBUG] Finished creating API proxy %q: %#v", d.Id(), res)

	if resourceApigeeApiRead(d, meta) != nil {
		return fmt.Errorf("Error reading API proxy at end of Create: %s", err)
	}

	d.Set("md5hash", localMd5Hash)
	d.Set("detect_md5hash", localMd5Hash)

	return nil
}

func resourceApigeeApiUpdate(d *schema.ResourceData, meta interface{}) error {
	//For how API proxy api is implemented, just treat an update as create, when the name is same, it will create a new revision
	return resourceApigeeApiCreate(d, meta)
}

func resourceApigeeApiRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ApigeeBasePath}}organizations/{{org_id}}/apis/{{name}}")
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] API proxy read url is: %s", url)

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}
	log.Printf("[DEBUG] resourceApigeeApiRead sendRequest")
	log.Printf("[DEBUG] resourceApigeeApiRead, url=, 	%s", url)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ApigeeApi %q", d.Id()))
	}
	log.Printf("[DEBUG] resourceApigeeApuRead sendRequest completed")
	previousLastModifiedAt := getApigeeApiLastModifiedAt(d)
	if err := d.Set("meta_data", flattenApigeeApiMetaData(res["metaData"], d, config)); err != nil {
		return fmt.Errorf("Error reading API proxy: %s", err)
	}
	currentLastModifiedAt := getApigeeApiLastModifiedAt(d)
	if err := d.Set("name", flattenApigeeApiName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading API proxy: %s", err)
	}
	if err := d.Set("revision", flattenApigeeApiRevision(res["revision"], d, config)); err != nil {
		return fmt.Errorf("Error reading API proxy: %s", err)
	}
	if err := d.Set("latest_revision_id", flattenApigeeApiLatestRevisionId(res["latestRevisionId"], d, config)); err != nil {
		return fmt.Errorf("Error reading API proxy: %s", err)
	}

	//setting hash to suggest update
	if previousLastModifiedAt != currentLastModifiedAt {
		d.Set("md5hash", "UNKNOWN")
		d.Set("detect_md5hash", "UNKNOWN")
	}
	return nil
}

func getApigeeApiLastModifiedAt(d *schema.ResourceData) string {

	metaDataRaw := d.Get("meta_data").([]interface{})
	if len(metaDataRaw) != 1 {
		//in Terraform Schema, a nest in object is implemented as an array of length one, even if it's technically an object
		return "UNKNOWN"
	}
	metaData := metaDataRaw[0].(map[string]interface{})
	if metaData == nil {
		return "UNKNOWN"
	}
	lastModifiedAt := metaData["last_modified_at"].(string)
	if lastModifiedAt == "" {
		return "UNKNOWN"
	}
	return lastModifiedAt
}

func resourceApigeeApiDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{ApigeeBasePath}}organizations/{{org_id}}/apis/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting API proxy %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Api")
	}

	log.Printf("[DEBUG] Finished deleting API proxy %q: %#v", d.Id(), res)
	return nil
}

func resourceApigeeApiImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"organizations/(?P<org_id>[^/]+)/apis/(?P<name>[^/]+)",
		"(?P<org_id>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "organizations/{{org_id}}/apis/{{name}}")

	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)
	log.Printf("[DEBUG] resourceApigeeApiImport, id=			 	%s", id)

	return []*schema.ResourceData{d}, nil
}

func flattenApigeeApiMetaData(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["created_at"] =
		flattenApigeeApiMetaDataCreatedAt(original["createdAt"], d, config)
	transformed["last_modified_at"] =
		flattenApigeeApiMetaDataLastModifiedAt(original["lastModifiedAt"], d, config)
	transformed["sub_type"] =
		flattenApigeeApiMetaDataSubType(original["subType"], d, config)
	return []interface{}{transformed}
}
func flattenApigeeApiMetaDataCreatedAt(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeApiMetaDataLastModifiedAt(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeApiMetaDataSubType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeApiName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeApiRevision(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeApiLatestRevisionId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandApigeeApiName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func apigeeApiDetectBundleUpdate(_ context.Context, diff *schema.ResourceDiff, v interface{}) bool {
	tmp, _ := diff.GetChange("detect_md5hash")
	oldBundleHash := tmp.(string)
	currentBundleHash := ""
	if config_bundle, ok := diff.GetOkExists("config_bundle"); ok {
		currentBundleHash = tpgresource.GetFileMd5Hash(config_bundle.(string))
	}
	log.Printf("[DEBUG] apigeeApiDetectUpdate detect_md5hash: %s -> %s", oldBundleHash, currentBundleHash)

	if oldBundleHash != currentBundleHash {
		return true
	}
	return diff.HasChange("config_bundle") || diff.HasChange("md5hash")
}
