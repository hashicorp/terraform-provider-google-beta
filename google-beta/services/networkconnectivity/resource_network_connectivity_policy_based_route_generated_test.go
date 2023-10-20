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

package networkconnectivity_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccNetworkConnectivityPolicyBasedRoute_networkConnectivityPolicyBasedRouteBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkConnectivityPolicyBasedRouteDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkConnectivityPolicyBasedRoute_networkConnectivityPolicyBasedRouteBasicExample(context),
			},
			{
				ResourceName:            "google_network_connectivity_policy_based_route.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccNetworkConnectivityPolicyBasedRoute_networkConnectivityPolicyBasedRouteBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_network_connectivity_policy_based_route" "default" {
  name = "tf-test-my-pbr%{random_suffix}"
  network = google_compute_network.my_network.id
  filter {
    protocol_version = "IPV4"
  }
  next_hop_other_routes = "DEFAULT_ROUTING"
}

resource "google_compute_network" "my_network" {
  name                    = "tf-test-my-network%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func TestAccNetworkConnectivityPolicyBasedRoute_networkConnectivityPolicyBasedRouteFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkConnectivityPolicyBasedRouteDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkConnectivityPolicyBasedRoute_networkConnectivityPolicyBasedRouteFullExample(context),
			},
			{
				ResourceName:            "google_network_connectivity_policy_based_route.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccNetworkConnectivityPolicyBasedRoute_networkConnectivityPolicyBasedRouteFullExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_network_connectivity_policy_based_route" "default" {
  name = "tf-test-my-pbr%{random_suffix}"
  description = "My routing policy"
  network = google_compute_network.my_network.id
  priority = 2302

  filter {
    protocol_version = "IPV4"
    ip_protocol = "UDP"
    src_range = "10.0.0.0/24"
    dest_range = "0.0.0.0/0"
  }
  next_hop_ilb_ip = google_compute_global_address.ilb.address

  virtual_machine {
    tags = ["restricted"]
  }

  labels = {
    env = "default"
  }
}

resource "google_compute_network" "my_network" {
  name                    = "tf-test-my-network%{random_suffix}"
  auto_create_subnetworks = false
}

# This example substitutes an arbitrary internal IP for an internal network
# load balancer for brevity. Consult https://cloud.google.com/load-balancing/docs/internal
# to set one up.
resource "google_compute_global_address" "ilb" {
  name = "tf-test-my-ilb%{random_suffix}"
}
`, context)
}

func testAccCheckNetworkConnectivityPolicyBasedRouteDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_network_connectivity_policy_based_route" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{NetworkConnectivityBasePath}}projects/{{project}}/locations/global/policyBasedRoutes/{{name}}")
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
				return fmt.Errorf("NetworkConnectivityPolicyBasedRoute still exists at %s", url)
			}
		}

		return nil
	}
}
