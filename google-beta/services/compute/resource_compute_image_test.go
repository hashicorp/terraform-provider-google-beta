// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	tpgcompute "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/compute"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	compute "google.golang.org/api/compute/v0.beta"
)

func TestAccComputeImage_withLicense(t *testing.T) {
	t.Parallel()

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeImageDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeImage_license("image-test-" + acctest.RandString(t, 10)),
			},
			{
				ResourceName:            "google_compute_image.foobar",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels"},
			},
		},
	})
}

func TestAccComputeImage_update(t *testing.T) {
	t.Parallel()

	var image compute.Image

	context := map[string]interface{}{
		"name":            "image-test-" + acctest.RandString(t, 10),
		"disk_image_path": "./test-fixtures/raw-disk-image.tar.gz",
		"bucket_one":      "tf-test-compute-image-bucket-" + acctest.RandString(t, 10),
		"bucket_two":      "tf-test-compute-image-bucket-" + acctest.RandString(t, 10),
	}

	// Only labels supports an update
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeImageDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeImage_basic(context),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeImageExists(
						t, "google_compute_image.foobar", &image),
					testAccCheckComputeImageContainsLabel(&image, "my-label", "my-label-value"),
				),
			},
			{
				Config: testAccComputeImage_update(context),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeImageExists(
						t, "google_compute_image.foobar", &image),
					testAccCheckComputeImageDoesNotContainLabel(&image, "my-label"),
					testAccCheckComputeImageContainsLabel(&image, "empty-label", "oh-look-theres-a-label-now"),
					testAccCheckComputeImageContainsLabel(&image, "new-field", "only-shows-up-when-updated"),
				),
			},
			{
				ResourceName:            "google_compute_image.foobar",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"raw_disk", "labels", "terraform_labels"},
			},
		},
	})
}

func TestAccComputeImage_basedondisk(t *testing.T) {
	t.Parallel()

	var image compute.Image

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeImageDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeImage_basedondisk(acctest.RandString(t, 10), acctest.RandString(t, 10)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeImageExists(
						t, "google_compute_image.foobar", &image),
					testAccCheckComputeImageHasSourceType(&image),
				),
			},
			{
				ResourceName:      "google_compute_image.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeImage_shieldedInstance_InitialState(t *testing.T) {
	t.Parallel()

	var image compute.Image
	imageName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeImageDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeImage_shieldedInstance_InitialState(imageName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeImageExists(
						t, "google_compute_image.foobar", &image),
					testAccCheckComputeImageHasShieldedInstanceInitialState(&image),
				),
			},
			{
				ResourceName:      "google_compute_image.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeImage_shieldedInstance_UpdatedState(t *testing.T) {
	t.Parallel()

	var image compute.Image
	imageName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeImageDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeImage_shieldedInstance_InitialState(imageName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeImageExists(
						t, "google_compute_image.foobar", &image),
					testAccCheckComputeImageHasShieldedInstanceInitialState(&image),
				),
			},
			{
				ResourceName:      "google_compute_image.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeImage_shieldedInstance_UpdatedState(imageName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeImageExists(
						t, "google_compute_image.foobar", &image),
					testAccCheckComputeImageHasShieldedInstanceUpdatedState(&image),
				),
			},
			{
				ResourceName:      "google_compute_image.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeImage_sourceImage(t *testing.T) {
	t.Parallel()

	var image compute.Image
	imageName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeImageDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeImage_sourceImage(imageName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeImageExists(
						t, "google_compute_image.foobar", &image),
					testAccCheckComputeImageHasSourceType(&image),
				),
			},
			{
				ResourceName:      "google_compute_image.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeImage_sourceSnapshot(t *testing.T) {
	t.Parallel()

	var image compute.Image

	diskName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	snapshotName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	imageName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeImageDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeImage_sourceSnapshot(diskName, snapshotName, imageName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeImageExists(
						t, "google_compute_image.foobar", &image),
					testAccCheckComputeImageHasSourceType(&image),
				),
			},
			{
				ResourceName:      "google_compute_image.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckComputeImageExists(t *testing.T, n string, image *compute.Image) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.Attributes["name"] == "" {
			return fmt.Errorf("No name is set")
		}

		config := acctest.GoogleProviderConfig(t)

		found, err := config.NewComputeClient(config.UserAgent).Images.Get(
			config.Project, rs.Primary.Attributes["name"]).Do()
		if err != nil {
			return err
		}

		if found.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Image not found")
		}

		*image = *found

		return nil
	}
}

func TestAccComputeImage_resolveImage(t *testing.T) {
	t.Parallel()

	var image compute.Image
	rand := acctest.RandString(t, 10)
	name := fmt.Sprintf("test-image-%s", rand)
	fam := fmt.Sprintf("test-image-family-%s", rand)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeImageDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeImage_resolving(name, fam),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeImageExists(
						t, "google_compute_image.foobar", &image),
					testAccCheckComputeImageResolution(t, "google_compute_image.foobar"),
				),
			},
		},
	})
}

