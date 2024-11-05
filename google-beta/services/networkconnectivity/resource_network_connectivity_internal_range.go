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

package networkconnectivity

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

func ResourceNetworkConnectivityInternalRange() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkConnectivityInternalRangeCreate,
		Read:   resourceNetworkConnectivityInternalRangeRead,
		Update: resourceNetworkConnectivityInternalRangeUpdate,
		Delete: resourceNetworkConnectivityInternalRangeDelete,

		Importer: &schema.ResourceImporter{
			State: resourceNetworkConnectivityInternalRangeImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the policy based route.`,
			},
			"network": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `Fully-qualified URL of the network that this route applies to, for example: projects/my-project/global/networks/my-network.`,
			},
			"peering": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: verify.ValidateEnum([]string{"FOR_SELF", "FOR_PEER", "NOT_SHARED"}),
				Description:  `The type of peering set for this internal range. Possible values: ["FOR_SELF", "FOR_PEER", "NOT_SHARED"]`,
			},
			"usage": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: verify.ValidateEnum([]string{"FOR_VPC", "EXTERNAL_TO_VPC", "FOR_MIGRATION"}),
				Description:  `The type of usage set for this InternalRange. Possible values: ["FOR_VPC", "EXTERNAL_TO_VPC", "FOR_MIGRATION"]`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `An optional description of this resource.`,
			},
			"ip_cidr_range": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: `The IP range that this internal range defines.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `User-defined labels.


**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"migration": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `Specification for migration with source and target resource names.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source": {
							Type:     schema.TypeString,
							Required: true,
							Description: `Resource path as an URI of the source resource, for example a subnet.
The project for the source resource should match the project for the
InternalRange.
An example /projects/{project}/regions/{region}/subnetworks/{subnet}`,
						},
						"target": {
							Type:     schema.TypeString,
							Required: true,
							Description: `Resource path of the target resource. The target project can be
different, as in the cases when migrating to peer networks. The resource
may not exist yet.
For example /projects/{project}/regions/{region}/subnetworks/{subnet}`,
						},
					},
				},
			},
			"overlaps": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Optional. Types of resources that are allowed to overlap with the current internal range. Possible values: ["OVERLAP_ROUTE_RANGE", "OVERLAP_EXISTING_SUBNET_RANGE"]`,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: verify.ValidateEnum([]string{"OVERLAP_ROUTE_RANGE", "OVERLAP_EXISTING_SUBNET_RANGE"}),
				},
			},
			"prefix_length": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: `An alternate to ipCidrRange. Can be set when trying to create a reservation that automatically finds a free range of the given size.
If both ipCidrRange and prefixLength are set, there is an error if the range sizes do not match. Can also be used during updates to change the range size.`,
			},
			"target_cidr_range": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `Optional. Can be set to narrow down or pick a different address space while searching for a free range.
