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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/compute/RouterNatAddress.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package compute

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func addressResourceNameSetFromSelfLinkSet(v interface{}) *schema.Set {
	if v == nil {
		return schema.NewSet(schema.HashString, nil)
	}
	vSet := v.(*schema.Set)
	ls := make([]interface{}, 0, vSet.Len())
	for _, v := range vSet.List() {
		if v == nil {
			continue
		}
		ls = append(ls, tpgresource.GetResourceNameFromSelfLink(v.(string)))
	}
	return schema.NewSet(schema.HashString, ls)
}

// drain_nat_ips MUST be set from (just set) previous values of nat_ips
// so this customizeDiff func makes sure drainNatIps values:
//   - aren't set at creation time
//   - are in old value of nat_ips but not in new values
func resourceComputeRouterNatAddressDrainNatIpsCustomDiff(_ context.Context, diff *schema.ResourceDiff, meta interface{}) error {
	o, n := diff.GetChange("drain_nat_ips")
	oSet := addressResourceNameSetFromSelfLinkSet(o)
	nSet := addressResourceNameSetFromSelfLinkSet(n)
	addDrainIps := nSet.Difference(oSet)

	// We don't care if there are no new drainNatIps
	if addDrainIps.Len() == 0 {
		return nil
	}

	// Resource hasn't been created yet - return error
	if diff.Id() == "" {
		return fmt.Errorf("New RouterNat cannot have drain_nat_ips, got values %+v", addDrainIps.List())
	}
	//
	o, n = diff.GetChange("nat_ips")
	oNatSet := addressResourceNameSetFromSelfLinkSet(o)
	nNatSet := addressResourceNameSetFromSelfLinkSet(n)

	// Resource is being updated - make sure new drainNatIps were in natIps prior d and no longer are in natIps.
	for _, v := range addDrainIps.List() {
		if !oNatSet.Contains(v) {
			return fmt.Errorf("drain_nat_ip %q was not previously set in nat_ips %+v", v.(string), oNatSet.List())
		}
		if nNatSet.Contains(v) {
			return fmt.Errorf("drain_nat_ip %q cannot be drained if still set in nat_ips %+v", v.(string), nNatSet.List())
		}
	}
	return nil
}

func resourceComputeRouterNatAddressDeleteOnlyNatIps(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	items, err := resourceComputeRouterNatAddressListForPatch(d, meta)
	if err != nil {
		return nil, err
	}

	idx, item, err := resourceComputeRouterNatAddressFindNestedObjectInList(d, meta, items)
	if err != nil {
		return nil, err
	}

	// Return error if item to update does not exist.
	if item == nil {
		return nil, fmt.Errorf("Unable to update RouterNatAddress %q - not found in list", d.Id())
	}

	if item["natIps"] != nil {
		croppedNatIps := item["natIps"].([]interface{})[:1]
		item["natIps"] = croppedNatIps
	}

	items[idx] = item
	// Return list with new item added
	res := map[string]interface{}{
		"nats": items,
	}
	return res, nil
}

func ResourceComputeRouterNatAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeRouterNatAddressCreate,
		Read:   resourceComputeRouterNatAddressRead,
		Update: resourceComputeRouterNatAddressUpdate,
		Delete: resourceComputeRouterNatAddressDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeRouterNatAddressImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			resourceComputeRouterNatAddressDrainNatIpsCustomDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"nat_ips": {
				Type:     schema.TypeSet,
				Required: true,
				Description: `Self-links of NAT IPs to be used in a Nat service. Only valid if the referenced RouterNat
natIpAllocateOption is set to MANUAL_ONLY.`,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				},
				Set: computeRouterNatIPsHash,
			},
			"router": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The name of the Cloud Router in which the referenced NAT service is configured.`,
			},
			"router_nat": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The name of the Nat service in which this address will be configured.`,
			},
			"drain_nat_ips": {
				Type:     schema.TypeSet,
				Optional: true,
				Description: `A list of URLs of the IP resources to be drained. These IPs must be
valid static external IPs that have been assigned to the NAT.`,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				},
				// Default schema.HashSchema is used.
			},
			"region": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `Region where the NAT service reside.`,
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

func resourceComputeRouterNatAddressCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	// A custom_create function similar to the generated code when using a nested_query, but replaces the encoder with a custom one instead of just injecting it;
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	natIpsProp, err := expandNestedComputeRouterNatAddressNatIps(d.Get("nat_ips"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("nat_ips"); ok || !reflect.DeepEqual(v, natIpsProp) {
		obj["natIps"] = natIpsProp
	}
	drainNatIpsProp, err := expandNestedComputeRouterNatAddressDrainNatIps(d.Get("drain_nat_ips"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("drain_nat_ips"); ok || !reflect.DeepEqual(v, drainNatIpsProp) {
		obj["drainNatIps"] = drainNatIpsProp
	}
	nameProp, err := expandNestedComputeRouterNatAddressRouterNat(d.Get("router_nat"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("router_nat"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	log.Printf("[DEBUG] Creating new RouterNatAddress: %#v", obj)

	obj, err = resourceComputeRouterNatAddressEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	lockName, err := tpgresource.ReplaceVars(d, config, "router/{{region}}/{{router}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/routers/{{router}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RouterNatAddress: %s", err)
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
		return fmt.Errorf("Error creating RouterNatAddress: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/regions/{{region}}/routers/{{router}}/{{router_nat}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ComputeOperationWaitTime(
		config, res, project, "Creating RouterNatAddress", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create RouterNatAddress: %s", err)
	}

	log.Printf("[DEBUG] Finished creating RouterNatAddress %q: %#v", d.Id(), res)

	return resourceComputeRouterNatAddressRead(d, meta)
}

func resourceComputeRouterNatAddressRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/routers/{{router}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RouterNatAddress: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ComputeRouterNatAddress %q", d.Id()))
	}

	res, err = flattenNestedComputeRouterNatAddress(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Object isn't there any more - remove it from the state.
		log.Printf("[DEBUG] Removing ComputeRouterNatAddress because it couldn't be matched.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading RouterNatAddress: %s", err)
	}

	if err := d.Set("nat_ips", flattenNestedComputeRouterNatAddressNatIps(res["natIps"], d, config)); err != nil {
		return fmt.Errorf("Error reading RouterNatAddress: %s", err)
	}
	if err := d.Set("drain_nat_ips", flattenNestedComputeRouterNatAddressDrainNatIps(res["drainNatIps"], d, config)); err != nil {
		return fmt.Errorf("Error reading RouterNatAddress: %s", err)
	}
	if err := d.Set("router_nat", flattenNestedComputeRouterNatAddressRouterNat(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading RouterNatAddress: %s", err)
	}

	return nil
}

func resourceComputeRouterNatAddressUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RouterNatAddress: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	natIpsProp, err := expandNestedComputeRouterNatAddressNatIps(d.Get("nat_ips"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("nat_ips"); ok || !reflect.DeepEqual(v, natIpsProp) {
		obj["natIps"] = natIpsProp
	}
	drainNatIpsProp, err := expandNestedComputeRouterNatAddressDrainNatIps(d.Get("drain_nat_ips"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("drain_nat_ips"); ok || !reflect.DeepEqual(v, drainNatIpsProp) {
		obj["drainNatIps"] = drainNatIpsProp
	}

	obj, err = resourceComputeRouterNatAddressUpdateEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	lockName, err := tpgresource.ReplaceVars(d, config, "router/{{region}}/{{router}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/routers/{{router}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating RouterNatAddress %q: %#v", d.Id(), obj)
	headers := make(http.Header)

	obj, err = resourceComputeRouterNatAddressPatchUpdateEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

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
		return fmt.Errorf("Error updating RouterNatAddress %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating RouterNatAddress %q: %#v", d.Id(), res)
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Updating RouterNatAddress", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceComputeRouterNatAddressRead(d, meta)
}

func resourceComputeRouterNatAddressDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RouterNatAddress: %s", err)
	}
	billingProject = project

	lockName, err := tpgresource.ReplaceVars(d, config, "router/{{region}}/{{router}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/routers/{{router}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	obj, err = resourceComputeRouterNatAddressPatchDeleteEncoder(d, meta, obj)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "RouterNatAddress")
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	// Since RouterNatAddress reopresents only the natIps field, we must make sure we only remove this value and not the entire nat
	obj, err = resourceComputeRouterNatAddressDeleteOnlyNatIps(d, meta, obj)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Deleting RouterNatAddress %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "RouterNatAddress")
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Deleting RouterNatAddress", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting RouterNatAddress %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeRouterNatAddressImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/routers/(?P<router>[^/]+)/(?P<router_nat>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<router>[^/]+)/(?P<router_nat>[^/]+)$",
		"^(?P<region>[^/]+)/(?P<router>[^/]+)/(?P<router_nat>[^/]+)$",
		"^(?P<router>[^/]+)/(?P<router_nat>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/regions/{{region}}/routers/{{router}}/{{router_nat}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNestedComputeRouterNatAddressNatIps(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertAndMapStringArr(v.([]interface{}), tpgresource.ConvertSelfLinkToV1)
}

func flattenNestedComputeRouterNatAddressDrainNatIps(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertAndMapStringArr(v.([]interface{}), tpgresource.ConvertSelfLinkToV1)
}

func flattenNestedComputeRouterNatAddressRouterNat(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func expandNestedComputeRouterNatAddressNatIps(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			return nil, fmt.Errorf("Invalid value for nat_ips: nil")
		}
		f, err := tpgresource.ParseRegionalFieldValue("addresses", raw.(string), "project", "region", "zone", d, config, true)
		if err != nil {
			return nil, fmt.Errorf("Invalid value for nat_ips: %s", err)
		}
		req = append(req, f.RelativeLink())
	}
	return req, nil
}

func expandNestedComputeRouterNatAddressDrainNatIps(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			return nil, fmt.Errorf("Invalid value for drain_nat_ips: nil")
		}
		f, err := tpgresource.ParseRegionalFieldValue("addresses", raw.(string), "project", "region", "zone", d, config, true)
		if err != nil {
			return nil, fmt.Errorf("Invalid value for drain_nat_ips: %s", err)
		}
		req = append(req, f.RelativeLink())
	}
	return req, nil
}

func expandNestedComputeRouterNatAddressRouterNat(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceComputeRouterNatAddressEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	currItems, err := resourceComputeRouterNatAddressListForPatch(d, meta)
	if err != nil {
		return nil, err
	}

	idx, found, err := resourceComputeRouterNatAddressFindNestedObjectInList(d, meta, currItems)
	if err != nil {
		return nil, err
	}

	// Merge new with existing item if item was already created (with the router nat resource).
	if found != nil {
		// Merge new object into old.
		for k, v := range obj {
			found[k] = v
		}
		currItems[idx] = found

		// Return list with new item added
		resPatch := map[string]interface{}{
			"nats": currItems,
		}

		return resPatch, nil
	}

	// Prevent creating a RouterNatAddress if no RouterNat has been found
	log.Printf("[WARNING] No RouterNat resource %+v found, preventing RouterNatAddress creation", obj)
	res := map[string]interface{}{
		"nats": nil,
	}

	return res, nil
}

func resourceComputeRouterNatAddressUpdateEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// Since we only want to change the handling of the CREATE function, this encoder just returns the unchanged obj value
	return obj, nil
}

func flattenNestedComputeRouterNatAddress(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	var v interface{}
	var ok bool

	v, ok = res["nats"]
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
		return nil, fmt.Errorf("expected list or map for value nats. Actual value: %v", v)
	}

	_, item, err := resourceComputeRouterNatAddressFindNestedObjectInList(d, meta, v.([]interface{}))
	if err != nil {
		return nil, err
	}
	return item, nil
}

func resourceComputeRouterNatAddressFindNestedObjectInList(d *schema.ResourceData, meta interface{}, items []interface{}) (index int, item map[string]interface{}, err error) {
	expectedRouterNat, err := expandNestedComputeRouterNatAddressRouterNat(d.Get("router_nat"), d, meta.(*transport_tpg.Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedRouterNat := flattenNestedComputeRouterNatAddressRouterNat(expectedRouterNat, d, meta.(*transport_tpg.Config))

	// Search list for this resource.
	for idx, itemRaw := range items {
		if itemRaw == nil {
			continue
		}
		item := itemRaw.(map[string]interface{})

		itemRouterNat := flattenNestedComputeRouterNatAddressRouterNat(item["name"], d, meta.(*transport_tpg.Config))
		// IsEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(tpgresource.IsEmptyValue(reflect.ValueOf(itemRouterNat)) && tpgresource.IsEmptyValue(reflect.ValueOf(expectedFlattenedRouterNat))) && !reflect.DeepEqual(itemRouterNat, expectedFlattenedRouterNat) {
			log.Printf("[DEBUG] Skipping item with name= %#v, looking for %#v)", itemRouterNat, expectedFlattenedRouterNat)
			continue
		}
		log.Printf("[DEBUG] Found item for resource %q: %#v)", d.Id(), item)
		return idx, item, nil
	}
	return -1, nil, nil
}

// PatchCreateEncoder handles creating request data to PATCH parent resource
// with list including new object.
func resourceComputeRouterNatAddressPatchCreateEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	currItems, err := resourceComputeRouterNatAddressListForPatch(d, meta)
	if err != nil {
		return nil, err
	}

	_, found, err := resourceComputeRouterNatAddressFindNestedObjectInList(d, meta, currItems)
	if err != nil {
		return nil, err
	}

	// Return error if item already created.
	if found != nil {
		return nil, fmt.Errorf("Unable to create RouterNatAddress, existing object already found: %+v", found)
	}

	// Return list with the resource to create appended
	res := map[string]interface{}{
		"nats": append(currItems, obj),
	}

	return res, nil
}

// PatchUpdateEncoder handles creating request data to PATCH parent resource
// with list including updated object.
func resourceComputeRouterNatAddressPatchUpdateEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	items, err := resourceComputeRouterNatAddressListForPatch(d, meta)
	if err != nil {
		return nil, err
	}

	idx, item, err := resourceComputeRouterNatAddressFindNestedObjectInList(d, meta, items)
	if err != nil {
		return nil, err
	}

	// Return error if item to update does not exist.
	if item == nil {
		return nil, fmt.Errorf("Unable to update RouterNatAddress %q - not found in list", d.Id())
	}

	// Copy over values for immutable fields
	obj["name"] = item["name"]
	// Merge any fields in item that aren't managed by this resource into obj
	// This is necessary because item might be managed by multiple resources.
	settableFields := map[string]struct{}{
		"natIps":      struct{}{},
		"drainNatIps": struct{}{},
	}
	for k, v := range item {
		if _, ok := settableFields[k]; !ok {
			obj[k] = v
		}
	}

	// Override old object with new
	items[idx] = obj

	// Return list with new item added
	res := map[string]interface{}{
		"nats": items,
	}

	return res, nil
}

// PatchDeleteEncoder handles creating request data to PATCH parent resource
// with list excluding object to delete.
func resourceComputeRouterNatAddressPatchDeleteEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	currItems, err := resourceComputeRouterNatAddressListForPatch(d, meta)
	if err != nil {
		return nil, err
	}

	idx, item, err := resourceComputeRouterNatAddressFindNestedObjectInList(d, meta, currItems)
	if err != nil {
		return nil, err
	}
	if item == nil {
		// Spoof 404 error for proper handling by Delete (i.e. no-op)
		return nil, tpgresource.Fake404("nested", "ComputeRouterNatAddress")
	}

	updatedItems := append(currItems[:idx], currItems[idx+1:]...)
	res := map[string]interface{}{
		"nats": updatedItems,
	}

	return res, nil
}

// ListForPatch handles making API request to get parent resource and
// extracting list of objects.
func resourceComputeRouterNatAddressListForPatch(d *schema.ResourceData, meta interface{}) ([]interface{}, error) {
	config := meta.(*transport_tpg.Config)
	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/routers/{{router}}")
	if err != nil {
		return nil, err
	}
	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return nil, err
	}

	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return nil, err
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   project,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return nil, err
	}
	var v interface{}
	var ok bool

	v, ok = res["nats"]
	if ok && v != nil {
		ls, lsOk := v.([]interface{})
		if !lsOk {
			return nil, fmt.Errorf(`expected list for nested field "nats"`)
		}
		return ls, nil
	}
	return nil, nil
}
