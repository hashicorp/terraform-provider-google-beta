// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccDNSManagedZone_dnsManagedZoneBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDNSManagedZoneDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDNSManagedZone_dnsManagedZoneBasicExample(context),
			},
			{
				ResourceName:      "google_dns_managed_zone.example-zone",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDNSManagedZone_dnsManagedZoneBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dns_managed_zone" "example-zone" {
  name        = "example-zone"
  dns_name    = "example-${random_id.rnd.hex}.com."
  description = "Example DNS zone"
  labels = {
    foo = "bar"
  }
}

resource "random_id" "rnd" {
  byte_length = 4
}
`, context)
}

func TestAccDNSManagedZone_dnsManagedZonePrivateExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDNSManagedZoneDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDNSManagedZone_dnsManagedZonePrivateExample(context),
			},
			{
				ResourceName:      "google_dns_managed_zone.private-zone",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDNSManagedZone_dnsManagedZonePrivateExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dns_managed_zone" "private-zone" {
  name        = "tf-test-private-zone%{random_suffix}"
  dns_name    = "private.example.com."
  description = "Example private DNS zone"
  labels = {
    foo = "bar"
  }

  visibility = "private"

  private_visibility_config {
    networks {
      network_url = google_compute_network.network-1.self_link
    }
    networks {
      network_url = google_compute_network.network-2.self_link
    }
  }
}

resource "google_compute_network" "network-1" {
  name                    = "tf-test-network-1%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_network" "network-2" {
  name                    = "tf-test-network-2%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func TestAccDNSManagedZone_dnsManagedZonePrivatePeeringExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckDNSManagedZoneDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDNSManagedZone_dnsManagedZonePrivatePeeringExample(context),
			},
		},
	})
}

func testAccDNSManagedZone_dnsManagedZonePrivatePeeringExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dns_managed_zone" "peering-zone" {
  provider = google-beta

  name        = "tf-test-peering-zone%{random_suffix}"
  dns_name    = "peering.example.com."
  description = "Example private DNS peering zone"

  visibility = "private"

  private_visibility_config {
    networks {
      network_url = google_compute_network.network-source.self_link
    }
  }

  peering_config {
    target_network {
      network_url = google_compute_network.network-target.self_link
    }
  }
}

resource "google_compute_network" "network-source" {
  provider = google-beta

  name                    = "tf-test-network-source%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_network" "network-target" {
  provider = google-beta

  name                    = "tf-test-network-target%{random_suffix}"
  auto_create_subnetworks = false
}

provider "google-beta" {
  region = "us-central1"
  zone   = "us-central1-a"
}
`, context)
}

func testAccCheckDNSManagedZoneDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_dns_managed_zone" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{DNSBasePath}}projects/{{project}}/managedZones/{{name}}")
			if err != nil {
				return err
			}

			_, err = sendRequest(config, "GET", "", url, nil)
			if err == nil {
				return fmt.Errorf("DNSManagedZone still exists at %s", url)
			}
		}

		return nil
	}
}
