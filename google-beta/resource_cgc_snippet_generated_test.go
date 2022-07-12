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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCGCSnippet_flaskGoogleCloudQuickstartExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCGCSnippet_flaskGoogleCloudQuickstartExample(context),
			},
			{
				ResourceName:            "google_compute_instance.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata", "metadata_startup_script"},
			},
		},
	})
}

func testAccCGCSnippet_flaskGoogleCloudQuickstartExample(context map[string]interface{}) string {
	return Nprintf(`
# Create a single Compute Engine instance
resource "google_compute_instance" "default" {
  name         = "tf-test-flask-vm%{random_suffix}"
  machine_type = "f1-micro"
  zone         = "us-west1-a"
  tags         = ["ssh"]

  metadata = {
    enable-oslogin = "TRUE"
  }
  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  # Install Flask
  metadata_startup_script = "sudo apt-get update; sudo apt-get install -yq build-essential python-pip rsync; pip install flask"

  network_interface {
    network = "default"

    access_config {
      # Include this section to give the VM an external IP address
    }
  }
}

resource "google_compute_firewall" "ssh" {
  name = "tf-test-allow-ssh%{random_suffix}"
  allow {
    ports    = ["22"]
    protocol = "tcp"
  }
  direction     = "INGRESS"
  network       = "default"
  priority      = 1000
  source_ranges = ["0.0.0.0/0"]
  target_tags   = ["ssh"]
}


# [START vpc_flask_quickstart_5000_fw]
resource "google_compute_firewall" "flask" {
  name    = "tf-test-flask-app-firewall%{random_suffix}"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["5000"]
  }
  source_ranges = ["0.0.0.0/0"]
}
# [END vpc_flask_quickstart_5000_fw]

# Create new multi-region storage bucket in the US
# with versioning enabled

resource "google_storage_bucket" "default" {
  name          = "tf-test-bucket-tfstate%{random_suffix}"
  force_destroy = false
  location      = "US"
  storage_class = "STANDARD"
  versioning {
    enabled = true
  }
}
`, context)
}

func TestAccCGCSnippet_sqlSqlserverVmInstanceExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccCGCSnippet_sqlSqlserverVmInstanceExample(context),
			},
			{
				ResourceName:      "google_compute_instance.sqlserver_vm",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCGCSnippet_sqlSqlserverVmInstanceExample(context map[string]interface{}) string {
	return Nprintf(`
# VPC network
resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-vpc-network%{random_suffix}"
  auto_create_subnetworks = false
}

# Subnet
resource "google_compute_subnetwork" "default" {
  provider      = google-beta
  name          = "tf-test-vpc-subnet%{random_suffix}"
  ip_cidr_range = "10.0.1.0/24"
  region        = "europe-west1"
  network       = google_compute_network.default.id
}

resource "google_compute_instance" "sqlserver_vm" {
  provider = google-beta
  name     = "tf-test-sqlserver-vm%{random_suffix}"
  boot_disk {
    auto_delete = true
    device_name = "persistent-disk-0"
    initialize_params {
      image = "windows-sql-cloud/sql-std-2019-win-2022"
      size  = 50
      type  = "pd-balanced"
    }
    mode   = "READ_WRITE"
  }
  machine_type = "n1-standard-4"
  zone         = "europe-west1-b"
  network_interface {
    access_config {
      network_tier = "PREMIUM"
    }
    network            = google_compute_network.default.id
    stack_type         = "IPV4_ONLY"
    subnetwork         = google_compute_subnetwork.default.id
  }
}

resource "google_compute_firewall" "sql_server_1433" {
  provider      = google-beta
  name          = "tf-test-sql-server-1433-3%{random_suffix}"
  allow {
    ports    = ["1433"]
    protocol = "tcp"
  }
  description   = "Allow SQL Server access from all sources on port 1433."
  direction     = "INGRESS"
  network       = google_compute_network.default.id
  priority      = 1000
  source_ranges = ["0.0.0.0/0"]
}
`, context)
}

func TestAccCGCSnippet_spotInstanceBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCGCSnippet_spotInstanceBasicExample(context),
			},
			{
				ResourceName:      "google_compute_instance.spot_vm_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCGCSnippet_spotInstanceBasicExample(context map[string]interface{}) string {
	return Nprintf(`

resource "google_compute_instance" "spot_vm_instance" {
  name         = "tf-test-spot-instance-name%{random_suffix}"
  machine_type = "f1-micro"
  zone         = "us-central1-c"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }
  
  scheduling {
      preemptible = true
      automatic_restart = false
      provisioning_model = "SPOT"
  }

  network_interface {
    # A default network is created for all GCP projects
    network = "default"
    access_config {
    }
  }
}

`, context)
}

