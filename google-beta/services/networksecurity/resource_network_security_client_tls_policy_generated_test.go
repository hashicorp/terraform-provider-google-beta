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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccNetworkSecurityClientTlsPolicy_networkSecurityClientTlsPolicyBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckNetworkSecurityClientTlsPolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecurityClientTlsPolicy_networkSecurityClientTlsPolicyBasicExample(context),
			},
			{
				ResourceName:            "google_network_security_client_tls_policy.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccNetworkSecurityClientTlsPolicy_networkSecurityClientTlsPolicyBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_network_security_client_tls_policy" "default" {
  provider               = google-beta
  name                   = "tf-test-my-client-tls-policy%{random_suffix}"
  labels                 = {
    foo = "bar"
  }
  description            = "my description"
  sni                    = "secure.example.com"
}
`, context)
}

func TestAccNetworkSecurityClientTlsPolicy_networkSecurityClientTlsPolicyAdvancedExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckNetworkSecurityClientTlsPolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecurityClientTlsPolicy_networkSecurityClientTlsPolicyAdvancedExample(context),
			},
			{
				ResourceName:            "google_network_security_client_tls_policy.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccNetworkSecurityClientTlsPolicy_networkSecurityClientTlsPolicyAdvancedExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_network_security_client_tls_policy" "default" {
  provider               = google-beta
  name                   = "tf-test-my-client-tls-policy%{random_suffix}"
  labels                 = {
    foo = "bar"
  }
  description            = "my description"
  client_certificate {
    certificate_provider_instance {
        plugin_instance = "google_cloud_private_spiffe"
      }
    }
  server_validation_ca {
    grpc_endpoint {
      target_uri = "unix:mypath"
    }
  }
  server_validation_ca {
    grpc_endpoint {
      target_uri = "unix:mypath1"
    }
  }
}
`, context)
}

func testAccCheckNetworkSecurityClientTlsPolicyDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_network_security_client_tls_policy" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/clientTlsPolicies/{{name}}")
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
				return fmt.Errorf("NetworkSecurityClientTlsPolicy still exists at %s", url)
			}
		}

		return nil
	}
}
