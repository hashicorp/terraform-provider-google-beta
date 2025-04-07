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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/vertexai/FeatureGroup.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/examples/base_configs/iam_test_file.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package vertexai_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccVertexAIFeatureGroupIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVertexAIFeatureGroupIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_vertex_ai_feature_group_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/featureGroups/%s roles/viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf_test_example_feature_group%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccVertexAIFeatureGroupIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_vertex_ai_feature_group_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/featureGroups/%s roles/viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf_test_example_feature_group%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccVertexAIFeatureGroupIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccVertexAIFeatureGroupIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_vertex_ai_feature_group_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/featureGroups/%s roles/viewer user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf_test_example_feature_group%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccVertexAIFeatureGroupIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVertexAIFeatureGroupIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_vertex_ai_feature_group_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_vertex_ai_feature_group_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/featureGroups/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf_test_example_feature_group%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccVertexAIFeatureGroupIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_vertex_ai_feature_group_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/featureGroups/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf_test_example_feature_group%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccVertexAIFeatureGroupIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
  resource "google_vertex_ai_feature_group" "feature_group" {
  name = "tf_test_example_feature_group%{random_suffix}"
  description = "A sample feature group"
  region = "us-central1"
  labels = {
      label-one = "value-one"
  }
  big_query {
    big_query_source {
        # The source table must have a column named 'feature_timestamp' of type TIMESTAMP.
        input_uri = "bq://${google_bigquery_table.sample_table.project}.${google_bigquery_table.sample_table.dataset_id}.${google_bigquery_table.sample_table.table_id}"
    }
    entity_id_columns = ["feature_id"]
  }
}

resource "google_bigquery_dataset" "sample_dataset" {
  dataset_id                  = "tf_test_job_load%{random_suffix}_dataset"
  friendly_name               = "test"
  description                 = "This is a test description"
  location                    = "US"
}

resource "google_bigquery_table" "sample_table" {
  deletion_protection = false
  dataset_id = google_bigquery_dataset.sample_dataset.dataset_id
  table_id   = "tf_test_job_load%{random_suffix}_table"

  schema = <<EOF
[
    {
        "name": "feature_id",
        "type": "STRING",
        "mode": "NULLABLE"
    },
    {
        "name": "feature_timestamp",
        "type": "TIMESTAMP",
        "mode": "NULLABLE"
    }
]
EOF
}

