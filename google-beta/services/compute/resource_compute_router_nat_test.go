// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccComputeRouterNat_basic(t *testing.T) {
	t.Parallel()

	project := envvar.GetTestProjectFromEnv()
	region := envvar.GetTestRegionFromEnv()

	testId := acctest.RandString(t, 10)
	routerName := fmt.Sprintf("tf-test-router-nat-%s", testId)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRouterNatDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRouterNatBasic(routerName),
			},
			{
				// implicitly full ImportStateId
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportStateId:     fmt.Sprintf("%s/%s/%s/%s", project, region, routerName, routerName),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportStateId:     fmt.Sprintf("%s/%s/%s", region, routerName, routerName),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportStateId:     fmt.Sprintf("%s/%s", routerName, routerName),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRouterNatKeepRouter(routerName),
				Check: testAccCheckComputeRouterNatDelete(
					t, "google_compute_router_nat.foobar"),
			},
		},
	})
}

func TestAccComputeRouterNat_update(t *testing.T) {
	t.Parallel()

	testId := acctest.RandString(t, 10)
	routerName := fmt.Sprintf("tf-test-router-nat-%s", testId)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRouterNatDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRouterNatBasicBeforeUpdate(routerName),
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRouterNatUpdated(routerName),
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRouterNatUpdateToNatIPsId(routerName),
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRouterNatUpdateToNatIPsName(routerName),
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRouterNatBasicBeforeUpdate(routerName),
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRouterNat_removeLogConfig(t *testing.T) {
	t.Parallel()

	testId := acctest.RandString(t, 10)
	routerName := fmt.Sprintf("tf-test-router-nat-%s", testId)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRouterNatDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRouterNatLogConfig(routerName),
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRouterNatLogConfigRemoved(routerName),
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRouterNat_withManualIpAndSubnetConfiguration(t *testing.T) {
	t.Parallel()

	testId := acctest.RandString(t, 10)
	routerName := fmt.Sprintf("tf-test-router-nat-%s", testId)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRouterNatDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRouterNatWithManualIpAndSubnetConfiguration(routerName),
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRouterNat_withPortAllocationMethods(t *testing.T) {
	t.Parallel()

	testId := acctest.RandString(t, 10)
	routerName := fmt.Sprintf("tf-test-router-nat-%s", testId)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRouterNatDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRouterNatWithAllocationMethod(routerName, false, true),
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRouterNatWithAllocationMethod(routerName, true, false),
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRouterNatWithAllocationMethod(routerName, false, false),
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRouterNatWithAllocationMethod(routerName, true, false),
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRouterNatWithAllocationMethod(routerName, false, true),
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRouterNatWithAllocationMethodWithParameters(routerName, false, true, 256, 8192),
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRouterNat_withNatIpsAndDrainNatIps(t *testing.T) {
	t.Parallel()

	testId := acctest.RandString(t, 10)
	routerName := fmt.Sprintf("tf-test-router-nat-%s", testId)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRouterNatDestroyProducer(t),
		Steps: []resource.TestStep{
			// (ERROR): Creation with drain nat IPs should fail
			{
				Config:      testAccComputeRouterNatWithOneDrainOneRemovedNatIps(routerName),
				ExpectError: regexp.MustCompile("New RouterNat cannot have drain_nat_ips"),
			},
			// Create NAT with three nat IPs
			{
				Config: testAccComputeRouterNatWithNatIps(routerName),
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// (ERROR) - Should not allow draining IPs still in natIps
			{
				Config:      testAccComputeRouterNatWithInvalidDrainNatIpsStillInNatIps(routerName),
				ExpectError: regexp.MustCompile("cannot be drained if still set in nat_ips"),
			},
			// natIps #1, #2, #3--> natIp #2, drainNatIp #3
			{
				Config: testAccComputeRouterNatWithOneDrainOneRemovedNatIps(routerName),
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// (ERROR): Should not be able to drain previously removed natIps (#1)
			{
				Config:      testAccComputeRouterNatWithInvalidDrainMissingNatIp(routerName),
				ExpectError: regexp.MustCompile("was not previously set in nat_ips"),
			},
		},
	})
}

func TestAccComputeRouterNat_withNatRules(t *testing.T) {
	t.Parallel()

	testId := acctest.RandString(t, 10)
	routerName := fmt.Sprintf("tf-test-router-nat-%s", testId)
	ruleDescription := acctest.RandString(t, 10)
	ruleDescriptionUpdate := acctest.RandString(t, 10)
	match := "inIpRange(destination.ip, '1.1.0.0/16') || inIpRange(destination.ip, '2.2.0.0/16')"
	matchUpdate := "destination.ip == '1.1.0.1' || destination.ip == '8.8.8.8'"

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRouterNatDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRouterNatRulesBasic_omitRules(routerName),
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRouterNatRulesBasic(routerName, 0, ruleDescription, match),
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRouterNatRulesBasic(routerName, 65000, ruleDescription, match),
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRouterNatRulesBasic(routerName, 100, ruleDescription, match),
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRouterNatRulesBasic(routerName, 100, ruleDescriptionUpdate, match),
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRouterNatRulesBasic(routerName, 100, ruleDescriptionUpdate, matchUpdate),
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRouterNatRulesWithSourceActiveAndDrainIps(routerName, 100, ruleDescriptionUpdate, matchUpdate),
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRouterNatRulesWithDrainIps(routerName, 100, ruleDescriptionUpdate, matchUpdate),
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRouterNatMultiRules(routerName),
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRouterNatRulesBasic_omitAction(routerName, 100, ruleDescriptionUpdate, matchUpdate),
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRouterNatRulesBasic_omitDescription(routerName, 100, matchUpdate),
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRouterNatMultiRulesWithIpId(routerName),
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRouterNatRulesBasic_omitRules(routerName),
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRouterNat_withEndpointTypes(t *testing.T) {
	t.Parallel()

	testId := acctest.RandString(t, 10)
	routerName := fmt.Sprintf("tf-test-router-nat-%s", testId)
	testResourceName := "google_compute_router_nat.foobar"

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRouterNatDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRouterNatBasic(routerName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(testResourceName, "endpoint_types.0", "ENDPOINT_TYPE_VM"),
				),
			},
			{
				ResourceName:      testResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRouterNatUpdateEndpointType(routerName, "ENDPOINT_TYPE_SWG"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(testResourceName, "endpoint_types.0", "ENDPOINT_TYPE_SWG"),
				),
			},
			{
				ResourceName:      testResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRouterNatUpdateEndpointType(routerName, "ENDPOINT_TYPE_VM"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(testResourceName, "endpoint_types.0", "ENDPOINT_TYPE_VM"),
				),
			},
			{
				ResourceName:      testResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRouterNatUpdateEndpointType(routerName, "ENDPOINT_TYPE_MANAGED_PROXY_LB"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(testResourceName, "endpoint_types.0", "ENDPOINT_TYPE_MANAGED_PROXY_LB"),
				),
			},
			{
				ResourceName:      testResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRouterNat_withPrivateNat(t *testing.T) {
	t.Parallel()

	project := envvar.GetTestProjectFromEnv()
	region := envvar.GetTestRegionFromEnv()

	testId := acctest.RandString(t, 10)
	routerName := fmt.Sprintf("tf-test-router-nat-%s", testId)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRouterNatDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRouterNatPrivateType(routerName),
			},
			{
				// implicitly full ImportStateId
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportStateId:     fmt.Sprintf("%s/%s/%s/%s", project, region, routerName, routerName),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportStateId:     fmt.Sprintf("%s/%s/%s", region, routerName, routerName),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportStateId:     fmt.Sprintf("%s/%s", routerName, routerName),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRouterNatKeepRouter(routerName),
				Check: testAccCheckComputeRouterNatDelete(
					t, "google_compute_router_nat.foobar"),
			},
		},
	})
}

func TestAccComputeRouterNat_withPrivateNatAndRules(t *testing.T) {
	t.Parallel()

	testId := acctest.RandString(t, 10)
	routerName := fmt.Sprintf("tf-test-router-private-nat-%s", testId)
	hubName := fmt.Sprintf("%s-hub", routerName)
	pEnv := envvar.GetTestProjectFromEnv()
	ruleDescription := acctest.RandString(t, 10)
	match := fmt.Sprintf("nexthop.hub == '//networkconnectivity.googleapis.com/projects/%s/locations/global/hubs/%s'", pEnv, hubName)
	activeRangesNetworkOne := "google_compute_subnetwork.subnet1.self_link"
	drainRangesEmpty := ""
	activeRangesNetworkTwoAndThree := "google_compute_subnetwork.subnet2.self_link,google_compute_subnetwork.subnet3.self_link"
	activeRangesNetworkThreeAndFour := "google_compute_subnetwork.subnet3.self_link,google_compute_subnetwork.subnet4.self_link"
	drainRangesNetworkOne := "google_compute_subnetwork.subnet1.self_link"
	drainRangesNetworkOneAndTwo := "google_compute_subnetwork.subnet1.self_link,google_compute_subnetwork.subnet2.self_link"

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRouterNatDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRouterNatRulesBasic_privateNatOmitRules(routerName, hubName),
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRouterNatRulesBasic_privateNatWithRuleAndActiveDrainRange(routerName, hubName, 100, ruleDescription, match, activeRangesNetworkOne, drainRangesEmpty),
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRouterNatRulesBasic_privateNatWithRuleAndActiveDrainRange(routerName, hubName, 100, ruleDescription, match, activeRangesNetworkTwoAndThree, drainRangesNetworkOne),
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRouterNatRulesBasic_privateNatWithRuleAndActiveDrainRange(routerName, hubName, 100, ruleDescription, match, activeRangesNetworkThreeAndFour, drainRangesNetworkOneAndTwo),
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRouterNatRulesBasic_privateNatWithRuleAndActiveDrainRange(routerName, hubName, 100, ruleDescription, match, activeRangesNetworkOne, drainRangesEmpty),
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRouterNatRulesBasic_privateNatOmitRules(routerName, hubName),
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRouterNat_withPrivateNatAndEmptyAction(t *testing.T) {
	t.Parallel()

	testId := acctest.RandString(t, 10)
	routerName := fmt.Sprintf("tf-test-router-private-nat-%s", testId)
	hubName := fmt.Sprintf("%s-hub", routerName)
	pEnv := envvar.GetTestProjectFromEnv()
	ruleDescription := acctest.RandString(t, 10)
	match := fmt.Sprintf("nexthop.hub == '//networkconnectivity.googleapis.com/projects/%s/locations/global/hubs/%s'", pEnv, hubName)
	activeRangesNetworkOne := "google_compute_subnetwork.subnet1.self_link"
	drainRangesEmpty := ""

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRouterNatDestroyProducer(t),
		Steps: []resource.TestStep{
			// (ERROR): Creation with empty action should fail
			{
				Config:      testAccComputeRouterNatRulesBasic_privateNatWithRuleAndEmptyAction(routerName, hubName, 100, ruleDescription, match),
				ExpectError: regexp.MustCompile("The rule for PRIVATE nat type must contain an action with source_nat_active_ranges set"),
			},
			// Create NAT with action and active ranges set
			{
				Config: testAccComputeRouterNatRulesBasic_privateNatWithRuleAndActiveDrainRange(routerName, hubName, 100, ruleDescription, match, activeRangesNetworkOne, drainRangesEmpty),
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// (ERROR) - Updating the rule by removing the action should fail
			{
				Config:      testAccComputeRouterNatRulesBasic_privateNatWithRuleAndEmptyAction(routerName, hubName, 100, ruleDescription, match),
				ExpectError: regexp.MustCompile("The rule for PRIVATE nat type must contain an action with source_nat_active_ranges set"),
			},
		},
	})
}

func TestAccComputeRouterNat_withPrivateNatAndEmptyActionActiveRanges(t *testing.T) {
	t.Parallel()

	testId := acctest.RandString(t, 10)
	routerName := fmt.Sprintf("tf-test-router-private-nat-%s", testId)
	hubName := fmt.Sprintf("%s-hub", routerName)
	pEnv := envvar.GetTestProjectFromEnv()
	ruleDescription := acctest.RandString(t, 10)
	match := fmt.Sprintf("nexthop.hub == '//networkconnectivity.googleapis.com/projects/%s/locations/global/hubs/%s'", pEnv, hubName)
	activeRangesNetworkOne := "google_compute_subnetwork.subnet1.self_link"
	drainRangesEmpty := ""

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRouterNatDestroyProducer(t),
		Steps: []resource.TestStep{
			// (ERROR): Creation with empty action active ranges should fail
			{
				Config:      testAccComputeRouterNatRulesBasic_privateNatWithRuleAndEmptyActionActiveRanges(routerName, hubName, 100, ruleDescription, match),
				ExpectError: regexp.MustCompile("The rule for PRIVATE nat type must contain an action with source_nat_active_ranges set"),
			},
			// Create NAT with action and active ranges set
			{
				Config: testAccComputeRouterNatRulesBasic_privateNatWithRuleAndActiveDrainRange(routerName, hubName, 100, ruleDescription, match, activeRangesNetworkOne, drainRangesEmpty),
			},
			{
				ResourceName:      "google_compute_router_nat.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// (ERROR) - Updating the rule by erasing the action active ranges should fail
			{
				Config:      testAccComputeRouterNatRulesBasic_privateNatWithRuleAndEmptyActionActiveRanges(routerName, hubName, 100, ruleDescription, match),
				ExpectError: regexp.MustCompile("The rule for PRIVATE nat type must contain an action with source_nat_active_ranges set"),
			},
		},
	})
}

func testAccCheckComputeRouterNatDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		config := acctest.GoogleProviderConfig(t)

		routersService := config.NewComputeClient(config.UserAgent).Routers

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_router" {
				continue
			}

			project, err := acctest.GetTestProject(rs.Primary, config)
			if err != nil {
				return err
			}

			region, err := acctest.GetTestRegion(rs.Primary, config)
			if err != nil {
				return err
			}

			routerName := rs.Primary.Attributes["router"]

			_, err = routersService.Get(project, region, routerName).Do()

			if err == nil {
				return fmt.Errorf("Error, Router %s in region %s still exists", routerName, region)
			}
		}

		return nil
	}
}

