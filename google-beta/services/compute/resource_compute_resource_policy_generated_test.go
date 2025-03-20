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

package compute_test

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

func TestAccComputeResourcePolicy_resourcePolicyBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeResourcePolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeResourcePolicy_resourcePolicyBasicExample(context),
			},
			{
				ResourceName:            "google_compute_resource_policy.foo",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region"},
			},
		},
	})
}

func testAccComputeResourcePolicy_resourcePolicyBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_resource_policy" "foo" {
  name   = "tf-test-gce-policy%{random_suffix}"
  region = "us-central1"
  snapshot_schedule_policy {
    schedule {
      daily_schedule {
        days_in_cycle = 1
        start_time    = "04:00"
      }
    }
  }
}
`, context)
}

func TestAccComputeResourcePolicy_resourcePolicyHourlyFormatExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeResourcePolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeResourcePolicy_resourcePolicyHourlyFormatExample(context),
			},
			{
				ResourceName:            "google_compute_resource_policy.foo",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region"},
			},
		},
	})
}

func testAccComputeResourcePolicy_resourcePolicyHourlyFormatExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_resource_policy" "foo" {
  name   = "tf-test-gce-policy%{random_suffix}"
  region = "us-central1"
  snapshot_schedule_policy {
    schedule {
      daily_schedule {
        days_in_cycle = 1
        start_time    = "4:00"
      }
    }
  }
}
`, context)
}

func TestAccComputeResourcePolicy_resourcePolicyFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeResourcePolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeResourcePolicy_resourcePolicyFullExample(context),
			},
			{
				ResourceName:            "google_compute_resource_policy.bar",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region"},
			},
		},
	})
}

func testAccComputeResourcePolicy_resourcePolicyFullExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_resource_policy" "bar" {
  name   = "tf-test-gce-policy%{random_suffix}"
  region = "us-central1"
  snapshot_schedule_policy {
    schedule {
      hourly_schedule {
        hours_in_cycle = 20
        start_time     = "23:00"
      }
    }
    retention_policy {
      max_retention_days    = 10
      on_source_disk_delete = "KEEP_AUTO_SNAPSHOTS"
    }
    snapshot_properties {
      labels = {
        my_label = "value"
      }
      storage_locations = ["us"]
      guest_flush       = true
    }
  }
}
`, context)
}

func TestAccComputeResourcePolicy_resourcePolicyPlacementPolicyExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeResourcePolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeResourcePolicy_resourcePolicyPlacementPolicyExample(context),
			},
			{
				ResourceName:            "google_compute_resource_policy.baz",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region"},
			},
		},
	})
}

func testAccComputeResourcePolicy_resourcePolicyPlacementPolicyExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_resource_policy" "baz" {
  name   = "tf-test-gce-policy%{random_suffix}"
  region = "us-central1"
  group_placement_policy {
    vm_count = 2
    collocation = "COLLOCATED"
  }
}
`, context)
}

func TestAccComputeResourcePolicy_resourcePolicyPlacementPolicyMaxDistanceExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeResourcePolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeResourcePolicy_resourcePolicyPlacementPolicyMaxDistanceExample(context),
			},
			{
				ResourceName:            "google_compute_resource_policy.baz",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region"},
			},
		},
	})
}

func testAccComputeResourcePolicy_resourcePolicyPlacementPolicyMaxDistanceExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_resource_policy" "baz" {
  name   = "tf-test-gce-policy%{random_suffix}"
  region = "us-central1"
  provider = google-beta
  group_placement_policy {
    vm_count = 2
    collocation = "COLLOCATED"
    max_distance = 2
  }
}
`, context)
}

func TestAccComputeResourcePolicy_resourcePolicyInstanceSchedulePolicyExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeResourcePolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeResourcePolicy_resourcePolicyInstanceSchedulePolicyExample(context),
			},
			{
				ResourceName:            "google_compute_resource_policy.hourly",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region"},
			},
		},
	})
}

func testAccComputeResourcePolicy_resourcePolicyInstanceSchedulePolicyExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_resource_policy" "hourly" {
  name   = "tf-test-gce-policy%{random_suffix}"
  region = "us-central1"
  description = "Start and stop instances"
  instance_schedule_policy {
    vm_start_schedule {
      schedule = "0 * * * *"
    }
    vm_stop_schedule {
      schedule = "15 * * * *"
    }
    time_zone = "US/Central"
  }
}
`, context)
}

func TestAccComputeResourcePolicy_resourcePolicySnapshotScheduleChainNameExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeResourcePolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeResourcePolicy_resourcePolicySnapshotScheduleChainNameExample(context),
			},
			{
				ResourceName:            "google_compute_resource_policy.hourly",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region"},
			},
		},
	})
}

