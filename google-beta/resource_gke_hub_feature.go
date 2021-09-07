// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: DCL     ***
//
// ----------------------------------------------------------------------------
//
//     This file is managed by Magic Modules (https://github.com/GoogleCloudPlatform/magic-modules)
//     and is based on the DCL (https://github.com/GoogleCloudPlatform/declarative-resource-client-library).
//     Changes will need to be made to the DCL or Magic Modules instead of here.
//
//     We are not currently able to accept contributions to this file. If changes
//     are required, please file an issue at https://github.com/hashicorp/terraform-provider-google/issues/new/choose
//
// ----------------------------------------------------------------------------

package google

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	dcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	gkehub "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkehub/beta"
)

func resourceGkeHubFeature() *schema.Resource {
	return &schema.Resource{
		Create: resourceGkeHubFeatureCreate,
		Read:   resourceGkeHubFeatureRead,
		Update: resourceGkeHubFeatureUpdate,
		Delete: resourceGkeHubFeatureDelete,

		Importer: &schema.ResourceImporter{
			State: resourceGkeHubFeatureImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The location for the resource",
			},

			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "GCP labels for this Feature.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The full, unique name of this Feature resource",
			},

			"project": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      "The project for the resource",
			},

			"spec": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.",
				MaxItems:    1,
				Elem:        GkeHubFeatureSpecSchema(),
			},

			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. When the Feature resource was created.",
			},

			"delete_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. When the Feature resource was deleted.",
			},

			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. When the Feature resource was last updated.",
			},
		},
	}
}

func GkeHubFeatureSpecSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"multiclusteringress": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Multicluster Ingress-specific spec.",
				MaxItems:    1,
				Elem:        GkeHubFeatureSpecMulticlusteringressSchema(),
			},
		},
	}
}

func GkeHubFeatureSpecMulticlusteringressSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"config_membership": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      "Fully-qualified Membership name which hosts the MultiClusterIngress CRD. Example: `projects/foo-proj/locations/global/memberships/bar`",
			},
		},
	}
}

func resourceGkeHubFeatureCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &gkehub.Feature{
		Location: dcl.String(d.Get("location").(string)),
		Labels:   checkStringMap(d.Get("labels")),
		Name:     dcl.String(d.Get("name").(string)),
		Project:  dcl.String(project),
		Spec:     expandGkeHubFeatureSpec(d.Get("spec")),
	}
	lockName, err := replaceVarsForId(d, config, "{{project}}/{{location}}/{{feature}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	id, err := replaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/features/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)
	createDirective := CreateDirective
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLGkeHubClient(config, userAgent, billingProject)
	res, err := client.ApplyFeature(context.Background(), obj, createDirective...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error creating Feature: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Feature %q: %#v", d.Id(), res)

	return resourceGkeHubFeatureRead(d, meta)
}

func resourceGkeHubFeatureRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &gkehub.Feature{
		Location: dcl.String(d.Get("location").(string)),
		Labels:   checkStringMap(d.Get("labels")),
		Name:     dcl.String(d.Get("name").(string)),
		Project:  dcl.String(project),
		Spec:     expandGkeHubFeatureSpec(d.Get("spec")),
	}

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLGkeHubClient(config, userAgent, billingProject)
	res, err := client.GetFeature(context.Background(), obj)
	if err != nil {
		resourceName := fmt.Sprintf("GkeHubFeature %q", d.Id())
		return handleNotFoundDCLError(err, d, resourceName)
	}

	if err = d.Set("location", res.Location); err != nil {
		return fmt.Errorf("error setting location in state: %s", err)
	}
	if err = d.Set("labels", res.Labels); err != nil {
		return fmt.Errorf("error setting labels in state: %s", err)
	}
	if err = d.Set("name", res.Name); err != nil {
		return fmt.Errorf("error setting name in state: %s", err)
	}
	if err = d.Set("project", res.Project); err != nil {
		return fmt.Errorf("error setting project in state: %s", err)
	}
	if err = d.Set("spec", flattenGkeHubFeatureSpec(res.Spec)); err != nil {
		return fmt.Errorf("error setting spec in state: %s", err)
	}
	if err = d.Set("create_time", res.CreateTime); err != nil {
		return fmt.Errorf("error setting create_time in state: %s", err)
	}
	if err = d.Set("delete_time", res.DeleteTime); err != nil {
		return fmt.Errorf("error setting delete_time in state: %s", err)
	}
	if err = d.Set("update_time", res.UpdateTime); err != nil {
		return fmt.Errorf("error setting update_time in state: %s", err)
	}

	return nil
}
func resourceGkeHubFeatureUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &gkehub.Feature{
		Location: dcl.String(d.Get("location").(string)),
		Labels:   checkStringMap(d.Get("labels")),
		Name:     dcl.String(d.Get("name").(string)),
		Project:  dcl.String(project),
		Spec:     expandGkeHubFeatureSpec(d.Get("spec")),
	}
	lockName, err := replaceVarsForId(d, config, "{{project}}/{{location}}/{{feature}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	directive := UpdateDirective
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLGkeHubClient(config, userAgent, billingProject)
	res, err := client.ApplyFeature(context.Background(), obj, directive...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error updating Feature: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Feature %q: %#v", d.Id(), res)

	return resourceGkeHubFeatureRead(d, meta)
}

func resourceGkeHubFeatureDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &gkehub.Feature{
		Location: dcl.String(d.Get("location").(string)),
		Labels:   checkStringMap(d.Get("labels")),
		Name:     dcl.String(d.Get("name").(string)),
		Project:  dcl.String(project),
		Spec:     expandGkeHubFeatureSpec(d.Get("spec")),
	}
	lockName, err := replaceVarsForId(d, config, "{{project}}/{{location}}/{{feature}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	log.Printf("[DEBUG] Deleting Feature %q", d.Id())
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLGkeHubClient(config, userAgent, billingProject)
	if err := client.DeleteFeature(context.Background(), obj); err != nil {
		return fmt.Errorf("Error deleting Feature: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting Feature %q", d.Id())
	return nil
}

func resourceGkeHubFeatureImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/features/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/features/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func expandGkeHubFeatureSpec(o interface{}) *gkehub.FeatureSpec {
	if o == nil {
		return gkehub.EmptyFeatureSpec
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return gkehub.EmptyFeatureSpec
	}
	obj := objArr[0].(map[string]interface{})
	return &gkehub.FeatureSpec{
		Multiclusteringress: expandGkeHubFeatureSpecMulticlusteringress(obj["multiclusteringress"]),
	}
}

func flattenGkeHubFeatureSpec(obj *gkehub.FeatureSpec) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"multiclusteringress": flattenGkeHubFeatureSpecMulticlusteringress(obj.Multiclusteringress),
	}

	return []interface{}{transformed}

}

func expandGkeHubFeatureSpecMulticlusteringress(o interface{}) *gkehub.FeatureSpecMulticlusteringress {
	if o == nil {
		return gkehub.EmptyFeatureSpecMulticlusteringress
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return gkehub.EmptyFeatureSpecMulticlusteringress
	}
	obj := objArr[0].(map[string]interface{})
	return &gkehub.FeatureSpecMulticlusteringress{
		ConfigMembership: dcl.String(obj["config_membership"].(string)),
	}
}

func flattenGkeHubFeatureSpecMulticlusteringress(obj *gkehub.FeatureSpecMulticlusteringress) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"config_membership": obj.ConfigMembership,
	}

	return []interface{}{transformed}

}
