// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
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

package redis

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceRedisCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceRedisClusterCreate,
		Read:   resourceRedisClusterRead,
		Update: resourceRedisClusterUpdate,
		Delete: resourceRedisClusterDelete,

		Importer: &schema.ResourceImporter{
			State: resourceRedisClusterImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(120 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
			tpgresource.DefaultProviderRegion,
		),

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
				Description: `Unique name of the resource in this scope including project and location using the form:
projects/{projectId}/locations/{locationId}/clusters/{clusterId}`,
			},
			"psc_configs": {
				Type:     schema.TypeList,
				Required: true,
				Description: `Required. Each PscConfig configures the consumer network where two
network addresses will be designated to the cluster for client access.
Currently, only one PscConfig is supported.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"network": {
							Type:     schema.TypeString,
							Required: true,
							Description: `Required. The consumer network where the network address of
the discovery endpoint will be reserved, in the form of
projects/{network_project_id_or_number}/global/networks/{network_id}.`,
						},
					},
				},
			},
			"shard_count": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: `Required. Number of shards for the Redis cluster.`,
			},
			"authorization_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"AUTH_MODE_UNSPECIFIED", "AUTH_MODE_IAM_AUTH", "AUTH_MODE_DISABLED", ""}),
				Description:  `Optional. The authorization mode of the Redis cluster. If not provided, auth feature is disabled for the cluster. Default value: "AUTH_MODE_DISABLED" Possible values: ["AUTH_MODE_UNSPECIFIED", "AUTH_MODE_IAM_AUTH", "AUTH_MODE_DISABLED"]`,
				Default:      "AUTH_MODE_DISABLED",
			},
			"node_type": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"REDIS_SHARED_CORE_NANO", "REDIS_HIGHMEM_MEDIUM", "REDIS_HIGHMEM_XLARGE", "REDIS_STANDARD_SMALL", ""}),
				Description: `The nodeType for the Redis cluster.
If not provided, REDIS_HIGHMEM_MEDIUM will be used as default Possible values: ["REDIS_SHARED_CORE_NANO", "REDIS_HIGHMEM_MEDIUM", "REDIS_HIGHMEM_XLARGE", "REDIS_STANDARD_SMALL"]`,
			},
			"redis_configs": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Configure Redis Cluster behavior using a subset of native Redis configuration parameters.
Please check Memorystore documentation for the list of supported parameters:
https://cloud.google.com/memorystore/docs/cluster/supported-instance-configurations`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"region": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `The name of the region of the Redis cluster.`,
			},
			"replica_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: `Optional. The number of replica nodes per shard.`,
			},
			"transit_encryption_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"TRANSIT_ENCRYPTION_MODE_UNSPECIFIED", "TRANSIT_ENCRYPTION_MODE_DISABLED", "TRANSIT_ENCRYPTION_MODE_SERVER_AUTHENTICATION", ""}),
				Description: `Optional. The in-transit encryption for the Redis cluster.
If not provided, encryption is disabled for the cluster. Default value: "TRANSIT_ENCRYPTION_MODE_DISABLED" Possible values: ["TRANSIT_ENCRYPTION_MODE_UNSPECIFIED", "TRANSIT_ENCRYPTION_MODE_DISABLED", "TRANSIT_ENCRYPTION_MODE_SERVER_AUTHENTICATION"]`,
				Default: "TRANSIT_ENCRYPTION_MODE_DISABLED",
			},
			"zone_distribution_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `Immutable. Zone distribution config for Memorystore Redis cluster.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": {
							Type:         schema.TypeString,
							Computed:     true,
							Optional:     true,
							ValidateFunc: verify.ValidateEnum([]string{"MULTI_ZONE", "SINGLE_ZONE", ""}),
							Description: `Immutable. The mode for zone distribution for Memorystore Redis cluster.
If not provided, MULTI_ZONE will be used as default Possible values: ["MULTI_ZONE", "SINGLE_ZONE"]`,
						},
						"zone": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Immutable. The zone for single zone Memorystore Redis cluster.`,
						},
					},
				},
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The timestamp associated with the cluster creation request. A timestamp in
RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional
digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".`,
			},
			"discovery_endpoints": {
				Type:     schema.TypeList,
				Computed: true,
				Description: `Output only. Endpoints created on each given network,
for Redis clients to connect to the cluster.
Currently only one endpoint is supported.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Output only. Network address of the exposed Redis endpoint used by clients to connect to the service.`,
						},
						"port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: `Output only. The port number of the exposed Redis endpoint.`,
						},
						"psc_config": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `Output only. Customer configuration for where the endpoint
is created and accessed from.`,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"network": {
										Type:     schema.TypeString,
										Optional: true,
										Description: `The consumer network where the network address of the discovery
endpoint will be reserved, in the form of
projects/{network_project_id}/global/networks/{network_id}.`,
									},
								},
							},
						},
					},
				},
			},
			"precise_size_gb": {
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: `Output only. Redis memory precise size in GB for the entire cluster.`,
			},
			"psc_connections": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `Output only. PSC connections for discovery of the cluster topology and accessing the cluster.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Output only. The IP allocated on the consumer network for the PSC forwarding rule.`,
						},
						"forwarding_rule": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Output only. The URI of the consumer side forwarding rule. Example: projects/{projectNumOrId}/regions/us-east1/forwardingRules/{resourceId}.`,
						},
						"network": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `The consumer network where the IP address resides, in the form of projects/{projectId}/global/networks/{network_id}.`,
						},
						"project_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Output only. The consumer projectId where the forwarding rule is created from.`,
						},
						"psc_connection_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Output only. The PSC connection id of the forwarding rule connected to the service attachment.`,
						},
					},
				},
			},
			"size_gb": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `Output only. Redis memory size in GB for the entire cluster.`,
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The current state of this cluster. Can be CREATING, READY, UPDATING, DELETING and SUSPENDED`,
			},
			"state_info": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `Output only. Additional information about the current state of the cluster.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"update_info": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `A nested object resource`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"target_replica_count": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: `Target number of replica nodes per shard.`,
									},
									"target_shard_count": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: `Target number of shards for redis cluster.`,
									},
								},
							},
						},
					},
				},
			},
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `System assigned, unique identifier for the cluster.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceRedisClusterCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	authorizationModeProp, err := expandRedisClusterAuthorizationMode(d.Get("authorization_mode"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("authorization_mode"); !tpgresource.IsEmptyValue(reflect.ValueOf(authorizationModeProp)) && (ok || !reflect.DeepEqual(v, authorizationModeProp)) {
		obj["authorizationMode"] = authorizationModeProp
	}
	transitEncryptionModeProp, err := expandRedisClusterTransitEncryptionMode(d.Get("transit_encryption_mode"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("transit_encryption_mode"); !tpgresource.IsEmptyValue(reflect.ValueOf(transitEncryptionModeProp)) && (ok || !reflect.DeepEqual(v, transitEncryptionModeProp)) {
		obj["transitEncryptionMode"] = transitEncryptionModeProp
	}
	nodeTypeProp, err := expandRedisClusterNodeType(d.Get("node_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("node_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(nodeTypeProp)) && (ok || !reflect.DeepEqual(v, nodeTypeProp)) {
		obj["nodeType"] = nodeTypeProp
	}
	zoneDistributionConfigProp, err := expandRedisClusterZoneDistributionConfig(d.Get("zone_distribution_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("zone_distribution_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(zoneDistributionConfigProp)) && (ok || !reflect.DeepEqual(v, zoneDistributionConfigProp)) {
		obj["zoneDistributionConfig"] = zoneDistributionConfigProp
	}
	pscConfigsProp, err := expandRedisClusterPscConfigs(d.Get("psc_configs"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("psc_configs"); !tpgresource.IsEmptyValue(reflect.ValueOf(pscConfigsProp)) && (ok || !reflect.DeepEqual(v, pscConfigsProp)) {
		obj["pscConfigs"] = pscConfigsProp
	}
	replicaCountProp, err := expandRedisClusterReplicaCount(d.Get("replica_count"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("replica_count"); ok || !reflect.DeepEqual(v, replicaCountProp) {
		obj["replicaCount"] = replicaCountProp
	}
	shardCountProp, err := expandRedisClusterShardCount(d.Get("shard_count"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("shard_count"); !tpgresource.IsEmptyValue(reflect.ValueOf(shardCountProp)) && (ok || !reflect.DeepEqual(v, shardCountProp)) {
		obj["shardCount"] = shardCountProp
	}
	redisConfigsProp, err := expandRedisClusterRedisConfigs(d.Get("redis_configs"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("redis_configs"); !tpgresource.IsEmptyValue(reflect.ValueOf(redisConfigsProp)) && (ok || !reflect.DeepEqual(v, redisConfigsProp)) {
		obj["redisConfigs"] = redisConfigsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{RedisBasePath}}projects/{{project}}/locations/{{region}}/clusters?clusterId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Cluster: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Cluster: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating Cluster: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/clusters/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = RedisOperationWaitTime(
		config, res, project, "Creating Cluster", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Cluster: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Cluster %q: %#v", d.Id(), res)

	return resourceRedisClusterRead(d, meta)
}

func resourceRedisClusterRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{RedisBasePath}}projects/{{project}}/locations/{{region}}/clusters/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Cluster: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("RedisCluster %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Cluster: %s", err)
	}

	region, err := tpgresource.GetRegion(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("region", region); err != nil {
		return fmt.Errorf("Error reading Cluster: %s", err)
	}

	if err := d.Set("create_time", flattenRedisClusterCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Cluster: %s", err)
	}
	if err := d.Set("state", flattenRedisClusterState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading Cluster: %s", err)
	}
	if err := d.Set("uid", flattenRedisClusterUid(res["uid"], d, config)); err != nil {
		return fmt.Errorf("Error reading Cluster: %s", err)
	}
	if err := d.Set("authorization_mode", flattenRedisClusterAuthorizationMode(res["authorizationMode"], d, config)); err != nil {
		return fmt.Errorf("Error reading Cluster: %s", err)
	}
	if err := d.Set("transit_encryption_mode", flattenRedisClusterTransitEncryptionMode(res["transitEncryptionMode"], d, config)); err != nil {
		return fmt.Errorf("Error reading Cluster: %s", err)
	}
	if err := d.Set("node_type", flattenRedisClusterNodeType(res["nodeType"], d, config)); err != nil {
		return fmt.Errorf("Error reading Cluster: %s", err)
	}
	if err := d.Set("zone_distribution_config", flattenRedisClusterZoneDistributionConfig(res["zoneDistributionConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading Cluster: %s", err)
	}
	if err := d.Set("discovery_endpoints", flattenRedisClusterDiscoveryEndpoints(res["discoveryEndpoints"], d, config)); err != nil {
		return fmt.Errorf("Error reading Cluster: %s", err)
	}
	if err := d.Set("psc_connections", flattenRedisClusterPscConnections(res["pscConnections"], d, config)); err != nil {
		return fmt.Errorf("Error reading Cluster: %s", err)
	}
	if err := d.Set("state_info", flattenRedisClusterStateInfo(res["stateInfo"], d, config)); err != nil {
		return fmt.Errorf("Error reading Cluster: %s", err)
	}
	if err := d.Set("replica_count", flattenRedisClusterReplicaCount(res["replicaCount"], d, config)); err != nil {
		return fmt.Errorf("Error reading Cluster: %s", err)
	}
	if err := d.Set("size_gb", flattenRedisClusterSizeGb(res["sizeGb"], d, config)); err != nil {
		return fmt.Errorf("Error reading Cluster: %s", err)
	}
	if err := d.Set("precise_size_gb", flattenRedisClusterPreciseSizeGb(res["preciseSizeGb"], d, config)); err != nil {
		return fmt.Errorf("Error reading Cluster: %s", err)
	}
	if err := d.Set("shard_count", flattenRedisClusterShardCount(res["shardCount"], d, config)); err != nil {
		return fmt.Errorf("Error reading Cluster: %s", err)
	}
	if err := d.Set("redis_configs", flattenRedisClusterRedisConfigs(res["redisConfigs"], d, config)); err != nil {
		return fmt.Errorf("Error reading Cluster: %s", err)
	}

	return nil
}

func resourceRedisClusterUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Cluster: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	pscConfigsProp, err := expandRedisClusterPscConfigs(d.Get("psc_configs"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("psc_configs"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, pscConfigsProp)) {
		obj["pscConfigs"] = pscConfigsProp
	}
	replicaCountProp, err := expandRedisClusterReplicaCount(d.Get("replica_count"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("replica_count"); ok || !reflect.DeepEqual(v, replicaCountProp) {
		obj["replicaCount"] = replicaCountProp
	}
	shardCountProp, err := expandRedisClusterShardCount(d.Get("shard_count"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("shard_count"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, shardCountProp)) {
		obj["shardCount"] = shardCountProp
	}
	redisConfigsProp, err := expandRedisClusterRedisConfigs(d.Get("redis_configs"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("redis_configs"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, redisConfigsProp)) {
		obj["redisConfigs"] = redisConfigsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{RedisBasePath}}projects/{{project}}/locations/{{region}}/clusters/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Cluster %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("psc_configs") {
		updateMask = append(updateMask, "pscConfigs")
	}

	if d.HasChange("replica_count") {
		updateMask = append(updateMask, "replicaCount")
	}

	if d.HasChange("shard_count") {
		updateMask = append(updateMask, "shardCount")
	}

	if d.HasChange("redis_configs") {
		updateMask = append(updateMask, "redisConfigs")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// if updateMask is empty we are not updating anything so skip the post
	if len(updateMask) > 0 {
		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "PATCH",
			Project:   billingProject,
			RawURL:    url,
			UserAgent: userAgent,
			Body:      obj,
			Timeout:   d.Timeout(schema.TimeoutUpdate),
			Headers:   headers,
		})

		if err != nil {
			return fmt.Errorf("Error updating Cluster %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating Cluster %q: %#v", d.Id(), res)
		}

		err = RedisOperationWaitTime(
			config, res, project, "Updating Cluster", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceRedisClusterRead(d, meta)
}

func resourceRedisClusterDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Cluster: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{RedisBasePath}}projects/{{project}}/locations/{{region}}/clusters/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting Cluster %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Cluster")
	}

	err = RedisOperationWaitTime(
		config, res, project, "Deleting Cluster", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Cluster %q: %#v", d.Id(), res)
	return nil
}

func resourceRedisClusterImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/clusters/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<region>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/clusters/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenRedisClusterCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenRedisClusterState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenRedisClusterUid(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenRedisClusterAuthorizationMode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenRedisClusterTransitEncryptionMode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenRedisClusterNodeType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenRedisClusterZoneDistributionConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["mode"] =
		flattenRedisClusterZoneDistributionConfigMode(original["mode"], d, config)
	transformed["zone"] =
		flattenRedisClusterZoneDistributionConfigZone(original["zone"], d, config)
	return []interface{}{transformed}
}
func flattenRedisClusterZoneDistributionConfigMode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenRedisClusterZoneDistributionConfigZone(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenRedisClusterDiscoveryEndpoints(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"address":    flattenRedisClusterDiscoveryEndpointsAddress(original["address"], d, config),
			"port":       flattenRedisClusterDiscoveryEndpointsPort(original["port"], d, config),
			"psc_config": flattenRedisClusterDiscoveryEndpointsPscConfig(original["pscConfig"], d, config),
		})
	}
	return transformed
}
func flattenRedisClusterDiscoveryEndpointsAddress(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenRedisClusterDiscoveryEndpointsPort(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenRedisClusterDiscoveryEndpointsPscConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["network"] =
		flattenRedisClusterDiscoveryEndpointsPscConfigNetwork(original["network"], d, config)
	return []interface{}{transformed}
}
func flattenRedisClusterDiscoveryEndpointsPscConfigNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenRedisClusterPscConnections(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"psc_connection_id": flattenRedisClusterPscConnectionsPscConnectionId(original["pscConnectionId"], d, config),
			"address":           flattenRedisClusterPscConnectionsAddress(original["address"], d, config),
			"forwarding_rule":   flattenRedisClusterPscConnectionsForwardingRule(original["forwardingRule"], d, config),
			"project_id":        flattenRedisClusterPscConnectionsProjectId(original["projectId"], d, config),
			"network":           flattenRedisClusterPscConnectionsNetwork(original["network"], d, config),
		})
	}
	return transformed
}
func flattenRedisClusterPscConnectionsPscConnectionId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenRedisClusterPscConnectionsAddress(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenRedisClusterPscConnectionsForwardingRule(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenRedisClusterPscConnectionsProjectId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenRedisClusterPscConnectionsNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenRedisClusterStateInfo(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["update_info"] =
		flattenRedisClusterStateInfoUpdateInfo(original["updateInfo"], d, config)
	return []interface{}{transformed}
}
func flattenRedisClusterStateInfoUpdateInfo(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["target_shard_count"] =
		flattenRedisClusterStateInfoUpdateInfoTargetShardCount(original["targetShardCount"], d, config)
	transformed["target_replica_count"] =
		flattenRedisClusterStateInfoUpdateInfoTargetReplicaCount(original["targetReplicaCount"], d, config)
	return []interface{}{transformed}
}
func flattenRedisClusterStateInfoUpdateInfoTargetShardCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenRedisClusterStateInfoUpdateInfoTargetReplicaCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenRedisClusterReplicaCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenRedisClusterSizeGb(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenRedisClusterPreciseSizeGb(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenRedisClusterShardCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenRedisClusterRedisConfigs(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandRedisClusterAuthorizationMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterTransitEncryptionMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterNodeType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterZoneDistributionConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMode, err := expandRedisClusterZoneDistributionConfigMode(original["mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["mode"] = transformedMode
	}

	transformedZone, err := expandRedisClusterZoneDistributionConfigZone(original["zone"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedZone); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["zone"] = transformedZone
	}

	return transformed, nil
}

func expandRedisClusterZoneDistributionConfigMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterZoneDistributionConfigZone(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterPscConfigs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedNetwork, err := expandRedisClusterPscConfigsNetwork(original["network"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedNetwork); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["network"] = transformedNetwork
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandRedisClusterPscConfigsNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterReplicaCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterShardCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterRedisConfigs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
