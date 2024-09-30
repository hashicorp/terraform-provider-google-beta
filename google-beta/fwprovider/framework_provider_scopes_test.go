// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package fwprovider_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// TestAccFwProvider_scopes is a series of acc tests asserting how the PF provider handles scopes arguments
// It is PF specific because the HCL used provisions PF-implemented resources
// It is a counterpart to TestAccSdkProvider_scopes
func TestAccFwProvider_scopes(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		// Configuring the provider using inputs
		"default scopes are used when there are no user inputs": testAccFwProvider_scopes_providerDefault,
		"scopes can be set in config":                           testAccFwProvider_scopes_setInConfig,
		//no ENVs to test

		// Schema-level validation
		"when scopes is set to an empty array in the config the value is ignored and default scopes are used": testAccFwProvider_scopes_emptyArray,

		// Usage
		//    Beta-only generation is needed because we need to access a PF-implemented data source linked to resource in an API.
		//    Currently this only exists in TPGB.
		"the scopes argument impacts provisioning resources": testAccFwProvider_scopes_usage,
	}

	for name, tc := range testCases {
		// shadow the tc variable into scope so that when
		// the loop continues, if t.Run hasn't executed tc(t)
		// yet, we don't have a race condition
		// see https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		tc := tc
		t.Run(name, func(t *testing.T) {
			tc(t)
		})
	}
}

func testAccFwProvider_scopes_providerDefault(t *testing.T) {
	acctest.SkipIfVcr(t) // Test doesn't interact with API

	acctest.VcrTest(t, resource.TestCase{
		// No PreCheck for checking ENVs
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFwProvider_scopes_unset(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.google_provider_config_plugin_framework.default", "scopes.#", fmt.Sprintf("%d", len(transport.DefaultClientScopes))),
					resource.TestCheckResourceAttr("data.google_provider_config_plugin_framework.default", "scopes.0", transport.DefaultClientScopes[0]),
					resource.TestCheckResourceAttr("data.google_provider_config_plugin_framework.default", "scopes.1", transport.DefaultClientScopes[1]),
				),
			},
		},
	})
}

func testAccFwProvider_scopes_setInConfig(t *testing.T) {
	acctest.SkipIfVcr(t) // Test doesn't interact with API

	scopes := []string{"https://www.googleapis.com/auth/cloud-platform"} // first of the two default scopes
	context := map[string]interface{}{
		"scopes": fmt.Sprintf("[\"%s\"]", scopes[0]),
	}

	acctest.VcrTest(t, resource.TestCase{
		// No PreCheck for checking ENVs
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFwProvider_scopes_inProviderBlock(context),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.google_provider_config_plugin_framework.default", "scopes.#", fmt.Sprintf("%d", len(scopes))),
					resource.TestCheckResourceAttr("data.google_provider_config_plugin_framework.default", "scopes.0", scopes[0]),
				),
			},
		},
	})
}

func testAccFwProvider_scopes_emptyArray(t *testing.T) {
	acctest.SkipIfVcr(t) // Test doesn't interact with API

	context := map[string]interface{}{
		"scopes": "[]",
	}

	acctest.VcrTest(t, resource.TestCase{
		// No PreCheck for checking ENVs
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFwProvider_scopes_inProviderBlock(context),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.google_provider_config_plugin_framework.default", "scopes.#", fmt.Sprintf("%d", len(transport.DefaultClientScopes))),
					resource.TestCheckResourceAttr("data.google_provider_config_plugin_framework.default", "scopes.0", transport.DefaultClientScopes[0]),
					resource.TestCheckResourceAttr("data.google_provider_config_plugin_framework.default", "scopes.1", transport.DefaultClientScopes[1]),
				),
			},
		},
	})
}

func testAccFwProvider_scopes_usage(t *testing.T) {
	acctest.SkipIfVcr(t) // Skip because Firebase is weird with VCR, and we have to use Firebase resources in the test

	// We include scopes that aren't sufficient to enable provisioning the resources in the config below
	context := map[string]interface{}{
		"scopes":        "[\"https://www.googleapis.com/auth/pubsub\"]",
		"random_suffix": acctest.RandString(t, 10),

		"bundle_id":    "apple.app." + acctest.RandString(t, 5),
		"display_name": "tf-test Display Name AppleAppConfig DataSource",
		"app_store_id": 12345,
		"team_id":      1234567890,
	}

	acctest.VcrTest(t, resource.TestCase{
		// No PreCheck for checking ENVs
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config:      testAccFwProvider_access_token_useScopes(context),
				ExpectError: regexp.MustCompile("Request had insufficient authentication scopes"),
			},
		},
	})
}

// testAccFwProvider_scopes_inProviderBlock allows setting the scopes argument in a provider block.
// This function uses data.google_provider_config_plugin_framework because it is implemented with the PF
func testAccFwProvider_scopes_inProviderBlock(context map[string]interface{}) string {
	return acctest.Nprintf(`
provider "google" {
	scopes = %{scopes}
}

data "google_provider_config_plugin_framework" "default" {}
`, context)
}

// testAccFwProvider_scopes_inEnvsOnly allows testing when the scopes argument is not set
func testAccFwProvider_scopes_unset() string {
	return `
data "google_provider_config_plugin_framework" "default" {}
`
}

func testAccFwProvider_access_token_useScopes(context map[string]interface{}) string {
	return acctest.Nprintf(`
provider "google" {} // default scopes used

provider "google" {
  alias = "under-scoped"
  scopes = %{scopes}
}

data "google_provider_config_plugin_framework" "default" {
}

resource "google_firebase_apple_app" "my_app_config" {
  project = data.google_provider_config_plugin_framework.default.project
  bundle_id = "%{bundle_id}"
  display_name = "%{display_name}"
  app_store_id = "%{app_store_id}"
  team_id = "%{team_id}"
}

// This is implemented with plugin-framework so tests our use of scopes in a PF specific way
data "google_firebase_apple_app_config" "my_app_config" {
  provider = google.under-scoped
  app_id = google_firebase_apple_app.my_app_config.app_id
}
`, context)
}
