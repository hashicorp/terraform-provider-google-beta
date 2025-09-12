## 6.49.3 (Unreleased)

BUG FIXES:
* compute: fixed a crash in `google_compute_security_policy` due to a changed API response for empty `match.0.expr_options` blocks ([#10715](https://github.com/hashicorp/terraform-provider-google-beta/pull/10715))

## 6.49.2 (August 22, 2025)
BUG FIXES:

* container: fixed issue where a failed creation on `google_container_node_pool` would result in an unrecoverable tainted state  ([#24077](https://github.com/hashicorp/terraform-provider-google/pull/24077))

## 6.49.1 (August 20, 2025)
BUG FIXES:
* secretmanager: fixed issue where upgrading to 6.49.0 would cause all `google_secret_manager_secret_version` resources to be recreated unless `secret_data_wo_version` was set ([#10574](https://github.com/hashicorp/terraform-provider-google-beta/pull/10574))


## 6.49.0 (August 19, 2025)

DEPRECATIONS:
* beyondcorp: `google_beyondcorp_application_iam_binding`, `google_beyondcorp_application_iam_member` and `google_beyondcorp_application_iam_policy` IAM resources, and the `google_beyondcorp_application_iam_policy ` datasource have been deprecated and will be removed in the upcoming major release ([#10532](https://github.com/hashicorp/terraform-provider-google-beta/pull/10532))
* tpu: deprecated `google_tpu_tensorflow_versions` data source. Use `google_tpu_v2_runtime_versions` instead. ([#10514](https://github.com/hashicorp/terraform-provider-google-beta/pull/10514))

FEATURES:
* **New Data Source:** `google_artifact_registry_tag` ([#10531](https://github.com/hashicorp/terraform-provider-google-beta/pull/10531))
* **New Data Source:** `google_artifact_registry_tags` ([#10518](https://github.com/hashicorp/terraform-provider-google-beta/pull/10518))
* **New Resource:** `google_dialogflow_convesation_profile` ([#10533](https://github.com/hashicorp/terraform-provider-google-beta/pull/10533))

IMPROVEMENTS:
* apikeys: added `service_account_email` to `google_apikeys_key` ([#10538](https://github.com/hashicorp/terraform-provider-google-beta/pull/10538))
* bigqueryreservation: added support for `scaling_mode` and `max_slots` properties on `google_bigquery_reservation` (beta) ([#10509](https://github.com/hashicorp/terraform-provider-google-beta/pull/10509))
* compute: added `advanced_options_config` field to `google_compute_region_security_policy` resource ([#10498](https://github.com/hashicorp/terraform-provider-google-beta/pull/10498))
* container: added `eviction_soft`, `eviction_soft_grace_period`, `eviction_minimum_reclaim`, `eviction_max_pod_grace_period_seconds`, `max_parallel_image_pulls`, `transparent_hugepage_enabled`, `transparent_hugepage_defrag` and `min_node_cpus` fields to `node_config` block of `google_container_node_pool` and `google_container_cluster` resources ([#10522](https://github.com/hashicorp/terraform-provider-google-beta/pull/10522))
* networkmanagement: added `subnet` and `network` fields to the `google_network_management_vpc_flow_logs_config` resource (beta) ([#10506](https://github.com/hashicorp/terraform-provider-google-beta/pull/10506))
* networkmanagement: added output-only field `target_resource_state` to the `google_network_management_vpc_flow_logs_config` resource ([#10506](https://github.com/hashicorp/terraform-provider-google-beta/pull/10506))
* resourcemanager: Added `management_project` and `configured_capabilities` fields to the `google_folder` resource. ([#10525](https://github.com/hashicorp/terraform-provider-google-beta/pull/10525))

BUG FIXES:
* cloud_tasks: correctly set `name` field to be required in `google_cloud_tasks_queue` resource ([#10534](https://github.com/hashicorp/terraform-provider-google-beta/pull/10534))
* clouddeploy: allowed sending `start_time` with default values in `weekly_windows` in  `google_clouddeploy_deploy_policy` resource. `start_time 00:00` means the policy will start at midnight. ([#10530](https://github.com/hashicorp/terraform-provider-google-beta/pull/10530))
* kms: `skip_initial_version_creation` field is no longer immutable in `google_kms_crypto_key`, but is still only settable at-creation ([#10526](https://github.com/hashicorp/terraform-provider-google-beta/pull/10526))
* netapp: fixed bug where `google_netapp_volume.large_capacity` was not properly marked as immutable, causing updates to fail (and making it impossible to change the field value after creation) ([#10541](https://github.com/hashicorp/terraform-provider-google-beta/pull/10541))
* networkconnectivity: added update support for `linked_vpc_network` in `google_network_connectivity_spoke` ([#10507](https://github.com/hashicorp/terraform-provider-google-beta/pull/10507))


## 6.48.0 (August 12, 2025)

FEATURES:
* **New Data Source:** `google_artifact_registry_package` ([#10490](https://github.com/hashicorp/terraform-provider-google-beta/pull/10490))
* **New Data Source:** `google_artifact_registry_repositories` ([#10494](https://github.com/hashicorp/terraform-provider-google-beta/pull/10494))
* **New Data Source:** `google_artifact_registry_version` ([#10468](https://github.com/hashicorp/terraform-provider-google-beta/pull/10468))
* **New Resource:** `google_dialogflow_cx_playbook` (initial basic support, full features to follow in a later release) ([#10485](https://github.com/hashicorp/terraform-provider-google-beta/pull/10485))
* **New Resource:** `google_vertexai_rag_engine_config` ([#10481](https://github.com/hashicorp/terraform-provider-google-beta/pull/10481))

IMPROVEMENTS:
* backupdr: added `log_retention_days` field to `google_backup_dr_backup_plan` resource ([#10463](https://github.com/hashicorp/terraform-provider-google-beta/pull/10463))
* compute: added `advanced_options_config` field to `google_compute_region_security_policy` resource ([#10498](https://github.com/hashicorp/terraform-provider-google-beta/pull/10498))
* compute: added `ha_policy` field to `google_compute_region_backend_service` resource ([#10493](https://github.com/hashicorp/terraform-provider-google-beta/pull/10493))
* compute: added the ability to use global target forwarding rule for `target_service` field in `google_compute_service_attachment` resource ([#10483](https://github.com/hashicorp/terraform-provider-google-beta/pull/10483))
* container: added `boot_disk` to `node_config` in `google_container_cluster` and `google_container_node_pool` resources ([#10457](https://github.com/hashicorp/terraform-provider-google-beta/pull/10457))
* container: added `node_config.kubelet_config.single_process_oom_kill` field to `google_container_node_pool` and `google_container_cluster` resources ([#10461](https://github.com/hashicorp/terraform-provider-google-beta/pull/10461))
* container: added in-place update support for `user_managed_keys_config` field in `google_container_cluster` resource ([#10475](https://github.com/hashicorp/terraform-provider-google-beta/pull/10475))
* dataproc: added `cluster_config.cluster_tier` field to `google_dataproc_cluster` resource ([#10453](https://github.com/hashicorp/terraform-provider-google-beta/pull/10453))
* gkeonprem: added `enable_advanced_cluster` field to `google_gkeonprem_vmware_admin_cluster` resource ([#10496](https://github.com/hashicorp/terraform-provider-google-beta/pull/10496))
* memorystore: added `allow_fewer_zones_deployment` field to `google_memorystore_instance` resource ([#10462](https://github.com/hashicorp/terraform-provider-google-beta/pull/10462))
* sql: added field `psa_write_endpoint` flag to `google_sql_database_instance` resource ([#10467](https://github.com/hashicorp/terraform-provider-google-beta/pull/10467))
* sql: added `network_attachment_uri` field to `google_sql_database_instance` resource ([#10484](https://github.com/hashicorp/terraform-provider-google-beta/pull/10484))
* sql: added `node_count` field to `sql_database_instance` resource, and added new value `READ_POOL_INSTANCE` enum to `instance_type` field of `sql_database_instance` resource ([#10487](https://github.com/hashicorp/terraform-provider-google-beta/pull/10487))
* storagetransfer: added `federated_identity_config` to `google_storage_transfer_job` resource ([#10489](https://github.com/hashicorp/terraform-provider-google-beta/pull/10489))
* storagetransfer: added `transfer_spec.aws_s3_data_source.cloudfront_domain` field to `google_storage_transfer_job` resource ([#10479](https://github.com/hashicorp/terraform-provider-google-beta/pull/10479))

BUG FIXES:
* accesscontextmanager: made `scopes` field as immutable for `access_context_manager_access_policy` resource ([#10478](https://github.com/hashicorp/terraform-provider-google-beta/pull/10478))
* bigquery: fixed handling of non-legacy roles for access block inside `google_bigquery_dataset` resource ([#10488](https://github.com/hashicorp/terraform-provider-google-beta/pull/10488))
* container: fixed an issue causing errors during updates to `node_config` to be suppressed in `google_container_cluster` and `google_container_node_pool` resources ([#10459](https://github.com/hashicorp/terraform-provider-google-beta/pull/10459))

## 6.47.0 (August 05, 2025)

DEPRECATIONS:
* compute: deprecated `network_self_link` field in `google_compute_subnetworks` data source. Use `network_name` instead. ([#10423](https://github.com/hashicorp/terraform-provider-google-beta/pull/10423))
* resourcemanager: deprecated `project` field in `google_service_account_key` data source. The field is non functional and can safely be removed from your configuration. ([#10442](https://github.com/hashicorp/terraform-provider-google-beta/pull/10442))

FEATURES:
* **New Data Source:** `google_artifact_registry_docker_images` ([#10422](https://github.com/hashicorp/terraform-provider-google-beta/pull/10422))
* **New Resource:** `google_apigee_security_action` ([#10407](https://github.com/hashicorp/terraform-provider-google-beta/pull/10407))
* **New Resource:** `google_developer_connect_insights_config` ([#10431](https://github.com/hashicorp/terraform-provider-google-beta/pull/10431))
* **New Resource:** `google_discovery_engine_cmek_config` ([#10416](https://github.com/hashicorp/terraform-provider-google-beta/pull/10416))
* **New Resource:** `google_iam_workforce_pool_iam_binding` ([#10426](https://github.com/hashicorp/terraform-provider-google-beta/pull/10426))
* **New Resource:** `google_iam_workforce_pool_iam_member` ([#10426](https://github.com/hashicorp/terraform-provider-google-beta/pull/10426))
* **New Resource:** `google_iam_workforce_pool_iam_policy` ([#10426](https://github.com/hashicorp/terraform-provider-google-beta/pull/10426))

IMPROVEMENTS:
* backupdr: added `backup_retention_inheritance` field to `google_backup_dr_backup_vault` resource ([#10446](https://github.com/hashicorp/terraform-provider-google-beta/pull/10446))
* bigqueryanalyticshub: added `commercial_info` and `delete_commercial` fields in `google_bigquery_analytics_hub_listing` resource ([#10415](https://github.com/hashicorp/terraform-provider-google-beta/pull/10415))
* bigqueryanalyticshub: added `discovery_type` field to `google_bigquery_analytics_hub_data_exchange` resource ([#10435](https://github.com/hashicorp/terraform-provider-google-beta/pull/10435))
* bigqueryanalyticshub: added `state`, `discovery_type`, and `allow_only_metadata_sharing` fields to `google_bigquery_analytics_hub_listing` resource ([#10435](https://github.com/hashicorp/terraform-provider-google-beta/pull/10435))
* cloudfunction: added `automatic_update_policy` and `on_deploy_update_policy` to `google_cloudfunctions_function` resource ([#10448](https://github.com/hashicorp/terraform-provider-google-beta/pull/10448))
* cloudrunv2: added `gpu_zonal_redundancy_disabled` field to `google_cloud_run_v2_job` resource. ([#10440](https://github.com/hashicorp/terraform-provider-google-beta/pull/10440))
* compute: added `labels` field to `google_compute_storage_pool` resource ([#10425](https://github.com/hashicorp/terraform-provider-google-beta/pull/10425))
* compute: added `network_name` field to `google_compute_subnetworks` data source ([#10423](https://github.com/hashicorp/terraform-provider-google-beta/pull/10423))
* container: added `ip_allocation_policy.additional_ip_ranges_config` field to `google_container_cluster` resource ([#10451](https://github.com/hashicorp/terraform-provider-google-beta/pull/10451))
* container: added `network_config.additional_node_network_configs.subnetwork` field to `google_container_node_pool` resource ([#10451](https://github.com/hashicorp/terraform-provider-google-beta/pull/10451))
* container: added `addons_config.lustre_csi_driver_config` field to `google_container_cluster` resource ([#10413](https://github.com/hashicorp/terraform-provider-google-beta/pull/10413))
* container: added support for `rbac_binding_config` in `google_container_cluster` ([#10441](https://github.com/hashicorp/terraform-provider-google-beta/pull/10441))
* dataproc: added `cluster_config.cluster_tier` field to `google_dataproc_cluster` resource ([#10453](https://github.com/hashicorp/terraform-provider-google-beta/pull/10453))
* looker: added `LOOKER_CORE_TRIAL_STANDARD`, `LOOKER_CORE_TRIAL_ENTERPRISE`, and `LOOKER_CORE_TRIAL_EMBED` editions to `google_looker_instance` resource. ([#10427](https://github.com/hashicorp/terraform-provider-google-beta/pull/10427))
* managedkafka: added `tls_config` field to `google_managed_kafka_cluster` resource ([#10420](https://github.com/hashicorp/terraform-provider-google-beta/pull/10420))
* memorystore: added `allow_fewer_zones_deployment` field to `google_redis_cluster` resource ([#10434](https://github.com/hashicorp/terraform-provider-google-beta/pull/10434))
* storage: added `deletion_policy` field to `google_storage_bucket_object` resource ([#10445](https://github.com/hashicorp/terraform-provider-google-beta/pull/10445))
* vertexai: added `custom_delete` field to `google_vertex_ai_endpoint_with_model_garden_deployment` resource ([#10430](https://github.com/hashicorp/terraform-provider-google-beta/pull/10430))

BUG FIXES:
* bigquery: fixed a crash in `google_bigquery_table` when configured as an external table with `parquet_options` ([#10438](https://github.com/hashicorp/terraform-provider-google-beta/pull/10438))
* cloudrunv2: fixed an issue where `manual_instance_count` was unable to set to `0` in `google_cloud_run_v2_worker_pool`. ([#10433](https://github.com/hashicorp/terraform-provider-google-beta/pull/10433))
* composer: fixed updates failing for `recovery_config` with explicitly disabled scheduled snapshots ([#10405](https://github.com/hashicorp/terraform-provider-google-beta/pull/10405))
* iap: fixed an issue where deleting `google_iap_settings` without setting `GOOGLE_PROJECT` incorrectly failed ([#10410](https://github.com/hashicorp/terraform-provider-google-beta/pull/10410))
* storage: removed client-side GCS name validations for `google_storage_bucket` ([#10406](https://github.com/hashicorp/terraform-provider-google-beta/pull/10406))

## 6.46.0 (July 29, 2025)

FEATURES:
* **New Data Source:** `google_storage_insights_dataset_config` ([#10402](https://github.com/hashicorp/terraform-provider-google-beta/pull/10402))
* **New Resource:** `google_apigee_api_product` ([#10378](https://github.com/hashicorp/terraform-provider-google-beta/pull/10378))
* **New Resource:** `google_discovery_engine_recommendation_engine` ([#10394](https://github.com/hashicorp/terraform-provider-google-beta/pull/10394))
* **New Resource:** `google_oracle_database_odb_network` ([#10383](https://github.com/hashicorp/terraform-provider-google-beta/pull/10383))
* **New Resource:**  `google_oracle_database_odb_subnet` ([#10396](https://github.com/hashicorp/terraform-provider-google-beta/pull/10396))
* **New Resource:** `google_storage_insights_dataset_config` ([#10401](https://github.com/hashicorp/terraform-provider-google-beta/pull/10401))

IMPROVEMENTS:
* compute: added `params.resourceManagerTags` field to the `google_compute_router` ([#10392](https://github.com/hashicorp/terraform-provider-google-beta/pull/10392))
* compute: added in-place update support for `provisioned_iops`, `provisioned_throughput`,  and `access_mode` fields in `google_compute_region_disk` resource ([#10397](https://github.com/hashicorp/terraform-provider-google-beta/pull/10397))
* dataproc: added `authentication_config` field to `google_dataproc_batch` and `google_dataproc_session_template` resource ([#10375](https://github.com/hashicorp/terraform-provider-google-beta/pull/10375))
* dataproc: added `idle_ttl` field to `google_dataproc_session_template` resource ([#10386](https://github.com/hashicorp/terraform-provider-google-beta/pull/10386))
* networkconnectivity: added field `allocation_options` to resource `google_network_connectivity_internal_range` ([#10390](https://github.com/hashicorp/terraform-provider-google-beta/pull/10390))
* oracledatabase: added `odb_network` and `odb_subnet` fields, and made `network` and `cidr` fields optional in `google_oracle_database_autonomous_database` resource ([#10389](https://github.com/hashicorp/terraform-provider-google-beta/pull/10389))
* oracledatabase: added `odb_network`, `odb_subnet` and `backup_odb_subnet` fields, and made `network`, `cidr` and `backup_subnet_cidr` fields optional in `google_oracle_database_cloud_vm_cluster` resource ([#10391](https://github.com/hashicorp/terraform-provider-google-beta/pull/10391))
* secretmanager: added `tags` field to `google_secret_manager_regional_secret` to allow setting tags for regional_secrets at creation time ([#10400](https://github.com/hashicorp/terraform-provider-google-beta/pull/10400))
* securesourcemanager: added `deletion_policy` field to `google_secure_source_manager_repository` resource ([#10395](https://github.com/hashicorp/terraform-provider-google-beta/pull/10395))
* workbench: added `enable_managed_euc` field to `google_workbench_instance` resource. ([#10388](https://github.com/hashicorp/terraform-provider-google-beta/pull/10388))
* workbench: added `reservation_affinity` field to `google_workbench_instance` resource. ([#10384](https://github.com/hashicorp/terraform-provider-google-beta/pull/10384))

BUG FIXES:
* composer: fixed updates failing for `google_composer_environment` `recovery_config` with explicitly disabled scheduled snapshots ([#10405](https://github.com/hashicorp/terraform-provider-google-beta/pull/10405))
* datastore: fixed a permadiff with `google_datastream_connection_profile`'s `create_without_validation` field ([#10403](https://github.com/hashicorp/terraform-provider-google-beta/pull/10403))
* memorystore: fixed bug to allow `google_memorystore_instance`  to be used with no provider default region or with a `location` that doesn't match the provider default region. ([#10380](https://github.com/hashicorp/terraform-provider-google-beta/pull/10380))
* networkconnectivity: fixed `instances[].ip_address` & `instances[].virtual_machine` fields in `linked_router_appliance_instances` block being incorrectly treated as immutable for `google_network_connectivity_spoke` resource ([#10399](https://github.com/hashicorp/terraform-provider-google-beta/pull/10399))
* resourcemanager: updated service account creation to prevent failures due to eventual consistency in `google_service_account` resource ([#10371](https://github.com/hashicorp/terraform-provider-google-beta/pull/10371))
* sql: fixed a provider crash when importing `google_sql_database` resource ([#10374](https://github.com/hashicorp/terraform-provider-google-beta/pull/10374))

## 6.45.0 (July 22, 2025)

DEPRECATIONS:
* gemini: deprecated the `disable_web_grounding` field in the `google_gemini_gemini_gcp_enablement_setting` resource ([#10338](https://github.com/hashicorp/terraform-provider-google-beta/pull/10338))

FEATURES:
* **New Resource:** `google_bigtable_schema_bundle` ([#10342](https://github.com/hashicorp/terraform-provider-google-beta/pull/10342))
* **New Resource:** `google_compute_preview_feature` ([#10364](https://github.com/hashicorp/terraform-provider-google-beta/pull/10364))
* **New Resource:** `google_dialogflow_cx_generator` ([#10348](https://github.com/hashicorp/terraform-provider-google-beta/pull/10348))
* **New Resource:** `google_model_armor_floorsetting` ([#10359](https://github.com/hashicorp/terraform-provider-google-beta/pull/10359))
* **New Resource:** `google_vertex_ai_endpoint_with_model_garden_deployment` ([#10365](https://github.com/hashicorp/terraform-provider-google-beta/pull/10365))

IMPROVEMENTS:
* accesscontextmanager: added `name` to `google_access_context_manager_gcp_user_access_binding` resource ([#10370](https://github.com/hashicorp/terraform-provider-google-beta/pull/10370))
* bigquery: added `ignore_auto_generated_schema` virtual field to `google_bigquery_table` resource to ignore server-added columns in the `schema` field ([#10366](https://github.com/hashicorp/terraform-provider-google-beta/pull/10366))
* compute: added `params.resourceManagerTags` field to the `google_compute_subnetwork` ([#10357](https://github.com/hashicorp/terraform-provider-google-beta/pull/10357))
* compute: added `mirrorPercent` field to `requestMirrorPolicy` in `defaultRouteAction`, `pathMatchers[].defaultRouteAction`, `pathMatchers[].pathRules[].routeAction`, and `pathMatchers[].routeRules[].routeAction` to `google_compute_region_url_map` resource ([#10351](https://github.com/hashicorp/terraform-provider-google-beta/pull/10351))
* compute: added `rule.match.src_secure_tags`, `rule.target_secure_tags`, `predefined_rules.match.src_secure_tags` and `predefined_rules.target_secure_tags` fields to `google_compute_firewall_policy_with_rules` resource ([#10367](https://github.com/hashicorp/terraform-provider-google-beta/pull/10367))
* dataproc: added `cluster_config.security_config.identity_config` field to `google_dataproc_cluster` resource ([#10352](https://github.com/hashicorp/terraform-provider-google-beta/pull/10352))
* dataproc: updated `cluster_config.gce_cluster_config.metadata` field to be computed in `google_dataproc_cluster` resource ([#10352](https://github.com/hashicorp/terraform-provider-google-beta/pull/10352))
* dialogflowcx: added `flexible` support to `google_dialogflow_cx_webhook` resource. ([#10339](https://github.com/hashicorp/terraform-provider-google-beta/pull/10339))
* gemini: added `web_grounding_type` field to `google_gemini_gemini_gcp_enablement_setting` resource ([#10338](https://github.com/hashicorp/terraform-provider-google-beta/pull/10338))
* netapp: added in-place update support for `allow_auto_tiering` field in `google_netapp_storage_pool` resource ([#10353](https://github.com/hashicorp/terraform-provider-google-beta/pull/10353))
* secretmanager: added `tags` field to `google_secret_manager_secret` to allow setting tags for secrets at creation time ([#10360](https://github.com/hashicorp/terraform-provider-google-beta/pull/10360))
* securesourcemanager: added `deletion_policy` field to `google_secure_source_manager_instance` resource ([#10349](https://github.com/hashicorp/terraform-provider-google-beta/pull/10349))
* sql: added `network_attachment_uri` field to `google_sql_database_instance` ([#10354](https://github.com/hashicorp/terraform-provider-google-beta/pull/10354))
* vmwareengine: added `GOOGLE_CLOUD_NETAPP_VOLUMES` peering type to resource `google_vmwareengine_network_peering` ([#10363](https://github.com/hashicorp/terraform-provider-google-beta/pull/10363))

BUG FIXES:
* modelarmor: fixed conflicting field validation for `filter_config.sdp_settings` on `google_model_armor_template` ([#10361](https://github.com/hashicorp/terraform-provider-google-beta/pull/10361))
* resourcemanager: updated service account creation to prevent failures due to eventual consistency in `google_service_account` resource ([#10371](https://github.com/hashicorp/terraform-provider-google-beta/pull/10371))

## 6.44.0 (July 16, 2025)

FEATURES:
* **New Data Source:** `google_compute_network_attachment` ([#10336](https://github.com/hashicorp/terraform-provider-google-beta/pull/10336))
* **New Data Source:** `google_firestore_document` ([#10321](https://github.com/hashicorp/terraform-provider-google-beta/pull/10321))
* **New Resource:** `google_backup_dr_service_config` ([#10320](https://github.com/hashicorp/terraform-provider-google-beta/pull/10320))
* **New Resource:** `google_bigquery_analytics_hub_data_exchange_subscription` ([#10328](https://github.com/hashicorp/terraform-provider-google-beta/pull/10328))

IMPROVEMENTS:
* apigee: added `access_logging_config` field to `google_apigee_instance` resource ([#10303](https://github.com/hashicorp/terraform-provider-google-beta/pull/10303))
* apigee: marked `access_logging_config` field immutable in `google_apigee_instance` resource ([#10337](https://github.com/hashicorp/terraform-provider-google-beta/pull/10337))
* backupdr: added in-place update support for `google_backup_dr_backup_plan` resource ([#10312](https://github.com/hashicorp/terraform-provider-google-beta/pull/10312))
* bigqueryanalyticshub: added `routine` field to `google_bigquery_analytics_hub_listing` resource ([#10327](https://github.com/hashicorp/terraform-provider-google-beta/pull/10327))
* compute: added `params.resource_manager_tags` field to `google_compute_firewall` resource ([#10304](https://github.com/hashicorp/terraform-provider-google-beta/pull/10304))
* compute: added `aggregate_reservation.vm_family`, `aggregate_reservation.reserved_resources.accelerator.accelerator_count`, `aggregate_reservation.reserved_resources.accelerator.accelerator_type` and `aggregate_reservation.workload_type` fields to `google_future_reservation` resource ([#10317](https://github.com/hashicorp/terraform-provider-google-beta/pull/10317))
* compute: added `application_aware_interconnect` and `aai_enabled` fields to `google_compute_interconnect` resource ([#10333](https://github.com/hashicorp/terraform-provider-google-beta/pull/10333))
* compute: added `load_balancing_scheme` field to `google_compute_backend_bucket` resource ([#10301](https://github.com/hashicorp/terraform-provider-google-beta/pull/10301))
* compute: added `provisioned_iops` and `provisioned_throughput` fields to `google_compute_region_disk` resource ([#10319](https://github.com/hashicorp/terraform-provider-google-beta/pull/10319))
* compute: added `request_body_inspection_size` field to `google_compute_security_policy` resource ([#10318](https://github.com/hashicorp/terraform-provider-google-beta/pull/10318))
* compute: added `specific_reservation.instance_properties.maintenance_interval`, `share_settings.projects` and `enable_emergent_maintenance` fields to `google_compute_reservation` resource ([#10329](https://github.com/hashicorp/terraform-provider-google-beta/pull/10329))
* firestore: added `tags` field to `google_firestore_database` resource ([#10335](https://github.com/hashicorp/terraform-provider-google-beta/pull/10335))
* securesourcemanager: added in-place update support for `description` field in `google_secure_source_manager_repository` resource ([#10325](https://github.com/hashicorp/terraform-provider-google-beta/pull/10325))
* storage: added `force_empty_content_type` field to `google_storage_bucket_object` resource ([#10334](https://github.com/hashicorp/terraform-provider-google-beta/pull/10334))

BUG FIXES:
* artifactregistry: fixed an issue where changes to `cleanup_policies` were not being applied correctly in `google_artifact_registry_repository` resource ([#10324](https://github.com/hashicorp/terraform-provider-google-beta/pull/10324))
* firebasehosting: skipped deletion of `google_firebase_hosting_site` resource of type `DEFAULT_SITE` ([#10305](https://github.com/hashicorp/terraform-provider-google-beta/pull/10305))
* iambeta: fixed perma-diff for `jwks_json` field when GCP normalizes JSON formatting in `google_iam_workload_identity_pool_provider` resource ([#10306](https://github.com/hashicorp/terraform-provider-google-beta/pull/10306))

## 6.43.0 (July 8, 2025)

DEPRECATIONS:
* iap: deprecated `google_iap_client` and `google_iap_brand` ([#10269](https://github.com/hashicorp/terraform-provider-google-beta/pull/10269))

FEATURES:
* **New Data Source:** `google_network_management_connectivity_test_run` ([#10300](https://github.com/hashicorp/terraform-provider-google-beta/pull/10300))
* **New Data Source:** `google_redis_cluster` ([#10273](https://github.com/hashicorp/terraform-provider-google-beta/pull/10273))
* **New Resource:** `google_contact_center_insights_analysis_rule` ([#10272](https://github.com/hashicorp/terraform-provider-google-beta/pull/10272))
* **New Resource:** `google_model_armor_template` ([#10270](https://github.com/hashicorp/terraform-provider-google-beta/pull/10270))

IMPROVEMENTS:
* bigquery: added `ignore_schema_changes` virtual field to `google_bigquery_table` resource. Only `dataPolicies` field is supported in `ignore_schema_changes` for now. ([#10299](https://github.com/hashicorp/terraform-provider-google-beta/pull/10299))
* billing: added `currency_code` to `google_billing_account` data source ([#10284](https://github.com/hashicorp/terraform-provider-google-beta/pull/10284))
* compute: added `params.resource_manager_tags` field to `google_compute_network` resource ([#10266](https://github.com/hashicorp/terraform-provider-google-beta/pull/10266))
* compute: added `load_balancing_scheme` field to `google_compute_backend_bucket` resource ([#10301](https://github.com/hashicorp/terraform-provider-google-beta/pull/10301))
* compute: added `params.resource_manager_tags` field to `google_compute_route` resource ([#10293](https://github.com/hashicorp/terraform-provider-google-beta/pull/10293))
* compute: added `update_strategy`  field to `google_compute_network_peering` resource ([#10275](https://github.com/hashicorp/terraform-provider-google-beta/pull/10275))
* container: added `secret_manager_config.rotation_config` field to `google_container_cluster` resource ([#10291](https://github.com/hashicorp/terraform-provider-google-beta/pull/10291))
* container: added `anonymous_authentication_config` field to `google_container_cluster` resource ([#10295](https://github.com/hashicorp/terraform-provider-google-beta/pull/10295))
* dataplex: added `suspended` field to `google_dataplex_datascan` resource ([#10276](https://github.com/hashicorp/terraform-provider-google-beta/pull/10276))
* discoveryengine: added `enable_table_annotation`, `enable_image_annotation`, `structured_content_types`, `exclude_html_elements`, `exclude_html_classes` and `exclude_html_ids` fields to `layout_parsing_config` of `google_discovery_engine_data_store` resource ([#10288](https://github.com/hashicorp/terraform-provider-google-beta/pull/10288))
* discoveryengine: added `kms_key_name` field to `google_discovery_engine_data_store` resource ([#10281](https://github.com/hashicorp/terraform-provider-google-beta/pull/10281))
* memorystore: added `managed_server_ca` field to `google_memorystore_instance` resource ([#10268](https://github.com/hashicorp/terraform-provider-google-beta/pull/10268))
* secretmanager: added `deletion_protection` field to `google_secret_manager_secret` resource to optionally make deleting them require an explicit intent ([#10289](https://github.com/hashicorp/terraform-provider-google-beta/pull/10289))
* secretmanager: added `fetch_secret_data` to `google_secret_manager_secret_version` to optionally skip fetching the secret data ([#10282](https://github.com/hashicorp/terraform-provider-google-beta/pull/10282))

BUG FIXES:
* compute: fixed `match` field in `google_compute_router_route_policy` resource to be marked as required ([#10298](https://github.com/hashicorp/terraform-provider-google-beta/pull/10298))
* compute: fixed an issue with `bgp_always_compare_med` in `google_compute_network` where it was unable to be set from `true` to `false` ([#10286](https://github.com/hashicorp/terraform-provider-google-beta/pull/10286))
* compute: made no replication status in `google_compute_disk_async_replication` a retryable error ([#10296](https://github.com/hashicorp/terraform-provider-google-beta/pull/10296))
* gkeonprem: fixed type of `load_balancer.0.bgp_lb_config.0.address_pools.0.manual_assign` in `google_gkeonprem_bare_metal_cluster`, making it a boolean instead of a string ([#10283](https://github.com/hashicorp/terraform-provider-google-beta/pull/10283))
* integrationconnectors: removed validation from auth configs in `google_integration_connectors_connection` resource ([#10267](https://github.com/hashicorp/terraform-provider-google-beta/pull/10267))


## 6.42.0 (July 1, 2025)

FEATURES:
* **New Resource:** `google_apihub_plugin_instance` ([#10225](https://github.com/hashicorp/terraform-provider-google-beta/pull/10225))
* **New Resource:** `google_apihub_plugin` ([#10254](https://github.com/hashicorp/terraform-provider-google-beta/pull/10254))
* **New Resource:** `google_compute_wire_group` ([#10255](https://github.com/hashicorp/terraform-provider-google-beta/pull/10255))
* **New Resource:** `google_dialogflow_cx_generative_settings` ([#10244](https://github.com/hashicorp/terraform-provider-google-beta/pull/10244))

IMPROVEMENTS:
* cloudidentity: added `create_ignore_already_exists` field to `google_cloud_identity_group_membership` resource ([#10229](https://github.com/hashicorp/terraform-provider-google-beta/pull/10229))
* cloudkms: added `etag` field to `google_kms_autokey_config` resource ([#10227](https://github.com/hashicorp/terraform-provider-google-beta/pull/10227))
* cloudrunv2: added `node_selector` field to `google_cloud_run_v2_job` resource ([#10234](https://github.com/hashicorp/terraform-provider-google-beta/pull/10234))
* compute: added `access_mode` field to `google_compute_region_disk` resource ([#10256](https://github.com/hashicorp/terraform-provider-google-beta/pull/10256))
* compute: added `match.src_secure_tags` and `target_secure_tags` fields to `google_compute_firewall_policy_rule` resource ([#10261](https://github.com/hashicorp/terraform-provider-google-beta/pull/10261))
* compute: added `params.resource_manager_tags` field to `google_compute_network` resource ([#10266](https://github.com/hashicorp/terraform-provider-google-beta/pull/10266))
* compute: added `policy_type` field to `google_compute_network_firewall_policy`, `google_compute_network_firewall_policy_with_rules`, `google_compute_region_network_firewall_policy`, and `google_compute_region_network_firewall_policy_with_rules` resources ([#10239](https://github.com/hashicorp/terraform-provider-google-beta/pull/10239))
* compute: added `resource_policies.workload_policy` field to `google_compute_instance_group_manager` resource ([#10265](https://github.com/hashicorp/terraform-provider-google-beta/pull/10265))
* container: added `confidential_nodes.confidential_instance_type` field to `google_container_cluster` resource ([#10257](https://github.com/hashicorp/terraform-provider-google-beta/pull/10257))
* container: added `gke_auto_upgrade_config` field to `google_container_cluster` resource ([#10258](https://github.com/hashicorp/terraform-provider-google-beta/pull/10258))
* container: added `node_config.confidential_nodes.confidential_instance_type` field to `google_container_node_pool` resource ([#10257](https://github.com/hashicorp/terraform-provider-google-beta/pull/10257))
* firestore: revoked deprecation of `deletion_policy` field in `google_firestore_database` resource ([#10251](https://github.com/hashicorp/terraform-provider-google-beta/pull/10251))
* iam_beta: added `attestation_rules` field to `google_iam_workload_identity_pool_managed_identity` resource ([#10250](https://github.com/hashicorp/terraform-provider-google-beta/pull/10250))
* memorystore: added `kms_key` field to `google_memorystore_instance` resource ([#10246](https://github.com/hashicorp/terraform-provider-google-beta/pull/10246))
* redis: added `effective_reserved_ip_range` field to `google_redis_instance` resource ([#10235](https://github.com/hashicorp/terraform-provider-google-beta/pull/10235))
* secretmanager: added `deletion_protection` field to `google_secret_manager_regional_secret` resource ([#10247](https://github.com/hashicorp/terraform-provider-google-beta/pull/10247))
* spanner: added `encryption_config.kms_key_name` field to `google_spanner_backup_schedule` resource ([#10230](https://github.com/hashicorp/terraform-provider-google-beta/pull/10230))
* storage: added `allow_cross_org_vpcs` and `allow_all_service_agent_access` fields to `google_storage_bucket` resource ([#10252](https://github.com/hashicorp/terraform-provider-google-beta/pull/10252))

BUG FIXES:
* bigqueryanalyticshub: supported in-place update for `log_linked_dataset_query_user_email` in `google_bigquery_analytics_hub_listing` and `google_bigquery_analytics_hub_data_exchange` resources. Once enabled, this feature cannot be disabled. ([#10241](https://github.com/hashicorp/terraform-provider-google-beta/pull/10241))
* bigquerydatatransfer: stopped surfacing persistent warnings recommending write-only field when using `secret_access_key` on `google_bigquery_data_transfer_config` ([#10263](https://github.com/hashicorp/terraform-provider-google-beta/pull/10263))
* memorystore: added the ability to set the `replica_count` field in `google_memorystore_instance` resource to 0 ([#10259](https://github.com/hashicorp/terraform-provider-google-beta/pull/10259))
* monitoring: made `description` and `displayName` optional and mutable in `google_monitoring_metric_descriptor` resource ([#10233](https://github.com/hashicorp/terraform-provider-google-beta/pull/10233))
* redis: fixed `reserved_ip_range` field not being populated for `google_redis_instance` data source ([#10235](https://github.com/hashicorp/terraform-provider-google-beta/pull/10235))
* secretmanager: stopped surfacing persistent warnings recommending write-only field when using `secret_data` on `google_secret_manager_secret_version` ([#10263](https://github.com/hashicorp/terraform-provider-google-beta/pull/10263))
* sql: stopped surfacing persistent warnings recommending write-only field when using `password` on `google_sql_user` ([#10263](https://github.com/hashicorp/terraform-provider-google-beta/pull/10263))
* workbench: added support for setting `serial-port-logging-enable` key in `metadata` field in `google_workbench_instance` resource ([#10253](https://github.com/hashicorp/terraform-provider-google-beta/pull/10253))

## 6.41.0 (June 24, 2025)

BREAKING CHANGES:
* lustre: added `per_unit_storage_throughput` as a required field to `google_lustre_instance` resource in response to a change in the API surface ([#10211](https://github.com/hashicorp/terraform-provider-google-beta/pull/10211))

FEATURES:
* **New Data Source:** `google_dataplex_data_quality_rules` ([#10189](https://github.com/hashicorp/terraform-provider-google-beta/pull/10189))
* **New Resource:** `google_apihub_plugin_instance` ([#10225](https://github.com/hashicorp/terraform-provider-google-beta/pull/10225))
* **New Resource:** `google_contact_center_insights_view` ([#10192](https://github.com/hashicorp/terraform-provider-google-beta/pull/10192))
* **New Resource:** `google_dataproc_session_template` ([#10204](https://github.com/hashicorp/terraform-provider-google-beta/pull/10204))
* **New Resource:** `google_dialogflow_encryption_spec` ([#10220](https://github.com/hashicorp/terraform-provider-google-beta/pull/10220))

IMPROVEMENTS:
* alloydb: added `network_config.allocated_ip_range_override` field to `google_alloydb_instance` resource ([#10216](https://github.com/hashicorp/terraform-provider-google-beta/pull/10216))
* bigqueryanalyticshub: added `log_linked_dataset_query_user_email` field to `google_bigquery_analytics_hub_data_exchange` resource ([#10200](https://github.com/hashicorp/terraform-provider-google-beta/pull/10200))
* bigqueryanalyticshub: added `log_linked_dataset_query_user_email` field to `google_bigquery_analytics_hub_listing_subscription` resource ([#10202](https://github.com/hashicorp/terraform-provider-google-beta/pull/10202))
* bigqueryanalyticshub: added `pubsub_topic` field to `google_bigquery_analytics_hub_listing` resource ([#10219](https://github.com/hashicorp/terraform-provider-google-beta/pull/10219))
* bigtable: added `row_key_schema` to `google_bigtable_table` resource ([#10222](https://github.com/hashicorp/terraform-provider-google-beta/pull/10222))
* cloudasset: added support for universe domain handling for `google_cloud_asset_resources_search_all` datasource. ([#10210](https://github.com/hashicorp/terraform-provider-google-beta/pull/10210))
* cloudquotas: added `inherited`  and `inherited_from` fields to `google_cloud_quotas_quota_adjuster_settings` resource ([#10223](https://github.com/hashicorp/terraform-provider-google-beta/pull/10223))
* compute: added `CROSS_SITE_NETWORK` option to `requested_features` field in `google_compute_interconnect` resource ([#10207](https://github.com/hashicorp/terraform-provider-google-beta/pull/10207))
* compute: added `TLS_JA4_FINGERPRINT` option to `enforce_on_key` field in `google_compute_region_security_policy`, `google_compute_security_policy`, and `google_compute_security_policy_rule` resources ([#10199](https://github.com/hashicorp/terraform-provider-google-beta/pull/10199))
* compute: added `send_propagated_connection_limit_if_zero` to `google_compute_service_attachment` to resolve an issue where `propagated_connection_limit` were not working for 0 value previously. Now setting `send_propagated_connection_limit_if_zero = true` will send `propagated_connection_limit = 0` when it's unset or set to `0`. ([#10213](https://github.com/hashicorp/terraform-provider-google-beta/pull/10213))
* compute: added `wire_groups` field to `google_compute_interconnect` resource ([#10207](https://github.com/hashicorp/terraform-provider-google-beta/pull/10207))
* container: added `performance_monitoring_unit` in node_config/advanced_machine_features to 'google_container_cluster' resource ([#10191](https://github.com/hashicorp/terraform-provider-google-beta/pull/10191))
* container: added `release_channel_upgrade_target_version` to `google_container_engine_versions` data source ([#10221](https://github.com/hashicorp/terraform-provider-google-beta/pull/10221))
* dataplex: added support for discovery scan in `google_dataplex_datascan` resource ([#10205](https://github.com/hashicorp/terraform-provider-google-beta/pull/10205))
* provider: added support for adc impersonation in different universes ([#10212](https://github.com/hashicorp/terraform-provider-google-beta/pull/10212))
* storage: added `source_md5hash` field in `google_storage_bucket_object` ([#10196](https://github.com/hashicorp/terraform-provider-google-beta/pull/10196))

BUG FIXES:
* compute: fixed `google_compute_firewall_policy_rule` staying disabled after apply with `disabled = false` ([#10215](https://github.com/hashicorp/terraform-provider-google-beta/pull/10215))
* compute: marked `name` in `google_compute_node_group`, `google_compute_node_template` as required as it was impossible to create successfully without a value ([#10224](https://github.com/hashicorp/terraform-provider-google-beta/pull/10224))
* sql: fixed an error in updating `connection_pool_config` in `google_sql_database_instance` ([#10218](https://github.com/hashicorp/terraform-provider-google-beta/pull/10218))
* tags: fixed perma-diff for `parent` field in `google_tags_location_tag_binding` resource ([#10217](https://github.com/hashicorp/terraform-provider-google-beta/pull/10217))

## 6.40.0 (June 17, 2025)

DEPRECATIONS:
* notebook: `google_notebook_runtime` is deprecated and will be removed in a future major release. Use `google_workbench_instance` instead. ([#10186](https://github.com/hashicorp/terraform-provider-google-beta/pull/10186))

FEATURES:
* **New Data Source:** `google_dataplex_data_quality_rules` ([#10189](https://github.com/hashicorp/terraform-provider-google-beta/pull/10189))
* **New Resource:** `google_dialogflow_cx_tool` ([#10154](https://github.com/hashicorp/terraform-provider-google-beta/pull/10154))

IMPROVEMENTS:
* backupdr: added 'supported_resource_types' field to `google_backup_dr_backup_plan` resource ([#10155](https://github.com/hashicorp/terraform-provider-google-beta/pull/10155))
* backupdr: added support for updating in-place to the `google_backup_dr_backup_plan_association` resource ([#10176](https://github.com/hashicorp/terraform-provider-google-beta/pull/10176))
* bigqueryanalyticshub: added `log_linked_dataset_query_user_email` field to `google_bigquery_analytics_hub_listing` resource ([#10177](https://github.com/hashicorp/terraform-provider-google-beta/pull/10177))
* compute: added `cipher_suite` block with phase1 and phase2 encryption configurations to `google_compute_vpn_tunnel` resource. ([#10188](https://github.com/hashicorp/terraform-provider-google-beta/pull/10188))
* compute: added `fingerprint` field in `google_compute_target_http_proxy` and `google_compute_target_https_proxy` resources. ([#10175](https://github.com/hashicorp/terraform-provider-google-beta/pull/10175))
* compute: added `headers`, `expected_output_url`, and `expected_redirect_response_code` fields to `test` in `google_compute_url_map` resource and made `service` field optional ([#10161](https://github.com/hashicorp/terraform-provider-google-beta/pull/10161))
* compute: added `path_matcher.default_route_action` fields to `google_compute_region_url_map` resource ([#10171](https://github.com/hashicorp/terraform-provider-google-beta/pull/10171))
* gkehub: added `custom_role` field to `google_gke_hub_scope_rbac_role_binding` resource ([#10151](https://github.com/hashicorp/terraform-provider-google-beta/pull/10151))
* integrationconnectors: added support for `log_config.level` for `google_integration_connectors_connection` ([#10170](https://github.com/hashicorp/terraform-provider-google-beta/pull/10170))
* netapp: added `enable_hot_tier_auto_resize` and `hot_tier_size_gib` fields to `google_netapp_storage_pool` resource ([#10153](https://github.com/hashicorp/terraform-provider-google-beta/pull/10153))
* netapp: added `tiering_policy.hot_tier_bypass_mode_enabled` field to `google_netapp_volume` resource ([#10153](https://github.com/hashicorp/terraform-provider-google-beta/pull/10153))
* networkconnectivity: added `psc_config.producer_instance_location` and `psc_config.allowed_google_producers_resource_hierarchy_level` fields to `google_network_connectivity_service_connection_policy` ([#10179](https://github.com/hashicorp/terraform-provider-google-beta/pull/10179))
* redis: added `managed_server_ca` to `google_redis_cluster` resource ([#10169](https://github.com/hashicorp/terraform-provider-google-beta/pull/10169))
* resourcemanager: allowed `dataproc-control.googleapis.com` and `stackdriverprovisioning.googleapis.com` services in `google_project_service` resource ([#10174](https://github.com/hashicorp/terraform-provider-google-beta/pull/10174))
* storage: removed the hardcoded 80m timeout used during `google_storage_bucket` deletion when removing an anywhere cache, polling instead. This should speed up deletion in these cases. ([#10160](https://github.com/hashicorp/terraform-provider-google-beta/pull/10160))
* vertexai: added `region` to `google_vertex_ai_index_endpoint_deployed_index` ([#10184](https://github.com/hashicorp/terraform-provider-google-beta/pull/10184))

BUG FIXES:
* beyondcorp: fixed the issue where `hubs.internet_gateway.assigned_ips` was not populated correctly in the `google_beyondcorp_security_gateway` resource ([#10182](https://github.com/hashicorp/terraform-provider-google-beta/pull/10182))
* compute: fixed `google_compute_router_nat` where changes to `auto_network_tier` are always showed after initial apply ([#10152](https://github.com/hashicorp/terraform-provider-google-beta/pull/10152))
* compute: fixed validation for `target_service` field in `google_compute_service_attachment` resource causing issues when targeting a `google_network_services_gateway` resource ([#10178](https://github.com/hashicorp/terraform-provider-google-beta/pull/10178))
* dataflow: fields `network`, `subnetwork`, `num_workers`, `max_num_workers` and `machine_type` will no longer cause permadiff on `dataflow_flex_template_job` ([#10168](https://github.com/hashicorp/terraform-provider-google-beta/pull/10168))
* dataproc: fixed a permadiff with "prodcurrent" and "prodprevious" within image subminor version for `google_dataproc_cluster` ([#10163](https://github.com/hashicorp/terraform-provider-google-beta/pull/10163))
* networksecurity: `marked google_network_security_address_group` `capacity` as immutable because it can't be updated in place. ([#10165](https://github.com/hashicorp/terraform-provider-google-beta/pull/10165))

## 6.39.0 (June 10, 2025)

FEATURES:
* **New Resource:** `google_apihub_curation` ([#10130](https://github.com/hashicorp/terraform-provider-google-beta/pull/10130))
* **New Resource:** `google_compute_interconnect_attachment_group` ([#10136](https://github.com/hashicorp/terraform-provider-google-beta/pull/10136))
* **New Resource:** `google_compute_interconnect_group` ([#10136](https://github.com/hashicorp/terraform-provider-google-beta/pull/10136))
* **New Resource:** `google_compute_snapshot_settings` ([#10133](https://github.com/hashicorp/terraform-provider-google-beta/pull/10133))

IMPROVEMENTS:
* apigee: added `client_ip_resolution_config` field to `google_apigee_environment` resource ([#10143](https://github.com/hashicorp/terraform-provider-google-beta/pull/10143))
* beyondcorp: added `delegating_service_account` field to `google_beyondcorp_security_gateway` resource ([#10114](https://github.com/hashicorp/terraform-provider-google-beta/pull/10114))
* bigquery: added `data_source_id` to update requests through `google_bigquery_data_transfer_config` ([#10126](https://github.com/hashicorp/terraform-provider-google-beta/pull/10126))
* cloudrunv2: added `google_cloud_run_v2_job` support for `depends_on` and `startup_probe` properties ([#10147](https://github.com/hashicorp/terraform-provider-google-beta/pull/10147))
* container: added `network_performance_config` field to `google_container_cluster` resource ([#10117](https://github.com/hashicorp/terraform-provider-google-beta/pull/10117))
* dataplex: added `catalog_publishing_enabled` field to `google_dataplex_datascan` resource ([#10141](https://github.com/hashicorp/terraform-provider-google-beta/pull/10141))
* datastream: added `network_attachment` support via `psc_interface_config` attribute in `google_datastream_private_connection` ([#10112](https://github.com/hashicorp/terraform-provider-google-beta/pull/10112))
* eventarc: made `network_attachment` optional in `google_eventarc_pipeline` ([#10125](https://github.com/hashicorp/terraform-provider-google-beta/pull/10125))
* gemini: added `disable_web_grounding` field to `google_gemini_gemini_gcp_enablement_setting` resource ([#10115](https://github.com/hashicorp/terraform-provider-google-beta/pull/10115))
* gemini: added `enable_data_sharing` field to `google_gemini_data_sharing_with_google_setting` resource ([#10144](https://github.com/hashicorp/terraform-provider-google-beta/pull/10144))
* gkehub2: added `spec.rbacrolebindingactuation` field to resource `google_gke_hub_feature` ([#10121](https://github.com/hashicorp/terraform-provider-google-beta/pull/10121))
* gkehub: added `custom_role` field to `google_gke_hub_scope_rbac_role_binding` resource ([#10151](https://github.com/hashicorp/terraform-provider-google-beta/pull/10151))
* gkeonprem: added `private_registry_config` field to `google_gkeonprem_vmware_admin_cluster` resource ([#10150](https://github.com/hashicorp/terraform-provider-google-beta/pull/10150))
* iambeta: enforced `workload_identity_pool_managed_identity_id` field validation per the documented specifications ([#10132](https://github.com/hashicorp/terraform-provider-google-beta/pull/10132))
* pubsub: added `message_transform` field to `google_pubsub_topic` resource ([#10137](https://github.com/hashicorp/terraform-provider-google-beta/pull/10137))
* pubsub: added `message_transforms` field to `google_pubsub_subscription` resource ([#10138](https://github.com/hashicorp/terraform-provider-google-beta/pull/10138))

BUG FIXES:
* bigquery: modified `google_bigquery_dataset_iam_member`  to no longer remove authorized views and routines ([#10145](https://github.com/hashicorp/terraform-provider-google-beta/pull/10145))
* colab: fixed perma-diff in `google_colab_runtime_template` caused by the API returning a non-null default value. ([#10127](https://github.com/hashicorp/terraform-provider-google-beta/pull/10127))
* colab: fixed perma-diff in `google_colab_runtime_template` caused by empty blocks. ([#10139](https://github.com/hashicorp/terraform-provider-google-beta/pull/10139))
* compute: fixed a permadiff in `network_profile` field of `google_compute_network` related to specifying partial self-links ([#10140](https://github.com/hashicorp/terraform-provider-google-beta/pull/10140))
* compute: fixed an issue where `google_compute_firewall_policy_with_rules.target_resources` could see a diff between the beta and v1 API in the resource's self-link ([#10142](https://github.com/hashicorp/terraform-provider-google-beta/pull/10142))
* container: fixed nodepool secondary range validation to allow the use of netmasks. ([#10128](https://github.com/hashicorp/terraform-provider-google-beta/pull/10128))
* gemini: removed overly restrictive `product` validation on `google_gemini_gemini_gcp_enablement_setting_binding`, `google_gemini_data_sharing_with_google_setting_binding`. New values like `GOOGLE_CLOUD_ASSIST` will now be accepted. ([#10146](https://github.com/hashicorp/terraform-provider-google-beta/pull/10146))

## 6.38.0 (June 4, 2025)

DEPRECATIONS:
* colab: deprecated  `post_startup_script_config` field in `google_colab_runtime_template` resource ([#10104](https://github.com/hashicorp/terraform-provider-google-beta/pull/10104))

FEATURES:
* **New Data Source:** `google_bigquery_datasets` ([#10095](https://github.com/hashicorp/terraform-provider-google-beta/pull/10095))
* **New Resource:** `google_dataplex_entry` ([#10086](https://github.com/hashicorp/terraform-provider-google-beta/pull/10086))

IMPROVEMENTS:
* compute: added `candidate_cloud_router_ip_address`, `candidate_customer_router_ip_address`, `candidate_cloud_router_ipv6_address`, and `candidate_customer_router_ipv6_address` fields to `google_compute_interconnect_attachment` resource ([#10092](https://github.com/hashicorp/terraform-provider-google-beta/pull/10092))
* compute: added `httpFilterConfigs` and `httpFilterMetadata` fields in `google_compute_url_map` resource ([#10101](https://github.com/hashicorp/terraform-provider-google-beta/pull/10101))
* compute: added `numeric_id` to `google_compute_region_instance_template` resource ([#10098](https://github.com/hashicorp/terraform-provider-google-beta/pull/10098))
* compute: added `source_subnetwork_ip_ranges_to_nat64` and `nat64_subnetwork` fields  in `google_compute_router_nat` resource ([#10106](https://github.com/hashicorp/terraform-provider-google-beta/pull/10106))
* datastream: added `psc_interface_config` field in  `google_datastream_private_connection` resource ([#23091](https://github.com/hashicorp/terraform-provider-google/pull/23091))
* dns: added `dns64_config` field  to `google_dns_policy` resource ([#10106](https://github.com/hashicorp/terraform-provider-google-beta/pull/10106))
* filestore: added `effective_replication.role` and `effective_replication.replicas.peer_instance` fields to `google_filestore_instance` resource ([#10087](https://github.com/hashicorp/terraform-provider-google-beta/pull/10087))
* networkconnectivity: added `IPV6` enum to `protocol_version` field in `google_network_connectivity_policy_based_route` resource ([#10099](https://github.com/hashicorp/terraform-provider-google-beta/pull/10099))
* netapp: added `backup_retention_policy.backup_minimum_enforced_retention_days`, `backup_retention_policy.daily_backup_immutable`, `backup_retention_policy.weekly_backup_immutable`, `backup_retention_policy.monthly_backup_immutable`, and `backup_retention_policy.manual_backup_immutable` fields to `google_netapp_backup_vault` ([#10110](https://github.com/hashicorp/terraform-provider-google/pull/10110))
* privateca: added support for setting default values for basic constraints for `google_privateca_certificate_template` via the `null_ca` and `zero_max_issuer_path_length` fields ([#22981](https://github.com/hashicorp/terraform-provider-google/pull/22981))
* privateca: added `name_constraints` field for `google_privateca_certificate_template` resource ([#22981](https://github.com/hashicorp/terraform-provider-google/pull/22981))
* provider: supported service account impersonation in different universes through credential file ([#10097](https://github.com/hashicorp/terraform-provider-google-beta/pull/10097))

BUG FIXES:
* colab: fixed perma-diff in `google_colab_runtime_template` caused by the API returning a non-null default value ([#10127](https://github.com/hashicorp/terraform-provider-google-beta/pull/10127))
* compute: fixed an issue where rules ordering in `google_compute_region_security_policy` caused a diff after apply ([#10105](https://github.com/hashicorp/terraform-provider-google-beta/pull/10105))
* filestore: fixed bug where `google_filestore_instance.initial_replication` field could not be set ([#10087](https://github.com/hashicorp/terraform-provider-google-beta/pull/10087))

## 6.37.0 (May 27, 2025)

FEATURES:
* **New Data Source:** `google_bigquery_table` ([#10076](https://github.com/hashicorp/terraform-provider-google-beta/pull/10076))
* **New Data Source:** `google_gke_hub_membership` ([#10075](https://github.com/hashicorp/terraform-provider-google-beta/pull/10075))
* **New Resource:** `google_apigee_security_monitoring_condition` ([#10063](https://github.com/hashicorp/terraform-provider-google-beta/pull/10063))
* **New Resource:** `google_beyondcorp_security_gateway_application` ([#10059](https://github.com/hashicorp/terraform-provider-google-beta/pull/10059))
* **New Resource:** `google_cloud_run_v2_worker_pool` ([#10054](https://github.com/hashicorp/terraform-provider-google-beta/pull/10054))
* **New Resource:** `google_compute_future_reservation` ([#10020](https://github.com/hashicorp/terraform-provider-google-beta/pull/10020))
* **New Resource:** `google_dataplex_glossary_category` ([#10016](https://github.com/hashicorp/terraform-provider-google-beta/pull/10016))
* **New Resource:** `google_dataplex_glossary_term` ([#10016](https://github.com/hashicorp/terraform-provider-google-beta/pull/10016))
* **New Resource:** `google_iam_workforce_pool_provider_key` ([#10070](https://github.com/hashicorp/terraform-provider-google-beta/pull/10070))
* **New Resource:** `google_iam_workload_identity_pool_managed_identity` ([#10081](https://github.com/hashicorp/terraform-provider-google-beta/pull/10081))
* **New Resource:** `google_iam_workload_identity_pool_namespace` ([#10044](https://github.com/hashicorp/terraform-provider-google-beta/pull/10044))
* **New Resource:** `google_managed_kafka_acl` ([#10067](https://github.com/hashicorp/terraform-provider-google-beta/pull/10067))

IMPROVEMENTS:
* alloydb: added `activation_policy` field to `google_alloydb_instance` resource ([#10010](https://github.com/hashicorp/terraform-provider-google-beta/pull/10010))
* compute: added `mirror_percent` field to `default_route_action.request_mirror_policy`, `path_matchers.default_route_action.request_mirror_policy`, `path_matchers.path_rules.route_action.request_mirror_policy`, and `path_matchers.route_rules.route_action.request_mirror_policy` in `google_compute_url_map` resource ([#10071](https://github.com/hashicorp/terraform-provider-google-beta/pull/10071))
* compute: added `network_pass_through_lb_traffic_policy.0.zonal_affinity.0.spillover`, `network_pass_through_lb_traffic_policy.0.zonal_affinity.0.spillover_ratio` and `dynamic_forwarding.0.ip_port_selection.0.enabled` to `google_compute_backend_service` resource ([#10056](https://github.com/hashicorp/terraform-provider-google-beta/pull/10056))
* compute: added in-place update support for `mtu` field in `google_compute_network` ([#10066](https://github.com/hashicorp/terraform-provider-google-beta/pull/10066))
* compute: added `subsetting.0.subset_size` and `dynamic_forwarding.0.ip_port_selection.0.enabled` to `google_compute_region_backend_service` resource ([#10056](https://github.com/hashicorp/terraform-provider-google-beta/pull/10056))
* container: added in-place update support for `ip_allocation_policy.stack_type` field in `google_container_cluster` resource ([#10037](https://github.com/hashicorp/terraform-provider-google-beta/pull/10037))
* container: added in-place update support for `enable_multi_networking` in `google_container_cluster` resource ([#10045](https://github.com/hashicorp/terraform-provider-google-beta/pull/10045))
* databasemigrationservice: added `create_without_validation` field to `google_database_migration_service_private_connection` resource ([#10046](https://github.com/hashicorp/terraform-provider-google-beta/pull/10046))
* dataflow: added `additional_pipeline_options` field to `google_dataflow_flex_template_job` resource ([#10040](https://github.com/hashicorp/terraform-provider-google-beta/pull/10040))
* filestore: added PSC fields to `google_filestore_instance` ([#10061](https://github.com/hashicorp/terraform-provider-google-beta/pull/10061))
* memorystore: added field `desired_auto_created_endpoints` for `google_memorystore_instance` resource ([#10031](https://github.com/hashicorp/terraform-provider-google-beta/pull/10031))
* netapp: added `hybrid_peering_details` and `hybrid_replication_type` fields to `google_netapp_volume_replication` resource ([#10077](https://github.com/hashicorp/terraform-provider-google-beta/pull/10077))
* netapp: added `hybrid_replication_parameters` fields to `google_netapp_volume` resource ([#10077](https://github.com/hashicorp/terraform-provider-google-beta/pull/10077))
* netblock: added `restricted-googleapis-with-directconnectivity` and `private-googleapis-with-directconnectivity` range_types to `google_netblock_ip_ranges` data source ([#10051](https://github.com/hashicorp/terraform-provider-google-beta/pull/10051))
* netblock: added ipv6 ranges for `restricted-googleapis` and `private-googleapis` range_types to `google_netblock_ip_ranges` data source ([#10051](https://github.com/hashicorp/terraform-provider-google-beta/pull/10051))
* privateca: added `name_constraints` field for `google_privateca_certificate_template` resource ([#10083](https://github.com/hashicorp/terraform-provider-google-beta/pull/10083))
* spanner: added field `instance_type` to the `google_spanner_instance` resource ([#10038](https://github.com/hashicorp/terraform-provider-google-beta/pull/10038))
* storage: added `ip_filter` to `google_storage_bucket` resource. ([#10078](https://github.com/hashicorp/terraform-provider-google-beta/pull/10078))

BUG FIXES:
* gemini: fixed permadiff on `product` field in `google_gemini_logging_setting_binding ` resource ([#10011](https://github.com/hashicorp/terraform-provider-google-beta/pull/10011))
* gemini: fixed permadiff on `product` field in `google_gemini_release_channel_setting_binding ` resource ([#10050](https://github.com/hashicorp/terraform-provider-google-beta/pull/10050))
* networkservices: fixed validation error when modifying the `cache_mode` field in `edge_cache_service` ([#10053](https://github.com/hashicorp/terraform-provider-google-beta/pull/10053))
* privateca: fixed issue preventing setting `0` and null values for basic constraints in the `google_privateca_certificate_template` resource via the addition of `null_ca` and `zero_max_issuer_path_length` fields ([#10083](https://github.com/hashicorp/terraform-provider-google-beta/pull/10083))
* vpcaccess: fixed an issue where Terraform config validation conditions could have erroneously invalidated existing `google_vpc_access_connector` resources ([#10018](https://github.com/hashicorp/terraform-provider-google-beta/pull/10018))

## 6.36.1 (May 21, 2025)

BUG FIXES: 
* compute: fixed forced instance recreation when adding a `attached_disk` with unset `force_attach` to `google_compute_instance` ([#10064](https://github.com/hashicorp/terraform-provider-google-beta/pull/10064))

## 6.36.0 (May 20, 2025)

* DEPRECATIONS:
* beyondcorp: deprecated `google_beyondcorp_application` ([#9968](https://github.com/hashicorp/terraform-provider-google-beta/pull/9968))
* firestore: deprecated `deletion_policy` field of `google_firestore_database` resource ([#9976](https://github.com/hashicorp/terraform-provider-google-beta/pull/9976))

FEATURES:
* **New Data Source:** `google_beyondcorp_security_gateway` ([#9996](https://github.com/hashicorp/terraform-provider-google-beta/pull/9996))
* **New Data Source:** `google_lustre_instance` ([#9978](https://github.com/hashicorp/terraform-provider-google-beta/pull/9978))
* **New Resource:** `google_bigquery_row_access_policy` ([#10004](https://github.com/hashicorp/terraform-provider-google-beta/pull/10004))
* **New Resource:** `google_dataplex_glossary` ([#9997](https://github.com/hashicorp/terraform-provider-google-beta/pull/9997))
* **New Resource:** `google_firebase_app_hosting_default_domain` ([#9966](https://github.com/hashicorp/terraform-provider-google-beta/pull/9966))
* **New Resource:** `google_firebase_app_hosting_domain` ([#9966](https://github.com/hashicorp/terraform-provider-google-beta/pull/9966))
* **New Resource:** `google_firebase_app_hosting_traffic` ([#9966](https://github.com/hashicorp/terraform-provider-google-beta/pull/9966))
* **New Resource:** `google_iam_workload_identity_pool_iam_*` ([#9990](https://github.com/hashicorp/terraform-provider-google-beta/pull/9990))

IMPROVEMENTS:
* beyondcorp: increased default timeouts on `google_beyondcorp_app_gateway ` operations from 20m to 40m ([#10003](https://github.com/hashicorp/terraform-provider-google-beta/pull/10003))
* bigtable: added `deletion_protection` field to `google_bigtable_logical_view` resource ([#9969](https://github.com/hashicorp/terraform-provider-google-beta/pull/9969))
* compute: added 'H2C' as a supported value for `protocol` in `google_compute_backend_service` and `google_compute_region_backend_service` ([#9994](https://github.com/hashicorp/terraform-provider-google-beta/pull/9994))
* compute: added `external_managed_backend_bucket_migration_state` and `external_managed_backend_bucket_migration_testing_percentage` to `google_compute_global_forwarding_rule` resource. ([#9985](https://github.com/hashicorp/terraform-provider-google-beta/pull/9985))
* compute: added `external_managed_migration_state` and `external_managed_migration_testing_percentage` to `google_compute_backend_service` resource. ([#9985](https://github.com/hashicorp/terraform-provider-google-beta/pull/9985))
* compute: added `force_attach` field to `boot_disk` and  `attached_disk` of  `google_compute_instance`  resource ([#9999](https://github.com/hashicorp/terraform-provider-google-beta/pull/9999))
* compute: added `numeric_id` to `google_compute_instance_template` resource ([#9975](https://github.com/hashicorp/terraform-provider-google-beta/pull/9975))
* compute: added the numeric id as `generated_id` attribute to the `google_compute_network_endpoint_group` ([#9984](https://github.com/hashicorp/terraform-provider-google-beta/pull/9984))
* compute: added update support for `load_balancing_scheme` in `google_compute_backend_service` and `google_compute_global_forwarding_rule` resources to allow migrating between classic and global external ALB ([#9985](https://github.com/hashicorp/terraform-provider-google-beta/pull/9985))
* container: added `in_transit_encryption_config` field in `google_container_cluster` resource ([#9972](https://github.com/hashicorp/terraform-provider-google-beta/pull/9972))
* container: allowed in-place update `node_config.windows_node_config` field in `google_container_cluster` and `google_container_node_pool` resource ([#9986](https://github.com/hashicorp/terraform-provider-google-beta/pull/9986))
* container: allowed in-place update for `node_config.storage_pools` field in `google_container_cluster` and `google_container_node_pool` resourcee ([#9967](https://github.com/hashicorp/terraform-provider-google-beta/pull/9967))
* dialogflowcx: added `event_handlers.trigger_fulfillment.enable_generative_fallback` field to `google_dialogflow_cx_flow` resource ([#9958](https://github.com/hashicorp/terraform-provider-google-beta/pull/9958))
* dialogflowcx: added `gen_app_builder_settings` field to `google_dialogflow_cx_agent` resource ([#9971](https://github.com/hashicorp/terraform-provider-google-beta/pull/9971))
* iambeta: added `mode`, `inline_certificate_issuance_config`, and `inline_trust_config` fields to `google_iam_workload_identity_pool` resource ([#9990](https://github.com/hashicorp/terraform-provider-google-beta/pull/9990))
* vmwareengine: increased `google_cloud_vmwareengine_private_cloud` timeout to 6 hours. ([#9974](https://github.com/hashicorp/terraform-provider-google-beta/pull/9974))

BUG FIXES:
* compute: added global retry for "resourceNotReady for Networks" 400 errors ([#9970](https://github.com/hashicorp/terraform-provider-google-beta/pull/9970))
* dialogflowcx: fixed an issue where `dialogflow_cx_custom_endpoint` is not correctedly handled ([#9995](https://github.com/hashicorp/terraform-provider-google-beta/pull/9995))
* iamoauthclient: marked `google_iam_oauth_client_credential.client_secret` as sensitive ([#9992](https://github.com/hashicorp/terraform-provider-google-beta/pull/9992))
* resourcemanager: fixed an issue in `google_projects` data source where the provider `universe_domain` did not overwrite the list URL ([#9964](https://github.com/hashicorp/terraform-provider-google-beta/pull/9964))

## 6.35.0 (May 13, 2025)

FEATURES:
* **New Resource:** `google_compute_cross_site_network` ([#9940](https://github.com/hashicorp/terraform-provider-google-beta/pull/9940))

IMPROVEMENTS:
* alloydb: added `psc_auto_connections` field to `google_alloydb_instance` resource ([#9938](https://github.com/hashicorp/terraform-provider-google-beta/pull/9938))
* apigee: added `s_sl_info.enforce` field in `google_apigee_target_server` resource ([#9922](https://github.com/hashicorp/terraform-provider-google-beta/pull/9922))
* bigquery: added `security_mode` option for `google_bigquery_routine` resource ([#9949](https://github.com/hashicorp/terraform-provider-google-beta/pull/9949))
* bigtable: added support for explicit disable automated backup on create for `google_bigtable_table` ([#9943](https://github.com/hashicorp/terraform-provider-google-beta/pull/9943))
* compute: added `guest_os_features` and `architecture` to `google_compute_instance_template` and `google_compute_region_instance_template` ([#9950](https://github.com/hashicorp/terraform-provider-google-beta/pull/9950))
* compute: added `grpc_tls_health_check` field to `google_compute_healthcheck` resource ([#9924](https://github.com/hashicorp/terraform-provider-google-beta/pull/9924))
* compute: allowed in-place updates for `subnetworks`, `description`, `producer_accept_lists`, and `producer_reject_lists` on `google_compute_network_attachment` ([#9926](https://github.com/hashicorp/terraform-provider-google-beta/pull/9926))
* dialogflowcx: added `knowledge_connector_settings` field to `google_dialogflow_cx_flow` and `google_dialogflow_cx_page` resources ([#9939](https://github.com/hashicorp/terraform-provider-google-beta/pull/9939))
* filestore: added `directory_services` field to `google_filestore_instance` ([#9919](https://github.com/hashicorp/terraform-provider-google-beta/pull/9919))
* netapp: added `backup_vault_type`, `backup_region`, `source_region`, `source_backup_vault`, and `destination_backup_vault` fields to `google_netapp_backup_vault` ([#9933](https://github.com/hashicorp/terraform-provider-google-beta/pull/9933))
* netapp: added `volume_region` and `backup_region` fields to `google_netapp_backup` ([#9933](https://github.com/hashicorp/terraform-provider-google-beta/pull/9933))
* networkconnectivity: added `immutability` field to `google_network_connectivity_internal_range` resource ([#9931](https://github.com/hashicorp/terraform-provider-google-beta/pull/9931))
* networkservices: added `flex_shielding` field to `google_network_services_edge_cache_origin` resource ([#9951](https://github.com/hashicorp/terraform-provider-google-beta/pull/9951))
* spanner: added field `default_time_zone` to `google_spanner_database` resource ([#9936](https://github.com/hashicorp/terraform-provider-google-beta/pull/9936))
* storage: added new field `content_hexsha512` and `content_base64sha512` in data source `google_storage_bucket_object_content` ([#9920](https://github.com/hashicorp/terraform-provider-google-beta/pull/9920))

BUG FIXES:
* gemini: fixed bug on `google_gemini_code_repository_index` where `force_destroy` field did nothing. ([#9952](https://github.com/hashicorp/terraform-provider-google-beta/pull/9952))
* privateca: removed requirement to specify `organization` for `google_privateca_certificate_authority` resource ([#9942](https://github.com/hashicorp/terraform-provider-google-beta/pull/9942))
* workbench: fixed some metadata changes not being reflected in `google_workbench_instance` ([#9927](https://github.com/hashicorp/terraform-provider-google-beta/pull/9927))

## 6.34.1 (May 12, 2025)

BUG FIXES:
* bigtable: fixed forced instance recreation due to addition of `cluster.node_scaling_factor` for `google_bigtable_instance` ([#9961](https://github.com/hashicorp/terraform-provider-google-beta/pull/9961))

## 6.34.0 (May 6, 2025)
DEPRECATIONS:
* tpu: deprecated `google_tpu_node` resource. `google_tpu_node` is deprecated and will be removed in a future major release. Use `google_tpu_v2_vm` instead. ([#9902](https://github.com/hashicorp/terraform-provider-google-beta/pull/9902))

FEATURES:
* **New Resource:** `google_apigee_security_profile_v2` ([#9895](https://github.com/hashicorp/terraform-provider-google-beta/pull/9895))
* **New Resource:** `google_resource_manager_capability` ([#9917](https://github.com/hashicorp/terraform-provider-google-beta/pull/9917))

IMPROVEMENTS:
* bigtable: added `cluster.node_scaling_factor` field to `google_bigtable_instance` resource ([#9907](https://github.com/hashicorp/terraform-provider-google-beta/pull/9907))
* cloudrunv2: added `scaling_mode` and `manual_instance_count` fields to `google_cloud_run_v2_service` resource ([#9908](https://github.com/hashicorp/terraform-provider-google-beta/pull/9908))
* filestore: added `directory_services` field to `google_filestore_instance` (beta) ([#9919](https://github.com/hashicorp/terraform-provider-google-beta/pull/9919))
* networkconnectivity: added `state_reason` field to `google_network_connectivity_spoke ` resource ([#9896](https://github.com/hashicorp/terraform-provider-google-beta/pull/9896))
* sql: added `connection_pool_config` field to the google_sql_database_instance resource ([#9918](https://github.com/hashicorp/terraform-provider-google-beta/pull/9918))
* vpcaccess: changed fields `min_instances`, `max_instances`, `machine_type` to allow update `google_vpc_access_connector` without without recreation. ([#9914](https://github.com/hashicorp/terraform-provider-google-beta/pull/9914))

BUG FIXES:
* compute: fixed the bug when validating the subnetwork project in `google_compute_instance` resource ([#9913](https://github.com/hashicorp/terraform-provider-google-beta/pull/9913))
* workbench: fixed a permadiff on `metadata` of `instance-region` in `google_workbench_instance` resource ([#9903](https://github.com/hashicorp/terraform-provider-google-beta/pull/9903))

## 6.33.0 (Apr 29, 2025)
FEATURES:
* **New Data Source:** `google_memcache_instance` ([#9864](https://github.com/hashicorp/terraform-provider-google-beta/pull/9864))
* **New Resource:** `google_bigtable_logical_view` ([#9876](https://github.com/hashicorp/terraform-provider-google-beta/pull/9876))
* **New Resource:** `google_bigtable_materialized_view` ([#9862](https://github.com/hashicorp/terraform-provider-google-beta/pull/9862))
* **New Resource:** `google_os_config_v2_policy_orchestrator_for_folder` ([#9841](https://github.com/hashicorp/terraform-provider-google-beta/pull/9841))

IMPROVEMENTS:
* beyondcorp: Added `upstreams` fields to `google_beyondcorp_application` resource ([#9890](https://github.com/hashicorp/terraform-provider-google-beta/pull/9890))
* compute: Added fields like `raw_key`, `rsa_encrypted_key`, `kms_key_service_account` to all relevant resources on `google_compute_instance_template` and `google_compute_region_instance_template` ([#9880](https://github.com/hashicorp/terraform-provider-google-beta/pull/9880))
* compute: added `disk_id` to `google_compute_region_disk` resource ([#9855](https://github.com/hashicorp/terraform-provider-google-beta/pull/9855))
* compute: marked `location` field as required in `google_compute_interconnect` resource ([#9865](https://github.com/hashicorp/terraform-provider-google-beta/pull/9865))
* container: added `data_cache_count` to `ephemeral_storage_local_ssd_config` for `google_container_node_pool` ([#9851](https://github.com/hashicorp/terraform-provider-google-beta/pull/9851))
* container: added update for `gvnic` to `google_container_node_pool` ([#9834](https://github.com/hashicorp/terraform-provider-google-beta/pull/9834))
* dataplex: added `notification_report` field to `google_dataplex_datascan` resource ([#9857](https://github.com/hashicorp/terraform-provider-google-beta/pull/9857))
* dns: added `target_name_servers.domainName` to resource`google_dns_managed_zone ` (beta) ([#9832](https://github.com/hashicorp/terraform-provider-google-beta/pull/9832))
* gkehub: added `configmanagement.config_sync.deployment_overrides` field to `google_gke_hub_feature_membership` resource ([#9828](https://github.com/hashicorp/terraform-provider-google-beta/pull/9828))
* identityplatform: Added `response_type` field to `google_identity_platform_oauth_idp_config` ([#9856](https://github.com/hashicorp/terraform-provider-google-beta/pull/9856))
* netapp: added `custom_performance_enabled`, `total_throughput_mibps`, and `total_iops` fields to `google_netapp_storage_pool` resource (beta) ([#9872](https://github.com/hashicorp/terraform-provider-google-beta/pull/9872))
* networkservices: added `metadata` field to `google_networkservices_lbtrafficextension` resource ([#9849](https://github.com/hashicorp/terraform-provider-google-beta/pull/9849))
* sql: added output-only field `dns_names` to `google_sql_database_instance` resource ([#9879](https://github.com/hashicorp/terraform-provider-google-beta/pull/9879))
* storage: added new fields `time_created` and `updated` in `google_storage_bucket` ([#9877](https://github.com/hashicorp/terraform-provider-google-beta/pull/9877))
* storagetransfer: added `transfer_spec.aws_s3_data_source.managed_private_network` field to `google_storage_transfer_job` resource ([#9886](https://github.com/hashicorp/terraform-provider-google-beta/pull/9886))

BUG FIXES:
* alloydb: stopped diffs when `google_alloydb_instance.network_config` is not specified as the API newly returns a value. Removing the field from config will no longer create a diff and will preserve the current value ([#9881](https://github.com/hashicorp/terraform-provider-google-beta/pull/9881))
* clouddeploy: allowed sending empty block for `rollback` field  in  `google_clouddeploy_automation` resource. ([#9878](https://github.com/hashicorp/terraform-provider-google-beta/pull/9878))
* compute: fixed an issue preventing `terms.priority` from being set to priority value 0 in `google_compute_router_route_policy` resource ([#9830](https://github.com/hashicorp/terraform-provider-google-beta/pull/9830))
* securesourcemanager: increased default timeouts on `google_secure_source_manager_instance` operations to 120m from 60m. Operations could take longer than an hour. ([#9868](https://github.com/hashicorp/terraform-provider-google-beta/pull/9868))
* sql: replaced the Terraform-based default value for `settings.disk_type` in `google_sql_database_instance` with a server-assigned default, allowing for compatibility with machine types that require `HyperDisk_Balanced` ([#9870](https://github.com/hashicorp/terraform-provider-google-beta/pull/9870))
* workstations: increased default timeouts on `google_workstations_workstation_cluster` operations to 120m from 60m. Operations could take longer than an hour. ([#9867](https://github.com/hashicorp/terraform-provider-google-beta/pull/9867))

## 6.32.0 (Apr 25, 2025)

IMPROVEMENTS:
* container: added `flex_start` to `node_config` in `google_container_cluster` and `google_container_node_pool` ([#9885](https://github.com/hashicorp/terraform-provider-google-beta/pull/9885))

## 6.31.1 (Apr 25, 2025)

BUG FIXES:
* storage: removed extra permission (storage.anywhereCaches.list) required for destroying a `resource_storage_bucket` ([#9842](https://github.com/hashicorp/terraform-provider-google-beta/pull/9842))

## 6.31.0 (Apr 22, 2025)

DEPRECATIONS:
* integrations: deprecated `run_as_service_account` field in `google_integrations_client` resource ([#9767](https://github.com/hashicorp/terraform-provider-google-beta/pull/9767))

FEATURES:
* **New Resource:** `google_compute_resource_policy_attachment` ([#9824](https://github.com/hashicorp/terraform-provider-google-beta/pull/9824))
* **New Resource:** `google_compute_storage_pool` ([#9786](https://github.com/hashicorp/terraform-provider-google-beta/pull/9786))
* **New Resource:** `google_gke_backup_backup_channel` ([#9819](https://github.com/hashicorp/terraform-provider-google-beta/pull/9819))
* **New Resource:** `google_gke_backup_restore_channel` ([#9819](https://github.com/hashicorp/terraform-provider-google-beta/pull/9819))
* **New Resource:** `google_iap_web_cloud_run_service_iam_binding` ([#9823](https://github.com/hashicorp/terraform-provider-google-beta/pull/9823))
* **New Resource:** `google_iap_web_cloud_run_service_iam_member` ([#9823](https://github.com/hashicorp/terraform-provider-google-beta/pull/9823))
* **New Resource:** `google_iap_web_cloud_run_service_iam_policy` ([#9823](https://github.com/hashicorp/terraform-provider-google-beta/pull/9823))
* **New Resource:** `google_storage_batch_operations_job` ([#9779](https://github.com/hashicorp/terraform-provider-google-beta/pull/9779))

IMPROVEMENTS:
* accesscontextmanager: added `scoped_access_settings` field to `gcp_user_access_binding` resource ([#9763](https://github.com/hashicorp/terraform-provider-google-beta/pull/9763))
* alloydb: added `assistive_experiences_enabled` field to `observabilityConfig` in `google_alloydb_instance` resource ([#9808](https://github.com/hashicorp/terraform-provider-google-beta/pull/9808))
* alloydb: added `machine_type` field to `google_alloydb_instance` resource ([#9795](https://github.com/hashicorp/terraform-provider-google-beta/pull/9795))
* artifactregistry: added `DEBIAN_SNAPSHOT` enum value to `repository_base` in `google_artifact_registry_repository` ([#9770](https://github.com/hashicorp/terraform-provider-google-beta/pull/9770))
* compute: added `log_config.optional_mode`, `log_config.optional_fields`, `backend.preference`, `max_stream_duration` and `cdn_policy.request_coalescing` fields to `google_compute_backend_service` resource ([#9818](https://github.com/hashicorp/terraform-provider-google-beta/pull/9818))
* container: added support for updating the `confidential_nodes` field in `google_container_node_pool` ([#9804](https://github.com/hashicorp/terraform-provider-google-beta/pull/9804))
* discoveryengine: added `allow_cross_region` field to `google_discovery_engine_chat_engine` resource ([#9782](https://github.com/hashicorp/terraform-provider-google-beta/pull/9782))
* gkehub: added `configmanagement.config_sync.deployment_overrides` field to `google_gke_hub_feature_membership` resource ([#9828](https://github.com/hashicorp/terraform-provider-google-beta/pull/9828))
* kms: added new enum values for `import_method` field in  `google_kms_key_ring_import_job` resource ([#9769](https://github.com/hashicorp/terraform-provider-google-beta/pull/9769))
* metastore: added `tags` field to `google_dataproc_metastore_service` resource to allow setting tags for services at creation time ([#9768](https://github.com/hashicorp/terraform-provider-google-beta/pull/9768))
* monitoring: added `log_check_failures` to `google_monitoring_uptime_check_config` ([#9794](https://github.com/hashicorp/terraform-provider-google-beta/pull/9794))
* networkconnectivity: added IPv6 support to `google_network_connectivity_internal_range` resource ([#9826](https://github.com/hashicorp/terraform-provider-google-beta/pull/9826))
* networkconnectivity: added `exclude_cidr_ranges` field to `google_network_connectivity_internal_range` resource ([#9778](https://github.com/hashicorp/terraform-provider-google-beta/pull/9778))
* privateca: added `backdate_duration` field to the `google_privateca_ca_pool` resource to add support for backdating the `not_before_time` of certificates ([#9812](https://github.com/hashicorp/terraform-provider-google-beta/pull/9812))
* redis: added `tags` field to `google_redis_instance` ([#9783](https://github.com/hashicorp/terraform-provider-google-beta/pull/9783))
* sql: added `custom_subject_alternative_names` field to `instances` resource ([#9799](https://github.com/hashicorp/terraform-provider-google-beta/pull/9799))
* sql: added `data_disk_provisioned_iops` and `data_disk_provisioned_throughput` fields to `google_sql_database_instance` resource ([#9822](https://github.com/hashicorp/terraform-provider-google-beta/pull/9822))
* sql: added `retain_backups_on_delete` field to `google_sql_database_instance` resource ([#9780](https://github.com/hashicorp/terraform-provider-google-beta/pull/9780))

BUG FIXES:
* colab: fixed perma-diff in `google_colab_runtime_template` caused by not returning default values. ([#9784](https://github.com/hashicorp/terraform-provider-google-beta/pull/9784))
* discoveryengine: fixed `google_discovery_engine_target_site` operations to allow for enough time to index before timing out ([#9800](https://github.com/hashicorp/terraform-provider-google-beta/pull/9800))
* compute: fixed perma-diff in `google_compute_network_firewall_policy_rule` when `security_profile_group` starts with `//` ([#9827](https://github.com/hashicorp/terraform-provider-google-beta/pull/9827))
* healthcare: made `google_healthcare_pipeline_job` wait for creation and update operation to complete ([#9785](https://github.com/hashicorp/terraform-provider-google-beta/pull/9785))
* identityplatform: fixed perma-diff in `google_identity_platform_config` when fields in `blocking_functions.forward_inbound_credentials` are set to `false` ([#9814](https://github.com/hashicorp/terraform-provider-google-beta/pull/9814))
* sql: added diff suppression for some version changes to`google_sql_database_instance`. Diffs for `database_version` for MySQL 8.0 will be suppressed when the version is updated by auto version upgrade.([#22356](https://github.com/hashicorp/terraform-provider-google/pull/22356))
* sql: fixed the issue of shortened version of failover_dr_replica_name causing unnecessary diff in `google_sql_database_instance` ([#9775](https://github.com/hashicorp/terraform-provider-google-beta/pull/9775))

## 6.30.0 (Apr 15, 2025)

FEATURES:
* **New Resource:** `google_developer_connect_account_connector` ([#9741](https://github.com/hashicorp/terraform-provider-google-beta/pull/9741))
* **New Resource:** `google_vertex_ai_feature_group_iam_*` ([#9735](https://github.com/hashicorp/terraform-provider-google-beta/pull/9735))
* **New Resource:** `google_vertex_ai_feature_online_store_iam_*` ([#9735](https://github.com/hashicorp/terraform-provider-google-beta/pull/9735))
* **New Resource:** `google_vertex_ai_feature_online_store_featureview_iam_*` ([#9735](https://github.com/hashicorp/terraform-provider-google-beta/pull/9735))

IMPROVEMENTS:
* cloudrunv2: added `iap_enabled` field to `google_cloud_run_v2_service` resource ([#9758](https://github.com/hashicorp/terraform-provider-google-beta/pull/9758))
* compute: added `source_disk_encryption_key.kms_key_self_link` and `source_disk_encryption_key.rsa_encrypted_key` fields to `google_compute_snapshot` resource ([#9730](https://github.com/hashicorp/terraform-provider-google-beta/pull/9730))
* compute: added `source_disk_encryption_key`, `source_image_encryption_key` and `source_snapshot_encryption_key` fields to `google_compute_image` resource ([#9730](https://github.com/hashicorp/terraform-provider-google-beta/pull/9730))
* databasemigrationservice: added `ssl.type` field to `google_database_migration_service_connection_profile` resource ([#9739](https://github.com/hashicorp/terraform-provider-google-beta/pull/9739))
* firestore: added `MONGODB_COMPATIBLE_API` enum option to `api_scope` field in `google_firestore_index` resource ([#9750](https://github.com/hashicorp/terraform-provider-google-beta/pull/9750))
* firestore: added `database_edition` field to `google_firestore_database` resource ([#9750](https://github.com/hashicorp/terraform-provider-google-beta/pull/9750))
* firestore: added `density` and `multikey` fields to `google_firestore_index` resource ([#9750](https://github.com/hashicorp/terraform-provider-google-beta/pull/9750))
* memorystore: added `managed_backup_source` and `gcs_source` fields to `google_memorystore_instance` resource ([#9753](https://github.com/hashicorp/terraform-provider-google-beta/pull/9753))
* monitoring: added `password_wo` write-only field and `password_wo_version` field to `google_monitoring_uptime_check_config` resource ([#9727](https://github.com/hashicorp/terraform-provider-google-beta/pull/9727))
* redis: added `managed_backup_source` and `gcs_source` fields to `google_redis_cluster` resource ([#9745](https://github.com/hashicorp/terraform-provider-google-beta/pull/9745))
* storage: added support for deleting pending caches present on bucket when setting `force_destory` to true in `google_storage_bucket` resource ([#9737](https://github.com/hashicorp/terraform-provider-google-beta/pull/9737))
* storagecontrol: added `trial_config` field to `google_storage_control_folder_intelligence_config` resource ([#9724](https://github.com/hashicorp/terraform-provider-google-beta/pull/9724))
* storagecontrol: added `trial_config` field to `google_storage_control_organization_intelligence_config` resource ([#9724](https://github.com/hashicorp/terraform-provider-google-beta/pull/9724))
* storagecontrol: added `trial_config` field to `google_storage_control_project_intelligence_config` resource ([#9724](https://github.com/hashicorp/terraform-provider-google-beta/pull/9724))

BUG FIXES:
* container: fixed perma-diff in `fleet` field when the `fleet.project` field being added is null or empty in `google_container_cluster` resource ([#9726](https://github.com/hashicorp/terraform-provider-google-beta/pull/9726))
* pubsub: fixed perma-diff by changing `allowed_persistence_regions` field to set in `google_pubsub_topic` resource ([#9743](https://github.com/hashicorp/terraform-provider-google-beta/pull/9743))

## 6.29.0 (Apr 8, 2025)

FEATURES:
* **New Resource:** `google_apigee_control_plane_access` ([#9709](https://github.com/hashicorp/terraform-provider-google-beta/pull/9709))
* **New Resource:** `google_clouddeploy_deploy_policy` ([#9694](https://github.com/hashicorp/terraform-provider-google-beta/pull/9694))
* **New Resource:** `google_folder_service_identity` ([#9703](https://github.com/hashicorp/terraform-provider-google-beta/pull/9703))
* **New Resource:** `google_os_config_v2_policy_orchestrator_for_organization` ([#9696](https://github.com/hashicorp/terraform-provider-google-beta/pull/9696))

IMPROVEMENTS:
* accesscontextmanager: added `session_settings` field to `gcp_user_access_binding` resource ([#9720](https://github.com/hashicorp/terraform-provider-google-beta/pull/9720))
* cloudedeploy: added `timed_promote_release_rule` and `repair_rollout_rule` fields to `google_clouddeploy_automation` resource ([#9694](https://github.com/hashicorp/terraform-provider-google-beta/pull/9694))
* compute: added `group_placement_policy.0.tpu_topology` field to `google_compute_resource_policy` resource. ([#9702](https://github.com/hashicorp/terraform-provider-google-beta/pull/9702))
* datastream: added support for creating streams for Salesforce source in `google_datastream_stream`. ([#9706](https://github.com/hashicorp/terraform-provider-google-beta/pull/9706))
* gkeonprem: added `enable_advanced_cluster` field to `google_gkeonprem_vmware_admin_cluster` resource ([#9693](https://github.com/hashicorp/terraform-provider-google-beta/pull/9693))
* gkeonprem: added `enable_advanced_cluster` field to `google_gkeonprem_vmware_cluster` resource ([#9693](https://github.com/hashicorp/terraform-provider-google-beta/pull/9693))
* memorystore: added `automated_backup_config` field to `google_memorystore_instance` resource ([#9708](https://github.com/hashicorp/terraform-provider-google-beta/pull/9708))
* netapp: added `tiering_policy` to `google_netapp_volume_replication` resource ([#9716](https://github.com/hashicorp/terraform-provider-google-beta/pull/9716))
* parametermanagerregional: added `kms_key_version` field to `google_parameter_manager_regional_parameter_version` resource and datasource ([#9712](https://github.com/hashicorp/terraform-provider-google-beta/pull/9712))
* parametermanagerregional: added `kms_key` field to `google_parameter_manager_regional_parameter` resource and `google_parameter_manager_regional_parameters` datasource ([#9712](https://github.com/hashicorp/terraform-provider-google-beta/pull/9712))
* redis: added `automated_backup_config` field to `google_redis_cluster` ([#9682](https://github.com/hashicorp/terraform-provider-google-beta/pull/9682))
* storage: added `md5hexhash` field in `google_storage_bucket_object` ([#9722](https://github.com/hashicorp/terraform-provider-google-beta/pull/9722))
* workbench: added `confidential_instance_config` field to `google_workbench_instance` resource ([#9688](https://github.com/hashicorp/terraform-provider-google-beta/pull/9688))

BUG FIXES:
* colab: fixed an issue where `google_colab_*` resources incorrectly required a provider-level region matching the resource location ([#9714](https://github.com/hashicorp/terraform-provider-google-beta/pull/9714))
* datastream: updated `private_key`to be mutable in `google_datastream_connection_profile` resource. ([#9689](https://github.com/hashicorp/terraform-provider-google-beta/pull/9689))
* gkehub: enabled partial results to be returned when a cloud region is unreachable in `google_gke_hub_feature ` ([#9715](https://github.com/hashicorp/terraform-provider-google-beta/pull/9715))

## 6.28.0 (Apr 1, 2025)

DEPRECATIONS:
* compute: deprecated `enable_flow_logs` in favor of `log_config` on `google_compute_subnetwork` resource.  If `log_config` is present, flow logs are enabled, and `enable_flow_logs` can be safely removed. ([#9679](https://github.com/hashicorp/terraform-provider-google-beta/pull/9679))
* containerregistry: Deprecated `google_container_registry` resource, and `google_container_registry_image` and `google_container_registry_repository` data sources. Use `google_artifact_registry_repository` instead. ([#9650](https://github.com/hashicorp/terraform-provider-google-beta/pull/9650))

FEATURES:
* **New Data Source:** `google_compute_region_backend_service` ([#9616](https://github.com/hashicorp/terraform-provider-google-beta/pull/9616))
* **New Data Source:** `google_organization_iam_custom_roles` ([#9628](https://github.com/hashicorp/terraform-provider-google-beta/pull/9628))
* **New Data Source:** `google_storage_control_folder_intelligence_config` ([#9655](https://github.com/hashicorp/terraform-provider-google-beta/pull/9655))
* **New Data Source:** `google_storage_control_organization_intelligence_config` ([#9655](https://github.com/hashicorp/terraform-provider-google-beta/pull/9655))
* **New Data Source:** `google_storage_control_project_intelligence_config` ([#9655](https://github.com/hashicorp/terraform-provider-google-beta/pull/9655))
* **New Resource:** `google_apigee_dns_zone` ([#9622](https://github.com/hashicorp/terraform-provider-google-beta/pull/9622))
* **New Resource:** `google_dataproc_metastore_database_iam_*` resources ([#9615](https://github.com/hashicorp/terraform-provider-google-beta/pull/9615))
* **New Resource:** `google_dataproc_metastore_table_iam_*` ([#9647](https://github.com/hashicorp/terraform-provider-google-beta/pull/9647))
* **New Resource:** `google_discovery_engine_sitemap` ([#9608](https://github.com/hashicorp/terraform-provider-google-beta/pull/9608))
* **New Resource:** `google_eventarc_enrollment` ([#9623](https://github.com/hashicorp/terraform-provider-google-beta/pull/9623))
* **New Resource:** `google_firebase_app_hosting_build` ([#9646](https://github.com/hashicorp/terraform-provider-google-beta/pull/9646))
* **New Resource:** `google_memorystore_instance_desired_user_created_endpoints` ([#9652](https://github.com/hashicorp/terraform-provider-google-beta/pull/9652))
* **New Resource:** `google_storage_control_folder_intelligence_config` ([#9644](https://github.com/hashicorp/terraform-provider-google-beta/pull/9644))
* **New Resource:** `google_storage_control_organization_intelligence_config` ([#9617](https://github.com/hashicorp/terraform-provider-google-beta/pull/9617))

IMPROVEMENTS:
* accesscontextmanager: added `roles` field to ingress and egress policies of `google_access_context_manager_service_perimeter*` resources ([#9661](https://github.com/hashicorp/terraform-provider-google-beta/pull/9661))
* cloudfunctions2: added `binary_authorization_policy` field to `google_cloudfunctions2_function` resource ([#9649](https://github.com/hashicorp/terraform-provider-google-beta/pull/9649))
* cloudrunv2: added `gpu_zonal_redundancy_disabled` field to `google_cloud_run_v2_service` resource ([#9639](https://github.com/hashicorp/terraform-provider-google-beta/pull/9639))
* compute: added `md5_authentication_keys` field to `google_compute_router` resource ([#9673](https://github.com/hashicorp/terraform-provider-google-beta/pull/9673))
* compute: added `EXTERNAL_IPV6_SUBNETWORK_CREATION` as a supported value for the `mode` field in `google_compute_public_delegated_prefix` resource ([#9630](https://github.com/hashicorp/terraform-provider-google-beta/pull/9630))
* compute: added `external_ipv6_prefix`, `stack_type`, and `ipv6_access_type` fields to `google_compute_subnetwork` data source ([#9660](https://github.com/hashicorp/terraform-provider-google-beta/pull/9660))
* compute: added `path_matchers.route_rules.custom_error_response_policy` field to `google_compute_url_map` resource ([#9656](https://github.com/hashicorp/terraform-provider-google-beta/pull/9656))
* compute: added `source_machine_image_encryption_key` field to `google_compute_instance_from_machine_image` resource ([#9632](https://github.com/hashicorp/terraform-provider-google-beta/pull/9632))
* compute: added `tls_settings` field to `google_compute_backend_service` resource ([#9654](https://github.com/hashicorp/terraform-provider-google-beta/pull/9654))
* compute: added several `boot_disk`, `attached_disk`, and `instance_encryption_key` fields to `google_compute_instance` and `google_compute_instance_template` resources ([#9669](https://github.com/hashicorp/terraform-provider-google-beta/pull/9669))
* compute: added `image_encryption_key.raw_key` and `image_encryption_key.rsa_encrypted_key` fields to `google_compute_image` resource ([#9669](https://github.com/hashicorp/terraform-provider-google-beta/pull/9669))
* compute: added `snapshot_encryption_key.rsa_encrypted_key` field to `google_compute_snapshot` resource ([#9669](https://github.com/hashicorp/terraform-provider-google-beta/pull/9669))
* container: added `disable_l4_lb_firewall_reconciliation` field to `google_container_cluster` resource ([#9648](https://github.com/hashicorp/terraform-provider-google-beta/pull/9648))
* datafusion: added `tags` field to `google_data_fusion_instance` resource to allow setting tags for instances at creation time ([#9609](https://github.com/hashicorp/terraform-provider-google-beta/pull/9609))
* datastream: added `blmt_config` field to `bigquery_destination_config` resource to enable support for BigLake Managed Tables streams ([#9677](https://github.com/hashicorp/terraform-provider-google-beta/pull/9677))
* datastream: added `secret_manager_stored_password` field to `google_datastream_connection_profile` resource ([#9633](https://github.com/hashicorp/terraform-provider-google-beta/pull/9633))
* identityplatform: added `disabled_user_signup` and `disabled_user_deletion` to `google_identity_platform_tenant` resource ([#9613](https://github.com/hashicorp/terraform-provider-google-beta/pull/9613))
* memorystore: added `psc_attachment_details` field to `google_memorystore_instance` resource, to enable use of the fine-grained resource `google_memorystore_instance_desired_user_created_connections` ([#9652](https://github.com/hashicorp/terraform-provider-google-beta/pull/9652))
* memorystore: added the `cross_cluster_replication_config` field to the `google_redis_cluster` resource ([#9670](https://github.com/hashicorp/terraform-provider-google-beta/pull/9670))
* metastore: added `deletion_protection` field to `google_dataproc_metastore_federation` resource ([#9674](https://github.com/hashicorp/terraform-provider-google-beta/pull/9674))
* networksecurity: added `antivirus_overrides` field to `google_network_security_security_profile` resource ([#9643](https://github.com/hashicorp/terraform-provider-google-beta/pull/9643))
* networksecurity: added `connected_deployment_groups` and `associations` fields to `google_network_security_mirroring_endpoint_group` resource ([#9606](https://github.com/hashicorp/terraform-provider-google-beta/pull/9606))
* networksecurity: added `locations` field to `google_network_security_mirroring_deployment_group` resource ([#9607](https://github.com/hashicorp/terraform-provider-google-beta/pull/9607))
* networksecurity: added `locations` field to `google_network_security_mirroring_endpoint_group_association` resource ([#9603](https://github.com/hashicorp/terraform-provider-google-beta/pull/9603))
* parametermanager: added `kms_key_version` field to `google_parameter_manager_parameter_version` resource and datasource ([#9642](https://github.com/hashicorp/terraform-provider-google-beta/pull/9642))
* parametermanager: added `kms_key` field to `google_parameter_manager_parameter` resource and `google_parameter_manager_parameters` datasource ([#9642](https://github.com/hashicorp/terraform-provider-google-beta/pull/9642))
* provider: added `external_credentials` block in `provider` ([#9658](https://github.com/hashicorp/terraform-provider-google-beta/pull/9658))
* redis: added `automated_backup_config` field to `google_redis_cluster` resource ([#9682](https://github.com/hashicorp/terraform-provider-google-beta/pull/9682))
* storage: added `content_base64` field in `google_storage_bucket_object_content` datasource ([#9638](https://github.com/hashicorp/terraform-provider-google-beta/pull/9638))

BUG FIXES:
* alloydb: added a mutex to `google_alloydb_cluster` to prevent conflicts among multiple cluster operations ([#9604](https://github.com/hashicorp/terraform-provider-google-beta/pull/9604))
* artifactregistry: fixed type assertion panic in `google_artifact_registry_repository` resource ([#9672](https://github.com/hashicorp/terraform-provider-google-beta/pull/9672))
* bigtable: fixed `automated_backup_policy` field for `google_bigtable_table` resource ([#9627](https://github.com/hashicorp/terraform-provider-google-beta/pull/9627))
* cloudrunv2: fixed the diffs for unchanged `template.template.containers.env` in `google_cloud_run_v2_job` resource ([#9681](https://github.com/hashicorp/terraform-provider-google-beta/pull/9681))
* compute: fixed a regression in `google_compute_subnetwork` where setting `log_config` would not enable flow logs without `enable_flow_logs` also being set to true. To enable or disable flow logs, please use `log_config`. `enable_flow_logs` is now deprecated and will be removed in the next major release. ([#9679](https://github.com/hashicorp/terraform-provider-google-beta/pull/9679))
* compute: fixed unable to update the `preview` field for `google_compute_region_security_policy_rule` resource ([#9614](https://github.com/hashicorp/terraform-provider-google-beta/pull/9614))
* compute: fixed unable to update the `preview` field for `google_compute_security_policy_rule` resource ([#9614](https://github.com/hashicorp/terraform-provider-google-beta/pull/9614))
* orgpolicy: fix permadiff in `google_org_policy_policy` when multiple rules are present ([#9611](https://github.com/hashicorp/terraform-provider-google-beta/pull/9611))
* resourcemanager: increased page size for list services api to help any teams hitting `ListEnabledRequestsPerMinutePerProject` quota issues ([#9637](https://github.com/hashicorp/terraform-provider-google-beta/pull/9637))
* spanner: fixed issue with applying changes in provider `default_labels` on `google_spanner_instance` resource ([#9629](https://github.com/hashicorp/terraform-provider-google-beta/pull/9629))
* storage: fixed `google_storage_anywhere_cache` to cancel long-running operations after create and update requests timeout ([#9625](https://github.com/hashicorp/terraform-provider-google-beta/pull/9625))
* workbench: fixed metadata permadiff in `google_workbench_instance` resource ([#9641](https://github.com/hashicorp/terraform-provider-google-beta/pull/9641))

## 6.27.0 (Mar 25, 2025)

FEATURES:
* **New Data Source:** `google_compute_images` ([#9556](https://github.com/hashicorp/terraform-provider-google-beta/pull/9556))
* **New Data Source:** `google_organization_iam_custom_role` ([#9577](https://github.com/hashicorp/terraform-provider-google-beta/pull/9577))
* **New Resource:** `google_lustre_instance` ([#9601](https://github.com/hashicorp/terraform-provider-google-beta/pull/9601))
* **New Resource:** `google_os_config_v2_policy_orchestrator` ([#9579](https://github.com/hashicorp/terraform-provider-google-beta/pull/9579))
* **New Resource:** `google_storage_control_project_intelligence_config` ([#9570](https://github.com/hashicorp/terraform-provider-google-beta/pull/9570))

IMPROVEMENTS:
* bigquery: added `secondary_location` and `replication_status` fields to support managed disaster recovery feature in `google_bigquery_reservation` ([#9575](https://github.com/hashicorp/terraform-provider-google-beta/pull/9575))
* clouddeploy: added `dns_endpoint` field to to `google_clouddeploy_target` resource ([#9553](https://github.com/hashicorp/terraform-provider-google-beta/pull/9553))
* compute: added `group_placement_policy.0.gpu_topology` field to  `google_compute_resource_policy` resource ([#9555](https://github.com/hashicorp/terraform-provider-google-beta/pull/9555))
* compute: added `shielded_instance_initial_state` structure to `google_compute_image` resource ([#9583](https://github.com/hashicorp/terraform-provider-google-beta/pull/9583))
* compute: added `LINK_TYPE_ETHERNET_400G_LR4` enum value to `link_type` field in `google_compute_interconnect` resource ([#9571](https://github.com/hashicorp/terraform-provider-google-beta/pull/9571))
* compute: added `architecture` and `guest_os_features` to `google_compute_instance` ([#9558](https://github.com/hashicorp/terraform-provider-google-beta/pull/9558))
* compute: added `instance_lifecycle_policy.on_failed_health_check` field in resources `google_compute_instance_group_manager` and `google_compute_region_instance_group_manager` ([#9598](https://github.com/hashicorp/terraform-provider-google-beta/pull/9598))
* compute: added `workload_policy.type`, `workload_policy.max_topology_distance` and `workload_policy.accelerator_topology` fields to `google_compute_resource_policy` resource ([#9599](https://github.com/hashicorp/terraform-provider-google-beta/pull/9599))
* container: added `ip_endpoints_config` field to `google_container_cluster` resource ([#9597](https://github.com/hashicorp/terraform-provider-google-beta/pull/9597))
* container: added `node_config.windows_node_config` field to `google_container_node_pool` resource. ([#9559](https://github.com/hashicorp/terraform-provider-google-beta/pull/9559))
* container: added `pod_autoscaling` field to `google_container_cluster` resource ([#9574](https://github.com/hashicorp/terraform-provider-google-beta/pull/9574))
* memorystore: added the `maintenance_policy` field to the `google_memorystore_instance` resource ([#9595](https://github.com/hashicorp/terraform-provider-google-beta/pull/9595))
* memorystore: enabled update support for `node_type` field in `google_memorystore_instance` resource ([#9568](https://github.com/hashicorp/terraform-provider-google-beta/pull/9568))
* networkmanagement: added `destination.forwarding_rule`, `destination.gke_master_cluster`, `destination.fqdn`, `destination.cloud_sql_instance`, `destination.redis_instance`, `destination.redis_cluster`,  fields to `google_network_management_connectivity_test` resource ([#9591](https://github.com/hashicorp/terraform-provider-google-beta/pull/9591))
* networkmanagement: added `round_trip`, `bypass_firewall_checks` fields to `google_network_management_connectivity_test` resource ([#9591](https://github.com/hashicorp/terraform-provider-google-beta/pull/9591))
* networkmanagement: added `source.gke_master_cluster`,  `source.cloud_sql_instance`, `source.cloud_function`, `source.app_engine_version`, `source.cloud_run_revision` fields to `google_network_management_connectivity_test` resource ([#9591](https://github.com/hashicorp/terraform-provider-google-beta/pull/9591))
* networksecurity: added `connected_deployment_group` and `associations` fields to `google_network_security_intercept_endpoint_group` resource ([#9586](https://github.com/hashicorp/terraform-provider-google-beta/pull/9586))
* networksecurity: added `locations` field to `google_network_security_intercept_deployment_group` resource ([#9578](https://github.com/hashicorp/terraform-provider-google-beta/pull/9578))
* networksecurity: added `locations` field to `google_network_security_intercept_endpoint_group_association` resource ([#9600](https://github.com/hashicorp/terraform-provider-google-beta/pull/9600))
* redis: added update support for `google_redis_cluster` `node_type` ([#9554](https://github.com/hashicorp/terraform-provider-google-beta/pull/9554))
* storage: added metadata_options in `google_storage_transfer_job` ([#9567](https://github.com/hashicorp/terraform-provider-google-beta/pull/9567))

BUG FIXES:
* bigqueryanalyticshub: fixed a bug in `google_bigquery_analytics_hub_listing_subscription` where a subscription using a different project than the dataset would not work ([#9596](https://github.com/hashicorp/terraform-provider-google-beta/pull/9596))
* cloudrun: fixed the perma-diffs for unchanged `template.spec.containers.env` in `google_cloud_run_service` resource ([#9572](https://github.com/hashicorp/terraform-provider-google-beta/pull/9572))
* cloudrunv2: fixed the perma-diffs for unchanged `template.containers.env` in `google_cloud_run_v2_service` resource ([#9572](https://github.com/hashicorp/terraform-provider-google-beta/pull/9572))
* compute: fixed the issue that user can't use regional disk in `google_compute_instance_template` ([#9569](https://github.com/hashicorp/terraform-provider-google-beta/pull/9569))
* dataflow: fixed a permadiff on `template_gcs_path` in `google_dataflow_job` resource ([#9564](https://github.com/hashicorp/terraform-provider-google-beta/pull/9564))
* storage: lowered the minimum required items for `custom_placement_config.data_locations` from 2 to 1, and removed the Terraform-enforced maximum item limit for the field in `google_storage_bucket` ([#9562](https://github.com/hashicorp/terraform-provider-google-beta/pull/9562))

## 6.26.0 (Mar 18, 2025)

FEATURES:
* **New Data Source:** `google_project_iam_custom_role` ([#9551](https://github.com/hashicorp/terraform-provider-google-beta/pull/9551))
* **New Data Source:** `google_project_iam_custom_roles` ([#9519](https://github.com/hashicorp/terraform-provider-google-beta/pull/9519))
* **New Resource:** `google_eventarc_pipeline` ([#9508](https://github.com/hashicorp/terraform-provider-google-beta/pull/9508))
* **New Resource:** `google_firebase_app_hosting_backend` ([#9531](https://github.com/hashicorp/terraform-provider-google-beta/pull/9531))
* **New Resource:** `google_managed_kafka_connect_cluster` ([#9552](https://github.com/hashicorp/terraform-provider-google-beta/pull/9552))
* **New Resource:** `google_managed_kafka_connector` ([#9552](https://github.com/hashicorp/terraform-provider-google-beta/pull/9552))

IMPROVEMENTS:
* alloydb: added `psc_config` field to ``google_alloydb_cluster` resource ([#9548](https://github.com/hashicorp/terraform-provider-google-beta/pull/9548))
* bigquery: added `table_metadata_view` query param to `google_bigquery_table` ([#9530](https://github.com/hashicorp/terraform-provider-google-beta/pull/9530))
* bigquery: added support for continuous query to `google_bigquery_job` ([#9520](https://github.com/hashicorp/terraform-provider-google-beta/pull/9520))
* clouddeploy: added `dns_endpoint` field to to `google_clouddeploy_target` resource ([#9553](https://github.com/hashicorp/terraform-provider-google-beta/pull/9553))
* compute: added `UNRESTRICTED` option to the `tls_early_data` field in the `google_compute_target_https_proxy` resource ([#9527](https://github.com/hashicorp/terraform-provider-google-beta/pull/9527))
* compute: added `enable_flow_logs` and `state` fields to `google_compute_subnetwork` resource ([#9541](https://github.com/hashicorp/terraform-provider-google-beta/pull/9541))
* container: added additional value `KCP_HPA` for `logging_config.enable_components` field in `google_container_cluster` resource ([#9529](https://github.com/hashicorp/terraform-provider-google-beta/pull/9529))
* dataform: added `deletion_policy` field to `google_dataform_repository` resource. Default value is `DELETE`. Setting `deletion_policy` to `FORCE` will delete any child resources of this repository as well. ([#9549](https://github.com/hashicorp/terraform-provider-google-beta/pull/9549))
* memorystore: added update support for `engine_version` field in `google_memorystore_instance` resource ([#9534](https://github.com/hashicorp/terraform-provider-google-beta/pull/9534))
* metastore: added `create_time` and `update_time` fields to `google_dataproc_metastore_federation` resource ([#9528](https://github.com/hashicorp/terraform-provider-google-beta/pull/9528))
* metastore: added `create_time` and `update_time` fields to `google_dataproc_metastore_service` resource ([#9523](https://github.com/hashicorp/terraform-provider-google-beta/pull/9523))
* networksecurity: added `not_operations` field to `google_network_security_authz_policy` resource ([#9511](https://github.com/hashicorp/terraform-provider-google-beta/pull/9511))
* networkservices: added `ip_version` and `envoy_headers` fields to `google_network_services_gateway` resource ([#9514](https://github.com/hashicorp/terraform-provider-google-beta/pull/9514))
* sql: increased `settings.insights_config.query_string_length` and `settings.insights_config.query_string_length` limits for Enterprise Plus edition `sql_database_instance` resource. ([#9539](https://github.com/hashicorp/terraform-provider-google-beta/pull/9539))
* storageinsights: added `parquet_options` field to `google_storage_insights_report_config` resource ([#9522](https://github.com/hashicorp/terraform-provider-google-beta/pull/9522))
* workflows: added `execution_history_level` field to `google_workflows_workflow` resource ([#9509](https://github.com/hashicorp/terraform-provider-google-beta/pull/9509))

BUG FIXES:
* accesscontextmanager: fixed panic on empty `access_policies` in `google_access_context_manager_access_policy` ([#9536](https://github.com/hashicorp/terraform-provider-google-beta/pull/9536))
* compute: adjusted mapped image names that were preventing usage of `fedora-coreos` in `google_compute_image` resource ([#9513](https://github.com/hashicorp/terraform-provider-google-beta/pull/9513))
* container: re-added `DNS_SCOPE_UNSPECIFIED` value to the `dns_config.cluster_dns_scope` field in `google_container_cluster` resource and suppressed diffs between `DNS_SCOPE_UNSPECIFIED` in config and empty/null in state ([#9547](https://github.com/hashicorp/terraform-provider-google-beta/pull/9547))
* discoveryengine: changed field `dataStoreIds` to mutable in `google_discovery_engine_search_engine` ([#9506](https://github.com/hashicorp/terraform-provider-google-beta/pull/9506))
* networksecurity: `min_tls_version` and `tls_feature_profile` fields updated to use the server assigned default and prevent a permadiff in `google_network_security_tls_inspection_policy` resource. ([#9514](https://github.com/hashicorp/terraform-provider-google-beta/pull/9514))
* oslogin: added a wait after creating `google_os_login_ssh_public_key` to allow for propagation ([#9546](https://github.com/hashicorp/terraform-provider-google-beta/pull/9546))
* spanner: fixed issue with disabling autoscaling in `google_spanner_instance` ([#9542](https://github.com/hashicorp/terraform-provider-google-beta/pull/9542))

## 6.25.0 (Mar 11, 2025)

NOTES:
* eventarc: `google_eventarc_channel` now uses MMv1 engine instead of DCL. ([#9488](https://github.com/hashicorp/terraform-provider-google-beta/pull/9488))
* workbench: increased create timeout for `google_workbench_instance` to 40mins. ([#9468](https://github.com/hashicorp/terraform-provider-google-beta/pull/9468))

FEATURES:
* **New Data Source:** `google_compute_region_ssl_policy` ([#9439](https://github.com/hashicorp/terraform-provider-google-beta/pull/9439))
* **New Resource:** `google_eventarc_google_api_source` ([#9492](https://github.com/hashicorp/terraform-provider-google-beta/pull/9492))
* **New Resource:** `google_iam_oauth_client_credential` ([#9491](https://github.com/hashicorp/terraform-provider-google-beta/pull/9491))
* **New Resource:** `google_iam_oauth_client` ([#9456](https://github.com/hashicorp/terraform-provider-google-beta/pull/9456))
* **New Resource:** `google_network_security_backend_authentication_config` ([#9481](https://github.com/hashicorp/terraform-provider-google-beta/pull/9481))

IMPROVEMENTS:
* alloydb: added `psc_instance_config.psc_interface_configs` field to `google_alloydb_instance` resource ([#9469](https://github.com/hashicorp/terraform-provider-google-beta/pull/9469))
* compute: added `create_snapshot_before_destroy` to `google_compute_disk` and `google_compute_region_disk` to enable creating a snapshot before disk deletion ([#9442](https://github.com/hashicorp/terraform-provider-google-beta/pull/9442))
* compute: added `custom_metrics` field to `google_compute_backend_service` and `google_compute_region_backend_service` ([#9473](https://github.com/hashicorp/terraform-provider-google-beta/pull/9473))
* compute: added `ip_collection` and `ipv6_gce_endpoint` fields to `google_compute_subnetwork` resource ([#9490](https://github.com/hashicorp/terraform-provider-google-beta/pull/9490))
* compute: added `log_config.optional_mode` and `log_config.optional_fields` fields to `google_compute_region_backend_service` resource ([#9484](https://github.com/hashicorp/terraform-provider-google-beta/pull/9484))
* compute: added `rsa_encrypted_key` to `google_compute_region_disk` ([#9442](https://github.com/hashicorp/terraform-provider-google-beta/pull/9442))
* compute: added `scheduling.termination_time` field to `google_compute_instance`, `google_compute_instance_from_machine_image`, `google_compute_instance_from_template`, `google_compute_instance_template`, and `google_compute_region_instance_template` resources ([#9479](https://github.com/hashicorp/terraform-provider-google-beta/pull/9479))
* compute: added update support for `firewall_policy`  in `google_compute_firewall_policy_association` resource. It is recommended to only perform this operation in combination with a protective lifecycle tag such as "create_before_destroy" or "prevent_destroy" on your previous `firewall_policy` resource in order to prevent situations where a target attachment has no associated policy. ([#9495](https://github.com/hashicorp/terraform-provider-google-beta/pull/9495))
* compute: made `purpose` field updatable in `google_compute_subnetwork`. ([#9489](https://github.com/hashicorp/terraform-provider-google-beta/pull/9489))
* container: added "JOBSET" as a supported value for `enable_components` in `google_container_cluster` resource ([#9453](https://github.com/hashicorp/terraform-provider-google-beta/pull/9453))
* datastream: added support for creating connection profiles for Salesforce in `google_datastream_connection_profile` ([#9482](https://github.com/hashicorp/terraform-provider-google-beta/pull/9482))
* firebasedataconnect: added `deletion_policy` field to `google_firebase_data_connect_service` resource ([#9496](https://github.com/hashicorp/terraform-provider-google-beta/pull/9496))
* networksecurity: added `description` field to `google_network_security_intercept_deployment`, `google_network_security_intercept_deployment_group`, `google_network_security_intercept_endpoint_group` resources ([#9474](https://github.com/hashicorp/terraform-provider-google-beta/pull/9474))
* networksecurity: added `description` field to `google_network_security_mirroring_deployment`, `google_network_security_mirroring_deployment_group`, `google_network_security_mirroring_endpoint_group` resources ([#9476](https://github.com/hashicorp/terraform-provider-google-beta/pull/9476))
* tpuv2: added `spot` field to `google_tpu_v2_vm` resource ([#9478](https://github.com/hashicorp/terraform-provider-google-beta/pull/9478))
* workstations: added `tags` field to `google_workstations_workstation_cluster` resource ([#9441](https://github.com/hashicorp/terraform-provider-google-beta/pull/9441))

BUG FIXES:
* backupdr: added missing `SUNDAY` option to `days_of_week` field in `google_backup_dr_backup_plan` resource ([#9446](https://github.com/hashicorp/terraform-provider-google-beta/pull/9446))
* compute: fixed `network_interface.internal_ipv6_prefix_length` not being set or read in Terraform state in `google_compute_instance` resource ([#9444](https://github.com/hashicorp/terraform-provider-google-beta/pull/9444))
* compute: fixed bug in `google_compute_router_nat` where `max_ports_per_vm` couldn't be unset once set. ([#9483](https://github.com/hashicorp/terraform-provider-google-beta/pull/9483))
* container: fixed perma-diff in `google_container_cluster` when `cluster_dns_scope` is unspecified ([#9443](https://github.com/hashicorp/terraform-provider-google-beta/pull/9443))
* networksecurity: added wait time on `google_network_security_gateway_security_policy_rule` resource when creating and deleting to prevent race conditions ([#9448](https://github.com/hashicorp/terraform-provider-google-beta/pull/9448))

## 6.24.0 (Mar 3, 2025)

NOTES:
* gemini: removed unsupported value `GEMINI_CLOUD_ASSIST` for field `product` in `google_gemini_logging_setting_binding` resource ([#9438](https://github.com/hashicorp/terraform-provider-google-beta/pull/9438))
* gemini: removed unsupported value `GEMINI_CODE_ASSIST` for field `product` in `google_gemini_data_sharing_with_google_setting_binding` resource (Beta) ([#9437](https://github.com/hashicorp/terraform-provider-google-beta/pull/9437))
* iam: added member value to the error message when member validation fails for google_project_iam_* ([#9406](https://github.com/hashicorp/terraform-provider-google-beta/pull/9406))

DEPRECATIONS:
* datacatalog: deprecated `google_data_catalog_entry` and `google_data_catalog_tag` resources. For steps to transition your Data Catalog users, workloads, and content to Dataplex Catalog, see https://cloud.google.com/dataplex/docs/transition-to-dataplex-catalog. ([#9393](https://github.com/hashicorp/terraform-provider-google-beta/pull/9393))
* notebooks: deprecated non-functional `google_notebooks_location` resource ([#9373](https://github.com/hashicorp/terraform-provider-google-beta/pull/9373))

FEATURES:
* **New Data Source:** `google_memorystore_instance` ([#9400](https://github.com/hashicorp/terraform-provider-google-beta/pull/9400))
* **New Resource:** `google_apihub_host_project_registration` ([#9419](https://github.com/hashicorp/terraform-provider-google-beta/pull/9419))
* **New Resource:** `google_compute_instant_snapshot` ([#9412](https://github.com/hashicorp/terraform-provider-google-beta/pull/9412))
* **New Resource:** `google_eventarc_message_bus` ([#9423](https://github.com/hashicorp/terraform-provider-google-beta/pull/9423))
* **New Resource:** `google_gemini_data_sharing_with_google_setting_binding` (GA) ([#9437](https://github.com/hashicorp/terraform-provider-google-beta/pull/9437))
* **New Resource:** `google_gemini_gcp_enablement_setting_binding` (GA) ([#9407](https://github.com/hashicorp/terraform-provider-google-beta/pull/9407))
* **New Resource:** `google_gemini_gemini_gcp_enablement_setting_binding` ([#9392](https://github.com/hashicorp/terraform-provider-google-beta/pull/9392))
* **New Resource:** `google_storage_anywhere_cache` ([#9389](https://github.com/hashicorp/terraform-provider-google-beta/pull/9389))

IMPROVEMENTS:
* compute: added `creation_timestamp`, `next_hop_peering`, ` warnings.code`, `warnings.message`, `warnings.data.key`, `warnings.data.value`, `next_hop_hub`, `route_type`, `as_paths.path_segment_type`, `as_paths.as_lists` and `route_status`  fields to `google_compute_route` resource ([#9386](https://github.com/hashicorp/terraform-provider-google-beta/pull/9386))
* compute: added `max_stream_duration` field to `google_compute_url_map` resource ([#9387](https://github.com/hashicorp/terraform-provider-google-beta/pull/9387))
* compute: added fields `architecture`, `source_instant_snapshot`, `source_storage_object`, `resource_manager_tags`  to `google_compute_disk`. ([#9412](https://github.com/hashicorp/terraform-provider-google-beta/pull/9412))
* container: added enum  value `UPGRADE_INFO_EVENT` for GKE notification filter in `google_container_cluster` resource ([#9421](https://github.com/hashicorp/terraform-provider-google-beta/pull/9421))
* iam: added `AZURE_AD_GROUPS_ID` field to `google_iam_workforce_pool_provider.extra_attributes_oauth2_client.attributes_type` resource ([#9433](https://github.com/hashicorp/terraform-provider-google-beta/pull/9433))
* networkconnectivity: added `policy_mode` field to `google_network_connectivity_hub` resource ([#9409](https://github.com/hashicorp/terraform-provider-google-beta/pull/9409))
* networkservices: added `location` field to `google_network_services_grpc_route` resource ([#9429](https://github.com/hashicorp/terraform-provider-google-beta/pull/9429))
* storagetransfer: added `logging_config` field to `google_storage_transfer_job` resource ([#9378](https://github.com/hashicorp/terraform-provider-google-beta/pull/9378))

BUG FIXES:
* bigquery: updated the `max_staleness` field in `google_bigquery_table` to be a computed field ([#9411](https://github.com/hashicorp/terraform-provider-google-beta/pull/9411))
* chronicle: fixed an error during resource creation with certain `run_frequency` configurations in `google_chronicle_rule_deployment` ([#9422](https://github.com/hashicorp/terraform-provider-google-beta/pull/9422))
* discoveryengine: fixed bug preventing creation of `google_discovery_engine_target_site` resources ([#9436](https://github.com/hashicorp/terraform-provider-google-beta/pull/9436))
* eventarc: fixed an issue where `google_eventarc_trigger` creation failed due to the region could not be parsed from the trigger's name ([#9383](https://github.com/hashicorp/terraform-provider-google-beta/pull/9383))
* gemini: fixed permadiff on `product` field in `google_gemini_data_sharing_with_google_setting_binding` resource (Beta) ([#9437](https://github.com/hashicorp/terraform-provider-google-beta/pull/9437))
* publicca: encoded `b64_mac_key` in base64url, instead of base64 in `google_public_ca_external_account_key` ([#9424](https://github.com/hashicorp/terraform-provider-google-beta/pull/9424))
* storage: fixed a 412 error returned on some `google_storage_bucket_iam_policy` deletions ([#9434](https://github.com/hashicorp/terraform-provider-google-beta/pull/9434))

## 6.23.0 (Feb 26, 2025)

NOTES:
* The `google_sql_user` resource now supports `password_wo` [write-only arguments](https://developer.hashicorp.com/terraform/language/v1.11.x/resources/ephemeral#write-only-arguments)
* The `google_bigquery_data_transfer_config` resource now supports `secret_access_key_wo` [write-only arguments](https://developer.hashicorp.com/terraform/language/v1.11.x/resources/ephemeral#write-only-arguments)
* The `google_secret_version` resource now supports `secret_data_wo` [write-only arguments](https://developer.hashicorp.com/terraform/language/v1.11.x/resources/ephemeral#write-only-arguments)

IMPROVEMENTS:
* sql: added `password_wo` and `password_wo_version` fields to `google_sql_user` resource ([#21616](https://github.com/hashicorp/terraform-provider-google/pull/21616))
* bigquerydatatransfer: added `secret_access_key_wo` and `secret_access_key_wo_version` fields to `google_bigquery_data_transfer_config` resource ([#21617](https://github.com/hashicorp/terraform-provider-google/pull/21617))
* secretmanager: added `secret_data_wo` and `secret_data_wo_version` fields to `google_secret_version` resource ([#21618](https://github.com/hashicorp/terraform-provider-google/pull/21618))

## 6.22.0 (Feb 24, 2025)

NOTES:
* provider: The Terraform Provider for Google Cloud's regular release date will move from Monday to Tuesday in early March. The 2025/03/10 release will be made on 2025/03/11.

DEPRECATIONS:
* datacatalog: deprecated `google_data_catalog_tag_template`. Use `google_dataplex_aspect_type` instead. For steps to transition your Data Catalog users, workloads, and content to Dataplex Catalog, see https://cloud.google.com/dataplex/docs/transition-to-dataplex-catalog. ([#9347](https://github.com/hashicorp/terraform-provider-google-beta/pull/9347))
* datacatalog: deperecated `google_data_catalog_entry_group`. Use `google_dataplex_entry_group` instead. For steps to transition your Data Catalog users, workloads, and content to Dataplex Catalog, see https://cloud.google.com/dataplex/docs/transition-to-dataplex-catalog. ([#9349](https://github.com/hashicorp/terraform-provider-google-beta/pull/9349))

FEATURES:
* **New Data Source:** `google_alloydb_cluster` ([#9361](https://github.com/hashicorp/terraform-provider-google-beta/pull/9361))
* **New Data Source:** `google_project_ancestry` ([#9326](https://github.com/hashicorp/terraform-provider-google-beta/pull/9326))
* **New Resource:** `google_gemini_data_sharing_with_google_setting_binding` ([#9356](https://github.com/hashicorp/terraform-provider-google-beta/pull/9356))
* **New Resource:** `google_spanner_instance_partition` ([#9354](https://github.com/hashicorp/terraform-provider-google-beta/pull/9354))

IMPROVEMENTS:
* compute: added `import_subnet_routes_with_public_ip` and `export_subnet_routes_with_public_ip` fields to `google_compute_network_peering_routes_config` resource ([#9320](https://github.com/hashicorp/terraform-provider-google-beta/pull/9320))
* developerconnect: added `bitbucket_cloud_config` and `bitbucket_data_center_config` fields to `google_developer_connect_connection` resource (ga) ([#9338](https://github.com/hashicorp/terraform-provider-google-beta/pull/9338))
* iam: added `extra_attributes_oauth2_client` field to `google_iam_workforce_pool_provider` resource ([#9336](https://github.com/hashicorp/terraform-provider-google-beta/pull/9336))
* redis: added `kms_key` field to `google_redis_cluster` resource ([#9334](https://github.com/hashicorp/terraform-provider-google-beta/pull/9334))
* tpuv2: added `network_config` field to `google_tpu_v2_queued_resource` resource ([#9332](https://github.com/hashicorp/terraform-provider-google-beta/pull/9332))

BUG FIXES:
* apigee: fixed error when deleting `google_apigee_organization` ([#9352](https://github.com/hashicorp/terraform-provider-google-beta/pull/9352))
* bigtable: fixed a bug where sometimes updating an instance's cluster list could result in an error if there was an existing cluster with autoscaling enabled ([#9368](https://github.com/hashicorp/terraform-provider-google-beta/pull/9368))
* chronicle: fixed bug setting `enabled` on creation in `google_chronicle_rule_deployment` ([#9343](https://github.com/hashicorp/terraform-provider-google-beta/pull/9343))
## 6.21.0 (Feb 18, 2025)

NOTES:
* provider: The Terraform Provider for Google Cloud's regular release date will move from Monday to Tuesday in early March. The 2025/03/10 release will be made on 2025/03/11.

FEATURES:
* **New Data Source:** `google_alloydb_instance` ([#9307](https://github.com/hashicorp/terraform-provider-google-beta/pull/9307))
* **New Resource:** `google_firebase_data_connect_service` ([#9304](https://github.com/hashicorp/terraform-provider-google-beta/pull/9304))
* **New Resource:** `google_gemini_data_sharing_with_google_setting` ([#9250](https://github.com/hashicorp/terraform-provider-google-beta/pull/9250))
* **New Resource:** `google_gemini_gemini_gcp_enablement_setting` (beta) ([#9253](https://github.com/hashicorp/terraform-provider-google-beta/pull/9253))
* **New Resource:** `google_gemini_logging_setting_binding` ([#9292](https://github.com/hashicorp/terraform-provider-google-beta/pull/9292))
* **New Resource:** `google_gemini_release_channel_setting_binding` ([#9287](https://github.com/hashicorp/terraform-provider-google-beta/pull/9287))
* **New Resource:** `google_netapp_volume_quota_rule` ([#9248](https://github.com/hashicorp/terraform-provider-google-beta/pull/9248))

IMPROVEMENTS:
* accesscontextmanager: added `etag` to access context manager directional policy resources `google_access_context_manager_service_perimeter_dry_run_egress_policy`, `google_access_context_manager_service_perimeter_dry_run_ingress_policy`, `google_access_context_manager_service_perimeter_egress_policy` and `google_access_context_manager_service_perimeter_ingress_policy` to prevent overriding changes ([#9302](https://github.com/hashicorp/terraform-provider-google-beta/pull/9302))
* accesscontextmanager: added `title` field to policy blocks under `google_access_context_manager_service_perimeter` and variants ([#9259](https://github.com/hashicorp/terraform-provider-google-beta/pull/9259))
* artifactregistry: set pageSize to 1000 to speedup `google_artifact_registry_docker_image` data source queries ([#9297](https://github.com/hashicorp/terraform-provider-google-beta/pull/9297))
* compute: added `graceful_shutdown` field to `google_compute_instance`, `google_compute_instance_template` and `google_compute_region_instance_template` resource ([#9278](https://github.com/hashicorp/terraform-provider-google-beta/pull/9278))
* compute: added `labels` field to `google_compute_ha_vpn_gateway` resource ([#9309](https://github.com/hashicorp/terraform-provider-google-beta/pull/9309))
* compute: added validation for disk names in `google_compute_disk` ([#9280](https://github.com/hashicorp/terraform-provider-google-beta/pull/9280))
* container: added new fields `container_log_max_size`, `container_log_max_files`, `image_gc_low_threshold_percent`, `image_gc_high_threshold_percent`, `image_minimum_gc_age`, `image_maximum_gc_age`, and `allowed_unsafe_sysctls` to `node_kubelet_config` block in `google_container_cluster` resource. ([#9274](https://github.com/hashicorp/terraform-provider-google-beta/pull/9274))
* monitoring: added `condition_sql` field to `google_monitoring_alert_policy` resource ([#9242](https://github.com/hashicorp/terraform-provider-google-beta/pull/9242))
* networkservices: added `location` field to `google_network_services_mesh` resource ([#9282](https://github.com/hashicorp/terraform-provider-google-beta/pull/9282))
* workstations: added update support to `persistent_directories.gce_pd.size_gb` and `persistent_directories.gce_pd.disk_type`  in `google_workstations_workstation_config` resource ([#9305](https://github.com/hashicorp/terraform-provider-google-beta/pull/9305))
* securitycenter: added `type`, `expiry_time` field to `google_scc_mute_config` resource ([#9273](https://github.com/hashicorp/terraform-provider-google-beta/pull/9273))

BUG FIXES:
* chronicle: fixed creation issues when optional fields were missing for `google_chronicle_rule_deployment` resource ([#9312](https://github.com/hashicorp/terraform-provider-google-beta/pull/9312))
* dns: fixed a bug where `google_dns_managed_zone` is unable to update with `service_directory_config` specified ([#9239](https://github.com/hashicorp/terraform-provider-google-beta/pull/9239))
* databasemigrationservice: fixed error details type on `google_database_migration_service_migration_job` ([#9244](https://github.com/hashicorp/terraform-provider-google-beta/pull/9244))
* networkservices: fixed a bug with `google_network_services_authz_extension.wire_format` sending an invalid default value by removing the Terraform default and letting the API set the default. ([#9245](https://github.com/hashicorp/terraform-provider-google-beta/pull/9245))

## 6.20.0 (Feb 10, 2025)

NOTES:
* provider: The Terraform Provider for Google Cloud's regular release date will move from Monday to Tuesday in early March. The 2025/03/10 release will be made on 2025/03/11.
* compute: `google_compute_firewall_policy` now uses MMv1 engine instead of DCL. ([#9228](https://github.com/hashicorp/terraform-provider-google-beta/pull/9228))

FEATURES:
* **New Data Source:** `google_beyondcorp_application_iam_policy` ([#9205](https://github.com/hashicorp/terraform-provider-google-beta/pull/9205))
* **New Data Source:** `google_parameter_manager_parameter_version_render` ([#9190](https://github.com/hashicorp/terraform-provider-google-beta/pull/9190))
* **New Data Source:** `google_parameter_manager_regional_parameter_version_render` ([#9232](https://github.com/hashicorp/terraform-provider-google-beta/pull/9232))
* **New Resource:** `google_beyondcorp_application` ([#9205](https://github.com/hashicorp/terraform-provider-google-beta/pull/9205))
* **New Resource:** `google_beyondcorp_application_iam_binding`  ([#9205](https://github.com/hashicorp/terraform-provider-google-beta/pull/9205))
* **New Resource:** `google_beyondcorp_application_iam_member`  ([#9205](https://github.com/hashicorp/terraform-provider-google-beta/pull/9205))
* **New Resource:** `google_beyondcorp_application_iam_policy`  ([#9205](https://github.com/hashicorp/terraform-provider-google-beta/pull/9205))
* **New Resource:** `google_bigquery_analytics_hub_listing_subscription` ([#9195](https://github.com/hashicorp/terraform-provider-google-beta/pull/9195))
* **New Resource:** `google_colab_notebook_execution` ([#9186](https://github.com/hashicorp/terraform-provider-google-beta/pull/9186))
* **New Resource:** `google_colab_schedule` ([#9226](https://github.com/hashicorp/terraform-provider-google-beta/pull/9226))
* **New Resource:** `google_compute_network_firewall_policy_packet_mirroring_rule` ([#9202](https://github.com/hashicorp/terraform-provider-google-beta/pull/9202))
* **New Resource:** `google_gemini_logging_setting` ([#9198](https://github.com/hashicorp/terraform-provider-google-beta/pull/9198))
* **New Resource:** `google_gemini_release_channel_setting` ([#9207](https://github.com/hashicorp/terraform-provider-google-beta/pull/9207))

IMPROVEMENTS:
* accesscontextmanager: added `resource` to `sources` in `egress_from` under resources `google_access_context_manager_service_perimeter`, `google_access_context_manager_service_perimeters`, `google_access_context_manager_service_perimeter_egress_policy`, `google_access_context_manager_service_perimeter_dry_run_egress_policy` ([#9196](https://github.com/hashicorp/terraform-provider-google-beta/pull/9196))
* cloudrunv2: added `base_image_uri` and `build_info` to `google_cloud_run_v2_service` ([#9229](https://github.com/hashicorp/terraform-provider-google-beta/pull/9229))
* colab: added `auto_upgrade` field to `google_colab_runtime` ([#9216](https://github.com/hashicorp/terraform-provider-google-beta/pull/9216))
* colab: added `software_config.post_startup_script_config` field to `google_colab_runtime_template` ([#9206](https://github.com/hashicorp/terraform-provider-google-beta/pull/9206))
* colab: added `desired_state` field to `google_colab_runtime`, making it startable/stoppable ([#9209](https://github.com/hashicorp/terraform-provider-google-beta/pull/9209))
* compute: added `ip_collection` field to `google_compute_forwarding_rule ` resource ([#9194](https://github.com/hashicorp/terraform-provider-google-beta/pull/9194))
* compute: added `mode` and `allocatable_prefix_length` fields to `google_compute_public_delegated_prefix` resource ([#9218](https://github.com/hashicorp/terraform-provider-google-beta/pull/9218))
* compute: allow parallelization of `google_compute_per_instance_config` and `google_compute_region_per_instance_config` deletions by not locking on the parent resource, but including instance name. ([#9181](https://github.com/hashicorp/terraform-provider-google-beta/pull/9181))
* container: added `auto_monitoring_config` field and subfields to the `google_container_cluster` resource ([#9224](https://github.com/hashicorp/terraform-provider-google-beta/pull/9224))
* filestore: added `initial_replication` field for peer instance configuration and `effective_replication` output for replication configuration output to `google_filestore_instance` ([#9200](https://github.com/hashicorp/terraform-provider-google-beta/pull/9200))
* memorystore: added `CLUSTER_DISABLED`  to `mode` field  in  `google_memorystore_instance` ([#9178](https://github.com/hashicorp/terraform-provider-google-beta/pull/9178))
* networkservices: added `compression_mode` and `allowed_methods` fields to `google_network_services_edge_cache_service` resource ([#9201](https://github.com/hashicorp/terraform-provider-google-beta/pull/9201))
* privateca: added `user_defined_access_urls` and subfields to `google_privateca_certificate_authority` resource to add support for custom CDP AIA URLs ([#9221](https://github.com/hashicorp/terraform-provider-google-beta/pull/9221))
* workbench: added `enable_third_party_identity` field to `google_workbench_instance` resource ([#9236](https://github.com/hashicorp/terraform-provider-google-beta/pull/9236))

## 6.19.0 (Feb 3, 2025)
NOTES:
* tpuv2: made service use the v2alpha1 Cloud TPU API version, which is used for Public Preview features ([#9131](https://github.com/hashicorp/terraform-provider-google-beta/pull/9131))

DEPRECATIONS:
* beyondcorp: deprecated `location` on `google_beyondcorp_security_gateway`. The only valid value is `global`, which is now also the default value. The field will be removed in a future major release. ([#9121](https://github.com/hashicorp/terraform-provider-google-beta/pull/9121))

FEATURES:
* **New Data Source:** `google_parameter_manager_parameter_version` ([#9154](https://github.com/hashicorp/terraform-provider-google-beta/pull/9154))
* **New Data Source:** `google_parameter_manager_parameters` ([#9148](https://github.com/hashicorp/terraform-provider-google-beta/pull/9148))
* **New Data Source:** `google_parameter_manager_regional_parameter_version` ([#9165](https://github.com/hashicorp/terraform-provider-google-beta/pull/9165))
* **New Resource:** `google_beyondcorp_security_gateway_iam_binding` ([#9169](https://github.com/hashicorp/terraform-provider-google-beta/pull/9169))
* **New Resource:** `google_beyondcorp_security_gateway_iam_member` ([#9169](https://github.com/hashicorp/terraform-provider-google-beta/pull/9169))
* **New Resource:** `google_beyondcorp_security_gateway_iam_policy` ([#9169](https://github.com/hashicorp/terraform-provider-google-beta/pull/9169))

IMPROVEMENTS:
* accesscontextmanager: added `etag` to `google_access_context_manager_service_perimeter_dry_run_resource` to prevent overriding list of resources ([#9120](https://github.com/hashicorp/terraform-provider-google-beta/pull/9120))
* bigquery: added `schema_foreign_type_info` field and related schema handling to `google_bigquery_table` resource (beta) ([#9122](https://github.com/hashicorp/terraform-provider-google-beta/pull/9122))
* compute: allowed parallelization of `google_compute_(region_)per_instance_config` by not locking on the parent resource, but including instance name. ([#9116](https://github.com/hashicorp/terraform-provider-google-beta/pull/9116))
* compute: added `network_profile` field to `google_compute_network` resource. ([#9135](https://github.com/hashicorp/terraform-provider-google-beta/pull/9135))
* compute: added `zero_advertised_route_priority` field to `google_compute_router_peer` ([#9133](https://github.com/hashicorp/terraform-provider-google-beta/pull/9133))
* container: added `max_run_duration` to `node_config` in `google_container_cluster` and `google_container_node_pool` ([#9163](https://github.com/hashicorp/terraform-provider-google-beta/pull/9163))
* dataproc: added `encryption_config` to `google_dataproc_workflow_template` ([#9168](https://github.com/hashicorp/terraform-provider-google-beta/pull/9168))
* gkehub2: added support for `fleet_default_member_config.config_management.config_sync.metrics_gcp_service_account_email` field to `google_gke_hub_feature` resource ([#9147](https://github.com/hashicorp/terraform-provider-google-beta/pull/9147))
* iam: added `prefix` and `regex` fields to `google_service_accounts` data source ([#9129](https://github.com/hashicorp/terraform-provider-google-beta/pull/9129))
* pubsub: added `ingestion_data_source_settings.aws_msk` and `ingestion_data_source_settings.confluent_cloud` fields to `google_pubsub_topic` resource ([#9114](https://github.com/hashicorp/terraform-provider-google-beta/pull/9114))
* spanner: added `encryption_config` field to  `google_spanner_backup_schedule` ([#9161](https://github.com/hashicorp/terraform-provider-google-beta/pull/9161))
* workflows: added `tags` and `workflow_tags` fields to `google_workflows_workflow` resource ([#9152](https://github.com/hashicorp/terraform-provider-google-beta/pull/9152))

BUG FIXES:
* alloydb: marked `google_alloydb_user.password` as sensitive ([#9124](https://github.com/hashicorp/terraform-provider-google-beta/pull/9124))
* beyondcorp: corrected `location` to always be global in `google_beyondcorp_security_gateway` ([#9121](https://github.com/hashicorp/terraform-provider-google-beta/pull/9121))
* cloudquotas: removed validation for `parent` in `google_cloud_quotas_quota_adjuster_settings` ([#9153](https://github.com/hashicorp/terraform-provider-google-beta/pull/9153))
* compute: made `google_compute_router_peer.advertised_route_priority` use server-side default if unset. To set the value to `0` you must also set `zero_advertised_route_priority = true`. ([#9133](https://github.com/hashicorp/terraform-provider-google-beta/pull/9133))
* container: fixed a diff caused by server-side set values for `node_config.resource_labels` ([#9171](https://github.com/hashicorp/terraform-provider-google-beta/pull/9171))
* container: marked `cluster_autoscaling.resource_limits.maximum` as required, as requests would fail if it was not set ([#9151](https://github.com/hashicorp/terraform-provider-google-beta/pull/9151))
* firestore: fixed error preventing deletion of wildcard fields in `google_firestore_field` ([#9140](https://github.com/hashicorp/terraform-provider-google-beta/pull/9140))
* netapp: fixed an issue where a diff on `zone` would be found if it was unspecified in `google_netapp_storage_pool` ([#9157](https://github.com/hashicorp/terraform-provider-google-beta/pull/9157))
* networksecurity: fixed sporadic-diff in `google_network_security_security_profile` ([#9162](https://github.com/hashicorp/terraform-provider-google-beta/pull/9162))
* spanner: fixed bug with `google_spanner_instance.force_destroy` not setting `billing_project` value correctly ([#9132](https://github.com/hashicorp/terraform-provider-google-beta/pull/9132))
* storage: fixed an issue where plans with a dependency on the `content` field in the `google_storage_bucket_object_content` data source could erroneously fail ([#9166](https://github.com/hashicorp/terraform-provider-google-beta/pull/9166))

## 6.18.1 (January 29, 2025)

BUG FIXES:
* container: fixed a diff caused by server-side set values for `node_config.resource_labels` ([#9171](https://github.com/hashicorp/terraform-provider-google-beta/pull/9171))

## 6.18.0 (January 27, 2025)

FEATURES:
* **New Data Source:** `google_compute_instance_template_iam_policy` ([#9085](https://github.com/hashicorp/terraform-provider-google-beta/pull/9085))
* **New Data Source:** `google_kms_key_handles` ([#9105](https://github.com/hashicorp/terraform-provider-google-beta/pull/9105))
* **New Data Source:** `google_organizations` ([#9093](https://github.com/hashicorp/terraform-provider-google-beta/pull/9093))
* **New Data Source:** `google_parameter_manager_parameter` ([#9084](https://github.com/hashicorp/terraform-provider-google-beta/pull/9084))
* **New Data Source:** `google_parameter_manager_regional_parameters` ([#9089](https://github.com/hashicorp/terraform-provider-google-beta/pull/9089))
* **New Resource:** `google_apihub_api_hub_instance` ([#9080](https://github.com/hashicorp/terraform-provider-google-beta/pull/9080))
* **New Resource:** `google_chronicle_retrohunt` ([#9090](https://github.com/hashicorp/terraform-provider-google-beta/pull/9090))
* **New Resource:** `google_colab_runtime` ([#9076](https://github.com/hashicorp/terraform-provider-google-beta/pull/9076))
* **New Resource:** `google_colab_runtime_template_iam_binding` ([#9091](https://github.com/hashicorp/terraform-provider-google-beta/pull/9091))
* **New Resource:** `google_colab_runtime_template_iam_member` ([#9091](https://github.com/hashicorp/terraform-provider-google-beta/pull/9091))
* **New Resource:** `google_colab_runtime_template_iam_policy` ([#9091](https://github.com/hashicorp/terraform-provider-google-beta/pull/9091))
* **New Resource:** `google_compute_instance_template_iam_binding` ([#9085](https://github.com/hashicorp/terraform-provider-google-beta/pull/9085))
* **New Resource:** `google_compute_instance_template_iam_member` ([#9085](https://github.com/hashicorp/terraform-provider-google-beta/pull/9085))
* **New Resource:** `google_compute_instance_template_iam_policy` ([#9085](https://github.com/hashicorp/terraform-provider-google-beta/pull/9085))
* **New Resource:** `google_parameter_manager_parameter_version` ([#9111](https://github.com/hashicorp/terraform-provider-google-beta/pull/9111))
* **New Resource:** `google_redis_cluster_user_created_connections` ([#9099](https://github.com/hashicorp/terraform-provider-google-beta/pull/9099))

IMPROVEMENTS:
* alloydb: added support for `skip_await_major_version_upgrade` field in `google_alloydb_cluster` resource, allowing for `major_version` to be updated ([#9066](https://github.com/hashicorp/terraform-provider-google-beta/pull/9066))
* apigee: added `properties` field to `google_apigee_environment` resource ([#9072](https://github.com/hashicorp/terraform-provider-google-beta/pull/9072))
* bug: added support for setting `custom_learned_route_priority` to 0 in 'google_compute_router_peer' by adding the `zero_custom_learned_route_priority` field ([#9083](https://github.com/hashicorp/terraform-provider-google-beta/pull/9083))
* cloudrunv2: added `build_config` to `google_cloud_run_v2_service` ([#9100](https://github.com/hashicorp/terraform-provider-google-beta/pull/9100))
* compute: added `dest_network_scope`, `src_network_scope` and `src_networks` fields to `google_compute_firewall_policy_rule` resource (beta) ([#9082](https://github.com/hashicorp/terraform-provider-google-beta/pull/9082))
* compute: added `dest_network_scope`, `src_network_scope` and `src_networks` fields to `google_compute_firewall_policy_with_rules` resource (beta) ([#9082](https://github.com/hashicorp/terraform-provider-google-beta/pull/9082))
* compute: added `dest_network_scope`, `src_network_scope` and `src_networks` fields to `google_compute_network_firewall_policy_rule` resource (beta) ([#9082](https://github.com/hashicorp/terraform-provider-google-beta/pull/9082))
* compute: added `dest_network_scope`, `src_network_scope` and `src_networks` fields to `google_compute_network_firewall_policy_with_rules` resource (beta) ([#9082](https://github.com/hashicorp/terraform-provider-google-beta/pull/9082))
* compute: added `dest_network_scope`, `src_network_scope` and `src_networks` fields to `google_compute_region_network_firewall_policy_rule` resource (beta) ([#9082](https://github.com/hashicorp/terraform-provider-google-beta/pull/9082))
* compute: added `dest_network_scope`, `src_network_scope` and `src_networks` fields to `google_compute_region_network_firewall_policy_with_rules` resource (beta) ([#9082](https://github.com/hashicorp/terraform-provider-google-beta/pull/9082))
* compute: added `pdp_scope` field to `google_compute_public_advertised_prefix` resource ([#9096](https://github.com/hashicorp/terraform-provider-google-beta/pull/9096))
* compute: adding `labels` field to `google_compute_interconnect_attachment` ([#9095](https://github.com/hashicorp/terraform-provider-google-beta/pull/9095))
* compute: fixed a issue where `custom_learned_route_priority` was accidentally set to 0 during updates in 'google_compute_router_peer' ([#9083](https://github.com/hashicorp/terraform-provider-google-beta/pull/9083))
* filestore: added support for `tags` field to `google_filestore_instance` resource ([#9086](https://github.com/hashicorp/terraform-provider-google-beta/pull/9086))
* networksecurity: added `custom_mirroring_profile` and `custom_intercept_profile` fields to `google_network_security_security_profile` and `google_network_security_security_profile_group`  resources ([#9110](https://github.com/hashicorp/terraform-provider-google-beta/pull/9110))
* pubsub: added `enforce_in_transit` fields to `google_pubsub_topic` resource ([#9069](https://github.com/hashicorp/terraform-provider-google-beta/pull/9069))
* pubsub: added `ingestion_data_source_settings.azure_event_hubs` field to `google_pubsub_topic` resource ([#9065](https://github.com/hashicorp/terraform-provider-google-beta/pull/9065))
* redis: added `psc_service_attachments` field to `google_redis_cluster` resource, to enable use of the fine-grained resource `google_redis_cluster_user_created_connections` ([#9099](https://github.com/hashicorp/terraform-provider-google-beta/pull/9099))

BUG FIXES:
* apigee: fixed `properties` field update on `google_apigee_environment` resource ([#9107](https://github.com/hashicorp/terraform-provider-google-beta/pull/9107))
* artifactregistry: fixed perma-diff in `google_artifact_registry_repository` ([#9109](https://github.com/hashicorp/terraform-provider-google-beta/pull/9109))
* compute: fixed failure when creating `google_compute_global_forwarding_rule` with labels targeting PSC endpoint ([#9106](https://github.com/hashicorp/terraform-provider-google-beta/pull/9106))
* container: fixed `additive_vpc_scope_dns_domain` being ignored in Autopilot cluster definition ([#9075](https://github.com/hashicorp/terraform-provider-google-beta/pull/9075))
* container: fixed propagation of `node_pool_defaults.node_config_defaults.insecure_kubelet_readonly_port_enabled` in node config. ([#9074](https://github.com/hashicorp/terraform-provider-google-beta/pull/9074))
* iam: fixed missing result by adding pagination for data source `google_service_accounts`. ([#9094](https://github.com/hashicorp/terraform-provider-google-beta/pull/9094))
* metastore: increased timeout on google_dataproc_metastore_service operations to 75m from 60m. This will expose server-returned reasons for operation failure instead of masking them with a Terraform timeout. ([#9102](https://github.com/hashicorp/terraform-provider-google-beta/pull/9102))
* resourcemanager: added a slightly longer wait (two 10s checks bumped to 15s) for issues with billing associations in `google_project`. Default network deletion should succeed more often. ([#9103](https://github.com/hashicorp/terraform-provider-google-beta/pull/9103))

## 6.17.0 (January 21, 2025)

FEATURES:
* **New Data Source:** `google_parameter_manager_regional_parameter` (beta) ([#9030](https://github.com/hashicorp/terraform-provider-google-beta/pull/9030))
* **New Resource:** `google_apigee_environment_addons_config` ([#9021](https://github.com/hashicorp/terraform-provider-google-beta/pull/9021))
* **New Resource:** `google_chronicle_reference_list` (beta) ([#9047](https://github.com/hashicorp/terraform-provider-google-beta/pull/9047))
* **New Resource:** `google_chronicle_rule_deployment` ([#9043](https://github.com/hashicorp/terraform-provider-google-beta/pull/9043))
* **New Resource:** `google_chronicle_rule` ([#9032](https://github.com/hashicorp/terraform-provider-google-beta/pull/9032))
* **New Resource:** `google_colab_runtime_template` ([#9050](https://github.com/hashicorp/terraform-provider-google-beta/pull/9050))
* **New Resource:** `google_edgenetwork_interconnect_attachment` ([#9024](https://github.com/hashicorp/terraform-provider-google-beta/pull/9024))
* **New Resource:** `google_parameter_manager_parameter` ([#9041](https://github.com/hashicorp/terraform-provider-google-beta/pull/9041))
* **New Resource:** `google_parameter_manager_regional_parameter_version` ([#9062](https://github.com/hashicorp/terraform-provider-google-beta/pull/9062))
* **New Resource:** `google_parameter_manager_regional_parameter` ([#9026](https://github.com/hashicorp/terraform-provider-google-beta/pull/9026))

IMPROVEMENTS:
* accesscontextmanager: added `etag` to `google_access_context_manager_service_perimeter_resource` to prevent overriding list of resources ([#9058](https://github.com/hashicorp/terraform-provider-google-beta/pull/9058))
* compute: added `BPS_100G` enum value to `bandwidth` field of `google_compute_interconnect_attachment`. ([#9040](https://github.com/hashicorp/terraform-provider-google-beta/pull/9040))
* compute: added support for `IPV6_ONLY` stack_type to `google_compute_subnetwork`, `google_compute_instance`, `google_compute_instance_template` and `google_compute_region_instance_template`. ([#9020](https://github.com/hashicorp/terraform-provider-google-beta/pull/9020))
* compute: promoted `bgp_best_path_selection_mode `,`bgp_bps_always_compare_med` and `bgp_bps_inter_region_cost ` fields in `google_compute_network` from Beta to Ga ([#9029](https://github.com/hashicorp/terraform-provider-google-beta/pull/9029))
* compute: promoted `next_hop_origin `,`next_hop_med ` and `next_hop_inter_region_cost ` output fields in `google_compute_route` form Beta to GA ([#9029](https://github.com/hashicorp/terraform-provider-google-beta/pull/9029))
* discoveryengine: added `advanced_site_search_config` field to `google_discovery_engine_data_store` resource ([#9060](https://github.com/hashicorp/terraform-provider-google-beta/pull/9060))
* gemini: added `force_destroy` field to resource `google_code_repository_index`, enabling deletion of the resource even when it has dependent RepositoryGroups ([#9036](https://github.com/hashicorp/terraform-provider-google-beta/pull/9036))
* networkservices: added in-place update support for `ports` field on `google_network_services_gateway` resource ([#9056](https://github.com/hashicorp/terraform-provider-google-beta/pull/9056))
* sql: `sql_source_representation_instance` now uses `string` representation of `databaseVersion` ([#9027](https://github.com/hashicorp/terraform-provider-google-beta/pull/9027))
* sql: added `replication_cluster` field to `google_sql_database_instance` resource ([#9044](https://github.com/hashicorp/terraform-provider-google-beta/pull/9044))
* sql: added support of switchover for MySQL and PostgreSQL in `google_sql_database_instance` resource ([#9044](https://github.com/hashicorp/terraform-provider-google-beta/pull/9044))
* workbench: changed `container_image` field of `google_workbench_instance` resource to modifiable. ([#9046](https://github.com/hashicorp/terraform-provider-google-beta/pull/9046))

BUG FIXES:
* apigee: fixed error 404 for `organization` update requests. ([#9022](https://github.com/hashicorp/terraform-provider-google-beta/pull/9022))
* artifactregistry: fixed `artifact_registry_repository` not accepting durations with 'm', 'h' or 'd' ([#9054](https://github.com/hashicorp/terraform-provider-google-beta/pull/9054))
* networkservices: fixed bug where `google_network_services_gateway` could not be updated in place ([#9056](https://github.com/hashicorp/terraform-provider-google-beta/pull/9056))
* storagetransfer: fixed a permadiff with `transfer_spec.aws_s3_data_source.aws_access_key` in `google_storage_transfer_job` ([#9019](https://github.com/hashicorp/terraform-provider-google-beta/pull/9019))

## 6.16.0 (January 13, 2025)

FEATURES:
* **New Data Source:** `google_kms_autokey_config` ([#8986](https://github.com/hashicorp/terraform-provider-google-beta/pull/8986))
* **New Resource:** `google_beyondcorp_security_gateway` ([#9017](https://github.com/hashicorp/terraform-provider-google-beta/pull/9017))
* **New Resource:** `google_chronicle_data_access_label` ([#8999](https://github.com/hashicorp/terraform-provider-google-beta/pull/8999))
* **New Resource:** `google_chronicle_data_access_scope` ([#9000](https://github.com/hashicorp/terraform-provider-google-beta/pull/9000))
* **New Resource:** `google_cloud_quotas_quota_adjuster_settings` ([#9005](https://github.com/hashicorp/terraform-provider-google-beta/pull/9005))

IMPROVEMENTS:
* chronicle: updated `watchlist_id` field to be optional in `google_chronicle_watchlist` resource ([#8988](https://github.com/hashicorp/terraform-provider-google-beta/pull/8988))
* developerconnect: added `crypto_key_config`, `github_enterprise_config`, `gitlab_config` , and `gitlab_enterprise_config` fields to `google_developer_connect_connection` resource ([#8998](https://github.com/hashicorp/terraform-provider-google-beta/pull/8998))
* dns: added `health_check` and `external_endpoints` fields to `google_dns_record_set` resource ([#9016](https://github.com/hashicorp/terraform-provider-google-beta/pull/9016))
* sql: added `server_ca_pool` field to `google_sql_database_instance` resource ([#9008](https://github.com/hashicorp/terraform-provider-google-beta/pull/9008))
* vmwareengine: allowed import of non-STANDARD private clouds in `google_vmwareengine_private_cloud` ([#9006](https://github.com/hashicorp/terraform-provider-google-beta/pull/9006))

BUG FIXES:
* dataproc: fixed boolean fields in `shielded_instance_config` in the `google_dataproc_cluster` resource ([#9003](https://github.com/hashicorp/terraform-provider-google-beta/pull/9003))
* gkeonprem: fixed permadiff on `vcenter` field in `google_gkeonprem_vmware_cluster` resource ([#9011](https://github.com/hashicorp/terraform-provider-google-beta/pull/9011))
* kms: fixed permadiff on `google_kms_autokey_config` by introducing a 5 second sleep post-create / post-update ([#8992](https://github.com/hashicorp/terraform-provider-google-beta/pull/8992))
* networkservices: fixed `google_network_services_gateway` resource so that it correctly waits for the router to be deleted on `terraform destroy` ([#8993](https://github.com/hashicorp/terraform-provider-google-beta/pull/8993))
* provider: fixed issue where `GOOGLE_CLOUD_QUOTA_PROJECT` env var would override explicit `billing_project` ([#9012](https://github.com/hashicorp/terraform-provider-google-beta/pull/9012))

## 6.15.0 (January 6, 2025)

NOTES:
* compute: `google_compute_firewall_policy_association` now uses MMv1 engine instead of DCL. ([#8948](https://github.com/hashicorp/terraform-provider-google-beta/pull/8948))

DEPRECATIONS:
* compute: deprecated `numeric_id` (string) field in `google_compute_network` resource. Use the new `network_id` (integer)  field instead ([#8915](https://github.com/hashicorp/terraform-provider-google-beta/pull/8915))

FEATURES:
* **New Data Source:** `google_gke_hub_feature` ([#8930](https://github.com/hashicorp/terraform-provider-google-beta/pull/8930))
* **New Data Source:** `google_kms_autokey_config` ([#8986](https://github.com/hashicorp/terraform-provider-google-beta/pull/8986))
* **New Data Source:** `google_kms_key_handle` ([#8933](https://github.com/hashicorp/terraform-provider-google-beta/pull/8933))
* **New Resource:** `google_gkeonprem_vmware_admin_cluster` ([#8932](https://github.com/hashicorp/terraform-provider-google-beta/pull/8932))
* **New Resource:** `google_chronicle_watchlist` ([#8983](https://github.com/hashicorp/terraform-provider-google-beta/pull/8983))
* **New Resource:** `google_network_security_intercept_endpoint_group_association` ([#8958](https://github.com/hashicorp/terraform-provider-google-beta/pull/8958))
* **New Resource:** `google_network_security_intercept_endpoint_group` ([#8912](https://github.com/hashicorp/terraform-provider-google-beta/pull/8912))
* **New Resource:** `google_storage_folder` ([#8961](https://github.com/hashicorp/terraform-provider-google-beta/pull/8961))

IMPROVEMENTS:
* artifactregistry: added `vulnerability_scanning_config` field to `google_artifact_registry_repository` resource ([#8934](https://github.com/hashicorp/terraform-provider-google-beta/pull/8934))
* bigquery: added `condition` field to `google_bigquery_dataset_access` resource ([#8921](https://github.com/hashicorp/terraform-provider-google-beta/pull/8921))
* bigquery: added `condition` field to `google_bigquery_dataset` resource ([#8921](https://github.com/hashicorp/terraform-provider-google-beta/pull/8921))
* bigquery: added `external_catalog_table_options` field to `google_bigquery_table` resource ([#8942](https://github.com/hashicorp/terraform-provider-google-beta/pull/8942))
* composer: added `airflow_metadata_retention_config` field to `google_composer_environment` ([#8963](https://github.com/hashicorp/terraform-provider-google-beta/pull/8963))
* compute: added back the validation for `target_service` field on the `google_compute_service_attachment` resource to validade a `ForwardingRule` or `Gateway` URL ([#8924](https://github.com/hashicorp/terraform-provider-google-beta/pull/8924))
* compute: added `availability_domain` field to `google_compute_instance`, `google_compute_instance_template` and `google_compute_region_instance_template` resources ([#8914](https://github.com/hashicorp/terraform-provider-google-beta/pull/8914))
* compute: added `network_id` (integer) field to `google_compute_network` resource and data source ([#8915](https://github.com/hashicorp/terraform-provider-google-beta/pull/8915))
* compute: added `preset_topology` field to `google_network_connectivity_hub` resource ([#8929](https://github.com/hashicorp/terraform-provider-google-beta/pull/8929))
* compute: added `subnetwork_id` field to `google_compute_subnetwork` data source ([#8893](https://github.com/hashicorp/terraform-provider-google-beta/pull/8893))
* compute: made setting resource policies for `google_compute_instance` outside of terraform or using `google_compute_disk_resource_policy_attachment` no longer affect the `boot_disk.initialize_params.resource_policies` field ([#8959](https://github.com/hashicorp/terraform-provider-google-beta/pull/8959))
* container: changed `google_container_cluster` to apply maintenance policy updates after upgrades during cluster update ([#8922](https://github.com/hashicorp/terraform-provider-google-beta/pull/8922))
* container: made nodepool concurrent operations scale better for `google_container_cluster` and `google_container_node_pool` resources ([#8943](https://github.com/hashicorp/terraform-provider-google-beta/pull/8943))
* datastream: added `gtid` and `binary_log_position` fields to `google_datastream_stream` resource ([#8967](https://github.com/hashicorp/terraform-provider-google-beta/pull/8967))
* developerconnect: added support for setting up a `google_developer_connect_connection` resource without specifying the `authorizer_credentials` field ([#8953](https://github.com/hashicorp/terraform-provider-google-beta/pull/8953))
* filestore: added `tags` field to `google_filestore_backup` to allow setting tags for backups at creation time ([#8928](https://github.com/hashicorp/terraform-provider-google-beta/pull/8928))
* networkconnectivity: added `group` field to `google_network_connectivity_spoke` resource ([#8909](https://github.com/hashicorp/terraform-provider-google-beta/pull/8909))
* parallelstore: added `deployment_type` field to `google_parallelstore_instance` resource ([#8939](https://github.com/hashicorp/terraform-provider-google-beta/pull/8939))
* storagetransfer: added `replication_spec` field to `google_storage_transfer_job` resource ([#8976](https://github.com/hashicorp/terraform-provider-google-beta/pull/8976))
* workbench: made `gcs-data-bucket` metadata key modifiable in `google_workbench_instance` resource ([#8936](https://github.com/hashicorp/terraform-provider-google-beta/pull/8936))
* workstations: added `source_workstation` field to `google_workstations_workstation` resource ([#8938](https://github.com/hashicorp/terraform-provider-google-beta/pull/8938))

BUG FIXES:
* accesscontextmanager: fixed permadiff due to reordering on `google_access_context_manager_service_perimeter_dry_run_egress_policy` `egress_from.identities` ([#8980](https://github.com/hashicorp/terraform-provider-google-beta/pull/8980))
* accesscontextmanager: fixed permadiff due to reordering on `google_access_context_manager_service_perimeter_dry_run_ingress_policy` `ingress_from.identities` ([#8980](https://github.com/hashicorp/terraform-provider-google-beta/pull/8980))
* accesscontextmanager: fixed permadiff due to reordering on `google_access_context_manager_service_perimeter_egress_policy` `egress_from.identities` ([#8980](https://github.com/hashicorp/terraform-provider-google-beta/pull/8980))
* accesscontextmanager: fixed permadiff due to reordering on `google_access_context_manager_service_perimeter_ingress_policy` `ingress_from.identities` ([#8980](https://github.com/hashicorp/terraform-provider-google-beta/pull/8980))
* apigee: fixed 404 error when updating `google_apigee_environment` ([#8949](https://github.com/hashicorp/terraform-provider-google-beta/pull/8949))
* bigquery: fixed DROP COLUMN error with bigquery flexible column names in `google_bigquery_table` ([#8982](https://github.com/hashicorp/terraform-provider-google-beta/pull/8982))
* compute: allowed Service Attachment with Project Number to be used as `google_compute_forwarding_rule.target` ([#8978](https://github.com/hashicorp/terraform-provider-google-beta/pull/8978))
* compute: fixed an issue where `terraform plan -refresh=false` with `google_compute_ha_vpn_gateway.gateway_ip_version` would plan a resource replacement if a full refresh had not been run yet. Terraform now assumes that the value is the default value, `IPV4`, until a refresh is completed. ([#8904](https://github.com/hashicorp/terraform-provider-google-beta/pull/8904))
* compute: fixed panic when zonal resize request fails on `google_compute_resize_request` ([#8941](https://github.com/hashicorp/terraform-provider-google-beta/pull/8941))
* compute: fixed perma-destroy for `psc_data` in `google_compute_region_network_endpoint_group` resource ([#8972](https://github.com/hashicorp/terraform-provider-google-beta/pull/8972))
* compute: fixed `google_compute_instance_guest_attributes` to return an empty list when queried values don't exist instead of throwing an error ([#8957](https://github.com/hashicorp/terraform-provider-google-beta/pull/8957))
* integrationconnectors: allowed `AUTH_TYPE_UNSPECIFIED` option in `google_integration_connectors_connection` resource to support non-standard auth types ([#8971](https://github.com/hashicorp/terraform-provider-google-beta/pull/8971))
* logging: fixed bug in `google_logging_project_bucket_config` when providing `project` in the format of `<project-id-only>` ([#8923](https://github.com/hashicorp/terraform-provider-google-beta/pull/8923))
* networkconnectivity: made `include_export_ranges` and `exclude_export_ranges` fields mutable in `google_network_connectivity_spoke` to avoid recreation of resources ([#8946](https://github.com/hashicorp/terraform-provider-google-beta/pull/8946))
* sql: fixed permadiff when `settings.data_cache_config` is set to false for `google_sql_database_instance` resource ([#8889](https://github.com/hashicorp/terraform-provider-google-beta/pull/8889))
* storage: made `resource_google_storage_bucket_object` generate diff for `md5hash`, `generation`, `crc32c` if content changes ([#8908](https://github.com/hashicorp/terraform-provider-google-beta/pull/8908))
* vertexai: made `contents_delta_uri` an optional field in `google_vertex_ai_index` ([#8969](https://github.com/hashicorp/terraform-provider-google-beta/pull/8969))
* workbench: fixed an issue where a server-added `metadata` tag of `"resource-url"` would not be ignored on `google_workbench_instance` ([#8927](https://github.com/hashicorp/terraform-provider-google-beta/pull/8927))

## 6.14.1 (December 18, 2024)

BUG FIXES:
* compute: fixed an issue where `google_compute_firewall_policy_rule` was incorrectly removed from the Terraform state ([#8940](https://github.com/hashicorp/terraform-provider-google-beta/pull/8940))

## 6.14.0 (December 16, 2024)

FEATURES:
* **New Resource:** `google_network_security_intercept_deployment_group` ([#8859](https://github.com/hashicorp/terraform-provider-google-beta/pull/8859))
* **New Resource:** `google_network_security_intercept_deployment` ([#8876](https://github.com/hashicorp/terraform-provider-google-beta/pull/8876))
* **New Resource:** `google_network_security_authz_policy` ([#8847](https://github.com/hashicorp/terraform-provider-google-beta/pull/8847))
* **New Resource:** `google_network_services_authz_extension` ([#8847](https://github.com/hashicorp/terraform-provider-google-beta/pull/8847))

IMPROVEMENTS:
* compute: `google_compute_instance` is no longer recreated when changing `boot_disk.auto_delete` ([#8837](https://github.com/hashicorp/terraform-provider-google-beta/pull/8837))
* compute: added `CA_ENTERPRISE_ANNUAL` option for field `cloud_armor_tier` in `google_compute_project_cloud_armor_tier` resource ([#8848](https://github.com/hashicorp/terraform-provider-google-beta/pull/8848))
* compute: added `network_tier` field to `google_compute_global_forwarding_rule` resource ([#8838](https://github.com/hashicorp/terraform-provider-google-beta/pull/8838))
* compute: made `metadata_startup_script` able to be updated via graceful switch in `google_compute_instance` ([#8888](https://github.com/hashicorp/terraform-provider-google-beta/pull/8888))
* firebasehosting: added `headers` field in `google_firebase_hosting_version` resource ([#8887](https://github.com/hashicorp/terraform-provider-google-beta/pull/8887))
* identityplatform: marked `quota.0.sign_up_quota_config` subfields conditionally required in `google_identity_platform_config` to move errors from apply time up to plan time, and clarified the rule in documentation ([#8869](https://github.com/hashicorp/terraform-provider-google-beta/pull/8869))
* networkconnectivity: added support for updating `linked_vpn_tunnels.include_import_ranges`, `linked_interconnect_attachments.include_import_ranges`, `linked_router_appliance_instances. instances` and `linked_router_appliance_instances.include_import_ranges` in `google_network_connectivity_spoke` ([#8883](https://github.com/hashicorp/terraform-provider-google-beta/pull/8883))
* orgpolicy: added `parameters` fields to `google_org_policy_policy` resource ([#8881](https://github.com/hashicorp/terraform-provider-google-beta/pull/8881))
* storage: added `hdfs_data_source` field to `google_storage_transfer_job` resource ([#8839](https://github.com/hashicorp/terraform-provider-google-beta/pull/8839))
* tpuv2: added `network_configs` and `network_config.queue_count` fields to `google_tpu_v2_vm` resource ([#8865](https://github.com/hashicorp/terraform-provider-google-beta/pull/8865))

BUG FIXES:
* accesscontextmanager: fixed an update bug in `google_access_context_manager_perimeter` by removing the broken output-only `etag` field in `google_access_context_manager_perimeter` and `google_access_context_manager_perimeters` ([#8891](https://github.com/hashicorp/terraform-provider-google-beta/pull/8911))
* compute: fixed permadiff on the `recaptcha_options` field for `google_compute_security_policy` resource ([#8861](https://github.com/hashicorp/terraform-provider-google-beta/pull/8861))
* compute: fixed issue where updating labels on `resource_google_compute_resource_policy` would fail because of a patch error with `guest_flush` ([#8874](https://github.com/hashicorp/terraform-provider-google-beta/pull/8874))
* networkconnectivity: fixed `linked_router_appliance_instances.instances.virtual_machine` and `linked_router_appliance_instances.instances.ip_address` attributes in `google_network_connectivity_spoke` to be correctly marked as required. Otherwise the request to create the resource will fail. ([#8883](https://github.com/hashicorp/terraform-provider-google-beta/pull/8883))
* privateca: fixed an issue which causes error when updating labels for activated sub-CA ([#8872](https://github.com/hashicorp/terraform-provider-google-beta/pull/8872))
* sql: fixed permadiff when 'settings.data_cache_config' is set to false for 'google_sql_database_instance' resource ([#8889](https://github.com/hashicorp/terraform-provider-google-beta/pull/8889))

## 6.13.0 (December 9, 2024)

NOTES:
* New [ephemeral resources](https://developer.hashicorp.com/terraform/language/v1.10.x/resources/ephemeral) `google_service_account_access_token`, `google_service_account_id_token`, `google_service_account_jwt`, `google_service_account_key` now support [ephemeral values](https://developer.hashicorp.com/terraform/language/v1.10.x/values/variables#exclude-values-from-state).
DEPRECATIONS:
* gkehub: deprecated `configmanagement.config_sync.metrics_gcp_service_account_email` in `google_gke_hub_feature_membership` resource ([#8827](https://github.com/hashicorp/terraform-provider-google-beta/pull/8827))

FEATURES:
* **New Ephemeral Resource:** `google_service_account_access_token` ([#20542](https://github.com/hashicorp/terraform-provider-google/pull/20542))
* **New Ephemeral Resource:** `google_service_account_id_token` ([#20542](https://github.com/hashicorp/terraform-provider-google/pull/20542))
* **New Ephemeral Resource:** `google_service_account_jwt` ([#20542](https://github.com/hashicorp/terraform-provider-google/pull/20542))
* **New Ephemeral Resource:** `google_service_account_key` ([#20542](https://github.com/hashicorp/terraform-provider-google/pull/20542))
* **New Data Source:** `google_backup_dr_backup_vault` ([#8775](https://github.com/hashicorp/terraform-provider-google-beta/pull/8775))
* **New Data Source:** `google_backup_dr_backup` (beta) ([#8762](https://github.com/hashicorp/terraform-provider-google-beta/pull/8762))
* **New Resource:** `google_gemini_code_repository_index` ([#8781](https://github.com/hashicorp/terraform-provider-google-beta/pull/8781))
* **New Resource:** `google_gemini_repository_group_iam_binding` (beta only) ([#8824](https://github.com/hashicorp/terraform-provider-google-beta/pull/8824))
* **New Resource:** `google_gemini_repository_group_iam_member` (beta only) ([#8824](https://github.com/hashicorp/terraform-provider-google-beta/pull/8824))
* **New Resource:** `google_gemini_repository_group_iam_policy` (beta only) ([#8824](https://github.com/hashicorp/terraform-provider-google-beta/pull/8824))
* **New Resource:** `google_gemini_repository_group` (beta only) ([#8824](https://github.com/hashicorp/terraform-provider-google-beta/pull/8824))
* **New Resource:** `google_iam_projects_policy_binding` (beta) ([#8756](https://github.com/hashicorp/terraform-provider-google-beta/pull/8756))
* **New Resource:** `google_network_security_mirroring_deployment` ([#8791](https://github.com/hashicorp/terraform-provider-google-beta/pull/8791))
* **New Resource:** `google_network_security_mirroring_deployment_group` ([#8791](https://github.com/hashicorp/terraform-provider-google-beta/pull/8791))
* **New Resource:** `google_network_security_mirroring_endpoint_group_association` ([#8791](https://github.com/hashicorp/terraform-provider-google-beta/pull/8791))
* **New Resource:** `google_network_security_mirroring_endpoint_group` ([#8791](https://github.com/hashicorp/terraform-provider-google-beta/pull/8791))
* **New Resource:** `google_tpu_v2_queued_resource` (beta) ([#8760](https://github.com/hashicorp/terraform-provider-google-beta/pull/8760))

IMPROVEMENTS:
* accesscontextmanager: added `etag` to `google_access_context_manager_service_perimeter` and `google_access_context_manager_service_perimeters` ([#8767](https://github.com/hashicorp/terraform-provider-google-beta/pull/8767))
* alloydb: increased default timeout on `google_alloydb_cluster` to 120m from 30m ([#8820](https://github.com/hashicorp/terraform-provider-google-beta/pull/8820))
* bigtable: added `row_affinity` field to `google_bigtable_app_profile` resource ([#8753](https://github.com/hashicorp/terraform-provider-google-beta/pull/8753))
* cloudbuild: added `private_service_connect` field to `google_cloudbuild_worker_pool` resource ([#8827](https://github.com/hashicorp/terraform-provider-google-beta/pull/8827))
* clouddeploy: added `associated_entities` field to `google_clouddeploy_target` resource ([#8827](https://github.com/hashicorp/terraform-provider-google-beta/pull/8827))
* clouddeploy: added `serial_pipeline.strategy.canary.runtime_config.kubernetes.gateway_service_mesh.route_destinations` field to `google_clouddeploy_delivery_pipeline` resource ([#8827](https://github.com/hashicorp/terraform-provider-google-beta/pull/8827))
* composer: added multiple composer 3 related fields to `google_composer_environment` (GA) ([#8784](https://github.com/hashicorp/terraform-provider-google-beta/pull/8784))
* compute: `google_compute_instance`, `google_compute_instance_template`, `google_compute_region_instance_template` now supports `advanced_machine_features.enable_uefi_networking` field ([#8805](https://github.com/hashicorp/terraform-provider-google-beta/pull/8805))
* compute: added support for specifying storage pool with name or partial url ([#8794](https://github.com/hashicorp/terraform-provider-google-beta/pull/8794))
* compute: added `numeric_id` to the `google_compute_network` data source ([#8821](https://github.com/hashicorp/terraform-provider-google-beta/pull/8821))
* compute: added `threshold_configs` field to `google_compute_security_policy` resource ([#8818](https://github.com/hashicorp/terraform-provider-google-beta/pull/8818))
* compute: added server generated id as `forwarding_rule_id` to `google_compute_global_forwarding_rule` ([#8736](https://github.com/hashicorp/terraform-provider-google-beta/pull/8736))
* compute: added server generated id as `health_check_id` to `google_region_health_check` ([#8736](https://github.com/hashicorp/terraform-provider-google-beta/pull/8736))
* compute: added server generated id as `instance_group_manager_id` to `google_instance_group_manager` ([#8736](https://github.com/hashicorp/terraform-provider-google-beta/pull/8736))
* compute: added server generated id as `instance_group_manager_id` to `google_region_instance_group_manager` ([#8736](https://github.com/hashicorp/terraform-provider-google-beta/pull/8736))
* compute: added server generated id as `network_endpoint_id` to `google_region_network_endpoint` ([#8736](https://github.com/hashicorp/terraform-provider-google-beta/pull/8736))
* compute: added server generated id as `subnetwork_id` to `google_subnetwork` ([#8736](https://github.com/hashicorp/terraform-provider-google-beta/pull/8736))
* compute: added the `psc_data` field to the `google_compute_region_network_endpoint_group` resource ([#8766](https://github.com/hashicorp/terraform-provider-google-beta/pull/8766))
* container: added `enterprise_config` field to `google_container_cluster` resource ([#8808](https://github.com/hashicorp/terraform-provider-google-beta/pull/8808))
* container: added `node_pool_autoconfig.linux_node_config.cgroup_mode` field to `google_container_cluster` resource ([#8771](https://github.com/hashicorp/terraform-provider-google-beta/pull/8771))
* dataproc: added `autotuning_config` and `cohort` fields to `google_dataproc_batch` ([#8740](https://github.com/hashicorp/terraform-provider-google-beta/pull/8740))
* dataproc: added `cluster_config.preemptible_worker_config.instance_flexibility_policy.provisioning_model_mix` field to `google_dataproc_cluster` resource ([#8732](https://github.com/hashicorp/terraform-provider-google-beta/pull/8732))
* dataproc: added `confidential_instance_config` field to `google_dataproc_cluster` resource ([#8790](https://github.com/hashicorp/terraform-provider-google-beta/pull/8790))
* discoveryengine: added `HEALTHCARE_FHIR` to `industry_vertical` field in `google_discovery_engine_search_engine` ([#8778](https://github.com/hashicorp/terraform-provider-google-beta/pull/8778))
* gkehub: added `configmanagement.config_sync.stop_syncing` field to `google_gke_hub_feature_membership` resource ([#8827](https://github.com/hashicorp/terraform-provider-google-beta/pull/8827))
* monitoring: added `disable_metric_validation` field to `google_monitoring_alert_policy` resource ([#8817](https://github.com/hashicorp/terraform-provider-google-beta/pull/8817))
* oracledatabase: added `deletion_protection` field to `google_oracle_database_autonomous_database` ([#8787](https://github.com/hashicorp/terraform-provider-google-beta/pull/8787))
* oracledatabase: added `deletion_protection` field to `google_oracle_database_cloud_exadata_infrastructure` ([#8788](https://github.com/hashicorp/terraform-provider-google-beta/pull/8788))
* oracledatabase: added `deletion_protection` field to `google_oracle_database_cloud_vm_cluster ` ([#8730](https://github.com/hashicorp/terraform-provider-google-beta/pull/8730))
* parallelstore: added `deployment_type` to `google_parallelstore_instance` ([#8769](https://github.com/hashicorp/terraform-provider-google-beta/pull/8769))
* resourcemanager: made `google_service_account` `email` and `member` fields available during plan ([#8799](https://github.com/hashicorp/terraform-provider-google-beta/pull/8799))

BUG FIXES:
* apigee: fixed error of update in `google_apigee_developer` resource ([#8728](https://github.com/hashicorp/terraform-provider-google-beta/pull/8728))
* apigee: made `google_apigee_organization` wait for deletion operation to complete. ([#8795](https://github.com/hashicorp/terraform-provider-google-beta/pull/8795))
* cloudfunctions: fixed issue when updating `vpc_connector_egress_settings` field for `google_cloudfunctions_function` resource. ([#8755](https://github.com/hashicorp/terraform-provider-google-beta/pull/8755))
* dataproc: ensured oneOf condition is honored when expanding the job configuration for Hive, Pig, Spark-sql, and Presto in `google_dataproc_job`. ([#8765](https://github.com/hashicorp/terraform-provider-google-beta/pull/8765))
* gkehub: fixed allowable value `INSTALLATION_UNSPECIFIED` in `template_library.installation` ([#8831](https://github.com/hashicorp/terraform-provider-google-beta/pull/8831))
* sql: fixed edition downgrade failure for an `ENTERPRISE_PLUS` instance with data cache enabled. ([#8731](https://github.com/hashicorp/terraform-provider-google-beta/pull/8731))

## 6.12.0 (November 18, 2024)

FEATURES:
* **New Data Source:** `google_access_context_manager_access_policy` ([#8676](https://github.com/hashicorp/terraform-provider-google-beta/pull/8676))
* **New Data Source:** `google_backup_dr_data_source` ([#8641](https://github.com/hashicorp/terraform-provider-google-beta/pull/8641))
* **New Resource:** `google_dataproc_gdc_spark_application` ([#8662](https://github.com/hashicorp/terraform-provider-google-beta/pull/8662))
* **New Resource:** `google_iam_folders_policy_binding` ([#8677](https://github.com/hashicorp/terraform-provider-google-beta/pull/8677))
* **New Resource:** `google_iam_organizations_policy_binding` ([#8679](https://github.com/hashicorp/terraform-provider-google-beta/pull/8679))

IMPROVEMENTS:
* artifactregistry: added `common_repository` field to `google_artifact_registry_repository` resource ([#8681](https://github.com/hashicorp/terraform-provider-google-beta/pull/8681))
* backupdr: added `access_restriction` field to`google_backup_dr_backup_vault` resource (beta) ([#8656](https://github.com/hashicorp/terraform-provider-google-beta/pull/8656))
* cloudrunv2: added `urls` output field to `google_cloud_run_v2_service` resource ([#8686](https://github.com/hashicorp/terraform-provider-google-beta/pull/8686))
* compute: added `IDPF` as a possible value for the `network_interface.nic_type` field in `google_compute_instance` resource ([#8664](https://github.com/hashicorp/terraform-provider-google-beta/pull/8664))
* compute: added `IDPF` as a possible value for the `guest_os_features.type` field in `google_compute_image` resource ([#8664](https://github.com/hashicorp/terraform-provider-google-beta/pull/8664))
* compute: added `replica_names` field to `sql_database_instance` resource ([#8637](https://github.com/hashicorp/terraform-provider-google-beta/pull/8637))
* filestore: added `performance_config` field to `google_filestore_instance` resource ([#8647](https://github.com/hashicorp/terraform-provider-google-beta/pull/8647))
* redis: added `persistence_config` to `google_redis_cluster`. ([#8643](https://github.com/hashicorp/terraform-provider-google-beta/pull/8643))
* securesourcemanager: added `workforce_identity_federation_config` field to `google_secure_source_manager_instance` resource ([#8670](https://github.com/hashicorp/terraform-provider-google-beta/pull/8670))
* spanner: added `default_backup_schedule_type` field to  `google_spanner_instance` ([#8644](https://github.com/hashicorp/terraform-provider-google-beta/pull/8644))
* sql: added `psc_auto_connections` fields to `google_sql_database_instance` resource ([#8682](https://github.com/hashicorp/terraform-provider-google-beta/pull/8682))

BUG FIXES:
* accesscontextmanager: fixed permadiff in perimeter `google_access_context_manager_service_perimeter_ingress_policy` and `google_access_context_manager_service_perimeter_egress_policy` resources when there are duplicate resources in the rules ([#8675](https://github.com/hashicorp/terraform-provider-google-beta/pull/8675))
* accesscontextmanager: fixed comparison of `identity_type` in `ingress_from` and `egress_from` when the `IDENTITY_TYPE_UNSPECIFIED` is set ([#8648](https://github.com/hashicorp/terraform-provider-google-beta/pull/8648))
* compute: fixed permadiff on attempted `type` field updates in `google_computer_security_policy`, updating this field will now force recreation of the resource ([#8689](https://github.com/hashicorp/terraform-provider-google-beta/pull/8689))
* identityplatform: fixed perma-diff in `google_identity_platform_config` ([#8663](https://github.com/hashicorp/terraform-provider-google-beta/pull/8663))

## 6.11.2 (November 15, 2024)

BUG FIXES:
* vertexai: fixed issue with google_vertex_ai_endpoint where upgrading to 6.11.0 would delete all traffic splits that were set outside Terraform (which was previously a required step for all meaningful use of this resource). ([#8708](https://github.com/hashicorp/terraform-provider-google-beta/pull/8708))

## 6.11.1 (November 12, 2024)

BUG FIXES:
* container: fixed diff on `google_container_cluster.user_managed_keys_config` field for resources that had not set it. ([#8687](https://github.com/hashicorp/terraform-provider-google-beta/pull/8687))
* container: marked `google_container_cluster.user_managed_keys_config` as immutable because it can't be updated in place. ([#8687](https://github.com/hashicorp/terraform-provider-google-beta/pull/8687))

## 6.11.0 (November 11, 2024)

NOTES:
* compute: migrated `google_compute_firewall_policy_rule` from DCL engine to MMv1 engine. ([#8604](https://github.com/hashicorp/terraform-provider-google-beta/pull/8604))

BREAKING CHANGES:
* looker: made `oauth_config` a required field in `google_looker_instance`, as creating this resource without that field always triggers an API error ([#8633](https://github.com/hashicorp/terraform-provider-google-beta/pull/8633))

DEPRECATIONS:
* backupdr: deprecated `force_delete` on `google_backup_dr_backup_vault`. Use `ignore_inactive_datasources` instead ([#8616](https://github.com/hashicorp/terraform-provider-google-beta/pull/8616))

FEATURES:
* **New Data Source:** `google_backup_dr_backup_plan_association` ([#8632](https://github.com/hashicorp/terraform-provider-google-beta/pull/8632))
* **New Data Source:** `google_backup_dr_backup_plan` ([#8603](https://github.com/hashicorp/terraform-provider-google-beta/pull/8603))
* **New Data Source:** `google_spanner_database` ([#8568](https://github.com/hashicorp/terraform-provider-google-beta/pull/8568))
* **New Resource:** `google_apigee_api` ([#8567](https://github.com/hashicorp/terraform-provider-google-beta/pull/8567))
* **New Resource:** `google_backup_dr_backup_plan_association` ([#8632](https://github.com/hashicorp/terraform-provider-google-beta/pull/8632))
* **New Resource:** `google_backup_dr_backup_plan` ([#8603](https://github.com/hashicorp/terraform-provider-google-beta/pull/8603))
* **New Resource:** `google_compute_region_resize_request` ([#8588](https://github.com/hashicorp/terraform-provider-google-beta/pull/8588))
* **New Resource:** `google_dataproc_gdc_application_environment` ([#8609](https://github.com/hashicorp/terraform-provider-google-beta/pull/8609))
* **New Resource:** `google_dataproc_gdc_service_instance` ([#8591](https://github.com/hashicorp/terraform-provider-google-beta/pull/8591))
* **New Resource:** `google_iam_principal_access_boundary_policy` ([#8634](https://github.com/hashicorp/terraform-provider-google-beta/pull/8634))
* **New Resource:** `google_network_management_vpc_flow_logs_config` ([#8623](https://github.com/hashicorp/terraform-provider-google-beta/pull/8623))

IMPROVEMENTS:
* apigee: added in-place update support for `google_apigee_env_references` ([#8621](https://github.com/hashicorp/terraform-provider-google-beta/pull/8621))
* apigee: added in-place update support for `google_apigee_environment` resource ([#8627](https://github.com/hashicorp/terraform-provider-google-beta/pull/8627))
* backupdr: added `ignore_inactive_datasources` and `ignore_backup_plan_references` fields to `google_backup_dr_backup_vault` resource ([#8616](https://github.com/hashicorp/terraform-provider-google-beta/pull/8616))
* bigquery: added `external_catalog_dataset_options` fields to `google_bigquery_dataset` resource ([#8558](https://github.com/hashicorp/terraform-provider-google-beta/pull/8558))
* cloudrunv2: added `gcs.mount_options` to `google_cloud_run_v2_service` and `google_cloud_run_v2_job` ([#8613](https://github.com/hashicorp/terraform-provider-google-beta/pull/8613))
* compute: added `rules` property to `google_compute_region_security_policy` resource ([#8574](https://github.com/hashicorp/terraform-provider-google-beta/pull/8574))
* compute: added `disks` field to `google_compute_node_template` resource ([#8620](https://github.com/hashicorp/terraform-provider-google-beta/pull/8620))
* compute: added `replica_names` field to `sql_database_instance` resource ([#8637](https://github.com/hashicorp/terraform-provider-google-beta/pull/8637))
* compute: added new field `instance_flexibility_policy` to resource `google_compute_region_instance_group_manager` ([#8581](https://github.com/hashicorp/terraform-provider-google-beta/pull/8581))
* compute: increased `google_compute_security_policy` timeouts from 20 minutes to 30 minutes ([#8589](https://github.com/hashicorp/terraform-provider-google-beta/pull/8589))
* container: added `control_plane_endpoints_config` field to `google_container_cluster` resource. ([#8630](https://github.com/hashicorp/terraform-provider-google-beta/pull/8630))
* container: added `parallelstore_csi_driver_config` field to `google_container_cluster` resource. ([#8607](https://github.com/hashicorp/terraform-provider-google-beta/pull/8607))
* container: added `user_managed_keys_config` field to `google_container_cluster` resource. ([#8562](https://github.com/hashicorp/terraform-provider-google-beta/pull/8562))
* firestore: allowed single field indexes to support `__name__ DESC` indexes in `google_firestore_index` resources ([#8576](https://github.com/hashicorp/terraform-provider-google-beta/pull/8576))
* privateca: added support for `google_privateca_certificate_authority` with type = "SUBORDINATE" to be activated into "STAGED" state ([#8560](https://github.com/hashicorp/terraform-provider-google-beta/pull/8560))
* spanner: added `default_backup_schedule_type` field to  `google_spanner_instance` ([#8644](https://github.com/hashicorp/terraform-provider-google-beta/pull/8644))
* vertexai: added `traffic_split`, `private_service_connect_config`, `predict_request_response_logging_config`, `dedicated_endpoint_enabled`, and `dedicated_endpoint_dns` fields to `google_vertex_ai_endpoint` resource ([#8619](https://github.com/hashicorp/terraform-provider-google-beta/pull/8619))
* workflows: added `deletion_protection` field to `google_workflows_workflow` resource ([#8563](https://github.com/hashicorp/terraform-provider-google-beta/pull/8563))

BUG FIXES:
* compute: fixed a diff based on server-side reordering of `match.src_address_groups` and `match.dest_address_groups` in `google_compute_network_firewall_policy_rule` ([#8592](https://github.com/hashicorp/terraform-provider-google-beta/pull/8592))
* compute: fixed permadiff on the `preconfigured_waf_config` field for `google_compute_security_policy` resource ([#8622](https://github.com/hashicorp/terraform-provider-google-beta/pull/8622))
* container: fixed in-place updates for `node_config.containerd_config` in `google_container_cluster` and `google_container_node_pool` ([#8566](https://github.com/hashicorp/terraform-provider-google-beta/pull/8566))

## 6.10.0 (November 4, 2024)

FEATURES:
* **New Data Source:** `google_compute_instance_guest_attributes` ([#8556](https://github.com/hashicorp/terraform-provider-google-beta/pull/8556))
* **New Data Source:** `google_service_accounts` ([#8532](https://github.com/hashicorp/terraform-provider-google-beta/pull/8532))
* **New Resource:** `google_iap_settings` ([#8548](https://github.com/hashicorp/terraform-provider-google-beta/pull/8548))

IMPROVEMENTS:
* apphub: added `GLOBAL` enum value to `scope.type` field in `google_apphub_application` resource ([#8504](https://github.com/hashicorp/terraform-provider-google-beta/pull/8504))
* assuredworkloads: added `workload_options` field to `google_assured_workloads_workload` resource ([#8495](https://github.com/hashicorp/terraform-provider-google-beta/pull/8495))
* backupdr: marked `networks` field optional in `google_backup_dr_management_server` resource ([#8594](https://github.com/hashicorp/terraform-provider-google-beta/pull/8594))
* bigquery: added `external_catalog_dataset_options` fields to `google_bigquery_dataset` resource (beta) ([#8558](https://github.com/hashicorp/terraform-provider-google-beta/pull/8558))
* bigquery: added descriptive validation errors for missing required fields in `google_bigquery_job` destination table configuration ([#8542](https://github.com/hashicorp/terraform-provider-google-beta/pull/8542))
* compute: `desired_status` on google_compute_instance can now be set to `TERMINATED` or `SUSPENDED` on instance creation ([#8515](https://github.com/hashicorp/terraform-provider-google-beta/pull/8515))
* compute: added `header_action` and `redirect_options` fields  to `google_compute_security_policy_rule` resource ([#8544](https://github.com/hashicorp/terraform-provider-google-beta/pull/8544))
* compute: added `interface.ipv6-address` field in `google_compute_external_vpn_gateway` resource ([#8552](https://github.com/hashicorp/terraform-provider-google-beta/pull/8552))
* compute: added plan-time validation to `name` on `google_compute_instance` ([#8520](https://github.com/hashicorp/terraform-provider-google-beta/pull/8520))
* compute: added support for `advanced_machine_features.turbo_mode` to `google_compute_instance`, `google_compute_instance_template`, and `google_compute_region_instance_template` ([#8551](https://github.com/hashicorp/terraform-provider-google-beta/pull/8551))
* container: added in-place update support for `labels`, `resource_manager_tags` and `workload_metadata_config` in `google_container_cluster.node_config` ([#8522](https://github.com/hashicorp/terraform-provider-google-beta/pull/8522))
* memorystore: added `mode` flag to `google_memorystore_instance` ([#8498](https://github.com/hashicorp/terraform-provider-google-beta/pull/8498))
* resourcemanager: added `disabled` to `google_service_account` datasource ([#8518](https://github.com/hashicorp/terraform-provider-google-beta/pull/8518))
* spanner: added `asymmetric_autoscaling_options` field to  `google_spanner_instance` ([#8503](https://github.com/hashicorp/terraform-provider-google-beta/pull/8503))
* sql: removed the client-side default of `ENTERPRISE` for `edition` in `google_sql_database_instance` so that `edition` is determined by the API when unset. This will cause new instances to use `ENTERPRISE_PLUS` as the default for POSTGRES_16. ([#8490](https://github.com/hashicorp/terraform-provider-google-beta/pull/8490))
* vmwareengine: added `autoscaling_settings` to `google_vmwareengine_private_cloud` resource ([#8529](https://github.com/hashicorp/terraform-provider-google-beta/pull/8529))

BUG FIXES:
* accesscontextmanager: fixed permadiff for perimeter ingress / egress rule resources ([#8526](https://github.com/hashicorp/terraform-provider-google-beta/pull/8526))
* compute: fixed an error in `google_compute_region_security_policy_rule` that prevented updating the default rule ([#8535](https://github.com/hashicorp/terraform-provider-google-beta/pull/8535))
* compute: fixed an error in `google_compute_security_policy_rule` that prevented updating the default rule ([#8535](https://github.com/hashicorp/terraform-provider-google-beta/pull/8535))
* container: fixed missing in-place updates for some `google_container_cluster.node_config` subfields ([#8522](https://github.com/hashicorp/terraform-provider-google-beta/pull/8522))

## 6.9.0 (October 28, 2024)

DEPRECATIONS:
* containerattached: deprecated `security_posture_config` field in `google_container_attached_cluster` resource ([#8446](https://github.com/hashicorp/terraform-provider-google-beta/pull/8446))

FEATURES:
* **New Data Source:** `google_oracle_database_autonomous_database` ([#8440](https://github.com/hashicorp/terraform-provider-google-beta/pull/8440))
* **New Data Source:** `google_oracle_database_autonomous_databases` ([#8438](https://github.com/hashicorp/terraform-provider-google-beta/pull/8438))
* **New Data Source:** `google_oracle_database_cloud_exadata_infrastructures` ([#8430](https://github.com/hashicorp/terraform-provider-google-beta/pull/8430))
* **New Data Source:** `google_oracle_database_cloud_vm_clusters` ([#8437](https://github.com/hashicorp/terraform-provider-google-beta/pull/8437))
* **New Resource:** `google_apigee_app_group` ([#8451](https://github.com/hashicorp/terraform-provider-google-beta/pull/8451))
* **New Resource:** `google_apigee_developer` ([#8445](https://github.com/hashicorp/terraform-provider-google-beta/pull/8445))
* **New Resource:** `google_network_connectivity_group` ([#8439](https://github.com/hashicorp/terraform-provider-google-beta/pull/8439))

IMPROVEMENTS:
* compute: `google_compute_network_firewall_policy_association` now uses MMv1 engine instead of DCL. ([#8489](https://github.com/hashicorp/terraform-provider-google-beta/pull/8489))
* compute: `google_compute_region_network_firewall_policy_association` now uses MMv1 engine instead of DCL. ([#8489](https://github.com/hashicorp/terraform-provider-google-beta/pull/8489))
* compute: added `creation_timestamp` field to `google_compute_instance`, `google_compute_instance_template`, `google_compute_region_instance_template` ([#8442](https://github.com/hashicorp/terraform-provider-google-beta/pull/8442))
* compute: added `key_revocation_action_type` to `google_compute_instance` and related resources ([#8473](https://github.com/hashicorp/terraform-provider-google-beta/pull/8473))
* looker: added `deletion_policy` to `google_looker_instance` to allow force-destroying instances with nested resources by setting `deletion_policy = FORCE` ([#8453](https://github.com/hashicorp/terraform-provider-google-beta/pull/8453))
* monitoring: added `alert_strategy.notification_prompts` field to `google_monitoring_alert_policy` ([#8457](https://github.com/hashicorp/terraform-provider-google-beta/pull/8457))
* storage: added `hierarchical_namespace` to `google_storage_bucket` resource ([#8428](https://github.com/hashicorp/terraform-provider-google-beta/pull/8428))
* sql: removed the client-side default of `ENTERPRISE` for `edition` in `google_sql_database_instance` so that `edition` is determined by the API when unset. This will cause new instances to use `ENTERPRISE_PLUS` as the default for POSTGRES_16. ([#8490](https://github.com/hashicorp/terraform-provider-google-beta/pull/8490))
* vmwareengine: added `autoscaling_settings` to `google_vmwareengine_cluster` resource ([#8477](https://github.com/hashicorp/terraform-provider-google-beta/pull/8477))
* workstations: added `max_usable_workstations` field to `google_workstations_workstation_config` resource. ([#8421](https://github.com/hashicorp/terraform-provider-google-beta/pull/8421))

BUG FIXES:
* compute: fixed an issue where immutable `distribution_zones` was incorrectly sent to the API when updating `distribution_policy_target_shape` in `google_compute_region_instance_group_manager` resource ([#8470](https://github.com/hashicorp/terraform-provider-google-beta/pull/8470))
* container: fixed a crash in `google_container_node_pool` caused by an occasional nil pointer ([#8452](https://github.com/hashicorp/terraform-provider-google-beta/pull/8452))
* essentialcontacts: fixed `google_essential_contacts_contact` import to include required parent field. ([#8423](https://github.com/hashicorp/terraform-provider-google-beta/pull/8423))
* sql: made `google_sql_database_instance.0.settings.0.data_cache_config` accept server-side changes when unset. When unset, no diffs will be created when instances change in `edition` and the feature is enabled or disabled as a result. ([#8485](https://github.com/hashicorp/terraform-provider-google-beta/pull/8485))
* storage: removed retry on 404s during refresh for `google_storage_bucket`, preventing hanging when refreshing deleted buckets ([#8478](https://github.com/hashicorp/terraform-provider-google-beta/pull/8478))


## 6.8.0 (October 21, 2024)

FEATURES:
* **New Data Source:** `google_oracle_database_cloud_exadata_infrastructure` ([#8407](https://github.com/hashicorp/terraform-provider-google-beta/pull/8407))
* **New Data Source:** `google_oracle_database_cloud_vm_cluster` ([#8410](https://github.com/hashicorp/terraform-provider-google-beta/pull/8410))
* **New Data Source:** `google_oracle_database_db_nodes` ([#8420](https://github.com/hashicorp/terraform-provider-google-beta/pull/8420))
* **New Data Source:** `google_oracle_database_db_servers` ([#8389](https://github.com/hashicorp/terraform-provider-google-beta/pull/8389))
* **New Resource:** `google_oracle_database_autonomous_database` ([#8411](https://github.com/hashicorp/terraform-provider-google-beta/pull/8411))
* **New Resource:** `google_oracle_database_cloud_exadata_infrastructure` ([#8371](https://github.com/hashicorp/terraform-provider-google-beta/pull/8371))
* **New Resource:** `google_oracle_database_cloud_vm_cluster` ([#8397](https://github.com/hashicorp/terraform-provider-google-beta/pull/8397))
* **New Resource:** `google_transcoder_job_template` ([#8406](https://github.com/hashicorp/terraform-provider-google-beta/pull/8406))
* **New Resource:** `google_transcoder_job` ([#8406](https://github.com/hashicorp/terraform-provider-google-beta/pull/8406))

IMPROVEMENTS:
* cloudfunctions: increased the timeouts to 20 minutes for `google_cloudfunctions_function` resource ([#8372](https://github.com/hashicorp/terraform-provider-google-beta/pull/8372))
* cloudrunv2: added `invoker_iam_disabled` field to `google_cloud_run_v2_service` ([#8395](https://github.com/hashicorp/terraform-provider-google-beta/pull/8395))
* compute: made `google_compute_network_firewall_policy_rule` use MMv1 engine instead of DCL. ([#8412](https://github.com/hashicorp/terraform-provider-google-beta/pull/8412))
* compute: made `google_compute_region_network_firewall_policy_rule` use MMv1 engine instead of DCL. ([#8412](https://github.com/hashicorp/terraform-provider-google-beta/pull/8412))
* compute: added `ip_address_selection_policy` field to `google_compute_backend_service` and `google_compute_region_backend_service`. ([#8413](https://github.com/hashicorp/terraform-provider-google-beta/pull/8413))
* compute: added `provisioned_throughput` field to `google_compute_instance_template` resource ([#8405](https://github.com/hashicorp/terraform-provider-google-beta/pull/8405))
* compute: added `provisioned_throughput` field to `google_compute_region_instance_template` resource ([#8405](https://github.com/hashicorp/terraform-provider-google-beta/pull/8405))
* container: `google_container_cluster` will now accept server-specified values for `node_pool_auto_config.0.node_kubelet_config` when it is not defined in configuration and will not detect drift. Note that this means that removing the value from configuration will now preserve old settings instead of reverting the old settings. ([#8385](https://github.com/hashicorp/terraform-provider-google-beta/pull/8385))
* container: added support for additional values `KCP_CONNECTION`, and `KCP_SSHD`in `google_container_cluster.logging_config` ([#8381](https://github.com/hashicorp/terraform-provider-google-beta/pull/8381))
* dialogflowcx: added `advanced_settings.logging_settings` and `advanced_settings.speech_settings` to `google_dialogflow_cx_agent` and `google_dialogflow_cx_flow` ([#8374](https://github.com/hashicorp/terraform-provider-google-beta/pull/8374))
* networkconnectivity: added `linked_producer_vpc_network` field to `google_network_connectivity_spoke` resource ([#8376](https://github.com/hashicorp/terraform-provider-google-beta/pull/8376))
* secretmanager: added `is_secret_data_base64` field to `google_secret_manager_secret_version` and `google_secret_manager_secret_version_access` datasources ([#8394](https://github.com/hashicorp/terraform-provider-google-beta/pull/8394))
* secretmanager: added `is_secret_data_base64` field to `google_secret_manager_regional_secret_version` and `google_secret_manager_regional_secret_version_access` datasources ([#8394](https://github.com/hashicorp/terraform-provider-google-beta/pull/8394))
* spanner: added `kms_key_names` to `encryption_config` in `google_spanner_database` ([#8403](https://github.com/hashicorp/terraform-provider-google-beta/pull/8403))
* workstations: added `max_usable_workstations` field to `google_workstations_workstation_config` resource ([#8421](https://github.com/hashicorp/terraform-provider-google-beta/pull/8421))
* workstations: added field `allowed_ports` to `google_workstations_workstation_config` ([#8402](https://github.com/hashicorp/terraform-provider-google-beta/pull/8402))

BUG FIXES:
* bigquery: fixed a regression that caused `google_bigquery_dataset_iam_*` resources to attempt to set deleted IAM members, thereby triggering an API error ([#8408](https://github.com/hashicorp/terraform-provider-google-beta/pull/8408))
* compute: fixed an issue in `google_compute_backend_service` and `google_compute_region_backend_service` to allow sending `false` for `iap.enabled` ([#8369](https://github.com/hashicorp/terraform-provider-google-beta/pull/8369))
* container: `node_config.linux_node_config`, `node_config.workload_metadata_config` and `node_config.kubelet_config` will now successfully send empty messages to the API when `terraform plan` indicates they are being removed, rather than null, which caused an error. The sole reliable case is `node_config.linux_node_config` when the block is removed, where there will still be a permadiff, but the update request that's triggered will no longer error and other changes displayed in the plan should go through. ([#8400](https://github.com/hashicorp/terraform-provider-google-beta/pull/8400))

## 6.7.0 (October 14, 2024)

FEATURES:
* **New Resource:** `google_healthcare_pipeline_job` ([#8330](https://github.com/hashicorp/terraform-provider-google-beta/pull/8330))
* **New Resource:** `google_secure_source_manager_branch_rule` ([#8360](https://github.com/hashicorp/terraform-provider-google-beta/pull/8360))

IMPROVEMENTS:
* container: `google_container_cluster` will now accept server-specified values for `node_pool_auto_config.0.node_kubelet_config` when it is not defined in configuration and will not detect drift. Note that this means that removing the value from configuration will now preserve old settings instead of reverting the old settings. ([#8385](https://github.com/hashicorp/terraform-provider-google-beta/pull/8385))
* discoveryengine: added `chat_engine_config.dialogflow_agent_to_link` field to `google_discovery_engine_chat_engine` resource ([#8333](https://github.com/hashicorp/terraform-provider-google-beta/pull/8333))
* networkconnectivity: added field `migration` to resource `google_network_connectivity_internal_range` ([#8350](https://github.com/hashicorp/terraform-provider-google-beta/pull/8350))
* networkservices: added `routing_mode` field to `google_network_services_gateway` resource ([#8355](https://github.com/hashicorp/terraform-provider-google-beta/pull/8355))

BUG FIXES:
* bigtable: fixed an error where BigTable IAM resources could be created with conditions but the condition was not stored in state ([#8334](https://github.com/hashicorp/terraform-provider-google-beta/pull/8334))
* container: fixed issue which caused to not being able to disable `enable_cilium_clusterwide_network_policy` field on `google_container_cluster`. ([#8338](https://github.com/hashicorp/terraform-provider-google-beta/pull/8338))
* container: fixed a diff triggered by a new API-side default value for `node_config.0.kubelet_config.0.insecure_kubelet_readonly_port_enabled`. Terraform will now accept server-specified values for `node_config.0.kubelet_config` when it is not defined in configuration and will not detect drift. Note that this means that removing the value from configuration will now preserve old settings instead of reverting the old settings. ([#8385](https://github.com/hashicorp/terraform-provider-google-beta/pull/8385))
* dataproc: fixed a bug in `google_dataproc_cluster` that prevented creation of clusters with `internal_ip_only` set to false ([#8363](https://github.com/hashicorp/terraform-provider-google-beta/pull/8363))
* iam: addressed `google_service_account` creation issues caused by the eventual consistency of the GCP IAM API by ignoring 403 errors returned on polling the service account after creation. ([#8336](https://github.com/hashicorp/terraform-provider-google-beta/pull/8336))
* logging: fixed the whitespace permadiff on `exclusions.filter` field in `google_logging_billing_account_sink`, `google_logging_folder_sink`, `google_logging_organization_sink` and `google_logging_project_sink` resources ([#8343](https://github.com/hashicorp/terraform-provider-google-beta/pull/8343))
* pubsub: fixed permadiff with configuring an empty `retry_policy` in `google_pubsub_subscription`.  This will result in `minimum_backoff` and `maximum_backoff` using server-side defaults. To use "immedate retry", do not specify a `retry_policy` block at all. ([#8365](https://github.com/hashicorp/terraform-provider-google-beta/pull/8365))
* secretmanager: fixed the issue of unpopulated fields `labels`, `annotations` and `version_destroy_ttl` in the terraform state for the `google_secret_manager_secrets` datasource ([#8346](https://github.com/hashicorp/terraform-provider-google-beta/pull/8346))

## 6.6.0 (October 7, 2024)

FEATURES:
* **New Resource:** `google_dataproc_batch` ([#8306](https://github.com/hashicorp/terraform-provider-google-beta/pull/8306))
* **New Resource:** `google_healthcare_pipeline_job` ([#8330](https://github.com/hashicorp/terraform-provider-google-beta/pull/8330))
* **New Resource:** `google_site_verification_owner` ([#8287](https://github.com/hashicorp/terraform-provider-google-beta/pull/8287))

IMPROVEMENTS:
* assuredworkloads: added `HEALTHCARE_AND_LIFE_SCIENCES_CONTROLS` and `HEALTHCARE_AND_LIFE_SCIENCES_CONTROLS_WITH_US_SUPPORT` enum values to `compliance_regime` in the `google_assured_workloads_workload` resource ([#8326](https://github.com/hashicorp/terraform-provider-google-beta/pull/8326))
* compute: added `bgp_best_path_selection_mode `,`bgp_bps_always_compare_med` and `bgp_bps_inter_region_cost ` fields to `google_compute_network` resource ([#8321](https://github.com/hashicorp/terraform-provider-google-beta/pull/8321))
* compute: added `next_hop_origin `,`next_hop_med ` and `next_hop_inter_region_cost ` output fields to `google_compute_route` resource ([#8321](https://github.com/hashicorp/terraform-provider-google-beta/pull/8321))
* compute: added enum `STATEFUL_COOKIE_AFFINITY` and `strong_session_affinity_cookie` field to `google_compute_backend_service` and `google_compute_region_backend_service` resource ([#8296](https://github.com/hashicorp/terraform-provider-google-beta/pull/8296))
* compute: added `TDX` instance option for `confidential_instance_type` in `google_compute_instance` ([#8320](https://github.com/hashicorp/terraform-provider-google-beta/pull/8320))
* containeraws: added `kubelet_config` field group to the `google_container_aws_node_pool` resource ([#8326](https://github.com/hashicorp/terraform-provider-google-beta/pull/8326))
* dataproc: switched to the v1 API for `google_dataproc_autoscaling_policy` resource ([#8306](https://github.com/hashicorp/terraform-provider-google-beta/pull/8306))
* pubsub: added GCS ingestion settings and platform log settings to `google_pubsub_topic` resource ([#8298](https://github.com/hashicorp/terraform-provider-google-beta/pull/8298))
* sourcerepo: added `create_ignore_already_exists` field to `google_sourcerepo_repository` resource ([#8329](https://github.com/hashicorp/terraform-provider-google-beta/pull/8329))
* sql: added in-place update support for `settings.time_zone` in `google_sql_database_instance` resource ([#8293](https://github.com/hashicorp/terraform-provider-google-beta/pull/8293))
* tags: increased maximum accepted input length for the `short_name` field in `google_tags_tag_key` and `google_tags_tag_value` resources ([#8324](https://github.com/hashicorp/terraform-provider-google-beta/pull/8324))

BUG FIXES:
* bigquery: fixed `google_bigquery_dataset_iam_member` to be able to delete itself and overwrite the existing iam members for bigquery dataset keeping the authorized datasets as they are. ([#8304](https://github.com/hashicorp/terraform-provider-google-beta/pull/8304))
* bigquery: fixed an error which could occur with service account field values containing non-lower-case characters in `google_bigquery_dataset_access` ([#8319](https://github.com/hashicorp/terraform-provider-google-beta/pull/8319))
* compute: fixed an issue where the `boot_disk.initialize_params.resource_policies` field in `google_compute_instance` forced a resource recreation when used in combination with `google_compute_disk_resource_policy_attachment` ([#8309](https://github.com/hashicorp/terraform-provider-google-beta/pull/8309))
* compute: fixed the issue that `labels` was not set when creating the resource `google_compute_interconnect` ([#8284](https://github.com/hashicorp/terraform-provider-google-beta/pull/8284))
* tags:  removed `google_tags_location_tag_binding` resource from the Terraform state when its parent resource has been removed outside of Terraform ([#8310](https://github.com/hashicorp/terraform-provider-google-beta/pull/8310))
* workbench: fixed a bug in the `google_workbench_instance` resource where the removal of `labels` was not functioning as expected. ([#8280](https://github.com/hashicorp/terraform-provider-google-beta/pull/8280))

## 6.5.0 (September 30, 2024)
DEPRECATIONS:
* compute: deprecated `macsec.pre_shared_keys.fail_open` field in `google_compute_interconnect` resource. Use the new `macsec.fail_open` field instead ([#8245](https://github.com/hashicorp/terraform-provider-google-beta/pull/8245))

FEATURES:
* **New Data Source:** `google_compute_region_instance_group_manager` ([#8259](https://github.com/hashicorp/terraform-provider-google-beta/pull/8259))
* **New Data Source:** `google_privileged_access_manager_entitlement` ([#8253](https://github.com/hashicorp/terraform-provider-google-beta/pull/8253))
* **New Data Source:** `google_secret_manager_regional_secret_version_access` ([#8220](https://github.com/hashicorp/terraform-provider-google-beta/pull/8220))
* **New Data Source:** `google_secret_manager_regional_secret_version` ([#8209](https://github.com/hashicorp/terraform-provider-google-beta/pull/8209))
* **New Data Source:** `google_secret_manager_regional_secrets` ([#8217](https://github.com/hashicorp/terraform-provider-google-beta/pull/8217))
* **New Resource:** `google_compute_region_network_firewall_policy_with_rules` ([#8225](https://github.com/hashicorp/terraform-provider-google-beta/pull/8225))
* **New Resource:** `google_compute_router_nat_address` ([#8227](https://github.com/hashicorp/terraform-provider-google-beta/pull/8227))
* **New Resource:** `google_logging_log_scope` ([#8235](https://github.com/hashicorp/terraform-provider-google-beta/pull/8235))

IMPROVEMENTS:
* apigee: added `activate` field to `google_apigee_nat_address` resource ([#8261](https://github.com/hashicorp/terraform-provider-google-beta/pull/8261))
* bigquery: added `biglake_configuration` field to `google_bigquery_table` resource to support BigLake Managed Tables ([#8221](https://github.com/hashicorp/terraform-provider-google-beta/pull/8221))
* cloudrun: added `node_selector` field to `google_cloud_run_service` resource ([#8216](https://github.com/hashicorp/terraform-provider-google-beta/pull/8216))
* cloudrunv2: added `node_selector` field to `google_cloud_run_v2_service` resource ([#8216](https://github.com/hashicorp/terraform-provider-google-beta/pull/8216))
* compute: added `existing_reservations` field to `google_compute_region_commitment` resource ([#8256](https://github.com/hashicorp/terraform-provider-google-beta/pull/8256))
* compute: added `host_error_timeout_seconds` field to `google_compute_instance` resource ([#8252](https://github.com/hashicorp/terraform-provider-google-beta/pull/8252))
* compute: added `hostname` field to `google_compute_instance` data source ([#8268](https://github.com/hashicorp/terraform-provider-google-beta/pull/8268))
* compute: added `initial_nat_ip` field to `google_compute_router_nat` resource ([#8227](https://github.com/hashicorp/terraform-provider-google-beta/pull/8227))
* compute: added `macsec.fail_open` field to `google_compute_interconnect` resource ([#8245](https://github.com/hashicorp/terraform-provider-google-beta/pull/8245))
* compute: added `SUSPENDED` as a possible value to `desired_state` field in `google_compute_instance` resource ([#8257](https://github.com/hashicorp/terraform-provider-google-beta/pull/8257))
* compute: added import support for `projects/{{project}}/meta-data/{{key}}` format for `google_compute_project_metadata_item` resource ([#8274](https://github.com/hashicorp/terraform-provider-google-beta/pull/8274))
* compute: marked `customer_name` and `location` fields as optional in `google_compute_interconnect` resource to support cross cloud interconnect ([#8279](https://github.com/hashicorp/terraform-provider-google-beta/pull/8279))
* container: added `linux_node_config.hugepages_config` field to `google_container_node_pool` resource ([#8210](https://github.com/hashicorp/terraform-provider-google-beta/pull/8210))
* looker: added `psc_enabled` and `psc_config` fields to `google_looker_instance` resource ([#8211](https://github.com/hashicorp/terraform-provider-google-beta/pull/8211))
* networkconnectivity: added `include_import_ranges` field to `google_network_connectivity_spoke` resource for `linked_vpn_tunnels`, `linked_interconnect_attachments` and `linked_router_appliance_instances` ([#8215](https://github.com/hashicorp/terraform-provider-google-beta/pull/8215))
* secretmanagerregional: added `version_aliases` field to `google_secret_manager_regional_secret` resource ([#8209](https://github.com/hashicorp/terraform-provider-google-beta/pull/8209))
* workbench: increased create timeout to 20 minutes for `google_workbench_instance` resource ([#8228](https://github.com/hashicorp/terraform-provider-google-beta/pull/8228))

BUG FIXES:
* bigquery: fixed in-place update of `google_bigquery_table` resource when `external_data_configuration.schema` field is set ([#8234](https://github.com/hashicorp/terraform-provider-google-beta/pull/8234))
* bigquerydatapolicy: fixed permadiff on `policy_tag` field in `google_bigquery_datapolicy_data_policy` resource ([#8239](https://github.com/hashicorp/terraform-provider-google-beta/pull/8239))
* composer: fixed `storage_config.bucket` field to support a bucket name with or without "gs://" prefix ([#8229](https://github.com/hashicorp/terraform-provider-google-beta/pull/8229))
* container: added support for setting `addons_config.gcp_filestore_csi_driver_config` and `enable_autopilot` in the same `google_container_cluster` ([#8260](https://github.com/hashicorp/terraform-provider-google-beta/pull/8260))
* container: fixed `node_config.kubelet_config` updates in `google_container_cluster` resource ([#8238](https://github.com/hashicorp/terraform-provider-google-beta/pull/8238))
* container: fixed a bug where specifying `node_pool_defaults.node_config_defaults` with `enable_autopilot = true` would cause `google_container_cluster` resource creation failure ([#8223](https://github.com/hashicorp/terraform-provider-google-beta/pull/8223))
* workbench: fixed a bug in the `google_workbench_instance` resource where the removal of `labels` was not functioning as expected ([#8280](https://github.com/hashicorp/terraform-provider-google-beta/pull/8280))


## 6.4.0 (September 23, 2024)

DEPRECATIONS:
* securitycenterv2: deprecated `google_scc_v2_organization_scc_big_query_exports`. Use `google_scc_v2_organization_scc_big_query_export` instead. ([#8166](https://github.com/hashicorp/terraform-provider-google-beta/pull/8166))

FEATURES:
* **New Data Source:** `google_secret_manager_regional_secret_version` ([#8209](https://github.com/hashicorp/terraform-provider-google-beta/pull/8209))
* **New Data Source:** `google_secret_manager_regional_secret` ([#8189](https://github.com/hashicorp/terraform-provider-google-beta/pull/8189))
* **New Resource:** `google_compute_firewall_policy_with_rules` ([#8181](https://github.com/hashicorp/terraform-provider-google-beta/pull/8181))
* **New Resource:** `google_database_migration_service_migration_job` ([#8187](https://github.com/hashicorp/terraform-provider-google-beta/pull/8187))
* **New Resource:** `google_discovery_engine_target_site` ([#8174](https://github.com/hashicorp/terraform-provider-google-beta/pull/8174))
* **New Resource:** `google_healthcare_workspace` ([#8179](https://github.com/hashicorp/terraform-provider-google-beta/pull/8179))
* **New Resource:** `google_scc_folder_scc_big_query_export` ([#8183](https://github.com/hashicorp/terraform-provider-google-beta/pull/8183))
* **New Resource:** `google_scc_organization_scc_big_query_export` ([#8172](https://github.com/hashicorp/terraform-provider-google-beta/pull/8172))
* **New Resource:** `google_scc_project_scc_big_query_export` ([#8173](https://github.com/hashicorp/terraform-provider-google-beta/pull/8173))
* **New Resource:** `google_scc_v2_organization_scc_big_query_export` ([#8166](https://github.com/hashicorp/terraform-provider-google/pull/8166))
* **New Resource:** `google_secret_manager_regional_secret_version` ([#8199](https://github.com/hashicorp/terraform-provider-google-beta/pull/8199))
* **New Resource:** `google_secret_manager_regional_secret` ([#8170](https://github.com/hashicorp/terraform-provider-google-beta/pull/8170))
* **New Resource:** `google_site_verification_web_resource` ([#8180](https://github.com/hashicorp/terraform-provider-google-beta/pull/8180))
* **New Resource:** `google_spanner_backup_schedule` ([#8160](https://github.com/hashicorp/terraform-provider-google-beta/pull/8160))

IMPROVEMENTS:
* alloydb: added `enable_outbound_public_ip` field to `google_alloydb_instance` resource ([#8156](https://github.com/hashicorp/terraform-provider-google-beta/pull/8156))
* apigee: added in-place update for `consumer_accept_list` field in `google_apigee_instance` resource ([#8155](https://github.com/hashicorp/terraform-provider-google-beta/pull/8155))
* compute: added `interface` field to `google_compute_attached_disk` resource ([#8154](https://github.com/hashicorp/terraform-provider-google-beta/pull/8154))
* compute: added in-place update in `google_compute_interconnect` resource except for `remote_location` and `requested_features` fields ([#8203](https://github.com/hashicorp/terraform-provider-google-beta/pull/8203))
* filestore: added `deletion_protection_enabled` and `deletion_protection_reason` fields to `google_filestore_instance` resource ([#8158](https://github.com/hashicorp/terraform-provider-google-beta/pull/8158))
* looker: added `fips_enabled` field to `google_looker_instance` resource ([#8206](https://github.com/hashicorp/terraform-provider-google-beta/pull/8206))
* metastore: added `deletion_protection` field to `google_dataproc_metastore_service` resource ([#8200](https://github.com/hashicorp/terraform-provider-google-beta/pull/8200))
* netapp: added `allow_auto_tiering` field to `google_netapp_storage_pool` resource ([#8163](https://github.com/hashicorp/terraform-provider-google-beta/pull/8163))
* netapp: added `tiering_policy` field to `google_netapp_volume` resource ([#8163](https://github.com/hashicorp/terraform-provider-google-beta/pull/8163))
* secretmanagerregional: added `version_aliases` field to `google_secret_manager_regional_secret` resource ([#8209](https://github.com/hashicorp/terraform-provider-google-beta/pull/8209))
* spanner: added `edition` field to `google_spanner_instance` resource ([#8160](https://github.com/hashicorp/terraform-provider-google-beta/pull/8160))

BUG FIXES:
* compute: fixed a permadiff on `iap` field in `google_compute_backend` and `google_compute_region_backend` resources ([#8204](https://github.com/hashicorp/terraform-provider-google-beta/pull/8204))
* container: fixed a bug where specifying `node_pool_defaults.node_config_defaults` with `enable_autopilot = true` will cause `google_container_cluster` resource creation failure ([#8223](https://github.com/hashicorp/terraform-provider-google-beta/pull/8223))
* container: fixed a permadiff on `node_config.gcfs_config` field in `google_container_cluster` and `google_container_node_pool` resources ([#8207](https://github.com/hashicorp/terraform-provider-google-beta/pull/8207))
* container: fixed the in-place update for `node_config.gcfs_config` in `google_container_cluster` and `google_container_node_pool` resources ([#8207](https://github.com/hashicorp/terraform-provider-google-beta/pull/8207))
* container: made `node_config.kubelet_config.cpu_manager_policy` field optional to fix its update in `google_container_cluster` resource ([#8171](https://github.com/hashicorp/terraform-provider-google-beta/pull/8171))
* dns: fixed a permadiff on `dnssec_config` field in `google_dns_managed_zone` resource ([#8165](https://github.com/hashicorp/terraform-provider-google-beta/pull/8165))
* pubsub: allowed `filter` field to contain line breaks in `google_pubsub_subscription` resource ([#8161](https://github.com/hashicorp/terraform-provider-google-beta/pull/8161))

## 6.3.0 (September 16, 2024)

FEATURES:
* **New Data Source:** `google_bigquery_tables` ([#8130](https://github.com/hashicorp/terraform-provider-google-beta/pull/8130))
* **New Resource:** `google_compute_network_firewall_policy_with_rules` ([#8118](https://github.com/hashicorp/terraform-provider-google-beta/pull/8118))
* **New Resource:** `google_developer_connect_connection` ([#8150](https://github.com/hashicorp/terraform-provider-google-beta/pull/8150))
* **New Resource:** `google_developer_connect_git_repository_link` ([#8150](https://github.com/hashicorp/terraform-provider-google-beta/pull/8150))
* **New Resource:** `google_memorystore_instance` ([#8126](https://github.com/hashicorp/terraform-provider-google-beta/pull/8126))

IMPROVEMENTS:
* compute: added `connected_endpoints.consumer_network` and `connected_endpoints.psc_connection_id` fields to `google_compute_service_attachment` resource ([#8148](https://github.com/hashicorp/terraform-provider-google-beta/pull/8148))
* compute: added `propagated_connection_limit` and `connected_endpoints.propagated_connection_count` fields to `google_compute_service_attachment` resource ([#8148](https://github.com/hashicorp/terraform-provider-google-beta/pull/8148))
* compute: added field `http_keep_alive_timeout_sec` to `google_region_compute_target_http_proxy` and `google_region_compute_target_http_proxy` resources ([#8151](https://github.com/hashicorp/terraform-provider-google-beta/pull/8151))
* compute: added support for `boot_disk.initialize_params.resource_policies` in `google_compute_instance` and `google_instance_template` ([#8134](https://github.com/hashicorp/terraform-provider-google-beta/pull/8134))
* container: added `storage_pools` to `node_config` in `google_container_cluster` and `google_container_node_pool` ([#8146](https://github.com/hashicorp/terraform-provider-google-beta/pull/8146))
* containerattached: added `security_posture_config` field to `google_container_attached_cluster` resource ([#8137](https://github.com/hashicorp/terraform-provider-google-beta/pull/8137))
* netapp: added `large_capacity` and `multiple_endpoints` to `google_netapp_volume` resource ([#8116](https://github.com/hashicorp/terraform-provider-google-beta/pull/8116))
* resourcemanager: added `tags` field to `google_folder` to allow setting tags for folders at creation time ([#8113](https://github.com/hashicorp/terraform-provider-google-beta/pull/8113))

BUG FIXES:
* compute: setting `network_ip` to "" will no longer cause diff and will be treated the same as `null` ([#8128](https://github.com/hashicorp/terraform-provider-google-beta/pull/8128))
* dataproc: updated `google_dataproc_cluster` to protect against handling nil `kerberos_config` values ([#8129](https://github.com/hashicorp/terraform-provider-google-beta/pull/8129))
* dns: added a mutex to `google_dns_record_set` to prevent conflicts when multiple resources attempt to operate on the same record set ([#8139](https://github.com/hashicorp/terraform-provider-google-beta/pull/8139))
* managedkafka: added 5 second wait post `google_managed_kafka_topic` creation to fix eventual consistency errors ([#8149](https://github.com/hashicorp/terraform-provider-google-beta/pull/8149))

## 6.2.0 (September 9, 2024)

FEATURES:
* **New Data Source:** `google_certificate_manager_certificates` ([#8099](https://github.com/hashicorp/terraform-provider-google-beta/pull/8099))
* **New Resource:** `google_backup_dr_backup_vault` ([#8083](https://github.com/hashicorp/terraform-provider-google-beta/pull/8083))
* **New Resource:** `google_scc_v2_folder_scc_big_query_export` ([#8079](https://github.com/hashicorp/terraform-provider-google-beta/pull/8079))
* **New Resource:** `google_scc_v2_project_scc_big_query_export` ([#8070](https://github.com/hashicorp/terraform-provider-google-beta/pull/8070))

IMPROVEMENTS:
* assuredworkload: added field `partner_service_billing_account` to `google_assured_workloads_workload` ([#8097](https://github.com/hashicorp/terraform-provider-google-beta/pull/8097))
* bigtable: added support for `column_family.type` in `google_bigtable_table` ([#8069](https://github.com/hashicorp/terraform-provider-google-beta/pull/8069))
* cloudrunv2: added `template.service_mesh` to `google_cloud_run_v2_service` ([#8096](https://github.com/hashicorp/terraform-provider-google-beta/pull/8096))
* compute: added `boot_disk.interface` field to `google_compute_instance` resource ([#8075](https://github.com/hashicorp/terraform-provider-google-beta/pull/8075))
* container: added `node_pool_auto_config.node_kublet_config.insecure_kubelet_readonly_port_enabled` field to `google_container_cluster`. ([#8076](https://github.com/hashicorp/terraform-provider-google-beta/pull/8076))
* container: added `insecure_kubelet_readonly_port_enabled` to `node_pool.node_config.kubelet_config` and `node_config.kubelet_config` in `google_container_node_pool` resource. ([#8071](https://github.com/hashicorp/terraform-provider-google-beta/pull/8071))
* container: added `insecure_kubelet_readonly_port_enabled` to `node_pool_defaults.node_config_defaults`, `node_pool.node_config.kubelet_config`, and `node_config.kubelet_config` in `google_container_cluster` resource. ([#8071](https://github.com/hashicorp/terraform-provider-google-beta/pull/8071))
* container: added support for in-place updates for `google_compute_node_pool.node_config.gcfs_config` and `google_container_cluster.node_config.gcfs_cluster` and `google_container_cluster.node_pool.node_config.gcfs_cluster` ([#8101](https://github.com/hashicorp/terraform-provider-google-beta/pull/8101))
* iambeta: added `x509` field to `google_iam_workload_identity_pool_provider ` resource ([#8110](https://github.com/hashicorp/terraform-provider-google-beta/pull/8110))
* networkconnectivity: added `include_export_ranges` to `google_network_connectivity_spoke` ([#8088](https://github.com/hashicorp/terraform-provider-google-beta/pull/8088))
* pubsub: added `cloud_storage_config.max_messages` and `cloud_storage_config.avro_config.use_topic_schema` fields to `google_pubsub_subscription` resource ([#8086](https://github.com/hashicorp/terraform-provider-google-beta/pull/8086))
* redis: added the `maintenance_policy` field to the `google_redis_cluster` resource ([#8087](https://github.com/hashicorp/terraform-provider-google-beta/pull/8087))
* resourcemanager: added `tags` field to `google_project` to allow setting tags for projects at creation time ([#8091](https://github.com/hashicorp/terraform-provider-google-beta/pull/8091))
* securitycenter: added support for empty `streaming_config.filter` values in `google_scc_notification_config` resources ([#8105](https://github.com/hashicorp/terraform-provider-google-beta/pull/8105))

BUG FIXES:
* compute: fixed `google_compute_interconnect` to support correct `available_features` option of `IF_MACSEC` ([#8082](https://github.com/hashicorp/terraform-provider-google-beta/pull/8082))
* compute: fixed a bug where `advertised_route_priority` was accidentally set to 0 during updates in `google_compute_router_peer` ([#8102](https://github.com/hashicorp/terraform-provider-google-beta/pull/8102))
* compute: fixed a permadiff caused by setting `start_time` in an incorrect `H:mm` format in `google_compute_resource_policies` resources ([#8067](https://github.com/hashicorp/terraform-provider-google-beta/pull/8067))
* compute: fixed `network_interface.subnetwork_project` validation to match with the project in `network_interface.subnetwork` field when `network_interface.subnetwork` has full self_link in `google_compute_instance` resource ([#8089](https://github.com/hashicorp/terraform-provider-google-beta/pull/8089))
* kms: updated the `google_kms_autokey_config` resource's `folder` field to accept values that are either full resource names (`folders/{folder_id}`) or just the folder id (`{folder_id}` only) ([#8100](https://github.com/hashicorp/terraform-provider-google-beta/pull/8100))
* storage: added retry support for 429 errors in `google_storage_bucket` resource ([#8092](https://github.com/hashicorp/terraform-provider-google-beta/pull/8092))


## 6.1.0 (September 4, 2024)

FEATURES:
* **New Data Source:** `google_kms_crypto_key_latest_version` ([#8032](https://github.com/hashicorp/terraform-provider-google-beta/pull/8032))
* **New Data Source:** `google_kms_crypto_key_versions` ([#8026](https://github.com/hashicorp/terraform-provider-google-beta/pull/8026))

IMPROVEMENTS:
* databasemigrationservice: added support in `google_database_migration_service_connection_profile` for creating DMS connection profiles that link to existing Cloud SQL instances/AlloyDB clusters. ([#8062](https://github.com/hashicorp/terraform-provider-google-beta/pull/8062))
* alloydb: added `subscription_type` and `trial_metadata` field to `google_alloydb_cluster` resource ([#8042](https://github.com/hashicorp/terraform-provider-google-beta/pull/8042))
* bigquery: added `encryption_configuration` field to `google_bigquery_data_transfer_config` resource ([#8045](https://github.com/hashicorp/terraform-provider-google-beta/pull/8045))
* bigqueryanalyticshub: added `selected_resources`, and `restrict_direct_table_access` to `google_bigquery_analytics_hub_listing` resource ([#8029](https://github.com/hashicorp/terraform-provider-google-beta/pull/8029))
* bigqueryanalyticshub: added `sharing_environment_config` to `google_bigquery_analytics_hub_data_exchange` resource ([#8029](https://github.com/hashicorp/terraform-provider-google-beta/pull/8029))
* cloudtasks: added `http_target` field to `google_cloud_tasks_queue` resource ([#8033](https://github.com/hashicorp/terraform-provider-google-beta/pull/8033))
* compute: added `accelerators` field to `google_compute_node_template` resource ([#8063](https://github.com/hashicorp/terraform-provider-google-beta/pull/8063))
* compute: allowed disabling `server_tls_policy` during update in `google_compute_target_https_proxy` resources ([#8023](https://github.com/hashicorp/terraform-provider-google-beta/pull/8023))
* datastream: added `transaction_logs` and `change_tables` to `datastream_stream` resource ([#8031](https://github.com/hashicorp/terraform-provider-google-beta/pull/8031))
* discoveryengine: added `chunking_config` and `layout_parsing_config` fields to `google_discovery_engine_data_store` resource ([#8049](https://github.com/hashicorp/terraform-provider-google-beta/pull/8049))
* dlp: added `inspect_template_modified_cadence` field to `big_query_target` and `cloud_sql_target` in `google_data_loss_prevention_discovery_config` resource ([#8054](https://github.com/hashicorp/terraform-provider-google-beta/pull/8054))
* dlp: added `tag_resources` field to `google_data_loss_prevention_discovery_config` resource ([#8054](https://github.com/hashicorp/terraform-provider-google-beta/pull/8054))

BUG FIXES:
* bigquery: fixed an error which could occur with email field values containing non-lower-case characters in `google_bigquery_dataset_access` resource ([#8039](https://github.com/hashicorp/terraform-provider-google-beta/pull/8039))
* bigqueryanalyticshub: made `bigquery_dataset` immutable in `google_bigquery_analytics_hub_listing` as it was not updatable in the API. Now modifying the field in Terraform will correctly recreate the resource rather than causing Terraform to report it would attempt an invalid update. ([#8029](https://github.com/hashicorp/terraform-provider-google-beta/pull/8029))
* container: fixed update inconsistency in `google_container_cluster` resource ([#8030](https://github.com/hashicorp/terraform-provider-google-beta/pull/8030))
* pubsub: fixed a validation bug that didn't allow empty filter definitions for `google_pubsub_subscription` resources ([#8055](https://github.com/hashicorp/terraform-provider-google-beta/pull/8055))
* resourcemanager: fixed a bug where data.google_client_config failed silently when inadequate credentials were used to configure the provider ([#8057](https://github.com/hashicorp/terraform-provider-google-beta/pull/8057))
* sql: fixed importing `google_sql_user` where `host` is an IPv4 CIDR ([#8028](https://github.com/hashicorp/terraform-provider-google-beta/pull/8028))
* sql: fixed overwriting of `name` field for IAM Group user for `google_sql_user` resource ([#8024](https://github.com/hashicorp/terraform-provider-google-beta/pull/8024))

## 6.0.1 (August 26, 2024)

BREAKING CHANGES:

* sql: removed `settings.ip_configuration.require_ssl` from `google_sql_database_instance` in favor of `settings.ip_configuration.ssl_mode`. This field was intended to be removed in 6.0.0. ([#8043](https://github.com/hashicorp/terraform-provider-google-beta/pull/8043))

## 6.0.0 (August 26, 2024)

[Terraform Google Provider 6.0.0 Upgrade Guide](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/version_6_upgrade)

BREAKING CHANGES:
* provider: changed provider labels to add the `goog-terraform-provisioned: true` label by default. ([#8004](https://github.com/hashicorp/terraform-provider-google-beta/pull/8004))
* activedirectory: added `deletion_protection` field to `google_active_directory_domain` resource. This field defaults to `true`, preventing accidental deletions. To delete the resource, you must first set `deletion_protection = false` before destroying the resource. ([#7837](https://github.com/hashicorp/terraform-provider-google-beta/pull/7837))
* alloydb: removed `network` in `google_alloy_db_cluster`. Use `network_config.network` instead. ([#7999](https://github.com/hashicorp/terraform-provider-google-beta/pull/7999))
* billing: revised the format of `id` for `google_billing_project_info` ([#7793](https://github.com/hashicorp/terraform-provider-google-beta/pull/7793)) 
* bigquery: added client-side validation to prevent table view creation if schema contains required fields for `google_bigquery_table` resource ([#7755](https://github.com/hashicorp/terraform-provider-google-beta/pull/7755))
* bigquery: removed `allow_resource_tags_on_deletion` from `google_bigquery_table`. Resource tags are now always allowed on table deletion. ([#7940](https://github.com/hashicorp/terraform-provider-google-beta/pull/7940))
* bigqueryreservation: removed `multi_region_auxiliary` from `google_bigquery_reservation` ([#7844](https://github.com/hashicorp/terraform-provider-google-beta/pull/7844))
* cloudrunv2: added `deletion_protection` field to `google_cloudrunv2_service` to make deleting them require an explicit intent.  This field defaults to `true`, preventing accidental deletions. To delete the resource, you must first set `deletion_protection = false` before destroying the resource. ([#7901](https://github.com/hashicorp/terraform-provider-google-beta/pull/7901))
* cloudrunv2: changed `liveness_probe` to no longer infer a default value from api on `google_cloud_run_v2_service`. Removing this field and applying the change will now remove liveness probe from the Cloud Run service. ([#7753](https://github.com/hashicorp/terraform-provider-google-beta/pull/7753))
* cloudrunv2: retyped `containers.env` to SET from ARRAY for `google_cloud_run_v2_service` and `google_cloud_run_v2_job`. ([#7812](https://github.com/hashicorp/terraform-provider-google-beta/pull/7812))
* composer: `ip_allocation_policy = []` in `google_composer_environment` is no longer valid configuration. Removing the field from configuration should not produce a diff. ([#8011](https://github.com/hashicorp/terraform-provider-google-beta/pull/8011))
* compute: added new required field `enabled` in `google_compute_backend_service` and `google_compute_region_backend_service` ([#7758](https://github.com/hashicorp/terraform-provider-google-beta/pull/7758))
* compute: revised and in some cases removed default values  of `connection_draining_timeout_sec`, `balancing_mode` and `outlier_detection` in `google_compute_region_backend_service` and `google_compute_backend_service`. ([#7723](https://github.com/hashicorp/terraform-provider-google-beta/pull/7723))
* compute: updated resource id for `compute_network_endpoints` ([#7806](https://github.com/hashicorp/terraform-provider-google-beta/pull/7806))
* compute: stopped the `certifcate_id` field in `google_compute_managed_ssl_certificate` resource being incorrectly marked as a user-configurable value when it should just be an output. ([#7936](https://github.com/hashicorp/terraform-provider-google-beta/pull/7936))
* compute: `guest_accelerator = []` is no longer valid configuration in `google_compute_instance`. To explicitly set an empty list of objects, set guest_accelerator.count = 0. ([#8011](https://github.com/hashicorp/terraform-provider-google-beta/pull/8011))
* compute: `google_compute_instance_from_template` and `google_compute_instance_from_machine_image` `network_interface.alias_ip_range, network_interface.access_config, attached_disk, guest_accelerator, service_account, scratch_disk` can no longer be set to an empty block `[]`. Removing the fields from configuration should not produce a diff. ([#8011](https://github.com/hashicorp/terraform-provider-google-beta/pull/8011))
* compute: `secondary_ip_ranges = []` in `google_compute_subnetwork` is no longer valid configuration. To set an explicitly empty list, use `send_secondary_ip_range_if_empty` and completely remove `secondary_ip_range` from config. ([#8011](https://github.com/hashicorp/terraform-provider-google-beta/pull/8011))
* container: made `advanced_datapath_observability_config.enable_relay` required in `google_container_cluster` ([#7930](https://github.com/hashicorp/terraform-provider-google-beta/pull/7930))
* container: removed deprecated field `advanced_datapath_observability_config.relay_mode` from `google_container_cluster` resource. Users are expected to use `enable_relay` field instead. ([#7930](https://github.com/hashicorp/terraform-provider-google-beta/pull/7930))
* container: three label-related fields are now in `google_container_cluster` resource. `resource_labels` field is non-authoritative and only manages the labels defined by the users on the resource through Terraform. The new output-only `terraform_labels` field merges the labels defined by the users on the resource through Terraform and the default labels configured on the provider. The new output-only `effective_labels` field lists all of labels present on the resource in GCP, including the labels configured through Terraform, the system, and other clients. ([#7932](https://github.com/hashicorp/terraform-provider-google-beta/pull/7932))
* container: made three fields `resource_labels`, `terraform_labels`, and `effective_labels` be present in `google_container_cluster` datasources. All three fields will have all of labels present on the resource in GCP including the labels configured through Terraform, the system, and other clients, equivalent to `effective_labels` on the resource. ([#7932](https://github.com/hashicorp/terraform-provider-google-beta/pull/7932))
* container: `guest_accelerator = []` is no longer valid configuration in `google_container_cluster` and `google_container_node_pool`. To explicitly set an empty list of objects, set guest_accelerator.count = 0. ([#8011](https://github.com/hashicorp/terraform-provider-google-beta/pull/8011))
* container: `guest_accelerator.gpu_driver_installation_config = []` and `guest_accelerator.gpu_sharing_config = []` are no longer valid configuration in `google_container_cluster` and `google_container_node_pool`. Removing the fields from configuration should not produce a diff. ([#8011](https://github.com/hashicorp/terraform-provider-google-beta/pull/8011))
* datastore: removed `google_datastore_index` in favor of `google_firestore_index` ([#7987](https://github.com/hashicorp/terraform-provider-google-beta/pull/7987))
* edgenetwork: three label-related fields are now in `google_edgenetwork_network ` and `google_edgenetwork_subnet` resources. `labels` field is non-authoritative and only manages the labels defined by the users on the resource through Terraform. The new output-only `terraform_labels` field merges the labels defined by the users on the resource through Terraform and the default labels configured on the provider. The new output-only `effective_labels` field lists all of labels present on the resource in GCP, including the labels configured through Terraform, the system, and other clients. ([#7932](https://github.com/hashicorp/terraform-provider-google-beta/pull/7932))
* identityplatform: removed resource `google_identity_platform_project_default_config` in favor of `google_identity_platform_project_config` ([#7880](https://github.com/hashicorp/terraform-provider-google-beta/pull/7880))
* integrations: removed `create_sample_workflows` and `provision_gmek` from `google_integrations_client` ([#7977](https://github.com/hashicorp/terraform-provider-google-beta/pull/7977))
* pubsub: allowed `schema_settings` in `google_pubsub_topic` to be removed ([#7674](https://github.com/hashicorp/terraform-provider-google-beta/pull/7674))
* redis: added a `deletion_protection_enabled` field to the `google_redis_cluster` resource. This field defaults to `true`, preventing accidental deletions. To delete the resource, you must first set `deletion_protection_enabled = false` before destroying the resource. ([#7995](https://github.com/hashicorp/terraform-provider-google-beta/pull/7995))
* resourcemanager: added `deletion_protection` field to `google_folder` to make deleting them require an explicit intent. Folder resources now cannot be destroyed unless `deletion_protection = false` is set for the resource. ([#7903](https://github.com/hashicorp/terraform-provider-google-beta/pull/7903))
* resourcemanager: made `deletion_policy` in `google_project` 'PREVENT' by default. This makes deleting them require an explicit intent. `google_project` resources cannot be destroyed unless `deletion_policy` is set to 'ABANDON' or 'DELETE' for the resource. ([#7946](https://github.com/hashicorp/terraform-provider-google-beta/pull/7946))
* storage: removed `no_age` field from  `lifecycle_rule.condition` in the `google_storage_bucket` resource ([#7923](https://github.com/hashicorp/terraform-provider-google-beta/pull/7923))
* sql: removed `settings.ip_configuration.require_ssl` in `google_sql_database_instance`. Please use `settings.ip_configuration.ssl_mode` instead. ([#7804](https://github.com/hashicorp/terraform-provider-google-beta/pull/7804))
* vpcaccess: removed default values for `min_throughput` and `min_instances` fields on `google_vpc_access_connector` and made them default to values returned from the API when not provided by users ([#7709](https://github.com/hashicorp/terraform-provider-google-beta/pull/7709))
* vpcaccess: added a conflicting fields restriction between `min_throughput` and `min_instances` fields on `google_vpc_access_connector` ([#7709](https://github.com/hashicorp/terraform-provider-google-beta/pull/7709))
* vpcaccess: added a conflicting fields restriction between `max_throughput` and `max_instances` fields on `google_vpc_access_connector` ([#7709](https://github.com/hashicorp/terraform-provider-google-beta/pull/7709))
* workstation: defaulted `host.gce_instance.disable_ssh` to true for `google_workstations_workstation_config` ([#7946](https://github.com/hashicorp/terraform-provider-google-beta/pull/7946))

IMPROVEMENTS:
* compute: added fields `reserved_internal_range` and `secondary_ip_ranges[].reserved_internal_range` to `google_compute_subnetwork` resource ([#7980](https://github.com/hashicorp/terraform-provider-google-beta/pull/7980))
* compute: changed the behavior of `name_prefix` in multiple Compute resources to allow for a longer max length of 54 characters. See the upgrade guide and resource documentation for more details. ([#7981](https://github.com/hashicorp/terraform-provider-google-beta/pull/7981))

BUG FIXES:
* compute: fixed an issue regarding sending `enabled` field by default for null `iap` message in `google_compute_backend_service` and `google_compute_region_backend_service` ([#7758](https://github.com/hashicorp/terraform-provider-google-beta/pull/7758))
