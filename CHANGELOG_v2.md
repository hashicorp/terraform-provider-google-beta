## 2.20.3 (March 10, 2020)

NOTES:
* `2.20.3` is a backport release, and some changes will not appear in `3.X` series releases until `3.12.0`.
To upgrade to `3.X` you will need to perform a large jump in versions, and it is _strongly_ advised that you attempt to upgrade to `3.X` instead of using this release.
* `2.20.3` is primarily a preventative fix, in anticipation of a change in API response messages adding a default value.

BUG FIXES:
* compute: fixed error when reading `google_compute_instance_template` resources with `network_interface[*].name` set. ([#1815](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1815))

## 2.20.2 (February 04, 2020)

BUG FIXES:
* bigtable: fixed diff for DEVELOPMENT instances that are returned from the API with one node ([#1704](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1704))

## 2.20.1 (December 13, 2019)

BUG FIXES:
* iam: Fixed a bug that causes badRequest errors on IAM resources due to deleted serviceAccount principals ([#1501](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1501))

## 2.20.2 (February 03, 2020)

BUG FIXES:
* bigtable: fixed diff for DEVELOPMENT instances that are returned from the API with one node ([#1704](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1704))

## 2.20.1 (December 13, 2019)

**Note**: 2.20.1 is a backport release. The changes in it are unavailable in 3.0.0-beta.1 through 3.2.0.

BUG FIXES:
* iam: Fixed a bug that causes badRequest errors on IAM resources due to deleted serviceAccount principals ([#1501](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1501))

## 2.20.0 (November 13, 2019)

BREAKING CHANGES:
* `google_compute_instance_iam_*` resources now support IAM Conditions. If any conditions had been created out of band before this release, take extra care to ensure they are present in your Terraform config so the provider doesn't try to create new bindings with no conditions. Terraform will show a diff that it is adding the condition to the resource, which is safe to apply. ([#1360](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1360))
* `google_iap_app_engine_version_iam_*` resources now support IAM Conditions. If any conditions had been created out of band before this release, take extra care to ensure they are present in your Terraform config so the provider doesn't try to create new bindings with no conditions. Terraform will show a diff that it is adding the condition to the resource, which is safe to apply. ([#1352](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1352))
* `google_iap_web_backend_service_iam_*` resources now support IAM Conditions. If any conditions had been created out of band before this release, take extra care to ensure they are present in your Terraform config so the provider doesn't try to create new bindings with no conditions. Terraform will show a diff that it is adding the condition to the resource, which is safe to apply. ([#1352](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1352))
* `google_project_iam_*` resources now support IAM Conditions. If any conditions had been created out of band before this release, take extra care to ensure they are present in your Terraform config so the provider doesn't try to create new bindings with no conditions. Terraform will show a diff that it is adding the condition to the resource, which is safe to apply. ([#1321](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1321))
* compute: the `backend.group` field is now required for `google_compute_region_backend_service`. Configurations without this would not have worked, so this isn't considered an API break. ([#1311](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1311))

FEATURES:
* **New Resource:** `google_data_fusion_instance` ([#1339](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1339))

IMPROVEMENTS:
* bigtable: added import support to `google_bigtable_table` ([#1350](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1350))
* compute: `load_balancing_scheme` for `google_compute_forwarding_rule` now accepts `INTERNAL_MANAGED` as a value. ([#1311](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1311))
* compute: added support for L7 ILB to google_compute_region_backend_service. ([#1311](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1311))
* compute: extended backend configuration options for `google_compute_region_backend_service` to include `backend.balancing_mode`, `backend.capacity_scaler`, `backend.max_connections`, `backend.max_connections_per_endpoint`, `backend.max_connections_per_instance`, `backend.max_rate`, `backend.max_rate_per_endpoint`, `backend.max_rate_per_instance`, and `backend.max_utilization` ([#1311](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1311))
* iam: changed the `id` for many IAM resources to the reference resource long name. Updated `instance_name` on `google_compute_instance_iam` and `subnetwork` on `google_compute_subnetwork` to their respective long names in state ([#1360](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1360))
* iap: added support for IAM Conditions to the `google_compute_instance_iam_*` resources ([#1360](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1360))
* iap: added support for IAM Conditions to the `google_iap_app_engine_version_iam_*` resources ([#1352](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1352))
* iap: added support for IAM Conditions to the `google_iap_web_backend_service_iam_*` resources ([#1352](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1352))
* logging: added `display_name` field to `google_logging_metric` resource ([#1344](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1344))
* monitoring: Added `validate_ssl` to `google_monitoring_uptime_check_config` ([#1243](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1243))
* project: added batching functionality to `google_project_service` read calls, so fewer API requests are made ([#1354](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1354))
* resourcemanager: added support for IAM Conditions to the `google_project_iam_*` resources ([#1321](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1321))
* storage: added notification_id field to `google_storage_notification` ([#1368](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1368))

BUG FIXES:
* compute: fixed issue where setting a 0 for `min_replicas` in `google_compute_autoscaler` and `google_compute_region_autoscaler` would set that field to its server-side default instead of 0. ([#1351](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1351))
* dns: fixed crash when `network` blocks are defined without `network_url`s ([#1345](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1345))
* google: used the correct update method for google_service_account.description ([#1362](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1362))
* logging: fixed issue where logging exclusion resources silently failed when being mutated in parallel ([#1329](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1329))

## 2.19.0 (November 05, 2019)

DEPRECATIONS:
* `compute`: deprecated `enable_flow_logs` on `google_compute_subnetwork`. The presence of the `log_config` block signals that flow logs are enabled for a subnetwork ([#1320](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1320))
* `compute`: deprecated `instance_template` for `google_compute_instance_group_manager` and `google_compute_region_instance_group_manager` . Use `version.instance_template` instead. ([#1309](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1309))
* `compute`: deprecated `update_strategy` for `google_compute_instance_group_manager` . Use `update_policy` instead. ([#1309](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1309))
* `container`: deprecated `google_container_cluster` `ip_allocation_policy.create_subnetwork`, `ip_allocation_policy.subnetwork_name`, `ip_allocation_policy.node_ipv4_cidr_block`. Define an explicit `google_compute_subnetwork` and use `subnetwork` instead. ([#1312](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1312))
* `container`: deprecated `google_container_cluster` `ip_allocation_policy.use_ip_aliases`. If it's set to true, remove it from your config. If false, remove `ip_allocation_policy` as a whole. ([#1312](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1312))
* `iam`: Deprecated `pgp_key` on `google_service_account_key` resource. See https://www.terraform.io/docs/extend/best-practices/sensitive-state.html for more information. ([#1326](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1326))

BREAKING CHANGES:
* `google_service_account_iam_*` resources now support IAM Conditions. If any conditions had been created out of band before this release, take extra care to ensure they are present in your Terraform config so the provider doesn't try to create new bindings with no conditions. Terraform will show a diff that it is adding the condition to the resource, which is safe to apply. ([#1188](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1188))

FEATURES:
* `compute`: added `google_compute_router` datasource ([#1233](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1233))

IMPROVEMENTS:
* `cloudbuild`: added ability to specify `name` for `cloud_build_trigger` to avoid name collisions when creating multiple triggers at once. ([#1277](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1277))
* `compute`: added support for multiple versions of `instance_template` and granular control of the update policies for `google_compute_instance_group_manager` and `google_compute_region_instance_group_manager`. ([#1309](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1309))
* `container`: added `taint` field in GKE resources to the GA `google` provider ([#1296](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1296))
* `container`: fix a diff created in the cloud console when `MaintenanceExclusions` are added. ([#1310](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1310))
* `container`: added `maintenance_policy.recurring_window` support to `google_container_cluster`, significantly increasing expressive range. ([#1292](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1292))
* `compute`: added `google_compute_instance` support for display device (Virtual Displays) ([#1313](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1313))
* `iam`: added support for IAM Conditions to the `google_service_account_iam_*` resources (beta provider only) ([#1188](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1188))
* `iam`: added `description` to `google_service_account`. ([#1291](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1291))

BUG FIXES:
* `appengine`: Resolved permadiff in `google_app_engine_domain_mapping.ssl_settings.certificate_id`. ([#1303](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1303))
* `storage`: Fixed error in `google_storage_bucket` where locked retention policies would cause a bucket to report failure on all updates (even though updates were applied correctly). ([#1307](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1307))
* `container`: Fixed nil reference to ShieldedNodes. ([#1314](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1314))

## 2.18.1 (October 25, 2019)

BUGS:
* `resourcemanager`: fixed deleting the default network in `google_project` ([#1299](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1299))

## 2.18.0 (October 23, 2019)

KNOWN ISSUES:
* `resourcemanager`: `google_project` `auto_create_network` is failing to delete networks when set to `false`. Use an earlier provider version to resolve.

DEPRECATIONS:
* `container`: The `kubernetes_dashboard` addon is deprecated for `google_container_cluster`. ([#1247](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1247))

FEATURES:
* **New Resource:** `google_app_engine_application_url_dispatch_rules` ([#1262](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1262))

IMPROVEMENTS:
* `all`: increased support for custom endpoints across the provider ([#1244](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1244))
* `appengine`: added the ability to delete the parent service of `google_app_engine_standard_app_version` ([#1222](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1222))
* `container`: Added `shielded_instance_config` attribute to `node_config` ([#1198](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1198))
* `container`: Allow the configuration of release channels when creating GKE clusters. ([#1260](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1260))
* `dataflow`: added `ip_configuration` option to `job`. ([#1284](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1284))
* `pubsub`: Added field `oidc_token` to `google_pubsub_subscription` ([#1265](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1265))
* `sql`: added `location` field to `backup_configuration` block in `google_sql_database_instance` ([#1282](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1282))

BUGS:
* `all`: fixed the custom endpoint version used by older legacy REST clients ([#1274](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1274))
* `bigquery`: fix issue with `google_bigquery_data_transfer_config` `params` crashing on boolean values ([#1263](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1263))
* `cloudrun`: fixed the apiVersion sent in `google_cloud_run_domain_mapping` requests ([#1251](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1251))
* `compute`: added support for updating multiple fields at once to `google_compute_subnetwork` ([#1269](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1269))
* `compute`: fixed diffs in `google_compute_instance_group`'s `network` field when equivalent values were specified ([#1286](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1286))
* `compute`: fixed issues updating `google_compute_instance_group`'s `instances` field when config/state values didn't match ([#1286](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1286))
* `iam`: fixed bug where IAM binding wouldn't replace members if they were deleted outside of terraform. ([#1272](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1272))
* `pubsub`: Fixed permadiff due to interaction of organization policies and `google_pubsub_topic`. ([#1281](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1281))

## 2.17.0 (October 08, 2019)

NOTES:
* An [upgrade guide](https://www.terraform.io/docs/providers/google/version_3_upgrade.html) has been started for the upcoming 3.0.0 release. ([#1220](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1220))
* `google_project_services` users of provider versions prior to `2.17.0` should update, as past versions of the provider will not handle an upcoming rename of `bigquery-json.googleapis.com` to `bigquery.googleapis.com` well. See https://github.com/terraform-providers/terraform-provider-google/issues/4590 for details. ([#1234](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1234))

DEPRECATIONS:
* `google_project_services` ([#1218](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1218))

FEATURES:
* **New Resource:** `google_bigtable_gc_policy` ([#1213](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1213))
* **New Resource:** `google_binary_authorization_attestor_iam_policy` ([#1166](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1166))
* **New Resource:** `google_compute_region_ssl_certificate` ([#1183](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1183))
* **New Resource:** `google_compute_region_target_http_proxy` ([#1183](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1183))
* **New Resource:** `google_compute_region_target_https_proxy` ([#1183](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1183))
* **New Resource:** `google_iap_app_engine_service_iam_*` ([#1205](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1205))
* **New Resource:** `google_iap_app_engine_version_iam_*` ([#1205](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1205))
* **New Resource:** `google_storage_bucket_access_control` ([#1177](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1177))

IMPROVEMENTS:
* all: made `monitoring-read` scope available. ([#1208](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1208))
* bigquery: added support for default customer-managed encryption keys (CMEK) for BigQuery datasets. ([#1081](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1081))
* bigtable: import support added to `google_bigtable_instance` ([#1224](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1224))
* cloudbuild: added `github` field in `google_cloudbuild_trigger`. ([#1229](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1229))
* container: moved `default_max_pods_per_node` to ga. ([#1235](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1235))
* containeranalysis: moved `google_containeranalysis_note` to ga ([#1166](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1166))
* projectservice: added mitigations for bigquery-json to bigquery rename in project service resources. ([#1234](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1234))

BUGS:
* cloudscheduler: Fixed permadiff for `app_engine_http_target.app_engine_routing` on `google_cloud_scheduler_job` ([#1131](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1131))
* compute: Added ability to set `quic_override` on `google_compute_https_target_proxy` to empty. ([#1219](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1219))
* compute: Fix bug where changes to `region_backend_service.backends.failover` was not detected. ([#1236](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1236))
* compute: fixed `google_compute_router_peer` to default if empty for `advertise_mode` ([#1163](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1163))
* compute: fixed perma-diff in `google_compute_router_nat` when referencing subnetwork via `name` ([#1194](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1194))
* compute: fixed perma-diff in `google_compute_router_nat` when referencing subnetwork via `name` ([#1194](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1194))
* container: fixed an overly-aggressive validation for `master_ipv4_cidr_block` in `google_container_cluster` ([#1211](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1211))

## 2.16.0 (September 24, 2019)

KNOWN ISSUES:
* Based on an upstream change, users of the `google_project_services` resource may have seen the `bigquery.googleapis.com` service added and the `bigquery-json.googleapis.com` service removed, causing a diff. This was later reverted, causing another diff. This issue is being tracked as https://github.com/terraform-providers/terraform-provider-google/issues/4590.

FEATURES:
* **New Resource**: `google_compute_region_url_map` is now available. To support this, the `protocol` for `google_compute_region_backend_service` can now be set to `HTTP`, `HTTPS`, `HTTP2`, and `SSL`. ([#1161](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1161))
* **New Resource**: Adds `google_runtimeconfig_config_iam_*` resources ([#1138](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1138))
* **New Resource**: Added `google_compute_resource_policy` and `google_compute_disk_resource_policy_attachment` to manage `google_compute_disk` resource policies as fine-grained resources ([#1085](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1085))

ENHANCEMENTS:
* composer: Add `python_version` and ability to set `image_version` in `google_composer_environment` in the GA provider ([#1143](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1143))
* compute: `google_compute_global_forwarding_rule` now supports `metadata_filters`. ([#1160](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1160))
* compute: `google_compute_backend_service` now supports `locality_lb_policy`, `outlier_detection`, `consistent_hash`, and `circuit_breakers`. ([#1118](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1118))
* compute: Add support for `guest_os_features` to resource `google_compute_image` ([#1156](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1156))
* compute: Added `drain_nat_ips` to `google_compute_router_nat` ([#1155](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1155))
* container: google_container_node_pool now supports node_locations to specify specific node zones. ([#1154](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1154))
* googleapis: `google_netblock_ip_ranges` data source now has a `private-googleapis` field, for the IP addresses used for Private Google Access for services that do not support VPC Service Controls API access. ([#1102](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1102))
* project: `google_project_iam_*` Properly set the `project` field in state ([#1158](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1158))

BUG FIXES:
* cloudiot: Fixed error where `subfolder_matches` were not set in `google_cloudiot_registry` `event_notification_configs` ([#1175](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1175))

## 2.15.0 (September 17, 2019)

FEATURES:
* **New Resource**: `google_iap_web_iam_binding/_member/_policy` are now available for managing IAP web IAM permissions ([#1044](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1044))
* **New Resource**: `google_iap_web_backend_service_binding/_member/_policy` are now available for managing IAM permissions on IAP enabled backend services ([#1044](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1044))
* **New Resource**: `google_iap_web_type_compute_iam_binding/_member/_policy` are now available for managing IAM permissions on IAP enabled compute services ([#1044](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1044))
* **New Resource**: `google_iap_web_type_app_engine_iam_binding/_member/_policy` are now available for managing IAM permissions on IAP enabled App Engine applications ([#1044](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1044))
* **New Resource**: Add the new resource `google_app_engine_domain_mapping` ([#1079](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1079))
* **New Resource**: `google_cloudfunctions_function_iam_policy`, `google_cloudfunctions_function_iam_binding`, and `google_cloudfunctions_function_iam_member` ([#1121](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1121))
* **New Resource**: `google_compute_reservation` allows you to reserve instance capacity in GCE. ([#1086](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1086))
* **New Resource**: `google_compute_region_health_check` is now available. This and `google_compute_health_check` now include additional support for HTTP2 health checks. ([#1058](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1058))

ENHANCEMENTS:
* compute: Added full routing options to `google_compute_router_peer` ([#1104](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1104))
* compute: add `tunnel_id` to `google_compute_vpn_tunnel` and `gateway_id` to `google_compute_vpn_gateway` ([#1106](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1106))
* compute: `google_compute_subnetwork` now includes the `purpose` and `role` fields. ([#1051](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1051))
* compute: add `purpose` field to `google_compute_address` ([#1115](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1115))
* compute: add `mode` option to `google_compute_instance.boot_disk` ([#1119](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1119))
* compute: `google_compute_firewall` does not show a diff if allowed or denied rules are specified with uppercase protocol values ([#1144](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1144))
* compute: Add support for the `log_config` block to `compute_backend_service` (Beta only) ([#1137](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1137))
* logging: added `metric_descriptor.unit` to `google_logging_metric` resource ([#1117](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1117))

BUG FIXES:
* all: More classes of generic HTTP errors are retried provider-wide. ([#1120](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1120))
* container: Fix error when `master_authorized_networks_config` is removed from the `google_container_cluster` configuration. ([#1133](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1133))
* iam: Make `google_service_account_` and `google_service_account_iam_*` validation less restrictive to allow for more default service accounts ([#1109](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1109))
* iam: set auditconfigs in state for google_\*\_iam_policy resources ([#1134](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1134))
* logging: `google_logging_metric` `explicit` bucket option can now be set ([#1096](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1096))
* pubsub: Add retry for Pubsub Topic creation when project is still initializing org policies ([#1094](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1094))
* servicenetworking: remove need for provider-level project to delete connection ([#1132](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1132))
* sql: Add more retries for operationInProgress 409 errors for `google_sql_database_instance` ([#1108](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1108))

MISC:
* The User-Agent header that Terraform sends has been updated to correctly report the version of Terraform being run, and has minorly changed the formatting on the Terraform string. ([#1107](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1107))


## 2.14.0 (August 28, 2019)

DEPRECATIONS:
* cloudiot: `resource_cloudiot_registry`'s `event_notification_config` field has been deprecated. ([#1064](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1064))

FEATURES:
* **New Resource**: `google_bigtable_app_profile` is now available ([#988](https://github.com/terraform-providers/terraform-provider-google-beta/issues/988))
* **New Resource**: `google_ml_engine_model` ([#957](https://github.com/terraform-providers/terraform-provider-google-beta/issues/957))
* **New Resource**: `google_dataproc_autoscaling_policy` ([#1078](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1078))
* **New Data Source**: `google_kms_secret_ciphertext` ([#1011](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1011))

ENHANCEMENTS:
* bigquery: Add support for clustering/partitioning to bigquery_table ([#1025](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1025))
* bigtable: `num_nodes` can now be updated in `google_bigtable_instance` ([#1067](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1067))
* cloudiot: `resource_cloudiot_registry` now has fields plural `event_notification_configs` and `log_level`, and `event_notification_config` has been deprecated. ([#1064](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1064))
* cloud_run: New output-only fields have been added to google_cloud_run_service' status. ([#1071](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1071))
* compute: Adding bandwidth attribute to interconnect attachment. ([#1016](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1016))
* compute: `google_compute_region_instance_group_manager.update_policy` now supports `instance_redistribution_type` ([#1073](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1073))
* compute: adds admin_enabled to google_compute_interconnect_attachment ([#1072](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1072))
* compute: The compute routes includes next_hop_ilb attribute support in beta. ([#1076](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1076))
* scheduler: Add support for `oauth_token` and `oidc_token` on resource `google_cloud_scheduler_job` ([#1024](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1024))

BUG FIXES:
* containerregistry: Correctly handle domain-scoped projects ([#1035](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1035))
* iam: Fixed regression in 2.13.0 for permadiff on empty members in IAM policy bindings. ([#1092](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1092))
* project: `google_project_iam_custom_role` now sets the project properly on import. ([#1089](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1089))
* sql: Added back a missing import format for `google_sql_database`. ([#1061](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1061))

## 2.13.0 (August 15, 2019)

KNOWN ISSUES:
* `bigtable`: `google_bigtable_instance` may cause a panic on Terraform `0.11`. This was resolved in `2.17.0`.

FEATURES:
* **New Resource**: added the `google_vpc_access_connector` resource and the `vpc_connector` option on the `google_cloudfunctions_function` resource. ([#1004](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1004))
* **New Resource**: Added `google_scc_source` resource for managing Cloud Security Command Center sources in Terraform ([#1033](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1033))
* **New Data Source**: `google_compute_network_endpoint_group`([#999](https://github.com/terraform-providers/terraform-provider-google-beta/issues/999))

ENHANCEMENTS:
* bigquery: Added support for `google_bigquery_data_transfer_config` (which include scheduled queries). ([#975](https://github.com/terraform-providers/terraform-provider-google-beta/issues/975))
* bigtable: `google_bigtable_instance` max number of `cluster` blocks is now 4 ([#995](https://github.com/terraform-providers/terraform-provider-google-beta/issues/995))
* binary_authorization: Added `globalPolicyEvaluationMode` to `google_binary_authorization_policy`. ([#987](https://github.com/terraform-providers/terraform-provider-google-beta/issues/987))
* cloudfunctions: Allow partial URIs in google_cloudfunctions_function event_trigger.resource ([#1009](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1009))
* compute: Enable update for `google_compute_router_nat` ([#979](https://github.com/terraform-providers/terraform-provider-google-beta/issues/979))
* netblock: extended `google_netblock_ip_ranges` to support multiple useful IP address ranges that have a special meaning on GCP. ([#986](https://github.com/terraform-providers/terraform-provider-google-beta/issues/986))
* project: Wrapped API requests with retries for `google_project`, `google_folder`, and `google_*_organization_policy` ([#971](https://github.com/terraform-providers/terraform-provider-google-beta/issues/971))
* project: IAM and service requests are now batched ([#1014](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1014))
* provider: allow provider's region to be specified as a self_link ([#1022](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1022))
* provider: Adds new provider-level field `user_project_override`, which allows billing, quota checks, and service enablement checks to occur against the project a resource is in instead of the project the credentials are from. ([#1010](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1010))
* pubsub: Pub/Sub topic geo restriction support. ([#989](https://github.com/terraform-providers/terraform-provider-google-beta/issues/989))

BUG FIXES:
* binary_authorization: don't diff when attestation authority note public keys don't have an ID in the config ([#1042](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1042))
* compute: instance descriptions will now be stored in state ([#990](https://github.com/terraform-providers/terraform-provider-google-beta/issues/990))
* container: `key_name` in `google_container_cluster.database_encryption` is no longer a required field. ([#1032](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1032))
* project: ignore errors when deleting a default network that doesn't exist ([#991](https://github.com/terraform-providers/terraform-provider-google-beta/issues/991))

## 2.12.0 (August 01, 2019)

FEATURES:
* **New Data Source**: `google_kms_crypto_key_version` - Provides access to KMS key version data with Google Cloud KMS. ([#964](https://github.com/terraform-providers/terraform-provider-google-beta/issues/964))
* **New Resource**: `google_cloud_run_service` - Set up a cloud run service ([#757](https://github.com/terraform-providers/terraform-provider-google-beta/issues/757))
* **New Resource**: `google_cloud_run_domain_mapping` - Allows custom domains to map to a cloud run service ([#757](https://github.com/terraform-providers/terraform-provider-google-beta/issues/757))

ENHANCEMENTS:
* binary_authorization: Add support for Cloud KMS PKIX keys to `binary_authorization_attestor`. ([#964](https://github.com/terraform-providers/terraform-provider-google-beta/issues/964))
* composer: Add private IP config for `google_composer_environment` ([#908](https://github.com/terraform-providers/terraform-provider-google-beta/issues/908))
* compute: add support for port_specification to resource `google_compute_health_check` ([#933](https://github.com/terraform-providers/terraform-provider-google-beta/issues/933))
* compute: Fixed import formats for `google_compute_network_endpoint` and add location-only import formats ([#947](https://github.com/terraform-providers/terraform-provider-google-beta/issues/947))
* compute: add support for `resource_policies` to resource `google_compute_disk` ([#960](https://github.com/terraform-providers/terraform-provider-google-beta/issues/960))
* compute: Support labelling for compute_instance boot_disks and compute_instance_template disks. ([#982](https://github.com/terraform-providers/terraform-provider-google-beta/issues/982))
* container: `workload_identity_config` in `google_container_cluster` can now be updated without recreating the cluster. ([#896](https://github.com/terraform-providers/terraform-provider-google-beta/issues/896))
* container: validate that master_ipv4_cidr_block is set if enable_private_nodes is true ([#948](https://github.com/terraform-providers/terraform-provider-google-beta/issues/948))
* dataflow: added support for user-defined `labels` on resource `google_dataflow_job` ([#970](https://github.com/terraform-providers/terraform-provider-google-beta/issues/970))
* dataproc: add support for `optional_components` to resource `resource_dataproc_cluster` ([#961](https://github.com/terraform-providers/terraform-provider-google-beta/issues/961))
* project: add checks to import to prevent importing by project number instead of id ([#954](https://github.com/terraform-providers/terraform-provider-google-beta/issues/954))
* storage: add support for `retention_policy` to resource `google_storage_bucket` ([#949](https://github.com/terraform-providers/terraform-provider-google-beta/issues/949))

BUG FIXES:
* access_context_manager: import format checking ([#952](https://github.com/terraform-providers/terraform-provider-google-beta/issues/952))
* dataproc: Suppress diff for `google_dataproc_cluster` `software_config.0.image_version` to prevent permadiff when server uses more specific versions of config value ([#969](https://github.com/terraform-providers/terraform-provider-google-beta/issues/969))
* organization: Add auditConfigs to update masks for setting org and folder IAM policy (`google_organization_iam_policy`, `google_folder_iam_policy`) ([#967](https://github.com/terraform-providers/terraform-provider-google-beta/issues/967))
* storage: `google_storage_bucket` Set website metadata during read ([#925](https://github.com/terraform-providers/terraform-provider-google-beta/issues/925))

## 2.11.0 (July 16, 2019)

NOTES:
* container: We have changed the way container clusters handle cluster state, and they should now wait until the cluster is ready when creating, updating, or refreshing cluster state. This is meant to decrease the frequency of errors where Terraform is operating on a cluster that isn't ready to be operated on. If this change causes a problem, please open an issue with as much information as you can provide, especially [debug logs](https://www.terraform.io/docs/internals/debugging.html). See [terraform-provider-google #3989](https://github.com/terraform-providers/terraform-provider-google/issues/3989) for more info.

FEATURES:
* **New Resources**: `google_bigtable_instance_iam_binding`, `google_bigtable_instance_iam_member`, and `google_bigtable_instance_iam_policy` are now available. ([#923](https://github.com/terraform-providers/terraform-provider-google-beta/issues/923))
* **New Resources**: `google_sourcerepo_repository_iam_*` Add support for source repo repository IAM resources ([#914](https://github.com/terraform-providers/terraform-provider-google-beta/issues/914))

ENHANCEMENTS:
* bigquery: Added support for `external_data_configuration` to `google_bigquery_table`. ([#696](https://github.com/terraform-providers/terraform-provider-google-beta/issues/696))
* compute: Avoid getting project if no diff found for google_compute_instance_template ([#932](https://github.com/terraform-providers/terraform-provider-google-beta/issues/932))
* firestore: `google_firestore_index` `query_scope` can have `COLLECTION_GROUP` specified. ([#919](https://github.com/terraform-providers/terraform-provider-google-beta/issues/919))

BUG FIXES:
* compute: Mark instance KMS self link field kms_key_self_link as computed ([#819](https://github.com/terraform-providers/terraform-provider-google-beta/issues/819))
* compute: Allow security policy to be removed from `google_backend_service` ([#916](https://github.com/terraform-providers/terraform-provider-google-beta/issues/916))
* container: `google_container_cluster` deeper nil checks to prevent crash on empty object ([#934](https://github.com/terraform-providers/terraform-provider-google-beta/issues/934))
* container: `google_container_cluster` keep clusters in state if they are created in an error state and don't get correctly cleaned up. ([#929](https://github.com/terraform-providers/terraform-provider-google-beta/issues/929))
* container: `google_container_node_pool` Correctly set nodepool autoscaling in state when disabled in the API ([#931](https://github.com/terraform-providers/terraform-provider-google-beta/issues/931))
* container: `google_container_cluster` will now wait to act until the cluster can be operated on, respecting timeouts. ([#927](https://github.com/terraform-providers/terraform-provider-google-beta/issues/927))
* monitoring: Fix diff in `google_monitoring_uptime_check_config` on a deprecated field. ([#944](https://github.com/terraform-providers/terraform-provider-google-beta/issues/944))
* service: `google_service_networking_connection` correctly delete the connection when the resource is destroyed. ([#935](https://github.com/terraform-providers/terraform-provider-google-beta/issues/935))
* spanner: Wait for spanner databases to create before returning. Don't wait for databases to delete before returning anymore. ([#922](https://github.com/terraform-providers/terraform-provider-google-beta/issues/922))
* storage: Fixed an issue where `google_storage_transfer_job` `schedule_end_date` caused requests to fail if unset. ([#936](https://github.com/terraform-providers/terraform-provider-google-beta/issues/936))
* storage: `google_storage_object_acl` Prevent panic when using interpolated object names. ([#917](https://github.com/terraform-providers/terraform-provider-google-beta/issues/917))


## 2.10.0 (July 02, 2019)

DEPRECATIONS:
* monitoring: Deprecated non-existent fields `is_internal` and `internal_checkers` from `google_monitoring_uptime_check_config`. ([#888](https://github.com/terraform-providers/terraform-provider-google-beta/issues/888))

FEATURES:
* **New Resource**: `google_compute_project_default_network_tier` ([#882](https://github.com/terraform-providers/terraform-provider-google-beta/issues/882))
* **New Resource** `google_healthcare_dataset_iam_binding` ([#899](https://github.com/terraform-providers/terraform-provider-google-beta/pull/899))
* **New Resource** `google_healthcare_dataset_iam_member` ([8#99](https://github.com/terraform-providers/terraform-provider-google-beta/pull/899))
* **New Resource** `google_healthcare_dataset_iam_policy` ([#899](https://github.com/terraform-providers/terraform-provider-google-beta/pull/899))
* **New Resource** `google_healthcare_dicom_store_iam_binding` ([#899](https://github.com/terraform-providers/terraform-provider-google-beta/pull/899))
* **New Resource** `google_healthcare_dicom_store_iam_member` ([#899](https://github.com/terraform-providers/terraform-provider-google-beta/pull/899))
* **New Resource** `google_healthcare_dicom_store_iam_policy` ([#899](https://github.com/terraform-providers/terraform-provider-google-beta/pull/899))
* **New Resource** `google_healthcare_fhir_store_iam_binding` ([#899](https://github.com/terraform-providers/terraform-provider-google-beta/pull/899))
* **New Resource** `google_healthcare_fhir_store_iam_member` ([#899](https://github.com/terraform-providers/terraform-provider-google-beta/pull/899))
* **New Resource** `google_healthcare_fhir_store_iam_policy` ([#899](https://github.com/terraform-providers/terraform-provider-google-beta/pull/899))
* **New Resource** `google_healthcare_hl7_v2_store_iam_binding` ([#899](https://github.com/terraform-providers/terraform-provider-google-beta/pull/899))
* **New Resource** `google_healthcare_hl7_v2_store_iam_member` ([#899](https://github.com/terraform-providers/terraform-provider-google-beta/pull/899))
* **New Resource** `google_healthcare_hl7_v2_store_iam_policy` ([#899](https://github.com/terraform-providers/terraform-provider-google-beta/pull/899))

ENHANCEMENTS:
* compute: Added fields for managing network endpoint group backends in `google_compute_backend_service`, including `max_connections_per_endpoint` and `max_rate_per_endpoint` ([#854](https://github.com/terraform-providers/terraform-provider-google-beta/issues/854))
* compute: Support custom timeouts in `google_compute_instance_group_manager` and `google_compute_region_instance_group_manager` ([#909](https://github.com/terraform-providers/terraform-provider-google-beta/issues/909))
* container: `node_config.sandbox_config` is supported on GKE node pool definitions, allowing you to configure GKE Sandbox. ([#863](https://github.com/terraform-providers/terraform-provider-google-beta/issues/863))
* container: `google_container_cluster` add support for GKE resource usage ([#825](https://github.com/terraform-providers/terraform-provider-google-beta/issues/825))
* folder: `google_folder` improve error message on delete ([#878](https://github.com/terraform-providers/terraform-provider-google-beta/issues/878))
* iam: sort bindings in `google_*_iam_policy` resources to get simpler diffs ([#881](https://github.com/terraform-providers/terraform-provider-google-beta/issues/881))
* kms: `google_kms_crypto_key` now supports labels. ([#885](https://github.com/terraform-providers/terraform-provider-google-beta/issues/885))
* pubsub: `google_pubsub_topic` supports KMS keys with `kms_key_name`. ([#894](https://github.com/terraform-providers/terraform-provider-google-beta/issues/894))

BUG FIXES:
* iam: the member field in iam_* resources is now case-insensitive ([#876](https://github.com/terraform-providers/terraform-provider-google-beta/issues/876))
* servicenetworking: `google_service_networking_connection` fix update ([#871](https://github.com/terraform-providers/terraform-provider-google-beta/issues/871))

## 2.9.1 (June 21, 2019)

BUG FIXES:
* kms: fix regression when reading existing `google_kms_crypto_key` resources ([#873](https://github.com/terraform-providers/terraform-provider-google-beta/issues/873))
* storage: `google_storage_bucket` fix for crash that occurs when running plan on old buckets ([#870](https://github.com/terraform-providers/terraform-provider-google-beta/issues/870))
* storage: `google_storage_bucket` allow updating bucket_policy_only to false ([#870](https://github.com/terraform-providers/terraform-provider-google-beta/issues/870))

## 2.9.0 (June 19, 2019)

FEATURES:
* **Custom Endpoint Support**: The Google provider supports custom endpoints, allowing you to use GCP-like APIs such as emulators. See the [Provider Reference](https://www.terraform.io/docs/providers/google/provider_reference.html) for details. ([#811](https://github.com/terraform-providers/terraform-provider-google-beta/issues/811))
* **New Resource**: `google_compute_resource_policy` is now available which can be used to schedule disk snapshots. ([#1850](https://github.com/GoogleCloudPlatform/magic-modules/pull/1850))
* **New Resource**: `google_compute_external_vpn_gateway` is now available which can be used to connect to external VPN gateways. ([#833](https://github.com/terraform-providers/terraform-provider-google-beta/issues/833))
* **New Resource** Network endpoint groups (`google_compute_network_endpoint_group`) and fine-grained resource endpoints (`google_compute_network_endpoint`) are now available. ([#781](https://github.com/terraform-providers/terraform-provider-google-beta/issues/781))

ENHANCEMENTS:
* increased default timeouts for `google_compute_instance`, `google_container_cluster`, `google_dataproc_cluster`, and `google_sql_database_instance` ([#862](https://github.com/terraform-providers/terraform-provider-google-beta/issues/862))
* container: `google_container_cluster` Stop guest_accelerator from having a permadiff for accelerators with `count=0` ([#851](https://github.com/terraform-providers/terraform-provider-google-beta/issues/851))
* container: `google_container_cluster` supports `authenticator_groups_config` to allow Google Groups-based authentication. ([#669](https://github.com/terraform-providers/terraform-provider-google-beta/issues/669))
* container: `google_container_cluster` supports `enable_intranode_visibility`. ([#801](https://github.com/terraform-providers/terraform-provider-google-beta/issues/801))
* container: `google_container_cluster` supports Workload Identity to access GCP APIs in GKE applications with `workload_identity_config`. ([#824](https://github.com/terraform-providers/terraform-provider-google-beta/issues/824))
* dataproc: `google_dataproc_cluster` supports `min_cpu_platform` ([#424](https://github.com/terraform-providers/terraform-provider-google-beta/issues/424)], [[#848](https://github.com/terraform-providers/terraform-provider-google-beta/issues/848))
* dns: `google_dns_record_set`: allow importing dns record sets in any project ([#853](https://github.com/terraform-providers/terraform-provider-google-beta/issues/853))
* kms: `kms_crypto_key` supports `purpose` ([#845](https://github.com/terraform-providers/terraform-provider-google-beta/issues/845))
* storage: `google_storage_bucket` now supports enabling `bucket_policy_only` access control. ([#1878](https://github.com/GoogleCloudPlatform/magic-modules/pull/1878))
* storage: IAM resources for storage buckets (`google_storage_bucket_iam_*`) now all support import ([#835](https://github.com/terraform-providers/terraform-provider-google-beta/issues/835))
* pubsub: `google_pubsub_topic` Updates for labels are now supported ([#832](https://github.com/terraform-providers/terraform-provider-google-beta/issues/832))

BUG FIXES:
* bigquery: `google_bigquery_dataset` Relax IAM role restrictions on BQ datasets ([#857](https://github.com/terraform-providers/terraform-provider-google-beta/issues/857))
* compute: `google_project_iam` When importing resources `project` no longer needs to be set in the config post import ([#805](https://github.com/terraform-providers/terraform-provider-google-beta/issues/805))
* compute: `google_sql_user` User's can now be updated to change their password ([#810](https://github.com/terraform-providers/terraform-provider-google-beta/issues/810))
* compute: `google_compute_instance_template` Fixed issue so project can now be specified by interpolated varibles. ([#816](https://github.com/terraform-providers/terraform-provider-google-beta/issues/816))
* compute: `google_compute_instance_template` Throw error when using incompatible disk fields instead of continual plan diff ([#812](https://github.com/terraform-providers/terraform-provider-google-beta/issues/812))
* compute: `google_compute_instance_from_template` Make sure disk type is expanded to a URL ([#771](https://github.com/terraform-providers/terraform-provider-google-beta/issues/771))
* comptue: `google_compute_instance_template` Attempt to put disks in state in the same order they were specified ([#771](https://github.com/terraform-providers/terraform-provider-google-beta/issues/771))
* container: `google_container_cluster` and `google_node_pool` now retry correctly when polling for status of an operation. ([#818](https://github.com/terraform-providers/terraform-provider-google-beta/issues/818))
* container: `google_container_cluster` `istio_config.auth` will no longer permadiff on `AUTH_NONE` when an auth method other than TLS is defined. ([#834](https://github.com/terraform-providers/terraform-provider-google-beta/issues/834))
* dns: `google_dns_record_set` overrides all existing record types on create, not just NS ([#850](https://github.com/terraform-providers/terraform-provider-google-beta/issues/850))
* monitoring: `google_monitoring_notification_channel` Allow setting enabled to false ([#864](https://github.com/terraform-providers/terraform-provider-google-beta/issues/864))
* pubsub: `google_pubsub_subscription` and `google_pubsub_topic` resources can be created inside VPC service controls. ([#827](https://github.com/terraform-providers/terraform-provider-google-beta/issues/827))
* redis: `google_redis_instance` Fall back to region from `location_id` when region isn't specified ([#847](https://github.com/terraform-providers/terraform-provider-google-beta/issues/847))

## 2.8.0 (June 04, 2019)

DEPRECATIONS:
* compute: The `auto_create_routes` field on `google_compute_network_peering` has been deprecated because it is not user configurable. ([#3394](https://github.com/terraform-providers/terraform-provider-google/issues/3394))

FEATURES:
* **New Resource**: `google_compute_ha_vpn_gateway` is now available. This is an alternative to `google_compute_vpn_gateway` that can be set up to provide higher availability. ([#704](https://github.com/terraform-providers/terraform-provider-google-beta/pull/704))
* **New Datasource**: `google_compute_ssl_certificate` ([#742](https://github.com/terraform-providers/terraform-provider-google-beta/pull/742))
* **New Datasource**: `google_composer_image_versions` ([#752](https://github.com/terraform-providers/terraform-provider-google-beta/pull/752))

ENHANCEMENTS:
* app_engine: Remove restrictive `app_engine_application` location validation. ([#760](https://github.com/terraform-providers/terraform-provider-google-beta/pull/760))
* compute: `google_compute_vpn_tunnel` supports HA fields `vpn_gateway`, `vpn_gateway_interface`, `peer_gcp_gateway`, `peer_external_gateway`, `vpn_gateway_interface` ([#704](https://github.com/terraform-providers/terraform-provider-google-beta/pull/704))
* compute: `google_container_cluster` add support for vertical pod autoscaling ([#749](https://github.com/terraform-providers/terraform-provider-google-beta/issues/749))
* compute: `google_compute_router_interface` now supports specifying an `interconnect_attachment`. ([#769](https://github.com/terraform-providers/terraform-provider-google-beta/pull/769))
* compute: `google_compute_router_nat` now supports specifying a `log_config` block. ([#743](https://github.com/terraform-providers/terraform-provider-google-beta/pull/743))
* compute: `google_compute_router_nat` now supports more import formats. ([#785](https://github.com/terraform-providers/terraform-provider-google-beta/pull/785))
* compute: `google_compute_network_peering` now supports importing/exporting custom routes ([#754](https://github.com/terraform-providers/terraform-provider-google-beta/pull/754))
* compute: `google_compute_backend_service` now supports self-managed internal load balancing ([#772](https://github.com/terraform-providers/terraform-provider-google-beta/issues/772))
* compute: `google_compute_region_backend_service` now supports failover policies  ([#789](https://github.com/terraform-providers/terraform-provider-google-beta/pull/789))
* compute: Add support for INTERNAL_SELF_MANAGED backend service. Changed Resources: `google_compute_backend_service`, `google_compute_global_forwarding_rule`. ([#772](https://github.com/terraform-providers/terraform-provider-google-beta/pull/772))
* composer: Make cloud composer environment image version updateable ([#741](https://github.com/terraform-providers/terraform-provider-google-beta/pull/741))
* container: `google_container_cluster` now supports `vertical_pod_autoscaling` ([#733](https://github.com/terraform-providers/terraform-provider-google-beta/pull/733))
* container: Expose the `services_ipv4_cidr` for `container_cluster`. ([#804](https://github.com/terraform-providers/terraform-provider-google-beta/pull/804))
* dataflow: `google_dataflow_job` now supports setting machine type ([#1862](https://github.com/GoogleCloudPlatform/magic-modules/pull/1862))
* dns: `google_dns_managed_zone` now supports DNSSec ([#737](https://github.com/terraform-providers/terraform-provider-google-beta/pull/737))
* kms: `google_kms_key_ring` is now autogenerated. ([#748](https://github.com/terraform-providers/terraform-provider-google-beta/pull/748))
* pubsub: `google_pubsub_subscription` supports setting an `expiration_policy` with no `ttl`. ([#783](https://github.com/terraform-providers/terraform-provider-google-beta/pull/783))

BUG FIXES:
* binauth: `google_binary_authorization_policy` can be used with attestors in another project. ([#778](https://github.com/terraform-providers/terraform-provider-google-beta/pull/778))
* compute: allow setting firewall priority to 0 ([#755](https://github.com/terraform-providers/terraform-provider-google-beta/pull/755))
* compute: Resolved an issue where `google_compute_region_backend_service` was unable to perform a state migration. ([#775](https://github.com/terraform-providers/terraform-provider-google-beta/pull/775))
* compute: allow empty metadata.startup-script on instances ([#776](https://github.com/terraform-providers/terraform-provider-google-beta/pull/776))
* compute: Fix flattened custom patchable resources in `google_compute_network`. ([#782](https://github.com/terraform-providers/terraform-provider-google-beta/pull/782))
* compute: `google_compute_vpn_tunnel` now supports sending an empty external gateway interface id. ([#759](https://github.com/terraform-providers/terraform-provider-google-beta/pull/759))
* container: allow AUTH_NONE in istio addon_config ([#664](https://github.com/terraform-providers/terraform-provider-google-beta/pull/664))
* container: allow going from no ip_allocation_policy to a blank-equivalent one ([#774](https://github.com/terraform-providers/terraform-provider-google-beta/pull/774))
* container: `google_container_cluster` will no longer diff unnecessarily on `issue_client_certificate`. ([#788](https://github.com/terraform-providers/terraform-provider-google-beta/pull/788))
* container: `google_container_cluster` can enable client certificates on GKE `1.12+` series releases. ([#788](https://github.com/terraform-providers/terraform-provider-google-beta/pull/788))
* container: `google_container_cluster` now retries the call to remove default node pools during cluster creation ([#799](https://github.com/terraform-providers/terraform-provider-google-beta/pull/799))
* storage: Fix occasional crash when updating storage buckets ([#706](https://github.com/terraform-providers/terraform-provider-google-beta/pull/706))

## 2.7.0 (May 21, 2019)

NOTE:
* Several resources were previously undocumented on the site or changelog; they should be added to both with this release. `google_compute_backend_bucket_signed_url_key` and `google_compute_backend_service_signed_url_key` were introduced in `2.4.0`.

BACKWARDS INCOMPATIBILITIES:
* cloudfunctions: `google_cloudfunctions_function.runtime` now has an explicit default value of `nodejs6`. Users who have a different value set in the API but the value undefined in their config will see a diff. ([#697](https://github.com/terraform-providers/terraform-provider-google-beta/issues/697))

FEATURES:
* **New Resources**: `google_compute_instance_iam_binding`, `google_compute_instance_iam_member`, and `google_compute_instance_iam_policy` are now available. ([#685](https://github.com/terraform-providers/terraform-provider-google-beta/pull/685))
* **New Resources**: IAM resources for Dataproc jobs and clusters (`google_dataproc_job_iam_policy`, `google_dataproc_job_iam_member`, `google_dataproc_job_iam_binding`, `google_dataproc_cluster_iam_policy`, `google_dataproc_cluster_iam_member`, `google_dataproc_cluster_iam_binding`) are now available. [#709](https://github.com/terraform-providers/terraform-provider-google-beta/pull/709)
* **New Resources**: `google_iap_tunnel_instance_iam_binding`, `google_iap_tunnel_instance_iam_member`, and `google_iap_tunnel_instance_iam_policy` are now available. ([#687](https://github.com/terraform-providers/terraform-provider-google-beta/issues/687))

ENHANCEMENTS:
* provider: Add GCP zone to `google_client_config` datasource ([#668](https://github.com/terraform-providers/terraform-provider-google-beta/issues/668))
* compute: Add support for creating instances with CMEK ([#698](https://github.com/terraform-providers/terraform-provider-google-beta/issues/698))
* compute: Can now specify project when importing instance groups.
* compute: `google_compute_instance` now supports `shielded_instance_config` for verifiable integrity of your VM instances. ([#711](https://github.com/terraform-providers/terraform-provider-google-beta/issues/711))
* compute: `google_compute_backend_service` now supports `HTTP2` protocol (beta API feature) [#708](https://github.com/terraform-providers/terraform-provider-google-beta/pull/708)
* compute: `google_compute_instance_template` now supports `shielded_instance_config` for verifiable integrity of your VM instances. ([#711](https://github.com/terraform-providers/terraform-provider-google-beta/issues/711))
* container: use the cluster subnet to look up the node cidr block ([#722](https://github.com/terraform-providers/terraform-provider-google-beta/issues/722))

BUG FIXES:
* cloudfunctions: `google_cloudfunctions_function.runtime` now has an explicit default value of `nodejs6`. ([#697](https://github.com/terraform-providers/terraform-provider-google-beta/issues/697))
* monitoring: updating `google_monitoring_alert_policy` is more likely to succeed ([#684](https://github.com/terraform-providers/terraform-provider-google-beta/issues/684))
* kms: `google_kms_crypto_key` now (in addition to marking all crypto key versions for destruction) correctly disables auto-rotation for destroyed keys ([#705](https://github.com/terraform-providers/terraform-provider-google-beta/issues/705))
* iam: Increase IAM custom role length validation to match API. ([#728](https://github.com/terraform-providers/terraform-provider-google-beta/issues/728))

## 2.6.0 (May 07, 2019)

KNOWN ISSUES:
* cloudfunctions: `google_cloudfunctions_function`s without a `runtime` set will fail to create due to an upstream API change. You can work around this by setting an explicit `runtime` in `2.X` series releases.

DEPRECATIONS:
* monitoring: `google_monitoring_alert_policy` `labels` was deprecated, as the field was never used and it was typed incorrectly. ([#635](https://github.com/terraform-providers/terraform-provider-google-beta/issues/635))

FEATURES:
* **New Datasource**: `google_compute_node_types` for sole-tenant node types is now available. ([#614](https://github.com/terraform-providers/terraform-provider-google-beta/pull/614))
* **New Resource**: `google_compute_node_group` for sole-tenant nodes is now available. ([#643](https://github.com/terraform-providers/terraform-provider-google-beta/pull/643))
* **New Resource**: `google_compute_node_template` for sole-tenant nodes is now available. ([#614](https://github.com/terraform-providers/terraform-provider-google-beta/pull/614))
* **New Resource**: `google_firestore_index` is now available to configure composite indexes on Firestore. ([#632](https://github.com/terraform-providers/terraform-provider-google-beta/issues/632))
* **New Resource**: `google_logging_metric` is now available to configure Stackdriver logs-based metrics. ([#1702](https://github.com/GoogleCloudPlatform/magic-modules/pull/1702))
* **New Resource**: `google_compute_network_endpoint_group` ([#630](https://github.com/terraform-providers/terraform-provider-google-beta/issues/630))
* **New Resource**: `google_security_scanner_scan_config` is now available for configuring scan runs with Cloud Security Scanner. ([#641](https://github.com/terraform-providers/terraform-provider-google-beta/issues/641))

ENHANCEMENTS:
* compute: `google_compute_subnetwork` now supports `log_config` to configure flow logs' logging behaviour. ([#619](https://github.com/terraform-providers/terraform-provider-google-beta/issues/619))
* container: `google_container_cluster` now supports `database_encryption` to configure etcd encryption. ([#649](https://github.com/terraform-providers/terraform-provider-google-beta/issues/649))
* dataflow: `google_dataflow_job`'s `network` and `subnetwork` can be configured. ([#631](https://github.com/terraform-providers/terraform-provider-google-beta/issues/631))
* monitoring: `google_monitoring_alert_policy` `user_labels` support was added. ([#635](https://github.com/terraform-providers/terraform-provider-google-beta/issues/635))
* compute: `google_compute_region_backend_service` is now generated with Magic Modules, adding configurable timeouts, multiple import formats, `creation_timestamp` output. ([#645](https://github.com/terraform-providers/terraform-provider-google-beta/issues/645))
* compute: `iam_compute_subnetwork` is now GA. ([#656](https://github.com/terraform-providers/terraform-provider-google-beta/issues/656))
* pubsub: `google_pubsub_subscription` now supports setting an `expiration_policy`. ([#1703](https://github.com/GoogleCloudPlatform/magic-modules/pull/1703))

BUG FIXES:
* bigquery: `google_bigquery_table` will work with a larger range of projects id formats. ([#658](https://github.com/terraform-providers/terraform-provider-google-beta/issues/658))
* cloudfunctions: `google_cloudfunctions_fucntion` no longer restricts an outdated list of `region`s ([#659](https://github.com/terraform-providers/terraform-provider-google-beta/issues/659))
* compute: `google_compute_instance` now retries updating metadata when fingerprints are mismatched. ([#583](https://github.com/terraform-providers/terraform-provider-google-beta/issues/583))
* compute: `google_compute_instance` and `google_compute_instance_template` now support node affinities for scheduling on sole tenant nodes [[#663](https://github.com/terraform-providers/terraform-provider-google-beta/issues/663)](https://github.com/terraform-providers/terraform-provider-google-beta/pull/663)
* compute: `google_compute_managed_ssl_certificate` will no longer diff when using an absolute FQDN. ([#591](https://github.com/terraform-providers/terraform-provider-google-beta/issues/591))
* compute: `google_compute_disk` resources using `google-beta` will properly detach users at deletion instead of failing. ([#640](https://github.com/terraform-providers/terraform-provider-google-beta/issues/640))
* compute: `google_compute_subnetwork.secondary_ip_ranges` doesn't cause a diff on out of band changes, allows updating to empty list of ranges. ([#3496](https://github.com/terraform-providers/terraform-provider-google-beta/issues/3496))
* container: `google_container_cluster` setting networks / subnetworks by name works with `location`. ([#634](https://github.com/terraform-providers/terraform-provider-google-beta/issues/634))
* container: `google_container_cluster` removed an overly restrictive validation restricting `node_pool` and `remove_default_node_pool` being specified at the same time. ([#637](https://github.com/terraform-providers/terraform-provider-google-beta/issues/637))
* storage: `data_source_google_storage_bucket_object` now correctly URL encodes the slashes in a file name ([#587](https://github.com/terraform-providers/terraform-provider-google-beta/issues/587))

## 2.5.1 (April 22, 2019)

BUG FIXES:
* compute: `google_compute_backend_service` handles empty/nil `iap` block created by previous providers properly. ([#622](https://github.com/terraform-providers/terraform-provider-google-beta/issues/622))
* compute: `google_compute_backend_service` allows multiple instance types in `backends.group` again. ([#625](https://github.com/terraform-providers/terraform-provider-google-beta/issues/625))
* dns: `google_dns_managed_zone` does not permadiff when visiblity is set to default and returned as empty from API ([#624](https://github.com/terraform-providers/terraform-provider-google-beta/issues/624))
* google_projects: Datasource `google_projects` now handles paginated results from listing projects ([#626](https://github.com/terraform-providers/terraform-provider-google-beta/pull/626))
* google_project_iam: `google_project_iam_policy/member/binding` now attempts to retry for read-only operations as well as retrying read-write operations([#620](https://github.com/terraform-providers/terraform-provider-google-beta/pull/620))
* kms: `google_kms_crypto_key.rotation_period` now can be an empty string to allow for unset behavior in modules ([#627](https://github.com/terraform-providers/terraform-provider-google-beta/pull/627))

## 2.5.0 (April 18, 2019)


KNOWN ISSUES:
* compute: `google_compute_subnetwork` will fail to reorder `secondary_ip_range` values at apply time
* compute: `google_compute_subnetwork`s used with a VPC-native GKE cluster will have a diff if that cluster creates secondary ranges automatically.

BACKWARDS INCOMPATIBILITIES:
* all: This is the first release to use the 0.12 SDK required for Terraform 0.12 support. Some provider behaviour may have changed as a result of changes made by the new SDK version.
* compute: `google_compute_instance_group` will not reconcile instances recreated within the same `terraform apply` due to underlying `0.12` SDK changes in the provider. ([#616](https://github.com/terraform-providers/terraform-provider-google-beta/issues/616))
* compute: `google_compute_subnetwork` will have a diff if `secondary_ip_range` values defined in config don't exactly match real state; if so, they will need to be reconciled. ([#3432](https://github.com/terraform-providers/terraform-provider-google-beta/issues/3432))
* container: `google_container_cluster` will have a diff if `master_authorized_networks.cidr_blocks` defined in config doesn't exactly match the real state; if so, it will need to be reconciled. ([#603](https://github.com/terraform-providers/terraform-provider-google-beta/issues/603))


BUG FIXES:
* container: `google_container_cluster` catch out of band changes to `master_authorized_networks.cidr_blocks`. ([#603](https://github.com/terraform-providers/terraform-provider-google-beta/issues/603))


## 2.4.1 (April 30, 2019)

NOTES: This 2.4.1 release is a bugfix release for 2.4.0. It backports the fixes applied in the 2.5.1 release to the 2.4.0 series.

BUG FIXES:
* compute: `google_compute_backend_service` handles empty/nil `iap` block created by previous providers properly. ([#622](https://github.com/terraform-providers/terraform-provider-google-beta/issues/622))
* compute: `google_compute_backend_service` allows multiple instance types in `backends.group` again. ([#625](https://github.com/terraform-providers/terraform-provider-google-beta/issues/625))
* dns: `google_dns_managed_zone` does not permadiff when visiblity is set to default and returned as empty from API ([#624](https://github.com/terraform-providers/terraform-provider-google-beta/issues/624))

## 2.4.0 (April 15, 2019)

KNOWN ISSUES:

* compute: `google_compute_backend_service` resources created with past provider versions won't work with `2.4.0`. You can pin your provider version or manually delete them and recreate them until this is resolved. (https://github.com/terraform-providers/terraform-provider-google/issues/3441)
* dns: `google_dns_managed_zone.visibility` will cause a diff if set to `public`. Setting it to `""` (defaulting to public) will work around this. (https://github.com/terraform-providers/terraform-provider-google/issues/3435)

BACKWARDS INCOMPATIBILITIES:
* accesscontextmanager: `google_access_context_manager_service_perimeter` `unrestricted_services` field was removed based on a removal in the underlying API. ([#576](https://github.com/terraform-providers/terraform-provider-google-beta/issues/576))

FEATURES:
* **New Resource**: `google_compute_backend_bucket_signed_url_key` is now available. ([#530](https://github.com/terraform-providers/terraform-provider-google-beta/issues/530))
* **New Resource**: `google_compute_backend_service_signed_url_key` is now available. ([#577](https://github.com/terraform-providers/terraform-provider-google-beta/issues/577))
* **New Datasource**: `google_service_account_access_token` is now available. ([#575](https://github.com/terraform-providers/terraform-provider-google-beta/issues/575))

ENHANCEMENTS:
* compute: `google_compute_backend_service` is now generated with Magic Modules, adding configurable timeouts, multiple import formats, `creation_timestamp` output. ([#569](https://github.com/terraform-providers/terraform-provider-google-beta/issues/569))
* compute: `google_compute_backend_service` now supports `load_balancing_scheme` and `cdn_policy.signed_url_cache_max_age_sec`. ([#584](https://github.com/terraform-providers/terraform-provider-google-beta/issues/584))
* compute: `google_compute_network` now supports `delete_default_routes_on_create` to delete pre-created routes at network creation time. ([#592](https://github.com/terraform-providers/terraform-provider-google-beta/issues/592))
* compute: `google_compute_autoscaler` now supports `metric.single_instance_assignment` ([#580](https://github.com/terraform-providers/terraform-provider-google-beta/issues/580))
* dns: `google_dns_policy` now supports `enable_logging`. ([#573](https://github.com/terraform-providers/terraform-provider-google-beta/issues/573))
* dns: `google_dns_managed_zone` now supports `peering_config` to enable DNS Peering. ([#572](https://github.com/terraform-providers/terraform-provider-google-beta/issues/572))

BUG FIXES:
* container: `google_container_cluster` will ignore out of band changes on `node_ipv4_cidr_block`. ([#558](https://github.com/terraform-providers/terraform-provider-google-beta/issues/558))
* container: `google_container_cluster` will now reject config with both `node_pool` and `remove_default_node_pool` defined ([#600](https://github.com/terraform-providers/terraform-provider-google-beta/issues/600))
* container: `google_container_cluster` will allow >20 `cidr_blocks` in `master_authorized_networks_config`. ([#594](https://github.com/terraform-providers/terraform-provider-google-beta/issues/594))
* netblock: `data.google_netblock_ip_ranges.cidr_blocks` will better handle ipv6 input. ([#590](https://github.com/terraform-providers/terraform-provider-google-beta/issues/590))
* sql: `google_sql_database_instance` will retry reads during Terraform refreshes if it hits a rate limit. ([#579](https://github.com/terraform-providers/terraform-provider-google-beta/issues/579))

## 2.3.0 (March 26, 2019)

DEPRECATIONS:
* container: `google_container_cluster` `zone` and `region` fields are deprecated in favour of `location`, `additional_zones` in favour of `node_locations`. ([#461](https://github.com/terraform-providers/terraform-provider-google-beta/issues/461))
* container: `google_container_node_pool` `zone` and `region` fields are deprecated in favour of `location`. ([#461](https://github.com/terraform-providers/terraform-provider-google-beta/issues/461))
* container: `data.google_container_cluster` `zone` and `region` fields are deprecated in favour of `location`. ([#461](https://github.com/terraform-providers/terraform-provider-google-beta/issues/461))
* container: `google_container_engine_versions` `zone` and `region` fields are deprecated in favour of `location`. ([#461](https://github.com/terraform-providers/terraform-provider-google-beta/issues/461))

FEATURES:
* **New Datasource**: `google_*_organization_policy` Adding datasources for folder and project org policy ([#468](https://github.com/terraform-providers/terraform-provider-google-beta/issues/468))

ENHANCEMENTS:
* compute: `google_compute_disk`, `google_compute_region_disk` now support `physical_block_size_bytes` ([#526](https://github.com/terraform-providers/terraform-provider-google-beta/issues/526))
* compute: `google_compute_vpn_tunnel will properly apply labels. ([#541](https://github.com/terraform-providers/terraform-provider-google-beta/issues/541))
* container: `google_container_cluster` adds a unified `location` field for regions and zones, `node_locations` to manage extra zones for multi-zonal clusters and specific zones for regional clusters. ([#461](https://github.com/terraform-providers/terraform-provider-google-beta/issues/461))
* container: `google_container_node_pool` adds a unified `location` field for regions and zones. ([#461](https://github.com/terraform-providers/terraform-provider-google-beta/issues/461))
* container: `data.google_container_cluster` adds a unified `location` field for regions and zones. ([#461](https://github.com/terraform-providers/terraform-provider-google-beta/issues/461))
* container: `google_container_engine_versions` adds a unified `location` field for regions and zones. ([#461](https://github.com/terraform-providers/terraform-provider-google-beta/issues/461))
* dataflow: `google_dataflow_job` has support for custom service accounts with `service_account_email`. ([#527](https://github.com/terraform-providers/terraform-provider-google-beta/issues/527))
* monitoring: `google_monitoring_uptime_check` will properly recreate to perform updates. ([#485](https://github.com/terraform-providers/terraform-provider-google-beta/issues/485))
* resourcemanager: `google_*_organization_policy` Add import support for folder and project organization_policies ([#512](https://github.com/terraform-providers/terraform-provider-google-beta/issues/512))
* sql: `google_sql_ssl_cert` Allow project to be specified at resource level ([#524](https://github.com/terraform-providers/terraform-provider-google-beta/issues/524))
* storage: `google_storage_bucket` avoids calls to the compute api during import ([#529](https://github.com/terraform-providers/terraform-provider-google-beta/issues/529))
* storage: `google_storage_bucket.storage_class` supports updating. ([#548](https://github.com/terraform-providers/terraform-provider-google-beta/issues/548))
* various: Some import formats that previously failed will now work as documented. ([#542](https://github.com/terraform-providers/terraform-provider-google-beta/issues/542))

BUG FIXES:
* compute: `google_compute_disk` will properly detach instances again. ([#538](https://github.com/terraform-providers/terraform-provider-google-beta/issues/538))
* container: `google_container_cluster`, `google_container_node_pool` properly suppress new GKE `1.12` `metadata` values. ([#522](https://github.com/terraform-providers/terraform-provider-google-beta/issues/522))
* various: Only 409 concurrent operation errors will be retried, and naming conflicts will not. ([#544](https://github.com/terraform-providers/terraform-provider-google-beta/issues/544))

## 2.2.0 (March 12, 2019)

KNOWN ISSUES:

* compute: `google_compute_disk` is unable to detach instances at deletion time.

---

FEATURES:
* **New Datasource**: `data.google_projects` for retrieving a list of projects based on a filter. ([#493](https://github.com/terraform-providers/terraform-provider-google-beta/issues/493))
* **New Resource**: `google_tpu_node` for Cloud TPU Nodes ([#494](https://github.com/terraform-providers/terraform-provider-google-beta/issues/494))
* **New Resource**: `google_dns_policy` for Cloud DNS policies. ([#488](https://github.com/terraform-providers/terraform-provider-google-beta/pull/488))

ENHANCEMENTS:
* compute: `google_compute_disk` and `google_compute_region_disk` will now detach themselves from a more up to date set of users at delete time. ([#480](https://github.com/terraform-providers/terraform-provider-google-beta/issues/480))
* compute: `google_compute_network` is now generated by Magic Modules, supporting configurable timeouts and more import formats. ([#509](https://github.com/terraform-providers/terraform-provider-google-beta/issues/509))
* compute: `google_compute_firewall` will validate the maximum size of service account lists at plan time. ([#508](https://github.com/terraform-providers/terraform-provider-google-beta/issues/508))
* container: `google_container_cluster` can now disable VPC Native clusters with `ip_allocation_policy.use_ip_aliases` ([#489](https://github.com/terraform-providers/terraform-provider-google-beta/issues/489))
* container: `data.google_container_engine_versions` supports `version_prefix` to allow fuzzy version matching. Using this field, Terraform can match the latest version of a major, minor, or patch release. ([#506](https://github.com/terraform-providers/terraform-provider-google-beta/issues/506))
* pubsub: `google_pubsub_subscription` now supports configuring `message_retention_duration` and `retain_acked_messages`. ([#503](https://github.com/terraform-providers/terraform-provider-google-beta/issues/503))

BUG FIXES:
* app_engine: `google_app_engine_application` correctly outputs `gcr_domain`.  ([#479](https://github.com/terraform-providers/terraform-provider-google-beta/issues/479))
* compute: `data.google_compute_subnetwork` outputs the `self_link` field again. ([#481](https://github.com/terraform-providers/terraform-provider-google-beta/issues/481))
* compute: `google_compute_attached_disk` is now removed from state if the instance was removed. ([#497](https://github.com/terraform-providers/terraform-provider-google-beta/issues/497))
* container: `google_container_cluster` private_cluster_config now has a diff suppress to prevent a permadiff for and allows for empty `master_ipv4_cidr_block`  ([#460](https://github.com/terraform-providers/terraform-provider-google-beta/issues/460))
* container: `google_container_cluster` import behavior fixed/documented for TF-state-only fields (`remove_default_node_pool`, `min_master_version`) ([#476](https://github.com/terraform-providers/terraform-provider-google-beta/issues/476)][[#487](https://github.com/terraform-providers/terraform-provider-google-beta/issues/487)][[#495](https://github.com/terraform-providers/terraform-provider-google-beta/issues/495))
* storagetransfer: `google_storage_transfer_job` will no longer crash when accessing nil dates. ([#499](https://github.com/terraform-providers/terraform-provider-google-beta/issues/499))

## 2.1.0 (February 26, 2019)

FEATURES:
* **New Resource**: Add support for `google_compute_managed_ssl_certificate`.  ([#458](https://github.com/terraform-providers/terraform-provider-google-beta/issues/458))
* **New Datasource**: `google_client_openid_userinfo` for retrieving the `email` used to authenticate with GCP. ([#459](https://github.com/terraform-providers/terraform-provider-google-beta/issues/459))

ENHANCEMENTS:
* compute: `data.google_compute_subnetwork` can now be addressed by `self_link` as an alternative to the existing `name`/`region`/`project` fields. ([#429](https://github.com/terraform-providers/terraform-provider-google-beta/issues/429))
* dns: Support for privately visible zones is added to `google_dns_managed_zone`. ([#268](https://github.com/terraform-providers/terraform-provider-google-beta/issues/268))
* pubsub: `google_pubsub_topic` is now generated using Magic Modules, adding Open in Cloud Shell examples, configurable timeouts, and the `labels` field. ([#432](https://github.com/terraform-providers/terraform-provider-google-beta/issues/432))
* pubsub: `google_pubsub_subscription` is now generated using Magic Modules, adding Open in Cloud Shell examples, configurable timeouts, update support, and the `labels` field. ([#432](https://github.com/terraform-providers/terraform-provider-google-beta/issues/432))
* sql: `google_sql_database_instance` now provides `public_ip_address` and `private_ip_address` outputs of the first public and private IP of the instance respectively. ([#454](https://github.com/terraform-providers/terraform-provider-google-beta/issues/454))


BUG FIXES:
* sql: `google_sql_database_instance` allows the empty string to be set for `private_network`. ([#454](https://github.com/terraform-providers/terraform-provider-google-beta/issues/454))

## 2.0.0 (February 12, 2019)

BACKWARDS INCOMPATIBILITIES:
* bigtable: `google_bigtable_instance` `zone` field is no longer inferred from the provider.
* bigtable: `google_bigtable_table` now reads `family` from the table's column family in Cloud Bigtable instead of creating a new column family ([#70](https://github.com/terraform-providers/terraform-provider-google-beta/issues/70))
* bigtable: `google_bigtable_instance.cluster.num_nodes` will fail at plan time if `DEVELOPMENT` instances have `num_nodes = "0"` set explicitly. If it has been set, unset the field. ([#82](https://github.com/terraform-providers/terraform-provider-google-beta/issues/82))
* cloudbuild: `google_cloudbuild_trigger.build.step.args` is now a list instead of space separated strings. ([#308](https://github.com/terraform-providers/terraform-provider-google-beta/issues/308))
* cloudfunctions: `google_cloudfunctions_function.retry_on_failure` has been removed. Use `event_trigger.failure_policy.retry` instead. ([#75](https://github.com/terraform-providers/terraform-provider-google-beta/issues/75))
* cloudfunctions: `google_cloudfunctions_function.trigger_bucket` and `google_cloudfunctions_function.trigger_topic` have been removed. Use `event trigger` instead. ([#30](https://github.com/terraform-providers/terraform-provider-google-beta/issues/30))
* composer: `google_composer_environment.node_config.zone` is now `Required`. ([#396](https://github.com/terraform-providers/terraform-provider-google-beta/issues/396))
* compute: `google_compute_instance`, `google_compute_instance_from_template` `metadata` field is now authoritative and will remove values not explicitly set in config. [[#2208](https://github.com/terraform-providers/terraform-provider-google-beta/issues/2208)](https://github.com/terraform-providers/terraform-provider-google/pull/2208)
* compute: `google_compute_region_instance_group_manager` field `update_strategy` is now deprecated in the beta provider only. It will only function in the `google` provider, ([#76](https://github.com/terraform-providers/terraform-provider-google-beta/issues/76))
* compute: `google_compute_global_forwarding_rule` field `labels` is now removed ([#81](https://github.com/terraform-providers/terraform-provider-google-beta/issues/81))
* compute: `google_compute_project_metadata` resource is now authoritative and will remove values not explicitly set in config. [[#2205](https://github.com/terraform-providers/terraform-provider-google-beta/issues/2205)](https://github.com/terraform-providers/terraform-provider-google/pull/2205)
* compute: `google_compute_url_map` resource is now authoritative and will remove values not explicitly set in config. [[#2245](https://github.com/terraform-providers/terraform-provider-google-beta/issues/2245)](https://github.com/terraform-providers/terraform-provider-google/pull/2245)
* compute: `google_compute_snapshot.snapshot_encryption_key_raw`, `google_compute_snapshot.snapshot_encryption_key_sha256`, `google_compute_snapshot.source_disk_encryption_key_raw`, `google_compute_snapshot.source_disk_encryption_key_sha256` fields are now removed. Use `google_compute_snapshot.snapshot_encryption_key.0.raw_key`, `google_compute_snapshot.snapshot_encryption_key.0.sha256`, `google_compute_snapshot.source_disk_encryption_key.0.raw_key`, `google_compute_snapshot.source_disk_encryption_key.0.sha256` instead. ([#202](https://github.com/terraform-providers/terraform-provider-google-beta/issues/202))
* compute: `google_compute_instance_group_manager` is no longer imported by the provider-level region. Set the appropriate provider-level zone instead. ([#248](https://github.com/terraform-providers/terraform-provider-google-beta/issues/248))
* compute: `google_compute_region_instance_group_manager.update_strategy` in the `google-beta` provider has been removed. ([#189](https://github.com/terraform-providers/terraform-provider-google-beta/issues/189))
* compute: `google_compute_instance`, `google_compute_instance_template`, `google_compute_instance_from_template` have had the `network_interface.address` field removed. ([#190](https://github.com/terraform-providers/terraform-provider-google-beta/issues/190))
* compute: `google_compute_instance` has had the `network_interface.access_config.assigned_nat_ip` field removed ([#48](https://github.com/terraform-providers/terraform-provider-google-beta/issues/48))
* compute: `google_compute_disk` is no longer imported by the provider-level region. Set the appropriate provider-level zone instead. ([#249](https://github.com/terraform-providers/terraform-provider-google-beta/issues/249))
* compute: `google_compute_router_nat.subnetwork.source_ip_ranges_to_nat` is now Required inside `subnetwork` blocks. ([#281](https://github.com/terraform-providers/terraform-provider-google-beta/issues/281))
* compute: `google_compute_ssl_certificate`'s `private_key` field is no longer stored in state in cleartext; it is now SHA256 encoded. ([#400](https://github.com/terraform-providers/terraform-provider-google-beta/issues/400))
* container: `google_container_cluster` fields (`private_cluster`, `master_ipv4_cidr_block`) are removed. Use `private_cluster_config` and `private_cluster_config.master_ipv4_cidr_block` instead. ([#78](https://github.com/terraform-providers/terraform-provider-google-beta/issues/78))
* container: `google_container_node_pool`'s `name_prefix` field has been restored and is no longer deprecated. ([#2975](https://github.com/terraform-providers/terraform-provider-google-beta/issues/2975))
* sql: `google_sql_database_instance` resource is now authoritative and will remove values not explicitly set in config. [[#2203](https://github.com/terraform-providers/terraform-provider-google-beta/issues/2203)](https://github.com/terraform-providers/terraform-provider-google/pull/2203)
* endpoints: `google_endpoints_service.protoc_output` was removed. Use `google_endpoints_service.protoc_output_base64` instead. ([#79](https://github.com/terraform-providers/terraform-provider-google-beta/issues/79))
* resourcemanager: `google_project_iam_policy` is now authoritative and will remove values not explicitly set in config. Several fields were removed that made it authoritative: `authoritative`, `restore_policy`, and `disable_project`. This resource is very dangerous! Ensure you are not using the removed fields (`authoritative`, `restore_policy`, `disable_project`). ([#25](https://github.com/terraform-providers/terraform-provider-google-beta/issues/25))
* resourcemanager: Datasource `google_service_account_key.service_account_id` has been removed. Use the `name` field instead. ([#80](https://github.com/terraform-providers/terraform-provider-google-beta/issues/80))
* resourcemanager: `google_project.app_engine` has been removed. Use the `google_app_engine_application` resource instead. ([#74](https://github.com/terraform-providers/terraform-provider-google-beta/issues/74))
* resourcemanager: `google_organization_custom_role.deleted` is now an output-only attribute. Use `terraform destroy`, or remove the resource from your config instead. ([#191](https://github.com/terraform-providers/terraform-provider-google-beta/issues/191))
* resourcemanager: `google_project_custom_role.deleted` is now an output-only attribute. Use `terraform destroy`, or remove the resource from your config instead. ([#199](https://github.com/terraform-providers/terraform-provider-google-beta/issues/199))
* serviceusage: `google_project_service` will now error instead of silently disabling dependent services if `disable_dependent_services` is unset. ([#384](https://github.com/terraform-providers/terraform-provider-google-beta/issues/384))
* storage: `google_storage_object_acl.role_entity` is now authoritative and will remove values not explicitly set in config. Use `google_storage_object_access_control` for fine-grained management. ([#26](https://github.com/terraform-providers/terraform-provider-google-beta/issues/26))
* storage: `google_storage_default_object_acl.role_entity` is now authoritative and will remove values not explicitly set in config. ([#47](https://github.com/terraform-providers/terraform-provider-google-beta/issues/47))
* iam: `google_*_iam_binding` Change all IAM bindings to be authoritative ([#291](https://github.com/terraform-providers/terraform-provider-google-beta/issues/291))

FEATURES:
* **New Resource**: `google_access_context_manager_access_policy` for managing the container for an organization's access levels. ([#96](https://github.com/terraform-providers/terraform-provider-google-beta/issues/96))
* **New Resource**: `google_access_context_manager_access_level` for managing an organization's access levels. ([#149](https://github.com/terraform-providers/terraform-provider-google-beta/issues/149))
* **New Resource**: `google_access_context_manager_service_perimeter` for managing service perimeters in an access policy. ([#246](https://github.com/terraform-providers/terraform-provider-google-beta/issues/246))
* **New Resource**: `google_app_engine_firewall_rule` ([#271](https://github.com/terraform-providers/terraform-provider-google-beta/issues/271)][[#336](https://github.com/terraform-providers/terraform-provider-google-beta/issues/336))
* **New Resource**: `google_monitoring_group` ([#120](https://github.com/terraform-providers/terraform-provider-google-beta/issues/120))
* **New Resource**: `google_project_iam_audit_config` ([#265](https://github.com/terraform-providers/terraform-provider-google-beta/issues/265))
* **New Resource**: `google_storage_transfer_job` for managing recurring storage transfers with Google Cloud Storage. ([#256](https://github.com/terraform-providers/terraform-provider-google-beta/issues/256))
* **New Resource**: `google_cloud_scheduler_job` for managing the cron job scheduling service with Google Cloud Scheduler. ([#378](https://github.com/terraform-providers/terraform-provider-google-beta/issues/378))
* **New Datasource**: `google_storage_bucket_object` ([#223](https://github.com/terraform-providers/terraform-provider-google-beta/issues/223))
* **New Datasource**: `google_storage_transfer_project_service_account` data source for retrieving the Storage Transfer service account for a project ([#247](https://github.com/terraform-providers/terraform-provider-google-beta/issues/247))
* **New Datasource**: `google_kms_crypto_key` ([#359](https://github.com/terraform-providers/terraform-provider-google-beta/issues/359))
* **New Datasource**: `google_kms_key_ring` ([#359](https://github.com/terraform-providers/terraform-provider-google-beta/issues/359))

ENHANCEMENTS:
* provider: Add `access_token` config option to allow Terraform to authenticate using short-lived Google OAuth 2.0 access token ([#330](https://github.com/terraform-providers/terraform-provider-google-beta/issues/330))
* bigquery: Add new locations `europe-west2` and `australia-southeast1` to valid location set for `google_bigquery_dataset` ([#41](https://github.com/terraform-providers/terraform-provider-google-beta/issues/41))
* bigquery: Add `default_partition_expiration_ms` field to `google_bigquery_dataset` resource. ([#127](https://github.com/terraform-providers/terraform-provider-google-beta/issues/127))
* bigquery: Add `delete_contents_on_destroy` field to `google_bigquery_dataset` resource. ([#413](https://github.com/terraform-providers/terraform-provider-google-beta/issues/413))
* bigquery: Add `time_partitioning.require_partition_filter` to `google_bigquery_table` resource. ([#324](https://github.com/terraform-providers/terraform-provider-google-beta/issues/324))
* bigquery: Allow more BigQuery regions ([#269](https://github.com/terraform-providers/terraform-provider-google-beta/issues/269))
* bigtable: Add `column_family` at create time to `google_bigtable_table`. [[#2228](https://github.com/terraform-providers/terraform-provider-google-beta/issues/2228)](https://github.com/terraform
