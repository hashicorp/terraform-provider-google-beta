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
//	Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/services/appengine/resource_app_engine_application_test.go
//
//	DO NOT EDIT this file directly. Any changes made to this file will be
//	overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
package appengine_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccAppEngineApplication_basic(t *testing.T) {
	t.Parallel()

	org := envvar.GetTestOrgFromEnv(t)
	pid := fmt.Sprintf("tf-test-%d", acctest.RandInt(t))
	billingAccount := envvar.GetTestBillingAccountFromEnv(t)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccAppEngineApplication_basic(pid, org, billingAccount),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("google_app_engine_application.acceptance", "url_dispatch_rule.#"),
					resource.TestCheckResourceAttrSet("google_app_engine_application.acceptance", "name"),
					resource.TestCheckResourceAttrSet("google_app_engine_application.acceptance", "code_bucket"),
					resource.TestCheckResourceAttrSet("google_app_engine_application.acceptance", "default_hostname"),
					resource.TestCheckResourceAttrSet("google_app_engine_application.acceptance", "default_bucket"),
				),
			},
			{
				ResourceName:      "google_app_engine_application.acceptance",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccAppEngineApplication_update(pid, org, billingAccount),
			},
			{
				ResourceName:      "google_app_engine_application.acceptance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAppEngineApplication_withIAP(t *testing.T) {
	t.Parallel()

	org := envvar.GetTestOrgFromEnv(t)
	pid := fmt.Sprintf("tf-test-%d", acctest.RandInt(t))
	billingAccount := envvar.GetTestBillingAccountFromEnv(t)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccAppEngineApplication_withIAP(pid, org, billingAccount),
			},
			{
				ResourceName:            "google_app_engine_application.acceptance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"iap.0.oauth2_client_secret"},
			},
		},
	})
}

func testAccAppEngineApplication_withIAP(pid, org, billingAccount string) string {
	return fmt.Sprintf(`
resource "google_project" "acceptance" {
  project_id = "%s"
  name       = "%s"
  org_id     = "%s"
  billing_account = "%s"
  deletion_policy = "DELETE"
}

resource "google_app_engine_application" "acceptance" {
  project        = google_project.acceptance.project_id
  auth_domain    = "hashicorptest.com"
  location_id    = "us-central"
  serving_status = "SERVING"

  iap {
    enabled              = false
    oauth2_client_id     = "test"
    oauth2_client_secret = "test"
  }
}
`, pid, pid, org, billingAccount)
}

func testAccAppEngineApplication_basic(pid, org, billingAccount string) string {
	return fmt.Sprintf(`
resource "google_project" "acceptance" {
  project_id = "%s"
  name       = "%s"
  org_id     = "%s"
  billing_account = "%s"
  deletion_policy = "DELETE"
}

resource "google_app_engine_application" "acceptance" {
  project        = google_project.acceptance.project_id
  auth_domain    = "hashicorptest.com"
  location_id    = "us-central"
  database_type  = "CLOUD_DATASTORE_COMPATIBILITY"
  serving_status = "SERVING"
}
`, pid, pid, org, billingAccount)
}

func testAccAppEngineApplication_update(pid, org, billingAccount string) string {
	return fmt.Sprintf(`
resource "google_project" "acceptance" {
  project_id = "%s"
  name       = "%s"
  org_id     = "%s"
  billing_account = "%s"
  deletion_policy = "DELETE"
}

resource "google_app_engine_application" "acceptance" {
  project        = google_project.acceptance.project_id
  auth_domain    = "tf-test.club"
  location_id    = "us-central"
  database_type  = "CLOUD_DATASTORE_COMPATIBILITY"
  serving_status = "USER_DISABLED"
}
`, pid, pid, org, billingAccount)
}
