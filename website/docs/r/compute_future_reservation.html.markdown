---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/compute/FutureReservation.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Compute Engine"
description: |-
  Represents a future reservation resource in Compute Engine.
---

# google_compute_future_reservation

Represents a future reservation resource in Compute Engine. Future reservations allow users
to reserve capacity for a specified time window, ensuring that resources are available
when needed.

Reservations apply only to Compute Engine, Cloud Dataproc, and Google
Kubernetes Engine VM usage.Reservations do not apply to `f1-micro` or
`g1-small` machine types, preemptible VMs, sole tenant nodes, or other
services not listed above
like Cloud SQL and Dataflow.

~> **Warning:** This resource is in beta, and should be used with the terraform-provider-google-beta provider.
See [Provider Versions](https://terraform.io/docs/providers/google/guides/provider_versions.html) for more details on beta resources.

To get more information about FutureReservation, see:

* [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/futureReservations)
* How-to Guides
    * [Future Reservations Guide](https://cloud.google.com/compute/docs/instances/future-reservations-overview)

## Example Usage - Future Reservation Basic


```hcl
resource "google_compute_future_reservation" "gce_future_reservation" {
  provider = google-beta
  name     = "gce-future-reservation"
  project  = "my-project-name"
  auto_delete_auto_created_reservations = true
  planning_status = "DRAFT"
  name_prefix = "fr-basic"
  time_window {
    start_time = "2025-11-01T00:00:00Z"
    end_time   = "2025-11-02T00:00:00Z"
  }
  specific_sku_properties {
    total_count = "1"

    instance_properties {
      machine_type = "e2-standard-2"
    }
  }
}
```
## Example Usage - Future Reservation Aggregate Reservation


```hcl
resource "google_compute_future_reservation" "gce_future_reservation" {
  provider = google-beta
  name     = "gce-future-reservation-aggregate-reservation"
  project  = "my-project-name"
  auto_delete_auto_created_reservations = true
  planning_status = "DRAFT"
  name_prefix = "fr-basic"
  time_window {
    start_time = "2025-11-01T00:00:00Z"
    end_time   = "2025-11-02T00:00:00Z"
  }
  aggregate_reservation {
    vm_family = "VM_FAMILY_CLOUD_TPU_DEVICE_CT3"
    workload_type = "UNSPECIFIED"
    reserved_resources {
      accelerator {
        accelerator_count = 32
        accelerator_type  = "projects/my-project-name/zones/us-central1-a/acceleratorTypes/ct3"
      }
    }
    reserved_resources {
      accelerator {
        accelerator_count = 2
        accelerator_type  = "projects/my-project-name/zones/us-central1-a/acceleratorTypes/ct3"
      }
    }
  }
}
```

## Argument Reference

The following arguments are supported:


* `time_window` -
  (Required)
  Time window for this Future Reservation.
  Structure is [documented below](#nested_time_window).

* `name` -
  (Required)
  Name of the resource. Provided by the client when the resource is
  created. The name must be 1-63 characters long, and comply with
  RFC1035. Specifically, the name must be 1-63 characters long and match
  the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the
  first character must be a lowercase letter, and all following
  characters must be a dash, lowercase letter, or digit, except the las
  character, which cannot be a dash.


* `description` -
  (Optional)
  An optional description of this resource.

* `share_settings` -
  (Optional)
  Settings for sharing the future reservation
  Structure is [documented below](#nested_share_settings).

* `name_prefix` -
  (Optional)
  Name prefix for the reservations to be created at the time of delivery. The name prefix must comply with RFC1035. Maximum allowed length for name prefix is 20. Automatically created reservations name format will be -date-####.

* `planning_status` -
  (Optional)
  Planning state before being submitted for evaluation
  Possible values are: `DRAFT`, `SUBMITTED`.

* `auto_delete_auto_created_reservations` -
  (Optional)
  Setting for enabling or disabling automatic deletion for auto-created reservation. If set to true, auto-created reservations will be deleted at Future Reservation's end time (default) or at user's defined timestamp if any of the [autoCreatedReservationsDeleteTime, autoCreatedReservationsDuration] values is specified. For keeping auto-created reservation indefinitely, this value should be set to false.

* `specific_reservation_required` -
  (Optional)
  Indicates whether the auto-created reservation can be consumed by VMs with affinity for "any" reservation. If the field is set, then only VMs that target the reservation by name can consume from the delivered reservation.

* `reservation_name` -
  (Optional)
  Name of reservations where the capacity is provisioned at the time of delivery of future reservations. If the reservation with the given name does not exist already, it is created automatically at the time of Approval with INACTIVE state till specified start-time. Either provide the reservationName or a namePrefix.

* `deployment_type` -
  (Optional)
  Type of the deployment requested as part of future reservation.
  Possible values are: `DENSE`, `FLEXIBLE`.

* `reservation_mode` -
  (Optional)
  The reservation mode which determines reservation-termination behavior and expected pricing.
  Possible values are: `CALENDAR`, `DEFAULT`.

* `commitment_info` -
  (Optional)
  If not present, then FR will not deliver a new commitment or update an existing commitment.
  Structure is [documented below](#nested_commitment_info).

* `scheduling_type` -
  (Optional)
  Maintenance information for this reservation
  Possible values are: `GROUPED`, `INDEPENDENT`.

* `specific_sku_properties` -
  (Optional)
  Future Reservation configuration to indicate instance properties and total count.
  Structure is [documented below](#nested_specific_sku_properties).

* `auto_created_reservations_delete_time` -
  (Optional)
  Future timestamp when the FR auto-created reservations will be deleted by Compute Engine.

* `auto_created_reservations_duration` -
  (Optional)
  Specifies the duration of auto-created reservations. It represents relative time to future reservation startTime when auto-created reservations will be automatically deleted by Compute Engine. Duration time unit is represented as a count of seconds and fractions of seconds at nanosecond resolution.
  Structure is [documented below](#nested_auto_created_reservations_duration).

* `aggregate_reservation` -
  (Optional)
  Aggregate reservation details for the future reservation.
  Structure is [documented below](#nested_aggregate_reservation).

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.



<a name="nested_time_window"></a>The `time_window` block supports:

* `start_time` -
  (Required)
  Start time of the future reservation in RFC3339 format.

* `end_time` -
  (Optional)
  End time of the future reservation in RFC3339 format.

* `duration` -
  (Optional)
  Duration of the future reservation
  Structure is [documented below](#nested_time_window_duration).


<a name="nested_time_window_duration"></a>The `duration` block supports:

* `seconds` -
  (Optional)
  Span of time at a resolution of a second. Must be from 0 to 315,576,000,000 inclusive.

* `nanos` -
  (Optional)
  Span of time that's a fraction of a second at nanosecond resolution.

<a name="nested_share_settings"></a>The `share_settings` block supports:

* `share_type` -
  (Optional)
  Type of sharing for this future reservation.
  Possible values are: `LOCAL`, `SPECIFIC_PROJECTS`.

* `projects` -
  (Optional)
  list of Project names to specify consumer projects for this shared-reservation. This is only valid when shareType's value is SPECIFIC_PROJECTS.

* `project_map` -
  (Optional)
  A map of project id and project config. This is only valid when shareType's value is SPECIFIC_PROJECTS.
  Structure is [documented below](#nested_share_settings_project_map).


<a name="nested_share_settings_project_map"></a>The `project_map` block supports:

* `id` - (Required) The identifier for this object. Format specified above.

* `project_id` -
  (Optional)
  The project ID, should be same as the key of this project config in the parent map.

<a name="nested_commitment_info"></a>The `commitment_info` block supports:

* `commitment_plan` -
  (Optional)
  Indicates if a Commitment needs to be created as part of FR delivery. If this field is not present, then no commitment needs to be created.
  Possible values are: `INVALID`, `THIRTY_SIX_MONTH`, `TWELVE_MONTH`.

* `commitment_name` -
  (Optional)
  name of the commitment where capacity is being delivered to.

* `previous_commitment_terms` -
  (Optional)
  Only applicable if FR is delivering to the same reservation. If set, all parent commitments will be extended to match the end date of the plan for this commitment.
  Possible values are: `EXTEND`.

<a name="nested_specific_sku_properties"></a>The `specific_sku_properties` block supports:

* `instance_properties` -
  (Optional)
  Properties of the SKU instances being reserved.
  Structure is [documented below](#nested_specific_sku_properties_instance_properties).

* `total_count` -
  (Optional)
  Total number of instances for which capacity assurance is requested at a future time period.

* `source_instance_template` -
  (Optional)
  The instance template that will be used to populate the ReservedInstanceProperties of the future reservation


<a name="nested_specific_sku_properties_instance_properties"></a>The `instance_properties` block supports:

* `machine_type` -
  (Optional)
  Specifies type of machine (name only) which has fixed number of vCPUs and fixed amount of memory. This also includes specifying custom machine type following custom-NUMBER_OF_CPUS-AMOUNT_OF_MEMORY pattern.

* `guest_accelerators` -
  (Optional)
  Specifies accelerator type and count.
  Structure is [documented below](#nested_specific_sku_properties_instance_properties_guest_accelerators).

* `min_cpu_platform` -
  (Optional)
  Minimum cpu platform the reservation.

* `local_ssds` -
  (Optional)
  Specifies amount of local ssd to reserve with each instance. The type of disk is local-ssd.
  Structure is [documented below](#nested_specific_sku_properties_instance_properties_local_ssds).

* `maintenance_freeze_duration_hours` -
  (Optional)
  Specifies the number of hours after reservation creation where instances using the reservation won't be scheduled for maintenance.

* `location_hint` -
  (Optional)
  An opaque location hint used to place the allocation close to other resources. This field is for use by internal tools that use the public API.

* `maintenance_interval` -
  (Optional)
  Specifies the frequency of planned maintenance events. The accepted values are: PERIODIC
  Possible values are: `PERIODIC`.


<a name="nested_specific_sku_properties_instance_properties_guest_accelerators"></a>The `guest_accelerators` block supports:

* `accelerator_type` -
  (Optional)
  Full or partial URL of the accelerator type resource to attach to this instance.

* `accelerator_count` -
  (Optional)
  The number of the guest accelerator cards exposed to this instance.

<a name="nested_specific_sku_properties_instance_properties_local_ssds"></a>The `local_ssds` block supports:

* `disk_size_gb` -
  (Optional)
  Specifies the size of the disk in base-2 GB.

* `interface` -
  (Optional)
  Specifies the disk interface to use for attaching this disk, which is either SCSI or NVME. The default is SCSI.
  Possible values are: `SCSI`, `NVME`.

<a name="nested_auto_created_reservations_duration"></a>The `auto_created_reservations_duration` block supports:

* `seconds` -
  (Optional)
  Span of time at a resolution of a second. Must be from 0 to 315,576,000,000 inclusive.

* `nanos` -
  (Optional)
  Span of time that's a fraction of a second at nanosecond resolution. Durations less than one second are represented with a 0 seconds field and a positive nanos field. Must be from 0 to 999,999,999 inclusive.

<a name="nested_aggregate_reservation"></a>The `aggregate_reservation` block supports:

* `vm_family` -
  (Optional)
  The VM family that all instances scheduled against this reservation must belong to.
  Possible values are: `VM_FAMILY_CLOUD_TPU_DEVICE_CT3`, `VM_FAMILY_CLOUD_TPU_LITE_DEVICE_CT5L`, `VM_FAMILY_CLOUD_TPU_LITE_POD_SLICE_CT5LP`, `VM_FAMILY_CLOUD_TPU_LITE_POD_SLICE_CT6E`, `VM_FAMILY_CLOUD_TPU_POD_SLICE_CT3P`, `VM_FAMILY_CLOUD_TPU_POD_SLICE_CT4P`, `VM_FAMILY_CLOUD_TPU_POD_SLICE_CT5P`.

* `reserved_resources` -
  (Required)
  futureReservations.list of reserved resources (CPUs, memory, accelerators).
  Structure is [documented below](#nested_aggregate_reservation_reserved_resources).

* `workload_type` -
  (Optional)
  The workload type of the instances that will target this reservation.
  Possible values are: `BATCH`, `SERVING`, `UNSPECIFIED`.


<a name="nested_aggregate_reservation_reserved_resources"></a>The `reserved_resources` block supports:

* `accelerator` -
  (Optional)
  Properties of accelerator resources in this reservation.
  Structure is [documented below](#nested_aggregate_reservation_reserved_resources_reserved_resources_accelerator).


<a name="nested_aggregate_reservation_reserved_resources_reserved_resources_accelerator"></a>The `accelerator` block supports:

* `accelerator_count` -
  (Optional)
  Number of accelerators of specified type.

* `accelerator_type` -
  (Optional)
  Full or partial URL to accelerator type. e.g. "projects/{PROJECT}/zones/{ZONE}/acceleratorTypes/ct4l"

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/zones/{{zone}}/futureReservations/{{name}}`

* `zone` -
  URL of the Zone where this future reservation resides.

* `creation_timestamp` -
  The creation timestamp for this future reservation in RFC3339 text format.

* `self_link_with_id` -
  Server-defined URL for this resource with the resource id.

* `status` -
  [Output only] Status of the Future Reservation
  Structure is [documented below](#nested_status).
* `self_link` - The URI of the created resource.


<a name="nested_status"></a>The `status` block contains:

* `procurement_status` -
  (Optional)
  Current state of this Future Reservation
  Possible values are: `APPROVED`, `CANCELLED`, `COMMITTED`, `DECLINED`, `DRAFTING`, `FAILED`, `FAILED_PARTIALLY_FULFILLED`, `FULFILLED`, `PENDING_AMENDMENT_APPROVAL`, `PENDING_APPROVAL`, `PROCURING`, `PROVISIONING`.

* `lock_time` -
  (Optional)
  Time when Future Reservation would become LOCKED, after which no modifications to Future Reservation will be allowed. Applicable only after the Future Reservation is in the APPROVED state. The lockTime is an RFC3339 string. The procurementStatus will transition to PROCURING state at this time.

* `auto_created_reservations` -
  (Optional)
  Fully qualified urls of the automatically created reservations at startTime.

* `fulfilled_count` -
  (Optional)
  This count indicates the fulfilled capacity so far. This is set during "PROVISIONING" state. This count also includes capacity delivered as part of existing matching reservations.

* `specific_sku_properties` -
  (Optional)
  Instance properties related to the Future Reservation.
  Structure is [documented below](#nested_status_specific_sku_properties).

* `amendment_status` -
  (Optional)
  The current status of the requested amendment.
  Possible values are: .

* `last_known_good_state` -
  (Optional)
  This field represents the future reservation before an amendment was requested. If the amendment is declined, the Future Reservation will be reverted to the last known good state. The last known good state is not set when updating a future reservation whose Procurement Status is DRAFTING.
  Structure is [documented below](#nested_status_last_known_good_state).


<a name="nested_status_specific_sku_properties"></a>The `specific_sku_properties` block supports:

* `source_instance_template_id` -
  (Optional)
  ID of the instance template used to populate the Future Reservation properties.

<a name="nested_status_last_known_good_state"></a>The `last_known_good_state` block supports:

* `future_reservation_specs` -
  (Optional)
  The previous instance-related properties of the Future Reservation.
  Structure is [documented below](#nested_status_last_known_good_state_future_reservation_specs).

* `procurement_status` -
  (Optional)
  The status of the last known good state for the Future Reservation
  Possible values are: .

* `name_prefix` -
  (Optional)
  The name prefix of the Future Reservation before an amendment was requested.

* `description` -
  (Optional)
  The description of the FutureReservation before an amendment was requested.

* `lock_time` -
  (Optional)
  The lock time of the FutureReservation before an amendment was requested.

* `existing_matching_usage_info` -
  (Optional)
  Represents the matching usage for the future reservation before an amendment was requested.
  Structure is [documented below](#nested_status_last_known_good_state_existing_matching_usage_info).


<a name="nested_status_last_known_good_state_future_reservation_specs"></a>The `future_reservation_specs` block supports:

* `specific_sku_properties` -
  (Optional)
  The previous instance related properties of the Future Reservation.
  Structure is [documented below](#nested_status_last_known_good_state_future_reservation_specs_specific_sku_properties).

* `time_window` -
  (Optional)
  [Output Only] The previous time window of the Future Reservation.
  Structure is [documented below](#nested_status_last_known_good_state_future_reservation_specs_time_window).

* `share_settings` -
  (Optional)
  The previous share settings of the Future Reservation.
  Structure is [documented below](#nested_status_last_known_good_state_future_reservation_specs_share_settings).


<a name="nested_status_last_known_good_state_future_reservation_specs_specific_sku_properties"></a>The `specific_sku_properties` block supports:

* `instance_properties` -
  (Optional)
  Properties of the SKU instances being reserved.
  Structure is [documented below](#nested_status_last_known_good_state_future_reservation_specs_specific_sku_properties_instance_properties).

* `total_count` -
  (Optional)
  Total number of instances for which capacity assurance is requested at a future time period.

* `source_instance_template` -
  (Optional)
  The instance template that will be used to populate the ReservedInstanceProperties of the future reservation


<a name="nested_status_last_known_good_state_future_reservation_specs_specific_sku_properties_instance_properties"></a>The `instance_properties` block supports:

* `machine_type` -
  (Optional)
  Specifies type of machine (name only) which has fixed number of vCPUs and fixed amount of memory. This also includes specifying custom machine type following custom-NUMBER_OF_CPUS-AMOUNT_OF_MEMORY pattern.

* `guest_accelerators` -
  (Optional)
  Specifies accelerator type and count.
  Structure is [documented below](#nested_status_last_known_good_state_future_reservation_specs_specific_sku_properties_instance_properties_guest_accelerators).

* `min_cpu_platform` -
  (Optional)
  Minimum CPU platform for the reservation.

* `local_ssds` -
  (Optional)
  Specifies amount of local ssd to reserve with each instance. The type of disk is local-ssd.
  Structure is [documented below](#nested_status_last_known_good_state_future_reservation_specs_specific_sku_properties_instance_properties_local_ssds).

* `maintenance_freeze_duration_hours` -
  (Optional)
  Specifies the number of hours after reservation creation where instances using the reservation won't be scheduled for maintenance.

* `location_hint` -
  (Optional)
  An opaque location hint used to place the allocation close to other resources. This field is for use by internal tools that use the public API.

* `maintenance_interval` -
  (Optional)
  Specifies the frequency of planned maintenance events. The accepted values are: PERIODIC.
  Possible values are: `PERIODIC`.


<a name="nested_status_last_known_good_state_future_reservation_specs_specific_sku_properties_instance_properties_guest_accelerators"></a>The `guest_accelerators` block supports:

* `accelerator_type` -
  (Optional)
  Full or partial URL of the accelerator type resource to attach to this instance.

* `accelerator_count` -
  (Optional)
  The number of the guest accelerator cards exposed to this instance.

<a name="nested_status_last_known_good_state_future_reservation_specs_specific_sku_properties_instance_properties_local_ssds"></a>The `local_ssds` block supports:

* `disk_size_gb` -
  (Optional)
  Specifies the size of the disk in base-2 GB.

* `interface` -
  (Optional)
  Specifies the disk interface to use for attaching this disk, which is either SCSI or NVME. The default is SCSI.
  Possible values are: `SCSI`, `NVME`.

<a name="nested_status_last_known_good_state_future_reservation_specs_time_window"></a>The `time_window` block supports:

* `start_time` -
  (Optional)
  Start time of the Future Reservation. The startTime is an RFC3339 string.

* `end_time` -
  (Optional)
  End time of the Future Reservation in RFC3339 format.

* `duration` -
  (Optional)
  Specifies the duration of the reservation.
  Structure is [documented below](#nested_status_last_known_good_state_future_reservation_specs_time_window_duration).


<a name="nested_status_last_known_good_state_future_reservation_specs_time_window_duration"></a>The `duration` block supports:

* `seconds` -
  (Optional)
  Span of time at a resolution of a second. Must be from 0 to 315,576,000,000 inclusive.

* `nanos` -
  (Optional)
  Span of time that's a fraction of a second at nanosecond resolution. Durations less than one second are represented with a 0 seconds field and a positive nanos field. Must be from 0 to 999,999,999 inclusive.

<a name="nested_status_last_known_good_state_future_reservation_specs_share_settings"></a>The `share_settings` block supports:

* `share_type` -
  (Optional)
  Type of sharing for this shared-reservation
  Possible values are: `LOCAL`, `ORGANIZATION`, `SPECIFIC_PROJECTS`.

* `projects` -
  (Optional)
  A futureReservations.list of Project names to specify consumer projects for this shared-reservation. This is only valid when shareType's value is SPECIFIC_PROJECTS.

* `project_map` -
  (Optional)
  A map of project id and project config. This is only valid when shareType's value is SPECIFIC_PROJECTS.
  Structure is [documented below](#nested_status_last_known_good_state_future_reservation_specs_share_settings_project_map).


<a name="nested_status_last_known_good_state_future_reservation_specs_share_settings_project_map"></a>The `project_map` block supports:

* `project` - (Required) The identifier for this object. Format specified above.

* `project_id` -
  (Optional)
  The project ID, should be same as the key of this project config in the parent map.

<a name="nested_status_last_known_good_state_existing_matching_usage_info"></a>The `existing_matching_usage_info` block supports:

* `count` -
  (Optional)
  Count representing minimum(FR totalCount, matching_reserved_capacity+matching_unreserved_instances).

* `time_stamp` -
  (Optional)
  Timestamp when the matching usage was calculated.

## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


FutureReservation can be imported using any of these accepted formats:

* `projects/{{project}}/zones/{{zone}}/futureReservations/{{name}}`
* `{{project}}/{{zone}}/{{name}}`
* `{{zone}}/{{name}}`
* `{{name}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import FutureReservation using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/zones/{{zone}}/futureReservations/{{name}}"
  to = google_compute_future_reservation.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), FutureReservation can be imported using one of the formats above. For example:

```
$ terraform import google_compute_future_reservation.default projects/{{project}}/zones/{{zone}}/futureReservations/{{name}}
$ terraform import google_compute_future_reservation.default {{project}}/{{zone}}/{{name}}
$ terraform import google_compute_future_reservation.default {{zone}}/{{name}}
$ terraform import google_compute_future_reservation.default {{name}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