func testAccCheckComputeRouterNatDelete(t *testing.T, n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config := acctest.GoogleProviderConfig(t)

		routersService := config.NewComputeClient(config.UserAgent).Routers

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_router_nat" {
				continue
			}

			project, err := acctest.GetTestProject(rs.Primary, config)
			if err != nil {
				return err
			}

			region, err := acctest.GetTestRegion(rs.Primary, config)
			if err != nil {
				return err
			}

			name := rs.Primary.Attributes["name"]
			routerName := rs.Primary.Attributes["router"]

			router, err := routersService.Get(project, region, routerName).Do()

			if err != nil {
				return fmt.Errorf("Error Reading Router %s: %s", routerName, err)
			}

			nats := router.Nats
			for _, nat := range nats {
				if nat.Name == name {
					return fmt.Errorf("Nat %s still exists on router %s/%s", name, region, router.Name)
				}
			}
		}

		return nil
	}
}

func testAccComputeRouterNatBasic(routerName string) string {
	return fmt.Sprintf(`
resource "google_compute_network" "foobar" {
  name = "%s-net"
}

resource "google_compute_subnetwork" "foobar" {
  name          = "%s-subnet"
  network       = google_compute_network.foobar.self_link
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
}

resource "google_compute_router" "foobar" {
  name    = "%s"
  region  = google_compute_subnetwork.foobar.region
  network = google_compute_network.foobar.self_link
}

resource "google_compute_router_nat" "foobar" {
  name                               = "%s"
  router                             = google_compute_router.foobar.name
  region                             = google_compute_router.foobar.region
  nat_ip_allocate_option             = "AUTO_ONLY"
  source_subnetwork_ip_ranges_to_nat = "ALL_SUBNETWORKS_ALL_IP_RANGES"
  log_config {
    enable = true
    filter = "ERRORS_ONLY"
  }
}
`, routerName, routerName, routerName, routerName)
}

