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

package alloydb_test

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

func TestAccAlloydbBackup_alloydbBackupBasicTestExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"network_name":  acctest.BootstrapSharedServiceNetworkingConnection(t, "alloydb-1"),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckAlloydbBackupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccAlloydbBackup_alloydbBackupBasicTestExample(context),
			},
			{
				ResourceName:            "google_alloydb_backup.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations", "backup_id", "labels", "location", "reconciling", "terraform_labels", "update_time"},
			},
		},
	})
}

func testAccAlloydbBackup_alloydbBackupBasicTestExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_alloydb_backup" "default" {
  location     = "us-central1"
  backup_id    = "tf-test-alloydb-backup%{random_suffix}"
  cluster_name = google_alloydb_cluster.default.name

  depends_on = [google_alloydb_instance.default]
}

resource "google_alloydb_cluster" "default" {
  cluster_id = "tf-test-alloydb-cluster%{random_suffix}"
  location   = "us-central1"
  network_config {
    network = data.google_compute_network.default.id
  }
}

resource "google_alloydb_instance" "default" {
  cluster       = google_alloydb_cluster.default.name
  instance_id   = "tf-test-alloydb-instance%{random_suffix}"
  instance_type = "PRIMARY"
}

data "google_compute_network" "default" {
  name = "%{network_name}"
}
`, context)
}

func TestAccAlloydbBackup_alloydbBackupFullTestExample(t *testing.T) {
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"network_name":  acctest.BootstrapSharedServiceNetworkingConnection(t, "alloydb-1"),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckAlloydbBackupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccAlloydbBackup_alloydbBackupFullTestExample(context),
			},
			{
				ResourceName:            "google_alloydb_backup.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations", "backup_id", "labels", "location", "reconciling", "terraform_labels", "update_time"},
			},
		},
	})
}

func testAccAlloydbBackup_alloydbBackupFullTestExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_alloydb_backup" "default" {
  location     = "us-central1"
  backup_id    = "tf-test-alloydb-backup%{random_suffix}"
  cluster_name = google_alloydb_cluster.default.name

  description = "example description"
  type = "ON_DEMAND"
  labels = {
    "label" = "key"
  }
  depends_on = [google_alloydb_instance.default]
}

resource "google_alloydb_cluster" "default" {
  cluster_id = "tf-test-alloydb-cluster%{random_suffix}"
  location   = "us-central1"
  network_config {
    network = data.google_compute_network.default.id
  }
}

resource "google_alloydb_instance" "default" {
  cluster       = google_alloydb_cluster.default.name
  instance_id   = "tf-test-alloydb-instance%{random_suffix}"
  instance_type = "PRIMARY"
}

data "google_compute_network" "default" {
  name = "%{network_name}"
}
`, context)
}

func testAccCheckAlloydbBackupDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_alloydb_backup" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{AlloydbBasePath}}projects/{{project}}/locations/{{location}}/backups/{{backup_id}}")
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
				return fmt.Errorf("AlloydbBackup still exists at %s", url)
			}
		}

		return nil
	}
}
