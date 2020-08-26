package google

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccHealthcareDatasetIdParsing(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		ImportId            string
		ExpectedError       bool
		ExpectedTerraformId string
		ExpectedDatasetId   string
		Config              *Config
	}{
		"id is in project/location/datasetName format": {
			ImportId:            "test-project/us-central1/test-dataset",
			ExpectedError:       false,
			ExpectedTerraformId: "test-project/us-central1/test-dataset",
			ExpectedDatasetId:   "projects/test-project/locations/us-central1/datasets/test-dataset",
		},
		"id is in domain:project/location/datasetName format": {
			ImportId:            "example.com:test-project/us-central1/test-dataset",
			ExpectedError:       false,
			ExpectedTerraformId: "example.com:test-project/us-central1/test-dataset",
			ExpectedDatasetId:   "projects/example.com:test-project/locations/us-central1/datasets/test-dataset",
		},
		"id is in location/datasetName format": {
			ImportId:            "us-central1/test-dataset",
			ExpectedError:       false,
			ExpectedTerraformId: "test-project/us-central1/test-dataset",
			ExpectedDatasetId:   "projects/test-project/locations/us-central1/datasets/test-dataset",
			Config:              &Config{Project: "test-project"},
		},
		"id is in location/datasetName format without project in config": {
			ImportId:      "us-central1/test-dataset",
			ExpectedError: true,
			Config:        &Config{Project: ""},
		},
	}

	for tn, tc := range cases {
		datasetId, err := parseHealthcareDatasetId(tc.ImportId, tc.Config)

		if tc.ExpectedError && err == nil {
			t.Fatalf("bad: %s, expected an error", tn)
		}

		if err != nil {
			if tc.ExpectedError {
				continue
			}
			t.Fatalf("bad: %s, err: %#v", tn, err)
		}

		if datasetId.terraformId() != tc.ExpectedTerraformId {
			t.Fatalf("bad: %s, expected Terraform ID to be `%s` but is `%s`", tn, tc.ExpectedTerraformId, datasetId.terraformId())
		}

		if datasetId.datasetId() != tc.ExpectedDatasetId {
			t.Fatalf("bad: %s, expected Dataset ID to be `%s` but is `%s`", tn, tc.ExpectedDatasetId, datasetId.datasetId())
		}
	}
}

func TestAccProjectServiceIdentity_basic(t *testing.T) {
	t.Parallel()

	resourceName := "google_project_service_identity.hc_sa"

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckHealthcareDatasetDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testGoogleProjectServiceIdentity_basic(),
			},
			{
				ResourceName: resourceName,
			},
		},
	})
}

func testGoogleProjectServiceIdentity_basic() string {
	return `
resource "google_project_service_identity" "hc_sa" {
  service = "healthcare.googleapis.com"
}

resource "google_project_iam_member" "hc_sa_bq_jobuser" {
    role    = "roles/bigquery.jobUser"
    member  = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-healthcare.iam.gserviceaccount.com"
}`
}
