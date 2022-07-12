// ----------------------------------------------------------------------------
//
//	***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//	This file is automatically generated by Magic Modules and manual
//	changes will be clobbered when the file is regenerated.
//
//	Please read more about how to change this file in
//	.github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------
package google

import (
	"fmt"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"
)

var IapAppEngineVersionIamSchema = map[string]*schema.Schema{
	"project": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"app_id": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	"service": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	"version_id": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type IapAppEngineVersionIamUpdater struct {
	project   string
	appId     string
	service   string
	versionId string
	d         TerraformResourceData
	Config    *Config
}

func IapAppEngineVersionIamUpdaterProducer(d TerraformResourceData, config *Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		if err := d.Set("project", project); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
	}
	values["project"] = project
	if v, ok := d.GetOk("app_id"); ok {
		values["appId"] = v.(string)
	}

	if v, ok := d.GetOk("service"); ok {
		values["service"] = v.(string)
	}

	if v, ok := d.GetOk("version_id"); ok {
		values["versionId"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/iap_web/appengine-(?P<appId>[^/]+)/services/(?P<service>[^/]+)/versions/(?P<versionId>[^/]+)", "(?P<project>[^/]+)/(?P<appId>[^/]+)/(?P<service>[^/]+)/(?P<versionId>[^/]+)", "(?P<appId>[^/]+)/(?P<service>[^/]+)/(?P<versionId>[^/]+)", "(?P<version>[^/]+)"}, d, config, d.Get("version_id").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &IapAppEngineVersionIamUpdater{
		project:   values["project"],
		appId:     values["appId"],
		service:   values["service"],
		versionId: values["versionId"],
		d:         d,
		Config:    config,
	}

	if err := d.Set("project", u.project); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("app_id", u.appId); err != nil {
		return nil, fmt.Errorf("Error setting app_id: %s", err)
	}
	if err := d.Set("service", u.service); err != nil {
		return nil, fmt.Errorf("Error setting service: %s", err)
	}
	if err := d.Set("version_id", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting version_id: %s", err)
	}

	return u, nil
}

func IapAppEngineVersionIdParseFunc(d *schema.ResourceData, config *Config) error {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		values["project"] = project
	}

	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/iap_web/appengine-(?P<appId>[^/]+)/services/(?P<service>[^/]+)/versions/(?P<versionId>[^/]+)", "(?P<project>[^/]+)/(?P<appId>[^/]+)/(?P<service>[^/]+)/(?P<versionId>[^/]+)", "(?P<appId>[^/]+)/(?P<service>[^/]+)/(?P<versionId>[^/]+)", "(?P<version>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &IapAppEngineVersionIamUpdater{
		project:   values["project"],
		appId:     values["appId"],
		service:   values["service"],
		versionId: values["versionId"],
		d:         d,
		Config:    config,
	}
	if err := d.Set("version_id", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting version_id: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *IapAppEngineVersionIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyAppEngineVersionUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	project, err := getProject(u.d, u.Config)
	if err != nil {
		return nil, err
	}
	var obj map[string]interface{}
	obj = map[string]interface{}{
		"options": map[string]interface{}{
			"requestedPolicyVersion": iamPolicyVersion,
		},
	}

	userAgent, err := generateUserAgentString(u.d, u.Config.userAgent)
	if err != nil {
		return nil, err
	}

	policy, err := sendRequest(u.Config, "POST", project, url, userAgent, obj)
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	out := &cloudresourcemanager.Policy{}
	err = Convert(policy, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a policy to a resource manager policy: {{err}}", err)
	}

	return out, nil
}

func (u *IapAppEngineVersionIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyAppEngineVersionUrl("setIamPolicy")
	if err != nil {
		return err
	}
	project, err := getProject(u.d, u.Config)
	if err != nil {
		return err
	}

	userAgent, err := generateUserAgentString(u.d, u.Config.userAgent)
	if err != nil {
		return err
	}

	_, err = sendRequestWithTimeout(u.Config, "POST", project, url, userAgent, obj, u.d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *IapAppEngineVersionIamUpdater) qualifyAppEngineVersionUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{IapBasePath}}%s:%s", fmt.Sprintf("projects/%s/iap_web/appengine-%s/services/%s/versions/%s", u.project, u.appId, u.service, u.versionId), methodIdentifier)
	url, err := replaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *IapAppEngineVersionIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/iap_web/appengine-%s/services/%s/versions/%s", u.project, u.appId, u.service, u.versionId)
}

func (u *IapAppEngineVersionIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-iap-appengineversion-%s", u.GetResourceId())
}

func (u *IapAppEngineVersionIamUpdater) DescribeResource() string {
	return fmt.Sprintf("iap appengineversion %q", u.GetResourceId())
}
