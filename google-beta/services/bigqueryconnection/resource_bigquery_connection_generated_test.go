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

package bigqueryconnection_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccBigqueryConnectionConnection_bigqueryConnectionCloudResourceExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		CheckDestroy: testAccCheckBigqueryConnectionConnectionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigqueryConnectionConnection_bigqueryConnectionCloudResourceExample(context),
			},
			{
				ResourceName:            "google_bigquery_connection.connection",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location"},
			},
		},
	})
}

func testAccBigqueryConnectionConnection_bigqueryConnectionCloudResourceExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_bigquery_connection" "connection" {
   connection_id = "tf-test-my-connection%{random_suffix}"
   location      = "US"
   friendly_name = "👋"
   description   = "a riveting description"
   cloud_resource {}
}
`, context)
}

func TestAccBigqueryConnectionConnection_bigqueryConnectionBasicExample(t *testing.T) {
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		CheckDestroy: testAccCheckBigqueryConnectionConnectionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigqueryConnectionConnection_bigqueryConnectionBasicExample(context),
			},
			{
				ResourceName:            "google_bigquery_connection.connection",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"cloud_sql.0.credential", "location"},
			},
		},
	})
}

func testAccBigqueryConnectionConnection_bigqueryConnectionBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_sql_database_instance" "instance" {
    name             = "tf-test-my-database-instance%{random_suffix}"
    database_version = "POSTGRES_11"
    region           = "us-central1"
    settings {
		tier = "db-f1-micro"
	}

    deletion_protection  = %{deletion_protection}
}

resource "google_sql_database" "db" {
    instance = google_sql_database_instance.instance.name
    name     = "db"
}

resource "random_password" "pwd" {
    length = 16
    special = false
}

resource "google_sql_user" "user" {
    name = "user%{random_suffix}"
    instance = google_sql_database_instance.instance.name
    password = random_password.pwd.result
}

resource "google_bigquery_connection" "connection" {
    friendly_name = "👋"
    description   = "a riveting description"
    location      = "US"
    cloud_sql {
        instance_id = google_sql_database_instance.instance.connection_name
        database    = google_sql_database.db.name
        type        = "POSTGRES"
        credential {
          username = google_sql_user.user.name
          password = google_sql_user.user.password
        }
    }
}
`, context)
}

func TestAccBigqueryConnectionConnection_bigqueryConnectionFullExample(t *testing.T) {
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		CheckDestroy: testAccCheckBigqueryConnectionConnectionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigqueryConnectionConnection_bigqueryConnectionFullExample(context),
			},
			{
				ResourceName:            "google_bigquery_connection.connection",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"cloud_sql.0.credential", "location"},
			},
		},
	})
}

func testAccBigqueryConnectionConnection_bigqueryConnectionFullExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_sql_database_instance" "instance" {
    name             = "tf-test-my-database-instance%{random_suffix}"
    database_version = "POSTGRES_11"
    region           = "us-central1"
    settings {
		tier = "db-f1-micro"
	}

    deletion_protection  = %{deletion_protection}
}

resource "google_sql_database" "db" {
    instance = google_sql_database_instance.instance.name
    name     = "db"
}

resource "random_password" "pwd" {
    length = 16
    special = false
}

resource "google_sql_user" "user" {
    name = "user%{random_suffix}"
    instance = google_sql_database_instance.instance.name
    password = random_password.pwd.result
}

resource "google_bigquery_connection" "connection" {
    connection_id = "tf-test-my-connection%{random_suffix}"
    location      = "US"
    friendly_name = "👋"
    description   = "a riveting description"
    cloud_sql {
        instance_id = google_sql_database_instance.instance.connection_name
        database    = google_sql_database.db.name
        type        = "POSTGRES"
        credential {
          username = google_sql_user.user.name
          password = google_sql_user.user.password
        }
    }
}
`, context)
}

func TestAccBigqueryConnectionConnection_bigqueryConnectionAwsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		CheckDestroy: testAccCheckBigqueryConnectionConnectionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigqueryConnectionConnection_bigqueryConnectionAwsExample(context),
			},
			{
				ResourceName:            "google_bigquery_connection.connection",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location"},
			},
		},
	})
}

func testAccBigqueryConnectionConnection_bigqueryConnectionAwsExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_bigquery_connection" "connection" {
   connection_id = "tf-test-my-connection%{random_suffix}"
   location      = "aws-us-east-1"
   friendly_name = "👋"
   description   = "a riveting description"
   aws { 
      access_role {
         iam_role_id =  "arn:aws:iam::999999999999:role/omnirole%{random_suffix}"
      }
   }
}
`, context)
}

func TestAccBigqueryConnectionConnection_bigqueryConnectionAzureExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		CheckDestroy: testAccCheckBigqueryConnectionConnectionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigqueryConnectionConnection_bigqueryConnectionAzureExample(context),
			},
			{
				ResourceName:            "google_bigquery_connection.connection",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location"},
			},
		},
	})
}

