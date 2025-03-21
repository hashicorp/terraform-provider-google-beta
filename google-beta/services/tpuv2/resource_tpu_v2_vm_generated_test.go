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

package tpuv2_test

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

func TestAccTpuV2Vm_tpuV2VmBasicExample(t *testing.T) {
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckTpuV2VmDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccTpuV2Vm_tpuV2VmBasicExample(context),
			},
			{
				ResourceName:            "google_tpu_v2_vm.tpu",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels", "zone"},
			},
		},
	})
}

func testAccTpuV2Vm_tpuV2VmBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_tpu_v2_runtime_versions" "available" {
  provider = google-beta
}

resource "google_tpu_v2_vm" "tpu" {
  provider = google-beta

  name = "tf-test-test-tpu%{random_suffix}"
  zone = "us-central1-c"

  runtime_version = "tpu-vm-tf-2.13.0"
}
`, context)
}

func TestAccTpuV2Vm_tpuV2VmFullExample(t *testing.T) {
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		CheckDestroy: testAccCheckTpuV2VmDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccTpuV2Vm_tpuV2VmFullExample(context),
			},
			{
				ResourceName:            "google_tpu_v2_vm.tpu",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels", "zone"},
			},
		},
	})
}

func testAccTpuV2Vm_tpuV2VmFullExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_tpu_v2_runtime_versions" "available" {
  provider = google-beta
}

data "google_tpu_v2_accelerator_types" "available" {
  provider = google-beta
}

resource "google_tpu_v2_vm" "tpu" {
  provider = google-beta

  name = "tf-test-test-tpu%{random_suffix}"
  zone = "us-central1-c"
  description = "Text description of the TPU."

  runtime_version  = "tpu-vm-tf-2.13.0"

  accelerator_config {
    type     = "V2"
    topology = "2x2"
  }

  cidr_block = "10.0.0.0/29"

  network_config {
    can_ip_forward      = true
    enable_external_ips = true
    network             = google_compute_network.network.id
    subnetwork          = google_compute_subnetwork.subnet.id
    queue_count         = 32
  }
  
  scheduling_config {
    preemptible = true
    spot = true
  }

  shielded_instance_config {
    enable_secure_boot = true
  }

  service_account {
    email = google_service_account.sa.email
    scope = [
      "https://www.googleapis.com/auth/cloud-platform",
    ]
  }

  data_disks {
    source_disk = google_compute_disk.disk.id
    mode        = "READ_ONLY"
  }

  labels = {
    foo = "bar"
  }

  metadata = {
    foo = "bar"
  }

  tags = ["foo"]

  depends_on = [time_sleep.wait_60_seconds]
}

resource "google_compute_subnetwork" "subnet" {
  provider = google-beta

  name          = "tf-test-tpu-subnet%{random_suffix}"
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.network.id
}

resource "google_compute_network" "network" {
  provider = google-beta

  name                    = "tf-test-tpu-net%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_service_account" "sa" {
  provider = google-beta

  account_id   = "tf-test-tpu-sa%{random_suffix}"
  display_name = "Test TPU VM"
}

resource "google_compute_disk" "disk" {
  provider = google-beta

  name  = "tf-test-tpu-disk%{random_suffix}"
  image = "debian-cloud/debian-11"
  size  = 10
  type  = "pd-ssd"
  zone  = "us-central1-c"
}

# Wait after service account creation to limit eventual consistency errors.
resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_service_account.sa]

  create_duration = "60s"
}
`, context)
}

func testAccCheckTpuV2VmDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_tpu_v2_vm" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{TpuV2BasePath}}projects/{{project}}/locations/{{zone}}/nodes/{{name}}")
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
				return fmt.Errorf("TpuV2Vm still exists at %s", url)
			}
		}

		return nil
	}
}
