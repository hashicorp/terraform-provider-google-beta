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

package networksecurity

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceNetworkSecurityServerTlsPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkSecurityServerTlsPolicyCreate,
		Read:   resourceNetworkSecurityServerTlsPolicyRead,
		Update: resourceNetworkSecurityServerTlsPolicyUpdate,
		Delete: resourceNetworkSecurityServerTlsPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceNetworkSecurityServerTlsPolicyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Name of the ServerTlsPolicy resource.`,
			},
			"allow_open": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: `This field applies only for Traffic Director policies. It is must be set to false for external HTTPS load balancer policies.
Determines if server allows plaintext connections. If set to true, server allows plain text connections. By default, it is set to false. This setting is not exclusive of other encryption modes. For example, if allowOpen and mtlsPolicy are set, server allows both plain text and mTLS connections. See documentation of other encryption modes to confirm compatibility.
Consider using it if you wish to upgrade in place your deployment to TLS while having mixed TLS and non-TLS traffic reaching port :80.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A free-text description of the resource. Max length 1024 characters.`,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `Set of label tags associated with the ServerTlsPolicy resource.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"location": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `The location of the server tls policy.
The default value is 'global'.`,
				Default: "global",
			},
			"mtls_policy": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `This field is required if the policy is used with external HTTPS load balancers. This field can be empty for Traffic Director.
Defines a mechanism to provision peer validation certificates for peer to peer authentication (Mutual TLS - mTLS). If not specified, client certificate will not be requested. The connection is treated as TLS and not mTLS. If allowOpen and mtlsPolicy are set, server allows both plain text and mTLS connections.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_validation_ca": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `Required if the policy is to be used with Traffic Director. For external HTTPS load balancers it must be empty.
Defines the mechanism to obtain the Certificate Authority certificate to validate the client certificate.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"certificate_provider_instance": {
										Type:     schema.TypeList,
										Optional: true,
										Description: `Optional if policy is to be used with Traffic Director. For external HTTPS load balancer must be empty.
Defines a mechanism to provision server identity (public and private keys). Cannot be combined with allowOpen as a permissive mode that allows both plain text and TLS is not supported.`,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"plugin_instance": {
													Type:        schema.TypeString,
													Required:    true,
													Description: `Plugin instance name, used to locate and load CertificateProvider instance configuration. Set to "google_cloud_private_spiffe" to use Certificate Authority Service certificate provider instance.`,
												},
											},
										},
										ExactlyOneOf: []string{},
									},
									"grpc_endpoint": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: `gRPC specific configuration to access the gRPC server to obtain the cert and private key.`,
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"target_uri": {
													Type:        schema.TypeString,
													Required:    true,
													Description: `The target URI of the gRPC endpoint. Only UDS path is supported, and should start with "unix:".`,
												},
											},
										},
										ExactlyOneOf: []string{},
									},
								},
							},
						},
						"client_validation_mode": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: verify.ValidateEnum([]string{"CLIENT_VALIDATION_MODE_UNSPECIFIED", "ALLOW_INVALID_OR_MISSING_CLIENT_CERT", "REJECT_INVALID", ""}),
							Description: `When the client presents an invalid certificate or no certificate to the load balancer, the clientValidationMode specifies how the client connection is handled.
Required if the policy is to be used with the external HTTPS load balancing. For Traffic Director it must be empty. Possible values: ["CLIENT_VALIDATION_MODE_UNSPECIFIED", "ALLOW_INVALID_OR_MISSING_CLIENT_CERT", "REJECT_INVALID"]`,
						},
						"client_validation_trust_config": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `Reference to the TrustConfig from certificatemanager.googleapis.com namespace.
If specified, the chain validation will be performed against certificates configured in the given TrustConfig.
Allowed only if the policy is to be used with external HTTPS load balancers.`,
						},
					},
				},
			},
			"server_certificate": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Defines a mechanism to provision client identity (public and private keys) for peer to peer authentication. The presence of this dictates mTLS.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate_provider_instance": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `Optional if policy is to be used with Traffic Director. For external HTTPS load balancer must be empty.
