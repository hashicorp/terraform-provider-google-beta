package google

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccComputeOrganizationSecurityPolicyRule_organizationSecurityPolicyRuleUpdateExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        getTestOrgFromEnv(t),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeOrganizationSecurityPolicyRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeOrganizationSecurityPolicyRule_organizationSecurityPolicyRulePreUpdateExample(context),
			},
			{
				ResourceName:      "google_compute_organization_security_policy_rule.policy",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeOrganizationSecurityPolicyRule_organizationSecurityPolicyRulePostUpdateExample(context),
			},
			{
				ResourceName:      "google_compute_organization_security_policy_rule.policy",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeOrganizationSecurityPolicyRule_organizationSecurityPolicyRulePreUpdateExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_organization_security_policy" "policy" {
  display_name = "tf-test%{random_suffix}"
  parent       = "organizations/%{org_id}"
}

resource "google_compute_organization_security_policy_rule" "policy" {

  policy_id = google_compute_organization_security_policy.policy.id
  action = "allow"

  direction = "INGRESS"
  enable_logging = true
  match {
    config {
      src_ip_ranges = ["192.168.0.0/16", "10.0.0.0/8"]
      layer4_config {
        ip_protocol = "tcp"
        ports = ["22"]
      }
      layer4_config {
        ip_protocol = "icmp"
      }
    }
  }
  priority = 100
}
`, context)
}

func testAccComputeOrganizationSecurityPolicyRule_organizationSecurityPolicyRulePostUpdateExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_organization_security_policy" "policy" {
  display_name = "tf-test%{random_suffix}"
  parent       = "organizations/%{org_id}"
}

resource "google_compute_organization_security_policy_rule" "policy" {

  policy_id = google_compute_organization_security_policy.policy.id
  action = "deny"

  direction = "INGRESS"
  enable_logging = false
  description = "Updated description"
  match {
    config {
      src_ip_ranges = ["172.16.0.0/12"]
      layer4_config {
        ip_protocol = "udp"
        ports = ["53"]
      }
    }
  }
  priority = 100
}
`, context)
}