func TestAccCGCSnippet_instanceCustomHostnameExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCGCSnippet_instanceCustomHostnameExample(context),
			},
			{
				ResourceName:      "google_compute_instance.custom_hostname_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCGCSnippet_instanceCustomHostnameExample(context map[string]interface{}) string {
	return Nprintf(`

resource "google_compute_instance" "custom_hostname_instance" {
  name         = "tf-test-custom-hostname-instance-name%{random_suffix}"
  machine_type = "f1-micro"
  zone = "us-central1-c"

  # Set a custom hostname below 
  hostname = "hashicorptest.com"
  
  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }
  network_interface {
    # A default network is created for all GCP projects
    network = "default"
    access_config {
    }
  }
}

`, context)
}

func TestAccCGCSnippet_computeReservationExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       getTestProjectFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCGCSnippet_computeReservationExample(context),
			},
		},
	})
}

func testAccCGCSnippet_computeReservationExample(context map[string]interface{}) string {
	return Nprintf(`

resource "google_compute_reservation" "gce_reservation_local" {
  name = "tf-test-gce-reservation-local%{random_suffix}"
  zone = "us-central1-c"
  project = "%{project}"

  share_settings {
    share_type = "LOCAL"
  }

  specific_reservation {
    count = 1
    instance_properties {
      machine_type     = "n2-standard-2"
    }
  }
}

`, context)
}

func TestAccCGCSnippet_sqlDatabaseInstanceSqlserverExample(t *testing.T) {
	skipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCGCSnippet_sqlDatabaseInstanceSqlserverExample(context),
			},
			{
				ResourceName:            "google_sql_database_instance.instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"deletion_protection", "root_password"},
			},
		},
	})
}

func testAccCGCSnippet_sqlDatabaseInstanceSqlserverExample(context map[string]interface{}) string {
	return Nprintf(`
# [START cloud_sql_sqlserver_instance_80_db_n1_s2]
resource "google_sql_database_instance" "instance" {
  name             = "tf-test-sqlserver-instance%{random_suffix}"
  region           = "us-central1"
  database_version = "SQLSERVER_2019_STANDARD"
  root_password = "INSERT-PASSWORD-HERE"
  settings {
    tier = "db-custom-2-7680"
  }
  deletion_protection =  "%{deletion_protection}"
}
# [END cloud_sql_sqlserver_instance_80_db_n1_s2]

resource "random_password" "pwd" {
    length = 16
    special = false
}

resource "google_sql_user" "user" {
    name = "user"
    instance = google_sql_database_instance.instance.name
    password = random_password.pwd.result
}
`, context)
}

func TestAccCGCSnippet_sqlSqlserverInstanceCloneExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCGCSnippet_sqlSqlserverInstanceCloneExample(context),
			},
			{
				ResourceName:            "google_sql_database_instance.clone",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"deletion_protection", "root_password", "clone"},
			},
		},
	})
}

