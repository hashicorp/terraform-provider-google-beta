package google

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/errwrap"
	version "github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	containerBeta "google.golang.org/api/container/v1beta1"
)

var (
	instanceGroupManagerURL = regexp.MustCompile(fmt.Sprintf("^https://www.googleapis.com/compute/v1/projects/(%s)/zones/([a-z0-9-]*)/instanceGroupManagers/([^/]*)", ProjectRegex))

	networkConfig = &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cidr_blocks": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				MaxItems: 20,
				Elem:     cidrBlockConfig,
			},
		},
	}
	cidrBlockConfig = &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cidr_block": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.CIDRNetwork(0, 32),
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}

	ipAllocationSubnetFields    = []string{"ip_allocation_policy.0.create_subnetwork", "ip_allocation_policy.0.subnetwork_name"}
	ipAllocationCidrBlockFields = []string{"ip_allocation_policy.0.cluster_ipv4_cidr_block", "ip_allocation_policy.0.services_ipv4_cidr_block", "ip_allocation_policy.0.node_ipv4_cidr_block"}
	ipAllocationRangeFields     = []string{"ip_allocation_policy.0.cluster_secondary_range_name", "ip_allocation_policy.0.services_secondary_range_name"}
)

func resourceContainerCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceContainerClusterCreate,
		Read:   resourceContainerClusterRead,
		Update: resourceContainerClusterUpdate,
		Delete: resourceContainerClusterDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		SchemaVersion: 1,
		MigrateState:  resourceContainerClusterMigrateState,

		Importer: &schema.ResourceImporter{
			State: resourceContainerClusterStateImporter,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					value := v.(string)

					if len(value) > 40 {
						errors = append(errors, fmt.Errorf(
							"%q cannot be longer than 40 characters", k))
					}
					if !regexp.MustCompile("^[a-z0-9-]+$").MatchString(value) {
						errors = append(errors, fmt.Errorf(
							"%q can only contain lowercase letters, numbers and hyphens", k))
					}
					if !regexp.MustCompile("^[a-z]").MatchString(value) {
						errors = append(errors, fmt.Errorf(
							"%q must start with a letter", k))
					}
					if !regexp.MustCompile("[a-z0-9]$").MatchString(value) {
						errors = append(errors, fmt.Errorf(
							"%q must end with a number or a letter", k))
					}
					return
				},
			},

			"location": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"zone", "region"},
			},

			"region": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				Deprecated:    "Use location instead",
				ConflictsWith: []string{"zone", "location"},
			},

			"zone": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				Deprecated:    "Use location instead",
				ConflictsWith: []string{"region", "location"},
			},

			"node_locations": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"additional_zones": {
				Type:       schema.TypeSet,
				Optional:   true,
				Computed:   true,
				Deprecated: "Use node_locations instead",
				Elem:       &schema.Schema{Type: schema.TypeString},
			},

			"addons_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"http_load_balancing": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"disabled": {
										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
						"horizontal_pod_autoscaling": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"disabled": {
										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
						"kubernetes_dashboard": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"disabled": {
										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
						"network_policy_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"disabled": {
										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
						"istio_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"disabled": {
										Type:     schema.TypeBool,
										Default:  false,
										Optional: true,
									},
									"auth": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringInSlice([]string{"AUTH_MUTUAL_TLS"}, false),
									},
								},
							},
						},
						"cloudrun_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"disabled": {
										Type:     schema.TypeBool,
										Default:  false,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"cluster_autoscaling": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"resource_limits": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"resource_type": {
										Type:     schema.TypeString,
										Required: true,
									},
									"minimum": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"maximum": {
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"cluster_ipv4_cidr": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: orEmpty(validateRFC1918Network(8, 32)),
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"enable_binary_authorization": {
				Default:  false,
				Type:     schema.TypeBool,
				Optional: true,
			},

			"enable_kubernetes_alpha": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Default:  false,
			},

			"enable_tpu": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Default:  false,
			},

			"enable_legacy_abac": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			"initial_node_count": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},

			"logging_service": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringInSlice([]string{"logging.googleapis.com", "logging.googleapis.com/kubernetes", "none"}, false),
			},

			"maintenance_policy": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"daily_maintenance_window": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"start_time": {
										Type:             schema.TypeString,
										Required:         true,
										ValidateFunc:     validateRFC3339Time,
										DiffSuppressFunc: rfc3339TimeDiffSuppress,
									},
									"duration": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},

			"master_auth": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"password": {
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
						},

						"username": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"client_certificate_config": {
							Type:             schema.TypeList,
							MaxItems:         1,
							Optional:         true,
							DiffSuppressFunc: masterAuthClientCertCfgSuppress,
							ForceNew:         true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"issue_client_certificate": {
										Type:             schema.TypeBool,
										Required:         true,
										ForceNew:         true,
										DiffSuppressFunc: masterAuthClientCertCfgSuppress,
									},
								},
							},
						},

						"client_certificate": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"client_key": {
							Type:      schema.TypeString,
							Computed:  true,
							Sensitive: true,
						},

						"cluster_ca_certificate": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"master_authorized_networks_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem:     networkConfig,
			},

			"min_master_version": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"monitoring_service": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringInSlice([]string{"monitoring.googleapis.com", "monitoring.googleapis.com/kubernetes", "none"}, false),
			},

			"network": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "default",
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},

			"network_policy": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"provider": {
							Type:             schema.TypeString,
							Default:          "PROVIDER_UNSPECIFIED",
							Optional:         true,
							ValidateFunc:     validation.StringInSlice([]string{"PROVIDER_UNSPECIFIED", "CALICO"}, false),
							DiffSuppressFunc: emptyOrDefaultStringSuppress("PROVIDER_UNSPECIFIED"),
						},
					},
				},
			},

			"node_config": schemaNodeConfig,

			"node_pool": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true, // TODO(danawillow): Add ability to add/remove nodePools
				Elem: &schema.Resource{
					Schema: schemaNodePool,
				},
			},

			"node_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"pod_security_policy_config": {
				DiffSuppressFunc: podSecurityPolicyCfgSuppress,
				Type:             schema.TypeList,
				Optional:         true,
				MaxItems:         1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},
					},
				},
			},

			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			"subnetwork": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},

			"endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"instance_group_urls": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"master_version": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"ip_allocation_policy": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"use_ip_aliases": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
							ForceNew: true,
						},

						// GKE creates subnetwork automatically
						"create_subnetwork": {
							Type:          schema.TypeBool,
							Optional:      true,
							ForceNew:      true,
							ConflictsWith: ipAllocationRangeFields,
						},

						"subnetwork_name": {
							Type:          schema.TypeString,
							Optional:      true,
							ForceNew:      true,
							ConflictsWith: ipAllocationRangeFields,
						},

						// GKE creates/deletes secondary ranges in VPC
						"cluster_ipv4_cidr_block": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							ConflictsWith:    ipAllocationRangeFields,
							DiffSuppressFunc: cidrOrSizeDiffSuppress,
						},
						"services_ipv4_cidr_block": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							ConflictsWith:    ipAllocationRangeFields,
							DiffSuppressFunc: cidrOrSizeDiffSuppress,
						},
						"node_ipv4_cidr_block": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							ConflictsWith:    ipAllocationRangeFields,
							DiffSuppressFunc: cidrOrSizeDiffSuppress,
						},

						// User manages secondary ranges manually
						"cluster_secondary_range_name": {
							Type:          schema.TypeString,
							Optional:      true,
							Computed:      true,
							ForceNew:      true,
							ConflictsWith: append(ipAllocationSubnetFields, ipAllocationCidrBlockFields...),
						},
						"services_secondary_range_name": {
							Type:          schema.TypeString,
							Optional:      true,
							Computed:      true,
							ForceNew:      true,
							ConflictsWith: append(ipAllocationSubnetFields, ipAllocationCidrBlockFields...),
						},
					},
				},
			},

			"remove_default_node_pool": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"private_cluster": {
				Removed:  "Use private_cluster_config.enable_private_nodes instead.",
				Computed: true,
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},

			"private_cluster_config": {
				Type:             schema.TypeList,
				Optional:         true,
				MaxItems:         1,
				Computed:         true,
				DiffSuppressFunc: containerClusterPrivateClusterConfigSuppress,
				ConflictsWith:    []string{"private_cluster", "master_ipv4_cidr_block"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_private_endpoint": {
							Type:             schema.TypeBool,
							Optional:         true,
							ForceNew:         true,
							DiffSuppressFunc: containerClusterPrivateClusterConfigSuppress,
						},
						"enable_private_nodes": {
							Type:             schema.TypeBool,
							Optional:         true,
							ForceNew:         true,
							DiffSuppressFunc: containerClusterPrivateClusterConfigSuppress,
						},
						"master_ipv4_cidr_block": {
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: orEmpty(validation.CIDRNetwork(28, 28)),
						},
						"private_endpoint": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"public_endpoint": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"master_ipv4_cidr_block": {
				Removed:  "Use private_cluster_config.master_ipv4_cidr_block instead.",
				Computed: true,
				Type:     schema.TypeString,
				Optional: true,
			},

			"resource_labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"tpu_ipv4_cidr_block": {
				Computed: true,
				Type:     schema.TypeString,
			},

			"default_max_pods_per_node": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
		},
	}
}

func resourceContainerClusterCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	location, err := getLocation(d, config)
	if err != nil {
		return err
	}

	clusterName := d.Get("name").(string)

	cluster := &containerBeta.Cluster{
		Name:                           clusterName,
		InitialNodeCount:               int64(d.Get("initial_node_count").(int)),
		MaintenancePolicy:              expandMaintenancePolicy(d.Get("maintenance_policy")),
		MasterAuthorizedNetworksConfig: expandMasterAuthorizedNetworksConfig(d.Get("master_authorized_networks_config")),
		InitialClusterVersion:          d.Get("min_master_version").(string),
		ClusterIpv4Cidr:                d.Get("cluster_ipv4_cidr").(string),
		Description:                    d.Get("description").(string),
		LegacyAbac: &containerBeta.LegacyAbac{
			Enabled:         d.Get("enable_legacy_abac").(bool),
			ForceSendFields: []string{"Enabled"},
		},
		LoggingService:          d.Get("logging_service").(string),
		MonitoringService:       d.Get("monitoring_service").(string),
		NetworkPolicy:           expandNetworkPolicy(d.Get("network_policy")),
		AddonsConfig:            expandClusterAddonsConfig(d.Get("addons_config")),
		EnableKubernetesAlpha:   d.Get("enable_kubernetes_alpha").(bool),
		IpAllocationPolicy:      expandIPAllocationPolicy(d.Get("ip_allocation_policy")),
		PodSecurityPolicyConfig: expandPodSecurityPolicyConfig(d.Get("pod_security_policy_config")),
		EnableTpu:               d.Get("enable_tpu").(bool),
		BinaryAuthorization: &containerBeta.BinaryAuthorization{
			Enabled:         d.Get("enable_binary_authorization").(bool),
			ForceSendFields: []string{"Enabled"},
		},
		Autoscaling:    expandClusterAutoscaling(d.Get("cluster_autoscaling"), d),
		MasterAuth:     expandMasterAuth(d.Get("master_auth")),
		ResourceLabels: expandStringMap(d, "resource_labels"),
	}

	if v, ok := d.GetOk("default_max_pods_per_node"); ok {
		cluster.DefaultMaxPodsConstraint = expandDefaultMaxPodsConstraint(v)
	}

	// Only allow setting node_version on create if it's set to the equivalent master version,
	// since `InitialClusterVersion` only accepts valid master-style versions.
	if v, ok := d.GetOk("node_version"); ok {
		// ignore -gke.X suffix for now. if it becomes a problem later, we can fix it.
		mv := strings.Split(cluster.InitialClusterVersion, "-")[0]
		nv := strings.Split(v.(string), "-")[0]
		if mv != nv {
			return fmt.Errorf("node_version and min_master_version must be set to equivalent values on create")
		}
	}

	if v, ok := d.GetOk("node_locations"); ok {
		locationsSet := v.(*schema.Set)
		if locationsSet.Contains(location) {
			return fmt.Errorf("when using a multi-zonal cluster, additional_zones should not contain the original 'zone'")
		}

		// GKE requires a full list of node locations
		// but when using a multi-zonal cluster our schema only asks for the
		// additional zones, so append the cluster location if it's a zone
		if isZone(location) {
			locationsSet.Add(location)
		}
		cluster.Locations = convertStringSet(locationsSet)
	} else if v, ok := d.GetOk("additional_zones"); ok {
		locationsSet := v.(*schema.Set)
		if locationsSet.Contains(location) {
			return fmt.Errorf("when using a multi-zonal cluster, additional_zones should not contain the original 'zone'")
		}

		// GKE requires a full list of node locations
		// but when using a multi-zonal cluster our schema only asks for the
		// additional zones, so append the cluster location if it's a zone
		if isZone(location) {
			locationsSet.Add(location)
		}
		cluster.Locations = convertStringSet(locationsSet)
	}

	if v, ok := d.GetOk("network"); ok {
		network, err := ParseNetworkFieldValue(v.(string), d, config)
		if err != nil {
			return err
		}
		cluster.Network = network.RelativeLink()
	}

	if v, ok := d.GetOk("subnetwork"); ok {
		subnetwork, err := ParseSubnetworkFieldValue(v.(string), d, config)
		if err != nil {
			return err
		}
		cluster.Subnetwork = subnetwork.RelativeLink()
	}

	nodePoolsCount := d.Get("node_pool.#").(int)
	if nodePoolsCount > 0 {
		nodePools := make([]*containerBeta.NodePool, 0, nodePoolsCount)
		for i := 0; i < nodePoolsCount; i++ {
			prefix := fmt.Sprintf("node_pool.%d.", i)
			nodePool, err := expandNodePool(d, prefix)
			if err != nil {
				return err
			}
			nodePools = append(nodePools, nodePool)
		}
		cluster.NodePools = nodePools
	} else {
		// Node Configs have default values that are set in the expand function,
		// but can only be set if node pools are unspecified.
		cluster.NodeConfig = expandNodeConfig([]interface{}{})
	}

	if v, ok := d.GetOk("node_config"); ok {
		cluster.NodeConfig = expandNodeConfig(v)
	}

	if v, ok := d.GetOk("private_cluster_config"); ok {
		cluster.PrivateClusterConfig = expandPrivateClusterConfig(v)
	}

	req := &containerBeta.CreateClusterRequest{
		Cluster: cluster,
	}

	mutexKV.Lock(containerClusterMutexKey(project, location, clusterName))
	defer mutexKV.Unlock(containerClusterMutexKey(project, location, clusterName))

	parent := fmt.Sprintf("projects/%s/locations/%s", project, location)
	var op *containerBeta.Operation
	err = retry(func() error {
		op, err = config.clientContainerBeta.Projects.Locations.Clusters.Create(parent, req).Do()
		return err
	})
	if err != nil {
		return err
	}

	d.SetId(clusterName)

	// Wait until it's created
	timeoutInMinutes := int(d.Timeout(schema.TimeoutCreate).Minutes())
	waitErr := containerOperationWait(config, op, project, location, "creating GKE cluster", timeoutInMinutes)
	if waitErr != nil {
		if deleteErr := cleanFailedContainerCluster(d, meta); deleteErr != nil {
			log.Printf("[WARN] Unable to clean up cluster from failed creation: %s", deleteErr)
		} else {
			log.Printf("[WARN] Verified failed creation of cluster %s was cleaned up", d.Id())
		}
		// The resource didn't actually create
		d.SetId("")
		return waitErr
	}

	log.Printf("[INFO] GKE cluster %s has been created", clusterName)

	if d.Get("remove_default_node_pool").(bool) {
		parent := fmt.Sprintf("%s/nodePools/%s", containerClusterFullName(project, location, clusterName), "default-pool")
		op, err = config.clientContainerBeta.Projects.Locations.Clusters.NodePools.Delete(parent).Do()
		if err != nil {
			return errwrap.Wrapf("Error deleting default node pool: {{err}}", err)
		}
		err = containerOperationWait(config, op, project, location, "removing default node pool", timeoutInMinutes)
		if err != nil {
			return errwrap.Wrapf("Error deleting default node pool: {{err}}", err)
		}
	}

	return resourceContainerClusterRead(d, meta)
}

func resourceContainerClusterRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	location, err := getLocation(d, config)
	if err != nil {
		return err
	}

	cluster := &containerBeta.Cluster{}
	err = resource.Retry(2*time.Minute, func() *resource.RetryError {
		name := containerClusterFullName(project, location, d.Get("name").(string))
		cluster, err = config.clientContainerBeta.Projects.Locations.Clusters.Get(name).Do()
		if err != nil {
			return resource.NonRetryableError(err)
		}
		if cluster.Status != "RUNNING" {
			return resource.RetryableError(fmt.Errorf("Cluster %q has status %q with message %q", d.Get("name"), cluster.Status, cluster.StatusMessage))
		}
		return nil
	})
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("Container Cluster %q", d.Get("name").(string)))
	}

	d.Set("name", cluster.Name)
	if err := d.Set("network_policy", flattenNetworkPolicy(cluster.NetworkPolicy)); err != nil {
		return err
	}

	d.Set("location", cluster.Location)
	if isZone(cluster.Location) {
		d.Set("zone", cluster.Location)
	} else {
		d.Set("region", cluster.Location)
	}

	locations := schema.NewSet(schema.HashString, convertStringArrToInterface(cluster.Locations))
	locations.Remove(cluster.Zone) // Remove the original zone since we only store additional zones
	d.Set("node_locations", locations)
	d.Set("additional_zones", locations)

	d.Set("endpoint", cluster.Endpoint)
	if err := d.Set("maintenance_policy", flattenMaintenancePolicy(cluster.MaintenancePolicy)); err != nil {
		return err
	}
	if err := d.Set("master_auth", flattenMasterAuth(cluster.MasterAuth)); err != nil {
		return err
	}
	if err := d.Set("master_authorized_networks_config", flattenMasterAuthorizedNetworksConfig(cluster.MasterAuthorizedNetworksConfig)); err != nil {
		return err
	}
	d.Set("initial_node_count", cluster.InitialNodeCount)
	d.Set("master_version", cluster.CurrentMasterVersion)
	d.Set("node_version", cluster.CurrentNodeVersion)
	d.Set("cluster_ipv4_cidr", cluster.ClusterIpv4Cidr)
	d.Set("description", cluster.Description)
	d.Set("enable_kubernetes_alpha", cluster.EnableKubernetesAlpha)
	d.Set("enable_legacy_abac", cluster.LegacyAbac.Enabled)
	d.Set("logging_service", cluster.LoggingService)
	d.Set("monitoring_service", cluster.MonitoringService)
	d.Set("network", cluster.NetworkConfig.Network)
	d.Set("subnetwork", cluster.NetworkConfig.Subnetwork)
	d.Set("enable_binary_authorization", cluster.BinaryAuthorization != nil && cluster.BinaryAuthorization.Enabled)
	d.Set("enable_tpu", cluster.EnableTpu)
	d.Set("tpu_ipv4_cidr_block", cluster.TpuIpv4CidrBlock)
	if err := d.Set("cluster_autoscaling", flattenClusterAutoscaling(cluster.Autoscaling)); err != nil {
		return err
	}
	if cluster.DefaultMaxPodsConstraint != nil {
		d.Set("default_max_pods_per_node", cluster.DefaultMaxPodsConstraint.MaxPodsPerNode)
	}
	if err := d.Set("node_config", flattenNodeConfig(cluster.NodeConfig)); err != nil {
		return err
	}
	d.Set("project", project)
	if err := d.Set("addons_config", flattenClusterAddonsConfig(cluster.AddonsConfig)); err != nil {
		return err
	}
	nps, err := flattenClusterNodePools(d, config, cluster.NodePools)
	if err != nil {
		return err
	}
	if err := d.Set("node_pool", nps); err != nil {
		return err
	}

	if err := d.Set("ip_allocation_policy", flattenIPAllocationPolicy(cluster.IpAllocationPolicy, d, config)); err != nil {
		return err
	}

	if err := d.Set("private_cluster_config", flattenPrivateClusterConfig(cluster.PrivateClusterConfig)); err != nil {
		return err
	}

	igUrls, err := getInstanceGroupUrlsFromManagerUrls(config, cluster.InstanceGroupUrls)
	if err != nil {
		return err
	}
	if err := d.Set("instance_group_urls", igUrls); err != nil {
		return err
	}

	if err := d.Set("pod_security_policy_config", flattenPodSecurityPolicyConfig(cluster.PodSecurityPolicyConfig)); err != nil {
		return err
	}

	d.Set("resource_labels", cluster.ResourceLabels)
	return nil
}

func resourceContainerClusterUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	location, err := getLocation(d, config)
	if err != nil {
		return err
	}

	clusterName := d.Get("name").(string)
	timeoutInMinutes := int(d.Timeout(schema.TimeoutUpdate).Minutes())

	d.Partial(true)

	lockKey := containerClusterMutexKey(project, location, clusterName)

	updateFunc := func(req *containerBeta.UpdateClusterRequest, updateDescription string) func() error {
		return func() error {
			name := containerClusterFullName(project, location, clusterName)
			op, err := config.clientContainerBeta.Projects.Locations.Clusters.Update(name, req).Do()
			if err != nil {
				return err
			}
			// Wait until it's updated
			return containerOperationWait(config, op, project, location, updateDescription, timeoutInMinutes)
		}
	}

	// The ClusterUpdate object that we use for most of these updates only allows updating one field at a time,
	// so we have to make separate calls for each field that we want to update. The order here is fairly arbitrary-
	// if the order of updating fields does matter, it is called out explicitly.
	if d.HasChange("master_authorized_networks_config") {
		c := d.Get("master_authorized_networks_config")
		req := &containerBeta.UpdateClusterRequest{
			Update: &containerBeta.ClusterUpdate{
				DesiredMasterAuthorizedNetworksConfig: expandMasterAuthorizedNetworksConfig(c),
			},
		}

		updateF := updateFunc(req, "updating GKE cluster master authorized networks")
		if err := lockedCall(lockKey, updateF); err != nil {
			return err
		}
		log.Printf("[INFO] GKE cluster %s master authorized networks config has been updated", d.Id())

		d.SetPartial("master_authorized_networks_config")
	}

	if d.HasChange("addons_config") {
		if ac, ok := d.GetOk("addons_config"); ok {
			req := &containerBeta.UpdateClusterRequest{
				Update: &containerBeta.ClusterUpdate{
					DesiredAddonsConfig: expandClusterAddonsConfig(ac),
				},
			}

			updateF := updateFunc(req, "updating GKE cluster addons")
			// Call update serially.
			if err := lockedCall(lockKey, updateF); err != nil {
				return err
			}

			log.Printf("[INFO] GKE cluster %s addons have been updated", d.Id())

			d.SetPartial("addons_config")
		}
	}
	if d.HasChange("enable_binary_authorization") {
		enabled := d.Get("enable_binary_authorization").(bool)
		req := &containerBeta.UpdateClusterRequest{
			Update: &containerBeta.ClusterUpdate{
				DesiredBinaryAuthorization: &containerBeta.BinaryAuthorization{
					Enabled:         enabled,
					ForceSendFields: []string{"Enabled"},
				},
			},
		}

		updateF := updateFunc(req, "updating GKE binary authorization")
		// Call update serially.
		if err := lockedCall(lockKey, updateF); err != nil {
			return err
		}

		log.Printf("[INFO] GKE cluster %s's binary authorization has been updated to %v", d.Id(), enabled)

		d.SetPartial("enable_binary_authorization")
	}

	if d.HasChange("cluster_autoscaling") {
		req := &containerBeta.UpdateClusterRequest{
			Update: &containerBeta.ClusterUpdate{
				DesiredClusterAutoscaling: expandClusterAutoscaling(d.Get("cluster_autoscaling"), d),
			}}

		updateF := updateFunc(req, "updating GKE cluster autoscaling")
		// Call update serially.
		if err := lockedCall(lockKey, updateF); err != nil {
			return err
		}

		log.Printf("[INFO] GKE cluster %s's cluster-wide autoscaling has been updated", d.Id())

		d.SetPartial("cluster_autoscaling")
	}

	if d.HasChange("maintenance_policy") {
		var req *containerBeta.SetMaintenancePolicyRequest
		if mp, ok := d.GetOk("maintenance_policy"); ok {
			req = &containerBeta.SetMaintenancePolicyRequest{
				MaintenancePolicy: expandMaintenancePolicy(mp),
			}
		} else {
			req = &containerBeta.SetMaintenancePolicyRequest{
				NullFields: []string{"MaintenancePolicy"},
			}
		}

		updateF := func() error {
			name := containerClusterFullName(project, location, clusterName)
			op, err := config.clientContainerBeta.Projects.Locations.Clusters.SetMaintenancePolicy(name, req).Do()

			if err != nil {
				return err
			}

			// Wait until it's updated
			return containerOperationWait(config, op, project, location, "updating GKE cluster maintenance policy", timeoutInMinutes)
		}

		// Call update serially.
		if err := lockedCall(lockKey, updateF); err != nil {
			return err
		}

		log.Printf("[INFO] GKE cluster %s maintenance policy has been updated", d.Id())

		d.SetPartial("maintenance_policy")
	}

	// we can only ever see a change to one of additional_zones and node_locations; because
	// thy conflict with each other and are each computed, Terraform will suppress the diff
	// on one of them even when migrating from one to the other.
	if d.HasChange("additional_zones") {
		azSetOldI, azSetNewI := d.GetChange("additional_zones")
		azSetNew := azSetNewI.(*schema.Set)
		azSetOld := azSetOldI.(*schema.Set)
		if azSetNew.Contains(location) {
			return fmt.Errorf("additional_zones should not contain the original 'zone'")
		}
		// Since we can't add & remove zones in the same request, first add all the
		// zones, then remove the ones we aren't using anymore.
		azSet := azSetOld.Union(azSetNew)

		if isZone(location) {
			azSet.Add(location)
		}

		req := &containerBeta.UpdateClusterRequest{
			Update: &containerBeta.ClusterUpdate{
				DesiredLocations: convertStringSet(azSet),
			},
		}

		updateF := updateFunc(req, "updating GKE cluster node locations")
		// Call update serially.
		if err := lockedCall(lockKey, updateF); err != nil {
			return err
		}

		if isZone(location) {
			azSetNew.Add(location)
		}
		if !azSet.Equal(azSetNew) {
			req = &containerBeta.UpdateClusterRequest{
				Update: &containerBeta.ClusterUpdate{
					DesiredLocations: convertStringSet(azSetNew),
				},
			}

			updateF := updateFunc(req, "updating GKE cluster node locations")
			// Call update serially.
			if err := lockedCall(lockKey, updateF); err != nil {
				return err
			}
		}

		log.Printf("[INFO] GKE cluster %s node locations have been updated to %v", d.Id(), azSet.List())

		d.SetPartial("additional_zones")
	} else if d.HasChange("node_locations") {
		azSetOldI, azSetNewI := d.GetChange("node_locations")
		azSetNew := azSetNewI.(*schema.Set)
		azSetOld := azSetOldI.(*schema.Set)
		if azSetNew.Contains(location) {
			return fmt.Errorf("for multi-zonal clusters, node_locations should not contain the primary 'zone'")
		}
		// Since we can't add & remove zones in the same request, first add all the
		// zones, then remove the ones we aren't using anymore.
		azSet := azSetOld.Union(azSetNew)

		if isZone(location) {
			azSet.Add(location)
		}

		req := &containerBeta.UpdateClusterRequest{
			Update: &containerBeta.ClusterUpdate{
				DesiredLocations: convertStringSet(azSet),
			},
		}

		updateF := updateFunc(req, "updating GKE cluster node locations")
		// Call update serially.
		if err := lockedCall(lockKey, updateF); err != nil {
			return err
		}

		if isZone(location) {
			azSetNew.Add(location)
		}
		if !azSet.Equal(azSetNew) {
			req = &containerBeta.UpdateClusterRequest{
				Update: &containerBeta.ClusterUpdate{
					DesiredLocations: convertStringSet(azSetNew),
				},
			}

			updateF := updateFunc(req, "updating GKE cluster node locations")
			// Call update serially.
			if err := lockedCall(lockKey, updateF); err != nil {
				return err
			}
		}

		log.Printf("[INFO] GKE cluster %s node locations have been updated to %v", d.Id(), azSet.List())

		d.SetPartial("node_locations")
	}

	if d.HasChange("enable_legacy_abac") {
		enabled := d.Get("enable_legacy_abac").(bool)
		req := &containerBeta.SetLegacyAbacRequest{
			Enabled:         enabled,
			ForceSendFields: []string{"Enabled"},
		}

		updateF := func() error {
			log.Println("[DEBUG] updating enable_legacy_abac")
			name := containerClusterFullName(project, location, clusterName)
			op, err := config.clientContainerBeta.Projects.Locations.Clusters.SetLegacyAbac(name, req).Do()
			if err != nil {
				return err
			}

			// Wait until it's updated
			err = containerOperationWait(config, op, project, location, "updating GKE legacy ABAC", timeoutInMinutes)
			log.Println("[DEBUG] done updating enable_legacy_abac")
			return err
		}

		// Call update serially.
		if err := lockedCall(lockKey, updateF); err != nil {
			return err
		}

		log.Printf("[INFO] GKE cluster %s legacy ABAC has been updated to %v", d.Id(), enabled)

		d.SetPartial("enable_legacy_abac")
	}

	if d.HasChange("monitoring_service") || d.HasChange("logging_service") {
		logging := d.Get("logging_service").(string)
		monitoring := d.Get("monitoring_service").(string)

		updateF := func() error {
			name := containerClusterFullName(project, location, clusterName)
			req := &containerBeta.UpdateClusterRequest{
				Update: &containerBeta.ClusterUpdate{
					DesiredMonitoringService: monitoring,
					DesiredLoggingService:    logging,
				},
			}
			op, err := config.clientContainerBeta.Projects.Locations.Clusters.Update(name, req).Do()
			if err != nil {
				return err
			}

			// Wait until it's updated
			return containerOperationWait(config, op, project, location, "updating GKE logging+monitoring service", timeoutInMinutes)
		}

		// Call update serially.
		if err := lockedCall(lockKey, updateF); err != nil {
			return err
		}

		log.Printf("[INFO] GKE cluster %s: logging service has been updated to %s, monitoring service has been updated to %s", d.Id(), logging, monitoring)
		d.SetPartial("logging_service")
		d.SetPartial("monitoring_service")
	}

	if d.HasChange("network_policy") {
		np := d.Get("network_policy")
		req := &containerBeta.SetNetworkPolicyRequest{
			NetworkPolicy: expandNetworkPolicy(np),
		}

		updateF := func() error {
			log.Println("[DEBUG] updating network_policy")
			name := containerClusterFullName(project, location, clusterName)
			op, err := config.clientContainerBeta.Projects.Locations.Clusters.SetNetworkPolicy(name, req).Do()
			if err != nil {
				return err
			}

			// Wait until it's updated
			err = containerOperationWait(config, op, project, location, "updating GKE cluster network policy", timeoutInMinutes)
			log.Println("[DEBUG] done updating network_policy")
			return err
		}

		// Call update serially.
		if err := lockedCall(lockKey, updateF); err != nil {
			return err
		}

		log.Printf("[INFO] Network policy for GKE cluster %s has been updated", d.Id())

		d.SetPartial("network_policy")

	}

	if n, ok := d.GetOk("node_pool.#"); ok {
		for i := 0; i < n.(int); i++ {
			nodePoolInfo, err := extractNodePoolInformationFromCluster(d, config, clusterName)
			if err != nil {
				return err
			}

			if err := nodePoolUpdate(d, meta, nodePoolInfo, fmt.Sprintf("node_pool.%d.", i), timeoutInMinutes); err != nil {
				return err
			}
		}
		d.SetPartial("node_pool")
	}

	// The master must be updated before the nodes
	if d.HasChange("min_master_version") {
		desiredMasterVersion := d.Get("min_master_version").(string)
		currentMasterVersion := d.Get("master_version").(string)
		des, err := version.NewVersion(desiredMasterVersion)
		if err != nil {
			return err
		}
		cur, err := version.NewVersion(currentMasterVersion)
		if err != nil {
			return err
		}

		// Only upgrade the master if the current version is lower than the desired version
		if cur.LessThan(des) {
			req := &containerBeta.UpdateClusterRequest{
				Update: &containerBeta.ClusterUpdate{
					DesiredMasterVersion: desiredMasterVersion,
				},
			}

			updateF := updateFunc(req, "updating GKE master version")
			// Call update serially.
			if err := lockedCall(lockKey, updateF); err != nil {
				return err
			}
			log.Printf("[INFO] GKE cluster %s: master has been updated to %s", d.Id(), desiredMasterVersion)
		}
		d.SetPartial("min_master_version")
	}

	// It's not super important that this come after updating the node pools, but it still seems like a better
	// idea than doing it before.
	if d.HasChange("node_version") {
		foundDefault := false
		if n, ok := d.GetOk("node_pool.#"); ok {
			for i := 0; i < n.(int); i++ {
				key := fmt.Sprintf("node_pool.%d.", i)
				if d.Get(key+"name").(string) == "default-pool" {
					desiredNodeVersion := d.Get("node_version").(string)
					req := &containerBeta.UpdateClusterRequest{
						Update: &containerBeta.ClusterUpdate{
							DesiredNodeVersion: desiredNodeVersion,
							DesiredNodePoolId:  "default-pool",
						},
					}
					updateF := updateFunc(req, "updating GKE default node pool node version")
					// Call update serially.
					if err := lockedCall(lockKey, updateF); err != nil {
						return err
					}
					log.Printf("[INFO] GKE cluster %s: default node pool has been updated to %s", d.Id(),
						desiredNodeVersion)
					foundDefault = true
				}
			}
		}

		if !foundDefault {
			return fmt.Errorf("node_version was updated but default-pool was not found. To update the version for a non-default pool, use the version attribute on that pool.")
		}

		d.SetPartial("node_version")
	}

	if d.HasChange("node_config") {
		if d.HasChange("node_config.0.image_type") {
			it := d.Get("node_config.0.image_type").(string)
			req := &containerBeta.UpdateClusterRequest{
				Update: &containerBeta.ClusterUpdate{
					DesiredImageType: it,
				},
			}

			updateF := func() error {
				name := containerClusterFullName(project, location, clusterName)
				op, err := config.clientContainerBeta.Projects.Locations.Clusters.Update(name, req).Do()
				if err != nil {
					return err
				}

				// Wait until it's updated
				return containerOperationWait(config, op, project, location, "updating GKE image type", timeoutInMinutes)
			}

			// Call update serially.
			if err := lockedCall(lockKey, updateF); err != nil {
				return err
			}

			log.Printf("[INFO] GKE cluster %s: image type has been updated to %s", d.Id(), it)
		}
		d.SetPartial("node_config")
	}

	if d.HasChange("master_auth") {
		var req *containerBeta.SetMasterAuthRequest
		if ma, ok := d.GetOk("master_auth"); ok {
			req = &containerBeta.SetMasterAuthRequest{
				Action: "SET_USERNAME",
				Update: expandMasterAuth(ma),
			}
		} else {
			req = &containerBeta.SetMasterAuthRequest{
				Action: "SET_USERNAME",
				Update: &containerBeta.MasterAuth{
					Username: "admin",
				},
			}
		}

		updateF := func() error {
			name := containerClusterFullName(project, location, clusterName)
			op, err := config.clientContainerBeta.Projects.Locations.Clusters.SetMasterAuth(name, req).Do()
			if err != nil {
				return err
			}

			// Wait until it's updated
			return containerOperationWait(config, op, project, location, "updating master auth", timeoutInMinutes)
		}

		// Call update serially.
		if err := lockedCall(lockKey, updateF); err != nil {
			return err
		}

		log.Printf("[INFO] GKE cluster %s: master auth has been updated", d.Id())
		d.SetPartial("master_auth")
	}

	if d.HasChange("pod_security_policy_config") {
		c := d.Get("pod_security_policy_config")
		req := &containerBeta.UpdateClusterRequest{
			Update: &containerBeta.ClusterUpdate{
				DesiredPodSecurityPolicyConfig: expandPodSecurityPolicyConfig(c),
			},
		}

		updateF := func() error {
			name := containerClusterFullName(project, location, clusterName)
			op, err := config.clientContainerBeta.Projects.Locations.Clusters.Update(name, req).Do()
			if err != nil {
				return err
			}
			// Wait until it's updated
			return containerOperationWait(config, op, project, location, "updating GKE cluster pod security policy config", timeoutInMinutes)
		}
		if err := lockedCall(lockKey, updateF); err != nil {
			return err
		}
		log.Printf("[INFO] GKE cluster %s pod security policy config has been updated", d.Id())

		d.SetPartial("pod_security_policy_config")
	}

	if d.HasChange("resource_labels") {
		resourceLabels := d.Get("resource_labels").(map[string]interface{})
		req := &containerBeta.SetLabelsRequest{
			ResourceLabels: convertStringMap(resourceLabels),
		}
		updateF := func() error {
			name := containerClusterFullName(project, location, clusterName)
			op, err := config.clientContainerBeta.Projects.Locations.Clusters.SetResourceLabels(name, req).Do()
			if err != nil {
				return err
			}

			// Wait until it's updated
			return containerOperationWait(config, op, project, location, "updating GKE resource labels", timeoutInMinutes)
		}

		// Call update serially.
		if err := lockedCall(lockKey, updateF); err != nil {
			return err
		}

		d.SetPartial("resource_labels")
	}

	if d.HasChange("remove_default_node_pool") && d.Get("remove_default_node_pool").(bool) {
		name := fmt.Sprintf("%s/nodePools/%s", containerClusterFullName(project, location, clusterName), "default-pool")
		op, err := config.clientContainerBeta.Projects.Locations.Clusters.NodePools.Delete(name).Do()
		if err != nil {
			if !isGoogleApiErrorWithCode(err, 404) {
				return errwrap.Wrapf("Error deleting default node pool: {{err}}", err)
			}
			log.Printf("[WARN] Container cluster %q default node pool already removed, no change", d.Id())
		} else {
			err = containerOperationWait(config, op, project, location, "removing default node pool", timeoutInMinutes)
			if err != nil {
				return errwrap.Wrapf("Error deleting default node pool: {{err}}", err)
			}
		}
	}

	d.Partial(false)

	return resourceContainerClusterRead(d, meta)
}

func resourceContainerClusterDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	location, err := getLocation(d, config)
	if err != nil {
		return err
	}

	clusterName := d.Get("name").(string)
	timeoutInMinutes := int(d.Timeout(schema.TimeoutDelete).Minutes())

	log.Printf("[DEBUG] Deleting GKE cluster %s", d.Get("name").(string))
	mutexKV.Lock(containerClusterMutexKey(project, location, clusterName))
	defer mutexKV.Unlock(containerClusterMutexKey(project, location, clusterName))

	var op *containerBeta.Operation
	var count = 0
	err = resource.Retry(30*time.Second, func() *resource.RetryError {
		count++

		name := containerClusterFullName(project, location, clusterName)
		op, err = config.clientContainerBeta.Projects.Locations.Clusters.Delete(name).Do()

		if err != nil {
			log.Printf("[WARNING] Cluster is still not ready to delete, retrying %s", clusterName)
			return resource.RetryableError(err)
		}

		if count == 15 {
			return resource.NonRetryableError(fmt.Errorf("Error retrying to delete cluster %s", clusterName))
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("Error deleting Cluster: %s", err)
	}

	// Wait until it's deleted
	waitErr := containerOperationWait(config, op, project, location, "deleting GKE cluster", timeoutInMinutes)
	if waitErr != nil {
		return waitErr
	}

	log.Printf("[INFO] GKE cluster %s has been deleted", d.Id())

	d.SetId("")

	return nil
}

// cleanFailedContainerCluster deletes clusters that failed but were
// created in an error state. Similar to resourceContainerClusterDelete
// but implemented in separate function as it doesn't try to lock already
// locked cluster state, does different error handling, and doesn't do retries.
func cleanFailedContainerCluster(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	location, err := getLocation(d, config)
	if err != nil {
		return err
	}

	clusterName := d.Get("name").(string)
	fullName := containerClusterFullName(project, location, clusterName)

	log.Printf("[DEBUG] Cleaning up failed GKE cluster %s", d.Get("name").(string))
	op, err := config.clientContainerBeta.Projects.Locations.Clusters.Delete(fullName).Do()
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("Container Cluster %q", d.Get("name").(string)))
	}

	// Wait until it's deleted
	timeoutInMinutes := int(d.Timeout(schema.TimeoutDelete).Minutes())
	waitErr := containerOperationWait(config, op, project, location, "deleting GKE cluster", timeoutInMinutes)
	if waitErr != nil {
		return waitErr
	}

	log.Printf("[INFO] GKE cluster %s has been deleted", d.Id())
	d.SetId("")
	return nil
}

