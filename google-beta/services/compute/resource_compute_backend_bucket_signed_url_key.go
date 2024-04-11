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

package compute

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
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceComputeBackendBucketSignedUrlKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeBackendBucketSignedUrlKeyCreate,
		Read:   resourceComputeBackendBucketSignedUrlKeyRead,
		Delete: resourceComputeBackendBucketSignedUrlKeyDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"backend_bucket": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The backend bucket this signed URL key belongs.`,
			},
			"key_value": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `128-bit key value used for signing the URL. The key value must be a
valid RFC 4648 Section 5 base64url encoded string.`,
				Sensitive: true,
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateRegexp(`^(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?)$`),
				Description:  `Name of the signed URL key.`,
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

func resourceComputeBackendBucketSignedUrlKeyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	keyNameProp, err := expandNestedComputeBackendBucketSignedUrlKeyName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(keyNameProp)) && (ok || !reflect.DeepEqual(v, keyNameProp)) {
		obj["keyName"] = keyNameProp
	}
	keyValueProp, err := expandNestedComputeBackendBucketSignedUrlKeyKeyValue(d.Get("key_value"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("key_value"); !tpgresource.IsEmptyValue(reflect.ValueOf(keyValueProp)) && (ok || !reflect.DeepEqual(v, keyValueProp)) {
		obj["keyValue"] = keyValueProp
	}
	backendBucketProp, err := expandNestedComputeBackendBucketSignedUrlKeyBackendBucket(d.Get("backend_bucket"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("backend_bucket"); !tpgresource.IsEmptyValue(reflect.ValueOf(backendBucketProp)) && (ok || !reflect.DeepEqual(v, backendBucketProp)) {
		obj["backendBucket"] = backendBucketProp
	}

	lockName, err := tpgresource.ReplaceVars(d, config, "signedUrlKey/{{project}}/backendBuckets/{{backend_bucket}}/")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/backendBuckets/{{backend_bucket}}/addSignedUrlKey")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new BackendBucketSignedUrlKey: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for BackendBucketSignedUrlKey: %s", err)
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
		return fmt.Errorf("Error creating BackendBucketSignedUrlKey: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/global/backendBuckets/{{backend_bucket}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ComputeOperationWaitTime(
		config, res, project, "Creating BackendBucketSignedUrlKey", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create BackendBucketSignedUrlKey: %s", err)
	}

	log.Printf("[DEBUG] Finished creating BackendBucketSignedUrlKey %q: %#v", d.Id(), res)

	return resourceComputeBackendBucketSignedUrlKeyRead(d, meta)
}

func resourceComputeBackendBucketSignedUrlKeyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/backendBuckets/{{backend_bucket}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for BackendBucketSignedUrlKey: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ComputeBackendBucketSignedUrlKey %q", d.Id()))
	}

	res, err = flattenNestedComputeBackendBucketSignedUrlKey(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Object isn't there any more - remove it from the state.
		log.Printf("[DEBUG] Removing ComputeBackendBucketSignedUrlKey because it couldn't be matched.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading BackendBucketSignedUrlKey: %s", err)
	}

	if err := d.Set("name", flattenNestedComputeBackendBucketSignedUrlKeyName(res["keyName"], d, config)); err != nil {
		return fmt.Errorf("Error reading BackendBucketSignedUrlKey: %s", err)
	}

	return nil
}

func resourceComputeBackendBucketSignedUrlKeyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for BackendBucketSignedUrlKey: %s", err)
	}
	billingProject = project

	lockName, err := tpgresource.ReplaceVars(d, config, "signedUrlKey/{{project}}/backendBuckets/{{backend_bucket}}/")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/backendBuckets/{{backend_bucket}}/deleteSignedUrlKey?keyName={{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting BackendBucketSignedUrlKey %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "BackendBucketSignedUrlKey")
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Deleting BackendBucketSignedUrlKey", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting BackendBucketSignedUrlKey %q: %#v", d.Id(), res)
	return nil
}

func flattenNestedComputeBackendBucketSignedUrlKeyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNestedComputeBackendBucketSignedUrlKeyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNestedComputeBackendBucketSignedUrlKeyKeyValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNestedComputeBackendBucketSignedUrlKeyBackendBucket(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("backendBuckets", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for backend_bucket: %s", err)
	}
	return f.RelativeLink(), nil
}

func flattenNestedComputeBackendBucketSignedUrlKey(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	var v interface{}
	var ok bool

	v, ok = res["cdnPolicy"]
	if !ok || v == nil {
		return nil, nil
	}
	res = v.(map[string]interface{})

	v, ok = res["signedUrlKeyNames"]
	if !ok || v == nil {
		return nil, nil
	}

	switch v.(type) {
	case []interface{}:
		break
	case map[string]interface{}:
		// Construct list out of single nested resource
		v = []interface{}{v}
	default:
		return nil, fmt.Errorf("expected list or map for value cdnPolicy.signedUrlKeyNames. Actual value: %v", v)
	}

	_, item, err := resourceComputeBackendBucketSignedUrlKeyFindNestedObjectInList(d, meta, v.([]interface{}))
	if err != nil {
		return nil, err
	}
	return item, nil
}

func resourceComputeBackendBucketSignedUrlKeyFindNestedObjectInList(d *schema.ResourceData, meta interface{}, items []interface{}) (index int, item map[string]interface{}, err error) {
	expectedName, err := expandNestedComputeBackendBucketSignedUrlKeyName(d.Get("name"), d, meta.(*transport_tpg.Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedName := flattenNestedComputeBackendBucketSignedUrlKeyName(expectedName, d, meta.(*transport_tpg.Config))

	// Search list for this resource.
	for idx, itemRaw := range items {
		if itemRaw == nil {
			continue
		}
		// List response only contains the ID - construct a response object.
		item := map[string]interface{}{
			"keyName": itemRaw,
		}

		itemName := flattenNestedComputeBackendBucketSignedUrlKeyName(item["keyName"], d, meta.(*transport_tpg.Config))
		// IsEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(tpgresource.IsEmptyValue(reflect.ValueOf(itemName)) && tpgresource.IsEmptyValue(reflect.ValueOf(expectedFlattenedName))) && !reflect.DeepEqual(itemName, expectedFlattenedName) {
			log.Printf("[DEBUG] Skipping item with keyName= %#v, looking for %#v)", itemName, expectedFlattenedName)
			continue
		}
		log.Printf("[DEBUG] Found item for resource %q: %#v)", d.Id(), item)
		return idx, item, nil
	}
	return -1, nil, nil
}
