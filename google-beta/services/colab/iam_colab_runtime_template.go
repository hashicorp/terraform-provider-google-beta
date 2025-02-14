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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/colab/RuntimeTemplate.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/iam_policy.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package colab

import (
	"fmt"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgiamresource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

var ColabRuntimeTemplateIamSchema = map[string]*schema.Schema{
	"project": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"location": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"runtime_template": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
	},
}

type ColabRuntimeTemplateIamUpdater struct {
	project         string
	location        string
	runtimeTemplate string
	d               tpgresource.TerraformResourceData
	Config          *transport_tpg.Config
}

func ColabRuntimeTemplateIamUpdaterProducer(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgiamresource.ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, _ := tpgresource.GetProject(d, config)
	if project != "" {
		if err := d.Set("project", project); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
	}
	values["project"] = project
	location, _ := tpgresource.GetLocation(d, config)
	if location != "" {
		if err := d.Set("location", location); err != nil {
			return nil, fmt.Errorf("Error setting location: %s", err)
		}
	}
	values["location"] = location
	if v, ok := d.GetOk("runtime_template"); ok {
		values["runtime_template"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := tpgresource.GetImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/notebookRuntimeTemplates/(?P<runtime_template>[^/]+)", "(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<runtime_template>[^/]+)", "(?P<location>[^/]+)/(?P<runtime_template>[^/]+)", "(?P<runtime_template>[^/]+)"}, d, config, d.Get("runtime_template").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &ColabRuntimeTemplateIamUpdater{
		project:         values["project"],
		location:        values["location"],
		runtimeTemplate: values["runtime_template"],
		d:               d,
		Config:          config,
	}

	if err := d.Set("project", u.project); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("location", u.location); err != nil {
		return nil, fmt.Errorf("Error setting location: %s", err)
	}
	if err := d.Set("runtime_template", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting runtime_template: %s", err)
	}

	return u, nil
}

func ColabRuntimeTemplateIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	values := make(map[string]string)

	project, _ := tpgresource.GetProject(d, config)
	if project != "" {
		values["project"] = project
	}

	location, _ := tpgresource.GetLocation(d, config)
	if location != "" {
		values["location"] = location
	}

	m, err := tpgresource.GetImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/notebookRuntimeTemplates/(?P<runtime_template>[^/]+)", "(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<runtime_template>[^/]+)", "(?P<location>[^/]+)/(?P<runtime_template>[^/]+)", "(?P<runtime_template>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &ColabRuntimeTemplateIamUpdater{
		project:         values["project"],
		location:        values["location"],
		runtimeTemplate: values["runtime_template"],
		d:               d,
		Config:          config,
	}
	if err := d.Set("runtime_template", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting runtime_template: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *ColabRuntimeTemplateIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyRuntimeTemplateUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	project, err := tpgresource.GetProject(u.d, u.Config)
	if err != nil {
		return nil, err
	}
	var obj map[string]interface{}

	userAgent, err := tpgresource.GenerateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return nil, err
	}

	policy, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    u.Config,
		Method:    "POST",
		Project:   project,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
	})
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	out := &cloudresourcemanager.Policy{}
	err = tpgresource.Convert(policy, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a policy to a resource manager policy: {{err}}", err)
	}

	return out, nil
}

func (u *ColabRuntimeTemplateIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := tpgresource.ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyRuntimeTemplateUrl("setIamPolicy")
	if err != nil {
		return err
	}
	project, err := tpgresource.GetProject(u.d, u.Config)
	if err != nil {
		return err
	}

	userAgent, err := tpgresource.GenerateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return err
	}

	_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    u.Config,
		Method:    "POST",
		Project:   project,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   u.d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *ColabRuntimeTemplateIamUpdater) qualifyRuntimeTemplateUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{ColabBasePath}}%s:%s", fmt.Sprintf("projects/%s/locations/%s/notebookRuntimeTemplates/%s", u.project, u.location, u.runtimeTemplate), methodIdentifier)
	url, err := tpgresource.ReplaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *ColabRuntimeTemplateIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/locations/%s/notebookRuntimeTemplates/%s", u.project, u.location, u.runtimeTemplate)
}

func (u *ColabRuntimeTemplateIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-colab-runtimetemplate-%s", u.GetResourceId())
}

func (u *ColabRuntimeTemplateIamUpdater) DescribeResource() string {
	return fmt.Sprintf("colab runtimetemplate %q", u.GetResourceId())
}
