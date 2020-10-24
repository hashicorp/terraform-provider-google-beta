package google

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/servicenetworking/v1"
)

var (
	// TODO(nickmaatgoogle): enable this feature once api release v0.34.0 is released
	secondaryIpRangeSpec = &schema.Resource{
		Schema: map[string]*schema.Schema{
			"range_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateRFC1035Name(1, 63),
				Description: `A name for the secondary IP range.  The name must be 1-63 characters long, and comply with RFC1035. 
The name must be unique within the subnetwork.`,
			},
			"ip_prefix_length": {
				Type:     schema.TypeInt,
				Required: true,
				Description: `The prefix length of the secondary IP range. 
Use CIDR range notation, such as "30"" to provision a secondary IP range with an "x.x.x.x/30" CIDR range. 
The IP address range is drawn from a pool of available ranges in the service consumer's allocated range.`,
			},
			"requested_address": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsIPv4Address,
				Description: `The starting address of a range. The address must be a valid IPv4 address in the x.x.x.x format. 
This value combined with the IP prefix range is the CIDR range for the secondary IP range. 
The range must be within the allocated range that is assigned to the private connection.`,
			},
		},
	}
)

func resourceServiceNetworkingSubnet() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceNetworkingSubnetCreate,
		Read:   resourceServiceNetworkingSubnetRead,
		Delete: resourceServiceNetworkingSubnetDelete,
		Importer: &schema.ResourceImporter{
			State: resourceServiceNetworkingSubnetImportState,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"network": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `Name of VPC network connected with service producers using VPC peering.`,
			},
			"service": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Provider peering service that is managing peering connectivity for a service provider organization. For Google services that support this functionality it is 'servicenetworking.googleapis.com'.`,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `A name for the new subnet. Naming follows the same constraints as 
[compute subnetworks](https://www.terraform.io/docs/providers/google/r/compute_subnetwork.html#name)`,
			},
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The GCP region for this subnetwork.`,
			},
			"consumer": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `A resource that represents the service consumer, such as "projects/123456". 
The project number can be different from the value in the consumer network parameter. 
For example, the network might be part of a Shared VPC network. 
In those cases, Service Networking validates that this resource belongs to that Shared VPC.`,
			},
			"ip_prefix_length": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
				Description: `The prefix length of the subnet's IP address range.
Use CIDR range notation, such as "30" to provision a subnet with an "x.x.x.x/30" CIDR range.
The IP address range is drawn from a pool of available ranges in the service consumer's allocated range.`,
			},
			"description": {
				Type:        schema.TypeString,
				Required:    false,
				ForceNew: true,
				Description: `Attach context onto this subnet resource.`,
			},
			"subnetwork_users": {
				Type:        schema.TypeList,
				Required:    false,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: `A list of members that are granted the "compute.networkUser" role on the subnet.`,
			},
			"requested_address": {
				Type:     schema.TypeString,
				Required: false,
				ForceNew: true,
				Description: `The starting address of a range. The address must be a valid IPv4 address in the x.x.x.x format. 
This value combined with the IP prefix range is the CIDR range for the subnet. 
The range must be within the allocated range that is assigned to the private connection.`,
				ValidateFunc: validation.IsIPv4Address,
			},
			// TODO(nickmaatgoogle): enable this feature once api release v0.34.0 is released
			"secondary_ip_range_specs": {
				Type:        schema.TypeList,
				Required:    false,
				ForceNew:    true,
				Elem:        secondaryIpRangeSpec,
				Description: `A list of secondary IP ranges to be created within the new subnetwork.`,
			},
			"validate": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: `If true, submits the inputs for validation prior to creating the a new subnetwork. 
Validation will perform basic sanity checking on permissions and connectivity prerequisites. 
Will fail resource creation faster with a minor startup delay.`,
			},
		},
	}
}

func resourceServiceNetworkingSubnetCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	region, err := getRegion(d, config)
	if err != nil {
		return err
	}

	network := d.Get("network").(string)
	serviceNetworkingNetworkName, err := retrieveServiceNetworkingNetworkName(d, config, network, userAgent)
	if err != nil {
		return errwrap.Wrapf("Failed to find Service Networking Connection, err: {{err}}", err)
	}

	subnetRequest := &servicenetworking.AddSubnetworkRequest{
		Consumer: d.Get("consumer").(string),
		ConsumerNetwork: serviceNetworkingNetworkName,
		Description: d.Get("description").(string),
		IpPrefixLength: int64(d.Get("ip_prefix_length").(int)),
		Region: region,
		RequestedAddress: d.Get("requested_address").(string),
		Subnetwork: d.Get("name").(string),
		SubnetworkUsers: convertStringArr(d.Get("subnetwork_users").([]interface{})),
		SecondaryIpRangeSpecs: expandSubnetworkSecondaryIpSpec(d.Get("secondary_ip_range_specs")),
	}

	parentService := formatParentService(d.Get("service").(string))

	op, err := config.NewServiceNetworkingClient(userAgent).Services.AddSubnetwork(parentService, subnetRequest).Do()
	if err != nil {
		return err
	}

	if err := serviceNetworkingOperationWaitTime(config, op, "Create Service Networking Subnetwork", userAgent, d.Timeout(schema.TimeoutCreate)); err != nil {
		return err
	}

	// NOTE(nickmaatgoogle): An out of band aspect of this API is that it returns metadata about the resource
	// inside the operation. The subnetwork returned will subsequently be managed by regular compute_subnetwork apis for read and delete.
	// So we record relevant metadata after the operation finishes.
	var subnetworkRes servicenetworking.Subnetwork
	if err := serviceNetworkOperationWaitTimeWithResponse(
		config, op, &subnetworkRes, "Create Service Networking Subnetwork", userAgent,
		d.Timeout(schema.TimeoutCreate)); err != nil {
		return nil
	}

	if len(subnetworkRes.Name) == 0 || len(subnetworkRes.Network) == 0  {
		return fmt.Errorf("Returned subnet resource is empty: %s", op.Name)
	}

	subnetwork := subnetworkId{
		ConsumerNetwork: network,
		HostNetwork: subnetworkRes.Network,
		Name: subnetworkRes.Name,
		Service: d.Get("service").(string),
	}
	d.SetId(subnetwork.Id())
	if err = recordSubnetworkMetadata(d, subnetworkRes); err != nil {
		return err
	}
	return resourceServiceNetworkingSubnetRead(d, meta)
}

func resourceServiceNetworkingSubnetRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	region, err := getRegion(d, config)
	if err != nil {
		return err
	}

	subnetwork, err := config.NewComputeClient(userAgent).Subnetworks.Get(project, region, d.Id()).Do()
	if err != nil {
		return err
	}

	if subnetwork == nil {
		d.SetId("")
		log.Printf("[WARNING] Failed to find Service Networking Subnetwork, network: %s", d.Id())
		return nil
	}

	//if err := d.Set("network", connectionId.Network); err != nil {
	//	return fmt.Errorf("Error setting network: %s", err)
	//}
	//if err := d.Set("service", connectionId.Service); err != nil {
	//	return fmt.Errorf("Error setting service: %s", err)
	//}
	//if err := d.Set("peering", connection.Peering); err != nil {
	//	return fmt.Errorf("Error setting peering: %s", err)
	//}
	//if err := d.Set("reserved_peering_ranges", connection.ReservedPeeringRanges); err != nil {
	//	return fmt.Errorf("Error setting reserved_peering_ranges: %s", err)
	//}
	return nil
}

func resourceServiceNetworkingSubnetUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	connectionId, err := parseConnectionId(d.Id())
	if err != nil {
		return errwrap.Wrapf("Unable to parse Service Networking Connection id, err: {{err}}", err)
	}

	parentService := formatParentService(connectionId.Service)

	if d.HasChange("reserved_peering_ranges") {
		network := d.Get("network").(string)
		serviceNetworkingNetworkName, err := retrieveServiceNetworkingNetworkName(d, config, network, userAgent)
		if err != nil {
			return errwrap.Wrapf("Failed to find Service Networking Connection, err: {{err}}", err)
		}

		connection := &servicenetworking.Connection{
			Network:               serviceNetworkingNetworkName,
			ReservedPeeringRanges: convertStringArr(d.Get("reserved_peering_ranges").([]interface{})),
		}

		// The API docs don't specify that you can do connections/-, but that's what gcloud does,
		// and it's easier than grabbing the connection name.
		op, err := config.NewServiceNetworkingClient(userAgent).Services.Connections.Patch(parentService+"/connections/-", connection).UpdateMask("reservedPeeringRanges").Force(true).Do()
		if err != nil {
			return err
		}
		if err := serviceNetworkingOperationWaitTime(config, op, "Update Service Networking Connection", userAgent, d.Timeout(schema.TimeoutUpdate)); err != nil {
			return err
		}
	}
	return resourceServiceNetworkingSubnetRead(d, meta)
}

func resourceServiceNetworkingSubnetDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	network := d.Get("network").(string)
	serviceNetworkingNetworkName, err := retrieveServiceNetworkingNetworkName(d, config, network, userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	peering := d.Get("peering").(string)
	obj["name"] = peering
	url := fmt.Sprintf("%s%s/removePeering", config.ComputeBasePath, serviceNetworkingNetworkName)

	networkFieldValue, err := ParseNetworkFieldValue(network, d, config)
	if err != nil {
		return errwrap.Wrapf("Failed to retrieve network field value, err: {{err}}", err)
	}

	project := networkFieldValue.Project
	res, err := sendRequestWithTimeout(config, "POST", project, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ServiceNetworkingSubnet %q", d.Id()))
	}

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWaitTime(
		config, op, project, "Updating Network", userAgent, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return err
	}

	d.SetId("")
	log.Printf("[INFO] Service network connection removed.")

	return nil
}

