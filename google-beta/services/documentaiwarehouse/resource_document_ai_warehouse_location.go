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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/documentaiwarehouse/Location.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package documentaiwarehouse

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceDocumentAIWarehouseLocation() *schema.Resource {
	return &schema.Resource{
		Create: resourceDocumentAIWarehouseLocationCreate,
		Read:   resourceDocumentAIWarehouseLocationRead,
		Delete: resourceDocumentAIWarehouseLocationDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"access_control_mode": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"ACL_MODE_DOCUMENT_LEVEL_ACCESS_CONTROL_GCI", "ACL_MODE_DOCUMENT_LEVEL_ACCESS_CONTROL_BYOID", "ACL_MODE_UNIVERSAL_ACCESS"}),
				Description:  `The access control mode for accessing the customer data. Possible values: ["ACL_MODE_DOCUMENT_LEVEL_ACCESS_CONTROL_GCI", "ACL_MODE_DOCUMENT_LEVEL_ACCESS_CONTROL_BYOID", "ACL_MODE_UNIVERSAL_ACCESS"]`,
			},
			"database_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"DB_INFRA_SPANNER", "DB_CLOUD_SQL_POSTGRES"}),
				Description:  `The type of database used to store customer data. Possible values: ["DB_INFRA_SPANNER", "DB_CLOUD_SQL_POSTGRES"]`,
			},
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The location in which the instance is to be provisioned. It takes the form projects/{projectNumber}/locations/{location}.`,
			},
			"project_number": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The unique identifier of the project.`,
			},
			"document_creator_default_role": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"DOCUMENT_ADMIN", "DOCUMENT_EDITOR", "DOCUMENT_VIEWER", ""}),
				Description:  `The default role for the person who create a document. Possible values: ["DOCUMENT_ADMIN", "DOCUMENT_EDITOR", "DOCUMENT_VIEWER"]`,
			},
			"kms_key": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The KMS key used for CMEK encryption. It is required that
the kms key is in the same region as the endpoint. The
same key will be used for all provisioned resources, if
encryption is available. If the kmsKey is left empty, no
encryption will be enforced.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceDocumentAIWarehouseLocationCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	databaseTypeProp, err := expandDocumentAIWarehouseLocationDatabaseType(d.Get("database_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("database_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(databaseTypeProp)) && (ok || !reflect.DeepEqual(v, databaseTypeProp)) {
		obj["databaseType"] = databaseTypeProp
	}
	accessControlModeProp, err := expandDocumentAIWarehouseLocationAccessControlMode(d.Get("access_control_mode"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("access_control_mode"); !tpgresource.IsEmptyValue(reflect.ValueOf(accessControlModeProp)) && (ok || !reflect.DeepEqual(v, accessControlModeProp)) {
		obj["accessControlMode"] = accessControlModeProp
	}
	kmsKeyProp, err := expandDocumentAIWarehouseLocationKmsKey(d.Get("kms_key"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("kms_key"); !tpgresource.IsEmptyValue(reflect.ValueOf(kmsKeyProp)) && (ok || !reflect.DeepEqual(v, kmsKeyProp)) {
		obj["kmsKey"] = kmsKeyProp
	}
	documentCreatorDefaultRoleProp, err := expandDocumentAIWarehouseLocationDocumentCreatorDefaultRole(d.Get("document_creator_default_role"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("document_creator_default_role"); !tpgresource.IsEmptyValue(reflect.ValueOf(documentCreatorDefaultRoleProp)) && (ok || !reflect.DeepEqual(v, documentCreatorDefaultRoleProp)) {
		obj["documentCreatorDefaultRole"] = documentCreatorDefaultRoleProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DocumentAIWarehouseBasePath}}projects/{{project_number}}/locations/{{location}}:initialize")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Location: %#v", obj)
	billingProject := ""

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
		return fmt.Errorf("Error creating Location: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project_number}}/locations/{{location}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = DocumentAIWarehouseOperationWaitTime(
		config, res, "Creating Location", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Location: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Location %q: %#v", d.Id(), res)

	return resourceDocumentAIWarehouseLocationRead(d, meta)
}

func resourceDocumentAIWarehouseLocationRead(d *schema.ResourceData, meta interface{}) error {
	// This resource could not be read from the API.
	return nil
}

func resourceDocumentAIWarehouseLocationDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARNING] DocumentAIWarehouse Location resources"+
		" cannot be deleted from Google Cloud. The resource %s will be removed from Terraform"+
		" state, but will still be present on Google Cloud.", d.Id())
	d.SetId("")

	return nil
}

func flattenDocumentAIWarehouseLocationDatabaseType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDocumentAIWarehouseLocationAccessControlMode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDocumentAIWarehouseLocationKmsKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDocumentAIWarehouseLocationDocumentCreatorDefaultRole(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandDocumentAIWarehouseLocationDatabaseType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDocumentAIWarehouseLocationAccessControlMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDocumentAIWarehouseLocationKmsKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDocumentAIWarehouseLocationDocumentCreatorDefaultRole(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
