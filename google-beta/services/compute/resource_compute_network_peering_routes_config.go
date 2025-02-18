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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/compute/NetworkPeeringRoutesConfig.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
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
)

func ResourceComputeNetworkPeeringRoutesConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeNetworkPeeringRoutesConfigCreate,
		Read:   resourceComputeNetworkPeeringRoutesConfigRead,
		Update: resourceComputeNetworkPeeringRoutesConfigUpdate,
		Delete: resourceComputeNetworkPeeringRoutesConfigDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeNetworkPeeringRoutesConfigImport,
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
			"export_custom_routes": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: `Whether to export the custom routes to the peer network.`,
			},
			"import_custom_routes": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: `Whether to import the custom routes to the peer network.`,
			},
			"network": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The name of the primary network for the peering.`,
			},
			"peering": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Name of the peering.`,
			},
			"export_subnet_routes_with_public_ip": {
				Type:     schema.TypeBool,
				Computed: true,
				Optional: true,
				Description: `Whether subnet routes with public IP range are exported.
IPv4 special-use ranges are always exported to peers and
are not controlled by this field.`,
			},
			"import_subnet_routes_with_public_ip": {
				Type:     schema.TypeBool,
				Computed: true,
				Optional: true,
				Description: `Whether subnet routes with public IP range are imported.
IPv4 special-use ranges are always imported from peers and
are not controlled by this field.`,
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

func resourceComputeNetworkPeeringRoutesConfigCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandNestedComputeNetworkPeeringRoutesConfigPeering(d.Get("peering"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("peering"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	exportCustomRoutesProp, err := expandNestedComputeNetworkPeeringRoutesConfigExportCustomRoutes(d.Get("export_custom_routes"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("export_custom_routes"); ok || !reflect.DeepEqual(v, exportCustomRoutesProp) {
		obj["exportCustomRoutes"] = exportCustomRoutesProp
	}
	importCustomRoutesProp, err := expandNestedComputeNetworkPeeringRoutesConfigImportCustomRoutes(d.Get("import_custom_routes"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("import_custom_routes"); ok || !reflect.DeepEqual(v, importCustomRoutesProp) {
		obj["importCustomRoutes"] = importCustomRoutesProp
	}
	exportSubnetRoutesWithPublicIpProp, err := expandNestedComputeNetworkPeeringRoutesConfigExportSubnetRoutesWithPublicIp(d.Get("export_subnet_routes_with_public_ip"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("export_subnet_routes_with_public_ip"); ok || !reflect.DeepEqual(v, exportSubnetRoutesWithPublicIpProp) {
		obj["exportSubnetRoutesWithPublicIp"] = exportSubnetRoutesWithPublicIpProp
	}
	importSubnetRoutesWithPublicIpProp, err := expandNestedComputeNetworkPeeringRoutesConfigImportSubnetRoutesWithPublicIp(d.Get("import_subnet_routes_with_public_ip"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("import_subnet_routes_with_public_ip"); ok || !reflect.DeepEqual(v, importSubnetRoutesWithPublicIpProp) {
		obj["importSubnetRoutesWithPublicIp"] = importSubnetRoutesWithPublicIpProp
	}

	obj, err = resourceComputeNetworkPeeringRoutesConfigEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	lockName, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/global/networks/{{network}}/peerings")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/networks/{{network}}/updatePeering")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new NetworkPeeringRoutesConfig: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for NetworkPeeringRoutesConfig: %s", err)
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
		return fmt.Errorf("Error creating NetworkPeeringRoutesConfig: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/global/networks/{{network}}/networkPeerings/{{peering}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ComputeOperationWaitTime(
		config, res, project, "Creating NetworkPeeringRoutesConfig", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create NetworkPeeringRoutesConfig: %s", err)
	}

	log.Printf("[DEBUG] Finished creating NetworkPeeringRoutesConfig %q: %#v", d.Id(), res)

	return resourceComputeNetworkPeeringRoutesConfigRead(d, meta)
}

func resourceComputeNetworkPeeringRoutesConfigRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/networks/{{network}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for NetworkPeeringRoutesConfig: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ComputeNetworkPeeringRoutesConfig %q", d.Id()))
	}

	res, err = flattenNestedComputeNetworkPeeringRoutesConfig(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Object isn't there any more - remove it from the state.
		log.Printf("[DEBUG] Removing ComputeNetworkPeeringRoutesConfig because it couldn't be matched.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading NetworkPeeringRoutesConfig: %s", err)
	}

	if err := d.Set("peering", flattenNestedComputeNetworkPeeringRoutesConfigPeering(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkPeeringRoutesConfig: %s", err)
	}
	if err := d.Set("export_custom_routes", flattenNestedComputeNetworkPeeringRoutesConfigExportCustomRoutes(res["exportCustomRoutes"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkPeeringRoutesConfig: %s", err)
	}
	if err := d.Set("import_custom_routes", flattenNestedComputeNetworkPeeringRoutesConfigImportCustomRoutes(res["importCustomRoutes"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkPeeringRoutesConfig: %s", err)
	}
	if err := d.Set("export_subnet_routes_with_public_ip", flattenNestedComputeNetworkPeeringRoutesConfigExportSubnetRoutesWithPublicIp(res["exportSubnetRoutesWithPublicIp"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkPeeringRoutesConfig: %s", err)
	}
	if err := d.Set("import_subnet_routes_with_public_ip", flattenNestedComputeNetworkPeeringRoutesConfigImportSubnetRoutesWithPublicIp(res["importSubnetRoutesWithPublicIp"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkPeeringRoutesConfig: %s", err)
	}

	return nil
}

func resourceComputeNetworkPeeringRoutesConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for NetworkPeeringRoutesConfig: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	nameProp, err := expandNestedComputeNetworkPeeringRoutesConfigPeering(d.Get("peering"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("peering"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	exportCustomRoutesProp, err := expandNestedComputeNetworkPeeringRoutesConfigExportCustomRoutes(d.Get("export_custom_routes"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("export_custom_routes"); ok || !reflect.DeepEqual(v, exportCustomRoutesProp) {
		obj["exportCustomRoutes"] = exportCustomRoutesProp
	}
	importCustomRoutesProp, err := expandNestedComputeNetworkPeeringRoutesConfigImportCustomRoutes(d.Get("import_custom_routes"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("import_custom_routes"); ok || !reflect.DeepEqual(v, importCustomRoutesProp) {
		obj["importCustomRoutes"] = importCustomRoutesProp
	}
	exportSubnetRoutesWithPublicIpProp, err := expandNestedComputeNetworkPeeringRoutesConfigExportSubnetRoutesWithPublicIp(d.Get("export_subnet_routes_with_public_ip"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("export_subnet_routes_with_public_ip"); ok || !reflect.DeepEqual(v, exportSubnetRoutesWithPublicIpProp) {
		obj["exportSubnetRoutesWithPublicIp"] = exportSubnetRoutesWithPublicIpProp
	}
	importSubnetRoutesWithPublicIpProp, err := expandNestedComputeNetworkPeeringRoutesConfigImportSubnetRoutesWithPublicIp(d.Get("import_subnet_routes_with_public_ip"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("import_subnet_routes_with_public_ip"); ok || !reflect.DeepEqual(v, importSubnetRoutesWithPublicIpProp) {
		obj["importSubnetRoutesWithPublicIp"] = importSubnetRoutesWithPublicIpProp
	}

	obj, err = resourceComputeNetworkPeeringRoutesConfigEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	lockName, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/global/networks/{{network}}/peerings")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/networks/{{network}}/updatePeering")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating NetworkPeeringRoutesConfig %q: %#v", d.Id(), obj)
	headers := make(http.Header)

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
		return fmt.Errorf("Error updating NetworkPeeringRoutesConfig %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating NetworkPeeringRoutesConfig %q: %#v", d.Id(), res)
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Updating NetworkPeeringRoutesConfig", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceComputeNetworkPeeringRoutesConfigRead(d, meta)
}

func resourceComputeNetworkPeeringRoutesConfigDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARNING] Compute NetworkPeeringRoutesConfig resources"+
		" cannot be deleted from Google Cloud. The resource %s will be removed from Terraform"+
		" state, but will still be present on Google Cloud.", d.Id())
	d.SetId("")

	return nil
}

func resourceComputeNetworkPeeringRoutesConfigImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/global/networks/(?P<network>[^/]+)/networkPeerings/(?P<peering>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<network>[^/]+)/(?P<peering>[^/]+)$",
		"^(?P<network>[^/]+)/(?P<peering>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/global/networks/{{network}}/networkPeerings/{{peering}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNestedComputeNetworkPeeringRoutesConfigPeering(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedComputeNetworkPeeringRoutesConfigExportCustomRoutes(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedComputeNetworkPeeringRoutesConfigImportCustomRoutes(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedComputeNetworkPeeringRoutesConfigExportSubnetRoutesWithPublicIp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedComputeNetworkPeeringRoutesConfigImportSubnetRoutesWithPublicIp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNestedComputeNetworkPeeringRoutesConfigPeering(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNestedComputeNetworkPeeringRoutesConfigExportCustomRoutes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNestedComputeNetworkPeeringRoutesConfigImportCustomRoutes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNestedComputeNetworkPeeringRoutesConfigExportSubnetRoutesWithPublicIp(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNestedComputeNetworkPeeringRoutesConfigImportSubnetRoutesWithPublicIp(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceComputeNetworkPeeringRoutesConfigEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// Stick request in a networkPeering block as in
	// https://cloud.google.com/compute/docs/reference/rest/v1/networks/updatePeering
	newObj := make(map[string]interface{})
	newObj["networkPeering"] = obj
	return newObj, nil
}

func flattenNestedComputeNetworkPeeringRoutesConfig(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	var v interface{}
	var ok bool

	v, ok = res["peerings"]
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
		return nil, fmt.Errorf("expected list or map for value peerings. Actual value: %v", v)
	}

	_, item, err := resourceComputeNetworkPeeringRoutesConfigFindNestedObjectInList(d, meta, v.([]interface{}))
	if err != nil {
		return nil, err
	}
	return item, nil
}

func resourceComputeNetworkPeeringRoutesConfigFindNestedObjectInList(d *schema.ResourceData, meta interface{}, items []interface{}) (index int, item map[string]interface{}, err error) {
	expectedPeering, err := expandNestedComputeNetworkPeeringRoutesConfigPeering(d.Get("peering"), d, meta.(*transport_tpg.Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedPeering := flattenNestedComputeNetworkPeeringRoutesConfigPeering(expectedPeering, d, meta.(*transport_tpg.Config))

	// Search list for this resource.
	for idx, itemRaw := range items {
		if itemRaw == nil {
			continue
		}
		item := itemRaw.(map[string]interface{})

		itemPeering := flattenNestedComputeNetworkPeeringRoutesConfigPeering(item["name"], d, meta.(*transport_tpg.Config))
		// IsEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(tpgresource.IsEmptyValue(reflect.ValueOf(itemPeering)) && tpgresource.IsEmptyValue(reflect.ValueOf(expectedFlattenedPeering))) && !reflect.DeepEqual(itemPeering, expectedFlattenedPeering) {
			log.Printf("[DEBUG] Skipping item with name= %#v, looking for %#v)", itemPeering, expectedFlattenedPeering)
			continue
		}
		log.Printf("[DEBUG] Found item for resource %q: %#v)", d.Id(), item)
		return idx, item, nil
	}
	return -1, nil, nil
}
