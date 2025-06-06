---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: Handwritten     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/website/docs/r/storage_managed_folder_iam.html.markdown
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Cloud Storage"
description: |-
  Collection of resources to manage IAM policy for Cloud Storage ManagedFolder
---

# IAM policy for Cloud Storage ManagedFolder
Three different resources help you manage your IAM policy for Cloud Storage ManagedFolder. Each of these resources serves a different use case:

* `google_storage_managed_folder_iam_policy`: Authoritative. Sets the IAM policy for the managedfolder and replaces any existing policy already attached.
* `google_storage_managed_folder_iam_binding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the managedfolder are preserved.
* `google_storage_managed_folder_iam_member`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the managedfolder are preserved.

A data source can be used to retrieve policy data in advent you do not need creation

* `google_storage_managed_folder_iam_policy`: Retrieves the IAM policy for the managedfolder

~> **Note:** `google_storage_managed_folder_iam_policy` **cannot** be used in conjunction with `google_storage_managed_folder_iam_binding` and `google_storage_managed_folder_iam_member` or they will fight over what your policy should be.

~> **Note:** `google_storage_managed_folder_iam_binding` resources **can be** used in conjunction with `google_storage_managed_folder_iam_member` resources **only if** they do not grant privilege to the same role.

~> **Note:**  This resource supports IAM Conditions but they have some known limitations which can be found [here](https://cloud.google.com/iam/docs/conditions-overview#limitations). Please review this article if you are having issues with IAM Conditions.


## google_storage_managed_folder_iam_policy

```hcl
data "google_iam_policy" "admin" {
  binding {
    role = "roles/storage.admin"
    members = [
      "user:jane@example.com",
    ]
  }
}

resource "google_storage_managed_folder_iam_policy" "policy" {
  bucket         = google_storage_managed_folder.folder.bucket
  managed_folder = google_storage_managed_folder.folder.name
  policy_data    = data.google_iam_policy.admin.policy_data
}
```

With IAM Conditions:

```hcl
data "google_iam_policy" "admin" {
  binding {
    role = "roles/storage.admin"
    members = [
      "user:jane@example.com",
    ]

    condition {
      title       = "expires_after_2019_12_31"
      description = "Expiring at midnight of 2019-12-31"
      expression  = "request.time < timestamp(\"2020-01-01T00:00:00Z\")"
    }
  }
}

resource "google_storage_managed_folder_iam_policy" "policy" {
  bucket         = google_storage_managed_folder.folder.bucket
  managed_folder = google_storage_managed_folder.folder.name
  policy_data    = data.google_iam_policy.admin.policy_data
}
```
## google_storage_managed_folder_iam_binding

```hcl
resource "google_storage_managed_folder_iam_binding" "binding" {
  bucket         = google_storage_managed_folder.folder.bucket
  managed_folder = google_storage_managed_folder.folder.name
  role           = "roles/storage.admin"
  members = [
    "user:jane@example.com",
  ]
}
```

With IAM Conditions:

```hcl
resource "google_storage_managed_folder_iam_binding" "binding" {
  bucket         = google_storage_managed_folder.folder.bucket
  managed_folder = google_storage_managed_folder.folder.name
  role           = "roles/storage.admin"
  members = [
    "user:jane@example.com",
  ]

  condition {
    title       = "expires_after_2019_12_31"
    description = "Expiring at midnight of 2019-12-31"
    expression  = "request.time < timestamp(\"2020-01-01T00:00:00Z\")"
  }
}
```
## google_storage_managed_folder_iam_member

```hcl
resource "google_storage_managed_folder_iam_member" "member" {
  bucket         = google_storage_managed_folder.folder.bucket
  managed_folder = google_storage_managed_folder.folder.name
  role           = "roles/storage.admin"
  member         = "user:jane@example.com"
}
```

With IAM Conditions:

```hcl
resource "google_storage_managed_folder_iam_member" "member" {
  bucket         = google_storage_managed_folder.folder.bucket
  managed_folder = google_storage_managed_folder.folder.name
  role           = "roles/storage.admin"
  member         = "user:jane@example.com"

  condition {
    title       = "expires_after_2019_12_31"
    description = "Expiring at midnight of 2019-12-31"
    expression  = "request.time < timestamp(\"2020-01-01T00:00:00Z\")"
  }
}
```

## Argument Reference

The following arguments are supported:

* `bucket` - (Required) The name of the bucket that contains the managed folder. Used to find the parent resource to bind the IAM policy to
* `managed_folder` - (Required) Used to find the parent resource to bind the IAM policy to

* `member/members` - (Required) Identities that will be granted the privilege in `role`.
  Each entry can have one of the following values:
  * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
  * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
  * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
  * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
  * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
  * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
  * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
  * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
  * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"

* `role` - (Required) The role that should be applied. Only one
    `google_storage_managed_folder_iam_binding` can be used per role. Note that custom roles must be of the format
    `[projects|organizations]/{parent-name}/roles/{role-name}`.

* `policy_data` - (Required only by `google_storage_managed_folder_iam_policy`) The policy data generated by
  a `google_iam_policy` data source.

* `condition` - (Optional) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
  Structure is documented below.

---

The `condition` block supports:

* `expression` - (Required) Textual representation of an expression in Common Expression Language syntax.

* `title` - (Required) A title for the expression, i.e. a short string describing its purpose.

* `description` - (Optional) An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.

~> **Warning:** Terraform considers the `role` and condition contents (`title`+`description`+`expression`) as the
  identifier for the binding. This means that if any part of the condition is changed out-of-band, Terraform will
  consider it to be an entirely different resource and will treat it as such.
## Attributes Reference

In addition to the arguments listed above, the following computed attributes are
exported:

* `etag` - (Computed) The etag of the IAM policy.

## Import

For all import syntaxes, the "resource in question" can take any of the following forms:

* b/{{bucket}}/managedFolders/{{managed_folder}}
* {{bucket}}/{{managed_folder}}

Any variables not passed in the import command will be taken from the provider configuration.

Cloud Storage managedfolder IAM resources can be imported using the resource identifiers, role, and member.

IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
```
$ terraform import google_storage_managed_folder_iam_member.editor "b/{{bucket}}/managedFolders/{{managed_folder}} roles/storage.objectViewer user:jane@example.com"
```

IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
```
$ terraform import google_storage_managed_folder_iam_binding.editor "b/{{bucket}}/managedFolders/{{managed_folder}} roles/storage.objectViewer"
```

IAM policy imports use the identifier of the resource in question, e.g.
```
$ terraform import google_storage_managed_folder_iam_policy.editor b/{{bucket}}/managedFolders/{{managed_folder}}
```

-> **Custom Roles** If you're importing a IAM resource with a custom role, make sure to use the
 full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
