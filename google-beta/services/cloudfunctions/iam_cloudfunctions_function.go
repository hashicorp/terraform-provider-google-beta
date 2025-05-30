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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/cloudfunctions/CloudFunction.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/iam_policy.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package cloudfunctions

import (
	"fmt"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgiamresource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

var CloudFunctionsCloudFunctionIamSchema = map[string]*schema.Schema{
	"project": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"region": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"cloud_function": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
	},
}

type CloudFunctionsCloudFunctionIamUpdater struct {
	project       string
	region        string
	cloudFunction string
	d             tpgresource.TerraformResourceData
	Config        *transport_tpg.Config
}

func CloudFunctionsCloudFunctionIamUpdaterProducer(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgiamresource.ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, _ := tpgresource.GetProject(d, config)
	if project != "" {
		if err := d.Set("project", project); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
	}
	values["project"] = project
	region, _ := tpgresource.GetRegion(d, config)
	if region != "" {
		if err := d.Set("region", region); err != nil {
			return nil, fmt.Errorf("Error setting region: %s", err)
		}
	}
	values["region"] = region
	if v, ok := d.GetOk("cloud_function"); ok {
		values["cloud_function"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := tpgresource.GetImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/functions/(?P<cloud_function>[^/]+)", "(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<cloud_function>[^/]+)", "(?P<region>[^/]+)/(?P<cloud_function>[^/]+)", "(?P<cloud_function>[^/]+)"}, d, config, d.Get("cloud_function").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &CloudFunctionsCloudFunctionIamUpdater{
		project:       values["project"],
		region:        values["region"],
		cloudFunction: values["cloud_function"],
		d:             d,
		Config:        config,
	}

	if err := d.Set("project", u.project); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("region", u.region); err != nil {
		return nil, fmt.Errorf("Error setting region: %s", err)
	}
	if err := d.Set("cloud_function", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting cloud_function: %s", err)
	}

	return u, nil
}

func CloudFunctionsCloudFunctionIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	values := make(map[string]string)

	project, _ := tpgresource.GetProject(d, config)
	if project != "" {
		values["project"] = project
	}

	region, _ := tpgresource.GetRegion(d, config)
	if region != "" {
		values["region"] = region
	}

	m, err := tpgresource.GetImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/functions/(?P<cloud_function>[^/]+)", "(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<cloud_function>[^/]+)", "(?P<region>[^/]+)/(?P<cloud_function>[^/]+)", "(?P<cloud_function>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &CloudFunctionsCloudFunctionIamUpdater{
		project:       values["project"],
		region:        values["region"],
		cloudFunction: values["cloud_function"],
		d:             d,
		Config:        config,
	}
	if err := d.Set("cloud_function", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting cloud_function: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *CloudFunctionsCloudFunctionIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyCloudFunctionUrl("getIamPolicy")
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
		Method:    "GET",
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

func (u *CloudFunctionsCloudFunctionIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := tpgresource.ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyCloudFunctionUrl("setIamPolicy")
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

func (u *CloudFunctionsCloudFunctionIamUpdater) qualifyCloudFunctionUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{CloudFunctionsBasePath}}%s:%s", fmt.Sprintf("projects/%s/locations/%s/functions/%s", u.project, u.region, u.cloudFunction), methodIdentifier)
	url, err := tpgresource.ReplaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *CloudFunctionsCloudFunctionIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/locations/%s/functions/%s", u.project, u.region, u.cloudFunction)
}

func (u *CloudFunctionsCloudFunctionIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-cloudfunctions-cloudfunction-%s", u.GetResourceId())
}

func (u *CloudFunctionsCloudFunctionIamUpdater) DescribeResource() string {
	return fmt.Sprintf("cloudfunctions cloudfunction %q", u.GetResourceId())
}
