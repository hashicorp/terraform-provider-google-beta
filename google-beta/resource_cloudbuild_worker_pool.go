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
	cloudbuild "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbuild/beta"
)

func resourceCloudbuildWorkerPool() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudbuildWorkerPoolCreate,
		Read:   resourceCloudbuildWorkerPoolRead,
		Update: resourceCloudbuildWorkerPoolUpdate,
		Delete: resourceCloudbuildWorkerPoolDelete,

		Importer: &schema.ResourceImporter{
			State: resourceCloudbuildWorkerPoolImport,
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
				Description: ``,
			},

			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},

			"network_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        CloudbuildWorkerPoolNetworkConfigSchema(),
			},

			"project": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      ``,
			},

			"worker_config": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        CloudbuildWorkerPoolWorkerConfigSchema(),
			},

			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},

			"delete_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},

			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},

			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},
		},
	}
}

func CloudbuildWorkerPoolNetworkConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"peered_network": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareResourceNames,
				Description:      ``,
			},
		},
	}
}

func CloudbuildWorkerPoolWorkerConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"disk_size_gb": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: ``,
			},

			"machine_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: ``,
			},

			"no_external_ip": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: ``,
			},
		},
	}
}

func resourceCloudbuildWorkerPoolCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &cloudbuild.WorkerPool{
		Location:      dcl.String(d.Get("location").(string)),
		Name:          dcl.String(d.Get("name").(string)),
		NetworkConfig: expandCloudbuildWorkerPoolNetworkConfig(d.Get("network_config")),
		Project:       dcl.String(project),
		WorkerConfig:  expandCloudbuildWorkerPoolWorkerConfig(d.Get("worker_config")),
	}

	id, err := replaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/workerPools/{{name}}")
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
	client := NewDCLCloudbuildClient(config, userAgent, billingProject)
	res, err := client.ApplyWorkerPool(context.Background(), obj, createDirective...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error creating WorkerPool: %s", err)
	}

	log.Printf("[DEBUG] Finished creating WorkerPool %q: %#v", d.Id(), res)

	return resourceCloudbuildWorkerPoolRead(d, meta)
}

func resourceCloudbuildWorkerPoolRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &cloudbuild.WorkerPool{
		Location:      dcl.String(d.Get("location").(string)),
		Name:          dcl.String(d.Get("name").(string)),
		NetworkConfig: expandCloudbuildWorkerPoolNetworkConfig(d.Get("network_config")),
		Project:       dcl.String(project),
		WorkerConfig:  expandCloudbuildWorkerPoolWorkerConfig(d.Get("worker_config")),
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
	client := NewDCLCloudbuildClient(config, userAgent, billingProject)
	res, err := client.GetWorkerPool(context.Background(), obj)
	if err != nil {
		resourceName := fmt.Sprintf("CloudbuildWorkerPool %q", d.Id())
		return handleNotFoundDCLError(err, d, resourceName)
	}

	if err = d.Set("location", res.Location); err != nil {
		return fmt.Errorf("error setting location in state: %s", err)
	}
	if err = d.Set("name", res.Name); err != nil {
		return fmt.Errorf("error setting name in state: %s", err)
	}
	if err = d.Set("network_config", flattenCloudbuildWorkerPoolNetworkConfig(res.NetworkConfig)); err != nil {
		return fmt.Errorf("error setting network_config in state: %s", err)
	}
	if err = d.Set("project", res.Project); err != nil {
		return fmt.Errorf("error setting project in state: %s", err)
	}
	if err = d.Set("worker_config", flattenCloudbuildWorkerPoolWorkerConfig(res.WorkerConfig)); err != nil {
		return fmt.Errorf("error setting worker_config in state: %s", err)
	}
	if err = d.Set("create_time", res.CreateTime); err != nil {
		return fmt.Errorf("error setting create_time in state: %s", err)
	}
	if err = d.Set("delete_time", res.DeleteTime); err != nil {
		return fmt.Errorf("error setting delete_time in state: %s", err)
	}
	if err = d.Set("state", res.State); err != nil {
		return fmt.Errorf("error setting state in state: %s", err)
	}
	if err = d.Set("update_time", res.UpdateTime); err != nil {
		return fmt.Errorf("error setting update_time in state: %s", err)
	}

	return nil
}
func resourceCloudbuildWorkerPoolUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &cloudbuild.WorkerPool{
		Location:      dcl.String(d.Get("location").(string)),
		Name:          dcl.String(d.Get("name").(string)),
		NetworkConfig: expandCloudbuildWorkerPoolNetworkConfig(d.Get("network_config")),
		Project:       dcl.String(project),
		WorkerConfig:  expandCloudbuildWorkerPoolWorkerConfig(d.Get("worker_config")),
	}
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
	client := NewDCLCloudbuildClient(config, userAgent, billingProject)
	res, err := client.ApplyWorkerPool(context.Background(), obj, directive...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error updating WorkerPool: %s", err)
	}

	log.Printf("[DEBUG] Finished creating WorkerPool %q: %#v", d.Id(), res)

	return resourceCloudbuildWorkerPoolRead(d, meta)
}

func resourceCloudbuildWorkerPoolDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &cloudbuild.WorkerPool{
		Location:      dcl.String(d.Get("location").(string)),
		Name:          dcl.String(d.Get("name").(string)),
		NetworkConfig: expandCloudbuildWorkerPoolNetworkConfig(d.Get("network_config")),
		Project:       dcl.String(project),
		WorkerConfig:  expandCloudbuildWorkerPoolWorkerConfig(d.Get("worker_config")),
	}

	log.Printf("[DEBUG] Deleting WorkerPool %q", d.Id())
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLCloudbuildClient(config, userAgent, billingProject)
	if err := client.DeleteWorkerPool(context.Background(), obj); err != nil {
		return fmt.Errorf("Error deleting WorkerPool: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting WorkerPool %q", d.Id())
	return nil
}

func resourceCloudbuildWorkerPoolImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/workerPools/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/workerPools/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func expandCloudbuildWorkerPoolNetworkConfig(o interface{}) *cloudbuild.WorkerPoolNetworkConfig {
	if o == nil {
		return cloudbuild.EmptyWorkerPoolNetworkConfig
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return cloudbuild.EmptyWorkerPoolNetworkConfig
	}
	obj := objArr[0].(map[string]interface{})
	return &cloudbuild.WorkerPoolNetworkConfig{
		PeeredNetwork: dcl.String(obj["peered_network"].(string)),
	}
}

func flattenCloudbuildWorkerPoolNetworkConfig(obj *cloudbuild.WorkerPoolNetworkConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"peered_network": obj.PeeredNetwork,
	}

	return []interface{}{transformed}

}

func expandCloudbuildWorkerPoolWorkerConfig(o interface{}) *cloudbuild.WorkerPoolWorkerConfig {
	if o == nil {
		return nil
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return nil
	}
	obj := objArr[0].(map[string]interface{})
	return &cloudbuild.WorkerPoolWorkerConfig{
		DiskSizeGb:   dcl.Int64(int64(obj["disk_size_gb"].(int))),
		MachineType:  dcl.String(obj["machine_type"].(string)),
		NoExternalIP: dcl.Bool(obj["no_external_ip"].(bool)),
	}
}

func flattenCloudbuildWorkerPoolWorkerConfig(obj *cloudbuild.WorkerPoolWorkerConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"disk_size_gb":   obj.DiskSizeGb,
		"machine_type":   obj.MachineType,
		"no_external_ip": obj.NoExternalIP,
	}

	return []interface{}{transformed}

}
