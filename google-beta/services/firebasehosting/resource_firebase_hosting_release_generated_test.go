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

package firebasehosting_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccFirebaseHostingRelease_firebasehostingReleaseInSiteExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseHostingRelease_firebasehostingReleaseInSiteExample(context),
			},
			{
				ResourceName:            "google_firebase_hosting_release.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"channel_id", "site_id", "version_name"},
			},
		},
	})
}

func testAccFirebaseHostingRelease_firebasehostingReleaseInSiteExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_firebase_hosting_site" "default" {
  provider = google-beta
  project  = "%{project_id}"
  site_id  = "tf-test-site-id%{random_suffix}"
}

resource "google_firebase_hosting_version" "default" {
  provider = google-beta
  site_id  = google_firebase_hosting_site.default.site_id
  config {
    redirects {
      glob = "/google/**"
      status_code = 302
      location = "https://www.google.com"
    }
  }
}

resource "google_firebase_hosting_release" "default" {
  provider     = google-beta
  site_id      = google_firebase_hosting_site.default.site_id
  version_name = google_firebase_hosting_version.default.name
  message      = "Test release"
}
`, context)
}

func TestAccFirebaseHostingRelease_firebasehostingReleaseInChannelExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseHostingRelease_firebasehostingReleaseInChannelExample(context),
			},
			{
				ResourceName:            "google_firebase_hosting_release.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"channel_id", "site_id", "version_name"},
			},
		},
	})
}

func testAccFirebaseHostingRelease_firebasehostingReleaseInChannelExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_firebase_hosting_site" "default" {
  provider = google-beta
  project  = "%{project_id}"
  site_id  = "tf-test-site-with-channel%{random_suffix}"
}

resource "google_firebase_hosting_version" "default" {
  provider = google-beta
  site_id  = google_firebase_hosting_site.default.site_id
  config {
    redirects {
      glob = "/google/**"
      status_code = 302
      location = "https://www.google.com"
    }
  }
}

resource "google_firebase_hosting_channel" "default" {
  provider   = google-beta
  site_id    = google_firebase_hosting_site.default.site_id
  channel_id = "tf-test-channel-id%{random_suffix}"
}

resource "google_firebase_hosting_release" "default" {
  provider     = google-beta
  site_id      = google_firebase_hosting_site.default.site_id
  channel_id   = google_firebase_hosting_channel.default.channel_id
  version_name = google_firebase_hosting_version.default.name
  message      = "Test release in channel"
}
`, context)
}

func TestAccFirebaseHostingRelease_firebasehostingReleaseDisableExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseHostingRelease_firebasehostingReleaseDisableExample(context),
			},
			{
				ResourceName:            "google_firebase_hosting_release.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"channel_id", "site_id", "version_name"},
			},
		},
	})
}

func testAccFirebaseHostingRelease_firebasehostingReleaseDisableExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_firebase_hosting_site" "default" {
  provider = google-beta
  project  = "%{project_id}"
  site_id  = "tf-test-site-id%{random_suffix}"
}

resource "google_firebase_hosting_release" "default" {
  provider = google-beta
  site_id  = google_firebase_hosting_site.default.site_id
  type     = "SITE_DISABLE"
  message  = "Take down site"
}
`, context)
}
