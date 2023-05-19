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

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccComputeBackendServiceIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/compute.admin",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendServiceIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_compute_backend_service_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/global/backendServices/%s roles/compute.admin", acctest.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-backend-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccComputeBackendServiceIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_compute_backend_service_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/global/backendServices/%s roles/compute.admin", acctest.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-backend-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeBackendServiceIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/compute.admin",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccComputeBackendServiceIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_compute_backend_service_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/global/backendServices/%s roles/compute.admin user:admin@hashicorptest.com", acctest.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-backend-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeBackendServiceIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/compute.admin",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendServiceIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_compute_backend_service_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_compute_backend_service_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/global/backendServices/%s", acctest.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-backend-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeBackendServiceIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_compute_backend_service_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/global/backendServices/%s", acctest.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-backend-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeBackendServiceIamBindingGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/compute.admin",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendServiceIamBinding_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_compute_backend_service_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/global/backendServices/%s roles/compute.admin %s", acctest.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-backend-service%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeBackendServiceIamBindingGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/compute.admin",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendServiceIamBinding_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_compute_backend_service_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/global/backendServices/%s roles/compute.admin", acctest.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-backend-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_compute_backend_service_iam_binding.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/global/backendServices/%s roles/compute.admin %s", acctest.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-backend-service%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_compute_backend_service_iam_binding.foo3",
				ImportStateId:     fmt.Sprintf("projects/%s/global/backendServices/%s roles/compute.admin %s", acctest.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-backend-service%s", context["random_suffix"]), context["condition_title_no_desc"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeBackendServiceIamMemberGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/compute.admin",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendServiceIamMember_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_compute_backend_service_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/global/backendServices/%s roles/compute.admin user:admin@hashicorptest.com %s", acctest.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-backend-service%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeBackendServiceIamMemberGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/compute.admin",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendServiceIamMember_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_compute_backend_service_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/global/backendServices/%s roles/compute.admin user:admin@hashicorptest.com", acctest.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-backend-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_compute_backend_service_iam_member.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/global/backendServices/%s roles/compute.admin user:admin@hashicorptest.com %s", acctest.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-backend-service%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_compute_backend_service_iam_member.foo3",
				ImportStateId:     fmt.Sprintf("projects/%s/global/backendServices/%s roles/compute.admin user:admin@hashicorptest.com %s", acctest.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-backend-service%s", context["random_suffix"]), context["condition_title_no_desc"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeBackendServiceIamPolicyGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/compute.admin",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	// Test should have 2 bindings: one with a description and one without. Any < chars are converted to a unicode character by the API.
	expectedPolicyData := Nprintf(`{"bindings":[{"condition":{"description":"%{condition_desc}","expression":"%{condition_expr}","title":"%{condition_title}"},"members":["user:admin@hashicorptest.com"],"role":"%{role}"},{"condition":{"expression":"%{condition_expr}","title":"%{condition_title}-no-description"},"members":["user:admin@hashicorptest.com"],"role":"%{role}"}]}`, context)
	expectedPolicyData = strings.Replace(expectedPolicyData, "<", "\\u003c", -1)

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendServiceIamPolicy_withConditionGenerated(context),
				Check: resource.ComposeAggregateTestCheckFunc(
					// TODO(SarahFrench) - uncomment once https://github.com/GoogleCloudPlatform/magic-modules/pull/6466 merged
					// resource.TestCheckResourceAttr("data.google_iam_policy.foo", "policy_data", expectedPolicyData),
					resource.TestCheckResourceAttr("google_compute_backend_service_iam_policy.foo", "policy_data", expectedPolicyData),
					resource.TestCheckResourceAttrWith("data.google_iam_policy.foo", "policy_data", checkGoogleIamPolicy),
				),
			},
			{
				ResourceName:      "google_compute_backend_service_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/global/backendServices/%s", acctest.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-backend-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeBackendServiceIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
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

resource "google_compute_backend_service_iam_member" "foo" {
  project = google_compute_backend_service.default.project
  name = google_compute_backend_service.default.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccComputeBackendServiceIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
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

resource "google_compute_backend_service_iam_policy" "foo" {
  project = google_compute_backend_service.default.project
  name = google_compute_backend_service.default.name
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_compute_backend_service_iam_policy" "foo" {
  project = google_compute_backend_service.default.project
  name = google_compute_backend_service.default.name
  depends_on = [
    google_compute_backend_service_iam_policy.foo
  ]
}
`, context)
}

func testAccComputeBackendServiceIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
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

resource "google_compute_backend_service_iam_policy" "foo" {
  project = google_compute_backend_service.default.project
  name = google_compute_backend_service.default.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccComputeBackendServiceIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
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

resource "google_compute_backend_service_iam_binding" "foo" {
  project = google_compute_backend_service.default.project
  name = google_compute_backend_service.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccComputeBackendServiceIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
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

resource "google_compute_backend_service_iam_binding" "foo" {
  project = google_compute_backend_service.default.project
  name = google_compute_backend_service.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}

func testAccComputeBackendServiceIamBinding_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
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

resource "google_compute_backend_service_iam_binding" "foo" {
  project = google_compute_backend_service.default.project
  name = google_compute_backend_service.default.name
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

func testAccComputeBackendServiceIamBinding_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
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

resource "google_compute_backend_service_iam_binding" "foo" {
  project = google_compute_backend_service.default.project
  name = google_compute_backend_service.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}

resource "google_compute_backend_service_iam_binding" "foo2" {
  project = google_compute_backend_service.default.project
  name = google_compute_backend_service.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}

resource "google_compute_backend_service_iam_binding" "foo3" {
  project = google_compute_backend_service.default.project
  name = google_compute_backend_service.default.name
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

func testAccComputeBackendServiceIamMember_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
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

resource "google_compute_backend_service_iam_member" "foo" {
  project = google_compute_backend_service.default.project
  name = google_compute_backend_service.default.name
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

func testAccComputeBackendServiceIamMember_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
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

resource "google_compute_backend_service_iam_member" "foo" {
  project = google_compute_backend_service.default.project
  name = google_compute_backend_service.default.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}

resource "google_compute_backend_service_iam_member" "foo2" {
  project = google_compute_backend_service.default.project
  name = google_compute_backend_service.default.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}

resource "google_compute_backend_service_iam_member" "foo3" {
  project = google_compute_backend_service.default.project
  name = google_compute_backend_service.default.name
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

func testAccComputeBackendServiceIamPolicy_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
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

resource "google_compute_backend_service_iam_policy" "foo" {
  project = google_compute_backend_service.default.project
  name = google_compute_backend_service.default.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}
