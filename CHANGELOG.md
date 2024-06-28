## 5.36.0 (Unreleased)

## 5.35.0 (June 24, 2024)

FEATURES:
* **New Data Source:** `google_artifact_registry_docker_image` ([#7544](https://github.com/hashicorp/terraform-provider-google-beta/pull/7544))
* **New Data Source:** `google_composer_user_workloads_config_map` ([#7519](https://github.com/hashicorp/terraform-provider-google-beta/pull/7519))
* **New Resource:** `google_service_networking_vpc_service_controls` ([#7545](https://github.com/hashicorp/terraform-provider-google-beta/pull/7545))

IMPROVEMENTS:
* bigquery: added `resource_tags` field to `google_bigquery_dataset` resource ([#7549](https://github.com/hashicorp/terraform-provider-google-beta/pull/7549))
* billingbudget: added `enable_project_level_recipients` field to `google_billing_budget` resource ([#7539](https://github.com/hashicorp/terraform-provider-google-beta/pull/7539))
* cloudrunv2: added fields `start_execution_token` and `run_execution_token` to resource `google_cloud_run_v2_job` ([#7525](https://github.com/hashicorp/terraform-provider-google-beta/pull/7525))
* compute: added `action_token_site_keys` and `session_token_site_keys` fields to `google_compute_security_policy` and `google_compute_security_policy_rule` resources ([#7520](https://github.com/hashicorp/terraform-provider-google-beta/pull/7520))
* dataprocmetastore: added `autoscaling_config` field to `google_dataproc_metastore_service` resource ([#7528](https://github.com/hashicorp/terraform-provider-google-beta/pull/7528))
* gkehub2: added `ENTERPRISE` option to `security_posture_config` field on `google_gke_hub_fleet` resource ([#7541](https://github.com/hashicorp/terraform-provider-google-beta/pull/7541))
* pubsub: added `bigquery_config.service_account_email` field to `google_pubsub_subscription` resource ([#7543](https://github.com/hashicorp/terraform-provider-google-beta/pull/7543))
* redis: added `maintenance_version` field to `google_redis_instance` ([#7527](https://github.com/hashicorp/terraform-provider-google-beta/pull/7527))
* storage: changed update behavior in `google_storage_bucket_object` to no longer delete to avoid object deletion on content update ([#7564](https://github.com/hashicorp/terraform-provider-google-beta/pull/7564))
* sql: added support for more MySQL values in `type` field of `google_sql_user` resource ([#7548](https://github.com/hashicorp/terraform-provider-google-beta/pull/7548))
* sql: increased timeouts on `google_sql_database_instance` to 90m to account for longer-running actions such as creation through cloning ([#7553](https://github.com/hashicorp/terraform-provider-google-beta/pull/7553))
* workbench: added update support to `gce_setup.boot_disk` and `gce_setup.data_disks` fields in `google_workbench_instance` resource ([#7566](https://github.com/hashicorp/terraform-provider-google-beta/pull/7566))

BUG FIXES:
* compute: updated `google_compute_instance` to force reboot if `min_node_cpus` is updated ([#7524](https://github.com/hashicorp/terraform-provider-google-beta/pull/7524))
* compute: fixed `description` field in `google_compute_firewall` to support empty/null values on update ([#7563](https://github.com/hashicorp/terraform-provider-google-beta/pull/7563))
* compute: fixed perma-diff on `google_compute_disk` for Ubuntu amd64 canonical LTS images ([#7522](https://github.com/hashicorp/terraform-provider-google-beta/pull/7522))
* storage: fixed lowercased `custom_placement_config` values in `google_storage_bucket` causing perma-destroy ([#7551](https://github.com/hashicorp/terraform-provider-google-beta/pull/7551))
* workbench: fixed issue where instance was not starting after an update in `google_workbench_instance` resource ([#7557](https://github.com/hashicorp/terraform-provider-google-beta/pull/7557))
* workbench: fixed perma-diff caused by empty `accelerator_configs` in `google_workbench_instance` resource ([#7557](https://github.com/hashicorp/terraform-provider-google-beta/pull/7557))

## 5.34.0 (June 17, 2024)

NOTES:
* compute: Updated field description of `connection_draining_timeout_sec`, `balancing_mode` and `outlier_detection` in `google_compute_region_backend_service` and `google_compute_backend_service`  to inform that default values will be changed in 6.0.0 ([#7513](https://github.com/hashicorp/terraform-provider-google-beta/pull/7513))

DEPRECATIONS:
* vertexai: deprecated beta field `embedding_management` for `google_vertex_ai_feature_online_store` resource ([#7473](https://github.com/hashicorp/terraform-provider-google-beta/pull/7473))

FEATURES:
* **New Data Source:** `google_composer_user_workloads_config_map` (beta) ([#7519](https://github.com/hashicorp/terraform-provider-google-beta/pull/7519))
* **New Data Source:** `google_composer_user_workloads_secret` (beta) ([#7474](https://github.com/hashicorp/terraform-provider-google-beta/pull/7474))
* **New Resource:** `google_composer_user_workloads_config_map` (beta) ([#7497](https://github.com/hashicorp/terraform-provider-google-beta/pull/7497))
* **New Resource:** `google_managed_kafka_cluster` (beta) ([#7477](https://github.com/hashicorp/terraform-provider-google-beta/pull/7477))
* **New Resource:** `google_managed_kafka_topic` (beta) ([#7503](https://github.com/hashicorp/terraform-provider-google-beta/pull/7503))
* **New Resource:** `google_netapp_backup` ([#7479](https://github.com/hashicorp/terraform-provider-google-beta/pull/7479))
* **New Resource:** `google_network_services_service_lb_policies` ([#7463](https://github.com/hashicorp/terraform-provider-google-beta/pull/7463))
* **New Resource:** `google_scc_management_folder_security_health_analytics_custom_module` ([#7483](https://github.com/hashicorp/terraform-provider-google-beta/pull/7483))
* **New Resource:** `google_scc_management_organization_project_security_health_analytics_custom_module` ([#7489](https://github.com/hashicorp/terraform-provider-google-beta/pull/7489))
* **New Resource:** `google_scc_management_organization_security_health_analytics_custom_module` ([#7493](https://github.com/hashicorp/terraform-provider-google-beta/pull/7493))

IMPROVEMENTS:
* alloydb: changed the resource `google_alloydb_instance` to be created directly with public IP enabled instead of creating the resource with public IP disabled and then enabling it ([#7469](https://github.com/hashicorp/terraform-provider-google-beta/pull/7469))
* bigtable: added `automated_backup_configuration` field to `google_bigtable_table` resource ([#7468](https://github.com/hashicorp/terraform-provider-google-beta/pull/7468))
* cloudbuildv2: added support for connecting to Bitbucket Data Center and Bitbucket Cloud with the `bitbucket_data_center_config` and `bitbucket_cloud_config` fields in `google_cloudbuildv2_connection` ([#7494](https://github.com/hashicorp/terraform-provider-google-beta/pull/7494))
* compute: added support for Port Mapping NEG and endpoint. New NEG type is a regional type of `GCE_VM_IP_PORTMAP` that requires endpoints with `instance`, `port` and `client_destination_port` (beta) ([#7471](https://github.com/hashicorp/terraform-provider-google-beta/pull/7471))
* compute: added update support to `ssl_policy` field in `google_compute_region_target_https_proxy` resource ([#7484](https://github.com/hashicorp/terraform-provider-google-beta/pull/7484))
* compute: removed enum validation on `guest_os_features.type` in `google_compute_disk` to allow for new features to be used without provider update ([#7465](https://github.com/hashicorp/terraform-provider-google-beta/pull/7465))
* compute: updated documentation of google_compute_target_https_proxy and google_compute_region_target_https_proxy ([#7481](https://github.com/hashicorp/terraform-provider-google-beta/pull/7481))
* container: added support for `security_posture_config.mode` value "ENTERPRISE" in `resource_container_cluster` ([#7467](https://github.com/hashicorp/terraform-provider-google-beta/pull/7467))
* discoveryengine: added `document_processing_config` field to `google_discovery_engine_data_store` resource ([#7475](https://github.com/hashicorp/terraform-provider-google-beta/pull/7475))
* edgecontainer: added `maintenance_exclusions` field to `google_edgecontainer_cluster` resource ([#7490](https://github.com/hashicorp/terraform-provider-google-beta/pull/7490))
* gkehub: added `prevent_drift` field to ConfigManagement `fleet_default_member_config` ([#7464](https://github.com/hashicorp/terraform-provider-google-beta/pull/7464))
* netapp: added `administrators` field to `google_netapp_active_directory` resource ([#7466](https://github.com/hashicorp/terraform-provider-google-beta/pull/7466))
* vertexai: promoted `optimized` field to GA for `google_vertex_ai_feature_online_store` resource ([#7473](https://github.com/hashicorp/terraform-provider-google-beta/pull/7473))
* workbench: updated the metadata keys managed by the backend. ([#7488](https://github.com/hashicorp/terraform-provider-google-beta/pull/7488))

BUG FIXES:
* compute: fixed an issue where `google_compute_instance_group_manager` with a pending operation was incorrectly removed due to the operation no longer being present in the backend ([#7498](https://github.com/hashicorp/terraform-provider-google-beta/pull/7498))
* compute: fixed issue where users could not create `google_compute_security_policy` resources with `layer_7_ddos_defense_config` explicitly disabled ([#7470](https://github.com/hashicorp/terraform-provider-google-beta/pull/7470))
* workbench: fixed a bug in the `google_workbench_instance` resource where specifying a network in some scenarios would cause instance creation to fail ([#7518](https://github.com/hashicorp/terraform-provider-google-beta/pull/7518))

## 5.33.0 (June 10, 2024)

DEPRECATIONS:
* healthcare: deprecated `notification_config` deprecated `notification_config` in `google_healthcare_fhir_store` resource. Use `notification_configs` instead. ([#7450](https://github.com/hashicorp/terraform-provider-google-beta/pull/7450))

FEATURES:
* **New Data Source:** `google_compute_security_policy` ([#7453](https://github.com/hashicorp/terraform-provider-google-beta/pull/7453))
* **New Resource:** `google_compute_project_cloud_armor_tier` ([#7456](https://github.com/hashicorp/terraform-provider-google-beta/pull/7456))
* **New Resource:** `google_network_services_service_lb_policies` ([#7463](https://github.com/hashicorp/terraform-provider-google-beta/pull/7463))
* **New Resource:** `google_scc_management_organization_event_threat_detection_custom_module` ([#7454](https://github.com/hashicorp/terraform-provider-google-beta/pull/7454))
* **New Resource:** `google_spanner_instance_config` ([#7459](https://github.com/hashicorp/terraform-provider-google-beta/pull/7459))

IMPROVEMENTS:
* appengine: added `flexible_runtime_settings` field to `google_app_engine_flexible_app_version` resource ([#7462](https://github.com/hashicorp/terraform-provider-google-beta/pull/7462))
* bigtable: added `force_destroy` field to `google_bigtable_instance` resource. This will force delete any backups present in the instance and allow the instance to be deleted. ([#7441](https://github.com/hashicorp/terraform-provider-google-beta/pull/7441))
* clouddeploy: added `execution_configs.verbose` field to `google_clouddeploy_target` resource ([#7442](https://github.com/hashicorp/terraform-provider-google-beta/pull/7442))
* compute: added `partner_metadata` field to `google_compute_instance_template` resource ([#7449](https://github.com/hashicorp/terraform-provider-google-beta/pull/7449))
* compute: added `partner_metadata` field to `google_compute_instance` resource ([#7449](https://github.com/hashicorp/terraform-provider-google-beta/pull/7449))
* compute: added `partner_metadata` field to `google_compute_regional_instance_template` resource ([#7449](https://github.com/hashicorp/terraform-provider-google-beta/pull/7449))
* compute: added `standby_policy`, `target_suspended_size` and  `target_stopped_size`  fields to `google_compute_instance_group_manager` and `google_compute_region_instance_group_manager` resources ([#7436](https://github.com/hashicorp/terraform-provider-google-beta/pull/7436))
* compute: added `storage_pool` field to `google_compute_disk` resource ([#7434](https://github.com/hashicorp/terraform-provider-google-beta/pull/7434))
* container: added `secret_manager_config` field to `google_container_cluster` resource ([#7448](https://github.com/hashicorp/terraform-provider-google-beta/pull/7448))
* dlp: added `secrets_discovery_target`, `cloud_sql_target.filter.database_resource_reference`, and `big_query_target.filter.table_reference` fields to `google_data_loss_prevention_discovery_config` resource ([#7461](https://github.com/hashicorp/terraform-provider-google-beta/pull/7461))
* gkebackup: added `backup_schedule.backup_config.permissive_mode` field to `google_gke_backup_backup_plan` resource ([#7430](https://github.com/hashicorp/terraform-provider-google-beta/pull/7430))
* gkebackup: added `restore_config.restore_order` field to `google_gke_backup_restore_plan` resource ([#7430](https://github.com/hashicorp/terraform-provider-google-beta/pull/7430))
* gkebackup: added `restore_config.volume_data_restore_policy_bindings` field to `google_gke_backup_restore_plan` resource ([#7430](https://github.com/hashicorp/terraform-provider-google-beta/pull/7430))
* gkebackup: added new enum values `MERGE_SKIP_ON_CONFLICT`, `MERGE_REPLACE_VOLUME_ON_CONFLICT` and `MERGE_REPLACE_ON_CONFLICT` to field `restore_config.namespaced_resource_restore_mode` in `google_gke_backup_restore_plan` resource ([#7430](https://github.com/hashicorp/terraform-provider-google-beta/pull/7430))
* healthcare: added `notification_config.send_for_bulk_import` field to `google_healthcare_dicom_store` ([#7457](https://github.com/hashicorp/terraform-provider-google-beta/pull/7457))
* integrationconnectors: added `endpoint_global_access` field to `google_integration_connectors_endpoint_attachment` resource ([#7443](https://github.com/hashicorp/terraform-provider-google-beta/pull/7443))
* netapp: added `backup_config` field to `google_netapp_volume` resource ([#7439](https://github.com/hashicorp/terraform-provider-google-beta/pull/7439))
* redis: added `zone_distribution_config` field to `google_redis_cluster` resource ([#7451](https://github.com/hashicorp/terraform-provider-google-beta/pull/7451))
* resourcemanager: added support for `range_type = "default-domains-netblocks"` in `google_netblock_ip_ranges` data source ([#7440](https://github.com/hashicorp/terraform-provider-google-beta/pull/7440))
* secretmanager: added support for IAM conditions in `google_secret_manager_secret_iam_*` resources ([#7444](https://github.com/hashicorp/terraform-provider-google-beta/pull/7444))
* workstations: added `boot_disk_size_gb`, `enable_nested_virtualization`, and `pool_size` to `host.gce_instance.boost_configs` in `google_workstations_workstation_config` resource ([#7452](https://github.com/hashicorp/terraform-provider-google-beta/pull/7452))

BUG FIXES:
* container: fixed `google_container_node_pool` crash if `node_config.secondary_boot_disks.mode` is not set ([#7460](https://github.com/hashicorp/terraform-provider-google-beta/pull/7460))
* dlp: removed `required` on `inspect_config.limits.max_findings_per_info_type.info_type` field to allow the use of default limit by not setting this field in `google_data_loss_prevention_inspect_template` resource ([#7438](https://github.com/hashicorp/terraform-provider-google-beta/pull/7438))
* provider: fixed application default credential and access token authorization when `universe_domain` is set ([#7433](https://github.com/hashicorp/terraform-provider-google-beta/pull/7433))

## 5.32.0 (June 3, 2024)

NOTES:
* privateca: converted `google_privateca_certificate_template` to now use the MMv1 engine instead of DCL ([#7409](https://github.com/hashicorp/terraform-provider-google-beta/pull/7409))

FEATURES:
* **New Resource:** `google_dataplex_entry_type` ([#7412](https://github.com/hashicorp/terraform-provider-google-beta/pull/7412))
* **New Resource:** `google_logging_log_view_iam_binding` ([#7420](https://github.com/hashicorp/terraform-provider-google-beta/pull/7420))
* **New Resource:** `google_logging_log_view_iam_member` ([#7420](https://github.com/hashicorp/terraform-provider-google-beta/pull/7420))
* **New Resource:** `google_logging_log_view_iam_policy` ([#7420](https://github.com/hashicorp/terraform-provider-google-beta/pull/7420))

IMPROVEMENTS:
* alloydb: added `psc_config` field to `google_alloydb_cluster` resource ([#7429](https://github.com/hashicorp/terraform-provider-google-beta/pull/7429))
* alloydb: added `psc_instance_config` field to `google_alloydb_instance` resource ([#7429](https://github.com/hashicorp/terraform-provider-google-beta/pull/7429))
* cloudrunv2: added `default_uri_disabled` field to `google_cloud_run_v2_service` resourceto ([#7422](https://github.com/hashicorp/terraform-provider-google-beta/pull/7422))
* compute: added `NONE` to acceptable options for `update_policy.minimal_action` field in `google_compute_instance_group_manager` resource ([#7417](https://github.com/hashicorp/terraform-provider-google-beta/pull/7417))
* sql: updated support for a new value `week5` in field `setting.maintenance_window.update_track` in `google_sql_database_instance` resource ([#7408](https://github.com/hashicorp/terraform-provider-google-beta/pull/7408))

BUG FIXES:
* cloudrunv2: added validation for `timeout` field to `google_cloud_run_v2_job`, `google_cloud_run_v2_service` resources ([#7426](https://github.com/hashicorp/terraform-provider-google-beta/pull/7426))
* compute: fixed permadiff in ordering of `advertised_ip_ranges.range` field on `google_compute_router` resource ([#7411](https://github.com/hashicorp/terraform-provider-google-beta/pull/7411))
* iam: added a 10 second sleep when creating `google_service_account` resource ([#7427](https://github.com/hashicorp/terraform-provider-google-beta/pull/7427))
* storage: fixed `google_storage_bucket.lifecycle_rule.condition` block fields  `days_since_noncurrent_time` and `days_since_custom_time`  and `num_newer_versions` were not working for 0 value. ([#7414](https://github.com/hashicorp/terraform-provider-google-beta/pull/7414))

## 5.31.0 (May 28, 2024)

FEATURES:
* **New Data Source:** `google_compute_subnetworks` ([#7371](https://github.com/hashicorp/terraform-provider-google-beta/pull/7371))
* **New Resource:** `google_dataplex_aspect_type` ([#7397](https://github.com/hashicorp/terraform-provider-google-beta/pull/7397))
* **New Resource:** `google_dataplex_entry_group` ([#7389](https://github.com/hashicorp/terraform-provider-google-beta/pull/7389))
* **New Resource:** `google_kms_autokey_config` ([#7385](https://github.com/hashicorp/terraform-provider-google-beta/pull/7385))
* **New Resource:** `google_kms_key_handle` ([#7385](https://github.com/hashicorp/terraform-provider-google-beta/pull/7385))
* **New Resource:** `google_network_services_lb_route_extension` ([#7394](https://github.com/hashicorp/terraform-provider-google-beta/pull/7394))

IMPROVEMENTS:
* appengine: added field `instance_ip_mode` to resource `google_app_engine_flexible_app_version` resource (beta) ([#7377](https://github.com/hashicorp/terraform-provider-google-beta/pull/7377))
* bigquery: added `external_data_configuration.bigtable_options` to `google_bigquery_table` ([#7387](https://github.com/hashicorp/terraform-provider-google-beta/pull/7387))
* cloudrun: added support for `nfs` to `google_cloudrun_service` (beta). ([#7381](https://github.com/hashicorp/terraform-provider-google-beta/pull/7381))
* composer: added support for importing `google_composer_user_workloads_secret` via the "{{environment}}/{{name}}" format. ([#7390](https://github.com/hashicorp/terraform-provider-google-beta/pull/7390))
* composer: improved timeouts for `google_composer_user_workloads_secret`. ([#7390](https://github.com/hashicorp/terraform-provider-google-beta/pull/7390))
* compute: added `TLS_JA3_FINGERPRINT` and `USER_IP` options in field `rate_limit_options.enforce_on_key` to `google_compute_security_policy` resource ([#7376](https://github.com/hashicorp/terraform-provider-google-beta/pull/7376))
* compute: added 'rateLimitOptions' field to 'google_compute_security_policy_rule' resource ([#7376](https://github.com/hashicorp/terraform-provider-google-beta/pull/7376))
* compute: changed `google_compute_region_ssl_policy`'s `region` field to optional and allow to be inferred from environment ([#7384](https://github.com/hashicorp/terraform-provider-google-beta/pull/7384))
* compute: added `on_instance_stop_action` field to `google_compute_instance`, `google_compute_instance_template`, and `google_compute_instance_from_machine_image ` resources (beta) ([#7392](https://github.com/hashicorp/terraform-provider-google-beta/pull/7392))
* compute: added `subnet_length` field to `google_compute_interconnect_attachment` resource ([#7388](https://github.com/hashicorp/terraform-provider-google-beta/pull/7388))
* container: added `containerd_config` field and subfields to `google_container_cluster` and `google_container_node_pool` resources, to allow those resources to access private image registries. ([#7372](https://github.com/hashicorp/terraform-provider-google-beta/pull/7372))
* container: allowed both `enable_autopilot` and `workload_identity_config` to be set in `google_container_cluster` resource. ([#7375](https://github.com/hashicorp/terraform-provider-google-beta/pull/7375))
* datastream: added `create_without_validation` field to `google_datastream_connection_profile`, `google_datastream_private_connection` and `google_datastream_stream` resources ([#7382](https://github.com/hashicorp/terraform-provider-google-beta/pull/7382))
* network-security: added `trust_config`, `min_tls_version`, `tls_feature_profile` and `custom_tls_features` fields to `google_network_security_tls_inspection_policy` resource ([#7368](https://github.com/hashicorp/terraform-provider-google-beta/pull/7368))
* networkservices: made field `load_balancing_scheme` immutable in resource `google_network_services_lb_traffic_extension`, as in-place updating is always failing ([#7394](https://github.com/hashicorp/terraform-provider-google-beta/pull/7394))
* networkservices: made required fields `extension_chains.extensions.authority ` and `extension_chains.extensions.timeout` optional in resource `google_network_services_lb_traffic_extension` ([#7394](https://github.com/hashicorp/terraform-provider-google-beta/pull/7394))
* networkservices: removed unsupported load balancing scheme `LOAD_BALANCING_SCHEME_UNSPECIFIED` from the field `load_balancing_scheme` in resource `google_network_services_lb_traffic_extension` ([#7394](https://github.com/hashicorp/terraform-provider-google-beta/pull/7394))
* pubsub: added `cloud_storage_config.filename_datetime_format` field to `google_pubsub_subscription` resource ([#7386](https://github.com/hashicorp/terraform-provider-google-beta/pull/7386))
* tpu: added `type` of `accelerator_config` to `google_tpu_v2_vm` resource ([#7369](https://github.com/hashicorp/terraform-provider-google-beta/pull/7369))

BUG FIXES:
* monitoring: fixed a permadiff with `monitored_resource.labels` property in the `google_monitoring_uptime_check_config` resource ([#7380](https://github.com/hashicorp/terraform-provider-google-beta/pull/7380))
* storage: fixed a bug where field `autoclass` block is generating permadiff whenever the block is removed from the config  in `google_storage_bucket` resource ([#7395](https://github.com/hashicorp/terraform-provider-google-beta/pull/7395))
* storagetransfer: fixed a permadiff with `transfer_spec.0.aws_s3_data_source.0.aws_access_key` `resource_storage_transfer_job` ([#7391](https://github.com/hashicorp/terraform-provider-google-beta/pull/7391))

## 5.30.0 (May 20, 2024)

FEATURES:
* **New Data Source:** `google_cloud_asset_resources_search_all` ([#7361](https://github.com/hashicorp/terraform-provider-google-beta/pull/7361))
* **New Resource:** `google_compute_interconnect` ([#7338](https://github.com/hashicorp/terraform-provider-google-beta/pull/7338))
* **New Resource:** `google_network_services_lb_traffic_extension` ([#7367](https://github.com/hashicorp/terraform-provider-google-beta/pull/7367))

IMPROVEMENTS:
* compute:  added `kms_key_name` field to the `google_bigquery_connection` resource ([#7335](https://github.com/hashicorp/terraform-provider-google-beta/pull/7335))
* compute: added `match.expr.expression` field to `google_compute_region_security_policy_rule` resource ([#7330](https://github.com/hashicorp/terraform-provider-google-beta/pull/7330))
* compute: added `auto_network_tier` field to `google_compute_router_nat` resource ([#7333](https://github.com/hashicorp/terraform-provider-google-beta/pull/7333))
* container: added `KUBELET` and `CADVISOR` options to `monitoring_config.enable_components` in `google_container_cluster` resource ([#7351](https://github.com/hashicorp/terraform-provider-google-beta/pull/7351))
* dataproc: added `local_ssd_interface` to `google_dataproc_cluster` resource ([#7366](https://github.com/hashicorp/terraform-provider-google-beta/pull/7366))
* datastream: added `sql_server_profile` to `google_datastream_connection_profile` resource ([#7339](https://github.com/hashicorp/terraform-provider-google-beta/pull/7339))
* dlp: added `cloud_sql_target` field to `google_data_loss_prevention_discovery_config` resource ([#7337](https://github.com/hashicorp/terraform-provider-google-beta/pull/7337))
* netapp: added `FLEX` value to field `service_level` in `google_netapp_storage_pool` resource ([#7350](https://github.com/hashicorp/terraform-provider-google-beta/pull/7350))
* networksecurity: added `trust_config`, `min_tls_version`, `tls_feature_profile` and `custom_tls_features` fields to `google_network_security_tls_inspection_policy` resource ([#7368](https://github.com/hashicorp/terraform-provider-google-beta/pull/7368))
* networkservices: supported in-place update for `gateway_security_policy` and `certificate_urls` fields in `google_network_services_gateway` resource ([#7348](https://github.com/hashicorp/terraform-provider-google-beta/pull/7348))

BUG FIXES:
* compute: fixed a perma-diff on `machine_type` field in `google_compute_instance` resource ([#7345](https://github.com/hashicorp/terraform-provider-google-beta/pull/7345))
* compute: fixed a perma-diff on `type` field in `google_compute_disk` resource ([#7345](https://github.com/hashicorp/terraform-provider-google-beta/pull/7345))
* storage: fixed update issue for `lifecycle_rule.condition.custom_time_before` and `lifecycle_rule.condition.noncurrent_time_before` in `google_storage_bucket` resource ([#7360](https://github.com/hashicorp/terraform-provider-google-beta/pull/7360))

## 5.29.1 (May 14, 2024)

BREAKING CHANGES:
* compute: removed `secondary_ip_range.reserved_internal_range` field from `google_compute_subnetwork` ([7363](https://github.com/hashicorp/terraform-provider-google-beta/pull/7363))

## 5.29.0 (May 13, 2024)

NOTES:
* compute: added documentation for md5_authentication_key field in google_compute_router_peer resource. The field was introduced in [v5.12.0](https://github.com/hashicorp/terraform-provider-google-beta/releases/tag/v5.12.0), but documentation was unintentionally omitted at that time. ([#7306](https://github.com/hashicorp/terraform-provider-google-beta/pull/7306))

FEATURES:
* **New Resource:** `google_bigtable_authorized_view` ([#7310](https://github.com/hashicorp/terraform-provider-google-beta/pull/7310))
* **New Resource:** `google_integration_connectors_managed_zone` ([#7320](https://github.com/hashicorp/terraform-provider-google-beta/pull/7320))
* **New Resource:** `google_network_connectivity_regional_endpoint` ([#7313](https://github.com/hashicorp/terraform-provider-google-beta/pull/7313))

IMPROVEMENTS:
* clouddeploy: added `custom_target` field to  `google_clouddeploy_target` resource ([#7309](https://github.com/hashicorp/terraform-provider-google-beta/pull/7309))
* clouddeploy: added `google_cloud_build_repo` to `custom_target_type` resource ([#7325](https://github.com/hashicorp/terraform-provider-google-beta/pull/7325))
* compute: added `preconfigured_waf_config` field to `google_compute_region_security_policy_rule` resource; ([#7324](https://github.com/hashicorp/terraform-provider-google-beta/pull/7324))
* compute: added `rate_limit_options` field to 'google_compute_region_security_policy_rule' resource; ([#7324](https://github.com/hashicorp/terraform-provider-google-beta/pull/7324))
* compute: added `security_profile_group`, `tls_inspect` to `google_compute_firewall_policy_rule` ([#7309](https://github.com/hashicorp/terraform-provider-google-beta/pull/7309))
* compute: added `security_profile_group`, `tls_inspect` to `google_compute_network_firewall_policy_rule` ([#7309](https://github.com/hashicorp/terraform-provider-google-beta/pull/7309))
* compute: added fields `reserved_internal_range` and `secondary_ip_ranges.reserved_internal_range` to `google_compute_subnetwork` resource ([#7318](https://github.com/hashicorp/terraform-provider-google-beta/pull/7318))
* container: added `dns_config.additive_vpc_scope_dns_domain` field to `google_container_cluster` resource ([#7321](https://github.com/hashicorp/terraform-provider-google-beta/pull/7321))
* container: added `enable_nested_virtualization` field to `google_container_node_pool` and `google_container_cluster` resource. ([#7314](https://github.com/hashicorp/terraform-provider-google-beta/pull/7314))
* iam: added `extra_attributes_oauth2_client` field to `google_iam_workforce_pool_provider` resource ([#7319](https://github.com/hashicorp/terraform-provider-google-beta/pull/7319))
* privateca: added `maximum_lifetime` field to  `google_privateca_certificate_template` resource ([#7309](https://github.com/hashicorp/terraform-provider-google-beta/pull/7309))

BUG FIXES:
* bigquery: added `allow_resource_tags_on_deletion` to `google_bigquery_table` to allow deletion of table when it still has associated resource tags ([#7327](https://github.com/hashicorp/terraform-provider-google-beta/pull/7327))

## 5.28.0 (May 6, 2024)

DEPRECATIONS:
* integrations: deprecated `create_sample_workflows` and `provision_gmek` fields in `google_integrations_client`. ([#7285](https://github.com/hashicorp/terraform-provider-google-beta/pull/7285))

FEATURES:
* **New Data Source:** `google_storage_buckets` ([#7291](https://github.com/hashicorp/terraform-provider-google-beta/pull/7291))
* **New Resource:** `google_compute_security_policy_rule` ([#7282](https://github.com/hashicorp/terraform-provider-google-beta/pull/7282))
* **New Resource:** `google_privileged_access_manager_entitlement` ([#7283](https://github.com/hashicorp/terraform-provider-google-beta/pull/7283))

IMPROVEMENTS:
* alloydb: added `maintenance_update_policy` field to `google_alloydb_cluster` resource ([#7288](https://github.com/hashicorp/terraform-provider-google-beta/pull/7288))
* container: added `node_config.secondary_boot_disks` field to `google_container_node_pool` ([#7292](https://github.com/hashicorp/terraform-provider-google-beta/pull/7292))
* integrations: added `create_sample_integrations` field to `google_integrations_client`, replacing deprecated field `create_sample_workflows`. ([#7285](https://github.com/hashicorp/terraform-provider-google-beta/pull/7285))
* redis: added `redis_configs` field to `google_redis_cluster` resource ([#7289](https://github.com/hashicorp/terraform-provider-google-beta/pull/7289))

BUG FIXES:
* dns: fixed bug where the deletion of `google_dns_managed_zone` resources was blocked by any associated SOA-type `google_dns_record_set` resources ([#7305](https://github.com/hashicorp/terraform-provider-google-beta/pull/7305))
* storage: fixed an issue where `google_storage_bucket_object` and `google_storage_bucket_objects` data sources would ignore custom endpoints ([#7287](https://github.com/hashicorp/terraform-provider-google-beta/pull/7287))

## 5.27.0 (Apr 30, 2024)

FEATURES:
* **New Data Source:** `google_storage_bucket_objects` ([#7270](https://github.com/hashicorp/terraform-provider-google-beta/pull/7270))
* **New Resource:** `google_composer_user_workloads_secret` ([#7257](https://github.com/hashicorp/terraform-provider-google-beta/pull/7257))
* **New Resource:** `google_compute_security_policy_rule` ([#7282](https://github.com/hashicorp/terraform-provider-google-beta/pull/7282))
* **New Resource:** `google_data_loss_prevention_discovery_config` ([#7252](https://github.com/hashicorp/terraform-provider-google-beta/pull/7252))
* **New Resource:** `google_integrations_auth_config` ([#7268](https://github.com/hashicorp/terraform-provider-google-beta/pull/7268))
* **New Resource:** `google_network_connectivity_internal_range` ([#7265](https://github.com/hashicorp/terraform-provider-google-beta/pull/7265))

IMPROVEMENTS:
* alloydb: added `network_config` field to `google_alloydb_instance` resource ([#7271](https://github.com/hashicorp/terraform-provider-google-beta/pull/7271))
* alloydb: added `public_ip_address` field  to `google_alloydb_instance` resource ([#7271](https://github.com/hashicorp/terraform-provider-google-beta/pull/7271))
* apigee: added `forward_proxy_uri` field to `google_apigee_environment` resource ([#7260](https://github.com/hashicorp/terraform-provider-google-beta/pull/7260))
* bigquerydatapolicy: added `data_masking_policy.routine` field to `google_bigquery_data_policy` resource ([#7250](https://github.com/hashicorp/terraform-provider-google-beta/pull/7250))
* compute: added `server_tls_policy` field to `google_compute_region_target_https_proxy` resource ([#7280](https://github.com/hashicorp/terraform-provider-google-beta/pull/7280))
* filestore: added `protocol` field to `google_filestore_instance` resource to support NFSv3 and NFSv4.1 ([#7254](https://github.com/hashicorp/terraform-provider-google-beta/pull/7254))
* firebasehosting: added `config.rewrites.path` field to `google_firebase_hosting_version` resource ([#7258](https://github.com/hashicorp/terraform-provider-google-beta/pull/7258))
* logging: added `intercept_children` field to `google_logging_organization_sink` and `google_logging_folder_sink` resources ([#7279](https://github.com/hashicorp/terraform-provider-google-beta/pull/7279))
* monitoring: added `service_agent_authentication` field to `google_monitoring_uptime_check_config` resource ([#7276](https://github.com/hashicorp/terraform-provider-google-beta/pull/7276))
* privateca: added `subject_key_id` field to `google_privateca_certificate` and `google_privateca_certificate_authority` resources ([#7273](https://github.com/hashicorp/terraform-provider-google-beta/pull/7273))
* secretmanager: added `version_destroy_ttl` field to `google_secret_manager_secret` resource ([#7253](https://github.com/hashicorp/terraform-provider-google-beta/pull/7253))

BUG FIXES:
* appengine: added suppression for a diff in `google_app_engine_standard_app_version.automatic_scaling` when the block is unset in configuration ([#7262](https://github.com/hashicorp/terraform-provider-google-beta/pull/7262))
* sql: fixed issues with updating the `enable_google_ml_integration` field in `google_sql_database_instance` resource ([#7249](https://github.com/hashicorp/terraform-provider-google-beta/pull/7249))

## 5.26.0 (Apr 22, 2024)

FEATURES:
* **New Resource:** `google_project_iam_member_remove` ([#7242](https://github.com/hashicorp/terraform-provider-google-beta/pull/7242))

IMPROVEMENTS:
* apigee: added support for `api_consumer_data_location`, `api_consumer_data_encryption_key_name`, and `control_plane_encryption_key_name` in `google_apigee_organization` ([#7245](https://github.com/hashicorp/terraform-provider-google-beta/pull/7245))
* artifactregistry: added `remote_repository_config.<facade>_repository.custom_repository.uri` field to `google_artifact_registry_repository` resource. ([#7230](https://github.com/hashicorp/terraform-provider-google-beta/pull/7230))
* bigquery: added `resource_tags` field to `google_bigquery_table` resource ([#7247](https://github.com/hashicorp/terraform-provider-google-beta/pull/7247))
* billing: added `ownership_scope` field to `google_billing_budget` resource ([#7239](https://github.com/hashicorp/terraform-provider-google-beta/pull/7239))
* cloudfunctions2: added `build_config.service_account` field to `google_cloudfunctions2_function` resource ([#7231](https://github.com/hashicorp/terraform-provider-google-beta/pull/7231))
* composer: fixed validation on `google_composer_environment` resource so it will identify a disallowed upgrade to Composer 3 before attempting to provide feedback that's specific to using Composer 3 ([#7213](https://github.com/hashicorp/terraform-provider-google-beta/pull/7213))
* compute: added `params.resource_manager_tags` field to `resource_compute_instance_group_manager` and `resource_compute_region_instance_group_manager` that enables to create these resources with tags (beta) ([#7226](https://github.com/hashicorp/terraform-provider-google-beta/pull/7226))
* resourcemanager: added the field `api_method` to datasource `google_active_folder` so you can use either `SEARCH` or `LIST` to find your folder ([#7248](https://github.com/hashicorp/terraform-provider-google-beta/pull/7248))
* storage: added labels validation to `google_storage_bucket` resource ([#7212](https://github.com/hashicorp/terraform-provider-google-beta/pull/7212))
* workstations: added output-only field `control_plane_ip` to `google_workstations_workstation_cluster` resource (beta) ([#7240](https://github.com/hashicorp/terraform-provider-google-beta/pull/7240))

BUG FIXES:
* apigee: fixed permadiff in ordering of `google_apigee_organization.properties.property`. ([#7234](https://github.com/hashicorp/terraform-provider-google-beta/pull/7234))
* cloudrun: fixed the bug that computed `metadata.0.labels` and `metadata.0.annotations` fields don't appear in terraform plan when creating resource `google_cloud_run_service` and `google_cloud_run_domain_mapping` ([#7217](https://github.com/hashicorp/terraform-provider-google-beta/pull/7217))
* dns: fixed bug where some methods of authentication didn't work when using `dns` data sources ([#7233](https://github.com/hashicorp/terraform-provider-google-beta/pull/7233))
* iam: fixed a bug that prevented setting `create_ignore_already_exists` on existing resources in `google_service_account`. ([#7236](https://github.com/hashicorp/terraform-provider-google-beta/pull/7236))
* sql: fixed issues with updating the `enable_google_ml_integration` field in `google_sql_database_instance` resource ([#7249](https://github.com/hashicorp/terraform-provider-google-beta/pull/7249))
* storage: added validation to `name` field in `google_storage_bucket` resource ([#7237](https://github.com/hashicorp/terraform-provider-google-beta/pull/7237))
* vmwareengine: fixed stretched cluster creation in `google_vmwareengine_private_cloud` ([#7246](https://github.com/hashicorp/terraform-provider-google-beta/pull/7246))

## 5.25.0 (Apr 15, 2024)

FEATURES:
* **New Data Source:** `google_tags_tag_keys` ([#7196](https://github.com/hashicorp/terraform-provider-google-beta/pull/7196))
* **New Data Source:** `google_tags_tag_values` ([#7196](https://github.com/hashicorp/terraform-provider-google-beta/pull/7196))
* **New Resource:** `google_parallelstore_instance` ([#7209](https://github.com/hashicorp/terraform-provider-google-beta/pull/7209))

IMPROVEMENTS:
* bigquery: added in-place schema column drop support for `google_bigquery_table` resource ([#7193](https://github.com/hashicorp/terraform-provider-google-beta/pull/7193))
* compute: added `endpoint_types` field to `google_compute_router_nat` resource ([#7190](https://github.com/hashicorp/terraform-provider-google-beta/pull/7190))
* compute: added `enable_ipv4`, `ipv4_nexthop_address` and `peer_ipv4_nexthop_address` fields to `google_compute_router_peer` resource ([#7207](https://github.com/hashicorp/terraform-provider-google-beta/pull/7207))
* compute: added `identifier_range` field to `google_compute_router` resource ([#7207](https://github.com/hashicorp/terraform-provider-google-beta/pull/7207))
* compute: added `ip_version` field to `google_compute_router_interface` resource ([#7207](https://github.com/hashicorp/terraform-provider-google-beta/pull/7207))
* compute: increased timeouts from 8 minutes to 20 minutes for `google_compute_security_policy` resource ([#7204](https://github.com/hashicorp/terraform-provider-google-beta/pull/7204))
* container: added `stateful_ha_config` field to `google_container_cluster` resource ([#7206](https://github.com/hashicorp/terraform-provider-google-beta/pull/7206))
* firestore: added `vector_config` field to `google_firestore_index` resource ([#7180](https://github.com/hashicorp/terraform-provider-google-beta/pull/7180))
* gkebackup: added `backup_schedule.rpo_config` field to `google_gke_backup_backup_plan` resource ([#7211](https://github.com/hashicorp/terraform-provider-google-beta/pull/7211))
* networksecurity: added `disabled` field to `google_network_security_firewall_endpoint_association` resource ([#7184](https://github.com/hashicorp/terraform-provider-google-beta/pull/7184))
* sql: added `enable_google_ml_integration` field to `google_sql_database_instance` resource ([#7208](https://github.com/hashicorp/terraform-provider-google-beta/pull/7208))
* storage: added labels validation to `google_storage_bucket` resource ([#7212](https://github.com/hashicorp/terraform-provider-google-beta/pull/7212))
* vmwareengine: added `preferred_zone` and `secondary_zone` fields to `google_vmwareengine_private_cloud` resource ([#7210](https://github.com/hashicorp/terraform-provider-google-beta/pull/7210))

BUG FIXES:
* networksecurity: fixed an issue where `google_network_security_firewall_endpoint_association` resource could not be created due to a bad parameter ([#7184](https://github.com/hashicorp/terraform-provider-google-beta/pull/7184))
* privateca: fixed permission issue by specifying signer certs chain when activating a sub-CA across regions for `google_privateca_certificate_authority` resource ([#7197](https://github.com/hashicorp/terraform-provider-google-beta/pull/7197))

## 5.24.0 (Apr 8, 2024)

IMPROVEMENTS:
* cloudrunv2: added `template.volumes.nfs` field to `google_cloud_run_v2_job` resource ([#7169](https://github.com/hashicorp/terraform-provider-google-beta/pull/7169))
* container: added `enable_cilium_clusterwide_network_policy` field to `google_container_cluster` resource ([#7171](https://github.com/hashicorp/terraform-provider-google-beta/pull/7171))
* container: added `node_pool_auto_config.resource_manager_tags` field to `google_container_cluster` resource ([#7162](https://github.com/hashicorp/terraform-provider-google-beta/pull/7162))
* gkeonprem: added `disable_bundled_ingress` field to `google_gkeonprem_vmware_cluster` resource ([#7163](https://github.com/hashicorp/terraform-provider-google-beta/pull/7163))
* redis: added `node_type` and `precise_size_gb` fields to `google_redis_cluster` ([#7174](https://github.com/hashicorp/terraform-provider-google-beta/pull/7174))
* storage: added `project_number` attribute to `google_storage_bucket` resource and data source ([#7164](https://github.com/hashicorp/terraform-provider-google-beta/pull/7164))
* storage: added ability to provide `project` argument to `google_storage_bucket` data source. This will not impact reading the resource's data, instead this helps users avoid calls to the Compute API within the data source. ([#7164](https://github.com/hashicorp/terraform-provider-google-beta/pull/7164))

BUG FIXES:
* appengine: fixed a crash in `google_app_engine_flexible_app_version` due to the `deployment` field not being returned by the API ([#7175](https://github.com/hashicorp/terraform-provider-google-beta/pull/7175))
* bigquery: fixed a crash when `google_bigquery_table` had a `primary_key.columns` entry set to `""` ([#7166](https://github.com/hashicorp/terraform-provider-google-beta/pull/7166))
* compute: fixed update scenarios on `google_compute_region_target_https_proxy` and `google_compute_target_https_proxy` resources. ([#7170](https://github.com/hashicorp/terraform-provider-google-beta/pull/7170))
* dataflow: fixed an issue where the provider would crash when `enable_streaming_engine` is set as a `parameter` value in `google_dataflow_flex_template_job` ([#7160](https://github.com/hashicorp/terraform-provider-google-beta/pull/7160))

## 5.23.0 (Apr 01, 2023)

NOTES:
* provider: introduced support for [provider-defined functions](https://developer.hashicorp.com/terraform/plugin/framework/functions). This feature is in Terraform v1.8.0+. ([#7153](https://github.com/hashicorp/terraform-provider-google-beta/pull/7153))

DEPRECATIONS:
* kms: deprecated `attestation.external_protection_level_options` in favor of `external_protection_level_options` in `google_kms_crypto_key_version` ([#7155](https://github.com/hashicorp/terraform-provider-google-beta/pull/7155))

FEATURES:
* **New Data Source:** `google_apphub_application` ([#7143](https://github.com/hashicorp/terraform-provider-google-beta/pull/7143))
* **New Resource:** `google_cloud_quotas_quota_preference` ([#7126](https://github.com/hashicorp/terraform-provider-google-beta/pull/7126))
* **New Resource:** `google_vertex_ai_deployment_resource_pool` ([#7158](https://github.com/hashicorp/terraform-provider-google-beta/pull/7158))
* **New Resource:** `google_integrations_client` ([#7129](https://github.com/hashicorp/terraform-provider-google-beta/pull/7129))

IMPROVEMENTS:
* bigquery: added `dataGovernanceType` to `google_bigquery_routine` resource ([#7149](https://github.com/hashicorp/terraform-provider-google-beta/pull/7149))
* bigquery: added support for `external_data_configuration.json_extension` to `google_bigquery_table` ([#7138](https://github.com/hashicorp/terraform-provider-google-beta/pull/7138))
* compute: added `cloud_router_ipv6_address`, `customer_router_ipv6_address` fields to `google_compute_interconnect_attachment` resource ([#7151](https://github.com/hashicorp/terraform-provider-google-beta/pull/7151))
* compute: added `generated_id` field to `google_compute_region_backend_service` resource ([#7128](https://github.com/hashicorp/terraform-provider-google-beta/pull/7128))
* integrations: added deletion support for `google_integrations_client` resource ([#7142](https://github.com/hashicorp/terraform-provider-google-beta/pull/7142))
* kms: added `crypto_key_backend` field to `google_kms_crypto_key` resource ([#7155](https://github.com/hashicorp/terraform-provider-google-beta/pull/7155))
* metastore: added `scheduled_backup` field to `google_dataproc_metastore_service` resource ([#7140](https://github.com/hashicorp/terraform-provider-google-beta/pull/7140))
* provider: added provider-defined function `name_from_id` for retrieving the short-form name of a resource from its self link or id ([#7153](https://github.com/hashicorp/terraform-provider-google-beta/pull/7153))
* provider: added provider-defined function `project_from_id` for retrieving the project id from a resource's self link or id ([#7153](https://github.com/hashicorp/terraform-provider-google-beta/pull/7153))
* provider: added provider-defined function `region_from_zone` for deriving a region from a zone's name ([#7153](https://github.com/hashicorp/terraform-provider-google-beta/pull/7153))
* provider: added provider-defined functions `location_from_id`, `region_from_id`, and `zone_from_id` for retrieving the location/region/zone names from a resource's self link or id ([#7153](https://github.com/hashicorp/terraform-provider-google-beta/pull/7153))

BUG FIXES:
* cloudrunv2: fixed Terraform state inconsistency when resource `google_cloud_run_v2_job` creation fails ([#7159](https://github.com/hashicorp/terraform-provider-google-beta/pull/7159))
* cloudrunv2: fixed Terraform state inconsistency when resource `google_cloud_run_v2_service` creation fails ([#7159](https://github.com/hashicorp/terraform-provider-google-beta/pull/7159))
* container: fixed `google_container_cluster` permadiff when `master_ipv4_cidr_block` is set for a private flexible cluster ([#7147](https://github.com/hashicorp/terraform-provider-google-beta/pull/7147))
* dataflow: fixed an issue where the provider would crash when `enableStreamingEngine` is set as a `parameter` value in `google_dataflow_flex_template_job` ([#7160](https://github.com/hashicorp/terraform-provider-google-beta/pull/7160))
* kms: added top-level `external_protection_level_options` field in `google_kms_crypto_key_version` resource ([#7155](https://github.com/hashicorp/terraform-provider-google-beta/pull/7155))

## 5.22.0 (Mar 26, 2024)

BREAKING CHANGES:
* networksecurity: added required field `billing_project_id` to `google_network_security_firewall_endpoint` resource. Any configuration without `billing_project_id` specified will cause resource creation fail (beta) ([#7124](https://github.com/hashicorp/terraform-provider-google-beta/pull/7124))

FEATURES:
* **New Data Source:** `google_cloud_quotas_quota_info` ([#7092](https://github.com/hashicorp/terraform-provider-google-beta/pull/7092))
* **New Data Source:** `google_cloud_quotas_quota_infos` ([#7116](https://github.com/hashicorp/terraform-provider-google-beta/pull/7116))
* **New Resource:** `google_access_context_manager_service_perimeter_dry_run_resource` ([#7115](https://github.com/hashicorp/terraform-provider-google-beta/pull/7115))

IMPROVEMENTS:
* accesscontextmanager: supported managing service perimeter dry run resources outside the perimeter via new resource `google_access_context_manager_service_perimeter_dry_run_resource` ([#7115](https://github.com/hashicorp/terraform-provider-google-beta/pull/7115))
* cloudrunv2: added plan-time validation to restrict number of ports to 1 in `google_cloud_run_v2_service` ([#7103](https://github.com/hashicorp/terraform-provider-google-beta/pull/7103))
* cloudrunv2: supported mounting Cloud Storage buckets using GCSFuse in `google_cloud_run_v2_job` ([#7102](https://github.com/hashicorp/terraform-provider-google-beta/pull/7102))
* composer: added field `count` to validate number of DAG processors in `google_composer_environment` ([#7120](https://github.com/hashicorp/terraform-provider-google-beta/pull/7120))
* compute: added enumeration value `SEV_LIVE_MIGRATABLE_V2` for the `guest_os_features` of `google_compute_disk` ([#7123](https://github.com/hashicorp/terraform-provider-google-beta/pull/7123))
* compute: added `status.all_instances_config.revision` field to `google_compute_instance_group_manager` and `google_compute_region_instance_group_manager` ([#7104](https://github.com/hashicorp/terraform-provider-google-beta/pull/7104))
* compute: added field `path_template_match` to resource `google_compute_region_url_map` ([#7094](https://github.com/hashicorp/terraform-provider-google-beta/pull/7094))
* compute: added field `path_template_rewrite` to resource `google_compute_region_url_map` ([#7094](https://github.com/hashicorp/terraform-provider-google-beta/pull/7094))
* pubsub: added `ingestion_data_source_settings` field to `google_pubsub_topic` resource ([#7111](https://github.com/hashicorp/terraform-provider-google-beta/pull/7111))
* storage: added 'soft_delete_policy' to 'google_storage_bucket' resource ([#7119](https://github.com/hashicorp/terraform-provider-google-beta/pull/7119))
* workstations: added `host.gceInstance.boostConfig` to `google_workstations_workstation_config` ([#7122](https://github.com/hashicorp/terraform-provider-google-beta/pull/7122))

BUG FIXES:
* accesscontextmanager: fixed an issue with `access_context_manager_service_perimeter_ingress_policy` and `access_context_manager_service_perimeter_egress_policy` where updates could not be applied after initial creation. Any updates applied to these resources will now involve their recreation. To ensure that new policies are added before old ones are removed, add a `lifecycle` block with `create_before_destroy = true` to your resource configuration alongside other updates. ([#7105](https://github.com/hashicorp/terraform-provider-google-beta/pull/7105))
* firebase: made the `google_firebase_android_app` resource's `package_name` field required and immutable. This prevents API errors encountered by users who attempted to update or leave that field unset in their configurations. ([#7100](https://github.com/hashicorp/terraform-provider-google-beta/pull/7100))
* spanner: removed validation function for the field `version_retention_period` in the resource `google_spanner_database` and directly returned error from backend ([#7117](https://github.com/hashicorp/terraform-provider-google-beta/pull/7117))

## 5.21.0 (Mar 18, 2024)

FEATURES:
* **New Data Source:** `google_apphub_discovered_service` ([#7080](https://github.com/hashicorp/terraform-provider-google-beta/pull/7080))
* **New Data Source:** `google_apphub_discovered_workload` ([#7083](https://github.com/hashicorp/terraform-provider-google-beta/pull/7083))
* **New Data Source:** `google_cloud_quotas_quota_info` ([#7092](https://github.com/hashicorp/terraform-provider-google-beta/pull/7092))
* **New Resource:** `google_apphub_workload` ([#7088](https://github.com/hashicorp/terraform-provider-google-beta/pull/7088))
* **New Resource:** `google_firebase_app_check_device_check_config` ([#7062](https://github.com/hashicorp/terraform-provider-google-beta/pull/7062))
* **New Resource:** `google_iap_tunnel_dest_group` ([#7072](https://github.com/hashicorp/terraform-provider-google-beta/pull/7072))
* **New Resource:** `google_kms_ekm_connection` ([#7059](https://github.com/hashicorp/terraform-provider-google-beta/pull/7059))
* **New Resource:** `google_apphub_application` ([#7051](https://github.com/hashicorp/terraform-provider-google-beta/pull/7051))
* **New Resource:** `google_apphub_service` ([#7090](https://github.com/hashicorp/terraform-provider-google-beta/pull/7090))
* **New Resource:** `google_apphub_service_project_attachment` ([#7073](https://github.com/hashicorp/terraform-provider-google-beta/pull/7073))
* **New Resource:** `google_network_security_firewall_endpoint_association` ([#7075](https://github.com/hashicorp/terraform-provider-google-beta/pull/7075))

IMPROVEMENTS:
* cloudrunv2: added support for `scaling.min_instance_count` in `google_cloud_run_v2_service`. ([#7053](https://github.com/hashicorp/terraform-provider-google-beta/pull/7053))
* firestore: added `cmek_config` field to `google_firestore_database` resource ([#7054](https://github.com/hashicorp/terraform-provider-google-beta/pull/7054))
* gkeonprem: allowed `vcenter_network` to be set in `google_gkeonprem_vmware_cluster`, previously it was output-only ([#7055](https://github.com/hashicorp/terraform-provider-google-beta/pull/7055))
* storagetransferservice: added field `transfer_spec.azure_blob_storage_data_source.credentials_secret` to `google_storage_transfer_job` ([#7091](https://github.com/hashicorp/terraform-provider-google-beta/pull/7091))
* workstations: added support for `ephemeral_directories` in `google_workstations_workstation_config` ([#7061](https://github.com/hashicorp/terraform-provider-google-beta/pull/7061))

BUG FIXES:
* compute: allowed sending empty values for `SERVERLESS` in `google_compute_region_network_endpoint_group` resource ([#7052](https://github.com/hashicorp/terraform-provider-google-beta/pull/7052))
* notebooks: fixed an issue where default tags would cause a diff recreating `google_notebooks_instance` resources ([#7086](https://github.com/hashicorp/terraform-provider-google-beta/pull/7086))
* storage: fixed an issue where two or more lifecycle rules with different values of `no_age` field always generates change in `google_storage_bucket` resource. ([#7060](https://github.com/hashicorp/terraform-provider-google-beta/pull/7060))

## 5.20.0 (Mar 11, 2024)

FEATURES:
* **New Resource:** `google_clouddeploy_custom_target_type_iam_*` ([#7029](https://github.com/hashicorp/terraform-provider-google-beta/pull/7029))

IMPROVEMENTS:
* certificatemanager: added `type` field to `google_certificate_manager_dns_authorization` resource ([#7036](https://github.com/hashicorp/terraform-provider-google-beta/pull/7036))
* compute: added the `network_url` attribute to the `consumer_accept_list`-block of the `google_compute_service_attachment` resource ([#7047](https://github.com/hashicorp/terraform-provider-google-beta/pull/7047))
* gkehub: added support for `policycontroller.policy_controller_hub_config.policy_content.bundles` and 
`policycontroller.policy_controller_hub_config.deployment_configs` fields to `google_gke_hub_feature_membership` ([#7043](https://github.com/hashicorp/terraform-provider-google-beta/pull/7043))

BUG FIXES:
* artifactregistry: fixed permadiff when `google_artifact_repository.docker_config` field is unset ([#7044](https://github.com/hashicorp/terraform-provider-google-beta/pull/7044))
* bigquery: corrected plan-time validation on `google_bigquery_dataset.dataset_id` ([#7032](https://github.com/hashicorp/terraform-provider-google-beta/pull/7032))
* kms: fixed issue where `google_kms_crypto_key_version.attestation.cert_chains` properties were incorrectly set to type string ([#7045](https://github.com/hashicorp/terraform-provider-google-beta/pull/7045))

## 5.19.0 (Mar 4, 2024)

FEATURES:
* **New Resource:** `google_clouddeploy_target_iam_*` ([#7012](https://github.com/hashicorp/terraform-provider-google-beta/pull/7012))

IMPROVEMENTS:
* bigquery: added `remote_function_options` field to `google_bigquery_routine` resource ([#7015](https://github.com/hashicorp/terraform-provider-google-beta/pull/7015))
* certificatemanager: added `location` field to `google_certificate_manager_dns_authorization` resource ([#7006](https://github.com/hashicorp/terraform-provider-google-beta/pull/7006))
* composer: added `composer_network_attachment` and modified `network`/`subnetwork` to support composer 3 in `google_composer_environment`  ([#7023](https://github.com/hashicorp/terraform-provider-google-beta/pull/7023))
* composer: added validations for composer 2/3 only fields in `google_composer_environment` ([#7008](https://github.com/hashicorp/terraform-provider-google-beta/pull/7008))
* compute: added `certificate_manager_certificates` field to `google_compute_region_target_https_proxy` resource ([#7010](https://github.com/hashicorp/terraform-provider-google-beta/pull/7010))
* gkehub2: added `namespace_labels` field to `google_gke_hub_scope` resource ([#7022](https://github.com/hashicorp/terraform-provider-google-beta/pull/7022))

BUG FIXES:
* resourcemanager: added a retry to deleting the default network when `auto_create_network` is false in `google_project` ([#7021](https://github.com/hashicorp/terraform-provider-google-beta/pull/7021))

## 5.18.0 (Feb 26, 2024)

BREAKING CHANGES:
* securityposture: marked `policy_sets` and `policy_sets.policies` required in `google_securityposture_posture`. API validation already enforced this, so no resources could be provisioned without these ([#6981](https://github.com/hashicorp/terraform-provider-google-beta/pull/6981))

FEATURES:
* **New Data Source:** `google_compute_forwarding_rules` ([#6997](https://github.com/hashicorp/terraform-provider-google-beta/pull/6997))
* **New Resource:** `google_firebase_app_check_app_attest_config` ([#6971](https://github.com/hashicorp/terraform-provider-google-beta/pull/6971))
* **New Resource:** `google_firebase_app_check_play_integrity_config` ([#6971](https://github.com/hashicorp/terraform-provider-google-beta/pull/6971))
* **New Resource:** `google_firebase_app_check_recaptcha_enterprise_config` ([#6989](https://github.com/hashicorp/terraform-provider-google-beta/pull/6989))
* **New Resource:** `google_firebase_app_check_recaptcha_v3_config` ([#6989](https://github.com/hashicorp/terraform-provider-google-beta/pull/6989))
* **New Resource:** `google_migration_center_preference_set` ([#6974](https://github.com/hashicorp/terraform-provider-google-beta/pull/6974))
* **New Resource:** `google_netapp_volume_replication` ([#7002](https://github.com/hashicorp/terraform-provider-google-beta/pull/7002))

IMPROVEMENTS:
* cloudfunctions: added output-only `version_id` field on `google_cloudfunctions_function` ([#6968](https://github.com/hashicorp/terraform-provider-google-beta/pull/6968))
* composer: supported patch versions of airflow on `google_composer_environment` ([#7000](https://github.com/hashicorp/terraform-provider-google-beta/pull/7000))
* compute: supported updating `network_interface.stack_type` field on `google_compute_instance` resource. ([#6977](https://github.com/hashicorp/terraform-provider-google-beta/pull/6977))
* container: added `node_config.resource_manager_tags` field to `google_container_cluster` resource ([#7001](https://github.com/hashicorp/terraform-provider-google-beta/pull/7001))
* container: added `node_config.resource_manager_tags` field to `google_container_node_pool` resource ([#7001](https://github.com/hashicorp/terraform-provider-google-beta/pull/7001))
* container: added output-only fields `membership_id` and  `membership_location` under `fleet` in `google_container_cluster` resource ([#6983](https://github.com/hashicorp/terraform-provider-google-beta/pull/6983))
* looker: added `custom_domain` field to `google_looker_instance ` resource ([#6979](https://github.com/hashicorp/terraform-provider-google-beta/pull/6979))
* netapp: added field `restore_parameters` and output-only fields `state`, `state_details` and `create_time` to `google_netapp_volume` resource ([#6976](https://github.com/hashicorp/terraform-provider-google-beta/pull/6976))
* workbench: added `container_image` field to `google_workbench_instance` resource ([#6988](https://github.com/hashicorp/terraform-provider-google-beta/pull/6988))
* workbench: added `shielded_instance_config` field to `google_workbench_instance` resource ([#6984](https://github.com/hashicorp/terraform-provider-google-beta/pull/6984))

BUG FIXES:
* bigquery: allowed users to set permissions for `principal`/`principalSets` (`iamMember`) in `google_bigquery_dataset_iam_member`. ([#6975](https://github.com/hashicorp/terraform-provider-google-beta/pull/6975))
* cloudfunctions2: fixed an issue where not specifying `event_config.trigger_region` in `google_cloudfunctions2_function` resulted in a permanent diff. The field now pulls a default value from the API when unset. ([#6991](https://github.com/hashicorp/terraform-provider-google-beta/pull/6991))
* compute: fixed perma-diff on `min_ports_per_vm` in `google_compute_router_nat` when the field is unset by making the field default to the API-set value ([#6993](https://github.com/hashicorp/terraform-provider-google-beta/pull/6993))
* dataflow: fixed crash in `google_dataflox_job` to return an error instead if a job's Environment field is nil when reading job information ([#6999](https://github.com/hashicorp/terraform-provider-google-beta/pull/6999))
* notebooks: changed `tag` field to default to the API's value if not specified in `google_notebooks_instance` ([#6986](https://github.com/hashicorp/terraform-provider-google-beta/pull/6986))

## 5.17.0 (Feb 20, 2024)

NOTES:
* cloudbuildv2: changed underlying actuation engine for `google_cloudbuildv2_connection`, there should be no user-facing impact ([#6943](https://github.com/hashicorp/terraform-provider-google-beta/pull/6943))

DEPRECATIONS:
* container: deprecated support for `relay_mode` field in `google_container_cluster.monitoring_config.advanced_datapath_observability_config` in favor of `enable_relay` field, `relay_mode` field will be removed a future major release ([#6960](https://github.com/hashicorp/terraform-provider-google-beta/pull/6960))

FEATURES:
* **New Resource:** `google_firebase_app_check_debug_token` ([#6953](https://github.com/hashicorp/terraform-provider-google-beta/pull/6953))
* **New Resource:** `google_network_security_firewall_endpoint` ([#6940](https://github.com/hashicorp/terraform-provider-google-beta/pull/6940))
* **New Resource:** `google_clouddeploy_custom_target_type` ([#6956](https://github.com/hashicorp/terraform-provider-google-beta/pull/6956))
* **New Resource:** `google_network_security_security_profile_group` ([#6961](https://github.com/hashicorp/terraform-provider-google-beta/pull/6961))

IMPROVEMENTS:
* cloudasset: allowed overriding the billing project for the `google_cloud_asset_resources_search_all` datasource ([#6941](https://github.com/hashicorp/terraform-provider-google-beta/pull/6941))
* clouddeploy: added support for `canary_revision_tags`, `prior_revision_tags`, `stable_revision_tags`, and `stable_cutback_duration` to `google_clouddeploy_delivery_pipeline` ([#6951](https://github.com/hashicorp/terraform-provider-google-beta/pull/6951))
* cloudfunctions: added `version_id` on `google_cloudfunctions_function` ([#6968](https://github.com/hashicorp/terraform-provider-google-beta/pull/6968))
* container: added support for `enable_relay` field to `google_container_cluster.monitoring_config.advanced_datapath_observability_config` ([#6960](https://github.com/hashicorp/terraform-provider-google-beta/pull/6960))
* eventarc: added support for `http_endpoint.uri` and `network_config.network_attachment` to `google_eventarc_trigger` ([#6951](https://github.com/hashicorp/terraform-provider-google-beta/pull/6951))
* healthcare: added `reject_duplicate_message` field to `google_healthcare_hl7_v2_store ` resource ([#6964](https://github.com/hashicorp/terraform-provider-google-beta/pull/6964))
* identityplatform: added `client`, `permissions`, `monitoring` and `mfa` fields to `google_identity_platform_config` ([#6944](https://github.com/hashicorp/terraform-provider-google-beta/pull/6944))
* notebooks: added `desired_state` field to `google_notebooks_instance` ([#6965](https://github.com/hashicorp/terraform-provider-google-beta/pull/6965))
* vertexai: added `feature_registry_source` field to `google_vertex_ai_feature_online_store_featureview` resource ([#6962](https://github.com/hashicorp/terraform-provider-google-beta/pull/6962))
* workbench: added `desired_state` field to `google_workbench_instance` resource ([#6966](https://github.com/hashicorp/terraform-provider-google-beta/pull/6966))
* workstations: added support for `disable_ssh` in `google_workstations_workstation_config` ([#6947](https://github.com/hashicorp/terraform-provider-google-beta/pull/6947))

BUG FIXES:
* compute: made `resource_manager_tags` updatable on `google_compute_instance_template` and `google_compute_region_instance_template` ([#6958](https://github.com/hashicorp/terraform-provider-google-beta/pull/6958))
* notebooks: prevented recreation of `google_notebooks_instance` when `kms_key` or `service_account_scopes` are changed server-side ([#6948](https://github.com/hashicorp/terraform-provider-google-beta/pull/6948))

## 5.16.0 (Feb 12, 2024)

FEATURES:
* **New Resource:** `google_clouddeploy_delivery_pipeline_iam_*` ([#6928](https://github.com/hashicorp/terraform-provider-google-beta/pull/6928))
* **New Resource:** `google_compute_instance_group_membership` ([#6933](https://github.com/hashicorp/terraform-provider-google-beta/pull/6933))
* **New Resource:** `google_discovery_engine_search_engine` ([#6919](https://github.com/hashicorp/terraform-provider-google-beta/pull/6919))
* **New Resource:** `google_firebase_app_check_service_config` ([#6921](https://github.com/hashicorp/terraform-provider-google-beta/pull/6921))

IMPROVEMENTS:
* bigquery: promoted `table_replication_info` field on `resource_bigquery_table` resource to GA  ([#6929](https://github.com/hashicorp/terraform-provider-google-beta/pull/6929))
* compute: added `confidential_instance_config.confidential_instance_type` field to `google_compute_instance`,  `google_compute_instance_template`, and `google_compute_region_instance_template` resources ([#6934](https://github.com/hashicorp/terraform-provider-google-beta/pull/6934))
* networksecurity: removed unused custom code from `google_network_security_address_group` ([#6931](https://github.com/hashicorp/terraform-provider-google-beta/pull/6931))
* provider: added an optional provider level label `goog-terraform-provisioned` to identify resources that were created by Terraform when viewing/editing these resources in other tools. ([#6924](https://github.com/hashicorp/terraform-provider-google-beta/pull/6924))

BUG FIXES:
* firebasehosting: fixed typing in `google_firebase_hosting_custom_domain` `issues.details` field ([#6926](https://github.com/hashicorp/terraform-provider-google-beta/pull/6926))

## 5.15.0 (Feb 5, 2024)

FEATURES:
* **New Data Source:** `google_compute_machine_types` ([#6903](https://github.com/hashicorp/terraform-provider-google-beta/pull/6903))
* **New Resource:** `google_blockchain_node_engine_blockchain_nodes` ([#6897](https://github.com/hashicorp/terraform-provider-google-beta/pull/6897))
* **New Resource:** `google_compute_region_network_endpoint` ([#6913](https://github.com/hashicorp/terraform-provider-google-beta/pull/6913))
* **New Resource:** `google_discovery_engine_chat_engine` ([#6918](https://github.com/hashicorp/terraform-provider-google-beta/pull/6918))
* **New Resource:** `google_discovery_engine_search_engine` ([#6919](https://github.com/hashicorp/terraform-provider-google-beta/pull/6919))
* **New Resource:** `google_netapp_volume_snapshot` ([#6914](https://github.com/hashicorp/terraform-provider-google-beta/pull/6914))

IMPROVEMENTS:
* compute: added `INTERNET_IP_PORT` and `INTERNET_FQDN_PORT` options for the `google_compute_region_network_endpoint_group` resource. ([#6913](https://github.com/hashicorp/terraform-provider-google-beta/pull/6913))
* compute: added `creation_timestamp` to `google_compute_instance_group_manager` and `google_compute_region_instance_group_manager`. ([#6904](https://github.com/hashicorp/terraform-provider-google-beta/pull/6904))
* compute: added `disk_id` attribute to `google_compute_disk` resource ([#6906](https://github.com/hashicorp/terraform-provider-google-beta/pull/6906))
* compute: added `stack_type` attribute for `google_compute_interconnect_attachment` resource. ([#6915](https://github.com/hashicorp/terraform-provider-google-beta/pull/6915))
* compute: updated the `google_compute_security_policy` resource's `json_parsing` field to accept the value `STANDARD_WITH_GRAPHQL` ([#6898](https://github.com/hashicorp/terraform-provider-google-beta/pull/6898))
* memcache: added `reserved_ip_range_id` field to `google_memcache_instance` resource ([#6901](https://github.com/hashicorp/terraform-provider-google-beta/pull/6901))
* netapp: added `deletion_policy` field to `google_netapp_volume` resource ([#6905](https://github.com/hashicorp/terraform-provider-google-beta/pull/6905))

BUG FIXES:
* alloydb: fixed an issue where `database_flags` in secondary `google_alloydb_instance` resources would cause a diff, as they are copied from the primary ([#6910](https://github.com/hashicorp/terraform-provider-google-beta/pull/6910))
* filestore: made `google_filestore_instance.source_backup` field configurable ([#6899](https://github.com/hashicorp/terraform-provider-google-beta/pull/6899))
* vmwareengine: fixed a bug to prevent recreation of existing [`google_vmwareengine_private_cloud`](https://registry.terraform.io/providers/hashicorp/google/latest/docs/data-sources/vmwareengine_private_cloud) resources when upgrading provider version from <5.10.0  ([#6911](https://github.com/hashicorp/terraform-provider-google-beta/pull/6911))

## 5.14.0 (Jan 29, 2024)

FEATURES:
* **New Resource:** `google_discovery_engine_data_store` ([#6892](https://github.com/hashicorp/terraform-provider-google-beta/pull/6892))
* **New Resource:** `google_securityposture_posture_deployment` ([#6893](https://github.com/hashicorp/terraform-provider-google-beta/pull/6893))
* **New Resource:** `google_securityposture_posture` ([#6890](https://github.com/hashicorp/terraform-provider-google-beta/pull/6890))

IMPROVEMENTS:
* cloudrun: added `template.spec.volumes.csi` field to `google_cloud_run_service` resource to support mounting Cloud Storage buckets using GCSFuse ([#6875](https://github.com/hashicorp/terraform-provider-google-beta/pull/6875))
* composer: added `data_retention_config` field to `google_composer_environment` resource ([#6877](https://github.com/hashicorp/terraform-provider-google-beta/pull/6877))
* logging: updated the `google_logging_project_bucket_config` resource to be created using the asynchronous create method ([#6883](https://github.com/hashicorp/terraform-provider-google-beta/pull/6883))
* pubsub: added `use_table_schema` field to `google_pubsub_subscription` resource ([#6881](https://github.com/hashicorp/terraform-provider-google-beta/pull/6881))
* vertexai: added `vector_search_config` field to `google_vertex_ai_feature_online_store_featureview` resource ([#6876](https://github.com/hashicorp/terraform-provider-google-beta/pull/6876))
* workflows: added `call_log_level` field to `google_workflows_workflow` resource ([#6878](https://github.com/hashicorp/terraform-provider-google-beta/pull/6878))
* workstations: added `readiness_checks` field to `google_workstations_workstation_config` resource ([#6895](https://github.com/hashicorp/terraform-provider-google-beta/pull/6895))

BUG FIXES:
* cloudfunctions2: fixed permadiff when `build_config.docker_repository` field is not specified on `google_cloudfunctions2_function` resource ([#6887](https://github.com/hashicorp/terraform-provider-google-beta/pull/6887))
* compute: fixed error when `iap` field is unset for `google_compute_region_backend_service` resource ([#6886](https://github.com/hashicorp/terraform-provider-google-beta/pull/6886))
* eventarc: fixed error when setting `destination.cloud_function` field on `google_eventarc_trigger` resource by making it output-only ([#6879](https://github.com/hashicorp/terraform-provider-google-beta/pull/6879))

## 5.13.0 (Jan 22, 2024)

NOTES:
* cloudbuildv2: changed underlying actuation engine for `google_cloudbuildv2_repository`, there should be no user-facing impact ([#6843](https://github.com/hashicorp/terraform-provider-google-beta/pull/6843))
* provider: added support for in-place update for `labels` and `terraform_labels` fields in immutable resources ([#6857](https://github.com/hashicorp/terraform-provider-google-beta/pull/6857))

FEATURES:
* **New Resource:** `google_netapp_backup_policy` ([#6839](https://github.com/hashicorp/terraform-provider-google-beta/pull/6839))
* **New Resource:** `google_netapp_volume` ([#6852](https://github.com/hashicorp/terraform-provider-google-beta/pull/6852))
* **New Resource:** `google_network_security_address_group_iam_*` ([#6859](https://github.com/hashicorp/terraform-provider-google-beta/pull/6859))
* **New Resource:** `google_network_security_security_profile` ([#6868](https://github.com/hashicorp/terraform-provider-google-beta/pull/6868))
* **New Resource:** `google_vertex_ai_feature_group_feature` ([#6861](https://github.com/hashicorp/terraform-provider-google-beta/pull/6861))

IMPROVEMENTS:
* alloydb: allowed `database_version` as an input on `google_alloydb_cluster` resource ([#6841](https://github.com/hashicorp/terraform-provider-google-beta/pull/6841))
* bigquery: added `spark_options` field to `google_bigquery_routine` resource ([#6867](https://github.com/hashicorp/terraform-provider-google-beta/pull/6867))
* bigquery: added support for replica materialized view in `google_bigquery_table` resource ([#6865](https://github.com/hashicorp/terraform-provider-google-beta/pull/6865))
* cloudrunv2: added `nfs` and `gcs` fields to `google_cloud_run_v2_service.template.volumes` ([#6845](https://github.com/hashicorp/terraform-provider-google-beta/pull/6845))
* cloudrunv2: added `tcp_socket` field to `google_cloud_run_v2.template.containers.liveness_probe` ([#6845](https://github.com/hashicorp/terraform-provider-google-beta/pull/6845))
* composer: added `enable_private_environment` and `enable_private_builds_only` fields to `google_composer_environment` resource ([#6870](https://github.com/hashicorp/terraform-provider-google-beta/pull/6870))
* compute: added `enable_confidential_compute` field to `google_compute_instance.boot_disk.initialize_params` ([#6842](https://github.com/hashicorp/terraform-provider-google-beta/pull/6842))
* gkehub2: added `clusterupgrade` field to `google_gke_hub_feature` resource ([#6836](https://github.com/hashicorp/terraform-provider-google-beta/pull/6836))
* healthcare: added `enable_history_modifications` field to `google_healthcare_fhir_store` resource ([#6864](https://github.com/hashicorp/terraform-provider-google-beta/pull/6864))
* notebooks: allowed `machine_type` and `accelerator_config` to be updatable on `google_notebooks_runtime` resource ([#6854](https://github.com/hashicorp/terraform-provider-google-beta/pull/6854))
* workstations: added `disable_tcp_connections` field to `google_workstations_workstation_config` resource ([#6863](https://github.com/hashicorp/terraform-provider-google-beta/pull/6863))

BUG FIXES:
* compute: fixed the bug that `max_ttl` is sent in API calls even it is removed from configuration when changing cache_mode to FORCE_CACHE_ALL in `google_compute_backend_bucket` resource ([#6847](https://github.com/hashicorp/terraform-provider-google-beta/pull/6847))
* networkservices: fixed a perma-diff on `addresses` field in `google_network_services_gateway` resource ([#6871](https://github.com/hashicorp/terraform-provider-google-beta/pull/6871))
* provider: fixed `universe_domain` behavior to correctly throw an error when explicitly configured `universe_domain` values did not match credentials assumed to be in the default universe ([#6860](https://github.com/hashicorp/terraform-provider-google-beta/pull/6860))
* spanner: fixed error when adding `autoscaling_config` to an existing `google_spanner_instance` resource ([#6869](https://github.com/hashicorp/terraform-provider-google-beta/pull/6869))

## 5.12.0 (Jan 16, 2024)

FEATURES:
* **New Data Source:** `google_dns_managed_zones` ([#6835](https://github.com/hashicorp/terraform-provider-google-beta/pull/6835))
* **New Data Source:** `google_filestore_instance` ([#6822](https://github.com/hashicorp/terraform-provider-google-beta/pull/6822))
* **New Data Source:** `google_vmwareengine_external_access_rule` ([#6811](https://github.com/hashicorp/terraform-provider-google-beta/pull/6811))
* **New Resource:** `google_clouddomains_registration` ([#6833](https://github.com/hashicorp/terraform-provider-google-beta/pull/6833))
* **New Resource:** `google_netapp_kmsconfig` ([#6831](https://github.com/hashicorp/terraform-provider-google-beta/pull/6831))
* **New Resource:** `google_vertex_ai_feature_online_store_featureview` ([#6821](https://github.com/hashicorp/terraform-provider-google-beta/pull/6821))
* **New Resource:** `google_vmwareengine_external_access_rule` ([#6811](https://github.com/hashicorp/terraform-provider-google-beta/pull/6811))

IMPROVEMENTS:
* compute: added `md5_authentication_key` field to `google_compute_router_peer` resource ([#6815](https://github.com/hashicorp/terraform-provider-google-beta/pull/6815))
* compute: added in-place update support to `params.resource_manager_tags` field in `google_compute_instance` resource ([#6828](https://github.com/hashicorp/terraform-provider-google-beta/pull/6828))
* compute: added in-place update support to `description` field in `google_compute_instance` resource ([#6804](https://github.com/hashicorp/terraform-provider-google-beta/pull/6804))
* gkehub: added `policycontroller` field to `google_gke_hub_feature_membership` resource ([#6813](https://github.com/hashicorp/terraform-provider-google-beta/pull/6813))
* gkehub2: added `clusterupgrade` field to `google_gke_hub_feature` resource ([#6836](https://github.com/hashicorp/terraform-provider-google-beta/pull/6836))
* gkeonprem: added in-place update support to `vsphere_config` field and added `host_groups` field in `google_gkeonprem_vmware_node_pool` resource ([#6802](https://github.com/hashicorp/terraform-provider-google-beta/pull/6802))
* iam: added `create_ignore_already_exists` field to `google_service_account` resource. If `ignore_create_already_exists` is set to true, resource creation would succeed when response error is 409 `ALREADY_EXISTS`. ([#6818](https://github.com/hashicorp/terraform-provider-google-beta/pull/6818))
* servicenetworking: added field `deletion_policy` to `google_service_networking_connection` ([#6830](https://github.com/hashicorp/terraform-provider-google-beta/pull/6830))
* sql: set `replica_configuration`, `ca_cert`, and `server_ca_cert` fields to be sensitive in `google_sql_instance` and `google_sql_ssl_cert` resources ([#6823](https://github.com/hashicorp/terraform-provider-google-beta/pull/6823))

BUG FIXES:
* bigquery: fixed perma-diff of `encryption_configuration` when API returns an empty object on `google_bigquery_table` resource ([#6817](https://github.com/hashicorp/terraform-provider-google-beta/pull/6817))
* compute: fixed an issue where the provider would `wait_for_instances` if set before deleting on `google_compute_instance_group_manager` and `google_compute_region_instance_group_manager` resources ([#6829](https://github.com/hashicorp/terraform-provider-google-beta/pull/6829))
* compute: fixed perma-diff that reordered `stateful_external_ip` and `stateful_internal_ip` blocks on `google_compute_instance_group_manager` and `google_compute_region_instance_group_manager` resources ([#6810](https://github.com/hashicorp/terraform-provider-google-beta/pull/6810))
* datapipeline: fixed perma-diff of `scheduler_service_account_email` when it's not explicitly specified in `google_data_pipeline_pipeline` resource ([#6814](https://github.com/hashicorp/terraform-provider-google-beta/pull/6814))
* edgecontainer: fixed resource import on `google_edgecontainer_vpn_connection` resource ([#6834](https://github.com/hashicorp/terraform-provider-google-beta/pull/6834))
* servicemanagement: fixed an issue where an inconsistent plan would be created when certain fields such as `openapi_config`, `grpc_config`, and `protoc_output_base64`, had computed values in `google_endpoints_service` resource ([#6832](https://github.com/hashicorp/terraform-provider-google-beta/pull/6832))
* storage: fixed an issue where retry timeout wasn't being utilized when creating `google_storage_bucket` resource ([#6806](https://github.com/hashicorp/terraform-provider-google-beta/pull/6806))

## 5.11.0 (Jan 08, 2024)

NOTES:
* compute: changed underlying actuation engine for `google_network_firewall_policy` and `google_region_network_firewall_policy`, there should be no user-facing impact ([#6776](https://github.com/hashicorp/terraform-provider-google-beta/pull/6776))
DEPRECATIONS:
* gkehub2: deprecated field `configmanagement.config_sync.oci.version` in `google_gke_hub_feature` resource ([#6764](https://github.com/hashicorp/terraform-provider-google-beta/pull/6764))

FEATURES:
* **New Data Source:** `google_compute_reservation` ([#6791](https://github.com/hashicorp/terraform-provider-google-beta/pull/6791))
* **New Resource:** `google_clouddeploy_automation` ([#6794](https://github.com/hashicorp/terraform-provider-google-beta/pull/6794))
* **New Resource:** `google_integration_connectors_endpoint_attachment` ([#6766](https://github.com/hashicorp/terraform-provider-google-beta/pull/6766))
* **New Resource:** `google_logging_folder_settings` ([#6754](https://github.com/hashicorp/terraform-provider-google-beta/pull/6754))
* **New Resource:** `google_logging_organization_settings` ([#6754](https://github.com/hashicorp/terraform-provider-google-beta/pull/6754))
* **New Resource:** `google_netapp_active_directory` ([#6781](https://github.com/hashicorp/terraform-provider-google-beta/pull/6781))
* **New Resource:** `google_vertex_ai_feature_online_store` ([#6779](https://github.com/hashicorp/terraform-provider-google-beta/pull/6779))
* **New Resource:** `google_vertex_ai_feature_group` ([#6780](https://github.com/hashicorp/terraform-provider-google-beta/pull/6780))
* **New Resource:** `google_netapp_backup_vault` ([#6793](https://github.com/hashicorp/terraform-provider-google-beta/pull/6793))

IMPROVEMENTS:
* bigqueryanalyticshub: added `restricted_export_config` field to `google_bigquery_analytics_hub_listing ` resource ([#6784](https://github.com/hashicorp/terraform-provider-google-beta/pull/6784))
* composer: added support for `composer_internal_ipv4_cidr_block` field to `google_composer_environment` ([#6761](https://github.com/hashicorp/terraform-provider-google-beta/pull/6761))
* composer: added `config.software_config.web_server_plugins_mode`, `config.workloads_config` and `dag_processor` fields to `google_composer_environment`. ([#6797](https://github.com/hashicorp/terraform-provider-google-beta/pull/6797))
* compute: added `provisioned_iops`and `provisioned_throughput` fields under `boot_disk.initialize_params` to `google_compute_instance` resource ([#6792](https://github.com/hashicorp/terraform-provider-google-beta/pull/6792))
* compute: added `resource_manager_tags` and `disk.resource_manager_tags` for `google_compute_instance_template` ([#6798](https://github.com/hashicorp/terraform-provider-google-beta/pull/6798))
* compute: added `resource_manager_tags` and `disk.resource_manager_tags` for `google_compute_region_instance_template` ([#6798](https://github.com/hashicorp/terraform-provider-google-beta/pull/6798))
* container: added `workload_alts_config` field to `google_container_cluster` resource ([#6762](https://github.com/hashicorp/terraform-provider-google-beta/pull/6762))
* dataproc: added `auxiliary_node_groups` field to `google_dataproc_cluster` resource ([#6753](https://github.com/hashicorp/terraform-provider-google-beta/pull/6753))
* edgecontainer: increased default timeout on `google_edgecontainer_cluster`, `google_edgecontainer_node_pool` to 480m from 60m ([#6796](https://github.com/hashicorp/terraform-provider-google-beta/pull/6796))
* gkehub2: added field `version` under `configmanagement` in `google_gke_hub_feature` resource ([#6764](https://github.com/hashicorp/terraform-provider-google-beta/pull/6764))
* kms: added output-only field `primary` to `google_kms_crypto_key` ([#6782](https://github.com/hashicorp/terraform-provider-google-beta/pull/6782))
* metastore: added `consumers.custom_routes_enabled` to `google_dataproc_metastore_service` ([#6767](https://github.com/hashicorp/terraform-provider-google-beta/pull/6767))
* sql: added support for IAM GROUP authentication in the `type` field of `google_sql_user` ([#6787](https://github.com/hashicorp/terraform-provider-google-beta/pull/6787))
* storagetransfer: made `name` field settable on `google_storage_transfer_job` ([#6777](https://github.com/hashicorp/terraform-provider-google-beta/pull/6777))

BUG FIXES:
* container: added check that `node_version` and `min_master_version` are the same on create of `google_container_cluster`, when running terraform plan ([#6763](https://github.com/hashicorp/terraform-provider-google-beta/pull/6763))
* container: fixed a bug where disabling PDCSI addon `gce_persistent_disk_csi_driver_config` during creation will result in permadiff in `google_container_cluster` resource ([#6751](https://github.com/hashicorp/terraform-provider-google-beta/pull/6751))
* container: fixed an issue in which migrating from the deprecated Binauthz enablement bool to the new evaluation mode enum inadvertently caused two cluster update events, instead of none. ([#6785](https://github.com/hashicorp/terraform-provider-google-beta/pull/6785))
* containerattached: fixed crash when updating a cluster to remove `admin_users` or `admin_groups` in `google_container_attached_cluster` ([#6786](https://github.com/hashicorp/terraform-provider-google-beta/pull/6786))
* dialogflowcx: fixed a permadiff in the `git_integration_settings` field of `google_diagflow_cx_agent` ([#6756](https://github.com/hashicorp/terraform-provider-google-beta/pull/6756))
* gkehub2: added field `version` under `configmanagement` in `google_gke_hub_feature` resource ([#6764](https://github.com/hashicorp/terraform-provider-google-beta/pull/6764))
* monitoring: fixed the index out of range crash in `dashboard_json` for the resource `google_monitoring_dashboard` ([#6750](https://github.com/hashicorp/terraform-provider-google-beta/pull/6750))

## 5.10.0 (Dec 18, 2023)

FEATURES:
* **New Data Source:** `google_compute_region_disk` ([#6726](https://github.com/hashicorp/terraform-provider-google-beta/pull/6726))
* **New Data Source:** `google_vmwareengine_external_address` ([#6714](https://github.com/hashicorp/terraform-provider-google-beta/pull/6714))
* **New Data Source:** `google_vmwareengine_subnet` ([#6715](https://github.com/hashicorp/terraform-provider-google-beta/pull/6715))
* **New Data Source:** `google_vmwareengine_vcenter_credentials` ([#6717](https://github.com/hashicorp/terraform-provider-google-beta/pull/6717))
* **New Resource:** `google_vmwareengine_external_address` ([#6714](https://github.com/hashicorp/terraform-provider-google-beta/pull/6714))
* **New Resource:** `google_vmwareengine_subnet` ([#6715](https://github.com/hashicorp/terraform-provider-google-beta/pull/6715))
* **New Resource:** `google_workbench_instance` ([#6739](https://github.com/hashicorp/terraform-provider-google-beta/pull/6739))
* **New Resource:** `google_workbench_instance_iam_*` ([#6739](https://github.com/hashicorp/terraform-provider-google-beta/pull/6739))

IMPROVEMENTS:
* bigquery: added `external_dataset_reference` field to `google_bigquery_dataset` resource ([#6716](https://github.com/hashicorp/terraform-provider-google-beta/pull/6716))
* compute: added `network_performance_config` field to `google_container_node_pool` resource to support GKE tier 1 networking ([#6719](https://github.com/hashicorp/terraform-provider-google-beta/pull/6719))
* compute: added `remove_instance_on_destroy` option to `google_compute_per_instance_config` resource ([#6724](https://github.com/hashicorp/terraform-provider-google-beta/pull/6724))
* compute: added `remove_instance_on_destroy` option to `google_compute_region_per_instance_config` resource ([#6724](https://github.com/hashicorp/terraform-provider-google-beta/pull/6724))
* container: added support for `network_performance_config.total_egress_bandwidth_tier` to support GKE tier 1 networking ([#6712](https://github.com/hashicorp/terraform-provider-google-beta/pull/6712))
* container: added support for in-place update for `machine_type`/`disk_type`/`disk_size_gb` in `google_container_node_pool` resource ([#6722](https://github.com/hashicorp/terraform-provider-google-beta/pull/6722))
* containerazure: added `config.labels` to `google_container_azure_node_pool` ([#6732](https://github.com/hashicorp/terraform-provider-google-beta/pull/6732))
* dataform: added `display_name`, `labels` and `npmrc_environment_variables_secret_version` fields to `google_dataform_repository` resource ([#6727](https://github.com/hashicorp/terraform-provider-google-beta/pull/6727))
* monitoring: added `severity` field to `google_monitoring_alert_policy` resource ([#6741](https://github.com/hashicorp/terraform-provider-google-beta/pull/6741))
* notebooks: added support for `labels` to `google_notebooks_runtime` ([#6746](https://github.com/hashicorp/terraform-provider-google-beta/pull/6746))
* orgpolicy: added `dry_run_spec` to `google_org_policy_policy` ([#6732](https://github.com/hashicorp/terraform-provider-google-beta/pull/6732))
* recaptchaenterprise: added `waf_settings` to `google_recaptcha_enterprise_key` ([#6732](https://github.com/hashicorp/terraform-provider-google-beta/pull/6732))
* securesourcemanager: added `host_config`, `state_note`, `kms_key`, and `private_config` fields to `google_secure_source_manager_instance` resource ([#6725](https://github.com/hashicorp/terraform-provider-google-beta/pull/6725))
* spanner: added `autoscaling_config.max_nodes` and `autoscaling_config.min_nodes` to `google_spanner_instance` ([#6748](https://github.com/hashicorp/terraform-provider-google-beta/pull/6748))
* storage: added `rpo` field to `google_storage_bucket` resource ([#6734](https://github.com/hashicorp/terraform-provider-google-beta/pull/6734))
* vmwareengine: added `type` field to `google_vmwareengine_private_cloud` resource ([#6744](https://github.com/hashicorp/terraform-provider-google-beta/pull/6744))
* workloadidentity: added `saml` block to `google_iam_workload_identity_pool_provider` resource ([#6718](https://github.com/hashicorp/terraform-provider-google-beta/pull/6718))

BUG FIXES:
* logging: fixed an issue where value change of `unique_writer_identity` on `google_logging_project_sink` does not trigger diff on dependent's usages of `writer_identity` ([#6742](https://github.com/hashicorp/terraform-provider-google-beta/pull/6742))

## 5.9.0 (Dec 11, 2023)

FEATURES:
* **New Data Source:** `google_logging_folder_settings` ([#6699](https://github.com/hashicorp/terraform-provider-google-beta/pull/6699))
* **New Data Source:** `google_logging_organization_settings` ([#6699](https://github.com/hashicorp/terraform-provider-google-beta/pull/6699))
* **New Data Source:** `google_logging_project_settings` ([#6699](https://github.com/hashicorp/terraform-provider-google-beta/pull/6699))
* **New Data Source:** `google_vmwareengine_network_policy` ([#6686](https://github.com/hashicorp/terraform-provider-google-beta/pull/6686))
* **New Data Source:** `google_vmwareengine_nsx_credentials` ([#6701](https://github.com/hashicorp/terraform-provider-google-beta/pull/6701))
* **New Resource:** `google_scc_event_threat_detection_custom_module` ([#6693](https://github.com/hashicorp/terraform-provider-google-beta/pull/6693))
* **New Resource:** `google_secure_source_manager_instance` ([#6685](https://github.com/hashicorp/terraform-provider-google-beta/pull/6685))
* **New Resource:** `google_vmwareengine_network_policy` ([#6686](https://github.com/hashicorp/terraform-provider-google-beta/pull/6686))

IMPROVEMENTS:
* bigqueryconnection: added `spark` support to `google_bigquery_connection` resource ([#6708](https://github.com/hashicorp/terraform-provider-google-beta/pull/6708))
* cloudidentity: added `expiry_detail` field to `google_cloud_identity_group_membership` resource ([#6689](https://github.com/hashicorp/terraform-provider-google-beta/pull/6689))
* container: added `queued_provisioning` field to `google_container_node_pool` resource ([#6678](https://github.com/hashicorp/terraform-provider-google-beta/pull/6678))
* gkehub: added `default_cluster_config` field to `google_gke_hub_fleet` resource ([#6683](https://github.com/hashicorp/terraform-provider-google-beta/pull/6683))
* gkehub: added `binary_authorization_config` field to `google_gke_hub_fleet` resource ([#6705](https://github.com/hashicorp/terraform-provider-google-beta/pull/6705))
* sql: added support for in-place updates to the `edition` field in `google_sql_database_instance` resource ([#6681](https://github.com/hashicorp/terraform-provider-google-beta/pull/6681))

BUG FIXES:
* artifactregistry: fixed permadiff due to unsorted `virtual_repository_config` array in `google_artifact_registry_repository` ([#6691](https://github.com/hashicorp/terraform-provider-google-beta/pull/6691))
* container: made `dns_config` field updatable on `google_container_cluster` resource ([#6695](https://github.com/hashicorp/terraform-provider-google-beta/pull/6695))
* dlp: added conflicting field validation in the `storage_config.timespan_config` block in `data_loss_prevention_job_trigger` resource ([#6680](https://github.com/hashicorp/terraform-provider-google-beta/pull/6680))
* dlp: updated the `storage_config.timespan_config.timestamp_field` field in `data_loss_prevention_job_trigger` to be optional ([#6680](https://github.com/hashicorp/terraform-provider-google-beta/pull/6680))
* firestore: added retries during creation of `google_firestore_index` resources to address retryable 409 code API errors ("Please retry, underlying data changed", and "Aborted due to cross-transaction contention") ([#6677](https://github.com/hashicorp/terraform-provider-google-beta/pull/6677), [#6702](https://github.com/hashicorp/terraform-provider-google-beta/pull/6702))
* storage: fixed unexpected `lifecycle_rule` conditions being added for `google_storage_bucket` ([#6711](https://github.com/hashicorp/terraform-provider-google-beta/pull/6711))

## 5.8.0 (Dec 4, 2023)

FEATURES:
* **New Data Source:** `google_vmwareengine_network_peering` ([#6675](https://github.com/hashicorp/terraform-provider-google-beta/pull/6675))
* **New Resource:** `google_dataform_repository_iam_*` (beta) ([#6648](https://github.com/hashicorp/terraform-provider-google-beta/pull/6648))
* **New Resource:** `google_migration_center_group` ([#6651](https://github.com/hashicorp/terraform-provider-google-beta/pull/6651))
* **New Resource:** `google_netapp_storage_pool` ([#6663](https://github.com/hashicorp/terraform-provider-google-beta/pull/6663))
* **New Resource:** `google_vertex_ai_endpoint_iam_*` (beta) ([#6657](https://github.com/hashicorp/terraform-provider-google-beta/pull/6657))
* **New Resource:** `google_vmwareengine_network_peering` ([#6675](https://github.com/hashicorp/terraform-provider-google-beta/pull/6675))

IMPROVEMENTS:
* artifactregistry: added `remote_repository_config.upstream_credentials` field to `google_artifact_registry_repository` resource ([#6658](https://github.com/hashicorp/terraform-provider-google-beta/pull/6658))
* cloudbuild: added fields `build.artifacts.maven_artifacts`, `build.artifacts.npm_packages `, and `build.artifacts.python_packages ` to resource `google_cloudbuild_trigger` ([#6650](https://github.com/hashicorp/terraform-provider-google-beta/pull/6650)
* composer: added `database_config.zone` field in `google_composer_environment` ([#6653](https://github.com/hashicorp/terraform-provider-google-beta/pull/6653))
* compute: added field `service_directory_registrations` to resource `google_compute_global_forwarding_rule` ([#6667](https://github.com/hashicorp/terraform-provider-google-beta/pull/6667))
* firestore: added virtual field `deletion_policy` to `google_firestore_database` ([#6664](https://github.com/hashicorp/terraform-provider-google-beta/pull/6664))
* firestore: enabled database deletion upon destroy for `google_firestore_database` ([#6664](https://github.com/hashicorp/terraform-provider-google-beta/pull/6664))
* gkehub2: added `policycontroller` field to `fleet_default_member_config` in `google_gke_hub_feature` ([#6649](https://github.com/hashicorp/terraform-provider-google-beta/pull/6649))
* iam: added `allowed_services`, `disable_programmatic_signin` fields to `google_iam_workforce_pool` resource ([#6666](https://github.com/hashicorp/terraform-provider-google-beta/pull/6666))
* vmwareengine: added `STANDARD` type support to `google_vmwareengine_network` resource ([#6669](https://github.com/hashicorp/terraform-provider-google-beta/pull/6669))

BUG FIXES:
* compute: fixed a permadiff caused by issues with ipv6 diff suppression in `google_compute_forwarding_rule` and `google_compute_global_forwarding_rule` ([#6652](https://github.com/hashicorp/terraform-provider-google-beta/pull/6652))
* firestore: fixed an issue where `google_firestore_database` could be deleted when `delete_protection_state` was `DELETE_PROTECTION_ENABLED` ([#6664](https://github.com/hashicorp/terraform-provider-google-beta/pull/6664))
* firestore: made resource creation retry for 409 errors with the text "Aborted due to cross-transaction contention" in `google_firestore_index ` ([#6677](https://github.com/hashicorp/terraform-provider-google-beta/pull/6677))

## 5.7.0 (Nov 20, 2023)

DEPRECATIONS:
* gkehub: deprecated `config_management.binauthz` in `google_gke_hub_feature_membership` ([#6646](https://github.com/hashicorp/terraform-provider-google-beta/pull/6646))

IMPROVEMENTS:
* bigtable: added `standard_isolation` and `standard_isolation.priority` fields to `google_bigtable_app_profile` resource ([#6621](https://github.com/hashicorp/terraform-provider-google-beta/pull/6621))
* containerattached: added `proxy_config` field to `google_container_attached_cluster` resource ([#6637](https://github.com/hashicorp/terraform-provider-google-beta/pull/6637))
* gkehub: added `membership_location` field to `google_gke_hub_feature_membership` resource ([#6646](https://github.com/hashicorp/terraform-provider-google-beta/pull/6646))
* logging: made the change to aqcuire and update the `google_logging_project_sink` resource that already exists at the desired location. These logging buckets cannot be removed so deleting this resource will remove the bucket config from your terraform state but will leave the logging bucket unchanged. ([#6632](https://github.com/hashicorp/terraform-provider-google-beta/pull/6632))
* memcache: added `MEMCACHE_1_6_15` as a possible value for `memcache_version` in `google_memcache_instance` resource ([#6642](https://github.com/hashicorp/terraform-provider-google-beta/pull/6642))
* monitoring: added error message to delete Alert Policies first on 400 response when deleting `google_monitoring_uptime_check_config` resource ([#6645](https://github.com/hashicorp/terraform-provider-google-beta/pull/6645))
* spanner: added `autoscaling_config` field to  `google_spanner_instance` resource ([#6616](https://github.com/hashicorp/terraform-provider-google-beta/pull/6616))

BUG FIXES:
* compute: changed `external_ipv6_prefix` field to not be output only in `google_compute_subnetwork` resource ([#6619](https://github.com/hashicorp/terraform-provider-google-beta/pull/6619))
* compute: fixed issue where `google_compute_attached_disk` would produce an error for certain zone configs ([#6620](https://github.com/hashicorp/terraform-provider-google-beta/pull/6620))
* edgecontainer: fixed update method of `google_edgecontainer_cluster` resource ([#6625](https://github.com/hashicorp/terraform-provider-google-beta/pull/6625))
* provider: fixed an issue where universe domains would not overwrite API endpoints ([#6636](https://github.com/hashicorp/terraform-provider-google-beta/pull/6636))
* resourcemanager: made `data_source_google_project_service` no longer return an error when the service is not enabled ([#6638](https://github.com/hashicorp/terraform-provider-google-beta/pull/6638))
* sql: `ssl_mode` field is not stored in terraform state if it has never been used in `google_sql_database_instance` resource ([#6622](https://github.com/hashicorp/terraform-provider-google-beta/pull/6622))

NOTES:
* dataproc: backfilled `terraform_labels` field for resource `google_dataproc_workflow_template`, so the resource recreation won't happen during provider upgrade from `4.x` to `5.7` ([#6634](https://github.com/hashicorp/terraform-provider-google-beta/pull/6634))
* provider: backfilled `terraform_labels` for some immutable resources, so the resource recreation won't happen during provider upgrade from `4.X` to `5.7` ([#6635](https://github.com/hashicorp/terraform-provider-google-beta/pull/6635))

## 5.6.0 (Nov 13, 2023)

FEATURES:
* **New Data Source:** `google_backup_dr_management_server` ([#6596](https://github.com/hashicorp/terraform-provider-google-beta/pull/6596))
* **New Resource:** `google_compute_instance_settings` ([#6615](https://github.com/hashicorp/terraform-provider-google-beta/pull/6615))
* **New Resource:** `google_integration_connectors_connection`([#6612](https://github.com/hashicorp/terraform-provider-google-beta/pull/6612))

IMPROVEMENTS:
* assuredworkloads: added `enable_sovereign_controls`, `partner`, `partner_permissions`, `violation_notifications_enabled`, and several other output-only fields to `google_assured_workloads_workloads` ([#6597](https://github.com/hashicorp/terraform-provider-google-beta/pull/6597))
* composer: added `storage_config` to `google_composer_environment` ([#6606](https://github.com/hashicorp/terraform-provider-google-beta/pull/6606))
* container: added `fleet` field to `google_container_cluster` resource ([#6610](https://github.com/hashicorp/terraform-provider-google-beta/pull/6610))
* containeraws: added `admin_groups` to `google_container_aws_cluster` ([#6597](https://github.com/hashicorp/terraform-provider-google-beta/pull/6597))
* containerazure: added `admin_groups` to `google_container_azure_cluster` ([#6597](https://github.com/hashicorp/terraform-provider-google-beta/pull/6597))
* dataproc: added support for `instance_flexibility_policy` in `google_dataproc_cluster` ([#6593](https://github.com/hashicorp/terraform-provider-google-beta/pull/6593))
* dialogflowcx: added `is_default_start_flow` field to `google_dialogflow_cx_flow` resource to allow management of default flow resources via Terraform ([#6600](https://github.com/hashicorp/terraform-provider-google-beta/pull/6600))
* dialogflowcx: added `is_default_welcome_intent` and `is_default_negative_intent` fields to `google_dialogflow_cx_intent` resource to allow management of default intent resources via Terraform ([#6600](https://github.com/hashicorp/terraform-provider-google-beta/pull/6600))
* gkehub: added `fleet_default_member_config` field to `google_gke_hub_feature` resource ([#6608](https://github.com/hashicorp/terraform-provider-google-beta/pull/6608))
* gkehub: added `metrics_gcp_service_account_email` to `google_gke_hub_feature_membership` ([#6597](https://github.com/hashicorp/terraform-provider-google-beta/pull/6597))
* logging: added `index_configs` field to `logging_bucket_config` resource ([#6598](https://github.com/hashicorp/terraform-provider-google-beta/pull/6598))
* logging: added `index_configs` field to `logging_project_bucket_config` resource ([#6598](https://github.com/hashicorp/terraform-provider-google-beta/pull/6598))
* monitoring: added `pings_count`, `user_labels`, and `custom_content_type` fields to `google_monitoring_uptime_check_config` resource ([#6594](https://github.com/hashicorp/terraform-provider-google-beta/pull/6594))
* spanner: added `autoscaling_config` field to  `google_spanner_instance` ([#6616](https://github.com/hashicorp/terraform-provider-google-beta/pull/6616))
* sql: added `ssl_mode` field to `google_sql_database_instance` resource ([#6579](https://github.com/hashicorp/terraform-provider-google-beta/pull/6579))
* vertexai: added `private_service_connect_config` to `google_vertex_ai_index_endpoint` ([#6614](https://github.com/hashicorp/terraform-provider-google-beta/pull/6614))
* workstations: added `domain_config` field to resource `google_workstations_workstation_cluster` (beta) ([#6609](https://github.com/hashicorp/terraform-provider-google-beta/pull/6609))

BUG FIXES:
* provider: made `terraform_labels` immutable in immutable resources to not block the upgrade. This will create a Terraform plan that recreates the resource on `4.X` -> `5.6.0` upgrade for affected resources. A mitigation to backfill the values during the upgrade is planned, and will release resource-by-resource. ([#6613](https://github.com/hashicorp/terraform-provider-google-beta/pull/6613))

## 5.5.0 (Nov 06, 2023)

FEATURES:
* **New Data Source:** `google_bigquery_dataset` ([#6570](https://github.com/hashicorp/terraform-provider-google-beta/pull/6570))

IMPROVEMENTS:
* alloydb: added `SECONDARY` as an option for `instance_type` field in `google_alloydb_instance` resource, to support creation of secondary instance inside a secondary cluster. ([#6583](https://github.com/hashicorp/terraform-provider-google-beta/pull/6583))
* alloydb: added `deletion_policy` field to `google_alloydb_cluster` resource, to allow force-destroying instances along with their cluster. This is necessary to delete secondary instances, which cannot be deleted otherwise. ([#6583](https://github.com/hashicorp/terraform-provider-google-beta/pull/6583))
* alloydb: added support to promote `google_alloydb_cluster` resources from secondary to primary ([#6589](https://github.com/hashicorp/terraform-provider-google-beta/pull/6589))
* alloydb: increased default timeout on `google_alloydb_instance` to 120m from 40m ([#6583](https://github.com/hashicorp/terraform-provider-google-beta/pull/6583))
* dataproc: added `instance_flexibility_policy` field ro `google_dataproc_cluster` resource ([#6593](https://github.com/hashicorp/terraform-provider-google-beta/pull/6593))
* monitoring: added `subject` field to `google_monitoring_alert_policy` resource ([#6590](https://github.com/hashicorp/terraform-provider-google-beta/pull/6590))
* storage: added `enable_object_retention` field to `google_storage_bucket` resource ([#6588](https://github.com/hashicorp/terraform-provider-google-beta/pull/6588))
* storage: added `retention` field to `google_storage_bucket_object` resource ([#6588](https://github.com/hashicorp/terraform-provider-google-beta/pull/6588))
* workflows: added `user_env_vars` field to `google_workflows_workflow` resource ([#6567](https://github.com/hashicorp/terraform-provider-google-beta/pull/6567))

BUG FIXES:
* compute: fixed an error when `maintenance_interval` is updated on `google_compute_instance_template` ([#6569](https://github.com/hashicorp/terraform-provider-google-beta/pull/6569))
* firestore: fixed an issue with creation of multiple `google_firestore_field` resources ([#6572](https://github.com/hashicorp/terraform-provider-google-beta/pull/6572))

## 5.4.0 (Oct 30, 2023)

DEPRECATIONS:
* bigquery: deprecated `cloud_spanner.use_serverless_analytics` on `google_bigquery_connection`. Use `cloud_spanner.use_data_boost` instead. ([#6539](https://github.com/hashicorp/terraform-provider-google-beta/pull/6539))

NOTES:
* provider: added `universe_domain` attribute as a provider attribute ([#6551](https://github.com/hashicorp/terraform-provider-google-beta/pull/6551))

BREAKING CHANGES:
* cloudrunv2: marked `location` field as required in resource `google_cloud_run_v2_job`. Any configuration without `location` specified will cause resource creation fail ([#6540](https://github.com/hashicorp/terraform-provider-google-beta/pull/6540))
* cloudrunv2: marked `location` field as required in resource `google_cloud_run_v2_service`. Any configuration without `location` specified will cause resource creation fail ([#6540](https://github.com/hashicorp/terraform-provider-google-beta/pull/6540))

FEATURES:
* **New Data Source:** `google_cloud_identity_group_lookup` ([#6530](https://github.com/hashicorp/terraform-provider-google-beta/pull/6530))
* **New Resource:** `google_network_connectivity_policy_based_route` ([#6552](https://github.com/hashicorp/terraform-provider-google-beta/pull/6552))
* **New Resource:** `google_pubsub_schema_iam_*` ([#6533](https://github.com/hashicorp/terraform-provider-google-beta/pull/6533))

IMPROVEMENTS:
* accesscontextmanager: added support for specifying `vpc_network_sources` to `google_access_context_manager_access_levels`, `google_access_context_manager_access_level`, and `google_access_context_manager_access_level_condition` ([#6553](https://github.com/hashicorp/terraform-provider-google-beta/pull/6553))
* apigee: added support for `type` in `google_apigee_environment` ([#6562](https://github.com/hashicorp/terraform-provider-google-beta/pull/6562))
* bigquery: added `cloud_spanner.database_role`, `cloud_spanner.use_data_boost`, and `cloud_spanner.max_parallelism` fields to `google_bigquery_connection` ([#6539](https://github.com/hashicorp/terraform-provider-google-beta/pull/6539))
* bigquery: added support for `iam_member` to `google_bigquery_dataset.access` ([#6550](https://github.com/hashicorp/terraform-provider-google-beta/pull/6550))
* compute: added `maintenance_interval` field to `google_compute_node_group` resource ([#6561](https://github.com/hashicorp/terraform-provider-google-beta/pull/6561))
* container: added `enable_confidential_storage` to `node_config` in `google_container_cluster` and `google_container_node_pool` ([#6531](https://github.com/hashicorp/terraform-provider-google-beta/pull/6531))
* container: added update support for `google_container_node_pool.node_config.taint` ([#6536](https://github.com/hashicorp/terraform-provider-google-beta/pull/6536))
* containerattached: added `admin_groups` field to `google_container_attached_cluster` resource ([#6537](https://github.com/hashicorp/terraform-provider-google-beta/pull/6537))
* dialogflowcx: added `advanced_settings` field to `google_dialogflow_cx_flow` resource ([#6543](https://github.com/hashicorp/terraform-provider-google-beta/pull/6543))
* dialogflowcx: added `advanced_settings` fields to `google_dialogflow_cx_page` resource ([#6543](https://github.com/hashicorp/terraform-provider-google-beta/pull/6543))
* dialogflowcx: added `advanced_settings`, `text_to_speech_settings`, `git_integration_settings` fields to `google_dialogflow_cx_agent` resource ([#6543](https://github.com/hashicorp/terraform-provider-google-beta/pull/6543))
* tpuv2: added `cidr_block`, `labels`, `tags`, `network_config`, `scheduling_config`, `shielded_instance_config`, `service_account` and `data_disks` fields to `google_tpu_v2_vm` ([#6555](https://github.com/hashicorp/terraform-provider-google-beta/pull/6555))
* tpuv2: added `accelerator_config` field to `google_tpu_v2_vm` resource ([#6559](https://github.com/hashicorp/terraform-provider-google-beta/pull/6559))

BUG FIXES:
* bigquery: fixed a bug when updating a `google_bigquery_dataset` that contained an `iamMember` access rule added out of band with Terraform ([#6550](https://github.com/hashicorp/terraform-provider-google-beta/pull/6550))
* bigqueryreservation: fixed bug of incorrect resource recreation when `capacity_commitment_id` is unspecified in resource `google_bigquery_capacity_commitment` ([#6548](https://github.com/hashicorp/terraform-provider-google-beta/pull/6548))
* cloudrunv2: made `annotations` field on the `google_cloud_run_v2_job` data source include all annotations present on the resource in GCP ([#6532](https://github.com/hashicorp/terraform-provider-google-beta/pull/6532))
* cloudrunv2: made `annotations` field on the `google_cloud_run_v2_service` data source include all annotations present on the resource in GCP ([#6532](https://github.com/hashicorp/terraform-provider-google-beta/pull/6532))
* cloudrunv2: made `labels` and `terraform labels` fields on the `google_cloud_run_v2_job` data source include all annotations present on the resource in GCP ([#6532](https://github.com/hashicorp/terraform-provider-google-beta/pull/6532))
* cloudrunv2: made `labels` and `terraform labels` fields on the `google_cloud_run_v2_service` data source include all annotations present on the resource in GCP ([#6532](https://github.com/hashicorp/terraform-provider-google-beta/pull/6532))
* edgecontainer: fixed an issue where the update endpoint for `google_edgecontainer_cluster` was incorrect. ([#6560](https://github.com/hashicorp/terraform-provider-google-beta/pull/6560))
* redis: allow `replica_count` to be set to zero in the `google_redis_cluster` resource ([#6534](https://github.com/hashicorp/terraform-provider-google-beta/pull/6534))

## 5.3.0 (Oct 23, 2023)

DEPRECATIONS:
* bigquery: deprecated `time_partitioning.require_partition_filter` in favor of new top level field `require_partition_filter` in resource `google_bigquery_table` ([#6496](https://github.com/hashicorp/terraform-provider-google-beta/pull/6496))

FEATURES:
* **New Data Source:** `google_cloud_run_v2_job` ([#6508](https://github.com/hashicorp/terraform-provider-google-beta/pull/6508))
* **New Data Source:** `google_cloud_run_v2_service` ([#6527](https://github.com/hashicorp/terraform-provider-google-beta/pull/6527))
* **New Data Source:** `google_compute_networks` ([#6498](https://github.com/hashicorp/terraform-provider-google-beta/pull/6498))

IMPROVEMENTS:
* cloudidentity: added `additional_group_keys` attribute to `google_cloud_identity_group` resource ([#6504](https://github.com/hashicorp/terraform-provider-google-beta/pull/6504))
* compute: added `enable_confidential_compute` field under `boot_disk.0.initialize_params` in `google_compute_instance` ([#6528](https://github.com/hashicorp/terraform-provider-google-beta/pull/6528))
* compute: added `internal_ipv6_range` to `google_compute_network` data source and `internal_ipv6_prefix` field to `data.google_compute_subnetwork` data source ([#6514](https://github.com/hashicorp/terraform-provider-google-beta/pull/6514))
* container: added support for `security_posture_config.vulnerability_mode` value `VULNERABILITY_ENTERPRISE` in `google_container_cluster` ([#6520](https://github.com/hashicorp/terraform-provider-google-beta/pull/6520))
* dataform: added `ssh_authentication_config` and `service_account` to `google_dataform_repository` resource ([#6480](https://github.com/hashicorp/terraform-provider-google-beta/pull/6480))
* dataproc: added `min_num_instances` field to `google_dataproc_cluster` resource ([#6503](https://github.com/hashicorp/terraform-provider-google-beta/pull/6503))
* logging: added `custom_writer_identity` field to `google_logging_project_sink` ([#6486](https://github.com/hashicorp/terraform-provider-google-beta/pull/6486))
* secretmanager: made `ttl` field mutable in `google_secret_manager_secret` ([#6521](https://github.com/hashicorp/terraform-provider-google-beta/pull/6521))
* storage: added `terminal_storage_class` to the `autoclass` field in `google_storage_bucket` resource ([#6519](https://github.com/hashicorp/terraform-provider-google-beta/pull/6519))

BUG FIXES:
* bigquerydatatransfer: fixed an error when updating `google_bigquery_data_transfer_config` related to incorrect update masks ([#6516](https://github.com/hashicorp/terraform-provider-google-beta/pull/6516))
* cloudrunv2: fixed a bug where `google_cloud_run_v2_service.custom_audiences` could not be set or updated properly ([#6482](https://github.com/hashicorp/terraform-provider-google-beta/pull/6482))
* compute: fixed an error during the deletion when post was set to 0 on `google_compute_global_network_endpoint` ([#6523](https://github.com/hashicorp/terraform-provider-google-beta/pull/6523))
* compute: fixed an issue with TTLs being sent for `google_compute_backend_service` when `cache_mode` is set to `USE_ORIGIN_HEADERS` ([#6499](https://github.com/hashicorp/terraform-provider-google-beta/pull/6499))
* container: fixed an issue where empty `autoscaling` block would crash the provider for `google_container_node_pool` ([#6483](https://github.com/hashicorp/terraform-provider-google-beta/pull/6483))
* dataflow: fixed a bug where resource updates returns an error if only `labels` has changes for batch `google_dataflow_job` and `google_dataflow_flex_template_job` ([#6502](https://github.com/hashicorp/terraform-provider-google-beta/pull/6502))
* dialogflowcx: fixed updating `google_dialogflow_cx_version`; updates will no longer time out. ([#6484](https://github.com/hashicorp/terraform-provider-google-beta/pull/6484))
* sql: fixed a bug where adding the `edition` field to a `google_sql_database_instance` resource that already existed and used ENTERPRISE edition resulted in a permant diff in plans ([#6485](https://github.com/hashicorp/terraform-provider-google-beta/pull/6485))
* sql: removed host validation to support IP address and DNS address in host in `google_sql_source_representation_instance` resource ([#6493](https://github.com/hashicorp/terraform-provider-google-beta/pull/6493))

## 5.2.0 (Oct 16, 2023)

FEATURES:
* **New Data Source:** `google_secret_manager_secrets` ([#6463](https://github.com/hashicorp/terraform-provider-google-beta/pull/6463))
* **New Resource:** `google_alloydb_user` ([#6454](https://github.com/hashicorp/terraform-provider-google-beta/pull/6454))
* **New Resource:** `google_firestore_backup_schedule` ([#6465](https://github.com/hashicorp/terraform-provider-google-beta/pull/6465))

IMPROVEMENTS:
* alloydb: added `cluster_type` and `secondary_config` fields to support secondary clusters in `google_alloydb_cluster` resource. ([#6474](https://github.com/hashicorp/terraform-provider-google-beta/pull/6474))
* compute: added `recreate_closed_psc` flag to support recreating the PSC Consumer forwarding rule if the `psc_connection_status` is closed on `google_compute_forwarding_rule`. ([#6468](https://github.com/hashicorp/terraform-provider-google-beta/pull/6468))
* compute: added `INTERNET_IP_PORT`, `INTERNET_FQDN_PORT`, `SERVERLESS`, and `PRIVATE_SERVICE_CONNECT` as acceptable values for the `network_endpoint_type` field for the `resource_compute_network_endpoint_group` resource ([#6472](https://github.com/hashicorp/terraform-provider-google-beta/pull/6472))
* compute: added `SEV_LIVE_MIGRATABLE_V2` to `guest_os_features` enum on `google_compute_image` resource. ([#6466](https://github.com/hashicorp/terraform-provider-google-beta/pull/6466))
* compute: added `allow_subnet_cidr_routes_overlap` field to `google_compute_subnetwork` resource ([#6445](https://github.com/hashicorp/terraform-provider-google-beta/pull/6445))
* dataform: added `ssh_authentication_config` and `service_account` to `google_dataform_repository` resource ([#6480](https://github.com/hashicorp/terraform-provider-google-beta/pull/6480))

BUG FIXES:
* alloydb: added `client_connection_config` field to `google_alloydb_instance` resource ([#6478](https://github.com/hashicorp/terraform-provider-google-beta/pull/6478))
* bigquery: removed mutual exclusivity checks for `view`, `materialized_view`, and `schema` for the `google_bigquery_table` resource ([#6471](https://github.com/hashicorp/terraform-provider-google-beta/pull/6471))
* compute: added `certificate_manager_certificates` field to `google_compute_target_https_proxy` resource ([#6460](https://github.com/hashicorp/terraform-provider-google-beta/pull/6460))
* compute: added validation to prevent setting empty `rule.action.source_nat_active_ranges` to `google_compute_router_nat` resource ([#6467](https://github.com/hashicorp/terraform-provider-google-beta/pull/6467))
* compute: fixed an issue where external `google_compute_global_address` can't be created when `network_tier` in `google_compute_project_default_network_tier` is set to `STANDARD` ([#6456](https://github.com/hashicorp/terraform-provider-google-beta/pull/6456))
* compute: fixed a false permadiff on `ip_address` when it is set to ipv6 on `google_compute_forwarding_rule` ([#6444](https://github.com/hashicorp/terraform-provider-google-beta/pull/6444))
* provider: fixed a bug where an update request was sent to services when updateMask is empty ([#6443](https://github.com/hashicorp/terraform-provider-google-beta/pull/6443))
* securitypolicy: fixed a bug where setting `advanced_options_config.user_ip_request_headers` field with empty value was not cleaning the list ([#6470](https://github.com/hashicorp/terraform-provider-google-beta/pull/6470))

## 5.1.0 (Oct 9, 2023)

FEATURES:
* **New Resource:** `google_database_migration_service_private_connection` ([#6436](https://github.com/hashicorp/terraform-provider-google-beta/pull/6436))))
* **New Resource:** `google_edgecontainer_cluster` ([#6406](https://github.com/hashicorp/terraform-provider-google-beta/pull/6406))
* **New Resource:** `google_edgecontainer_node_pool` ([#6406](https://github.com/hashicorp/terraform-provider-google-beta/pull/6406))
* **New Resource:** `google_edgecontainer_vpn_connection` ([#6406](https://github.com/hashicorp/terraform-provider-google-beta/pull/6406))
* **New Resource:** `google_firebase_hosting_custom_domain` ([#6409](https://github.com/hashicorp/terraform-provider-google-beta/pull/6409))
* **New Resource:** `google_gke_hub_fleet` ([#6417](https://github.com/hashicorp/terraform-provider-google-beta/pull/6417))

IMPROVEMENTS:
* compute: added `device_name` field to `scratch_disk` block of `google_compute_instance` resource ([#6401](https://github.com/hashicorp/terraform-provider-google-beta/pull/6401))
* container: added `node_config.linux_node_config.cgroup_mode` field to `google_container_node_pool` ([#6435](https://github.com/hashicorp/terraform-provider-google-beta/pull/6435))
* databasemigrationservice: added support for `oracle` profiles to `google_database_migration_service_connection_profile` ([#6426](https://github.com/hashicorp/terraform-provider-google-beta/pull/6426))
* firestore: added `api_scope` field to `google_firestore_index` resource ([#6424](https://github.com/hashicorp/terraform-provider-google-beta/pull/6424))
* gkehub: added `location` field to `google_gke_hub_membership_iam_*` resources ([#6437](https://github.com/hashicorp/terraform-provider-google-beta/pull/6437))
* gkehub: added `location` field to `google_gke_hub_membership` resource ([#6437](https://github.com/hashicorp/terraform-provider-google-beta/pull/6437))
* gkeonprem: added update-in-place support for `vcenter` fields in `google_gkeonprem_vmware_cluster` ([#6418](https://github.com/hashicorp/terraform-provider-google-beta/pull/6418))
* identityplatform: added `sms_region_config` to the resource `google_identity_platform_config` ([#6398](https://github.com/hashicorp/terraform-provider-google-beta/pull/6398))

BUG FIXES:
* dns: fixed record set configuration parsing in `google_dns_record_set` ([#6397](https://github.com/hashicorp/terraform-provider-google-beta/pull/6397))
* provider: fixed an issue where the plugin-framework implementation of the provider handled default region values that were self-links differently to the SDK implementation. This issue is not believed to have affected users because of downstream functions that turn self links into region names. ([#6432](https://github.com/hashicorp/terraform-provider-google-beta/pull/6432))
* provider: fixed a bug that caused update requests to be sent for resources with a `terraform_labels` field even if no fields were updated ([#6443](https://github.com/hashicorp/terraform-provider-google-beta/pull/6443))

## 5.0.0 (Oct 2, 2023)

KNOWN ISSUES:

* Updating some resources post-upgrade results in an error like "The update_mask in the Update{{Resource}}Request must be set". This should be resolved in `5.1.0`, see https://github.com/hashicorp/terraform-provider-google/issues/16091 for details.

[Terraform Google Provider 5.0.0 Upgrade Guide](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/version_5_upgrade)

NOTES:
* provider: some provider default values are now shown at plan-time ([#6188](https://github.com/hashicorp/terraform-provider-google-beta/pull/6188))

LABELS REWORK:
* provider: default labels configured on the provider through the new `default_labels` field are now supported. The default labels configured on the provider will be applied to all of the resources with standard `labels` field.
* provider: resources with labels - three label-related fields are now in all of the resources with standard `labels` field. `labels` field is non-authoritative and only manages the labels defined by the users on the resource through Terraform. The new output-only `terraform_labels` field merges the labels defined by the users on the resource through Terraform and the default labels configured on the provider. The new output-only `effective_labels` field lists all of labels present on the resource in GCP, including the labels configured through Terraform, the system, and other clients.
* provider: resources with annotations - two annotation-related fields are now in all of the resources with standard `annotations` field. The `annotations` field is non-authoritative and only manages the annotations defined by the users on the resource through Terraform. The new output-only `effective_annotations` field lists all of annotations present on the resource in GCP, including the annotations configured through Terraform, the system, and other clients.
* provider: datasources with labels - three fields `labels`, `terraform_labels`, and `effective_labels` are now present in most resource-based datasources. All three fields have all of labels present on the resource in GCP including the labels configured through Terraform, the system, and other clients, equivalent to `effective_labels` on the resource.
* provider: datasources with annotations - both `annotations` and `effective_annotations` are now present in most resource-based datasources. Both fields have all of annotations present on the resource in GCP including the annotations configured through Terraform, the system, and other clients, equivalent to `effective_annotations` on the resource.

BREAKING CHANGES:
* provider: added provider-level validation so these fields are not set as empty strings in a user's config: `credentials`, `access_token`, `impersonate_service_account`, `project`, `billing_project`, `region`, `zone` ([#6358](https://github.com/hashicorp/terraform-provider-google-beta/pull/6358))
* provider: fixed many import functions throughout the provider that matched a subset of the provided input when possible. Now, the GCP resource id supplied to "terraform import" must match exactly. ([#6364](https://github.com/hashicorp/terraform-provider-google-beta/pull/6364))
* provider: made data sources return errors on 404s when applicable instead of silently failing ([#6241](https://github.com/hashicorp/terraform-provider-google-beta/pull/6241))
* provider: made empty strings in the provider configuration block no longer be ignored when configuring the provider ([#6358](https://github.com/hashicorp/terraform-provider-google-beta/pull/6358))
* accesscontextmanager: changed multiple array fields to sets where appropriate to prevent duplicates and fix diffs caused by server side reordering. ([#6217](https://github.com/hashicorp/terraform-provider-google-beta/pull/6217))
* bigquery: added more input validations for `google_bigquery_table` schema ([#5975](https://github.com/hashicorp/terraform-provider-google-beta/pull/5975))
* bigquery: made `routine_type` required for `google_bigquery_routine` ([#6080](https://github.com/hashicorp/terraform-provider-google-beta/pull/6080))
* cloudfunction2: made `location` required on `google_cloudfunctions2_function` ([#6260](https://github.com/hashicorp/terraform-provider-google-beta/pull/6260))
* cloudiot: removed deprecated datasource `google_cloudiot_registry_iam_policy` ([#6206](https://github.com/hashicorp/terraform-provider-google-beta/pull/6206))
* cloudiot: removed deprecated resource `google_cloudiot_device` ([#6206](https://github.com/hashicorp/terraform-provider-google-beta/pull/6206))
* cloudiot: removed deprecated resource  `google_cloudiot_registry` ([#6206](https://github.com/hashicorp/terraform-provider-google-beta/pull/6206))
* cloudiot: removed deprecated resource `google_cloudiot_registry_iam_*` ([#6206](https://github.com/hashicorp/terraform-provider-google-beta/pull/6206))
* cloudrunv2: removed deprecated field `liveness_probe.tcp_socket` from `google_cloud_run_v2_service` resource. ([#6029](https://github.com/hashicorp/terraform-provider-google-beta/pull/6029))
* cloudrunv2: removed deprecated fields `startup_probe` and `liveness_probe` from `google_cloud_run_v2_job` resource. ([#6029](https://github.com/hashicorp/terraform-provider-google-beta/pull/6029))
* cloudrunv2: retyped `volumes.cloud_sql_instance.instances` to SET from ARRAY for `google_cloud_run_v2_service` ([#6261](https://github.com/hashicorp/terraform-provider-google-beta/pull/6261))
* compute: made `google_compute_node_group` require one of `initial_size` or `autoscaling_policy` fields configured upon resource creation ([#6384](https://github.com/hashicorp/terraform-provider-google-beta/pull/6384))
* compute: made `size` in `google_compute_node_group` an output only field. ([#6384](https://github.com/hashicorp/terraform-provider-google-beta/pull/6384))
* compute: removed default value for `rule.rate_limit_options.encorce_on_key` on resource `google_compute_security_policy` ([#6174](https://github.com/hashicorp/terraform-provider-google-beta/pull/6174))
* compute: retyped `consumer_accept_lists` to a SET from an ARRAY type for `google_compute_service_attachment` ([#6369](https://github.com/hashicorp/terraform-provider-google-beta/pull/6369))
* container: added `deletion_protection` to `google_container_cluster` which is enabled to `true` by default. When enabled, this field prevents Terraform from deleting the resource. ([#6391](https://github.com/hashicorp/terraform-provider-google-beta/pull/6391))
* container: changed `management.auto_repair` and `management.auto_upgrade` defaults to true in `google_container_node_pool` ([#6329](https://github.com/hashicorp/terraform-provider-google-beta/pull/6329))
* container: changed `networking_mode` default to `VPC_NATIVE` for newly created `google_container_cluster` resources ([#6402](https://github.com/hashicorp/terraform-provider-google-beta/pull/6402))
* container: removed `enable_binary_authorization` in `google_container_cluster` ([#6285](https://github.com/hashicorp/terraform-provider-google-beta/pull/6285))
* container: removed default for `logging_variant` in `google_container_node_pool` ([#6329](https://github.com/hashicorp/terraform-provider-google-beta/pull/6329))
* container: removed default value in `network_policy.provider` in `google_container_cluster` ([#6323](https://github.com/hashicorp/terraform-provider-google-beta/pull/6323))
* container: removed the behaviour that `google_container_cluster` will delete the cluster if it's created in an error state. Instead, it will mark the cluster as tainted, allowing manual inspection and intervention. To proceed with deletion, run another `terraform apply`. ([#6301](https://github.com/hashicorp/terraform-provider-google-beta/pull/6301))
* container: reworked the `taint` field in `google_container_cluster` and `google_container_node_pool` to only manage a subset of taint keys based on those already in state. Most existing resources are unaffected, unless they use `sandbox_config`- see upgrade guide for details. ([#6351](https://github.com/hashicorp/terraform-provider-google-beta/pull/6351))
* dataplex: removed `data_profile_result` and `data_quality_result` from `google_dataplex_scan` ([#6070](https://github.com/hashicorp/terraform-provider-google-beta/pull/6070))
* firebase: changed `deletion_policy` default to `DELETE` for `google_firebase_web_app`. ([#6018](https://github.com/hashicorp/terraform-provider-google-beta/pull/6018))
* firebase: removed `google_firebase_project_location` ([#6223](https://github.com/hashicorp/terraform-provider-google-beta/pull/6223))
* gameservices: removed Terraform support for `gameservices` ([#6112](https://github.com/hashicorp/terraform-provider-google-beta/pull/6112))
* logging: changed the default value of `unique_writer_identity` from `false` to `true` in `google_logging_project_sink`. ([#6210](https://github.com/hashicorp/terraform-provider-google-beta/pull/6210))
* logging: made `growth_factor`, `num_finite_buckets`, and `scale` required for `google_logging_metric` ([#6173](https://github.com/hashicorp/terraform-provider-google-beta/pull/6173))
* looker: removed `LOOKER_MODELER` as a possible value in `google_looker_instance.platform_edition` ([#6349](https://github.com/hashicorp/terraform-provider-google-beta/pull/6349))
* monitoring: fixed perma-diffs in `google_monitoring_dashboard.dashboard_json` by suppressing values returned by the API that are not in configuration ([#6392](https://github.com/hashicorp/terraform-provider-google-beta/pull/6392))
* monitoring: made `labels` immutable in `google_monitoring_metric_descriptor` ([#6372](https://github.com/hashicorp/terraform-provider-google-beta/pull/6372))
* privateca: removed deprecated fields `config_values`, `pem_certificates` from `google_privateca_certificate` ([#6097](https://github.com/hashicorp/terraform-provider-google-beta/pull/6097))
* secretmanager: removed `automatic` field in `google_secret_manager_secret` resource ([#6279](https://github.com/hashicorp/terraform-provider-google-beta/pull/6279))
* servicenetworking: used Create instead of Patch to create `google_service_networking_connection` ([#6222](https://github.com/hashicorp/terraform-provider-google-beta/pull/6222))
* servicenetworking: used the `deleteConnection` method to delete the resource `google_service_networking_connection` ([#6332](https://github.com/hashicorp/terraform-provider-google-beta/pull/6332))

FEATURES:
* **New Resource:** `google_scc_folder_custom_module` ([#6367](https://github.com/hashicorp/terraform-provider-google-beta/pull/6367))
* **New Resource:** `google_scc_organization_custom_module` ([#6390](https://github.com/hashicorp/terraform-provider-google-beta/pull/6390))

IMPROVEMENTS:
* alloydb: added additional fields to `google_alloydb_instance` and `google_alloydb_backup` ([#6363](https://github.com/hashicorp/terraform-provider-google-beta/pull/6363))
* artifactregistry: added support for remote APT and YUM repositories to `google_artifact_registry_repository` ([#6362](https://github.com/hashicorp/terraform-provider-google-beta/pull/6362))
* baremetal: made delete a noop for the resource `google_bare_metal_admin_cluster` to better align with actual behavior ([#6388](https://github.com/hashicorp/terraform-provider-google-beta/pull/6388))
* bigtable: added `state` output attribute to `google_bigtable_instance` clusters ([#6353](https://github.com/hashicorp/terraform-provider-google-beta/pull/6353))
* compute: made `google_compute_node_group` mutable ([#6384](https://github.com/hashicorp/terraform-provider-google-beta/pull/6384))
* compute: added `network_interface.security_policy` field to `google_compute_instance` resource ([#6343](https://github.com/hashicorp/terraform-provider-google-beta/pull/6343))
* compute: added `type` field to `google_compute_router_nat` resource ([#6331](https://github.com/hashicorp/terraform-provider-google-beta/pull/6371)) 
* compute: added `rules.action.source_nat_active_ranges` and `rules.action.source_nat_drain_ranges` field to `google_compute_router_nat` resource ([#6331](https://github.com/hashicorp/terraform-provider-google-beta/pull/6371)) 
* compute: added `network_attachment` to `google_compute_instance` ([#6331](https://github.com/hashicorp/terraform-provider-google-beta/pull/6331))
* container: added the `effective_taints` attribute to `google_container_cluster` and `google_container_node_pool`, outputting all known taint values ([#6351](https://github.com/hashicorp/terraform-provider-google-beta/pull/6351))
* container: allowed setting `addons_config.gcs_fuse_csi_driver_config` on `google_container_cluster` with `enable_autopilot: true`. ([#6378](https://github.com/hashicorp/terraform-provider-google-beta/pull/6378))
* containeraws: added `binary_authorization` to `google_container_aws_cluster` ([#6373](https://github.com/hashicorp/terraform-provider-google-beta/pull/6373))
* containeraws: added `update_settings` to `google_container_aws_node_pool` ([#6373](https://github.com/hashicorp/terraform-provider-google-beta/pull/6373))
* osconfig: added `week_day_of_month.day_offset` field to the `google_os_config_patch_deployment` resource ([#6379](https://github.com/hashicorp/terraform-provider-google-beta/pull/6379))
* secretmanager: allowed update for `rotation.rotation_period` field in `google_secret_manager_secret` resource ([#6345](https://github.com/hashicorp/terraform-provider-google-beta/pull/6345))
* sql: added `preferred_zone` field to `google_sql_database_instance` resource ([#6360](https://github.com/hashicorp/terraform-provider-google-beta/pull/6360))
* storagetransfer: added `event_stream` field to `google_storage_transfer_job` resource ([#6382](https://github.com/hashicorp/terraform-provider-google-beta/pull/6382))
* workstations: added `replica_zones`, `service_account_scopes`, and `enable_audit_agent` to `google_workstations_workstation_config` (beta) ([#6355](https://github.com/hashicorp/terraform-provider-google-beta/pull/6355))

BUG FIXES:
* bigquery: fixed diff suppression in `external_data_configuration.connection_id` in `google_bigquery_table` ([#6368](https://github.com/hashicorp/terraform-provider-google-beta/pull/6368))
* bigquery: fixed view and materialized view creation when schema is specified in `google_bigquery_table` ([#6034](https://github.com/hashicorp/terraform-provider-google-beta/pull/6034))
* bigtable: avoided re-creation of `google_bigtable_instance` when cluster is still updating and storage type changed ([#6353](https://github.com/hashicorp/terraform-provider-google-beta/pull/6353))
* bigtable: fixed a bug where dynamically created clusters would incorrectly run into duplication error in `google_bigtable_instance` ([#6338](https://github.com/hashicorp/terraform-provider-google-beta/pull/6338))
* compute: added default value to `metric.filter` in the resource `google_compute_autoscaler` (beta) ([#6082](https://github.com/hashicorp/terraform-provider-google-beta/pull/6082))
* compute: removed the default value for field `reconcile_connections ` in resource `google_compute_service_attachment`, the field will now default to a value returned by the API when not set in configuration ([#6322](https://github.com/hashicorp/terraform-provider-google-beta/pull/6322))
* compute: replaced incorrect default value for `enable_endpoint_independent_mapping` with APIs default in resource `google_compute_router_nat` ([#6053](https://github.com/hashicorp/terraform-provider-google-beta/pull/6053))
* container: fixed an issue in `google_container_node_pool` where empty `linux_node_config.sysctls` would crash the provider ([#6339](https://github.com/hashicorp/terraform-provider-google-beta/pull/6339))
* dataflow: fixed issue causing error message when max_workers and num_workers were supplied via parameters in `google_dataflow_flex_template_job` ([#6357](https://github.com/hashicorp/terraform-provider-google-beta/pull/6357))
* dataflow: fixed max_workers read value permanently displaying as 0 in `google_dataflow_flex_template_job` ([#6357](https://github.com/hashicorp/terraform-provider-google-beta/pull/6357))
* dataflow: fixed permadiff when SdkPipeline values are supplied via parameters in `google_dataflow_flex_template_job` ([#6357](https://github.com/hashicorp/terraform-provider-google-beta/pull/6357))
* firebase: made `google_firebase_rules.release` immutable ([#6373](https://github.com/hashicorp/terraform-provider-google-beta/pull/6373))
* identityplayform: fixed a potential perma-diff for `sign_in` in `google_identity_platform_config` resource ([#6317](https://github.com/hashicorp/terraform-provider-google-beta/pull/6317))
* monitoring: fixed an issue where `metadata` was not able to be updated in `google_monitoring_metric_descriptor` ([#6372](https://github.com/hashicorp/terraform-provider-google-beta/pull/6372))
* monitoring: fixed bug where importing `google_monitoring_notification_channel` failed when no default project was supplied in provider configuration or through environment variables ([#6327](https://github.com/hashicorp/terraform-provider-google-beta/pull/6327))
* secretmanager: fixed an issue in `google_secretmanager_secret` where replacing `replication.automatic` with `replication.auto` would destroy and recreate the resource ([#6325](https://github.com/hashicorp/terraform-provider-google-beta/pull/6325))
* sql: fixed diffs when re-ordering existing `database_flags` in `google_sql_database_instance` ([#6172](https://github.com/hashicorp/terraform-provider-google-beta/pull/6172))
* tags: fixed import failure on `google_tags_tag_binding` ([#6383](https://github.com/hashicorp/terraform-provider-google-beta/pull/6383))
* vertexai: made `contents_delta_uri` a required field in `google_vertex_ai_index` as omitting it would result in an error ([#6374](https://github.com/hashicorp/terraform-provider-google-beta/pull/6374))
* workstations: fixed in-place updates of `host.gce_instance.accelerators` in `google_workstation_config` ([#6354](https://github.com/hashicorp/terraform-provider-google-beta/pull/6354))

## 4.85.0 (June 12, 2024)

NOTES:
* The `4.85.0` release backports configuration for the retention period for Cloud Storage soft delete (https://cloud.google.com/resources/storage/soft-delete-announce) so that customers who have not yet upgraded to `5.22.0`+ are able to configure the retention period of objects in their buckets. By upgrading to this version and configuring or otherwise interacting with the `google_storage_bucket.soft_delete_policy` values, you will need to upgrade directly to `5.22.0`+ from `4.85.0` when upgrading to `5.X` in the future.

IMPROVEMENTS:
* storage: added `soft_delete_policy` to `google_storage_bucket` resource ([#7119](https://github.com/hashicorp/terraform-provider-google-beta/pull/7119))

## 4.84.0 (September 26, 2023)

DEPRECATIONS:
* alloydb: deprecated `network` field in favor of `network_config` on `google_alloydb_cluster`. ([#6297](https://github.com/hashicorp/terraform-provider-google-beta/pull/6297))
* identityplayform: deprecated `google_identity_platform_project_default_config` resource. Use `google_identity_platform_config` resource instead ([#6293](https://github.com/hashicorp/terraform-provider-google-beta/pull/6293))

FEATURES:
* **New Data Source:** `google_certificate_manager_certificate_map` ([#6316](https://github.com/hashicorp/terraform-provider-google-beta/pull/6316))
* **New Resource:** `google_artifact_registry_vpcsc_config` ([#6265](https://github.com/hashicorp/terraform-provider-google-beta/pull/6265))
* **New Resource:** `google_dialogflow_cx_security_settings` ([#6300](https://github.com/hashicorp/terraform-provider-google-beta/pull/6300))
* **New Resource:** `google_gke_backup_restore_plan` ([#6278](https://github.com/hashicorp/terraform-provider-google-beta/pull/6278))
* **New Resource:** `google_scc_project_custom_module` ([#6315](https://github.com/hashicorp/terraform-provider-google-beta/pull/6315))
* **New Resource:** `google_tpu_v2_vm` ([#6264](https://github.com/hashicorp/terraform-provider-google-beta/pull/6264))
* **New Resource:** `google_edgenetwork_network` ([#6305](https://github.com/hashicorp/terraform-provider-google-beta/pull/6305))
* **New Resource:** `google_edgenetwork_subnet` ([#6305](https://github.com/hashicorp/terraform-provider-google-beta/pull/6305))

IMPROVEMENTS:
* alloydb: added `network_config` field to support named IP ranges on `google_alloydb_cluster`. ([#6297](https://github.com/hashicorp/terraform-provider-google-beta/pull/6297))
* cloudrunv2: added fields `network_interfaces` to resource `google_cloud_run_v2_job` to support Direct VPC egress. ([#6287](https://github.com/hashicorp/terraform-provider-google-beta/pull/6287))
* cloudrunv2: added fields `network_interfaces` to resource `google_cloud_run_v2_service` to support Direct VPC egress. ([#6287](https://github.com/hashicorp/terraform-provider-google-beta/pull/6287))
* compute: updated the `autoscaling_policy.mode` to accept `ONLY_SCALE_OUT` on `google_compute_autoscaler` ([#6304](https://github.com/hashicorp/terraform-provider-google-beta/pull/6304))
* compute: added `server_tls_policy` argument to `google_compute_target_https_proxy` resource ([#6269](https://github.com/hashicorp/terraform-provider-google-beta/pull/6269))
* compute: added `member` attribute to `google_compute_default_service_account` datasource ([#6311](https://github.com/hashicorp/terraform-provider-google-beta/pull/6311))
* compute: added output field `internal_ipv6_prefix` to `google_compute_subnetwork` resource ([#6306](https://github.com/hashicorp/terraform-provider-google-beta/pull/6306))
* container: added `node_config.fast_socket` field to `google_container_node_pool` ([#6289](https://github.com/hashicorp/terraform-provider-google-beta/pull/6289))
* containeraws: added support for `auto_repair` in `google_container_aws_node_pool` ([#6282](https://github.com/hashicorp/terraform-provider-google-beta/pull/6282))
* containerazure: added support for `auto_repair` in `google_container_azure_node_pool` ([#6282](https://github.com/hashicorp/terraform-provider-google-beta/pull/6282))
* filestore: added support for the `"ZONAL"` value to `tier` in `google_filestore_instance` ([#6303](https://github.com/hashicorp/terraform-provider-google-beta/pull/6303))
* firestore: added `delete_protection_state` field to `google_firestore_database` resource. ([#6295](https://github.com/hashicorp/terraform-provider-google-beta/pull/6295))
* identityplatform: added `sign-in` field to `google_identity_platform_config` resource ([#6293](https://github.com/hashicorp/terraform-provider-google-beta/pull/6293))
* networkconnectivity: added support for `linked_vpc_network` in `google_network_connectivity_spoke` ([#6282](https://github.com/hashicorp/terraform-provider-google-beta/pull/6282))
* networkservices: increased default timeout for `google_network_services_edge_cache_origin` to 120m from 60m ([#6275](https://github.com/hashicorp/terraform-provider-google-beta/pull/6275))
* networkservices: increased default timeout for `google_network_services_edge_cache_service` to 60m from 30m ([#6281](https://github.com/hashicorp/terraform-provider-google-beta/pull/6281))
* secretmanager: added `is_secret_data_base64` field to `google_secret_manager_secret_version` resource ([#6273](https://github.com/hashicorp/terraform-provider-google-beta/pull/6273))
* workstations: added `env` field to `google_workstations_workstation` resource ([#6258](https://github.com/hashicorp/terraform-provider-google-beta/pull/6258))

BUG FIXES:
* bigquery: updated documentation for `google_bigquery_table.time_partitioning.expiration_ms` ([#6290](https://github.com/hashicorp/terraform-provider-google-beta/pull/6290))
* bigtable: added a read timeout to `google_bigtable_instance` ([#6276](https://github.com/hashicorp/terraform-provider-google-beta/pull/6276))
* bigtable: improved regional reliability when instance overlaps a downed region in the resource `google_bigtable_instance` ([#6313](https://github.com/hashicorp/terraform-provider-google-beta/pull/6313))
* eventarc: resolved permadiff on `google_eventarc_trigger.event_data_content_type` by defaulting to the value returned by the API if not set in the configuration. ([#6282](https://github.com/hashicorp/terraform-provider-google-beta/pull/6282))
* identityplatform: fixed a potential perma-diff for `sign_in` in `google_identity_platform_config` resource ([#6317](https://github.com/hashicorp/terraform-provider-google-beta/pull/6317))
* monitoring: fixed scaling issues when deploying terraform changes with many `google_monitoring_monitored_project` ([#6259](https://github.com/hashicorp/terraform-provider-google-beta/pull/6259))
* monitoring: fixed validation of `service_id` on `google_monitoring_custom_service` and `slo_id` on `google_monitoring_slo` ([#6266](https://github.com/hashicorp/terraform-provider-google-beta/pull/6266))
* osconfig: fixed no more than one setting is allowed under `patch_config.windows_update` on `google_os_config_patch_deployment` ([#6314](https://github.com/hashicorp/terraform-provider-google-beta/pull/6314))
* provider: addressed a bug where configuring the provider with unknown values did not behave as expected ([#6312](https://github.com/hashicorp/terraform-provider-google-beta/pull/6312))
* provider: fixed the provider so it resumes ignoring empty strings set in the `provider` block ([#6268](https://github.com/hashicorp/terraform-provider-google-beta/pull/6268))
* secretmanager: replaced the panic block with an error in import function of `google_secret_manager_secret_version` resource ([#6296](https://github.com/hashicorp/terraform-provider-google-beta/pull/6296))
* secretmanager: fixed an issue in `google_secretmanager_secret` where replacing `replication.automatic` with `replication.auto` would destroy and recreate the resource ([#6325](https://github.com/hashicorp/terraform-provider-google-beta/pull/6325))

## 4.83.0 (September 18, 2023)

DEPRECATIONS:
* secretmanager: deprecated `automatic` field on `google_secret_manager_secret`. Use `auto` instead. ([#6237](https://github.com/hashicorp/terraform-provider-google-beta/pull/6237))

FEATURES:
* **New Resource:** `google_biglake_table` ([#6205](https://github.com/hashicorp/terraform-provider-google-beta/pull/6205))
* **New Resource:** `google_data_pipeline_pipeline` ([#6236](https://github.com/hashicorp/terraform-provider-google-beta/pull/6236))
* **New Resource:** `google_dialogflow_cx_test_case` ([#6249](https://github.com/hashicorp/terraform-provider-google-beta/pull/6249))
* **New Resource:** `google_storage_insights_report_config` ([#6253](https://github.com/hashicorp/terraform-provider-google-beta/pull/6253))
* **New Resource:** `google_apigee_target_server` ([#6215](https://github.com/hashicorp/terraform-provider-google-beta/pull/6215))

IMPROVEMENTS:
* bigquery: added `allow_non_incremental_definition` to `google_bigquery_table` resource ([#6248](https://github.com/hashicorp/terraform-provider-google-beta/pull/6248))
* bigquery: added `table_constraints` field to `google_bigquery_table` resource ([#6250](https://github.com/hashicorp/terraform-provider-google-beta/pull/6250))
* compute: added internal IPV6 support for `google_compute_address` and `google_compute_instance` resources ([#6232](https://github.com/hashicorp/terraform-provider-google-beta/pull/6232))
* containerattached: added `binary_authorization` field to `google_container_attached_cluster` resource ([#6256](https://github.com/hashicorp/terraform-provider-google-beta/pull/6256))
* containeraws: added update support for `config.instance_type` in `container_aws_node_pool` ([#6282](https://github.com/hashicorp/terraform-provider-google-beta/pull/6282))
* firestore: added `point_in_time_recovery_enablement` field to `google_firestore_database` resource ([#6239](https://github.com/hashicorp/terraform-provider-google-beta/pull/6239))
* firestore: added `update_time` and `uid` fields to `google_firestore_database` resource ([#6257](https://github.com/hashicorp/terraform-provider-google-beta/pull/6257))
* gkehub2: added `labels`, `namespace_labels` fields to `google_gke_hub_namespace` resource ([#6202](https://github.com/hashicorp/terraform-provider-google-beta/pull/6202))
* gkehub: added `labels` fields to `google_gke_hub_membership_binding` resource ([#6216](https://github.com/hashicorp/terraform-provider-google-beta/pull/6216))
* gkehub: added `labels` fields to `google_gke_hub_scope` resource ([#6243](https://github.com/hashicorp/terraform-provider-google-beta/pull/6243))
* gkeonprem: added `upgrade_policy` and `binary_authorization` fields in `google_gkeonprem_bare_metal_cluster` resource (beta) ([#6224](https://github.com/hashicorp/terraform-provider-google-beta/pull/6224))
* gkeonprem: added `upgrade_policy` field in `google_gkeonprem_vmware_cluster` resource (beta) ([#6224](https://github.com/hashicorp/terraform-provider-google-beta/pull/6224))
* secretmanager: added `auto` field to `google_secret_manager_secret` resource ([#6237](https://github.com/hashicorp/terraform-provider-google-beta/pull/6237))
* secretmanager: added `deletion_policy` field to `google_secret_manager_secret_version` resource ([#6252](https://github.com/hashicorp/terraform-provider-google-beta/pull/6252))
* storage: supported in-place update for `autoclass` field in `google_storage_bucket` resource ([#6233](https://github.com/hashicorp/terraform-provider-google-beta/pull/6233))
* vertexai: added `public_endpoint_enabled` to `google_vertex_ai_index_endpoint` ([#6208](https://github.com/hashicorp/terraform-provider-google-beta/pull/6208))
* workstations: added `env` field to `google_workstations_workstation` resource (beta) ([#6258](https://github.com/hashicorp/terraform-provider-google-beta/pull/6258))

BUG FIXES:
* bigquerydatatransfer: fixed a bug when importing `location` of `google_bigquery_data_transfer_config` ([#6203](https://github.com/hashicorp/terraform-provider-google-beta/pull/6203))
* container: fixed a bug where `additional_pod_network_configs` was not sent correctly in `google_container_node_pool` ([#6211](https://github.com/hashicorp/terraform-provider-google-beta/pull/6211))
* container: fixed concurrent ops' quota-error to be retriable in `google_container_node_pool ` ([#6254](https://github.com/hashicorp/terraform-provider-google-beta/pull/6254))
* eventarc: resolved permadiff on `event_content_type` in `eventarc_trigger`, the field will now default to a value returned by the API when not set in configuration ([#6282](https://github.com/hashicorp/terraform-provider-google-beta/pull/6282))
* pipeline: fixed issue where certain `google_dataflow_job` instances would crash the provider ([#6255](https://github.com/hashicorp/terraform-provider-google-beta/pull/6255))
* provider: fixed a bug where `user_project_override` would not be not used correctly when provisioning resources implemented using the plugin framework. Currently there are no resources implemented this way, so no-one should have been impacted. ([#6230](https://github.com/hashicorp/terraform-provider-google-beta/pull/6230))
* pubsub: fixed issue where setting `no_wrapper.write_metadata` to false wasn't passed to the API for `google_pubsub_subscription` ([#6219](https://github.com/hashicorp/terraform-provider-google-beta/pull/6219))
* serviceaccount: added retries for reads after `google_service_account` creation if 403 Forbidden is returned. ([#6221](https://github.com/hashicorp/terraform-provider-google-beta/pull/6221))
* storage: fixed the failure in building a plan when a `content` value is expected on `google_storage_bucket_object_content` ([#6204](https://github.com/hashicorp/terraform-provider-google-beta/pull/6204)

## 4.82.0 (September 11, 2023)

IMPROVEMENTS:
* compute: added in-place update support for field `enable_proxy_protocol` in `google_compute_service_attachment` resource ([#6192](https://github.com/hashicorp/terraform-provider-google-beta/pull/6192))
* compute: added in-place update support for field `reconcile_connections` in `google_compute_service_attachment` resource ([#6187](https://github.com/hashicorp/terraform-provider-google-beta/pull/6187))
* compute: added in-place update support for field `allowPscGlobalAccess` in `google_compute_forwarding_rule` resource ([#6179](https://github.com/hashicorp/terraform-provider-google-beta/pull/6179))
* container: added additional options for field `monitoring_config.enable_components` in `google_container_cluster` resource ([#6198](https://github.com/hashicorp/terraform-provider-google-beta/pull/6198))
* gkehub: added `labels` field to `google_gke_hub_scope_rbac_role_binding` resource ([#6200](https://github.com/hashicorp/terraform-provider-google-beta/pull/6200))
* logging: added in-place update support for field `unique_writer_identity` in `google_logging_project_sink` resource ([#6193](https://github.com/hashicorp/terraform-provider-google-beta/pull/6193))
* networkconnectivity: added `psc_connections.error.details` field to `google_network_connectivity_service_connection_policy` resource ([#6197](https://github.com/hashicorp/terraform-provider-google-beta/pull/6197))
* secretmanager: added in-place update support for field `replication.user_managed.replicas.customer_managed_encryption` in `google_secret_manager_secret` resource ([#6177](https://github.com/hashicorp/terraform-provider-google-beta/pull/6177))

BUG FIXES:
* bigquery: made `params.destination_table_name_template` and `params.data_path` immutable as updating these fields if value of `data_source_id` is `amazon_s3` in `google_bigquery_data_transfer_config` resource ([#6195](https://github.com/hashicorp/terraform-provider-google-beta/pull/6195))
* compute: fixed a crash when empty is given to `all_instances_config` in `google_compute_region_instance_group_manager` resource ([#6191](https://github.com/hashicorp/terraform-provider-google-beta/pull/6191))
* dns: fixed hash function for `network_url` in `google_dns_managed_zone` and `google_dns_policy` resources to make sure that the private DNS zone or DNS policy can be attatched to all of the networks in different projects, even though the network name is the same across of those projects ([#6199](https://github.com/hashicorp/terraform-provider-google-beta/pull/6199))
* servicedirectory: made `location` immutable as updating this field in `google_service_directory_namespace` resource ([#6182](https://github.com/hashicorp/terraform-provider-google-beta/pull/6182))

## 4.81.0 (September 05, 2023)
FEATURES:
* **New Resource:** `google_biglake_catalog` ([#6152](https://github.com/hashicorp/terraform-provider-google-beta/pull/6152))
* **New Resource:** `google_redis_cluster` ([#6158](https://github.com/hashicorp/terraform-provider-google-beta/pull/6158))
* **New Resource:** `google_biglake_database` ([#6161](https://github.com/hashicorp/terraform-provider-google-beta/pull/6161))
* **New Resource:** `google_compute_network_attachment` ([#6159](https://github.com/hashicorp/terraform-provider-google-beta/pull/6159))
* **New Resource:** `google_gke_hub_membership_binding` ([#6170](https://github.com/hashicorp/terraform-provider-google-beta/pull/6170))
* **New Resource:** `google_gke_hub_namespace` ([#6170](https://github.com/hashicorp/terraform-provider-google-beta/pull/6170))
* **New Resource:** `google_gke_hub_scope` ([#6170](https://github.com/hashicorp/terraform-provider-google-beta/pull/6170))
* **New Resource:** `google_gke_hub_scope_iam_member` ([#6170](https://github.com/hashicorp/terraform-provider-google-beta/pull/6170))
* **New Resource:** `google_gke_hub_scope_iam_policy` ([#6170](https://github.com/hashicorp/terraform-provider-google-beta/pull/6170))
* **New Resource:** `google_gke_hub_membership_binding` ([#6170](https://github.com/hashicorp/terraform-provider-google-beta/pull/6170))
* **New Resource:** `google_gke_hub_scope_rbac_role_binding` ([#6170](https://github.com/hashicorp/terraform-provider-google-beta/pull/6170))

IMPROVEMENTS:
* compute: made the field `distribution_policy_target_shape` of `google_compute_region_instance_group_manager` not cause recreation of the resource. ([#6156](https://github.com/hashicorp/terraform-provider-google-beta/pull/6156))
* container: added `enable_fqdn_network_policy` field to `google_container_cluster` ([#6157](https://github.com/hashicorp/terraform-provider-google-beta/pull/6157))
* container: added `node_config.confidential_compute` field to `google_container_node_pool` resource ([#6166](https://github.com/hashicorp/terraform-provider-google-beta/pull/6166))
* datastream: allowed `password` of `google_datastream_connection_profile` to be mutable. ([#6140](https://github.com/hashicorp/terraform-provider-google-beta/pull/6140))
* dialogflowcx: added `response_type`, `channel`, `payload`, `conversation_success`, `output_audio_text`, `live_agent_handoff`, `play_audo`, `telephony_transfer_call`, `reprompt_event_handlers`, `set_parameter_actions`, and `conditional_cases` fields to `google_dialogflow_cx_page` resource ([#6168](https://github.com/hashicorp/terraform-provider-google-beta/pull/6168))
* dialogflowcx: added `response_type`, `channel`, `payload`, `conversation_success`, `output_audio_text`, `live_agent_handoff`, `play_audo`, `telephony_transfer_call`, `set_parameter_actions`, and `conditional_cases` fields to `google_dialogflow_cx_flow` resource ([#6168](https://github.com/hashicorp/terraform-provider-google-beta/pull/6168))
* iam: added `web_sso_config.additional_scopes` field to `google_iam_workforce_pool_provider` resource. ([#6145](https://github.com/hashicorp/terraform-provider-google-beta/pull/6145))
* iamworkforcepool: added `jwksJson` field to `WorkforcePoolProvider` resource ([#6153](https://github.com/hashicorp/terraform-provider-google-beta/pull/6153))
* monitoring: added `synthetic_monitor` to `google_monitoring_uptime_check_config` resource ([#6148](https://github.com/hashicorp/terraform-provider-google-beta/pull/6148))
* provider: improved error message when resource creation fails to to invalid API response ([#6149](https://github.com/hashicorp/terraform-provider-google-beta/pull/6149))

BUG FIXES:
* cloudrunv2: changed `template.volumes.secret.items.mode` field in `google_cloud_run_v2_job` resource to a non-required field. ([#6154](https://github.com/hashicorp/terraform-provider-google-beta/pull/6154))
* cloudrunv2: changed `template.volumes.secret.items.mode` field in `google_cloud_run_v2_service` resource to a non-required field. ([#6154](https://github.com/hashicorp/terraform-provider-google-beta/pull/6154))
* filestore: fixed a bug causing permadiff on `reserved_ip_range` field in `google_filestore_instance` ([#6143](https://github.com/hashicorp/terraform-provider-google-beta/pull/6143))
* identityplatform: fixed a permadiff on `authorized_domains` in `google_identity_platform_config` resource ([#6137](https://github.com/hashicorp/terraform-provider-google-beta/pull/6137))


## 4.80.0 (August 28, 2023)

DEPRECATIONS:
* dataplex: deprecated the following `google_dataplex_datascan` fields: `dataProfileResult` and `dataQualityResult` ([#6090](https://github.com/hashicorp/terraform-provider-google-beta/pull/6090))
* firebase: deprecated `google_firebase_project_location` in favor of `google_firebase_storage_bucket` and `google_firestore_database` ([#6087](https://github.com/hashicorp/terraform-provider-google-beta/pull/6087))

FEATURES:
* **New Data Source:** `google_sql_database_instance_latest_recovery_time` ([#6109](https://github.com/hashicorp/terraform-provider-google-beta/pull/6109))
* **New Resource:** `google_certificate_manager_trust_config` ([#6118](https://github.com/hashicorp/terraform-provider-google-beta/pull/6118))
* **New Resource:** `google_compute_region_security_policy_rule` ([#6086](https://github.com/hashicorp/terraform-provider-google-beta/pull/6086))
* **New Resource:** `google_gke_hub_membership_rbac_role_binding` ([#6103](https://github.com/hashicorp/terraform-provider-google-beta/pull/6103))
* **New Resource:** `google_iam_deny_policy` (ga only) ([#6125](https://github.com/hashicorp/terraform-provider-google-beta/pull/6125))
* **New Resource:** dataform_repository_workflow_config (beta) ([#6102](https://github.com/hashicorp/terraform-provider-google-beta/pull/6102))
* **New Resource:** google_bigquery_bi_reservation ([#6088](https://github.com/hashicorp/terraform-provider-google-beta/pull/6088))

IMPROVEMENTS:
* alloydb: added `restore_backup_source` and `restore_continuous_backup_source` fields to support restore feature in `google_alloydb_cluster` resource. ([#6129](https://github.com/hashicorp/terraform-provider-google-beta/pull/6129))
* artifactregistry: added `cleanup_policies` and `cleanup_policy_dry_run` fields to resource `google_artifact_registry_repository` ([#6117](https://github.com/hashicorp/terraform-provider-google-beta/pull/6117))
* compute: added `security_policy` field to `google_compute_target_instance` resource ([#6122](https://github.com/hashicorp/terraform-provider-google-beta/pull/6122))
* compute: added support for `security_policy` field to `google_compute_target_pool` ([#6124](https://github.com/hashicorp/terraform-provider-google-beta/pull/6124))
* compute: added support for `user_defined_fields` to `google_compute_region_security_policy` ([#6086](https://github.com/hashicorp/terraform-provider-google-beta/pull/6086))
* compute: added support for specifying regional disks for `google_compute_instance` `boot_disk.source` ([#6132](https://github.com/hashicorp/terraform-provider-google-beta/pull/6132))
* container: added `additional_pod_ranges_config` field to `google_container_cluster` resource ([#6133](https://github.com/hashicorp/terraform-provider-google-beta/pull/6133))
* dataplex: added fields `data_profile_spec.post_scan_actions`, `data_profile_spec.include_fields` and `data_profile_spec.exclude_fields` ([#6104](https://github.com/hashicorp/terraform-provider-google-beta/pull/6104))
* dns: added support for removing the networks block from the configuration in resource `google_dns_response_policy` ([#6111](https://github.com/hashicorp/terraform-provider-google-beta/pull/6111))
* firebase: added `api_key_id` field to `google_firebase_web_app`, `google_firebase_android_app`, and `google_firebase_apple_app`. ([#6127](https://github.com/hashicorp/terraform-provider-google-beta/pull/6127))
* gkeonprem: automatically set `ignore_errors` to true in `google_gkeonprem_bare_metal_admin_cluster` delete calls ([#6095](https://github.com/hashicorp/terraform-provider-google-beta/pull/6095))
* sql: added `psc_config` , `psc_service_attachment_link`, and `dns_name` fields to `google_sql_database_instance` ([#6119](https://github.com/hashicorp/terraform-provider-google-beta/pull/6119))
* workstations: added `enable_nested_virtualization` field to `google_workstations_workstation_config` resource ([#6123](https://github.com/hashicorp/terraform-provider-google-beta/pull/6123))

BUG FIXES:
* bigquery: added support to unset policy tags in table schema ([#6106](https://github.com/hashicorp/terraform-provider-google-beta/pull/6106))
* bigtable: fixed permadiff in `google_bigtable_gc_policy.gc_rules` when `max_age` is specified using increments larger than hours ([#6131](https://github.com/hashicorp/terraform-provider-google-beta/pull/6131))
* bigtable: fixed permadiff in `google_bigtable_gc_policy.gc_rules` when `mode` is specified ([#6131](https://github.com/hashicorp/terraform-provider-google-beta/pull/6131))
* container: updated `resource_container_cluster` to ignore `dns_config` diff when `enable_autopilot = true` ([#6108](https://github.com/hashicorp/terraform-provider-google-beta/pull/6108))
* containeraws: allowed `config.labels` to be updatable in `google_container_aws_node_pool` ([#6120](https://github.com/hashicorp/terraform-provider-google-beta/pull/6120))
* containerazure: added diff suppression for case changes of enum values in `google_container_azure_cluster` ([#6096](https://github.com/hashicorp/terraform-provider-google-beta/pull/6096))


## 4.79.0 (August 21, 2023)
FEATURES:
* **New Resource:** `google_backup_dr_management_server` ([#6054](https://github.com/hashicorp/terraform-provider-google-beta/pull/6054))
* **New Resource:** `google_compute_region_security_policy_rule` ([#6086](https://github.com/hashicorp/terraform-provider-google-beta/pull/6086))

IMPROVEMENTS:
* cloudbuild: added `git_file_source.bitbucket_server_config` and `source_to_build.bitbucket_server_config` fields to `google_cloudbuild_trigger` resource ([#6051](https://github.com/hashicorp/terraform-provider-google-beta/pull/6051))
* cloudrunv2: added the following output only fields to `google_cloud_run_v2_job` and `google_cloud_run_v2_service` resources: `create_time`, `update_time`, `delete_time`, `expire_time`, `creator` and `last_modifier` ([#6067](https://github.com/hashicorp/terraform-provider-google-beta/pull/6067))
* composer: added `config.private_environment_config.connection_type` field to `google_composer_environment` resource ([#6043](https://github.com/hashicorp/terraform-provider-google-beta/pull/6043))
* compute: added `disk.provisioned_iops` field to `google_compute_instance_template` and `google_compute_region_instance_template` resources ([#6071](https://github.com/hashicorp/terraform-provider-google-beta/pull/6071))
* compute: added `advanced_options_config.user_ip_request_headers` field to `google_compute_security_policy` resource ([#6048](https://github.com/hashicorp/terraform-provider-google-beta/pull/6048))
* compute: added `user_defined_fields` field to `google_compute_region_security_policy` resource ([#6086](https://github.com/hashicorp/terraform-provider-google-beta/pull/6086))
* databasemigrationservice: added `edition` field to `google_database_migration_service_connection_profile` resource ([#6074](https://github.com/hashicorp/terraform-provider-google-beta/pull/6074))
* dns: allowed `globalL7ilb` value for the `routing_policy.load_balancer_type` field in `google_dns_record_set` resource ([#6084](https://github.com/hashicorp/terraform-provider-google-beta/pull/6084))
* gkeonprem: added `control_plane_node.vsphere_config.storage_policy_name` and `vcenter.storage_policy_name` fields to `google_gkeonprem_vmware_cluster` resource ([#6072](https://github.com/hashicorp/terraform-provider-google-beta/pull/6072))
* healthcare: added `default_search_handling_strict` field to `google_healthcare_fhir_store` resource ([#6078](https://github.com/hashicorp/terraform-provider-google-beta/pull/6078))
* metastore: added `scaling_config` field to `google_dataproc_metastore_service` resource ([#6052](https://github.com/hashicorp/terraform-provider-google-beta/pull/6052))
* secretmanager: added `version_aliases` field to `google_secret_manager_secret` resource ([#6058](https://github.com/hashicorp/terraform-provider-google-beta/pull/6058))

BUG FIXES:
* alloydb: fixed a permadiff on `google_alloydb_cluster` when `backup_window`, `enabled` or `location` fields are unset ([#6036](https://github.com/hashicorp/terraform-provider-google-beta/pull/6036))
* containeraws: fixed permadiffs on `google_container_aws_cluster` and `google_container_aws_node_pool` resources ([#6060](https://github.com/hashicorp/terraform-provider-google-beta/pull/6060))
* dataplex: fixed a bug when importing `google_dataplex_datascan` after running a job ([#6047](https://github.com/hashicorp/terraform-provider-google-beta/pull/6047))
* dns: changed `private_visibility_config.networks` from `required` to requiring at least one of `private_visibility_config.networks` or `private_visibility_config.gke_clusters` in `google_dns_managed_zone` resource ([#6035](https://github.com/hashicorp/terraform-provider-google-beta/pull/6035))



## 4.78.0 (August 15, 2023)

FEATURES:
* **New Resource:** `google_billing_project_info` ([#6015](https://github.com/hashicorp/terraform-provider-google-beta/pull/6015))
* **New Resource:** `google_dataform_repository_release_config` ([#6009](https://github.com/hashicorp/terraform-provider-google-beta/pull/6009))
* **New Resource:** `google_network_connectivity_service_connection_policy` ([#6000](https://github.com/hashicorp/terraform-provider-google-beta/pull/6000))

IMPROVEMENTS:
* alloydb: added `continuous_backup_config` and `continuous_backup_info` fields to `cluster` resource ([#5996](https://github.com/hashicorp/terraform-provider-google-beta/pull/5996))
* bigquery: added `external_data_configuration.file_set_spec_type` to `google_bigquery_table` ([#6017](https://github.com/hashicorp/terraform-provider-google-beta/pull/6017))
* bigquery: added `max_staleness` to `google_bigquery_table` ([#6010](https://github.com/hashicorp/terraform-provider-google-beta/pull/6010))
* billingbudget: added `resource_ancestors` field to `google_billing_budget` resource ([#6008](https://github.com/hashicorp/terraform-provider-google-beta/pull/6008))
* cloudfunctions2: added support for GCF Gen2 CMEK ([#6004](https://github.com/hashicorp/terraform-provider-google-beta/pull/6004))
* cloudidentity: added field `type` to `google_cloud_identity_group_memberships` ([#6013](https://github.com/hashicorp/terraform-provider-google-beta/pull/6013))
* compute: added `subnetwork` field to the resource `google_compute_global_forwarding_rule` ([#6026](https://github.com/hashicorp/terraform-provider-google-beta/pull/6026))
* compute: added support for `INTERNAL_MANAGED` to the field `load_balancing_scheme` in the resource `google_compute_backend_service` ([#6026](https://github.com/hashicorp/terraform-provider-google-beta/pull/6026))
* compute: added support for `INTERNAL_MANAGED` to the field `load_balancing_scheme` in the resource `google_compute_global_forwarding_rule` ([#6026](https://github.com/hashicorp/terraform-provider-google-beta/pull/6026))
* compute: added support for `ip_version` to `google_compute_forwarding_rule` ([#6006](https://github.com/hashicorp/terraform-provider-google-beta/pull/6006))
* container: marked `master_ipv4_cidr_block` as not required when `private_endpoint_subnetwork` is provided for `google_container_cluster` ([#6025](https://github.com/hashicorp/terraform-provider-google-beta/pull/6025))
* container: added support for `advanced_datapath_observability_config` to `google_container_cluster` ([#6027](https://github.com/hashicorp/terraform-provider-google-beta/pull/6027))
* eventarc: added field `event_data_content_type` to `google_eventarc_trigger` ([#6032](https://github.com/hashicorp/terraform-provider-google-beta/pull/6032))
* healthcare: added `send_previous_resource_on_delete` field to `notification_configs` of `google_healthcare_fhir_store` ([#5999](https://github.com/hashicorp/terraform-provider-google-beta/pull/5999))
* pubsub: added `cloud_storage_config` field to `google_pubsub_subscription` resource ([#6024](https://github.com/hashicorp/terraform-provider-google-beta/pull/6024))
* secretmanager: added `annotations` field to `google_secret_manager_secret` resource ([#6007](https://github.com/hashicorp/terraform-provider-google-beta/pull/6007))
* workstations: added `private_cluster_config.allowed_projects` arguments to `google_workstations_workstation_cluster` ([#6021](https://github.com/hashicorp/terraform-provider-google-beta/pull/6021))

BUG FIXES:
* certificatemanager: added recreation behavior to the `google_certificate_manager_certificate` resource when its location changes ([#6031](https://github.com/hashicorp/terraform-provider-google-beta/pull/6031))
* cloudfunctions2: fixed creation failure state inconsistency in `google_cloudfunctions2_function` ([#6023](https://github.com/hashicorp/terraform-provider-google-beta/pull/6023))
* monitoring: updated `evaluation_interval` on `condition_prometheus_query_language` to be optional ([#6028](https://github.com/hashicorp/terraform-provider-google-beta/pull/6028))

## 4.77.0 (August 7, 2023)

NOTES:
* vpcaccess: reverted the ability to update the number of instances for resource `google_vpc_access_connector` ([#5957](https://github.com/hashicorp/terraform-provider-google-beta/pull/5957))

FEATURES:
* **New Resource:** `google_document_ai_warehouse_document_schema` ([#5965](https://github.com/hashicorp/terraform-provider-google-beta/pull/5965))
* **New Resource:** `google_document_ai_warehouse_location` ([#5965](https://github.com/hashicorp/terraform-provider-google-beta/pull/5965))

IMPROVEMENTS:
* alloydb: added `continuous_backup_config` and `continuous_backup_info` fields to `cluster` resource ([#5996](https://github.com/hashicorp/terraform-provider-google-beta/pull/5996))
* cloudbuild: removed the validation function for the values of `machine_type` field on the `google_cloudbuild_trigger` resource ([#5985](https://github.com/hashicorp/terraform-provider-google-beta/pull/5985))
* compute: added future_limit in quota exceeded error details for compute resources. ([#5982](https://github.com/hashicorp/terraform-provider-google-beta/pull/5982))
* compute: added `enable_strong_affinity` field to `google_compute_region_backend_service` (beta) ([#5962](https://github.com/hashicorp/terraform-provider-google-beta/pull/5962))
* compute: added `ipv6_endpoint_type` and `ip_version` to `google_compute_address` ([#5986](https://github.com/hashicorp/terraform-provider-google-beta/pull/5986))
* compute: added `network_interface.ipv6_access_config.external_ipv6_prefix_length` to `google_compute_instance` ([#5986](https://github.com/hashicorp/terraform-provider-google-beta/pull/5986))
* compute: added `network_interface.ipv6_access_config.name` to `google_compute_instance` ([#5986](https://github.com/hashicorp/terraform-provider-google-beta/pull/5986))
* compute: added a new type `GLOBAL_MANAGED_PROXY` for the field `purpose` in the resource `google_compute_subnetwork` ([#5981](https://github.com/hashicorp/terraform-provider-google-beta/pull/5981))
* compute: added protocol type: UNSPECIFIED in `google_compute_backend_service` as per [release note](https://cloud.google.com/load-balancing/docs/release-notes#July_24_2023) ([#5967](https://github.com/hashicorp/terraform-provider-google-beta/pull/5967))
* compute: added `local_ssd_recovery_timeout` field to `google_compute_instance` resource ([#5968](https://github.com/hashicorp/terraform-provider-google-beta/pull/5968))
* compute: added `local_ssd_recovery_timeout` field to `google_compute_instance_template` resource ([#5968](https://github.com/hashicorp/terraform-provider-google-beta/pull/5968))
* compute: added `local_ssd_recovery_timeout` field to `google_compute_regional_instance_template` resource ([#5968](https://github.com/hashicorp/terraform-provider-google-beta/pull/5968))
* compute: made `network_interface.ipv6_access_config.external_ipv6` configurable in `google_compute_instance` ([#5986](https://github.com/hashicorp/terraform-provider-google-beta/pull/5986))
* container: added `enable_k8s_beta_apis.enabled_apis` field to `google_container_cluster` ([#5961](https://github.com/hashicorp/terraform-provider-google-beta/pull/5961))
* container: added `node_config.host_maintenance_policy` field to `google_container_cluster` and `google_container_node_pool` ([#5983](https://github.com/hashicorp/terraform-provider-google-beta/pull/5983))
* container: added `placement_policy.policy_name` field to `google_container_node_pool` resource ([#5994](https://github.com/hashicorp/terraform-provider-google-beta/pull/5994))
* container: unsuppressed `private_cluster_config` when `master_global_access_config` is set in `google_container_cluster` ([#5995](https://github.com/hashicorp/terraform-provider-google-beta/pull/5995))
* container: allowed `enabled_private_endpoint` to be settable on creation for PSC-based clusters ([#5989](https://github.com/hashicorp/terraform-provider-google-beta/pull/5989))
* gkeonprem: added taint on failed resource creation for `google_gkeonprem_bare_metal_admin_cluster` ([#5990](https://github.com/hashicorp/terraform-provider-google-beta/pull/5990))
* gkeonprem: increased timeout for resources `google_gkeonprem_bare_metal_cluster` and `google_gkeonprem_bare_metal_admin_cluster` ([#5990](https://github.com/hashicorp/terraform-provider-google-beta/pull/5990))
* identityplayform: added support for `blocking_functions` `quota` and `authorized_domains` in `google_identity_platform_config` ([#5964](https://github.com/hashicorp/terraform-provider-google-beta/pull/5964))
* monitoring: added update support for `period` in `google_monitoring_uptime_check_config` ([#5959](https://github.com/hashicorp/terraform-provider-google-beta/pull/5959))
* pubsub: added `no_wrapper` field to `google_pubsub_subscription` resource ([#5972](https://github.com/hashicorp/terraform-provider-google-beta/pull/5972))
* workstations: added `accelerators` field to `google_workstations_workstation_config` resource ([#5991](https://github.com/hashicorp/terraform-provider-google-beta/pull/5991))

BUG FIXES:
* bigquery: fixed a bug in update support for several fields in `google_bigquery_data_transfer_config` ([#5987](https://github.com/hashicorp/terraform-provider-google-beta/pull/5987))
* cloudfunctions2: fixed an issue where `google_cloudfunctions2_function.build_config.source.storage_source.generation` created a diff when not set in config ([#5992](https://github.com/hashicorp/terraform-provider-google-beta/pull/5992))
* firebasedatabase: fixed empty `database_url` output attribute ([#5988](https://github.com/hashicorp/terraform-provider-google-beta/pull/5988))
* monitoring: fixed an issue in `google_monitoring_monitored_project` where project numbers were not accepted for `name` ([#5955](https://github.com/hashicorp/terraform-provider-google-beta/pull/5955))
* vpcaccess: reverted new behaviour introduced by resource `google_vpc_access_connector` in `4.75.0`. `min_throughput` and `max_throughput` fields lost their default value, and customers could not make deployment due to that change. ([#5957](https://github.com/hashicorp/terraform-provider-google-beta/pull/5957))

## 4.76.0 (July 31, 2023)

FEATURES:
* **New Resource:** `google_dataplex_task` ([#5914](https://github.com/hashicorp/terraform-provider-google-beta/pull/5914))
* **New Resource:** `google_iap_web_region_backend_service_iam_binding` ([#5944](https://github.com/hashicorp/terraform-provider-google-beta/pull/5944))
* **New Resource:** `google_iap_web_region_backend_service_iam_member` ([#5944](https://github.com/hashicorp/terraform-provider-google-beta/pull/5944))
* **New Resource:** `google_iap_web_region_backend_service_iam_policy` ([#5944](https://github.com/hashicorp/terraform-provider-google-beta/pull/5944))

IMPROVEMENTS:
* compute: added `security_policy` field to `google_compute_region_backend_service` resource ([#5924](https://github.com/hashicorp/terraform-provider-google-beta/pull/5924))
* cloudrun: added `status.traffic` output fields to `google_cloud_run_service` resource ([#5943](https://github.com/hashicorp/terraform-provider-google-beta/pull/5943))
* cloudrunv2: added field `custom_audiences` to resource `google_cloud_run_v2_service ` ([#5935](https://github.com/hashicorp/terraform-provider-google-beta/pull/5935))
* composer: added support for updating `resilience_mode` in `google_composer_environment` ([#5921](https://github.com/hashicorp/terraform-provider-google-beta/pull/5921))
* compute: added `reconcile_connections` for `google_compute_service_attachment`. ([#5945](https://github.com/hashicorp/terraform-provider-google-beta/pull/5945))
* container : added `gcs_fuse_csi_driver_config` field to `addons_config` in `google_container_cluster` resource. ([#5946](https://github.com/hashicorp/terraform-provider-google-beta/pull/5946))
* container: added `allow_net_admin` field to `google_container_cluster` resource ([#5940](https://github.com/hashicorp/terraform-provider-google-beta/pull/5940))
* container: added multi-NIC network for `google_container_cluster` and `google_container_node_pool`. ([#5949](https://github.com/hashicorp/terraform-provider-google-beta/pull/5949))
* container: allowed user to set up to 20 maintenance exclusions for `google_container_cluster` resource ([#5947](https://github.com/hashicorp/terraform-provider-google-beta/pull/5947))
* healthcare: added `last_updated_partition_config` field to `google_healthcare_fhir_store` resource ([#5937](https://github.com/hashicorp/terraform-provider-google-beta/pull/5937))
* monitoring: added `condition_prometheus_query_language` field to `google_monitoring_alert_policy` resource ([#5952](https://github.com/hashicorp/terraform-provider-google-beta/pull/5952))
* networkservices: made `scope` field optional in `google_network_services_gateway` resource ([#5939](https://github.com/hashicorp/terraform-provider-google-beta/pull/5939))
* spanner: added `enable_drop_protection` to `google_spanner_database` resource([#5942](https://github.com/hashicorp/terraform-provider-google-beta/pull/5942))

BUG FIXES:
* alloydb: fixed permadiffs when setting 0 as start time (midnight) for `automated_backup_policy` in `google_alloydb_cluster` resource ([#5913](https://github.com/hashicorp/terraform-provider-google-beta/pull/5913))
* artifactregistry: fixed reading back maven_config state in `google_artifact_registry_repository` ([#5936](https://github.com/hashicorp/terraform-provider-google-beta/pull/5936))
* cloudtasks: suppressed time-unit permadiffs on `google_cloud_tasks_queue` min and max backoff settings ([#5920](https://github.com/hashicorp/terraform-provider-google-beta/pull/5920))
* cloudrun: fixed the bug where default system labels set in `service.spec.template.metadata.labels` were treated as a diff. ([#5953](https://github.com/hashicorp/terraform-provider-google-beta/pull/5953))
* compute: fixed wrongly triggered recreation on changes of `enforce_on_key_configs` on `google_compute_security_policy` ([#5928](https://github.com/hashicorp/terraform-provider-google-beta/pull/5928))
* monitoring: fixed an issue in `google_monitoring_monitored_project` where project numbers were not accepted for `name` ([#5955](https://github.com/hashicorp/terraform-provider-google-beta/pull/5955))


## 4.75.1 (July 27, 2023)

BUG FIXES:

* vpcaccess: reverted new behaviour introduced by resource `google_vpc_access_connector` in `4.75.0`. `min_throughput` and `max_throughput` fields lost their default value, and customers could not make deployment due to that change.

* vpcaccess: reverted the ability to update the number of instances for resource `google_vpc_access_connector`

## 4.75.0 (July 24, 2023)

FEATURES:
* **New Resource:** `google_looker_instance` ([#5903](https://github.com/hashicorp/terraform-provider-google-beta/pull/5903))

IMPROVEMENTS:
* apigee: added `disable_vpc_peering` field to `google_apigee_organization` resource ([#5901](https://github.com/hashicorp/terraform-provider-google-beta/pull/5901))
* bigquery: added `external_data_configuration.json_options` and `external_data_configuration.parquet_options` fields to `google_bigquery_table` ([#5906](https://github.com/hashicorp/terraform-provider-google-beta/pull/5906))
* bigtable: added `change_stream_retention` field to `google_bigtable_table.table` resource ([#5880](https://github.com/hashicorp/terraform-provider-google-beta/pull/5880))
* compute: added `most_recent` argument to `google_compute_image` datasource ([#5902](https://github.com/hashicorp/terraform-provider-google-beta/pull/5902))
* compute: added field `enable_confidential_compute` for `google_compute_disk` resource ([#5897](https://github.com/hashicorp/terraform-provider-google-beta/pull/5897))
* container: added `gpu_driver_installation_config.gpu_driver_version` field to `google_container_node_pool` ([#5899](https://github.com/hashicorp/terraform-provider-google-beta/pull/5899))
* gkebackup: added `state` and `state_reason` output-only fields to `google_gkebackup_backupplan` resource ([#5909](https://github.com/hashicorp/terraform-provider-google-beta/pull/5909))
* healthcare: added `complex_data_type_reference_parsing ` field to `google_healthcare_fhir_store` resource ([#5884](https://github.com/hashicorp/terraform-provider-google-beta/pull/5884))
* networkservices: increased max_size to 20 for both `included_query_parameters` and `excluded_query_parameters` on `google_network_services_edge_cache_service` ([#5889](https://github.com/hashicorp/terraform-provider-google-beta/pull/5889))
* vpcaccess: added support for updates to `google_vpc_access_connector` resource ([#5894](https://github.com/hashicorp/terraform-provider-google-beta/pull/5894))

BUG FIXES:
* alloydb: fixed `google_alloydb_cluster` handling of automated backup policy midnight start time ([#5913](https://github.com/hashicorp/terraform-provider-google-beta/pull/5913))
* compute: fixed logic when unsetting `google_compute_instance.min_cpu_platform` and switching to a `machine_type` that does not support `min_cpu_platform` at the same time ([#5911](https://github.com/hashicorp/terraform-provider-google-beta/pull/5911))
* tags: fixed race condition when modifying `google_tags_location_tag_binding` ([#5904](https://github.com/hashicorp/terraform-provider-google-beta/pull/5904))

## 4.74.0 (July 18, 2023)

IMPROVEMENTS:

* bigquery: added `storage_billing_model` argument to `google_bigquery_dataset` ([#5868](https://github.com/hashicorp/terraform-provider-google-beta/pull/5868))
* bigquery: added `external_data_configuration.metadata_cache_mode` and `external_data_configuration.object_metadata` to `google_bigquery_table` ([#5856](https://github.com/hashicorp/terraform-provider-google-beta/pull/5856))
* bigquery: made `external_data_configuration.source_fomat` optional in `google_bigquery_table` ([#5856](https://github.com/hashicorp/terraform-provider-google-beta/pull/5856))
* certificatemanager: added `issuance_config` field to `google_certificate_manager_certificate` resource ([#5860](https://github.com/hashicorp/terraform-provider-google-beta/pull/5860))
* cloudbuildv2: added `gitlab_config` field to `google_cloudbuildv2_connection` resource ([#5848](https://github.com/hashicorp/terraform-provider-google-beta/pull/5848))
* compute: added field `http_keep_alive_timeout_sec` to resource `google_compute_target_http_proxy` ([#5864](https://github.com/hashicorp/terraform-provider-google-beta/pull/5864))
* compute: added field `http_keep_alive_timeout_sec` to resource `google_compute_target_https_proxy` ([#5864](https://github.com/hashicorp/terraform-provider-google-beta/pull/5864))
* compute: added support for updating labels in `google_compute_external_vpn_gateway` ([#5875](https://github.com/hashicorp/terraform-provider-google-beta/pull/5875))
* container: added field `tpu_topology` under `placement_policy` in resource `google_container_node_pool` ([#5871](https://github.com/hashicorp/terraform-provider-google-beta/pull/5871))
* iamworkforcepool: added `oidc.client_secret` field to `google_iam_workforce_pool_provider` and new enum values `CODE` and `MERGE_ID_TOKEN_OVER_USER_INFO_CLAIMS` to `oidc.web_sso_config.response_type` and `oidc.web_sso_config.assertion_claims_behavior` respectively ([#5853](https://github.com/hashicorp/terraform-provider-google-beta/pull/5853))
* sql: added `settings.data_cache_config` to `sql_database_instance` resource. ([#5869](https://github.com/hashicorp/terraform-provider-google-beta/pull/5869))
* sql: added `settings.edition` field to `sql_database_instance` resource. ([#5869](https://github.com/hashicorp/terraform-provider-google-beta/pull/5869))
* vertexai: supported`shard_size` in `google_vertex_ai_index` ([#5874](https://github.com/hashicorp/terraform-provider-google-beta/pull/5874))

BUG FIXES:
* compute: made `google_compute_router_peer.peer_ip_address` optional ([#5855](https://github.com/hashicorp/terraform-provider-google-beta/pull/5855))
* redis: fixed issue with `google_redis_instance` populating output-only field `maintenance_schedule`. ([#5852](https://github.com/hashicorp/terraform-provider-google-beta/pull/5852))
* orgpolicy: fixed forcing recreation on imported state for `google_org_policy_policy` ([#5873](https://github.com/hashicorp/terraform-provider-google-beta/pull/5873))
* osconfig: fixed validation of file resource `state` fields in `google_os_config_os_policy_assignment` ([#5863](https://github.com/hashicorp/terraform-provider-google-beta/pull/5863))

## 4.73.2 (July 17, 2023)

BUG FIXES:
* monitoring: fixed an issue which occurred when `name` field of `google_monitoring_monitored_project` was long-form

## 4.73.1 (July 13, 2023)

BUG FIXES:
* monitoring: fixed an issue causing `google_monitoring_monitored_project` to appear to be deleted

## 4.73.0 (July 10, 2023)

FEATURES:
* **New Resource:** `google_firebase_extensions_instance` ([#5832](https://github.com/hashicorp/terraform-provider-google-beta/pull/5832))

IMPROVEMENTS:
* compute: added the `no_automate_dns_zone` field to `google_compute_forwarding_rule`. ([#5842](https://github.com/hashicorp/terraform-provider-google-beta/pull/5842))
* compute: promoted `google_compute_disk_async_replication` resource to GA. ([#5843](https://github.com/hashicorp/terraform-provider-google-beta/pull/5843))
* compute: promoted `async_primary_disk` field in `google_compute_disk` resource to GA. ([#5843](https://github.com/hashicorp/terraform-provider-google-beta/pull/5843))
* compute: promoted `async_primary_disk` field in `google_compute_region_disk` resource to GA. ([#5843](https://github.com/hashicorp/terraform-provider-google-beta/pull/5843))
* compute: promoted `disk_consistency_group_policy` field in `google_compute_resource_policy` resource to GA. ([#5843](https://github.com/hashicorp/terraform-provider-google-beta/pull/5843))
* resourcemanager: fixed handling of `google_service_account_id_token` when authenticated with GCE metadata credentials ([#5825](https://github.com/hashicorp/terraform-provider-google-beta/pull/5825))

BUG FIXES:
* networkservices: increased default timeout for `google_network_services_edge_cache_keyset` to 90m ([#5839](https://github.com/hashicorp/terraform-provider-google-beta/pull/5839))

## 4.72.1 (July 6, 2023)

BUG FIXES:
* compute: fixed an issue in `google_compute_instance_template` where initialize params stopped the `disk.disk_size_gb` field being used ([#5849](https://github.com/hashicorp/terraform-provider-google-beta/pull/5849))

## 4.72.0 (July 3, 2023)

FEATURES:
* **New Resource:** `google_public_ca_external_account_key` ([#5813](https://github.com/hashicorp/terraform-provider-google-beta/pull/5813))
* **New Resource:** `google_compute_network_edge_security_service` ([#5808](https://github.com/hashicorp/terraform-provider-google-beta/pull/5808))
* **New Resource:** `google_compute_region_security_policy` ([#5808](https://github.com/hashicorp/terraform-provider-google-beta/pull/5808))

IMPROVEMENTS:
* cloudrunv2: added annotations support to google_cloud_run_v2_job ([#5795](https://github.com/hashicorp/terraform-provider-google-beta/pull/5795))
* compute: added `provisioned_throughput` field to `google_compute_disk` used by `hyperdisk-throughput` pd type ([#5814](https://github.com/hashicorp/terraform-provider-google-beta/pull/5814))
* compute: added field `http_keep_alive_timeout_sec` to resource `google_compute_target_http_proxy` ([#5818](https://github.com/hashicorp/terraform-provider-google-beta/pull/5818))
* compute: added field `http_keep_alive_timeout_sec` to resource `google_compute_target_https_proxy` ([#5818](https://github.com/hashicorp/terraform-provider-google-beta/pull/5818))
* container: added field `security_posture_config` to resource `google_container_cluster` ([#5821](https://github.com/hashicorp/terraform-provider-google-beta/pull/5821))
* logging: added support for `locked` to `google_logging_project_bucket_config` ([#5811](https://github.com/hashicorp/terraform-provider-google-beta/pull/5811))

BUG FIXES:
* bigquery: fixed an issue where api default value for `edition` field of `google_bigquery_reservation` was not handled ([#5800](https://github.com/hashicorp/terraform-provider-google-beta/pull/5800))
* cloudfunction2: fixed permadiffs of some fields of `service_config` in `google_cloudfunctions2_function` resource ([#5810](https://github.com/hashicorp/terraform-provider-google-beta/pull/5810))
* compute: fixed an issue with setting project field to long form in `google_compute_forwarding_rule` and `google_compute_global_forwarding_rule` ([#5820](https://github.com/hashicorp/terraform-provider-google-beta/pull/5820))
* gkehub: fixed an issue with setting project field to long form in `google_gke_hub_feature` ([#5820](https://github.com/hashicorp/terraform-provider-google-beta/pull/5820)

## 4.71.0 (June 27, 2023)

FEATURES:
* **New Resource:** `google_gke_hub_feature_iam_*` ([#5782](https://github.com/hashicorp/terraform-provider-google-beta/pull/5782))
* **New Resource:** `google_gke_hub_feature` ([#5782](https://github.com/hashicorp/terraform-provider-google-beta/pull/5782))
* **New Resource:** `google_vmwareengine_cluster` ([#5784](https://github.com/hashicorp/terraform-provider-google-beta/pull/5784))
* **New Resource:** `google_vmwareengine_private_cloud` ([#5784](https://github.com/hashicorp/terraform-provider-google-beta/pull/5784))

IMPROVEMENTS:
* apigee: added output-only field `apigee_project_id` to resource `google_apigee_organization` ([#5781](https://github.com/hashicorp/terraform-provider-google-beta/pull/5781))
* bigtable: increased default timeout for instance operations to 1 hour in resoure `google_bigtable_instance` ([#5779](https://github.com/hashicorp/terraform-provider-google-beta/pull/5779))
* composer: added field `resilience_mode` to resource `google_composer_environment` ([#5790](https://github.com/hashicorp/terraform-provider-google-beta/pull/5790))
* compute: added support for `params.resource_manager_tags` and `boot_disk.initialize_params.resource_manager_tags` to resource `google_compute_instance` ([#5787](https://github.com/hashicorp/terraform-provider-google-beta/pull/5787))
* bigquerydatatransfer: made field `service_account_name` mutable in resource `google_bigquery_data_transfer_config` ([#5777](https://github.com/hashicorp/terraform-provider-google-beta/pull/5777))
* iambeta: added field `jwks_json` to resource `google_iam_workload_identity_pool_provider` ([#5789](https://github.com/hashicorp/terraform-provider-google-beta/pull/5789))

BUG FIXES:
* bigtable: validated that `cluster_id` values are unique within resource `google_bigtable_instance` ([#5778](https://github.com/hashicorp/terraform-provider-google-beta/pull/5778))
* storage: fixed a bug that caused a permadiff when the `autoclass.enabled` field was explicitly set to false in resource `google_storage_bucket` ([#5776](https://github.com/hashicorp/terraform-provider-google-beta/pull/5776))

## 4.70.0 (June 20, 2023)

FEATURES:
* **New Resource:** `google_compute_network_endpoints` ([#5756](https://github.com/hashicorp/terraform-provider-google-beta/pull/5756))
* **New Resource:** `vertex_ai_index_endpoint` ([#5738](https://github.com/hashicorp/terraform-provider-google-beta/pull/5738))

IMPROVEMENTS:
* bigtable: added 20 minutes timeout support to `google_bigtable_gc_policy` ([#5752](https://github.com/hashicorp/terraform-provider-google-beta/pull/5752))
* cloudfunctions2: added `url` output field to `google_cloudfunctions2_function` ([#5745](https://github.com/hashicorp/terraform-provider-google-beta/pull/5745))
* compute: added field `network_attachment` to `google_compute_instance_template` ([#5761](https://github.com/hashicorp/terraform-provider-google-beta/pull/5761))
* compute: surfaced additional information about quota exceeded errors for compute resources. ([#5763](https://github.com/hashicorp/terraform-provider-google-beta/pull/5763))
* compute: added `path_template_match` and `path_template_rewrite` to `google_compute_url_map`. ([#5760](https://github.com/hashicorp/terraform-provider-google-beta/pull/5760))
* container: added `sole_tenant_config` to `node_config` in `google_container_node_pool` and `google_container_cluster` ([#5774](https://github.com/hashicorp/terraform-provider-google-beta/pull/5774))
* dataform: added field `workspace_compilation_overrides` to resource `google_dataform_repository` (beta) ([#5736](https://github.com/hashicorp/terraform-provider-google-beta/pull/5736))
* dlp: added `crypto_hash_config` to `google_data_loss_prevention_deidentify_template` ([#5757](https://github.com/hashicorp/terraform-provider-google-beta/pull/5757))
* dlp: added `trigger_id` field to `google_data_loss_prevention_job_trigger` ([#5773](https://github.com/hashicorp/terraform-provider-google-beta/pull/5773))
* dlp: added missing file types `POWERPOINT` and `EXCEL` in `inspect_job.storage_config.cloud_storage_options.file_types` enum to `google_data_loss_prevention_job_trigger` resource ([#5749](https://github.com/hashicorp/terraform-provider-google-beta/pull/5749))
* dlp: added multiple `sensitivity_score` field to `google_data_loss_prevention_deidentify_template` resource ([#5764](https://github.com/hashicorp/terraform-provider-google-beta/pull/5764))
* dlp: added multiple `sensitivity_score` field to `google_data_loss_prevention_inspect_template` resource ([#5758](https://github.com/hashicorp/terraform-provider-google-beta/pull/5758))
* dlp: added multiple `sensitivity_score` field to `google_data_loss_prevention_job_trigger` resource ([#5765](https://github.com/hashicorp/terraform-provider-google-beta/pull/5765))
* pubsub: allowed `definition` field of `google_pubsub_schema` updatable. (https://cloud.google.com/pubsub/docs/schemas#commit-schema-revision) ([#5750](https://github.com/hashicorp/terraform-provider-google-beta/pull/5750))
* sql: added `POSTGRES_15` to version docs for `database_version` field to `google_sql_database_instance` ([#5772](https://github.com/hashicorp/terraform-provider-google-beta/pull/5772))
* vpcaccess: Added `connected_projects` attribute to resource `google_vpc_access_connector`. ([#5734](https://github.com/hashicorp/terraform-provider-google-beta/pull/5734))

BUG FIXES:
* filestore: fixed an issue on multiple resources where non-retryable quota errors were considered retryable ([#5744](https://github.com/hashicorp/terraform-provider-google-beta/pull/5744))
* vertexai: made `google_vertex_ai_featurestore_entitytype_feature` always use region corresponding to parent's region ([#5739](https://github.com/hashicorp/terraform-provider-google-beta/pull/5739))

## 4.69.1 (June 12, 2023)

NOTE:
* Added a new user guide to the provider documentation ([#5768](https://github.com/hashicorp/terraform-provider-google-beta/pull/5768))

## 4.69.0 (June 12, 2023)

FEATURES:
* **New Data Source:** `google_vmwareengine_network` ([#5725](https://github.com/hashicorp/terraform-provider-google-beta/pull/5725))
* **New Resource:** `google_access_context_manager_service_perimeter_egress_policy` ([#5723](https://github.com/hashicorp/terraform-provider-google-beta/pull/5723))
* **New Resource:** `google_access_context_manager_service_perimeter_ingress_policy` ([#5723](https://github.com/hashicorp/terraform-provider-google-beta/pull/5723))
* **New Resource:** `google_certificate_manager_certificate_issuance_config` ([#5712](https://github.com/hashicorp/terraform-provider-google-beta/pull/5712))
* **New Resource:** `google_dataplex_datascan` ([#5707](https://github.com/hashicorp/terraform-provider-google-beta/pull/5707))
* **New Resource:** `google_dataplex_datascan_iam_*` ([#5731](https://github.com/hashicorp/terraform-provider-google-beta/pull/5731))
* **New Resource:** `google_vmwareengine_network` ([#5725](https://github.com/hashicorp/terraform-provider-google-beta/pull/5725))

IMPROVEMENTS:
* billing: added `lookup_projects` to `google_billing_account` datasource that skips reading the list of associated projects ([#5721](https://github.com/hashicorp/terraform-provider-google-beta/pull/5721))
* dlp: added `info_type_transformations` block in the `record_transformations` field to `google_data_loss_prevention_deidentify_template` resource. ([#5729](https://github.com/hashicorp/terraform-provider-google-beta/pull/5729))
* dlp: added `redact_config`, `fixed_size_bucketing_config`, `bucketing_config`, `time_part_config` and `date_shift_config`  fields to `google_data_loss_prevention_deidentify_template` resource ([#5711](https://github.com/hashicorp/terraform-provider-google-beta/pull/5711))
* dlp: added `stored_info_type_id` field to `google_data_loss_prevention_stored_info_type` resource ([#5708](https://github.com/hashicorp/terraform-provider-google-beta/pull/5708))
* dlp: added `template_id` field to `google_data_loss_prevention_deidentify_template` and `google_data_loss_prevention_inspect_template` ([#5726](https://github.com/hashicorp/terraform-provider-google-beta/pull/5726))
* dlp: changed `actions` field from required to optional in `google_data_loss_prevention_job_trigger` resource ([#5716](https://github.com/hashicorp/terraform-provider-google-beta/pull/5716))
* gkehub: added field `fleet_observability` to `google_gke_hub_feature` ([#5715](https://github.com/hashicorp/terraform-provider-google-beta/pull/5715))
* kms: removed validation for `purpose` in `google_kms_crypto_key` to allow newly added values for the field ([#5713](https://github.com/hashicorp/terraform-provider-google-beta/pull/5713))
* networkservices: added necessary fields to `google_network_services_gateway` to make it compatible with secure web proxy ([#5701](https://github.com/hashicorp/terraform-provider-google-beta/pull/5701))
* pubsub: allowed `schema_settings` of `google_pubsub_topic` to change without deleting and recreating the resource ([#5724](https://github.com/hashicorp/terraform-provider-google-beta/pull/5724))
* vertexai: increased `google_vertex_ai_metadata_store` creation timeout to 40 minutes ([#5728](https://github.com/hashicorp/terraform-provider-google-beta/pull/5728))

BUG FIXES:
* networkservices: fixed a bug where modifying non-updatable fields `scope` in `google_network_services_gateway` would fail with API errors; now updating them will recreate the resource ([#5701](https://github.com/hashicorp/terraform-provider-google-beta/pull/5701))
* tags: fixed providing `projects/<project_id` to `parent` causing recreation on `google_tags_tag_key` ([#5718](https://github.com/hashicorp/terraform-provider-google-beta/pull/5718))

## 4.68.0 (June 5, 2023)

FEATURES:
* **New Resource:** `google_container_analysis_note_iam_*` ([#5676](https://github.com/hashicorp/terraform-provider-google-beta/pull/5676))

IMPROVEMENTS:
* dlp: added `included_fields` and `excluded_fields` fields to `google_data_loss_prevention_job_trigger` ([#5687](https://github.com/hashicorp/terraform-provider-google-beta/pull/5687))
* dns: added `regionalL7ilb` enum support to the `routing_policy.load_balancer_type` field in `google_dns_record_set` ([#5678](https://github.com/hashicorp/terraform-provider-google-beta/pull/5678))
* workstations: added `idle_timeout` and `running_timeout` fields in `google_workstations_workstation_config` ([#5673](https://github.com/hashicorp/terraform-provider-google-beta/pull/5673))
* workstations: added update support for `persistent_directories.reclaim_policy` and `persistent_directories.source_snapshot` fields in `google_workstations_workstation_config` ([#5695](https://github.com/hashicorp/terraform-provider-google-beta/pull/5695))

BUG FIXES:
* accesscontextmanager: fixed incorrect validations for `spec` and `status` in `google_access_context_manager_service_perimeter` ([#5675](https://github.com/hashicorp/terraform-provider-google-beta/pull/5675))
* alloydb: increased timeouts for `google_alloydb_instance` from 20m to 40m ([#5681](https://github.com/hashicorp/terraform-provider-google-beta/pull/5681))
* apigee: fixed bug where updating `config_bundle` in `google_apigee_sharedflow` that's attached to `google_apigee_sharedflow_deployment` causes an error ([#5683](https://github.com/hashicorp/terraform-provider-google-beta/pull/5683))
* compute: increased timeout for `compute_security_policy` from 4m to 8m ([#5680](https://github.com/hashicorp/terraform-provider-google-beta/pull/5680))
* dataproc: fixed crash when reading `google_dataproc_cluster.virtual_cluster_config` ([#5689](https://github.com/hashicorp/terraform-provider-google-beta/pull/5689))

## 4.67.0 (May 30, 2023)

FEATURES:
* **New Data Source:** `google_*_iam_policy` ([#5661](https://github.com/hashicorp/terraform-provider-google-beta/pull/5661))
* **New Data Source:** `google_vertex_ai_index` ([#5649](https://github.com/hashicorp/terraform-provider-google-beta/pull/5649))

IMPROVEMENTS:
* cloudrun: added `template.spec.volumes.empty_dir` and `template.spec.containers.name` fields to `google_cloud_run_service` ([#5654](https://github.com/hashicorp/terraform-provider-google-beta/pull/5654))
* compute: added `guest_os_features` and `licenses` fields to `google_compute_disk` and `google_compute_region_disk` ([#5659](https://github.com/hashicorp/terraform-provider-google-beta/pull/5659))
* datastream: added `mysql_source_config.max_concurrent_backfill_tasks` field to `google_datastream_stream` ([#5648](https://github.com/hashicorp/terraform-provider-google-beta/pull/5648))
* firebase: added additional import formats for `google_firebase_webapp` ([#5647](https://github.com/hashicorp/terraform-provider-google-beta/pull/5647))
* notebooks: added update support for `google_notebooks_instance.metadata` field ([#5655](https://github.com/hashicorp/terraform-provider-google-beta/pull/5655))
* privateca: added `encoding_format` field to `google_privateca_ca_pool` ([#5662](https://github.com/hashicorp/terraform-provider-google-beta/pull/5662))

BUG FIXES:
* apigee: increased `google_apigee_organization` timeout defaults to 45m from 20m ([#5652](https://github.com/hashicorp/terraform-provider-google-beta/pull/5652))
* cloudresourcemanager: added retries to handle internal error: type: "googleapis.com" subject: "160009" ([#5685](https://github.com/hashicorp/terraform-provider-google-beta/pull/5685))
* cloudrun: fixed a permadiff for `metadata.annotation` in `google_cloud_run_service` ([#5651](https://github.com/hashicorp/terraform-provider-google-beta/pull/5651))
* container: fixed a crash scenario in `google_container_node_pool` ([#5671](https://github.com/hashicorp/terraform-provider-google-beta/pull/5671))
* gkeonprem: changed `hostname` (under `ip_block`) from required to optional for `google_gkeonprem_vmware_cluster` ([#5670](https://github.com/hashicorp/terraform-provider-google-beta/pull/5670))
* serviceusage: added retries to handle internal error: type: "googleapis.com" subject: "160009" when activating services ([#5685](https://github.com/hashicorp/terraform-provider-google-beta/pull/5685))

## 4.66.0 (May 22, 2023)

NOTE:
* Upgraded to Go 1.19.9 ([#5623](https://github.com/hashicorp/terraform-provider-google-beta/pull/5623))

FEATURES:
* **New Resource:** `google_network_security_server_tls_policy` ([#5619](https://github.com/hashicorp/terraform-provider-google-beta/pull/5619))

IMPROVEMENTS:
* bigquery: added `ICEBERG` as an enum for `external_data_configuration.source_format` field in `google_bigquery_table` ([#5622](https://github.com/hashicorp/terraform-provider-google-beta/pull/5622))
* cloudfunctions: added `status` attribute to the `google_cloudfunctions_function` resource and data source ([#5625](https://github.com/hashicorp/terraform-provider-google-beta/pull/5625))
* compute: added `storage_location` field in `google_compute_image` resource ([#5644](https://github.com/hashicorp/terraform-provider-google-beta/pull/5644))
* compute: added support for additional machine types in `google_compute_region_commitment` ([#5633](https://github.com/hashicorp/terraform-provider-google-beta/pull/5633))
* dataflow: added multiple fields to `google_dataflow_flex_template_job` ([#5635](https://github.com/hashicorp/terraform-provider-google-beta/pull/5635))
* monitoring: added `forecast_options` field to `google_monitoring_alert_policy` resource ([#5642](https://github.com/hashicorp/terraform-provider-google-beta/pull/5642))
* monitoring: added `notification_channel_strategy` field to `google_monitoring_alert_policy` resource ([#5624](https://github.com/hashicorp/terraform-provider-google-beta/pull/5624))
* sql: added `advanced_machine_features` field in `google_sql_database_instance` ([#5639](https://github.com/hashicorp/terraform-provider-google-beta/pull/5639))
* storagetransfer: added field `path` to `transfer_spec.aws_s3_data_source` in `google_storage_transfer_job` ([#5641](https://github.com/hashicorp/terraform-provider-google-beta/pull/5641))
* workstations: added support for `source_snapshot` in `google_workstations_workstation_config` ([#5636](https://github.com/hashicorp/terraform-provider-google-beta/pull/5636))

BUG FIXES:
* artifactregistry: fixed new repositories ignoring the provider region if location is unset in `google_artifact_registry_repository`. ([#5637](https://github.com/hashicorp/terraform-provider-google-beta/pull/5637))
* compute: fixed permadiff on `log_config.sample_rate` of `google_compute_backend_service` ([#5631](https://github.com/hashicorp/terraform-provider-google-beta/pull/5631))
* container: fixed permadiff on `gateway_api_config.channel` of `google_container_cluster` ([#5626](https://github.com/hashicorp/terraform-provider-google-beta/pull/5626))
* dataflow: fixed inconsistent final plan when labels are added to `google_dataflow_job` ([#5634](https://github.com/hashicorp/terraform-provider-google-beta/pull/5634))
* provider: fixed an issue where mtls transports were not used consistently(initial implementation in v4.65.0, reverted in v4.65.1) ([#5645](https://github.com/hashicorp/terraform-provider-google-beta/pull/5645))
* storage: fixed inconsistent final plan when labels are added to `google_storage_bucket` ([#5634](https://github.com/hashicorp/terraform-provider-google-beta/pull/5634))

## 4.65.2 (May 16, 2023)

BUG FIXES:
* provider: fixed an issue where `google_client_config` datasource return `null` for all attributes when region or zone is unset in provider config

## 4.65.1 (May 15, 2023)

BUG FIXES:
* provider: fixed an issue where `google_client_config` datasource return `null` for `access_token`

## 4.65.0 (May 15, 2023)

FEATURES:
* **New Data Source:** `google_datastream_static_ips` ([#5587](https://github.com/hashicorp/terraform-provider-google-beta/pull/5587))
* **New Resource:** `google_compute_disk_async_replication` ([#5588](https://github.com/hashicorp/terraform-provider-google-beta/pull/5588))
* **New Resource:** `google_firestore_field` ([#5603](https://github.com/hashicorp/terraform-provider-google-beta/pull/5603))
* **New Resource:** `google_gkeonprem_bare_metal_cluster` ([#5594](https://github.com/hashicorp/terraform-provider-google-beta/pull/5594))
* **New Resource:** `google_gkeonprem_bare_metal_node_pool` ([#5602](https://github.com/hashicorp/terraform-provider-google-beta/pull/5602))
* **New Resource:** `google_network_security_tls_inspection_policy` ([#5615](https://github.com/hashicorp/terraform-provider-google-beta/pull/5615))

IMPROVEMENTS:
* bigquery: added general field `load.parquet_options` to `google_bigquery_job` ([#5592](https://github.com/hashicorp/terraform-provider-google-beta/pull/5592))
* cloudbuild: added `allow_failure` and `allow_exit_codes` to `build.step` in `google_cloudbuild_trigger` resource ([#5593](https://github.com/hashicorp/terraform-provider-google-beta/pull/5593))
* cloudbuild: added `git_file_source.repository` and `source_to_build.repository` fields to `google_cloudbuild_trigger` resource (beta) ([#5597](https://github.com/hashicorp/terraform-provider-google-beta/pull/5597))
* cloudrunv2: added `template.containers.depends_on` and `template.volumes.empty_dir` to `google_cloud_run_v2_service`. ([#5613](https://github.com/hashicorp/terraform-provider-google-beta/pull/5613))
* cloudrunv2: added `template.template.volumes.empty_dir` to `google_cloud_run_v2_job`. ([#5613](https://github.com/hashicorp/terraform-provider-google-beta/pull/5613))
* compute: added enumeration values `SEV_SNP_CAPABLE`, `SUSPEND_RESUME_COMPATIBLE`, `TDX_CAPABLE` for the `guest_os_features` of `google_compute_image` ([#5604](https://github.com/hashicorp/terraform-provider-google-beta/pull/5604))
* compute: added support for `stack_type` to `google_compute_network_peering` ([#5601](https://github.com/hashicorp/terraform-provider-google-beta/pull/5601))
* container: added `gcs_fuse_csi_driver_config` to `google_container_cluster` resource. ([#5616](https://github.com/hashicorp/terraform-provider-google-beta/pull/5616))
* dlp: added `publish_to_stackdriver` field to `google_data_loss_prevention_job_trigger` resource ([#5610](https://github.com/hashicorp/terraform-provider-google-beta/pull/5610))
* network_security: added `tls_inspection_policy` field to `google_network_security_gateway_security_policy` ([#5615](https://github.com/hashicorp/terraform-provider-google-beta/pull/5615))

BUG FIXES:
* certificatemanager: fixed an issue where `self_managed.pem_certificate` and `self_managed.pem_certificate` can't be updated on `google_certificate_manager_certificate` ([#5606](https://github.com/hashicorp/terraform-provider-google-beta/pull/5606))
* compute: fixed crash on `terraform destroy -refresh=false` for instance group managers with `wait_for_instances = "true"` if the instance group manager was not found ([#5614](https://github.com/hashicorp/terraform-provider-google-beta/pull/5614))
* container: fixed node auto-provisioning not working when `auto_provisioning_defaults.management` is not provided on `google_container_cluster` ([#5605](https://github.com/hashicorp/terraform-provider-google-beta/pull/5605))
* provider: fixed an issue where mtls transports were not used consistently ([#5618](https://github.com/hashicorp/terraform-provider-google-beta/pull/5618))

## 4.64.0 (May 8, 2023)

FEATURES:
* **New Data Source:** `google_alloydb_locations` ([#5507](https://github.com/hashicorp/terraform-provider-google-beta/pull/5507))
* **New Data Source:** `google_sql_tiers` ([#5548](https://github.com/hashicorp/terraform-provider-google-beta/pull/5548))
* **New Resource:** `google_access_context_manager_egress_policy` ([#5525](https://github.com/hashicorp/terraform-provider-google-beta/pull/5525))
* **New Resource:** `google_database_migration_service_connection_profile` ([#5527](https://github.com/hashicorp/terraform-provider-google-beta/pull/5527))
* **New Resource:** `google_gkeonprem_vmware_cluster` ([#5533](https://github.com/hashicorp/terraform-provider-google-beta/pull/5533))
* **New Resource:** `google_gkeonprem_vmware_node_pool` ([#5579](https://github.com/hashicorp/terraform-provider-google-beta/pull/5579))
* **New Resource:** `google_network_security_address_group` ([#5539](https://github.com/hashicorp/terraform-provider-google-beta/pull/5539))
* **New Resource:** `google_network_security_authorization_policy` ([#5582](https://github.com/hashicorp/terraform-provider-google-beta/pull/5582))
* **New Resource:** `google_network_services_grpc_route` ([#5572](https://github.com/hashicorp/terraform-provider-google-beta/pull/5572))
* **New Resource:** `google_network_services_service_binding` ([#5536](https://github.com/hashicorp/terraform-provider-google-beta/pull/5536))
* **New Resource:** `google_networksecurity_client_tls_policy` ([#5561](https://github.com/hashicorp/terraform-provider-google-beta/pull/5561))
* **New Resource:** `google_networkservices_endpoint_policy` ([#5542](https://github.com/hashicorp/terraform-provider-google-beta/pull/5542))
* **New Resource:** `google_networkservices_tls_route` ([#5524](https://github.com/hashicorp/terraform-provider-google-beta/pull/5524))
* **New Resource:** `google_workstations_workstation_config_iam` ([#5512](https://github.com/hashicorp/terraform-provider-google-beta/pull/5512))
* **New Resource:** `google_workstations_workstation_iam` ([#5512](https://github.com/hashicorp/terraform-provider-google-beta/pull/5512))

IMPROVEMENTS:
* alloydb: added `encryption_config` and `encryption_info` fields in `google_alloydb_cluster`, to allow CMEK encryption of the cluster's data. ([#5551](https://github.com/hashicorp/terraform-provider-google-beta/pull/5551))
* alloydb: added support for CMEK in `google_alloydb_backup` resource ([#5549](https://github.com/hashicorp/terraform-provider-google-beta/pull/5549))
* alloydb: added the `encryption_config` field inside the `automated_backup_policy` block in`google_alloydb_cluster`, to allow CMEK encryption of automated backups. ([#5551](https://github.com/hashicorp/terraform-provider-google-beta/pull/5551))
* certificatemanager: added `location` field to `certificatemanager` certificate resource ([#5554](https://github.com/hashicorp/terraform-provider-google-beta/pull/5554))
* cloudrun: added field `port` to `http_get` to resource `google_cloud_run_service` ([#5510](https://github.com/hashicorp/terraform-provider-google-beta/pull/5510))
* cloudrunv2: added field `port` to `http_get` to resource `google_cloud_run_v2_service` ([#5510](https://github.com/hashicorp/terraform-provider-google-beta/pull/5510))
* cloudrunv2: added field `startupCpuBoost` to resource `google_cloud_run_v2_service` ([#5521](https://github.com/hashicorp/terraform-provider-google-beta/pull/5521))
* cloudrunv2: added support for `session_affinity` to `google_cloud_run_v2_service` ([#5518](https://github.com/hashicorp/terraform-provider-google-beta/pull/5518))
* compute: added `allow_psc_global_access` to `google_compute_forwarding_rule` resource ([#5523](https://github.com/hashicorp/terraform-provider-google-beta/pull/5523))
* compute: added `dest_fqdns`, `dest_region_codes`, `dest_threat_intelligences`, `src_fqdns`, `src_region_codes`, and `src_threat_intelligences` to `google_compute_firewall_policy_rule` resource. ([#5523](https://github.com/hashicorp/terraform-provider-google-beta/pull/5523))
* compute: added `source_ip_ranges` and `base_forwarding_rule` to `google_compute_forwarding_rule` resource ([#5523](https://github.com/hashicorp/terraform-provider-google-beta/pull/5523))
* compute: added `bypass_cache_on_request_headers` to `cdn_policy` in `google_compute_backend_service` resource ([#5563](https://github.com/hashicorp/terraform-provider-google-beta/pull/5563))
* compute: added `dest_address_groups` and `src_address_groups` fields to `google_compute_firewall_policy_rule` and `google_compute_network_firewall_policy_rule` ([#5530](https://github.com/hashicorp/terraform-provider-google-beta/pull/5530))
* compute: added new field `async_primary_disk` to `google_compute_disk` and `google_compute_region_disk` ([#5553](https://github.com/hashicorp/terraform-provider-google-beta/pull/5553))
* compute: added new field `disk_consistency_group_policy` to `google_compute_resource_policy` ([#5553](https://github.com/hashicorp/terraform-provider-google-beta/pull/5553))
* compute: added support for IPv6 prefix exchange in `google_compute_router_peer` ([#5531](https://github.com/hashicorp/terraform-provider-google-beta/pull/5531))
* compute: made `network_firewall_policy_enforcement_order` field mutable in `google_compute_network`. ([#5516](https://github.com/hashicorp/terraform-provider-google-beta/pull/5516))
* dlp: added `exclude_by_hotword` exclusion rule to `google_data_loss_prevention_inspect_template` resource ([#5555](https://github.com/hashicorp/terraform-provider-google-beta/pull/5555))
* dlp: added `image_transformations` field to `google_data_loss_prevention_deidentify_template` resource ([#5556](https://github.com/hashicorp/terraform-provider-google-beta/pull/5556))
* dlp: added `inspectConfig` field to `google_data_loss_prevention_job_trigger` resource ([#5535](https://github.com/hashicorp/terraform-provider-google-beta/pull/5535))
* dlp: added `replace_dictionary_config` field to `info_type_transformations` in `google_data_loss_prevention_deidentify_template` resource ([#5556](https://github.com/hashicorp/terraform-provider-google-beta/pull/5556))
* dlp: added `surrogate_type` custom type to `google_data_loss_prevention_inspect_template` resource ([#5555](https://github.com/hashicorp/terraform-provider-google-beta/pull/5555))
* dlp: added `version` field for multiple `info_type` blocks to `google_data_loss_prevention_inspect_template` resource ([#5555](https://github.com/hashicorp/terraform-provider-google-beta/pull/5555))
* sql: Added support for Postgres in `google_sql_source_representation_instance` ([#5557](https://github.com/hashicorp/terraform-provider-google-beta/pull/5557))
* vertexai: added `region` field to `google_vertex_ai_endpoint` ([#5514](https://github.com/hashicorp/terraform-provider-google-beta/pull/5514))
* workflows: added `crypto_key_name` field to `google_workflows_workflow` resource ([#5509](https://github.com/hashicorp/terraform-provider-google-beta/pull/5509))
* workstations: supported in-place update for `host` and `container` in `google_workstations_workstation_config` ([#5585](https://github.com/hashicorp/terraform-provider-google-beta/pull/5585))

BUG FIXES:
* cloudplatform: added validation for `role_id` on `google_organization_iam_custom_role` ([#5569](https://github.com/hashicorp/terraform-provider-google-beta/pull/5569))
* compute: fixed an import bug for `google_compute_router_interface` that happened when project was not set in the provider configuration or via environment variable ([#5508](https://github.com/hashicorp/terraform-provider-google-beta/pull/5508))
* dns: fixed bug in `google_dns_keys` data source where list attributes could not be used at plan-time ([#5546](https://github.com/hashicorp/terraform-provider-google-beta/pull/5546))
* firebase: specified required argument `bundle_id` in `google_firebase_apple_app` ([#5577](https://github.com/hashicorp/terraform-provider-google-beta/pull/5577))
* workstations: fixed an issue where modifying `persistent_directories` and `encryption_key` would fail with API errors; now updating them will recreate the resource ([#5585](https://github.com/hashicorp/terraform-provider-google-beta/pull/5585))
* workstations: fixed an issue where unsetting the container working directory in `google_workstations_workstations_config` was not propagated to the underlying resource ([#5585](https://github.com/hashicorp/terraform-provider-google-beta/pull/5585))

## 4.63.1 (April 26, 2023)

BUG FIXES:
* bigtable: fixed plan failure because of an unused zone being unavailable

## 4.63.0 (April 24, 2023)

NOTES:
* alloydb: changed `location` from `optional` to `required` for `google_alloydb_cluster` and `google_alloydb_backup` resources. `location` had previously been marked as optional, but operations failed if it was omitted, and there was no way for `location` to be inherited from the provider configuration or from an environment variable. This means there was no way to have a working configuration without `location` specified. ([#5492](https://github.com/hashicorp/terraform-provider-google-beta/pull/5492), [#5494](https://github.com/hashicorp/terraform-provider-google-beta/pull/5494))
* workflows: updated api version from v1beta1 to v1 (beta) ([#5482](https://github.com/hashicorp/terraform-provider-google-beta/pull/5482))

FEATURES:
* **New Resource:** `google_access_context_manager_ingress_policy` ([#5474](https://github.com/hashicorp/terraform-provider-google-beta/pull/5474))
* **New Resource:** `google_compute_public_advertised_prefix` ([#5476](https://github.com/hashicorp/terraform-provider-google-beta/pull/5476))
* **New Resource:** `google_compute_public_delegated_prefix` ([#5476](https://github.com/hashicorp/terraform-provider-google-beta/pull/5476))
* **New Resource:** `google_compute_region_commitment` ([#5473](https://github.com/hashicorp/terraform-provider-google-beta/pull/5473))
* **New Resource:** `google_network_services_http_route` ([#5471](https://github.com/hashicorp/terraform-provider-google-beta/pull/5471))
* **New Resource:** `google_network_services_tcp_route` (beta) ([#5497](https://github.com/hashicorp/terraform-provider-google-beta/pull/5497))

IMPROVEMENTS:
* dlp: added `inspect_job.actions.job_notification_emails` and `inspect_job.actions.deidentify`  fields to `google_data_loss_prevention_job_trigger` resource ([#5477](https://github.com/hashicorp/terraform-provider-google-beta/pull/5477))
* dlp: added `triggers.manual` and `inspect_job.storage_config.hybrid_options` to `google_data_loss_prevention_job_trigger` ([#5490](https://github.com/hashicorp/terraform-provider-google-beta/pull/5490))
* iam: added `oidc.web_sso_config` field to `google_iam_workforce_pool_provider` ([#5491](https://github.com/hashicorp/terraform-provider-google-beta/pull/5491))

BUG FIXES:
* alloydb: changed `weekly_schedule` (under `automated_backup_policy`) from required to optional for `google_alloydb_cluster` ([#5495](https://github.com/hashicorp/terraform-provider-google-beta/pull/5495))
* compute: fixed an issue with TTLs being sent when `USE_ORIGIN_HEADERS` is set in `google_compute_backend_bucket` ([#5488](https://github.com/hashicorp/terraform-provider-google-beta/pull/5488))
* networkservices: increased default timeouts for `google_network_services_edge_cache_keyset` to 60m (from 30m) ([#5481](https://github.com/hashicorp/terraform-provider-google-beta/pull/5481))
* sql: fixed an issue that prevented setting `enable_private_path_for_google_cloud_services` to `false` in `google_sql_database_instance`([#5484](https://github.com/hashicorp/terraform-provider-google-beta/pull/5484))

## 4.62.1 (April 19, 2023)

BUG FIXES:
* compute: fixed a diff that occurred when `stack_type` was unset on `google_compute_ha_vpn_gateway` ([#5479](https://github.com/hashicorp/terraform-provider-google-beta/pull/5479))

## 4.62.0 (April 17, 2023)

FEATURES:
* **New Data Source:** `google_compute_region_instance_template` ([#5467](https://github.com/hashicorp/terraform-provider-google-beta/pull/5467))
* **New Resource:** `google_compute_region_instance_template` ([#5467](https://github.com/hashicorp/terraform-provider-google-beta/pull/5467))
* **New Resource:** `google_logging_linked_dataset` ([#5459](https://github.com/hashicorp/terraform-provider-google-beta/pull/5459))

IMPROVEMENTS:
* cloudasset: added `OS_INVENTORY` value to `content_type` for `google_cloud_asset_*_feed` ([#5465](https://github.com/hashicorp/terraform-provider-google-beta/pull/5465))
* clouddeploy: added canary deployment fields for resource `google_clouddeploy_delivery_pipeline` ([#5451](https://github.com/hashicorp/terraform-provider-google-beta/pull/5451))
* compute: supported region instance template in`source_instance_template` field of `google_compute_instance_from_template` resource ([#5467](https://github.com/hashicorp/terraform-provider-google-beta/pull/5467))
* container: added `pod_cidr_overprovision_config` field to `google_container_cluster` and  `google_container_node_pool` resources. ([#5468](https://github.com/hashicorp/terraform-provider-google-beta/pull/5468))
* orgpolicy: accepted variable cases for booleans such as true, True, and TRUE in `google_org_policy_policy` ([#5443](https://github.com/hashicorp/terraform-provider-google-beta/pull/5443))

BUG FIXES:
* cloudidentity: fixed immutability issue on `initialGroupConfig` field for resource `google_cloud_identity_group` ([#5456](https://github.com/hashicorp/terraform-provider-google-beta/pull/5456))
* provider: fixed an error resulting from leaving `batching.send_after` unspecified and `batching` specified ([#5460](https://github.com/hashicorp/terraform-provider-google-beta/pull/5460))
* provider: fixed bug where `credentials` field could not be set as an empty string ([#5466](https://github.com/hashicorp/terraform-provider-google-beta/pull/5466))
* vertex: increased the default timeout for `google_vertex_ai_index` to 180m ([#5450](https://github.com/hashicorp/terraform-provider-google-beta/pull/5450))

## 4.61.0 (April 10, 2023)

BREAKING CHANGES:
* cloudrunv2: set a default value of 3 for `max_retries` in `google_cloud_run_v2_job`. This should match the API's existing default, but may show a diff at plan time in limited circumstances as drift is now detected ([#5432](https://github.com/hashicorp/terraform-provider-google-beta/pull/5432))

FEATURES:
* **New Data Source:** `google_firebase_android_app_config` ([#5425](https://github.com/hashicorp/terraform-provider-google-beta/pull/5425))
* **New Resource:** `google_apigee_keystores_aliases_pkcs12` ([#5411](https://github.com/hashicorp/terraform-provider-google-beta/pull/5411))
* **New Resource:** `google_apigee_keystores_aliases_self_signed_cert` ([#5394](https://github.com/hashicorp/terraform-provider-google-beta/pull/5394))
* **New Resource:** `google_network_security_url_lists` ([#5439](https://github.com/hashicorp/terraform-provider-google-beta/pull/5439))
* **New Resource:** `google_network_services_mesh` ([#5393](https://github.com/hashicorp/terraform-provider-google-beta/pull/5393))
* **New Resource:** `google_network_security_gateway_security_policy` (beta) ([#5434](https://github.com/hashicorp/terraform-provider-google-beta/pull/5434))
* **New Resource:** `google_network_security_gateway_security_policy_rule` (beta) ([#5434](https://github.com/hashicorp/terraform-provider-google-beta/pull/5434))

IMPROVEMENTS:
* alloydb: added update support for `initial_user` and `automated_backup_policy.weekly_schedule` to `google_alloydb_cluster` ([#5420](https://github.com/hashicorp/terraform-provider-google-beta/pull/5420))
* artifactregistry: added support for tag immutability ([#5427](https://github.com/hashicorp/terraform-provider-google-beta/pull/5427))
* artifactregistry: promoted `mode`, `virtual_repository_config`, and `remote_repository_config` to GA ([#5426](https://github.com/hashicorp/terraform-provider-google-beta/pull/5426))
* bigqueryreservation: added `edition` and `autoscale` to `google_bigquery_reservation` and `edition` to `bigquery_capacity_commitment` ([#5399](https://github.com/hashicorp/terraform-provider-google-beta/pull/5399))
* compute: added support for `SEV_LIVE_MIGRATABLE` to `guest_os_features.type` in `google_compute_image` ([#5424](https://github.com/hashicorp/terraform-provider-google-beta/pull/5424))
* compute: added support for `stack_type` to `google_compute_ha_vpn_gateway` ([#5395](https://github.com/hashicorp/terraform-provider-google-beta/pull/5395))
* container: added support for `ephemeral_storage_local_ssd_config` to `google_container_cluster.node_config`, `google_container_cluster.node_pools.node_config`, `google_container_node_pool.node_config` ([#5400](https://github.com/hashicorp/terraform-provider-google-beta/pull/5400))
* dlp: Changed `dictionary`, `regex`, `regex.group_indexes` and `large_custom_dictionary` fields in `google_data_loss_prevention_stored_info_type` to be update-in-place ([#5428](https://github.com/hashicorp/terraform-provider-google-beta/pull/5428))
* logging: added support for `disabled` to `google_logging_metric` ([#5423](https://github.com/hashicorp/terraform-provider-google-beta/pull/5423))
* networkservices: increased the max count for `route_rule` to 200 on `google_network_services_edge_cache_service` ([#5433](https://github.com/hashicorp/terraform-provider-google-beta/pull/5433))
* storagetransfer: added support for 'last_modified_since' and 'last_modified_before' fields to 'google_storage_transfer_job' resource ([#5398](https://github.com/hashicorp/terraform-provider-google-beta/pull/5398))

BUG FIXES:
* bigquery: fixed the import logic in `google_bigquery_capacity_commitment` ([#5435](https://github.com/hashicorp/terraform-provider-google-beta/pull/5435))
* cloudrunv2: fixed the bug where setting `max_retries` to 0 in `google_cloud_run_v2_job` was not respected. ([#5432](https://github.com/hashicorp/terraform-provider-google-beta/pull/5432))
* container: fixed a bug creating a diff adding a `stack_type` when GKE omitted `stackType` in API responses from older GKE clusters ([#5429](https://github.com/hashicorp/terraform-provider-google-beta/pull/5429))
* dataproc: fixed validation of `optional_components` ([#5410](https://github.com/hashicorp/terraform-provider-google-beta/pull/5410))
* provider: fixed an issue where the `USER_PROJECT_OVERRIDE` environment variable was not being read ([#5441](https://github.com/hashicorp/terraform-provider-google-beta/pull/5441))
* provider: fixed an issue where the provider crashed when "batching" was set in `4.60.0`/`4.60.1` ([#5440](https://github.com/hashicorp/terraform-provider-google-beta/pull/5440))

## 4.60.2 (April 6, 2023)

BUG FIXES:
* provider: fixed an issue where the provider crashed when "batching" was set in `4.60.0`/`4.60.1`
* provider: fixed an issue where the `USER_PROJECT_OVERRIDE` environment variable was not being read

## 4.60.1 (April 5, 2023)

BUG FIXES:
* container: fixed a bug creating a diff adding a `stack_type` when GKE omitted `stackType` in API responses from older GKE clusters

## 4.60.0 (April 4, 2023)

FEATURES:
* **New Resource:** `google_apigee_keystores_aliases_key_cert_file` ([#5386](https://github.com/hashicorp/terraform-provider-google-beta/pull/5386))

IMPROVEMENTS:
* compute: added `address_type`, `network`, `network_tier`, `prefix_length`, `purpose`, `subnetwork` and `users` field for `google_compute_address` and `google_compute_global_address` datasource ([#5363](https://github.com/hashicorp/terraform-provider-google-beta/pull/5363))
* compute: added `network_firewall_policy_enforcement_order` field to `google_compute_network` resource ([#5375](https://github.com/hashicorp/terraform-provider-google-beta/pull/5375))
* compute: added output-only attribute `self_link_unique` for `google_compute_instance_template` to point to the unique id of the resource instead of its name ([#5384](https://github.com/hashicorp/terraform-provider-google-beta/pull/5384))
* container: added `stack_type` field to `google_container_cluster` resource ([#5364](https://github.com/hashicorp/terraform-provider-google-beta/pull/5364))
* container: added `advanced_machine_features` field to `google_container_cluster` resource ([#5371](https://github.com/hashicorp/terraform-provider-google-beta/pull/5371))
* networkservice: updated the max number of `host_rule` on `google_network_services_edge_cache_service` ([#5376](https://github.com/hashicorp/terraform-provider-google-beta/pull/5376))
* sql: added support of single-database-recovery for SQL Server PITR with `database_names` attribute to `google_sql_instance` ([#5366](https://github.com/hashicorp/terraform-provider-google-beta/pull/5366))

BUG FIXES:
* cloudrun: fixed race condition when polling for status during an update of a `google_cloud_run_service` ([#5365](https://github.com/hashicorp/terraform-provider-google-beta/pull/5365))
* cloudsql: fixed the error in any subsequent apply on `google_sql_user` after its `google_sql_database_instance` is deleted ([#5369](https://github.com/hashicorp/terraform-provider-google-beta/pull/5369))
* datacatalog: fixed `google_data_catalog_tag` only allowing 10 tags by increasing the page size to 1000 ([#5362](https://github.com/hashicorp/terraform-provider-google-beta/pull/5362))
* firebase: fixed `google_firebase_project` to succeed on apply when the project already has firebase enabled ([#5379](https://github.com/hashicorp/terraform-provider-google-beta/pull/5379))


## 4.59.0 (March 28, 2023)

FEATURES:
* **New Resource:** `google_dataplex_asset_iam_*` ([#5348](https://github.com/hashicorp/terraform-provider-google-beta/pull/5348))
* **New Resource:** `google_dataplex_lake_iam_*` ([#5348](https://github.com/hashicorp/terraform-provider-google-beta/pull/5348))
* **New Resource:** `google_dataplex_zone_iam_*` ([#5348](https://github.com/hashicorp/terraform-provider-google-beta/pull/5348))
* **New Resource:** `google_network_services_gateway` ([#5355](https://github.com/hashicorp/terraform-provider-google-beta/pull/5355))

IMPROVEMENTS:
* auth: added support for oauth2 token exchange over mTLS ([#5343](https://github.com/hashicorp/terraform-provider-google-beta/pull/5343))
* bigquery: added `is_case_insensitive` and `default_collation` fields to `google_bigquery_dataset` resource ([#5342](https://github.com/hashicorp/terraform-provider-google-beta/pull/5342))
* compute: added `scratch_disk.size` field on `google_compute_instance` ([#5358](https://github.com/hashicorp/terraform-provider-google-beta/pull/5358))
* compute: added 3000 as allowable value for `disk_size_gb` for SCRATCH disks in `google_compute_instance_template` ([#5358](https://github.com/hashicorp/terraform-provider-google-beta/pull/5358))
* compute: added `WEIGHED_MAGLEV` to `locality_lb_policy` enum for backend service resources ([#5353](https://github.com/hashicorp/terraform-provider-google-beta/pull/5353))
* container: added `local_nvme_ssd_block` to `node_config` block in the `google_container_node_pool` ([#5335](https://github.com/hashicorp/terraform-provider-google-beta/pull/5335))
* logging: added `enable_analytics` field to `google_logging_project_bucket_config` ([#5347](https://github.com/hashicorp/terraform-provider-google-beta/pull/5347))
* networkservices: updated max allowed items to 25 for `expose_headers`, `allow_headers`, `request_header_to_remove`, `request_header_to_add`, `response_header_to_add` and `response_header_to_remove` of `google_network_services_edge_cache_service` ([#5346](https://github.com/hashicorp/terraform-provider-google-beta/pull/5346))
* networkservices: updated max allowed items to 25 for `request_headers_to_add` of `google_network_services_edge_cache_origin` ([#5346](https://github.com/hashicorp/terraform-provider-google-beta/pull/5346))

BUG FIXES:
* certificatemanager: fixed `managed.dns_authorizations` not being included during import of `google_certificate_manager_certificate` ([#5325](https://github.com/hashicorp/terraform-provider-google-beta/pull/5325))
* certificatemanager: fixed a bug where modifying non-updatable fields `hostname` and `matcher` in `google_certificate_manager_certificate_map_entry` would fail with API errors; now updating them will recreate the resource ([#5327](https://github.com/hashicorp/terraform-provider-google-beta/pull/5327))
* compute: fixed bug where `enforce_on_key_name` could not be unset on `google_compute_security_policy` ([#5326](https://github.com/hashicorp/terraform-provider-google-beta/pull/5326))
* datastream: fixed bug where field `dataset_id` could not utilize the id from bigquery directly ([#5331](https://github.com/hashicorp/terraform-provider-google-beta/pull/5331))
* workstations: fixed permadiff on `service_account` of `google_workstations_workstation_config` ([#5323](https://github.com/hashicorp/terraform-provider-google-beta/pull/5323))

## 4.58.0 (March 21, 2023)

FEATURES:
* **New Resource:** `google_apigee_sharedflow` ([#5300](https://github.com/hashicorp/terraform-provider-google-beta/pull/5300))
* **New Resource:** `google_apigee_sharedflow_deployment` ([#5300](https://github.com/hashicorp/terraform-provider-google-beta/pull/5300))
* **New Resource:** `google_apigee_flowhook` ([#5300](https://github.com/hashicorp/terraform-provider-google-beta/pull/5300))

IMPROVEMENTS:
* datafusion: added support for `accelerators` field to `google_datafusion_instance` resource. ([#5304](https://github.com/hashicorp/terraform-provider-google-beta/pull/5304))
* privateca: added support for X.509 name constraints to `google_privateca_pool`, `google_privateca_certificate`, and `google_privateca_certificate_authority` ([#5317](https://github.com/hashicorp/terraform-provider-google-beta/pull/5317))

BUG FIXES:
* alloydb: fixed permadiff on `automated_backup_policy.weekly_schedule` of `google_alloydb_cluster` ([#5305](https://github.com/hashicorp/terraform-provider-google-beta/pull/5305))
* bigquery: fixed a permadiff when `friendly_name` is removed from `google_bigquery_dataset` ([#5319](https://github.com/hashicorp/terraform-provider-google-beta/pull/5319))
* redis: fixed a bug causing diff detection on `reserved_ip_range` in `google_redis_instance` ([#5310](https://github.com/hashicorp/terraform-provider-google-beta/pull/5310))

## 4.57.0 (March 13, 2023)

FEATURES:
* **New Resource:** `google_access_context_manager_authorized_orgs_desc` ([#5292](https://github.com/hashicorp/terraform-provider-google-beta/pull/5292))
* **New Resource:** `google_bigquery_capacity_commitment` ([#5282](https://github.com/hashicorp/terraform-provider-google-beta/pull/5282))
* **New Resource:** `google_workstations_workstation` ([#5273](https://github.com/hashicorp/terraform-provider-google-beta/pull/5273))
* **New Resource:** `google_apigee_env_keystore` ([#5267](https://github.com/hashicorp/terraform-provider-google-beta/pull/5267))
* **New Resource:** `google_apigee_env_references` ([#5267](https://github.com/hashicorp/terraform-provider-google-beta/pull/5267))

IMPROVEMENTS:
* artifactregistry: added field `virtual_repository_config` and `remote_repository_config` to `google_artifact_registry_repository` ([#5289](https://github.com/hashicorp/terraform-provider-google-beta/pull/5289))
* bigquerydatapolicy: updated api version from v1beta1 to v1 and made it possible to use additional data policies. ([#5291](https://github.com/hashicorp/terraform-provider-google-beta/pull/5291))
* compute: added `maintenance_interval` field to `google_compute_instance_template` and `google_compute_instance` ([#5297](https://github.com/hashicorp/terraform-provider-google-beta/pull/5297))

BUG FIXES:
* cloudidentity: fixed an issue on `google_cloud_identity_group` `initial_group_config` field when importing ([#5266](https://github.com/hashicorp/terraform-provider-google-beta/pull/5266))
* compute: fixed the error of invalid value for field `failover_policy` when UDP is selected on `google_compute_region_backend_service` ([#5280](https://github.com/hashicorp/terraform-provider-google-beta/pull/5280))
* firebase: allowed specifying a `project` field on datasources for `google_firebase_android_app`, `google_firebase_web_app`, and `google_firebase_apple_app`. ([#5293](https://github.com/hashicorp/terraform-provider-google-beta/pull/5293))
* tags: fixed a bug preventing use of `google_tags_location_tag_binding` with zonal parent resources ([#5269](https://github.com/hashicorp/terraform-provider-google-beta/pull/5269))

## 4.56.0 (March 6, 2023)

DEPRECATIONS:
* gkehub: deprecated `mesh.control_plane` in `google_gke_hub_feature_membership`. Use `mesh.management` instead ([#5258](https://github.com/hashicorp/terraform-provider-google-beta/pull/5258))

FEATURES:
* **New Resource:** google_scc_mute_config ([#5241](https://github.com/hashicorp/terraform-provider-google-beta/pull/5241))
* **New Resource:** google_workstations_workstation_config ([#5246](https://github.com/hashicorp/terraform-provider-google-beta/pull/5246))

IMPROVEMENTS:
* cloudbuild: added `peered_network_ip_range` field to `google_cloudbuild_worker_pool` resource ([#5258](https://github.com/hashicorp/terraform-provider-google-beta/pull/5258))
* clouddeploy: added `multi_target` field to `google_clouddeploy_target` resource ([#5258](https://github.com/hashicorp/terraform-provider-google-beta/pull/5258))
* cloudrun: added `template.0.containers0.liveness_probe.grpc`, `template.0.containers0.startup_probe.grpc` fields to `google_cloud_run_v2_service` resource ([#5259](https://github.com/hashicorp/terraform-provider-google-beta/pull/5259))
* compute: added `internal_ip` and `external_ip` to `google_compute_per_instance_config` and `google_compute_region_per_instance_config` (beta) ([#5256](https://github.com/hashicorp/terraform-provider-google-beta/pull/5256))
* compute: added `max_distance` field to `resource-policy` resource ([#5257](https://github.com/hashicorp/terraform-provider-google-beta/pull/5257))
* compute: added field `deletion_policy` to resource `google_compute_shared_vpc_service_project` ([#5243](https://github.com/hashicorp/terraform-provider-google-beta/pull/5243))
* container: added field `protect_config` to `google_container_cluster` (beta) ([#5255](https://github.com/hashicorp/terraform-provider-google-beta/pull/5255))
* containerazure: added `azure_services_authentication` to `google_container_azure_cluster` ([#5258](https://github.com/hashicorp/terraform-provider-google-beta/pull/5258))
* networkservices: increased maximum `allow_origins` from 5 to 25 on `network_services_edge_cache_service` ([#5239](https://github.com/hashicorp/terraform-provider-google-beta/pull/5239))
* storagetransfer: added general field `sink_agent_pool_name` and `source_agent_pool_name` to `google_storage_transfer_job` ([#5262](https://github.com/hashicorp/terraform-provider-google-beta/pull/5262))

BUG FIXES:
* artifactregistry: fixed an issue where `google-beta` used an outdated beta API rather than the GA service API. New format values like "KFP" will now be accepted by both providers. ([#5247](https://github.com/hashicorp/terraform-provider-google-beta/pull/5247))
* cloudfunctions: fixed no diff found on `event_trigger.resource` of `google_cloudfunctions_function` ([#5261](https://github.com/hashicorp/terraform-provider-google-beta/pull/5261))
* dataproc: fixed an issue where `master_config.num_instances` would not force recreation when changed in `google_dataproc_cluster` ([#5251](https://github.com/hashicorp/terraform-provider-google-beta/pull/5251))
* spanner: fixed the error when updating `deletion_protection` on `google_spanner_database` ([#5242](https://github.com/hashicorp/terraform-provider-google-beta/pull/5242))
* spanner: fixed the error when updating `force_destroy` on `google_spanner_instance` ([#5242](https://github.com/hashicorp/terraform-provider-google-beta/pull/5242))

## 4.55.0 (February 27, 2023)

FEATURES:
* **New Resource:** `google_cloudbuild_bitbucket_server_config` ([#5218](https://github.com/hashicorp/terraform-provider-google-beta/pull/5218))
* **New Resource:** `google_firebase_hosting_release` ([#5229](https://github.com/hashicorp/terraform-provider-google-beta/pull/5229))
* **New Resource:** `google_firebase_hosting_version` ([#5229](https://github.com/hashicorp/terraform-provider-google-beta/pull/5229))

IMPROVEMENTS:
* container: added support for `node_config.kubelet_config.pod_pids_limit` on `google_container_node_pool` ([#5217](https://github.com/hashicorp/terraform-provider-google-beta/pull/5217))
* storage: changed the default create timeout of `google_storage_bucket` to 10m from 4m ([#5222](https://github.com/hashicorp/terraform-provider-google-beta/pull/5222))

BUG FIXES:
* container: fixed a crash when leaving `placement_policy` blank on `google_container_node_pool` ([#5233](https://github.com/hashicorp/terraform-provider-google-beta/pull/5233))
* serviceusage: removed unneeded Optional schema behaviour from `email` field in `google_project_service_identity` ([#5226](https://github.com/hashicorp/terraform-provider-google-beta/pull/5226))

## 4.54.0 (February 23, 2023)

FEATURES:
* **New Data Source:** `google_firebase_hosting_channel` ([#5188](https://github.com/hashicorp/terraform-provider-google-beta/pull/5188))
* **New Data Source:** `google_logging_sink` ([#5207](https://github.com/hashicorp/terraform-provider-google-beta/pull/5207))
* **New Data Source:** `google_sql_databases` ([#5204](https://github.com/hashicorp/terraform-provider-google-beta/pull/5204))

IMPROVEMENTS:
* cloudbuild: added `bitbucket_server_trigger_config` field to `google_cloudbuild_trigger` resource ([#5198](https://github.com/hashicorp/terraform-provider-google-beta/pull/5198))
* cloudbuild: added `github.enterprise_config_resource_name` field to `google_cloudbuild_trigger` resource ([#5205](https://github.com/hashicorp/terraform-provider-google-beta/pull/5205))
* compute: added field `rsa_encrypted_key` to `google_compute_disk` resource ([#5187](https://github.com/hashicorp/terraform-provider-google-beta/pull/5187))
* compute: added additional fields to `rules[].rate_limit_options` in `google_compute_security_policy` to support Cloud Armor Rate Limit Options ([#5193](https://github.com/hashicorp/terraform-provider-google-beta/pull/5193))
* sql: added replica promotion support to `google_sql_database_instance`. This change will allow users to promote read replica as stand alone primary instance. ([#5184](https://github.com/hashicorp/terraform-provider-google-beta/pull/5184))

BUG FIXES:
* bigquery: fixed permadiff on `max_time_travel_hours` of `google_bigquery_dataset` ([#5190](https://github.com/hashicorp/terraform-provider-google-beta/pull/5190))
* compute: added possibility to remove `stateful_disk` in `compute_instance_group_manager` and `compute_region_instance_group_manager`. ([#5203](https://github.com/hashicorp/terraform-provider-google-beta/pull/5203))
* sql: fixed an issue with updating the `settings.activation_policy` field in `google_sql_database_instance` ([#5202](https://github.com/hashicorp/terraform-provider-google-beta/pull/5202))

## 4.53.1 (February 14, 2023)

BUG FIXES:
* provider: fixed crash when trying to configure the provider with invalid credentials

## 4.53.0 (February 13, 2023)

FEATURES:
* **New Resource:** `google_apigee_addons_config` ([#5171](https://github.com/hashicorp/terraform-provider-google-beta/pull/5171))
* **New Resource:** `google_cloudbuildv2_connection_iam_binding` ([#5158](https://github.com/hashicorp/terraform-provider-google-beta/pull/5158))
* **New Resource:** `google_cloudbuildv2_connection_iam_member` ([#5158](https://github.com/hashicorp/terraform-provider-google-beta/pull/5158))
* **New Resource:** `google_cloudbuildv2_connection_iam_policy` ([#5158](https://github.com/hashicorp/terraform-provider-google-beta/pull/5158))
* **New Resource:** `google_firestore_database` ([#5181](https://github.com/hashicorp/terraform-provider-google-beta/pull/5181))
* **New Resource:** `google_workstations_workstation_cluster` ([#5154](https://github.com/hashicorp/terraform-provider-google-beta/pull/5154))

IMPROVEMENTS:
* compute: added `resource_policies` field to `google_compute_instance_template` ([#5182](https://github.com/hashicorp/terraform-provider-google-beta/pull/5182))
* compute: added field `force_update_on_repair` to `instance_lifecycle_policy` ([#5172](https://github.com/hashicorp/terraform-provider-google-beta/pull/5172))
* compute: added field `instance_lifecycle_policy` to `google_compute_instance_group_manager` and `google_compute_region_instance_group_manager` ([#5172](https://github.com/hashicorp/terraform-provider-google-beta/pull/5172))
* compute: added the `labels` field to the `google_compute_external_vpn_gateway` resource ([#5162](https://github.com/hashicorp/terraform-provider-google-beta/pull/5162))
* datastream: added `postgresql_source_config` & `oracle_source_config` in `google_datastream_stream` ([#5166](https://github.com/hashicorp/terraform-provider-google-beta/pull/5166))
* datastream: added support for creating `google_datastream_stream` with `desired_state=RUNNING` ([#5166](https://github.com/hashicorp/terraform-provider-google-beta/pull/5166))
* datastream: exposed validation errors during `google_datastream_stream` creation ([#5166](https://github.com/hashicorp/terraform-provider-google-beta/pull/5166))
* firebase: marked `deletion_policy` as updatable without recreation on `google_firebase_android_app` and `google_firebase_apple_app` ([#5163](https://github.com/hashicorp/terraform-provider-google-beta/pull/5163))
* sql: added `enable_private_path_for_google_cloud_services` field to `google_sql_database_instance` resource ([#5177](https://github.com/hashicorp/terraform-provider-google-beta/pull/5177))
* vertex_ai: added `offline_storage_ttl_days` to `google_vertex_ai_featurestore_entitytype` resource ([#5178](https://github.com/hashicorp/terraform-provider-google-beta/pull/5178))
* vertex_ai: added `online_storage_ttl_days` to `google_vertex_ai_featurestore` resource ([#5178](https://github.com/hashicorp/terraform-provider-google-beta/pull/5178))
* vertex_ai: added the field `description` to `google_vertex_ai_featurestore_entitytype` ([#5161](https://github.com/hashicorp/terraform-provider-google-beta/pull/5161))

BUG FIXES:
* composer: fixed an issue with cleaning up environments created in an error state ([#5164](https://github.com/hashicorp/terraform-provider-google-beta/pull/5164))
* compute: fixed wrong maximum limit description for possible VPC MTUs ([#5180](https://github.com/hashicorp/terraform-provider-google-beta/pull/5180))
* datafusion: fixed `version` can't be updated on `google_data_fusion_instance` ([#5175](https://github.com/hashicorp/terraform-provider-google-beta/pull/5175))

## 4.52.0 (February 6, 2023)

FEATURES:
* **New Data Source:** `google_secret_manager_secret_version_access` ([#5147](https://github.com/hashicorp/terraform-provider-google-beta/pull/5147))
* **New Resource:** `google_cloudbuildv2_connection` ([#5140](https://github.com/hashicorp/terraform-provider-google-beta/pull/5140))
* **New Resource:** `google_cloudbuildv2_repository` ([#5140](https://github.com/hashicorp/terraform-provider-google-beta/pull/5140))
* **New Resource:** `google_workstations_workstation_cluster` ([#5154](https://github.com/hashicorp/terraform-provider-google-beta/pull/5154))

IMPROVEMENTS:
* bigquery: added support for federated Azure identities to BigQuery Omni connections. ([#5150](https://github.com/hashicorp/terraform-provider-google-beta/pull/5150))
* bigquery: added `cloud_spanner.use_serverless_analytics` field ([#5139](https://github.com/hashicorp/terraform-provider-google-beta/pull/5139))
* bigquery: added `cloud_sql.service_account_id` and `azure.identity` output fields ([#5139](https://github.com/hashicorp/terraform-provider-google-beta/pull/5139))
* cloudbuild: added field `repository_event_config` to resource `trigger` ([#5142](https://github.com/hashicorp/terraform-provider-google-beta/pull/5142))
* compute: added `locality_lb_policies` field to `google_compute_backend_service` ([#5146](https://github.com/hashicorp/terraform-provider-google-beta/pull/5146))
* sql: updated the `settings.deletion_protection_enabled` property documentation. ([#13581](https://github.com/hashicorp/terraform-provider-google-beta/pull/13581))
* sql: made `root_password` field updatable in `google_sql_database_instance` ([#5133](https://github.com/hashicorp/terraform-provider-google-beta/pull/5133))

BUG FIXES:
* cloudfunctions: updated max_instances field to take API's result as default value ([#13575](https://github.com/hashicorp/terraform-provider-google-beta/pull/13575))
* container: fixed an issue with resuming failed cluster creation ([#5136](https://github.com/hashicorp/terraform-provider-google-beta/pull/5136))
* datacatalog: fixed the import failure when the `project` is different from the default on `google_data_catalog_taxonomy` ([#5145](https://github.com/hashicorp/terraform-provider-google-beta/pull/5145))
* secretmanager: fixed incorrect required_with for topics in `google_secret_managed_secret` ([#5149](https://github.com/hashicorp/terraform-provider-google-beta/pull/5149))

## 4.51.0 (January 30, 2023)

DEPRECATIONS:
* cloudrunv2: deprecated `liveness_probe.tcp_socket` field from `google_cloud_run_v2_service` resource as it is not supported by the API and it will be removed in a future major release ([#5128](https://github.com/hashicorp/terraform-provider-google-beta/pull/5128))
* cloudrunv2: deprecated `startup_probe` and `liveness_probe` fields from `google_cloud_run_v2_job` resource as they are not supported by the API and they will be removed in a future major release ([#5118](https://github.com/hashicorp/terraform-provider-google-beta/pull/5118))

FEATURES:
* **New Resource:** `google_iam_access_boundary_policy` ([#5130](https://github.com/hashicorp/terraform-provider-google-beta/pull/5130))
* **New Resource:** `google_tags_location_tag_bindings` ([#5115](https://github.com/hashicorp/terraform-provider-google-beta/pull/5115))

IMPROVEMENTS:
* cloudbuild: added `github_enterprise_config` fields to `google_cloudbuild_trigger` resource. ([#5110](https://github.com/hashicorp/terraform-provider-google-beta/pull/5110))
* cloudrunV2: added `annotations` to `google_cloud_run_v2_service` resource ([#5108](https://github.com/hashicorp/terraform-provider-google-beta/pull/5108))
* composer: Added field `cloud_data_lineage_integration` to resource `google_composer_environment` (beta) ([#5109](https://github.com/hashicorp/terraform-provider-google-beta/pull/5109))
* compute: added `tcp_time_wait_timeout_sec` field to `google_compute_router_nat` resource ([#5123](https://github.com/hashicorp/terraform-provider-google-beta/pull/5123))
* compute: Added fields to resource `google_compute_security_policy` to support Cloud Armor Auto Deploy (beta) ([#5116](https://github.com/hashicorp/terraform-provider-google-beta/pull/5116))
* compute: added `share_settings` field to the `google_compute_node_group` resource. ([#5113](https://github.com/hashicorp/terraform-provider-google-beta/pull/5113))
* containerattached: added `deletion_policy` field to `google_container_attached_cluster` resource. ([#5121](https://github.com/hashicorp/terraform-provider-google-beta/pull/5121))
* datastream: added `customer_managed_encryption_key` and `destination_config.bigquery_destination_config.source_hierarchy_datasets.dataset_template.kms_key_name` fields to `google_datastream_stream` resource ([#5120](https://github.com/hashicorp/terraform-provider-google-beta/pull/5120))
* dlp: added `publish_findings_to_cloud_data_catalog` and `publish_summary_to_cscc` to `google_data_loss_prevention_job_trigger` resource ([#5127](https://github.com/hashicorp/terraform-provider-google-beta/pull/5127))
* sql: added point_in_time_recovery_enabled for SQLServer in `google_sql_database_instance` ([#5124](https://github.com/hashicorp/terraform-provider-google-beta/pull/5124))
* spanner: added support for IAM conditions with `google_spanner_database_iam_member` and `google_spanner_instance_iam_member` ([#5125](https://github.com/hashicorp/terraform-provider-google-beta/pull/5125))
* sql: added additional fields to `google_sql_source_representation_instance`([#5114](https://github.com/hashicorp/terraform-provider-google-beta/pull/5114))

BUG FIXES:
* bigquery: fixed bug where valid iam member values for bigquery were prevented from actuation by validation ([#5111](https://github.com/hashicorp/terraform-provider-google-beta/pull/5111))
* bigquery: fixed permadiff on `external_data_configuration.connection_id` of `google_bigquery_table` ([#5126](https://github.com/hashicorp/terraform-provider-google-beta/pull/5126))
* gke: fixed the error of Invalid address to set on `config_connector_config` of the data source `google_container_cluster` ([#5131](https://github.com/hashicorp/terraform-provider-google-beta/pull/5131))
* google_project: fixes misleading examples that could cause `firebase:enabled` label to be accidentally removed. ([#5122](https://github.com/hashicorp/terraform-provider-google-beta/pull/5122))

## 4.50.0 (January 23, 2023)

FEATURES:
* **New Data Source:** `google_compute_network_peering` ([#5092](https://github.com/hashicorp/terraform-provider-google-beta/pull/5092))
* **New Data Source:** `google_compute_router_nat` ([#5091](https://github.com/hashicorp/terraform-provider-google-beta/pull/5091))
* **New Resource:** `google_cloud_run_v2_job_iam_binding` ([#5099](https://github.com/hashicorp/terraform-provider-google-beta/pull/5099))
* **New Resource:** `google_cloud_run_v2_job_iam_member` ([#5099](https://github.com/hashicorp/terraform-provider-google-beta/pull/5099))
* **New Resource:** `google_cloud_run_v2_job_iam_policy` ([#5099](https://github.com/hashicorp/terraform-provider-google-beta/pull/5099))
* **New Resource:** `google_cloud_run_v2_service_iam_binding` ([#5099](https://github.com/hashicorp/terraform-provider-google-beta/pull/5099))
* **New Resource:** `google_cloud_run_v2_service_iam_member` ([#5099](https://github.com/hashicorp/terraform-provider-google-beta/pull/5099))
* **New Resource:** `google_cloud_run_v2_service_iam_policy` ([#5099](https://github.com/hashicorp/terraform-provider-google-beta/pull/5099))
* **New Resource:** `google_gke_backup_backup_plan_iam_binding` ([#5107](https://github.com/hashicorp/terraform-provider-google-beta/pull/5107))
* **New Resource:** `google_gke_backup_backup_plan_iam_member` ([#5107](https://github.com/hashicorp/terraform-provider-google-beta/pull/5107))
* **New Resource:** `google_gke_backup_backup_plan_iam_policy` ([#5107](https://github.com/hashicorp/terraform-provider-google-beta/pull/5107))

IMPROVEMENTS:
* bigquery_table - added `reference_file_schema_uri` ([#5100](https://github.com/hashicorp/terraform-provider-google-beta/pull/5100))
* billingbudget: made fields `credit_types` and `subaccounts` updatable for `google_billing_budget` ([#5087](https://github.com/hashicorp/terraform-provider-google-beta/pull/5087))
* cloudrunV2: added `annotations` to `CloudRunV2_service` resource ([#5108](https://github.com/hashicorp/terraform-provider-google-beta/pull/5108))
* composer: added `recovery_config` in `google_composer_environment` resource ([#5105](https://github.com/hashicorp/terraform-provider-google-beta/pull/5105))
* compute: added support for 'edge_security_policy' field to 'google_compute_backend_service' resource. ([#5101](https://github.com/hashicorp/terraform-provider-google-beta/pull/5101))
* compute: added `max_run_duration` field to `google_compute_instance` and `google_compute_instance_template` resource (beta) ([#5096](https://github.com/hashicorp/terraform-provider-google-beta/pull/5096))
* dataproc: added support for `dataproc_metric_config` to resource `google_dataproc_cluster` ([#5093](https://github.com/hashicorp/terraform-provider-google-beta/pull/5093))
* dlp: added all subfields under `deidentify_template.record_transformations.field_transformations.primitive_transformation` to `google_data_loss_prevention_deidentify_template` ([#5104](https://github.com/hashicorp/terraform-provider-google-beta/pull/5104))
* sql: changed the default create timeout of `google_sql_database_instance` to 40m from 30m ([#5094](https://github.com/hashicorp/terraform-provider-google-beta/pull/5094))

BUG FIXES:
* certificatemanager: removed incorrect indication that the `self_managed` field in `google_certificate_manager_certificate` was treated as sensitive, and marked `self_managed.pem_private_key` as sensitive ([#5106](https://github.com/hashicorp/terraform-provider-google-beta/pull/5106))
* cloudplatform: fixed the error with header `X-Goog-User-Project` on `google_client_openid_userinfo` ([#5090](https://github.com/hashicorp/terraform-provider-google-beta/pull/5090))
* cloudsql: fixed `disk_type` can't be updated on `google_sql_database_instance` ([#5095](https://github.com/hashicorp/terraform-provider-google-beta/pull/5095))
* vertexai: fixed updating value_type in google_vertex_ai_featurestore_entitytype_feature ([#5098](https://github.com/hashicorp/terraform-provider-google-beta/pull/5098))

## 4.49.0 (January 17, 2023)

FEATURES:
* **New Data Source:** `google_project_service` ([#5067](https://github.com/hashicorp/terraform-provider-google-beta/pull/5067))
* **New Data Source:** `google_sql_database_instances` ([#5066](https://github.com/hashicorp/terraform-provider-google-beta/pull/5066))
* **New Data Source:** `google_container_attached_install_manifest` ([#5073](https://github.com/hashicorp/terraform-provider-google-beta/pull/5073))
* **New Data Source:** `google_container_attached_install_manifest` ([#5080](https://github.com/hashicorp/terraform-provider-google-beta/pull/5080))
* **New Data Source:** `google_container_attached_versions` ([#5073](https://github.com/hashicorp/terraform-provider-google-beta/pull/5073))
* **New Resource:** `google_datastream_stream` ([#5045](https://github.com/hashicorp/terraform-provider-google-beta/pull/5045))

IMPROVEMENTS:
* android_app: added general fields `sha1_hashes`, `sha256_hashes` and `etag` to `google_firebase_android_app`. ([#5074](https://github.com/hashicorp/terraform-provider-google-beta/pull/5074))
* cloudids: added `threat_exception` field to `google_cloud_ids_endpoint` ([#5072](https://github.com/hashicorp/terraform-provider-google-beta/pull/5072))
* composer: added `triggerer` field in `google_composer_environment` ([#5055](https://github.com/hashicorp/terraform-provider-google-beta/pull/5055))
* compute: enabled deletion for `stateful_ips` fields in `instance_group_manager` and `region_instance_group_manager`. ([#5064](https://github.com/hashicorp/terraform-provider-google-beta/pull/5064))
* compute: added field `expire_time` to resource `google_compute_region_ssl_certificate` ([#5049](https://github.com/hashicorp/terraform-provider-google-beta/pull/5049))
* compute: added field `expire_time` to resource `google_compute_ssl_certificate` ([#5049](https://github.com/hashicorp/terraform-provider-google-beta/pull/5049))
* container: added `release_channel_latest_version` in `google_container_engine_versions` datasource ([#5044](https://github.com/hashicorp/terraform-provider-google-beta/pull/5044))
* container: added `google_container_aws_node_pool` `autoscaling_metrics_collection` field ([#5084](https://github.com/hashicorp/terraform-provider-google-beta/pull/5084))
* container: added update support for `google_container_aws_node_pool` `tags` field ([#5084](https://github.com/hashicorp/terraform-provider-google-beta/pull/5084))
* dataproc: added support for `node_group_affinity.` in `google_dataproc_cluster` ([#5053](https://github.com/hashicorp/terraform-provider-google-beta/pull/5053))
* dataproc: added support for `reservation_affinity` in `google_dataproc_cluster` ([#5050](https://github.com/hashicorp/terraform-provider-google-beta/pull/5050))
* dlp: added field 'identifyingFields' to 'bigQueryOptions' for creating DLP jobs. ([#5085](https://github.com/hashicorp/terraform-provider-google-beta/pull/5085))
* metastore: added `telemetry_config` field to `google_dataproc_metastore_service` ([#5065](https://github.com/hashicorp/terraform-provider-google-beta/pull/5065))
* sql: added the ability to set `point_in_time_recovery_enabled` flag in `google_sql_database_instance` for `SQLSERVER` instance, since the API supports it now. ([#5079](https://github.com/hashicorp/terraform-provider-google-beta/pull/5079))
* sql: added `instance_type` field to `google_sql_database_instance` resource ([#5057](https://github.com/hashicorp/terraform-provider-google-beta/pull/5057))
* vertexai: added `scaling` field in `google_vertex_ai_featurestore` ([#5081](https://github.com/hashicorp/terraform-provider-google-beta/pull/5081))

BUG FIXES:
* android_app: modified the `package_name` field suffix to always start with a letter in `google_firebase_android_app`. ([#5074](https://github.com/hashicorp/terraform-provider-google-beta/pull/5074))
* bigqueryconnection: fixed a bug where `aws.access_role.iam_role_id` cannot be updated on `google_bigquery_connection` ([#5083](https://github.com/hashicorp/terraform-provider-google-beta/pull/5083))
* container: fixed a bug preventing updates to `master_global_access_config` in `google_container_cluster` ([#5043](https://github.com/hashicorp/terraform-provider-google-beta/pull/5043))
* container: fixed perma-diff based on a server set taint `kubernetes.io/arch` on `google_container_node_pool` ([#5054](https://github.com/hashicorp/terraform-provider-google-beta/pull/5054))
* spanner: fixed crash when `google_spanner_database.ddl` item was nil ([#5071](https://github.com/hashicorp/terraform-provider-google-beta/pull/5071))

## 4.48.0 (January 9, 2023)

FEATURES:
* **New Data Source:** `google_beyondcorp_app_connection` ([#5025](https://github.com/hashicorp/terraform-provider-google-beta/pull/5025))
* **New Data Source:** `google_beyondcorp_app_connector` ([#5008](https://github.com/hashicorp/terraform-provider-google-beta/pull/5008))
* **New Data Source:** `google_beyondcorp_app_gateway` ([#5008](https://github.com/hashicorp/terraform-provider-google-beta/pull/5008))
* **New Data Source:** `google_cloudbuild_trigger` ([#5017](https://github.com/hashicorp/terraform-provider-google-beta/pull/5017))
* **New Data Source:** `google_compute_instance_group_manager` ([#5002](https://github.com/hashicorp/terraform-provider-google-beta/pull/5002))
* **New Data Source:** `google_firebase_apple_app_config` ([#5031](https://github.com/hashicorp/terraform-provider-google-beta/pull/5031))
* **New Data Source:** `google_firebase_apple_app` ([#4978](https://github.com/hashicorp/terraform-provider-google-beta/pull/4978))
* **New Data Source:** `google_pubsub_subscription` ([#5001](https://github.com/hashicorp/terraform-provider-google-beta/pull/5001))
* **New Data Source:** `google_sql_database` ([#5038](https://github.com/hashicorp/terraform-provider-google-beta/pull/5038))
* **New Resource:** `google_apigee_sync_authorization` ([#5015](https://github.com/hashicorp/terraform-provider-google-beta/pull/5015))
* **New Resource:** `google_beyondcorp_app_connection` ([#5012](https://github.com/hashicorp/terraform-provider-google-beta/pull/5012))
* **New Resource:** `google_container_attached_cluster` ([#5037](https://github.com/hashicorp/terraform-provider-google-beta/pull/5037))
* **New Resource:** `google_dns_managed_zone_iam_*` ([#5007](https://github.com/hashicorp/terraform-provider-google-beta/pull/5007))
* **New Resource:** `google_firebase_database_instance`([#5019](https://github.com/hashicorp/terraform-provider-google-beta/pull/5019))

IMPROVEMENTS:
* cloudfunctions2: added `available_cpu` and `max_instance_request_concurrency` to support concurrency in `google_cloudfunctions2_function` ([#5011](https://github.com/hashicorp/terraform-provider-google-beta/pull/5011))
* gkehub: added support for `configmanagement.config_sync.oci` field to `google_gke_hub_feature_membership` resource([#5013](https://github.com/hashicorp/terraform-provider-google-beta/pull/5013))
* compute: added support for local IP ranges in `google_compute_firewall` ([#4979](https://github.com/hashicorp/terraform-provider-google-beta/pull/4979))
* compute: added `router_appliance_instance` field to `google_compute_router_bgp_peer` ([#5035](https://github.com/hashicorp/terraform-provider-google-beta/pull/5035))
* compute: added support for `generated_id` field in `google_compute_backend_service` to get the value of `id` defined by the server ([#4981](https://github.com/hashicorp/terraform-provider-google-beta/pull/4981))
* compute: added support for `image_encryption_key` to `google_compute_image` ([#4989](https://github.com/hashicorp/terraform-provider-google-beta/pull/4989))
* compute: added support for `source_snapshot`, `source_snapshot_encyption_key`, and `source_image_encryption_key` to `google_compute_instance_template` ([#4989](https://github.com/hashicorp/terraform-provider-google-beta/pull/4989))
* compute: added stateful ip fields `stateful_internal_ip` and `stateful_external_ip` to `google_compute_instance_group_manager` ([#4992](https://github.com/hashicorp/terraform-provider-google-beta/pull/4992))
* container: added `gateway_api_config` block to `google_container_cluster` resource for supporting the gke gateway api controller ([#4976](https://github.com/hashicorp/terraform-provider-google-beta/pull/4976))
* container: supported in-place update for `labels` in `google_container_node_pool` ([#4998](https://github.com/hashicorp/terraform-provider-google-beta/pull/4998))
* dataproc: added support for `SPOT` option for `preemptibility` in `google_dataproc_cluster` ([#5024](https://github.com/hashicorp/terraform-provider-google-beta/pull/5024))
* dlp: added field `deidentify_config.record_transformations.field_transformations` to `google_data_loss_prevention_deidentify_template` ([#4996](https://github.com/hashicorp/terraform-provider-google-beta/pull/4996))
* dlp: added field `deidentify_config.record_transformations.record_suppressions` to `google_data_loss_prevention_deidentify_template` ([#5004](https://github.com/hashicorp/terraform-provider-google-beta/pull/5004))
* dlp: added `version` field to `google_data_loss_prevention_inspect_template` resource ([#5032](https://github.com/hashicorp/terraform-provider-google-beta/pull/5032))
* osconfig: added support for `skip_await_rollout` in `google_os_config_os_policy_assignment` ([#5026](https://github.com/hashicorp/terraform-provider-google-beta/pull/5026))
* sql: added [new deletion protection](https://cloud.google.com/sql/docs/mysql/deletion-protection) feature `deletion_protection_enabled` in `google_sql_database_instance` to guard against deletion from all surfaces ([#4987](https://github.com/hashicorp/terraform-provider-google-beta/pull/4987))
* sql: made `settings.sql_server_audit_config.bucket` field in `google_sql_database_instance` to be optional. ([#4988](https://github.com/hashicorp/terraform-provider-google-beta/pull/4988))
* storagetransfer: supported in-place update for `schedule` in `google_storage_transfer_job` ([#4993](https://github.com/hashicorp/terraform-provider-google-beta/pull/4993))

BUG FIXES:
* bigquery: fixed a permadiff on `labels` of `google_bigquery_dataset` when it is referenced in `google_dataplex_asset` ([#5022](https://github.com/hashicorp/terraform-provider-google-beta/pull/5022))
* compute: fixed a permadiff on `private_ip_google_access` of `google_compute_subnetwork` ([#4983](https://github.com/hashicorp/terraform-provider-google-beta/pull/4983))
* compute: fixed an issue where `enable_dynamic_port_allocation` was not able to set to `false` in `google_compute_router_nat` ([#4982](https://github.com/hashicorp/terraform-provider-google-beta/pull/4982))
* container: fixed a permadiff on `location_policy` of `google_container_cluster` and `google_container_node_pool` ([#4997](https://github.com/hashicorp/terraform-provider-google-beta/pull/4997))
* identityplatform: fixed issues with `google_identity_platform_config` creation  ([#5005](https://github.com/hashicorp/terraform-provider-google-beta/pull/5005))
* resourcemanager: fixed the `google_project` datasource silently returning empty results when the project was not found or not in the ACTIVE state. Now, an error will be surfaced instead. ([#5029](https://github.com/hashicorp/terraform-provider-google-beta/pull/5029))
* sql: fixed `sql_database_instance` leaking root users ([#4991](https://github.com/hashicorp/terraform-provider-google-beta/pull/4991))

## 4.47.0 (December 21, 2022)

NOTES:
* sql: fixed an issue where `google_sql_database` was abandoned by default as of version `4.45.0`. Users who have upgraded to `4.45.0` or `4.46.0` will see a diff when running their next `terraform apply` after upgrading this version, indicating the `deletion_policy` field's value has changed from `"ABANDON"` to `"DELETE"`. This will create a no-op call against the API, but can otherwise be safely applied. ([#4972](https://github.com/hashicorp/terraform-provider-google-beta/pull/4972))

FEATURES:
* **New Resource:** `google_alloydb_backup` ([#4959](https://github.com/hashicorp/terraform-provider-google-beta/pull/4959))
* **New Resource:** `google_filestore_backup` ([#4963](https://github.com/hashicorp/terraform-provider-google-beta/pull/4963))

IMPROVEMENTS:
* bigtable: added `deletion_protection` field to `google_bigtable_table` ([#4975](https://github.com/hashicorp/terraform-provider-google-beta/pull/4975))
* compute: Made `google_compute_subnetwork.ipv6_access_type` field updatable in-place ([#4965](https://github.com/hashicorp/terraform-provider-google-beta/pull/4965))
* container: added `auto_provisioning_defaults.cluster_autoscaling.upgrade_settings` in `google_container_cluster` ([#4958](https://github.com/hashicorp/terraform-provider-google-beta/pull/4958))
* container: added `gateway_api_config` block to `google_container_cluster` resource for supporting the gke gateway api controller ([#4976](https://github.com/hashicorp/terraform-provider-google-beta/pull/4976))
* datacatalog: added update support for `fields` in `google_data_catalog_tag_template` ([#4968](https://github.com/hashicorp/terraform-provider-google-beta/pull/4968))
* iam: added plan-time validation for IAM members ([#4960](https://github.com/hashicorp/terraform-provider-google-beta/pull/4960))
* logging: added `bucket_name` field to `google_logging_metric` ([#4964](https://github.com/hashicorp/terraform-provider-google-beta/pull/4964))
* logging: made `metric_descriptor` field optional for `google_logging_metric` ([#4971](https://github.com/hashicorp/terraform-provider-google-beta/pull/4971))

BUG FIXES:
* composer: fixed a crash when updating `ip_allocation_policy` of `google_composer_environment` ([#4956](https://github.com/hashicorp/terraform-provider-google-beta/pull/4956))
* sql: fixed an issue where `google_sql_database` was abandoned by default as of version `4.45.0`. Users who have upgraded to `4.45.0` or `4.46.0` will see a diff when running their next `terraform apply` after upgrading this version, indicating the `deletion_policy` field's value has changed from `"ABANDON"` to `"DELETE"`. This will create a no-op call against the API, but can otherwise be safely applied. ([#4972](https://github.com/hashicorp/terraform-provider-google-beta/pull/4972))

## 4.46.0 (December 12, 2022)

FEATURES:
* **New Data Source:** `google_firebase_android_app` ([#4955](https://github.com/hashicorp/terraform-provider-google-beta/pull/4955))
* **New Resource:** `google_cloud_run_v2_job` ([#4937](https://github.com/hashicorp/terraform-provider-google-beta/pull/4937))
* **New Resource:** `google_cloud_run_v2_service` ([#4942](https://github.com/hashicorp/terraform-provider-google-beta/pull/4942))
* **New Resource:** `google_gke_backup_backup_plan` ([#4948](https://github.com/hashicorp/terraform-provider-google-beta/pull/4948))
* **New Resource:** `google_firebase_storage_bucket` ([#4951](https://github.com/hashicorp/terraform-provider-google-beta/pull/4951))

IMPROVEMENTS:
* network_services: added `origin_override_action` and `origin_redirect` to `google_network_services_edge_cache_origin` ([#4936](https://github.com/hashicorp/terraform-provider-google-beta/pull/4936))
* bigquerydatatransfer: recreate `google_bigquery_data_transfer_config` for Cloud Storage transfers when immutable params `data_path_template` and `destination_table_name_template` are changed ([#4929](https://github.com/hashicorp/terraform-provider-google-beta/pull/4929))
* compute: added fields to resource `google_compute_security_policy` to support Cloud Armor bot management ([#4938](https://github.com/hashicorp/terraform-provider-google-beta/pull/4938))
* container: added support for concurrent node pool mutations on a cluster. Previously, node pool mutations were restricted to run synchronously clientside. NOTE: While this feature is supported in Terraform from this release onwards, only a limited number of GCP projects will support this behavior initially. The provider will automatically process mutations concurrently as the feature rolls out generally. ([#4947](https://github.com/hashicorp/terraform-provider-google-beta/pull/4947))
* metastore: added general field `network_config` to `google_dataproc_metastore_service` ([#4952](https://github.com/hashicorp/terraform-provider-google-beta/pull/4952))
* storage: added support for `autoclass` in `google_storage_bucket` resource ([#4953](https://github.com/hashicorp/terraform-provider-google-beta/pull/4953))

BUG FIXES:
* alloydb: made `machine_config.cpu_count` updatable on `google_alloydb_instance` ([#4930](https://github.com/hashicorp/terraform-provider-google-beta/pull/4930))
* composer: fixed a crash when updating `ip_allocation_policy` of `google_composer_environment` ([#4956](https://github.com/hashicorp/terraform-provider-google-beta/pull/4956))
* container: fixed GKE permadiff/thrashing when `update_settings. max_surge` or `update_settings. max_unavailable` values are updating on `google_container_node_pool` ([#4945](https://github.com/hashicorp/terraform-provider-google-beta/pull/4945))
* datastream: fixed `google_datastream_private_connection` ignoring failures during creation ([#4939](https://github.com/hashicorp/terraform-provider-google-beta/pull/4939))
* firebase: fixed permadiff on the field `deletion_policy` of `google_firebase_apple_app` ([#4954](https://github.com/hashicorp/terraform-provider-google-beta/pull/4954))
* kms: fixed issues with deleting crypto key versions in states other than ENABLED ([#4943](https://github.com/hashicorp/terraform-provider-google-beta/pull/4943))

## 4.45.0 (December 5, 2022)

FEATURES:
* **New Data Source:** `google_logging_project_cmek_settings` ([#4902](https://github.com/hashicorp/terraform-provider-google-beta/pull/4902))
* **New Resource:** `google_iam_workforce_pool_provider` ([#4922](https://github.com/hashicorp/terraform-provider-google-beta/pull/4922))
* **New Resource:** `google_vertex_ai_tensorboard` ([#4896](https://github.com/hashicorp/terraform-provider-google-beta/pull/4896))
* **New Resource:** `google_data_fusion_instance_iam_binding` ([#4926](https://github.com/hashicorp/terraform-provider-google-beta/pull/4926))
* **New Resource:** `google_data_fusion_instance_iam_member` ([#4926](https://github.com/hashicorp/terraform-provider-google-beta/pull/4926))
* **New Resource:** `google_data_fusion_instance_iam_policy` ([#4926](https://github.com/hashicorp/terraform-provider-google-beta/pull/4926))
* **New Resource:** `google_eventarc_google_channel_config` ([#4905](https://github.com/hashicorp/terraform-provider-google-beta/pull/4905))
* **New Resource:** `google_vertex_ai_index` ([#4923](https://github.com/hashicorp/terraform-provider-google-beta/pull/4923))
* **New Resource:** `google_vertex_ai_featurestore_entitytype_iam_binding` ([#4920](https://github.com/hashicorp/terraform-provider-google-beta/pull/4920))
* **New Resource:** `google_vertex_ai_featurestore_entitytype_iam_member` ([#4920](https://github.com/hashicorp/terraform-provider-google-beta/pull/4920))
* **New Resource:** `google_vertex_ai_featurestore_entitytype_iam_policy` ([#4920](https://github.com/hashicorp/terraform-provider-google-beta/pull/4920))


IMPROVEMENTS:
* anthos-fleet-management: added option `mesh: control_plane` to resource `google_gke_hub_feature_membership`. ([#4927](https://github.com/hashicorp/terraform-provider-google-beta/pull/4927))
* bigquerydatatransfer: made `google_bigquery_data_transfer_config` recreate for Cloud Storage transfers when immutable params `data_path_template` and `destination_table_name_template` are changed ([#4929](https://github.com/hashicorp/terraform-provider-google-beta/pull/4929))
* bigtable: added support for abandoning GC policy ([#4897](https://github.com/hashicorp/terraform-provider-google-beta/pull/4897))
* cloudsql: added `connector_enforcement` field to `google_sql_database_instance` resource ([#4894](https://github.com/hashicorp/terraform-provider-google-beta/pull/4894))
* compute: added `default_route_action.cors_policy` field to `google_compute_region_url_map` resource ([#4895](https://github.com/hashicorp/terraform-provider-google-beta/pull/4895))
* compute: added `default_route_action.fault_injection_policy` field to `google_compute_region_url_map` resource ([#4895](https://github.com/hashicorp/terraform-provider-google-beta/pull/4895))
* compute: added `default_route_action.timeout` field to `google_compute_region_url_map` resource ([#4895](https://github.com/hashicorp/terraform-provider-google-beta/pull/4895))
* compute: added `default_route_action.url_rewrite` field to `google_compute_region_url_map` resource ([#4895](https://github.com/hashicorp/terraform-provider-google-beta/pull/4895))
* compute: added `include_http_headers` field to the `cdn_policy` field of `google_compute_backend_service` resource ([#4912](https://github.com/hashicorp/terraform-provider-google-beta/pull/4912))
* compute: added field `list_managed_instances_results` to `google_compute_instance_group_manager` and `google_compute_region_instance_group_manager` ([#4903](https://github.com/hashicorp/terraform-provider-google-beta/pull/4903))
* compute: added subnetwork and private_ip_address arguments to resource_compute_router_interface ([#4914](https://github.com/hashicorp/terraform-provider-google-beta/pull/4914))
* container: added `resource_labels` field to `node_config` resource ([#4913](https://github.com/hashicorp/terraform-provider-google-beta/pull/4913))
* container: added field `enable_private_nodes` in `network_config` to `google_container_node_pool` ([#4921](https://github.com/hashicorp/terraform-provider-google-beta/pull/4921))
* container: added field `gcp_public_cidrs_access_enabled` and `private_endpoint_subnetwork` to `google_container_cluster` ([#4921](https://github.com/hashicorp/terraform-provider-google-beta/pull/4921))
* container: added update support for `enable_private_endpoint` and `enable_private_nodes` in `google_container_cluster` ([#4921](https://github.com/hashicorp/terraform-provider-google-beta/pull/4921))
* datafusion: added `api_endpoint` and `p4_service_account ` attributes to `google_data_fusion_instance` ([#4926](https://github.com/hashicorp/terraform-provider-google-beta/pull/4926))
* datafusion: added `zone`, `display_name`, `crypto_key_config`, `event_publish_config`, and `enable_rbac` args to `google_data_fusion_instance` ([#4926](https://github.com/hashicorp/terraform-provider-google-beta/pull/4926))
* logging: added `cmek_settings` field to `google_logging_project_bucket_config` resource ([#4902](https://github.com/hashicorp/terraform-provider-google-beta/pull/4902))
* sql: added 'deny_maintenance_period' field for 'google_sql_database_instance' within which 'end_date', 'start_date' and 'time' fields are present. ([#4915](https://github.com/hashicorp/terraform-provider-google-beta/pull/4915))
* sql: added field `deletion_policy` to resource `google_sql_database` ([#4916](https://github.com/hashicorp/terraform-provider-google-beta/pull/4916))

BUG FIXES:
* alloydb: fixed permdiff on `display_name` of `google_alloydb_instance` ([#4925](https://github.com/hashicorp/terraform-provider-google-beta/pull/4925))
* compute: fixed a failure in updating `most_disruptive_allowed_action` on `google_compute_per_instance_config` and `google_compute_region_per_instance_config` ([#4898](https://github.com/hashicorp/terraform-provider-google-beta/pull/4898))
* compute: fixed the error when `metadata` and `machine_type` are updated while `metadata_startup_script` was already provided on `google_compute_instance` ([#4901](https://github.com/hashicorp/terraform-provider-google-beta/pull/4901))
* container: fixed the inability to update `authenticator_groups_config` on `google_container_cluster` ([#4918](https://github.com/hashicorp/terraform-provider-google-beta/pull/4918))
* container: fixed the data source `google_container_cluster` to return an error if it does not exist ([#4900](https://github.com/hashicorp/terraform-provider-google-beta/pull/4900))
* sql: fixed `googe_sql_database_instance` to include `backup_configuration` in initial create request ([#4911](https://github.com/hashicorp/terraform-provider-google-beta/pull/4911))
* storage: fixed permdiff when `website`, `website.main_page_suffix`, `website.not_found_page` are removed on `google_storage_bucket` ([#4899](https://github.com/hashicorp/terraform-provider-google-beta/pull/4899))



## 4.44.1 (November 22, 2022)

NOTES: No changes, only released to keep this provider in sync with the `google` provider

## 4.44.0 (November 21, 2022)

FEATURES:
* **New Data Source:** `google_cloud_asset_resources_search_all` ([#4891](https://github.com/hashicorp/terraform-provider-google-beta/pull/4891))
* **New Resource:** `google_alloydb_instance` ([#4857](https://github.com/hashicorp/terraform-provider-google-beta/pull/4857))
* **New Resource:** `google_beyondcorp_app_connector` ([#4866](https://github.com/hashicorp/terraform-provider-google-beta/pull/4866))
* **New Resource:** `google_beyondcorp_app_gateway` ([#4866](https://github.com/hashicorp/terraform-provider-google-beta/pull/4866))
* **New Resource:** `google_compute_network_firewall_policy_association` ([#4868](https://github.com/hashicorp/terraform-provider-google-beta/pull/4868))
* **New Resource:** `google_compute_network_firewall_policy_rule` ([#4880](https://github.com/hashicorp/terraform-provider-google-beta/pull/4880))
* **New Resource:** `google_compute_network_firewall_policy` ([#4851](https://github.com/hashicorp/terraform-provider-google-beta/pull/4851))
* **New Resource:** `google_compute_region_network_firewall_policy_association` ([#4868](https://github.com/hashicorp/terraform-provider-google-beta/pull/4868))
* **New Resource:** `google_compute_region_network_firewall_policy_rule` ([#4880](https://github.com/hashicorp/terraform-provider-google-beta/pull/4880))
* **New Resource:** `google_compute_region_network_firewall_policy` ([#4851](https://github.com/hashicorp/terraform-provider-google-beta/pull/4851))
* **New Resource:** `google_eventarc_channel` ([#4876](https://github.com/hashicorp/terraform-provider-google-beta/pull/4876))
* **New Resource:** `google_firebase_apple_app` ([#4887](https://github.com/hashicorp/terraform-provider-google-beta/pull/4887))
* **New Resource:** `google_firebase_hosting_channel` ([#4890](https://github.com/hashicorp/terraform-provider-google-beta/pull/4890))
* **New Resource:** `google_firebase_hosting_site` ([#4846](https://github.com/hashicorp/terraform-provider-google-beta/pull/4846))
* **New Resource:** `google_identity_platform_project_default_config` ([#4853](https://github.com/hashicorp/terraform-provider-google-beta/pull/4853))
* **New Resource:** `google_kms_crypto_key_versions` ([#4831](https://github.com/hashicorp/terraform-provider-google-beta/pull/4831))
* **New Resource:** `google_storage_transfer_agent_pool` ([#4835](https://github.com/hashicorp/terraform-provider-google-beta/pull/4835))

IMPROVEMENTS:
* bigquery: supported authorized routines on resource `bigquery_dataset` and `bigquery_dataset_access` ([#4855](https://github.com/hashicorp/terraform-provider-google-beta/pull/4855))
* clouddeploy: added execution_configs.execution_timeout in target resource. ([#4849](https://github.com/hashicorp/terraform-provider-google-beta/pull/4849))
* clouddeploy: added support for Cloud Run Targets ([#4849](https://github.com/hashicorp/terraform-provider-google-beta/pull/4849))
* clouddeploy: added support for Deployment Verification standard strategy ([#4849](https://github.com/hashicorp/terraform-provider-google-beta/pull/4849))
* cloudidentity: made security label settable by making labels updatable in `google_cloud_identity_groups` ([#4834](https://github.com/hashicorp/terraform-provider-google-beta/pull/4834))
* cloudrun: added field `liveness_probe.grpc` and `startup_probe.grpc` to resource `google_cloud_run_service` ([#4863](https://github.com/hashicorp/terraform-provider-google-beta/pull/4863))
* cloudsql: added `connector_enforcement` field to `google_sql_database_instance` resource ([#4894](https://github.com/hashicorp/terraform-provider-google-beta/pull/4894))
* compute: added optional `redundant_interface` argument to `google_compute_router_interface` resource ([#4881](https://github.com/hashicorp/terraform-provider-google-beta/pull/4881))
* compute: added `default_route_action.request_mirror_policy` field to `google_compute_region_url_map` resource ([#4879](https://github.com/hashicorp/terraform-provider-google-beta/pull/4879))
* compute: added `default_route_action.retry_policy` field to `google_compute_region_url_map` resource ([#4879](https://github.com/hashicorp/terraform-provider-google-beta/pull/4879))
* compute: added `default_route_action.weighted_backend_services` field to `google_compute_region_url_map` resource ([#4879](https://github.com/hashicorp/terraform-provider-google-beta/pull/4879))
* compute: added `preconfigured_waf_config` block to `google_compute_security_policy` resource ([#4852](https://github.com/hashicorp/terraform-provider-google-beta/pull/4852))
* compute: modified machine_type field in compute instance resource to accept short name. ([#4849](https://github.com/hashicorp/terraform-provider-google-beta/pull/4849))
* container: added `node_config.logging_variant` to `google_container_node_pool`. ([#4889](https://github.com/hashicorp/terraform-provider-google-beta/pull/4889))
* container: added `node_pool_defaults.node_config_defaults.logging_variant`, `node_pool.node_config.logging_variant`, and `node_config.logging_variant` to `google_container_cluster`. ([#4889](https://github.com/hashicorp/terraform-provider-google-beta/pull/4889))
* container: added support for Shielded Instance configuration for node auto-provisioning to `google_container_cluster` ([#4833](https://github.com/hashicorp/terraform-provider-google-beta/pull/4833))
* container: added management attribute to the google_container_cluster resource ([#4862](https://github.com/hashicorp/terraform-provider-google-beta/pull/4862))
* container: added field `blue_green_settings` to `google_container_node_pool` ([#4860](https://github.com/hashicorp/terraform-provider-google-beta/pull/4860))
* container: added field `strategy` to `google_container_node_pool` ([#4860](https://github.com/hashicorp/terraform-provider-google-beta/pull/4860))
* container: added support for additional values `APISERVER`, `CONTROLLER_MANAGER`, and `SCHEDULER` in `google_container_cluster.monitoring_config` ([#4854](https://github.com/hashicorp/terraform-provider-google-beta/pull/4854))
* datafusion: added `enable_rbac` field to `google_data_fusion_instance` resource ([#4864](https://github.com/hashicorp/terraform-provider-google-beta/pull/4864))
* dlp: added fields `rows_limit`, `rows_limit_percent`, and `sample_method` to `big_query_options` in `google_data_loss_prevention_job_trigger` ([#4856](https://github.com/hashicorp/terraform-provider-google-beta/pull/4856))
* dlp: added pubsub action to `google_data_loss_prevention_job_trigger` ([#4832](https://github.com/hashicorp/terraform-provider-google-beta/pull/4832))
* dns: added `gke_clusters` field to `google_dns_managed_zone` resource ([#4888](https://github.com/hashicorp/terraform-provider-google-beta/pull/4888))
* dns: added `gke_clusters` field to `google_dns_response_policy` resource ([#4888](https://github.com/hashicorp/terraform-provider-google-beta/pull/4888))
* eventarc: added field `channel` to `google_eventarc_trigger` ([#4876](https://github.com/hashicorp/terraform-provider-google-beta/pull/4876))
* gkehub: added `mesh` field and `management` subfield to resource `feature_membership` ([#4867](https://github.com/hashicorp/terraform-provider-google-beta/pull/4867))
* networkservices: added `aws_v4_authentication ` field to `google_network_services_edge_cache_origin ` to support S3-compatible Origins ([#4875](https://github.com/hashicorp/terraform-provider-google-beta/pull/4875))
* networkservices: added `signed_token_options` and `add_signatures` field to `google_network_services_edge_cache_service` and `validation_shared_keys` to `google_network_services_edge_cache_keyset` to support dual-token authentication ([#4884](https://github.com/hashicorp/terraform-provider-google-beta/pull/4884))
* sql: added `query_plan_per_minute` field to `insights_config` in `google_sql_database_instance` resource ([#4840](https://github.com/hashicorp/terraform-provider-google-beta/pull/4840))
* vertexai: added fields to `vertex_ai_featurestore_entitytype` to support feature value monitoring ([#4859](https://github.com/hashicorp/terraform-provider-google-beta/pull/4859))

BUG FIXES:
* apigee: fixed permadiff on consumer_accept_list for `google_apigee_instance` ([#4883](https://github.com/hashicorp/terraform-provider-google-beta/pull/4883))
* appengine: fixed permadiff on serviceaccount for 'google_app_engine_flexible_app_version' ([#4858](https://github.com/hashicorp/terraform-provider-google-beta/pull/4858))
* bigtable: updated ForceNew logic for `kms_key_name` ([#4873](https://github.com/hashicorp/terraform-provider-google-beta/pull/4873))
* bigtable: updated the error handling logic to remove the resource on resource not found error only ([#4841](https://github.com/hashicorp/terraform-provider-google-beta/pull/4841))
* billingbudget: fixed a bug where `budget_filter.credit_types_treatment` in `google_billing_budget` resource was not updating. ([#4836](https://github.com/hashicorp/terraform-provider-google-beta/pull/4836))
* cloudbuild: fixed a failure when BITBUCKET is provided for `repo_type` on `google_cloudbuild_trigger` ([#4878](https://github.com/hashicorp/terraform-provider-google-beta/pull/4878))
* cloudids: fixed `endpoint_forwarding_rule` and `endpoint_ip` attributes for `google_cloud_ids_endpoint` ([#4843](https://github.com/hashicorp/terraform-provider-google-beta/pull/4843))
* compute: fixed perma-diff on `google_compute_disk` for new amd64 images ([#4847](https://github.com/hashicorp/terraform-provider-google-beta/pull/4847))
* compute: made `target_https_proxy` possible to set `ssl_certificates` and `certificate_map` in `google_compute_target_https_proxy`  at the same time ([#4839](https://github.com/hashicorp/terraform-provider-google-beta/pull/4839))
* container: fixed a bug where `cluster_autoscaling.auto_provisioning_defaults.service_account` can not be set when `enable_autopilot = true` for `google_container_cluster` ([#4877](https://github.com/hashicorp/terraform-provider-google-beta/pull/4877))
* dialogflowcx: fixed a deployment issue for `google_dialogflow_cx_version` and `google_dialogflow_cx_environment` when they are deployed to a non-global location ([#4869](https://github.com/hashicorp/terraform-provider-google-beta/pull/4869))
* dns: fixed apply failure when `description` is set to empty string on `google_dns_managed_zone` ([#4837](https://github.com/hashicorp/terraform-provider-google-beta/pull/4837))
* provider: fixed a crash during provider authentication for certain environments ([#4892](https://github.com/hashicorp/terraform-provider-google-beta/pull/4892))
* storage: fixed a crash when `log_bucket` is updated with empty body on `google_storage_bucket` ([#4893](https://github.com/hashicorp/terraform-provider-google-beta/pull/4893))
* vertexai: made google_vertex_ai_featurestore_entitytype always use regional endpoint corresponding to parent's region ([#4845](https://github.com/hashicorp/terraform-provider-google-beta/pull/4845))

## 4.43.0 (November 7, 2022)

FEATURES:
* **New Resource:** `google_kms_crypto_key_version`([#4831](https://github.com/hashicorp/terraform-provider-google-beta/pull/4831))

## 4.42.1 (November 2, 2022)

BUG FIXES:
* storage: fixed a crash in `google_storage_bucket` when upgrading provider to version `4.42.0` with `lifecycle_rule.condition.age` unset ([#4828](https://github.com/hashicorp/terraform-provider-google-beta/pull/4828))

## 4.42.0 (October 31, 2022)

FEATURES:
* **New Data Source:** `google_compute_addresses` ([#4802](https://github.com/hashicorp/terraform-provider-google-beta/pull/4802))
* **New Data Source:** `google_compute_region_network_endpoint_group` ([#4811](https://github.com/hashicorp/terraform-provider-google-beta/pull/4811))
* **New Resource:** `google_alloydb_cluster` ([#4780](https://github.com/hashicorp/terraform-provider-google-beta/pull/4780))
* **New Resource:** `google_dataform_repository` (beta) ([#4801](https://github.com/hashicorp/terraform-provider-google-beta/pull/4801))
* **New Resource:** `google_firebase_android_app` ([#4814](https://github.com/hashicorp/terraform-provider-google-beta/pull/4814))
* **New Resource:** `google_iam_workforce_pool` ([#4818](https://github.com/hashicorp/terraform-provider-google-beta/pull/4818))
* **New Resource:** `google_monitoring_generic_service` ([#4789](https://github.com/hashicorp/terraform-provider-google-beta/pull/4789))
* **New Resource:** `google_scc_source_iam_binding` ([#4806](https://github.com/hashicorp/terraform-provider-google-beta/pull/4806))
* **New Resource:** `google_scc_source_iam_member` ([#4806](https://github.com/hashicorp/terraform-provider-google-beta/pull/4806))
* **New Resource:** `google_scc_source_iam_policy` ([#4806](https://github.com/hashicorp/terraform-provider-google-beta/pull/4806))
* **New Resource:** `google_vertex_ai_endpoint` ([#4815](https://github.com/hashicorp/terraform-provider-google-beta/pull/4815))
* **New Resource:** `google_vertex_ai_featurestore_iam_binding` ([#4825](https://github.com/hashicorp/terraform-provider-google-beta/pull/4825))
* **New Resource:** `google_vertex_ai_featurestore_iam_member` ([#4825](https://github.com/hashicorp/terraform-provider-google-beta/pull/4825))
* **New Resource:** `google_vertex_ai_featurestore_iam_policy` ([#4825](https://github.com/hashicorp/terraform-provider-google-beta/pull/4825))

IMPROVEMENTS:
* appengine: added `member` field to `google_app_engine_default_service_account` datasource ([#4779](https://github.com/hashicorp/terraform-provider-google-beta/pull/4779))
* bigquery: added `max_time_travel_hours` field in `google_bigquery_dataset` resource ([#4803](https://github.com/hashicorp/terraform-provider-google-beta/pull/4803))
* bigquery: added `member` field to `google_bigquery_default_service_account` datasource ([#4779](https://github.com/hashicorp/terraform-provider-google-beta/pull/4779))
* cloudbuild: added `script` field to `google_cloudbuild_trigger` resource ([#4807](https://github.com/hashicorp/terraform-provider-google-beta/pull/4807))
* cloudplatform: validated `project_id` for `google_project` data-source ([#4810](https://github.com/hashicorp/terraform-provider-google-beta/pull/4810))
* cloudrun: added field `liveness_probe` to resource `google_cloud_run_service` ([#4788](https://github.com/hashicorp/terraform-provider-google-beta/pull/4788))
* cloudrun: added field `startup_probe` to resource `google_cloud_run_service` ([#4773](https://github.com/hashicorp/terraform-provider-google-beta/pull/4773))
* compute: added `source_disk` field to `google_compute_disk` and `google_compute_region_disk` resource ([#4783](https://github.com/hashicorp/terraform-provider-google-beta/pull/4783))
* compute: added general field `rules` to `google_compute_router_nat` ([#4797](https://github.com/hashicorp/terraform-provider-google-beta/pull/4797))
* container: added `disk_size` and `disk_type` fields to `google_container_cluster.cluster_autoscaling.auto_provisioning_defaults` ([#4786](https://github.com/hashicorp/terraform-provider-google-beta/pull/4786))
* container: added support for in-place update of `node_config.0.tags` for `google_container_node_pool` resource ([#4781](https://github.com/hashicorp/terraform-provider-google-beta/pull/4781))
* datastream: added `private_connectivity` field to `google_datastream_connection_profile` ([#4808](https://github.com/hashicorp/terraform-provider-google-beta/pull/4808))
* dns: added `enable_geo_fencing` to `routing_policy` block of `google_dns_record_set` resource ([#4816](https://github.com/hashicorp/terraform-provider-google-beta/pull/4816))
* dns: added `health_checked_targets` to `wrr` and `geo` blocks of `google_dns_record_set` resource ([#4816](https://github.com/hashicorp/terraform-provider-google-beta/pull/4816))
* dns: added `primary_backup` to `routing_policy` block of `google_dns_record_set` resource ([#4816](https://github.com/hashicorp/terraform-provider-google-beta/pull/4816))
* firebase: added `app_urls` field to `google_firebase_web_app` ([#4798](https://github.com/hashicorp/terraform-provider-google-beta/pull/4798))
* firebase: added deletion support and new field `deletion_policy` for `google_firebase_web_app` ([#4796](https://github.com/hashicorp/terraform-provider-google-beta/pull/4796))
* privateca: added a new field `skip_grace_period` to skip the grace period when deleting a CertificateAuthority. ([#4784](https://github.com/hashicorp/terraform-provider-google-beta/pull/4784))
* serviceaccount: added `member` field to `google_service_account` resource and datasource ([#4779](https://github.com/hashicorp/terraform-provider-google-beta/pull/4779))
* sql: added `time_zone` field in `google_sql_database_instance` ([#4774](https://github.com/hashicorp/terraform-provider-google-beta/pull/4774))
* storage: added `member` field to `google_storage_project_service_account` and `google_storage_transfer_project_service_account` datasource ([#4779](https://github.com/hashicorp/terraform-provider-google-beta/pull/4779))

BUG FIXES:
* compute: made `vm_count` in `google_compute_resource_policy` optional ([#4792](https://github.com/hashicorp/terraform-provider-google-beta/pull/4792))
* container: fixed inability to update `datapath_provider` on `google_container_cluster` by making field changes trigger resource recreation ([#4824](https://github.com/hashicorp/terraform-provider-google-beta/pull/4824))
* pubsub: ensured topics are recreated when their schemas change. ([#4791](https://github.com/hashicorp/terraform-provider-google-beta/pull/4791))
* redis: updated `persistence_config.rdb_snapshot_period` to optional in the `google_redis_instance` resource. ([#4821](https://github.com/hashicorp/terraform-provider-google-beta/pull/4821))

## 4.41.0 (October 17, 2022)

KNOWN ISSUES:
* container: This release introduced a new field, `node_config.0.guest_accelerator.0.gpu_sharing_config`, to an https://www.terraform.io/language/attr-as-blocks field (`node_config.0.guest_accelerator`). As detailed on the linked page, this may cause issues for modules and/or formats other than HCL.

BREAKING CHANGES:
* sql: updated `google_sql_user.sql_server_user_details` to be read only. Any configuration attempting to set this field is invalid and will cause the provider to fail during plan time. ([#4764](https://github.com/hashicorp/terraform-provider-google-beta/pull/4764))

FEATURES:
* **New Resource:**  `google_cloud_ids_endpoint` ([#4765](https://github.com/hashicorp/terraform-provider-google-beta/pull/4765))
* **New Resource:** `google_bigquery_analytics_hub_listing_iam_binding` ([#4771](https://github.com/hashicorp/terraform-provider-google-beta/pull/4771))
* **New Resource:** `google_bigquery_analytics_hub_listing_iam_member` ([#4771](https://github.com/hashicorp/terraform-provider-google-beta/pull/4771))
* **New Resource:** `google_bigquery_analytics_hub_listing_iam_policy` ([#4771](https://github.com/hashicorp/terraform-provider-google-beta/pull/4771))
* **New Resource:** `google_bigquery_analytics_hub_listing` ([#4771](https://github.com/hashicorp/terraform-provider-google-beta/pull/4771))

IMPROVEMENTS:
* appengine: added support for `service_account` field to `google_app_engine_standard_app_version` resource ([#4757](https://github.com/hashicorp/terraform-provider-google-beta/pull/4757))
* bigquery: added `avro_options` field to `google_bigquery_table` resource ([#4768](https://github.com/hashicorp/terraform-provider-google-beta/pull/4768))
* cloudrun: added field `startup_probe` to resource `google_cloud_run_service` ([#4773](https://github.com/hashicorp/terraform-provider-google-beta/pull/4773))
* compute: added `node_config.0.guest_accelerator.0.gpu_sharing_config` field to `google_container_node_pool` resource ([#4758](https://github.com/hashicorp/terraform-provider-google-beta/pull/4758))
* datafusion: added `crypto_key_config` field to `google_data_fusion_instance` resource ([#4761](https://github.com/hashicorp/terraform-provider-google-beta/pull/4761))
* filestore: removed constraint that forced multiple `google_filestore_instance` creations to occur serially ([#4770](https://github.com/hashicorp/terraform-provider-google-beta/pull/4770))

BUG FIXES:
* kms: fixed apply failure when `google_kms_crypto_key` is removed after its versions were destroyed earlier ([#4769](https://github.com/hashicorp/terraform-provider-google-beta/pull/4769))
* monitoring: fixed a bug causing a perma-diff in `google_monitoring_alert_policy` when `cross_series_reducer` was set to "REDUCE_NONE" ([#4763](https://github.com/hashicorp/terraform-provider-google-beta/pull/4763))

## 4.40.0 (October 10, 2022)

FEATURES:
* **New Data Source:** `google_cloudfunctions2_function` ([#4732](https://github.com/hashicorp/terraform-provider-google-beta/pull/4732))
* **New Data Source:** `google_compute_snapshot` ([#4731](https://github.com/hashicorp/terraform-provider-google-beta/pull/4731))
* **New Resource:** `google_compute_region_target_tcp_proxy` ([#4749](https://github.com/hashicorp/terraform-provider-google-beta/pull/4749))
* **New Resource:** `google_identity_platform_config` ([#4729](https://github.com/hashicorp/terraform-provider-google-beta/pull/4729))
* **New Resource:** `google_bigquery_datapolicy_data_policy` ([#4754](https://github.com/hashicorp/terraform-provider-google-beta/pull/4754))
* **New Resource:** `google_bigquery_datapolicy_data_policy_iam_binding` ([#4754](https://github.com/hashicorp/terraform-provider-google-beta/pull/4754))
* **New Resource:** `google_bigquery_datapolicy_data_policy_iam_member` ([#4754](https://github.com/hashicorp/terraform-provider-google-beta/pull/4754))
* **New Resource:** `google_bigquery_datapolicy_data_policy_iam_policy` ([#4754](https://github.com/hashicorp/terraform-provider-google-beta/pull/4754))
* **New Resource:** `google_org_policy_custom_constraint` ([#4741](https://github.com/hashicorp/terraform-provider-google-beta/pull/4741))
* **New Resource:** `google_vertex_ai_featurestore_entitytype_feature` ([#4736](https://github.com/hashicorp/terraform-provider-google-beta/pull/4736))

IMPROVEMENTS:
* bigqueryreservation: added `concurrency` and `multiRegionAuxiliary` to `google_bigquery_reservation` ([#4739](https://github.com/hashicorp/terraform-provider-google-beta/pull/4739))
* bigtable: added additional retry GC policy operations with a longer poll interval to avoid quota issues ([#4750](https://github.com/hashicorp/terraform-provider-google-beta/pull/4750))
* bigtable: improved error messaging ([#4746](https://github.com/hashicorp/terraform-provider-google-beta/pull/4746))
* compute: added support for `compression_mode` field in `google_compute_backend_bucket` and `google_compute_backend_service` resource ([#4733](https://github.com/hashicorp/terraform-provider-google-beta/pull/4733))
* dataflow : added support of `labels` to resource `google_dataflow_flextemplate_job` ([#4748](https://github.com/hashicorp/terraform-provider-google-beta/pull/4748))
* datastream: added field `bigquery_profile` to `google_datastream_connection_profile` ([#4742](https://github.com/hashicorp/terraform-provider-google-beta/pull/4742))
* dns: added general field `cloud_logging_config` to `google_dns_managed_zone` ([#4734](https://github.com/hashicorp/terraform-provider-google-beta/pull/4734))
* metastore: added bigquery support for `google_dataproc_metastore_service` ([#4753](https://github.com/hashicorp/terraform-provider-google-beta/pull/4753))
* storage: added `custom_placement_config` field to `google_storage_bucket` resource to support custom dual-region GCS buckets ([#4752](https://github.com/hashicorp/terraform-provider-google-beta/pull/4752))
* sql: added  `password_policy` field to `google_sql_user` resource ([#4730](https://github.com/hashicorp/terraform-provider-google-beta/pull/4730))

BUG FIXES:
* storage: fixed a bug where user specified labels get overwritten by Dataplex auto generated labels ([#4743](https://github.com/hashicorp/terraform-provider-google-beta/pull/4743))
* storagetransfer: fixed a crash in `google_storagetransfer_job` refreshes when `transfer_schedule` was empty ([#4745](https://github.com/hashicorp/terraform-provider-google-beta/pull/4745))


## 4.39.0 (October 3, 2022)

FEATURES:
* **New Data Source:** `google_artifact_registry_repository` ([#4714](https://github.com/hashicorp/terraform-provider-google-beta/pull/4714))
* **New Resource:** `google_identity_platform_config` ([#4729](https://github.com/hashicorp/terraform-provider-google-beta/pull/4729))

IMPROVEMENTS:
* certificatemanager: added public/private PEM fields `pem_certificate` / `pem_private_key` and deprecated `certificate_pem` / `private_key_pem` ([#4728](https://github.com/hashicorp/terraform-provider-google-beta/pull/4728))
* clouddeploy: added `serial_pipeline.stages.strategy` field to `google_clouddeploy_delivery_pipeline` ([#4707](https://github.com/hashicorp/terraform-provider-google-beta/pull/4707))
* container: added `notification_config.pubsub.filter` field to `google_container_cluster` ([#4718](https://github.com/hashicorp/terraform-provider-google-beta/pull/4718))
* eventarc: added `channels` and `conditions` fields to `google_eventarc_trigger` ([#4707](https://github.com/hashicorp/terraform-provider-google-beta/pull/4707))
* healthcare: added `notification_configs ` field to `google_healthcare_fhir_store` resource ([#4720](https://github.com/hashicorp/terraform-provider-google-beta/pull/4720))
* iap: added ability to import `google_iap_brand` using ID using {{project}}/{{brand_id}} format ([#4712](https://github.com/hashicorp/terraform-provider-google-beta/pull/4712))
* secretmanager: added output field 'version' to resource 'secret_manager_secret_version' ([#4724](https://github.com/hashicorp/terraform-provider-google-beta/pull/4724))
* sql: added `maintenance_version` and `available_maintenance_versions` fields to `google_sql_database_instance` resource ([#4725](https://github.com/hashicorp/terraform-provider-google-beta/pull/4725))
* storagetransfer: added `notification_config` field to `google_storage_transfer_job` resource ([#4709](https://github.com/hashicorp/terraform-provider-google-beta/pull/4709))
* tags: added `purpose` and `purpose_data` properties to `google_tags_tag_key` ([#4721](https://github.com/hashicorp/terraform-provider-google-beta/pull/4721))

BUG FIXES:
* bigquery: fixed a bug where `allow_quoted_newlines` and `allow_jagged_rows` could not be set to false on `google_bigquery_table` ([#4711](https://github.com/hashicorp/terraform-provider-google-beta/pull/4711))
* cloudfunction: fixed inability to update `docker_repository` and `kms_key_name` on `google_cloudfunctions_function` ([#4727](https://github.com/hashicorp/terraform-provider-google-beta/pull/4727))
* compute: fixed inability to manage Cloud Armor `adaptive_protection_config` on `google_compute_security_policy` ([#4726](https://github.com/hashicorp/terraform-provider-google-beta/pull/4726))
* container: fixed a bug where upgrading provider version breaks on `node_pool_auto_config` or `node_pool_defaults` ([#4706](https://github.com/hashicorp/terraform-provider-google-beta/pull/4706))
* iam: fixed diffs between `policy_data` from `google_iam_policy` data source and policy data in API responses ([#4722](https://github.com/hashicorp/terraform-provider-google-beta/pull/4722))
* iam: fixed permadiff resulting from empty fields being sent in requests to set conditional IAM policies ([#4723](https://github.com/hashicorp/terraform-provider-google-beta/pull/4723))
* secretmanager: fixed a bug where `google_secret_manager_secret_version` that was destroyed outside of Terraform would not be recreated on apply ([#4719](https://github.com/hashicorp/terraform-provider-google-beta/pull/4719))
* storagetransfer: fixed a crash in `google_storagetransfer_job` when `transfer_schedule` is empty ([#4745](https://github.com/hashicorp/terraform-provider-google-beta/pull/4745))
 
## 4.38.0 (September 26, 2022)

FEATURES:
* **New Data Source:** `google_vpc_access_connector` ([#4693](https://github.com/hashicorp/terraform-provider-google-beta/pull/4693))
* **New Resource:** `google_datastream_private_connection` ([#4691](https://github.com/hashicorp/terraform-provider-google-beta/pull/4691))

IMPROVEMENTS:
* appengine: Added `egress_setting` for field `vpc_access_connector` to `google_app_engine_standard_app_version` ([#4701](https://github.com/hashicorp/terraform-provider-google-beta/pull/4701))
* bigquery: added `json_extension` field to the `load` block of `google_bigquery_job` resource ([#4699](https://github.com/hashicorp/terraform-provider-google-beta/pull/4699))
* cloudfunctions: Added `build_worker_pool` to `google_cloudfunctions_function` ([#4696](https://github.com/hashicorp/terraform-provider-google-beta/pull/4696))
* compute: added `json_custom_config` field to `google_compute_security_policy` resource ([#4703](https://github.com/hashicorp/terraform-provider-google-beta/pull/4703))
* redis: Added `persistence_config` field to the `google_redis_instance` resource. ([#4688](https://github.com/hashicorp/terraform-provider-google-beta/pull/4688))
* storage: added support for `overwriteWhen` field to `transfer_options` in `google_storage_transfer_job` resource ([#4690](https://github.com/hashicorp/terraform-provider-google-beta/pull/4690))

BUG FIXES:
* bigtable: added drift detection on `gc_rules` for `google_bigtable_gc_policy` ([#4687](https://github.com/hashicorp/terraform-provider-google-beta/pull/4687))
* compute: fixed the inability to update `most_disruptive_allowed_action` for both `google_compute_per_instance_config` and `google_compute_region_per_instance_config` ([#4685](https://github.com/hashicorp/terraform-provider-google-beta/pull/4685))
* container: fixed allow passing empty list to monitoring_config and logging_config in `google_container_cluster` ([#4700](https://github.com/hashicorp/terraform-provider-google-beta/pull/4700))
* sql: fixed a bug causing a perma-diff on `disk_type` due to API values being downcased ([#4686](https://github.com/hashicorp/terraform-provider-google-beta/pull/4686))
* storage: fixed the inability to set 0 for `lifecycle_rule.condition.age` on `google_storage_bucket` ([#4698](https://github.com/hashicorp/terraform-provider-google-beta/pull/4698))

## 4.37.0 (September 19, 2022)

FEATURES:
* **New Resource:** `google_apigee_nat_address` ([#4676](https://github.com/hashicorp/terraform-provider-google-beta/pull/4676))
* **New Resource:** `google_dialogflow_cx_webhook` ([#4667](https://github.com/hashicorp/terraform-provider-google-beta/pull/4667))
* **New Resource:** `google_filestore_snapshot` ([#4661](https://github.com/hashicorp/terraform-provider-google-beta/pull/4661))

IMPROVEMENTS:
* apigee: added read-only field `connection_state` to `google_apigee_endpoint_attachment` ([#4668](https://github.com/hashicorp/terraform-provider-google-beta/pull/4668))
* bigtable: added support for `autoscaling_config.storage_target` to `google_bigtable_instance` ([#4671](https://github.com/hashicorp/terraform-provider-google-beta/pull/4671))
* cloudbuild: added support for `BITBUCKET` option to `git_source.repo_type` in `google_cloudbuild_trigger` ([#4679](https://github.com/hashicorp/terraform-provider-google-beta/pull/4679))
* dns: added in validation for trailing dot at end of DNS record name ([#4674](https://github.com/hashicorp/terraform-provider-google-beta/pull/4674))
* project: added validation for field `project_id` in `google_project` datasource. ([#4684](https://github.com/hashicorp/terraform-provider-google-beta/pull/4684))
* serviceaccount: added `expires_in` attribute for generating `exp` claim to `google_service_account_jwt` datasource. ([#4677](https://github.com/hashicorp/terraform-provider-google-beta/pull/4677))

BUG FIXES:
* notebooks: fixed perma-diff in `google_notebooks_instance` ([#4664](https://github.com/hashicorp/terraform-provider-google-beta/pull/4664))
* privateca: fixed an issue that blocked subordinate CA data sources when `state` was not `AWAITING_USER_ACTIVATION` ([#4672](https://github.com/hashicorp/terraform-provider-google-beta/pull/4672))
* storage: fixed permdiff on the field `versioning` of `google_storage_bucket` ([#4665](https://github.com/hashicorp/terraform-provider-google-beta/pull/4665))

## 4.36.0 (September 12, 2022)

FEATURES:
* **New Resource:** `google_bigquery_analytics_hub_data_exchange_iam_binding` ([#4656](https://github.com/hashicorp/terraform-provider-google-beta/pull/4656))
* **New Resource:** `google_bigquery_analytics_hub_data_exchange_iam_member` ([#4656](https://github.com/hashicorp/terraform-provider-google-beta/pull/4656))
* **New Resource:** `google_bigquery_analytics_hub_data_exchange_iam_policy` ([#4656](https://github.com/hashicorp/terraform-provider-google-beta/pull/4656))
* **New Resource:** `google_bigquery_analytics_hub_data_exchange` ([#4656](https://github.com/hashicorp/terraform-provider-google-beta/pull/4656))
* **New Resource:** `google_datastream_connection_profile` ([#4657](https://github.com/hashicorp/terraform-provider-google-beta/pull/4657))

IMPROVEMENTS:
* appengine: added field `service_account` to `google_app_engine_flexible_app_version` ([#4653](https://github.com/hashicorp/terraform-provider-google-beta/pull/4653))
* bigtable: increased timeout in `google_bigtable_table` creation. ([#4655](https://github.com/hashicorp/terraform-provider-google-beta/pull/4655))
* cloudbuild: added `location` field to `google_cloudbuild_trigger` resource ([#4646](https://github.com/hashicorp/terraform-provider-google-beta/pull/4646))
* compute: added `certificate_map` to `compute_target_ssl_proxy` resource ([#4654](https://github.com/hashicorp/terraform-provider-google-beta/pull/4654))
* compute: added field `chain_name` to `google_compute_resource_policy.snapshot_properties` ([#4660](https://github.com/hashicorp/terraform-provider-google-beta/pull/4660))
* compute: added field `chain_name` to resource `google_compute_snapshot` ([#4660](https://github.com/hashicorp/terraform-provider-google-beta/pull/4660))
* container: added `autoscaling.total_min_node_count`, `autoscaling.total_max_node_count`, and `autoscaling.location_policy` to `google_container_cluster.node_pool` ([#4649](https://github.com/hashicorp/terraform-provider-google-beta/pull/4649))
* container: added `autoscaling.total_min_node_count`, `autoscaling.total_max_node_count`, and `autoscaling.location_policy` to `google_container_node_pool` resource ([#4649](https://github.com/hashicorp/terraform-provider-google-beta/pull/4649))
* container: added field `node_pool_defaults` to `resource_container_cluster`. ([#4648](https://github.com/hashicorp/terraform-provider-google-beta/pull/4648))
* dataproc: added option `shielded_instance_config` to resource `google_dataproc_workflow_template`. ([#4647](https://github.com/hashicorp/terraform-provider-google-beta/pull/4647))
* metastore: extended default timeouts for `google_dataproc_metastore_service` from 40m to 60m ([#4652](https://github.com/hashicorp/terraform-provider-google-beta/pull/4652))
* pubsub: made `google_pubsub_subscription.enable_exactly_once_delivery` mutable so that it updates subscription without recreation. ([#4645](https://github.com/hashicorp/terraform-provider-google-beta/pull/4645))

## 4.35.0 (September 6, 2022)

IMPROVEMENTS:
* apigee: added support for `nodeConfig` in `google_apigee_environment` ([#4632](https://github.com/hashicorp/terraform-provider-google-beta/pull/4632))
* apigee: added a `properties` field to `google_apigee_organization` ([#4644](https://github.com/hashicorp/terraform-provider-google-beta/pull/4644))
* cloudfunctions2: added `secret_environment_variables` and `secret_volumes` to `google_cloudfunctions2_function` ([#4641](https://github.com/hashicorp/terraform-provider-google-beta/pull/4641))
* compute: added support for param `visible_core_count` in `google_compute_instance` and `google_compute_instance_template` under `advanced_machine_features` ([#4635](https://github.com/hashicorp/terraform-provider-google-beta/pull/4635))
* compute: added support documentation links to error messages for certain Compute Operation errors. ([#4642](https://github.com/hashicorp/terraform-provider-google-beta/pull/4642))
* container: added `service_external_ips_config` support to `cluster_container` resource. ([#4639](https://github.com/hashicorp/terraform-provider-google-beta/pull/4639))
* container: added `enable_cost_allocation` to `google_container_cluster` ([#4640](https://github.com/hashicorp/terraform-provider-google-beta/pull/4640))
* dns: added `behavior` field to `google_dns_response_policy_rule` resource ([#4637](https://github.com/hashicorp/terraform-provider-google-beta/pull/4637))
* monitoring: added `force_delete` field to `google_monitoring_notification_channel` resource ([#4638](https://github.com/hashicorp/terraform-provider-google-beta/pull/4638))
* pubsub: made `enable_exactly_once_delivery` mutable so that it updates subscription in-place and avoids recreation of the subscription. ([#4645](https://github.com/hashicorp/terraform-provider-google-beta/pull/4645))
* vertexai: added `encryption_spec` field to `google_vertex_ai_featurestore` resource (beta) ([#4643](https://github.com/hashicorp/terraform-provider-google-beta/pull/4643))

BUG FIXES:
* compute: fixed the `id` format of the data source `google_compute_instance` ([#4636](https://github.com/hashicorp/terraform-provider-google-beta/pull/4636))

## 4.34.0 (August 29, 2022)
NOTES:
* updated Bigtable go client version from 1.13 to 1.16. ([#4613](https://github.com/hashicorp/terraform-provider-google-beta/pull/4613))

IMPROVEMENTS:
* apigee: added support for specifying retention when deleting `google_apigee_organization` ([#4604](https://github.com/hashicorp/terraform-provider-google-beta/pull/4604))
* appengine: added `app_engine_apis` field to `google_app_engine_standard_app_version` resource ([#4607](https://github.com/hashicorp/terraform-provider-google-beta/pull/4607))
* compute: improved error messaging for compute errors ([#4602](https://github.com/hashicorp/terraform-provider-google-beta/pull/4602))
* container: added general field `reservation_affinity` to `google_container_node_pool` ([#4622](https://github.com/hashicorp/terraform-provider-google-beta/pull/4622))
* container: added field `auto_provisioning_network_tags` to `google_container_cluster` (beta) ([#4611](https://github.com/hashicorp/terraform-provider-google-beta/pull/4611))
* sql: added support for major version upgrade to `google_sql_database_instance ` resource  ([#4606](https://github.com/hashicorp/terraform-provider-google-beta/pull/4606))

BUG FIXES:
* bigtable: fixed comparing column family name when reading a GC policy. ([#4624](https://github.com/hashicorp/terraform-provider-google-beta/pull/4624))
* bigtable: passed `isTopeLevel` in getGCPolicyFromJSON() instead of hardcoding it to true. ([#4615](https://github.com/hashicorp/terraform-provider-google-beta/pull/4615))
* cloud iam: made `denial_condition` optional on `google_iam_deny_policy` ([#4617](https://github.com/hashicorp/terraform-provider-google-beta/pull/4617))

## 4.33.0 (August 22, 2022)
IMPROVEMENTS:
* container: added update support for `authenticator_groups_config` in `google_container_cluster` ([#4591](https://github.com/hashicorp/terraform-provider-google-beta/pull/4591))
* dataflow: added ability to import `google_dataflow_job` ([#4595](https://github.com/hashicorp/terraform-provider-google-beta/pull/4595))
* dns: added `managed_zone_id` attribute to `google_dns_managed_zone` data source ([#4593](https://github.com/hashicorp/terraform-provider-google-beta/pull/4593))
* metastore: added `metadata_integration` and `hive_metastore_config.auxiliary_versions` fields to `google_dataproc_metastore_service` resource ([#4598](https://github.com/hashicorp/terraform-provider-google-beta/pull/4598))
* monitoring: added `accepted_response_status_codes` to `google_monitoring_uptime_check_config` ([#4594](https://github.com/hashicorp/terraform-provider-google-beta/pull/4594))
* sql: added `password_validation_policy` field to `google_cloud_sql` resource ([#4597](https://github.com/hashicorp/terraform-provider-google-beta/pull/4597))

BUG FIXES:
* bigquery: removed force replacement for `display_name` on `google_bigquery_data_transfer_config` ([#4592](https://github.com/hashicorp/terraform-provider-google-beta/pull/4592))
* compute: fixed permadiff for `instance_termination_action` in `google_compute_instance_template` ([#4590](https://github.com/hashicorp/terraform-provider-google-beta/pull/4590))

## 4.32.0 (August 15, 2022)

NOTES:
* Updated to Golang 1.18 ([#4564](https://github.com/hashicorp/terraform-provider-google-beta/pull/4564))

FEATURES:
* **New Resource:** `google_dataplex_asset` ([#4543](https://github.com/hashicorp/terraform-provider-google-beta/pull/4543))
* **New Resource:** `google_gke_hub_membership_iam_binding` ([#4583](https://github.com/hashicorp/terraform-provider-google-beta/pull/4583))
* **New Resource:** `google_gke_hub_membership_iam_member` ([#4583](https://github.com/hashicorp/terraform-provider-google-beta/pull/4583))
* **New Resource:** `google_gke_hub_membership_iam_policy` ([#4583](https://github.com/hashicorp/terraform-provider-google-beta/pull/4583))

IMPROVEMENTS:
* certificatemanager: added `state`, `authorization_attempt_info` and `provisioning_issue` output fields to `google_certificate_manager_certificate` ([#4548](https://github.com/hashicorp/terraform-provider-google-beta/pull/4548))
* cloudfunctions2: added field `event_filters` to resource `google_cloudfunctions2_function` ([#4547](https://github.com/hashicorp/terraform-provider-google-beta/pull/4547))
* compute: added `certificate_map` to `compute_target_https_proxy` resource ([#4550](https://github.com/hashicorp/terraform-provider-google-beta/pull/4550))
* compute: added validation for name field on `google_compute_network` ([#4579](https://github.com/hashicorp/terraform-provider-google-beta/pull/4579))
* compute: made `port` optional in `google_compute_network_endpoint` to allow network endpoints to be associated with `GCE_VM_IP` network endpoint groups ([#4575](https://github.com/hashicorp/terraform-provider-google-beta/pull/4575))
* container: added support for additional values `APISERVER`, `CONTROLLER_MANAGER`, and `SCHEDULER` in `google_container_cluster.monitoring_config` ([#4565](https://github.com/hashicorp/terraform-provider-google-beta/pull/4565))
* gkehub: added `monitoring` and `mutation_enabled` fields to resource `feature_membership` ([#4572](https://github.com/hashicorp/terraform-provider-google-beta/pull/4572))
* gkehub: added better support for import for `google_gke_hub_membership` ([#4542](https://github.com/hashicorp/terraform-provider-google-beta/pull/4542))
* pubsub: added `bigquery_config` to `google_pubsub_subscription` ([#4545](https://github.com/hashicorp/terraform-provider-google-beta/pull/4545))
* scheduler: added `paused` field to `google_cloud_scheduler_job` ([#4535](https://github.com/hashicorp/terraform-provider-google-beta/pull/4535))
* scheduler: added `state` output field to `google_cloud_scheduler_job` ([#4535](https://github.com/hashicorp/terraform-provider-google-beta/pull/4535))

BUG FIXES:
* apigee: fixed an issue where `google_apigee_instance` creation would fail due to multiple concurrent instances ([#4584](https://github.com/hashicorp/terraform-provider-google-beta/pull/4584))
* billingbudget: fixed a bug where `google_billing_budget.budget_filter.services` was not updating. ([#4577](https://github.com/hashicorp/terraform-provider-google-beta/pull/4577))
* compute: fixed perma-diff on `google_compute_disk` for new arm64 images ([#4533](https://github.com/hashicorp/terraform-provider-google-beta/pull/4533))
* dataflow: fixed bug where permadiff would show on `google_dataflow_job.additional_experiments` ([#4576](https://github.com/hashicorp/terraform-provider-google-beta/pull/4576))
* storage: fixed a bug in `google_storage_bucket` where `name` was incorrectly validated. ([#4566](https://github.com/hashicorp/terraform-provider-google-beta/pull/4566))

## 4.31.0 (Aug 1, 2022)

FEATURES:
* **New Resource:** `google_dataplex_zone` ([#4511](https://github.com/hashicorp/terraform-provider-google-beta/pull/4511))

IMPROVEMENTS:
* bucket: added support for `matches_prefix` and `matches_suffix` in `condition` of a `lifecycle_rule` in  `google_storage_bucket` ([#4527](https://github.com/hashicorp/terraform-provider-google-beta/pull/4527))
* compute: added `network` and `subnetwork` fields to `google_compute_region_network_endpoint_group` for PSC. ([#4528](https://github.com/hashicorp/terraform-provider-google-beta/pull/4528))
* container: added field `boot_disk_kms_key` to `auto_provisioning_defaults` in `google_container_cluster` ([#4524](https://github.com/hashicorp/terraform-provider-google-beta/pull/4524))
* notebooks: added `bootDiskType` support for `PD_EXTREME` in `google_notebooks_instance` ([#4530](https://github.com/hashicorp/terraform-provider-google-beta/pull/4530))
* notebooks: added `softwareConfig.upgradeable`, `softwareConfig.postStartupScriptBehavior`, `softwareConfig.kernels` in `google_notebooks_runtime` ([#4530](https://github.com/hashicorp/terraform-provider-google-beta/pull/4530))
* storage: added name validation for `google_storage_bucket` ([#4532](https://github.com/hashicorp/terraform-provider-google-beta/pull/4532))

BUG FIXES:
* compute: fixed perma-diff on `google_compute_disk` for new arm64 images ([#4533](https://github.com/hashicorp/terraform-provider-google-beta/pull/4533))
* dns: fixed a bug where `google_dns_record_set` would create an inconsistent plan when using interpolated values in `rrdatas` ([#4515](https://github.com/hashicorp/terraform-provider-google-beta/pull/4515))
* kms: fixed setting of resource id post-import for `google_kms_crypto_key` ([#4520](https://github.com/hashicorp/terraform-provider-google-beta/pull/4520))
* provider: fixed a bug where user-agent was showing "dev" rather than the provider version ([#4509](https://github.com/hashicorp/terraform-provider-google-beta/pull/4509))

## 4.30.0 (July 25, 2022)

FEATURES:
* **New Data Source:** `google_service_account_jwt` ([#4489](https://github.com/hashicorp/terraform-provider-google-beta/pull/4489))
* **New Resource:** `google_certificate_map_entry` ([#4501](https://github.com/hashicorp/terraform-provider-google-beta/pull/4501))
* **New Resource:** `google_certificate_map` ([#4501](https://github.com/hashicorp/terraform-provider-google-beta/pull/4501))
* **New Resource:** `google_compute_backend_bucket_iam_binding` ([#4484](https://github.com/hashicorp/terraform-provider-google-beta/pull/4484))
* **New Resource:** `google_compute_backend_bucket_iam_member` ([#4484](https://github.com/hashicorp/terraform-provider-google-beta/pull/4484))
* **New Resource:** `google_compute_backend_bucket_iam_policy` ([#4484](https://github.com/hashicorp/terraform-provider-google-beta/pull/4484))
* **New Resource:** `google_dataproc_metastore_federation` ([#4482](https://github.com/hashicorp/terraform-provider-google-beta/pull/4482))
* **New Resource:** `google_dataproc_metastore_federation_iam_binding` ([#4482](https://github.com/hashicorp/terraform-provider-google-beta/pull/4482))
* **New Resource:** `google_dataproc_metastore_federation_iam_member` ([#4482](https://github.com/hashicorp/terraform-provider-google-beta/pull/4482))
* **New Resource:** `google_dataproc_metastore_federation_iam_policy` ([#4482](https://github.com/hashicorp/terraform-provider-google-beta/pull/4482))

IMPROVEMENTS:
* billingbudget: made `thresholdRules` optional in `google_billing_budget` ([#4480](https://github.com/hashicorp/terraform-provider-google-beta/pull/4480))
* compute: added `instance_termination_action` field to `google_compute_instance_template` resource to support Spot VM termination action ([#4488](https://github.com/hashicorp/terraform-provider-google-beta/pull/4488))
* compute: added `instance_termination_action` field to `google_compute_instance` resource to support Spot VM termination action ([#4488](https://github.com/hashicorp/terraform-provider-google-beta/pull/4488))
* compute: added `request_coalescing` and `bypass_cache_on_request_headers` fields to `compute_backend_bucket` ([#4484](https://github.com/hashicorp/terraform-provider-google-beta/pull/4484))
* compute: added field `all_instances_config` to `google_compute_instance_group_manager` and `google_compute_region_instance_group_manager` ([#4506](https://github.com/hashicorp/terraform-provider-google-beta/pull/4506))
* compute: added support for `esp` protocol in `google_compute_packet_mirroring.filters.ip_protocols` ([#4496](https://github.com/hashicorp/terraform-provider-google-beta/pull/4496))
* monitoring: added `evaluation_missing_data` field to `google_monitoring_alert_policy` ([#4502](https://github.com/hashicorp/terraform-provider-google-beta/pull/4502))
* notebooks: added field `reserved_ip_range` to `google_notebooks_runtime` ([#4492](https://github.com/hashicorp/terraform-provider-google-beta/pull/4492))

BUG FIXES:
* bigtable: fixed an incorrect diff when adding two or more clusters ([#4490](https://github.com/hashicorp/terraform-provider-google-beta/pull/4490))
* compute: allowed properly updating `adaptive_protection_config` in `compute_security_policy` ([#4478](https://github.com/hashicorp/terraform-provider-google-beta/pull/4478))
* notebooks: fixed a bug where`google_notebooks_runtime` can't be updated ([#4492](https://github.com/hashicorp/terraform-provider-google-beta/pull/4492))
* sql: fixed an issue in `google_sql_database_instance` where updates would fail because of the `collation` field ([#4505](https://github.com/hashicorp/terraform-provider-google-beta/pull/4505))

## 4.29.0 (July 18, 2022)

FEATURES:
* **New Resource:** `google_cloudiot_registry_iam_binding` ([#4452](https://github.com/hashicorp/terraform-provider-google-beta/pull/4452))
* **New Resource:** `google_cloudiot_registry_iam_member` ([#4452](https://github.com/hashicorp/terraform-provider-google-beta/pull/4452))
* **New Resource:** `google_cloudiot_registry_iam_policy` ([#4452](https://github.com/hashicorp/terraform-provider-google-beta/pull/4452))
* **New Resource:** `google_compute_snapshot_iam_binding` ([#4445](https://github.com/hashicorp/terraform-provider-google-beta/pull/4445))
* **New Resource:** `google_compute_snapshot_iam_member` ([#4445](https://github.com/hashicorp/terraform-provider-google-beta/pull/4445))
* **New Resource:** `google_compute_snapshot_iam_policy` ([#4445](https://github.com/hashicorp/terraform-provider-google-beta/pull/4445))

IMPROVEMENTS:
* container: added `binauthz_evaluation_mode` field to `resource_container_cluster`. ([#4451](https://github.com/hashicorp/terraform-provider-google-beta/pull/4451))
* kms: added support for MAC value in `google_kms_crypto_key.purpose` ([#4458](https://github.com/hashicorp/terraform-provider-google-beta/pull/4458))
* metastore: added `databaseType`, `releaseChannel`, and `hiveMetastoreConfig.endpointProtocol` arguments ([#4443](https://github.com/hashicorp/terraform-provider-google-beta/pull/4443))

BUG FIXES:
* bigquery: fixed case-sensitivity for `user_by_email` and `group_by_email` on `google_bigquery_dataset_access` ([#4446](https://github.com/hashicorp/terraform-provider-google-beta/pull/4446))
* clouddeploy: fixed permadiff on `execution_configs` in `google_clouddeploy_target` resource ([#4450](https://github.com/hashicorp/terraform-provider-google/pull/4450))
* cloudscheduler: fixed a diff on the last slash of uri on `google_cloud_scheduler_job` ([#4444](https://github.com/hashicorp/terraform-provider-google-beta/pull/4444))
* compute: fixed force recreation on `provisioned_iops` of `google_compute_disk` ([#4464](https://github.com/hashicorp/terraform-provider-google-beta/pull/4464))
* compute: fixed missing `network_interface.0.ipv6_access_config.0.external_ipv6` output on `google_compute_instance` ([#4470](https://github.com/hashicorp/terraform-provider-google-beta/pull/4470))
* documentai: fixed a bug where eu region could not be utilized for documentai resources ([#4472](https://github.com/hashicorp/terraform-provider-google-beta/pull/4472))
* gkehub: fixed a bug where `issuer` can't be updated on `google_gke_hub_membership` ([#4471](https://github.com/hashicorp/terraform-provider-google-beta/pull/4471))

## 4.28.0 (July 11, 2022)

FEATURES:
* **New Resource:** google_bigquery_connection_iam_binding ([#4437](https://github.com/hashicorp/terraform-provider-google-beta/pull/4437))
* **New Resource:** google_bigquery_connection_iam_member ([#4437](https://github.com/hashicorp/terraform-provider-google-beta/pull/4437))
* **New Resource:** google_bigquery_connection_iam_policy ([#4437](https://github.com/hashicorp/terraform-provider-google-beta/pull/4437))
* **New Resource:** google_cloud_tasks_queue_iam_binding ([#4427](https://github.com/hashicorp/terraform-provider-google-beta/pull/4427))
* **New Resource:** google_cloud_tasks_queue_iam_member ([#4427](https://github.com/hashicorp/terraform-provider-google-beta/pull/4427))
* **New Resource:** google_cloud_tasks_queue_iam_policy ([#4427](https://github.com/hashicorp/terraform-provider-google-beta/pull/4427))
* **New Resource:** google_dataproc_autoscaling_policy_iam_binding ([#4441](https://github.com/hashicorp/terraform-provider-google-beta/pull/4441))
* **New Resource:** google_dataproc_autoscaling_policy_iam_member ([#4441](https://github.com/hashicorp/terraform-provider-google-beta/pull/4441))
* **New Resource:** google_dataproc_autoscaling_policy_iam_policy ([#4441](https://github.com/hashicorp/terraform-provider-google-beta/pull/4441))
* **New Resource:** google_dataproc_metastore_service_iam_binding ([#4416](https://github.com/hashicorp/terraform-provider-google-beta/pull/4416))
* **New Resource:** google_dataproc_metastore_service_iam_member ([#4416](https://github.com/hashicorp/terraform-provider-google-beta/pull/4416))
* **New Resource:** google_dataproc_metastore_service_iam_policy ([#4416](https://github.com/hashicorp/terraform-provider-google-beta/pull/4416))

IMPROVEMENTS:
* bigquery: fixed a permadiff in `google_bigquery_job.query. destination_table` ([#4401](https://github.com/hashicorp/terraform-provider-google-beta/pull/4401))
* billing: added `calendar_period` and `custom_period` fields to `google_billing_budget` ([#4429](https://github.com/hashicorp/terraform-provider-google-beta/pull/4429))
* cloudsql: added attribute `project` to data source `google_sql_backup_run` ([#4402](https://github.com/hashicorp/terraform-provider-google-beta/pull/4402))
* composer: added CMEK, PUPI and IP_masq_agent support for Composer 2 in `google_composer_environment` resource ([#4430](https://github.com/hashicorp/terraform-provider-google-beta/pull/4430))
* compute: added `max_ports_per_vm` field to `google_compute_router_nat` resource ([#4400](https://github.com/hashicorp/terraform-provider-google-beta/pull/4400))
* compute: added `GCE_VM_IP` support to `google_compute_network_endpoint_group` resource. ([#4434](https://github.com/hashicorp/terraform-provider-google-beta/pull/4434))
* privateca: added support to subordinate CA activation ([#4422](https://github.com/hashicorp/terraform-provider-google-beta/pull/4422))
* redis: added CMEK key field `customer_managed_key` in `google_redis_instance ` ([#4435](https://github.com/hashicorp/terraform-provider-google-beta/pull/4435))
* spanner: added field `version_retention_period` to `google_spanner_database` resource ([#4424](https://github.com/hashicorp/terraform-provider-google-beta/pull/4424))
* sql: added `settings.location_preference.secondary_zone` field in `google_sql_database_instance` ([#4433](https://github.com/hashicorp/terraform-provider-google-beta/pull/4433))
* sql: added `sql_server_audit_config` field in `google_sql_database_instance` ([#4403](https://github.com/hashicorp/terraform-provider-google-beta/pull/4403))

BUG FIXES:
* composer: fixed a problem with updating Cloud Composer's `scheduler_count` field (https://github.com/hashicorp/terraform-provider-google/issues/11940) ([#4408](https://github.com/hashicorp/terraform-provider-google-beta/pull/4408))
* composer: fixed permadiff on `private_environment_config.cloud_composer_connection_subnetwork` ([#4411](https://github.com/hashicorp/terraform-provider-google-beta/pull/4411))
* container: fixed an issue where `node_config.min_cpu_platform` could cause a perma-diff in `google_container_cluster` ([#4426](https://github.com/hashicorp/terraform-provider-google-beta/pull/4426))
* filestore: fixed a case where `google_filestore_instance.networks.network` would incorrectly see a diff between state and config when the network `id` format was used ([#4431](https://github.com/hashicorp/terraform-provider-google-beta/pull/4431))
* serviceusage: fixed an issue where `google_project_service_identity` didn't handle service identities without emails correctly ([#4432](https://github.com/hashicorp/terraform-provider-google-beta/pull/4432))


## 4.27.0 (June 27, 2022)

IMPROVEMENTS:
* clouddeploy: added `suspend` field to `google_clouddeploy_delivery_pipeline` resource ([#4394](https://github.com/hashicorp/terraform-provider-google-beta/pull/4394))
* compute: added maxPortsPerVm field to `google_compute_router_nat` resource ([#4400](https://github.com/hashicorp/terraform-provider-google-beta/pull/4400))
* compute: added `psc_connection_id` and `psc_connection_status` output fields to `google_compute_forwarding_rule` and `google_compute_global_forwarding_rule` resources ([#4392](https://github.com/hashicorp/terraform-provider-google-beta/pull/4392))
* container: added `tpu_config` to `google_container_cluster` (beta only) ([#4390](https://github.com/hashicorp/terraform-provider-google-beta/pull/4390))
* containeraws: made `config.instance_type` field updatable in `google_container_aws_node_pool` ([#4392](https://github.com/hashicorp/terraform-provider-google-beta/pull/4392))

BUG FIXES:
* compute: fixed default handling for `enable_dynamic_port_allocation ` to be managed by the api ([#4391](https://github.com/hashicorp/terraform-provider-google-beta/pull/4391))
* vertexai: Fixed a bug where terraform crashes when `force_destroy` is set in `google_vertex_ai_featurestore` resource ([#4398](https://github.com/hashicorp/terraform-provider-google-beta/pull/4398))

## 4.26.0 (June 21, 2022)

FEATURES:
* **New Resource:** `google_cloudfunctions2_function_iam_binding` ([#4377](https://github.com/hashicorp/terraform-provider-google-beta/pull/4377))
* **New Resource:** `google_cloudfunctions2_function_iam_member` ([#4377](https://github.com/hashicorp/terraform-provider-google-beta/pull/4377))
* **New Resource:** `google_cloudfunctions2_function_iam_policy` ([#4377](https://github.com/hashicorp/terraform-provider-google-beta/pull/4377))
* **New Resource:** `google_compute_region_ssl_policy` ([#4376](https://github.com/hashicorp/terraform-provider-google-beta/pull/4376))
* **New Resource:** `google_documentai_processor` ([#4389](https://github.com/hashicorp/terraform-provider-google-beta/pull/4389))
* **New Resource:** `google_documentai_processor_default_version` ([#4389](https://github.com/hashicorp/terraform-provider-google-beta/pull/4389))

IMPROVEMENTS:
* accesscontextmanager: Added `external_resources` to `egress_to` in `google_access_context_manager_service_perimeter` and `google_access_context_manager_service_perimeters` resource ([#4378](https://github.com/hashicorp/terraform-provider-google-beta/pull/4378))
* apigateway: Added `grpc_services` and `managed_service_configs` to `google_api_gateway_api_config` ([#4388](https://github.com/hashicorp/terraform-provider-google-beta/pull/4388))
* cloudbuild: Added `include_build_logs` to `google_cloudbuild_trigger` ([#4380](https://github.com/hashicorp/terraform-provider-google-beta/pull/4380))
* compute: Added `ssl_policy` field to `google_compute_region_target_https_proxy` ([#4376](https://github.com/hashicorp/terraform-provider-google-beta/pull/4376))
* container: Added `managed_prometheus` to `monitoring_config` in `google_container_cluster` ([#4373](https://github.com/hashicorp/terraform-provider-google-beta/pull/4373))
* container: Added `tpu_config` to `google_container_cluster` ([#4390](https://github.com/hashicorp/terraform-provider-google-beta/pull/4390))

BUG FIXES:
* dns: Fixed a bug where `google_dns_record_set` resource can not be changed from default routing to Geo routing policy. ([#4383](https://github.com/hashicorp/terraform-provider-google-beta/pull/4383))
* sql: Fixed a bug where `google_sql_database_instance` would fail if a replica was created, with an encryption key, in a different region than the master instance. ([#4379](https://github.com/hashicorp/terraform-provider-google-beta/pull/4379))

## 4.25.0 (June 15, 2022)

IMPROVEMENTS:
* bigquery: added `connection_id` to `external_data_configuration` for `google_bigquery_table` ([#4365](https://github.com/hashicorp/terraform-provider-google-beta/pull/4365))
* cloudfunctions2: added support for configuring `service_account_email` to `google_cloudfunctions2_function` resource ([#4367](https://github.com/hashicorp/terraform-provider-google-beta/pull/4367))
* compute: added `advanced_options_config` to `google_compute_security_policy` ([#4354](https://github.com/hashicorp/terraform-provider-google-beta/pull/4354))
* compute: added `cache_key_policy` field to `google_compute_backend_bucket` resource ([#4349](https://github.com/hashicorp/terraform-provider-google-beta/pull/4349))
* compute: added `include_named_cookies` to `cdn_policy` on `compute_backend_service` resource ([#4358](https://github.com/hashicorp/terraform-provider-google-beta/pull/4358))
* compute: added internal IPv6 support on `google_compute_network` and `google_compute_subnetwork` ([#4368](https://github.com/hashicorp/terraform-provider-google-beta/pull/4368))
* container: added `managed_prometheus` to `monitoring_config` in `google_container_cluster` ([#4373](https://github.com/hashicorp/terraform-provider-google-beta/pull/4373))
* container: added `spot` field to `node_config` sub-resource ([#4350](https://github.com/hashicorp/terraform-provider-google-beta/pull/4350))
* gkehub: added `prevent_drift` field to `google_gke_hub_feature_membership` resource ([#4370](https://github.com/hashicorp/terraform-provider-google-beta/pull/4370))
* monitoring: added support for JSONPath content matchers to `google_monitoring_uptime_check_config` resource ([#4361](https://github.com/hashicorp/terraform-provider-google-beta/pull/4361))
* monitoring: added support for `user_labels` to `google_monitoring_slo` resource ([#4363](https://github.com/hashicorp/terraform-provider-google-beta/pull/4363))
* sql: added `sql_server_user_details` field to `google_sql_user` resource ([#4364](https://github.com/hashicorp/terraform-provider-google-beta/pull/4364))

BUG FIXES:
* certificatemanager: fixed bug where `DEFAULT` scope would permadiff and force replace the certificate. ([#4356](https://github.com/hashicorp/terraform-provider-google-beta/pull/4356))
* dns: fixed perma-diff for updated labels in `google_dns_managed_zone` ([#4372](https://github.com/hashicorp/terraform-provider-google-beta/pull/4372))
* storagetransfer: fixed perm diff on transfer_options for `google_storage_transfer_job` ([#4357](https://github.com/hashicorp/terraform-provider-google-beta/pull/4357))

## 4.24.0 (June 6, 2022)

IMPROVEMENTS:
* compute: added `cache_key_policy` field to `google_compute_backend_bucket` resource ([#4349](https://github.com/hashicorp/terraform-provider-google-beta/pull/4349))

## 4.23.0 (June 1, 2022)

FEATURES:
* **New Data Source:** `google_tags_tag_key` ([#4337](https://github.com/hashicorp/terraform-provider-google-beta/pull/4337))
* **New Data Source:** `google_tags_tag_value` ([#4337](https://github.com/hashicorp/terraform-provider-google-beta/pull/4337))
* **New Resource:** `google_dataplex_lake` ([#4341](https://github.com/hashicorp/terraform-provider-google-beta/pull/4341))

IMPROVEMENTS:
* bigqueryconnection: updated connection types to support v1 ga ([#4323](https://github.com/hashicorp/terraform-provider-google-beta/pull/4323))
* cloudfunctions: added docker registry support for Cloud Functions ([#4324](https://github.com/hashicorp/terraform-provider-google-beta/pull/4324))
* memcache: added `maintenance_policy` and `maintenance_schedule` to `google_memcache_instance` ([#4338](https://github.com/hashicorp/terraform-provider-google-beta/pull/4338))
* service-directory: marked network field immutable in `google_service_directory_endpoint` ([#4334](https://github.com/hashicorp/terraform-provider-google-beta/pull/4334))

BUG FIXES:
* binaryauthorization: fixed permadiff in `google_binary_authorization_attestor` ([#4325](https://github.com/hashicorp/terraform-provider-google-beta/pull/4325))
* service: added re-polling for service account after creation, 404s sometimes due to [eventual consistency](https://cloud.google.com/iam/docs/overview#consistency) ([#4333](https://github.com/hashicorp/terraform-provider-google-beta/pull/4333))

## 4.22.0 (May 24, 2022)

NOTE: Due to technical difficulties encountered in the release process, the `4.22.0` release for `google-beta` occurred several hours after the corresponding `google` provider release.

FEATURES:
* **New Resource:** `google_certificate_manager_certificate` ([#4301](https://github.com/hashicorp/terraform-provider-google-beta/pull/4301))
* **New Resource:** `google_certificate_manager_dns_authorization` ([#4301](https://github.com/hashicorp/terraform-provider-google-beta/pull/4301))
* **New Resource:** `google_clouddeploy_delivery_pipeline` ([#4288](https://github.com/hashicorp/terraform-provider-google-beta/pull/4288))
* **New Resource:** `google_clouddeploy_target` ([#4288](https://github.com/hashicorp/terraform-provider-google-beta/pull/4288))

IMPROVEMENTS:
* bigquery: added connection of type cloud_resource for `google_bigquery_connection` ([#4312](https://github.com/hashicorp/terraform-provider-google-beta/pull/4312))
* cloudfunctions: added `https_trigger_security_level` to `google_cloudfunctions_function` ([#4295](https://github.com/hashicorp/terraform-provider-google-beta/pull/4295))
* cloudrun: added `traffic.tag` and `traffic.url` fields to `google_cloud_run_service` ([#4283](https://github.com/hashicorp/terraform-provider-google-beta/pull/4283))
* compute: added `enable_dynamic_port_allocation` to `google_compute_router_nat` ([#4316](https://github.com/hashicorp/terraform-provider-google-beta/pull/4316))
* compute: added field `update_policy.most_disruptive_allowed_action` to `google_compute_instance_group_manager` and `google_compute_region_instance_group_manager` ([#4282](https://github.com/hashicorp/terraform-provider-google-beta/pull/4282))
* compute: added support for NEG type `PRIVATE_SERVICE_CONNECT` in `NetworkEndpointGroup` ([#4303](https://github.com/hashicorp/terraform-provider-google-beta/pull/4303))
* compute: added support for `domain_names` attribute in `google_compute_service_attachment` ([#4313](https://github.com/hashicorp/terraform-provider-google-beta/pull/4313))
* compute: added value `REFRESH` to field update_policy.minimal_action` in `google_compute_instance_group_manager` and `google_compute_region_instance_group_manager` ([#4282](https://github.com/hashicorp/terraform-provider-google-beta/pull/4282))
* container: added field `exclusion_options` to `google_container_cluster` ([#4291](https://github.com/hashicorp/terraform-provider-google-beta/pull/4291))
* monitoring: added `checker_type` field to `google_monitoring_uptime_check_config` resource ([#4302](https://github.com/hashicorp/terraform-provider-google-beta/pull/4302))
* privateca: add a new field `desired_state` to manage CertificateAuthority state. ([#4279](https://github.com/hashicorp/terraform-provider-google-beta/pull/4279))
* sql: added `active_directory_config` field in `google_sql_database_instance` ([#4298](https://github.com/hashicorp/terraform-provider-google-beta/pull/4298))
* sql: removed requirement that Cloud SQL Insight is only allowed for Postgres in `google_sql_database_instance` ([#4310](https://github.com/hashicorp/terraform-provider-google-beta/pull/4310))

BUG FIXES:
* cloudfunctions: fixed an issue where `google_cloudfunctions2_function` would not update ([#4278](https://github.com/hashicorp/terraform-provider-google-beta/pull/4278))
* compute: fixed extra diffs generated on `google_security_policy` `rules` when modifying a rule ([#4287](https://github.com/hashicorp/terraform-provider-google-beta/pull/4287))
* container: fixed Autopilot cluster couldn't omit master ipv4 cidr in `google_container_cluster` ([#4280](https://github.com/hashicorp/terraform-provider-google-beta/pull/4280))
* resourcemanager: fixed a bug in wrongly writing to state when creation failed on `google_project_organization_policy` ([#4297](https://github.com/hashicorp/terraform-provider-google-beta/pull/4297))
* storage: not specifying `content` or `source` for `google_storage_bucket_object` now fails at plan-time instead of apply-time. ([#4292](https://github.com/hashicorp/terraform-provider-google-beta/pull/4292))

## 4.21.0 (May 16, 2022)

IMPROVEMENTS:
* cloudfunctions: added CMEK support for Cloud Functions ([#4272](https://github.com/hashicorp/terraform-provider-google-beta/pull/4272))
* compute: added `service_directory_registrations` to `google_compute_forwarding_rule` resource ([#4276](https://github.com/hashicorp/terraform-provider-google-beta/pull/4276))
* compute: removed validation checking against a fixed set of persistent disk types ([#4273](https://github.com/hashicorp/terraform-provider-google-beta/pull/4273))
* container: removed validation checking against a fixed set of persistent disk types ([#4273](https://github.com/hashicorp/terraform-provider-google-beta/pull/4273))
* containeraws: added `image_type` and `instance_placement` to `google_container_aws_node_pool` resource ([#4276](https://github.com/hashicorp/terraform-provider-google-beta/pull/4276))
* containeraws: added `instance_placement` and `logging_config` to `google_container_aws_cluster` resource ([#4276](https://github.com/hashicorp/terraform-provider-google-beta/pull/4276))
* containeraws: added `proxy_config` to `google_container_aws_node_pool` resource ([#4276](https://github.com/hashicorp/terraform-provider-google-beta/pull/4276))
* containerazure: added `image_type` to `google_container_azure_node_pool` resource ([#4276](https://github.com/hashicorp/terraform-provider-google-beta/pull/4276))
* containerazure: added `logging_config` to `google_container_azure_cluster` resource ([#4276](https://github.com/hashicorp/terraform-provider-google-beta/pull/4276))
* containerazure: added `proxy_config` to `google_container_azure_node_pool` resource ([#4276](https://github.com/hashicorp/terraform-provider-google-beta/pull/4276))
* dataproc: removed validation checking against a fixed set of persistent disk types ([#4273](https://github.com/hashicorp/terraform-provider-google-beta/pull/4273))
* dns: added `routing_policy` to `google_dns_record_set` resource ([#4265](https://github.com/hashicorp/terraform-provider-google-beta/pull/4265))

BUG FIXES:
* cloudfunctions: fixed an issue where `google_cloudfunctions2_function` would not update ([#4278](https://github.com/hashicorp/terraform-provider-google-beta/pull/4278))
* compute: fixed a crash in `google_compute_instance` when the instance is deleted outside of Terraform ([#4262](https://github.com/hashicorp/terraform-provider-google-beta/pull/4262))
* provider: removed printing credentials to the console if malformed JSON is given ([#4266](https://github.com/hashicorp/terraform-provider-google-beta/pull/4266))

## 4.20.0 (May 2, 2022)

NOTES:
* `google_privateca_certificate_authority` resources now cannot be destroyed unless `deletion_protection = false` is set in state for the resource. ([#4241](https://github.com/hashicorp/terraform-provider-google-beta/pull/4241))

FEATURES:
* **New Data Source:** `google_compute_disk` ([#4255](https://github.com/hashicorp/terraform-provider-google-beta/pull/4255))

IMPROVEMENTS:
* apigee: `consumer_accept_list` and `service_attachment` to `google_apigee_instance`. ([#4260](https://github.com/hashicorp/terraform-provider-google-beta/pull/4260))
* compute: added `subsetting` field to `google_compute_region_backend_service` ([#4246](https://github.com/hashicorp/terraform-provider-google-beta/pull/4246))
* privateca: added `deletion_protection` for `google_privateca_certificate_authority`. ([#4241](https://github.com/hashicorp/terraform-provider-google-beta/pull/4241))
* privateca: added new output fields on `google_privateca_certificate` including `issuer_certificate_authority`, `pem_certificate_chain` and `certificate_description.x509_description` ([#4242](https://github.com/hashicorp/terraform-provider-google-beta/pull/4242))
* redis: added multi read replica field `read_replicas_mode` and `secondary_ip_range` in `google_redis_instance` ([#4259](https://github.com/hashicorp/terraform-provider-google-beta/pull/4259))

BUG FIXES:
* compute: fixed a crash when `compute.instance` is not found ([#4262](https://github.com/hashicorp/terraform-provider-google-beta/pull/4262))
* provider: removed printing credentials to the console if malformed JSON is given ([#4266](https://github.com/hashicorp/terraform-provider-google-beta/pull/4266))
* sql: fixed bug where `encryption_key_name` was not being propagated to the API. ([#4261](https://github.com/hashicorp/terraform-provider-google-beta/pull/4261))

## 4.19.0 (April 25, 2022)

IMPROVEMENTS:
* cloudbuild: made `CLOUD_LOGGING_ONLY` available as a cloud build logging option. ([#4224](https://github.com/hashicorp/terraform-provider-google-beta/pull/4224))
* compute: added `redirect_options` field for `google_compute_security_policy` rules ([#4217](https://github.com/hashicorp/terraform-provider-google-beta/pull/4217))
* compute: added `FIXED_STANDARD` and `STANDARD` as valid values to the field `network_interface.0.access_configs.0.network_tier` of  `google_compute_instance_template` resource ([#4233](https://github.com/hashicorp/terraform-provider-google-beta/pull/4233))
* compute: added `FIXED_STANDARD` and `STANDARD` as valid values to the field `network_interface.0.access_configs.0.network_tier` of  `google_compute_instance` resource ([#4233](https://github.com/hashicorp/terraform-provider-google-beta/pull/4233))
* compute: added passing `exceed_redirect_options` field for `google_compute_security_policy` rules ([#4238](https://github.com/hashicorp/terraform-provider-google-beta/pull/4238))
* container: added `gke_backup_agent_config` in `addons_config` to `google_container_cluster` (beta) ([#4231](https://github.com/hashicorp/terraform-provider-google-beta/pull/4231))
* filestore: added `kms_key_name` field to `google_filestore_instance` resource to support CMEK ([#11493](https://github.com/hashicorp/terraform-provider-google/pull/11493))
* logging: made `google_logging_*_bucket_config` deletable ([#4234](https://github.com/hashicorp/terraform-provider-google-beta/pull/4234))
* notebooks: updated `container_images` on `google_notebooks_runtime` to default to the value returned by the API if not set ([#4216](https://github.com/hashicorp/terraform-provider-google-beta/pull/4216))
* provider: modified request retry logic to retry all per-minute quota limits returned with a 403 error code. Previously, only read requests were retried. This will generally affect Google Compute Engine resources. ([#4223](https://github.com/hashicorp/terraform-provider-google-beta/pull/4223))

BUG FIXES:
* bigquery: fixed a bug where `encryption_configuration.kms_key_name` stored the version rather than the key name. ([#4221](https://github.com/hashicorp/terraform-provider-google-beta/pull/4221))
* compute: fixed url_mask required mis-annotation in `google_compute_region_network_endpoint_group`, making it optional ([#4227](https://github.com/hashicorp/terraform-provider-google-beta/pull/4227))
* spanner: fixed escaping of database names with Postgres dialect in `google_spanner_database` ([#4228](https://github.com/hashicorp/terraform-provider-google-beta/pull/4228))

## 4.18.0 (April 18, 2022)

FEATURES:
* **New Resource:** `google_privateca_certificate_template_iam_binding` ([#4201](https://github.com/hashicorp/terraform-provider-google-beta/pull/4201))
* **New Resource:** `google_privateca_certificate_template_iam_member` ([#4201](https://github.com/hashicorp/terraform-provider-google-beta/pull/4201))
* **New Resource:** `google_privateca_certificate_template_iam_policy` ([#4201](https://github.com/hashicorp/terraform-provider-google-beta/pull/4201))

IMPROVEMENTS:
* bigtable: added `gc_rules` to `google_bigtable_gc_policy` resource. ([#4212](https://github.com/hashicorp/terraform-provider-google-beta/pull/4212))
* dialogflow: added support for location based dialogflow resources ([#4206](https://github.com/hashicorp/terraform-provider-google-beta/pull/4206))
* metastore: added support for encryption_config during service creation. ([#4204](https://github.com/hashicorp/terraform-provider-google-beta/pull/4204))
* privateca: support update on CertificateAuthority and Certificate ([#4207](https://github.com/hashicorp/terraform-provider-google-beta/pull/4207))

BUG FIXES:
* Update mutex on google_apigee_instance_attachment to lock on org_id. ([#4203](https://github.com/hashicorp/terraform-provider-google-beta/pull/4203))
* vpcaccess: fixed an issue where `google_vpc_access_connector` would be repeatedly recreated when `network` was not specified ([#4205](https://github.com/hashicorp/terraform-provider-google-beta/pull/4205))

## 4.17.0 (April 11, 2022)

FEATURES:
* **New Data Source:** `google_access_approval_folder_service_account` ([#4179](https://github.com/hashicorp/terraform-provider-google-beta/pull/4179))
* **New Data Source:** `google_access_approval_organization_service_account` ([#4179](https://github.com/hashicorp/terraform-provider-google-beta/pull/4179))
* **New Data Source:** `google_access_approval_project_service_account` ([#4179](https://github.com/hashicorp/terraform-provider-google-beta/pull/4179))
* **New Resource:** `google_access_context_manager_access_policy_iam_binding` ([#4180](https://github.com/hashicorp/terraform-provider-google-beta/pull/4180))
* **New Resource:** `google_access_context_manager_access_policy_iam_member` ([#4180](https://github.com/hashicorp/terraform-provider-google-beta/pull/4180))
* **New Resource:** `google_access_context_manager_access_policy_iam_policy` ([#4180](https://github.com/hashicorp/terraform-provider-google-beta/pull/4180))
* **New Resource:** `google_endpoints_service_consumers_iam_binding` ([#4160](https://github.com/hashicorp/terraform-provider-google-beta/pull/4160))
* **New Resource:** `google_endpoints_service_consumers_iam_member` ([#4160](https://github.com/hashicorp/terraform-provider-google-beta/pull/4160))
* **New Resource:** `google_endpoints_service_consumers_iam_policy` ([#4160](https://github.com/hashicorp/terraform-provider-google-beta/pull/4160))
* **New Resource:** `google_iam_deny_policy` ([#4194](https://github.com/hashicorp/terraform-provider-google-beta/pull/4194))

IMPROVEMENTS:
* access approval: added `active_key_version`, `ancestor_has_active_key_version`, and `invalid_key_version` fields to `google_folder_access_approval_settings`, `google_organization_access_approval_settings`, and `google_project_access_approval_settings` resources ([#4179](https://github.com/hashicorp/terraform-provider-google-beta/pull/4179))
* access context manager: added support for scoped policies in `google_access_context_manager_access_policy` ([#4180](https://github.com/hashicorp/terraform-provider-google-beta/pull/4180))
* apigee: added `deployment_type` and `api_proxy_type` to `google_apigee_environment` ([#4177](https://github.com/hashicorp/terraform-provider-google-beta/pull/4177))
* bigtable: updated the examples to show users can create all 3 different flavors of AppProfile ([#4172](https://github.com/hashicorp/terraform-provider-google-beta/pull/4172))
* cloudbuild: added `approval_config` to `google_cloudbuild_trigger` ([#4162](https://github.com/hashicorp/terraform-provider-google-beta/pull/4162))
* composer: added support for `airflow-1` and `airflow-2` aliases in image version argument ([#4185](https://github.com/hashicorp/terraform-provider-google-beta/pull/4185))
* dataflow: added `skip_wait_on_job_termination` attribute to `google_dataflow_job` and `google_dataflow_flex_template_job` resources (issue #10559) ([#4196](https://github.com/hashicorp/terraform-provider-google-beta/pull/4196))
* dataproc: added `presto_config` to `dataproc_job` ([#4171](https://github.com/hashicorp/terraform-provider-google-beta/pull/4171))
* healthcare: added support V3 parser version for Healthcare HL7 stores. ([#4189](https://github.com/hashicorp/terraform-provider-google-beta/pull/4189))
* healthcare: added support for `ANALYTICS_V2 `and `LOSSLESS` BigQueryDestination schema types to `google_healthcare_fhir_store` ([#4186](https://github.com/hashicorp/terraform-provider-google-beta/pull/4186))
* os-config: added field `migInstancesAllowed` to resource `os_config_patch_deployment` ([#4195](https://github.com/hashicorp/terraform-provider-google-beta/pull/4195))
* privateca: added support for IAM conditions to CaPool ([#4170](https://github.com/hashicorp/terraform-provider-google-beta/pull/4170))
* pubsub: added `enable_exactly_once_delivery` to `google_pubsub_subscription` ([#4166](https://github.com/hashicorp/terraform-provider-google-beta/pull/4166))
* spanner: added support for setting database_dialect on `google_spanner_database` ([#4158](https://github.com/hashicorp/terraform-provider-google-beta/pull/4158))

BUG FIXES:
* redis: fixed an issue where older redis instances had a dangerous diff on the field `read_replicas_mode`, adding a default of `READ_REPLICAS_DISABLED`. Now, if the field is not set in config, the value of the field will keep the old value from state. ([#4184](https://github.com/hashicorp/terraform-provider-google-beta/pull/4184))
* tags: fixed issue where tags could not be applied sequentially to the same parent in `google_tags_tag_binding` ([#4191](https://github.com/hashicorp/terraform-provider-google-beta/pull/4191))

## 4.16.0 (April 4, 2022)

FEATURES:
* **New Data Source:** `google_dataproc_metastore_service` ([#4155](https://github.com/hashicorp/terraform-provider-google-beta/pull/4155))
* **New Resource:** `google_firebaserules_release` ([#4132](https://github.com/hashicorp/terraform-provider-google-beta/pull/4132))
* **New Resource:** `google_firebaserules_ruleset` ([#4132](https://github.com/hashicorp/terraform-provider-google-beta/pull/4132))

IMPROVEMENTS:
* bigtable: added support for `autoscaling_config` to `google_bigtable_instance` ([#4150](https://github.com/hashicorp/terraform-provider-google-beta/pull/4150))
* composer: Added support for `composer-1` and `composer-2` aliases in image version argument ([#4131](https://github.com/hashicorp/terraform-provider-google-beta/pull/4131))
* compute: added support for attaching a `edge_security_policy` to `google_compute_backend_bucket` ([#4154](https://github.com/hashicorp/terraform-provider-google-beta/pull/4154))
* compute: added support for field `type` to `google_compute_security_policy` ([#4154](https://github.com/hashicorp/terraform-provider-google-beta/pull/4154))
* eventarc: added gke and workflows destination for eventarc trigger resource. ([#4152](https://github.com/hashicorp/terraform-provider-google-beta/pull/4152))
* networkservices: added `included_cookie_names` to cache key policy configuration ([#4147](https://github.com/hashicorp/terraform-provider-google-beta/pull/4147))
* spanner: added support for setting database_dialect on `google_spanner_database` ([#4158](https://github.com/hashicorp/terraform-provider-google-beta/pull/4158))
* storagetransfer: added `repeat_interval` field to `google_storage_transfer_job` resource ([#4144](https://github.com/hashicorp/terraform-provider-google-beta/pull/4144))

BUG FIXES:
* apikeys: fixed a bug where `google_apikeys_key.key_string` was not being set. ([#4139](https://github.com/hashicorp/terraform-provider-google-beta/pull/4139))
* container: fixed a bug where `google_container_cluster.authenticator_groups_config` could not be set in tandem with `enable_autopilot` ([#4140](https://github.com/hashicorp/terraform-provider-google-beta/pull/4140))
* iam: fixed an issue where special identifiers `allAuthenticatedUsers` and `allUsers` were flattened to lower case in IAM members. ([#4156](https://github.com/hashicorp/terraform-provider-google-beta/pull/4156))
* logging: fixed bug where `google_logging_project_bucket_config` would erroneously write to state after it errored out and wasn't actually created. ([#4141](https://github.com/hashicorp/terraform-provider-google-beta/pull/4141))
* monitoring: fixed a permadiff when `google_monitoring_uptime_check_config.http_check.path` does not begin with "/" ([#4135](https://github.com/hashicorp/terraform-provider-google-beta/pull/4135))
* osconfig: fixed a bug where `recurring_schedule.time_of_day` can not be set to 12am exact time in `google_os_config_patch_deployment` resource ([#4127](https://github.com/hashicorp/terraform-provider-google-beta/pull/4127))
* sql: fixed bug where permadiff of `encryption_key_name` would show on `google_sql_database_instance` for replica instances. ([#4130](https://github.com/hashicorp/terraform-provider-google-beta/pull/4130))
* storage: fixed a bug where `google_storage_bucket` data source would retry for 20 min when bucket was not found. ([#4129](https://github.com/hashicorp/terraform-provider-google-beta/pull/4129))
* storage: fixed bug where `google_storage_transfer_job` that was deleted outside of Terraform would not be recreated on apply. ([#4138](https://github.com/hashicorp/terraform-provider-google-beta/pull/4138))


## 4.15.0 (March 21, 2022)

FEATURES:
* **New Resource:** google_logging_log_view ([#4125](https://github.com/hashicorp/terraform-provider-google-beta/pull/4125))

IMPROVEMENTS:
* apigee: added `billing_type` attribute to `google_apigee_organization` resource. ([#4126](https://github.com/hashicorp/terraform-provider-google-beta/pull/4126))
* networkservices: added `disable_http2` property to `google_network_services_edge_cache_service` resource ([#4119](https://github.com/hashicorp/terraform-provider-google-beta/pull/4119))
* networkservices: updated `google_network_services_edge_cache_origin` resource to read and write the `timeout` property, including a new `read_timeout` field. ([#4122](https://github.com/hashicorp/terraform-provider-google-beta/pull/4122))
* networkservices: updated `google_network_services_edge_cache_origin` to retry_conditions to include `FORBIDDEN` ([#4122](https://github.com/hashicorp/terraform-provider-google-beta/pull/4122))

BUG FIXES:
* dataproc: fixed a crash when `logging_config` only contains `nil` entry  in `google_dataproc_workflow_template` ([#4124](https://github.com/hashicorp/terraform-provider-google-beta/pull/4124))
* sql: fixed crash when one of `settings.database_flags` is nil. ([#4123](https://github.com/hashicorp/terraform-provider-google-beta/pull/4123))

## 4.14.0 (March 14, 2022)

FEATURES:
* **New Resource:** `google_bigqueryreservation_assignment` ([#4098](https://github.com/hashicorp/terraform-provider-google-beta/pull/4098))
* **New Resource:** `google_apikeys_key` ([#4114](https://github.com/hashicorp/terraform-provider-google-beta/pull/4114))

IMPROVEMENTS:
* artifactregistry: added maven config for `google_artifact_registry_repository` ([#4112](https://github.com/hashicorp/terraform-provider-google-beta/pull/4112))
* cloudbuild: added support for manual builds, git source for webhook/pubsub triggered builds and filter field ([#4100](https://github.com/hashicorp/terraform-provider-google-beta/pull/4100))
* container: added support for gvnic to `google_container_node_pool` ([#4111](https://github.com/hashicorp/terraform-provider-google-beta/pull/4111))
* dataproc: added `preemptibility` field to the `preemptible_worker_config` of `google_dataproc_cluster` ([#4107](https://github.com/hashicorp/terraform-provider-google-beta/pull/4107))
* serviceusage: supported `force` behavior for deleting consumer quota override ([#4094](https://github.com/hashicorp/terraform-provider-google-beta/pull/4094))

BUG FIXES:
* dataproc: fixed a crash when `logging_config` only contains `nil` entry  in `google_dataproc_job` ([#4108](https://github.com/hashicorp/terraform-provider-google-beta/pull/4108))

## 4.13.0 (March 7, 2022)

FEATURES:
* **New Resource:** `google_apigee_endpoint_attachment` ([#4074](https://github.com/hashicorp/terraform-provider-google-beta/pull/4074))
* **New Resource:** `google_cloudfunctions2_function` ([#4093](https://github.com/hashicorp/terraform-provider-google-beta/pull/4093))
* **New Resource:** `google_region_backend_service_iam_*` ([#4088](https://github.com/hashicorp/terraform-provider-google-beta/pull/4088))
* **New Datasource:** `google_dns_record_set` ([#4085](https://github.com/hashicorp/terraform-provider-google-beta/pull/4085))
* **New Datasource:** `google_privateca_certificate_authority` ([#4087](https://github.com/hashicorp/terraform-provider-google-beta/pull/4087))

IMPROVEMENTS:
* compute: added support for `keepalive_interval` to `google_compute_router.bgp` ([#4089](https://github.com/hashicorp/terraform-provider-google-beta/pull/4089))
* compute: added update support for `google_compute_reservation.share_settings` ([#4092](https://github.com/hashicorp/terraform-provider-google-beta/pull/4092))
* storagetransfer: added attribute `subject_id` to data source `google_storage_transfer_project_service_account` ([#4073](https://github.com/hashicorp/terraform-provider-google-beta/pull/4073))

BUG FIXES:
* composer: allow region to be undefined in configuration for `google_composer_environment` ([#4083](https://github.com/hashicorp/terraform-provider-google-beta/pull/4083))
* container: fixed a bug where `vertical_pod_autoscaling` would cause autopilot clusters to recreate ([#4076](https://github.com/hashicorp/terraform-provider-google-beta/pull/4076))

## 4.12.0 (February 28, 2022)

NOTE:
* updated to go 1.16.14 ([#4066](https://github.com/hashicorp/terraform-provider-google-beta/pull/4066))

FEATURES:
* **New Resource:** `dns_response_policy*` ([#4046](https://github.com/hashicorp/terraform-provider-google-beta/pull/4046))
* **New Resource:** `dns_response_policy_rule*` ([#4046](https://github.com/hashicorp/terraform-provider-google-beta/pull/4046))

DEPRECATIONS:
* datafusion: deprecated `service_account` in `google_datafusion_instance`. Use `tenant_project_id` instead to extract the tenant project ID (beta) ([#4045](https://github.com/hashicorp/terraform-provider-google-beta/pull/4045))

IMPROVEMENTS:
* bigquery: added support for authorized datasets to `google_bigquery_dataset.access` and `google_bigquery_dataset_access` ([#4047](https://github.com/hashicorp/terraform-provider-google-beta/pull/4047))
* bigtable: added `multi_cluster_routing_cluster_ids` fields to `google_bigtable_app_profile` ([#4051](https://github.com/hashicorp/terraform-provider-google-beta/pull/4051))
* compute: added field `serverless_deployment` to `google_compute_network_endpoint_group` (beta only) for API Gateway resources ([#4041](https://github.com/hashicorp/terraform-provider-google-beta/pull/4041))
* compute: updated `instance` attribute for `google_compute_network_endpoint` to be optional, as Hybrid connectivity NEGs use network endpoints with just IP and Port. ([#4068](https://github.com/hashicorp/terraform-provider-google-beta/pull/4068))
* compute: added `NON_GCP_PRIVATE_IP_PORT` value for `network_endpoint_type` in the `google_compute_network_endpoint_group` resource ([#4068](https://github.com/hashicorp/terraform-provider-google-beta/pull/4068))
* compute: added `provisioning_model` field to `google_compute_instance_template ` resource to support Spot VM(beta) ([#4033](https://github.com/hashicorp/terraform-provider-google-beta/pull/4033))
* compute: added `provisioning_model` field to `google_compute_instance` resource to support Spot VM(beta) ([#4033](https://github.com/hashicorp/terraform-provider-google-beta/pull/4033))
* container: Add support for GKE Compact Placement ([#4043](https://github.com/hashicorp/terraform-provider-google-beta/pull/4043))
* datafusion: added support for `tenant_project_id` and `gcs_bucket` in `google_datafusion_instance` resource. ([#4045](https://github.com/hashicorp/terraform-provider-google-beta/pull/4045))
* provider: added retries for `ReadRequest` errors incorrectly coded as `403` errors, particularly in Google Compute Engine ([#4064](https://github.com/hashicorp/terraform-provider-google-beta/pull/4064))

BUG FIXES:
* apigee: fixed a bug where multiple `google_apigee_instance` could not be used on the same `google_apigee_organization` ([#4059](https://github.com/hashicorp/terraform-provider-google-beta/pull/4059))
* compute: corrected an issue in `google_compute_security_policy` where only alpha values for certain enums were accepted ([#4049](https://github.com/hashicorp/terraform-provider-google-beta/pull/4049))
* compute: fixed permadiff in `google_compute_instance.scheduling.provisioning_model` ([#4044](https://github.com/hashicorp/terraform-provider-google-beta/pull/4044))
* compute: fixed permadiff in `google_compute_instance_template.scheduling.provisioning_model` ([#4052](https://github.com/hashicorp/terraform-provider-google-beta/pull/4052))

## 4.11.0 (February 16, 2022)

IMPROVEMENTS:
* cloudfunctions: Added SecretManager integration support to `google_cloudfunctions_function`. ([#4040](https://github.com/hashicorp/terraform-provider-google-beta/pull/4040))
* compute: Added field `serverless_deployment` to `google_compute_network_endpoint_group` ([#4041](https://github.com/hashicorp/terraform-provider-google-beta/pull/4041))
* dataproc: increased the default timeout for `google_dataproc_cluster` from 20m to 45m ([#4027](https://github.com/hashicorp/terraform-provider-google-beta/pull/4027))
* sql: added field `clone.allocated_ip_range` to support address range picker for clone in resource `google_sql_database_instance` ([#4037](https://github.com/hashicorp/terraform-provider-google-beta/pull/4037))
* storagetransfer: added support for POSIX data source and data sink to `google_storage_transfer_job` via `transfer_spec.posix_data_source` and `transfer_spec.posix_data_sink` fields ([#4029](https://github.com/hashicorp/terraform-provider-google-beta/pull/4029))

BUG FIXES:
* cloudrun: updated `containers.ports.container_port` to be optional instead of required on `google_cloud_run_service` ([#4030](https://github.com/hashicorp/terraform-provider-google-beta/pull/4030))
* compute: marked `project` field optional in `google_compute_instance_template` data source ([#4031](https://github.com/hashicorp/terraform-provider-google-beta/pull/4031))

## 4.10.0 (February 7, 2022)

FEATURES:
* **New Resource:** `google_backend_service_iam_*` ([#4021](https://github.com/hashicorp/terraform-provider-google-beta/pull/4021))

IMPROVEMENTS:
* compute: added `EXTERNAL_MANAGED` as option for `load_balancing_scheme` in `google_compute_global_forwarding_rule` resource ([#4011](https://github.com/hashicorp/terraform-provider-google-beta/pull/4011))
* compute: added field `rate_limit_options` to `google_compute_security_policy` rules ([#4020](https://github.com/hashicorp/terraform-provider-google-beta/pull/4020))
* container: added support for image type configuration on the GKE Node Auto-provisioning ([#4023](https://github.com/hashicorp/terraform-provider-google-beta/pull/4023))
* container: added support for GCPFilestoreCSIDriver addon to `google_container_cluster` resource. ([#4015](https://github.com/hashicorp/terraform-provider-google-beta/pull/4015))
* dataproc: increased the default timeout for `google_dataproc_cluster` from 20m to 45m ([#4027](https://github.com/hashicorp/terraform-provider-google-beta/pull/4027))
* redis: added `maintenance_policy` and `maintenance_schedule` to `google_redis_instance` ([#4010](https://github.com/hashicorp/terraform-provider-google-beta/pull/4010))
* vpcaccess: updated field `network` in `google_vpc_access_connector` to accept `self_link` or `name` ([#4013](https://github.com/hashicorp/terraform-provider-google-beta/pull/4013))

BUG FIXES:
* storage: fixed bug where the provider crashes when `Object.owner` is missing when using `google_storage_object_acl` ([#4019](https://github.com/hashicorp/terraform-provider-google-beta/pull/4019))

## 4.9.0 (January 31, 2022)

BREAKING CHANGES:
* cloudrun: changed the `location` of `google_cloud_run_service` so that modifying the `location` field will recreate the resource rather than causing Terraform to report it would attempt an invalid update ([#3998](https://github.com/hashicorp/terraform-provider-google-beta/pull/3998))

IMPROVEMENTS:
* provider: changed the default timeout for many resources to 20 minutes, the current Terraform default, where it was less than 20 minutes previously ([#4002](https://github.com/hashicorp/terraform-provider-google-beta/pull/4002))
* redis: added `maintenance_policy` and `maintenance_schedule` to `google_redis_instance` ([#4010](https://github.com/hashicorp/terraform-provider-google-beta/pull/4010))
* storage: added field `transfer_spec.aws_s3_data_source.role_arn` to `google_storage_transfer_job` ([#3999](https://github.com/hashicorp/terraform-provider-google-beta/pull/3999))

BUG FIXES:
* cloudrun: fixed a bug where changing the non-updatable `location` of a `google_cloud_run_service` would not force resource recreation ([#3998](https://github.com/hashicorp/terraform-provider-google-beta/pull/3998))
* compute: fixed a bug where `google_compute_firewall` would incorrectly find `source_ranges` to be empty during validation ([#4008](https://github.com/hashicorp/terraform-provider-google-beta/pull/4008))
* notebooks: fixed permadiff in `google_notebooks_runtime.software_config` ([#3997](https://github.com/hashicorp/terraform-provider-google-beta/pull/3997))

## 4.8.0 (January 24, 2022)

BREAKING CHANGES:
* dlp: renamed the `characters_to_ignore.character_to_skip` field to `characters_to_ignore.characters_to_skip` in `google_data_loss_prevention_deidentify_template`. Any affected configurations will have been failing with an error at apply time already. ([#3983](https://github.com/hashicorp/terraform-provider-google-beta/pull/3983))

FEATURES:
* **New Resource:** `google_network_connectivity_spoke` ([#3987](https://github.com/hashicorp/terraform-provider-google-beta/pull/3987))

IMPROVEMENTS:
* apigee: added `ip_range` field to `google_apigee_instance` ([#3989](https://github.com/hashicorp/terraform-provider-google-beta/pull/3989))
* cloudrun: added support for `default_mode` and `mode` settings for created files within `secrets` in `google_cloud_run_service` ([#3984](https://github.com/hashicorp/terraform-provider-google-beta/pull/3984))
* compute: Added `share_settings` in `google_compute_reservation` ([#3980](https://github.com/hashicorp/terraform-provider-google-beta/pull/3980))

BUG FIXES:
* all: Fixed operation polling to support custom endpoints. ([#3986](https://github.com/hashicorp/terraform-provider-google-beta/pull/3986))
* cloudrun: Fixed permadiff in `google_cloud_run_service`'s `template.spec.service_account_name`. ([#3993](https://github.com/hashicorp/terraform-provider-google-beta/pull/3993))
* dlp: Fixed typo in name of `characters_to_ignore.characters_to_skip` field for `google_data_loss_prevention_deidentify_template` ([#3983](https://github.com/hashicorp/terraform-provider-google-beta/pull/3983))
* storagetransfer: fixed bug where `schedule` was required, but really it is optional. ([#3995](https://github.com/hashicorp/terraform-provider-google-beta/pull/3995))

## 4.7.0 (January 19, 2022)

IMPROVEMENTS:
* compute: added `EXTERNAL_MANAGED` as option for `load_balancing_scheme` in `google_compute_backend_service` resource ([#3975](https://github.com/hashicorp/terraform-provider-google-beta/pull/3975))
* container: promoted `dns_config` field of `google_container_cluster` to GA ([#3978](https://github.com/hashicorp/terraform-provider-google-beta/pull/3978))
* monitoring: added `conditionMatchedLog` and `alertStrategy` fields to `google_monitoring_alert_policy` resource ([#3968](https://github.com/hashicorp/terraform-provider-google-beta/pull/3968))

## 4.6.0 (January 10, 2022)

BREAKING CHANGES:
* pubsub: changed `google_pubsub_schema` so that modifiying fields will recreate the resource rather than causing Terraform to report it would attempt an invalid update ([#3933](https://github.com/hashicorp/terraform-provider-google-beta/pull/3933))

FEATURES:
* **New Resource:** `google_apigee_nat_address` ([#3941](https://github.com/hashicorp/terraform-provider-google-beta/pull/3941))
* **New Resource:** `google_network_connectivity_hub` ([#3947](https://github.com/hashicorp/terraform-provider-google-beta/pull/3947))

IMPROVEMENTS:
* bigquery: added ability to create a table with both a schema and view simultaneously to `google_bigquery_table` ([#3950](https://github.com/hashicorp/terraform-provider-google-beta/pull/3950))
* cloud_composer: Added support for Cloud Composer master authorized networks flag ([#3937](https://github.com/hashicorp/terraform-provider-google-beta/pull/3937))
* container: Added field `identity_service_config` to `google_container_cluster` ([#3957](https://github.com/hashicorp/terraform-provider-google-beta/pull/3957))
* osconfig: Added daily os config patch deployments ([#3945](https://github.com/hashicorp/terraform-provider-google-beta/pull/3945))
* storage: added configurable read timeout to `google_storage_bucket` ([#3938](https://github.com/hashicorp/terraform-provider-google-beta/pull/3938))

BUG FIXES:
* billingbudget: fixed a bug where `google_billing_budget.budget_filter.labels` was not updating. ([#3932](https://github.com/hashicorp/terraform-provider-google-beta/pull/3932))
* compute: fixed scenario where `region_instance_group_manager` would not start update if `wait_for_instances` was set and initial status was not `STABLE` ([#3949](https://github.com/hashicorp/terraform-provider-google-beta/pull/3949))
* healthcare: Added back `self_link` functionality which was accidentally removed in `4.0.0` release. ([#3946](https://github.com/hashicorp/terraform-provider-google-beta/pull/3946))
* pubsub: fixed update failure when attempting to change non-updatable resource `google_pubsub_schema` ([#3933](https://github.com/hashicorp/terraform-provider-google-beta/pull/3933))
* storage: fixed a bug where `google_storage_bucket.lifecycle_rule.condition.days_since_custom_time` was not updating. ([#3936](https://github.com/hashicorp/terraform-provider-google-beta/pull/3936))
* vpcaccess: Added back `self_link` functionality which was accidentally removed in `4.0.0` release. ([#3946](https://github.com/hashicorp/terraform-provider-google-beta/pull/3946))

## 4.5.0 (December 20, 2021)

FEATURES:
* **New Data Source:** google_container_aws_versions ([#3928](https://github.com/hashicorp/terraform-provider-google-beta/pull/3928))
* **New Data Source:** google_container_azure_versions ([#3928](https://github.com/hashicorp/terraform-provider-google-beta/pull/3928))
* **New Resource:** google_container_aws_cluster ([#3928](https://github.com/hashicorp/terraform-provider-google-beta/pull/3928))
* **New Resource:** google_container_aws_node_pool ([#3928](https://github.com/hashicorp/terraform-provider-google-beta/pull/3928))
* **New Resource:** google_container_azure_client ([#3928](https://github.com/hashicorp/terraform-provider-google-beta/pull/3928))
* **New Resource:** google_container_azure_cluster ([#3928](https://github.com/hashicorp/terraform-provider-google-beta/pull/3928))
* **New Resource:** google_container_azure_node_pool ([#3928](https://github.com/hashicorp/terraform-provider-google-beta/pull/3928))


IMPROVEMENTS:
* bigquery: added the `return_table_type` field to `google_bigquery_routine` ([#3922](https://github.com/hashicorp/terraform-provider-google-beta/pull/3922))
* cloudbuild: added support for `available_secrets` to `google_cloudbuild_trigger` ([#3907](https://github.com/hashicorp/terraform-provider-google-beta/pull/3907))
* cloudfunctions: added support for `min_instances` to `google_cloudfunctions_function` ([#3904](https://github.com/hashicorp/terraform-provider-google-beta/pull/3904))
* composer: added support for Private Service Connect by adding field `cloud_composer_connection_subnetwork` in `google_composer_environment` ([#3912](https://github.com/hashicorp/terraform-provider-google-beta/pull/3912))
* compute: fixed bug where `google_compute_instance`'s `can_ip_forward` could not be updated without recreating or restarting the instance. ([#3920](https://github.com/hashicorp/terraform-provider-google-beta/pull/3920))
* compute: added field `public_access_prevention` to resource `bucket` ([#3919](https://github.com/hashicorp/terraform-provider-google-beta/pull/3919))
* compute: added support for regional external HTTP(S) load balancer ([#3916](https://github.com/hashicorp/terraform-provider-google-beta/pull/3916))
* privateca: added support for setting default values for basic constraints for `google_privateca_certificate`, `google_privateca_certificate_authority`, and `google_privateca_ca_pool` via the `non_ca` and `zero_max_issuer_path_length` fields ([#3902](https://github.com/hashicorp/terraform-provider-google-beta/pull/3902))
* provider: enabled gRPC requests and response logging ([#3910](https://github.com/hashicorp/terraform-provider-google-beta/pull/3910))

BUG FIXES:
* assuredworkloads: fixed a bug preventing `google_assured_workloads_workload` from being created in any region other than us-central1 ([#3925](https://github.com/hashicorp/terraform-provider-google-beta/pull/3925))

## 4.4.0 (December 13, 2021)

DEPRECATIONS:
* filestore: deprecated `zone` on `google_filestore_instance` in favor of `location` to allow for regional instances ([#3887](https://github.com/hashicorp/terraform-provider-google-beta/pull/3887))

FEATURES:
* **New Resource:** `google_os_config_os_policy_assignment` ([#3892](https://github.com/hashicorp/terraform-provider-google-beta/pull/3892))
* **New Resource:** `google_recaptcha_enterprise_key` ([#3890](https://github.com/hashicorp/terraform-provider-google-beta/pull/3890))

IMPROVEMENTS:
* filestore: added support for `ENTERPRISE` value on `google_filestore_instance` `tier` ([#3887](https://github.com/hashicorp/terraform-provider-google-beta/pull/3887))
* privateca: added support for setting default values for basic constraints for `google_privateca_certificate`, `google_privateca_certificate_authority`, and `google_privateca_ca_pool` via the `non_ca` and `zero_max_issuer_path_length` fields ([#3902](https://github.com/hashicorp/terraform-provider-google-beta/pull/3902))
* sql: added field `allocated_ip_range` to resource `google_sql_database_instance` ([#3897](https://github.com/hashicorp/terraform-provider-google-beta/pull/3897))

BUG FIXES:
* compute: fixed incorrectly failing validation for `INTERNAL_MANAGED` `google_compute_region_backend_service`. ([#3888](https://github.com/hashicorp/terraform-provider-google-beta/pull/3888))
* compute: fixed scenario where `instance_group_manager` would not start update if `wait_for_instances` was set and initial status was not `STABLE` ([#3893](https://github.com/hashicorp/terraform-provider-google-beta/pull/3893))
* container: fixed the `ROUTES` value for the `networking_mode` field in `google_container_cluster`. A recent API change unintentionally changed the default to a `VPC_NATIVE` cluster, and removed the ability to create a `ROUTES`-based one. Provider versions prior to this one will default to `VPC_NATIVE` due to this change, and are unable to create `ROUTES` clusters. ([#3896](https://github.com/hashicorp/terraform-provider-google-beta/pull/3896))

## 4.3.0 (December 7, 2021)

FEATURES:
* **New Data Source:** `google_compute_router_status` ([#3859](https://github.com/hashicorp/terraform-provider-google-beta/pull/3859))
* **New Data Source:** `google_folders` ([#3886](https://github.com/hashicorp/terraform-provider-google-beta/pull/3886))
* **New Resource:** `google_notebooks_runtime` ([#3878](https://github.com/hashicorp/terraform-provider-google-beta/pull/3878))
* **New Resource:** `google_vertex_ai_metadata_store` ([#3885](https://github.com/hashicorp/terraform-provider-google-beta/pull/3885))

IMPROVEMENTS
* apigee: Added IAM support for `google_apigee_environment`. ([#3871](https://github.com/hashicorp/terraform-provider-google-beta/pull/3871)):
* apigee: Added supported values for 'peeringCidrRange' in `google_apigee_instance`. ([#3880](https://github.com/hashicorp/terraform-provider-google-beta/pull/3880))
* cloudbuild: added display_name and annotations to google_cloudbuild_worker_pool for compatibility with new GA. ([#3873](https://github.com/hashicorp/terraform-provider-google-beta/pull/3873))
* container: added `node_group` to `node_config` for container clusters and node pools to support sole tenancy ([#3881](https://github.com/hashicorp/terraform-provider-google-beta/pull/3881))
* container: added `spot` field to `node_config` sub-resource ([#3863](https://github.com/hashicorp/terraform-provider-google-beta/pull/3863))
* redis: Added Multi read replica field `replicaCount `, `nodes`,  `readEndpoint`, `readEndpointPort`, `readReplicasMode` in `google_redis_instance` ([#3870](https://github.com/hashicorp/terraform-provider-google-beta/pull/3870))

BUG FIXES:
* essentialcontacts: marked updating `email` in `google_essential_contacts_contact` as requiring recreation ([#3864](https://github.com/hashicorp/terraform-provider-google-beta/pull/3864))
* privateca: fixed crlAccessUrls in `CertificateAuthority ` ([#3861](https://github.com/hashicorp/terraform-provider-google-beta/pull/3861))

## 4.2.1 (December 3, 2021)

BUG FIXES:
* provider: reverted a requirement in v4.2.0 for Terraform 0.13 and above. This release should be compatible with Terraform 0.12.31

## 4.2.0 (December 2, 2021)

FEATURES:
* **New Data Source:** `google_compute_router_status` ([#3859](https://github.com/hashicorp/terraform-provider-google-beta/pull/3859))

IMPROVEMENTS:
* compute: added support for `queue_count` to `google_compute_instance.network_interface` and `google_compute_instance_template.network_interface` ([#3857](https://github.com/hashicorp/terraform-provider-google-beta/pull/3857))

BUG FIXES:
* all: fixed an issue where some documentation for new resources was not showing up in the GA provider if it was beta-only. ([#3848](https://github.com/hashicorp/terraform-provider-google-beta/pull/3848))
* bigquery: fixed update failure when attempting to change non-updatable fields in `google_bigquery_routine`. ([#3849](https://github.com/hashicorp/terraform-provider-google-beta/pull/3849))
* compute: fixed a bug that would cause `google_instance_from_machine_image` to fail with a resourceInUseByAnotherResource error ([#3855](https://github.com/hashicorp/terraform-provider-google-beta/pull/3855))
* compute: fixed a bug when `cache_mode` is set to FORCE_CACHE_ALL on `google_compute_backend_bucket` ([#3858](https://github.com/hashicorp/terraform-provider-google-beta/pull/3858))
* compute: fixed a perma-diff on `google_compute_region_health_check` when `log_config.enable` is set to false ([#3853](https://github.com/hashicorp/terraform-provider-google-beta/pull/3853))
* servicedirectory: added support for vpc network configuration in `google_service_directory_endpoint`. ([#3856](https://github.com/hashicorp/terraform-provider-google-beta/pull/3856))

## 4.1.0 (November 15, 2021)

IMPROVEMENTS:
* compute: Added `bfd` to `google_compute_router_peer` ([#3822](https://github.com/hashicorp/terraform-provider-google-beta/pull/3822))
* container: added `gcfs_config` to `node_config` of `google_container_node_pool` resource ([#3828](https://github.com/hashicorp/terraform-provider-google-beta/pull/3828))
* provider: added retries for the `resourceNotReady` error returned when attempting to add resources to a recently-modified subnetwork ([#3827](https://github.com/hashicorp/terraform-provider-google-beta/pull/3827))
* pubsub: added `message_retention_duration` field to `google_pubsub_topic` ([#3831](https://github.com/hashicorp/terraform-provider-google-beta/pull/3831))

BUG FIXES:
* apigee: fixed a bug where multiple `google_apigee_instance_attachment` could not be used on the same `google_apigee_instance` ([#3838](https://github.com/hashicorp/terraform-provider-google-beta/pull/3838))
* bigquery: fixed a bug following import where schema is empty on `google_bigquery_table` ([#3839](https://github.com/hashicorp/terraform-provider-google-beta/pull/3839))
* billingbudget: fixed unable to provide `labels` on `google_billing_budget` ([#3823](https://github.com/hashicorp/terraform-provider-google-beta/pull/3823))
* compute: allowed `source_disk` to accept full image path on `google_compute_snapshot` ([#3835](https://github.com/hashicorp/terraform-provider-google-beta/pull/3835))
* compute: fixed a bug in `google_compute_firewall` that would cause changes in `source_ranges` to not correctly be applied ([#3834](https://github.com/hashicorp/terraform-provider-google-beta/pull/3834))
* logging: fixed a bug with updating `description` on `google_logging_project_sink`, `google_logging_folder_sink` and `google_logging_organization_sink` ([#3826](https://github.com/hashicorp/terraform-provider-google-beta/pull/3826))

## 4.0.0 (November 02, 2021)

NOTES:
* compute: Google Compute Engine resources will now call the endpoint appropriate to the provider version rather than the beta endpoint by default ([#3787](https://github.com/hashicorp/terraform-provider-google-beta/pull/3787))
* container: Google Kubernetes Engine resources will now call the endpoint appropriate to the provider version rather than the beta endpoint by default ([#3788](https://github.com/hashicorp/terraform-provider-google-beta/pull/3788))

BREAKING CHANGES:
* appengine: marked `google_app_engine_standard_app_version` `entrypoint` as required ([#3784](https://github.com/hashicorp/terraform-provider-google-beta/pull/3784))
* compute: removed the ability to specify the `trace-append` or `trace-ro` as scopes in `google_compute_instance`, use `trace` instead ([#3759](https://github.com/hashicorp/terraform-provider-google-beta/pull/3759))
* compute: changed `advanced_machine_features` on `google_compute_instance_template` to track changes when the block is undefined in a user's config ([#3786](https://github.com/hashicorp/terraform-provider-google-beta/pull/3786))
* compute: changed `source_ranges` in `google_compute_firewall_rule` to track changes when it is not set in a config file ([#3791](https://github.com/hashicorp/terraform-provider-google-beta/pull/3791))
* compute: changed the import / drift detection behaviours for `metadata_startup_script`, `metadata.startup-script` in `google_compute_instance`. Now, `metadata.startup-script` will be set by default, and `metadata_startup_script` will only be set if present. ([#3765](https://github.com/hashicorp/terraform-provider-google-beta/pull/3765))
* compute: removed `source_disk_link` field from `google_compute_snapshot` ([#3783](https://github.com/hashicorp/terraform-provider-google-beta/pull/3783))
* container: `instance_group_urls` has been removed in favor of `node_pool.instance_group_urls` ([#3796](https://github.com/hashicorp/terraform-provider-google-beta/pull/3796))
* container: changed default for `enable_shielded_nodes` to true for `google_container_cluster` ([#3773](https://github.com/hashicorp/terraform-provider-google-beta/pull/3773))
* container: made `master_auth.client_certificate_config` required ([#3794](https://github.com/hashicorp/terraform-provider-google-beta/pull/3794))
* container: removed `master_auth.username` and `master_auth.password` from `google_container_cluster` ([#3794](https://github.com/hashicorp/terraform-provider-google-beta/pull/3794))
* container: removed `workload_metadata_configuration.node_metadata` in favor of `workload_metadata_configuration.mode` in `google_container_cluster` ([#3772](https://github.com/hashicorp/terraform-provider-google-beta/pull/3772))
* container: removed the `workload_identity_config.0.identity_namespace` field from `google_container_cluster`, use `workload_identity_config.0.workload_pool` instead ([#3776](https://github.com/hashicorp/terraform-provider-google-beta/pull/3776))
* kms: removed `self_link` field from `google_kms_crypto_key` and `google_kms_key_ring` ([#3783](https://github.com/hashicorp/terraform-provider-google-beta/pull/3783))
* project: removed ability to specify `bigquery-json.googleapis.com`, the provider will no longer convert it as the upstream API migration is finished. Use `bigquery.googleapis.com` instead. ([#3751](https://github.com/hashicorp/terraform-provider-google-beta/pull/3751))
* provider: changed `credentials`, `access_token` precedence so that `credentials` values in configuration take precedence over `access_token` values assigned through environment variables ([#3766](https://github.com/hashicorp/terraform-provider-google-beta/pull/3766))
* provider: removed redundant default scopes. The provider's default scopes when authenticating with credentials are now exclusively "https://www.googleapis.com/auth/cloud-platform" and "https://www.googleapis.com/auth/userinfo.email". ([#3756](https://github.com/hashicorp/terraform-provider-google-beta/pull/3756))
* pubsub: removed `path` from `google_pubsub_subscription` ([#3777](https://github.com/hashicorp/terraform-provider-google-beta/pull/3777))
* pubsub: removed `path` field from `google_pubsub_subscription` ([#3783](https://github.com/hashicorp/terraform-provider-google-beta/pull/3783))
* resourcemanager: made `google_project` remove `org_id` and `folder_id` from state when they are removed from config ([#3754](https://github.com/hashicorp/terraform-provider-google-beta/pull/3754))
* resourcemanager: changed the `project` field to `Required` in all `google_project_iam_*` resources ([#3767](https://github.com/hashicorp/terraform-provider-google-beta/pull/3767))
* sql: added drift detection to the following `google_sql_database_instance` fields: `activation_policy` (defaults `ALWAYS`), `availability_type` (defaults `ZONAL`), `disk_type` (defaults `PD_SSD`), `encryption_key_name` ([#3778](https://github.com/hashicorp/terraform-provider-google-beta/pull/3778))
* sql: changed the `database_version` field to `Required` in `google_sql_database_instance` resource ([#3770](https://github.com/hashicorp/terraform-provider-google-beta/pull/3770))
* sql: removed the following `google_sql_database_instance` fields: `authorized_gae_applications`, `crash_safe_replication`, `replication_type` ([#3778](https://github.com/hashicorp/terraform-provider-google-beta/pull/3778))
* storage: removed `bucket_policy_only` from `google_storage_bucket` ([#3769](https://github.com/hashicorp/terraform-provider-google-beta/pull/3769))
* storage: changed the `location` field to required in `google_storage_bucket` ([#3771](https://github.com/hashicorp/terraform-provider-google-beta/pull/3771))

VALIDATION CHANGES:
* bigquery: at least one of `statement_timeout_ms`, `statement_byte_budget`, or `key_result_statement` is required on `google_bigquery_job.query.script_options.` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* bigquery: exactly one of `query`, `load`, `copy` or `extract` is required on `google_bigquery_job` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* bigquery: exactly one of `source_table` or `source_model` is required on `google_bigquery_job.extract` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* cloudbuild: exactly one of `branch_name`, `commit_sha` or `tag_name` is required on `google_cloudbuild_trigger.build.source.repo_source` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* compute: at least one of `fixed_delay` or `percentage` is required on `google_compute_url_map.default_route_action.fault_injection_policy.delay` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* compute: at least one of `fixed` or `percent` is required on `google_compute_autoscaler.autoscaling_policy.scale_down_control.max_scaled_down_replicas` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* compute: at least one of `fixed` or `percent` is required on `google_compute_autoscaler.autoscaling_policy.scale_in_control.max_scaled_in_replicas` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* compute: at least one of `fixed` or `percent` is required on `google_compute_region_autoscaler.autoscaling_policy.scale_down_control.max_scaled_down_replicas` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* compute: at least one of `fixed` or `percent` is required on `google_compute_region_autoscaler.autoscaling_policy.scale_in_control.max_scaled_in_replicas` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* compute: at least one of `max_scaled_down_replicas` or `time_window_sec` is required on `google_compute_autoscaler.autoscaling_policy.scale_down_control` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* compute: at least one of `max_scaled_down_replicas` or `time_window_sec` is required on `google_compute_region_autoscaler.autoscaling_policy.scale_down_control` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* compute: at least one of `max_scaled_in_replicas` or `time_window_sec` is required on `google_compute_autoscaler.autoscaling_policy.scale_in_control.0.` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* compute: at least one of `max_scaled_in_replicas` or `time_window_sec` is required on `google_compute_region_autoscaler.autoscaling_policy.scale_in_control.0.` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* compute: required one of `source_tags`, `source_ranges` or `source_service_accounts` on INGRESS `google_compute_firewall` resources ([#3750](https://github.com/hashicorp/terraform-provider-google-beta/pull/3750))
* dlp: at least one of `start_time` or `end_time` is required on `google_data_loss_prevention_trigger.inspect_job.storage_config.timespan_config` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* dlp: exactly one of `url` or `regex_file_set` is required on `google_data_loss_prevention_trigger.inspect_job.storage_config.cloud_storage_options.file_set` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* resourcemanager: added conflict between `org_id`, `folder_id` at plan time in `google_project` ([#3754](https://github.com/hashicorp/terraform-provider-google-beta/pull/3754))
* osconfig: at least one of `linux_exec_step_config` or `windows_exec_step_config` is required on `google_os_config_patch_deployment.patch_config.post_step` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* osconfig: at least one of `linux_exec_step_config` or `windows_exec_step_config` is required on `google_os_config_patch_deployment.patch_config.pre_step` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* osconfig: at least one of `reboot_config`, `apt`, `yum`, `goo` `zypper`, `windows_update`, `pre_step` or `pre_step` is required on `google_os_config_patch_deployment.patch_config` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* osconfig: at least one of `security`, `minimal`, `excludes` or `exclusive_packages` is required on `google_os_config_patch_deployment.patch_config.yum` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* osconfig: at least one of `type`, `excludes` or `exclusive_packages` is required on `google_os_config_patch_deployment.patch_config.apt` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* osconfig: at least one of `with_optional`, `with_update`, `categories`, `severities`, `excludes` or `exclusive_patches` is required on `google_os_config_patch_deployment.patch_config.zypper` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* osconfig: exactly one of `classifications`, `excludes` or `exclusive_patches` is required on `google_os_config_patch_deployment.inspect_job.patch_config.windows_update` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* spanner: at least one of `num_nodes` or `processing_units` is required on `google_spanner_instance` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))

IMPROVEMENTS:
* container: added `managed_instance_group_urls` to `google_container_node_pool` to replace `instance_group_urls` on `google_container_cluster` ([#3815](https://github.com/hashicorp/terraform-provider-google-beta/pull/3815))
* kms: added support for EKM to `google_kms_crypto_key.protection_level` ([#3763](https://github.com/hashicorp/terraform-provider-google-beta/pull/3763))
* project: added support for `billing_project` on `google_project_service` ([#3768](https://github.com/hashicorp/terraform-provider-google-beta/pull/3768))
* spanner: increased the default timeout on `google_spanner_instance` operations from 4 minutes to 20 minutes, significantly reducing the likelihood that resources will time out ([#3789](https://github.com/hashicorp/terraform-provider-google-beta/pull/3789))

BUG FIXES:
* bigquery: fixed a bug of cannot add required fields to an existing schema on `google_bigquery_table` ([#3781](https://github.com/hashicorp/terraform-provider-google-beta/pull/3781))
* compute: fixed a bug in updating multiple `ttl` fields on `google_compute_backend_bucket` ([#3757](https://github.com/hashicorp/terraform-provider-google-beta/pull/3757))
* compute: fixed a perma-diff on `subnetwork` when it is optional on `google_compute_network_endpoint_group` ([#3780](https://github.com/hashicorp/terraform-provider-google-beta/pull/3780))
* compute: fixed perma-diff bug on `log_config.enable` of both `google_compute_backend_service` and `google_compute_region_backend_service` ([#3760](https://github.com/hashicorp/terraform-provider-google-beta/pull/3760))
* compute: fixed the `google_compute_instance_group_manager.update_policy.0.min_ready_sec` field so that updating it to `0` works ([#3810](https://github.com/hashicorp/terraform-provider-google-beta/pull/3810))
* compute: fixed the `google_compute_region_instance_group_manager.update_policy.0.min_ready_sec` field so that updating it to `0` works ([#3810](https://github.com/hashicorp/terraform-provider-google-beta/pull/3810))
* spanner: fixed the schema for `data.google_spanner_instance` so that non-configurable fields are considered outputs ([#3804](https://github.com/hashicorp/terraform-provider-google-beta/pull/3804))
