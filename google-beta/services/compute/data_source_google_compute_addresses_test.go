// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
// ----------------------------------------------------------------------------
//
//	***     AUTO GENERATED CODE    ***    Type: Handwritten     ***
//
// ----------------------------------------------------------------------------
//
//	This code is generated by Magic Modules using the following:
//
//	Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/services/compute/data_source_google_compute_addresses_test.go.tmpl
//
//	DO NOT EDIT this file directly. Any changes made to this file will be
//	overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
package compute_test

import (
	"fmt"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccDataSourceComputeAddresses(t *testing.T) {
	t.Parallel()

	addressName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))

	region := "europe-west8"
	region_bis := "asia-east1"
	dsName := "regional_addresses"
	dsFullName := fmt.Sprintf("data.google_compute_addresses.%s", dsName)
	dsAllName := "all_addresses"
	dsAllFullName := fmt.Sprintf("data.google_compute_addresses.%s", dsAllName)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceComputeAddressesConfig(addressName, region, region_bis),
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceComputeAddressesRegionSpecificCheck(t, addressName, dsFullName, region),
					testAccDataSourceComputeAddressesAllRegionsCheck(t, addressName, dsAllFullName, region, region_bis),
				),
			},
		},
	})
}

func testAccDataSourceComputeAddressesAllRegionsCheck(t *testing.T, address_name string, data_source_name string, expected_region string, expected_region_bis string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expected_addresses := buildAddressesList(3, address_name, expected_region)
		expected_addresses = append(expected_addresses, buildAddressesList(3, address_name, expected_region_bis)...)

		return testDataSourceAdressContains(s, data_source_name, expected_addresses)
	}
}

func testAccDataSourceComputeAddressesRegionSpecificCheck(t *testing.T, address_name string, data_source_name string, expected_region string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expected_addresses := buildAddressesList(3, address_name, expected_region)
		return testDataSourceAdressContains(s, data_source_name, expected_addresses)
	}
}

func testAccDataSourceComputeAddressesConfig(addressName, region, region_bis string) string {
	return fmt.Sprintf(`
locals { 
	region = "%s"
	region_bis  = "%s"
	address_name = "%s"
}

resource "google_compute_address" "address" {
  count = 3

  region = local.region
  name = "${local.address_name}-${local.region}-${count.index}"
  labels = {
	mykey = "myvalue"
  }
}

resource "google_compute_address" "address_region_bis" {
  count = 3

  region = local.region_bis
  name = "${local.address_name}-${local.region_bis}-${count.index}"
  labels = {
	mykey = "myvalue"
  }
}

data "google_compute_addresses" "regional_addresses" {
	filter = "name:${local.address_name}-*"
	depends_on = [google_compute_address.address]
	region = local.region
}

data "google_compute_addresses" "all_addresses" {
	filter = "name:${local.address_name}-*"
	depends_on = [google_compute_address.address, google_compute_address.address_region_bis]
}
`, region, region_bis, addressName)
}

type expectedAddress struct {
	name   string
	region string
}

func (r expectedAddress) checkAddressMatch(index int, attrs map[string]string) (bool, error) {
	map_name := fmt.Sprintf("addresses.%d.name", index)
	address_name := attrs[map_name]

	if address_name != r.name {
		return false, nil
	}

	map_region := fmt.Sprintf("addresses.%d.region", index)
	region, found := attrs[map_region]
	if !found {
		return false, fmt.Errorf("%s doesn't exists", map_region)
	}
	if region != r.region {
		return false, fmt.Errorf("Unexpected region: got %s expected %s", region, r.region)
	}

	map_label := fmt.Sprintf("addresses.%d.labels.mykey", index)
	label_value, found := attrs[map_label]
	if !found {
		return false, fmt.Errorf("label with key 'mykey' not found for %s", address_name)
	}
	if label_value != "myvalue" {
		return false, fmt.Errorf("label value of 'mykey' not equal to 'myvalue' for %s, got %s", address_name, label_value)
	}

	return true, nil
}

func testDataSourceAdressContains(state *terraform.State, data_source_name string, addresses []expectedAddress) error {
	ds, ok := state.RootModule().Resources[data_source_name]
	if !ok {
		return fmt.Errorf("root module has no resource called %s", data_source_name)
	}

	ds_attr := ds.Primary.Attributes

	addresses_length := len(addresses)

	if ds_attr["addresses.#"] != fmt.Sprintf("%d", addresses_length) {
		return fmt.Errorf("addresses.# is not equal to %d", addresses_length)
	}

	for address_index := 0; address_index < addresses_length; address_index++ {
		has_match := false
		for j := 0; j < len(addresses); j++ {
			match, err := addresses[j].checkAddressMatch(address_index, ds_attr)
			if err != nil {
				return err
			} else {
				if match {
					has_match = true
					addresses = removeExpectedAddress(addresses, j)
					break
				}
			}
		}
		if !has_match {
			return fmt.Errorf("unexpected address at index %d", address_index) // TODO improve
		}
	}

	if len(addresses) != 0 {
		return fmt.Errorf("%+v not found in data source", addresses)
	}
	return nil
}

func buildAddressesList(numberofAddresses int, addressName string, region string) []expectedAddress {
	var addresses []expectedAddress
	for i := 0; i < numberofAddresses; i++ {
		addresses = append(addresses, expectedAddress{
			name:   fmt.Sprintf("%s-%s-%d", addressName, region, i),
			region: region,
		})
	}
	return addresses
}

func removeExpectedAddress(s []expectedAddress, i int) []expectedAddress {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