// container engine's API currently mistakenly returns the instance group manager's
// URL instead of the instance group's URL in its responses. This shim detects that
// error, and corrects it, by fetching the instance group manager URL and retrieving
// the instance group manager, then using that to look up the instance group URL, which
// is then substituted.
//
// This should be removed when the API response is fixed.
func getInstanceGroupUrlsFromManagerUrls(config *Config, igmUrls []string) ([]string, error) {
	instanceGroupURLs := make([]string, 0, len(igmUrls))
	for _, u := range igmUrls {
		if !instanceGroupManagerURL.MatchString(u) {
			instanceGroupURLs = append(instanceGroupURLs, u)
			continue
		}
		matches := instanceGroupManagerURL.FindStringSubmatch(u)
		instanceGroupManager, err := config.clientCompute.InstanceGroupManagers.Get(matches[1], matches[2], matches[3]).Do()
		if err != nil {
			return nil, fmt.Errorf("Error reading instance group manager returned as an instance group URL: %s", err)
		}
		instanceGroupURLs = append(instanceGroupURLs, instanceGroupManager.InstanceGroup)
	}
	return instanceGroupURLs, nil
}

func expandClusterAddonsConfig(configured interface{}) *containerBeta.AddonsConfig {
	l := configured.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	config := l[0].(map[string]interface{})
	ac := &containerBeta.AddonsConfig{}

	if v, ok := config["http_load_balancing"]; ok && len(v.([]interface{})) > 0 {
		addon := v.([]interface{})[0].(map[string]interface{})
		ac.HttpLoadBalancing = &containerBeta.HttpLoadBalancing{
			Disabled:        addon["disabled"].(bool),
			ForceSendFields: []string{"Disabled"},
		}
	}

	if v, ok := config["horizontal_pod_autoscaling"]; ok && len(v.([]interface{})) > 0 {
		addon := v.([]interface{})[0].(map[string]interface{})
		ac.HorizontalPodAutoscaling = &containerBeta.HorizontalPodAutoscaling{
			Disabled:        addon["disabled"].(bool),
			ForceSendFields: []string{"Disabled"},
		}
	}

	if v, ok := config["kubernetes_dashboard"]; ok && len(v.([]interface{})) > 0 {
		addon := v.([]interface{})[0].(map[string]interface{})
		ac.KubernetesDashboard = &containerBeta.KubernetesDashboard{
			Disabled:        addon["disabled"].(bool),
			ForceSendFields: []string{"Disabled"},
		}
	}

	if v, ok := config["network_policy_config"]; ok && len(v.([]interface{})) > 0 {
		addon := v.([]interface{})[0].(map[string]interface{})
		ac.NetworkPolicyConfig = &containerBeta.NetworkPolicyConfig{
			Disabled:        addon["disabled"].(bool),
			ForceSendFields: []string{"Disabled"},
		}
	}

	if v, ok := config["istio_config"]; ok && len(v.([]interface{})) > 0 {
		addon := v.([]interface{})[0].(map[string]interface{})
		ac.IstioConfig = &containerBeta.IstioConfig{
			Disabled:        addon["disabled"].(bool),
			Auth:            addon["auth"].(string),
			ForceSendFields: []string{"Disabled"},
		}
	}

	if v, ok := config["cloudrun_config"]; ok && len(v.([]interface{})) > 0 {
		addon := v.([]interface{})[0].(map[string]interface{})
		ac.CloudRunConfig = &containerBeta.CloudRunConfig{
			Disabled:        addon["disabled"].(bool),
			ForceSendFields: []string{"Disabled"},
		}
	}

	return ac
}

func expandIPAllocationPolicy(configured interface{}) *containerBeta.IPAllocationPolicy {
	l := configured.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	config := l[0].(map[string]interface{})

	return &containerBeta.IPAllocationPolicy{
		UseIpAliases: config["use_ip_aliases"].(bool),

		CreateSubnetwork: config["create_subnetwork"].(bool),
		SubnetworkName:   config["subnetwork_name"].(string),

		ClusterIpv4CidrBlock:  config["cluster_ipv4_cidr_block"].(string),
		ServicesIpv4CidrBlock: config["services_ipv4_cidr_block"].(string),
		NodeIpv4CidrBlock:     config["node_ipv4_cidr_block"].(string),

		ClusterSecondaryRangeName:  config["cluster_secondary_range_name"].(string),
		ServicesSecondaryRangeName: config["services_secondary_range_name"].(string),
		ForceSendFields:            []string{"UseIpAliases"},
	}
}

func expandMaintenancePolicy(configured interface{}) *containerBeta.MaintenancePolicy {
	l := configured.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	maintenancePolicy := l[0].(map[string]interface{})
	dailyMaintenanceWindow := maintenancePolicy["daily_maintenance_window"].([]interface{})[0].(map[string]interface{})
	startTime := dailyMaintenanceWindow["start_time"].(string)
	return &containerBeta.MaintenancePolicy{
		Window: &containerBeta.MaintenanceWindow{
			DailyMaintenanceWindow: &containerBeta.DailyMaintenanceWindow{
				StartTime: startTime,
			},
		},
	}
}

