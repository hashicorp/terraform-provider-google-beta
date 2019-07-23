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

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccMonitoringGroup_monitoringGroupBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitoringGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMonitoringGroup_monitoringGroupBasicExample(context),
			},
			{
				ResourceName:      "google_monitoring_group.basic",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccMonitoringGroup_monitoringGroupBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_monitoring_group" "basic" {
  display_name = "New Test Group%{random_suffix}"

  filter = "resource.metadata.region=\"europe-west2\""
}
`, context)
}

func TestAccMonitoringGroup_monitoringGroupSubgroupExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitoringGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMonitoringGroup_monitoringGroupSubgroupExample(context),
			},
			{
				ResourceName:      "google_monitoring_group.subgroup",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccMonitoringGroup_monitoringGroupSubgroupExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_monitoring_group" "parent" {
  display_name = "New Test SubGroup%{random_suffix}"
  filter = "resource.metadata.region=\"europe-west2\""
}

resource "google_monitoring_group" "subgroup" {
  display_name = "New Test SubGroup%{random_suffix}"
  filter = "resource.metadata.region=\"europe-west2\""
  parent_name =  "${google_monitoring_group.parent.name}"
}
`, context)
}

func testAccCheckMonitoringGroupDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_monitoring_group" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(config, rs, "{{MonitoringBasePath}}{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", url, nil)
		if err == nil {
			return fmt.Errorf("MonitoringGroup still exists at %s", url)
		}
	}

	return nil
}
