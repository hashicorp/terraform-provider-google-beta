---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: Handwritten     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/website/docs/d/compute_router.html.markdown
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Compute Engine"
description: |-
  Get a Cloud Router within GCE.
---

# google_compute_router

Get a router within GCE from its name and VPC.

## Example Usage

```hcl
data "google_compute_router" "my-router" {
  name   = "myrouter-us-east1"
  network = "my-network"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the router.

* `network` - (Required) The VPC network on which this router lives.

* `project` - (Optional) The ID of the project in which the resource
    belongs. If it is not provided, the provider project is used.

* `region` - (Optional) The region this router has been created in. If
    unspecified, this defaults to the region configured in the provider.


## Attributes Reference

See [google_compute_router](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_router) resource for details of the available attributes.