func TestAccComputeImage_imageEncryptionKey(t *testing.T) {
	t.Parallel()

	kmsKey := acctest.BootstrapKMSKeyInLocation(t, "us-central1")
	kmsKeyName := tpgresource.GetResourceNameFromSelfLink(kmsKey.CryptoKey.Name)
	kmsRingName := tpgresource.GetResourceNameFromSelfLink(kmsKey.KeyRing.Name)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeInstanceTemplateDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeImage_imageEncryptionKey(kmsRingName, kmsKeyName, acctest.RandString(t, 10)),
			},
			{
				ResourceName:      "google_compute_image.image",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeImage_sourceImageEncryptionKey(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":     acctest.RandString(t, 10),
		"kms_key_self_link": acctest.BootstrapKMSKeyInLocation(t, "us-central1").CryptoKey.Name,
		"raw_key":           "SGVsbG8gZnJvbSBHb29nbGUgQ2xvdWQgUGxhdGZvcm0=",
		"rsa_encrypted_key": "fB6BS8tJGhGVDZDjGt1pwUo2wyNbkzNxgH1avfOtiwB9X6oPG94gWgenygitnsYJyKjdOJ7DyXLmxwQOSmnCYCUBWdKCSssyLV5907HL2mb5TfqmgHk5JcArI/t6QADZWiuGtR+XVXqiLa5B9usxFT2BTmbHvSKfkpJ7McCNc/3U0PQR8euFRZ9i75o/w+pLHFMJ05IX3JB0zHbXMV173PjObiV3ItSJm2j3mp5XKabRGSA5rmfMnHIAMz6stGhcuom6+bMri2u/axmPsdxmC6MeWkCkCmPjaKsVz1+uQUNCJkAnzesluhoD+R6VjFDm4WI7yYabu4MOOAOTaQXdEg==",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeInstanceTemplateDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeImage_sourceImageEncryptionKey(context),
			},
		},
	})
}

func TestAccComputeImage_sourceSnapshotEncryptionKey(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":     acctest.RandString(t, 10),
		"kms_key_self_link": acctest.BootstrapKMSKeyInLocation(t, "us-central1").CryptoKey.Name,
		"raw_key":           "SGVsbG8gZnJvbSBHb29nbGUgQ2xvdWQgUGxhdGZvcm0=",
		"rsa_encrypted_key": "fB6BS8tJGhGVDZDjGt1pwUo2wyNbkzNxgH1avfOtiwB9X6oPG94gWgenygitnsYJyKjdOJ7DyXLmxwQOSmnCYCUBWdKCSssyLV5907HL2mb5TfqmgHk5JcArI/t6QADZWiuGtR+XVXqiLa5B9usxFT2BTmbHvSKfkpJ7McCNc/3U0PQR8euFRZ9i75o/w+pLHFMJ05IX3JB0zHbXMV173PjObiV3ItSJm2j3mp5XKabRGSA5rmfMnHIAMz6stGhcuom6+bMri2u/axmPsdxmC6MeWkCkCmPjaKsVz1+uQUNCJkAnzesluhoD+R6VjFDm4WI7yYabu4MOOAOTaQXdEg==",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeInstanceTemplateDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeImage_sourceSnapshotEncryptionKey(context),
			},
		},
	})
}

func TestAccComputeImage_sourceDiskEncryptionKey(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":     acctest.RandString(t, 10),
		"kms_key_self_link": acctest.BootstrapKMSKeyInLocation(t, "us-central1").CryptoKey.Name,
		"raw_key":           "SGVsbG8gZnJvbSBHb29nbGUgQ2xvdWQgUGxhdGZvcm0=",
		"rsa_encrypted_key": "fB6BS8tJGhGVDZDjGt1pwUo2wyNbkzNxgH1avfOtiwB9X6oPG94gWgenygitnsYJyKjdOJ7DyXLmxwQOSmnCYCUBWdKCSssyLV5907HL2mb5TfqmgHk5JcArI/t6QADZWiuGtR+XVXqiLa5B9usxFT2BTmbHvSKfkpJ7McCNc/3U0PQR8euFRZ9i75o/w+pLHFMJ05IX3JB0zHbXMV173PjObiV3ItSJm2j3mp5XKabRGSA5rmfMnHIAMz6stGhcuom6+bMri2u/axmPsdxmC6MeWkCkCmPjaKsVz1+uQUNCJkAnzesluhoD+R6VjFDm4WI7yYabu4MOOAOTaQXdEg==",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeInstanceTemplateDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeImage_sourceDiskEncryptionKey(context),
			},
		},
	})
}

func testAccCheckComputeImageResolution(t *testing.T, n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config := acctest.GoogleProviderConfig(t)
		project := config.Project

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Resource not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}
		if rs.Primary.Attributes["name"] == "" {
			return fmt.Errorf("No image name is set")
		}
		if rs.Primary.Attributes["family"] == "" {
			return fmt.Errorf("No image family is set")
		}
		if rs.Primary.Attributes["self_link"] == "" {
			return fmt.Errorf("No self_link is set")
		}

		name := rs.Primary.Attributes["name"]
		family := rs.Primary.Attributes["family"]
		link := rs.Primary.Attributes["self_link"]

		latestDebian, err := config.NewComputeClient(config.UserAgent).Images.GetFromFamily("debian-cloud", "debian-11").Do()
		if err != nil {
			return fmt.Errorf("Error retrieving latest debian: %s", err)
		}

		images := map[string]string{
			"family/" + latestDebian.Family:                            "projects/debian-cloud/global/images/family/" + latestDebian.Family,
			"projects/debian-cloud/global/images/" + latestDebian.Name: "projects/debian-cloud/global/images/" + latestDebian.Name,
			latestDebian.Family:                                        "projects/debian-cloud/global/images/family/" + latestDebian.Family,
			latestDebian.Name:                                          "projects/debian-cloud/global/images/" + latestDebian.Name,
			latestDebian.SelfLink:                                      latestDebian.SelfLink,

			"global/images/" + name:          "global/images/" + name,
			"global/images/family/" + family: "global/images/family/" + family,
			name:                             "global/images/" + name,
			family:                           "global/images/family/" + family,
			"family/" + family:               "global/images/family/" + family,
			project + "/" + name:             "projects/" + project + "/global/images/" + name,
			project + "/" + family:           "projects/" + project + "/global/images/family/" + family,
			link:                             link,
		}

		for input, expectation := range images {
			result, err := tpgcompute.ResolveImage(config, project, input, config.UserAgent)
			if err != nil {
				return fmt.Errorf("Error resolving input %s to image: %+v\n", input, err)
			}
			if result != expectation {
				return fmt.Errorf("Expected input '%s' to resolve to '%s', it resolved to '%s' instead.\n", input, expectation, result)
			}
		}
		return nil
	}
}

