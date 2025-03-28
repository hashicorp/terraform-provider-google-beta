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

package dataprocmetastore_test

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

func TestAccDataprocMetastoreFederation_dataprocMetastoreFederationBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataprocMetastoreFederationDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocMetastoreFederation_dataprocMetastoreFederationBasicExample(context),
			},
			{
				ResourceName:            "google_dataproc_metastore_federation.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"deletion_protection", "federation_id", "labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccDataprocMetastoreFederation_dataprocMetastoreFederationBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dataproc_metastore_federation" "default" {
  location      = "us-central1"
  federation_id = "tf-test-metastore-fed%{random_suffix}"
  version       = "3.1.2"

  backend_metastores {
    rank           = "1"
    name           = google_dataproc_metastore_service.default.id
    metastore_type = "DATAPROC_METASTORE" 
  }
}

resource "google_dataproc_metastore_service" "default" {
  service_id = "tf-test-metastore-service%{random_suffix}"
  location   = "us-central1"
  tier       = "DEVELOPER"


  hive_metastore_config {
    version           = "3.1.2"
    endpoint_protocol = "GRPC"
  }
  deletion_protection = false
}
`, context)
}

func TestAccDataprocMetastoreFederation_dataprocMetastoreFederationBigqueryExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataprocMetastoreFederationDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocMetastoreFederation_dataprocMetastoreFederationBigqueryExample(context),
			},
			{
				ResourceName:            "google_dataproc_metastore_federation.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"federation_id", "labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccDataprocMetastoreFederation_dataprocMetastoreFederationBigqueryExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dataproc_metastore_federation" "default" {
  location      = "us-central1"
  federation_id = "tf-test-metastore-fed%{random_suffix}"
  version       = "3.1.2"

  backend_metastores {
    rank           = "2"
    name           = data.google_project.project.id
    metastore_type = "BIGQUERY" 
  }

  backend_metastores {
    rank           = "1"
    name           = google_dataproc_metastore_service.default.id
    metastore_type = "DATAPROC_METASTORE" 
  }
}

resource "google_dataproc_metastore_service" "default" {
  service_id = "tf-test-metastore-service%{random_suffix}"
  location   = "us-central1"
  tier       = "DEVELOPER"


  hive_metastore_config {
    version           = "3.1.2"
    endpoint_protocol = "GRPC"
  }
}

data "google_project" "project" {}
`, context)
}

func testAccCheckDataprocMetastoreFederationDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_dataproc_metastore_federation" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{DataprocMetastoreBasePath}}projects/{{project}}/locations/{{location}}/federations/{{federation_id}}")
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
				return fmt.Errorf("DataprocMetastoreFederation still exists at %s", url)
			}
		}

		return nil
	}
}