Defines a mechanism to provision server identity (public and private keys). Cannot be combined with allowOpen as a permissive mode that allows both plain text and TLS is not supported.`,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"plugin_instance": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `Plugin instance name, used to locate and load CertificateProvider instance configuration. Set to "google_cloud_private_spiffe" to use Certificate Authority Service certificate provider instance.`,
									},
								},
							},
							ExactlyOneOf: []string{},
						},
						"grpc_endpoint": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `gRPC specific configuration to access the gRPC server to obtain the cert and private key.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"target_uri": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `The target URI of the gRPC endpoint. Only UDS path is supported, and should start with "unix:".`,
									},
								},
							},
							ExactlyOneOf: []string{},
						},
					},
				},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the ServerTlsPolicy was created in UTC.`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
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
				Description: `Time the ServerTlsPolicy was updated in UTC.`,
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

func resourceNetworkSecurityServerTlsPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkSecurityServerTlsPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	allowOpenProp, err := expandNetworkSecurityServerTlsPolicyAllowOpen(d.Get("allow_open"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("allow_open"); !tpgresource.IsEmptyValue(reflect.ValueOf(allowOpenProp)) && (ok || !reflect.DeepEqual(v, allowOpenProp)) {
		obj["allowOpen"] = allowOpenProp
	}
	serverCertificateProp, err := expandNetworkSecurityServerTlsPolicyServerCertificate(d.Get("server_certificate"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("server_certificate"); !tpgresource.IsEmptyValue(reflect.ValueOf(serverCertificateProp)) && (ok || !reflect.DeepEqual(v, serverCertificateProp)) {
		obj["serverCertificate"] = serverCertificateProp
	}
	mtlsPolicyProp, err := expandNetworkSecurityServerTlsPolicyMtlsPolicy(d.Get("mtls_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("mtls_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(mtlsPolicyProp)) && (ok || !reflect.DeepEqual(v, mtlsPolicyProp)) {
		obj["mtlsPolicy"] = mtlsPolicyProp
	}
	labelsProp, err := expandNetworkSecurityServerTlsPolicyEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/serverTlsPolicies?serverTlsPolicyId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ServerTlsPolicy: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ServerTlsPolicy: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating ServerTlsPolicy: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/serverTlsPolicies/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = NetworkSecurityOperationWaitTime(
		config, res, project, "Creating ServerTlsPolicy", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create ServerTlsPolicy: %s", err)
	}

	log.Printf("[DEBUG] Finished creating ServerTlsPolicy %q: %#v", d.Id(), res)

	return resourceNetworkSecurityServerTlsPolicyRead(d, meta)
}

func resourceNetworkSecurityServerTlsPolicyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/serverTlsPolicies/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ServerTlsPolicy: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("NetworkSecurityServerTlsPolicy %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ServerTlsPolicy: %s", err)
	}

	if err := d.Set("create_time", flattenNetworkSecurityServerTlsPolicyCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServerTlsPolicy: %s", err)
	}
	if err := d.Set("update_time", flattenNetworkSecurityServerTlsPolicyUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServerTlsPolicy: %s", err)
	}
	if err := d.Set("labels", flattenNetworkSecurityServerTlsPolicyLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServerTlsPolicy: %s", err)
	}
	if err := d.Set("description", flattenNetworkSecurityServerTlsPolicyDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServerTlsPolicy: %s", err)
	}
	if err := d.Set("allow_open", flattenNetworkSecurityServerTlsPolicyAllowOpen(res["allowOpen"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServerTlsPolicy: %s", err)
	}
	if err := d.Set("server_certificate", flattenNetworkSecurityServerTlsPolicyServerCertificate(res["serverCertificate"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServerTlsPolicy: %s", err)
	}
	if err := d.Set("mtls_policy", flattenNetworkSecurityServerTlsPolicyMtlsPolicy(res["mtlsPolicy"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServerTlsPolicy: %s", err)
	}
	if err := d.Set("terraform_labels", flattenNetworkSecurityServerTlsPolicyTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServerTlsPolicy: %s", err)
	}
	if err := d.Set("effective_labels", flattenNetworkSecurityServerTlsPolicyEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServerTlsPolicy: %s", err)
	}

	return nil
}

func resourceNetworkSecurityServerTlsPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ServerTlsPolicy: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkSecurityServerTlsPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	allowOpenProp, err := expandNetworkSecurityServerTlsPolicyAllowOpen(d.Get("allow_open"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("allow_open"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, allowOpenProp)) {
		obj["allowOpen"] = allowOpenProp
	}
	serverCertificateProp, err := expandNetworkSecurityServerTlsPolicyServerCertificate(d.Get("server_certificate"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("server_certificate"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, serverCertificateProp)) {
		obj["serverCertificate"] = serverCertificateProp
	}
	mtlsPolicyProp, err := expandNetworkSecurityServerTlsPolicyMtlsPolicy(d.Get("mtls_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("mtls_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, mtlsPolicyProp)) {
		obj["mtlsPolicy"] = mtlsPolicyProp
	}
	labelsProp, err := expandNetworkSecurityServerTlsPolicyEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/serverTlsPolicies/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ServerTlsPolicy %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("allow_open") {
		updateMask = append(updateMask, "allowOpen")
	}

	if d.HasChange("server_certificate") {
		updateMask = append(updateMask, "serverCertificate")
	}

	if d.HasChange("mtls_policy") {
		updateMask = append(updateMask, "mtlsPolicy")
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

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
	})

	if err != nil {
		return fmt.Errorf("Error updating ServerTlsPolicy %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating ServerTlsPolicy %q: %#v", d.Id(), res)
	}

	err = NetworkSecurityOperationWaitTime(
		config, res, project, "Updating ServerTlsPolicy", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceNetworkSecurityServerTlsPolicyRead(d, meta)
}

func resourceNetworkSecurityServerTlsPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ServerTlsPolicy: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/serverTlsPolicies/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting ServerTlsPolicy %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "ServerTlsPolicy")
	}

	err = NetworkSecurityOperationWaitTime(
		config, res, project, "Deleting ServerTlsPolicy", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting ServerTlsPolicy %q: %#v", d.Id(), res)
	return nil
}

func resourceNetworkSecurityServerTlsPolicyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/serverTlsPolicies/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/serverTlsPolicies/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNetworkSecurityServerTlsPolicyCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityServerTlsPolicyUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityServerTlsPolicyLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenNetworkSecurityServerTlsPolicyDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityServerTlsPolicyAllowOpen(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityServerTlsPolicyServerCertificate(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["grpc_endpoint"] =
		flattenNetworkSecurityServerTlsPolicyServerCertificateGrpcEndpoint(original["grpcEndpoint"], d, config)
	transformed["certificate_provider_instance"] =
		flattenNetworkSecurityServerTlsPolicyServerCertificateCertificateProviderInstance(original["certificateProviderInstance"], d, config)
	return []interface{}{transformed}
}
func flattenNetworkSecurityServerTlsPolicyServerCertificateGrpcEndpoint(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["target_uri"] =
		flattenNetworkSecurityServerTlsPolicyServerCertificateGrpcEndpointTargetUri(original["targetUri"], d, config)
	return []interface{}{transformed}
}
func flattenNetworkSecurityServerTlsPolicyServerCertificateGrpcEndpointTargetUri(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityServerTlsPolicyServerCertificateCertificateProviderInstance(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["plugin_instance"] =
		flattenNetworkSecurityServerTlsPolicyServerCertificateCertificateProviderInstancePluginInstance(original["pluginInstance"], d, config)
	return []interface{}{transformed}
}
func flattenNetworkSecurityServerTlsPolicyServerCertificateCertificateProviderInstancePluginInstance(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityServerTlsPolicyMtlsPolicy(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["client_validation_mode"] =
		flattenNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationMode(original["clientValidationMode"], d, config)
	transformed["client_validation_trust_config"] =
		flattenNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationTrustConfig(original["clientValidationTrustConfig"], d, config)
	transformed["client_validation_ca"] =
		flattenNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCa(original["clientValidationCa"], d, config)
	return []interface{}{transformed}
}
func flattenNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationMode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationTrustConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCa(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"grpc_endpoint":                 flattenNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint(original["grpcEndpoint"], d, config),
			"certificate_provider_instance": flattenNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance(original["certificateProviderInstance"], d, config),
		})
	}
	return transformed
}
func flattenNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["target_uri"] =
		flattenNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpointTargetUri(original["targetUri"], d, config)
	return []interface{}{transformed}
}
func flattenNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpointTargetUri(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["plugin_instance"] =
		flattenNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstancePluginInstance(original["pluginInstance"], d, config)
	return []interface{}{transformed}
}
func flattenNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstancePluginInstance(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityServerTlsPolicyTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenNetworkSecurityServerTlsPolicyEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNetworkSecurityServerTlsPolicyDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityServerTlsPolicyAllowOpen(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityServerTlsPolicyServerCertificate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedGrpcEndpoint, err := expandNetworkSecurityServerTlsPolicyServerCertificateGrpcEndpoint(original["grpc_endpoint"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGrpcEndpoint); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["grpcEndpoint"] = transformedGrpcEndpoint
	}

	transformedCertificateProviderInstance, err := expandNetworkSecurityServerTlsPolicyServerCertificateCertificateProviderInstance(original["certificate_provider_instance"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCertificateProviderInstance); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["certificateProviderInstance"] = transformedCertificateProviderInstance
	}

	return transformed, nil
}

func expandNetworkSecurityServerTlsPolicyServerCertificateGrpcEndpoint(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTargetUri, err := expandNetworkSecurityServerTlsPolicyServerCertificateGrpcEndpointTargetUri(original["target_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTargetUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["targetUri"] = transformedTargetUri
	}

	return transformed, nil
}

func expandNetworkSecurityServerTlsPolicyServerCertificateGrpcEndpointTargetUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityServerTlsPolicyServerCertificateCertificateProviderInstance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPluginInstance, err := expandNetworkSecurityServerTlsPolicyServerCertificateCertificateProviderInstancePluginInstance(original["plugin_instance"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPluginInstance); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["pluginInstance"] = transformedPluginInstance
	}

	return transformed, nil
}

func expandNetworkSecurityServerTlsPolicyServerCertificateCertificateProviderInstancePluginInstance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityServerTlsPolicyMtlsPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedClientValidationMode, err := expandNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationMode(original["client_validation_mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedClientValidationMode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["clientValidationMode"] = transformedClientValidationMode
	}

	transformedClientValidationTrustConfig, err := expandNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationTrustConfig(original["client_validation_trust_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedClientValidationTrustConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["clientValidationTrustConfig"] = transformedClientValidationTrustConfig
	}

	transformedClientValidationCa, err := expandNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCa(original["client_validation_ca"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedClientValidationCa); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["clientValidationCa"] = transformedClientValidationCa
	}

	return transformed, nil
}

func expandNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationTrustConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCa(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedGrpcEndpoint, err := expandNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint(original["grpc_endpoint"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedGrpcEndpoint); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["grpcEndpoint"] = transformedGrpcEndpoint
		}

		transformedCertificateProviderInstance, err := expandNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance(original["certificate_provider_instance"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCertificateProviderInstance); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["certificateProviderInstance"] = transformedCertificateProviderInstance
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTargetUri, err := expandNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpointTargetUri(original["target_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTargetUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["targetUri"] = transformedTargetUri
	}

	return transformed, nil
}

func expandNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpointTargetUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPluginInstance, err := expandNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstancePluginInstance(original["plugin_instance"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPluginInstance); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["pluginInstance"] = transformedPluginInstance
	}

	return transformed, nil
}

func expandNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstancePluginInstance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityServerTlsPolicyEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
