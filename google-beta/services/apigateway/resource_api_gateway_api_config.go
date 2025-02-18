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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/apigateway/ApiConfig.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package apigateway

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceApiGatewayApiConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceApiGatewayApiConfigCreate,
		Read:   resourceApiGatewayApiConfigRead,
		Update: resourceApiGatewayApiConfigUpdate,
		Delete: resourceApiGatewayApiConfigDelete,

		Importer: &schema.ResourceImporter{
			State: resourceApiGatewayApiConfigImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"api": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The API to attach the config to.`,
			},
			"api_config_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `Identifier to assign to the API Config. Must be unique within scope of the parent resource(api).`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: `A user-visible name for the API.`,
			},
			"gateway_config": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Description: `Immutable. Gateway specific configuration.
If not specified, backend authentication will be set to use OIDC authentication using the default compute service account`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backend_config": {
							Type:        schema.TypeList,
							Required:    true,
							Description: `Backend settings that are applied to all backends of the Gateway.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"google_service_account": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
										Description: `Google Cloud IAM service account used to sign OIDC tokens for backends that have authentication configured
(https://cloud.google.com/service-infrastructure/docs/service-management/reference/rest/v1/services.configs#backend).`,
									},
								},
							},
						},
					},
				},
			},
			"grpc_services": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `gRPC service definition files. If specified, openapiDocuments must not be included.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"file_descriptor_set": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							Description: `Input only. File descriptor set, generated by protoc.
To generate, use protoc with imports and source info included. For an example test.proto file, the following command would put the value in a new file named out.pb.

$ protoc --include_imports --include_source_info test.proto -o out.pb`,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"contents": {
										Type:         schema.TypeString,
										Required:     true,
										ForceNew:     true,
										ValidateFunc: verify.ValidateBase64String,
										Description:  `Base64 encoded content of the file.`,
									},
									"path": {
										Type:        schema.TypeString,
										Required:    true,
										ForceNew:    true,
										Description: `The file path (full or relative path). This is typically the path of the file when it is uploaded.`,
									},
								},
							},
						},
						"source": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Uncompiled proto files associated with the descriptor set, used for display purposes (server-side compilation is not supported). These should match the inputs to 'protoc' command used to generate fileDescriptorSet.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"contents": {
										Type:         schema.TypeString,
										Required:     true,
										ForceNew:     true,
										ValidateFunc: verify.ValidateBase64String,
										Description:  `Base64 encoded content of the file.`,
									},
									"path": {
										Type:        schema.TypeString,
										Required:    true,
										ForceNew:    true,
										Description: `The file path (full or relative path). This is typically the path of the file when it is uploaded.`,
									},
								},
							},
						},
					},
				},
				ExactlyOneOf: []string{"openapi_documents", "grpc_services"},
				RequiredWith: []string{"managed_service_configs"},
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Resource labels to represent user-provided metadata.


**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"managed_service_configs": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `Optional. Service Configuration files. At least one must be included when using gRPC service definitions. See https://cloud.google.com/endpoints/docs/grpc/grpc-service-config#service_configuration_overview for the expected file contents.
If multiple files are specified, the files are merged with the following rules: * All singular scalar fields are merged using "last one wins" semantics in the order of the files uploaded. * Repeated fields are concatenated. * Singular embedded messages are merged using these rules for nested fields.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"contents": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: `Base64 encoded content of the file.`,
						},
						"path": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: `The file path (full or relative path). This is typically the path of the file when it is uploaded.`,
						},
					},
				},
				RequiredWith: []string{"grpc_services"},
			},
			"openapi_documents": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `OpenAPI specification documents. If specified, grpcServices and managedServiceConfigs must not be included.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"document": {
							Type:        schema.TypeList,
							Required:    true,
							Description: `The OpenAPI Specification document file.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"contents": {
										Type:         schema.TypeString,
										Required:     true,
										ForceNew:     true,
										ValidateFunc: verify.ValidateBase64String,
										Description:  `Base64 encoded content of the file.`,
									},
									"path": {
										Type:        schema.TypeString,
										Required:    true,
										ForceNew:    true,
										Description: `The file path (full or relative path). This is typically the path of the file when it is uploaded.`,
									},
								},
							},
						},
					},
				},
				ExactlyOneOf: []string{"openapi_documents", "grpc_services"},
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
				Description: `The resource name of the API Config.`,
			},
			"service_config_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The ID of the associated Service Config (https://cloud.google.com/service-infrastructure/docs/glossary#config).`,
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"api_config_id_prefix": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"api_config_id"},
				Description:   `Creates a unique name beginning with the specified prefix. If this and api_config_id are unspecified, a random value is chosen for the name.`,
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

func resourceApiGatewayApiConfigCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandApiGatewayApiConfigDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	gatewayConfigProp, err := expandApiGatewayApiConfigGatewayConfig(d.Get("gateway_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("gateway_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(gatewayConfigProp)) && (ok || !reflect.DeepEqual(v, gatewayConfigProp)) {
		obj["gatewayConfig"] = gatewayConfigProp
	}
	openapiDocumentsProp, err := expandApiGatewayApiConfigOpenapiDocuments(d.Get("openapi_documents"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("openapi_documents"); !tpgresource.IsEmptyValue(reflect.ValueOf(openapiDocumentsProp)) && (ok || !reflect.DeepEqual(v, openapiDocumentsProp)) {
		obj["openapiDocuments"] = openapiDocumentsProp
	}
	grpcServicesProp, err := expandApiGatewayApiConfigGrpcServices(d.Get("grpc_services"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("grpc_services"); !tpgresource.IsEmptyValue(reflect.ValueOf(grpcServicesProp)) && (ok || !reflect.DeepEqual(v, grpcServicesProp)) {
		obj["grpcServices"] = grpcServicesProp
	}
	managedServiceConfigsProp, err := expandApiGatewayApiConfigManagedServiceConfigs(d.Get("managed_service_configs"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("managed_service_configs"); !tpgresource.IsEmptyValue(reflect.ValueOf(managedServiceConfigsProp)) && (ok || !reflect.DeepEqual(v, managedServiceConfigsProp)) {
		obj["managedServiceConfigs"] = managedServiceConfigsProp
	}
	labelsProp, err := expandApiGatewayApiConfigEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	obj, err = resourceApiGatewayApiConfigEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ApiGatewayBasePath}}projects/{{project}}/locations/global/apis/{{api}}/configs?apiConfigId={{api_config_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ApiConfig: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ApiConfig: %s", err)
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
		return fmt.Errorf("Error creating ApiConfig: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = ApiGatewayOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating ApiConfig", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create ApiConfig: %s", err)
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating ApiConfig %q: %#v", d.Id(), res)

	return resourceApiGatewayApiConfigRead(d, meta)
}

func resourceApiGatewayApiConfigRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ApiGatewayBasePath}}projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config_id}}?view=FULL")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ApiConfig: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ApiGatewayApiConfig %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ApiConfig: %s", err)
	}

	if err := d.Set("name", flattenApiGatewayApiConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading ApiConfig: %s", err)
	}
	if err := d.Set("display_name", flattenApiGatewayApiConfigDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading ApiConfig: %s", err)
	}
	if err := d.Set("service_config_id", flattenApiGatewayApiConfigServiceConfigId(res["serviceConfigId"], d, config)); err != nil {
		return fmt.Errorf("Error reading ApiConfig: %s", err)
	}
	if err := d.Set("labels", flattenApiGatewayApiConfigLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading ApiConfig: %s", err)
	}
	if err := d.Set("openapi_documents", flattenApiGatewayApiConfigOpenapiDocuments(res["openapiDocuments"], d, config)); err != nil {
		return fmt.Errorf("Error reading ApiConfig: %s", err)
	}
	if err := d.Set("managed_service_configs", flattenApiGatewayApiConfigManagedServiceConfigs(res["managedServiceConfigs"], d, config)); err != nil {
		return fmt.Errorf("Error reading ApiConfig: %s", err)
	}
	if err := d.Set("terraform_labels", flattenApiGatewayApiConfigTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading ApiConfig: %s", err)
	}
	if err := d.Set("effective_labels", flattenApiGatewayApiConfigEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading ApiConfig: %s", err)
	}

	return nil
}

func resourceApiGatewayApiConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ApiConfig: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandApiGatewayApiConfigDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	openapiDocumentsProp, err := expandApiGatewayApiConfigOpenapiDocuments(d.Get("openapi_documents"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("openapi_documents"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, openapiDocumentsProp)) {
		obj["openapiDocuments"] = openapiDocumentsProp
	}
	grpcServicesProp, err := expandApiGatewayApiConfigGrpcServices(d.Get("grpc_services"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("grpc_services"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, grpcServicesProp)) {
		obj["grpcServices"] = grpcServicesProp
	}
	managedServiceConfigsProp, err := expandApiGatewayApiConfigManagedServiceConfigs(d.Get("managed_service_configs"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("managed_service_configs"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, managedServiceConfigsProp)) {
		obj["managedServiceConfigs"] = managedServiceConfigsProp
	}
	labelsProp, err := expandApiGatewayApiConfigEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	obj, err = resourceApiGatewayApiConfigEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ApiGatewayBasePath}}projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ApiConfig %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("openapi_documents") {
		updateMask = append(updateMask, "openapiDocuments")
	}

	if d.HasChange("grpc_services") {
		updateMask = append(updateMask, "grpcServices")
	}

	if d.HasChange("managed_service_configs") {
		updateMask = append(updateMask, "managedServiceConfigs")
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
			return fmt.Errorf("Error updating ApiConfig %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating ApiConfig %q: %#v", d.Id(), res)
		}

		err = ApiGatewayOperationWaitTime(
			config, res, project, "Updating ApiConfig", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceApiGatewayApiConfigRead(d, meta)
}

func resourceApiGatewayApiConfigDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ApiConfig: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{ApiGatewayBasePath}}projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting ApiConfig %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "ApiConfig")
	}

	err = ApiGatewayOperationWaitTime(
		config, res, project, "Deleting ApiConfig", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting ApiConfig %q: %#v", d.Id(), res)
	return nil
}

func resourceApiGatewayApiConfigImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/global/apis/(?P<api>[^/]+)/configs/(?P<api_config_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<api>[^/]+)/(?P<api_config_id>[^/]+)$",
		"^(?P<api>[^/]+)/(?P<api_config_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenApiGatewayApiConfigName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApiGatewayApiConfigDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApiGatewayApiConfigServiceConfigId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApiGatewayApiConfigLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenApiGatewayApiConfigOpenapiDocuments(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"document": flattenApiGatewayApiConfigOpenapiDocumentsDocument(original["document"], d, config),
		})
	}
	return transformed
}
func flattenApiGatewayApiConfigOpenapiDocumentsDocument(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["path"] =
		flattenApiGatewayApiConfigOpenapiDocumentsDocumentPath(original["path"], d, config)
	transformed["contents"] =
		flattenApiGatewayApiConfigOpenapiDocumentsDocumentContents(original["contents"], d, config)
	return []interface{}{transformed}
}
func flattenApiGatewayApiConfigOpenapiDocumentsDocumentPath(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApiGatewayApiConfigOpenapiDocumentsDocumentContents(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApiGatewayApiConfigManagedServiceConfigs(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"path":     flattenApiGatewayApiConfigManagedServiceConfigsPath(original["path"], d, config),
			"contents": flattenApiGatewayApiConfigManagedServiceConfigsContents(original["contents"], d, config),
		})
	}
	return transformed
}
func flattenApiGatewayApiConfigManagedServiceConfigsPath(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApiGatewayApiConfigManagedServiceConfigsContents(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApiGatewayApiConfigTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenApiGatewayApiConfigEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandApiGatewayApiConfigDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApiGatewayApiConfigGatewayConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedBackendConfig, err := expandApiGatewayApiConfigGatewayConfigBackendConfig(original["backend_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBackendConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["backendConfig"] = transformedBackendConfig
	}

	return transformed, nil
}

func expandApiGatewayApiConfigGatewayConfigBackendConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedGoogleServiceAccount, err := expandApiGatewayApiConfigGatewayConfigBackendConfigGoogleServiceAccount(original["google_service_account"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGoogleServiceAccount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["googleServiceAccount"] = transformedGoogleServiceAccount
	}

	return transformed, nil
}

func expandApiGatewayApiConfigGatewayConfigBackendConfigGoogleServiceAccount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApiGatewayApiConfigOpenapiDocuments(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDocument, err := expandApiGatewayApiConfigOpenapiDocumentsDocument(original["document"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDocument); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["document"] = transformedDocument
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandApiGatewayApiConfigOpenapiDocumentsDocument(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPath, err := expandApiGatewayApiConfigOpenapiDocumentsDocumentPath(original["path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPath); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["path"] = transformedPath
	}

	transformedContents, err := expandApiGatewayApiConfigOpenapiDocumentsDocumentContents(original["contents"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedContents); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["contents"] = transformedContents
	}

	return transformed, nil
}

func expandApiGatewayApiConfigOpenapiDocumentsDocumentPath(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApiGatewayApiConfigOpenapiDocumentsDocumentContents(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApiGatewayApiConfigGrpcServices(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedFileDescriptorSet, err := expandApiGatewayApiConfigGrpcServicesFileDescriptorSet(original["file_descriptor_set"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedFileDescriptorSet); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["fileDescriptorSet"] = transformedFileDescriptorSet
		}

		transformedSource, err := expandApiGatewayApiConfigGrpcServicesSource(original["source"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSource); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["source"] = transformedSource
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandApiGatewayApiConfigGrpcServicesFileDescriptorSet(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPath, err := expandApiGatewayApiConfigGrpcServicesFileDescriptorSetPath(original["path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPath); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["path"] = transformedPath
	}

	transformedContents, err := expandApiGatewayApiConfigGrpcServicesFileDescriptorSetContents(original["contents"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedContents); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["contents"] = transformedContents
	}

	return transformed, nil
}

func expandApiGatewayApiConfigGrpcServicesFileDescriptorSetPath(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApiGatewayApiConfigGrpcServicesFileDescriptorSetContents(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApiGatewayApiConfigGrpcServicesSource(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedPath, err := expandApiGatewayApiConfigGrpcServicesSourcePath(original["path"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPath); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["path"] = transformedPath
		}

		transformedContents, err := expandApiGatewayApiConfigGrpcServicesSourceContents(original["contents"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedContents); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["contents"] = transformedContents
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandApiGatewayApiConfigGrpcServicesSourcePath(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApiGatewayApiConfigGrpcServicesSourceContents(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApiGatewayApiConfigManagedServiceConfigs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedPath, err := expandApiGatewayApiConfigManagedServiceConfigsPath(original["path"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPath); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["path"] = transformedPath
		}

		transformedContents, err := expandApiGatewayApiConfigManagedServiceConfigsContents(original["contents"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedContents); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["contents"] = transformedContents
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandApiGatewayApiConfigManagedServiceConfigsPath(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApiGatewayApiConfigManagedServiceConfigsContents(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApiGatewayApiConfigEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func resourceApiGatewayApiConfigEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	var apiConfigId string
	if v, ok := d.GetOk("api_config_id"); ok {
		apiConfigId = v.(string)
	} else if v, ok := d.GetOk("api_config_id_prefix"); ok {
		apiConfigId = id.PrefixedUniqueId(v.(string))
	} else {
		apiConfigId = id.UniqueId()
	}

	if err := d.Set("api_config_id", apiConfigId); err != nil {
		return nil, fmt.Errorf("Error setting api_config_id: %s", err)
	}
	return obj, nil
}
