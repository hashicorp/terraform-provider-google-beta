// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
// ----------------------------------------------------------------------------
//
//	***     AUTO GENERATED CODE    ***    Type: Handwritten     ***
//
// ----------------------------------------------------------------------------
//
//	This code is generated by Magic Modules using the following:
//
//	Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/services/compute/data_source_google_compute_regions.go.tmpl
//
//	DO NOT EDIT this file directly. Any changes made to this file will be
//	overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
package compute

import (
	"fmt"
	"log"
	"sort"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	compute "google.golang.org/api/compute/v0.beta"
)

func DataSourceGoogleComputeRegions() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGoogleComputeRegionsRead,
		Schema: map[string]*schema.Schema{
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"names": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"UP", "DOWN"}, false),
			},
		},
	}
}

func dataSourceGoogleComputeRegionsRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return err
	}
	filter := ""
	if s, ok := d.GetOk("status"); ok {
		filter = fmt.Sprintf(" (status eq %s)", s)
	}

	call := config.NewComputeClient(userAgent).Regions.List(project).Filter(filter)

	resp, err := call.Do()
	if err != nil {
		return err
	}

	regions := flattenRegions(resp.Items)
	log.Printf("[DEBUG] Received Google Compute Regions: %q", regions)

	if err := d.Set("names", regions); err != nil {
		return fmt.Errorf("Error setting names: %s", err)
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error setting project: %s", err)
	}
	d.SetId(fmt.Sprintf("projects/%s", project))

	return nil
}

func flattenRegions(regions []*compute.Region) []string {
	result := make([]string, len(regions))
	for i, region := range regions {
		result[i] = region.Name
	}
	sort.Strings(result)
	return result
}