// Like basic but with extra resources
func testAccComputeRouterNatBasicBeforeUpdate(routerName string) string {
	return fmt.Sprintf(`
resource "google_compute_router" "foobar" {
  name    = "%s"
  region  = google_compute_subnetwork.foobar.region
  network = google_compute_network.foobar.self_link
}

resource "google_compute_network" "foobar" {
  name = "%s-net"
}

resource "google_compute_subnetwork" "foobar" {
  name          = "%s-subnet"
  network       = google_compute_network.foobar.self_link
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
}

resource "google_compute_address" "foobar" {
  name   = "%s-addr"
  region = google_compute_subnetwork.foobar.region
}

resource "google_compute_router_nat" "foobar" {
  name                               = "%s"
  router                             = google_compute_router.foobar.name
  region                             = google_compute_router.foobar.region
  nat_ip_allocate_option             = "AUTO_ONLY"
  source_subnetwork_ip_ranges_to_nat = "ALL_SUBNETWORKS_ALL_IP_RANGES"

  log_config {
    enable = true
    filter = "ERRORS_ONLY"
  }
}
`, routerName, routerName, routerName, routerName, routerName)
}

func testAccComputeRouterNatUpdated(routerName string) string {
	return fmt.Sprintf(`
resource "google_compute_router" "foobar" {
  name    = "%s"
  region  = google_compute_subnetwork.foobar.region
  network = google_compute_network.foobar.self_link
}

resource "google_compute_network" "foobar" {
  name = "%s-net"
}

resource "google_compute_subnetwork" "foobar" {
  name          = "%s-subnet"
  network       = google_compute_network.foobar.self_link
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
}

resource "google_compute_address" "foobar" {
  name   = "%s-addr"
  region = google_compute_subnetwork.foobar.region
}

resource "google_compute_router_nat" "foobar" {
  name   = "%s"
  router = google_compute_router.foobar.name
  region = google_compute_router.foobar.region

  nat_ip_allocate_option = "MANUAL_ONLY"
  nat_ips                = [google_compute_address.foobar.self_link]

  source_subnetwork_ip_ranges_to_nat = "LIST_OF_SUBNETWORKS"

  subnetwork {
    name                    = google_compute_subnetwork.foobar.self_link
    source_ip_ranges_to_nat = ["ALL_IP_RANGES"]
  }

  udp_idle_timeout_sec             = 60
  icmp_idle_timeout_sec            = 60
  tcp_established_idle_timeout_sec = 1600
  tcp_transitory_idle_timeout_sec  = 60
  tcp_time_wait_timeout_sec        = 60

  log_config {
    enable = true
    filter = "TRANSLATIONS_ONLY"
  }
}
`, routerName, routerName, routerName, routerName, routerName)
}

