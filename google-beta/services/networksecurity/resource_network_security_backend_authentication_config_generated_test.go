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

package networksecurity_test

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

func TestAccNetworkSecurityBackendAuthenticationConfig_networkSecurityBackendAuthenticationConfigBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckNetworkSecurityBackendAuthenticationConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecurityBackendAuthenticationConfig_networkSecurityBackendAuthenticationConfigBasicExample(context),
			},
			{
				ResourceName:            "google_network_security_backend_authentication_config.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "location", "name", "terraform_labels"},
			},
		},
	})
}

func testAccNetworkSecurityBackendAuthenticationConfig_networkSecurityBackendAuthenticationConfigBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_network_security_backend_authentication_config" "default" {
  provider = google-beta
  name             = "tf-test-my-backend-authentication-config%{random_suffix}"
  labels           = {
    foo = "bar"
  }
  description      = "my description"
  well_known_roots = "PUBLIC_ROOTS"
}
`, context)
}

func TestAccNetworkSecurityBackendAuthenticationConfig_networkSecurityBackendAuthenticationConfigFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckNetworkSecurityBackendAuthenticationConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecurityBackendAuthenticationConfig_networkSecurityBackendAuthenticationConfigFullExample(context),
			},
			{
				ResourceName:            "google_network_security_backend_authentication_config.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "location", "name", "terraform_labels"},
			},
		},
	})
}

func testAccNetworkSecurityBackendAuthenticationConfig_networkSecurityBackendAuthenticationConfigFullExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_certificate_manager_certificate" "certificate" {
  provider = google-beta
  name     = "tf-test-my-certificate%{random_suffix}"
  labels   = {
    foo = "bar"
  }
  location    = "global"
  self_managed {
    pem_certificate = file("test-fixtures/cert.pem")
    pem_private_key = file("test-fixtures/key.pem")
  }
  scope       = "CLIENT_AUTH"
}

resource "google_certificate_manager_trust_config" "trust_config" {
  provider    = google-beta
  name        = "tf-test-my-trust-config%{random_suffix}"
  description = "sample description for the trust config"
  location    = "global"

  trust_stores {
    trust_anchors { 
      pem_certificate = file("test-fixtures/cert.pem")
    }
    intermediate_cas { 
      pem_certificate = file("test-fixtures/cert.pem")
    }
  }

  labels = {
    foo = "bar"
  }
}

resource "google_network_security_backend_authentication_config" "default" {
  provider = google-beta
  name     = "tf-test-my-backend-authentication-config%{random_suffix}"
  labels   = {
    bar = "foo"
  }
  location           = "global"
  description        = "my description"
  well_known_roots   = "PUBLIC_ROOTS"
  client_certificate = google_certificate_manager_certificate.certificate.id
  trust_config       = google_certificate_manager_trust_config.trust_config.id
}
`, context)
}

func testAccCheckNetworkSecurityBackendAuthenticationConfigDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_network_security_backend_authentication_config" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/backendAuthenticationConfigs/{{name}}")
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
				return fmt.Errorf("NetworkSecurityBackendAuthenticationConfig still exists at %s", url)
			}
		}

		return nil
	}
}
