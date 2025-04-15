// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
)

// Merged IamBinding, IamMember, and IamPolicy tests into a single test
// to avoid storage pool's limit: create a maximum of 5 storage pools per hour
// https://cloud.google.com/compute/docs/disks/storage-pools#sp_limitations

func TestAccComputeStoragePoolIam(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/compute.viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// create storage pool for testing
				Config: testAccComputeStoragePoolWithoutIamBalanced(context),
			},
			{
				Config: testAccComputeStoragePoolIamBinding_basic(context),
			},
			{
				ResourceName:      "google_compute_storage_pool_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/zones/%s/storagePools/tf-test-storage-pool-balanced-%s roles/compute.viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestZoneFromEnv(), context["random_suffix"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccComputeStoragePoolIamBinding_update(context),
			},
			{
				ResourceName:      "google_compute_storage_pool_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/zones/%s/storagePools/tf-test-storage-pool-balanced-%s roles/compute.viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestZoneFromEnv(), context["random_suffix"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// remove Iam Binding
				Config: testAccComputeStoragePoolWithoutIamBalanced(context),
			},
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccComputeStoragePoolIamMember_basic(context),
			},
			{
				ResourceName:      "google_compute_storage_pool_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/zones/%s/storagePools/tf-test-storage-pool-balanced-%s roles/compute.viewer user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), envvar.GetTestZoneFromEnv(), context["random_suffix"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// remove Iam Member
				Config: testAccComputeStoragePoolWithoutIamBalanced(context),
			},
			{
				Config: testAccComputeStoragePoolIamPolicy_basic(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_compute_storage_pool_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_compute_storage_pool_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/zones/%s/storagePools/tf-test-storage-pool-balanced-%s", envvar.GetTestProjectFromEnv(), envvar.GetTestZoneFromEnv(), context["random_suffix"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeStoragePoolIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_compute_storage_pool_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/zones/%s/storagePools/tf-test-storage-pool-balanced-%s", envvar.GetTestProjectFromEnv(), envvar.GetTestZoneFromEnv(), context["random_suffix"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeStoragePoolIam_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/compute.viewer",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2020_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2021-01-01T00:00:00Z\")`,
	}

	// Test should have 2 bindings: one with a description and one without. Any < chars are converted to a unicode character by the API.
	expectedPolicyData := acctest.Nprintf(`{"bindings":[{"condition":{"description":"%{condition_desc}","expression":"%{condition_expr}","title":"%{condition_title}"},"members":["user:admin@hashicorptest.com"],"role":"%{role}"},{"condition":{"expression":"%{condition_expr_no_desc}","title":"%{condition_title_no_desc}"},"members":["user:admin@hashicorptest.com"],"role":"%{role}"}]}`, context)
	expectedPolicyData = strings.Replace(expectedPolicyData, "<", "\\u003c", -1)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// create storage pool for testing
				Config: testAccComputeStoragePoolWithoutIamThroughput(context),
			},
			{
				Config: testAccComputeStoragePoolIamBinding_withCondition(context),
			},
			{
				ResourceName:      "google_compute_storage_pool_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/zones/%s/storagePools/tf-test-storage-pool-throughput-%s roles/compute.viewer %s", envvar.GetTestProjectFromEnv(), envvar.GetTestZoneFromEnv(), context["random_suffix"], context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// remove Iam Binding
				Config: testAccComputeStoragePoolWithoutIamThroughput(context),
			},
			{
				Config: testAccComputeStoragePoolIamBinding_withAndWithoutCondition(context),
			},
			{
				ResourceName:      "google_compute_storage_pool_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/zones/%s/storagePools/tf-test-storage-pool-throughput-%s roles/compute.viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestZoneFromEnv(), context["random_suffix"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_compute_storage_pool_iam_binding.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/zones/%s/storagePools/tf-test-storage-pool-throughput-%s roles/compute.viewer %s", envvar.GetTestProjectFromEnv(), envvar.GetTestZoneFromEnv(), context["random_suffix"], context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_compute_storage_pool_iam_binding.foo3",
				ImportStateId:     fmt.Sprintf("projects/%s/zones/%s/storagePools/tf-test-storage-pool-throughput-%s roles/compute.viewer %s", envvar.GetTestProjectFromEnv(), envvar.GetTestZoneFromEnv(), context["random_suffix"], context["condition_title_no_desc"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// remove Iam Binding
				Config: testAccComputeStoragePoolWithoutIamThroughput(context),
			},
			{
				Config: testAccComputeStoragePoolIamMember_withCondition(context),
			},
			{
				ResourceName:      "google_compute_storage_pool_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/zones/%s/storagePools/tf-test-storage-pool-throughput-%s roles/compute.viewer user:admin@hashicorptest.com %s", envvar.GetTestProjectFromEnv(), envvar.GetTestZoneFromEnv(), context["random_suffix"], context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// remove Iam Binding
				Config: testAccComputeStoragePoolWithoutIamThroughput(context),
			},
			{
				Config: testAccComputeStoragePoolIamMember_withAndWithoutCondition(context),
			},
			{
				ResourceName:      "google_compute_storage_pool_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/zones/%s/storagePools/tf-test-storage-pool-throughput-%s roles/compute.viewer user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), envvar.GetTestZoneFromEnv(), context["random_suffix"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_compute_storage_pool_iam_member.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/zones/%s/storagePools/tf-test-storage-pool-throughput-%s roles/compute.viewer user:admin@hashicorptest.com %s", envvar.GetTestProjectFromEnv(), envvar.GetTestZoneFromEnv(), context["random_suffix"], context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_compute_storage_pool_iam_member.foo3",
				ImportStateId:     fmt.Sprintf("projects/%s/zones/%s/storagePools/tf-test-storage-pool-throughput-%s roles/compute.viewer user:admin@hashicorptest.com %s", envvar.GetTestProjectFromEnv(), envvar.GetTestZoneFromEnv(), context["random_suffix"], context["condition_title_no_desc"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// remove Iam Member
				Config: testAccComputeStoragePoolWithoutIamThroughput(context),
			},
			{
				Config: testAccComputeStoragePoolIamPolicy_withCondition(context),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("google_compute_storage_pool_iam_policy.foo", "policy_data", expectedPolicyData),
					resource.TestCheckResourceAttrWith("data.google_iam_policy.foo", "policy_data", tpgresource.CheckGoogleIamPolicy),
				),
			},
			{
				ResourceName:      "google_compute_storage_pool_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/zones/%s/storagePools/tf-test-storage-pool-throughput-%s", envvar.GetTestProjectFromEnv(), envvar.GetTestZoneFromEnv(), context["random_suffix"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeStoragePoolWithoutIamBalanced(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_storage_pool" "test-storage-pool-balanced" {
  name = "tf-test-storage-pool-balanced-%{random_suffix}"

  description = "Hyperdisk Balanced storage pool"

  capacity_provisioning_type   = "STANDARD"
  pool_provisioned_capacity_gb = "10240"

  performance_provisioning_type = "STANDARD"
  pool_provisioned_iops         = "10000"
  pool_provisioned_throughput   = "1024"

  storage_pool_type = data.google_compute_storage_pool_types.balanced.self_link

  deletion_protection = false

  zone = "us-central1-a"
}

data "google_project" "project" {}

data "google_compute_storage_pool_types" "balanced" {
  zone = "us-central1-a"
	storage_pool_type = "hyperdisk-balanced"
}
`, context)
}

func testAccComputeStoragePoolWithoutIamThroughput(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_storage_pool" "test-storage-pool-throughput" {
  name = "tf-test-storage-pool-throughput-%{random_suffix}"

  description = "Hyperdisk Throughput storage pool"

  capacity_provisioning_type   = "STANDARD"
  pool_provisioned_capacity_gb = "10240"

  performance_provisioning_type = "STANDARD"
  pool_provisioned_throughput   = "100"

  storage_pool_type = data.google_compute_storage_pool_types.throughput.self_link

  deletion_protection = false

  zone = "us-central1-a"
}

data "google_project" "project" {}

data "google_compute_storage_pool_types" "throughput" {
  zone = "us-central1-a"
	storage_pool_type = "hyperdisk-throughput"
}
`, context)
}

func testAccComputeStoragePoolIamMember_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_storage_pool" "test-storage-pool-balanced" {
  name = "tf-test-storage-pool-balanced-%{random_suffix}"

  description = "Hyperdisk Balanced storage pool"

  capacity_provisioning_type   = "STANDARD"
  pool_provisioned_capacity_gb = "10240"

  performance_provisioning_type = "STANDARD"
  pool_provisioned_iops         = "10000"
  pool_provisioned_throughput   = "1024"

  storage_pool_type = data.google_compute_storage_pool_types.balanced.self_link

  deletion_protection = false

  zone = "us-central1-a"
}

data "google_project" "project" {}

data "google_compute_storage_pool_types" "balanced" {
  zone = "us-central1-a"
	storage_pool_type = "hyperdisk-balanced"
}

resource "google_compute_storage_pool_iam_member" "foo" {
  project = google_compute_storage_pool.test-storage-pool-balanced.project
  zone = "us-central1-a"
  name = google_compute_storage_pool.test-storage-pool-balanced.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccComputeStoragePoolIamPolicy_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_storage_pool" "test-storage-pool-balanced" {
  name = "tf-test-storage-pool-balanced-%{random_suffix}"

  description = "Hyperdisk Balanced storage pool"

  capacity_provisioning_type   = "STANDARD"
  pool_provisioned_capacity_gb = "10240"

  performance_provisioning_type = "STANDARD"
  pool_provisioned_iops         = "10000"
  pool_provisioned_throughput   = "1024"

  storage_pool_type = data.google_compute_storage_pool_types.balanced.self_link

  deletion_protection = false

  zone = "us-central1-a"
}

data "google_project" "project" {}

data "google_compute_storage_pool_types" "balanced" {
  zone = "us-central1-a"
	storage_pool_type = "hyperdisk-balanced"
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_compute_storage_pool_iam_policy" "foo" {
  project = google_compute_storage_pool.test-storage-pool-balanced.project
  zone = "us-central1-a"
  name = google_compute_storage_pool.test-storage-pool-balanced.name
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_compute_storage_pool_iam_policy" "foo" {
  project = google_compute_storage_pool.test-storage-pool-balanced.project
  zone = "us-central1-a"
  name = google_compute_storage_pool.test-storage-pool-balanced.name
  depends_on = [
    google_compute_storage_pool_iam_policy.foo
  ]
}
`, context)
}

func testAccComputeStoragePoolIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_storage_pool" "test-storage-pool-balanced" {
  name = "tf-test-storage-pool-balanced-%{random_suffix}"

  description = "Hyperdisk Balanced storage pool"

  capacity_provisioning_type   = "STANDARD"
  pool_provisioned_capacity_gb = "10240"

  performance_provisioning_type = "STANDARD"
  pool_provisioned_iops         = "10000"
  pool_provisioned_throughput   = "1024"

  storage_pool_type = data.google_compute_storage_pool_types.balanced.self_link

  deletion_protection = false

  zone = "us-central1-a"
}

data "google_project" "project" {}

data "google_compute_storage_pool_types" "balanced" {
  zone = "us-central1-a"
	storage_pool_type = "hyperdisk-balanced"
}

data "google_iam_policy" "foo" {
}

resource "google_compute_storage_pool_iam_policy" "foo" {
  project = google_compute_storage_pool.test-storage-pool-balanced.project
  zone = "us-central1-a"
  name = google_compute_storage_pool.test-storage-pool-balanced.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccComputeStoragePoolIamBinding_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_storage_pool" "test-storage-pool-balanced" {
  name = "tf-test-storage-pool-balanced-%{random_suffix}"

  description = "Hyperdisk Balanced storage pool"

  capacity_provisioning_type   = "STANDARD"
  pool_provisioned_capacity_gb = "10240"

  performance_provisioning_type = "STANDARD"
  pool_provisioned_iops         = "10000"
  pool_provisioned_throughput   = "1024"

  storage_pool_type = data.google_compute_storage_pool_types.balanced.self_link

  deletion_protection = false

  zone = "us-central1-a"
}

data "google_project" "project" {}

data "google_compute_storage_pool_types" "balanced" {
  zone = "us-central1-a"
	storage_pool_type = "hyperdisk-balanced"
}

resource "google_compute_storage_pool_iam_binding" "foo" {
  project = google_compute_storage_pool.test-storage-pool-balanced.project
  zone = "us-central1-a"
  name = google_compute_storage_pool.test-storage-pool-balanced.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccComputeStoragePoolIamBinding_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_storage_pool" "test-storage-pool-balanced" {
  name = "tf-test-storage-pool-balanced-%{random_suffix}"

  description = "Hyperdisk Balanced storage pool"

  capacity_provisioning_type   = "STANDARD"
  pool_provisioned_capacity_gb = "10240"

  performance_provisioning_type = "STANDARD"
  pool_provisioned_iops         = "10000"
  pool_provisioned_throughput   = "1024"

  storage_pool_type = data.google_compute_storage_pool_types.balanced.self_link

  deletion_protection = false

  zone = "us-central1-a"
}

data "google_project" "project" {}

data "google_compute_storage_pool_types" "balanced" {
  zone = "us-central1-a"
	storage_pool_type = "hyperdisk-balanced"
}

resource "google_compute_storage_pool_iam_binding" "foo" {
  project = google_compute_storage_pool.test-storage-pool-balanced.project
  zone = "us-central1-a"
  name = google_compute_storage_pool.test-storage-pool-balanced.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}

func testAccComputeStoragePoolIamBinding_withCondition(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_storage_pool" "test-storage-pool-throughput" {
  name = "tf-test-storage-pool-throughput-%{random_suffix}"

  description = "Hyperdisk Throughput storage pool"

  capacity_provisioning_type   = "STANDARD"
  pool_provisioned_capacity_gb = "10240"

  performance_provisioning_type = "STANDARD"
  pool_provisioned_throughput   = "100"

  storage_pool_type = data.google_compute_storage_pool_types.throughput.self_link

  deletion_protection = false

  zone = "us-central1-a"
}

data "google_project" "project" {}

data "google_compute_storage_pool_types" "throughput" {
  zone = "us-central1-a"
	storage_pool_type = "hyperdisk-throughput"
}

resource "google_compute_storage_pool_iam_binding" "foo" {
  project = google_compute_storage_pool.test-storage-pool-throughput.project
  zone = "us-central1-a"
  name = google_compute_storage_pool.test-storage-pool-throughput.name
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

func testAccComputeStoragePoolIamBinding_withAndWithoutCondition(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_storage_pool" "test-storage-pool-throughput" {
  name = "tf-test-storage-pool-throughput-%{random_suffix}"

  description = "Hyperdisk Throughput storage pool"

  capacity_provisioning_type   = "STANDARD"
  pool_provisioned_capacity_gb = "10240"

  performance_provisioning_type = "STANDARD"
  pool_provisioned_throughput   = "100"

  storage_pool_type = data.google_compute_storage_pool_types.throughput.self_link

  deletion_protection = false

  zone = "us-central1-a"
}

data "google_project" "project" {}

data "google_compute_storage_pool_types" "throughput" {
  zone = "us-central1-a"
	storage_pool_type = "hyperdisk-throughput"
}

resource "google_compute_storage_pool_iam_binding" "foo" {
  project = google_compute_storage_pool.test-storage-pool-throughput.project
  zone = "us-central1-a"
  name = google_compute_storage_pool.test-storage-pool-throughput.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}

resource "google_compute_storage_pool_iam_binding" "foo2" {
  project = google_compute_storage_pool.test-storage-pool-throughput.project
  zone = "us-central1-a"
  name = google_compute_storage_pool.test-storage-pool-throughput.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }

  # Fixes flakiness in VCR replay test by ordering create/destroy.
  depends_on = [google_compute_storage_pool_iam_binding.foo]
}

resource "google_compute_storage_pool_iam_binding" "foo3" {
  project = google_compute_storage_pool.test-storage-pool-throughput.project
  zone = "us-central1-a"
  name = google_compute_storage_pool.test-storage-pool-throughput.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    # Check that lack of description doesn't cause any issues
    # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
    title       = "%{condition_title_no_desc}"
    expression  = "%{condition_expr_no_desc}"
  }

  depends_on = [google_compute_storage_pool_iam_binding.foo2]
}
`, context)
}

func testAccComputeStoragePoolIamMember_withCondition(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_storage_pool" "test-storage-pool-throughput" {
  name = "tf-test-storage-pool-throughput-%{random_suffix}"

  description = "Hyperdisk Throughput storage pool"

  capacity_provisioning_type   = "STANDARD"
  pool_provisioned_capacity_gb = "10240"

  performance_provisioning_type = "STANDARD"
  pool_provisioned_throughput   = "100"

  storage_pool_type = data.google_compute_storage_pool_types.throughput.self_link

  deletion_protection = false

  zone = "us-central1-a"
}

data "google_project" "project" {}

data "google_compute_storage_pool_types" "throughput" {
  zone = "us-central1-a"
	storage_pool_type = "hyperdisk-throughput"
}

resource "google_compute_storage_pool_iam_member" "foo" {
  project = google_compute_storage_pool.test-storage-pool-throughput.project
  zone = "us-central1-a"
  name = google_compute_storage_pool.test-storage-pool-throughput.name
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

func testAccComputeStoragePoolIamMember_withAndWithoutCondition(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_storage_pool" "test-storage-pool-throughput" {
  name = "tf-test-storage-pool-throughput-%{random_suffix}"

  description = "Hyperdisk Throughput storage pool"

  capacity_provisioning_type   = "STANDARD"
  pool_provisioned_capacity_gb = "10240"

  performance_provisioning_type = "STANDARD"
  pool_provisioned_throughput   = "100"

    storage_pool_type = data.google_compute_storage_pool_types.throughput.self_link

  deletion_protection = false

  zone = "us-central1-a"
}

data "google_project" "project" {}

data "google_compute_storage_pool_types" "throughput" {
  zone = "us-central1-a"
	storage_pool_type = "hyperdisk-throughput"
}

resource "google_compute_storage_pool_iam_member" "foo" {
  project = google_compute_storage_pool.test-storage-pool-throughput.project
  zone = "us-central1-a"
  name = google_compute_storage_pool.test-storage-pool-throughput.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}

resource "google_compute_storage_pool_iam_member" "foo2" {
  project = google_compute_storage_pool.test-storage-pool-throughput.project
  zone = "us-central1-a"
  name = google_compute_storage_pool.test-storage-pool-throughput.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }

  depends_on = [google_compute_storage_pool_iam_member.foo]
}

resource "google_compute_storage_pool_iam_member" "foo3" {
  project = google_compute_storage_pool.test-storage-pool-throughput.project
  zone = "us-central1-a"
  name = google_compute_storage_pool.test-storage-pool-throughput.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    # Check that lack of description doesn't cause any issues
    # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
    title       = "%{condition_title_no_desc}"
    expression  = "%{condition_expr_no_desc}"
  }
  depends_on = [google_compute_storage_pool_iam_member.foo2]
}
`, context)
}

func testAccComputeStoragePoolIamPolicy_withCondition(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_storage_pool" "test-storage-pool-throughput" {
  name = "tf-test-storage-pool-throughput-%{random_suffix}"

  description = "Hyperdisk Throughput storage pool"

  capacity_provisioning_type   = "STANDARD"
  pool_provisioned_capacity_gb = "10240"

  performance_provisioning_type = "STANDARD"
  pool_provisioned_throughput   = "100"

    storage_pool_type = data.google_compute_storage_pool_types.throughput.self_link

  deletion_protection = false

  zone = "us-central1-a"
}

data "google_project" "project" {}

data "google_compute_storage_pool_types" "throughput" {
  zone = "us-central1-a"
	storage_pool_type = "hyperdisk-throughput"
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

resource "google_compute_storage_pool_iam_policy" "foo" {
  project = google_compute_storage_pool.test-storage-pool-throughput.project
  zone = "us-central1-a"
  name = google_compute_storage_pool.test-storage-pool-throughput.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}
