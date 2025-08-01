---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/artifactregistry/VPCSCConfig.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Artifact Registry"
description: |-
  The Artifact Registry VPC SC config that applies to a Project.
---

# google_artifact_registry_vpcsc_config

The Artifact Registry VPC SC config that applies to a Project.

~> **Warning:** This resource is in beta, and should be used with the terraform-provider-google-beta provider.
See [Provider Versions](https://terraform.io/docs/providers/google/guides/provider_versions.html) for more details on beta resources.

To get more information about VPCSCConfig, see:

* [API documentation](https://cloud.google.com/artifact-registry/docs/reference/rest/v1/VPCSCConfig)

~> **Note:** VPC SC configs are automatically created for a given location. Creating a
resource of this type will acquire and update the resource that already
exists at the location. Deleting this resource will remove the config from
your Terraform state but leave the resource as is.
## Example Usage - Artifact Registry Vpcsc Config


```hcl
resource "google_artifact_registry_vpcsc_config" "my-config" {
  provider      = google-beta
  location      = "us-central1"
  vpcsc_policy   = "ALLOW"
}
```

## Argument Reference

The following arguments are supported:



* `vpcsc_policy` -
  (Optional)
  The VPC SC policy for project and location.
  Possible values are: `DENY`, `ALLOW`.

* `location` -
  (Optional)
  The name of the location this config is located in.

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.



## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/{{location}}/vpcscConfig`

* `name` -
  The name of the project's VPC SC Config.
  Always of the form: projects/{project}/location/{location}/vpcscConfig


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


VPCSCConfig can be imported using any of these accepted formats:

* `projects/{{project}}/locations/{{location}}/vpcscConfig/{{name}}`
* `{{project}}/{{location}}/{{name}}`
* `{{location}}/{{name}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import VPCSCConfig using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/locations/{{location}}/vpcscConfig/{{name}}"
  to = google_artifact_registry_vpcsc_config.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), VPCSCConfig can be imported using one of the formats above. For example:

```
$ terraform import google_artifact_registry_vpcsc_config.default projects/{{project}}/locations/{{location}}/vpcscConfig/{{name}}
$ terraform import google_artifact_registry_vpcsc_config.default {{project}}/{{location}}/{{name}}
$ terraform import google_artifact_registry_vpcsc_config.default {{location}}/{{name}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
