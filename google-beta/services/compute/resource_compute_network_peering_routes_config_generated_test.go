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

package compute_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccComputeNetworkPeeringRoutesConfig_networkPeeringRoutesConfigBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeNetworkPeeringRoutesConfig_networkPeeringRoutesConfigBasicExample(context),
			},
			{
				ResourceName:            "google_compute_network_peering_routes_config.peering_primary_routes",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network"},
			},
		},
	})
}

func testAccComputeNetworkPeeringRoutesConfig_networkPeeringRoutesConfigBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network_peering_routes_config" "peering_primary_routes" {
  peering = google_compute_network_peering.peering_primary.name
  network = google_compute_network.network_primary.name

  import_custom_routes                = true
  export_custom_routes                = true
  import_subnet_routes_with_public_ip = true
  export_subnet_routes_with_public_ip = true
}

resource "google_compute_network_peering" "peering_primary" {
  name         = "tf-test-primary-peering%{random_suffix}"
  network      = google_compute_network.network_primary.id
  peer_network = google_compute_network.network_secondary.id

  import_custom_routes = true
  export_custom_routes = true
  import_subnet_routes_with_public_ip = true
  export_subnet_routes_with_public_ip = true
}

resource "google_compute_network_peering" "peering_secondary" {
  name         = "tf-test-secondary-peering%{random_suffix}"
  network      = google_compute_network.network_secondary.id
  peer_network = google_compute_network.network_primary.id
}

resource "google_compute_network" "network_primary" {
  name                    = "tf-test-primary-network%{random_suffix}"
  auto_create_subnetworks = "false"
}

resource "google_compute_network" "network_secondary" {
  name                    = "tf-test-secondary-network%{random_suffix}"
  auto_create_subnetworks = "false"
}
`, context)
}

func TestAccComputeNetworkPeeringRoutesConfig_networkPeeringRoutesConfigGkeExample(t *testing.T) {
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeNetworkPeeringRoutesConfig_networkPeeringRoutesConfigGkeExample(context),
			},
			{
				ResourceName:            "google_compute_network_peering_routes_config.peering_gke_routes",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network"},
			},
		},
	})
}

func testAccComputeNetworkPeeringRoutesConfig_networkPeeringRoutesConfigGkeExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network_peering_routes_config" "peering_gke_routes" {
  peering = google_container_cluster.private_cluster.private_cluster_config[0].peering_name
  network = google_compute_network.container_network.name

  import_custom_routes                = true
  export_custom_routes                = true
  import_subnet_routes_with_public_ip = true
  export_subnet_routes_with_public_ip = true
}

resource "google_compute_network" "container_network" {
  name                    = "tf-test-container-network%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "container_subnetwork" {
  name                     = "tf-test-container-subnetwork%{random_suffix}"
  region                   = "us-central1"
  network                  = google_compute_network.container_network.name
  ip_cidr_range            = "10.0.36.0/24"
  private_ip_google_access = true

  secondary_ip_range {
    range_name    = "pod"
    ip_cidr_range = "10.0.0.0/19"
  }

  secondary_ip_range {
    range_name    = "svc"
    ip_cidr_range = "10.0.32.0/22"
  }
}

resource "google_container_cluster" "private_cluster" {
  name               = "tf-test-private-cluster%{random_suffix}"
  location           = "us-central1-a"
  initial_node_count = 1

  network    = google_compute_network.container_network.name
  subnetwork = google_compute_subnetwork.container_subnetwork.name

  private_cluster_config {
    enable_private_endpoint = true
    enable_private_nodes    = true
    master_ipv4_cidr_block  = "10.42.0.0/28"
  }

  master_authorized_networks_config {}

  ip_allocation_policy {
    cluster_secondary_range_name  = google_compute_subnetwork.container_subnetwork.secondary_ip_range[0].range_name
    services_secondary_range_name = google_compute_subnetwork.container_subnetwork.secondary_ip_range[1].range_name
  }
  deletion_protection  = %{deletion_protection}
}
`, context)
}
