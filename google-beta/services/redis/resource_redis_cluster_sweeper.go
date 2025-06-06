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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/redis/Cluster.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/sweeper_file.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package redis

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/sweeper"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func init() {
	// Initialize base sweeper object
	s := &sweeper.Sweeper{
		Name:           "google_redis_cluster",
		ListAndAction:  listAndActionRedisCluster,
		DeleteFunction: testSweepRedisCluster,
	}

	// Register the sweeper
	sweeper.AddTestSweepers(s)
}

func testSweepRedisCluster(_ string) error {
	return listAndActionRedisCluster(deleteResourceRedisCluster)
}

func listAndActionRedisCluster(action sweeper.ResourceAction) error {
	var lastError error
	resourceName := "RedisCluster"
	log.Printf("[INFO][SWEEPER_LOG] Starting sweeper for %s", resourceName)

	// Prepare configurations to iterate over
	var configs []*tpgresource.ResourceDataMock
	t := &testing.T{}
	billingId := envvar.GetTestBillingAccountFromEnv(t)
	// Build URL substitution maps individually to ensure proper formatting
	intermediateValues := make([]map[string]string, 3)
	intermediateValues[0] = map[string]string{}
	intermediateValues[0]["region"] = "us-central1"
	intermediateValues[1] = map[string]string{}
	intermediateValues[1]["region"] = "us-east1"
	intermediateValues[2] = map[string]string{}
	intermediateValues[2]["region"] = "europe-west1"

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
		listTemplate := strings.Split("https://redis.googleapis.com/v1beta1/projects/{{project}}/locations/{{region}}/clusters", "?")[0]
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
		resourceList, ok := res["clusters"]
		if ok {
			log.Printf("[INFO][SWEEPER_LOG] Found resources under expected key 'clusters'")
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

func deleteResourceRedisCluster(config *transport_tpg.Config, d *tpgresource.ResourceDataMock, obj map[string]interface{}) error {
	var deletionerror error
	resourceName := "RedisCluster"
	var name string
	if obj["name"] == nil {
		log.Printf("[INFO][SWEEPER_LOG] %s resource name was nil", resourceName)
		return fmt.Errorf("%s resource name was nil", resourceName)
	}

	name = tpgresource.GetResourceNameFromSelfLink(obj["name"].(string))

	// Skip resources that shouldn't be sweeped
	if !sweeper.IsSweepableTestResource(name) {
		return nil
	}

	deleteTemplate := "https://redis.googleapis.com/v1beta1/projects/{{project}}/locations/{{region}}/clusters/{{name}}"

	url, err := tpgresource.ReplaceVars(d, config, deleteTemplate)
	if err != nil {
		log.Printf("[INFO][SWEEPER_LOG] error preparing delete url: %s", err)
		deletionerror = err
	}
	// Ensure required field value before deletion
	targetValueStr := "false"

	// Convert target value to the appropriate type for comparison
	var targetValue interface{} = targetValueStr
	if targetValueStr == "true" {
		targetValue = true
	} else if targetValueStr == "false" {
		targetValue = false
	} else if intVal, err := strconv.Atoi(targetValueStr); err == nil {
		targetValue = intVal
	}

	// Parse the field path to handle nested fields
	fieldPath := strings.Split("deletionProtectionEnabled", ".")
	fieldName := fieldPath[0]

	// By default, assume we don't need to update
	needsUpdate := false
	fieldExists := false

	// Check if the field exists and if its value needs updating
	if len(fieldPath) == 1 {
		// Simple field at the top level
		if currentValue, hasValue := obj[fieldName]; hasValue {
			fieldExists = true
			// Only update if the value doesn't match
			if !reflect.DeepEqual(currentValue, targetValue) {
				needsUpdate = true
			}
		}
	} else {
		// Nested field
		if currentObj, hasTopLevel := obj[fieldName]; hasTopLevel {
			if nestedObj, ok := currentObj.(map[string]interface{}); ok {
				// Try to navigate through the nested structure
				current := nestedObj
				pathExists := true

				// Navigate through intermediate levels
				for i := 1; i < len(fieldPath)-1; i++ {
					if nextObj, hasNext := current[fieldPath[i]]; hasNext {
						if nextLevel, ok := nextObj.(map[string]interface{}); ok {
							current = nextLevel
						} else {
							// Not a map, can't continue navigation
							pathExists = false
							break
						}
					} else {
						// Field doesn't exist, can't continue navigation
						pathExists = false
						break
					}
				}

				// If we successfully navigated the path, check the final field
				if pathExists {
					finalFieldName := fieldPath[len(fieldPath)-1]
					if currentValue, exists := current[finalFieldName]; exists {
						fieldExists = true
						// Update only if the value doesn't match
						if !reflect.DeepEqual(currentValue, targetValue) {
							needsUpdate = true
						}
					}
				}
			}
		}
	}

	// Only proceed with update if the field exists and needs updating
	if fieldExists && needsUpdate {
		log.Printf("[INFO][SWEEPER_LOG] Ensuring %s is set to %v for %s resource: %s",
			"deletionProtectionEnabled", targetValue, resourceName, name)

		// Build URL for the update
		updateURL, err := tpgresource.ReplaceVars(d, config, deleteTemplate)
		if err != nil {
			log.Printf("[INFO][SWEEPER_LOG] error preparing update url: %s", err)
			return err
		}
		updateURL = updateURL + name

		// Create update object based on configuration
		var updateObj map[string]interface{}
		// Create object structure with just the field that needs to be updated
		if len(fieldPath) == 1 {
			// Simple field
			updateObj = map[string]interface{}{
				fieldName: targetValue,
			}
		} else {
			// For nested fields, create the object structure
			updateObj = make(map[string]interface{})
			currentObj := make(map[string]interface{})
			updateObj[fieldName] = currentObj

			// Build the nested structure
			for i := 1; i < len(fieldPath)-1; i++ {
				nestedObj := make(map[string]interface{})
				currentObj[fieldPath[i]] = nestedObj
				currentObj = nestedObj
			}

			// Set the final field
			currentObj[fieldPath[len(fieldPath)-1]] = targetValue
		}

		// Add update mask parameter
		updateURL, err = transport_tpg.AddQueryParams(updateURL, map[string]string{"updateMask": "deletionProtectionEnabled"})
		if err != nil {
			log.Printf("[INFO][SWEEPER_LOG] error adding query parameters: %s", err)
			return err
		}

		// Send the update request using the resource's update verb
		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "PATCH",
			Project:   config.Project,
			RawURL:    updateURL,
			UserAgent: config.UserAgent,
			Body:      updateObj,
		})

		if err != nil {
			log.Printf("[INFO][SWEEPER_LOG] Error ensuring field value: %s", err)
			return err
		}

		// Wait for the operation to complete using the resource's operation wait function
		err = RedisOperationWaitTime(
			config, res, config.Project, fmt.Sprintf("Ensuring %s value", "deletionProtectionEnabled"),
			config.UserAgent, time.Minute*5)

		if err != nil {
			log.Printf("[INFO][SWEEPER_LOG] Error waiting for operation to complete: %s", err)
			return err
		}

		log.Printf("[INFO][SWEEPER_LOG] Successfully updated %s for %s resource: %s",
			"deletionProtectionEnabled", resourceName, name)
	} else if !fieldExists {
		log.Printf("[INFO][SWEEPER_LOG] Field %s not found in resource, skipping update",
			"deletionProtectionEnabled")
	} else {
		log.Printf("[INFO][SWEEPER_LOG] Field %s already set to desired value, no update needed",
			"deletionProtectionEnabled")
	}
	url = url + name

	// Don't wait on operations as we may have a lot to delete
	_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   config.Project,
		RawURL:    url,
		UserAgent: config.UserAgent,
	})
	if err != nil {
		log.Printf("[INFO][SWEEPER_LOG] Error deleting for url %s : %s", url, err)
		deletionerror = err
	} else {
		log.Printf("[INFO][SWEEPER_LOG] Sent delete request for %s resource: %s", resourceName, name)
	}

	return deletionerror
}
