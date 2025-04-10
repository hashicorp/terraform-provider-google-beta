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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/iap/Brand.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package iap

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

func ResourceIapBrand() *schema.Resource {
	return &schema.Resource{
		Create: resourceIapBrandCreate,
		Read:   resourceIapBrandRead,
		Delete: resourceIapBrandDelete,

		Importer: &schema.ResourceImporter{
			State: resourceIapBrandImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"application_title": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Application name displayed on OAuth consent screen.`,
			},
			"support_email": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Support email displayed on the OAuth consent screen. Can be either a
user or group email. When a user email is specified, the caller must
be the user with the associated email address. When a group email is
specified, the caller can be either a user or a service account which
is an owner of the specified group in Cloud Identity.`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Output only. Identifier of the brand, in the format 'projects/{project_number}/brands/{brand_id}'
NOTE: The name can also be expressed as 'projects/{project_id}/brands/{brand_id}', e.g. when importing.
NOTE: The brand identification corresponds to the project number as only one
brand can be created per project.`,
			},
			"org_internal_only": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: `Whether the brand is only intended for usage inside the GSuite organization only.`,
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

func resourceIapBrandCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	supportEmailProp, err := expandIapBrandSupportEmail(d.Get("support_email"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("support_email"); !tpgresource.IsEmptyValue(reflect.ValueOf(supportEmailProp)) && (ok || !reflect.DeepEqual(v, supportEmailProp)) {
		obj["supportEmail"] = supportEmailProp
	}
	applicationTitleProp, err := expandIapBrandApplicationTitle(d.Get("application_title"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("application_title"); !tpgresource.IsEmptyValue(reflect.ValueOf(applicationTitleProp)) && (ok || !reflect.DeepEqual(v, applicationTitleProp)) {
		obj["applicationTitle"] = applicationTitleProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IapBasePath}}projects/{{project}}/brands")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Brand: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Brand: %s", err)
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
		return fmt.Errorf("Error creating Brand: %s", err)
	}
	// Set computed resource properties from create API response so that they're available on the subsequent Read
	// call.
	if err := d.Set("name", flattenIapBrandName(res["name"], d, config)); err != nil {
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
		respBody, ok := res["response"]
		if !ok {
			return fmt.Errorf("Create response didn't contain critical fields. Create may not have succeeded.")
		}

		name, ok = respBody.(map[string]interface{})["name"]
		if !ok {
			return fmt.Errorf("Create response didn't contain critical fields. Create may not have succeeded.")
		}
	}
	if err := d.Set("name", name.(string)); err != nil {
		return fmt.Errorf("Error setting name: %s", err)
	}
	d.SetId(name.(string))

	err = transport_tpg.PollingWaitTime(resourceIapBrandPollRead(d, meta), transport_tpg.PollCheckForExistence, "Creating Brand", d.Timeout(schema.TimeoutCreate), 5)
	if err != nil {
		return fmt.Errorf("Error waiting to create Brand: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Brand %q: %#v", d.Id(), res)

	return resourceIapBrandRead(d, meta)
}

func resourceIapBrandPollRead(d *schema.ResourceData, meta interface{}) transport_tpg.PollReadFunc {
	return func() (map[string]interface{}, error) {
		config := meta.(*transport_tpg.Config)

		url, err := tpgresource.ReplaceVars(d, config, "{{IapBasePath}}{{name}}")

		if err != nil {
			return nil, err
		}

		billingProject := ""

		project, err := tpgresource.GetProject(d, config)
		if err != nil {
			return nil, fmt.Errorf("Error fetching project for Brand: %s", err)
		}
		billingProject = project

		// err == nil indicates that the billing_project value was found
		if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
			billingProject = bp
		}

		userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
		if err != nil {
			return nil, err
		}

		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "GET",
			Project:   billingProject,
			RawURL:    url,
			UserAgent: userAgent,
		})
		if err != nil {
			return res, err
		}
		return res, nil
	}
}

func resourceIapBrandRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IapBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Brand: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("IapBrand %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Brand: %s", err)
	}

	if err := d.Set("support_email", flattenIapBrandSupportEmail(res["supportEmail"], d, config)); err != nil {
		return fmt.Errorf("Error reading Brand: %s", err)
	}
	if err := d.Set("application_title", flattenIapBrandApplicationTitle(res["applicationTitle"], d, config)); err != nil {
		return fmt.Errorf("Error reading Brand: %s", err)
	}
	if err := d.Set("org_internal_only", flattenIapBrandOrgInternalOnly(res["orgInternalOnly"], d, config)); err != nil {
		return fmt.Errorf("Error reading Brand: %s", err)
	}
	if err := d.Set("name", flattenIapBrandName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Brand: %s", err)
	}

	return nil
}

func resourceIapBrandDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARNING] Iap Brand resources"+
		" cannot be deleted from Google Cloud. The resource %s will be removed from Terraform"+
		" state, but will still be present on Google Cloud.", d.Id())
	d.SetId("")

	return nil
}

func resourceIapBrandImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := tpgresource.ParseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}

	nameParts := strings.Split(d.Get("name").(string), "/")
	if len(nameParts) != 4 && len(nameParts) != 2 {
		return nil, fmt.Errorf(
			"Saw %s when the name is expected to have either shape %s or %s",
			d.Get("name"),
			"projects/{{project}}/brands/{{name}}",
			"{{project}}/{{name}}",
		)
	}

	var project string
	if len(nameParts) == 4 {
		project = nameParts[1]
	}
	if len(nameParts) == 2 {
		project = nameParts[0] // Different index

		// Set `name` (and `id`) as a 4-part format so Read func produces valid URL
		brand := nameParts[1]
		name := fmt.Sprintf("projects/%s/brands/%s", project, brand)
		if err := d.Set("name", name); err != nil {
			return nil, fmt.Errorf("Error setting name: %s", err)
		}
		d.SetId(name)
	}

	if err := d.Set("project", project); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	return []*schema.ResourceData{d}, nil
}

func flattenIapBrandSupportEmail(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIapBrandApplicationTitle(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIapBrandOrgInternalOnly(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIapBrandName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandIapBrandSupportEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIapBrandApplicationTitle(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
