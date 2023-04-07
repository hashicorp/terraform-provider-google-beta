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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceNotebooksLocation() *schema.Resource {
	return &schema.Resource{
		Create: resourceNotebooksLocationCreate,
		Read:   resourceNotebooksLocationRead,
		Update: resourceNotebooksLocationUpdate,
		Delete: resourceNotebooksLocationDelete,

		Importer: &schema.ResourceImporter{
			State: resourceNotebooksLocationImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Name of the Location resource.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceNotebooksLocationCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandNotebooksLocationName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	url, err := ReplaceVars(d, config, "{{NotebooksBasePath}}projects/{{project}}/locations")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Location: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Location: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Location: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "projects/{{project}}/locations/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = NotebooksOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating Location", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create Location: %s", err)
	}

	if err := d.Set("name", flattenNotebooksLocationName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = ReplaceVars(d, config, "projects/{{project}}/locations/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Location %q: %#v", d.Id(), res)

	return resourceNotebooksLocationRead(d, meta)
}

func resourceNotebooksLocationRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{NotebooksBasePath}}projects/{{project}}/locations/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Location: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("NotebooksLocation %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Location: %s", err)
	}

	if err := d.Set("name", flattenNotebooksLocationName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Location: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading Location: %s", err)
	}

	return nil
}

func resourceNotebooksLocationUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Location: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	nameProp, err := expandNotebooksLocationName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	url, err := ReplaceVars(d, config, "{{NotebooksBasePath}}projects/{{project}}/locations/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Location %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "PUT", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Location %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Location %q: %#v", d.Id(), res)
	}

	err = NotebooksOperationWaitTime(
		config, res, project, "Updating Location", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceNotebooksLocationRead(d, meta)
}

func resourceNotebooksLocationDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Location: %s", err)
	}
	billingProject = project

	url, err := ReplaceVars(d, config, "{{NotebooksBasePath}}projects/{{project}}/locations/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Location %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Location")
	}

	err = NotebooksOperationWaitTime(
		config, res, project, "Deleting Location", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Location %q: %#v", d.Id(), res)
	return nil
}

func resourceNotebooksLocationImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "projects/{{project}}/locations/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNotebooksLocationName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func expandNotebooksLocationName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
