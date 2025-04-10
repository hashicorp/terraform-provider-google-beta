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

package redis_test

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

func TestAccRedisCluster_redisClusterHaExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection_enabled": false,
		"random_suffix":               acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckRedisClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccRedisCluster_redisClusterHaExample(context),
			},
			{
				ResourceName:            "google_redis_cluster.cluster-ha",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"gcs_source", "managed_backup_source", "name", "psc_configs", "region"},
			},
		},
	})
}

func testAccRedisCluster_redisClusterHaExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_redis_cluster" "cluster-ha" {
  name           = "tf-test-ha-cluster%{random_suffix}"
  shard_count    = 3
  psc_configs {
    network = google_compute_network.consumer_net.id
  }
  region = "us-central1"
  replica_count = 1
  node_type = "REDIS_SHARED_CORE_NANO"
  transit_encryption_mode = "TRANSIT_ENCRYPTION_MODE_DISABLED"
  authorization_mode = "AUTH_MODE_DISABLED"
  redis_configs = {
    maxmemory-policy	= "volatile-ttl"
  }
  deletion_protection_enabled = %{deletion_protection_enabled}

  zone_distribution_config {
    mode = "MULTI_ZONE"
  }
  maintenance_policy {
    weekly_maintenance_window {
      day = "MONDAY"
      start_time {
        hours = 1
        minutes = 0
        seconds = 0
        nanos = 0
      }
    }
  }
  depends_on = [
    google_network_connectivity_service_connection_policy.default
  ]
}

resource "google_network_connectivity_service_connection_policy" "default" {
  name = "tf-test-my-policy%{random_suffix}"
  location = "us-central1"
  service_class = "gcp-memorystore-redis"
  description   = "my basic service connection policy"
  network = google_compute_network.consumer_net.id
  psc_config {
    subnetworks = [google_compute_subnetwork.consumer_subnet.id]
  }
}

resource "google_compute_subnetwork" "consumer_subnet" {
  name          = "tf-test-my-subnet%{random_suffix}"
  ip_cidr_range = "10.0.0.248/29"
  region        = "us-central1"
  network       = google_compute_network.consumer_net.id
}

resource "google_compute_network" "consumer_net" {
  name                    = "tf-test-my-network%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func TestAccRedisCluster_redisClusterHaSingleZoneExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection_enabled": false,
		"random_suffix":               acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckRedisClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccRedisCluster_redisClusterHaSingleZoneExample(context),
			},
			{
				ResourceName:            "google_redis_cluster.cluster-ha-single-zone",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"gcs_source", "managed_backup_source", "name", "psc_configs", "region"},
			},
		},
	})
}

func testAccRedisCluster_redisClusterHaSingleZoneExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_redis_cluster" "cluster-ha-single-zone" {
  name           = "tf-test-ha-cluster-single-zone%{random_suffix}"
  shard_count    = 3
  psc_configs {
    network = google_compute_network.consumer_net.id
  }
  region = "us-central1"
  zone_distribution_config {
    mode = "SINGLE_ZONE"
    zone = "us-central1-f"
  }
  maintenance_policy {
    weekly_maintenance_window {
      day = "MONDAY"
      start_time {
        hours = 1
        minutes = 0
        seconds = 0
        nanos = 0
      }
    }
  }
  deletion_protection_enabled = %{deletion_protection_enabled}
  depends_on = [
    google_network_connectivity_service_connection_policy.default
  ]

}

resource "google_network_connectivity_service_connection_policy" "default" {
  name = "tf-test-my-policy%{random_suffix}"
  location = "us-central1"
  service_class = "gcp-memorystore-redis"
  description   = "my basic service connection policy"
  network = google_compute_network.consumer_net.id
  psc_config {
    subnetworks = [google_compute_subnetwork.consumer_subnet.id]
  }
}

resource "google_compute_subnetwork" "consumer_subnet" {
  name          = "tf-test-my-subnet%{random_suffix}"
  ip_cidr_range = "10.0.0.248/29"
  region        = "us-central1"
  network       = google_compute_network.consumer_net.id
}