func expandClusterAutoscaling(configured interface{}, d *schema.ResourceData) *containerBeta.ClusterAutoscaling {
	l, ok := configured.([]interface{})
	if !ok || l == nil || len(l) == 0 || l[0] == nil {
		// Before master version 1.11.2, we must send 'nil' values if autoscaling isn't
		// turned on - the cluster will return an error even if we're setting
		// EnableNodeAutoprovisioning to false.
		cmv, err := version.NewVersion(d.Get("master_version").(string))
		if err != nil {
			log.Printf("[DEBUG] Could not parse master_version into version (%q), trying min_master_version.", d.Get("master_version").(string))
			cmv, err = version.NewVersion(d.Get("min_master_version").(string))
			if err != nil {
				log.Printf("[DEBUG] Could not parse min_master_version into version (%q), assuming we are not already using cluster autoscaling.", d.Get("min_master_version").(string))
				// This deserves a little explanation.  The only reason we would ever want to send
				// `EnableNodeAutoprovisioning: false` is because we think we might need to
				// disable it (e.g. it is already enabled).  Otherwise, there is no difference
				// between sending `nil` and sending `EnableNodeAutoprovisioning: false`.
				// The only circumstance in which neither master_version nor min_master_version
				// can be parsed into version objects would be if the user has not set either one,
				// and we have not yet had a `read` call.  e.g. first-time creates, and possibly
				// some circumstance related to import.  It is probably safe to assume that
				// we are not going to be changing cluster autoscaling from on to off in those
				// circumstances.  Therefore, if we don't know what version we're running, and
				// the user has not requested cluster autoscaling, we'll fail "safe" and not touch
				// it.
				cmv, _ = version.NewVersion("0.0.0")
			}
		}
		dmv, _ := version.NewVersion("1.11.2")
		if cmv.LessThan(dmv) {
			return nil
		} else {
			return &containerBeta.ClusterAutoscaling{
				EnableNodeAutoprovisioning: false,
				ForceSendFields:            []string{"EnableNodeAutoprovisioning"},
			}
		}
	}
	r := &containerBeta.ClusterAutoscaling{}
	if config, ok := l[0].(map[string]interface{}); ok {
		r.EnableNodeAutoprovisioning = config["enabled"].(bool)
		if limits, ok := config["resource_limits"]; ok {
			if lmts, ok := limits.([]interface{}); ok {
				for _, v := range lmts {
					limit := v.(map[string]interface{})
					r.ResourceLimits = append(r.ResourceLimits, &containerBeta.ResourceLimit{
						ResourceType: limit["resource_type"].(string),
						// Here we're relying on *not* setting ForceSendFields for 0-values.
						Minimum: int64(limit["minimum"].(int)),
						Maximum: int64(limit["maximum"].(int)),
					})
				}
			}
		}
	}
	return r
}

func expandMasterAuth(configured interface{}) *containerBeta.MasterAuth {
	l := configured.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	masterAuth := l[0].(map[string]interface{})
	result := &containerBeta.MasterAuth{
		Username: masterAuth["username"].(string),
		Password: masterAuth["password"].(string),
	}
	if _, ok := masterAuth["client_certificate_config"]; ok {
		if len(masterAuth["client_certificate_config"].([]interface{})) > 0 {
			clientCertificateConfig := masterAuth["client_certificate_config"].([]interface{})[0].(map[string]interface{})
			if _, ok := clientCertificateConfig["issue_client_certificate"]; ok {
				result.ClientCertificateConfig = &containerBeta.ClientCertificateConfig{
					IssueClientCertificate: clientCertificateConfig["issue_client_certificate"].(bool),
				}
			}
		}
	}
	return result
}

func expandMasterAuthorizedNetworksConfig(configured interface{}) *containerBeta.MasterAuthorizedNetworksConfig {
	l := configured.([]interface{})
	if len(l) == 0 {
		return nil
	}
	result := &containerBeta.MasterAuthorizedNetworksConfig{
		Enabled: true,
	}
	if config, ok := l[0].(map[string]interface{}); ok {
		if _, ok := config["cidr_blocks"]; ok {
			cidrBlocks := config["cidr_blocks"].(*schema.Set).List()
			result.CidrBlocks = make([]*containerBeta.CidrBlock, 0)
			for _, v := range cidrBlocks {
				cidrBlock := v.(map[string]interface{})
				result.CidrBlocks = append(result.CidrBlocks, &containerBeta.CidrBlock{
					CidrBlock:   cidrBlock["cidr_block"].(string),
					DisplayName: cidrBlock["display_name"].(string),
				})
			}
		}
	}
	return result
}

func expandNetworkPolicy(configured interface{}) *containerBeta.NetworkPolicy {
	l := configured.([]interface{})
	if len(l) == 0 {
		return nil
	}
	result := &containerBeta.NetworkPolicy{}
	config := l[0].(map[string]interface{})
	if enabled, ok := config["enabled"]; ok && enabled.(bool) {
		result.Enabled = true
		if provider, ok := config["provider"]; ok {
			result.Provider = provider.(string)
		}
	}
	return result
}

func expandPrivateClusterConfig(configured interface{}) *containerBeta.PrivateClusterConfig {
	l := configured.([]interface{})
	if len(l) == 0 {
		return nil
	}
	config := l[0].(map[string]interface{})
	return &containerBeta.PrivateClusterConfig{
		EnablePrivateEndpoint: config["enable_private_endpoint"].(bool),
		EnablePrivateNodes:    config["enable_private_nodes"].(bool),
		MasterIpv4CidrBlock:   config["master_ipv4_cidr_block"].(string),
		ForceSendFields:       []string{"EnablePrivateEndpoint", "EnablePrivateNodes", "MasterIpv4CidrBlock"},
	}
}

func expandPodSecurityPolicyConfig(configured interface{}) *containerBeta.PodSecurityPolicyConfig {
	l := configured.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	config := l[0].(map[string]interface{})
	return &containerBeta.PodSecurityPolicyConfig{
		Enabled:         config["enabled"].(bool),
		ForceSendFields: []string{"Enabled"},
	}
}

func expandDefaultMaxPodsConstraint(v interface{}) *containerBeta.MaxPodsConstraint {
	if v == nil {
		return nil
	}

	return &containerBeta.MaxPodsConstraint{
		MaxPodsPerNode: int64(v.(int)),
	}
}

func flattenNetworkPolicy(c *containerBeta.NetworkPolicy) []map[string]interface{} {
	result := []map[string]interface{}{}
	if c != nil {
		result = append(result, map[string]interface{}{
			"enabled":  c.Enabled,
			"provider": c.Provider,
		})
	} else {
		// Explicitly set the network policy to the default.
		result = append(result, map[string]interface{}{
			"enabled":  false,
			"provider": "PROVIDER_UNSPECIFIED",
		})
	}
	return result
}

func flattenClusterAddonsConfig(c *containerBeta.AddonsConfig) []map[string]interface{} {
	result := make(map[string]interface{})
	if c == nil {
		return nil
	}
	if c.HorizontalPodAutoscaling != nil {
		result["horizontal_pod_autoscaling"] = []map[string]interface{}{
			{
				"disabled": c.HorizontalPodAutoscaling.Disabled,
			},
		}
	}
	if c.HttpLoadBalancing != nil {
		result["http_load_balancing"] = []map[string]interface{}{
			{
				"disabled": c.HttpLoadBalancing.Disabled,
			},
		}
	}
	if c.KubernetesDashboard != nil {
		result["kubernetes_dashboard"] = []map[string]interface{}{
			{
				"disabled": c.KubernetesDashboard.Disabled,
			},
		}
	}
	if c.NetworkPolicyConfig != nil {
		result["network_policy_config"] = []map[string]interface{}{
			{
				"disabled": c.NetworkPolicyConfig.Disabled,
			},
		}
	}

	if c.IstioConfig != nil {
		result["istio_config"] = []map[string]interface{}{
			{
				"disabled": c.IstioConfig.Disabled,
				"auth":     c.IstioConfig.Auth,
			},
		}
	}

	if c.CloudRunConfig != nil {
		result["cloudrun_config"] = []map[string]interface{}{
			{
				"disabled": c.CloudRunConfig.Disabled,
			},
		}
	}
	return []map[string]interface{}{result}
}

func flattenClusterNodePools(d *schema.ResourceData, config *Config, c []*containerBeta.NodePool) ([]map[string]interface{}, error) {
	nodePools := make([]map[string]interface{}, 0, len(c))

	for i, np := range c {
		nodePool, err := flattenNodePool(d, config, np, fmt.Sprintf("node_pool.%d.", i))
		if err != nil {
			return nil, err
		}
		nodePools = append(nodePools, nodePool)
	}

	return nodePools, nil
}

func flattenPrivateClusterConfig(c *containerBeta.PrivateClusterConfig) []map[string]interface{} {
	if c == nil {
		return nil
	}
	return []map[string]interface{}{
		{
			"enable_private_endpoint": c.EnablePrivateEndpoint,
			"enable_private_nodes":    c.EnablePrivateNodes,
			"master_ipv4_cidr_block":  c.MasterIpv4CidrBlock,
			"private_endpoint":        c.PrivateEndpoint,
			"public_endpoint":         c.PublicEndpoint,
		},
	}
}