func testAccComputeResourcePolicy_resourcePolicySnapshotScheduleChainNameExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_resource_policy" "hourly" {
  name   = "tf-test-gce-policy%{random_suffix}"
  region = "us-central1"
  description = "chain name snapshot"
  snapshot_schedule_policy {
    schedule {
      hourly_schedule {
        hours_in_cycle = 20
        start_time     = "23:00"
      }
    }
    retention_policy {
      max_retention_days    = 14
      on_source_disk_delete = "KEEP_AUTO_SNAPSHOTS"
    }
    snapshot_properties {
      labels = {
        my_label = "value"
      }
      storage_locations = ["us"]
      guest_flush       = true
      chain_name = "test-schedule-chain-name"
    }
  }
}
`, context)
}

func TestAccComputeResourcePolicy_resourcePolicyConsistencyGroupExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeResourcePolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeResourcePolicy_resourcePolicyConsistencyGroupExample(context),
			},
			{
				ResourceName:            "google_compute_resource_policy.cgroup",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region"},
			},
		},
	})
}

func testAccComputeResourcePolicy_resourcePolicyConsistencyGroupExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_resource_policy" "cgroup" {
  name   = "tf-test-gce-policy%{random_suffix}"
  region = "europe-west1"
  disk_consistency_group_policy {
    enabled = true
  }
}
`, context)
}

func TestAccComputeResourcePolicy_resourcePolicyWorkloadPolicyExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeResourcePolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeResourcePolicy_resourcePolicyWorkloadPolicyExample(context),
			},
			{
				ResourceName:            "google_compute_resource_policy.bar",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region"},
			},
		},
	})
}

func testAccComputeResourcePolicy_resourcePolicyWorkloadPolicyExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_resource_policy" "bar" {
  name   = "tf-test-gce-policy%{random_suffix}"
  region = "europe-west1"
  provider = google-beta
  workload_policy {
    type = "HIGH_AVAILABILITY"
  }
}
`, context)
}

func TestAccComputeResourcePolicy_resourcePolicyWorkloadPolicyAcceleratorTopologyExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeResourcePolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeResourcePolicy_resourcePolicyWorkloadPolicyAcceleratorTopologyExample(context),
			},
			{
				ResourceName:            "google_compute_resource_policy.bar",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region"},
			},
		},
	})
}

func testAccComputeResourcePolicy_resourcePolicyWorkloadPolicyAcceleratorTopologyExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_resource_policy" "bar" {
  name   = "tf-test-gce-policy%{random_suffix}"
  region = "europe-west1"
  provider = google-beta
  workload_policy {
    type = "HIGH_THROUGHPUT"
    accelerator_topology = "SOME NEW TOPOLOGY"
  }
}
`, context)
}

func TestAccComputeResourcePolicy_resourcePolicyWorkloadPolicyMaxTopologyDistanceExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeResourcePolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeResourcePolicy_resourcePolicyWorkloadPolicyMaxTopologyDistanceExample(context),
			},
			{
				ResourceName:            "google_compute_resource_policy.bar",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region"},
			},
		},
	})
}

func testAccComputeResourcePolicy_resourcePolicyWorkloadPolicyMaxTopologyDistanceExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_resource_policy" "bar" {
  name   = "tf-test-gce-policy%{random_suffix}"
  region = "europe-west1"
  provider = google-beta
  workload_policy {
    type = "HIGH_THROUGHPUT"
    max_topology_distance = "BLOCK"
  }
}
`, context)
}

func TestAccComputeResourcePolicy_resourcePolicyPlacementPolicyGpuTopologyExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeResourcePolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeResourcePolicy_resourcePolicyPlacementPolicyGpuTopologyExample(context),
			},
			{
				ResourceName:            "google_compute_resource_policy.baz",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region"},
			},
		},
	})
}

func testAccComputeResourcePolicy_resourcePolicyPlacementPolicyGpuTopologyExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_resource_policy" "baz" {
  provider = google-beta
  name   = "tf-test-gce-policy%{random_suffix}"
  region = "europe-west9"
  group_placement_policy {
    vm_count = 2
    collocation = "COLLOCATED"
    gpu_topology = "1x72"
  }
}
`, context)
}

func testAccCheckComputeResourcePolicyDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_resource_policy" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/resourcePolicies/{{name}}")
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
				return fmt.Errorf("ComputeResourcePolicy still exists at %s", url)
			}
		}

		return nil
	}
}
