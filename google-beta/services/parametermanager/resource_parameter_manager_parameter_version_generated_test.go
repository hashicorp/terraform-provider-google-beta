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

package parametermanager_test

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

func TestAccParameterManagerParameterVersion_parameterVersionBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckParameterManagerParameterVersionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccParameterManagerParameterVersion_parameterVersionBasicExample(context),
			},
			{
				ResourceName:            "google_parameter_manager_parameter_version.parameter-version-basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parameter", "parameter_version_id"},
			},
		},
	})
}

func testAccParameterManagerParameterVersion_parameterVersionBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_parameter_manager_parameter" "parameter-basic" {
  provider = google-beta
  parameter_id = "parameter%{random_suffix}"
}

resource "google_parameter_manager_parameter_version" "parameter-version-basic" {
  provider = google-beta
  parameter = google_parameter_manager_parameter.parameter-basic.id
  parameter_version_id = "tf_test_parameter_version%{random_suffix}"
  parameter_data = "app-parameter-version-data"
}
`, context)
}

func TestAccParameterManagerParameterVersion_parameterVersionWithJsonFormatExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckParameterManagerParameterVersionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccParameterManagerParameterVersion_parameterVersionWithJsonFormatExample(context),
			},
			{
				ResourceName:            "google_parameter_manager_parameter_version.parameter-version-with-json-format",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parameter", "parameter_version_id"},
			},
		},
	})
}

func testAccParameterManagerParameterVersion_parameterVersionWithJsonFormatExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_parameter_manager_parameter" "parameter-basic" {
  provider = google-beta
  parameter_id = "parameter%{random_suffix}"
  format = "JSON"
}

resource "google_parameter_manager_parameter_version" "parameter-version-with-json-format" {
  provider = google-beta
  parameter = google_parameter_manager_parameter.parameter-basic.id
  parameter_version_id = "tf_test_parameter_version%{random_suffix}"
  parameter_data = jsonencode({
    "key1": "val1",
    "key2": "val2"
  })
}
`, context)
}

func TestAccParameterManagerParameterVersion_parameterVersionWithKmsKeyExample(t *testing.T) {
	t.Parallel()
	acctest.BootstrapIamMembers(t, []acctest.IamMember{
		{
			Member: "serviceAccount:service-{project_number}@gcp-sa-pm.iam.gserviceaccount.com",
			Role:   "roles/cloudkms.cryptoKeyEncrypterDecrypter",
		},
	})

	context := map[string]interface{}{
		"kms_key":       acctest.BootstrapKMSKey(t).CryptoKey.Name,
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckParameterManagerParameterVersionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccParameterManagerParameterVersion_parameterVersionWithKmsKeyExample(context),
			},
			{
				ResourceName:            "google_parameter_manager_parameter_version.parameter-version-with-kms-key",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parameter", "parameter_version_id"},
			},
		},
	})
}

func testAccParameterManagerParameterVersion_parameterVersionWithKmsKeyExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_project" "project" {
  provider = google-beta
}

resource "google_parameter_manager_parameter" "parameter-basic" {
  provider  = google-beta
  parameter_id = "parameter%{random_suffix}"

  kms_key = "%{kms_key}"
}

resource "google_parameter_manager_parameter_version" "parameter-version-with-kms-key" {
  provider = google-beta
  parameter = google_parameter_manager_parameter.parameter-basic.id
  parameter_version_id = "tf_test_parameter_version%{random_suffix}"
  parameter_data = "app-parameter-version-data"
}
`, context)
}

func TestAccParameterManagerParameterVersion_parameterVersionWithYamlFormatExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckParameterManagerParameterVersionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccParameterManagerParameterVersion_parameterVersionWithYamlFormatExample(context),
			},
			{
				ResourceName:            "google_parameter_manager_parameter_version.parameter-version-with-yaml-format",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parameter", "parameter_version_id"},
			},
		},
	})
}

func testAccParameterManagerParameterVersion_parameterVersionWithYamlFormatExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_parameter_manager_parameter" "parameter-basic" {
  provider = google-beta
  parameter_id = "parameter%{random_suffix}"
  format = "YAML"
}

resource "google_parameter_manager_parameter_version" "parameter-version-with-yaml-format" {
  provider = google-beta
  parameter = google_parameter_manager_parameter.parameter-basic.id
  parameter_version_id = "tf_test_parameter_version%{random_suffix}"
  parameter_data = yamlencode({
    "key1": "val1",
    "key2": "val2"
  })
}
`, context)
}

func testAccCheckParameterManagerParameterVersionDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_parameter_manager_parameter_version" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ParameterManagerBasePath}}{{parameter}}/versions/{{parameter_version_id}}")
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
				return fmt.Errorf("ParameterManagerParameterVersion still exists at %s", url)
			}
		}

		return nil
	}
}
