package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccServiceNetworkingRoleBinding_basic(t *testing.T) {
	t.Parallel()

	project := getTestProjectFromEnv()

	network := BootstrapSharedTestNetwork(t, "service-networking-role-binding-basic")
	addr := fmt.Sprintf("tf-test-%s", randString(t, 10))
	service := "servicenetworking.googleapis.com"

	connectionResources := testAccServiceNetworkingRoleBindingConnection(network, addr, service)

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testServiceNetworkingRoleBindingDestroy(t, service, network),
		Steps: []resource.TestStep{
			{
				Config: testAccServiceNetworkingRoleBinding1(connectionResources, network, service, project),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("google_service_networking_role_binding.roles",
						"policy_binding.#", "1"),
					resource.TestCheckResourceAttr("google_service_networking_role_binding.roles",
						"policy_binding.0.role", "roles/container.hostServiceAgentUser"),
				),
			},
			{
				Config: testAccServiceNetworkingRoleBinding2(connectionResources, network, service, project),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("google_service_networking_role_binding.roles",
						"policy_binding.#", "2"),
					resource.TestCheckResourceAttr("google_service_networking_role_binding.roles",
						"policy_binding.0.role", "roles/container.hostServiceAgentUser"),
					resource.TestCheckResourceAttr("google_service_networking_role_binding.roles",
						"policy_binding.1.role", "roles/compute.securityAdmin"),
				),
			},
		},
	})

}

func testServiceNetworkingRoleBindingDestroy(t *testing.T, parent, network string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config := googleProviderConfig(t)
		parentService := "services/" + parent
		networkName := fmt.Sprintf("projects/%s/global/networks/%s", getTestProjectFromEnv(), network)

		response, err := config.NewServiceNetworkingClient(config.userAgent).Services.Connections.List(parentService).
			Network(networkName).Do()
		if err != nil {
			return err
		}

		for _, c := range response.Connections {
			if c.Network == networkName {
				return fmt.Errorf("Found %s which should have been destroyed.", networkName)
			}
		}

		return nil
	}
}

func testAccServiceNetworkingRoleBindingConnection(networkName, addressRangeName, serviceName string) string {
	return fmt.Sprintf(`
data "google_compute_network" "servicenet" {
  name = "%s"
}

resource "google_compute_global_address" "foobar" {
  name          = "%s"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = data.google_compute_network.servicenet.self_link
}

resource "google_service_networking_connection" "foobar" {
  network                 = data.google_compute_network.servicenet.self_link
  service                 = "%s"
  reserved_peering_ranges = [google_compute_global_address.foobar.name]
}
`, networkName, addressRangeName, serviceName)
}

func testAccServiceNetworkingRoleBinding1(connectionBlock, networkName, serviceName, projectId string) string {
	return fmt.Sprintf(`
%s

data "google_project" "consumer_project" {
  project_id = "%s"
}

resource "google_service_networking_role_binding" "roles" {
  provider         = google-beta

	consumer_network = "projects/${data.google_project.consumer_project.number}/global/networks/%s"
  service          = "%s"
  policy_binding   {
    member = "serviceAccount:${data.google_project.consumer_project.number}@cloudservices.gserviceaccount.com"
    role   = "roles/container.hostServiceAgentUser"
  }
}
`, connectionBlock, projectId, networkName, serviceName)
}

func testAccServiceNetworkingRoleBinding2(connectionBlock, networkName, serviceName, projectId string) string {
	return fmt.Sprintf(`
%s

data "google_project" "consumer_project" {
  project_id = "%s"
}

resource "google_service_networking_role_binding" "roles" {
  provider         = google-beta

	consumer_network = "projects/${data.google_project.consumer_project.number}/global/networks/%s"
  service          = "%s"
  policy_binding   {
    member = "serviceAccount:${data.google_project.consumer_project.number}@cloudservices.gserviceaccount.com"
    role   = "roles/container.hostServiceAgentUser"
  }
  policy_binding   {
    member = "serviceAccount:${data.google_project.consumer_project.number}@cloudservices.gserviceaccount.com"
    role   = "roles/compute.securityAdmin"
  }
}
`, connectionBlock, projectId, networkName, serviceName)
}
