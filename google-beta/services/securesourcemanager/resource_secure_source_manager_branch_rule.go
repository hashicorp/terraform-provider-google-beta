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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/securesourcemanager/BranchRule.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package securesourcemanager

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

func ResourceSecureSourceManagerBranchRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecureSourceManagerBranchRuleCreate,
		Read:   resourceSecureSourceManagerBranchRuleRead,
		Update: resourceSecureSourceManagerBranchRuleUpdate,
		Delete: resourceSecureSourceManagerBranchRuleDelete,

		Importer: &schema.ResourceImporter{
			State: resourceSecureSourceManagerBranchRuleImport,
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
			"branch_rule_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The ID for the BranchRule.`,
			},
			"include_pattern": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The BranchRule matches branches based on the specified regular expression. Use .* to match all branches.`,
			},
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The location for the Repository.`,
			},
			"repository_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The ID for the Repository.`,
			},
			"allow_stale_reviews": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `Determines if allow stale reviews or approvals before merging to the branch.`,
			},
			"disabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `Determines if the branch rule is disabled or not.`,
			},
			"minimum_approvals_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: `The minimum number of approvals required for the branch rule to be matched.`,
			},
			"minimum_reviews_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: `The minimum number of reviews required for the branch rule to be matched.`,
			},
			"require_comments_resolved": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `Determines if require comments resolved before merging to the branch.`,
			},
			"require_linear_history": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `Determines if require linear history before merging to the branch.`,
			},
			"require_pull_request": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `Determines if the branch rule requires a pull request or not.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the BranchRule was created in UTC.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The resource name for the BranchRule.`,
			},
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Unique identifier of the BranchRule.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the BranchRule was updated in UTC.`,
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

func resourceSecureSourceManagerBranchRuleCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	includePatternProp, err := expandSecureSourceManagerBranchRuleIncludePattern(d.Get("include_pattern"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("include_pattern"); !tpgresource.IsEmptyValue(reflect.ValueOf(includePatternProp)) && (ok || !reflect.DeepEqual(v, includePatternProp)) {
		obj["includePattern"] = includePatternProp
	}
	disabledProp, err := expandSecureSourceManagerBranchRuleDisabled(d.Get("disabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(disabledProp)) && (ok || !reflect.DeepEqual(v, disabledProp)) {
		obj["disabled"] = disabledProp
	}
	requirePullRequestProp, err := expandSecureSourceManagerBranchRuleRequirePullRequest(d.Get("require_pull_request"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("require_pull_request"); !tpgresource.IsEmptyValue(reflect.ValueOf(requirePullRequestProp)) && (ok || !reflect.DeepEqual(v, requirePullRequestProp)) {
		obj["requirePullRequest"] = requirePullRequestProp
	}
	minimumReviewsCountProp, err := expandSecureSourceManagerBranchRuleMinimumReviewsCount(d.Get("minimum_reviews_count"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("minimum_reviews_count"); !tpgresource.IsEmptyValue(reflect.ValueOf(minimumReviewsCountProp)) && (ok || !reflect.DeepEqual(v, minimumReviewsCountProp)) {
		obj["minimumReviewsCount"] = minimumReviewsCountProp
	}
	minimumApprovalsCountProp, err := expandSecureSourceManagerBranchRuleMinimumApprovalsCount(d.Get("minimum_approvals_count"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("minimum_approvals_count"); !tpgresource.IsEmptyValue(reflect.ValueOf(minimumApprovalsCountProp)) && (ok || !reflect.DeepEqual(v, minimumApprovalsCountProp)) {
		obj["minimumApprovalsCount"] = minimumApprovalsCountProp
	}
	requireCommentsResolvedProp, err := expandSecureSourceManagerBranchRuleRequireCommentsResolved(d.Get("require_comments_resolved"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("require_comments_resolved"); !tpgresource.IsEmptyValue(reflect.ValueOf(requireCommentsResolvedProp)) && (ok || !reflect.DeepEqual(v, requireCommentsResolvedProp)) {
		obj["requireCommentsResolved"] = requireCommentsResolvedProp
	}
	allowStaleReviewsProp, err := expandSecureSourceManagerBranchRuleAllowStaleReviews(d.Get("allow_stale_reviews"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("allow_stale_reviews"); !tpgresource.IsEmptyValue(reflect.ValueOf(allowStaleReviewsProp)) && (ok || !reflect.DeepEqual(v, allowStaleReviewsProp)) {
		obj["allowStaleReviews"] = allowStaleReviewsProp
	}
	requireLinearHistoryProp, err := expandSecureSourceManagerBranchRuleRequireLinearHistory(d.Get("require_linear_history"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("require_linear_history"); !tpgresource.IsEmptyValue(reflect.ValueOf(requireLinearHistoryProp)) && (ok || !reflect.DeepEqual(v, requireLinearHistoryProp)) {
		obj["requireLinearHistory"] = requireLinearHistoryProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecureSourceManagerBasePath}}projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}/branchRules?branch_rule_id={{branch_rule_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new BranchRule: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for BranchRule: %s", err)
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
		return fmt.Errorf("Error creating BranchRule: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}/branchRules/{{branch_rule_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = SecureSourceManagerOperationWaitTime(
		config, res, project, "Creating BranchRule", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create BranchRule: %s", err)
	}

	log.Printf("[DEBUG] Finished creating BranchRule %q: %#v", d.Id(), res)

	return resourceSecureSourceManagerBranchRuleRead(d, meta)
}

func resourceSecureSourceManagerBranchRuleRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecureSourceManagerBasePath}}projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}/branchRules/{{branch_rule_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for BranchRule: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("SecureSourceManagerBranchRule %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading BranchRule: %s", err)
	}

	if err := d.Set("name", flattenSecureSourceManagerBranchRuleName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading BranchRule: %s", err)
	}
	if err := d.Set("uid", flattenSecureSourceManagerBranchRuleUid(res["uid"], d, config)); err != nil {
		return fmt.Errorf("Error reading BranchRule: %s", err)
	}
	if err := d.Set("create_time", flattenSecureSourceManagerBranchRuleCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading BranchRule: %s", err)
	}
	if err := d.Set("update_time", flattenSecureSourceManagerBranchRuleUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading BranchRule: %s", err)
	}
	if err := d.Set("include_pattern", flattenSecureSourceManagerBranchRuleIncludePattern(res["includePattern"], d, config)); err != nil {
		return fmt.Errorf("Error reading BranchRule: %s", err)
	}
	if err := d.Set("disabled", flattenSecureSourceManagerBranchRuleDisabled(res["disabled"], d, config)); err != nil {
		return fmt.Errorf("Error reading BranchRule: %s", err)
	}
	if err := d.Set("require_pull_request", flattenSecureSourceManagerBranchRuleRequirePullRequest(res["requirePullRequest"], d, config)); err != nil {
		return fmt.Errorf("Error reading BranchRule: %s", err)
	}
	if err := d.Set("minimum_reviews_count", flattenSecureSourceManagerBranchRuleMinimumReviewsCount(res["minimumReviewsCount"], d, config)); err != nil {
		return fmt.Errorf("Error reading BranchRule: %s", err)
	}
	if err := d.Set("minimum_approvals_count", flattenSecureSourceManagerBranchRuleMinimumApprovalsCount(res["minimumApprovalsCount"], d, config)); err != nil {
		return fmt.Errorf("Error reading BranchRule: %s", err)
	}
	if err := d.Set("require_comments_resolved", flattenSecureSourceManagerBranchRuleRequireCommentsResolved(res["requireCommentsResolved"], d, config)); err != nil {
		return fmt.Errorf("Error reading BranchRule: %s", err)
	}
	if err := d.Set("allow_stale_reviews", flattenSecureSourceManagerBranchRuleAllowStaleReviews(res["allowStaleReviews"], d, config)); err != nil {
		return fmt.Errorf("Error reading BranchRule: %s", err)
	}
	if err := d.Set("require_linear_history", flattenSecureSourceManagerBranchRuleRequireLinearHistory(res["requireLinearHistory"], d, config)); err != nil {
		return fmt.Errorf("Error reading BranchRule: %s", err)
	}

	return nil
}

func resourceSecureSourceManagerBranchRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for BranchRule: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	includePatternProp, err := expandSecureSourceManagerBranchRuleIncludePattern(d.Get("include_pattern"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("include_pattern"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, includePatternProp)) {
		obj["includePattern"] = includePatternProp
	}
	disabledProp, err := expandSecureSourceManagerBranchRuleDisabled(d.Get("disabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, disabledProp)) {
		obj["disabled"] = disabledProp
	}
	requirePullRequestProp, err := expandSecureSourceManagerBranchRuleRequirePullRequest(d.Get("require_pull_request"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("require_pull_request"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, requirePullRequestProp)) {
		obj["requirePullRequest"] = requirePullRequestProp
	}
	minimumReviewsCountProp, err := expandSecureSourceManagerBranchRuleMinimumReviewsCount(d.Get("minimum_reviews_count"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("minimum_reviews_count"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, minimumReviewsCountProp)) {
		obj["minimumReviewsCount"] = minimumReviewsCountProp
	}
	minimumApprovalsCountProp, err := expandSecureSourceManagerBranchRuleMinimumApprovalsCount(d.Get("minimum_approvals_count"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("minimum_approvals_count"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, minimumApprovalsCountProp)) {
		obj["minimumApprovalsCount"] = minimumApprovalsCountProp
	}
	requireCommentsResolvedProp, err := expandSecureSourceManagerBranchRuleRequireCommentsResolved(d.Get("require_comments_resolved"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("require_comments_resolved"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, requireCommentsResolvedProp)) {
		obj["requireCommentsResolved"] = requireCommentsResolvedProp
	}
	allowStaleReviewsProp, err := expandSecureSourceManagerBranchRuleAllowStaleReviews(d.Get("allow_stale_reviews"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("allow_stale_reviews"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, allowStaleReviewsProp)) {
		obj["allowStaleReviews"] = allowStaleReviewsProp
	}
	requireLinearHistoryProp, err := expandSecureSourceManagerBranchRuleRequireLinearHistory(d.Get("require_linear_history"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("require_linear_history"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, requireLinearHistoryProp)) {
		obj["requireLinearHistory"] = requireLinearHistoryProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecureSourceManagerBasePath}}projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}/branchRules/{{branch_rule_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating BranchRule %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("include_pattern") {
		updateMask = append(updateMask, "includePattern")
	}

	if d.HasChange("disabled") {
		updateMask = append(updateMask, "disabled")
	}

	if d.HasChange("require_pull_request") {
		updateMask = append(updateMask, "requirePullRequest")
	}

	if d.HasChange("minimum_reviews_count") {
		updateMask = append(updateMask, "minimumReviewsCount")
	}

	if d.HasChange("minimum_approvals_count") {
		updateMask = append(updateMask, "minimumApprovalsCount")
	}

	if d.HasChange("require_comments_resolved") {
		updateMask = append(updateMask, "requireCommentsResolved")
	}

	if d.HasChange("allow_stale_reviews") {
		updateMask = append(updateMask, "allowStaleReviews")
	}

	if d.HasChange("require_linear_history") {
		updateMask = append(updateMask, "requireLinearHistory")
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
			return fmt.Errorf("Error updating BranchRule %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating BranchRule %q: %#v", d.Id(), res)
		}

	}

	return resourceSecureSourceManagerBranchRuleRead(d, meta)
}

func resourceSecureSourceManagerBranchRuleDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for BranchRule: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{SecureSourceManagerBasePath}}projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}/branchRules/{{branch_rule_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting BranchRule %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "BranchRule")
	}

	err = SecureSourceManagerOperationWaitTime(
		config, res, project, "Deleting BranchRule", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting BranchRule %q: %#v", d.Id(), res)
	return nil
}

func resourceSecureSourceManagerBranchRuleImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/repositories/(?P<repository_id>[^/]+)/branchRules/(?P<branch_rule_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<repository_id>[^/]+)/(?P<branch_rule_id>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<repository_id>[^/]+)/(?P<branch_rule_id>[^/]+)$",
		"^(?P<branch_rule_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}/branchRules/{{branch_rule_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenSecureSourceManagerBranchRuleName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecureSourceManagerBranchRuleUid(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecureSourceManagerBranchRuleCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecureSourceManagerBranchRuleUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecureSourceManagerBranchRuleIncludePattern(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecureSourceManagerBranchRuleDisabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecureSourceManagerBranchRuleRequirePullRequest(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecureSourceManagerBranchRuleMinimumReviewsCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenSecureSourceManagerBranchRuleMinimumApprovalsCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenSecureSourceManagerBranchRuleRequireCommentsResolved(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecureSourceManagerBranchRuleAllowStaleReviews(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecureSourceManagerBranchRuleRequireLinearHistory(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandSecureSourceManagerBranchRuleIncludePattern(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecureSourceManagerBranchRuleDisabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecureSourceManagerBranchRuleRequirePullRequest(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecureSourceManagerBranchRuleMinimumReviewsCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecureSourceManagerBranchRuleMinimumApprovalsCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecureSourceManagerBranchRuleRequireCommentsResolved(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecureSourceManagerBranchRuleAllowStaleReviews(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecureSourceManagerBranchRuleRequireLinearHistory(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