func testAccComputeRouterNatUpdateEndpointType(routerName string, endpointType string) string {
	return fmt.Sprintf(`
resource "google_compute_network" "foobar" {
  name = "%[1]s-net"
}

resource "google_compute_subnetwork" "foobar" {
  name          = "%[1]s-subnet"
  network       = google_compute_network.foobar.self_link
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
}

resource "google_compute_router" "foobar" {
  name    = "%[1]s"
  region  = google_compute_subnetwork.foobar.region
  network = google_compute_network.foobar.self_link
}

resource "google_compute_router_nat" "foobar" {
  name                               = "%[1]s"
  router                             = google_compute_router.foobar.name
  region                             = google_compute_router.foobar.region
  nat_ip_allocate_option             = "AUTO_ONLY"
  source_subnetwork_ip_ranges_to_nat = "ALL_SUBNETWORKS_ALL_IP_RANGES"
  endpoint_types                     = [ "%[2]s" ]
  log_config {
    enable = true
    filter = "ERRORS_ONLY"
  }
}
`, routerName, endpointType)
}

func testAccComputeRouterNatUpdateToNatIPsId(routerName string) string {
	return fmt.Sprintf(`
resource "google_compute_router" "foobar" {
name    = "%s"
region  = google_compute_subnetwork.foobar.region
network = google_compute_network.foobar.self_link
}

resource "google_compute_network" "foobar" {
name = "%s-net"
}
resource "google_compute_subnetwork" "foobar" {
name          = "%s-subnet"
network       = google_compute_network.foobar.self_link
ip_cidr_range = "10.0.0.0/16"
region        = "us-central1"
}

resource "google_compute_address" "foobar" {
name   = "%s-addr"
region = google_compute_subnetwork.foobar.region
}

resource "google_compute_router_nat" "foobar" {
  name   = "%s"
  router = google_compute_router.foobar.name
  region = google_compute_router.foobar.region

  nat_ip_allocate_option = "MANUAL_ONLY"
  nat_ips                = [google_compute_address.foobar.id]

  source_subnetwork_ip_ranges_to_nat = "LIST_OF_SUBNETWORKS"

  subnetwork {
    name                    = google_compute_subnetwork.foobar.self_link
    source_ip_ranges_to_nat = ["ALL_IP_RANGES"]
  }

  udp_idle_timeout_sec             = 60
  icmp_idle_timeout_sec            = 60
  tcp_established_idle_timeout_sec = 1600
  tcp_transitory_idle_timeout_sec  = 60
  tcp_time_wait_timeout_sec        = 60

  log_config {
    enable = true
    filter = "TRANSLATIONS_ONLY"
  }
}
`, routerName, routerName, routerName, routerName, routerName)
}

func testAccComputeRouterNatUpdateToNatIPsName(routerName string) string {
	return fmt.Sprintf(`
resource "google_compute_router" "foobar" {
name    = "%s"
region  = google_compute_subnetwork.foobar.region
network = google_compute_network.foobar.self_link
}

resource "google_compute_network" "foobar" {
name = "%s-net"
}
resource "google_compute_subnetwork" "foobar" {
name          = "%s-subnet"
network       = google_compute_network.foobar.self_link
ip_cidr_range = "10.0.0.0/16"
region        = "us-central1"
}

resource "google_compute_address" "foobar" {
name   = "%s-addr"
region = google_compute_subnetwork.foobar.region
}

resource "google_compute_router_nat" "foobar" {
  name   = "%s"
  router = google_compute_router.foobar.name
  region = google_compute_router.foobar.region

  nat_ip_allocate_option = "MANUAL_ONLY"
  nat_ips                = [google_compute_address.foobar.name]

  source_subnetwork_ip_ranges_to_nat = "LIST_OF_SUBNETWORKS"

  subnetwork {
    name                    = google_compute_subnetwork.foobar.self_link
    source_ip_ranges_to_nat = ["ALL_IP_RANGES"]
  }

  udp_idle_timeout_sec             = 60
  icmp_idle_timeout_sec            = 60
  tcp_established_idle_timeout_sec = 1600
  tcp_transitory_idle_timeout_sec  = 60
  tcp_time_wait_timeout_sec        = 60

  log_config {
    enable = true
    filter = "TRANSLATIONS_ONLY"
  }
}
`, routerName, routerName, routerName, routerName, routerName)
}

func testAccComputeRouterNatWithManualIpAndSubnetConfiguration(routerName string) string {
	return fmt.Sprintf(`
resource "google_compute_network" "foobar" {
  name                    = "%s-net"
  auto_create_subnetworks = "false"
}

resource "google_compute_subnetwork" "foobar" {
  name          = "%s-subnet"
  network       = google_compute_network.foobar.self_link
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
}

resource "google_compute_address" "foobar" {
  name   = "%s-router-nat-addr"
  region = google_compute_subnetwork.foobar.region
}

resource "google_compute_router" "foobar" {
  name    = "%s"
  region  = google_compute_subnetwork.foobar.region
  network = google_compute_network.foobar.self_link
  bgp {
    asn = 64514
  }
}

resource "google_compute_router_nat" "foobar" {
  name                               = "%s"
  router                             = google_compute_router.foobar.name
  region                             = google_compute_router.foobar.region
  nat_ip_allocate_option             = "MANUAL_ONLY"
  nat_ips                            = [google_compute_address.foobar.self_link]
  source_subnetwork_ip_ranges_to_nat = "LIST_OF_SUBNETWORKS"
  subnetwork {
    name                    = google_compute_subnetwork.foobar.name
    source_ip_ranges_to_nat = ["ALL_IP_RANGES"]
  }
}
`, routerName, routerName, routerName, routerName, routerName)
}

func testAccComputeRouterNatWithAllocationMethod(routerName string, enableEndpointIndependentMapping, enableDynamicPortAllocation bool) string {
	return fmt.Sprintf(`
resource "google_compute_network" "foobar" {
  name                    = "%s-net"
  auto_create_subnetworks = "false"
}

resource "google_compute_subnetwork" "foobar" {
  name          = "%s-subnet"
  network       = google_compute_network.foobar.self_link
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
}

resource "google_compute_address" "foobar" {
  name   = "%s-router-nat-addr"
  region = google_compute_subnetwork.foobar.region
}

resource "google_compute_router" "foobar" {
  name    = "%s"
  region  = google_compute_subnetwork.foobar.region
  network = google_compute_network.foobar.self_link
  bgp {
    asn = 64514
  }
}

resource "google_compute_router_nat" "foobar" {
  name                               = "%s"
  router                             = google_compute_router.foobar.name
  region                             = google_compute_router.foobar.region
  nat_ip_allocate_option             = "MANUAL_ONLY"
  nat_ips                            = [google_compute_address.foobar.self_link]
  source_subnetwork_ip_ranges_to_nat = "LIST_OF_SUBNETWORKS"
  subnetwork {
    name                    = google_compute_subnetwork.foobar.name
    source_ip_ranges_to_nat = ["ALL_IP_RANGES"]
  }
  enable_endpoint_independent_mapping = %t
  enable_dynamic_port_allocation = %t
}
`, routerName, routerName, routerName, routerName, routerName, enableEndpointIndependentMapping, enableDynamicPortAllocation)
}