func testAccCheckComputeImageContainsLabel(image *compute.Image, key string, value string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		v, ok := image.Labels[key]
		if !ok {
			return fmt.Errorf("Expected label with key '%s' not found", key)
		}
		if v != value {
			return fmt.Errorf("Incorrect label value for key '%s': expected '%s' but found '%s'", key, value, v)
		}
		return nil
	}
}

func testAccCheckComputeImageDoesNotContainLabel(image *compute.Image, key string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if v, ok := image.Labels[key]; ok {
			return fmt.Errorf("Expected no label for key '%s' but found one with value '%s'", key, v)
		}

		return nil
	}
}

func testAccCheckComputeImageHasShieldedInstanceInitialState(image *compute.Image) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if image.ShieldedInstanceInitialState == nil {
			return fmt.Errorf("Expected ShieldedInstanceInitialState to be present")
		}

		// Check PK values
		if image.ShieldedInstanceInitialState.Pk == nil {
			return fmt.Errorf("Expected PK to be present")
		}
		if image.ShieldedInstanceInitialState.Pk.FileType != "X509" {
			return fmt.Errorf("Expected PK FileType to be X509, got %s", image.ShieldedInstanceInitialState.Pk.FileType)
		}

		// Check KEK values
		if len(image.ShieldedInstanceInitialState.Keks) != 1 {
			return fmt.Errorf("Expected 1 KEK entry, got %d", len(image.ShieldedInstanceInitialState.Keks))
		}
		if image.ShieldedInstanceInitialState.Keks[0].FileType != "X509" {
			return fmt.Errorf("Expected KEK FileType to be X509, got %s", image.ShieldedInstanceInitialState.Keks[0].FileType)
		}

		// Check DB values
		if len(image.ShieldedInstanceInitialState.Dbs) != 1 {
			return fmt.Errorf("Expected 1 DB entry, got %d", len(image.ShieldedInstanceInitialState.Dbs))
		}
		if image.ShieldedInstanceInitialState.Dbs[0].FileType != "X509" {
			return fmt.Errorf("Expected DB FileType to be X509, got %s", image.ShieldedInstanceInitialState.Dbs[0].FileType)
		}

		// Check DBX values
		if len(image.ShieldedInstanceInitialState.Dbxs) != 1 {
			return fmt.Errorf("Expected 1 DBX entry, got %d", len(image.ShieldedInstanceInitialState.Dbxs))
		}
		if image.ShieldedInstanceInitialState.Dbxs[0].FileType != "X509" {
			return fmt.Errorf("Expected DBX FileType to be X509, got %s", image.ShieldedInstanceInitialState.Dbxs[0].FileType)
		}

		return nil
	}
}

func testAccCheckComputeImageHasShieldedInstanceUpdatedState(image *compute.Image) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if image.ShieldedInstanceInitialState == nil {
			return fmt.Errorf("Expected ShieldedInstanceInitialState to be present")
		}

		// Check PK values - now BIN
		if image.ShieldedInstanceInitialState.Pk == nil {
			return fmt.Errorf("Expected PK to be present")
		}
		if image.ShieldedInstanceInitialState.Pk.FileType != "BIN" {
			return fmt.Errorf("Expected PK FileType to be BIN, got %s", image.ShieldedInstanceInitialState.Pk.FileType)
		}

		// Check KEK values - now BIN
		if len(image.ShieldedInstanceInitialState.Keks) != 1 {
			return fmt.Errorf("Expected 1 KEK entry, got %d", len(image.ShieldedInstanceInitialState.Keks))
		}
		if image.ShieldedInstanceInitialState.Keks[0].FileType != "BIN" {
			return fmt.Errorf("Expected KEK FileType to be BIN, got %s", image.ShieldedInstanceInitialState.Keks[0].FileType)
		}

		// Check DB values - now BIN
		if len(image.ShieldedInstanceInitialState.Dbs) != 1 {
			return fmt.Errorf("Expected 1 DB entry, got %d", len(image.ShieldedInstanceInitialState.Dbs))
		}
		if image.ShieldedInstanceInitialState.Dbs[0].FileType != "BIN" {
			return fmt.Errorf("Expected DB FileType to be BIN, got %s", image.ShieldedInstanceInitialState.Dbs[0].FileType)
		}

		// Check DBX values - now missing
		if len(image.ShieldedInstanceInitialState.Dbxs) != 0 {
			return fmt.Errorf("Expected no DBX entry, got %d", len(image.ShieldedInstanceInitialState.Dbxs))
		}

		return nil
	}
}

func testAccCheckComputeImageHasSourceType(image *compute.Image) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if image.SourceType == "" {
			return fmt.Errorf("No source disk")
		}
		return nil
	}
}

