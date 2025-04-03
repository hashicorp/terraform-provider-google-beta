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

package netapp_test

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

func TestAccNetappVolumeReplication_netappVolumeReplicationCreateExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"network_name":  acctest.BootstrapSharedServiceNetworkingConnection(t, "gcnv-network-config-1", acctest.ServiceNetworkWithParentService("netapp.servicenetworking.goog")),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetappVolumeReplicationDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetappVolumeReplication_netappVolumeReplicationCreateExample(context),
			},
			{
				ResourceName:            "google_netapp_volume_replication.test_replication",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"delete_destination_volume", "destination_volume_parameters", "force_stopping", "labels", "location", "name", "replication_enabled", "terraform_labels", "volume_name", "wait_for_mirror"},
			},
		},
	})
}

func testAccNetappVolumeReplication_netappVolumeReplicationCreateExample(context map[string]interface{}) string {
	return acctest.Nprintf(`

data "google_compute_network" "default" {
  name = "%{network_name}"
}

resource "google_netapp_storage_pool" "source_pool" {
  name          = "tf-test-source-pool%{random_suffix}"
  location      = "us-central1"
  service_level = "PREMIUM"
  capacity_gib  = 2048
  network       = data.google_compute_network.default.id
}

resource "google_netapp_storage_pool" "destination_pool" {
  name          = "tf-test-destination-pool%{random_suffix}"
  location      = "us-west2"
  service_level = "PREMIUM"
  capacity_gib  = 2048
  network       = data.google_compute_network.default.id
  allow_auto_tiering = true
}

resource "google_netapp_volume" "source_volume" {
  location     = google_netapp_storage_pool.source_pool.location
  name         = "tf-test-source-volume%{random_suffix}"
  capacity_gib = 100
  share_name   = "tf-test-source-volume%{random_suffix}"
  storage_pool = google_netapp_storage_pool.source_pool.name
  protocols = [
    "NFSV3"
  ]
  deletion_policy = "FORCE"
}

resource "google_netapp_volume_replication" "test_replication" {
  depends_on           = [google_netapp_volume.source_volume]
  location             = google_netapp_volume.source_volume.location
  volume_name          = google_netapp_volume.source_volume.name
  name                 = "tf-test-test-replication%{random_suffix}"
  replication_schedule = "EVERY_10_MINUTES"
  description          = "This is a replication resource"
  destination_volume_parameters {
    storage_pool = google_netapp_storage_pool.destination_pool.id
    volume_id    = "tf-test-destination-volume%{random_suffix}"
    # Keeping the share_name of source and destination the same
    # simplifies implementing client failover concepts
    share_name  = "tf-test-source-volume%{random_suffix}"
    description = "This is a replicated volume"
    tiering_policy {
        cooling_threshold_days = 20
        tier_action = "ENABLED"
    }
  }
  # WARNING: Setting delete_destination_volume to true, will delete the
  # CURRENT destination volume if the replication is deleted. Omit the field 
  # or set delete_destination_volume=false to avoid accidental volume deletion.
  delete_destination_volume = true
  wait_for_mirror = true
}
`, context)
}

func testAccCheckNetappVolumeReplicationDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_netapp_volume_replication" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{NetappBasePath}}projects/{{project}}/locations/{{location}}/volumes/{{volume_name}}/replications/{{name}}")
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
				return fmt.Errorf("NetappVolumeReplication still exists at %s", url)
			}
		}

		return nil
	}
}