func testAccComputeRouterNatWithAllocationMethodWithParameters(routerName string, enableEndpointIndependentMapping, enableDynamicPortAllocation bool, minPortsPerVm, maxPortsPerVm uint32) string {
	return fmt.Sprintf(`
resource "google_compute_network" "foobar" {
  name                    = "%s-net"
  auto_create_subnetworks = "false"
}

resource "google_compute_subnetwork" "foobar" {
  name          = "%s-subnet"
  network       = google_compute_network.foobar.self_link
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
}

resource "google_compute_address" "foobar" {
  name   = "%s-router-nat-addr"
  region = google_compute_subnetwork.foobar.region
}

resource "google_compute_router" "foobar" {
  name    = "%s"
  region  = google_compute_subnetwork.foobar.region
  network = google_compute_network.foobar.self_link
  bgp {
    asn = 64514
  }
}

resource "google_compute_router_nat" "foobar" {
  name                               = "%s"
  router                             = google_compute_router.foobar.name
  region                             = google_compute_router.foobar.region
  nat_ip_allocate_option             = "MANUAL_ONLY"
  nat_ips                            = [google_compute_address.foobar.self_link]
  source_subnetwork_ip_ranges_to_nat = "LIST_OF_SUBNETWORKS"
  subnetwork {
    name                    = google_compute_subnetwork.foobar.name
    source_ip_ranges_to_nat = ["ALL_IP_RANGES"]
  }
  enable_endpoint_independent_mapping = %t
  enable_dynamic_port_allocation = %t
  min_ports_per_vm = %d
  max_ports_per_vm = %d
}
`, routerName, routerName, routerName, routerName, routerName, enableEndpointIndependentMapping, enableDynamicPortAllocation, minPortsPerVm, maxPortsPerVm)
}

func testAccComputeRouterNatBaseResourcesWithNatIps(routerName string) string {
	return fmt.Sprintf(`
resource "google_compute_network" "foobar" {
  name                    = "%s-net"
  auto_create_subnetworks = "false"
}

resource "google_compute_subnetwork" "foobar" {
  name          = "%s-subnet"
  network       = google_compute_network.foobar.self_link
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
}

resource "google_compute_address" "addr1" {
  name   = "%s-addr1"
  region = google_compute_subnetwork.foobar.region
}

resource "google_compute_address" "addr2" {
  name   = "%s-addr2"
  region = google_compute_subnetwork.foobar.region
}

resource "google_compute_address" "addr3" {
  name   = "%s-addr3"
  region = google_compute_subnetwork.foobar.region
}

resource "google_compute_address" "addr4" {
  name   = "%s-addr4"
  region = google_compute_subnetwork.foobar.region
}

resource "google_compute_router" "foobar" {
  name     = "%s"
  region   = google_compute_subnetwork.foobar.region
  network  = google_compute_network.foobar.self_link
}
`, routerName, routerName, routerName, routerName, routerName, routerName, routerName)
}

func testAccComputeRouterNatWithNatIps(routerName string) string {
	return fmt.Sprintf(`
%s

resource "google_compute_router_nat" "foobar" {
  name     = "%s"
  router   = google_compute_router.foobar.name
  region   = google_compute_router.foobar.region

  nat_ip_allocate_option = "MANUAL_ONLY"
  nat_ips = [
    google_compute_address.addr1.self_link,
    google_compute_address.addr2.self_link,
    google_compute_address.addr3.self_link,
  ]

  source_subnetwork_ip_ranges_to_nat = "LIST_OF_SUBNETWORKS"
  subnetwork {
    name                    = google_compute_subnetwork.foobar.self_link
    source_ip_ranges_to_nat = ["ALL_IP_RANGES"]
  }
}
`, testAccComputeRouterNatBaseResourcesWithNatIps(routerName), routerName)
}

func testAccComputeRouterNatWithOneDrainOneRemovedNatIps(routerName string) string {
	return fmt.Sprintf(`
%s

resource "google_compute_router_nat" "foobar" {
  name     = "%s"
  router   = google_compute_router.foobar.name
  region   = google_compute_router.foobar.region

  source_subnetwork_ip_ranges_to_nat = "LIST_OF_SUBNETWORKS"
  subnetwork {
    name                    = google_compute_subnetwork.foobar.self_link
    source_ip_ranges_to_nat = ["ALL_IP_RANGES"]
  }

  nat_ip_allocate_option = "MANUAL_ONLY"
  nat_ips = [
    google_compute_address.addr2.self_link,
  ]

  drain_nat_ips = [
    google_compute_address.addr3.self_link,
  ]
}
`, testAccComputeRouterNatBaseResourcesWithNatIps(routerName), routerName)
}

func testAccComputeRouterNatWithInvalidDrainMissingNatIp(routerName string) string {
	return fmt.Sprintf(`
%s

resource "google_compute_router_nat" "foobar" {
  name     = "%s"
  router   = google_compute_router.foobar.name
  region   = google_compute_router.foobar.region

  source_subnetwork_ip_ranges_to_nat = "LIST_OF_SUBNETWORKS"
  subnetwork {
    name                    = google_compute_subnetwork.foobar.self_link
    source_ip_ranges_to_nat = ["ALL_IP_RANGES"]
  }

  nat_ip_allocate_option = "MANUAL_ONLY"
  nat_ips = [
    google_compute_address.addr2.self_link,
  ]

  drain_nat_ips = [
    google_compute_address.addr1.self_link,
    google_compute_address.addr3.self_link,
  ]
}
`, testAccComputeRouterNatBaseResourcesWithNatIps(routerName), routerName)
}

func testAccComputeRouterNatWithInvalidDrainNatIpsStillInNatIps(routerName string) string {
	return fmt.Sprintf(`
%s

resource "google_compute_router_nat" "foobar" {
  name     = "%s"
  router   = google_compute_router.foobar.name
  region   = google_compute_router.foobar.region

  source_subnetwork_ip_ranges_to_nat = "LIST_OF_SUBNETWORKS"
  subnetwork {
    name                    = google_compute_subnetwork.foobar.self_link
    source_ip_ranges_to_nat = ["ALL_IP_RANGES"]
  }

  nat_ip_allocate_option = "MANUAL_ONLY"
  nat_ips = [
    google_compute_address.addr1.self_link,
    google_compute_address.addr2.self_link,
    google_compute_address.addr3.self_link,
  ]

  drain_nat_ips = [
    google_compute_address.addr3.self_link,
  ]
}
`, testAccComputeRouterNatBaseResourcesWithNatIps(routerName), routerName)
}

