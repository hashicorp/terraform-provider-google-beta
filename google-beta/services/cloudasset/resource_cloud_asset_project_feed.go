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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/cloudasset/ProjectFeed.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package cloudasset

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
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceCloudAssetProjectFeed() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudAssetProjectFeedCreate,
		Read:   resourceCloudAssetProjectFeedRead,
		Update: resourceCloudAssetProjectFeedUpdate,
		Delete: resourceCloudAssetProjectFeedDelete,

		Importer: &schema.ResourceImporter{
			State: resourceCloudAssetProjectFeedImport,
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
			"feed_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `This is the client-assigned asset feed identifier and it needs to be unique under a specific parent.`,
			},
			"feed_output_config": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `Output configuration for asset feed destination.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pubsub_destination": {
							Type:        schema.TypeList,
							Required:    true,
							Description: `Destination on Cloud Pubsub.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"topic": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `Destination on Cloud Pubsub topic.`,
									},
								},
							},
						},
					},
				},
			},
			"asset_names": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `A list of the full names of the assets to receive updates. You must specify either or both of
assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"asset_types": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `A list of types of the assets to receive updates. You must specify either or both of assetNames
and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
the feed. For example: "compute.googleapis.com/Disk"
See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
supported asset types.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"billing_project": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The project whose identity will be used when sending messages to the
destination pubsub topic. It also specifies the project for API
enablement check, quota, and billing. If not specified, the resource's
project will be used.`,
			},
			"condition": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `A condition which determines whether an asset update should be published. If specified, an asset
