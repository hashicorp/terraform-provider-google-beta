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
//	Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/services/iambeta/resource_iam_workload_identity_pool_namespace_test.go.tmpl
//
//	DO NOT EDIT this file directly. Any changes made to this file will be
//	overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
package iambeta_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccIAMBetaWorkloadIdentityPoolNamespace_minimal(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckIAMBetaWorkloadIdentityPoolNamespaceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIAMBetaWorkloadIdentityPoolNamespace_minimal(context),
			},
			{
				ResourceName:            "google_iam_workload_identity_pool_namespace.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"workload_identity_pool_id", "workload_identity_pool_namespace_id"},
			},
			{
				Config: testAccIAMBetaWorkloadIdentityPoolNamespace_updated(context),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("google_iam_workload_identity_pool_namespace.example", plancheck.ResourceActionUpdate),
					},
				},
			},
			{
				ResourceName:            "google_iam_workload_identity_pool_namespace.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"workload_identity_pool_id", "workload_identity_pool_namespace_id"},
			},
		},
	})
}

func TestAccIAMBetaWorkloadIdentityPoolNamespace_full(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckIAMBetaWorkloadIdentityPoolNamespaceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIAMBetaWorkloadIdentityPoolNamespace_full(context),
			},
			{
				ResourceName:            "google_iam_workload_identity_pool_namespace.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"workload_identity_pool_id", "workload_identity_pool_namespace_id"},
			},
			{
				Config: testAccIAMBetaWorkloadIdentityPoolNamespace_updated(context),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("google_iam_workload_identity_pool_namespace.example", plancheck.ResourceActionUpdate),
					},
				},
			},
			{
				ResourceName:            "google_iam_workload_identity_pool_namespace.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"workload_identity_pool_id", "workload_identity_pool_namespace_id"},
			},
		},
	})
}

func testAccIAMBetaWorkloadIdentityPoolNamespace_minimal(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_iam_workload_identity_pool" "pool" {
  provider = google-beta

  workload_identity_pool_id = "tf-test-example-pool%{random_suffix}"
  mode                      = "TRUST_DOMAIN"
}

resource "google_iam_workload_identity_pool_namespace" "example" {
  provider = google-beta

  workload_identity_pool_id           = google_iam_workload_identity_pool.pool.workload_identity_pool_id
  workload_identity_pool_namespace_id = "tf-test-example-namespace%{random_suffix}"
}
`, context)
}

func testAccIAMBetaWorkloadIdentityPoolNamespace_full(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_iam_workload_identity_pool" "pool" {
  provider = google-beta

  workload_identity_pool_id = "tf-test-example-pool%{random_suffix}"
  mode                      = "TRUST_DOMAIN"
}

resource "google_iam_workload_identity_pool_namespace" "example" {
  provider = google-beta

  workload_identity_pool_id           = google_iam_workload_identity_pool.pool.workload_identity_pool_id
  workload_identity_pool_namespace_id = "tf-test-example-namespace%{random_suffix}"
  description                         = "Example Namespace in a Workload Identity Pool"
  disabled                            = true
}
`, context)
}

func testAccIAMBetaWorkloadIdentityPoolNamespace_updated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_iam_workload_identity_pool" "pool" {
  provider = google-beta

  workload_identity_pool_id = "tf-test-example-pool%{random_suffix}"
  mode                      = "TRUST_DOMAIN"
}

resource "google_iam_workload_identity_pool_namespace" "example" {
  provider = google-beta

  workload_identity_pool_id           = google_iam_workload_identity_pool.pool.workload_identity_pool_id
  workload_identity_pool_namespace_id = "tf-test-example-namespace%{random_suffix}"
  description                         = "Updated Namespace in a Workload Identity Pool"
  disabled                            = false
}
`, context)
}
