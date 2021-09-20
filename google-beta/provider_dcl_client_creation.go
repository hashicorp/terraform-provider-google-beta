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
	dcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"

	assuredworkloads "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/assuredworkloads/beta"
	cloudbuild "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbuild/beta"
	cloudresourcemanager "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudresourcemanager/beta"
	compute "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta"
	dataproc "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc/beta"
	eventarc "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/eventarc/beta"
	gkehub "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkehub/beta"
	orgpolicy "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/orgpolicy/beta"
	privateca "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/privateca/beta"
)

func NewDCLAssuredWorkloadsClient(config *Config, userAgent, billingProject string) *assuredworkloads.Client {
	configOptions := []dcl.ConfigOption{
		dcl.WithHTTPClient(config.client),
		dcl.WithUserAgent(userAgent),
		dcl.WithLogger(dclLogger{}),
		dcl.WithBasePath(config.AssuredWorkloadsBasePath),
	}

	if config.UserProjectOverride {
		configOptions = append(configOptions, dcl.WithUserProjectOverride())
		if billingProject != "" {
			configOptions = append(configOptions, dcl.WithBillingProject(billingProject))
		}
	}

	dclConfig := dcl.NewConfig(configOptions...)
	return assuredworkloads.NewClient(dclConfig)
}

func NewDCLCloudbuildClient(config *Config, userAgent, billingProject string) *cloudbuild.Client {
	configOptions := []dcl.ConfigOption{
		dcl.WithHTTPClient(config.client),
		dcl.WithUserAgent(userAgent),
		dcl.WithLogger(dclLogger{}),
		dcl.WithBasePath(config.CloudBuildWorkerPoolBasePath),
	}

	if config.UserProjectOverride {
		configOptions = append(configOptions, dcl.WithUserProjectOverride())
		if billingProject != "" {
			configOptions = append(configOptions, dcl.WithBillingProject(billingProject))
		}
	}

	dclConfig := dcl.NewConfig(configOptions...)
	return cloudbuild.NewClient(dclConfig)
}

func NewDCLCloudResourceManagerClient(config *Config, userAgent, billingProject string) *cloudresourcemanager.Client {
	configOptions := []dcl.ConfigOption{
		dcl.WithHTTPClient(config.client),
		dcl.WithUserAgent(userAgent),
		dcl.WithLogger(dclLogger{}),
		dcl.WithBasePath(config.CloudResourceManagerBasePath),
	}

	if config.UserProjectOverride {
		configOptions = append(configOptions, dcl.WithUserProjectOverride())
		if billingProject != "" {
			configOptions = append(configOptions, dcl.WithBillingProject(billingProject))
		}
	}

	dclConfig := dcl.NewConfig(configOptions...)
	return cloudresourcemanager.NewClient(dclConfig)
}

func NewDCLComputeClient(config *Config, userAgent, billingProject string) *compute.Client {
	configOptions := []dcl.ConfigOption{
		dcl.WithHTTPClient(config.client),
		dcl.WithUserAgent(userAgent),
		dcl.WithLogger(dclLogger{}),
		dcl.WithBasePath(config.ComputeBasePath),
	}

	if config.UserProjectOverride {
		configOptions = append(configOptions, dcl.WithUserProjectOverride())
		if billingProject != "" {
			configOptions = append(configOptions, dcl.WithBillingProject(billingProject))
		}
	}

	dclConfig := dcl.NewConfig(configOptions...)
	return compute.NewClient(dclConfig)
}

func NewDCLDataprocClient(config *Config, userAgent, billingProject string) *dataproc.Client {
	configOptions := []dcl.ConfigOption{
		dcl.WithHTTPClient(config.client),
		dcl.WithUserAgent(userAgent),
		dcl.WithLogger(dclLogger{}),
		dcl.WithBasePath(config.DataprocBasePath),
	}

	if config.UserProjectOverride {
		configOptions = append(configOptions, dcl.WithUserProjectOverride())
		if billingProject != "" {
			configOptions = append(configOptions, dcl.WithBillingProject(billingProject))
		}
	}

	dclConfig := dcl.NewConfig(configOptions...)
	return dataproc.NewClient(dclConfig)
}

func NewDCLEventarcClient(config *Config, userAgent, billingProject string) *eventarc.Client {
	configOptions := []dcl.ConfigOption{
		dcl.WithHTTPClient(config.client),
		dcl.WithUserAgent(userAgent),
		dcl.WithLogger(dclLogger{}),
		dcl.WithBasePath(config.EventarcBasePath),
	}

	if config.UserProjectOverride {
		configOptions = append(configOptions, dcl.WithUserProjectOverride())
		if billingProject != "" {
			configOptions = append(configOptions, dcl.WithBillingProject(billingProject))
		}
	}

	dclConfig := dcl.NewConfig(configOptions...)
	return eventarc.NewClient(dclConfig)
}

func NewDCLGkeHubClient(config *Config, userAgent, billingProject string) *gkehub.Client {
	configOptions := []dcl.ConfigOption{
		dcl.WithHTTPClient(config.client),
		dcl.WithUserAgent(userAgent),
		dcl.WithLogger(dclLogger{}),
		dcl.WithBasePath(config.GkeHubBasePath),
	}

	if config.UserProjectOverride {
		configOptions = append(configOptions, dcl.WithUserProjectOverride())
		if billingProject != "" {
			configOptions = append(configOptions, dcl.WithBillingProject(billingProject))
		}
	}

	dclConfig := dcl.NewConfig(configOptions...)
	return gkehub.NewClient(dclConfig)
}

func NewDCLOrgPolicyClient(config *Config, userAgent, billingProject string) *orgpolicy.Client {
	configOptions := []dcl.ConfigOption{
		dcl.WithHTTPClient(config.client),
		dcl.WithUserAgent(userAgent),
		dcl.WithLogger(dclLogger{}),
		dcl.WithBasePath(config.OrgPolicyBasePath),
	}

	if config.UserProjectOverride {
		configOptions = append(configOptions, dcl.WithUserProjectOverride())
		if billingProject != "" {
			configOptions = append(configOptions, dcl.WithBillingProject(billingProject))
		}
	}

	dclConfig := dcl.NewConfig(configOptions...)
	return orgpolicy.NewClient(dclConfig)
}

func NewDCLPrivatecaClient(config *Config, userAgent, billingProject string) *privateca.Client {
	configOptions := []dcl.ConfigOption{
		dcl.WithHTTPClient(config.client),
		dcl.WithUserAgent(userAgent),
		dcl.WithLogger(dclLogger{}),
		dcl.WithBasePath(config.PrivatecaBasePath),
	}

	if config.UserProjectOverride {
		configOptions = append(configOptions, dcl.WithUserProjectOverride())
		if billingProject != "" {
			configOptions = append(configOptions, dcl.WithBillingProject(billingProject))
		}
	}

	dclConfig := dcl.NewConfig(configOptions...)
	return privateca.NewClient(dclConfig)
}