func testAccComputeImage_resolving(name, family string) string {
	return fmt.Sprintf(`
data "google_compute_image" "my_image" {
  family  = "debian-11"
  project = "debian-cloud"
}

resource "google_compute_disk" "foobar" {
  name  = "%s"
  zone  = "us-central1-a"
  image = data.google_compute_image.my_image.self_link
}

resource "google_compute_image" "foobar" {
  name        = "%s"
  family      = "%s"
  source_disk = google_compute_disk.foobar.self_link
}
`, name, name, family)
}

func testAccComputeImage_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_storage_bucket" "bucket_one" {
  name     = "%{bucket_one}"
  location = "US"
  uniform_bucket_level_access = true
}

resource "google_storage_bucket_object" "object" {
  name   = "raw-disk-image.tar.gz"
  bucket = google_storage_bucket.bucket_one.name
  source = "%{disk_image_path}"
}

resource "google_compute_image" "foobar" {
  name        = "%{name}"
  description = "description-test"
  family      = "family-test"
  raw_disk {
    source = "https://${google_storage_bucket.bucket_one.name}.storage.googleapis.com/${google_storage_bucket_object.object.name}"
  }
  labels = {
    my-label    = "my-label-value"
  }
}
`, context)
}

func testAccComputeImage_license(name string) string {
	return fmt.Sprintf(`
data "google_compute_image" "my_image" {
  family  = "debian-11"
  project = "debian-cloud"
}

resource "google_compute_disk" "foobar" {
  name  = "disk-test-%s"
  zone  = "us-central1-a"
  image = data.google_compute_image.my_image.self_link
}

resource "google_compute_image" "foobar" {
  name        = "%s"
  description = "description-test"
  source_disk = google_compute_disk.foobar.self_link

  labels = {
    my-label    = "my-label-value"
  }
  licenses = [
    "https://www.googleapis.com/compute/v1/projects/debian-cloud/global/licenses/debian-11-bullseye",
  ]
}
`, name, name)
}

func testAccComputeImage_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_storage_bucket" "bucket_two" {
  name     = "%{bucket_two}"
  location = "US"
  uniform_bucket_level_access = true
}

resource "google_storage_bucket_object" "object" {
  name   = "raw-disk-image.tar.gz"
  bucket = google_storage_bucket.bucket_two.name
  source = "%{disk_image_path}"
}

resource "google_compute_image" "foobar" {
  name        = "%{name}"
  description = "description-test"
  family      = "family-test"
  raw_disk {
    source = "https://${google_storage_bucket.bucket_two.name}.storage.googleapis.com/${google_storage_bucket_object.object.name}"
  }
  labels = {
    empty-label = "oh-look-theres-a-label-now"
    new-field   = "only-shows-up-when-updated"
  }
}
`, context)
}

func testAccComputeImage_basedondisk(diskName, imageName string) string {
	return fmt.Sprintf(`
data "google_compute_image" "my_image" {
  family  = "debian-11"
  project = "debian-cloud"
}

resource "google_compute_disk" "foobar" {
  name  = "disk-test-%s"
  zone  = "us-central1-a"
  image = data.google_compute_image.my_image.self_link
}

resource "google_compute_image" "foobar" {
  name        = "image-test-%s"
  source_disk = google_compute_disk.foobar.self_link
}
`, diskName, imageName)
}

