// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/iap/WebBackendService.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/examples/base_configs/iam_test_file.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package iap_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
)

func TestAccIapWebBackendServiceIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/iap.httpsResourceAccessor",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebBackendServiceIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_backend_service_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute/services/%s roles/iap.httpsResourceAccessor", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-backend-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccIapWebBackendServiceIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_backend_service_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute/services/%s roles/iap.httpsResourceAccessor", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-backend-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebBackendServiceIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/iap.httpsResourceAccessor",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccIapWebBackendServiceIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_backend_service_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute/services/%s roles/iap.httpsResourceAccessor user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-backend-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebBackendServiceIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/iap.httpsResourceAccessor",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebBackendServiceIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_iap_web_backend_service_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_iap_web_backend_service_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute/services/%s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-backend-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccIapWebBackendServiceIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_iap_web_backend_service_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute/services/%s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-backend-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebBackendServiceIamBindingGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/iap.httpsResourceAccessor",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebBackendServiceIamBinding_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_backend_service_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute/services/%s roles/iap.httpsResourceAccessor %s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-backend-service%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebBackendServiceIamBindingGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/iap.httpsResourceAccessor",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebBackendServiceIamBinding_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_backend_service_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute/services/%s roles/iap.httpsResourceAccessor", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-backend-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_iap_web_backend_service_iam_binding.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute/services/%s roles/iap.httpsResourceAccessor %s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-backend-service%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_iap_web_backend_service_iam_binding.foo3",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute/services/%s roles/iap.httpsResourceAccessor %s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-backend-service%s", context["random_suffix"]), context["condition_title_no_desc"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebBackendServiceIamMemberGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/iap.httpsResourceAccessor",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebBackendServiceIamMember_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_backend_service_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute/services/%s roles/iap.httpsResourceAccessor user:admin@hashicorptest.com %s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-backend-service%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebBackendServiceIamMemberGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/iap.httpsResourceAccessor",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebBackendServiceIamMember_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_backend_service_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute/services/%s roles/iap.httpsResourceAccessor user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-backend-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_iap_web_backend_service_iam_member.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute/services/%s roles/iap.httpsResourceAccessor user:admin@hashicorptest.com %s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-backend-service%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_iap_web_backend_service_iam_member.foo3",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute/services/%s roles/iap.httpsResourceAccessor user:admin@hashicorptest.com %s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-backend-service%s", context["random_suffix"]), context["condition_title_no_desc"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebBackendServiceIamPolicyGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/iap.httpsResourceAccessor",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	// Test should have 2 bindings: one with a description and one without. Any < chars are converted to a unicode character by the API.
	expectedPolicyData := acctest.Nprintf(`{"bindings":[{"condition":{"description":"%{condition_desc}","expression":"%{condition_expr}","title":"%{condition_title}"},"members":["user:admin@hashicorptest.com"],"role":"%{role}"},{"condition":{"expression":"%{condition_expr}","title":"%{condition_title}-no-description"},"members":["user:admin@hashicorptest.com"],"role":"%{role}"}]}`, context)
	expectedPolicyData = strings.Replace(expectedPolicyData, "<", "\\u003c", -1)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebBackendServiceIamPolicy_withConditionGenerated(context),
				Check: resource.ComposeAggregateTestCheckFunc(
					// TODO - uncomment once https://github.com/GoogleCloudPlatform/magic-modules/pull/6466 merged
					// resource.TestCheckResourceAttr("data.google_iam_policy.foo", "policy_data", expectedPolicyData),
					resource.TestCheckResourceAttr("google_iap_web_backend_service_iam_policy.foo", "policy_data", expectedPolicyData),
					resource.TestCheckResourceAttrWith("data.google_iam_policy.foo", "policy_data", tpgresource.CheckGoogleIamPolicy),
				),
			},
			{
				ResourceName:      "google_iap_web_backend_service_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute/services/%s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-backend-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccIapWebBackendServiceIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_backend_service" "default" {
  name          = "tf-test-backend-service%{random_suffix}"
  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  name               = "tf-test-health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_iap_web_backend_service_iam_member" "foo" {
  project = google_compute_backend_service.default.project
  web_backend_service = google_compute_backend_service.default.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccIapWebBackendServiceIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_backend_service" "default" {
  name          = "tf-test-backend-service%{random_suffix}"
  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  name               = "tf-test-health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_iap_web_backend_service_iam_policy" "foo" {
  project = google_compute_backend_service.default.project
  web_backend_service = google_compute_backend_service.default.name
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_iap_web_backend_service_iam_policy" "foo" {
  project = google_compute_backend_service.default.project
  web_backend_service = google_compute_backend_service.default.name
  depends_on = [
    google_iap_web_backend_service_iam_policy.foo
  ]
}
`, context)
}

func testAccIapWebBackendServiceIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_backend_service" "default" {
  name          = "tf-test-backend-service%{random_suffix}"
  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  name               = "tf-test-health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

data "google_iam_policy" "foo" {
}

resource "google_iap_web_backend_service_iam_policy" "foo" {
  project = google_compute_backend_service.default.project
  web_backend_service = google_compute_backend_service.default.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccIapWebBackendServiceIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_backend_service" "default" {
  name          = "tf-test-backend-service%{random_suffix}"
  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  name               = "tf-test-health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_iap_web_backend_service_iam_binding" "foo" {
  project = google_compute_backend_service.default.project
  web_backend_service = google_compute_backend_service.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccIapWebBackendServiceIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_backend_service" "default" {
  name          = "tf-test-backend-service%{random_suffix}"
  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  name               = "tf-test-health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_iap_web_backend_service_iam_binding" "foo" {
  project = google_compute_backend_service.default.project
  web_backend_service = google_compute_backend_service.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}

func testAccIapWebBackendServiceIamBinding_withConditionGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_backend_service" "default" {
  name          = "tf-test-backend-service%{random_suffix}"
  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  name               = "tf-test-health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_iap_web_backend_service_iam_binding" "foo" {
  project = google_compute_backend_service.default.project
  web_backend_service = google_compute_backend_service.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccIapWebBackendServiceIamBinding_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_backend_service" "default" {
  name          = "tf-test-backend-service%{random_suffix}"
  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  name               = "tf-test-health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_iap_web_backend_service_iam_binding" "foo" {
  project = google_compute_backend_service.default.project
  web_backend_service = google_compute_backend_service.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}

resource "google_iap_web_backend_service_iam_binding" "foo2" {
  project = google_compute_backend_service.default.project
  web_backend_service = google_compute_backend_service.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}

resource "google_iap_web_backend_service_iam_binding" "foo3" {
  project = google_compute_backend_service.default.project
  web_backend_service = google_compute_backend_service.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    # Check that lack of description doesn't cause any issues
    # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
    title       = "%{condition_title_no_desc}"
    expression  = "%{condition_expr_no_desc}"
  }
}
`, context)
}

func testAccIapWebBackendServiceIamMember_withConditionGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_backend_service" "default" {
  name          = "tf-test-backend-service%{random_suffix}"
  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  name               = "tf-test-health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_iap_web_backend_service_iam_member" "foo" {
  project = google_compute_backend_service.default.project
  web_backend_service = google_compute_backend_service.default.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccIapWebBackendServiceIamMember_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_backend_service" "default" {
  name          = "tf-test-backend-service%{random_suffix}"
  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  name               = "tf-test-health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_iap_web_backend_service_iam_member" "foo" {
  project = google_compute_backend_service.default.project
  web_backend_service = google_compute_backend_service.default.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}

resource "google_iap_web_backend_service_iam_member" "foo2" {
  project = google_compute_backend_service.default.project
  web_backend_service = google_compute_backend_service.default.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}

resource "google_iap_web_backend_service_iam_member" "foo3" {
  project = google_compute_backend_service.default.project
  web_backend_service = google_compute_backend_service.default.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    # Check that lack of description doesn't cause any issues
    # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
    title       = "%{condition_title_no_desc}"
    expression  = "%{condition_expr_no_desc}"
  }
}
`, context)
}

func testAccIapWebBackendServiceIamPolicy_withConditionGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_backend_service" "default" {
  name          = "tf-test-backend-service%{random_suffix}"
  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  name               = "tf-test-health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
    condition {
      # Check that lack of description doesn't cause any issues
      # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
      title       = "%{condition_title_no_desc}"
      expression  = "%{condition_expr_no_desc}"
    }
  }
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
    condition {
      title       = "%{condition_title}"
      description = "%{condition_desc}"
      expression  = "%{condition_expr}"
    }
  }
}

resource "google_iap_web_backend_service_iam_policy" "foo" {
  project = google_compute_backend_service.default.project
  web_backend_service = google_compute_backend_service.default.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}
