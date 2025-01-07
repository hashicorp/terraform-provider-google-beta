// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package cloudquotas_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccCloudQuotasQuotaAdjusterSettings_cloudQuotasQuotaAdjusterSettingsUpdate(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"enablement":         "ENABLED",
		"org_id":             envvar.GetTestOrgFromEnv(t),
		"billing_account":    envvar.GetTestBillingAccountFromEnv(t),
		"updated_enablement": "DISABLED",
		"random_suffix":      acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudQuotasQuotaAdjusterSettings_basic(context),
			},
			{
				ResourceName:            "google_cloud_quotas_quota_adjuster_settings.adjuster_settings",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent"},
			},
			{
				Config: testAccCloudQuotasQuotaAdjusterSettings_updateEnabelement(context),
			},
			{
				ResourceName:            "google_cloud_quotas_quota_adjuster_settings.adjuster_settings",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent"},
			},
		},
	})
}

func testAccCloudQuotasQuotaAdjusterSettings_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
	data "google_project" "project" {
		provider      = google-beta
	}

	resource "google_project_service" "cloudquotas" {
		provider      = google-beta
		project = data.google_project.project.project_id
		service 	= "cloudquotas.googleapis.com"
	}

	resource "google_cloud_quotas_quota_adjuster_settings" "adjuster_settings" {
		provider      = google-beta
		parent        = "projects/${data.google_project.project.number}"
		enablement    = "%{enablement}"
		depends_on = [google_project_service.cloudquotas]
	}
`, context)
}

func testAccCloudQuotasQuotaAdjusterSettings_updateEnabelement(context map[string]interface{}) string {
	return acctest.Nprintf(`
	data "google_project" "project" {
		provider      = google-beta
	}

	resource "google_project_service" "cloudquotas" {
		provider      = google-beta
		project = data.google_project.project.project_id
		service 	= "cloudquotas.googleapis.com"
	}	

	resource "google_cloud_quotas_quota_adjuster_settings" "adjuster_settings" {
		provider      = google-beta
		parent        = "projects/${data.google_project.project.number}"
		enablement    = "%{updated_enablement}"
		depends_on = [google_project_service.cloudquotas]
	}
`, context)
}