func testAccCGCSnippet_sqlSqlserverInstanceCloneExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_sql_database_instance" "source" {
  name             = "tf-test-sqlserver-instance-source-name%{random_suffix}"
  region           = "us-central1"
  database_version = "SQLSERVER_2017_STANDARD"
  root_password = "INSERT-PASSWORD-HERE"
  settings {
    tier = "db-custom-2-7680"
  }
  deletion_protection =  "%{deletion_protection}"
}

resource "google_sql_database_instance" "clone" {
  name             = "tf-test-sqlserver-instance-clone-name%{random_suffix}"
  region           = "us-central1"
  database_version = "SQLSERVER_2017_STANDARD"
  root_password = "INSERT-PASSWORD-HERE"
  clone {
    source_instance_name = google_sql_database_instance.source.id
  }
  deletion_protection =  "%{deletion_protection}"
}
`, context)
}

func TestAccCGCSnippet_sqlSqlserverInstanceBackupExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCGCSnippet_sqlSqlserverInstanceBackupExample(context),
			},
			{
				ResourceName:            "google_sql_database_instance.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"deletion_protection", "root_password"},
			},
		},
	})
}

func testAccCGCSnippet_sqlSqlserverInstanceBackupExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_sql_database_instance" "default" {
  name             = "tf-test-sqlserver-instance-backup%{random_suffix}"
  region           = "us-central1"
  database_version = "SQLSERVER_2019_STANDARD"
  root_password = "INSERT-PASSWORD-HERE"
  settings {
    tier = "db-custom-2-7680"
    backup_configuration {
      enabled                        = true
      start_time                     = "20:55"
    }
  }
  deletion_protection =  "%{deletion_protection}"
}
`, context)
}

func TestAccCGCSnippet_sqlSqlserverInstanceAuthorizedNetworkExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCGCSnippet_sqlSqlserverInstanceAuthorizedNetworkExample(context),
			},
			{
				ResourceName:            "google_sql_database_instance.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"deletion_protection", "root_password"},
			},
		},
	})
}

func testAccCGCSnippet_sqlSqlserverInstanceAuthorizedNetworkExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_sql_database_instance" "default" {
  name = "tf-test-sqlserver-instance-with-authorized-network%{random_suffix}"
  region           = "us-central1"
  database_version = "SQLSERVER_2019_STANDARD"
  root_password = "INSERT-PASSWORD-HERE"
  settings {
    tier = "db-custom-2-7680"
    ip_configuration {
      authorized_networks {
        name = "Network Name"
        value = "192.0.2.0/24"
        expiration_time = "3021-11-15T16:19:00.094Z"
      }
    }
  }
  deletion_protection = "%{deletion_protection}"
}
`, context)
}

func TestAccCGCSnippet_sqlSqlserverInstanceBackupLocationExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCGCSnippet_sqlSqlserverInstanceBackupLocationExample(context),
			},
			{
				ResourceName:            "google_sql_database_instance.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"root_password", "deletion_protection"},
			},
		},
	})
}

func testAccCGCSnippet_sqlSqlserverInstanceBackupLocationExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_sql_database_instance" "default" {
  name             = "tf-test-sqlserver-instance-with-backup-location%{random_suffix}"
  region           = "us-central1"
  database_version = "SQLSERVER_2019_STANDARD"
  root_password = "INSERT-PASSWORD-HERE"
  settings {
    tier = "db-custom-2-7680"
    backup_configuration {
      enabled                        = true
      location                       = "us-central1"
    }
  }
  deletion_protection =  "%{deletion_protection}"
}
`, context)
}

func TestAccCGCSnippet_sqlSqlserverInstanceBackupRetentionExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCGCSnippet_sqlSqlserverInstanceBackupRetentionExample(context),
			},
			{
				ResourceName:            "google_sql_database_instance.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"root_password", "deletion_protection"},
			},
		},
	})
}

