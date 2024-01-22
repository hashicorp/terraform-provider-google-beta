// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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

package artifactregistry_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccArtifactRegistryRepository_artifactRegistryRepositoryBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckArtifactRegistryRepositoryDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccArtifactRegistryRepository_artifactRegistryRepositoryBasicExample(context),
			},
			{
				ResourceName:            "google_artifact_registry_repository.my-repo",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"repository_id", "location", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccArtifactRegistryRepository_artifactRegistryRepositoryBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_artifact_registry_repository" "my-repo" {
  location      = "us-central1"
  repository_id = "tf-test-my-repository%{random_suffix}"
  description   = "example docker repository%{random_suffix}"
  format        = "DOCKER"
}
`, context)
}

func TestAccArtifactRegistryRepository_artifactRegistryRepositoryDockerExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckArtifactRegistryRepositoryDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccArtifactRegistryRepository_artifactRegistryRepositoryDockerExample(context),
			},
			{
				ResourceName:            "google_artifact_registry_repository.my-repo",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"repository_id", "location", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccArtifactRegistryRepository_artifactRegistryRepositoryDockerExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_artifact_registry_repository" "my-repo" {
  location      = "us-central1"
  repository_id = "tf-test-my-repository%{random_suffix}"
  description   = "example docker repository%{random_suffix}"
  format        = "DOCKER"

  docker_config {
    immutable_tags = true
  }
}
`, context)
}

func TestAccArtifactRegistryRepository_artifactRegistryRepositoryCmekExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"kms_key_name":  acctest.BootstrapKMSKeyInLocation(t, "us-central1").CryptoKey.Name,
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckArtifactRegistryRepositoryDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccArtifactRegistryRepository_artifactRegistryRepositoryCmekExample(context),
			},
			{
				ResourceName:            "google_artifact_registry_repository.my-repo",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"repository_id", "location", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccArtifactRegistryRepository_artifactRegistryRepositoryCmekExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_artifact_registry_repository" "my-repo" {
  location      = "us-central1"
  repository_id = "tf-test-my-repository%{random_suffix}"
  description   = "example docker repository with cmek"
  format        = "DOCKER"
  kms_key_name  = "%{kms_key_name}"
  depends_on = [
    google_kms_crypto_key_iam_member.crypto_key
  ]
}

resource "google_kms_crypto_key_iam_member" "crypto_key" {
  crypto_key_id = "%{kms_key_name}"
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member        = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-artifactregistry.iam.gserviceaccount.com"
}

data "google_project" "project" {}
`, context)
}

func TestAccArtifactRegistryRepository_artifactRegistryRepositoryVirtualExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckArtifactRegistryRepositoryDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccArtifactRegistryRepository_artifactRegistryRepositoryVirtualExample(context),
			},
			{
				ResourceName:            "google_artifact_registry_repository.my-repo",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"repository_id", "location", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccArtifactRegistryRepository_artifactRegistryRepositoryVirtualExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_artifact_registry_repository" "my-repo-upstream-1" {
  location      = "us-central1"
  repository_id = "tf-test-my-repository-upstream%{random_suffix}-1"
  description   = "example docker repository (upstream source)%{random_suffix} 1"
  format        = "DOCKER"
}

resource "google_artifact_registry_repository" "my-repo-upstream-2" {
  location      = "us-central1"
  repository_id = "tf-test-my-repository-upstream%{random_suffix}-2"
  description   = "example docker repository (upstream source)%{random_suffix} 2"
  format        = "DOCKER"
}

resource "google_artifact_registry_repository" "my-repo" {
  depends_on    = []
  location      = "us-central1"
  repository_id = "tf-test-my-repository%{random_suffix}"
  description   = "example virtual docker repository%{random_suffix}"
  format        = "DOCKER"
  mode          = "VIRTUAL_REPOSITORY"
  virtual_repository_config {
    upstream_policies {
      id          = "tf-test-my-repository-upstream%{random_suffix}-1"
      repository  = google_artifact_registry_repository.my-repo-upstream-1.id
      priority    = 20
    }
    upstream_policies {
      id          = "tf-test-my-repository-upstream%{random_suffix}-2"
      repository  = google_artifact_registry_repository.my-repo-upstream-2.id
      priority    = 10
    }
  }
}
`, context)
}

func TestAccArtifactRegistryRepository_artifactRegistryRepositoryRemoteExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckArtifactRegistryRepositoryDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccArtifactRegistryRepository_artifactRegistryRepositoryRemoteExample(context),
			},
			{
				ResourceName:            "google_artifact_registry_repository.my-repo",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"repository_id", "location", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccArtifactRegistryRepository_artifactRegistryRepositoryRemoteExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_artifact_registry_repository" "my-repo" {
  location      = "us-central1"
  repository_id = "tf-test-my-repository%{random_suffix}"
  description   = "example remote docker repository%{random_suffix}"
  format        = "DOCKER"
  mode          = "REMOTE_REPOSITORY"
  remote_repository_config {
    description = "docker hub"
    docker_repository {
      public_repository = "DOCKER_HUB"
    }
  }
}
`, context)
}

func TestAccArtifactRegistryRepository_artifactRegistryRepositoryRemoteAptExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckArtifactRegistryRepositoryDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccArtifactRegistryRepository_artifactRegistryRepositoryRemoteAptExample(context),
			},
			{
				ResourceName:            "google_artifact_registry_repository.my-repo",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"repository_id", "location", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccArtifactRegistryRepository_artifactRegistryRepositoryRemoteAptExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_artifact_registry_repository" "my-repo" {
  location      = "us-central1"
  repository_id = "tf-test-debian-buster%{random_suffix}"
  description   = "example remote apt repository%{random_suffix}"
  format        = "APT"
  mode          = "REMOTE_REPOSITORY"
  remote_repository_config {
    description = "Debian buster remote repository"
    apt_repository {
      public_repository {
        repository_base = "DEBIAN"
        repository_path = "debian/dists/buster"
      }
    }
  }
}
`, context)
}

