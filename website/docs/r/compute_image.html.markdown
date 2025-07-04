---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/compute/Image.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Compute Engine"
description: |-
  Represents an Image resource.
---

# google_compute_image

Represents an Image resource.

Google Compute Engine uses operating system images to create the root
persistent disks for your instances. You specify an image when you create
an instance. Images contain a boot loader, an operating system, and a
root file system. Linux operating system images are also capable of
running containers on Compute Engine.

Images can be either public or custom.

Public images are provided and maintained by Google, open-source
communities, and third-party vendors. By default, all projects have
access to these images and can use them to create instances.  Custom
images are available only to your project. You can create a custom image
from root persistent disks and other images. Then, use the custom image
to create an instance.


To get more information about Image, see:

* [API documentation](https://cloud.google.com/compute/docs/reference/v1/images)
* How-to Guides
    * [Official Documentation](https://cloud.google.com/compute/docs/images)

~> **Warning:** All arguments including the following potentially sensitive
values will be stored in the raw state as plain text: `image_encryption_key.raw_key`, `image_encryption_key.rsa_encrypted_key`, `source_disk_encryption_key.raw_key`, `source_disk_encryption_key.rsa_encrypted_key`, `source_image_encryption_key.raw_key`, `source_image_encryption_key.rsa_encrypted_key`, `source_snapshot_encryption_key.raw_key`, `source_snapshot_encryption_key.rsa_encrypted_key`.
[Read more about sensitive data in state](https://www.terraform.io/language/state/sensitive-data).

<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=image_basic&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Image Basic


```hcl
data "google_compute_image" "debian" {
  family  = "debian-12"
  project = "debian-cloud"
}

resource "google_compute_disk" "persistent" {
  name  = "example-disk"
  image = data.google_compute_image.debian.self_link
  size  = 10
  type  = "pd-ssd"
  zone  = "us-central1-a"
}

resource "google_compute_image" "example" {
  name = "example-image"

  source_disk = google_compute_disk.persistent.id
}
```
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=image_guest_os&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Image Guest Os


```hcl
data "google_compute_image" "debian" {
  family  = "debian-12"
  project = "debian-cloud"
}

resource "google_compute_disk" "persistent" {
  name  = "example-disk"
  image = data.google_compute_image.debian.self_link
  size  = 10
  type  = "pd-ssd"
  zone  = "us-central1-a"
}

resource "google_compute_image" "example" {
  name = "example-image"

  source_disk = google_compute_disk.persistent.id

  guest_os_features {
    type = "UEFI_COMPATIBLE"
  }

  guest_os_features {
    type = "VIRTIO_SCSI_MULTIQUEUE"
  }

  guest_os_features {
    type = "GVNIC"
  }

  guest_os_features {
    type = "SEV_CAPABLE"
  }

  guest_os_features {
    type = "SEV_LIVE_MIGRATABLE_V2"
  }
}
```
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=image_basic_storage_location&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Image Basic Storage Location


```hcl
data "google_compute_image" "debian" {
  family  = "debian-12"
  project = "debian-cloud"
}

resource "google_compute_disk" "persistent" {
  name  = "example-disk"
  image = data.google_compute_image.debian.self_link
  size  = 10
  type  = "pd-ssd"
  zone  = "us-central1-a"
}

resource "google_compute_image" "example" {
  name = "example-sl-image"

  source_disk = google_compute_disk.persistent.id
  storage_locations = ["us-central1"]
}
```

## Argument Reference

The following arguments are supported:


* `name` -
  (Required)
  Name of the resource; provided by the client when the resource is
  created. The name must be 1-63 characters long, and comply with
  RFC1035. Specifically, the name must be 1-63 characters long and
  match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means
  the first character must be a lowercase letter, and all following
  characters must be a dash, lowercase letter, or digit, except the
  last character, which cannot be a dash.


* `description` -
  (Optional)
  An optional description of this resource. Provide this property when
  you create the resource.

* `storage_locations` -
  (Optional)
  Cloud Storage bucket storage location of the image
  (regional or multi-regional).
  Reference link: https://cloud.google.com/compute/docs/reference/rest/v1/images

* `disk_size_gb` -
  (Optional)
  Size of the image when restored onto a persistent disk (in GB).

* `family` -
  (Optional)
  The name of the image family to which this image belongs. You can
  create disks by specifying an image family instead of a specific
  image name. The image family always returns its latest image that is
  not deprecated. The name of the image family must comply with
  RFC1035.

* `guest_os_features` -
  (Optional)
  A list of features to enable on the guest operating system.
  Applicable only for bootable images.
  Structure is [documented below](#nested_guest_os_features).

* `image_encryption_key` -
  (Optional)
  Encrypts the image using a customer-supplied encryption key.
  After you encrypt an image with a customer-supplied key, you must
  provide the same key if you use the image later (e.g. to create a
  disk from the image)
  Structure is [documented below](#nested_image_encryption_key).

* `labels` -
  (Optional)
  Labels to apply to this Image.
  **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
  Please refer to the field `effective_labels` for all of the labels present on the resource.

* `licenses` -
  (Optional)
  Any applicable license URI.

* `raw_disk` -
  (Optional)
  The parameters of the raw disk image.
  Structure is [documented below](#nested_raw_disk).

* `source_disk` -
  (Optional)
  The source disk to create this image based on.
  You must provide either this property or the
  rawDisk.source property but not both to create an image.

* `source_disk_encryption_key` -
  (Optional)
  The customer-supplied encryption key of the source disk. Required if
  the source disk is protected by a customer-supplied encryption key.
  Structure is [documented below](#nested_source_disk_encryption_key).

* `source_image` -
  (Optional)
  URL of the source image used to create this image. In order to create an image, you must provide the full or partial
  URL of one of the following:
  * The selfLink URL
  * This property
  * The rawDisk.source URL
  * The sourceDisk URL

* `source_image_encryption_key` -
  (Optional)
  The customer-supplied encryption key of the source image. Required if
  the source image is protected by a customer-supplied encryption key.
  Structure is [documented below](#nested_source_image_encryption_key).

* `source_snapshot` -
  (Optional)
  URL of the source snapshot used to create this image.
  In order to create an image, you must provide the full or partial URL of one of the following:
  * The selfLink URL
  * This property
  * The sourceImage URL
  * The rawDisk.source URL
  * The sourceDisk URL

* `shielded_instance_initial_state` -
  (Optional)
  Set the secure boot keys of shielded instance.
  Structure is [documented below](#nested_shielded_instance_initial_state).

* `source_snapshot_encryption_key` -
  (Optional)
  The customer-supplied encryption key of the source snapshot. Required if
  the source snapshot is protected by a customer-supplied encryption key.
  Structure is [documented below](#nested_source_snapshot_encryption_key).

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.



<a name="nested_guest_os_features"></a>The `guest_os_features` block supports:

* `type` -
  (Required)
  The type of supported feature. Read [Enabling guest operating system features](https://cloud.google.com/compute/docs/images/create-delete-deprecate-private-images#guest-os-features) to see a list of available options.
  Possible values are: `MULTI_IP_SUBNET`, `SECURE_BOOT`, `SEV_CAPABLE`, `UEFI_COMPATIBLE`, `VIRTIO_SCSI_MULTIQUEUE`, `WINDOWS`, `GVNIC`, `IDPF`, `SEV_LIVE_MIGRATABLE`, `SEV_SNP_CAPABLE`, `SUSPEND_RESUME_COMPATIBLE`, `TDX_CAPABLE`, `SEV_LIVE_MIGRATABLE_V2`.

<a name="nested_image_encryption_key"></a>The `image_encryption_key` block supports:

* `kms_key_self_link` -
  (Optional)
  The self link of the encryption key that is stored in Google Cloud
  KMS.

* `kms_key_service_account` -
  (Optional)
  The service account being used for the encryption request for the
  given KMS key. If absent, the Compute Engine default service
  account is used.

* `raw_key` -
  (Optional)
  Specifies a 256-bit customer-supplied encryption key, encoded in
  RFC 4648 base64 to either encrypt or decrypt this resource.
  **Note**: This property is sensitive and will not be displayed in the plan.

* `rsa_encrypted_key` -
  (Optional)
  Specifies a 256-bit customer-supplied encryption key, encoded in
  RFC 4648 base64 to either encrypt or decrypt this resource.
  **Note**: This property is sensitive and will not be displayed in the plan.

<a name="nested_raw_disk"></a>The `raw_disk` block supports:

* `container_type` -
  (Optional)
  The format used to encode and transmit the block device, which
  should be TAR. This is just a container and transmission format
  and not a runtime format. Provided by the client when the disk
  image is created.
  Default value is `TAR`.
  Possible values are: `TAR`.

* `sha1` -
  (Optional)
  An optional SHA1 checksum of the disk image before unpackaging.
  This is provided by the client when the disk image is created.

* `source` -
  (Required)
  The full Google Cloud Storage URL where disk storage is stored
  You must provide either this property or the sourceDisk property
  but not both.

<a name="nested_source_disk_encryption_key"></a>The `source_disk_encryption_key` block supports:

* `raw_key` -
  (Optional)
  Specifies a 256-bit customer-supplied encryption key, encoded in
  RFC 4648 base64 to either encrypt or decrypt this resource.
  **Note**: This property is sensitive and will not be displayed in the plan.

* `rsa_encrypted_key` -
  (Optional)
  Specifies an RFC 4648 base64 encoded, RSA-wrapped 2048-bit
  customer-supplied encryption key to either encrypt or decrypt
  this resource. You can provide either the rawKey or the rsaEncryptedKey.
  **Note**: This property is sensitive and will not be displayed in the plan.

* `kms_key_self_link` -
  (Optional)
  The self link of the encryption key used to decrypt this resource. Also called KmsKeyName
  in the cloud console. Your project's Compute Engine System service account
  (`service-{{PROJECT_NUMBER}}@compute-system.iam.gserviceaccount.com`) must have
  `roles/cloudkms.cryptoKeyEncrypterDecrypter` to use this feature.
  See https://cloud.google.com/compute/docs/disks/customer-managed-encryption#encrypt_a_new_persistent_disk_with_your_own_keys

* `kms_key_service_account` -
  (Optional)
  The service account being used for the encryption request for the
  given KMS key. If absent, the Compute Engine default service
  account is used.

<a name="nested_source_image_encryption_key"></a>The `source_image_encryption_key` block supports:

* `raw_key` -
  (Optional)
  Specifies a 256-bit customer-supplied encryption key, encoded in
  RFC 4648 base64 to either encrypt or decrypt this resource.
  **Note**: This property is sensitive and will not be displayed in the plan.

* `rsa_encrypted_key` -
  (Optional)
  Specifies an RFC 4648 base64 encoded, RSA-wrapped 2048-bit
  customer-supplied encryption key to either encrypt or decrypt
  this resource. You can provide either the rawKey or the rsaEncryptedKey.
  **Note**: This property is sensitive and will not be displayed in the plan.

* `kms_key_self_link` -
  (Optional)
  The self link of the encryption key used to decrypt this resource. Also called KmsKeyName
  in the cloud console. Your project's Compute Engine System service account
  (`service-{{PROJECT_NUMBER}}@compute-system.iam.gserviceaccount.com`) must have
  `roles/cloudkms.cryptoKeyEncrypterDecrypter` to use this feature.
  See https://cloud.google.com/compute/docs/disks/customer-managed-encryption#encrypt_a_new_persistent_disk_with_your_own_keys

* `kms_key_service_account` -
  (Optional)
  The service account being used for the encryption request for the
  given KMS key. If absent, the Compute Engine default service
  account is used.

<a name="nested_shielded_instance_initial_state"></a>The `shielded_instance_initial_state` block supports:

* `pk` -
  (Optional)
  The Platform Key (PK).
  Structure is [documented below](#nested_shielded_instance_initial_state_pk).

* `keks` -
  (Optional)
  The Key Exchange Key (KEK).
  Structure is [documented below](#nested_shielded_instance_initial_state_keks).

* `dbs` -
  (Optional)
  The Key Database (db).
  Structure is [documented below](#nested_shielded_instance_initial_state_dbs).

* `dbxs` -
  (Optional)
  The forbidden key database (dbx).
  Structure is [documented below](#nested_shielded_instance_initial_state_dbxs).


<a name="nested_shielded_instance_initial_state_pk"></a>The `pk` block supports:

* `content` -
  (Required)
  The raw content in the secure keys file.
  A base64-encoded string.

* `file_type` -
  (Optional)
  The file type of source file.

<a name="nested_shielded_instance_initial_state_keks"></a>The `keks` block supports:

* `content` -
  (Required)
  The raw content in the secure keys file.
  A base64-encoded string.

* `file_type` -
  (Optional)
  The file type of source file.

<a name="nested_shielded_instance_initial_state_dbs"></a>The `dbs` block supports:

* `content` -
  (Required)
  The raw content in the secure keys file.
  A base64-encoded string.

* `file_type` -
  (Optional)
  The file type of source file.

<a name="nested_shielded_instance_initial_state_dbxs"></a>The `dbxs` block supports:

* `content` -
  (Required)
  The raw content in the secure keys file.
  A base64-encoded string.

* `file_type` -
  (Optional)
  The file type of source file.

<a name="nested_source_snapshot_encryption_key"></a>The `source_snapshot_encryption_key` block supports:

* `raw_key` -
  (Optional)
  Specifies a 256-bit customer-supplied encryption key, encoded in
  RFC 4648 base64 to either encrypt or decrypt this resource.
  **Note**: This property is sensitive and will not be displayed in the plan.

* `rsa_encrypted_key` -
  (Optional)
  Specifies an RFC 4648 base64 encoded, RSA-wrapped 2048-bit
  customer-supplied encryption key to either encrypt or decrypt
  this resource. You can provide either the rawKey or the rsaEncryptedKey.
  **Note**: This property is sensitive and will not be displayed in the plan.

* `kms_key_self_link` -
  (Optional)
  The self link of the encryption key used to decrypt this resource. Also called KmsKeyName
  in the cloud console. Your project's Compute Engine System service account
  (`service-{{PROJECT_NUMBER}}@compute-system.iam.gserviceaccount.com`) must have
  `roles/cloudkms.cryptoKeyEncrypterDecrypter` to use this feature.
  See https://cloud.google.com/compute/docs/disks/customer-managed-encryption#encrypt_a_new_persistent_disk_with_your_own_keys

* `kms_key_service_account` -
  (Optional)
  The service account being used for the encryption request for the
  given KMS key. If absent, the Compute Engine default service
  account is used.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/global/images/{{name}}`

* `archive_size_bytes` -
  Size of the image tar.gz archive stored in Google Cloud Storage (in
  bytes).

* `creation_timestamp` -
  Creation timestamp in RFC3339 text format.

* `label_fingerprint` -
  The fingerprint used for optimistic locking of this resource. Used
  internally during updates.

* `terraform_labels` -
  The combination of labels configured directly on the resource
   and default labels configured on the provider.

* `effective_labels` -
  All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.
* `self_link` - The URI of the created resource.


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


Image can be imported using any of these accepted formats:

* `projects/{{project}}/global/images/{{name}}`
* `{{project}}/{{name}}`
* `{{name}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Image using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/global/images/{{name}}"
  to = google_compute_image.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), Image can be imported using one of the formats above. For example:

```
$ terraform import google_compute_image.default projects/{{project}}/global/images/{{name}}
$ terraform import google_compute_image.default {{project}}/{{name}}
$ terraform import google_compute_image.default {{name}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