resource "google_vertex_ai_feature_group_iam_member" "foo" {
  region = google_vertex_ai_feature_group.feature_group.region
  feature_group = google_vertex_ai_feature_group.feature_group.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccVertexAIFeatureGroupIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
  resource "google_vertex_ai_feature_group" "feature_group" {
  name = "tf_test_example_feature_group%{random_suffix}"
  description = "A sample feature group"
  region = "us-central1"
  labels = {
      label-one = "value-one"
  }
  big_query {
    big_query_source {
        # The source table must have a column named 'feature_timestamp' of type TIMESTAMP.
        input_uri = "bq://${google_bigquery_table.sample_table.project}.${google_bigquery_table.sample_table.dataset_id}.${google_bigquery_table.sample_table.table_id}"
    }
    entity_id_columns = ["feature_id"]
  }
}

resource "google_bigquery_dataset" "sample_dataset" {
  dataset_id                  = "tf_test_job_load%{random_suffix}_dataset"
  friendly_name               = "test"
  description                 = "This is a test description"
  location                    = "US"
}

resource "google_bigquery_table" "sample_table" {
  deletion_protection = false
  dataset_id = google_bigquery_dataset.sample_dataset.dataset_id
  table_id   = "tf_test_job_load%{random_suffix}_table"

  schema = <<EOF
[
    {
        "name": "feature_id",
        "type": "STRING",
        "mode": "NULLABLE"
    },
    {
        "name": "feature_timestamp",
        "type": "TIMESTAMP",
        "mode": "NULLABLE"
    }
]
EOF
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_vertex_ai_feature_group_iam_policy" "foo" {
  region = google_vertex_ai_feature_group.feature_group.region
  feature_group = google_vertex_ai_feature_group.feature_group.name
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_vertex_ai_feature_group_iam_policy" "foo" {
  region = google_vertex_ai_feature_group.feature_group.region
  feature_group = google_vertex_ai_feature_group.feature_group.name
  depends_on = [
    google_vertex_ai_feature_group_iam_policy.foo
  ]
}
`, context)
}

func testAccVertexAIFeatureGroupIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
  resource "google_vertex_ai_feature_group" "feature_group" {
  name = "tf_test_example_feature_group%{random_suffix}"
  description = "A sample feature group"
  region = "us-central1"
  labels = {
      label-one = "value-one"
  }
  big_query {
    big_query_source {
        # The source table must have a column named 'feature_timestamp' of type TIMESTAMP.
        input_uri = "bq://${google_bigquery_table.sample_table.project}.${google_bigquery_table.sample_table.dataset_id}.${google_bigquery_table.sample_table.table_id}"
    }
    entity_id_columns = ["feature_id"]
  }
}

resource "google_bigquery_dataset" "sample_dataset" {
  dataset_id                  = "tf_test_job_load%{random_suffix}_dataset"
  friendly_name               = "test"
  description                 = "This is a test description"
  location                    = "US"
}

resource "google_bigquery_table" "sample_table" {
  deletion_protection = false
  dataset_id = google_bigquery_dataset.sample_dataset.dataset_id
  table_id   = "tf_test_job_load%{random_suffix}_table"

  schema = <<EOF
[
    {
        "name": "feature_id",
        "type": "STRING",
        "mode": "NULLABLE"
    },
    {
        "name": "feature_timestamp",
        "type": "TIMESTAMP",
        "mode": "NULLABLE"
    }
]
EOF
}

data "google_iam_policy" "foo" {
}

resource "google_vertex_ai_feature_group_iam_policy" "foo" {
  region = google_vertex_ai_feature_group.feature_group.region
  feature_group = google_vertex_ai_feature_group.feature_group.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccVertexAIFeatureGroupIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
  resource "google_vertex_ai_feature_group" "feature_group" {
  name = "tf_test_example_feature_group%{random_suffix}"
  description = "A sample feature group"
  region = "us-central1"
  labels = {
      label-one = "value-one"
  }
  big_query {
    big_query_source {
        # The source table must have a column named 'feature_timestamp' of type TIMESTAMP.
        input_uri = "bq://${google_bigquery_table.sample_table.project}.${google_bigquery_table.sample_table.dataset_id}.${google_bigquery_table.sample_table.table_id}"
    }
    entity_id_columns = ["feature_id"]
  }
}

resource "google_bigquery_dataset" "sample_dataset" {
  dataset_id                  = "tf_test_job_load%{random_suffix}_dataset"
  friendly_name               = "test"
  description                 = "This is a test description"
  location                    = "US"
}

resource "google_bigquery_table" "sample_table" {
  deletion_protection = false
  dataset_id = google_bigquery_dataset.sample_dataset.dataset_id
  table_id   = "tf_test_job_load%{random_suffix}_table"

  schema = <<EOF
[
    {
        "name": "feature_id",
        "type": "STRING",
        "mode": "NULLABLE"
    },
    {
        "name": "feature_timestamp",
        "type": "TIMESTAMP",
        "mode": "NULLABLE"
    }
]
EOF
}

resource "google_vertex_ai_feature_group_iam_binding" "foo" {
  region = google_vertex_ai_feature_group.feature_group.region
  feature_group = google_vertex_ai_feature_group.feature_group.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccVertexAIFeatureGroupIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
  resource "google_vertex_ai_feature_group" "feature_group" {
  name = "tf_test_example_feature_group%{random_suffix}"
  description = "A sample feature group"
  region = "us-central1"
  labels = {
      label-one = "value-one"
  }
  big_query {
    big_query_source {
        # The source table must have a column named 'feature_timestamp' of type TIMESTAMP.
        input_uri = "bq://${google_bigquery_table.sample_table.project}.${google_bigquery_table.sample_table.dataset_id}.${google_bigquery_table.sample_table.table_id}"
    }
    entity_id_columns = ["feature_id"]
  }
}

resource "google_bigquery_dataset" "sample_dataset" {
  dataset_id                  = "tf_test_job_load%{random_suffix}_dataset"
  friendly_name               = "test"
  description                 = "This is a test description"
  location                    = "US"
}

resource "google_bigquery_table" "sample_table" {
  deletion_protection = false
  dataset_id = google_bigquery_dataset.sample_dataset.dataset_id
  table_id   = "tf_test_job_load%{random_suffix}_table"

  schema = <<EOF
[
    {
        "name": "feature_id",
        "type": "STRING",
        "mode": "NULLABLE"
    },
    {
        "name": "feature_timestamp",
        "type": "TIMESTAMP",
        "mode": "NULLABLE"
    }
]
EOF
}

resource "google_vertex_ai_feature_group_iam_binding" "foo" {
  region = google_vertex_ai_feature_group.feature_group.region
  feature_group = google_vertex_ai_feature_group.feature_group.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