func testAccBigqueryConnectionConnection_bigqueryConnectionAzureExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_bigquery_connection" "connection" {
   connection_id = "tf-test-my-connection%{random_suffix}"
   location      = "azure-eastus2"
   friendly_name = "👋"
   description   = "a riveting description"
   azure {
      customer_tenant_id = "tf-test-customer-tenant-id%{random_suffix}"
      federated_application_client_id = "tf-test-b43eeeee-eeee-eeee-eeee-a480155501ce%{random_suffix}"
   }
}
`, context)
}

func TestAccBigqueryConnectionConnection_bigqueryConnectionCloudspannerExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		CheckDestroy: testAccCheckBigqueryConnectionConnectionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigqueryConnectionConnection_bigqueryConnectionCloudspannerExample(context),
			},
			{
				ResourceName:            "google_bigquery_connection.connection",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location"},
			},
		},
	})
}

func testAccBigqueryConnectionConnection_bigqueryConnectionCloudspannerExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_bigquery_connection" "connection" {
   connection_id = "tf-test-my-connection%{random_suffix}"
   location      = "US"
   friendly_name = "👋"
   description   = "a riveting description"
   cloud_spanner { 
      database = "projects/project/instances/instance/databases/database%{random_suffix}"
      database_role = "tf_test_database_role%{random_suffix}"
   }
}
`, context)
}

func TestAccBigqueryConnectionConnection_bigqueryConnectionCloudspannerDataboostExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		CheckDestroy: testAccCheckBigqueryConnectionConnectionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigqueryConnectionConnection_bigqueryConnectionCloudspannerDataboostExample(context),
			},
			{
				ResourceName:            "google_bigquery_connection.connection",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location"},
			},
		},
	})
}

func testAccBigqueryConnectionConnection_bigqueryConnectionCloudspannerDataboostExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_bigquery_connection" "connection" {
   connection_id = "tf-test-my-connection%{random_suffix}"
   location      = "US"
   friendly_name = "👋"
   description   = "a riveting description"
   cloud_spanner { 
      database        = "projects/project/instances/instance/databases/database%{random_suffix}"
      use_parallelism = true
      use_data_boost  = true
      max_parallelism = 100
   }
}
`, context)
}

func TestAccBigqueryConnectionConnection_bigqueryConnectionSparkExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckBigqueryConnectionConnectionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigqueryConnectionConnection_bigqueryConnectionSparkExample(context),
			},
			{
				ResourceName:            "google_bigquery_connection.connection",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location"},
			},
		},
	})
}

func testAccBigqueryConnectionConnection_bigqueryConnectionSparkExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_bigquery_connection" "connection" {
   connection_id = "tf-test-my-connection%{random_suffix}"
   location      = "US"
   friendly_name = "👋"
   description   = "a riveting description"
   spark {
      spark_history_server_config {
         dataproc_cluster = google_dataproc_cluster.basic.id
      }
   }
}

resource "google_dataproc_cluster" "basic" {
   name   = "tf-test-my-connection%{random_suffix}"
   region = "us-central1"

   cluster_config {
     # Keep the costs down with smallest config we can get away with
     software_config {
       override_properties = {
         "dataproc:dataproc.allow.zero.workers" = "true"
       }
     }
 
     master_config {
       num_instances = 1
       machine_type  = "e2-standard-2"
       disk_config {
         boot_disk_size_gb = 35
       }
     }
   }   
 }
`, context)
}

func TestAccBigqueryConnectionConnection_bigqueryConnectionSqlWithCmekExample(t *testing.T) {
	t.Parallel()
	acctest.BootstrapIamMembers(t, []acctest.IamMember{
		{
			Member: "serviceAccount:bq-{project_number}@bigquery-encryption.iam.gserviceaccount.com",
			Role:   "roles/cloudkms.cryptoKeyEncrypterDecrypter",
		},
	})

	context := map[string]interface{}{
		"deletion_protection": false,
		"kms_key_name":        acctest.BootstrapKMSKey(t).CryptoKey.Name,
		"random_suffix":       acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckBigqueryConnectionConnectionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigqueryConnectionConnection_bigqueryConnectionSqlWithCmekExample(context),
			},
			{
				ResourceName:            "google_bigquery_connection.bq-connection-cmek",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"cloud_sql.0.credential", "location"},
			},
		},
	})
}

func testAccBigqueryConnectionConnection_bigqueryConnectionSqlWithCmekExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_sql_database_instance" "instance" {
  name             = "tf-test-my-database-instance%{random_suffix}"
  region           = "us-central1"

  database_version = "POSTGRES_11"
  settings {
    tier = "db-f1-micro"
  }

  deletion_protection  = %{deletion_protection}
}

resource "google_sql_database" "db" {
  instance = google_sql_database_instance.instance.name
  name     = "db"
}

resource "google_sql_user" "user" {
  name = "user%{random_suffix}"
  instance = google_sql_database_instance.instance.name
  password = "tf-test-my-password%{random_suffix}"
}

resource "google_bigquery_connection" "bq-connection-cmek" {
  friendly_name = "👋"
  description   = "a riveting description"
  location      = "US"
  kms_key_name  = "%{kms_key_name}"
  cloud_sql {
    instance_id = google_sql_database_instance.instance.connection_name
    database    = google_sql_database.db.name
    type        = "POSTGRES"
    credential {
      username = google_sql_user.user.name
      password = google_sql_user.user.password
    }
  }
}
`, context)
}

func testAccCheckBigqueryConnectionConnectionDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_bigquery_connection" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{BigqueryConnectionBasePath}}projects/{{project}}/locations/{{location}}/connections/{{connection_id}}")
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
				return fmt.Errorf("BigqueryConnectionConnection still exists at %s", url)
			}
		}

		return nil
	}
}
