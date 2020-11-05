package google

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/servicenetworking/v1"
)

var (
	serviceNetworkingPolicyBinding = &schema.Resource{
		Schema: map[string]*schema.Schema{
			"role": {
				Type:     schema.TypeString,
				Required: true,
				Description: `Role to be applied on the shared VPC host. Only allowlisted roles can be used at the specified granularity. These role must be allowed in organization setup.

Common roles required to use ServiceNetworking with GKE are:
* "roles/container.hostServiceAgentUser" 
* "roles/compute.securityAdmin"`,
			},
			"member": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Member to bind the role with. e.g. "user:myuser@mydomain.com", "serviceAccount:my-service-account@iam.gserviceaccount.com"`,
			},
		},
	}
)

// The policies cannot be deleted once added. Deletion is performed when the shared network is removed by the client.
func resourceServiceNetworkingRoleBinding() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceNetworkingRoleBindingCreateUpdate,
		Update: resourceServiceNetworkingRoleBindingCreateUpdate,
		Read:   resourceServiceNetworkingRoleBindingRead,
		Delete: resourceServiceNetworkingRoleBindingDelete,

		Schema: map[string]*schema.Schema{
			"service": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Provider peering service that is managing peering connectivity for a service provider organization. For Google services that support this functionality it is 'servicenetworking.googleapis.com'.`,
			},

			// NOTE(nickmaatgoogle): An out of band aspect of this API is that it uses a unique formatting of network
			// different from the standard self_link URI. This API uses project numbers only.
			// Most service providers will not have permissions to call resource manager on a consumer's project to convert project-id
			// into project-num. So we enforce that only project numbers are used.
			"consumer_network": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateServiceNetworkingNetworkResource(),
				Description: `Name of VPC network connected with service producers using VPC peering.
The network must use project number and be of the format: projects/12345/global/networks/consumernetwork`,
			},

			"policy_binding": {
				Type:     schema.TypeList,
				Required: true,
				// Activate the "Attributes as Blocks" processing mode
				ConfigMode:  schema.SchemaConfigModeAttr,
				Elem:        serviceNetworkingPolicyBinding,
				Description: `List of policy bindings to add to shared VPC host project.`,
			},

			"project": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: `The ID of the project in which the resource belongs. If it is not provided, the provider project is used.`,
			},
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: `The GCP region for this subnetwork. If it is not provided, the provider project is used.`,
			},
		},
	}
}

func resourceServiceNetworkingRoleBindingDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")
	return nil
}

func resourceServiceNetworkingRoleBindingRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	policyBinding, err := parsePolicyBindingId(d.Id())
	if err != nil {
		return errwrap.Wrapf("Unable to parse Service Networking Subnetwork id, err: {{err}}", err)
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	region, err := getRegion(d, config)
	if err != nil {
		return err
	}

	if err := d.Set("consumer_network", policyBinding.ConsumerNetwork); err != nil {
		return fmt.Errorf("Error setting consumer_network: %s", err)
	}
	if err := d.Set("service", policyBinding.Service); err != nil {
		return fmt.Errorf("Error setting service: %s", err)
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading resource project: %s", err)
	}
	if err := d.Set("region", region); err != nil {
		return fmt.Errorf("Error reading resource region: %s", err)
	}

	return nil
}

