// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package parallelstore_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccParallelstoreInstance_parallelstoreInstanceBasicExample_update(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckParallelstoreInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccParallelstoreInstance_parallelstoreInstanceBasicExample_basic(context),
			},
			{
				ResourceName:            "google_parallelstore_instance.instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "instance_id", "labels", "terraform_labels"},
			},
			{
				Config: testAccParallelstoreInstance_parallelstoreInstanceBasicExample_update(context),
			},
			{
				ResourceName:            "google_parallelstore_instance.instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "instance_id", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccParallelstoreInstance_parallelstoreInstanceBasicExample_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_parallelstore_instance" "instance" {
  instance_id = "instance%{random_suffix}"
  location = "us-central1-a"
  description = "test instance"
  capacity_gib = 12000
  network = google_compute_network.network.name
  reserved_ip_range = google_compute_global_address.private_ip_alloc.name
  labels = {
    test = "value"
  }
  provider = google-beta
  depends_on = [google_service_networking_connection.default]
}

resource "google_compute_network" "network" {
  name                    = "network%{random_suffix}"
  auto_create_subnetworks = true
  mtu = 8896
  provider = google-beta
}



# Create an IP address
resource "google_compute_global_address" "private_ip_alloc" {
  name          = "address%{random_suffix}"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 24
  network       = google_compute_network.network.id
  provider = google-beta
}

# Create a private connection
resource "google_service_networking_connection" "default" {
  network                 = google_compute_network.network.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.private_ip_alloc.name]
  provider = google-beta
}
`, context)
}

func testAccParallelstoreInstance_parallelstoreInstanceBasicExample_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_parallelstore_instance" "instance" {
  instance_id = "instance%{random_suffix}"
  location = "us-central1-a"
  description = "test instance updated"
  capacity_gib = 12000
  network = google_compute_network.network.name

  labels = {
    test = "value23"
  }
  provider = google-beta
  depends_on = [google_service_networking_connection.default]
}

resource "google_compute_network" "network" {
  name                    = "network%{random_suffix}"
  auto_create_subnetworks = true
  mtu = 8896
  provider = google-beta
}



# Create an IP address
resource "google_compute_global_address" "private_ip_alloc" {
  name          = "address%{random_suffix}"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 24
  network       = google_compute_network.network.id
  provider = google-beta
}

# Create a private connection
resource "google_service_networking_connection" "default" {
  network                 = google_compute_network.network.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.private_ip_alloc.name]
  provider = google-beta
}
`, context)
}
