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

package networksecurity

import (
	"context"
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/sweeper"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func init() {
	sweeper.AddTestSweepers("NetworkSecurityInterceptDeployment", testSweepNetworkSecurityInterceptDeployment)
}

func testSweepNetworkSecurityInterceptDeployment(_ string) error {
	var deletionerror error
	resourceName := "NetworkSecurityInterceptDeployment"
	log.Printf("[INFO][SWEEPER_LOG] Starting sweeper for %s", resourceName)
	// Using URL substitutions for region/zone pairs
	substitutions := []struct {
		region string
		zone   string
	}{
		{region: "us-central1-a", zone: ""},
	}

	// Iterate through each substitution
	for _, sub := range substitutions {
		config, err := sweeper.SharedConfigForRegion(sub.region)
		if err != nil {
			log.Printf("[INFO][SWEEPER_LOG] error getting shared config for region: %s", err)
			return err
		}

		err = config.LoadAndValidate(context.Background())
		if err != nil {
			log.Printf("[INFO][SWEEPER_LOG] error loading: %s", err)
			return err
		}

		t := &testing.T{}
		billingId := envvar.GetTestBillingAccountFromEnv(t)

		// Set fallback values for empty region/zone
		if sub.region == "" {
			log.Printf("[INFO][SWEEPER_LOG] Empty region provided, falling back to us-central1")
			sub.region = "us-central1"
		}
		if sub.zone == "" {
			log.Printf("[INFO][SWEEPER_LOG] Empty zone provided, falling back to us-central1-a")
			sub.zone = "us-central1-a"
		}

		// Setup variables to replace in list template
		d := &tpgresource.ResourceDataMock{
			FieldsInSchema: map[string]interface{}{
				"project":         config.Project,
				"region":          sub.region,
				"location":        sub.region,
				"zone":            sub.zone,
				"billing_account": billingId,
			},
		}

		listTemplate := strings.Split("https://networksecurity.googleapis.com/v1beta1/projects/{{project}}/locations/{{location}}/interceptDeployments", "?")[0]
		listUrl, err := tpgresource.ReplaceVars(d, config, listTemplate)
		if err != nil {
			log.Printf("[INFO][SWEEPER_LOG] error preparing sweeper list url: %s", err)
			return err
		}

		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "GET",
			Project:   config.Project,
			RawURL:    listUrl,
			UserAgent: config.UserAgent,
		})
		if err != nil {
			log.Printf("[INFO][SWEEPER_LOG] Error in response from request %s: %s", listUrl, err)
			return err
		}

		resourceList, ok := res["interceptDeployments"]
		if !ok {
			log.Printf("[INFO][SWEEPER_LOG] Nothing found in response.")
			continue
		}
		rl := resourceList.([]interface{})

		log.Printf("[INFO][SWEEPER_LOG] Found %d items in %s list response.", len(rl), resourceName)
		// Keep count of items that aren't sweepable for logging.
		nonPrefixCount := 0
		for _, ri := range rl {
			obj := ri.(map[string]interface{})
			var name string
			// Id detected in the delete URL, attempt to use id.
			if obj["id"] != nil {
				name = tpgresource.GetResourceNameFromSelfLink(obj["id"].(string))
			} else if obj["name"] != nil {
				name = tpgresource.GetResourceNameFromSelfLink(obj["name"].(string))
			} else {
				log.Printf("[INFO][SWEEPER_LOG] %s resource name and id were nil", resourceName)
				return err
			}

			// Skip resources that shouldn't be sweeped
			if !sweeper.IsSweepableTestResource(name) {
				nonPrefixCount++
				continue
			}

			deleteTemplate := "https://networksecurity.googleapis.com/v1beta1/projects/{{project}}/locations/{{location}}/interceptDeployments/{{intercept_deployment_id}}"

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
		}

		if nonPrefixCount > 0 {
			log.Printf("[INFO][SWEEPER_LOG] %d items were non-sweepable and skipped.", nonPrefixCount)
		}
	}

	return deletionerror
}