func testAccComputeRouterNatRulesBasic_omitRules(routerName string) string {
	return fmt.Sprintf(`
%s

resource "google_compute_router_nat" "foobar" {
  name     = "%s"
  router   = google_compute_router.foobar.name
  region   = google_compute_router.foobar.region

  nat_ip_allocate_option = "MANUAL_ONLY"
  nat_ips                = [google_compute_address.addr1.self_link]


  source_subnetwork_ip_ranges_to_nat = "LIST_OF_SUBNETWORKS"
  subnetwork {
    name                    = google_compute_subnetwork.foobar.id
    source_ip_ranges_to_nat = ["ALL_IP_RANGES"]
  }

  enable_endpoint_independent_mapping = false
}
`, testAccComputeRouterNatBaseResourcesWithNatIps(routerName), routerName)
}

func testAccComputeRouterNatRulesBasic_omitAction(routerName string, ruleNumber int, ruleDescription string, ruleMatch string) string {
	return fmt.Sprintf(`
%s

resource "google_compute_router_nat" "foobar" {
  name     = "%s"
  router   = google_compute_router.foobar.name
  region   = google_compute_router.foobar.region

  nat_ip_allocate_option = "MANUAL_ONLY"
  nat_ips                = [google_compute_address.addr1.self_link]


  source_subnetwork_ip_ranges_to_nat = "LIST_OF_SUBNETWORKS"
  subnetwork {
    name                    = google_compute_subnetwork.foobar.id
    source_ip_ranges_to_nat = ["ALL_IP_RANGES"]
  }

  rules {
    rule_number = %d
    description = "%s"
    match       = "%s"
  }

  enable_endpoint_independent_mapping = false
}
`, testAccComputeRouterNatBaseResourcesWithNatIps(routerName), routerName, ruleNumber, ruleDescription, ruleMatch)
}

func testAccComputeRouterNatRulesBasic_omitDescription(routerName string, ruleNumber int, ruleMatch string) string {
	return fmt.Sprintf(`
%s

resource "google_compute_router_nat" "foobar" {
  name     = "%s"
  router   = google_compute_router.foobar.name
  region   = google_compute_router.foobar.region

  nat_ip_allocate_option = "MANUAL_ONLY"
  nat_ips                = [google_compute_address.addr1.self_link]


  source_subnetwork_ip_ranges_to_nat = "LIST_OF_SUBNETWORKS"
  subnetwork {
    name                    = google_compute_subnetwork.foobar.id
    source_ip_ranges_to_nat = ["ALL_IP_RANGES"]
  }

  rules {
    rule_number = %d
    match       = "%s"
    action {
      source_nat_active_ips = [google_compute_address.addr2.self_link, google_compute_address.addr3.self_link]
    }
  }

  enable_endpoint_independent_mapping = false
}
`, testAccComputeRouterNatBaseResourcesWithNatIps(routerName), routerName, ruleNumber, ruleMatch)
}

func testAccComputeRouterNatRulesBasic(routerName string, ruleNumber int, ruleDescription string, ruleMatch string) string {
	return fmt.Sprintf(`
%s

resource "google_compute_router_nat" "foobar" {
  name     = "%s"
  router   = google_compute_router.foobar.name
  region   = google_compute_router.foobar.region

  nat_ip_allocate_option = "MANUAL_ONLY"
  nat_ips                = [google_compute_address.addr1.self_link]


  source_subnetwork_ip_ranges_to_nat = "LIST_OF_SUBNETWORKS"
  subnetwork {
    name                    = google_compute_subnetwork.foobar.id
    source_ip_ranges_to_nat = ["ALL_IP_RANGES"]
  }

  rules {
    rule_number = %d
    description = "%s"
    match       = "%s"
    action {
      source_nat_active_ips = [google_compute_address.addr2.self_link, google_compute_address.addr3.self_link]
    }
  }

  enable_endpoint_independent_mapping = false
}
`, testAccComputeRouterNatBaseResourcesWithNatIps(routerName), routerName, ruleNumber, ruleDescription, ruleMatch)
}

func testAccComputeRouterNatRulesWithSourceActiveAndDrainIps(routerName string, ruleNumber int, ruleDescription string, ruleMatch string) string {
	return fmt.Sprintf(`
%s

resource "google_compute_router_nat" "foobar" {
  name     = "%s"
  router   = google_compute_router.foobar.name
  region   = google_compute_router.foobar.region

  nat_ip_allocate_option = "MANUAL_ONLY"
  nat_ips                = [google_compute_address.addr1.self_link]

  source_subnetwork_ip_ranges_to_nat = "LIST_OF_SUBNETWORKS"
  subnetwork {
    name                    = google_compute_subnetwork.foobar.id
    source_ip_ranges_to_nat = ["ALL_IP_RANGES"]
  }

  rules {
    rule_number = %d
    description = "%s"
    match       = "%s"
    action {
      source_nat_active_ips = [google_compute_address.addr2.self_link]
      source_nat_drain_ips = [google_compute_address.addr3.self_link]
    }
  }

  enable_endpoint_independent_mapping = false
}
`, testAccComputeRouterNatBaseResourcesWithNatIps(routerName), routerName, ruleNumber, ruleDescription, ruleMatch)
}

func testAccComputeRouterNatRulesWithDrainIps(routerName string, ruleNumber int, ruleDescription string, ruleMatch string) string {
	return fmt.Sprintf(`
%s

resource "google_compute_router_nat" "foobar" {
  name     = "%s"
  router   = google_compute_router.foobar.name
  region   = google_compute_router.foobar.region

  nat_ip_allocate_option = "MANUAL_ONLY"
  nat_ips                = [google_compute_address.addr1.self_link]

  source_subnetwork_ip_ranges_to_nat = "LIST_OF_SUBNETWORKS"
  subnetwork {
    name                    = google_compute_subnetwork.foobar.id
    source_ip_ranges_to_nat = ["ALL_IP_RANGES"]
  }

  rules {
    rule_number = %d
    description = "%s"
    match       = "%s"
    action {
      source_nat_drain_ips = [google_compute_address.addr2.self_link]
    }
  }

  enable_endpoint_independent_mapping = false
}
`, testAccComputeRouterNatBaseResourcesWithNatIps(routerName), routerName, ruleNumber, ruleDescription, ruleMatch)
}

