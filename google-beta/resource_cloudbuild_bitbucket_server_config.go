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

package google

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceCloudBuildBitbucketServerConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudBuildBitbucketServerConfigCreate,
		Read:   resourceCloudBuildBitbucketServerConfigRead,
		Update: resourceCloudBuildBitbucketServerConfigUpdate,
		Delete: resourceCloudBuildBitbucketServerConfigDelete,

		Importer: &schema.ResourceImporter{
			State: resourceCloudBuildBitbucketServerConfigImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Immutable. API Key that will be attached to webhook. Once this field has been set, it cannot be changed.
Changing this field will result in deleting/ recreating the resource.`,
			},
			"config_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The ID to use for the BitbucketServerConfig, which will become the final component of the BitbucketServerConfig's resource name.`,
			},
			"host_uri": {
				Type:     schema.TypeString,
				Required: true,
				Description: `Immutable. The URI of the Bitbucket Server host. Once this field has been set, it cannot be changed.
If you need to change it, please create another BitbucketServerConfig.`,
			},
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The location of this bitbucket server config.`,
			},
			"secrets": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `Secret Manager secrets needed by the config.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"admin_access_token_version_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The resource name for the admin access token's secret version.`,
						},
						"read_access_token_version_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The resource name for the read access token's secret version.`,
						},
						"webhook_secret_version_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
							Description: `Immutable. The resource name for the webhook secret's secret version. Once this field has been set, it cannot be changed.
Changing this field will result in deleting/ recreating the resource.`,
						},
					},
				},
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Username of the account Cloud Build will use on Bitbucket Server.`,
			},
			"connected_repositories": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: `Connected Bitbucket Server repositories for this config.`,
				Elem:        cloudbuildBitbucketServerConfigConnectedRepositoriesSchema(),
				// Default schema.HashSchema is used.
			},
			"peered_network": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `The network to be used when reaching out to the Bitbucket Server instance. The VPC network must be enabled for private service connection.
This should be set if the Bitbucket Server instance is hosted on-premises and not reachable by public internet. If this field is left empty,
no network peering will occur and calls to the Bitbucket Server instance will be made over the public internet. Must be in the format
projects/{project}/global/networks/{network}, where {project} is a project number or id and {network} is the name of a VPC network in the project.`,
			},
			"ssl_ca": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `SSL certificate to use for requests to Bitbucket Server. The format should be PEM format but the extension can be one of .pem, .cer, or .crt.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The resource name for the config.`,
			},
			"webhook_key": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. UUID included in webhook requests. The UUID is used to look up the corresponding config.`,
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

