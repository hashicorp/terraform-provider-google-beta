// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package apigee_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccApigeeEnvironment_apigeeEnvironmentPatchUpdateTestExampleUpdate(t *testing.T) {
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"org_id":          envvar.GetTestOrgFromEnv(t),
		"billing_account": envvar.GetTestBillingAccountFromEnv(t),
		"random_suffix":   acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckApigeeEnvironmentDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccApigeeEnvironment_apigeeEnvironmentPatchUpdateTestExample(context),
			},
			{
				ResourceName:            "google_apigee_environment.apigee_environment",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"org_id"},
			},
			{
				Config: testAccApigeeEnvironment_apigeeEnvironmentPatchUpdateTestExampleUpdate(context),
			},
			{
				ResourceName:            "google_apigee_environment.apigee_environment",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"org_id"},
			},
		},
	})
}

func testAccApigeeEnvironment_apigeeEnvironmentPatchUpdateTestExampleUpdate(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_project" "project" {
  provider = google-beta

  project_id      = "tf-test%{random_suffix}"
  name            = "tf-test%{random_suffix}"
  org_id          = "%{org_id}"
  billing_account = "%{billing_account}"
  deletion_policy = "DELETE"
}

resource "google_project_service" "apigee" {
  provider = google-beta

  project = google_project.project.project_id
  service = "apigee.googleapis.com"
}

resource "google_project_service" "compute" {
  provider = google-beta

  project = google_project.project.project_id
  service = "compute.googleapis.com"
}

resource "google_project_service" "servicenetworking" {
  provider = google-beta

  project = google_project.project.project_id
  service = "servicenetworking.googleapis.com"
}

resource "google_project_service" "kms" {
  provider = google-beta

  project = google_project.project.project_id
  service = "cloudkms.googleapis.com"
}

resource "google_compute_network" "apigee_network" {
  provider = google-beta

  name       = "apigee-network"
  project    = google_project.project.project_id
  depends_on = [google_project_service.compute]
}

resource "google_compute_global_address" "apigee_range" {
  provider = google-beta

  name          = "tf-test-apigee-range%{random_suffix}"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = google_compute_network.apigee_network.id
  project       = google_project.project.project_id
}

resource "google_service_networking_connection" "apigee_vpc_connection" {
  provider = google-beta

  network                 = google_compute_network.apigee_network.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.apigee_range.name]
  depends_on              = [google_project_service.servicenetworking]
}

resource "google_kms_key_ring" "apigee_keyring" {
  provider = google-beta

  name       = "apigee-keyring"
  location   = "us-central1"
  project    = google_project.project.project_id
  depends_on = [google_project_service.kms]
}

resource "google_kms_crypto_key" "apigee_key" {
  provider = google-beta

  name            = "apigee-key"
  key_ring        = google_kms_key_ring.apigee_keyring.id
}

resource "google_project_service_identity" "apigee_sa" {
  provider = google-beta

  project = google_project.project.project_id
  service = google_project_service.apigee.service
}

resource "google_kms_crypto_key_iam_member" "apigee_sa_keyuser" {
  provider = google-beta

  crypto_key_id = google_kms_crypto_key.apigee_key.id
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"

  member = google_project_service_identity.apigee_sa.member
}

resource "google_apigee_organization" "apigee_org" {
  provider = google-beta

  analytics_region                     = "us-central1"
  project_id                           = google_project.project.project_id
  authorized_network                   = google_compute_network.apigee_network.id
  billing_type                         = "PAYG"
  runtime_database_encryption_key_name = google_kms_crypto_key.apigee_key.id

  depends_on = [
    google_service_networking_connection.apigee_vpc_connection,
    google_project_service.apigee,
    google_kms_crypto_key_iam_member.apigee_sa_keyuser,
  ]
}

resource "google_apigee_environment" "apigee_environment" {
  provider = google-beta

  org_id       = google_apigee_organization.apigee_org.id
  name         = "tf-test%{random_suffix}"
  description  = "Apigee Environment"
  display_name = "tf-test%{random_suffix}"
  type         = "INTERMEDIATE"
  forward_proxy_uri = "http://test:456"
}
`, context)
}
