// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package firebasehosting_test

import (
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccFirebaseHostingCustomDomain_firebasehostingCustomdomainUpdate(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    envvar.GetTestProjectFromEnv(),
		"site_id":       envvar.GetTestProjectFromEnv(),
		"custom_domain": "update.source.domain.com",
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckFirebaseHostingCustomDomainDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseHostingCustomDomain_firebasehostingCustomdomainBeforeUpdate(context),
			},
			{
				ResourceName:            "google_firebase_hosting_custom_domain.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"site_id", "custom_domain", "wait_dns_verification"},
			},
			{
				Config: testAccFirebaseHostingCustomDomain_firebasehostingCustomdomainAfterUpdate(context),
			},
			{
				ResourceName:            "google_firebase_hosting_custom_domain.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"site_id", "custom_domain", "wait_dns_verification"},
			},
		},
	})
}

func testAccFirebaseHostingCustomDomain_firebasehostingCustomdomainBeforeUpdate(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_firebase_hosting_custom_domain" "default" {
  provider = google-beta

  project  = "%{project_id}"
  site_id = "%{site_id}"
  custom_domain = "%{custom_domain}"
  cert_preference = "GROUPED"
  redirect_target = "destination.domain.com"

  wait_dns_verification = false
}
`, context)
}

func testAccFirebaseHostingCustomDomain_firebasehostingCustomdomainAfterUpdate(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_firebase_hosting_custom_domain" "default" {
  provider = google-beta

  project  = "%{project_id}"
  site_id = "%{site_id}"
  custom_domain = "%{custom_domain}"
  cert_preference = "PROJECT_GROUPED"
  redirect_target = "destination2.domain.com"

  wait_dns_verification = false
}
`, context)
}