func testAccCGCSnippet_sqlSqlserverInstanceBackupRetentionExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_sql_database_instance" "default" {
  name             = "tf-test-sqlserver-instance-backup-retention%{random_suffix}"
  region           = "us-central1"
  database_version = "SQLSERVER_2019_STANDARD"
  root_password = "INSERT-PASSWORD-HERE"
  settings {
    tier = "db-custom-2-7680"
    backup_configuration {
      enabled                        = true
      backup_retention_settings {
        retained_backups               = 365
        retention_unit                 = "COUNT"
      }
    }
  }
  deletion_protection =  "%{deletion_protection}"
}
`, context)
}

func TestAccCGCSnippet_sqlSqlserverInstanceReplicaExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCGCSnippet_sqlSqlserverInstanceReplicaExample(context),
			},
			{
				ResourceName:            "google_sql_database_instance.read_replica",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"deletion_protection", "root_password"},
			},
		},
	})
}

func testAccCGCSnippet_sqlSqlserverInstanceReplicaExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_sql_database_instance" "primary" {
  name             = "tf-test-sqlserver-primary-instance-name%{random_suffix}"
  region           = "europe-west4"
  database_version = "SQLSERVER_2019_ENTERPRISE"
  root_password    = "INSERT-PASSWORD-HERE"
  settings {
    tier = "db-custom-2-7680"
    backup_configuration {
      enabled = "true"
    }
  }
  deletion_protection = "%{deletion_protection}"
}

resource "google_sql_database_instance" "read_replica" {
  name                 = "tf-test-sqlserver-replica-instance-name%{random_suffix}"
  master_instance_name = google_sql_database_instance.primary.name
  region               = "europe-west4"
  database_version     = "SQLSERVER_2019_ENTERPRISE"
  root_password        = "INSERT-PASSWORD-HERE"
  replica_configuration {
    failover_target = false
  }

  settings {
    tier              = "db-custom-2-7680"
    availability_type = "ZONAL"
    disk_size         = "100"
  }
  deletion_protection = "%{deletion_protection}"
}
`, context)
}

func TestAccCGCSnippet_sqlSqlserverInstancePublicIpExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCGCSnippet_sqlSqlserverInstancePublicIpExample(context),
			},
			{
				ResourceName:            "google_sql_database_instance.sqlserver_public_ip_instance_name",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"deletion_protection", "root_password"},
			},
		},
	})
}

func testAccCGCSnippet_sqlSqlserverInstancePublicIpExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_sql_database_instance" "sqlserver_public_ip_instance_name" {
  name                 = "tf-test-sqlserver-public-ip-instance-name%{random_suffix}"
  region               = "europe-west4"
  database_version     = "SQLSERVER_2019_ENTERPRISE"
  root_password        = "INSERT-PASSWORD-HERE"
  settings {
    tier              = "db-custom-2-7680"
    availability_type = "ZONAL"
    ip_configuration {
      # Add optional authorized networks
      # Update to match the customer's networks
      authorized_networks {
        name  = "test-net-3"
        value = "203.0.113.0/24"
      }
      # Enable public IP
      ipv4_enabled = true
    }
  }
  deletion_protection =  "%{deletion_protection}"
}
`, context)
}

func TestAccCGCSnippet_sqlSqlserverInstancePrivateIpExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCGCSnippet_sqlSqlserverInstancePrivateIpExample(context),
			},
			{
				ResourceName:            "google_sql_database_instance.instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"root_password", "deletion_protection"},
			},
		},
	})
}

func testAccCGCSnippet_sqlSqlserverInstancePrivateIpExample(context map[string]interface{}) string {
	return Nprintf(`

resource "google_compute_network" "private_network" {
  name                    = "tf-test-private-network%{random_suffix}"
  auto_create_subnetworks = "false"
}

resource "google_compute_global_address" "private_ip_address" {
  name          = "tf-test-private-ip-address%{random_suffix}"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = google_compute_network.private_network.id
}

resource "google_service_networking_connection" "private_vpc_connection" {
  network                 = google_compute_network.private_network.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.private_ip_address.name]
}