func recordSubnetworkMetadata(d *schema.ResourceData, res servicenetworking.Subnetwork) error {
	if err := d.Set("ip_cidr_range", res.IpCidrRange); err != nil {
		return fmt.Errorf("Error setting ip_cidr_range: %s", err)
	}

	if err := d.Set("secondary_ip_ranges", res.SecondaryIpRanges); err != nil {
		return fmt.Errorf("Error setting secondary_ip_ranges: %s", err)
	}

	if err := d.Set("outside_allocation", res.OutsideAllocation); err != nil {
		return fmt.Errorf("Error setting outside_allocation: %s", err)
	}

	return nil
}

func resourceServiceNetworkingSubnetImportState(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	connectionId, err := parseConnectionId(d.Id())
	if err != nil {
		return nil, err
	}

	if err := d.Set("network", connectionId.Network); err != nil {
		return nil, fmt.Errorf("Error setting network: %s", err)
	}
	if err := d.Set("service", connectionId.Service); err != nil {
		return nil, fmt.Errorf("Error setting service: %s", err)
	}
	return []*schema.ResourceData{d}, nil
}

//
func subnetworkFullName(project, region, network string) string {
	return fmt.Sprintf("projects/%s/regions/%s/subnetworks/%s", project, region, network)
}

func expandSubnetworkSecondaryIpSpec(configured interface{}) []*servicenetworking.SecondaryIpRangeSpec {
	l := configured.([]interface{})
	if len(l) == 0 {
		return nil
	}
	result := make([]*servicenetworking.SecondaryIpRangeSpec, 0, len(l))

	for _, v := range l {
		spec := v.(map[string]interface{})
		result = append(result, &servicenetworking.SecondaryIpRangeSpec{
			IpPrefixLength: int64(spec["ip_prefix_length"].(int)),
			RangeName: spec["range_name"].(string),
			RequestedAddress: spec["requested_address"].(string),
 		})
	}

	return result
}

// NOTE(nickmaatgoogle): The Subnet resource in managed by this API doesn't have an Id field, so inorder
// to support the Read method, we create an Id using the tuple(ConsumerNetwork, HostNetwork, Name, Service).
type subnetworkId struct {
	ConsumerNetwork string
	HostNetwork string
	Name string
	Service string
}

func (id *subnetworkId) Id() string {
	return fmt.Sprintf("%s:%s:%s:%s",
		url.QueryEscape(id.ConsumerNetwork), url.QueryEscape(id.HostNetwork), url.QueryEscape(id.Name), url.QueryEscape(id.Service))
}

func parseSubnetworkId(id string) (*subnetworkId, error) {
	res := strings.Split(id, ":")

	if len(res) != 4 {
		return nil, fmt.Errorf("Failed to parse service networking subnetwork id, value: %s", id)
	}

	consumerNetwork, err := url.QueryUnescape(res[0])
	if err != nil {
		return nil, errwrap.Wrapf("Failed to parse service networking connection id, invalid network, err: {{err}}", err)
	} else if len(consumerNetwork) == 0 {
		return nil, fmt.Errorf("Failed to parse service networking subnetwork id, empty consumer network")
	}

	hostNetwork, err := url.QueryUnescape(res[1])
	if err != nil {
		return nil, errwrap.Wrapf("Failed to parse service networking subnetwork id, invalid host network, err: {{err}}", err)
	} else if len(hostNetwork) == 0 {
		return nil, fmt.Errorf("Failed to parse service networking subnetwork id, empty host network")
	}

	name, err := url.QueryUnescape(res[2])
	if err != nil {
		return nil, errwrap.Wrapf("Failed to parse service networking subnetwork id, invalid service, err: {{err}}", err)
	} else if len(name) == 0 {
		return nil, fmt.Errorf("Failed to parse service networking subnetwork id, empty name")
	}

	service, err := url.QueryUnescape(res[3])
	if err != nil {
		return nil, errwrap.Wrapf("Failed to parse service networking subnetwork id, invalid service, err: {{err}}", err)
	} else if len(service) == 0 {
		return nil, fmt.Errorf("Failed to parse service networking subnetwork id, empty service")
	}

	return &subnetworkId{
		ConsumerNetwork: consumerNetwork,
		HostNetwork: hostNetwork,
		Name: name,
		Service: service,
	}, nil
}