func TestAccArtifactRegistryRepository_artifactRegistryRepositoryRemoteYumExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckArtifactRegistryRepositoryDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccArtifactRegistryRepository_artifactRegistryRepositoryRemoteYumExample(context),
			},
			{
				ResourceName:            "google_artifact_registry_repository.my-repo",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"repository_id", "location", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccArtifactRegistryRepository_artifactRegistryRepositoryRemoteYumExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_artifact_registry_repository" "my-repo" {
  location      = "us-central1"
  repository_id = "tf-test-centos-8%{random_suffix}"
  description   = "example remote yum repository%{random_suffix}"
  format        = "YUM"
  mode          = "REMOTE_REPOSITORY"
  remote_repository_config {
    description = "Centos 8 remote repository"
    yum_repository {
      public_repository {
        repository_base = "CENTOS"
        repository_path = "centos/8-stream/BaseOS/x86_64/os"
      }
    }
  }
}
`, context)
}

func TestAccArtifactRegistryRepository_artifactRegistryRepositoryCleanupExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckArtifactRegistryRepositoryDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccArtifactRegistryRepository_artifactRegistryRepositoryCleanupExample(context),
			},
			{
				ResourceName:            "google_artifact_registry_repository.my-repo",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"repository_id", "location", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccArtifactRegistryRepository_artifactRegistryRepositoryCleanupExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_artifact_registry_repository" "my-repo" {
  location      = "us-central1"
  repository_id = "tf-test-my-repository%{random_suffix}"
  description   = "example docker repository with cleanup policies%{random_suffix}"
  format        = "DOCKER"
  cleanup_policy_dry_run = false
  cleanup_policies {
    id     = "delete-prerelease"
    action = "DELETE"
    condition {
      tag_state    = "TAGGED"
      tag_prefixes = ["alpha", "v0"]
      older_than   = "2592000s"
    }
  }
  cleanup_policies {
    id     = "keep-tagged-release"
    action = "KEEP"
    condition {
      tag_state             = "TAGGED"
      tag_prefixes          = ["release"]
      package_name_prefixes = ["webapp", "mobile"]
    }
  }
  cleanup_policies {
    id     = "keep-minimum-versions"
    action = "KEEP"
    most_recent_versions {
      package_name_prefixes = ["webapp", "mobile", "sandbox"]
      keep_count            = 5
    }
  }
}
`, context)
}

func TestAccArtifactRegistryRepository_artifactRegistryRepositoryRemoteCustomExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckArtifactRegistryRepositoryDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccArtifactRegistryRepository_artifactRegistryRepositoryRemoteCustomExample(context),
			},
			{
				ResourceName:            "google_artifact_registry_repository.my-repo",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"repository_id", "location", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccArtifactRegistryRepository_artifactRegistryRepositoryRemoteCustomExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_project" "project" {}

resource "google_secret_manager_secret" "tf-test-example-custom-remote-secret%{random_suffix}" {
  secret_id = "tf-test-example-secret%{random_suffix}"
  replication {
    auto {}
  }
}

resource "google_secret_manager_secret_version" "tf-test-example-custom-remote-secret%{random_suffix}_version" {
  secret = google_secret_manager_secret.tf-test-example-custom-remote-secret%{random_suffix}.id
  secret_data = "tf-test-remote-password%{random_suffix}"
}

resource "google_secret_manager_secret_iam_member" "secret-access" {
  secret_id = google_secret_manager_secret.tf-test-example-custom-remote-secret%{random_suffix}.id
  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-artifactregistry.iam.gserviceaccount.com"
}

resource "google_artifact_registry_repository" "my-repo" {
  location      = "us-central1"
  repository_id = "tf-test-example-custom-remote%{random_suffix}"
  description   = "example remote docker repository with credentials%{random_suffix}"
  format        = "DOCKER"
  mode          = "REMOTE_REPOSITORY"
  remote_repository_config {
    description = "docker hub with custom credentials"
    docker_repository {
      public_repository = "DOCKER_HUB"
    }
    upstream_credentials {
      username_password_credentials {
        username = "tf-test-remote-username%{random_suffix}"
        password_secret_version = google_secret_manager_secret_version.tf-test-example-custom-remote-secret%{random_suffix}_version.name
      }
    }
  }
}
`, context)
}

func testAccCheckArtifactRegistryRepositoryDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_artifact_registry_repository" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ArtifactRegistryBasePath}}projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("ArtifactRegistryRepository still exists at %s", url)
			}
		}

		return nil
	}
}