func testAccComputeImage_shieldedInstance_InitialState(imageName string) string {
	return fmt.Sprintf(`
data "google_compute_image" "my_image" {
  family  = "debian-11"
  project = "debian-cloud"
}

resource "google_compute_image" "foobar" {
  name         = "%s"
  source_image = data.google_compute_image.my_image.self_link
  shielded_instance_initial_state {
    pk {
      content   = "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURDVENDQWZHZ0F3SUJBZ0lVQVNIVCtxS0Z0MjNweTFXbmk3UWFDYi9ibzh3d0RRWUpLb1pJaHZjTkFRRUwKQlFBd0ZERVNNQkFHQTFVRUF3d0pibTlpYjJSNUlGQkxNQjRYRFRJMU1ERXdPREV4TURZMU5Gb1hEVE0xTURFdwpOakV4TURZMU5Gb3dGREVTTUJBR0ExVUVBd3dKYm05aWIyUjVJRkJMTUlJQklqQU5CZ2txaGtpRzl3MEJBUUVGCkFBT0NBUThBTUlJQkNnS0NBUUVBbkNab3Ezdk92bTk4K2FVdW14SFpsOXZ2R3pVc2FvZkUvZnJWcVRrMDVFTHYKVkJKM0ovUWpKY0VpNGtxdnVGT3Rnb1UwcnFpOG81aW5IMnpyQThtQkp1akVEekRJS1NLRFhZVThmVVk1Skw5Qwo0a1RQT0xIeEFQeVR2QjJ5dzAzUkhpUTJPUE5YY2JEU3VHcnBIMVlSb29aMnhNUzZtek5TOUxEN2Y1RXg1cnpUCmpsaHhCY1RqUUhnakptTmZQVk5aWWMwcy9yVjZTbUViYTRraGlWem1VaE9NVzhFZ2ZuVUc2d1VBTG5hZVFpaUMKdmFITi9YdTd5WjVjS3ovTFIyeDRJTFJ6WUNCZXVpdjk1ZHdGem8yNk13L1Z2Q29wUGxZdkdEKzI5Y1dXaEFVMgp0d2VyUGQrL2c1dWkxNTEwdVA4Zk0rc0pCMFJmT3hHMDloS2tVUFR4Y1FJREFRQUJvMU13VVRBZEJnTlZIUTRFCkZnUVVvVWtFRVNTcVVKZi9hRXYxSnI3bStDQTFPVEV3SHdZRFZSMGpCQmd3Rm9BVW9Va0VFU1NxVUpmL2FFdjEKSnI3bStDQTFPVEV3RHdZRFZSMFRBUUgvQkFVd0F3RUIvekFOQmdrcWhraUc5dzBCQVFzRkFBT0NBUUVBakRmTwoxMVc1U09KcndveUZwS0ozRm9qV2RNT3Z6TXBGQzJCZkFDQkxlMS9jR29QTUxIUWNEVVIra21FN2V5L042eEE2CjlPUGhXdmxjZ1g2KzlkUThaTzEzbGl0Zmtwakg2UVlUUEp0ZHE0OHFpeFVLV1U3MGNhUURoL2RsRnZlVG00bDUKdE8zSFlDQ0NnNjlCY05neUtkcGZxMFpxVnNidWRyNlFlN2RCbEJkQ1dNSUo1R0pzWnlnNmFSZUhoSnhrb3ZtYgpMeHgwR2lFMUprNlVDRU9vUFo4Z3l6VmUvS3QyL1VJUXgvNlg1RFFvODZmWmNNQ3lpWENNcDZWOXdZcmpXcDFxCnlOQVljRDBiSUZZbEZlMHJlK0h4WTZ3bUcwSjlUa2tNWVFjRTlVVk4waThHZWNwM3YvcHIxeVExM0VseEFVTS8KbEt1LzZNU1R3bTdzNzdrRWJBPT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
      file_type = "X509"
    }
    dbs {
      content   = "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURDVENDQWZHZ0F3SUJBZ0lVUmlvUEVOY1dYM21qdzJYZEFSMlo2T1dUb0p3d0RRWUpLb1pJaHZjTkFRRUwKQlFBd0ZERVNNQkFHQTFVRUF3d0pibTlpYjJSNUlFUkNNQjRYRFRJMU1ERXdPREV4TURZMU5Gb1hEVE0xTURFdwpOakV4TURZMU5Gb3dGREVTTUJBR0ExVUVBd3dKYm05aWIyUjVJRVJDTUlJQklqQU5CZ2txaGtpRzl3MEJBUUVGCkFBT0NBUThBTUlJQkNnS0NBUUVBdjJndll6b0dDaEF2QzhKcEZpYjRmZFFpLytjemtYak91RnJydTBFMFV2anUKc003RVNvMi9sc0hEcjBJbm5QRzRUOVA5Zm51NWV4enlaeDRjc3lOeVJuYUo5ZVMycmZTZjF4aHNJWmdNNlB3UApmSm9yd09mRkdUYVQwVnZrOFhCT2VPN2hhMnRuN1diZkpnV28wQWlKSGpvY3JoZkhOckdIY3dkcGZKbTFNWHNXCllWMUNaa25nNWczWUpSRzJpUVdoTnZQMjVRcnN5d3ZnRGZLUUNvMnZZbjZlL1grT2MwOG14Ykp3MHlhcVNFRUEKTHg0Q0tFNUNhd2ppVlVGWCsxQm1uRG4yOVhUUHZSbG44UG82S3pacnFYV1VNV1FxOWNZOUJocGt0QXVVbDZsbwpaRUN2V20zUFhjOElXZkZ4LzJEY2dyYjNwb1BRZmRjUnM0bGNzVWxXTlFJREFRQUJvMU13VVRBZEJnTlZIUTRFCkZnUVVNak04WURHb3VGRVBHQ0RSS2Vmb2VsYmJKbDh3SHdZRFZSMGpCQmd3Rm9BVU1qTThZREdvdUZFUEdDRFIKS2Vmb2VsYmJKbDh3RHdZRFZSMFRBUUgvQkFVd0F3RUIvekFOQmdrcWhraUc5dzBCQVFzRkFBT0NBUUVBR2wrVQpoNk42Q1Z2UWRUT1lBN3dOQk9zVk4wOHUySkZiN2k2Y0NxTnZ5ZVFKSk9lV0h4Vmthb25aUG5NR0I4NFRrODJFCmRwUVR3aFJVcXpxdVZtbk8xcnhXODZmZ3lveTU5NE80TFpzNjhHV2tEYkVXVzJ2K0lzRnBkZjBncS9TaXhjUmQKbklrWThVTWRJOGt3aEhRMEVYUnZ1MFVLOVBuS0xBdy9GenQyaThnbmk1UXg3R2xralJnOGdrMXAyTDJQWTdXSwpENi94dDBUcjlUVkprRXYvUHROdTBCQ0pTQ3JTbWcxSGJPN1lXQU1CemoyL1VVcUM4d3BxWVdGOUl1WDhTcGRaCjJvTEJuVkNnd1N4aTU4QkVUY2pINk5BOEVSenJScUVsQm1EYVJ6eUREQThYVUw3N1RvM3VIdytaRjNMcUgwUkoKRlF4UjN0VXBYK1pxdWZ3M3VRPT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
      file_type = "X509"
    }
    dbxs {
      content   = "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"
      file_type = "X509"
    }
    keks {
      content   = "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURDekNDQWZPZ0F3SUJBZ0lVYWMzZG1qREh5bURKVmFnUHpZNWM2bUsweEU0d0RRWUpLb1pJaHZjTkFRRUwKQlFBd0ZURVRNQkVHQTFVRUF3d0tibTlpYjJSNUlFdEZTekFlRncweU5UQXhNRGd4TVRBMk5UUmFGdzB6TlRBeApNRFl4TVRBMk5UUmFNQlV4RXpBUkJnTlZCQU1NQ201dlltOWtlU0JMUlVzd2dnRWlNQTBHQ1NxR1NJYjNEUUVCCkFRVUFBNElCRHdBd2dnRUtBb0lCQVFES2F5K0JXcDAyM0JpL201cUFwcXg5Z3JBSUFsQzgrRGRZMnpJbFNEb3MKblozSi8vMUxwWlU1N3BubzZuMG1GbmhrMGE5SWFRN1F4Z3Z6bWJzVGhmLzVvVWs3ZHl6Qk91S2NDenV2dnJGNQpCYWVPdUg4V3JLRS9zVUpuK0wvVXZubmVtN3BUaWtyemgxTjZuRVIvaHJkSUwvUzVHM09OT0E2clg1MDQ4UDhTCm0xT3kxYmlBaHNWY1RqVDZTNWFPMzJSWnNIUEJxZS9PbWJwRTZEVjNDSzJ3OFFpYTNvYUdtZFBySEdzNjZXa1IKNExHamhLMUVFaWlvUFFhQ2o4Zkk4SjBvcW1IOHF4WnE5enVQVXZuVkZ6VG12KzBLMmFhWFJJSVpCN1Q4RENqMwpPSS8rdEw3MDJ2M3BUTC8yNGoxOEVrN2RPY2hYdlhtN1VhSW5MZHJFeE9HN0FnTUJBQUdqVXpCUk1CMEdBMVVkCkRnUVdCQlF0Nk02M3p3TmdMS3lCckJ2d2dFR0RURXlJNGpBZkJnTlZIU01FR0RBV2dCUXQ2TTYzendOZ0xLeUIKckJ2d2dFR0RURXlJNGpBUEJnTlZIUk1CQWY4RUJUQURBUUgvTUEwR0NTcUdTSWIzRFFFQkN3VUFBNElCQVFBYQpPSmdKZHVHK2dTMU5ZM3VIM2JVOUhZbjJXMi9KU0tWc1NaYnh5NnlZaHpNZXJES2xrd1RVQ21iR1lRbUp1K1hXCmhZS0p6K2FCdlBhWXZ2ZDFRb3lXVlFBd2F3UjlVcm1ubW9US2UzM3ZKUkg4emVYOVZhK3g2UzR3R2Z4LzQwbFMKeUFBZXNHeDB4Z0h5dGkyRUJWM0t1WlM2ZHd0RjAybS90dmZpSExUVzh4NHh3UFNyTlNFNHY4V3loeG1OUXk2MwpId2lOWWlydGYrTENoWnlqSzMyMHlkbnBkM28rNXJzZW9ob0o4Z05TbHhSSTlwb1Fka1BWRnROYVgwTjNhd3lCCjdwZHFVeU9XU2xpRHVpdXVZYi9XTGNuQlk4blZKUy9kTFMzRUVJem1yaSs5aDlVb2o1U0MwWk9FQ1pBdG5SYk8KK3JYUWpkTkkyb0FRMTFXRGhRQTEKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
      file_type = "X509"
    }
}
}
`, imageName)
}