resource "google_compute_network" "consumer_net" {
  name                    = "tf-test-my-network%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func TestAccRedisCluster_redisClusterSecondaryExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"primary_cluster_deletion_protection_enabled":   false,
		"secondary_cluster_deletion_protection_enabled": false,
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckRedisClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccRedisCluster_redisClusterSecondaryExample(context),
			},
			{
				ResourceName:            "google_redis_cluster.secondary_cluster",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"gcs_source", "managed_backup_source", "name", "psc_configs", "region"},
			},
		},
	})
}

func testAccRedisCluster_redisClusterSecondaryExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
// Primary cluster
resource "google_redis_cluster" "primary_cluster" {
  name          = "tf-test-my-primary-cluster%{random_suffix}"
  region        = "us-east1"
  psc_configs {
    network = google_compute_network.consumer_net.id
  }

  // Settings that should match on primary and secondary clusters. 
  // If you define a setting here, ensure that the secondary clusters also define it with the same values. 
  // Please see https://cloud.google.com/memorystore/docs/cluster/about-cross-region-replication#settings_copied_from_the_primary_during_instance_creation for the complete list of such settings.
  authorization_mode = "AUTH_MODE_DISABLED"
  transit_encryption_mode = "TRANSIT_ENCRYPTION_MODE_DISABLED"
  shard_count   = 3
  redis_configs = {
    maxmemory-policy = "volatile-ttl"
  }
  node_type = "REDIS_HIGHMEM_MEDIUM"
  persistence_config {
    mode = "RDB"
    rdb_config {
      rdb_snapshot_period = "ONE_HOUR"
      rdb_snapshot_start_time = "2024-10-02T15:01:23Z"
    }
  }

  // Settings that can have different values on primary and secondary clusters.
  // Please see https://cloud.google.com/memorystore/docs/cluster/about-cross-region-replication#override_allowed_during_instance_creation for the complete list of such settings.
  zone_distribution_config {
    mode = "MULTI_ZONE"
  }
  replica_count = 1
  maintenance_policy {
    weekly_maintenance_window {
      day = "MONDAY"
      start_time {
        hours = 1
        minutes = 0
        seconds = 0
        nanos = 0
      }
    }
  }
  deletion_protection_enabled = %{primary_cluster_deletion_protection_enabled}

  depends_on = [
    google_network_connectivity_service_connection_policy.primary_cluster_region_scp
  ]
}


// Secondary cluster
resource "google_redis_cluster" "secondary_cluster" {
  name          = "tf-test-my-secondary-cluster%{random_suffix}"
  region        = "europe-west1"
  psc_configs {
    network = google_compute_network.consumer_net.id
  }

  // Settings that should match on primary and secondary clusters. 
  // If you defined a setting here for primary, ensure the secondary clusters also define it with the same values. 
  // Please see https://cloud.google.com/memorystore/docs/cluster/about-cross-region-replication#settings_copied_from_the_primary_during_instance_creation for the complete list of such settings.
  authorization_mode = "AUTH_MODE_DISABLED"
  transit_encryption_mode = "TRANSIT_ENCRYPTION_MODE_DISABLED"
  shard_count   = 3
  redis_configs = {
    maxmemory-policy = "volatile-ttl"
  }
  node_type = "REDIS_HIGHMEM_MEDIUM"
  persistence_config {
    mode = "RDB"
    rdb_config {
      rdb_snapshot_period = "ONE_HOUR"
      rdb_snapshot_start_time = "2024-10-02T15:01:23Z"
    }
  }

  // Settings that can be different on primary and secondary clusters.
  // Please see https://cloud.google.com/memorystore/docs/cluster/about-cross-region-replication#override_allowed_during_instance_creation for the complete list of such settings.
  zone_distribution_config {
    mode = "MULTI_ZONE"
  }
  replica_count = 2
  maintenance_policy {
    weekly_maintenance_window {
      day = "WEDNESDAY"
      start_time {
        hours = 1
        minutes = 0
        seconds = 0
        nanos = 0
      }
    }
  }
  deletion_protection_enabled = %{secondary_cluster_deletion_protection_enabled}

  // Cross cluster replication config
  cross_cluster_replication_config {
    cluster_role = "SECONDARY"
    primary_cluster {
      cluster = google_redis_cluster.primary_cluster.id
    }
  }

  depends_on = [
    google_network_connectivity_service_connection_policy.secondary_cluster_region_scp
  ]
}


