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
//	Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/services/firebase/resource_firebase_project_test.go.tmpl
//
//	DO NOT EDIT this file directly. Any changes made to this file will be
//	overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
package firebase_test

import (
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccFirebaseProject_destroyAndReapply(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        envvar.GetTestOrgFromEnv(t),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseProject_firebaseProjectBasicExample(context),
			},
			{
				Config: testAccFirebaseProject_firebaseProjectBasicExampleDestroyed(context),
			},
			{
				Config: testAccFirebaseProject_firebaseProjectBasicExample(context),
			},
			{
				ResourceName:      "google_firebase_project.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccFirebaseProject_firebaseProjectBasicExampleDestroyed(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_project" "default" {
  provider = google-beta

  project_id = "tf-test-my-project%{random_suffix}"
  name       = "tf-test-my-project%{random_suffix}"
  org_id     = "%{org_id}"
  deletion_policy = "DELETE"

  labels = {
    "firebase" = "enabled"
  }
}
`, context)
}