func testAccComputeImage_shieldedInstance_UpdatedState(imageName string) string {
	return fmt.Sprintf(`
data "google_compute_image" "my_image" {
  family  = "debian-11"
  project = "debian-cloud"
}

resource "google_compute_image" "foobar" {
  name         = "%s"
  source_image = data.google_compute_image.my_image.self_link
  shielded_instance_initial_state {
    pk {
      content   = "MIIEowIBAAKCAQEA8ybZA+UQZvwK4kX1PdBkYYPGEx/SvHpCU2FaRHHs7OnmqVHI+KnANgPPpqQYHmtTwdZuGZPk0iU5/JkPJJ1PnNBXA9RlGrKVHXa+OVlgJPiJP5qDgqK5n2E9IF5kJHhylG5w4RXXspRWBHJELGALQz/Bu2qZCeiKh7FslZxHEJTs4ynDLJUxZBhRHBKKog/mL2nEcX1hx4LgCXzVJvzgxI+uDxHM+GMeNPtPQTX9Kt0zC7l0ajQpKY6Nyul5T+HvqF7a+4GaD1glXI2XUYv/WKo5Aj0KGzDhqBtXgFhSVtAEKpk9/+UvWoHMtEjRFQk+2WPdPZsqGVYo8Aw0ZLBCdQIDAQABAoIBAQDXiBXJqZDk0tC4"
      file_type = "BIN"
    }
    dbs {
      content   = "MIIEowIBAAKCAQEA8ybZA+UQZvwK4kX1PdBkYYPGEx/SvHpCU2FaRHHs7OnmqVHI+KnANgPPpqQYHmtTwdZuGZPk0iU5/JkPJJ1PnNBXA9RlGrKVHXa+OVlgJPiJP5qDgqK5n2E9IF5kJHhylG5w4RXXspRWBHJELGALQz/Bu2qZCeiKh7FslZxHEJTs4ynDLJUxZBhRHBKKog/mL2nEcX1hx4LgCXzVJvzgxI+uDxHM+GMeNPtPQTX9Kt0zC7l0ajQpKY6Nyul5T+HvqF7a+4GaD1glXI2XUYv/WKo5Aj0KGzDhqBtXgFhSVtAEKpk9/+UvWoHMtEjRFQk+2WPdPZsqGVYo8Aw0ZLBCdQIDAQABAoIBAQDXiBXJqZDk0tC4"
      file_type = "BIN"
    }
    keks {
      content   = "MIIEowIBAAKCAQEA8ybZA+UQZvwK4kX1PdBkYYPGEx/SvHpCU2FaRHHs7OnmqVHI+KnANgPPpqQYHmtTwdZuGZPk0iU5/JkPJJ1PnNBXA9RlGrKVHXa+OVlgJPiJP5qDgqK5n2E9IF5kJHhylG5w4RXXspRWBHJELGALQz/Bu2qZCeiKh7FslZxHEJTs4ynDLJUxZBhRHBKKog/mL2nEcX1hx4LgCXzVJvzgxI+uDxHM+GMeNPtPQTX9Kt0zC7l0ajQpKY6Nyul5T+HvqF7a+4GaD1glXI2XUYv/WKo5Aj0KGzDhqBtXgFhSVtAEKpk9/+UvWoHMtEjRFQk+2WPdPZsqGVYo8Aw0ZLBCdQIDAQABAoIBAQDXiBXJqZDk0tC4"
      file_type = "BIN"
    }
  }
}
`, imageName)
}

