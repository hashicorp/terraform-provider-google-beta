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

package firebaseappcheck_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccFirebaseAppCheckServiceConfig_firebaseAppCheckServiceConfigOffExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    envvar.GetTestProjectFromEnv(),
		"service_id":    "firestore.googleapis.com",
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckFirebaseAppCheckServiceConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseAppCheckServiceConfig_firebaseAppCheckServiceConfigOffExample(context),
			},
			{
				ResourceName:            "google_firebase_app_check_service_config.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"service_id"},
			},
		},
	})
}

func testAccFirebaseAppCheckServiceConfig_firebaseAppCheckServiceConfigOffExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_project_service" "appcheck" {
  project = "%{project_id}"
  service = "firebaseappcheck.googleapis.com"
  disable_on_destroy = false
}

resource "google_firebase_app_check_service_config" "default" {
  project = "%{project_id}"
  service_id = "%{service_id}"

  depends_on = [google_project_service.appcheck]
}
`, context)
}

func TestAccFirebaseAppCheckServiceConfig_firebaseAppCheckServiceConfigEnforcedExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    envvar.GetTestProjectFromEnv(),
		"service_id":    "firebasestorage.googleapis.com",
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckFirebaseAppCheckServiceConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseAppCheckServiceConfig_firebaseAppCheckServiceConfigEnforcedExample(context),
			},
			{
				ResourceName:            "google_firebase_app_check_service_config.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"service_id"},
			},
		},
	})
}

func testAccFirebaseAppCheckServiceConfig_firebaseAppCheckServiceConfigEnforcedExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_project_service" "appcheck" {
  project = "%{project_id}"
  service = "firebaseappcheck.googleapis.com"
  disable_on_destroy = false
}

resource "google_firebase_app_check_service_config" "default" {
  project = "%{project_id}"
  service_id = "%{service_id}"
  enforcement_mode = "ENFORCED"

  depends_on = [google_project_service.appcheck]
}
`, context)
}

func TestAccFirebaseAppCheckServiceConfig_firebaseAppCheckServiceConfigUnenforcedExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    envvar.GetTestProjectFromEnv(),
		"service_id":    "identitytoolkit.googleapis.com",
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckFirebaseAppCheckServiceConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseAppCheckServiceConfig_firebaseAppCheckServiceConfigUnenforcedExample(context),
			},
			{
				ResourceName:            "google_firebase_app_check_service_config.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"service_id"},
			},
		},
	})
}

func testAccFirebaseAppCheckServiceConfig_firebaseAppCheckServiceConfigUnenforcedExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_project_service" "appcheck" {
  project = "%{project_id}"
  service = "firebaseappcheck.googleapis.com"
  disable_on_destroy = false
}

resource "google_firebase_app_check_service_config" "default" {
  project = "%{project_id}"
  service_id = "%{service_id}"
  enforcement_mode = "UNENFORCED"

  depends_on = [google_project_service.appcheck]
}
`, context)
}

func testAccCheckFirebaseAppCheckServiceConfigDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_firebase_app_check_service_config" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{FirebaseAppCheckBasePath}}projects/{{project}}/services/{{service_id}}")
			if err != nil {
				return err
			}

			billingProject := envvar.GetTestProjectFromEnv()

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err != nil {
				return err
			}

			// empty enforcementMode is equivalent to absence.
			if v := res["enforcementMode"]; v != nil {
				return fmt.Errorf("FirebaseAppCheckServiceConfig still exists at %s", url)
			}
		}

		return nil
	}
}
