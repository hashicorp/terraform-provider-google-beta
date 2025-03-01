---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/networksecurity/MirroringDeploymentGroup.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Network Security"
description: |-
  A Deployment Group represents the collector deployments across different zones within an organization.
---

# google_network_security_mirroring_deployment_group

A Deployment Group represents the collector deployments across different zones within an organization.

~> **Warning:** This resource is in beta, and should be used with the terraform-provider-google-beta provider.
See [Provider Versions](https://terraform.io/docs/providers/google/guides/provider_versions.html) for more details on beta resources.

To get more information about MirroringDeploymentGroup, see:

* [API documentation](https://cloud.google.com/network-security-integration/docs/reference/rest/v1beta1/projects.locations.mirroringDeploymentGroups)
* How-to Guides
    * [Mirroring deployment group overview](https://cloud.google.com/network-security-integration/docs/out-of-band/deployment-groups-overview)

<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=network_security_mirroring_deployment_group_basic&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Network Security Mirroring Deployment Group Basic


```hcl
resource "google_compute_network" "network" {
  provider                = google-beta
  name                    = "example-network"
  auto_create_subnetworks = false
}

resource "google_network_security_mirroring_deployment_group" "default" {
  provider                      = google-beta
  mirroring_deployment_group_id = "example-dg"
  location                      = "global"
  network                       = google_compute_network.network.id
  labels = {
    foo = "bar"
  }
}
```

## Argument Reference

The following arguments are supported:


* `network` -
  (Required)
  Required. Immutable. The network that is being used for the deployment. Format is:
  projects/{project}/global/networks/{network}.

* `location` -
  (Required)
  Resource ID segment making up resource `name`. It identifies the resource within its parent collection as described in https://google.aip.dev/122. See documentation for resource type `networksecurity.googleapis.com/MirroringDeploymentGroup`.

* `mirroring_deployment_group_id` -
  (Required)
  Required. Id of the requesting object
  If auto-generating Id server-side, remove this field and
  mirroring_deployment_group_id from the method_signature of Create RPC


- - -


* `labels` -
  (Optional)
  Optional. Labels as key value pairs 
  **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
  Please refer to the field `effective_labels` for all of the labels present on the resource.

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.


## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/{{location}}/mirroringDeploymentGroups/{{mirroring_deployment_group_id}}`

* `name` -
  Immutable. Identifier. Then name of the MirroringDeploymentGroup.

* `create_time` -
  Output only. [Output only] Create time stamp

* `update_time` -
  Output only. [Output only] Update time stamp

* `connected_endpoint_groups` -
  Output only. The list of Mirroring Endpoint Groups that are connected to this resource.
  Structure is [documented below](#nested_connected_endpoint_groups).

* `state` -
  Output only. Current state of the deployment group. 
   Possible values:
   STATE_UNSPECIFIED
  ACTIVE
  CREATING
  DELETING

* `reconciling` -
  Output only. Whether reconciling is in progress, recommended per
  https://google.aip.dev/128.

* `terraform_labels` -
  The combination of labels configured directly on the resource
   and default labels configured on the provider.

* `effective_labels` -
  All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.


<a name="nested_connected_endpoint_groups"></a>The `connected_endpoint_groups` block contains:

* `name` -
  (Output)
  Output only. A connected mirroring endpoint group.

## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


MirroringDeploymentGroup can be imported using any of these accepted formats:

* `projects/{{project}}/locations/{{location}}/mirroringDeploymentGroups/{{mirroring_deployment_group_id}}`
* `{{project}}/{{location}}/{{mirroring_deployment_group_id}}`
* `{{location}}/{{mirroring_deployment_group_id}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import MirroringDeploymentGroup using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/locations/{{location}}/mirroringDeploymentGroups/{{mirroring_deployment_group_id}}"
  to = google_network_security_mirroring_deployment_group.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), MirroringDeploymentGroup can be imported using one of the formats above. For example:

```
$ terraform import google_network_security_mirroring_deployment_group.default projects/{{project}}/locations/{{location}}/mirroringDeploymentGroups/{{mirroring_deployment_group_id}}
$ terraform import google_network_security_mirroring_deployment_group.default {{project}}/{{location}}/{{mirroring_deployment_group_id}}
$ terraform import google_network_security_mirroring_deployment_group.default {{location}}/{{mirroring_deployment_group_id}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