resource "google_network_connectivity_service_connection_policy" "primary_cluster_region_scp" {
  name = "tf-test-mypolicy-primary-cluster%{random_suffix}"
  location = "us-east1"
  service_class = "gcp-memorystore-redis"
  description   = "Primary cluster service connection policy"
  network = google_compute_network.consumer_net.id
  psc_config {
    subnetworks = [google_compute_subnetwork.primary_cluster_consumer_subnet.id]
  }
}

resource "google_compute_subnetwork" "primary_cluster_consumer_subnet" {
  name          = "tf-test-mysubnet-primary-cluster%{random_suffix}"
  ip_cidr_range = "10.0.1.0/29"
  region        = "us-east1"
  network       = google_compute_network.consumer_net.id
}


resource "google_network_connectivity_service_connection_policy" "secondary_cluster_region_scp" {
  name = "tf-test-mypolicy-secondary-cluster%{random_suffix}"
  location = "europe-west1"
  service_class = "gcp-memorystore-redis"
  description   = "Secondary cluster service connection policy"
  network = google_compute_network.consumer_net.id
  psc_config {
    subnetworks = [google_compute_subnetwork.secondary_cluster_consumer_subnet.id]
  }
}

resource "google_compute_subnetwork" "secondary_cluster_consumer_subnet" {
  name          = "tf-test-mysubnet-secondary-cluster%{random_suffix}"
  ip_cidr_range = "10.0.2.0/29"
  region        = "europe-west1"
  network       = google_compute_network.consumer_net.id
}

resource "google_compute_network" "consumer_net" {
  name                    = "mynetwork%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func TestAccRedisCluster_redisClusterRdbExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection_enabled": false,
		"random_suffix":               acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckRedisClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccRedisCluster_redisClusterRdbExample(context),
			},
			{
				ResourceName:            "google_redis_cluster.cluster-rdb",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"gcs_source", "managed_backup_source", "name", "psc_configs", "region"},
			},
		},
	})
}

func testAccRedisCluster_redisClusterRdbExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_redis_cluster" "cluster-rdb" {
  name           = "tf-test-rdb-cluster%{random_suffix}"
  shard_count    = 3
  psc_configs {
    network = google_compute_network.consumer_net.id
  }
  region = "us-central1"
  replica_count = 0
  node_type = "REDIS_SHARED_CORE_NANO"
  transit_encryption_mode = "TRANSIT_ENCRYPTION_MODE_DISABLED"
  authorization_mode = "AUTH_MODE_DISABLED"
  redis_configs = {
    maxmemory-policy	= "volatile-ttl"
  }
  deletion_protection_enabled = %{deletion_protection_enabled}

  zone_distribution_config {
    mode = "MULTI_ZONE"
  }
  maintenance_policy {
    weekly_maintenance_window {
      day = "MONDAY"
      start_time {
        hours = 1
        minutes = 0
        seconds = 0
        nanos = 0
      }
    }
  }
  persistence_config { 
    mode = "RDB"
    rdb_config {
      rdb_snapshot_period = "ONE_HOUR"
      rdb_snapshot_start_time = "2024-10-02T15:01:23Z"
    }
  }
  depends_on = [
    google_network_connectivity_service_connection_policy.default
  ]
}

resource "google_network_connectivity_service_connection_policy" "default" {
  name = "tf-test-my-policy%{random_suffix}"
  location = "us-central1"
  service_class = "gcp-memorystore-redis"
  description   = "my basic service connection policy"
  network = google_compute_network.consumer_net.id
  psc_config {
    subnetworks = [google_compute_subnetwork.consumer_subnet.id]
  }
}

resource "google_compute_subnetwork" "consumer_subnet" {
  name          = "tf-test-my-subnet%{random_suffix}"
  ip_cidr_range = "10.0.0.248/29"
  region        = "us-central1"
  network       = google_compute_network.consumer_net.id
}

resource "google_compute_network" "consumer_net" {
  name                    = "tf-test-my-network%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func TestAccRedisCluster_redisClusterAofExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection_enabled": false,
		"random_suffix":               acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckRedisClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccRedisCluster_redisClusterAofExample(context),
			},
			{
				ResourceName:            "google_redis_cluster.cluster-aof",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"gcs_source", "managed_backup_source", "name", "psc_configs", "region"},
			},
		},
	})
}

