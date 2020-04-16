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

func TestAccComputeNodeTemplate_nodeTemplateBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeNodeTemplateDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeNodeTemplate_nodeTemplateBasicExample(context),
			},
			{
				ResourceName:      "google_compute_node_template.template",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeNodeTemplate_nodeTemplateBasicExample(context map[string]interface{}) string {
	return Nprintf(`
data "google_compute_node_types" "central1a" {
  zone = "us-central1-a"
}

resource "google_compute_node_template" "template" {
  name      = "tf-test-soletenant-tmpl%{random_suffix}"
  region    = "us-central1"
  node_type = data.google_compute_node_types.central1a.names[0]
}
`, context)
}

func TestAccComputeNodeTemplate_nodeTemplateServerBindingExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckComputeNodeTemplateDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeNodeTemplate_nodeTemplateServerBindingExample(context),
			},
		},
	})
}

func testAccComputeNodeTemplate_nodeTemplateServerBindingExample(context map[string]interface{}) string {
	return Nprintf(`
provider "google-beta" {
  region = "us-central1"
  zone   = "us-central1-a"
}

data "google_compute_node_types" "central1a" {
  provider = google-beta
  zone     = "us-central1-a"
}

resource "google_compute_node_template" "template" {
  provider = google-beta

  name      = "tf-test-soletenant-with-licenses%{random_suffix}"
  region    = "us-central1"
  node_type = data.google_compute_node_types.central1a.names[0]

  node_affinity_labels = {
    foo = "baz"
  }

  server_binding {
    type = "RESTART_NODE_ON_MINIMAL_SERVERS"
  }
}
`, context)
}

func testAccCheckComputeNodeTemplateDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_node_template" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/nodeTemplates/{{name}}")
			if err != nil {
				return err
			}

			_, err = sendRequest(config, "GET", "", url, nil)
			if err == nil {
				return fmt.Errorf("ComputeNodeTemplate still exists at %s", url)
			}
		}

		return nil
	}
}
