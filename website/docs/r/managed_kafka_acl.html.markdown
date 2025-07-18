---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/managedkafka/Acl.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Managed Kafka"
description: |-
  A Managed Service for Apache Kafka ACL.
---

# google_managed_kafka_acl

A Managed Service for Apache Kafka ACL. Apache Kafka is a trademark owned by the Apache Software Foundation.



<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=managedkafka_acl_basic&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Managedkafka Acl Basic


```hcl
resource "google_managed_kafka_cluster" "cluster" {
  cluster_id = "my-cluster"
  location = "us-central1"
  capacity_config {
    vcpu_count = 3
    memory_bytes = 3221225472
  }
  gcp_config {
    access_config {
      network_configs {
        subnet = "projects/${data.google_project.project.number}/regions/us-central1/subnetworks/default"
      }
    }
  }
}

resource "google_managed_kafka_acl" "example" {
  acl_id = "topic/mytopic"
  cluster = google_managed_kafka_cluster.cluster.cluster_id
  location = "us-central1"
    acl_entries {
      principal = "User:admin@my-project.iam.gserviceaccount.com"
      permission_type = "ALLOW"
      operation = "ALL"
      host = "*"
    }
    acl_entries {
      principal = "User:producer-client@my-project.iam.gserviceaccount.com"
      permission_type = "ALLOW"
      operation = "WRITE"
      host = "*"
    }
}

data "google_project" "project" {
}
```

## Argument Reference

The following arguments are supported:


* `acl_entries` -
  (Required)
  The acl entries that apply to the resource pattern. The maximum number of allowed entries is 100.
  Structure is [documented below](#nested_acl_entries).

* `location` -
  (Required)
  ID of the location of the Kafka resource. See https://cloud.google.com/managed-kafka/docs/locations for a list of supported locations.

* `cluster` -
  (Required)
  The cluster name.

* `acl_id` -
  (Required)
  The ID to use for the acl, which will become the final component of the acl's name. The structure of `aclId` defines the Resource Pattern (resource_type, resource_name, pattern_type) of the acl. `aclId` is structured like one of the following:
  For acls on the cluster: `cluster`
  For acls on a single resource within the cluster: `topic/{resource_name}` `consumerGroup/{resource_name}` `transactionalId/{resource_name}`
  For acls on all resources that match a prefix: `topicPrefixed/{resource_name}` `consumerGroupPrefixed/{resource_name}` `transactionalIdPrefixed/{resource_name}`
  For acls on all resources of a given type (i.e. the wildcard literal '*''): `allTopics` (represents `topic/*`) `allConsumerGroups` (represents `consumerGroup/*`) `allTransactionalIds` (represents `transactionalId/*`).


* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.



<a name="nested_acl_entries"></a>The `acl_entries` block supports:

* `principal` -
  (Required)
  The principal. Specified as Google Cloud account, with the Kafka StandardAuthorizer prefix User:". For example: "User:test-kafka-client@test-project.iam.gserviceaccount.com". Can be the wildcard "User:*" to refer to all users.

* `permission_type` -
  (Optional)
  The permission type. Accepted values are (case insensitive): ALLOW, DENY.

* `operation` -
  (Required)
  The operation type. Allowed values are (case insensitive): ALL, READ,
  WRITE, CREATE, DELETE, ALTER, DESCRIBE, CLUSTER_ACTION, DESCRIBE_CONFIGS,
  ALTER_CONFIGS, and IDEMPOTENT_WRITE. See https://kafka.apache.org/documentation/#operations_resources_and_protocols
  for valid combinations of resource_type and operation for different Kafka API requests.

* `host` -
  (Optional)
  The host. Must be set to "*" for Managed Service for Apache Kafka.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/{{location}}/clusters/{{cluster}}/acls/{{acl_id}}`

* `name` -
  The name of the acl. The `ACL_ID` segment is used when connecting directly to the cluster. Must be in the format `projects/PROJECT_ID/locations/LOCATION/clusters/CLUSTER_ID/acls/ACL_ID`.

* `etag` -
  `etag` is used for concurrency control. An `etag` is returned in the
  response to `GetAcl` and `CreateAcl`. Callers are required to put that etag
  in the request to `UpdateAcl` to ensure that their change will be applied
  to the same version of the acl that exists in the Kafka Cluster.
  A terminal 'T' character in the etag indicates that the AclEntries were
  truncated due to repeated field limits.

* `resource_type` -
  The acl resource type derived from the name. One of: CLUSTER, TOPIC, GROUP, TRANSACTIONAL_ID.

* `resource_name` -
  The acl resource name derived from the name. For cluster resource_type, this is always "kafka-cluster". Can be the wildcard literal "*".

* `pattern_type` -
  The acl pattern type derived from the name. One of: LITERAL, PREFIXED.


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


Acl can be imported using any of these accepted formats:

* `projects/{{project}}/locations/{{location}}/clusters/{{cluster}}/acls/{{acl_id}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Acl using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/locations/{{location}}/clusters/{{cluster}}/acls/{{acl_id}}"
  to = google_managed_kafka_acl.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), Acl can be imported using one of the formats above. For example:

```
$ terraform import google_managed_kafka_acl.default projects/{{project}}/locations/{{location}}/clusters/{{cluster}}/acls/{{acl_id}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
