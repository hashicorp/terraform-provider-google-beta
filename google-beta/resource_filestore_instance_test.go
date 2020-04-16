package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccFilestoreInstance_update(t *testing.T) {
	t.Parallel()

	name := fmt.Sprintf("tf-test-%d", randInt(t))

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFilestoreInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFilestoreInstance_update(name),
			},
			{
				ResourceName:      "google_filestore_instance.instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccFilestoreInstance_update2(name),
			},
			{
				ResourceName:      "google_filestore_instance.instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccFilestoreInstance_update(name string) string {
	return fmt.Sprintf(`
resource "google_filestore_instance" "instance" {
  name = "tf-instance-%s"
  zone = "us-central1-b"
  file_shares {
    capacity_gb = 2660
    name        = "share"
  }
  networks {
    network = "default"
    modes   = ["MODE_IPV4"]
  }
  labels = {
    baz = "qux"
  }
  tier        = "PREMIUM"
  description = "An instance created during testing."
}
`, name)
}

func testAccFilestoreInstance_update2(name string) string {
	return fmt.Sprintf(`
resource "google_filestore_instance" "instance" {
  name = "tf-instance-%s"
  zone = "us-central1-b"
  file_shares {
    capacity_gb = 2760
    name        = "share"
  }
  networks {
    network = "default"
    modes   = ["MODE_IPV4"]
  }
  tier        = "PREMIUM"
  description = "A modified instance created during testing."
}
`, name)
}
