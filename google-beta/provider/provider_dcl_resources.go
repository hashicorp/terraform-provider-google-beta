// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

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

package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/apikeys"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/assuredworkloads"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/cloudbuild"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/clouddeploy"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/containeraws"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/containerazure"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/dataplex"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/dataproc"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/firebaserules"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/gkehub"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/recaptchaenterprise"
)

var dclResources = map[string]*schema.Resource{
	"google_apikeys_key":                   apikeys.ResourceApikeysKey(),
	"google_assured_workloads_workload":    assuredworkloads.ResourceAssuredWorkloadsWorkload(),
	"google_cloudbuild_worker_pool":        cloudbuild.ResourceCloudbuildWorkerPool(),
	"google_clouddeploy_delivery_pipeline": clouddeploy.ResourceClouddeployDeliveryPipeline(),
	"google_clouddeploy_target":            clouddeploy.ResourceClouddeployTarget(),
	"google_container_aws_cluster":         containeraws.ResourceContainerAwsCluster(),
	"google_container_aws_node_pool":       containeraws.ResourceContainerAwsNodePool(),
	"google_container_azure_client":        containerazure.ResourceContainerAzureClient(),
	"google_container_azure_cluster":       containerazure.ResourceContainerAzureCluster(),
	"google_container_azure_node_pool":     containerazure.ResourceContainerAzureNodePool(),
	"google_dataplex_asset":                dataplex.ResourceDataplexAsset(),
	"google_dataplex_lake":                 dataplex.ResourceDataplexLake(),
	"google_dataplex_zone":                 dataplex.ResourceDataplexZone(),
	"google_dataproc_workflow_template":    dataproc.ResourceDataprocWorkflowTemplate(),
	"google_firebaserules_release":         firebaserules.ResourceFirebaserulesRelease(),
	"google_firebaserules_ruleset":         firebaserules.ResourceFirebaserulesRuleset(),
	"google_gke_hub_feature_membership":    gkehub.ResourceGkeHubFeatureMembership(),
	"google_recaptcha_enterprise_key":      recaptchaenterprise.ResourceRecaptchaEnterpriseKey(),
}
