---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/iap/TunnelDestGroup.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Identity-Aware Proxy"
description: |-
  Tunnel destination groups represent resources that have the same tunnel access restrictions.
---

# google_iap_tunnel_dest_group

Tunnel destination groups represent resources that have the same tunnel access restrictions.


To get more information about TunnelDestGroup, see:

* [API documentation](https://cloud.google.com/iap/docs/reference/rest/v1/projects.iap_tunnel.locations.destGroups)
* How-to Guides
    * [Set up IAP TCP forwarding with an IP address or hostname in a Google Cloud or non-Google Cloud environment](https://cloud.google.com/iap/docs/tcp-by-host)

<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=iap_destgroup&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Iap Destgroup


```hcl
resource "google_iap_tunnel_dest_group" "dest_group" {
  region = "us-central1"
  group_name = "testgroup%{random_suffix}"
  cidrs = [
    "10.1.0.0/16",
    "192.168.10.0/24",
  ]
}
```

## Argument Reference

The following arguments are supported:


* `group_name` -
  (Required)
  Unique tunnel destination group name.


* `cidrs` -
  (Optional)
  List of CIDRs that this group applies to.

* `fqdns` -
  (Optional)
  List of FQDNs that this group applies to.

* `region` -
  (Optional)
  The region of the tunnel group. Must be the same as the network resources in the group.

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.



## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/iap_tunnel/locations/{{region}}/destGroups/{{group_name}}`

* `name` -
  Full resource name.


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


TunnelDestGroup can be imported using any of these accepted formats:

* `projects/{{project}}/iap_tunnel/locations/{{region}}/destGroups/{{group_name}}`
* `{{project}}/iap_tunnel/locations/{{region}}/destGroups/{{group_name}}`
* `{{project}}/{{region}}/{{group_name}}`
* `{{region}}/destGroups/{{group_name}}`
* `{{region}}/{{group_name}}`
* `{{group_name}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import TunnelDestGroup using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/iap_tunnel/locations/{{region}}/destGroups/{{group_name}}"
  to = google_iap_tunnel_dest_group.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), TunnelDestGroup can be imported using one of the formats above. For example:

```
$ terraform import google_iap_tunnel_dest_group.default projects/{{project}}/iap_tunnel/locations/{{region}}/destGroups/{{group_name}}
$ terraform import google_iap_tunnel_dest_group.default {{project}}/iap_tunnel/locations/{{region}}/destGroups/{{group_name}}
$ terraform import google_iap_tunnel_dest_group.default {{project}}/{{region}}/{{group_name}}
$ terraform import google_iap_tunnel_dest_group.default {{region}}/destGroups/{{group_name}}
$ terraform import google_iap_tunnel_dest_group.default {{region}}/{{group_name}}
$ terraform import google_iap_tunnel_dest_group.default {{group_name}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