func testAccRedisCluster_redisClusterAofExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_redis_cluster" "cluster-aof" {
  name           = "tf-test-aof-cluster%{random_suffix}"
  shard_count    = 3
  psc_configs {
    network = google_compute_network.consumer_net.id
  }
  region = "us-central1"
  replica_count = 0
  node_type = "REDIS_SHARED_CORE_NANO"
  transit_encryption_mode = "TRANSIT_ENCRYPTION_MODE_DISABLED"
  authorization_mode = "AUTH_MODE_DISABLED"
  redis_configs = {
    maxmemory-policy	= "volatile-ttl"
  }
  deletion_protection_enabled = %{deletion_protection_enabled}

  zone_distribution_config {
    mode = "MULTI_ZONE"
  }
  maintenance_policy {
    weekly_maintenance_window {
      day = "MONDAY"
      start_time {
        hours = 1
        minutes = 0
        seconds = 0
        nanos = 0
      }
    }
  }
  persistence_config { 
    mode = "AOF"
    aof_config {
      append_fsync = "EVERYSEC"
    }
  }
  depends_on = [
    google_network_connectivity_service_connection_policy.default
  ]
}

resource "google_network_connectivity_service_connection_policy" "default" {
  name = "tf-test-my-policy%{random_suffix}"
  location = "us-central1"
  service_class = "gcp-memorystore-redis"
  description   = "my basic service connection policy"
  network = google_compute_network.consumer_net.id
  psc_config {
    subnetworks = [google_compute_subnetwork.consumer_subnet.id]
  }
}

resource "google_compute_subnetwork" "consumer_subnet" {
  name          = "tf-test-my-subnet%{random_suffix}"
  ip_cidr_range = "10.0.0.248/29"
  region        = "us-central1"
  network       = google_compute_network.consumer_net.id
}

resource "google_compute_network" "consumer_net" {
  name                    = "tf-test-my-network%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func TestAccRedisCluster_redisClusterCmekExample(t *testing.T) {
	t.Parallel()
	acctest.BootstrapIamMembers(t, []acctest.IamMember{
		{
			Member: "serviceAccount:service-{project_number}@cloud-redis.iam.gserviceaccount.com",
			Role:   "roles/cloudkms.cryptoKeyEncrypterDecrypter",
		},
	})

	context := map[string]interface{}{
		"deletion_protection_enabled": false,
		"kms_key_name":                acctest.BootstrapKMSKeyInLocation(t, "us-central1").CryptoKey.Name,
		"random_suffix":               acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckRedisClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccRedisCluster_redisClusterCmekExample(context),
			},
			{
				ResourceName:            "google_redis_cluster.cluster-cmek",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"gcs_source", "managed_backup_source", "name", "psc_configs", "region"},
			},
		},
	})
}

func testAccRedisCluster_redisClusterCmekExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_redis_cluster" "cluster-cmek" {
  name           = "tf-test-cmek-cluster%{random_suffix}"
  shard_count    = 3
  psc_configs {
    network = google_compute_network.consumer_net.id
  }
  kms_key = "%{kms_key_name}"
  region = "us-central1"
  deletion_protection_enabled = %{deletion_protection_enabled}
  depends_on = [
    google_network_connectivity_service_connection_policy.default
  ]
}


data "google_project" "project" {
}

resource "google_network_connectivity_service_connection_policy" "default" {
  name = "tf-test-my-policy%{random_suffix}"
  location = "us-central1"
  service_class = "gcp-memorystore-redis"
  description   = "my basic service connection policy"
  network = google_compute_network.consumer_net.id
  psc_config {
    subnetworks = [google_compute_subnetwork.consumer_subnet.id]
  }
}

resource "google_compute_subnetwork" "consumer_subnet" {
  name          = "tf-test-my-subnet%{random_suffix}"
  ip_cidr_range = "10.0.0.248/29"
  region        = "us-central1"
  network       = google_compute_network.consumer_net.id
}

resource "google_compute_network" "consumer_net" {
  name                    = "tf-test-my-network%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func testAccCheckRedisClusterDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_redis_cluster" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{RedisBasePath}}projects/{{project}}/locations/{{region}}/clusters/{{name}}")
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
				return fmt.Errorf("RedisCluster still exists at %s", url)
			}
		}

		return nil
	}
}
