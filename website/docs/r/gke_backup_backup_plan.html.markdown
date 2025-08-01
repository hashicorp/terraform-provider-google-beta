---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/gkebackup/BackupPlan.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Backup for GKE"
description: |-
  Represents a Backup Plan instance.
---

# google_gke_backup_backup_plan

Represents a Backup Plan instance.


To get more information about BackupPlan, see:

* [API documentation](https://cloud.google.com/kubernetes-engine/docs/add-on/backup-for-gke/reference/rest/v1/projects.locations.backupPlans)
* How-to Guides
    * [Official Documentation](https://cloud.google.com/kubernetes-engine/docs/add-on/backup-for-gke)

## Example Usage - Gkebackup Backupplan Basic


```hcl
resource "google_container_cluster" "primary" {
  name               = "basic-cluster"
  location           = "us-central1"
  initial_node_count = 1
  workload_identity_config {
    workload_pool = "my-project-name.svc.id.goog"
  }
  addons_config {
    gke_backup_agent_config {
      enabled = true
    }
  }
  deletion_protection  = true
  network       = "default"
  subnetwork    = "default"
}

resource "google_gke_backup_backup_plan" "basic" {
  name = "basic-plan"
  cluster = google_container_cluster.primary.id
  location = "us-central1"
  backup_config {
    include_volume_data = true
    include_secrets = true
    all_namespaces = true
  }
}
```
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=gkebackup_backupplan_autopilot&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Gkebackup Backupplan Autopilot


```hcl
resource "google_container_cluster" "primary" {
  name               = "autopilot-cluster"
  location           = "us-central1"
  enable_autopilot = true
  ip_allocation_policy {   
  }
  release_channel {
    channel = "RAPID"
  }
  addons_config {
    gke_backup_agent_config {
      enabled = true
    }
  }
  deletion_protection  = true
  network       = "default"
  subnetwork    = "default"
}

resource "google_gke_backup_backup_plan" "autopilot" {
  name = "autopilot-plan"
  cluster = google_container_cluster.primary.id
  location = "us-central1"
  backup_config {
    include_volume_data = true
    include_secrets = true
    all_namespaces = true
  }
}
```
## Example Usage - Gkebackup Backupplan Cmek


```hcl
resource "google_container_cluster" "primary" {
  name               = "cmek-cluster"
  location           = "us-central1"
  initial_node_count = 1
  workload_identity_config {
    workload_pool = "my-project-name.svc.id.goog"
  }
  addons_config {
    gke_backup_agent_config {
      enabled = true
    }
  }
  deletion_protection  = true
  network       = "default"
  subnetwork    = "default"
}

resource "google_gke_backup_backup_plan" "cmek" {
  name = "cmek-plan"
  cluster = google_container_cluster.primary.id
  location = "us-central1"
  backup_config {
    include_volume_data = true
    include_secrets = true
    selected_namespaces {
      namespaces = ["default", "test"]
    }
    encryption_key {
      gcp_kms_encryption_key = google_kms_crypto_key.crypto_key.id
    }
  }
}

resource "google_kms_crypto_key" "crypto_key" {
  name     = "backup-key"
  key_ring = google_kms_key_ring.key_ring.id
}

resource "google_kms_key_ring" "key_ring" {
  name     = "backup-key"
  location = "us-central1"
}
```
## Example Usage - Gkebackup Backupplan Full


```hcl
resource "google_container_cluster" "primary" {
  name               = "full-cluster"
  location           = "us-central1"
  initial_node_count = 1
  workload_identity_config {
    workload_pool = "my-project-name.svc.id.goog"
  }
  addons_config {
    gke_backup_agent_config {
      enabled = true
    }
  }
  deletion_protection  = true
  network       = "default"
  subnetwork    = "default"
}

resource "google_gke_backup_backup_plan" "full" {
  name = "full-plan"
  cluster = google_container_cluster.primary.id
  location = "us-central1"
  retention_policy {
    backup_delete_lock_days = 30
    backup_retain_days = 180
  }
  backup_schedule {
    cron_schedule = "0 9 * * 1"
  }
  backup_config {
    include_volume_data = true
    include_secrets = true
    selected_applications {
      namespaced_names {
        name = "app1"
        namespace = "ns1"
      }
      namespaced_names {
        name = "app2"
        namespace = "ns2"
      }
    }
  }
}
```
## Example Usage - Gkebackup Backupplan Permissive


```hcl
resource "google_container_cluster" "primary" {
  name               = "permissive-cluster"
  location           = "us-central1"
  initial_node_count = 1
  workload_identity_config {
    workload_pool = "my-project-name.svc.id.goog"
  }
  addons_config {
    gke_backup_agent_config {
      enabled = true
    }
  }
  deletion_protection  = true
  network       = "default"
  subnetwork    = "default"
}

resource "google_gke_backup_backup_plan" "permissive" {
  name = "permissive-plan"
  cluster = google_container_cluster.primary.id
  location = "us-central1"
  retention_policy {
    backup_delete_lock_days = 30
    backup_retain_days = 180
  }
  backup_schedule {
    cron_schedule = "0 9 * * 1"
  }
  backup_config {
    include_volume_data = true
    include_secrets = true
    permissive_mode = true
    selected_applications {
      namespaced_names {
        name = "app1"
        namespace = "ns1"
      }
      namespaced_names {
        name = "app2"
        namespace = "ns2"
      }
    }
  }
}
```
## Example Usage - Gkebackup Backupplan Rpo Daily Window


```hcl
resource "google_container_cluster" "primary" {
  name               = "rpo-daily-cluster"
  location           = "us-central1"
  initial_node_count = 1
  workload_identity_config {
    workload_pool = "my-project-name.svc.id.goog"
  }
  addons_config {
    gke_backup_agent_config {
      enabled = true
    }
  }
  deletion_protection  = true
  network       = "default"
  subnetwork    = "default"
}

resource "google_gke_backup_backup_plan" "rpo_daily_window" {
  name = "rpo-daily-window"
  cluster = google_container_cluster.primary.id
  location = "us-central1"
  retention_policy {
    backup_delete_lock_days = 30
    backup_retain_days = 180
  }
  backup_schedule {
    paused = true
    rpo_config {
      target_rpo_minutes=1440
      exclusion_windows {
        start_time  {
          hours = 12
        }
        duration = "7200s"
        daily = true
      }
      exclusion_windows {
        start_time  {
          hours = 8
          minutes = 40
          seconds = 1
          nanos = 100
        }
        duration = "3600s"
        single_occurrence_date {
          year = 2024
          month = 3
          day = 16
        }
      }
    }
  }
  backup_config {
    include_volume_data = true
    include_secrets = true
    all_namespaces = true
  }
}
```
## Example Usage - Gkebackup Backupplan Rpo Weekly Window


```hcl
resource "google_container_cluster" "primary" {
  name               = "rpo-weekly-cluster"
  location           = "us-central1"
  initial_node_count = 1
  workload_identity_config {
    workload_pool = "my-project-name.svc.id.goog"
  }
  addons_config {
    gke_backup_agent_config {
      enabled = true
    }
  }
  deletion_protection  = true
  network       = "default"
  subnetwork    = "default"
}

resource "google_gke_backup_backup_plan" "rpo_weekly_window" {
  name = "rpo-weekly-window"
  cluster = google_container_cluster.primary.id
  location = "us-central1"
  retention_policy {
    backup_delete_lock_days = 30
    backup_retain_days = 180
  }
  backup_schedule {
    paused = true
    rpo_config {
      target_rpo_minutes=1440
      exclusion_windows {
        start_time  {
          hours = 1
          minutes = 23
        }
        duration = "1800s"
        days_of_week {
          days_of_week = ["MONDAY", "THURSDAY"]
        }
      }
      exclusion_windows {
        start_time  {
          hours = 12
        }
        duration = "3600s"
        single_occurrence_date {
          year = 2024
          month = 3
          day = 17
        }
      }
      exclusion_windows {
        start_time  {
          hours = 8
          minutes = 40
        }
        duration = "600s"
        single_occurrence_date {
          year = 2024
          month = 3
          day = 18
        }
      }
    }
  }
  backup_config {
    include_volume_data = true
    include_secrets = true
    all_namespaces = true
  }
}
```

## Argument Reference

The following arguments are supported:


* `name` -
  (Required)
  The full name of the BackupPlan Resource.

* `cluster` -
  (Required)
  The source cluster from which Backups will be created via this BackupPlan.

* `location` -
  (Required)
  The region of the Backup Plan.


* `description` -
  (Optional)
  User specified descriptive string for this BackupPlan.

* `retention_policy` -
  (Optional)
  RetentionPolicy governs lifecycle of Backups created under this plan.
  Structure is [documented below](#nested_retention_policy).

* `labels` -
  (Optional)
  Description: A set of custom labels supplied by the user.
  A list of key->value pairs.
  Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.

  **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
  Please refer to the field `effective_labels` for all of the labels present on the resource.

* `backup_schedule` -
  (Optional)
  Defines a schedule for automatic Backup creation via this BackupPlan.
  Structure is [documented below](#nested_backup_schedule).

* `deactivated` -
  (Optional)
  This flag indicates whether this BackupPlan has been deactivated.
  Setting this field to True locks the BackupPlan such that no further updates will be allowed
  (except deletes), including the deactivated field itself. It also prevents any new Backups
  from being created via this BackupPlan (including scheduled Backups).

* `backup_config` -
  (Optional)
  Defines the configuration of Backups created via this BackupPlan.
  Structure is [documented below](#nested_backup_config).

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.



<a name="nested_retention_policy"></a>The `retention_policy` block supports:

* `backup_delete_lock_days` -
  (Optional)
  Minimum age for a Backup created via this BackupPlan (in days).
  Must be an integer value between 0-90 (inclusive).
  A Backup created under this BackupPlan will not be deletable
  until it reaches Backup's (create time + backup_delete_lock_days).
  Updating this field of a BackupPlan does not affect existing Backups.
  Backups created after a successful update will inherit this new value.

* `backup_retain_days` -
  (Optional)
  The default maximum age of a Backup created via this BackupPlan.
  This field MUST be an integer value >= 0 and <= 365. If specified,
  a Backup created under this BackupPlan will be automatically deleted
  after its age reaches (createTime + backupRetainDays).
  If not specified, Backups created under this BackupPlan will NOT be
  subject to automatic deletion. Updating this field does NOT affect
  existing Backups under it. Backups created AFTER a successful update
  will automatically pick up the new value.
  NOTE: backupRetainDays must be >= backupDeleteLockDays.
  If cronSchedule is defined, then this must be <= 360 * the creation interval.
  If rpo_config is defined, then this must be
  <= 360 * targetRpoMinutes/(1440minutes/day)

* `locked` -
  (Optional)
  This flag denotes whether the retention policy of this BackupPlan is locked.
  If set to True, no further update is allowed on this policy, including
  the locked field itself.

<a name="nested_backup_schedule"></a>The `backup_schedule` block supports:

* `cron_schedule` -
  (Optional)
  A standard cron string that defines a repeating schedule for
  creating Backups via this BackupPlan.
  This is mutually exclusive with the rpoConfig field since at most one
  schedule can be defined for a BackupPlan.
  If this is defined, then backupRetainDays must also be defined.

* `paused` -
  (Optional)
  This flag denotes whether automatic Backup creation is paused for this BackupPlan.

* `rpo_config` -
  (Optional)
  Defines the RPO schedule configuration for this BackupPlan. This is mutually
  exclusive with the cronSchedule field since at most one schedule can be defined
  for a BackupPLan. If this is defined, then backupRetainDays must also be defined.
  Structure is [documented below](#nested_backup_schedule_rpo_config).


<a name="nested_backup_schedule_rpo_config"></a>The `rpo_config` block supports:

* `target_rpo_minutes` -
  (Required)
  Defines the target RPO for the BackupPlan in minutes, which means the target
  maximum data loss in time that is acceptable for this BackupPlan. This must be
  at least 60, i.e., 1 hour, and at most 86400, i.e., 60 days.

* `exclusion_windows` -
  (Optional)
  User specified time windows during which backup can NOT happen for this BackupPlan.
  Backups should start and finish outside of any given exclusion window. Note: backup
  jobs will be scheduled to start and finish outside the duration of the window as
  much as possible, but running jobs will not get canceled when it runs into the window.
  All the time and date values in exclusionWindows entry in the API are in UTC. We
  only allow <=1 recurrence (daily or weekly) exclusion window for a BackupPlan while no
  restriction on number of single occurrence windows.
  Structure is [documented below](#nested_backup_schedule_rpo_config_exclusion_windows).


<a name="nested_backup_schedule_rpo_config_exclusion_windows"></a>The `exclusion_windows` block supports:

* `start_time` -
  (Required)
  Specifies the start time of the window using time of the day in UTC.
  Structure is [documented below](#nested_backup_schedule_rpo_config_exclusion_windows_exclusion_windows_start_time).

* `duration` -
  (Required)
  Specifies duration of the window in seconds with up to nine fractional digits,
  terminated by 's'. Example: "3.5s". Restrictions for duration based on the
  recurrence type to allow some time for backup to happen:
    - single_occurrence_date:  no restriction
    - daily window: duration < 24 hours
    - weekly window:
      - days of week includes all seven days of a week: duration < 24 hours
      - all other weekly window: duration < 168 hours (i.e., 24 * 7 hours)

* `single_occurrence_date` -
  (Optional)
  No recurrence. The exclusion window occurs only once and on this date in UTC.
  Only one of singleOccurrenceDate, daily and daysOfWeek may be set.
  Structure is [documented below](#nested_backup_schedule_rpo_config_exclusion_windows_exclusion_windows_single_occurrence_date).

* `daily` -
  (Optional)
  The exclusion window occurs every day if set to "True".
  Specifying this field to "False" is an error.
  Only one of singleOccurrenceDate, daily and daysOfWeek may be set.

* `days_of_week` -
  (Optional)
  The exclusion window occurs on these days of each week in UTC.
  Only one of singleOccurrenceDate, daily and daysOfWeek may be set.
  Structure is [documented below](#nested_backup_schedule_rpo_config_exclusion_windows_exclusion_windows_days_of_week).


<a name="nested_backup_schedule_rpo_config_exclusion_windows_exclusion_windows_start_time"></a>The `start_time` block supports:

* `hours` -
  (Optional)
  Hours of day in 24 hour format.

* `minutes` -
  (Optional)
  Minutes of hour of day.

* `seconds` -
  (Optional)
  Seconds of minutes of the time.

* `nanos` -
  (Optional)
  Fractions of seconds in nanoseconds.

<a name="nested_backup_schedule_rpo_config_exclusion_windows_exclusion_windows_single_occurrence_date"></a>The `single_occurrence_date` block supports:

* `year` -
  (Optional)
  Year of the date.

* `month` -
  (Optional)
  Month of a year.

* `day` -
  (Optional)
  Day of a month.

<a name="nested_backup_schedule_rpo_config_exclusion_windows_exclusion_windows_days_of_week"></a>The `days_of_week` block supports:

* `days_of_week` -
  (Optional)
  A list of days of week.
  Each value may be one of: `MONDAY`, `TUESDAY`, `WEDNESDAY`, `THURSDAY`, `FRIDAY`, `SATURDAY`, `SUNDAY`.

<a name="nested_backup_config"></a>The `backup_config` block supports:

* `include_volume_data` -
  (Optional)
  This flag specifies whether volume data should be backed up when PVCs are
  included in the scope of a Backup.

* `include_secrets` -
  (Optional)
  This flag specifies whether Kubernetes Secret resources should be included
  when they fall into the scope of Backups.

* `encryption_key` -
  (Optional)
  This defines a customer managed encryption key that will be used to encrypt the "config"
  portion (the Kubernetes resources) of Backups created via this plan.
  Structure is [documented below](#nested_backup_config_encryption_key).

* `all_namespaces` -
  (Optional)
  If True, include all namespaced resources.

* `selected_namespaces` -
  (Optional)
  If set, include just the resources in the listed namespaces.
  Structure is [documented below](#nested_backup_config_selected_namespaces).

* `selected_applications` -
  (Optional)
  A list of namespaced Kubernetes Resources.
  Structure is [documented below](#nested_backup_config_selected_applications).

* `permissive_mode` -
  (Optional)
  This flag specifies whether Backups will not fail when
  Backup for GKE detects Kubernetes configuration that is
  non-standard or requires additional setup to restore.


<a name="nested_backup_config_encryption_key"></a>The `encryption_key` block supports:

* `gcp_kms_encryption_key` -
  (Required)
  Google Cloud KMS encryption key. Format: projects/*/locations/*/keyRings/*/cryptoKeys/*

<a name="nested_backup_config_selected_namespaces"></a>The `selected_namespaces` block supports:

* `namespaces` -
  (Required)
  A list of Kubernetes Namespaces.

<a name="nested_backup_config_selected_applications"></a>The `selected_applications` block supports:

* `namespaced_names` -
  (Required)
  A list of namespaced Kubernetes resources.
  Structure is [documented below](#nested_backup_config_selected_applications_namespaced_names).


<a name="nested_backup_config_selected_applications_namespaced_names"></a>The `namespaced_names` block supports:

* `namespace` -
  (Required)
  The namespace of a Kubernetes Resource.

* `name` -
  (Required)
  The name of a Kubernetes Resource.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/{{location}}/backupPlans/{{name}}`

* `uid` -
  Server generated, unique identifier of UUID format.

* `etag` -
  etag is used for optimistic concurrency control as a way to help prevent simultaneous
  updates of a backup plan from overwriting each other. It is strongly suggested that
  systems make use of the 'etag' in the read-modify-write cycle to perform BackupPlan updates
  in order to avoid race conditions: An etag is returned in the response to backupPlans.get,
  and systems are expected to put that etag in the request to backupPlans.patch or
  backupPlans.delete to ensure that their change will be applied to the same version of the resource.

* `protected_pod_count` -
  The number of Kubernetes Pods backed up in the last successful Backup created via this BackupPlan.

* `state` -
  The State of the BackupPlan.

* `state_reason` -
  Detailed description of why BackupPlan is in its current state.

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


BackupPlan can be imported using any of these accepted formats:

* `projects/{{project}}/locations/{{location}}/backupPlans/{{name}}`
* `{{project}}/{{location}}/{{name}}`
* `{{location}}/{{name}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import BackupPlan using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/locations/{{location}}/backupPlans/{{name}}"
  to = google_gke_backup_backup_plan.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), BackupPlan can be imported using one of the formats above. For example:

```
$ terraform import google_gke_backup_backup_plan.default projects/{{project}}/locations/{{location}}/backupPlans/{{name}}
$ terraform import google_gke_backup_backup_plan.default {{project}}/{{location}}/{{name}}
$ terraform import google_gke_backup_backup_plan.default {{location}}/{{name}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
