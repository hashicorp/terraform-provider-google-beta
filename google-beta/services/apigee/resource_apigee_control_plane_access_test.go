// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
// ----------------------------------------------------------------------------
//
//	***     AUTO GENERATED CODE    ***    Type: Handwritten     ***
//
// ----------------------------------------------------------------------------
//
//	This code is generated by Magic Modules using the following:
//
//	Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/services/apigee/resource_apigee_control_plane_access_test.go
//
//	DO NOT EDIT this file directly. Any changes made to this file will be
//	overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
package apigee_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccApigeeControlPlaneAccess_update(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"billing_account": envvar.GetTestBillingAccountFromEnv(t),
		"org_id":          envvar.GetTestOrgFromEnv(t),
		"random_suffix":   acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccApigeeControlPlaneAccess_full(context),
			},
			{
				ResourceName:            "google_apigee_control_plane_access.apigee_control_plane_access",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name"},
			},
			{
				Config: testAccApigeeControlPlaneAccess_update(context),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("google_apigee_control_plane_access.apigee_control_plane_access", plancheck.ResourceActionUpdate),
					},
				},
			},
			{
				ResourceName:            "google_apigee_control_plane_access.apigee_control_plane_access",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name"},
			},
		},
	})
}

func testAccApigeeControlPlaneAccess_full(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_project" "project" {
  project_id      = "tf-test-my-project%{random_suffix}"
  name            = "tf-test-my-project%{random_suffix}"
  org_id          = "%{org_id}"
  billing_account = "%{billing_account}"
  deletion_policy = "DELETE"
}

resource "google_project_service" "apigee" {
  project = google_project.project.project_id
  service = "apigee.googleapis.com"
}

resource "google_apigee_organization" "apigee_org" {
  analytics_region   = "us-central1"
  project_id         = google_project.project.project_id

  runtime_type       = "HYBRID"
  depends_on         = [google_project_service.apigee]
}

resource "google_service_account" "service_account" {
  account_id   = "sa-%{random_suffix}"
  display_name = "Service Account"
}

resource "google_apigee_control_plane_access" "apigee_control_plane_access" {
  name       = google_apigee_organization.apigee_org.name
  synchronizer_identities = [
    "serviceAccount:${google_service_account.service_account.email}",
  ]
  analytics_publisher_identities = [
    "serviceAccount:${google_service_account.service_account.email}",
  ]
}
`, context)
}

func testAccApigeeControlPlaneAccess_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_project" "project" {
  project_id      = "tf-test-my-project%{random_suffix}"
  name            = "tf-test-my-project%{random_suffix}"
  org_id          = "%{org_id}"
  billing_account = "%{billing_account}"
  deletion_policy = "DELETE"
}

resource "google_project_service" "apigee" {
  project = google_project.project.project_id
  service = "apigee.googleapis.com"
}

resource "google_apigee_organization" "apigee_org" {
  analytics_region   = "us-central1"
  project_id         = google_project.project.project_id

  runtime_type       = "HYBRID"
  depends_on         = [google_project_service.apigee]
}

resource "google_apigee_control_plane_access" "apigee_control_plane_access" {
  name       = google_apigee_organization.apigee_org.name
  synchronizer_identities = []
  analytics_publisher_identities = []
}
`, context)
}
