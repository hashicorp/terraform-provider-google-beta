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
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccComputeForwardingRule_forwardingRuleExternallbExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckComputeForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeForwardingRule_forwardingRuleExternallbExample(context),
			},
		},
	})
}

func testAccComputeForwardingRule_forwardingRuleExternallbExample(context map[string]interface{}) string {
	return Nprintf(`
// Forwarding rule for External Network Load Balancing using Backend Services
resource "google_compute_forwarding_rule" "default" {
  provider              = google-beta
  name                  = "tf-test-website-forwarding-rule%{random_suffix}"
  region                = "us-central1"
  port_range            = 80
  backend_service       = google_compute_region_backend_service.backend.id
}
resource "google_compute_region_backend_service" "backend" {
  provider              = google-beta
  name                  = "tf-test-website-backend%{random_suffix}"
  region                = "us-central1"
  load_balancing_scheme = "EXTERNAL"
  health_checks         = [google_compute_region_health_check.hc.id]
}
resource "google_compute_region_health_check" "hc" {
  provider           = google-beta
  name               = "check-tf-test-website-backend%{random_suffix}"
  check_interval_sec = 1
  timeout_sec        = 1
  region             = "us-central1"

  tcp_health_check {
    port = "80"
  }
}
`, context)
}

func TestAccComputeForwardingRule_forwardingRuleGlobalInternallbExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeForwardingRule_forwardingRuleGlobalInternallbExample(context),
			},
			{
				ResourceName:            "google_compute_forwarding_rule.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"backend_service", "network", "subnetwork", "region"},
			},
		},
	})
}

func testAccComputeForwardingRule_forwardingRuleGlobalInternallbExample(context map[string]interface{}) string {
	return Nprintf(`
// Forwarding rule for Internal Load Balancing
resource "google_compute_forwarding_rule" "default" {
  name                  = "tf-test-website-forwarding-rule%{random_suffix}"
  region                = "us-central1"
  load_balancing_scheme = "INTERNAL"
  backend_service       = google_compute_region_backend_service.backend.id
  all_ports             = true
  allow_global_access   = true
  network               = google_compute_network.default.name
  subnetwork            = google_compute_subnetwork.default.name
}
resource "google_compute_region_backend_service" "backend" {
  name                  = "tf-test-website-backend%{random_suffix}"
  region                = "us-central1"
  health_checks         = [google_compute_health_check.hc.id]
}
resource "google_compute_health_check" "hc" {
  name               = "check-tf-test-website-backend%{random_suffix}"
  check_interval_sec = 1
  timeout_sec        = 1
  tcp_health_check {
    port = "80"
  }
}
resource "google_compute_network" "default" {
  name = "tf-test-website-net%{random_suffix}"
  auto_create_subnetworks = false
}
resource "google_compute_subnetwork" "default" {
  name          = "tf-test-website-net%{random_suffix}"
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.default.id
}
`, context)
}

func TestAccComputeForwardingRule_forwardingRuleBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeForwardingRule_forwardingRuleBasicExample(context),
			},
			{
				ResourceName:            "google_compute_forwarding_rule.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"backend_service", "network", "subnetwork", "region"},
			},
		},
	})
}

func testAccComputeForwardingRule_forwardingRuleBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_forwarding_rule" "default" {
  name       = "tf-test-website-forwarding-rule%{random_suffix}"
  target     = google_compute_target_pool.default.id
  port_range = "80"
}

