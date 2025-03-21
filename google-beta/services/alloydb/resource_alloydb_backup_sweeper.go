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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/alloydb/Backup.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/sweeper_file.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package alloydb

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/sweeper"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func init() {
	// Initialize base sweeper object
	s := &sweeper.Sweeper{
		Name:           "google_alloydb_backup",
		ListAndAction:  listAndActionAlloydbBackup,
		DeleteFunction: testSweepAlloydbBackup,
	}

	// Register the sweeper
	sweeper.AddTestSweepers(s)
}

func testSweepAlloydbBackup(_ string) error {
	return listAndActionAlloydbBackup(deleteResourceAlloydbBackup)
}

func listAndActionAlloydbBackup(action sweeper.ResourceAction) error {
	var lastError error
	resourceName := "AlloydbBackup"
	log.Printf("[INFO][SWEEPER_LOG] Starting sweeper for %s", resourceName)

	// Prepare configurations to iterate over
	var configs []*tpgresource.ResourceDataMock
	t := &testing.T{}
	billingId := envvar.GetTestBillingAccountFromEnv(t)
	// Default single config
	intermediateValues := []map[string]string{
		{
			"region": "us-central1",
		},
	}

	// Create configs from intermediate values
	for _, values := range intermediateValues {
		mockConfig := &tpgresource.ResourceDataMock{
			FieldsInSchema: map[string]interface{}{
				"project":         envvar.GetTestProjectFromEnv(),
				"billing_account": billingId,
			},
		}

		// Apply all provided values
		for key, value := range values {
			mockConfig.FieldsInSchema[key] = value
		}

		// Set fallback values for common fields
		region, hasRegion := mockConfig.FieldsInSchema["region"].(string)
		if !hasRegion {
			region = "us-central1"
			mockConfig.FieldsInSchema["region"] = region
		}

		if _, hasLocation := mockConfig.FieldsInSchema["location"]; !hasLocation {
			mockConfig.FieldsInSchema["location"] = region
		}

		if _, hasZone := mockConfig.FieldsInSchema["zone"]; !hasZone {
			mockConfig.FieldsInSchema["zone"] = region + "-a"
		}

		configs = append(configs, mockConfig)
	}

	// Process all configurations (either from parent resources or direct substitutions)
	for _, mockConfig := range configs {
		// Get region from config
		region := sweeper.GetFieldOrDefault(mockConfig, "region", "us-central1")

		// Create shared config for this region
		config, err := sweeper.SharedConfigForRegion(region)
		if err != nil {
			log.Printf("[INFO][SWEEPER_LOG] error getting shared config for region: %s", err)
			lastError = err
			continue
		}

		err = config.LoadAndValidate(context.Background())
		if err != nil {
			log.Printf("[INFO][SWEEPER_LOG] error loading: %s", err)
			lastError = err
			continue
		}

		// Prepare list URL
		listTemplate := strings.Split("https://alloydb.googleapis.com/v1beta/projects/{{project}}/locations/{{location}}/backups", "?")[0]
		listUrl, err := tpgresource.ReplaceVars(mockConfig, config, listTemplate)
		if err != nil {
			log.Printf("[INFO][SWEEPER_LOG] error preparing sweeper list url: %s", err)
			lastError = err
			continue
		}

		// Log additional info for parent-based resources
		log.Printf("[INFO][SWEEPER_LOG] Listing %s resources at %s", resourceName, listUrl)

		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "GET",
			Project:   config.Project,
			RawURL:    listUrl,
			UserAgent: config.UserAgent,
		})
		if err != nil {
			log.Printf("[INFO][SWEEPER_LOG] Error in response from request %s: %s", listUrl, err)
			lastError = err
			continue
		}

		// First try the expected resource key
		resourceList, ok := res["backups"]
		if ok {
			log.Printf("[INFO][SWEEPER_LOG] Found resources under expected key 'backups'")
		} else {
			// Next, try the common "items" pattern
			resourceList, ok = res["items"]
			if ok {
				log.Printf("[INFO][SWEEPER_LOG] Found resources under standard 'items' key")
			} else {
				log.Printf("[INFO][SWEEPER_LOG] no resources found")
				continue
			}
		}
		rl := resourceList.([]interface{})

		log.Printf("[INFO][SWEEPER_LOG] Found %d items in %s list response.", len(rl), resourceName)
		// Keep count of items that aren't sweepable for logging.
		nonPrefixCount := 0
		for _, ri := range rl {
			obj, ok := ri.(map[string]interface{})
			if !ok {
				log.Printf("[INFO][SWEEPER_LOG] Item was not a map: %T", ri)
				continue
			}

			if err := action(config, mockConfig, obj); err != nil {
				log.Printf("[INFO][SWEEPER_LOG] Error in action: %s", err)
				lastError = err
			} else {
				nonPrefixCount++
			}
		}
	}

	return lastError
}

func deleteResourceAlloydbBackup(config *transport_tpg.Config, d *tpgresource.ResourceDataMock, obj map[string]interface{}) error {
	var deletionerror error
	resourceName := "AlloydbBackup"
	var name string
	// Id detected in the delete URL, attempt to use id.
	if obj["id"] != nil {
		name = tpgresource.GetResourceNameFromSelfLink(obj["id"].(string))
	} else if obj["name"] != nil {
		name = tpgresource.GetResourceNameFromSelfLink(obj["name"].(string))
	} else {
		log.Printf("[INFO][SWEEPER_LOG] %s resource name and id were nil", resourceName)
		return fmt.Errorf("%s resource name was nil", resourceName)
	}

	// Skip resources that shouldn't be sweeped
	if !sweeper.IsSweepableTestResource(name) {
		return nil
	}

	deleteTemplate := "https://alloydb.googleapis.com/v1beta/projects/{{project}}/locations/{{location}}/backups/{{backup_id}}"

	deleteUrl, err := tpgresource.ReplaceVars(d, config, deleteTemplate)
	if err != nil {
		log.Printf("[INFO][SWEEPER_LOG] error preparing delete url: %s", err)
		deletionerror = err
	}
	deleteUrl = deleteUrl + name

	// Don't wait on operations as we may have a lot to delete
	_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   config.Project,
		RawURL:    deleteUrl,
		UserAgent: config.UserAgent,
	})
	if err != nil {
		log.Printf("[INFO][SWEEPER_LOG] Error deleting for url %s : %s", deleteUrl, err)
		deletionerror = err
	} else {
		log.Printf("[INFO][SWEEPER_LOG] Sent delete request for %s resource: %s", resourceName, name)
	}

	return deletionerror
}
