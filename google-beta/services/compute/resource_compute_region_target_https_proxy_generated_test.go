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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccComputeRegionTargetHttpsProxy_regionTargetHttpsProxyBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionTargetHttpsProxyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionTargetHttpsProxy_regionTargetHttpsProxyBasicExample(context),
			},
			{
				ResourceName:            "google_compute_region_target_https_proxy.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ssl_policy", "url_map", "region"},
			},
		},
	})
}

func testAccComputeRegionTargetHttpsProxy_regionTargetHttpsProxyBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_target_https_proxy" "default" {
  region           = "us-central1"
  name             = "tf-test-test-proxy%{random_suffix}"
  url_map          = google_compute_region_url_map.default.id
  ssl_certificates = [google_compute_region_ssl_certificate.default.id]
}

resource "google_compute_region_ssl_certificate" "default" {
  region      = "us-central1"
  name        = "tf-test-my-certificate%{random_suffix}"
  private_key = file("test-fixtures/test.key")
  certificate = file("test-fixtures/test.crt")
}

resource "google_compute_region_url_map" "default" {
  region      = "us-central1"
  name        = "tf-test-url-map%{random_suffix}"
  description = "a description"

  default_service = google_compute_region_backend_service.default.id

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name            = "allpaths"
    default_service = google_compute_region_backend_service.default.id

    path_rule {
      paths   = ["/*"]
      service = google_compute_region_backend_service.default.id
    }
  }
}

resource "google_compute_region_backend_service" "default" {
  region      = "us-central1"
  name        = "tf-test-backend-service%{random_suffix}"
  protocol    = "HTTP"
  load_balancing_scheme = "INTERNAL_MANAGED"
  timeout_sec = 10

  health_checks = [google_compute_region_health_check.default.id]
}

resource "google_compute_region_health_check" "default" {
  region = "us-central1"
  name   = "tf-test-http-health-check%{random_suffix}"
  http_health_check {
    port = 80
  }
}
`, context)
}

func TestAccComputeRegionTargetHttpsProxy_regionTargetHttpsProxyCertificateManagerCertificateExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionTargetHttpsProxyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionTargetHttpsProxy_regionTargetHttpsProxyCertificateManagerCertificateExample(context),
			},
			{
				ResourceName:            "google_compute_region_target_https_proxy.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ssl_policy", "url_map", "region"},
			},
		},
	})
}

func testAccComputeRegionTargetHttpsProxy_regionTargetHttpsProxyCertificateManagerCertificateExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_target_https_proxy" "default" {
  name                             = "tf-test-target-http-proxy%{random_suffix}"
  url_map                          = google_compute_region_url_map.default.id
  certificate_manager_certificates =  ["//certificatemanager.googleapis.com/${google_certificate_manager_certificate.default.id}"] # [google_certificate_manager_certificate.default.id] is also acceptable
}

resource "google_certificate_manager_certificate" "default" {
  name              = "tf-test-my-certificate%{random_suffix}"
  location          = "us-central1"
  self_managed {
    pem_certificate = file("test-fixtures/cert.pem")
    pem_private_key = file("test-fixtures/private-key.pem")                                                                                                                
  }
}

resource "google_compute_region_url_map" "default" {
  name            = "tf-test-url-map%{random_suffix}"
  default_service = google_compute_region_backend_service.default.id
  region          = "us-central1"
}

resource "google_compute_region_backend_service" "default" {
  name                  = "tf-test-backend-service%{random_suffix}"
  region                = "us-central1"
  protocol              = "HTTPS"
  timeout_sec           = 30
  load_balancing_scheme = "INTERNAL_MANAGED"
}
`, context)
}

func testAccCheckComputeRegionTargetHttpsProxyDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_region_target_https_proxy" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/targetHttpsProxies/{{name}}")
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
				return fmt.Errorf("ComputeRegionTargetHttpsProxy still exists at %s", url)
			}
		}

		return nil
	}
}