func testAccComputeRouterNatMultiRules(routerName string) string {
	return fmt.Sprintf(`
%s

resource "google_compute_router_nat" "foobar" {
  name     = "%s"
  router   = google_compute_router.foobar.name
  region   = google_compute_router.foobar.region

  nat_ip_allocate_option = "MANUAL_ONLY"
  nat_ips                = [google_compute_address.addr1.self_link]


  source_subnetwork_ip_ranges_to_nat = "LIST_OF_SUBNETWORKS"
  subnetwork {
    name                    = google_compute_subnetwork.foobar.id
    source_ip_ranges_to_nat = ["ALL_IP_RANGES"]
  }

  rules {
    rule_number = 100
    description = "a"
    match       = "destination.ip == '1.1.1.1' || destination.ip == '2.2.2.2'"
    action {
      source_nat_active_ips = [google_compute_address.addr2.self_link]
    }
  }

  rules {
    rule_number = 5000
    description = "b"
    match       = "destination.ip == '3.3.3.3' || destination.ip == '4.4.4.4'"
    action {
      source_nat_active_ips = [google_compute_address.addr3.self_link]
    }
  }

  rules {
    rule_number = 300
    description = "c"
    match       = "destination.ip == '5.5.5.5' || destination.ip == '8.8.8.8'"
    action {
      source_nat_active_ips = [google_compute_address.addr4.self_link]
    }
  }

  enable_endpoint_independent_mapping = false
}
`, testAccComputeRouterNatBaseResourcesWithNatIps(routerName), routerName)
}

func testAccComputeRouterNatMultiRulesWithIpId(routerName string) string {
	return fmt.Sprintf(`
%s

resource "google_compute_router_nat" "foobar" {
  name     = "%s"
  router   = google_compute_router.foobar.name
  region   = google_compute_router.foobar.region

  nat_ip_allocate_option = "MANUAL_ONLY"
  nat_ips                = [google_compute_address.addr1.id]


  source_subnetwork_ip_ranges_to_nat = "LIST_OF_SUBNETWORKS"
  subnetwork {
    name                    = google_compute_subnetwork.foobar.id
    source_ip_ranges_to_nat = ["ALL_IP_RANGES"]
  }

  rules {
    rule_number = 100
    description = "a"
    match       = "destination.ip == '1.1.1.1' || destination.ip == '2.2.2.2'"
    action {
      source_nat_active_ips = [google_compute_address.addr2.id]
    }
  }

  rules {
    rule_number = 5000
    description = "b"
    match       = "destination.ip == '3.3.3.3' || destination.ip == '4.4.4.4'"
    action {
      source_nat_active_ips = [google_compute_address.addr3.id]
    }
  }

  rules {
    rule_number = 300
    description = "c"
    match       = "destination.ip == '5.5.5.5' || destination.ip == '8.8.8.8'"
    action {
      source_nat_active_ips = [google_compute_address.addr4.id]
    }
  }

  enable_endpoint_independent_mapping = false
}
`, testAccComputeRouterNatBaseResourcesWithNatIps(routerName), routerName)
}

func testAccComputeRouterNatKeepRouter(routerName string) string {
	return fmt.Sprintf(`
resource "google_compute_network" "foobar" {
  name                    = "%s"
  auto_create_subnetworks = "false"
}

resource "google_compute_subnetwork" "foobar" {
  name          = "%s"
  network       = google_compute_network.foobar.self_link
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
}

resource "google_compute_router" "foobar" {
  name    = "%s"
  region  = google_compute_subnetwork.foobar.region
  network = google_compute_network.foobar.self_link
}
`, routerName, routerName, routerName)
}

func testAccComputeRouterNatLogConfig(routerName string) string {
	return fmt.Sprintf(`
resource "google_compute_network" "foobar" {
  name = "%s-net"
}

resource "google_compute_subnetwork" "foobar" {
  name          = "%s-subnet"
  network       = google_compute_network.foobar.self_link
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
}

resource "google_compute_router" "foobar" {
  name    = "%s"
  region  = google_compute_subnetwork.foobar.region
  network = google_compute_network.foobar.self_link
}

resource "google_compute_router_nat" "foobar" {
  name                               = "%s"
  router                             = google_compute_router.foobar.name
  region                             = google_compute_router.foobar.region
  nat_ip_allocate_option             = "AUTO_ONLY"
  source_subnetwork_ip_ranges_to_nat = "ALL_SUBNETWORKS_ALL_IP_RANGES"
  log_config {
    enable = false
    filter = "ALL"
  }
}
`, routerName, routerName, routerName, routerName)
}

func testAccComputeRouterNatLogConfigRemoved(routerName string) string {
	return fmt.Sprintf(`
resource "google_compute_network" "foobar" {
  name = "%s-net"
}

resource "google_compute_subnetwork" "foobar" {
  name          = "%s-subnet"
  network       = google_compute_network.foobar.self_link
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
}

resource "google_compute_router" "foobar" {
  name    = "%s"
  region  = google_compute_subnetwork.foobar.region
  network = google_compute_network.foobar.self_link
}

resource "google_compute_router_nat" "foobar" {
  name                               = "%s"
  router                             = google_compute_router.foobar.name
  region                             = google_compute_router.foobar.region
  nat_ip_allocate_option             = "AUTO_ONLY"
  source_subnetwork_ip_ranges_to_nat = "ALL_SUBNETWORKS_ALL_IP_RANGES"
}
`, routerName, routerName, routerName, routerName)
}

func testAccComputeRouterNatPrivateType(routerName string) string {
	return fmt.Sprintf(`
resource "google_compute_network" "foobar" {
  name                    = "%s-net"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "foobar" {
  name          = "%s-subnet"
  network       = google_compute_network.foobar.self_link
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
  purpose          = "PRIVATE_NAT"
}

resource "google_compute_router" "foobar" {
  name    = "%s"
  region  = google_compute_subnetwork.foobar.region
  network = google_compute_network.foobar.self_link
}

resource "google_compute_router_nat" "foobar" {
  name                               = "%s"
  router                             = google_compute_router.foobar.name
  region                             = google_compute_router.foobar.region
  source_subnetwork_ip_ranges_to_nat = "LIST_OF_SUBNETWORKS"
  type = "PRIVATE" 
  enable_dynamic_port_allocation = false
  enable_endpoint_independent_mapping = false
  min_ports_per_vm = 32
  subnetwork {
    name                    = google_compute_subnetwork.foobar.id
    source_ip_ranges_to_nat = ["ALL_IP_RANGES"]
  }
}
`, routerName, routerName, routerName, routerName)
}

