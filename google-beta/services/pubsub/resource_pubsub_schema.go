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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/pubsub/Schema.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package pubsub

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourcePubsubSchema() *schema.Resource {
	return &schema.Resource{
		Create: resourcePubsubSchemaCreate,
		Read:   resourcePubsubSchemaRead,
		Update: resourcePubsubSchemaUpdate,
		Delete: resourcePubsubSchemaDelete,

		Importer: &schema.ResourceImporter{
			State: resourcePubsubSchemaImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The ID to use for the schema, which will become the final component of the schema's resource name.`,
			},
			"definition": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `The definition of the schema.
This should contain a string representing the full definition of the schema
that is a valid schema definition of the type specified in type. Changes
to the definition commit new [schema revisions](https://cloud.google.com/pubsub/docs/commit-schema-revision).
A schema can only have up to 20 revisions, so updates that fail with an
error indicating that the limit has been reached require manually
[deleting old revisions](https://cloud.google.com/pubsub/docs/delete-schema-revision).`,
			},
			"type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: verify.ValidateEnum([]string{"TYPE_UNSPECIFIED", "PROTOCOL_BUFFER", "AVRO", ""}),
				Description:  `The type of the schema definition Default value: "TYPE_UNSPECIFIED" Possible values: ["TYPE_UNSPECIFIED", "PROTOCOL_BUFFER", "AVRO"]`,
				Default:      "TYPE_UNSPECIFIED",
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

func resourcePubsubSchemaCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	typeProp, err := expandPubsubSchemaType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	definitionProp, err := expandPubsubSchemaDefinition(d.Get("definition"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("definition"); !tpgresource.IsEmptyValue(reflect.ValueOf(definitionProp)) && (ok || !reflect.DeepEqual(v, definitionProp)) {
		obj["definition"] = definitionProp
	}
	nameProp, err := expandPubsubSchemaName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{PubsubBasePath}}projects/{{project}}/schemas?schemaId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Schema: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Schema: %s", err)
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
		return fmt.Errorf("Error creating Schema: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/schemas/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Schema %q: %#v", d.Id(), res)

	return resourcePubsubSchemaRead(d, meta)
}

func resourcePubsubSchemaPollRead(d *schema.ResourceData, meta interface{}) transport_tpg.PollReadFunc {
	return func() (map[string]interface{}, error) {
		config := meta.(*transport_tpg.Config)

		url, err := tpgresource.ReplaceVars(d, config, "{{PubsubBasePath}}projects/{{project}}/schemas/{{name}}")

		if err != nil {
			return nil, err
		}

		billingProject := ""

		project, err := tpgresource.GetProject(d, config)
		if err != nil {
			return nil, fmt.Errorf("Error fetching project for Schema: %s", err)
		}
		billingProject = project

		// err == nil indicates that the billing_project value was found
		if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
			billingProject = bp
		}

		userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
		if err != nil {
			return nil, err
		}

		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "GET",
			Project:   billingProject,
			RawURL:    url,
			UserAgent: userAgent,
		})
		if err != nil {
			return res, err
		}
		return res, nil
	}
}

func resourcePubsubSchemaRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{PubsubBasePath}}projects/{{project}}/schemas/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Schema: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("PubsubSchema %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Schema: %s", err)
	}

	if err := d.Set("type", flattenPubsubSchemaType(res["type"], d, config)); err != nil {
		return fmt.Errorf("Error reading Schema: %s", err)
	}
	if err := d.Set("definition", flattenPubsubSchemaDefinition(res["definition"], d, config)); err != nil {
		return fmt.Errorf("Error reading Schema: %s", err)
	}
	if err := d.Set("name", flattenPubsubSchemaName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Schema: %s", err)
	}

	return nil
}

func resourcePubsubSchemaUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Schema: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	typeProp, err := expandPubsubSchemaType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	definitionProp, err := expandPubsubSchemaDefinition(d.Get("definition"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("definition"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, definitionProp)) {
		obj["definition"] = definitionProp
	}
	nameProp, err := expandPubsubSchemaName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	obj, err = resourcePubsubSchemaUpdateEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{PubsubBasePath}}projects/{{project}}/schemas/{{name}}:commit")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Schema %q: %#v", d.Id(), obj)
	headers := make(http.Header)

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
		Timeout:   d.Timeout(schema.TimeoutUpdate),
		Headers:   headers,
	})

	if err != nil {
		return fmt.Errorf("Error updating Schema %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Schema %q: %#v", d.Id(), res)
	}

	return resourcePubsubSchemaRead(d, meta)
}

func resourcePubsubSchemaDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Schema: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{PubsubBasePath}}projects/{{project}}/schemas/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting Schema %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "Schema")
	}

	err = transport_tpg.PollingWaitTime(resourcePubsubSchemaPollRead(d, meta), transport_tpg.PollCheckForAbsence, "Deleting Schema", d.Timeout(schema.TimeoutCreate), 10)
	if err != nil {
		return fmt.Errorf("Error waiting to delete Schema: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting Schema %q: %#v", d.Id(), res)
	return nil
}

func resourcePubsubSchemaImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/schemas/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/schemas/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenPubsubSchemaType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenPubsubSchemaDefinition(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenPubsubSchemaName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.GetResourceNameFromSelfLink(v.(string))
}

func expandPubsubSchemaType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSchemaDefinition(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSchemaName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return tpgresource.GetResourceNameFromSelfLink(v.(string)), nil
}

func resourcePubsubSchemaUpdateEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	newObj := make(map[string]interface{})
	newObj["name"] = d.Id()
	obj["name"] = d.Id()
	newObj["schema"] = obj
	return newObj, nil
}