func resourceServiceNetworkingRoleBindingCreateUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	consumerNetworkName := d.Get("consumer_network").(string)
	peeringService := d.Get("service").(string)
	parentService := formatParentService(peeringService)
	policyBindings := expandServiceNetworkingPolicyBinding(d.Get("policy_binding"))

	if len(policyBindings) == 0 {
		return errors.New("cannot create resource with no policies, expected at least 1 policy")
	}

	rolesRequest := &servicenetworking.AddRolesRequest{
		ConsumerNetwork: consumerNetworkName,
		PolicyBinding:   policyBindings,
	}

	op, err := config.NewServiceNetworkingClient(userAgent).Services.Roles.Add(parentService, rolesRequest).Do()
	if err != nil {
		return err
	}

	var rolesResponse servicenetworking.AddRolesResponse
	if err := serviceNetworkOperationWaitTimeWithRolesResponse(config, op, &rolesResponse, "Adding Policy Binding to Service Networking Shared VPC", userAgent, d.Timeout(schema.TimeoutCreate)); err != nil {
		return err
	}

	if len(rolesResponse.PolicyBinding) == 0 {
		return errors.New("empty policy binding returned, expected at least 1 policy")
	}

	if err := recordPolicies(d, rolesResponse); err != nil {
		return err
	}
	policyBinding := &policyBindingId{
		ConsumerNetwork: consumerNetworkName,
		Service:         peeringService,
	}
	d.SetId(policyBinding.Id())

	return resourceServiceNetworkingRoleBindingRead(d, meta)
}

func recordPolicies(d *schema.ResourceData, response servicenetworking.AddRolesResponse) error {
	log.Printf("[DEBUG] %d new roles added.", len(response.PolicyBinding))
	if err := d.Set("policy_binding", flattenServiceNetworkingPolicyBinding(response.PolicyBinding)); err != nil {
		return fmt.Errorf("Error parsing PolicyBinding Response: %s", err)
	}

	return nil
}

func expandServiceNetworkingPolicyBinding(configured interface{}) []*servicenetworking.PolicyBinding {
	l := configured.([]interface{})
	if len(l) == 0 {
		return nil
	}
	result := make([]*servicenetworking.PolicyBinding, 0, len(l))

	for _, v := range l {
		spec := v.(map[string]interface{})
		result = append(result, &servicenetworking.PolicyBinding{
			Member: spec["member"].(string),
			Role:   spec["role"].(string),
		})
	}

	return result
}

func flattenServiceNetworkingPolicyBinding(policyBindings []*servicenetworking.PolicyBinding) interface{} {
	transformed := make([]interface{}, 0, len(policyBindings))
	for _, policyBinding := range policyBindings {
		transformed = append(transformed, map[string]interface{}{
			"member": policyBinding.Member,
			"role":   policyBinding.Role,
		})
	}
	return transformed
}

const networkResourcePattern = "^projects/[-1-9][-0-9]+/global/.+$"

func validateServiceNetworkingNetworkResource() schema.SchemaValidateFunc {
	return validateRegexp(networkResourcePattern)
}

// NOTE(nickmaatgoogle): The policy binding resource managed by this API doesn't have an Id field, in order
// to support the Read method, we create an Id using the tuple(ConsumerNetwork, Service).
type policyBindingId struct {
	ConsumerNetwork string
	Service         string
}

func (id policyBindingId) Id() string {
	return fmt.Sprintf("%s:%s", url.QueryEscape(id.ConsumerNetwork), url.QueryEscape(id.Service))
}

func parsePolicyBindingId(id string) (*policyBindingId, error) {
	res := strings.Split(id, ":")

	if len(res) != 2 {
		return nil, fmt.Errorf("Failed to parse service networking policy binding id, value: %s", id)
	}

	network, err := url.QueryUnescape(res[0])
	if err != nil {
		return nil, errwrap.Wrapf("Failed to parse service networking policy id, invalid consumer network, err: {{err}}", err)
	} else if len(network) == 0 {
		return nil, fmt.Errorf("Failed to parse service networking policy id, empty consumer network")
	}

	service, err := url.QueryUnescape(res[1])
	if err != nil {
		return nil, errwrap.Wrapf("Failed to parse service networking policy id, invalid service, err: {{err}}", err)
	} else if len(service) == 0 {
		return nil, fmt.Errorf("Failed to parse service networking policy id, empty service")
	}

	return &policyBindingId{
		ConsumerNetwork: network,
		Service:         service,
	}, nil
}
