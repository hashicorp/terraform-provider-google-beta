package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccServiceNetworkingSubnetwork_create(t *testing.T) {
	t.Parallel()

	project := getTestProjectFromEnv()

	consumerNetwork := BootstrapSharedTestNetwork(t, "service-networking-subnetwork-consumer")
	addr := fmt.Sprintf("tf-test-%s", randString(t, 10))
	service := "servicenetworking.googleapis.com"
	subnetworkName := "subnetwork"

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testServiceNetworkingSubnetworkDestroy(t, service, consumerNetwork),
		Steps: []resource.TestStep{
			{
				Config: testAccServiceNetworkingSubnetwork(consumerNetwork, addr, service, subnetworkName, project),
			},
			{
				ResourceName:      "google_service_networking_subnetwork.subnetwork",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testServiceNetworkingSubnetworkDestroy(t *testing.T, parent, network string) resource.TestCheckFunc {
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

func testAccServiceNetworkingSubnetwork(networkName, addressRangeName, serviceName, subnetworkName, projectId string) string {
	return fmt.Sprintf(`
data "google_compute_network" "consumer_network" {
 name = "%s"
}

data "google_project" "consumer_project" {
	project_id = "%s"
}

resource "google_compute_global_address" "foobar" {
 name          = "%s"
 purpose       = "VPC_PEERING"
 address_type  = "INTERNAL"
 prefix_length = 16
 network       = data.google_compute_network.consumer_network.self_link
}

resource "google_service_networking_connection" "connection" {
 network                 = data.google_compute_network.consumer_network.self_link
 service                 = "%s"
 reserved_peering_ranges = [google_compute_global_address.foobar.name]
}

resource "google_service_networking_subnetwork" "subnetwork" {
  provider         = google-beta

	consumer_network = "projects/${data.google_project.consumer_project.number}/global/networks/%[1]s"
  consumer    		 = data.google_project.consumer_project.number
	
  service		  		 = "%s"
	name        	   = "%s"
	region      		 = "us-central1"	
  
  subnetwork_users = [
    "serviceAccount:${data.google_project.consumer_project.number}@cloudservices.gserviceaccount.com"
  ]

	ip_prefix_length = 24
	secondary_ip_range_spec {
		range_name = "secondary-pods"
		ip_prefix_length = 20
	}
  secondary_ip_range_spec {
		range_name = "secondary-nodes"
		ip_prefix_length = 20
	}

	validate = true
}
`, networkName, projectId, addressRangeName, serviceName, serviceName, subnetworkName)
}
