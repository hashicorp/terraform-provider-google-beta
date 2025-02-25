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
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccComputeDisk_diskBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeDiskDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeDisk_diskBasicExample(context),
			},
			{
				ResourceName:            "google_compute_disk.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"architecture", "interface", "labels", "params", "snapshot", "source_storage_object", "terraform_labels", "type", "zone"},
			},
		},
	})
}

func testAccComputeDisk_diskBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_disk" "default" {
  name  = "tf-test-test-disk%{random_suffix}"
  type  = "pd-ssd"
  zone  = "us-central1-a"
  image = "debian-11-bullseye-v20220719"
  labels = {
    environment = "dev"
  }
  physical_block_size_bytes = 4096
}
`, context)
}

func TestAccComputeDisk_diskAsyncExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeDiskDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeDisk_diskAsyncExample(context),
			},
			{
				ResourceName:            "google_compute_disk.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"architecture", "interface", "labels", "params", "snapshot", "source_storage_object", "terraform_labels", "type", "zone"},
			},
		},
	})
}

func testAccComputeDisk_diskAsyncExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_disk" "primary" {
  name  = "tf-test-async-test-disk%{random_suffix}"
  type  = "pd-ssd"
  zone  = "us-central1-a"

  physical_block_size_bytes = 4096
}

resource "google_compute_disk" "secondary" {
  name  = "tf-test-async-secondary-test-disk%{random_suffix}"
  type  = "pd-ssd"
  zone  = "us-east1-c"

  async_primary_disk {
    disk = google_compute_disk.primary.id
  }

  physical_block_size_bytes = 4096
}
`, context)
}

func TestAccComputeDisk_diskFeaturesExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeDiskDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeDisk_diskFeaturesExample(context),
			},
			{
				ResourceName:            "google_compute_disk.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"architecture", "interface", "labels", "params", "snapshot", "source_storage_object", "terraform_labels", "type", "zone"},
			},
		},
	})
}

func testAccComputeDisk_diskFeaturesExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_disk" "default" {
  name  = "tf-test-test-disk-features%{random_suffix}"
  type  = "pd-ssd"
  zone  = "us-central1-a"
  labels = {
    environment = "dev"
  }

  guest_os_features {
    type = "SECURE_BOOT"
  }

  guest_os_features {
    type = "MULTI_IP_SUBNET"
  }

  guest_os_features {
    type = "WINDOWS"
  }

  licenses = ["https://www.googleapis.com/compute/v1/projects/windows-cloud/global/licenses/windows-server-core"]

  physical_block_size_bytes = 4096
}
`, context)
}

func testAccCheckComputeDiskDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_disk" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/disks/{{name}}")
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
				return fmt.Errorf("ComputeDisk still exists at %s", url)
			}
		}

		return nil
	}
}
