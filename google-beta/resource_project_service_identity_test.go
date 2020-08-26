package google

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

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
