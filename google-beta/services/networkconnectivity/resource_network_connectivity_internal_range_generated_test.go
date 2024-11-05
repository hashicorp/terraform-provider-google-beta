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

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccNetworkConnectivityInternalRange_networkConnectivityInternalRangesBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkConnectivityInternalRangeDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkConnectivityInternalRange_networkConnectivityInternalRangesBasicExample(context),
			},
			{
				ResourceName:            "google_network_connectivity_internal_range.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "name", "network", "terraform_labels"},
			},
		},
	})
}

func testAccNetworkConnectivityInternalRange_networkConnectivityInternalRangesBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_network_connectivity_internal_range" "default" {
  name    = "basic%{random_suffix}"
  description = "Test internal range"
  network = google_compute_network.default.self_link
  usage   = "FOR_VPC"
  peering = "FOR_SELF"
  ip_cidr_range = "10.0.0.0/24"

  labels  = {
    label-a: "b"
  }
}

resource "google_compute_network" "default" {
  name                    = "tf-test-internal-ranges%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func TestAccNetworkConnectivityInternalRange_networkConnectivityInternalRangesAutomaticReservationExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkConnectivityInternalRangeDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkConnectivityInternalRange_networkConnectivityInternalRangesAutomaticReservationExample(context),
			},
			{
				ResourceName:            "google_network_connectivity_internal_range.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "name", "network", "terraform_labels"},
			},
		},
	})
}

func testAccNetworkConnectivityInternalRange_networkConnectivityInternalRangesAutomaticReservationExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_network_connectivity_internal_range" "default" {
  name    = "tf-test-automatic-reservation%{random_suffix}"
  network = google_compute_network.default.id
  usage   = "FOR_VPC"
  peering = "FOR_SELF"
  prefix_length = 24
  target_cidr_range = [
    "192.16.0.0/16"
  ]
}

resource "google_compute_network" "default" {
  name                    = "tf-test-internal-ranges%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func TestAccNetworkConnectivityInternalRange_networkConnectivityInternalRangesExternalRangesExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkConnectivityInternalRangeDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkConnectivityInternalRange_networkConnectivityInternalRangesExternalRangesExample(context),
			},
			{
				ResourceName:            "google_network_connectivity_internal_range.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "name", "network", "terraform_labels"},
			},
		},
	})
}

func testAccNetworkConnectivityInternalRange_networkConnectivityInternalRangesExternalRangesExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_network_connectivity_internal_range" "default" {
  name    = "tf-test-external-ranges%{random_suffix}"
  network = google_compute_network.default.id
  usage   = "EXTERNAL_TO_VPC"
  peering = "FOR_SELF"
  ip_cidr_range = "172.16.0.0/24"

  labels  = {
    external-reserved-range: "on-premises"
  }
}

resource "google_compute_network" "default" {
  name                    = "tf-test-internal-ranges%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func TestAccNetworkConnectivityInternalRange_networkConnectivityInternalRangesReserveWithOverlapExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkConnectivityInternalRangeDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkConnectivityInternalRange_networkConnectivityInternalRangesReserveWithOverlapExample(context),
			},
			{
				ResourceName:            "google_network_connectivity_internal_range.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "name", "network", "terraform_labels"},
			},
		},
	})
}

func testAccNetworkConnectivityInternalRange_networkConnectivityInternalRangesReserveWithOverlapExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_network_connectivity_internal_range" "default" {
  name    = "tf-test-overlap-range%{random_suffix}"
  description = "Test internal range"
  network = google_compute_network.default.id
  usage   = "FOR_VPC"
  peering = "FOR_SELF"
  ip_cidr_range = "10.0.0.0/30"

  overlaps = [
    "OVERLAP_EXISTING_SUBNET_RANGE"
  ]

  depends_on = [
    google_compute_subnetwork.default
  ]
}

resource "google_compute_network" "default" {
  name                    = "tf-test-internal-ranges%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  name          = "overlapping-subnet"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.id
}
`, context)
}

func TestAccNetworkConnectivityInternalRange_networkConnectivityInternalRangesMigrationExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkConnectivityInternalRangeDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkConnectivityInternalRange_networkConnectivityInternalRangesMigrationExample(context),
			},
			{
				ResourceName:            "google_network_connectivity_internal_range.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "name", "network", "terraform_labels"},
			},
		},
	})
}

func testAccNetworkConnectivityInternalRange_networkConnectivityInternalRangesMigrationExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_network_connectivity_internal_range" "default" {
  name          = "migration%{random_suffix}"
  description   = "Test internal range"
  network       = google_compute_network.default.self_link
  usage         = "FOR_MIGRATION"
  peering       = "FOR_SELF"
  ip_cidr_range = "10.1.0.0/16"
  migration {
    source = google_compute_subnetwork.source.self_link
    target = "projects/${data.google_project.target_project.project_id}/regions/us-central1/subnetworks/target-subnet"
  }
}

resource "google_compute_network" "default" {
  name                    = "tf-test-internal-ranges%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "source" {
  name          = "tf-test-source-subnet%{random_suffix}"
  ip_cidr_range = "10.1.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.default.name
}

data "google_project" "target_project" {
}
`, context)
}

func testAccCheckNetworkConnectivityInternalRangeDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_network_connectivity_internal_range" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{NetworkConnectivityBasePath}}projects/{{project}}/locations/global/internalRanges/{{name}}")
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
				return fmt.Errorf("NetworkConnectivityInternalRange still exists at %s", url)
			}
		}

		return nil
	}
}