If not set, defaults to the "10.0.0.0/8" address space. This can be used to search in other rfc-1918 address spaces like "172.16.0.0/12" and "192.168.0.0/16" or non-rfc-1918 address spaces used in the VPC.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"users": {
				Type:     schema.TypeList,
				Computed: true,
				Description: `Output only. The list of resources that refer to this internal range.
Resources that use the internal range for their range allocation are referred to as users of the range.
Other resources mark themselves as users while doing so by creating a reference to this internal range. Having a user, based on this reference, prevents deletion of the internal range referred to. Can be empty.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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

func resourceNetworkConnectivityInternalRangeCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkConnectivityInternalRangeDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	ipCidrRangeProp, err := expandNetworkConnectivityInternalRangeIpCidrRange(d.Get("ip_cidr_range"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ip_cidr_range"); !tpgresource.IsEmptyValue(reflect.ValueOf(ipCidrRangeProp)) && (ok || !reflect.DeepEqual(v, ipCidrRangeProp)) {
		obj["ipCidrRange"] = ipCidrRangeProp
	}
	networkProp, err := expandNetworkConnectivityInternalRangeNetwork(d.Get("network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	usageProp, err := expandNetworkConnectivityInternalRangeUsage(d.Get("usage"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("usage"); !tpgresource.IsEmptyValue(reflect.ValueOf(usageProp)) && (ok || !reflect.DeepEqual(v, usageProp)) {
		obj["usage"] = usageProp
	}
	peeringProp, err := expandNetworkConnectivityInternalRangePeering(d.Get("peering"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("peering"); !tpgresource.IsEmptyValue(reflect.ValueOf(peeringProp)) && (ok || !reflect.DeepEqual(v, peeringProp)) {
		obj["peering"] = peeringProp
	}
	prefixLengthProp, err := expandNetworkConnectivityInternalRangePrefixLength(d.Get("prefix_length"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("prefix_length"); !tpgresource.IsEmptyValue(reflect.ValueOf(prefixLengthProp)) && (ok || !reflect.DeepEqual(v, prefixLengthProp)) {
		obj["prefixLength"] = prefixLengthProp
	}
	targetCidrRangeProp, err := expandNetworkConnectivityInternalRangeTargetCidrRange(d.Get("target_cidr_range"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("target_cidr_range"); !tpgresource.IsEmptyValue(reflect.ValueOf(targetCidrRangeProp)) && (ok || !reflect.DeepEqual(v, targetCidrRangeProp)) {
		obj["targetCidrRange"] = targetCidrRangeProp
	}
	overlapsProp, err := expandNetworkConnectivityInternalRangeOverlaps(d.Get("overlaps"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("overlaps"); !tpgresource.IsEmptyValue(reflect.ValueOf(overlapsProp)) && (ok || !reflect.DeepEqual(v, overlapsProp)) {
		obj["overlaps"] = overlapsProp
	}
	migrationProp, err := expandNetworkConnectivityInternalRangeMigration(d.Get("migration"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("migration"); !tpgresource.IsEmptyValue(reflect.ValueOf(migrationProp)) && (ok || !reflect.DeepEqual(v, migrationProp)) {
		obj["migration"] = migrationProp
	}
	labelsProp, err := expandNetworkConnectivityInternalRangeEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkConnectivityBasePath}}projects/{{project}}/locations/global/internalRanges?internalRangeId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new InternalRange: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for InternalRange: %s", err)
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
		return fmt.Errorf("Error creating InternalRange: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/global/internalRanges/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = NetworkConnectivityOperationWaitTime(
		config, res, project, "Creating InternalRange", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create InternalRange: %s", err)
	}

	log.Printf("[DEBUG] Finished creating InternalRange %q: %#v", d.Id(), res)

	return resourceNetworkConnectivityInternalRangeRead(d, meta)
}

func resourceNetworkConnectivityInternalRangeRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkConnectivityBasePath}}projects/{{project}}/locations/global/internalRanges/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for InternalRange: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("NetworkConnectivityInternalRange %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading InternalRange: %s", err)
	}

	if err := d.Set("labels", flattenNetworkConnectivityInternalRangeLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading InternalRange: %s", err)
	}
	if err := d.Set("description", flattenNetworkConnectivityInternalRangeDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading InternalRange: %s", err)
	}
	if err := d.Set("ip_cidr_range", flattenNetworkConnectivityInternalRangeIpCidrRange(res["ipCidrRange"], d, config)); err != nil {
		return fmt.Errorf("Error reading InternalRange: %s", err)
	}
	if err := d.Set("network", flattenNetworkConnectivityInternalRangeNetwork(res["network"], d, config)); err != nil {
		return fmt.Errorf("Error reading InternalRange: %s", err)
	}
	if err := d.Set("usage", flattenNetworkConnectivityInternalRangeUsage(res["usage"], d, config)); err != nil {
		return fmt.Errorf("Error reading InternalRange: %s", err)
	}
	if err := d.Set("peering", flattenNetworkConnectivityInternalRangePeering(res["peering"], d, config)); err != nil {
		return fmt.Errorf("Error reading InternalRange: %s", err)
	}
	if err := d.Set("prefix_length", flattenNetworkConnectivityInternalRangePrefixLength(res["prefixLength"], d, config)); err != nil {
		return fmt.Errorf("Error reading InternalRange: %s", err)
	}
	if err := d.Set("target_cidr_range", flattenNetworkConnectivityInternalRangeTargetCidrRange(res["targetCidrRange"], d, config)); err != nil {
		return fmt.Errorf("Error reading InternalRange: %s", err)
	}
	if err := d.Set("users", flattenNetworkConnectivityInternalRangeUsers(res["users"], d, config)); err != nil {
		return fmt.Errorf("Error reading InternalRange: %s", err)
	}
	if err := d.Set("overlaps", flattenNetworkConnectivityInternalRangeOverlaps(res["overlaps"], d, config)); err != nil {
		return fmt.Errorf("Error reading InternalRange: %s", err)
	}
	if err := d.Set("migration", flattenNetworkConnectivityInternalRangeMigration(res["migration"], d, config)); err != nil {
		return fmt.Errorf("Error reading InternalRange: %s", err)
	}
	if err := d.Set("terraform_labels", flattenNetworkConnectivityInternalRangeTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading InternalRange: %s", err)
	}
	if err := d.Set("effective_labels", flattenNetworkConnectivityInternalRangeEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading InternalRange: %s", err)
	}

	return nil
}

func resourceNetworkConnectivityInternalRangeUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for InternalRange: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkConnectivityInternalRangeDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	ipCidrRangeProp, err := expandNetworkConnectivityInternalRangeIpCidrRange(d.Get("ip_cidr_range"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ip_cidr_range"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, ipCidrRangeProp)) {
		obj["ipCidrRange"] = ipCidrRangeProp
	}
	networkProp, err := expandNetworkConnectivityInternalRangeNetwork(d.Get("network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	usageProp, err := expandNetworkConnectivityInternalRangeUsage(d.Get("usage"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("usage"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, usageProp)) {
		obj["usage"] = usageProp
	}
	peeringProp, err := expandNetworkConnectivityInternalRangePeering(d.Get("peering"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("peering"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, peeringProp)) {
		obj["peering"] = peeringProp
	}
	prefixLengthProp, err := expandNetworkConnectivityInternalRangePrefixLength(d.Get("prefix_length"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("prefix_length"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, prefixLengthProp)) {
		obj["prefixLength"] = prefixLengthProp
	}
	targetCidrRangeProp, err := expandNetworkConnectivityInternalRangeTargetCidrRange(d.Get("target_cidr_range"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("target_cidr_range"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, targetCidrRangeProp)) {
		obj["targetCidrRange"] = targetCidrRangeProp
	}
	overlapsProp, err := expandNetworkConnectivityInternalRangeOverlaps(d.Get("overlaps"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("overlaps"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, overlapsProp)) {
		obj["overlaps"] = overlapsProp
	}
	labelsProp, err := expandNetworkConnectivityInternalRangeEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkConnectivityBasePath}}projects/{{project}}/locations/global/internalRanges/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating InternalRange %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("ip_cidr_range") {
		updateMask = append(updateMask, "ipCidrRange")
	}

	if d.HasChange("network") {
		updateMask = append(updateMask, "network")
	}

	if d.HasChange("usage") {
		updateMask = append(updateMask, "usage")
	}

	if d.HasChange("peering") {
		updateMask = append(updateMask, "peering")
	}

	if d.HasChange("prefix_length") {
		updateMask = append(updateMask, "prefixLength")
	}

	if d.HasChange("target_cidr_range") {
		updateMask = append(updateMask, "targetCidrRange")
	}

	if d.HasChange("overlaps") {
		updateMask = append(updateMask, "overlaps")
	}

	if d.HasChange("effective_labels") {
		updateMask = append(updateMask, "labels")
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
			return fmt.Errorf("Error updating InternalRange %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating InternalRange %q: %#v", d.Id(), res)
		}

		err = NetworkConnectivityOperationWaitTime(
			config, res, project, "Updating InternalRange", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceNetworkConnectivityInternalRangeRead(d, meta)
}

func resourceNetworkConnectivityInternalRangeDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for InternalRange: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkConnectivityBasePath}}projects/{{project}}/locations/global/internalRanges/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting InternalRange %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "InternalRange")
	}

	err = NetworkConnectivityOperationWaitTime(
		config, res, project, "Deleting InternalRange", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting InternalRange %q: %#v", d.Id(), res)
	return nil
}

func resourceNetworkConnectivityInternalRangeImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/global/internalRanges/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/global/internalRanges/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNetworkConnectivityInternalRangeLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenNetworkConnectivityInternalRangeDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityInternalRangeIpCidrRange(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityInternalRangeNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func flattenNetworkConnectivityInternalRangeUsage(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityInternalRangePeering(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityInternalRangePrefixLength(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenNetworkConnectivityInternalRangeTargetCidrRange(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityInternalRangeUsers(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityInternalRangeOverlaps(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityInternalRangeMigration(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["source"] =
		flattenNetworkConnectivityInternalRangeMigrationSource(original["source"], d, config)
	transformed["target"] =
		flattenNetworkConnectivityInternalRangeMigrationTarget(original["target"], d, config)
	return []interface{}{transformed}
}
func flattenNetworkConnectivityInternalRangeMigrationSource(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityInternalRangeMigrationTarget(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityInternalRangeTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("terraform_labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenNetworkConnectivityInternalRangeEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNetworkConnectivityInternalRangeDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityInternalRangeIpCidrRange(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityInternalRangeNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityInternalRangeUsage(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityInternalRangePeering(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityInternalRangePrefixLength(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityInternalRangeTargetCidrRange(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityInternalRangeOverlaps(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityInternalRangeMigration(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSource, err := expandNetworkConnectivityInternalRangeMigrationSource(original["source"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSource); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["source"] = transformedSource
	}

	transformedTarget, err := expandNetworkConnectivityInternalRangeMigrationTarget(original["target"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTarget); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["target"] = transformedTarget
	}

	return transformed, nil
}

func expandNetworkConnectivityInternalRangeMigrationSource(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityInternalRangeMigrationTarget(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityInternalRangeEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