resource "google_compute_target_pool" "default" {
  name = "tf-test-website-target-pool%{random_suffix}"
}
`, context)
}

func TestAccComputeForwardingRule_forwardingRuleL3DefaultExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckComputeForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeForwardingRule_forwardingRuleL3DefaultExample(context),
			},
		},
	})
}

func testAccComputeForwardingRule_forwardingRuleL3DefaultExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_forwarding_rule" "fwd_rule" {
  provider        = google-beta
  name            = "tf-test-l3-forwarding-rule%{random_suffix}"
  backend_service = google_compute_region_backend_service.service.id
  ip_protocol     = "L3_DEFAULT"
  all_ports       = true
}

resource "google_compute_region_backend_service" "service" {
  provider              = google-beta
  region                = "us-central1"
  name                  = "service%{random_suffix}"
  health_checks         = [google_compute_region_health_check.health_check.id]
  protocol              = "UNSPECIFIED"
  load_balancing_scheme = "EXTERNAL"
}

resource "google_compute_region_health_check" "health_check" {
  provider           = google-beta
  name               = "tf-test-health-check%{random_suffix}"
  region             = "us-central1"

  tcp_health_check {
    port = 80
  }
}
`, context)
}

func TestAccComputeForwardingRule_forwardingRuleInternallbExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeForwardingRule_forwardingRuleInternallbExample(context),
			},
			{
				ResourceName:            "google_compute_forwarding_rule.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"backend_service", "network", "subnetwork", "region"},
			},
		},
	})
}

func testAccComputeForwardingRule_forwardingRuleInternallbExample(context map[string]interface{}) string {
	return Nprintf(`
// Forwarding rule for Internal Load Balancing
resource "google_compute_forwarding_rule" "default" {
  name   = "tf-test-website-forwarding-rule%{random_suffix}"
  region = "us-central1"

  load_balancing_scheme = "INTERNAL"
  backend_service       = google_compute_region_backend_service.backend.id
  all_ports             = true
  network               = google_compute_network.default.name
  subnetwork            = google_compute_subnetwork.default.name
}

resource "google_compute_region_backend_service" "backend" {
  name          = "tf-test-website-backend%{random_suffix}"
  region        = "us-central1"
  health_checks = [google_compute_health_check.hc.id]
}

resource "google_compute_health_check" "hc" {
  name               = "check-tf-test-website-backend%{random_suffix}"
  check_interval_sec = 1
  timeout_sec        = 1

  tcp_health_check {
    port = "80"
  }
}

resource "google_compute_network" "default" {
  name                    = "tf-test-website-net%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  name          = "tf-test-website-net%{random_suffix}"
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.default.id
}
`, context)
}

func TestAccComputeForwardingRule_forwardingRuleHttpLbExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckComputeForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeForwardingRule_forwardingRuleHttpLbExample(context),
			},
		},
	})
}

