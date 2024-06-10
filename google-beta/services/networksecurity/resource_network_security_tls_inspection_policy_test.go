// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package networksecurity_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccNetworkSecurityTlsInspectionPolicy_update(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"projectNumber": envvar.GetTestProjectNumberFromEnv(),
		"randomSuffix":  acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckNetworkSecurityTlsInspectionPolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecurityTlsInspectionPolicy_basic(context),
			},
			{
				ResourceName:      "google_network_security_tls_inspection_policy.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccNetworkSecurityTlsInspectionPolicy_update(context),
			},
			{
				ResourceName:      "google_network_security_tls_inspection_policy.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccNetworkSecurityTlsInspectionPolicy_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_privateca_ca_pool" "default" {
  provider = google-beta
  name     = "tf-test-cap-%{randomSuffix}"
  location = "us-central1"
  tier     = "DEVOPS"

  publishing_options {
    publish_ca_cert = false
    publish_crl     = false
  }

  issuance_policy {
    maximum_lifetime = "1209600s"
    baseline_values {
      ca_options {
        is_ca = false
      }
      key_usage {
        base_key_usage {}
        extended_key_usage {
          server_auth = true
        }
      }
    }
  }
}

resource "google_privateca_certificate_authority" "default" {
  provider                               = google-beta
  pool                                   = google_privateca_ca_pool.default.name
  certificate_authority_id               = "tf-test-ca-%{randomSuffix}"
  location                               = "us-central1"
  lifetime                               = "86400s"
  type                                   = "SELF_SIGNED"
  deletion_protection                    = false
  skip_grace_period                      = true
  ignore_active_certificates_on_deletion = true

  config {
    subject_config {
      subject {
        organization = "Test LLC"
        common_name  = "my-ca"
      }
    }
    x509_config {
      ca_options {
        is_ca = true
      }
      key_usage {
        base_key_usage {
          cert_sign = true
          crl_sign  = true
        }
        extended_key_usage {
          server_auth = false
        }
      }
    }
  }

  key_spec {
    algorithm = "RSA_PKCS1_4096_SHA256"
  }
}

resource "google_project_service_identity" "default" {
  provider = google-beta
  service  = "networksecurity.googleapis.com"
}

resource "google_privateca_ca_pool_iam_member" "default" {
  provider = google-beta
  ca_pool  = google_privateca_ca_pool.default.id
  role     = "roles/privateca.certificateManager"
  member   = "serviceAccount:${google_project_service_identity.default.email}"
}

resource "google_certificate_manager_trust_config" "default" {
  provider    = google-beta
  name        = "tf-test-tc-%{randomSuffix}"
  description = "sample trust config description"
  location    = "us-central1"

  trust_stores {
    trust_anchors {
      pem_certificate = file("test-fixtures/ca_cert.pem")
    }
    intermediate_cas {
      pem_certificate = file("test-fixtures/ca_cert.pem")
    }
  }
}

resource "google_network_security_tls_inspection_policy" "default" {
  provider              = google-beta
  name                  = "tf-test-tip-%{randomSuffix}"
  location              = "us-central1"
  ca_pool               = google_privateca_ca_pool.default.id
  exclude_public_ca_set = false
  min_tls_version       = "TLS_1_0"
  trust_config          = google_certificate_manager_trust_config.default.id
  tls_feature_profile   = "PROFILE_CUSTOM"

  custom_tls_features = [
    "TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA",
    "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256",
    "TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA",
    "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
    "TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256",
    "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA",
    "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
    "TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA",
    "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
    "TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256",
    "TLS_RSA_WITH_3DES_EDE_CBC_SHA",
    "TLS_RSA_WITH_AES_128_CBC_SHA",
    "TLS_RSA_WITH_AES_128_GCM_SHA256",
    "TLS_RSA_WITH_AES_256_CBC_SHA",
    "TLS_RSA_WITH_AES_256_GCM_SHA384"
  ]

  depends_on = [
    google_privateca_certificate_authority.default,
    google_privateca_ca_pool_iam_member.default
  ]
}
`, context)
}

func testAccNetworkSecurityTlsInspectionPolicy_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_privateca_ca_pool" "default" {
  provider = google-beta
  name     = "tf-test-cap-%{randomSuffix}"
  location = "us-central1"
  tier     = "DEVOPS"

  publishing_options {
    publish_ca_cert = false
    publish_crl     = false
  }

  issuance_policy {
    maximum_lifetime = "1209600s"
    baseline_values {
      ca_options {
        is_ca = false
      }
      key_usage {
        base_key_usage {}
        extended_key_usage {
          server_auth = true
        }
      }
    }
  }
}

resource "google_privateca_ca_pool" "default_updated" {
  provider = google-beta
  name     = "tf-test-cap-updated-%{randomSuffix}"
  location = "us-central1"
  tier     = "DEVOPS"

  publishing_options {
    publish_ca_cert = false
    publish_crl     = false
  }

  issuance_policy {
    maximum_lifetime = "1209600s"
    baseline_values {
      ca_options {
        is_ca = false
      }
      key_usage {
        base_key_usage {}
        extended_key_usage {
          server_auth = true
        }
      }
    }
  }
}

resource "google_privateca_certificate_authority" "default" {
  provider                               = google-beta
  pool                                   = google_privateca_ca_pool.default.name
  certificate_authority_id               = "tf-test-ca-%{randomSuffix}"
  location                               = "us-central1"
  lifetime                               = "86400s"
  type                                   = "SELF_SIGNED"
  deletion_protection                    = false
  skip_grace_period                      = true
  ignore_active_certificates_on_deletion = true

  config {
    subject_config {
      subject {
        organization = "Test LLC"
        common_name  = "my-ca"
      }
    }
    x509_config {
      ca_options {
        is_ca = true
      }
      key_usage {
        base_key_usage {
          cert_sign = true
          crl_sign  = true
        }
        extended_key_usage {
          server_auth = false
        }
      }
    }
  }

  key_spec {
    algorithm = "RSA_PKCS1_4096_SHA256"
  }
}

resource "google_privateca_certificate_authority" "default_updated" {
  provider                               = google-beta
  pool                                   = google_privateca_ca_pool.default_updated.name
  certificate_authority_id               = "tf-test-ca-%{randomSuffix}"
  location                               = "us-central1"
  lifetime                               = "86400s"
  type                                   = "SELF_SIGNED"
  deletion_protection                    = false
  skip_grace_period                      = true
  ignore_active_certificates_on_deletion = true

  config {
    subject_config {
      subject {
        organization = "Test LLC"
        common_name  = "my-ca"
      }
    }
    x509_config {
      ca_options {
        is_ca = true
      }
      key_usage {
        base_key_usage {
          cert_sign = true
          crl_sign  = true
        }
        extended_key_usage {
          server_auth = false
        }
      }
    }
  }

  key_spec {
    algorithm = "RSA_PKCS1_4096_SHA256"
  }
}

resource "google_project_service_identity" "default" {
  provider = google-beta
  service  = "networksecurity.googleapis.com"
}

resource "google_privateca_ca_pool_iam_member" "default" {
  provider = google-beta
  ca_pool  = google_privateca_ca_pool.default.id
  role     = "roles/privateca.certificateManager"
  member   = "serviceAccount:${google_project_service_identity.default.email}"
}

resource "google_privateca_ca_pool_iam_member" "default_updated" {
  provider = google-beta
  ca_pool  = google_privateca_ca_pool.default_updated.id
  role     = "roles/privateca.certificateManager"
  member   = "serviceAccount:${google_project_service_identity.default.email}"
}

resource "google_certificate_manager_trust_config" "default" {
  provider    = google-beta
  name        = "tf-test-tc-%{randomSuffix}"
  description = "sample trust config description"
  location    = "us-central1"

  trust_stores {
    trust_anchors {
      pem_certificate = file("test-fixtures/ca_cert.pem")
    }
    intermediate_cas {
      pem_certificate = file("test-fixtures/ca_cert.pem")
    }
  }
}

resource "google_certificate_manager_trust_config" "default_updated" {
  provider    = google-beta
  name        = "tf-test-tc-updated-%{randomSuffix}"
  description = "another sample trust config description"
  location    = "us-central1"

  trust_stores {
    trust_anchors {
      pem_certificate = file("test-fixtures/ca_cert.pem")
    }
    intermediate_cas {
      pem_certificate = file("test-fixtures/ca_cert.pem")
    }
  }
}

resource "google_network_security_tls_inspection_policy" "default" {
  provider              = google-beta
  name                  = "tf-test-tip-%{randomSuffix}"
  location              = "us-central1"
  description           = "my tls inspection policy updated"
  ca_pool               = google_privateca_ca_pool.default_updated.id
  exclude_public_ca_set = true
  min_tls_version       = "TLS_1_2"
  trust_config          = google_certificate_manager_trust_config.default_updated.id

  depends_on  = [
    google_privateca_certificate_authority.default_updated,
    google_privateca_ca_pool_iam_member.default_updated
  ]
}
`, context)
}
