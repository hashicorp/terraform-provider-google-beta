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

package chronicle_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccChronicleWatchlist_chronicleWatchlistBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"chronicle_id":  envvar.GetTestChronicleInstanceIdFromEnv(t),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckChronicleWatchlistDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccChronicleWatchlist_chronicleWatchlistBasicExample(context),
			},
			{
				ResourceName:            "google_chronicle_watchlist.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"instance", "location", "watchlist_id"},
			},
		},
	})
}

func testAccChronicleWatchlist_chronicleWatchlistBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_chronicle_watchlist" "example" {
  provider = "google-beta"
  location = "us"
  instance = "%{chronicle_id}"
  watchlist_id = "tf-test-watchlist-name%{random_suffix}"
  description = "tf-test-watchlist-description%{random_suffix}"
  display_name = "tf-test-watchlist-name%{random_suffix}"
  multiplying_factor = 1
  entity_population_mechanism {
    manual {

    }
  }
  watchlist_user_preferences {
    pinned = true
  }
}
`, context)
}

func testAccCheckChronicleWatchlistDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_chronicle_watchlist" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ChronicleBasePath}}projects/{{project}}/locations/{{location}}/instances/{{instance}}/watchlists/{{watchlist_id}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("ChronicleWatchlist still exists at %s", url)
			}
		}

		return nil
	}
}
