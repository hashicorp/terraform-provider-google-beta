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
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceFirebaseAppleApp() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirebaseAppleAppCreate,
		Read:   resourceFirebaseAppleAppRead,
		Update: resourceFirebaseAppleAppUpdate,
		Delete: resourceFirebaseAppleAppDelete,

		Importer: &schema.ResourceImporter{
			State: resourceFirebaseAppleAppImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The user-assigned display name of the App.`,
			},
			"app_store_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The automatically generated Apple ID assigned to the Apple app by Apple in the Apple App Store.`,
			},
			"bundle_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The canonical bundle ID of the Apple app as it would appear in the Apple AppStore.`,
			},
			"team_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The Apple Developer Team ID associated with the App in the App Store.`,
			},
			"app_id": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The globally unique, Firebase-assigned identifier of the App.
This identifier should be treated as an opaque token, as the data format is not specified.`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The fully qualified resource name of the App, for example:
projects/projectId/iosApps/appId`,
			},
			"deletion_policy": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DELETE",
				Description: `(Optional) Set to 'ABANDON' to allow the Apple to be untracked from terraform state
rather than deleted upon 'terraform destroy'. This is useful because the Apple may be
serving traffic. Set to 'DELETE' to delete the Apple. Defaults to 'DELETE'.`,
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

func resourceFirebaseAppleAppCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandFirebaseAppleAppDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	bundleIdProp, err := expandFirebaseAppleAppBundleId(d.Get("bundle_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("bundle_id"); !isEmptyValue(reflect.ValueOf(bundleIdProp)) && (ok || !reflect.DeepEqual(v, bundleIdProp)) {
		obj["bundleId"] = bundleIdProp
	}
	appStoreIdProp, err := expandFirebaseAppleAppAppStoreId(d.Get("app_store_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("app_store_id"); !isEmptyValue(reflect.ValueOf(appStoreIdProp)) && (ok || !reflect.DeepEqual(v, appStoreIdProp)) {
		obj["appStoreId"] = appStoreIdProp
	}
	teamIdProp, err := expandFirebaseAppleAppTeamId(d.Get("team_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("team_id"); !isEmptyValue(reflect.ValueOf(teamIdProp)) && (ok || !reflect.DeepEqual(v, teamIdProp)) {
		obj["teamId"] = teamIdProp
	}

	url, err := replaceVars(d, config, "{{FirebaseBasePath}}projects/{{project}}/iosApps")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new AppleApp: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for AppleApp: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating AppleApp: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = FirebaseOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating AppleApp", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create AppleApp: %s", err)
	}

	if err := d.Set("name", flattenFirebaseAppleAppName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating AppleApp %q: %#v", d.Id(), res)

	return resourceFirebaseAppleAppRead(d, meta)
}

func resourceFirebaseAppleAppRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{FirebaseBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for AppleApp: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("FirebaseAppleApp %q", d.Id()))
	}

	// Explicitly set virtual fields to default values if unset
	if _, ok := d.GetOkExists("deletion_policy"); !ok {
		if err := d.Set("deletion_policy", "DELETE"); err != nil {
			return fmt.Errorf("Error setting deletion_policy: %s", err)
		}
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading AppleApp: %s", err)
	}

	if err := d.Set("name", flattenFirebaseAppleAppName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading AppleApp: %s", err)
	}
	if err := d.Set("display_name", flattenFirebaseAppleAppDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading AppleApp: %s", err)
	}
	if err := d.Set("app_id", flattenFirebaseAppleAppAppId(res["appId"], d, config)); err != nil {
		return fmt.Errorf("Error reading AppleApp: %s", err)
	}
	if err := d.Set("bundle_id", flattenFirebaseAppleAppBundleId(res["bundleId"], d, config)); err != nil {
		return fmt.Errorf("Error reading AppleApp: %s", err)
	}
	if err := d.Set("app_store_id", flattenFirebaseAppleAppAppStoreId(res["appStoreId"], d, config)); err != nil {
		return fmt.Errorf("Error reading AppleApp: %s", err)
	}
	if err := d.Set("team_id", flattenFirebaseAppleAppTeamId(res["teamId"], d, config)); err != nil {
		return fmt.Errorf("Error reading AppleApp: %s", err)
	}

	return nil
}

func resourceFirebaseAppleAppUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for AppleApp: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandFirebaseAppleAppDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	bundleIdProp, err := expandFirebaseAppleAppBundleId(d.Get("bundle_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("bundle_id"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, bundleIdProp)) {
		obj["bundleId"] = bundleIdProp
	}
	appStoreIdProp, err := expandFirebaseAppleAppAppStoreId(d.Get("app_store_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("app_store_id"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, appStoreIdProp)) {
		obj["appStoreId"] = appStoreIdProp
	}
	teamIdProp, err := expandFirebaseAppleAppTeamId(d.Get("team_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("team_id"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, teamIdProp)) {
		obj["teamId"] = teamIdProp
	}

	url, err := replaceVars(d, config, "{{FirebaseBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating AppleApp %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("bundle_id") {
		updateMask = append(updateMask, "bundleId")
	}

	if d.HasChange("app_store_id") {
		updateMask = append(updateMask, "appStoreId")
	}

	if d.HasChange("team_id") {
		updateMask = append(updateMask, "teamId")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating AppleApp %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating AppleApp %q: %#v", d.Id(), res)
	}

	return resourceFirebaseAppleAppRead(d, meta)
}

func resourceFirebaseAppleAppDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	// Handwritten
	obj := make(map[string]interface{})
	if d.Get("deletion_policy") == "DELETE" {
		obj["immediate"] = true
	} else {
		fmt.Printf("Skip deleting App %q due to deletion_policy: %q\n", d.Id(), d.Get("deletion_policy"))
		return nil
	}
	// End of Handwritten
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for App: %s", err)
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{FirebaseBasePath}}{{name}}:remove")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Deleting App %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "App")
	}

	err = FirebaseOperationWaitTime(
		config, res, project, "Deleting App", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting App %q: %#v", d.Id(), res)

	// This is useful if the Delete operation returns before the Get operation
	// during post-test destroy shows the completed state of the resource.
	time.Sleep(5 * time.Second)

	return nil
}

func resourceFirebaseAppleAppImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	config := meta.(*Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := parseImportId([]string{"(?P<project>[^ ]+) (?P<name>[^ ]+)", "(?P<name>[^ ]+)"}, d, config); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func flattenFirebaseAppleAppName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFirebaseAppleAppDisplayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFirebaseAppleAppAppId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFirebaseAppleAppBundleId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFirebaseAppleAppAppStoreId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFirebaseAppleAppTeamId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandFirebaseAppleAppDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseAppleAppBundleId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseAppleAppAppStoreId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseAppleAppTeamId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
