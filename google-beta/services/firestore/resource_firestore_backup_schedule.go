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

package firestore

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceFirestoreBackupSchedule() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirestoreBackupScheduleCreate,
		Read:   resourceFirestoreBackupScheduleRead,
		Update: resourceFirestoreBackupScheduleUpdate,
		Delete: resourceFirestoreBackupScheduleDelete,

		Importer: &schema.ResourceImporter{
			State: resourceFirestoreBackupScheduleImport,
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
			"retention": {
				Type:     schema.TypeString,
				Required: true,
				Description: `At what relative time in the future, compared to its creation time, the backup should be deleted, e.g. keep backups for 7 days.
A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".

For a daily backup recurrence, set this to a value up to 7 days. If you set a weekly backup recurrence, set this to a value up to 14 weeks.`,
			},
			"daily_recurrence": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `For a schedule that runs daily.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{},
				},
				ExactlyOneOf: []string{"daily_recurrence", "weekly_recurrence"},
			},
			"database": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The Firestore database id. Defaults to '"(default)"'.`,
				Default:     "(default)",
			},
			"weekly_recurrence": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `For a schedule that runs weekly on a specific day.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"day": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: verify.ValidateEnum([]string{"DAY_OF_WEEK_UNSPECIFIED", "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY", ""}),
							Description:  `The day of week to run. Possible values: ["DAY_OF_WEEK_UNSPECIFIED", "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"]`,
						},
					},
				},
				ExactlyOneOf: []string{"weekly_recurrence", "daily_recurrence"},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The unique backup schedule identifier across all locations and databases for the given project. Format:
'projects/{{project}}/databases/{{database}}/backupSchedules/{{backupSchedule}}'`,
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

func resourceFirestoreBackupScheduleCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	retentionProp, err := expandFirestoreBackupScheduleRetention(d.Get("retention"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("retention"); !tpgresource.IsEmptyValue(reflect.ValueOf(retentionProp)) && (ok || !reflect.DeepEqual(v, retentionProp)) {
		obj["retention"] = retentionProp
	}
	dailyRecurrenceProp, err := expandFirestoreBackupScheduleDailyRecurrence(d.Get("daily_recurrence"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("daily_recurrence"); ok || !reflect.DeepEqual(v, dailyRecurrenceProp) {
		obj["dailyRecurrence"] = dailyRecurrenceProp
	}
	weeklyRecurrenceProp, err := expandFirestoreBackupScheduleWeeklyRecurrence(d.Get("weekly_recurrence"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("weekly_recurrence"); !tpgresource.IsEmptyValue(reflect.ValueOf(weeklyRecurrenceProp)) && (ok || !reflect.DeepEqual(v, weeklyRecurrenceProp)) {
		obj["weeklyRecurrence"] = weeklyRecurrenceProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{FirestoreBasePath}}projects/{{project}}/databases/{{database}}/backupSchedules")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new BackupSchedule: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for BackupSchedule: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating BackupSchedule: %s", err)
	}
	if err := d.Set("name", flattenFirestoreBackupScheduleName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/databases/{{database}}/backupSchedules/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating BackupSchedule %q: %#v", d.Id(), res)

	return resourceFirestoreBackupScheduleRead(d, meta)
}

func resourceFirestoreBackupScheduleRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{FirestoreBasePath}}projects/{{project}}/databases/{{database}}/backupSchedules/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for BackupSchedule: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("FirestoreBackupSchedule %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading BackupSchedule: %s", err)
	}

	if err := d.Set("name", flattenFirestoreBackupScheduleName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading BackupSchedule: %s", err)
	}
	if err := d.Set("retention", flattenFirestoreBackupScheduleRetention(res["retention"], d, config)); err != nil {
		return fmt.Errorf("Error reading BackupSchedule: %s", err)
	}
	if err := d.Set("daily_recurrence", flattenFirestoreBackupScheduleDailyRecurrence(res["dailyRecurrence"], d, config)); err != nil {
		return fmt.Errorf("Error reading BackupSchedule: %s", err)
	}
	if err := d.Set("weekly_recurrence", flattenFirestoreBackupScheduleWeeklyRecurrence(res["weeklyRecurrence"], d, config)); err != nil {
		return fmt.Errorf("Error reading BackupSchedule: %s", err)
	}

	return nil
}

func resourceFirestoreBackupScheduleUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for BackupSchedule: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	retentionProp, err := expandFirestoreBackupScheduleRetention(d.Get("retention"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("retention"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, retentionProp)) {
		obj["retention"] = retentionProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{FirestoreBasePath}}projects/{{project}}/databases/{{database}}/backupSchedules/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating BackupSchedule %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("retention") {
		updateMask = append(updateMask, "retention")
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
		})

		if err != nil {
			return fmt.Errorf("Error updating BackupSchedule %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating BackupSchedule %q: %#v", d.Id(), res)
		}

	}

	return resourceFirestoreBackupScheduleRead(d, meta)
}

func resourceFirestoreBackupScheduleDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for BackupSchedule: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{FirestoreBasePath}}projects/{{project}}/databases/{{database}}/backupSchedules/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	log.Printf("[DEBUG] Deleting BackupSchedule %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "BackupSchedule")
	}

	log.Printf("[DEBUG] Finished deleting BackupSchedule %q: %#v", d.Id(), res)
	return nil
}

func resourceFirestoreBackupScheduleImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/databases/(?P<database>[^/]+)/backupSchedules/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<database>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<database>[^/]+)/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/databases/{{database}}/backupSchedules/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenFirestoreBackupScheduleName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func flattenFirestoreBackupScheduleRetention(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirestoreBackupScheduleDailyRecurrence(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	transformed := make(map[string]interface{})
	return []interface{}{transformed}
}

func flattenFirestoreBackupScheduleWeeklyRecurrence(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["day"] =
		flattenFirestoreBackupScheduleWeeklyRecurrenceDay(original["day"], d, config)
	return []interface{}{transformed}
}
func flattenFirestoreBackupScheduleWeeklyRecurrenceDay(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandFirestoreBackupScheduleRetention(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreBackupScheduleDailyRecurrence(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}

	if l[0] == nil {
		transformed := make(map[string]interface{})
		return transformed, nil
	}
	transformed := make(map[string]interface{})

	return transformed, nil
}

func expandFirestoreBackupScheduleWeeklyRecurrence(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDay, err := expandFirestoreBackupScheduleWeeklyRecurrenceDay(original["day"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDay); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["day"] = transformedDay
	}

	return transformed, nil
}

func expandFirestoreBackupScheduleWeeklyRecurrenceDay(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
