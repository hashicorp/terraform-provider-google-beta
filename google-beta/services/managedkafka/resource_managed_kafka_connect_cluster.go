// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/managedkafka/ConnectCluster.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package managedkafka

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
)

func ResourceManagedKafkaConnectCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceManagedKafkaConnectClusterCreate,
		Read:   resourceManagedKafkaConnectClusterRead,
		Update: resourceManagedKafkaConnectClusterUpdate,
		Delete: resourceManagedKafkaConnectClusterDelete,

		Importer: &schema.ResourceImporter{
			State: resourceManagedKafkaConnectClusterImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"capacity_config": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `A capacity configuration of a Kafka cluster.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"memory_bytes": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The memory to provision for the cluster in bytes. The CPU:memory ratio (vCPU:GiB) must be between 1:1 and 1:8. Minimum: 3221225472 (3 GiB).`,
						},
						"vcpu_count": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The number of vCPUs to provision for the cluster. The minimum is 3.`,
						},
					},
				},
			},
			"connect_cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The ID to use for the Connect Cluster, which will become the final component of the connect cluster's name. This value is structured like: 'my-connect-cluster-id'.`,
			},
			"gcp_config": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `Configuration properties for a Kafka Connect cluster deployed to Google Cloud Platform.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_config": {
							Type:        schema.TypeList,
							Required:    true,
							Description: `The configuration of access to the Kafka Connect cluster.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"network_configs": {
										Type:        schema.TypeList,
										Required:    true,
										Description: `Virtual Private Cloud (VPC) subnets where IP addresses for the Kafka Connect cluster are allocated. To make the connect cluster available in a VPC, you must specify at least one subnet per network. You must specify between 1 and 10 subnets. Additional subnets may be specified with additional 'network_configs' blocks.`,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"primary_subnet": {
													Type:             schema.TypeString,
													Required:         true,
													DiffSuppressFunc: tpgresource.ProjectNumberDiffSuppress,
													Description:      `VPC subnet to make available to the Kafka Connect cluster. Structured like: projects/{project}/regions/{region}/subnetworks/{subnet_id}. It is used to create a Private Service Connect (PSC) interface for the Kafka Connect workers. It must be located in the same region as the Kafka Connect cluster. The CIDR range of the subnet must be within the IPv4 address ranges for private networks, as specified in RFC 1918. The primary subnet CIDR range must have a minimum size of /22 (1024 addresses).`,
												},
												"additional_subnets": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: `Additional subnets may be specified. They may be in another region, but must be in the same VPC network. The Connect workers can communicate with network endpoints in either the primary or additional subnets.`,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"dns_domain_names": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: `Additional DNS domain names from the subnet's network to be made visible to the Connect Cluster. When using MirrorMaker2, it's necessary to add the bootstrap address's dns domain name of the target cluster to make it visible to the connector. For example: my-kafka-cluster.us-central1.managedkafka.my-project.cloud.goog`,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"kafka_cluster": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The name of the Kafka cluster this Kafka Connect cluster is attached to. Structured like: 'projects/PROJECT_ID/locations/LOCATION/clusters/CLUSTER_ID'.`,
			},
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `ID of the location of the Kafka Connect resource. See https://cloud.google.com/managed-kafka/docs/locations for a list of supported locations.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `List of label KEY=VALUE pairs to add. Keys must start with a lowercase character and contain only hyphens (-), underscores ( ), lowercase characters, and numbers. Values must contain only hyphens (-), underscores ( ), lowercase characters, and numbers.

**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time when the cluster was created.`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The name of the connect cluster. Structured like: 'projects/PROJECT_ID/locations/LOCATION/connectClusters/CONNECT_CLUSTER_ID'.`,
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The current state of the connect cluster. Possible values: 'STATE_UNSPECIFIED', 'CREATING', 'ACTIVE', 'DELETING'.`,
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time when the cluster was last updated.`,
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

func resourceManagedKafkaConnectClusterCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	kafkaClusterProp, err := expandManagedKafkaConnectClusterKafkaCluster(d.Get("kafka_cluster"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("kafka_cluster"); !tpgresource.IsEmptyValue(reflect.ValueOf(kafkaClusterProp)) && (ok || !reflect.DeepEqual(v, kafkaClusterProp)) {
		obj["kafkaCluster"] = kafkaClusterProp
	}
	capacityConfigProp, err := expandManagedKafkaConnectClusterCapacityConfig(d.Get("capacity_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("capacity_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(capacityConfigProp)) && (ok || !reflect.DeepEqual(v, capacityConfigProp)) {
		obj["capacityConfig"] = capacityConfigProp
	}
	gcpConfigProp, err := expandManagedKafkaConnectClusterGcpConfig(d.Get("gcp_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("gcp_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(gcpConfigProp)) && (ok || !reflect.DeepEqual(v, gcpConfigProp)) {
		obj["gcpConfig"] = gcpConfigProp
	}
	labelsProp, err := expandManagedKafkaConnectClusterEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ManagedKafkaBasePath}}projects/{{project}}/locations/{{location}}/connectClusters?connectClusterId={{connect_cluster_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ConnectCluster: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ConnectCluster: %s", err)
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
		return fmt.Errorf("Error creating ConnectCluster: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/connectClusters/{{connect_cluster_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = ManagedKafkaOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating ConnectCluster", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create ConnectCluster: %s", err)
	}

	if err := d.Set("name", flattenManagedKafkaConnectClusterName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/connectClusters/{{connect_cluster_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating ConnectCluster %q: %#v", d.Id(), res)

	return resourceManagedKafkaConnectClusterRead(d, meta)
}

func resourceManagedKafkaConnectClusterRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ManagedKafkaBasePath}}projects/{{project}}/locations/{{location}}/connectClusters/{{connect_cluster_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ConnectCluster: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ManagedKafkaConnectCluster %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ConnectCluster: %s", err)
	}

	if err := d.Set("name", flattenManagedKafkaConnectClusterName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading ConnectCluster: %s", err)
	}
	if err := d.Set("kafka_cluster", flattenManagedKafkaConnectClusterKafkaCluster(res["kafkaCluster"], d, config)); err != nil {
		return fmt.Errorf("Error reading ConnectCluster: %s", err)
	}
	if err := d.Set("create_time", flattenManagedKafkaConnectClusterCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading ConnectCluster: %s", err)
	}
	if err := d.Set("update_time", flattenManagedKafkaConnectClusterUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading ConnectCluster: %s", err)
	}
	if err := d.Set("labels", flattenManagedKafkaConnectClusterLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading ConnectCluster: %s", err)
	}
	if err := d.Set("capacity_config", flattenManagedKafkaConnectClusterCapacityConfig(res["capacityConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading ConnectCluster: %s", err)
	}
	if err := d.Set("gcp_config", flattenManagedKafkaConnectClusterGcpConfig(res["gcpConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading ConnectCluster: %s", err)
	}
	if err := d.Set("state", flattenManagedKafkaConnectClusterState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading ConnectCluster: %s", err)
	}
	if err := d.Set("terraform_labels", flattenManagedKafkaConnectClusterTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading ConnectCluster: %s", err)
	}
	if err := d.Set("effective_labels", flattenManagedKafkaConnectClusterEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading ConnectCluster: %s", err)
	}

	return nil
}

func resourceManagedKafkaConnectClusterUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ConnectCluster: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	kafkaClusterProp, err := expandManagedKafkaConnectClusterKafkaCluster(d.Get("kafka_cluster"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("kafka_cluster"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, kafkaClusterProp)) {
		obj["kafkaCluster"] = kafkaClusterProp
	}
	capacityConfigProp, err := expandManagedKafkaConnectClusterCapacityConfig(d.Get("capacity_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("capacity_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, capacityConfigProp)) {
		obj["capacityConfig"] = capacityConfigProp
	}
	gcpConfigProp, err := expandManagedKafkaConnectClusterGcpConfig(d.Get("gcp_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("gcp_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, gcpConfigProp)) {
		obj["gcpConfig"] = gcpConfigProp
	}
	labelsProp, err := expandManagedKafkaConnectClusterEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ManagedKafkaBasePath}}projects/{{project}}/locations/{{location}}/connectClusters/{{connect_cluster_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ConnectCluster %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("kafka_cluster") {
		updateMask = append(updateMask, "kafkaCluster")
	}

	if d.HasChange("capacity_config") {
		updateMask = append(updateMask, "capacityConfig")
	}

	if d.HasChange("gcp_config") {
		updateMask = append(updateMask, "gcpConfig")
	}

	if d.HasChange("effective_labels") {
		updateMask = append(updateMask, "labels")
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
			return fmt.Errorf("Error updating ConnectCluster %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating ConnectCluster %q: %#v", d.Id(), res)
		}

		err = ManagedKafkaOperationWaitTime(
			config, res, project, "Updating ConnectCluster", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceManagedKafkaConnectClusterRead(d, meta)
}

func resourceManagedKafkaConnectClusterDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ConnectCluster: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{ManagedKafkaBasePath}}projects/{{project}}/locations/{{location}}/connectClusters/{{connect_cluster_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting ConnectCluster %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "ConnectCluster")
	}

	err = ManagedKafkaOperationWaitTime(
		config, res, project, "Deleting ConnectCluster", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting ConnectCluster %q: %#v", d.Id(), res)
	return nil
}

func resourceManagedKafkaConnectClusterImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/connectClusters/(?P<connect_cluster_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<connect_cluster_id>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<connect_cluster_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/connectClusters/{{connect_cluster_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenManagedKafkaConnectClusterName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenManagedKafkaConnectClusterKafkaCluster(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenManagedKafkaConnectClusterCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenManagedKafkaConnectClusterUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenManagedKafkaConnectClusterLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenManagedKafkaConnectClusterCapacityConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["vcpu_count"] =
		flattenManagedKafkaConnectClusterCapacityConfigVcpuCount(original["vcpuCount"], d, config)
	transformed["memory_bytes"] =
		flattenManagedKafkaConnectClusterCapacityConfigMemoryBytes(original["memoryBytes"], d, config)
	return []interface{}{transformed}
}
func flattenManagedKafkaConnectClusterCapacityConfigVcpuCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenManagedKafkaConnectClusterCapacityConfigMemoryBytes(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenManagedKafkaConnectClusterGcpConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["access_config"] =
		flattenManagedKafkaConnectClusterGcpConfigAccessConfig(original["accessConfig"], d, config)
	return []interface{}{transformed}
}
func flattenManagedKafkaConnectClusterGcpConfigAccessConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["network_configs"] =
		flattenManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigs(original["networkConfigs"], d, config)
	return []interface{}{transformed}
}
func flattenManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigs(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"primary_subnet":     flattenManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsPrimarySubnet(original["primarySubnet"], d, config),
			"additional_subnets": flattenManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsAdditionalSubnets(original["additionalSubnets"], d, config),
			"dns_domain_names":   flattenManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsDnsDomainNames(original["dnsDomainNames"], d, config),
		})
	}
	return transformed
}
func flattenManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsPrimarySubnet(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsAdditionalSubnets(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsDnsDomainNames(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenManagedKafkaConnectClusterState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenManagedKafkaConnectClusterTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("terraform_labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenManagedKafkaConnectClusterEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandManagedKafkaConnectClusterKafkaCluster(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandManagedKafkaConnectClusterCapacityConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedVcpuCount, err := expandManagedKafkaConnectClusterCapacityConfigVcpuCount(original["vcpu_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVcpuCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["vcpuCount"] = transformedVcpuCount
	}

	transformedMemoryBytes, err := expandManagedKafkaConnectClusterCapacityConfigMemoryBytes(original["memory_bytes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMemoryBytes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["memoryBytes"] = transformedMemoryBytes
	}

	return transformed, nil
}

func expandManagedKafkaConnectClusterCapacityConfigVcpuCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandManagedKafkaConnectClusterCapacityConfigMemoryBytes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandManagedKafkaConnectClusterGcpConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAccessConfig, err := expandManagedKafkaConnectClusterGcpConfigAccessConfig(original["access_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAccessConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["accessConfig"] = transformedAccessConfig
	}

	return transformed, nil
}

func expandManagedKafkaConnectClusterGcpConfigAccessConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNetworkConfigs, err := expandManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigs(original["network_configs"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNetworkConfigs); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["networkConfigs"] = transformedNetworkConfigs
	}

	return transformed, nil
}

func expandManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedPrimarySubnet, err := expandManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsPrimarySubnet(original["primary_subnet"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPrimarySubnet); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["primarySubnet"] = transformedPrimarySubnet
		}

		transformedAdditionalSubnets, err := expandManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsAdditionalSubnets(original["additional_subnets"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAdditionalSubnets); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["additionalSubnets"] = transformedAdditionalSubnets
		}

		transformedDnsDomainNames, err := expandManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsDnsDomainNames(original["dns_domain_names"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDnsDomainNames); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["dnsDomainNames"] = transformedDnsDomainNames
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsPrimarySubnet(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsAdditionalSubnets(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsDnsDomainNames(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandManagedKafkaConnectClusterEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