func flattenIPAllocationPolicy(c *containerBeta.IPAllocationPolicy, d *schema.ResourceData, config *Config) []map[string]interface{} {
	if c == nil {
		return nil
	}
	node_cidr_block := ""
	if c.SubnetworkName != "" {
		subnetwork, err := ParseSubnetworkFieldValue(c.SubnetworkName, d, config)
		if err == nil {
			sn, err := config.clientCompute.Subnetworks.Get(subnetwork.Project, subnetwork.Region, subnetwork.Name).Do()
			if err == nil {
				node_cidr_block = sn.IpCidrRange
			}
		} else {
			log.Printf("[WARN] Unable to parse subnetwork name, got error while trying to get new subnetwork: %s", err)
		}
	}
	return []map[string]interface{}{
		{
			"use_ip_aliases": c.UseIpAliases,

			"create_subnetwork": c.CreateSubnetwork,
			"subnetwork_name":   c.SubnetworkName,

			"cluster_ipv4_cidr_block":  c.ClusterIpv4CidrBlock,
			"services_ipv4_cidr_block": c.ServicesIpv4CidrBlock,
			"node_ipv4_cidr_block":     node_cidr_block,

			"cluster_secondary_range_name":  c.ClusterSecondaryRangeName,
			"services_secondary_range_name": c.ServicesSecondaryRangeName,
		},
	}
}

func flattenMaintenancePolicy(mp *containerBeta.MaintenancePolicy) []map[string]interface{} {
	if mp == nil {
		return nil
	}
	return []map[string]interface{}{
		{
			"daily_maintenance_window": []map[string]interface{}{
				{
					"start_time": mp.Window.DailyMaintenanceWindow.StartTime,
					"duration":   mp.Window.DailyMaintenanceWindow.Duration,
				},
			},
		},
	}
}

func flattenMasterAuth(ma *containerBeta.MasterAuth) []map[string]interface{} {
	if ma == nil {
		return nil
	}
	masterAuth := []map[string]interface{}{
		{
			"username":               ma.Username,
			"password":               ma.Password,
			"client_certificate":     ma.ClientCertificate,
			"client_key":             ma.ClientKey,
			"cluster_ca_certificate": ma.ClusterCaCertificate,
		},
	}
	if len(ma.ClientCertificate) == 0 {
		masterAuth[0]["client_certificate_config"] = []map[string]interface{}{
			{"issue_client_certificate": false},
		}
	}
	return masterAuth
}

func flattenClusterAutoscaling(a *containerBeta.ClusterAutoscaling) []map[string]interface{} {
	r := make(map[string]interface{})
	if a == nil || !a.EnableNodeAutoprovisioning {
		r["enabled"] = false
	} else {
		resourceLimits := make([]interface{}, 0, len(a.ResourceLimits))
		for _, rl := range a.ResourceLimits {
			resourceLimits = append(resourceLimits, map[string]interface{}{
				"resource_type": rl.ResourceType,
				"minimum":       rl.Minimum,
				"maximum":       rl.Maximum,
			})
		}
		r["resource_limits"] = resourceLimits
		r["enabled"] = true
	}
	return []map[string]interface{}{r}
}

func flattenMasterAuthorizedNetworksConfig(c *containerBeta.MasterAuthorizedNetworksConfig) []map[string]interface{} {
	if c == nil {
		return nil
	}
	result := make(map[string]interface{})
	if c.Enabled {
		cidrBlocks := make([]interface{}, 0, len(c.CidrBlocks))
		for _, v := range c.CidrBlocks {
			cidrBlocks = append(cidrBlocks, map[string]interface{}{
				"cidr_block":   v.CidrBlock,
				"display_name": v.DisplayName,
			})
		}
		result["cidr_blocks"] = schema.NewSet(schema.HashResource(cidrBlockConfig), cidrBlocks)
	}
	return []map[string]interface{}{result}
}

func flattenPodSecurityPolicyConfig(c *containerBeta.PodSecurityPolicyConfig) []map[string]interface{} {
	if c == nil {
		return nil
	}
	return []map[string]interface{}{
		{
			"enabled": c.Enabled,
		},
	}
}

func resourceContainerClusterStateImporter(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)

	parts := strings.Split(d.Id(), "/")
	var project, location, clusterName string
	switch len(parts) {
	case 2:
		location = parts[0]
		clusterName = parts[1]
	case 3:
		project = parts[0]
		location = parts[1]
		clusterName = parts[2]
	default:
		return nil, fmt.Errorf("Invalid container cluster specifier. Expecting {location}/{name} or {project}/{location}/{name}")
	}

	if len(project) == 0 {
		var err error
		project, err = getProject(d, config)
		if err != nil {
			return nil, err
		}
	}
	d.Set("project", project)

	d.Set("location", location)
	if isZone(location) {
		d.Set("zone", location)
	} else {
		d.Set("region", location)
	}

	d.Set("name", clusterName)
	d.SetId(clusterName)

	return []*schema.ResourceData{d}, nil
}

func containerClusterMutexKey(project, location, clusterName string) string {
	return fmt.Sprintf("google-container-cluster/%s/%s/%s", project, location, clusterName)
}

func containerClusterFullName(project, location, cluster string) string {
	return fmt.Sprintf("projects/%s/locations/%s/clusters/%s", project, location, cluster)
}

func extractNodePoolInformationFromCluster(d *schema.ResourceData, config *Config, clusterName string) (*NodePoolInformation, error) {
	project, err := getProject(d, config)
	if err != nil {
		return nil, err
	}

	location, err := getLocation(d, config)
	if err != nil {
		return nil, err
	}

	return &NodePoolInformation{
		project:  project,
		location: location,
		cluster:  d.Get("name").(string),
	}, nil
}

func cidrOrSizeDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	// If the user specified a size and the API returned a full cidr block, suppress.
	return strings.HasPrefix(new, "/") && strings.HasSuffix(old, new)
}

// We want to suppress diffs for empty or default client certificate configs, i.e:
// 	[{ "issue_client_certificate": true}] --> []
//  [] -> [{ "issue_client_certificate": true}]
func masterAuthClientCertCfgSuppress(k, old, new string, r *schema.ResourceData) bool {
	var clientConfig map[string]interface{}
	if v, ok := r.GetOk("master_auth"); ok {
		masterAuths := v.([]interface{})
		masterAuth := masterAuths[0].(map[string]interface{})
		cfgs := masterAuth["client_certificate_config"].([]interface{})
		if len(cfgs) > 0 {
			clientConfig = cfgs[0].(map[string]interface{})
		}
	}

	if strings.HasSuffix(k, "client_certificate_config.#") && old == "0" && new == "1" {
		// nil --> { "issue_client_certificate": true }
		if issueCert, ok := clientConfig["issue_client_certificate"]; ok {
			return issueCert.(bool)
		}
	}

	return strings.HasSuffix(k, ".issue_client_certificate") && old == "" && new == "true"
}

// We want to suppress diffs for empty/disabled private cluster config.
func containerClusterPrivateClusterConfigSuppress(k, old, new string, d *schema.ResourceData) bool {
	o, n := d.GetChange("private_cluster_config.0.enable_private_endpoint")
	suppressEndpoint := !o.(bool) && !n.(bool)

	o, n = d.GetChange("private_cluster_config.0.enable_private_nodes")
	suppressNodes := !o.(bool) && !n.(bool)

	if k == "private_cluster_config.0.enable_private_endpoint" {
		return suppressEndpoint
	} else if k == "private_cluster_config.0.enable_private_nodes" {
		return suppressNodes
	} else if k == "private_cluster_config.#" {
		return suppressEndpoint && suppressNodes
	}
	return false
}

func podSecurityPolicyCfgSuppress(k, old, new string, r *schema.ResourceData) bool {
	if k == "pod_security_policy_config.#" && old == "1" && new == "0" {
		if v, ok := r.GetOk("pod_security_policy_config"); ok {
			cfgList := v.([]interface{})
			if len(cfgList) > 0 {
				d := cfgList[0].(map[string]interface{})
				// Suppress if old value was {enabled == false}
				return !d["enabled"].(bool)
			}
		}
	}
	return false
}
