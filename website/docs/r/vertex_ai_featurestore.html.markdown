---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/vertexai/Featurestore.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Vertex AI"
description: |-
  A collection of DataItems and Annotations on them.
---

# google_vertex_ai_featurestore

A collection of DataItems and Annotations on them.


To get more information about Featurestore, see:

* [API documentation](https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featurestores)
* How-to Guides
    * [Official Documentation](https://cloud.google.com/vertex-ai/docs)

## Example Usage - Vertex Ai Featurestore


```hcl
resource "google_vertex_ai_featurestore" "featurestore" {
  name     = "terraform"
  labels = {
    foo = "bar"
  }
  region   = "us-central1"
  online_serving_config {
    fixed_node_count = 2
  }
  encryption_spec {
    kms_key_name = "kms-name"
  }
  force_destroy = true
}
```
## Example Usage - Vertex Ai Featurestore With Beta Fields


```hcl
resource "google_vertex_ai_featurestore" "featurestore" {
  provider = google-beta
  name     = "terraform2"
  labels = {
    foo = "bar"
  }
  region   = "us-central1"
  online_serving_config {
    fixed_node_count = 2
  }
  encryption_spec {
    kms_key_name = "kms-name"
  }
  online_storage_ttl_days = 30
  force_destroy = true
}
```
## Example Usage - Vertex Ai Featurestore Scaling


```hcl
resource "google_vertex_ai_featurestore" "featurestore" {
  name     = "terraform3"
  labels = {
    foo = "bar"
  }
  region   = "us-central1"
  online_serving_config {
    scaling {
      min_node_count = 2
      max_node_count = 10
    }
  }
  encryption_spec {
    kms_key_name = "kms-name"
  }
  force_destroy = true
}
```

## Argument Reference

The following arguments are supported:



* `name` -
  (Optional)
  The name of the Featurestore. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.

* `labels` -
  (Optional)
  A set of key/value label pairs to assign to this Featurestore.

  **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
  Please refer to the field `effective_labels` for all of the labels present on the resource.

* `online_serving_config` -
  (Optional)
  Config for online serving resources.
  Structure is [documented below](#nested_online_serving_config).

* `online_storage_ttl_days` -
  (Optional, [Beta](https://terraform.io/docs/providers/google/guides/provider_versions.html))
  TTL in days for feature values that will be stored in online serving storage. The Feature Store online storage periodically removes obsolete feature values older than onlineStorageTtlDays since the feature generation time. Note that onlineStorageTtlDays should be less than or equal to offlineStorageTtlDays for each EntityType under a featurestore. If not set, default to 4000 days

* `encryption_spec` -
  (Optional)
  If set, both of the online and offline data storage will be secured by this key.
  Structure is [documented below](#nested_encryption_spec).

* `region` -
  (Optional)
  The region of the dataset. eg us-central1

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.

* `force_destroy` - (Optional) If set to true, any EntityTypes and Features for this Featurestore will also be deleted


<a name="nested_online_serving_config"></a>The `online_serving_config` block supports:

* `fixed_node_count` -
  (Optional)
  The number of nodes for each cluster. The number of nodes will not scale automatically but can be scaled manually by providing different values when updating.

* `scaling` -
  (Optional)
  Online serving scaling configuration. Only one of fixedNodeCount and scaling can be set. Setting one will reset the other.
  Structure is [documented below](#nested_online_serving_config_scaling).


<a name="nested_online_serving_config_scaling"></a>The `scaling` block supports:

* `min_node_count` -
  (Required)
  The minimum number of nodes to scale down to. Must be greater than or equal to 1.

* `max_node_count` -
  (Required)
  The maximum number of nodes to scale up to. Must be greater than minNodeCount, and less than or equal to 10 times of 'minNodeCount'.

<a name="nested_encryption_spec"></a>The `encryption_spec` block supports:

* `kms_key_name` -
  (Required)
  The Cloud KMS resource identifier of the customer managed encryption key used to protect a resource. Has the form: projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key. The key needs to be in the same region as where the compute resource is created.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/{{region}}/featurestores/{{name}}`

* `etag` -
  Used to perform consistent read-modify-write updates.

* `create_time` -
  The timestamp of when the featurestore was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.

* `update_time` -
  The timestamp of when the featurestore was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.

* `terraform_labels` -
  The combination of labels configured directly on the resource
   and default labels configured on the provider.

* `effective_labels` -
  All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


Featurestore can be imported using any of these accepted formats:

* `projects/{{project}}/locations/{{region}}/featurestores/{{name}}`
* `{{project}}/{{region}}/{{name}}`
* `{{region}}/{{name}}`
* `{{name}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Featurestore using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/locations/{{region}}/featurestores/{{name}}"
  to = google_vertex_ai_featurestore.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), Featurestore can be imported using one of the formats above. For example:

```
$ terraform import google_vertex_ai_featurestore.default projects/{{project}}/locations/{{region}}/featurestores/{{name}}
$ terraform import google_vertex_ai_featurestore.default {{project}}/{{region}}/{{name}}
$ terraform import google_vertex_ai_featurestore.default {{region}}/{{name}}
$ terraform import google_vertex_ai_featurestore.default {{name}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