func testAccComputeForwardingRule_forwardingRuleHttpLbExample(context map[string]interface{}) string {
	return Nprintf(`
// Forwarding rule for Internal Load Balancing
resource "google_compute_forwarding_rule" "default" {
  provider = google-beta
  depends_on = [google_compute_subnetwork.proxy]
  name   = "tf-test-website-forwarding-rule%{random_suffix}"
  region = "us-central1"

  ip_protocol           = "TCP"
  load_balancing_scheme = "INTERNAL_MANAGED"
  port_range            = "80"
  target                = google_compute_region_target_http_proxy.default.id
  network               = google_compute_network.default.id
  subnetwork            = google_compute_subnetwork.default.id
  network_tier          = "PREMIUM"
}

resource "google_compute_region_target_http_proxy" "default" {
  provider = google-beta

  region  = "us-central1"
  name    = "tf-test-website-proxy%{random_suffix}"
  url_map = google_compute_region_url_map.default.id
}

resource "google_compute_region_url_map" "default" {
  provider = google-beta

  region          = "us-central1"
  name            = "tf-test-website-map%{random_suffix}"
  default_service = google_compute_region_backend_service.default.id
}

resource "google_compute_region_backend_service" "default" {
  provider = google-beta

  load_balancing_scheme = "INTERNAL_MANAGED"

  backend {
    group = google_compute_region_instance_group_manager.rigm.instance_group
    balancing_mode = "UTILIZATION"
    capacity_scaler = 1.0
  }

  region      = "us-central1"
  name        = "tf-test-website-backend%{random_suffix}"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_region_health_check.default.id]
}

data "google_compute_image" "debian_image" {
  provider = google-beta
  family   = "debian-9"
  project  = "debian-cloud"
}

resource "google_compute_region_instance_group_manager" "rigm" {
  provider = google-beta
  region   = "us-central1"
  name     = "tf-test-website-rigm%{random_suffix}"
  version {
    instance_template = google_compute_instance_template.instance_template.id
    name              = "primary"
  }
  base_instance_name = "internal-glb"
  target_size        = 1
}

resource "google_compute_instance_template" "instance_template" {
  provider     = google-beta
  name         = "template-tf-test-website-backend%{random_suffix}"
  machine_type = "e2-medium"

  network_interface {
    network = google_compute_network.default.id
    subnetwork = google_compute_subnetwork.default.id
  }

  disk {
    source_image = data.google_compute_image.debian_image.self_link
    auto_delete  = true
    boot         = true
  }

  tags = ["allow-ssh", "load-balanced-backend"]
}

resource "google_compute_region_health_check" "default" {
  depends_on = [google_compute_firewall.fw4]
  provider = google-beta

  region = "us-central1"
  name   = "tf-test-website-hc%{random_suffix}"
  http_health_check {
    port_specification = "USE_SERVING_PORT"
  }
}

resource "google_compute_firewall" "fw1" {
  provider = google-beta
  name = "tf-test-website-fw%{random_suffix}-1"
  network = google_compute_network.default.id
  source_ranges = ["10.1.2.0/24"]
  allow {
    protocol = "tcp"
  }
  allow {
    protocol = "udp"
  }
  allow {
    protocol = "icmp"
  }
  direction = "INGRESS"
}

resource "google_compute_firewall" "fw2" {
  depends_on = [google_compute_firewall.fw1]
  provider = google-beta
  name = "tf-test-website-fw%{random_suffix}-2"
  network = google_compute_network.default.id
  source_ranges = ["0.0.0.0/0"]
  allow {
    protocol = "tcp"
    ports = ["22"]
  }
  target_tags = ["allow-ssh"]
  direction = "INGRESS"
}

resource "google_compute_firewall" "fw3" {
  depends_on = [google_compute_firewall.fw2]
  provider = google-beta
  name = "tf-test-website-fw%{random_suffix}-3"
  network = google_compute_network.default.id
  source_ranges = ["130.211.0.0/22", "35.191.0.0/16"]
  allow {
    protocol = "tcp"
  }
  target_tags = ["load-balanced-backend"]
  direction = "INGRESS"
}

resource "google_compute_firewall" "fw4" {
  depends_on = [google_compute_firewall.fw3]
  provider = google-beta
  name = "tf-test-website-fw%{random_suffix}-4"
  network = google_compute_network.default.id
  source_ranges = ["10.129.0.0/26"]
  target_tags = ["load-balanced-backend"]
  allow {
    protocol = "tcp"
    ports = ["80"]
  }
  allow {
    protocol = "tcp"
    ports = ["443"]
  }
  allow {
    protocol = "tcp"
    ports = ["8000"]
  }
  direction = "INGRESS"
}

resource "google_compute_network" "default" {
  provider = google-beta
  name                    = "tf-test-website-net%{random_suffix}"
  auto_create_subnetworks = false
  routing_mode = "REGIONAL"
}

resource "google_compute_subnetwork" "default" {
  provider = google-beta
  name          = "tf-test-website-net%{random_suffix}-default"
  ip_cidr_range = "10.1.2.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.id
}

resource "google_compute_subnetwork" "proxy" {
  provider = google-beta
  name          = "tf-test-website-net%{random_suffix}-proxy"
  ip_cidr_range = "10.129.0.0/26"
  region        = "us-central1"
  network       = google_compute_network.default.id
  purpose       = "INTERNAL_HTTPS_LOAD_BALANCER"
  role          = "ACTIVE"
}
`, context)
}

func testAccCheckComputeForwardingRuleDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_forwarding_rule" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/forwardingRules/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("ComputeForwardingRule still exists at %s", url)
			}
		}

		return nil
	}
}
