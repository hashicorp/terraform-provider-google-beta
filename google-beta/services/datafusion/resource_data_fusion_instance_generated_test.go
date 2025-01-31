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

package datafusion_test

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

func TestAccDataFusionInstance_dataFusionInstanceBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"prober_test_run": `options = { prober_test_run = "true" }`,
		"random_suffix":   acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataFusionInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataFusionInstance_dataFusionInstanceBasicExample(context),
			},
			{
				ResourceName:            "google_data_fusion_instance.basic_instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "network_config.0.private_service_connect_config.0.unreachable_cidr_block", "region", "terraform_labels"},
			},
		},
	})
}

func testAccDataFusionInstance_dataFusionInstanceBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_data_fusion_instance" "basic_instance" {
  name   = "tf-test-my-instance%{random_suffix}"
  region = "us-central1"
  type   = "BASIC"
  %{prober_test_run}
}
`, context)
}

func TestAccDataFusionInstance_dataFusionInstanceFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"prober_test_run": `options = { prober_test_run = "true" }`,
		"random_suffix":   acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataFusionInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataFusionInstance_dataFusionInstanceFullExample(context),
			},
			{
				ResourceName:            "google_data_fusion_instance.extended_instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "network_config.0.private_service_connect_config.0.unreachable_cidr_block", "region", "terraform_labels"},
			},
		},
	})
}

func testAccDataFusionInstance_dataFusionInstanceFullExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_data_fusion_instance" "extended_instance" {
  name                          = "tf-test-my-instance%{random_suffix}"
  description                   = "My Data Fusion instance"
  display_name                  = "My Data Fusion instance"
  region                        = "us-central1"
  type                          = "BASIC"
  enable_stackdriver_logging    = true
  enable_stackdriver_monitoring = true
  private_instance              = true
  dataproc_service_account      = data.google_app_engine_default_service_account.default.email

  labels = {
    example_key = "example_value"
  }

  network_config {
    network       = "default"
    ip_allocation = "${google_compute_global_address.private_ip_alloc.address}/${google_compute_global_address.private_ip_alloc.prefix_length}"
  }

  accelerators {
    accelerator_type = "CDC"
    state = "ENABLED"
  }
  %{prober_test_run}
}

data "google_app_engine_default_service_account" "default" {
}

resource "google_compute_network" "network" {
  name = "tf-test-datafusion-full-network%{random_suffix}"
}

resource "google_compute_global_address" "private_ip_alloc" {
  name          = "tf-test-datafusion-ip-alloc%{random_suffix}"
  address_type  = "INTERNAL"
  purpose       = "VPC_PEERING"
  prefix_length = 22
  network       = google_compute_network.network.id
}
`, context)
}

func TestAccDataFusionInstance_dataFusionInstancePscExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"prober_test_run": `options = { prober_test_run = "true" }`,
		"random_suffix":   acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataFusionInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataFusionInstance_dataFusionInstancePscExample(context),
			},
			{
				ResourceName:            "google_data_fusion_instance.psc_instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "network_config.0.private_service_connect_config.0.unreachable_cidr_block", "region", "terraform_labels"},
			},
		},
	})
}

func testAccDataFusionInstance_dataFusionInstancePscExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_data_fusion_instance" "psc_instance" {
  name             = "tf-test-psc-instance%{random_suffix}"
  region           = "us-central1"
  type             = "BASIC"
  private_instance = true

  network_config {
    connection_type = "PRIVATE_SERVICE_CONNECT_INTERFACES"
    private_service_connect_config {
      network_attachment     = google_compute_network_attachment.psc.id
      unreachable_cidr_block = "192.168.0.0/25"
    }
  }

  %{prober_test_run}
}

resource "google_compute_network" "psc" {
  name                    = "tf-test-datafusion-psc-network%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "psc" {
  name   = "tf-test-datafusion-psc-subnet%{random_suffix}"
  region = "us-central1"

  network       = google_compute_network.psc.id
  ip_cidr_range = "10.0.0.0/16"
}

