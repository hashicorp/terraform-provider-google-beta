---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/accesscontextmanager/IngressPolicy.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Access Context Manager (VPC Service Controls)"
description: |-
  This resource has been deprecated, please refer to ServicePerimeterIngressPolicy.
---

# google_access_context_manager_ingress_policy

This resource has been deprecated, please refer to ServicePerimeterIngressPolicy.


To get more information about IngressPolicy, see:

* [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies.servicePerimeters#ingresspolicy)

## Argument Reference

The following arguments are supported:


* `resource` -
  (Required)
  A GCP resource that is inside of the service perimeter.

* `ingress_policy_name` -
  (Required)
  The name of the Service Perimeter to add this resource to.




## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `{{ingress_policy_name}}/{{resource}}`

* `access_policy_id` -
  The name of the Access Policy this resource belongs to.


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


IngressPolicy can be imported using any of these accepted formats:

* `{{ingress_policy_name}}/{{resource}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import IngressPolicy using one of the formats above. For example:

```tf
import {
  id = "{{ingress_policy_name}}/{{resource}}"
  to = google_access_context_manager_ingress_policy.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), IngressPolicy can be imported using one of the formats above. For example:

```
$ terraform import google_access_context_manager_ingress_policy.default {{ingress_policy_name}}/{{resource}}
```
