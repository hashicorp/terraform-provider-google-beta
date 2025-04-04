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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/servicemanagement/ServiceConsumers.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/iam_policy.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package servicemanagement

import (
	"fmt"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgiamresource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

var ServiceManagementServiceConsumersIamSchema = map[string]*schema.Schema{
	"service_name": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	"consumer_project": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
	},
}

type ServiceManagementServiceConsumersIamUpdater struct {
	serviceName     string
	consumerProject string
	d               tpgresource.TerraformResourceData
	Config          *transport_tpg.Config
}

func ServiceManagementServiceConsumersIamUpdaterProducer(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgiamresource.ResourceIamUpdater, error) {
	values := make(map[string]string)

	if v, ok := d.GetOk("service_name"); ok {
		values["service_name"] = v.(string)
	}

	if v, ok := d.GetOk("consumer_project"); ok {
		values["consumer_project"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := tpgresource.GetImportIdQualifiers([]string{"services/(?P<service_name>[^/]+)/consumers/(?P<consumer_project>[^/]+)", "(?P<service_name>[^/]+)/(?P<consumer_project>[^/]+)", "(?P<consumer_project>[^/]+)"}, d, config, d.Get("consumer_project").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &ServiceManagementServiceConsumersIamUpdater{
		serviceName:     values["service_name"],
		consumerProject: values["consumer_project"],
		d:               d,
		Config:          config,
	}

	if err := d.Set("service_name", u.serviceName); err != nil {
		return nil, fmt.Errorf("Error setting service_name: %s", err)
	}
	if err := d.Set("consumer_project", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting consumer_project: %s", err)
	}

	return u, nil
}

func ServiceManagementServiceConsumersIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	values := make(map[string]string)

	m, err := tpgresource.GetImportIdQualifiers([]string{"services/(?P<service_name>[^/]+)/consumers/(?P<consumer_project>[^/]+)", "(?P<service_name>[^/]+)/(?P<consumer_project>[^/]+)", "(?P<consumer_project>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &ServiceManagementServiceConsumersIamUpdater{
		serviceName:     values["service_name"],
		consumerProject: values["consumer_project"],
		d:               d,
		Config:          config,
	}
	if err := d.Set("consumer_project", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting consumer_project: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *ServiceManagementServiceConsumersIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyServiceConsumersUrl("getIamPolicy")
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

func (u *ServiceManagementServiceConsumersIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := tpgresource.ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyServiceConsumersUrl("setIamPolicy")
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

func (u *ServiceManagementServiceConsumersIamUpdater) qualifyServiceConsumersUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{ServiceManagementBasePath}}%s:%s", fmt.Sprintf("services/%s/consumers/%s", u.serviceName, u.consumerProject), methodIdentifier)
	url, err := tpgresource.ReplaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *ServiceManagementServiceConsumersIamUpdater) GetResourceId() string {
	return fmt.Sprintf("services/%s/consumers/%s", u.serviceName, u.consumerProject)
}

func (u *ServiceManagementServiceConsumersIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-servicemanagement-serviceconsumers-%s", u.GetResourceId())
}

func (u *ServiceManagementServiceConsumersIamUpdater) DescribeResource() string {
	return fmt.Sprintf("servicemanagement serviceconsumers %q", u.GetResourceId())
}