will be returned only when the expression evaluates to true. When set, expression field
must be a valid CEL expression on a TemporalAsset with name temporal_asset. Example: a Feed with
expression "temporal_asset.deleted == true" will only publish Asset deletions. Other fields of
condition are optional.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"expression": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `Textual representation of an expression in Common Expression Language syntax.`,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `Description of the expression. This is a longer text which describes the expression,
e.g. when hovered over it in a UI.`,
						},
						"location": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `String indicating the location of the expression for error reporting, e.g. a file
name and a position in the file.`,
						},
						"title": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `Title for the expression, i.e. a short string describing its purpose.
This can be used e.g. in UIs which allow to enter the expression.`,
						},
					},
				},
			},
			"content_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: verify.ValidateEnum([]string{"CONTENT_TYPE_UNSPECIFIED", "RESOURCE", "IAM_POLICY", "ORG_POLICY", "OS_INVENTORY", "ACCESS_POLICY", ""}),
				Description:  `Asset content type. If not specified, no content but the asset name and type will be returned. Possible values: ["CONTENT_TYPE_UNSPECIFIED", "RESOURCE", "IAM_POLICY", "ORG_POLICY", "OS_INVENTORY", "ACCESS_POLICY"]`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The format will be projects/{projectNumber}/feeds/{client-assigned_feed_identifier}.`,
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

func resourceCloudAssetProjectFeedCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	assetNamesProp, err := expandCloudAssetProjectFeedAssetNames(d.Get("asset_names"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("asset_names"); !tpgresource.IsEmptyValue(reflect.ValueOf(assetNamesProp)) && (ok || !reflect.DeepEqual(v, assetNamesProp)) {
		obj["assetNames"] = assetNamesProp
	}
	assetTypesProp, err := expandCloudAssetProjectFeedAssetTypes(d.Get("asset_types"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("asset_types"); !tpgresource.IsEmptyValue(reflect.ValueOf(assetTypesProp)) && (ok || !reflect.DeepEqual(v, assetTypesProp)) {
		obj["assetTypes"] = assetTypesProp
	}
	contentTypeProp, err := expandCloudAssetProjectFeedContentType(d.Get("content_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("content_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(contentTypeProp)) && (ok || !reflect.DeepEqual(v, contentTypeProp)) {
		obj["contentType"] = contentTypeProp
	}
	feedOutputConfigProp, err := expandCloudAssetProjectFeedFeedOutputConfig(d.Get("feed_output_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("feed_output_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(feedOutputConfigProp)) && (ok || !reflect.DeepEqual(v, feedOutputConfigProp)) {
		obj["feedOutputConfig"] = feedOutputConfigProp
	}
	conditionProp, err := expandCloudAssetProjectFeedCondition(d.Get("condition"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("condition"); !tpgresource.IsEmptyValue(reflect.ValueOf(conditionProp)) && (ok || !reflect.DeepEqual(v, conditionProp)) {
		obj["condition"] = conditionProp
	}

	obj, err = resourceCloudAssetProjectFeedEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{CloudAssetBasePath}}projects/{{project}}/feeds?feedId={{feed_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ProjectFeed: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ProjectFeed: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	// Send the project ID in the X-Goog-User-Project header.
	origUserProjectOverride := config.UserProjectOverride
	config.UserProjectOverride = true
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
		return fmt.Errorf("Error creating ProjectFeed: %s", err)
	}
	if err := d.Set("name", flattenCloudAssetProjectFeedName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Restore the original value of user_project_override.
	config.UserProjectOverride = origUserProjectOverride

	log.Printf("[DEBUG] Finished creating ProjectFeed %q: %#v", d.Id(), res)

	return resourceCloudAssetProjectFeedRead(d, meta)
}

func resourceCloudAssetProjectFeedRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{CloudAssetBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ProjectFeed: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("CloudAssetProjectFeed %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ProjectFeed: %s", err)
	}

	if err := d.Set("name", flattenCloudAssetProjectFeedName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading ProjectFeed: %s", err)
	}
	if err := d.Set("asset_names", flattenCloudAssetProjectFeedAssetNames(res["assetNames"], d, config)); err != nil {
		return fmt.Errorf("Error reading ProjectFeed: %s", err)
	}
	if err := d.Set("asset_types", flattenCloudAssetProjectFeedAssetTypes(res["assetTypes"], d, config)); err != nil {
		return fmt.Errorf("Error reading ProjectFeed: %s", err)
	}
	if err := d.Set("content_type", flattenCloudAssetProjectFeedContentType(res["contentType"], d, config)); err != nil {
		return fmt.Errorf("Error reading ProjectFeed: %s", err)
	}
	if err := d.Set("feed_output_config", flattenCloudAssetProjectFeedFeedOutputConfig(res["feedOutputConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading ProjectFeed: %s", err)
	}
	if err := d.Set("condition", flattenCloudAssetProjectFeedCondition(res["condition"], d, config)); err != nil {
		return fmt.Errorf("Error reading ProjectFeed: %s", err)
	}

	return nil
}

func resourceCloudAssetProjectFeedUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ProjectFeed: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	assetNamesProp, err := expandCloudAssetProjectFeedAssetNames(d.Get("asset_names"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("asset_names"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, assetNamesProp)) {
		obj["assetNames"] = assetNamesProp
	}
	assetTypesProp, err := expandCloudAssetProjectFeedAssetTypes(d.Get("asset_types"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("asset_types"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, assetTypesProp)) {
		obj["assetTypes"] = assetTypesProp
	}
	contentTypeProp, err := expandCloudAssetProjectFeedContentType(d.Get("content_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("content_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, contentTypeProp)) {
		obj["contentType"] = contentTypeProp
	}
	feedOutputConfigProp, err := expandCloudAssetProjectFeedFeedOutputConfig(d.Get("feed_output_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("feed_output_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, feedOutputConfigProp)) {
		obj["feedOutputConfig"] = feedOutputConfigProp
	}
	conditionProp, err := expandCloudAssetProjectFeedCondition(d.Get("condition"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("condition"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, conditionProp)) {
		obj["condition"] = conditionProp
	}

	obj, err = resourceCloudAssetProjectFeedEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{CloudAssetBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ProjectFeed %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("asset_names") {
		updateMask = append(updateMask, "assetNames")
	}

	if d.HasChange("asset_types") {
		updateMask = append(updateMask, "assetTypes")
	}

	if d.HasChange("content_type") {
		updateMask = append(updateMask, "contentType")
	}

	if d.HasChange("feed_output_config") {
		updateMask = append(updateMask, "feedOutputConfig")
	}

	if d.HasChange("condition") {
		updateMask = append(updateMask, "condition")
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
			return fmt.Errorf("Error updating ProjectFeed %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating ProjectFeed %q: %#v", d.Id(), res)
		}

	}

	return resourceCloudAssetProjectFeedRead(d, meta)
}

func resourceCloudAssetProjectFeedDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ProjectFeed: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{CloudAssetBasePath}}{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting ProjectFeed %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "ProjectFeed")
	}

	log.Printf("[DEBUG] Finished deleting ProjectFeed %q: %#v", d.Id(), res)
	return nil
}

func resourceCloudAssetProjectFeedImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	if err := d.Set("name", d.Id()); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

func flattenCloudAssetProjectFeedName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudAssetProjectFeedAssetNames(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudAssetProjectFeedAssetTypes(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudAssetProjectFeedContentType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudAssetProjectFeedFeedOutputConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["pubsub_destination"] =
		flattenCloudAssetProjectFeedFeedOutputConfigPubsubDestination(original["pubsubDestination"], d, config)
	return []interface{}{transformed}
}
func flattenCloudAssetProjectFeedFeedOutputConfigPubsubDestination(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["topic"] =
		flattenCloudAssetProjectFeedFeedOutputConfigPubsubDestinationTopic(original["topic"], d, config)
	return []interface{}{transformed}
}
func flattenCloudAssetProjectFeedFeedOutputConfigPubsubDestinationTopic(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudAssetProjectFeedCondition(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["expression"] =
		flattenCloudAssetProjectFeedConditionExpression(original["expression"], d, config)
	transformed["title"] =
		flattenCloudAssetProjectFeedConditionTitle(original["title"], d, config)
	transformed["description"] =
		flattenCloudAssetProjectFeedConditionDescription(original["description"], d, config)
	transformed["location"] =
		flattenCloudAssetProjectFeedConditionLocation(original["location"], d, config)
	return []interface{}{transformed}
}
func flattenCloudAssetProjectFeedConditionExpression(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudAssetProjectFeedConditionTitle(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudAssetProjectFeedConditionDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudAssetProjectFeedConditionLocation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandCloudAssetProjectFeedAssetNames(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetProjectFeedAssetTypes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetProjectFeedContentType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetProjectFeedFeedOutputConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPubsubDestination, err := expandCloudAssetProjectFeedFeedOutputConfigPubsubDestination(original["pubsub_destination"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPubsubDestination); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["pubsubDestination"] = transformedPubsubDestination
	}

	return transformed, nil
}

func expandCloudAssetProjectFeedFeedOutputConfigPubsubDestination(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTopic, err := expandCloudAssetProjectFeedFeedOutputConfigPubsubDestinationTopic(original["topic"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTopic); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["topic"] = transformedTopic
	}

	return transformed, nil
}

func expandCloudAssetProjectFeedFeedOutputConfigPubsubDestinationTopic(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetProjectFeedCondition(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedExpression, err := expandCloudAssetProjectFeedConditionExpression(original["expression"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExpression); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["expression"] = transformedExpression
	}

	transformedTitle, err := expandCloudAssetProjectFeedConditionTitle(original["title"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTitle); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["title"] = transformedTitle
	}

	transformedDescription, err := expandCloudAssetProjectFeedConditionDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	transformedLocation, err := expandCloudAssetProjectFeedConditionLocation(original["location"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["location"] = transformedLocation
	}

	return transformed, nil
}

func expandCloudAssetProjectFeedConditionExpression(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetProjectFeedConditionTitle(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetProjectFeedConditionDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetProjectFeedConditionLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceCloudAssetProjectFeedEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// Remove the "folders/" prefix from the folder ID
	if folder, ok := d.GetOkExists("folder"); ok {
		if err := d.Set("folder_id", strings.TrimPrefix(folder.(string), "folders/")); err != nil {
			return nil, fmt.Errorf("Error setting folder_id: %s", err)
		}
	}
	// The feed object must be under the "feed" attribute on the request.
	newObj := make(map[string]interface{})
	newObj["feed"] = obj
	return newObj, nil
}