func testAccComputeRouterNatBaseResourcesWithPrivateNatSubnetworks(routerName, hubName string) string {
	return fmt.Sprintf(`
resource "google_compute_network" "foobar" {
  name                    = "%s-net"
  auto_create_subnetworks = "false"
}

resource "google_compute_subnetwork" "subnet1" {
  name          = "%s-subnet1"
  network       = google_compute_network.foobar.self_link
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
  purpose       = "PRIVATE_NAT"
}

resource "google_compute_subnetwork" "subnet2" {
  name          = "%s-subnet2"
  network       = google_compute_network.foobar.self_link
  ip_cidr_range = "10.10.1.0/24"
  region        = "us-central1"
  purpose       = "PRIVATE_NAT"
}

resource "google_compute_subnetwork" "subnet3" {
  name          = "%s-subnet3"
  network       = google_compute_network.foobar.self_link
  ip_cidr_range = "10.158.1.0/24"
  region        = "us-central1"
  purpose       = "PRIVATE_NAT"
}

resource "google_compute_subnetwork" "subnet4" {
  name          = "%s-subnet4"
  network       = google_compute_network.foobar.self_link
  ip_cidr_range = "10.168.1.0/24"
  region        = "us-central1"
  purpose       = "PRIVATE_NAT"
}

resource "google_network_connectivity_hub" "foobar" {
  name        = "%s"
  description = "vpc hub for inter vpc nat"
}

resource "google_network_connectivity_spoke" "primary" {
  name        = "%s-spoke"
  location    = "global"
  description = "vpc spoke for inter vpc nat"
  hub =  google_network_connectivity_hub.foobar.id
  linked_vpc_network {
    exclude_export_ranges = [
      "10.10.0.0/16"
    ]
    uri = google_compute_network.foobar.self_link
  }
}

resource "google_compute_router" "foobar" {
  name     = "%s"
  region   = google_compute_subnetwork.subnet1.region
  network  = google_compute_network.foobar.self_link
  depends_on = [
    google_network_connectivity_spoke.primary
  ]
}
`, routerName, routerName, routerName, routerName, routerName, hubName, routerName, routerName)
}

func testAccComputeRouterNatRulesBasic_privateNatOmitRules(routerName, hubName string) string {
	return fmt.Sprintf(`
%s

resource "google_compute_router_nat" "foobar" {
  name                                = "%s"
  router                              = google_compute_router.foobar.name
  region                              = google_compute_router.foobar.region
  source_subnetwork_ip_ranges_to_nat  = "LIST_OF_SUBNETWORKS"
  type                                = "PRIVATE" 
  enable_dynamic_port_allocation      = false
  enable_endpoint_independent_mapping = false
  min_ports_per_vm = 32
  subnetwork {
    name                    = google_compute_subnetwork.subnet1.id
    source_ip_ranges_to_nat = ["ALL_IP_RANGES"]
  }
}
`, testAccComputeRouterNatBaseResourcesWithPrivateNatSubnetworks(routerName, hubName), routerName)
}

func testAccComputeRouterNatRulesBasic_privateNatWithRuleAndActiveDrainRange(routerName, hubName string, ruleNumber int, ruleDescription, match, activeRanges, drainRanges string) string {
	return fmt.Sprintf(`
%s

resource "google_compute_router_nat" "foobar" {
  name                                = "%s"
  router                              = google_compute_router.foobar.name
  region                              = google_compute_router.foobar.region
  source_subnetwork_ip_ranges_to_nat  = "LIST_OF_SUBNETWORKS"
  type                                = "PRIVATE" 
  enable_dynamic_port_allocation      = false
  enable_endpoint_independent_mapping = false
  min_ports_per_vm = 32
  subnetwork {
    name                    = google_compute_subnetwork.subnet1.id
    source_ip_ranges_to_nat = ["ALL_IP_RANGES"]
  }

  rules {
    rule_number = %d
    description = "%s"
    match       = "%s"
    action {
      source_nat_active_ranges = [%s]
      source_nat_drain_ranges = [%s]
    }
  }
}
`, testAccComputeRouterNatBaseResourcesWithPrivateNatSubnetworks(routerName, hubName), routerName, ruleNumber, ruleDescription, match, activeRanges, drainRanges)
}

func testAccComputeRouterNatRulesBasic_privateNatWithRuleAndEmptyAction(routerName, hubName string, ruleNumber int, ruleDescription, match string) string {
	return fmt.Sprintf(`
%s

resource "google_compute_router_nat" "foobar" {
  name                                = "%s"
  router                              = google_compute_router.foobar.name
  region                              = google_compute_router.foobar.region
  source_subnetwork_ip_ranges_to_nat  = "LIST_OF_SUBNETWORKS"
  type                                = "PRIVATE" 
  enable_dynamic_port_allocation      = false
  enable_endpoint_independent_mapping = false
  min_ports_per_vm = 32
  subnetwork {
    name                    = google_compute_subnetwork.subnet1.id
    source_ip_ranges_to_nat = ["ALL_IP_RANGES"]
  }

  rules {
    rule_number = %d
    description = "%s"
    match       = "%s"
    # action is missing
  }
}
`, testAccComputeRouterNatBaseResourcesWithPrivateNatSubnetworks(routerName, hubName), routerName, ruleNumber, ruleDescription, match)
}

func testAccComputeRouterNatRulesBasic_privateNatWithRuleAndEmptyActionActiveRanges(routerName, hubName string, ruleNumber int, ruleDescription, match string) string {
	return fmt.Sprintf(`
%s

resource "google_compute_router_nat" "foobar" {
  name                                = "%s"
  router                              = google_compute_router.foobar.name
  region                              = google_compute_router.foobar.region
  source_subnetwork_ip_ranges_to_nat  = "LIST_OF_SUBNETWORKS"
  type                                = "PRIVATE" 
  enable_dynamic_port_allocation      = false
  enable_endpoint_independent_mapping = false
  min_ports_per_vm = 32
  subnetwork {
    name                    = google_compute_subnetwork.subnet1.id
    source_ip_ranges_to_nat = ["ALL_IP_RANGES"]
  }

  rules {
    rule_number = %d
    description = "%s"
    match       = "%s"
    action {
      source_nat_active_ranges = []
    }
  }
}
`, testAccComputeRouterNatBaseResourcesWithPrivateNatSubnetworks(routerName, hubName), routerName, ruleNumber, ruleDescription, match)
}
