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
//	Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/services/managedkafka/resource_managed_kafka_acl_test.go
//
//	DO NOT EDIT this file directly. Any changes made to this file will be
//	overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
package managedkafka_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccManagedKafkaAcl_update(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckManagedKafkaAclDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccManagedKafkaAcl_full(context),
			},
			{
				ResourceName:            "google_managed_kafka_acl.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"cluster", "location", "acl_id"},
			},
			{
				Config: testAccManagedKafkaAcl_update(context),
			},
			{
				ResourceName:            "google_managed_kafka_acl.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"cluster", "location", "acl_id"},
			},
		},
	})
}

func testAccManagedKafkaAcl_full(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_managed_kafka_cluster" "example" {
  cluster_id = "tf-test-my-cluster%{random_suffix}"
  location = "us-central1"
  capacity_config {
    vcpu_count = 3
    memory_bytes = 3221225472
  }
  gcp_config {
    access_config {
      network_configs {
        subnet = "projects/${data.google_project.project.number}/regions/us-central1/subnetworks/default"
      }
    }
  }
}

resource "google_managed_kafka_acl" "example" {
  cluster = google_managed_kafka_cluster.example.cluster_id
  acl_id = "topic/tf-test-my-acl%{random_suffix}"
  location = "us-central1"
   acl_entries {
      principal = "User:admin@my-project.iam.gserviceaccount.com"
      permission_type = "ALLOW"
      operation = "ALL"
      host = "*"
    }
    acl_entries {
      principal = "User:producer-client@my-project.iam.gserviceaccount.com"
      permission_type = "ALLOW"
      operation = "WRITE"
      host = "*"
    }
}

data "google_project" "project" {
}
`, context)
}

func testAccManagedKafkaAcl_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_managed_kafka_cluster" "example" {
  cluster_id = "tf-test-my-cluster%{random_suffix}"
  location = "us-central1"
  capacity_config {
    vcpu_count = 3
    memory_bytes = 3221225472
  }
  gcp_config {
    access_config {
      network_configs {
        subnet = "projects/${data.google_project.project.number}/regions/us-central1/subnetworks/default"
      }
    }
  }
}

resource "google_managed_kafka_acl" "example" {
  cluster = google_managed_kafka_cluster.example.cluster_id
  acl_id = "topic/tf-test-my-acl%{random_suffix}"
  location = "us-central1"
  acl_entries {
      principal = "User:admin@project.iam.gserviceaccount.com"
      permission_type = "ALLOW"
      operation = "ALL"
      host = "*"
    }
    acl_entries {
      principal = "User:producer-client@my-project.iam.gserviceaccount.com"
      permission_type = "ALLOW"
      operation = "WRITE"
      host = "*"
    }
    acl_entries {
      principal = "User:producer-client@my-project.iam.gserviceaccount.com"
      permission_type = "ALLOW"
      operation = "CREATE"
      host = "*"
    }
}

data "google_project" "project" {
}
`, context)
}