func testAccComputeImage_sourceImage(imageName string) string {
	return fmt.Sprintf(`
data "google_compute_image" "my_image" {
  family  = "debian-11"
  project = "debian-cloud"
}

resource "google_compute_image" "foobar" {
  name         = "%s"
  source_image = data.google_compute_image.my_image.self_link
}
`, imageName)
}

func testAccComputeImage_sourceSnapshot(diskName, snapshotName, imageName string) string {
	return fmt.Sprintf(`
data "google_compute_image" "my_image" {
  family  = "debian-11"
  project = "debian-cloud"
}

resource "google_compute_disk" "foobar" {
  name  = "%s"
  image = data.google_compute_image.my_image.self_link
  size  = 10
  type  = "pd-ssd"
  zone  = "us-central1-a"
}

resource "google_compute_snapshot" "foobar" {
  name        = "%s"
  source_disk = google_compute_disk.foobar.name
  zone        = "us-central1-a"
}

resource "google_compute_image" "foobar" {
  name            = "%s"
  source_snapshot = google_compute_snapshot.foobar.self_link
}
`, diskName, snapshotName, imageName)
}

func testAccComputeImage_sourceDiskEncryptionKey(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_compute_image" "debian" {
  family  = "debian-11"
  project = "debian-cloud"
}

resource "google_compute_disk" "src-disk-kms" {
  name  = "tf-test-src-disk-kms-%{random_suffix}"
  image = data.google_compute_image.debian.self_link
  size  = 10
  type  = "pd-ssd"
  zone  = "us-central1-a"

  disk_encryption_key {
	kms_key_self_link = "%{kms_key_self_link}"
  }
}

resource "google_compute_disk" "src-disk-raw" {
  name  = "tf-test-src-disk-raw-%{random_suffix}"
  image = data.google_compute_image.debian.self_link
  size  = 10
  type  = "pd-ssd"
  zone  = "us-central1-a"

  disk_encryption_key {
	raw_key = "%{raw_key}"
  }
}

resource "google_compute_disk" "src-disk-rsa" {
  name  = "tf-test-src-disk-rsa-%{random_suffix}"
  image = data.google_compute_image.debian.self_link
  size  = 10
  type  = "pd-ssd"
  zone  = "us-central1-a"

  disk_encryption_key {
	rsa_encrypted_key = "%{rsa_encrypted_key}"
  }
}

resource "google_compute_image" "foobar-kms" {
	name         = "tf-test-image-kms-%{random_suffix}"
	source_disk = google_compute_disk.src-disk-kms.self_link
	source_disk_encryption_key {
		kms_key_self_link = "%{kms_key_self_link}"
	}
}

resource "google_compute_image" "foobar-raw" {
	name         = "tf-test-image-raw-%{random_suffix}"
	source_disk = google_compute_disk.src-disk-raw.self_link
	source_disk_encryption_key {
		raw_key = "%{raw_key}"
	}
}

resource "google_compute_image" "foobar-rsa" {
	name         = "tf-test-image-rsa-%{random_suffix}"
	source_disk = google_compute_disk.src-disk-rsa.self_link
	source_disk_encryption_key {
		rsa_encrypted_key = "%{rsa_encrypted_key}"
		kms_key_service_account = data.google_compute_default_service_account.default.email
	}
}

data "google_compute_default_service_account" "default" {
}
`, context)
}

func testAccComputeImage_sourceImageEncryptionKey(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_compute_image" "debian" {
  family  = "debian-11"
  project = "debian-cloud"
}

resource "google_compute_disk" "src_disk" {
  name  = "tf-test-src-disk-%{random_suffix}"
  image = data.google_compute_image.debian.self_link
  size  = 10
  type  = "pd-ssd"
  zone  = "us-central1-a"

  disk_encryption_key {
	kms_key_self_link = "%{kms_key_self_link}"
  }
}

resource "google_compute_image" "src-image-kms" {
  name         = "tf-test-src-kms-%{random_suffix}"
  source_disk  = google_compute_disk.src_disk.self_link
  image_encryption_key {
	kms_key_self_link = "%{kms_key_self_link}"
  }
}

resource "google_compute_image" "src-image-raw" {
  name         = "tf-test-src-raw-%{random_suffix}"
  source_disk  = google_compute_disk.src_disk.self_link
  image_encryption_key {
	raw_key = "%{raw_key}"
  }
}

resource "google_compute_image" "src-image-rsa" {
  name         = "tf-test-src-rsa-%{random_suffix}"
  source_disk  = google_compute_disk.src_disk.self_link
  image_encryption_key {
	rsa_encrypted_key = "%{rsa_encrypted_key}"
  }
}

resource "google_compute_image" "foobar-kms" {
	name         = "tf-test-image-kms-%{random_suffix}"
	source_image = google_compute_image.src-image-kms.self_link
	source_image_encryption_key {
		kms_key_self_link = "%{kms_key_self_link}"
	}
}

resource "google_compute_image" "foobar-raw" {
	name         = "tf-test-image-raw-%{random_suffix}"
	source_image = google_compute_image.src-image-raw.self_link
	source_image_encryption_key {
		raw_key = "%{raw_key}"
	}
}

resource "google_compute_image" "foobar-rsa" {
	name         = "tf-test-image-rsa-%{random_suffix}"
	source_image = google_compute_image.src-image-rsa.self_link
	source_image_encryption_key {
		rsa_encrypted_key = "%{rsa_encrypted_key}"
		kms_key_service_account = data.google_compute_default_service_account.default.email
	}
}

data "google_compute_default_service_account" "default" {
}
`, context)
}

