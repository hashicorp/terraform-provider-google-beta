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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/firebaseappcheck/RecaptchaEnterpriseConfig.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package firebaseappcheck

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

func ResourceFirebaseAppCheckRecaptchaEnterpriseConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirebaseAppCheckRecaptchaEnterpriseConfigCreate,
		Read:   resourceFirebaseAppCheckRecaptchaEnterpriseConfigRead,
		Update: resourceFirebaseAppCheckRecaptchaEnterpriseConfigUpdate,
		Delete: resourceFirebaseAppCheckRecaptchaEnterpriseConfigDelete,

		Importer: &schema.ResourceImporter{
			State: resourceFirebaseAppCheckRecaptchaEnterpriseConfigImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"app_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The ID of an
[Web App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.webApps#WebApp.FIELDS.app_id).`,
			},
			"site_key": {
				Type:     schema.TypeString,
				Required: true,
				Description: `The score-based site key created in reCAPTCHA Enterprise used to invoke reCAPTCHA and generate the reCAPTCHA tokens for your application.

**Important**: This is not the siteSecret (as it is in reCAPTCHA v3), but rather your score-based reCAPTCHA Enterprise site key.`,
			},
			"token_ttl": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				Description: `Specifies the duration for which App Check tokens exchanged from reCAPTCHA Enterprise artifacts will be valid.
If unset, a default value of 1 hour is assumed. Must be between 30 minutes and 7 days, inclusive.

A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The relative resource name of the reCAPTCHA Enterprise configuration object`,
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

func resourceFirebaseAppCheckRecaptchaEnterpriseConfigCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	tokenTtlProp, err := expandFirebaseAppCheckRecaptchaEnterpriseConfigTokenTtl(d.Get("token_ttl"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("token_ttl"); !tpgresource.IsEmptyValue(reflect.ValueOf(tokenTtlProp)) && (ok || !reflect.DeepEqual(v, tokenTtlProp)) {
		obj["tokenTtl"] = tokenTtlProp
	}
	siteKeyProp, err := expandFirebaseAppCheckRecaptchaEnterpriseConfigSiteKey(d.Get("site_key"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("site_key"); !tpgresource.IsEmptyValue(reflect.ValueOf(siteKeyProp)) && (ok || !reflect.DeepEqual(v, siteKeyProp)) {
		obj["siteKey"] = siteKeyProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseAppCheckBasePath}}projects/{{project}}/apps/{{app_id}}/recaptchaEnterpriseConfig?updateMask=tokenTtl,siteKey")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new RecaptchaEnterpriseConfig: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RecaptchaEnterpriseConfig: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating RecaptchaEnterpriseConfig: %s", err)
	}
	if err := d.Set("name", flattenFirebaseAppCheckRecaptchaEnterpriseConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/apps/{{app_id}}/recaptchaEnterpriseConfig")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating RecaptchaEnterpriseConfig %q: %#v", d.Id(), res)

	return resourceFirebaseAppCheckRecaptchaEnterpriseConfigRead(d, meta)
}

func resourceFirebaseAppCheckRecaptchaEnterpriseConfigRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseAppCheckBasePath}}projects/{{project}}/apps/{{app_id}}/recaptchaEnterpriseConfig")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RecaptchaEnterpriseConfig: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("FirebaseAppCheckRecaptchaEnterpriseConfig %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading RecaptchaEnterpriseConfig: %s", err)
	}

	if err := d.Set("name", flattenFirebaseAppCheckRecaptchaEnterpriseConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading RecaptchaEnterpriseConfig: %s", err)
	}
	if err := d.Set("token_ttl", flattenFirebaseAppCheckRecaptchaEnterpriseConfigTokenTtl(res["tokenTtl"], d, config)); err != nil {
		return fmt.Errorf("Error reading RecaptchaEnterpriseConfig: %s", err)
	}
	if err := d.Set("site_key", flattenFirebaseAppCheckRecaptchaEnterpriseConfigSiteKey(res["siteKey"], d, config)); err != nil {
		return fmt.Errorf("Error reading RecaptchaEnterpriseConfig: %s", err)
	}

	return nil
}

func resourceFirebaseAppCheckRecaptchaEnterpriseConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RecaptchaEnterpriseConfig: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	tokenTtlProp, err := expandFirebaseAppCheckRecaptchaEnterpriseConfigTokenTtl(d.Get("token_ttl"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("token_ttl"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, tokenTtlProp)) {
		obj["tokenTtl"] = tokenTtlProp
	}
	siteKeyProp, err := expandFirebaseAppCheckRecaptchaEnterpriseConfigSiteKey(d.Get("site_key"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("site_key"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, siteKeyProp)) {
		obj["siteKey"] = siteKeyProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseAppCheckBasePath}}projects/{{project}}/apps/{{app_id}}/recaptchaEnterpriseConfig")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating RecaptchaEnterpriseConfig %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("token_ttl") {
		updateMask = append(updateMask, "tokenTtl")
	}

	if d.HasChange("site_key") {
		updateMask = append(updateMask, "siteKey")
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
			return fmt.Errorf("Error updating RecaptchaEnterpriseConfig %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating RecaptchaEnterpriseConfig %q: %#v", d.Id(), res)
		}

	}

	return resourceFirebaseAppCheckRecaptchaEnterpriseConfigRead(d, meta)
}

func resourceFirebaseAppCheckRecaptchaEnterpriseConfigDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARNING] FirebaseAppCheck RecaptchaEnterpriseConfig resources"+
		" cannot be deleted from Google Cloud. The resource %s will be removed from Terraform"+
		" state, but will still be present on Google Cloud.", d.Id())
	d.SetId("")

	return nil
}

func resourceFirebaseAppCheckRecaptchaEnterpriseConfigImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/apps/(?P<app_id>[^/]+)/recaptchaEnterpriseConfig$",
		"^(?P<project>[^/]+)/(?P<app_id>[^/]+)$",
		"^(?P<app_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/apps/{{app_id}}/recaptchaEnterpriseConfig")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenFirebaseAppCheckRecaptchaEnterpriseConfigName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseAppCheckRecaptchaEnterpriseConfigTokenTtl(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseAppCheckRecaptchaEnterpriseConfigSiteKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandFirebaseAppCheckRecaptchaEnterpriseConfigTokenTtl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseAppCheckRecaptchaEnterpriseConfigSiteKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