resource "google_compute_network_attachment" "psc" {
  name                  = "tf-test-datafusion-psc-attachment%{random_suffix}"
  region                = "us-central1"
  connection_preference = "ACCEPT_AUTOMATIC"

  subnetworks = [
    google_compute_subnetwork.psc.self_link
  ]
}
`, context)
}

func TestAccDataFusionInstance_dataFusionInstanceCmekExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataFusionInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataFusionInstance_dataFusionInstanceCmekExample(context),
			},
			{
				ResourceName:            "google_data_fusion_instance.cmek",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "network_config.0.private_service_connect_config.0.unreachable_cidr_block", "region", "terraform_labels"},
			},
		},
	})
}

func testAccDataFusionInstance_dataFusionInstanceCmekExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_data_fusion_instance" "cmek" {
  name   = "tf-test-my-instance%{random_suffix}"
  region = "us-central1"
  type   = "BASIC"

  crypto_key_config {
    key_reference = google_kms_crypto_key.crypto_key.id
  }

  depends_on = [google_kms_crypto_key_iam_member.crypto_key_member]
}

resource "google_kms_crypto_key" "crypto_key" {
  name     = "tf-test-my-instance%{random_suffix}"
  key_ring = google_kms_key_ring.key_ring.id
}

resource "google_kms_key_ring" "key_ring" {
  name     = "tf-test-my-instance%{random_suffix}"
  location = "us-central1"
}

resource "google_kms_crypto_key_iam_member" "crypto_key_member" {
  crypto_key_id = google_kms_crypto_key.crypto_key.id
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"

  member = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-datafusion.iam.gserviceaccount.com"
}

data "google_project" "project" {}
`, context)
}

func TestAccDataFusionInstance_dataFusionInstanceEnterpriseExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"prober_test_run": `options = { prober_test_run = "true" }`,
		"random_suffix":   acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataFusionInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataFusionInstance_dataFusionInstanceEnterpriseExample(context),
			},
			{
				ResourceName:            "google_data_fusion_instance.enterprise_instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "network_config.0.private_service_connect_config.0.unreachable_cidr_block", "region", "terraform_labels"},
			},
		},
	})
}

func testAccDataFusionInstance_dataFusionInstanceEnterpriseExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_data_fusion_instance" "enterprise_instance" {
  name = "tf-test-my-instance%{random_suffix}"
  region = "us-central1"
  type = "ENTERPRISE"
  enable_rbac = true
  %{prober_test_run}
}
`, context)
}

func TestAccDataFusionInstance_dataFusionInstanceEventExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataFusionInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataFusionInstance_dataFusionInstanceEventExample(context),
			},
			{
				ResourceName:            "google_data_fusion_instance.event",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "network_config.0.private_service_connect_config.0.unreachable_cidr_block", "region", "terraform_labels"},
			},
		},
	})
}

func testAccDataFusionInstance_dataFusionInstanceEventExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_data_fusion_instance" "event" {
  name    = "tf-test-my-instance%{random_suffix}"
  region  = "us-central1"
  type    = "BASIC"

  event_publish_config {
    enabled = true
    topic   = google_pubsub_topic.event.id
  }
}

resource "google_pubsub_topic" "event" {
  name = "tf-test-my-instance%{random_suffix}"
}
`, context)
}

func TestAccDataFusionInstance_dataFusionInstanceZoneExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataFusionInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataFusionInstance_dataFusionInstanceZoneExample(context),
			},
			{
				ResourceName:            "google_data_fusion_instance.zone",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "network_config.0.private_service_connect_config.0.unreachable_cidr_block", "region", "terraform_labels"},
			},
		},
	})
}

func testAccDataFusionInstance_dataFusionInstanceZoneExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_data_fusion_instance" "zone" {
  name   = "tf-test-my-instance%{random_suffix}"
  region = "us-central1"
  zone   = "us-central1-a"
  type   = "DEVELOPER"
}
`, context)
}

func testAccCheckDataFusionInstanceDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_data_fusion_instance" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{DataFusionBasePath}}projects/{{project}}/locations/{{region}}/instances/{{name}}")
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
				return fmt.Errorf("DataFusionInstance still exists at %s", url)
			}
		}

		return nil
	}
}