func testAccComputeImage_sourceSnapshotEncryptionKey(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_compute_image" "debian" {
  family  = "debian-11"
  project = "debian-cloud"
}

resource "google_compute_disk" "src_disk-kms" {
  name  = "tf-test-src-disk-kms-%{random_suffix}"
  image = data.google_compute_image.debian.self_link
  size  = 10
  type  = "pd-ssd"
  zone  = "us-central1-a"

  disk_encryption_key {
	kms_key_self_link = "%{kms_key_self_link}"
  }
}

resource "google_compute_disk" "src_disk-raw" {
  name  = "tf-test-src-disk-raw-%{random_suffix}"
  image = data.google_compute_image.debian.self_link
  size  = 10
  type  = "pd-ssd"
  zone  = "us-central1-a"

  disk_encryption_key {
	raw_key = "%{raw_key}"
  }
}

resource "google_compute_disk" "src_disk-rsa" {
  name  = "tf-test-src-disk-rsa-%{random_suffix}"
  image = data.google_compute_image.debian.self_link
  size  = 10
  type  = "pd-ssd"
  zone  = "us-central1-a"

  disk_encryption_key {
	rsa_encrypted_key = "%{rsa_encrypted_key}"
  }
}

resource "google_compute_snapshot" "src-snapshot-kms" {
  name  = "tf-test-src-snapshot-kms-%{random_suffix}"
  source_disk  = google_compute_disk.src_disk-kms.self_link
  zone  = "us-central1-a"

  snapshot_encryption_key {
	kms_key_self_link = "%{kms_key_self_link}"
  }
}

resource "google_compute_snapshot" "src-snapshot-raw" {
  name  = "tf-test-src-snapshot-raw-%{random_suffix}"
  source_disk  = google_compute_disk.src_disk-raw.self_link
  zone  = "us-central1-a"

  snapshot_encryption_key {
	raw_key = "%{raw_key}"
  }

  source_disk_encryption_key {
	raw_key = "%{raw_key}"
  }
}

resource "google_compute_snapshot" "src-snapshot-rsa" {
  name  = "tf-test-src-snapshot-rsa-%{random_suffix}"
  source_disk  = google_compute_disk.src_disk-rsa.self_link
  zone  = "us-central1-a"

  snapshot_encryption_key {
	rsa_encrypted_key = "%{rsa_encrypted_key}"
  }

  source_disk_encryption_key {
	rsa_encrypted_key = "%{rsa_encrypted_key}"
  }
}

resource "google_compute_image" "foobar-kms" {
	name         = "tf-test-image-kms-%{random_suffix}"
	source_snapshot = google_compute_snapshot.src-snapshot-kms.self_link
	source_snapshot_encryption_key {
		kms_key_self_link = "%{kms_key_self_link}"
	}
}

resource "google_compute_image" "foobar-raw" {
	name         = "tf-test-image-raw-%{random_suffix}"
	source_snapshot = google_compute_snapshot.src-snapshot-raw.self_link
	source_snapshot_encryption_key {
		raw_key = "%{raw_key}"
	}
}

resource "google_compute_image" "foobar-rsa" {
	name         = "tf-test-image-rsa-%{random_suffix}"
	source_snapshot = google_compute_snapshot.src-snapshot-rsa.self_link
	source_snapshot_encryption_key {
		rsa_encrypted_key = "%{rsa_encrypted_key}"
		kms_key_service_account = data.google_compute_default_service_account.default.email
	}
}

data "google_compute_default_service_account" "default" {
}
`, context)
}

func testAccComputeImage_imageEncryptionKey(kmsRingName, kmsKeyName, suffix string) string {
	return fmt.Sprintf(`
data "google_kms_key_ring" "ring" {
  name     = "%s"
  location = "us-central1"
}

data "google_kms_crypto_key" "key" {
  name     = "%s"
  key_ring = data.google_kms_key_ring.ring.id
}

resource "google_service_account" "test" {
  account_id   = "tf-test-sa-%s"
  display_name = "KMS Ops Account"
}

resource "google_kms_crypto_key_iam_member" "crypto_key" {
  crypto_key_id = data.google_kms_crypto_key.key.id
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member        = "serviceAccount:${google_service_account.test.email}"
}

data "google_compute_image" "debian" {
  family  = "debian-11"
  project = "debian-cloud"
}

resource "google_compute_image" "image" {
  name         = "tf-test-image-%s"
  source_image = data.google_compute_image.debian.self_link
  image_encryption_key {
    kms_key_self_link       = data.google_kms_crypto_key.key.id
    kms_key_service_account = google_service_account.test.email
  }
  depends_on = [
    google_kms_crypto_key_iam_member.crypto_key
  ]
}
`, kmsRingName, kmsKeyName, suffix, suffix)
}
