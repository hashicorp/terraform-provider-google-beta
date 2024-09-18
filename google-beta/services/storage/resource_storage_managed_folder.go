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

package storage

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceStorageManagedFolder() *schema.Resource {
	return &schema.Resource{
		Create: resourceStorageManagedFolderCreate,
		Read:   resourceStorageManagedFolderRead,
		Update: resourceStorageManagedFolderUpdate,
		Delete: resourceStorageManagedFolderDelete,

		Importer: &schema.ResourceImporter{
			State: resourceStorageManagedFolderImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"bucket": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The name of the bucket that contains the managed folder.`,
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateRegexp(`/$`),
				Description: `The name of the managed folder expressed as a path. Must include
trailing '/'. For example, 'example_dir/example_dir2/'.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The timestamp at which this managed folder was created.`,
			},
			"metageneration": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The metadata generation of the managed folder.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The timestamp at which this managed folder was most recently updated.`,
			},
			"force_destroy": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: `Allows the deletion of a managed folder even if contains
objects. If a non-empty managed folder is deleted, any objects
within the folder will remain in a simulated folder with the
same name.`,
				Default: false,
			},

			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceStorageManagedFolderCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	bucketProp, err := expandStorageManagedFolderBucket(d.Get("bucket"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("bucket"); !tpgresource.IsEmptyValue(reflect.ValueOf(bucketProp)) && (ok || !reflect.DeepEqual(v, bucketProp)) {
		obj["bucket"] = bucketProp
	}
	nameProp, err := expandStorageManagedFolderName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{StorageBasePath}}b/{{bucket}}/managedFolders")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ManagedFolder: %#v", obj)
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
		return fmt.Errorf("Error creating ManagedFolder: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{bucket}}/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating ManagedFolder %q: %#v", d.Id(), res)

	return resourceStorageManagedFolderRead(d, meta)
}

func resourceStorageManagedFolderRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{StorageBasePath}}b/{{bucket}}/managedFolders/{{%name}}")
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("StorageManagedFolder %q", d.Id()))
	}

	// Explicitly set virtual fields to default values if unset
	if _, ok := d.GetOkExists("force_destroy"); !ok {
		if err := d.Set("force_destroy", false); err != nil {
			return fmt.Errorf("Error setting force_destroy: %s", err)
		}
	}

	if err := d.Set("create_time", flattenStorageManagedFolderCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading ManagedFolder: %s", err)
	}
	if err := d.Set("update_time", flattenStorageManagedFolderUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading ManagedFolder: %s", err)
	}
	if err := d.Set("metageneration", flattenStorageManagedFolderMetageneration(res["metageneration"], d, config)); err != nil {
		return fmt.Errorf("Error reading ManagedFolder: %s", err)
	}
	if err := d.Set("bucket", flattenStorageManagedFolderBucket(res["bucket"], d, config)); err != nil {
		return fmt.Errorf("Error reading ManagedFolder: %s", err)
	}
	if err := d.Set("name", flattenStorageManagedFolderName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading ManagedFolder: %s", err)
	}
	if err := d.Set("self_link", tpgresource.ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading ManagedFolder: %s", err)
	}

	return nil
}

func resourceStorageManagedFolderUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	_ = config

	// we can only get here if force_destroy was updated
	if d.Get("force_destroy") != nil {
		if err := d.Set("force_destroy", d.Get("force_destroy")); err != nil {
			return fmt.Errorf("Error updating force_destroy: %s", err)
		}
	}

	// all other fields are immutable, don't do anything else
	return nil
}

func resourceStorageManagedFolderDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{StorageBasePath}}b/{{bucket}}/managedFolders/{{%name}}?allowNonEmpty={{force_destroy}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting ManagedFolder %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "ManagedFolder")
	}

	log.Printf("[DEBUG] Finished deleting ManagedFolder %q: %#v", d.Id(), res)
	return nil
}

func resourceStorageManagedFolderImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^(?P<bucket>[^/]+)/managedFolders/(?P<name>.+)$",
		"^(?P<bucket>[^/]+)/(?P<name>.+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "{{bucket}}/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Explicitly set virtual fields to default values on import
	if err := d.Set("force_destroy", false); err != nil {
		return nil, fmt.Errorf("Error setting force_destroy: %s", err)
	}

	return []*schema.ResourceData{d}, nil
}

func flattenStorageManagedFolderCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenStorageManagedFolderUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenStorageManagedFolderMetageneration(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenStorageManagedFolderBucket(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func flattenStorageManagedFolderName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandStorageManagedFolderBucket(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandStorageManagedFolderName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