resource "google_sql_database_instance" "instance" {
  name             = "tf-test-private-ip-sql-instance%{random_suffix}"
  region           = "us-central1"
  database_version = "SQLSERVER_2019_STANDARD"
  root_password        = "INSERT-PASSWORD-HERE"

  depends_on = [google_service_networking_connection.private_vpc_connection]

  settings {
    tier = "db-custom-2-7680"
    ip_configuration {
      ipv4_enabled    = "false"
      private_network = google_compute_network.private_network.id
    }
  }
  deletion_protection = "false"
}

`, context)
}

func TestAccCGCSnippet_sqlSqlserverInstanceFlagsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCGCSnippet_sqlSqlserverInstanceFlagsExample(context),
			},
			{
				ResourceName:            "google_sql_database_instance.instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"deletion_protection", "root_password"},
			},
		},
	})
}

func testAccCGCSnippet_sqlSqlserverInstanceFlagsExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_sql_database_instance" "instance" {
  name             = "tf-test-sqlserver-instance%{random_suffix}"
  region           = "us-central1"
  database_version = "SQLSERVER_2019_STANDARD"
  root_password = "INSERT-PASSWORD-HERE"
  settings {
    database_flags {
      name  = "1204"
      value = "on"
    }
    database_flags {
      name  = "remote access"
      value = "on"
    }
    database_flags {
      name  = "remote query timeout (s)"
      value = "300"
    }
    tier = "db-custom-2-7680"
  }
  deletion_protection = "%{deletion_protection}"
}
`, context)
}

func TestAccCGCSnippet_sqlInstanceCmekExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccCGCSnippet_sqlInstanceCmekExample(context),
			},
			{
				ResourceName:            "google_sql_database_instance.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"deletion_protection", "root_password"},
			},
		},
	})
}

func testAccCGCSnippet_sqlInstanceCmekExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project_service_identity" "gcp_sa_cloud_sql" {
  provider = google-beta
  service  = "sqladmin.googleapis.com"
}

resource "google_kms_key_ring" "keyring" {
  provider = google-beta
  name     = "tf-test-keyring-name%{random_suffix}"
  location = "us-central1"
}

resource "google_kms_crypto_key" "key" {
  provider = google-beta
  name     = "tf-test-crypto-key-name%{random_suffix}"
  key_ring = google_kms_key_ring.keyring.id
  purpose  = "ENCRYPT_DECRYPT"
}

resource "google_kms_crypto_key_iam_binding" "crypto_key" {
  provider      = google-beta
  crypto_key_id = google_kms_crypto_key.key.id
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"

  members = [
    "serviceAccount:${google_project_service_identity.gcp_sa_cloud_sql.email}",
  ]
}

resource "google_sql_database_instance" "mysql_instance_with_cmek" {
  name                = "tf-test-mysql-instance-cmek%{random_suffix}"
  provider            = google-beta
  region              = "us-central1"
  database_version    = "MYSQL_8_0"
  encryption_key_name = google_kms_crypto_key.key.id
  settings {
    tier = "db-n1-standard-2"
  }
  deletion_protection =  "%{deletion_protection}"
}

resource "google_sql_database_instance" "postgres_instance_with_cmek" {
  name                = "tf-test-postgres-instance-cmek%{random_suffix}"
  provider            = google-beta
  region              = "us-central1"
  database_version    = "POSTGRES_14"
  encryption_key_name = google_kms_crypto_key.key.id
  settings {
    tier = "db-custom-2-7680"
  }
  deletion_protection =  "%{deletion_protection}"
}