func cloudbuildBitbucketServerConfigConnectedRepositoriesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"project_key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Identifier for the project storing the repository.`,
			},
			"repo_slug": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Identifier for the repository.`,
			},
		},
	}
}

func resourceCloudBuildBitbucketServerConfigCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	hostUriProp, err := expandCloudBuildBitbucketServerConfigHostUri(d.Get("host_uri"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("host_uri"); !isEmptyValue(reflect.ValueOf(hostUriProp)) && (ok || !reflect.DeepEqual(v, hostUriProp)) {
		obj["hostUri"] = hostUriProp
	}
	secretsProp, err := expandCloudBuildBitbucketServerConfigSecrets(d.Get("secrets"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("secrets"); !isEmptyValue(reflect.ValueOf(secretsProp)) && (ok || !reflect.DeepEqual(v, secretsProp)) {
		obj["secrets"] = secretsProp
	}
	usernameProp, err := expandCloudBuildBitbucketServerConfigUsername(d.Get("username"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("username"); !isEmptyValue(reflect.ValueOf(usernameProp)) && (ok || !reflect.DeepEqual(v, usernameProp)) {
		obj["username"] = usernameProp
	}
	apiKeyProp, err := expandCloudBuildBitbucketServerConfigApiKey(d.Get("api_key"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("api_key"); !isEmptyValue(reflect.ValueOf(apiKeyProp)) && (ok || !reflect.DeepEqual(v, apiKeyProp)) {
		obj["apiKey"] = apiKeyProp
	}
	connectedRepositoriesProp, err := expandCloudBuildBitbucketServerConfigConnectedRepositories(d.Get("connected_repositories"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("connected_repositories"); !isEmptyValue(reflect.ValueOf(connectedRepositoriesProp)) && (ok || !reflect.DeepEqual(v, connectedRepositoriesProp)) {
		obj["connectedRepositories"] = connectedRepositoriesProp
	}
	peeredNetworkProp, err := expandCloudBuildBitbucketServerConfigPeeredNetwork(d.Get("peered_network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("peered_network"); !isEmptyValue(reflect.ValueOf(peeredNetworkProp)) && (ok || !reflect.DeepEqual(v, peeredNetworkProp)) {
		obj["peeredNetwork"] = peeredNetworkProp
	}
	sslCaProp, err := expandCloudBuildBitbucketServerConfigSslCa(d.Get("ssl_ca"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ssl_ca"); !isEmptyValue(reflect.ValueOf(sslCaProp)) && (ok || !reflect.DeepEqual(v, sslCaProp)) {
		obj["sslCa"] = sslCaProp
	}

	obj, err = resourceCloudBuildBitbucketServerConfigEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{CloudBuildBasePath}}projects/{{project}}/locations/{{location}}/bitbucketServerConfigs?bitbucketServerConfigId={{config_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new BitbucketServerConfig: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for BitbucketServerConfig: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating BitbucketServerConfig: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/bitbucketServerConfigs/{{config_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = CloudBuildOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating BitbucketServerConfig", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create BitbucketServerConfig: %s", err)
	}

	if err := d.Set("name", flattenCloudBuildBitbucketServerConfigName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/bitbucketServerConfigs/{{config_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating BitbucketServerConfig without connected repos: %q: %#v", d.Id(), res)

	if v, ok := d.GetOkExists("connected_repositories"); !isEmptyValue(reflect.ValueOf(connectedRepositoriesProp)) && (ok || !reflect.DeepEqual(v, connectedRepositoriesProp)) {
		connectedReposPropArray, ok := connectedRepositoriesProp.([]interface{})
		if !ok {
			return fmt.Errorf("Error reading connected_repositories")
		}

		requests := make([]interface{}, len(connectedReposPropArray))
		for i := 0; i < len(connectedReposPropArray); i++ {
			connectedRepo := make(map[string]interface{})
			connectedRepo["parent"] = id
			connectedRepo["repo"] = connectedReposPropArray[i]

			connectedRepoRequest := make(map[string]interface{})
			connectedRepoRequest["parent"] = id
			connectedRepoRequest["bitbucketServerConnectedRepository"] = connectedRepo

			requests[i] = connectedRepoRequest
		}
		obj = make(map[string]interface{})
		obj["requests"] = requests

		url, err = ReplaceVars(d, config, "{{CloudBuildBasePath}}projects/{{project}}/locations/{{location}}/bitbucketServerConfigs/{{config_id}}/connectedRepositories:batchCreate")
		if err != nil {
			return err
		}

		res, err = SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
		if err != nil {
			return fmt.Errorf("Error creating connected_repositories: %s", err)
		}

		err = CloudBuildOperationWaitTime(
			config, res, project, "Creating connected_repositories on BitbucketServerConfig", userAgent,
			d.Timeout(schema.TimeoutCreate))
		if err != nil {
			return fmt.Errorf("Error waiting to create connected_repositories: %s", err)
		}
	} else {
		log.Printf("[DEBUG] No connected repositories found to create: %#v", connectedRepositoriesProp)
	}

	log.Printf("[DEBUG] Finished creating BitbucketServerConfig %q: %#v", d.Id(), res)

	return resourceCloudBuildBitbucketServerConfigRead(d, meta)
}

func resourceCloudBuildBitbucketServerConfigRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{CloudBuildBasePath}}projects/{{project}}/locations/{{location}}/bitbucketServerConfigs/{{config_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for BitbucketServerConfig: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("CloudBuildBitbucketServerConfig %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading BitbucketServerConfig: %s", err)
	}

	if err := d.Set("name", flattenCloudBuildBitbucketServerConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading BitbucketServerConfig: %s", err)
	}
	if err := d.Set("host_uri", flattenCloudBuildBitbucketServerConfigHostUri(res["hostUri"], d, config)); err != nil {
		return fmt.Errorf("Error reading BitbucketServerConfig: %s", err)
	}
	if err := d.Set("secrets", flattenCloudBuildBitbucketServerConfigSecrets(res["secrets"], d, config)); err != nil {
		return fmt.Errorf("Error reading BitbucketServerConfig: %s", err)
	}
	if err := d.Set("username", flattenCloudBuildBitbucketServerConfigUsername(res["username"], d, config)); err != nil {
		return fmt.Errorf("Error reading BitbucketServerConfig: %s", err)
	}
	if err := d.Set("webhook_key", flattenCloudBuildBitbucketServerConfigWebhookKey(res["webhookKey"], d, config)); err != nil {
		return fmt.Errorf("Error reading BitbucketServerConfig: %s", err)
	}
	if err := d.Set("api_key", flattenCloudBuildBitbucketServerConfigApiKey(res["apiKey"], d, config)); err != nil {
		return fmt.Errorf("Error reading BitbucketServerConfig: %s", err)
	}
	if err := d.Set("connected_repositories", flattenCloudBuildBitbucketServerConfigConnectedRepositories(res["connectedRepositories"], d, config)); err != nil {
		return fmt.Errorf("Error reading BitbucketServerConfig: %s", err)
	}
	if err := d.Set("peered_network", flattenCloudBuildBitbucketServerConfigPeeredNetwork(res["peeredNetwork"], d, config)); err != nil {
		return fmt.Errorf("Error reading BitbucketServerConfig: %s", err)
	}
	if err := d.Set("ssl_ca", flattenCloudBuildBitbucketServerConfigSslCa(res["sslCa"], d, config)); err != nil {
		return fmt.Errorf("Error reading BitbucketServerConfig: %s", err)
	}

	return nil
}

func resourceCloudBuildBitbucketServerConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for BitbucketServerConfig: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	hostUriProp, err := expandCloudBuildBitbucketServerConfigHostUri(d.Get("host_uri"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("host_uri"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, hostUriProp)) {
		obj["hostUri"] = hostUriProp
	}
	secretsProp, err := expandCloudBuildBitbucketServerConfigSecrets(d.Get("secrets"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("secrets"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, secretsProp)) {
		obj["secrets"] = secretsProp
	}
	usernameProp, err := expandCloudBuildBitbucketServerConfigUsername(d.Get("username"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("username"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, usernameProp)) {
		obj["username"] = usernameProp
	}
	connectedRepositoriesProp, err := expandCloudBuildBitbucketServerConfigConnectedRepositories(d.Get("connected_repositories"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("connected_repositories"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, connectedRepositoriesProp)) {
		obj["connectedRepositories"] = connectedRepositoriesProp
	}
	peeredNetworkProp, err := expandCloudBuildBitbucketServerConfigPeeredNetwork(d.Get("peered_network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("peered_network"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, peeredNetworkProp)) {
		obj["peeredNetwork"] = peeredNetworkProp
	}
	sslCaProp, err := expandCloudBuildBitbucketServerConfigSslCa(d.Get("ssl_ca"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ssl_ca"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, sslCaProp)) {
		obj["sslCa"] = sslCaProp
	}

	obj, err = resourceCloudBuildBitbucketServerConfigEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{CloudBuildBasePath}}projects/{{project}}/locations/{{location}}/bitbucketServerConfigs/{{config_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating BitbucketServerConfig %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("host_uri") {
		updateMask = append(updateMask, "hostUri")
	}

	if d.HasChange("secrets") {
		updateMask = append(updateMask, "secrets")
	}

	if d.HasChange("username") {
		updateMask = append(updateMask, "username")
	}

	if d.HasChange("connected_repositories") {
		updateMask = append(updateMask, "connectedRepositories")
	}

	if d.HasChange("peered_network") {
		updateMask = append(updateMask, "peeredNetwork")
	}

	if d.HasChange("ssl_ca") {
		updateMask = append(updateMask, "sslCa")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	// remove connectedRepositories from updateMask
	for i, field := range updateMask {
		if field == "connectedRepositories" {
			updateMask = append(updateMask[:i], updateMask[i+1:]...)
			break
		}
	}
	// reconstruct url
	url, err = ReplaceVars(d, config, "{{CloudBuildBasePath}}projects/{{project}}/locations/{{location}}/bitbucketServerConfigs/{{config_id}}")
	if err != nil {
		return err
	}
	url, err = AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating BitbucketServerConfig %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating BitbucketServerConfig %q: %#v", d.Id(), res)
	}

	err = CloudBuildOperationWaitTime(
		config, res, project, "Updating BitbucketServerConfig", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	if d.HasChange("connected_repositories") {
		o, n := d.GetChange("connected_repositories")
		oReposSet, ok := o.(*schema.Set)
		if !ok {
			return fmt.Errorf("Error reading old connected repositories")
		}
		nReposSet, ok := n.(*schema.Set)
		if !ok {
			return fmt.Errorf("Error reading new connected repositories")
		}

		removeRepos := oReposSet.Difference(nReposSet).List()
		createRepos := nReposSet.Difference(oReposSet).List()

		url, err = ReplaceVars(d, config, "{{CloudBuildBasePath}}projects/{{project}}/locations/{{location}}/bitbucketServerConfigs/{{config_id}}:removeBitbucketServerConnectedRepository")
		if err != nil {
			return err
		}

		// send remove repo requests.
		for _, repo := range removeRepos {
			obj := make(map[string]interface{})
			obj["connectedRepository"] = repo
			res, err = SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
			if err != nil {
				return fmt.Errorf("Error removing connected_repositories: %s", err)
			}
		}

		// if repos to create, prepare and send batchCreate request
		if len(createRepos) > 0 {
			parent, err := ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/bitbucketServerConfigs/{{config_id}}")
			if err != nil {
				return fmt.Errorf("Error constructing id: %s", err)
			}
			var requests []interface{}
			for _, repo := range createRepos {
				connectedRepo := make(map[string]interface{})
				connectedRepo["parent"] = parent
				connectedRepo["repo"] = repo

				connectedRepoRequest := make(map[string]interface{})
				connectedRepoRequest["parent"] = parent
				connectedRepoRequest["bitbucketServerConnectedRepository"] = connectedRepo

				requests = append(requests, connectedRepoRequest)
			}
			obj = make(map[string]interface{})
			obj["requests"] = requests

			url, err = ReplaceVars(d, config, "{{CloudBuildBasePath}}projects/{{project}}/locations/{{location}}/bitbucketServerConfigs/{{config_id}}/connectedRepositories:batchCreate")
			if err != nil {
				return err
			}

			res, err = SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
			if err != nil {
				return fmt.Errorf("Error creating connected_repositories: %s", err)
			}

			err = CloudBuildOperationWaitTime(
				config, res, project, "Updating connected_repositories on BitbucketServerConfig", userAgent,
				d.Timeout(schema.TimeoutUpdate))
			if err != nil {
				return fmt.Errorf("Error waiting to create connected_repositories: %s", err)
			}
		}
	} else {
		log.Printf("[DEBUG] connected_repositories have no changes")
	}
	return resourceCloudBuildBitbucketServerConfigRead(d, meta)
}

func resourceCloudBuildBitbucketServerConfigDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for BitbucketServerConfig: %s", err)
	}
	billingProject = project

	url, err := ReplaceVars(d, config, "{{CloudBuildBasePath}}projects/{{project}}/locations/{{location}}/bitbucketServerConfigs/{{config_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting BitbucketServerConfig %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "BitbucketServerConfig")
	}

	err = CloudBuildOperationWaitTime(
		config, res, project, "Deleting BitbucketServerConfig", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting BitbucketServerConfig %q: %#v", d.Id(), res)
	return nil
}

func resourceCloudBuildBitbucketServerConfigImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/bitbucketServerConfigs/(?P<config_id>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<config_id>[^/]+)",
		"(?P<location>[^/]+)/(?P<config_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/bitbucketServerConfigs/{{config_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenCloudBuildBitbucketServerConfigName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudBuildBitbucketServerConfigHostUri(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudBuildBitbucketServerConfigSecrets(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["admin_access_token_version_name"] =
		flattenCloudBuildBitbucketServerConfigSecretsAdminAccessTokenVersionName(original["adminAccessTokenVersionName"], d, config)
	transformed["read_access_token_version_name"] =
		flattenCloudBuildBitbucketServerConfigSecretsReadAccessTokenVersionName(original["readAccessTokenVersionName"], d, config)
	transformed["webhook_secret_version_name"] =
		flattenCloudBuildBitbucketServerConfigSecretsWebhookSecretVersionName(original["webhookSecretVersionName"], d, config)
	return []interface{}{transformed}
}
func flattenCloudBuildBitbucketServerConfigSecretsAdminAccessTokenVersionName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudBuildBitbucketServerConfigSecretsReadAccessTokenVersionName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudBuildBitbucketServerConfigSecretsWebhookSecretVersionName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudBuildBitbucketServerConfigUsername(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudBuildBitbucketServerConfigWebhookKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudBuildBitbucketServerConfigApiKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudBuildBitbucketServerConfigConnectedRepositories(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(schema.HashResource(cloudbuildBitbucketServerConfigConnectedRepositoriesSchema()), []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"project_key": flattenCloudBuildBitbucketServerConfigConnectedRepositoriesProjectKey(original["projectKey"], d, config),
			"repo_slug":   flattenCloudBuildBitbucketServerConfigConnectedRepositoriesRepoSlug(original["repoSlug"], d, config),
		})
	}
	return transformed
}
func flattenCloudBuildBitbucketServerConfigConnectedRepositoriesProjectKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudBuildBitbucketServerConfigConnectedRepositoriesRepoSlug(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudBuildBitbucketServerConfigPeeredNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudBuildBitbucketServerConfigSslCa(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandCloudBuildBitbucketServerConfigHostUri(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildBitbucketServerConfigSecrets(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAdminAccessTokenVersionName, err := expandCloudBuildBitbucketServerConfigSecretsAdminAccessTokenVersionName(original["admin_access_token_version_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAdminAccessTokenVersionName); val.IsValid() && !isEmptyValue(val) {
		transformed["adminAccessTokenVersionName"] = transformedAdminAccessTokenVersionName
	}

	transformedReadAccessTokenVersionName, err := expandCloudBuildBitbucketServerConfigSecretsReadAccessTokenVersionName(original["read_access_token_version_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedReadAccessTokenVersionName); val.IsValid() && !isEmptyValue(val) {
		transformed["readAccessTokenVersionName"] = transformedReadAccessTokenVersionName
	}

	transformedWebhookSecretVersionName, err := expandCloudBuildBitbucketServerConfigSecretsWebhookSecretVersionName(original["webhook_secret_version_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWebhookSecretVersionName); val.IsValid() && !isEmptyValue(val) {
		transformed["webhookSecretVersionName"] = transformedWebhookSecretVersionName
	}

	return transformed, nil
}

func expandCloudBuildBitbucketServerConfigSecretsAdminAccessTokenVersionName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildBitbucketServerConfigSecretsReadAccessTokenVersionName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildBitbucketServerConfigSecretsWebhookSecretVersionName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildBitbucketServerConfigUsername(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildBitbucketServerConfigApiKey(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildBitbucketServerConfigConnectedRepositories(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedProjectKey, err := expandCloudBuildBitbucketServerConfigConnectedRepositoriesProjectKey(original["project_key"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedProjectKey); val.IsValid() && !isEmptyValue(val) {
			transformed["projectKey"] = transformedProjectKey
		}

		transformedRepoSlug, err := expandCloudBuildBitbucketServerConfigConnectedRepositoriesRepoSlug(original["repo_slug"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRepoSlug); val.IsValid() && !isEmptyValue(val) {
			transformed["repoSlug"] = transformedRepoSlug
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCloudBuildBitbucketServerConfigConnectedRepositoriesProjectKey(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildBitbucketServerConfigConnectedRepositoriesRepoSlug(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildBitbucketServerConfigPeeredNetwork(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildBitbucketServerConfigSslCa(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceCloudBuildBitbucketServerConfigEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// connectedRepositories is needed for batchCreate on the config after creation.
	delete(obj, "connectedRepositories")
	return obj, nil
}