resource "google_sql_database_instance" "default" {
  name                = "tf-test-sqlserver-instance-cmek%{random_suffix}"
  provider            = google-beta
  region              = "us-central1"
  database_version    = "SQLSERVER_2019_STANDARD"
  root_password       = "INSERT-PASSWORD-HERE "
  encryption_key_name = google_kms_crypto_key.key.id
  settings {
    tier = "db-custom-2-7680"
  }
  deletion_protection =  "%{deletion_protection}"
}
`, context)
}

func TestAccCGCSnippet_storageNewBucketExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCGCSnippet_storageNewBucketExample(context),
			},
			{
				ResourceName:      "google_storage_bucket.static",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCGCSnippet_storageNewBucketExample(context map[string]interface{}) string {
	return Nprintf(`
# Create new storage bucket in the US multi-region
# with coldline storage
resource "google_storage_bucket" "static" {
  name          = "tf-test-new-bucket%{random_suffix}"
  location      = "US"
  storage_class = "COLDLINE"

  uniform_bucket_level_access = true
} 

# Upload files
# Discussion about using tf to upload a large number of objects
# https://stackoverflow.com/questions/68455132/terraform-copy-multiple-files-to-bucket-at-the-same-time-bucket-creation

# The text object in Cloud Storage
resource "google_storage_bucket_object" "default" {
  name         = "tf-test-new-object%{random_suffix}"
# Uncomment and add valid path to an object.
#  source       = "/path/to/an/object"
  content      = "Data as string to be uploaded"
  content_type = "text/plain"
  bucket       = google_storage_bucket.static.id
}

# Get object metadata
data "google_storage_bucket_object" "default" {
  name         = google_storage_bucket_object.default.name
  bucket       = google_storage_bucket.static.id
}

output "object_metadata" {
  value        = data.google_storage_bucket_object.default
}

# Get bucket metadata
data "google_storage_bucket" "default" {
  name         = google_storage_bucket.static.id
}

output "bucket_metadata" {
  value        = data.google_storage_bucket.default
}
`, context)
}

func TestAccCGCSnippet_storageMakeDataPublicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccCGCSnippet_storageMakeDataPublicExample(context),
			},
			{
				ResourceName:      "google_storage_bucket.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCGCSnippet_storageMakeDataPublicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "default" {
  provider                    = google-beta
  name                        = "tf-test-example-bucket-name%{random_suffix}"
  location                    = "US"
  uniform_bucket_level_access = true
}

# Make bucket public
resource "google_storage_bucket_iam_member" "member" {
  provider = google-beta
  bucket   = google_storage_bucket.default.name
  role     = "roles/storage.objectViewer"
  member   = "allUsers"
}
`, context)
}

func TestAccCGCSnippet_storagePubsubNotificationsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccCGCSnippet_storagePubsubNotificationsExample(context),
			},
			{
				ResourceName:      "google_pubsub_topic.topic",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCGCSnippet_storagePubsubNotificationsExample(context map[string]interface{}) string {
	return Nprintf(`
// Create a Pub/Sub notification.
resource "google_storage_notification" "notification" {
  provider       = google-beta
  bucket         = google_storage_bucket.bucket.name
  payload_format = "JSON_API_V1"
  topic          = google_pubsub_topic.topic.id
  depends_on = [google_pubsub_topic_iam_binding.binding]
}

// Enable notifications by giving the correct IAM permission to the unique service account.
data "google_storage_project_service_account" "gcs_account" {
  provider = google-beta
}

// Create a Pub/Sub topic.
resource "google_pubsub_topic_iam_binding" "binding" {
  provider = google-beta
  topic    = google_pubsub_topic.topic.id
  role     = "roles/pubsub.publisher"
  members  = ["serviceAccount:${data.google_storage_project_service_account.gcs_account.email_address}"]
}

// Create a new storage bucket.
resource "google_storage_bucket" "bucket" {
  name     = "tf-test-example-bucket-name%{random_suffix}"
  provider = google-beta
  location = "US"
  uniform_bucket_level_access = true
}

resource "google_pubsub_topic" "topic" {
  name     = "tf_test_your_topic_name%{random_suffix}"
  provider = google-beta
}
`, context)
}
