## 7.46.2 (October 23, 2026)

NOTES:
* Updated to Go 1.26.10 for security fixes to the crypto/x509 package described under https://github.com/golang/go/issues?q=milestone%3AGo1.26.10+label%3ACherryPickApproved

## 7.46.1 (September 4, 2026)

BUG FIXES:
* compute: fix permadiff regression when iap is omitted from `google_compute_backend_service` ([#13041](https://github.com/hashicorp/terraform-provider-google-beta/pull/13041))

## 7.46.0 (August 26, 2026)

DEPRECATIONS:
* beyondcorp: deprecated `google_beyondcorp_app_connection`, `google_beyondcorp_app_connector`, and `google_beyondcorp_app_gateway` resources and data sources. Use `google_beyondcorp_security_gateway` and `google_beyondcorp_security_gateway_application` instead. ([#12939](https://github.com/hashicorp/terraform-provider-google-beta/pull/12939))

FEATURES:
* **New Data Source:** `google_memorystore_acl_policy` ([#12947](https://github.com/hashicorp/terraform-provider-google-beta/pull/12947))
* **New Data Source:** `google_redis_cluster_acl_policy` ([#12948](https://github.com/hashicorp/terraform-provider-google-beta/pull/12948))
* **New List Resource:** `google_migration_center_assets_export_job` ([#12929](https://github.com/hashicorp/terraform-provider-google-beta/pull/12929))
* **New List Resource:** `google_migration_center_discovery_client` ([#12929](https://github.com/hashicorp/terraform-provider-google-beta/pull/12929))
* **New List Resource:** `google_migration_center_group` ([#12929](https://github.com/hashicorp/terraform-provider-google-beta/pull/12929))
* **New List Resource:** `google_migration_center_import_job` ([#12929](https://github.com/hashicorp/terraform-provider-google-beta/pull/12929))
* **New List Resource:** `google_migration_center_preference_set` ([#12929](https://github.com/hashicorp/terraform-provider-google-beta/pull/12929))
* **New List Resource:** `google_migration_center_report_config` ([#12929](https://github.com/hashicorp/terraform-provider-google-beta/pull/12929))
* **New List Resource:** `google_migration_center_source` ([#12929](https://github.com/hashicorp/terraform-provider-google-beta/pull/12929))
* **New List Resource:** `google_network_services_authz_extension` ([#12941](https://github.com/hashicorp/terraform-provider-google-beta/pull/12941))
* **New List Resource:** `google_network_services_multicast_consumer_association` ([#12941](https://github.com/hashicorp/terraform-provider-google-beta/pull/12941))
* **New List Resource:** `google_network_services_multicast_domain_activation` ([#12941](https://github.com/hashicorp/terraform-provider-google-beta/pull/12941))
* **New List Resource:** `google_network_services_multicast_domain_group` ([#12941](https://github.com/hashicorp/terraform-provider-google-beta/pull/12941))
* **New List Resource:** `google_network_services_multicast_domain` ([#12941](https://github.com/hashicorp/terraform-provider-google-beta/pull/12941))
* **New List Resource:** `google_network_services_multicast_group_consumer_activation` ([#12941](https://github.com/hashicorp/terraform-provider-google-beta/pull/12941))
* **New List Resource:** `google_network_services_multicast_group_producer_activation` ([#12941](https://github.com/hashicorp/terraform-provider-google-beta/pull/12941))
* **New List Resource:** `google_network_services_multicast_group_range_activation` ([#12941](https://github.com/hashicorp/terraform-provider-google-beta/pull/12941))
* **New List Resource:** `google_network_services_multicast_group_range` ([#12941](https://github.com/hashicorp/terraform-provider-google-beta/pull/12941))
* **New List Resource:** `google_network_services_multicast_producer_association` ([#12941](https://github.com/hashicorp/terraform-provider-google-beta/pull/12941))
* **New Resource:** `google_memorystore_acl_policy` ([#12947](https://github.com/hashicorp/terraform-provider-google-beta/pull/12947))
* **New Resource:** `google_redis_cluster_acl_policy` ([#12948](https://github.com/hashicorp/terraform-provider-google-beta/pull/12948))
* **New Resource:** `google_vertex_ai_agent_anomaly_detection_scope` ([#12916](https://github.com/hashicorp/terraform-provider-google-beta/pull/12916))

IMPROVEMENTS:
* biglake: added `serde_info` field to `google_biglake_table` resource ([#12953](https://github.com/hashicorp/terraform-provider-google-beta/pull/12953))
* ces: added `connector_toolset` and `timeout` fields to `google_ces_toolset` resource ([#12928](https://github.com/hashicorp/terraform-provider-google-beta/pull/12928))
* compute: added write-only arguments for IAP `oauth2_client_id` and `oauth2_client_secret` to `google_compute_backend_service` resource ([#12915](https://github.com/hashicorp/terraform-provider-google-beta/pull/12915))
* container: added `opportunistic_maintenance_strategy` to `host_maintenance_policy` in `node_config` for `google_container_cluster` and `google_container_node_pool` ([#12925](https://github.com/hashicorp/terraform-provider-google-beta/pull/12925))
* discoveryengine: made `google_discovery_engine_search_engine` `search_engine_config.required_subscription_tier` updatable ([#12954](https://github.com/hashicorp/terraform-provider-google-beta/pull/12954))
* securesourcemanager: added `PULL_REQUEST_COMMENT` enum to `events` field in `google_secure_source_manager_hook` ([#12932](https://github.com/hashicorp/terraform-provider-google-beta/pull/12932))
* sql: added `replication_lag_max_seconds` to `google_sql_database_instance` ([#12919](https://github.com/hashicorp/terraform-provider-google-beta/pull/12919))
* vertexai: added `disable_natural_language_memories` and `generate_memories_examples` fields to `google_vertex_ai_reasoning_engine` ([#12949](https://github.com/hashicorp/terraform-provider-google-beta/pull/12949))

BUG FIXES:
* compute: fixed truncation of results at 500 images in `google_compute_images` data source ([#12934](https://github.com/hashicorp/terraform-provider-google-beta/pull/12934))
* container: fixed a permadiff on `enable_private_endpoint` and `master_global_access_config.enabled` in `google_container_cluster` when `control_plane_endpoints_config.ip_endpoints_config.enabled` is set to `false` ([#12944](https://github.com/hashicorp/terraform-provider-google-beta/pull/12944))
* sql: fixed `google_sql_user` returning `Missing Resource Identity After Read` when the parent Cloud SQL instance is stopped ([#12940](https://github.com/hashicorp/terraform-provider-google-beta/pull/12940))


## 7.45.0 (August 18, 2026)

FEATURES:
* **New Data Source:** `google_iam_workload_identity_pool_openid_config` ([#12903](https://github.com/hashicorp/terraform-provider-google-beta/pull/12903))
* **New Resource:** `google_agentic_applications_analyst_agent_persona` ([#12895](https://github.com/hashicorp/terraform-provider-google-beta/pull/12895))
* **New Resource:** `google_firestore_change_stream` ([#12908](https://github.com/hashicorp/terraform-provider-google-beta/pull/12908))

IMPROVEMENTS:
* accesscontextmanager: added `dry_run_access_levels` and `principal` fields to `google_access_context_manager_gcp_user_access_binding` resource ([#12885](https://github.com/hashicorp/terraform-provider-google-beta/pull/12885))
* accesscontextmanager: updated `group_key` to be optional and conflict with `principal` on `google_access_context_manager_gcp_user_access_binding` resource ([#12885](https://github.com/hashicorp/terraform-provider-google-beta/pull/12885))
* bigqueryreservation: added `labels` field to `google_bigquery_reservation` resource ([#12852](https://github.com/hashicorp/terraform-provider-google-beta/pull/12852))
* certificatemanager: enabled write-only support for `pem_private_key` on `google_certificate_manager_certificate` resource ([#12905](https://github.com/hashicorp/terraform-provider-google-beta/pull/12905))
* ces: added `snippets_config` field to `data_store_tool.modality_configs` and `service_directory_config` field to `python_function` in `google_ces_tool` ([#12878](https://github.com/hashicorp/terraform-provider-google-beta/pull/12878))
* chronicle: added `schedule_customizations` field to `google_chronicle_rule_deployment` resource ([#12856](https://github.com/hashicorp/terraform-provider-google-beta/pull/12856))
* cloudrunv2: added `templates.sandboxes` field to `google_cloud_run_v2_service` resource ([#12868](https://github.com/hashicorp/terraform-provider-google-beta/pull/12868))
* colab: added `custom_environment_spec.shielded_instance_config` and `workbench_runtime.vm_image` fields to `google_colab_notebook_execution` resource ([#12892](https://github.com/hashicorp/terraform-provider-google-beta/pull/12892))
* securityscanner: added `static_ip_scan` field to `google_security_scanner_scan_config` resource ([#12899](https://github.com/hashicorp/terraform-provider-google-beta/pull/12899))
* sql: added `cpu_utilization_threshold_percent`, `memory_usage_threshold_percent`, `history_list_length_threshold_count`, `semaphore_wait_threshold_count`, `transaction_lock_wait_threshold_count`, `transaction_kill_threshold_seconds`, `transaction_kill_type`, and `transaction_kill_excluded_user_hosts` fields to `settings.performance_capture_config` in `google_sql_database_instance` resource ([#12910](https://github.com/hashicorp/terraform-provider-google-beta/pull/12910))
* vertexai: added `spec.container_spec.port` field to `google_vertex_ai_reasoning_engine` resource ([#12881](https://github.com/hashicorp/terraform-provider-google-beta/pull/12881))
* workbench: added `compute_instance_id` field to `google_workbench_instance` resource ([#12891](https://github.com/hashicorp/terraform-provider-google-beta/pull/12891))

BUG FIXES:
* alloydb: fixed an issue where `updateMask` URL parameter was dropped during cluster updates (e.g. `database_version` upgrade) due to variable shadowing ([#12909](https://github.com/hashicorp/terraform-provider-google-beta/pull/12909))
* appengine: fixed permadiff in `google_app_engine_standard_app_version` ([#12863](https://github.com/hashicorp/terraform-provider-google-beta/pull/12863))
* apphub: fixed panic in `google_apphub_service` and `google_apphub_workload` resources when parsing `extended_metadata` containing MCP server definitions ([#12883](https://github.com/hashicorp/terraform-provider-google-beta/pull/12883))
* bigquery: fixed an issue where updating `google_bigquery_dataset` overwrote fine-grained IAM permissions ([#12893](https://github.com/hashicorp/terraform-provider-google-beta/pull/12893))
* cloudsecuritycompliance: fixed state drift on `supported_enforcement_modes` in `google_cloud_security_compliance_framework` resource ([#12851](https://github.com/hashicorp/terraform-provider-google-beta/pull/12851))
* colab: fixed drift detection on `direct_notebook_source.content` field in `google_colab_notebook_execution` resource ([#12892](https://github.com/hashicorp/terraform-provider-google-beta/pull/12892))
* compute: fixed a panic in `google_compute_shared_vpc_service_project` during `terraform plan`/`refresh` when the shared VPC link had been removed outside of Terraform ([#12869](https://github.com/hashicorp/terraform-provider-google-beta/pull/12869))
* compute: fixed permadiff on `adaptive_protection_config.layer_7_ddos_defense_config.enable` in `google_compute_security_policy` when field is not set in config ([#12870](https://github.com/hashicorp/terraform-provider-google-beta/pull/12870))
* compute: fixed permadiffs for `google_compute_disk` on Fedora CoreOS images ([#12853](https://github.com/hashicorp/terraform-provider-google-beta/pull/12853))
* networksecurity: fixed `google_network_security_gateway_security_policy` to force replacement when `name` or `location` is modified ([#12894](https://github.com/hashicorp/terraform-provider-google-beta/pull/12894))
* networkservices: fixed `google_network_services_gateway` to force replacement when `name` or `location` is modified ([#12902](https://github.com/hashicorp/terraform-provider-google-beta/pull/12902))

## 7.44.0 (August 12, 2026)

FEATURES:
* **New List Resource:** `google_bigquery_dataset_iam_member` ([#12809](https://github.com/hashicorp/terraform-provider-google-beta/pull/12809))
* **New List Resource:** `google_bigquery_table` ([#12807](https://github.com/hashicorp/terraform-provider-google-beta/pull/12807))
* **New Resource:** `google_vertex_ai_evaluation_metric` ([#12840](https://github.com/hashicorp/terraform-provider-google-beta/pull/12840))

IMPROVEMENTS:
* bigquery: added resource identity support to `google_bigquery_table` ([#12806](https://github.com/hashicorp/terraform-provider-google-beta/pull/12806))
* cloudrunv2: added `cpu_utilization` and `concurrency_utilization` fields to `google_cloud_run_v2_service` revision scaling configuration ([#12814](https://github.com/hashicorp/terraform-provider-google-beta/pull/12814))
* container: added `high_scale_checkpointing_config` block to `addons_config` in `google_container_cluster` ([#12844](https://github.com/hashicorp/terraform-provider-google-beta/pull/12844))
* dataproc: added `attached_disk_config` to `disk_config` to support attached disks in the `google_dataproc_cluster` resource ([#12816](https://github.com/hashicorp/terraform-provider-google-beta/pull/12816))
* memorystore: documented `TOKEN_AUTH` as a `google-beta`-only value for `authorization_mode` field on `google_memorystore_instance` resource ([#12810](https://github.com/hashicorp/terraform-provider-google-beta/pull/12810))
* networkservices: added `allow_global_access` field to `google_network_services_gateway` ([#12815](https://github.com/hashicorp/terraform-provider-google-beta/pull/12815))
* oracledatabase: added `identity_connector` to `google_oracle_database_exadb_vm_cluster` for CMEK support ([#12817](https://github.com/hashicorp/terraform-provider-google-beta/pull/12817))
* securesourcemanager: added `service_account` and `scan_config` fields to `google_secure_source_manager_repository` ([#12837](https://github.com/hashicorp/terraform-provider-google-beta/pull/12837))
* sql: added `password_secret_version` and `user` fields into `google_sql_provision_script` resource ([#12821](https://github.com/hashicorp/terraform-provider-google-beta/pull/12821))
* sql: removed `ForceNew` config from `disk_type` field in `google_sql_database_instance`, allowing it to change in future without requiring the database instance to be destroyed and recreated ([#12818](https://github.com/hashicorp/terraform-provider-google-beta/pull/12818))
* vertexai: added `generation_trigger_config` block to `google_vertex_ai_reasoning_engine` ([#12831](https://github.com/hashicorp/terraform-provider-google-beta/pull/12831))
* workstations: added `idle_action` field to `google_workstations_workstation_config` ([#12804](https://github.com/hashicorp/terraform-provider-google-beta/pull/12804))

BUG FIXES:
* backupdr: fixed issue where `google_backup_dr_restore_workload` dropped `resource_manager_tags` during restore requests, causing tags shown in plan to not be applied to restored resources ([#12812](https://github.com/hashicorp/terraform-provider-google-beta/pull/12812))
* biglakeiceberg: fixed a permadiff in `google_biglake_iceberg_table` by suppressing diffs on `location` when the API-returned path contains a suffix folder ([#12808](https://github.com/hashicorp/terraform-provider-google-beta/pull/12808))
* biglakeiceberg: fixed permadiff on `write.parquet.compression-codec` in `google_biglake_iceberg_table` ([#12836](https://github.com/hashicorp/terraform-provider-google-beta/pull/12836))
* bigquery: fixed a provider panic when reading incomplete IAM conditions on `google_bigquery_dataset_iam_member` ([#12819](https://github.com/hashicorp/terraform-provider-google-beta/pull/12819))
* cloudrunv2: fixed permadiff by setting `template.scaling.max_instance_count` to computed in `google_cloud_run_v2_service` ([#12829](https://github.com/hashicorp/terraform-provider-google-beta/pull/12829))
* cloudrunv2: fixed permadiff in `template.containers.resources.limits` block in `google_cloud_run_v2_service` resource ([#12845](https://github.com/hashicorp/terraform-provider-google-beta/pull/12845))
* cloudsecuritycompliance: fixed state drift on `supported_enforcement_modes` in `google_cloud_security_compliance_framework` resource ([#12851](https://github.com/hashicorp/terraform-provider-google-beta/pull/12851))
* container: fixed a permadiff where `ignore_node_count_changes` was not persisted in `google_container_cluster.node_pool` ([#12805](https://github.com/hashicorp/terraform-provider-google-beta/pull/12805))
* container: fixed an issue where `google_container_node_pool` and `google_container_cluster` failed to invalidate their Instance Group Manager cache when a resize occurred or when `ignore_node_count_changes` was active ([#12811](https://github.com/hashicorp/terraform-provider-google-beta/pull/12811))
* networkconnectivity: made `network` field in `google_network_connectivity_transport` optional ([#12823](https://github.com/hashicorp/terraform-provider-google-beta/pull/12823))
* recaptchaenterprise: fixed updates to `google_recaptcha_enterprise_key` so existing `challenge_settings.action_settings` entries are preserved when the action map changes ([#12848](https://github.com/hashicorp/terraform-provider-google-beta/pull/12848))
* servicenetworking: fixed `google_service_networking_connection` ignoring the configured `delete` timeout, which used the `create` timeout instead ([#12824](https://github.com/hashicorp/terraform-provider-google-beta/pull/12824))

## 7.43.0 (August 4, 2026)

NOTES:
* docs(workflows): added warning to `source_contents` field on `google_workflows_workflow` noting that it will become required in version 8.0.0 ([#12788](https://github.com/hashicorp/terraform-provider-google-beta/pull/12788))
* firestore: clarified which resource to use for which kind of index in Standard and Enterprise editions ([#12793](https://github.com/hashicorp/terraform-provider-google-beta/pull/12793))

FEATURES:
* **New List Resource:** `google_compute_instance` ([#12797](https://github.com/hashicorp/terraform-provider-google-beta/pull/12797))
* **New List Resource:** `google_discovery_engine_assistant` ([#12790](https://github.com/hashicorp/terraform-provider-google-beta/pull/12790))
* **New List Resource:** `google_discovery_engine_chat_engine` ([#12790](https://github.com/hashicorp/terraform-provider-google-beta/pull/12790))
* **New List Resource:** `google_discovery_engine_cmek_config` ([#12790](https://github.com/hashicorp/terraform-provider-google-beta/pull/12790))
* **New List Resource:** `google_discovery_engine_control` ([#12790](https://github.com/hashicorp/terraform-provider-google-beta/pull/12790))
* **New List Resource:** `google_discovery_engine_data_store` ([#12790](https://github.com/hashicorp/terraform-provider-google-beta/pull/12790))
* **New List Resource:** `google_discovery_engine_license_config` ([#12790](https://github.com/hashicorp/terraform-provider-google-beta/pull/12790))
* **New List Resource:** `google_discovery_engine_recommendation_engine` ([#12790](https://github.com/hashicorp/terraform-provider-google-beta/pull/12790))
* **New List Resource:** `google_discovery_engine_schema` ([#12790](https://github.com/hashicorp/terraform-provider-google-beta/pull/12790))
* **New List Resource:** `google_discovery_engine_search_engine` ([#12790](https://github.com/hashicorp/terraform-provider-google-beta/pull/12790))
* **New List Resource:** `google_discovery_engine_serving_config` ([#12790](https://github.com/hashicorp/terraform-provider-google-beta/pull/12790))
* **New List Resource:** `google_discovery_engine_sitemap` ([#12790](https://github.com/hashicorp/terraform-provider-google-beta/pull/12790))
* **New List Resource:** `google_discovery_engine_target_site` ([#12790](https://github.com/hashicorp/terraform-provider-google-beta/pull/12790))
* **New List Resource:** `google_discovery_engine_user_store` ([#12790](https://github.com/hashicorp/terraform-provider-google-beta/pull/12790))
* **New List Resource:** `google_project_iam_custom_role` ([#12791](https://github.com/hashicorp/terraform-provider-google-beta/pull/12791))
* **New List Resource:** `google_pubsub_subscription_iam_member` ([#12798](https://github.com/hashicorp/terraform-provider-google-beta/pull/12798))
* **New List Resource:** `google_service_account_iam_member` ([#12802](https://github.com/hashicorp/terraform-provider-google-beta/pull/12802))
* **New Resource:** `google_cloud_support_support_event_subscription` ([#12783](https://github.com/hashicorp/terraform-provider-google-beta/pull/12783))
* **New Resource:** `google_compute_region_network_policy_traffic_classification_rule` ([#12778](https://github.com/hashicorp/terraform-provider-google-beta/pull/12778))
* **New Resource:** `google_netapp_trial` ([#12777](https://github.com/hashicorp/terraform-provider-google-beta/pull/12777))

IMPROVEMENTS:
* apigee: added `service_account` field to `google_apigee_api_deployment` resource ([#12801](https://github.com/hashicorp/terraform-provider-google-beta/pull/12801))
* apihub: added `source_project_id` field to `google_apihub_plugin_instance` resource ([#12794](https://github.com/hashicorp/terraform-provider-google-beta/pull/12794))
* artifactregistry: added `no_cache` field to `google_artifact_registry_repository.remote_repository_config` ([#12779](https://github.com/hashicorp/terraform-provider-google-beta/pull/12779))
* bigquery: added `data_governance_tags_info` field to `google_bigquery_table` resource ([#12796](https://github.com/hashicorp/terraform-provider-google-beta/pull/12796))
* bigqueryanalyticshub: added `proposer` field to `google_bigquery_analytics_hub_query_template` ([#12784](https://github.com/hashicorp/terraform-provider-google-beta/pull/12784))
* bigquerydatapolicyv2: added `data_governance_tag` field to `google_bigquery_datapolicyv2_data_policy` resource ([#12766](https://github.com/hashicorp/terraform-provider-google-beta/pull/12766))
* bigqueryreservation: added `principal` field to `google_bigquery_reservation_assignment` resource ([#12781](https://github.com/hashicorp/terraform-provider-google-beta/pull/12781))
* cloudrunv: added `sandbox_launcher` field to the containers of  `google_cloud_run_service` resource ([#12786](https://github.com/hashicorp/terraform-provider-google-beta/pull/12786))
* cloudrunv2: added `sandbox_launcher` field to the containers of  `google_cloud_run_v2_service` resource ([#12786](https://github.com/hashicorp/terraform-provider-google-beta/pull/12786))
* discoveryengine: added `acl_enabled` field to `google_discovery_engine_data_store` resource ([#12768](https://github.com/hashicorp/terraform-provider-google-beta/pull/12768))
* privateca: made `config.subject_config.subject.organization` optional in `google_privateca_certificate` ([#12767](https://github.com/hashicorp/terraform-provider-google-beta/pull/12767))
* vertexai: added `traffic_config` field and live traffic split update support to `google_vertex_ai_reasoning_engine` ([#12775](https://github.com/hashicorp/terraform-provider-google-beta/pull/12775))
* workstations: added `idle_action` field to `google_workstations_workstation_config` ([#12804](https://github.com/hashicorp/terraform-provider-google-beta/pull/12804))

BUG FIXES:
* apigee: fixed `continue_on_error` argument being dropped in `google_apigee_flowhook`, aligning provider behavior with the Apigee API ([#12803](https://github.com/hashicorp/terraform-provider-google-beta/pull/12803))
* gkehub2: made field `spec.workloadidentity.scopeTenancyPool` in resource `google_gke_hub_feature` not required. ([#12774](https://github.com/hashicorp/terraform-provider-google-beta/pull/12774))

## 7.42.0 (July 28, 2026)

NOTES:
* compute: migrated `google_compute_region_instance_template` resource to use direct HTTP rather than a client library ([#12752](https://github.com/hashicorp/terraform-provider-google-beta/pull/12752))

DEPRECATIONS:
* vertexai: deprecated `google_vertex_ai_schedule`, an accidentally-added duplicate resource; use `google_colab_schedule` instead. ([#12739](https://github.com/hashicorp/terraform-provider-google-beta/pull/12739))

FEATURES:
* **New List Resource:** `google_service_account_key` ([#12750](https://github.com/hashicorp/terraform-provider-google-beta/pull/12750))
* **New Resource:** `google_agent_identity_auth_provider` ([#12762](https://github.com/hashicorp/terraform-provider-google-beta/pull/12762))
* **New Resource:** `google_apihub_runtime_project_attachment` ([#12764](https://github.com/hashicorp/terraform-provider-google-beta/pull/12764))
* **New Resource:** `google_chronicle_big_query_export` ([#12737](https://github.com/hashicorp/terraform-provider-google-beta/pull/12737))
* **New Resource:** `google_vector_search_data_object` ([#12755](https://github.com/hashicorp/terraform-provider-google-beta/pull/12755))
* **New Resource:** `google_vertex_ai_persistent_resource` ([#12756](https://github.com/hashicorp/terraform-provider-google-beta/pull/12756))

IMPROVEMENTS:
* bigquery: added `table_type` field to `google_bigquery_routine` resource ([#12761](https://github.com/hashicorp/terraform-provider-google-beta/pull/12761))
* colab: added `catch_up`, `create_pipeline_job_request`, `create_time`, `last_pause_time`, `last_resume_time`, `last_scheduled_run_response`, `max_concurrent_active_run_count`, `next_run_time`, `started_run_count`, `update_time` fields, and sub-fields under `create_notebook_execution_job_request.notebook_execution_job` (`create_time`, `custom_environment_spec`, `encryption_spec`, `job_state`, `kernel_name`, `labels`, `name`, `schedule_resource_name`, `workbench_runtime`) and under `create_notebook_execution_job_request` (`notebook_execution_job_id`, `parent`) to `google_colab_schedule` resource ([#12739](https://github.com/hashicorp/terraform-provider-google-beta/pull/12739))
* composer: added `config.node_config.traffic_routing_config` field to `google_composer_environment` resource ([#12733](https://github.com/hashicorp/terraform-provider-google-beta/pull/12733))
* compute: added `effective_location` field to `google_compute_interconnect` resource ([#12742](https://github.com/hashicorp/terraform-provider-google-beta/pull/12742))
* compute: added `request_headers` and `response_headers` fields to `log_config` on `google_compute_backend_service` and `google_compute_region_backend_service` resources ([#12746](https://github.com/hashicorp/terraform-provider-google-beta/pull/12746))
* compute: added identity support to `google_compute_instance`, allowing resource import using an `identity` block ([#12754](https://github.com/hashicorp/terraform-provider-google-beta/pull/12754))
* compute: changed `location` field to mutable for `google_compute_interconnect` resource ([#12742](https://github.com/hashicorp/terraform-provider-google-beta/pull/12742))
* container: added `addons_config.node_readiness_config` field to `google_container_cluster` resource ([#12743](https://github.com/hashicorp/terraform-provider-google-beta/pull/12743))
* container: added `rollback_safe_upgrade`, `desired_emulated_version`, and `emulated_version` fields to `google_container_cluster` resource ([#12757](https://github.com/hashicorp/terraform-provider-google-beta/pull/12757))
* container: increased default timeout to 2 hours for `google_container_node_pool` resource ([#12727](https://github.com/hashicorp/terraform-provider-google-beta/pull/12727))
* dataproc: added `confidential_instance_type` field to `google_dataproc_cluster` resource ([#12725](https://github.com/hashicorp/terraform-provider-google-beta/pull/12725))
* gkehub: added `min_control_plane_version`, `min_node_version`, `target_control_plane_version`, `target_node_version`, and `operational_state` fields to `google_gke_hub_rollout_sequence` resource ([#12749](https://github.com/hashicorp/terraform-provider-google-beta/pull/12749))
* hypercomputecluster: increased default timeouts for `google_hypercomputecluster_cluster` to 120 minutes ([#12763](https://github.com/hashicorp/terraform-provider-google-beta/pull/12763))
* modelarmor: added `template_metadata.filter_version_selector` field to `google_model_armor_template` resource ([#12736](https://github.com/hashicorp/terraform-provider-google-beta/pull/12736))
* sql: added identity support to `google_sql_user` for `terraform query` support ([#12748](https://github.com/hashicorp/terraform-provider-google-beta/pull/12748))

BUG FIXES:
* bigtable: fixed an issue where `bigtable_custom_endpoint` and `universe_domain` were ignored when creating Bigtable resources ([#12738](https://github.com/hashicorp/terraform-provider-google-beta/pull/12738))
* compute: fixed an issue where diffs in `google_compute_security_policy` were not detected ([#12745](https://github.com/hashicorp/terraform-provider-google-beta/pull/12745))
* gkehub: fixed `rollout_creation_scope` and `upgrade_types` fields in `google_gke_hub_rollout_sequence` resource ([#12749](https://github.com/hashicorp/terraform-provider-google-beta/pull/12749))
* osconfig: added client-side validation to ensure `resource_hierarchy_selector` and `location_selector` are not set at the same time in `google_os_config_v2_policy_orchestrator`, `google_os_config_v2_policy_orchestrator_for_folder`, and `google_os_config_v2_policy_orchestrator_for_organization` ([#12740](https://github.com/hashicorp/terraform-provider-google-beta/pull/12740))
* secretmanager: fixed an issue where `google_secret_manager_secret_version` would fail at apply time if neither `secret_data` nor `secret_data_wo` was set ([#12744](https://github.com/hashicorp/terraform-provider-google-beta/pull/12744))
* sql: fixed issue where updates to `settings.ip_configuration.psc_config.allowed_consumer_projects` in `google_sql_database_instance` were silently ignored on in-place updates ([#12759](https://github.com/hashicorp/terraform-provider-google-beta/pull/12759))
* vertexai: fixed `google_vertex_ai_endpoint_with_model_garden_deployment` destroying and recreating the endpoint when `min_replica_count`, `max_replica_count`, `required_replica_count`, or `autoscaling_metric_specs` changed ([#12735](https://github.com/hashicorp/terraform-provider-google-beta/pull/12735))

## 7.41.0 (July 17, 2026)

FEATURES:
* **New List Resource:** `google_compute_backend_bucket_signed_url_key` ([#12713](https://github.com/hashicorp/terraform-provider-google-beta/pull/12713))
* **New List Resource:** `google_compute_backend_service_signed_url_key` ([#12713](https://github.com/hashicorp/terraform-provider-google-beta/pull/12713))
* **New List Resource:** `google_compute_network_firewall_policy` ([#12713](https://github.com/hashicorp/terraform-provider-google-beta/pull/12713))
* **New List Resource:** `google_compute_network_firewall_policy_association` ([#12713](https://github.com/hashicorp/terraform-provider-google-beta/pull/12713))
* **New List Resource:** `google_compute_network_firewall_policy_packet_mirroring_rule` ([#12713](https://github.com/hashicorp/terraform-provider-google-beta/pull/12713))
* **New List Resource:** `google_compute_preview_feature` ([#12713](https://github.com/hashicorp/terraform-provider-google-beta/pull/12713))
* **New List Resource:** `google_compute_public_advertised_prefix` ([#12713](https://github.com/hashicorp/terraform-provider-google-beta/pull/12713))
* **New List Resource:** `google_compute_region_backend_bucket` ([#12713](https://github.com/hashicorp/terraform-provider-google-beta/pull/12713))
* **New List Resource:** `google_compute_region_network_firewall_policy` ([#12713](https://github.com/hashicorp/terraform-provider-google-beta/pull/12713))
* **New Resource:** `google_biglake_hive_table` ([#12721](https://github.com/hashicorp/terraform-provider-google-beta/pull/12721))
* **New Resource:** `google_compute_global_vm_extension_policy` ([#12673](https://github.com/hashicorp/terraform-provider-google-beta/pull/12673))

IMPROVEMENTS:
* accesscontextmanager: added `allowed_service_patterns` and `service_patterns_enforcement_scopes` fields to `google_access_context_manager_service_perimeter` to support VPC Service Controls for non-Google Cloud APIs ([#12705](https://github.com/hashicorp/terraform-provider-google-beta/pull/12705))
* accesscontextmanager: added `psc_endpoint` block to `sources` in `ingress_from` and `egress_from` under `google_access_context_manager_service_perimeter` resource and its variants ([#12678](https://github.com/hashicorp/terraform-provider-google-beta/pull/12678))
* backupdr: added `backup_blocked_by_vault_access_restriction` attribute to `google_backup_dr_data_source` data source ([#12717](https://github.com/hashicorp/terraform-provider-google-beta/pull/12717))
* backupdr: added `force_update_access_restriction` field to `google_backup_dr_backup_vault` resource ([#12717](https://github.com/hashicorp/terraform-provider-google-beta/pull/12717))
* backupdr: added `max_custom_on_demand_retention_days` to the Cloud SQL backup plan example for `google_backup_dr_backup_plan` resource ([#12700](https://github.com/hashicorp/terraform-provider-google-beta/pull/12700))
* backupdr: added update support for `access_restriction` to `google_backup_dr_backup_vault` resource ([#12717](https://github.com/hashicorp/terraform-provider-google-beta/pull/12717))
* certificatemanager: added in-place update support for the `self_managed` certificate data (`pem_certificate` and `pem_private_key`) on `google_certificate_manager_certificate` resource ([#12696](https://github.com/hashicorp/terraform-provider-google-beta/pull/12696))
* cloudrunv2: added `tags` field to `google_cloud_run_v2_service` and `google_cloud_run_v2_job` resources ([#12690](https://github.com/hashicorp/terraform-provider-google-beta/pull/12690))
* compute: added 3500GB and 7000GB SSD partition size support to `google_compute_instance_template` resource ([#12708](https://github.com/hashicorp/terraform-provider-google-beta/pull/12708))
* compute: added `FLEX_START` and `RESERVATION_BOUND` support to `google_compute_instance`, `google_compute_instance_template`, and `google_compute_region_instance_template` resources ([#12720](https://github.com/hashicorp/terraform-provider-google-beta/pull/12720))
* container: added support for updating `node_image_config` and `image_type` fields at the same time on `google_container_cluster` and `google_container_node_pool` resources ([#12674](https://github.com/hashicorp/terraform-provider-google-beta/pull/12674))
* container: increased default timeout to 2 hours for `google_container_node_pool` resource ([#12727](https://github.com/hashicorp/terraform-provider-google-beta/pull/12727))
* dataform: added `git_repository_link` field to `google_dataform_repository` resource ([#12683](https://github.com/hashicorp/terraform-provider-google-beta/pull/12683))
* dataproc: added `confidential_instance_type` field to `google_dataproc_cluster` resource ([#12725](https://github.com/hashicorp/terraform-provider-google-beta/pull/12725))
* dataproc: added `instance_selection.disk_config` field to `google_dataproc_cluster` resource ([#12698](https://github.com/hashicorp/terraform-provider-google-beta/pull/12698))
* discoveryengine: added `enable_llm_layout_parsing` and `enable_get_processed_document` fields to `google_discovery_engine_data_store` resource ([#12675](https://github.com/hashicorp/terraform-provider-google-beta/pull/12675))
* sql: added `include_replicas_for_major_version_upgrade` field to `google_sql_database_instance` resource ([#12702](https://github.com/hashicorp/terraform-provider-google-beta/pull/12702))
* sql: added `instance_auto_dns_status` and `write_endpoint_auto_dns_status` output fields to the `psc_auto_connections` block in `google_sql_database_instance` resource ([#12692](https://github.com/hashicorp/terraform-provider-google-beta/pull/12692))
* sql: added `switch_transaction_logs_to_cloud_storage_enabled` field to `google_sql_database_instance` resource ([#12686](https://github.com/hashicorp/terraform-provider-google-beta/pull/12686))
* vertexai: added `traffic_config`, `url`, `agent_server_mode`, `agent_gateway_config`, `keep_alive_probe`, `dedicated_ingress_endpoint_enabled`, `agent_config_source`, `agent_card`, `example_store`, `build_spec`, and `example_store_config` fields to `google_vertex_ai_reasoning_engine` resource ([#12671](https://github.com/hashicorp/terraform-provider-google-beta/pull/12671))
* workbench: added `enable_deletion_protection` field to `google_workbench_instance` resource ([#12712](https://github.com/hashicorp/terraform-provider-google-beta/pull/12712))
* workbench: added `resource_policies` field to `google_workbench_instance` resource ([#12710](https://github.com/hashicorp/terraform-provider-google-beta/pull/12710))
* workbench: added support for `min_cpu_platform` in `google_workbench_instance` resource ([#12723](https://github.com/hashicorp/terraform-provider-google-beta/pull/12723))
* workstations: added `instance_metadata` field to `google_workstations_workstation_config` resource ([#12699](https://github.com/hashicorp/terraform-provider-google-beta/pull/12699))

BUG FIXES:
* compute: fixed permadiff on `existing_reservations` field of `google_compute_reservation_region_commitment` resource after upgrading ([#12709](https://github.com/hashicorp/terraform-provider-google-beta/pull/12709))
* compute: fixed permadiff on `keepalive_interval` in `bgp` block of `google_compute_router` resource when set to default value ([#12676](https://github.com/hashicorp/terraform-provider-google-beta/pull/12676))
* resourcemanager: fixed validation of `target_service_account` and `delegates` in `google_service_account_access_token`, `google_service_account_id_token`, and `google_service_account_jwt` data sources, and of `name` in `google_service_account_key` data source, to reject identifiers that contain path separators ([#12679](https://github.com/hashicorp/terraform-provider-google-beta/pull/12679))
* securityposture: fixed bug where `enforce` field in `google_securityposture_posture` resource was always set, causing failures for list constraints ([#12715](https://github.com/hashicorp/terraform-provider-google-beta/pull/12715))
* vmwareengine: fixed `google_vmwareengine_private_cloud` updates by sending the correct `update_mask` parameter ([#12716](https://github.com/hashicorp/terraform-provider-google-beta/pull/12716))



## 7.40.0 (July 14, 2026)
DEPRECATIONS:
* storage: the `admit-on-second-miss` value for `google_storage_anywhere_cache.admission_policy` is deprecated and will be removed in a future major release. The backend will ignore this attribute and treat it as `admit-on-first-miss`. ([#12632](https://github.com/hashicorp/terraform-provider-google-beta/pull/12632))

NOTES:
* compute: migrated `google_compute_instance` code related to advanced machine features to use direct HTTP rather than a client library ([#12610](https://github.com/hashicorp/terraform-provider-google-beta/pull/12610))

FEATURES:
* **New Data Source:** `google_data_catalog_taxonomy` ([#12655](https://github.com/hashicorp/terraform-provider-google-beta/pull/12655))
* **New Data Source:** `google_oracle_database_exascale_db_storage_vault` ([#12664](https://github.com/hashicorp/terraform-provider-google-beta/pull/12664))
* **New List Resource:** 'google_project' ([#12602](https://github.com/hashicorp/terraform-provider-google-beta/pull/12602))
* **New List Resource:** `google_compute_instant_snapshot` ([#12661](https://github.com/hashicorp/terraform-provider-google-beta/pull/12661))
* **New List Resource:** `google_compute_region_instant_snapshot` ([#12661](https://github.com/hashicorp/terraform-provider-google-beta/pull/12661))
* **New List Resource:** `google_compute_region_target_http_proxy` ([#12661](https://github.com/hashicorp/terraform-provider-google-beta/pull/12661))
* **New List Resource:** `google_compute_region_target_tcp_proxy` ([#12661](https://github.com/hashicorp/terraform-provider-google-beta/pull/12661))
* **New List Resource:** `google_compute_region_url_map` ([#12661](https://github.com/hashicorp/terraform-provider-google-beta/pull/12661))
* **New List Resource:** `google_compute_rollout_plan` ([#12661](https://github.com/hashicorp/terraform-provider-google-beta/pull/12661))
* **New List Resource:** `google_compute_target_grpc_proxy` ([#12661](https://github.com/hashicorp/terraform-provider-google-beta/pull/12661))
* **New List Resource:** `google_compute_target_http_proxy` ([#12661](https://github.com/hashicorp/terraform-provider-google-beta/pull/12661))
* **New List Resource:** `google_compute_target_ssl_proxy` ([#12661](https://github.com/hashicorp/terraform-provider-google-beta/pull/12661))
* **New List Resource:** `google_compute_target_tcp_proxy` ([#12661](https://github.com/hashicorp/terraform-provider-google-beta/pull/12661))
* **New List Resource:** `google_dns_managed_zone` ([#12662](https://github.com/hashicorp/terraform-provider-google-beta/pull/12662))
* **New List Resource:** `google_oracle_database_exascale_db_storage_vaults` ([#12664](https://github.com/hashicorp/terraform-provider-google-beta/pull/12664))
* **New Resource:** `google_bigquery_analytics_hub_query_template` ([#12669](https://github.com/hashicorp/terraform-provider-google-beta/pull/12669))
* **New Resource:** `google_chronicle_environment_group` ([#12627](https://github.com/hashicorp/terraform-provider-google-beta/pull/12627))
* **New Resource:** `google_chronicle_soar_domain` ([#12635](https://github.com/hashicorp/terraform-provider-google-beta/pull/12635))
* **New Resource:** `google_iap_agent_registry_agent_iam_binding` ([#12651](https://github.com/hashicorp/terraform-provider-google-beta/pull/12651))
* **New Resource:** `google_iap_agent_registry_agent_iam_member` ([#12651](https://github.com/hashicorp/terraform-provider-google-beta/pull/12651))
* **New Resource:** `google_iap_agent_registry_agent_iam_policy` ([#12651](https://github.com/hashicorp/terraform-provider-google-beta/pull/12651))
* **New Resource:** `google_iap_agent_registry_endpoint_iam_binding` ([#12651](https://github.com/hashicorp/terraform-provider-google-beta/pull/12651))
* **New Resource:** `google_iap_agent_registry_endpoint_iam_member` ([#12651](https://github.com/hashicorp/terraform-provider-google-beta/pull/12651))
* **New Resource:** `google_iap_agent_registry_endpoint_iam_policy` ([#12651](https://github.com/hashicorp/terraform-provider-google-beta/pull/12651))
* **New Resource:** `google_iap_agent_registry_mcp_server_iam_binding` ([#12651](https://github.com/hashicorp/terraform-provider-google-beta/pull/12651))
* **New Resource:** `google_iap_agent_registry_mcp_server_iam_member` ([#12651](https://github.com/hashicorp/terraform-provider-google-beta/pull/12651))
* **New Resource:** `google_iap_agent_registry_mcp_server_iam_policy` ([#12651](https://github.com/hashicorp/terraform-provider-google-beta/pull/12651))
* **New Resource:** `google_tags_tag_binding_collection` ([#12619](https://github.com/hashicorp/terraform-provider-google-beta/pull/12619))
* **New Resource:** `google_vector_search_index` ([#12656](https://github.com/hashicorp/terraform-provider-google-beta/pull/12656))
* **New Resource:** `google_vertex_ai_semantic_governance_policy_engine` ([#12624](https://github.com/hashicorp/terraform-provider-google-beta/pull/12624))
* **New Resource:** `google_chronicle_data_export` ([#12657](https://github.com/hashicorp/terraform-provider-google-beta/pull/12657))

IMPROVEMENTS:
* agentregistry: added `name` field to `google_agent_registry_binding` resource ([#12629](https://github.com/hashicorp/terraform-provider-google-beta/pull/12629))
* agentregistry: added `name` field to `google_agent_registry_service` resource ([#12629](https://github.com/hashicorp/terraform-provider-google-beta/pull/12629))
* appengine: added `app_engine_bundled_services` field to `google_app_engine_standard_app_version` resource ([#12634](https://github.com/hashicorp/terraform-provider-google-beta/pull/12634))
* biglakeiceberg: add support for `CATALOG_TYPE_FEDERATED` with `federated_catalog_options` to `google_biglake_iceberg_catalog` ([#12659](https://github.com/hashicorp/terraform-provider-google-beta/pull/12659))
* compute: added `target_type` and `target_forwarding_rules` to `google_compute_region_network_firewall_policy_with_rules` resource ([#12607](https://github.com/hashicorp/terraform-provider-google-beta/pull/12607))
* compute: added `workload_identity_config` fields to `resource_compute_instance` and `resource_compute_instance_template` resources ([#12670](https://github.com/hashicorp/terraform-provider-google-beta/pull/12670))
* container: added `ANY_RESERVATION_THEN_FAIL` option to `consume_reservation_type` field in `google_container_cluster` and `google_container_node_pool` resources ([#12606](https://github.com/hashicorp/terraform-provider-google-beta/pull/12606))
* container: added `custom_node_init` configuration block to `node_config` (supporting Cloud Storage and Secret Manager) for both `google_container_cluster` and `google_container_node_pool` resources. ([#12666](https://github.com/hashicorp/terraform-provider-google-beta/pull/12666))
* container: added `maintenance_policy` field to `google_container_node_pool` resource ([#12638](https://github.com/hashicorp/terraform-provider-google-beta/pull/12638))
* container: added `recurring_maintenance_window` field to `google_container_cluster` resource ([#12647](https://github.com/hashicorp/terraform-provider-google-beta/pull/12647))
* dataplex: added `catalog_publishing_enabled` field to `google_dataplex_datascan` resource ([#12652](https://github.com/hashicorp/terraform-provider-google-beta/pull/12652))
* dlp: added `inspect_config.min_likelihood_per_info_type` to `google_data_loss_prevention_inspect_template` ([#12654](https://github.com/hashicorp/terraform-provider-google-beta/pull/12654))
* firestore: added `skip_wait` field to `google_firestore_field` resource, skipping the wait for index creation ([#12642](https://github.com/hashicorp/terraform-provider-google-beta/pull/12642))
* oracledatabase: added support for configuring Exascale-based VM clusters on top of dedicated storage vaults via the `exascale_db_storage_vault` parameter and `storage_management_type` to determine if VM Cluster is ASM or EXASCALE in `google_oracle_database_cloud_vm_cluster` ([#12626](https://github.com/hashicorp/terraform-provider-google-beta/pull/12626))
* oracledatabase: added `exadata_infrastructure` field to `google_oracle_database_exascale_db_storage_vault` resource ([#12616](https://github.com/hashicorp/terraform-provider-google-beta/pull/12616))
* oracledatabase: added `exascale_db_storage_vault` and `storage_management_type` fields to `google_oracle_database_cloud_vm_clusters` data source ([#12664](https://github.com/hashicorp/terraform-provider-google-beta/pull/12664))
* sql: added `enforce_new_sql_network_architecture` field to `google_sql_database_instance` resource ([#12653](https://github.com/hashicorp/terraform-provider-google-beta/pull/12653))
* sql: added `psc_auto_connection_policy_enabled` field and output-only `service_connection_policy` and `service_connection_policy_creation_result` fields to `google_sql_database_instance` resource ([#12645](https://github.com/hashicorp/terraform-provider-google-beta/pull/12645))
* storagetransfer: added `private_network_service` to resource `google_storage_transfer_job` ([#12617](https://github.com/hashicorp/terraform-provider-google-beta/pull/12617))
* vertex_ai: added `customization_configs` and `structured_memory_configs` fields to `google_vertex_ai_reasoning_engine` resource ([#12618](https://github.com/hashicorp/terraform-provider-google-beta/pull/12618))
* vertexai: added `traffic_config`, `url`, `agent_server_mode`, `agent_gateway_config`, `keep_alive_probe`, `dedicated_ingress_endpoint_enabled`, `agent_config_source`, `agent_card`, `example_store`, `build_spec`, `traffic_config`, and `example_store_config` fields to `google_vertex_ai_reasoning_engine` ([#12671](https://github.com/hashicorp/terraform-provider-google-beta/pull/12671))

BUG FIXES:
* bigtable: fixed a bug where `row_affinity` updates did not persist on `google_bigtable_app_profile` ([#12636](https://github.com/hashicorp/terraform-provider-google-beta/pull/12636))
* bug: fixed labels diff in `google_project` ([#12649](https://github.com/hashicorp/terraform-provider-google-beta/pull/12649))
* chronicle: added retries for "409: sync mutate calls cannot be queued" error for `google_cloud_scheduler_job` ([#12637](https://github.com/hashicorp/terraform-provider-google-beta/pull/12637))
* cloudscheudler: added retries for the Cloud Scheduler "409: sync mutate calls cannot be queued" error on Scheduler's Job resource mutations. ([#12613](https://github.com/hashicorp/terraform-provider-google-beta/pull/12613))
* compute: fixed an issue where `preview = false` updates for `google_compute_organization_security_policy_rule` were omitted from API requests. ([#12643](https://github.com/hashicorp/terraform-provider-google-beta/pull/12643))
* compute: fixed bug where it wasn't possible to disable `enable_proxy_protocol` on `google_compute_service_attachment` resource ([#12668](https://github.com/hashicorp/terraform-provider-google-beta/pull/12668))
* datastream: fixed a bug in update functionality in `google_datastream_connection_profile` `mongodb_profile.additional_options` ([#12660](https://github.com/hashicorp/terraform-provider-google-beta/pull/12660))
* eventarc: fixed a type mismatch when an `google_eventarc_trigger` resource returns non-empty `conditions`. ([#12646](https://github.com/hashicorp/terraform-provider-google-beta/pull/12646))
* filestore: aligned `google_filestore_instance` resource timeouts with the Filestore service instance operations TTLs ([#12630](https://github.com/hashicorp/terraform-provider-google-beta/pull/12630))
* gemini: fixed truncated timeouts in `google_gemini_code_tools_setting`, `google_gemini_data_sharing_with_google_setting_binding`, `google_gemini_gemini_gcp_enablement_setting_binding`, and `google_gemini_release_channel_setting_binding` ([#12640](https://github.com/hashicorp/terraform-provider-google-beta/pull/12640))
* hypercomputecluster: fixed 20-minute timeout limit during `google_hypercomputecluster_cluster` resource creation ([#12621](https://github.com/hashicorp/terraform-provider-google-beta/pull/12621))
* logging: fixed an issue where errors on update would not be propagated in `google_logging_project_bucket_config` ([#12603](https://github.com/hashicorp/terraform-provider-google-beta/pull/12603))
* observability: fixed unintentionally long timeouts in `google_observability_folder_settings`, `google_observability_organization_settings`, and `google_observability_project_settings` ([#12640](https://github.com/hashicorp/terraform-provider-google-beta/pull/12640))
* oracledatabase: fixed early client-side timeouts and aligned default schema timeouts with backend async polling limits on `google_oracle_database_exadb_vm_cluster`, `google_oracle_database_odb_network`, `google_oracle_database_odb_subnet`, `google_oracle_database_goldengate_connection`, `google_oracle_database_goldengate_deployment`, and `google_oracle_database_goldengate_connection_assignment`. ([#12648](https://github.com/hashicorp/terraform-provider-google-beta/pull/12648))
* oracledatabase: fixed truncated timeouts in `google_oracle_database_exadb_vm_cluster`, `google_oracle_database_goldengate_connection`, and `google_oracle_database_odb_subnet` ([#12640](https://github.com/hashicorp/terraform-provider-google-beta/pull/12640))
* privilegedaccessmanager: fixed a permadiff on `google_privileged_access_manager_entitlement` for entitlements created without an approval workflow ([#12644](https://github.com/hashicorp/terraform-provider-google-beta/pull/12644))
* provider: fixed validation of `external_credentials.identity_token` to reject malformed JWTs containing empty segments ([#12663](https://github.com/hashicorp/terraform-provider-google-beta/pull/12663))

## 7.39.0 (June 30, 2026)
NOTES:
* compute: migrated `google_compute_instance_template` resource partially to use direct HTTP rather than a client library ([#12576](https://github.com/hashicorp/terraform-provider-google-beta/pull/12576))
* compute: migrated `google_compute_network_peering` resource to use direct HTTP rather than a client library ([#12587](https://github.com/hashicorp/terraform-provider-google-beta/pull/12587))

FEATURES:
* **New Data Source:** `google_agent_registry_agent` ([#12590](https://github.com/hashicorp/terraform-provider-google-beta/pull/12590))
* **New Data Source:** `google_agent_registry_endpoint` ([#12590](https://github.com/hashicorp/terraform-provider-google-beta/pull/12590))
* **New Data Source:** `google_agent_registry_mcp_server` ([#12590](https://github.com/hashicorp/terraform-provider-google-beta/pull/12590))
* **New Data Source:** `google_compute_instance_groups` ([#12554](https://github.com/hashicorp/terraform-provider-google-beta/pull/12554))
* **New Data Source:** `google_storage_control_folder_intelligence_findings_summary` ([#12585](https://github.com/hashicorp/terraform-provider-google-beta/pull/12585))
* **New Data Source:** `google_storage_control_organization_intelligence_findings_summary` ([#12585](https://github.com/hashicorp/terraform-provider-google-beta/pull/12585))
* **New Data Source:** `google_storage_control_project_intelligence_findings_summary` ([#12585](https://github.com/hashicorp/terraform-provider-google-beta/pull/12585))
* **New Resource:** `google_agent_registry_binding` ([#12590](https://github.com/hashicorp/terraform-provider-google-beta/pull/12590))
* **New Resource:** `google_agent_registry_service` ([#12590](https://github.com/hashicorp/terraform-provider-google-beta/pull/12590))
* **New Resource:** `google_artifact_registry_project_config` ([#12575](https://github.com/hashicorp/terraform-provider-google-beta/pull/12575))
* **New Resource:** `google_biglake_hive_database` ([#12580](https://github.com/hashicorp/terraform-provider-google-beta/pull/12580))
* **New Resource:** `google_compute_bulk_per_instance_config` ([#12592](https://github.com/hashicorp/terraform-provider-google-beta/pull/12592))
* **New Resource:** `google_compute_firewall_policy_iam_binding` ([#12552](https://github.com/hashicorp/terraform-provider-google-beta/pull/12552))
* **New Resource:** `google_compute_firewall_policy_iam_member` ([#12552](https://github.com/hashicorp/terraform-provider-google-beta/pull/12552))
* **New Resource:** `google_compute_firewall_policy_iam_policy` ([#12552](https://github.com/hashicorp/terraform-provider-google-beta/pull/12552))
* **New Resource:** `google_compute_network_firewall_policy_iam_binding` ([#12552](https://github.com/hashicorp/terraform-provider-google-beta/pull/12552))
* **New Resource:** `google_compute_network_firewall_policy_iam_member` ([#12552](https://github.com/hashicorp/terraform-provider-google-beta/pull/12552))
* **New Resource:** `google_compute_network_firewall_policy_iam_policy` ([#12552](https://github.com/hashicorp/terraform-provider-google-beta/pull/12552))
* **New Resource:** `google_compute_region_network_firewall_policy_iam_binding` ([#12552](https://github.com/hashicorp/terraform-provider-google-beta/pull/12552))
* **New Resource:** `google_compute_region_network_firewall_policy_iam_member` ([#12552](https://github.com/hashicorp/terraform-provider-google-beta/pull/12552))
* **New Resource:** `google_compute_region_network_firewall_policy_iam_policy` ([#12552](https://github.com/hashicorp/terraform-provider-google-beta/pull/12552))
* **New Resource:** `google_compute_region_network_policy` ([#12579](https://github.com/hashicorp/terraform-provider-google-beta/pull/12579))
* **New Resource:** `google_compute_zone_vm_extension_policy` ([#12595](https://github.com/hashicorp/terraform-provider-google-beta/pull/12595))
* **New Resource:** `google_iap_agent_registry_iam_binding` ([#12593](https://github.com/hashicorp/terraform-provider-google-beta/pull/12593))
* **New Resource:** `google_iap_agent_registry_iam_member` ([#12593](https://github.com/hashicorp/terraform-provider-google-beta/pull/12593))
* **New Resource:** `google_iap_agent_registry_iam_policy` ([#12593](https://github.com/hashicorp/terraform-provider-google-beta/pull/12593))
* **New Resource:** `google_iap_location_web_iam_binding` ([#12593](https://github.com/hashicorp/terraform-provider-google-beta/pull/12593))
* **New Resource:** `google_iap_location_web_iam_member` ([#12593](https://github.com/hashicorp/terraform-provider-google-beta/pull/12593))
* **New Resource:** `google_iap_location_web_iam_policy` ([#12593](https://github.com/hashicorp/terraform-provider-google-beta/pull/12593))
* **New Resource:** `google_oracle_database_cloud_exadata_infrastructure_exascale_config` ([#12594](https://github.com/hashicorp/terraform-provider-google-beta/pull/12594))
* **New Resource:** `google_vertex_ai_model_garden_enable_model` ([#12577](https://github.com/hashicorp/terraform-provider-google-beta/pull/12577))
* **New List Resource:** `google_bigquery_dataset` ([#12571](https://github.com/hashicorp/terraform-provider-google-beta/pull/12571))
* **New List Resource:** `google_compute_cross_site_network` ([#12584](https://github.com/hashicorp/terraform-provider-google-beta/pull/12584))
* **New List Resource:** `google_compute_external_vpn_gateway` ([#12584](https://github.com/hashicorp/terraform-provider-google-beta/pull/12584))
* **New List Resource:** `google_compute_global_network_endpoint_group` ([#12584](https://github.com/hashicorp/terraform-provider-google-beta/pull/12584))
* **New List Resource:** `google_compute_ha_vpn_gateway` ([#12584](https://github.com/hashicorp/terraform-provider-google-beta/pull/12584))
* **New List Resource:** `google_compute_interconnect_attachment_group` ([#12584](https://github.com/hashicorp/terraform-provider-google-beta/pull/12584))
* **New List Resource:** `google_compute_interconnect_group` ([#12584](https://github.com/hashicorp/terraform-provider-google-beta/pull/12584))
* **New List Resource:** `google_compute_public_delegated_prefix` ([#12584](https://github.com/hashicorp/terraform-provider-google-beta/pull/12584))
* **New List Resource:** `google_compute_region_commitment` ([#12584](https://github.com/hashicorp/terraform-provider-google-beta/pull/12584))
* **New List Resource:** `google_compute_region_network_endpoint_group` ([#12584](https://github.com/hashicorp/terraform-provider-google-beta/pull/12584))
* **New List Resource:** `google_compute_vpn_gateway` ([#12584](https://github.com/hashicorp/terraform-provider-google-beta/pull/12584))
* **New List Resource:** `google_compute_wire_group` ([#12584](https://github.com/hashicorp/terraform-provider-google-beta/pull/12584))
* **New List Resource:** `google_folder_iam_member` ([#12565](https://github.com/hashicorp/terraform-provider-google-beta/pull/12565))
* **New List Resource:** `google_kms_crypto_key_version` ([#12572](https://github.com/hashicorp/terraform-provider-google-beta/pull/12572))
* **New List Resource:** `google_project` ([#12602](https://github.com/hashicorp/terraform-provider-google-beta/pull/12602))
* **New List Resource:** `google_project_service` ([#12562](https://github.com/hashicorp/terraform-provider-google-beta/pull/12562))

IMPROVEMENTS:
* bigquery: added `external_runtime_options.container_request_concurrency` field to `google_bigquery_routine` resource ([#12591](https://github.com/hashicorp/terraform-provider-google-beta/pull/12591))
* compute: added `instance_lifecycle_policy.on_failed_health_check` field in resources `google_compute_instance_group_manager` and `google_compute_region_instance_group_manager` (ga) ([#12564](https://github.com/hashicorp/terraform-provider-google-beta/pull/12564))
* container: added new fields `shutdown_grace_period_seconds` and `shutdown_grace_period_critical_pods_seconds` to `node_kubelet_config` block. ([#12581](https://github.com/hashicorp/terraform-provider-google-beta/pull/12581))
* dataplex: added `icon` field to `google_dataplex_data_product` resource ([#12559](https://github.com/hashicorp/terraform-provider-google-beta/pull/12559))
* dataplex: added `name` field to `google_dataplex_data_product_data_asset` resource ([#12586](https://github.com/hashicorp/terraform-provider-google-beta/pull/12586))
* dlp: added `allow_limited_availability_info_types` to `google_data_loss_prevention_inspect_template` ([#12589](https://github.com/hashicorp/terraform-provider-google-beta/pull/12589))
* networkservices: added `forward_attributes` field to `google_network_services_lb_edge_extension`, `google_network_services_lb_route_extension`, and `google_network_services_lb_traffic_extension` resources ([#12578](https://github.com/hashicorp/terraform-provider-google-beta/pull/12578))

BUG FIXES:
* compute: fixed a panic in `google_compute_project_metadata` and `google_compute_project_metadata_item` when project common instance metadata items contain null/empty values. ([#12574](https://github.com/hashicorp/terraform-provider-google-beta/pull/12574))
* compute: fixed a validation error on `google_compute_instance` (`Provisioned IOPS cannot be specified with disk type pd-balanced`) that occurred during updates on instances with Hyperdisk Balanced boot disks. ([#12550](https://github.com/hashicorp/terraform-provider-google-beta/pull/12550))
* dataproc: fixed a bug where changing `policy_id` on `google_dataproc_autoscaling_policy` planned an in-place update and failed; it now correctly forces resource replacement (destroy and recreate). ([#12597](https://github.com/hashicorp/terraform-provider-google-beta/pull/12597))
* iam: fixed ephemeral `google_service_account_key` producing a 404 due to duplicate `/keys` in the URL when `fetch_key = true` and `name` is provided ([#12553](https://github.com/hashicorp/terraform-provider-google-beta/pull/12553))

## 7.38.0 (June 23, 2026)
NOTES:
* bigquery: migrated `google_bigquery_table` resource to use direct HTTP rather than a client library ([#12520](https://github.com/hashicorp/terraform-provider-google-beta/pull/12520))
* compute: migrate `resource_compute_instance_template_test.go.tmpl` resource to use direct HTTP rather than a client library ([#12492](https://github.com/hashicorp/terraform-provider-google-beta/pull/12492))
* compute: migrated `google_compute_resource_compute_instance`  to use direct HTTP rather than a client library ([#12530](https://github.com/hashicorp/terraform-provider-google-beta/pull/12530))
* compute: migrated parts of `google_compute_instance` and shared instance functions to use direct HTTP ([#12483](https://github.com/hashicorp/terraform-provider-google-beta/pull/12483))
* migrated compute_instance_helpers interfaces ([#12506](https://github.com/hashicorp/terraform-provider-google-beta/pull/12506))

FEATURES:
* **New Data Source:** `google_storage_control_project_intelligence_finding_revision` ([#12523](https://github.com/hashicorp/terraform-provider-google-beta/pull/12523))
* **New Data Source:** `google_storage_control_project_intelligence_finding_revisions` ([#12523](https://github.com/hashicorp/terraform-provider-google-beta/pull/12523))
* **New Resource:** `google_biglake_hive_catalog` ([#12509](https://github.com/hashicorp/terraform-provider-google-beta/pull/12509))
* **New Resource:** `google_chronicle_findings_refinement_deployment` ([#12532](https://github.com/hashicorp/terraform-provider-google-beta/pull/12532))
* **New Resource:** `google_chronicle_parser_extension` ([#12517](https://github.com/hashicorp/terraform-provider-google-beta/pull/12517))
* **New Resource:** `google_dataplex_metadata_feed` ([#12537](https://github.com/hashicorp/terraform-provider-google-beta/pull/12537))
* **New Resource:** `google_vertex_ai_schedule` ([#12510](https://github.com/hashicorp/terraform-provider-google-beta/pull/12510))
* **New Resource:** `google_vertex_ai_tensorboard_run` ([#12524](https://github.com/hashicorp/terraform-provider-google-beta/pull/12524))
* **New List Resource:** `google_compute_address` ([#12528](https://github.com/hashicorp/terraform-provider-google-beta/pull/12528))
* **New List Resource:** `google_compute_cross_site_network` ([#12497](https://github.com/hashicorp/terraform-provider-google-beta/pull/12497))
* **New List Resource:** `google_compute_https_health_check` ([#12528](https://github.com/hashicorp/terraform-provider-google-beta/pull/12528))
* **New List Resource:** `google_compute_machine_image` ([#12528](https://github.com/hashicorp/terraform-provider-google-beta/pull/12528))
* **New List Resource:** `google_compute_node_template` ([#12528](https://github.com/hashicorp/terraform-provider-google-beta/pull/12528))
* **New List Resource:** `google_compute_packet_mirroring` ([#12528](https://github.com/hashicorp/terraform-provider-google-beta/pull/12528))
* **New List Resource:** `google_compute_region_autoscaler` ([#12528](https://github.com/hashicorp/terraform-provider-google-beta/pull/12528))
* **New List Resource:** `google_compute_region_composite_health_check` ([#12528](https://github.com/hashicorp/terraform-provider-google-beta/pull/12528))
* **New List Resource:** `google_compute_region_health_aggregation_policy` ([#12528](https://github.com/hashicorp/terraform-provider-google-beta/pull/12528))
* **New List Resource:** `google_compute_region_health_source` ([#12528](https://github.com/hashicorp/terraform-provider-google-beta/pull/12528))
* **New List Resource:** `google_project_iam_member` ([#12516](https://github.com/hashicorp/terraform-provider-google-beta/pull/12516))
* **New List Resource:** `google_pubsub_topic` ([#12525](https://github.com/hashicorp/terraform-provider-google-beta/pull/12525))
* **New List Resource:** `google_secret_manager_secret` ([#12521](https://github.com/hashicorp/terraform-provider-google-beta/pull/12521))

IMPROVEMENTS:
* apigee: added `consumer_key` and `consumer_secret` fields to `google_apigee_developer_app` to allow specifying a static credential ([#12487](https://github.com/hashicorp/terraform-provider-google-beta/pull/12487))
* artifactregistry: added update support to `upstream_credentials` in `google_artifact_registry_repository` ([#12486](https://github.com/hashicorp/terraform-provider-google-beta/pull/12486))
* biglakeiceberg: added `CATALOG_TYPE_BIGLAKE` enum to `catalog_type` field and added `restricted_locations_config.restricted_locations` field in `google_biglake_iceberg_catalog ` resource ([#12534](https://github.com/hashicorp/terraform-provider-google-beta/pull/12534))
* biglakeiceberg: added `sort_order` field to `google_biglake_iceberg_table` resource ([#12498](https://github.com/hashicorp/terraform-provider-google-beta/pull/12498))
* ces: added `timeout` and `tool_fake_config` fields to `google_ces_tool` and `google_ces_toolset` resource ([#12518](https://github.com/hashicorp/terraform-provider-google-beta/pull/12518))
* compute: added `params.resource_manager_tags` field to `google_compute_snapshot` resource ([#12502](https://github.com/hashicorp/terraform-provider-google-beta/pull/12502))
* compute: made `network_endpoints.ip_address` optional in `google_compute_network_endpoints` resource to support attaching endpoints to a network endpoint group of type `GCE_VM_IP_DEDICATED_BACKEND` ([#12503](https://github.com/hashicorp/terraform-provider-google-beta/pull/12503))
* container: added `dataplane_optimization_mode` in `google_container_cluster` ([#12494](https://github.com/hashicorp/terraform-provider-google-beta/pull/12494))
* container: added `ignore_node_count_changes` field to `google_container_cluster` and `google_container_node_pool` resources. When set to true, the provider ignores drift via external node count changes and skips related IGM API queries, resolving long plan times on clusters with a large number of instance groups. ([#12511](https://github.com/hashicorp/terraform-provider-google-beta/pull/12511))
* container: added `skip_node_pool_refresh` field to `google_container_cluster` resource. When set to true, the `google_container_cluster` skips refreshing and setting `node_pools` from the API, resolving long plan times on clusters with a large number of node pools. Note that this results in `node_pools` being set to an empty list in state ([#12511](https://github.com/hashicorp/terraform-provider-google-beta/pull/12511))
* container: added `taint_config` block to `google_container_cluster` and `google_container_node_pool` ([#12505](https://github.com/hashicorp/terraform-provider-google-beta/pull/12505))
* container: improved GKE node pool read performance by caching instance group metadata longer ([#12511](https://github.com/hashicorp/terraform-provider-google-beta/pull/12511))
* datastream: added `additional_options` field to `google_datastream_connection_profile` resource ([#12526](https://github.com/hashicorp/terraform-provider-google-beta/pull/12526))
* gkehub: added support for specifying `soak_duration` in minutes, hours and days to the `google_gke_hub_rollout_sequence` resource ([#12512](https://github.com/hashicorp/terraform-provider-google-beta/pull/12512))
* iamworkforcepool: added write-only support for `oidc.client_secret` in `google_iam_workforce_pool_provider` ([#12500](https://github.com/hashicorp/terraform-provider-google-beta/pull/12500))
* kms: added resource identity support for `google_kms_crypto_key_version` resource ([#12504](https://github.com/hashicorp/terraform-provider-google-beta/pull/12504))
* networkservices: added `dns_peering_config` field to `google_network_services_agent_gateway` resource ([#12481](https://github.com/hashicorp/terraform-provider-google-beta/pull/12481))
* sql: added `mode`, `dns_servers`, `admin_credential_secret_name`, and `organizational_unit` fields to `active_directory_config` block in `google_sql_database_instance` resource for SQL Server instances ([#12495](https://github.com/hashicorp/terraform-provider-google-beta/pull/12495))
* sql: added `performance_capture_config` to `google_sql_database_instance` ([#12536](https://github.com/hashicorp/terraform-provider-google-beta/pull/12536))
* storage: added `lifecycle_rule.condition.size_above_bytes` and `lifecycle_rule.condition.size_below_bytes` fields to `google_storage_bucket` resource ([#12490](https://github.com/hashicorp/terraform-provider-google-beta/pull/12490))

BUG FIXES:
* apigee: `google_apigee_developer_app` now updates `api_products` and `scopes` on the existing credential instead of creating a new credential (consumer key) on update ([#12533](https://github.com/hashicorp/terraform-provider-google-beta/pull/12533))
* biglake: allow `location` to be set on `google_biglake_iceberg_namespace` ([#12482](https://github.com/hashicorp/terraform-provider-google-beta/pull/12482))
* biglake: fixed creation failure of `google_biglake_iceberg_table` resource when the referenced `google_biglake_iceberg_catalog` has `credential_mode` set to `CREDENTIAL_MODE_VENDED_CREDENTIALS` due to a missing `X-Iceberg-Access-Delegation` header ([#12514](https://github.com/hashicorp/terraform-provider-google-beta/pull/12514))
* compute: fixed broken import of `share_settings` on `google_compute_reservation` ([#12527](https://github.com/hashicorp/terraform-provider-google-beta/pull/12527))
* datastream: fixed a positional diff when adding objects to the `salesforce_source_config.include_objects` field in `google_datastream_stream` resource ([#12531](https://github.com/hashicorp/terraform-provider-google-beta/pull/12531))
* iamworkforcepool: marked sensitive and ignore_read as true for `security_token` in `google_iam_workforce_pool_provider_scim_token` resource ([#12480](https://github.com/hashicorp/terraform-provider-google-beta/pull/12480))
* networkconnectivity: fixed `google_network_connectivity_regional_endpoint` being recreated on every apply when `address` is set to a resource URI ([#12529](https://github.com/hashicorp/terraform-provider-google-beta/pull/12529))
* networkservices: fixed `name` field expansion for `google_network_services_agent_gateway` resources so that short names are automatically expanded to full resource names, preventing API validation errors on create and update. ([#12513](https://github.com/hashicorp/terraform-provider-google-beta/pull/12513))

## 7.37.0 (June 16, 2026)

NOTES:
* compute: migrated `EnableDisplay` fields in `google_compute_instance` resources to use direct HTTP rather than a client library ([#12456](https://github.com/hashicorp/terraform-provider-google-beta/pull/12456))
* compute: migrated `desired_status` block and `startInstanceOperation` in `resource_compute_instance.go.tmpl` to use direct HTTP rather than a client library ([#12437](https://github.com/hashicorp/terraform-provider-google-beta/pull/12437))
* compute: migrated `getInstance`, `getDisk`, `Delete` and the `setMetadata` update block in `resource_compute_instance.go.tmpl` to use direct HTTP rather than a client library ([#12418](https://github.com/hashicorp/terraform-provider-google-beta/pull/12418))
* compute: migrated part of `google_compute_instance` to use direct HTTP rather than a client library ([#12463](https://github.com/hashicorp/terraform-provider-google-beta/pull/12463))

DEPRECATIONS:
* cloudrunv2: deprecated `http_get.http_headers.port` field in container startup probe and liveness probe in `google_cloud_run_v2_worker_pool` resource because it is not supported in Cloud Run API. This field will be removed in a future major release. ([#12476](https://github.com/hashicorp/terraform-provider-google-beta/pull/12476))
* cloudsecuritycompliance: deprecated the `organization` field on `google_cloud_security_compliance_cloud_control`, `google_cloud_security_compliance_framework`, and `google_cloud_security_compliance_framework_deployment`. Use `parent` instead ([#12449](https://github.com/hashicorp/terraform-provider-google-beta/pull/12449))
* networkservices: deprecated `protocols` on `google_network_services_agent_gateway` ([#12478](https://github.com/hashicorp/terraform-provider-google-beta/pull/12478))

FEATURES:
* **New Data Source:** `google_oracle_database_goldengate_deployment_versions` ([#12451](https://github.com/hashicorp/terraform-provider-google-beta/pull/12451))
* **New Data Source:** `google_storage_control_project_intelligence_finding` ([#12444](https://github.com/hashicorp/terraform-provider-google-beta/pull/12444))
* **New Data Source:** `google_storage_control_project_intelligence_findings` ([#12444](https://github.com/hashicorp/terraform-provider-google-beta/pull/12444))
* **New Resource:** `google_chronicle_parser` ([#12477](https://github.com/hashicorp/terraform-provider-google-beta/pull/12477))
* **New Resource:** `google_chronicle_soar_network` ([#12439](https://github.com/hashicorp/terraform-provider-google-beta/pull/12439))
* **New Resource:** `google_migration_center_import_data_file` ([#12423](https://github.com/hashicorp/terraform-provider-google-beta/pull/12423))
* **New Resource:** `google_network_services_agent_gateway` ([#12479](https://github.com/hashicorp/terraform-provider-google-beta/pull/12479))
* **New Resource:** `google_vertex_ai_tensorboard_experiment` ([#12472](https://github.com/hashicorp/terraform-provider-google-beta/pull/12472))
* **New List Resource:** `google_bigquery_dataset_access` ([#12441](https://github.com/hashicorp/terraform-provider-google-beta/pull/12441))
* **New List Resource:** `google_cloud_scheduler_job` ([#12441](https://github.com/hashicorp/terraform-provider-google-beta/pull/12441))
* **New List Resource:** `google_dns_record_set` ([#12467](https://github.com/hashicorp/terraform-provider-google-beta/pull/12467))
* **New List Resource:** `google_monitoring_alert_policy` ([#12441](https://github.com/hashicorp/terraform-provider-google-beta/pull/12441))
* **New List Resource:** `google_pubsub_subscription` ([#12441](https://github.com/hashicorp/terraform-provider-google-beta/pull/12441))

IMPROVEMENTS:
* apigee: added new resource `google_apigee_environment_debugmask` for managing Apigee environment debug masks ([#12421](https://github.com/hashicorp/terraform-provider-google-beta/pull/12421))
* backupdr: added support for `use_project_service_account` flag in `google_backup_dr_restore_workload` disk and compute restores ([#12473](https://github.com/hashicorp/terraform-provider-google-beta/pull/12473))
* cloudrunv2: added `http_get.http_headers.name` field to container startup probe and liveness probe in `google_cloud_run_v2_worker_pool` resource ([#12476](https://github.com/hashicorp/terraform-provider-google-beta/pull/12476))
* cloudrunv2: added `template.client` and `template.client_version` fields to `google_cloud_run_v2_worker_pool` resource ([#12440](https://github.com/hashicorp/terraform-provider-google-beta/pull/12440))
* cloudsecuritycompliance: added support for project parent to `google_cloud_security_compliance_cloud_control`, `google_cloud_security_compliance_framework`, and `google_cloud_security_compliance_framework_deployment` via the new `parent` field. The `organization` field has been deprecated. ([#12449](https://github.com/hashicorp/terraform-provider-google-beta/pull/12449))
* compute: added `params.resource_manager_tags` field to `google_compute_reservation` resource ([#12450](https://github.com/hashicorp/terraform-provider-google-beta/pull/12450))
* compute: added data sources for `google_compute_target_http_proxy`, `google_compute_target_https_proxy`, `google_compute_region_target_http_proxy`, and `google_compute_region_target_https_proxy` ([#12447](https://github.com/hashicorp/terraform-provider-google-beta/pull/12447))
* container: added `addons_config.slurm_operator_config` field to `google_container_cluster` resource ([#12445](https://github.com/hashicorp/terraform-provider-google-beta/pull/12445))
* container: added `node_image_config` field to `google_container_node_pool` and `google_container_cluster` resources ([#12470](https://github.com/hashicorp/terraform-provider-google-beta/pull/12470))
* databasemigrationservice: added `state` and `stop_on_warnings` fields to `google_database_migration_service_migration_job` resource ([#12430](https://github.com/hashicorp/terraform-provider-google-beta/pull/12430))
* dns: added resource identity support for `google_dns_record_set` resource ([#12467](https://github.com/hashicorp/terraform-provider-google-beta/pull/12467))
* networksecurity: added `network_rules` field on `google_network_security_authz_policy` resource ([#12488](https://github.com/hashicorp/terraform-provider-google-beta/pull/12488))
* pubsub: added `first_revision_id` and `last_revision_id` fields to `google_pubsub_topic` resource ([#12420](https://github.com/hashicorp/terraform-provider-google-beta/pull/12420))
* sql: added `settings.ip_configuration.psc_config.psc_auto_dns_enabled` and `settings.ip_configuration.psc_config.psc_write_endpoint_dns_enabled` fields to `google_sql_database_instance` resource ([#12454](https://github.com/hashicorp/terraform-provider-google-beta/pull/12454))

BUG FIXES:
* apigee: fixed `google_apigee_api` not detecting local bundle changes due to a missing default on `detect_md5hash`, and fixed the test sweeper's list URL ([#12466](https://github.com/hashicorp/terraform-provider-google-beta/pull/12466))
* apigee: fixed `google_apigee_security_action` update failure by enabling PATCH-based updates now that the Apigee Security Actions API supports mutations ([#12448](https://github.com/hashicorp/terraform-provider-google-beta/pull/12448))
* apigee: fixed a perma-diff for `api_products` and `scopes` fields in `google_apigee_developer_app` resource when updating them with multiple items ([#12464](https://github.com/hashicorp/terraform-provider-google-beta/pull/12464))
* apigee: fixed an issue where the resource would attempt recreation if the `key_expires_in` field was set in `google_apigee_developer_app` resource ([#12457](https://github.com/hashicorp/terraform-provider-google-beta/pull/12457))
* backupdr: updated the Beta provider to point to the `v1beta` API endpoint ([#12461](https://github.com/hashicorp/terraform-provider-google-beta/pull/12461))
* ces: fixed persistent diff in `google_ces_guardrail` when `llm_prompt_security` is configured with `default_settings` ([#12446](https://github.com/hashicorp/terraform-provider-google-beta/pull/12446))
* cloudrun: fixed a permadiff for the `run.googleapis.com/gpu-zonal-redundancy-disabled` annotation in `google_cloud_run_service` ([#12462](https://github.com/hashicorp/terraform-provider-google-beta/pull/12462))
* cloudrunv2: fixed bug where only one `http_get.http_headers` block could be specified in container startup probe and liveness probe in `google_cloud_run_v2_worker_pool` resource ([#12476](https://github.com/hashicorp/terraform-provider-google-beta/pull/12476))
* compute: fixed an issue in `google_compute_subnetwork` where `secondary_ip_range` entries linked to an `internal_range` could not be removed and adding new ranges would sometimes fail due to positional shifts ([#27175](https://github.com/hashicorp/terraform-provider-google/issues/27175)) ([#12422](https://github.com/hashicorp/terraform-provider-google-beta/pull/12422))
* compute: fixed diff when using `existing_reservations` field in `google_region_commitment` ([#12453](https://github.com/hashicorp/terraform-provider-google-beta/pull/12453))
* compute: fixed rules in `google_compute_security_policy` being unnecessarily recreated due to TypeSet hash instability ([#12436](https://github.com/hashicorp/terraform-provider-google-beta/pull/12436))
* sql: fixed inconsistent result after apply error when adding `users` of type `CLOUD_IAM_GROUP` with capitalized domain names for MySQL ([#12460](https://github.com/hashicorp/terraform-provider-google-beta/pull/12460))
* storage: fixed OOM issue for `google_storage_bucket` `force_destroy` by limiting the number of outstanding tasks to 2000 ([#12455](https://github.com/hashicorp/terraform-provider-google-beta/pull/12455))

## 7.36.0 (June 9, 2026)

FEATURES:
* **New Data Source:** `google_apigee_instance` ([#12393](https://github.com/hashicorp/terraform-provider-google-beta/pull/12393))
* **New Data Source:** `google_oracle_database_goldengate_deployment_types` ([#12361](https://github.com/hashicorp/terraform-provider-google-beta/pull/12361))
* **New Resource:** `google_apigee_datastore` ([#12358](https://github.com/hashicorp/terraform-provider-google-beta/pull/12358))
* **New Resource:** `google_chronicle_customlist` ([#12378](https://github.com/hashicorp/terraform-provider-google-beta/pull/12378))
* **New Resource:** `google_chronicle_environment` ([#12385](https://github.com/hashicorp/terraform-provider-google-beta/pull/12385))
* **New Resource:** `google_chronicle_findings_refinement` ([#12348](https://github.com/hashicorp/terraform-provider-google-beta/pull/12348))
* **New Resource:** `google_discovery_engine_search_engine_iam_binding` ([#12409](https://github.com/hashicorp/terraform-provider-google-beta/pull/12409))
* **New Resource:** `google_iam_organization_access_policy` ([#12398](https://github.com/hashicorp/terraform-provider-google-beta/pull/12398))
* **New Resource:** `google_license_manager_configuration` ([#12413](https://github.com/hashicorp/terraform-provider-google-beta/pull/12413))
* **New Resource:** `google_migration_center_import_job` ([#12351](https://github.com/hashicorp/terraform-provider-google-beta/pull/12351))
* **New List Resource:** `google_compute_disk` ([#12359](https://github.com/hashicorp/terraform-provider-google-beta/pull/12359))
* **New List Resource:** `google_compute_image` ([#12359](https://github.com/hashicorp/terraform-provider-google-beta/pull/12359))
* **New List Resource:** `google_compute_snapshot` ([#12359](https://github.com/hashicorp/terraform-provider-google-beta/pull/12359))
* **New List Resource:** `google_storage_hmac_key` ([#12364](https://github.com/hashicorp/terraform-provider-google-beta/pull/12364))

IMPROVEMENTS:
* accesscontextmanager: added in-place update for `egress_from` and `egress_to` fields in `google_access_context_manager_service_perimeter_egress_policy` resource ([#12397](https://github.com/hashicorp/terraform-provider-google-beta/pull/12397))
* accesscontextmanager: added in-place update for `egress_from` and `egress_to` fields in `google_access_context_manager_service_perimeter_ingress_policy` resource ([#12397](https://github.com/hashicorp/terraform-provider-google-beta/pull/12397))
* bigquery: added IAM support (`google_bigquery_routine_iam_policy`, `google_bigquery_routine_iam_binding`, `google_bigquery_routine_iam_member`) for `google_bigquery_routine` ([#12410](https://github.com/hashicorp/terraform-provider-google-beta/pull/12410))
* bigtable: added `automated_backup_policy.locations` field in `google_bigtable_table` resource ([#12365](https://github.com/hashicorp/terraform-provider-google-beta/pull/12365))
* ces: added `agent_tool`, `file_search_tool`, and `widget_tool` support to the `google_ces_tool` resource ([#12392](https://github.com/hashicorp/terraform-provider-google-beta/pull/12392))
* ces: added `google_search_tool.prompt_config` and `data_store_tool.data_store_source` fields in `google_ces_tool` resource ([#12392](https://github.com/hashicorp/terraform-provider-google-beta/pull/12392))
* ces: exposed `remote_agent_tool`, `connector_tool`, and `mcp_tool` as read-only (output-only) attributes in `google_ces_tool` ([#12392](https://github.com/hashicorp/terraform-provider-google-beta/pull/12392))
* container: added `node_creation_config` field to `google_container_cluster` resource ([#12408](https://github.com/hashicorp/terraform-provider-google-beta/pull/12408))
* container: added `node_drain_config.pdb_timeout_duration`  and `node_drain_config.grace_termination_duration` fields to `google_container_node_pool` and `google_container_cluster` resources ([#12401](https://github.com/hashicorp/terraform-provider-google-beta/pull/12401))
* data_catalog: added `RICHTEXT` to allowed values of `primitive_type` on `google_data_catalog_tag_template` fields. ([#12383](https://github.com/hashicorp/terraform-provider-google-beta/pull/12383))
* dataform: added `disabled` field to `google_dataform_repository_workflow_config` resource ([#12394](https://github.com/hashicorp/terraform-provider-google-beta/pull/12394))
* dataplex: added IAM support for `google_dataplex_data_product` (`iam_policy`, `iam_binding`, `iam_member`) ([#12371](https://github.com/hashicorp/terraform-provider-google-beta/pull/12371))
* dataplex: added `access_approval_config` field to `google_dataplex_data_product` ([#12371](https://github.com/hashicorp/terraform-provider-google-beta/pull/12371))
* firebaseailogic: added `regional_propagation_disabled` field to `google_firebase_ai_logic_prompt_template` and `google_firebase_ai_logic_prompt_template_lock` resources ([#12349](https://github.com/hashicorp/terraform-provider-google-beta/pull/12349))
* hypercomputecluster: marked `network_resources` field as required in `google_hypercomputecluster_cluster` resource to align with API validation ([#12374](https://github.com/hashicorp/terraform-provider-google-beta/pull/12374))
* securesourcemanager: added `psc_allowed_projects` field to `google_secure_source_manager_instance` resource ([#12402](https://github.com/hashicorp/terraform-provider-google-beta/pull/12402))
* securityscanner: added `ignore_http_status_errors` field to `google_security_scanner_scan_config` resource ([#12352](https://github.com/hashicorp/terraform-provider-google-beta/pull/12352))
* workbench: added `NVIDIA_RTX6000` to the supported `gce_setup.accelerator_configs.type` values on `google_workbench_instance` ([#12415](https://github.com/hashicorp/terraform-provider-google-beta/pull/12415))

BUG FIXES:
* apigee: send zero values for `ip_header_index` in `google_apigee_environment` ([#12381](https://github.com/hashicorp/terraform-provider-google-beta/pull/12381))
* backupdr: fixed an issue where `google_backup_dr_restore_workload` did not use the correct API JSON names for networking/reservation fields ([#12391](https://github.com/hashicorp/terraform-provider-google-beta/pull/12391))
* compute: fixed an issue where updating `connection_limit` in the `consumer_accept_lists` block of `google_compute_service_attachment` would not trigger a resource update. ([#12395](https://github.com/hashicorp/terraform-provider-google-beta/pull/12395))
* compute: fixed regional backend reference in  `google_compute_regional_url_map` ([#12411](https://github.com/hashicorp/terraform-provider-google-beta/pull/12411))
* dlp: fixed error when reading `google_data_loss_prevention_discovery_config` caused by nested error details ([#12380](https://github.com/hashicorp/terraform-provider-google-beta/pull/12380))
* firebaseailogic: fixed perma-diff in `google_firebase_ai_logic_config` `template_only` field on `google_firebase_ai_logic_config` resource ([#12375](https://github.com/hashicorp/terraform-provider-google-beta/pull/12375))
* sql: fixed permadiff on `connection_pool_config` when `connection_pooling_enabled` is set to `false` ([#12417](https://github.com/hashicorp/terraform-provider-google-beta/pull/12417))
* storage: fixed missing identity error when updating values in `google_storage_bucket` resource ([#12356](https://github.com/hashicorp/terraform-provider-google-beta/pull/12356))
* tags: fixed `google_tags_location_tag_binding` failing with `Operation location does not match service location 'global'` during creation ([#12379](https://github.com/hashicorp/terraform-provider-google-beta/pull/12379))
* vertexai: fixed `terraform import` of `google_vertex_ai_index_endpoint_deployed_index` failing with "Cannot determine region" when provider-level `region`/`zone` is unset ([#12399](https://github.com/hashicorp/terraform-provider-google-beta/pull/12399))

## 7.35.0 (June 2, 2026)

FEATURES:
* **New Data Source:** `google_oracle_database_goldengate_connection_types` ([#12329](https://github.com/hashicorp/terraform-provider-google-beta/pull/12329))
* **New Resource:** `google_ces_security_settings` ([#12327](https://github.com/hashicorp/terraform-provider-google-beta/pull/12327))
* **New Resource:** `google_chronicle_findings_refinement` ([#12348](https://github.com/hashicorp/terraform-provider-google-beta/pull/12348))
* **New Resource:** `google_iam_folder_access_policy` ([#12326](https://github.com/hashicorp/terraform-provider-google-beta/pull/12326))
* **New Resource:** `google_iam_project_access_policy` ([#12317](https://github.com/hashicorp/terraform-provider-google-beta/pull/12317))
* **New Resource:** `google_migration_center_discovery_client` ([#12333](https://github.com/hashicorp/terraform-provider-google-beta/pull/12333))
* **New Resource:** `google_migration_center_report` ([#12322](https://github.com/hashicorp/terraform-provider-google-beta/pull/12322))
* **New Resource:** `google_oracle_database_goldengate_connection_assignment` ([#12328](https://github.com/hashicorp/terraform-provider-google-beta/pull/12328))
* **New Resource:** `google_oracle_database_goldengate_connection` ([#12344](https://github.com/hashicorp/terraform-provider-google-beta/pull/12344))
* **New Resource:** `google_oracle_database_goldengate_deployment` ([#12336](https://github.com/hashicorp/terraform-provider-google-beta/pull/12336))
* **New List Resource:** `google_compute_firewall` ([#12323](https://github.com/hashicorp/terraform-provider-google-beta/pull/12323))
* **New List Resource:** `google_compute_global_address` ([#12323](https://github.com/hashicorp/terraform-provider-google-beta/pull/12323))
* **New List Resource:** `google_compute_subnetwork` ([#12323](https://github.com/hashicorp/terraform-provider-google-beta/pull/12323))
* **New List Resource:** `google_sql_database` ([#12325](https://github.com/hashicorp/terraform-provider-google-beta/pull/12325))

IMPROVEMENTS:
* compute: added `params.resource_manager_tags` field to `google_compute_future_reservation` resource (beta) ([#12332](https://github.com/hashicorp/terraform-provider-google-beta/pull/12332))
* compute: added `target_type` and `target_forwarding_rules` fields to `google_compute_network_firewall_policy_rule` resource ([#12314](https://github.com/hashicorp/terraform-provider-google-beta/pull/12314))
* container: added `crash_loop_back_off.max_container_restart_period` field to `google_container_node_pool` and `google_container_cluster` resources ([#12335](https://github.com/hashicorp/terraform-provider-google-beta/pull/12335))
* container: added additional value `KCP_VPA` for `logging_config.enable_components` field to `google_container_cluster` resource ([#12320](https://github.com/hashicorp/terraform-provider-google-beta/pull/12320))
* dataplex: added `service_account` support to `google_dataplex_data_product` access group principals ([#12345](https://github.com/hashicorp/terraform-provider-google-beta/pull/12345))
* firestore: added `ttl_config.expiration_offset` field to `google_firestore_field` resource ([#12346](https://github.com/hashicorp/terraform-provider-google-beta/pull/12346))
* netapp: marked `ontap_source.snapshot_uuid` field as Computed in `google_netapp_backup` resource ([#12341](https://github.com/hashicorp/terraform-provider-google-beta/pull/12341))
* networkmanagement: added `gke_pod` and `network_type` fields to `google_network_management_connectivity_test` resource ([#12342](https://github.com/hashicorp/terraform-provider-google-beta/pull/12342))

BUG FIXES:
* resourcemanager: fixed a bug where ephemeral `google_service_account_key` failed on deletion if the parent service account had already been deleted ([#12316](https://github.com/hashicorp/terraform-provider-google-beta/pull/12316))
* storage: fixed missing identity error when updating values in `google_storage_bucket` ([#12356](https://github.com/hashicorp/terraform-provider-google-beta/pull/12356))

## 7.34.0 (May 27, 2026)

NOTES:
* compute: migrated `google_compute_region_instance_template` to use direct HTTP rather than a client library ([#27471](https://github.com/hashicorp/terraform-provider-google/pull/27471))
* compute: migrated `google_compute_instance_group_manager` resource to use direct HTTP rather than a client library ([#27441](https://github.com/hashicorp/terraform-provider-google/pull/27441))

FEATURES:
* **New Data Source:** `google_compute_service_attachment` ([#12305](https://github.com/hashicorp/terraform-provider-google-beta/pull/12305))
* **New Data Source:** `google_oracle_database_goldengate_deployment_environments` ([#12287](https://github.com/hashicorp/terraform-provider-google-beta/pull/12287))
* **New Resource:** `google_config_deployment` ([#12254](https://github.com/hashicorp/terraform-provider-google-beta/pull/12254))
* **New Resource:** `google_dialogflow_sip_trunk` ([#12268](https://github.com/hashicorp/terraform-provider-google-beta/pull/12268))
* **New Resource:** `google_migration_center_assets_export_job` ([#12266](https://github.com/hashicorp/terraform-provider-google-beta/pull/12266))
* **New Resource:** `google_migration_center_report_config` ([#12243](https://github.com/hashicorp/terraform-provider-google-beta/pull/12243))
* **New Resource:** `google_migration_center_settings` ([#12265](https://github.com/hashicorp/terraform-provider-google-beta/pull/12265))
* **New Resource:** `google_migration_center_source` ([#12284](https://github.com/hashicorp/terraform-provider-google-beta/pull/12284))

IMPROVEMENTS:
* bigtable: added `edition` field to `google_bigtable_instance` resource ([#12294](https://github.com/hashicorp/terraform-provider-google-beta/pull/12294))
* ces: added `fail_open` field to `llm_prompt_security` block in `google_ces_guardrail` resource ([#12285](https://github.com/hashicorp/terraform-provider-google-beta/pull/12285))
* ces: added read-only `fail_open` field to `llm_prompt_security` block in `google_ces_app_version` resource ([#12285](https://github.com/hashicorp/terraform-provider-google-beta/pull/12285))
* compute: added `alias_ipv6` field to `google_compute_instance` and `google_compute_instance_template` (beta) ([#12288](https://github.com/hashicorp/terraform-provider-google-beta/pull/12288))
* compute: added `ipVersion` and `ipCollection` fields to `secondaryIpRange` field in `google_compute_subnetwork` resource ([#12249](https://github.com/hashicorp/terraform-provider-google-beta/pull/12249))
* compute: added `post_quantum_key_exchange` field to `google_compute_ssl_policy` and `google_compute_region_ssl_policy` resources ([#12276](https://github.com/hashicorp/terraform-provider-google-beta/pull/12276))
* compute: added support in the `google_compute_network` datasource for looking up a network by `self_link` in addition to `name` ([#27509]
* container: added `agent_sandbox_config` field to `google_container_cluster` resource ([#12279](https://github.com/hashicorp/terraform-provider-google-beta/pull/12279))
* container: added `node_config.gpudirect_strategy` and `node_pool.node_config.gpudirect_strategy` to `cluster` resource, added `node_config.gpudirect_strategy` to `node_pool` resource ([#12283](https://github.com/hashicorp/terraform-provider-google-beta/pull/12283))
* dataflow: added  `create_ignore_already_exists`  field to  `google_dataflow_flex_template_job` resource to handle 409 conflicts ([#12274](https://github.com/hashicorp/terraform-provider-google-beta/pull/12274))
* datafusion: added `maintenance_policy` field to `google_data_fusion_instance` resource ([#12269](https://github.com/hashicorp/terraform-provider-google-beta/pull/12269))
* dialogflow: added `sip_config` field to `google_dialogflow_conversation_profile` resource (beta) ([#12259](https://github.com/hashicorp/terraform-provider-google-beta/pull/12259))
* firebasestorage: added `bucket.bucket_id` field to `google_firebase_storage_default_bucket` resource (beta) ([#12297](https://github.com/hashicorp/terraform-provider-google-beta/pull/12297))
* iam: add resource identity support for `iam_member` resources ([#12236](https://github.com/hashicorp/terraform-provider-google-beta/pull/12236))
* networkconnectivity: promoted `google_network_connectivity_transport` resource to GA ([#12256](https://github.com/hashicorp/terraform-provider-google-beta/pull/12256))
* oracledatabase: added `identity_connector` to `google_oracle_database_cloud_vm_cluster` for CMEK support ([#12252](https://github.com/hashicorp/terraform-provider-google-beta/pull/12252))
* project: add Resource Identity support to `google_project_iam_binding` ([#12290](https://github.com/hashicorp/terraform-provider-google-beta/pull/12290))
* project: add Resource Identity support to `google_project_iam_policy` ([#12291](https://github.com/hashicorp/terraform-provider-google-beta/pull/12291))
* sql: promoted Hyperdisk fields, `data_disk_provisioned_iops` and `data_disk_provisioned_throughput`  to GA ([#12253](https://github.com/hashicorp/terraform-provider-google-beta/pull/12253))

BUG FIXES:
* bigtable: fixed an issue where `bigtable_custom_endpoint` and `universe_domain` were ignored when creating Bigtable resources. ([#12300](https://github.com/hashicorp/terraform-provider-google-beta/pull/12300))
* compute: fixed an issue in `google_compute_subnetwork` where `secondary_ip_range` entries linked to an `internal_range` could not be removed and adding new ranges would sometimes fail due to positional shifts ([#27175](https://github.com/hashicorp/terraform-provider-google/issues/27175)) ([#12298](https://github.com/hashicorp/terraform-provider-google-beta/pull/12298))
* compute: marked encryption keys as immutable and sensitive across compute and backupdr resources ([#12295](https://github.com/hashicorp/terraform-provider-google-beta/pull/12295))
* dialogflow: corrected `AUDIOENCODING_SPEEX_WITH_HEADER_BYTE` ENUM value to `AUDIO_ENCODING_SPEEX_WITH_HEADER_BYTE` for `audio_encoding` field in `google_dialogflow_conversation_profile` resource ([#12259](https://github.com/hashicorp/terraform-provider-google-beta/pull/12259))
* resourcemanager: resolved a one-time diff for `deletion_policy` that would occur on existing and imported `google_project_service` resources following upgrading to v7.32.0 ([#12281](https://github.com/hashicorp/terraform-provider-google-beta/pull/12281))

## 7.33.0 (May 19, 2026)

NOTES:
* compute: migrate `google_compute_target_pool` resource to use direct HTTP rather than a client library ([#12212](https://github.com/hashicorp/terraform-provider-google-beta/pull/12212))
* compute: migrated `google_compute_instance_group_manager` resource  to use direct HTTP rather than a client library ([#12206](https://github.com/hashicorp/terraform-provider-google-beta/pull/12206))
* compute: migrated `google_compute_project_default_network_tier` resource to use direct HTTP rather than a client library ([#12201](https://github.com/hashicorp/terraform-provider-google-beta/pull/12201))
* compute: migrated `google_compute_router_status` data source to use direct HTTP rather than a client library ([#12174](https://github.com/hashicorp/terraform-provider-google-beta/pull/12174))
* compute: migrated `google_compute_instance_group_manager` resource to use direct HTTP rather than a client library ([#12216](https://github.com/hashicorp/terraform-provider-google-beta/pull/12216))
* compute: partially migrated `google_compute_instance` resource to use direct HTTP rather then a client library ([#12205](https://github.com/hashicorp/terraform-provider-google-beta/pull/12205))

FEATURES:
* **New Data Source:** `google_logging_log_view` ([#12226](https://github.com/hashicorp/terraform-provider-google-beta/pull/12226))
* **New Resource:** `google_apigee_data_collector` ([#12190](https://github.com/hashicorp/terraform-provider-google-beta/pull/12190))
* **New Resource:** `google_contact_center_insights_encryption_spec` ([#12225](https://github.com/hashicorp/terraform-provider-google-beta/pull/12225))

IMPROVEMENTS:
* backupdr: added `guest_flush` field to `google_backup_dr_backup_plan` resource and `google_backup_dr_backup` data source. ([#12229](https://github.com/hashicorp/terraform-provider-google-beta/pull/12229))
* backupdr: added `guest_flush` field to `google_backup_dr_backup_plan` resource and `google_backup_dr_backup` data source. ([#12230](https://github.com/hashicorp/terraform-provider-google-beta/pull/12230))
* ces: added `security_settings` field to `google_ces_deployment` resource ([#12227](https://github.com/hashicorp/terraform-provider-google-beta/pull/12227))
* ces: added `tool_execution_mode` field to `google_ces_app` resource ([#12221](https://github.com/hashicorp/terraform-provider-google-beta/pull/12221))
* compute: added `stabilization_period` field to `google_compute_autoscaler` and `google_compute_region_autoscaler` resources ([#12232](https://github.com/hashicorp/terraform-provider-google-beta/pull/12232))
* compute: added support for "ARP_BROADCAST_PRIMARY_RANGE" values to the `resolve_subnet_mask` field in `google_compute_subnetwork` resource ([#12176](https://github.com/hashicorp/terraform-provider-google-beta/pull/12176))
* compute: added support for "GCE_VM_IP_DEDICATED_BACKEND" to the  `network_endpoint_type` field in `google_compute_network_endpoint_group` resource ([#12176](https://github.com/hashicorp/terraform-provider-google-beta/pull/12176))
* compute: migrated `data_source_google_compute_regions` to use direct HTTP rather than a client library ([#12202](https://github.com/hashicorp/terraform-provider-google-beta/pull/12202))
* container: added `pod_snapshot_config` field to `google_container_cluster` resource (GA) ([#12196](https://github.com/hashicorp/terraform-provider-google-beta/pull/12196))
* container: added `secret_sync_config` field to `google_container_cluster` resource (ga) ([#12215](https://github.com/hashicorp/terraform-provider-google-beta/pull/12215))
* databasemigrationservice: added `database` and `private_connectivity` fields to `google_database_migration_service_connection_profile` resource ([#12203](https://github.com/hashicorp/terraform-provider-google-beta/pull/12203))
* databasemigrationservice: added `postgres_homogeneous_config` field to `google_database_migration_service_migration_job` resource ([#12203](https://github.com/hashicorp/terraform-provider-google-beta/pull/12203))
* databasemigrationservice: added `psc_interface_config` field to `google_database_migration_service_private_connection` resource ([#12184](https://github.com/hashicorp/terraform-provider-google-beta/pull/12184))
* hypercomputecluster: added `terminal_storage_class` and `per_unit_storage_throughput` fields to the `google_hypercomputecluster_cluster` resource ([#12234](https://github.com/hashicorp/terraform-provider-google-beta/pull/12234))
* netapp: added `ontap_source` field to `google_netapp_backup` resource (beta) ([#12231](https://github.com/hashicorp/terraform-provider-google-beta/pull/12231))
* provider: support for a `deletion_policy` field has been added to almost all resources in the provider. Details on its usage can be found within individual resource documentation if supported. ([#12183](https://github.com/hashicorp/terraform-provider-google-beta/pull/12183))
* storagebatchoperations: added `description` field to `google_storage_batch_operations_job` resource ([#12207](https://github.com/hashicorp/terraform-provider-google-beta/pull/12207))
* workstations: added `workstation_authorization_url` and `workstation_launch_url` fields to the `google_workstations_workstation_cluster ` resource. ([#12185](https://github.com/hashicorp/terraform-provider-google-beta/pull/12185))

BUG FIXES:
* apigee: fixed forced replacement when importing `google_apigee_sharedflow_deployment` resource, where `service_account` read as null ([#12228](https://github.com/hashicorp/terraform-provider-google-beta/pull/12228))
* bigqueryconnection: fixed an issue where `configuration.authentication.username_password.password.secret_type` is not populated and a diff on `configuration.authentication.username_password.username` after import in `google_bigquery_connection` resource ([#12179](https://github.com/hashicorp/terraform-provider-google-beta/pull/12179))
* bigqueryreservation: Fixed `google_bigquery_reservation_assignment` returning a confusing 404 error when `reservation` is a bare name and `location` is not set ([#12210](https://github.com/hashicorp/terraform-provider-google-beta/pull/12210))
* ces: updated supported values for `channel_type`, `modality`, and `theme` in `google_ces_deployment` ([#12227](https://github.com/hashicorp/terraform-provider-google-beta/pull/12227))
* compute: updated `google_compute_forwarding_rule` resource to properly prompt for resource recreation when updating the `target` field between different "serviceAttachments", rather than having an in-place update blocked by an API error. ([#12214](https://github.com/hashicorp/terraform-provider-google-beta/pull/12214))
* modelarmor: fixed permadiff and `REQUEST_FIELD_MISSING` error when `template_metadata` is omitted from `google_model_armor_template` ([#12222](https://github.com/hashicorp/terraform-provider-google-beta/pull/12222))
* networkconnectivity: fixed an issue where `google_network_connectivity_destination` was not recognizing the `name` field as mapping to an API value ([#12224](https://github.com/hashicorp/terraform-provider-google-beta/pull/12224))
* networkconnectivity: fixed an issue where `google_network_connectivity_multicloud_data_transfer_config` was not recognizing the `name` field as mapping to an API value ([#12224](https://github.com/hashicorp/terraform-provider-google-beta/pull/12224))
* resourcemanager: added verification polling to `google_service_account` updates to ensure the resource is consistent before succeeding ([#12217](https://github.com/hashicorp/terraform-provider-google-beta/pull/12217))

## 7.32.0 (May 12, 2026)
NOTES:
* compute: migrated `google_compute_instance_from_machine` resource  to use direct HTTP rather than a client library ([#12152](https://github.com/hashicorp/terraform-provider-google-beta/pull/12152))
* compute: migrated `google_compute_instance_group_manager` resource to use direct HTTP rather than a client library ([#12151](https://github.com/hashicorp/terraform-provider-google-beta/pull/12151))
* compute: migrated `google_compute_zones` data source to use direct HTTP rather than a client library ([#12153](https://github.com/hashicorp/terraform-provider-google-beta/pull/12153))
* compute: migrated `google_compute_project_metadata_item` resource to use direct HTTP rather than a client library ([#12117](https://github.com/hashicorp/terraform-provider-google-beta/pull/12117))

FEATURES:
* **New Data Source:** `google_compute_region_instant_snapshot_iam_policy` ([#12172](https://github.com/hashicorp/terraform-provider-google-beta/pull/12172))
* **New Resource:** `google_compute_region_instant_snapshot_iam_binding` ([#12172](https://github.com/hashicorp/terraform-provider-google-beta/pull/12172))
* **New Resource:** `google_compute_region_instant_snapshot_iam_member` ([#12172](https://github.com/hashicorp/terraform-provider-google-beta/pull/12172))
* **New Resource:** `google_compute_region_instant_snapshot_iam_policy` ([#12172](https://github.com/hashicorp/terraform-provider-google-beta/pull/12172))
* **New Resource:** `google_compute_region_instant_snapshot` ([#12172](https://github.com/hashicorp/terraform-provider-google-beta/pull/12172))

IMPROVEMENTS:
* ces: added `scenario` field to `google_ces_evaluation` resource ([#12125](https://github.com/hashicorp/terraform-provider-google-beta/pull/12125))
* compute: added `IDPF` value to `nic_type` in `resource_compute_instance_template` ([#12142](https://github.com/hashicorp/terraform-provider-google-beta/pull/12142))
* compute: added `IDPF` value to `nic_type` in `resource_compute_instance` ([#12142](https://github.com/hashicorp/terraform-provider-google-beta/pull/12142))
* compute: added `IDPF` value to `nic_type` in `resource_compute_region_instance_template` ([#12142](https://github.com/hashicorp/terraform-provider-google-beta/pull/12142))
* compute: added `address_id` field to `google_compute_address` resource ([#12126](https://github.com/hashicorp/terraform-provider-google-beta/pull/12126))
* compute: added `advanced_options_config` field on `google_compute_organization_security_policy` resource ([#12147](https://github.com/hashicorp/terraform-provider-google-beta/pull/12147))
* compute: added `image`, `source_image_encryption_key`, and `source_image_id` fields to `google_compute_region_disk` resource. This field is currently behind an allowlist. ([#12141](https://github.com/hashicorp/terraform-provider-google-beta/pull/12141))
* compute: added `replica_zones` field to `google_compute_instance` resource ([#12150](https://github.com/hashicorp/terraform-provider-google-beta/pull/12150))
* compute: added `request_body` field on `google_compute_security_policy_rule` resource ([#12144](https://github.com/hashicorp/terraform-provider-google-beta/pull/12144))
* compute: added update support for `ip_collection` field to `google_compute_subnetwork` resource ([#12157](https://github.com/hashicorp/terraform-provider-google-beta/pull/12157))
* discoveryengine: added `config_id` attribute to `google_discovery_engine_widget_config` ([#12169](https://github.com/hashicorp/terraform-provider-google-beta/pull/12169))
* gkehub: added `auto_upgrade_config` field to `google_gke_hub_rollout_sequence` resource ([#12155](https://github.com/hashicorp/terraform-provider-google-beta/pull/12155))
* networksecurity: added support for project `parent` values to `google_network_security_firewall_endpoint` ([#12132](https://github.com/hashicorp/terraform-provider-google-beta/pull/12132))
* recaptchaenterprise: added `POLICY_BASED_CHALLENGE` value to `integration_type` field and added new `challenge_settings` field to `google_recaptcha_enterprise_key` ([#12131](https://github.com/hashicorp/terraform-provider-google-beta/pull/12131))
* redis: added new node types supported in `google_redis_cluster`. ([#12140](https://github.com/hashicorp/terraform-provider-google-beta/pull/12140))
* resourcemanager: add `private_key` and `private_key_type` fields to ephemeral `google_service_account_key` resource ([#12170](https://github.com/hashicorp/terraform-provider-google-beta/pull/12170))
* storage: added `ingest_on_write` field for `google_storage_anywhere_cache` resource ([#12162](https://github.com/hashicorp/terraform-provider-google-beta/pull/12162))
* workstations: added `gce_hd` field to `google_workstations_workstation_config` resource ([#12118](https://github.com/hashicorp/terraform-provider-google-beta/pull/12118))

BUG FIXES:
* cloudfunctions2: fixed bug where `all_traffic_on_latest_revision = false` was ignored in `google_cloudfunctions2_function` ([#12148](https://github.com/hashicorp/terraform-provider-google-beta/pull/12148))
* compute: fixed permadiff when removing `preconfigured_waf_config` from a `google_compute_security_policy` rule ([#12167](https://github.com/hashicorp/terraform-provider-google-beta/pull/12167))

## 7.31.0 (May 5, 2026)
NOTES:
* compute: migrated `google_compute_instance.network_interface` field to use direct HTTP rather than a client library ([#12074](https://github.com/hashicorp/terraform-provider-google-beta/pull/12074))
* compute: migrated `google_compute_image` datasource to use direct HTTP rather then a client library ([#12103](https://github.com/hashicorp/terraform-provider-google-beta/pull/12103))
* compute: migrated `partner_metadata` field on `google_compute_instance`, `google_compute_instance_template`, and `google_compute_region_instance_template` to use direct HTTP rather than a client library ([#12090](https://github.com/hashicorp/terraform-provider-google-beta/pull/12090))
* compute: migrated `google_compute_node_types` data source to use direct HTTP rather than a client library ([#12108](https://github.com/hashicorp/terraform-provider-google-beta/pull/12108))
* compute: migrated `google_compute_region_instance_group` data source to use direct HTTP rather than a client library ([#12102](https://github.com/hashicorp/terraform-provider-google-beta/pull/12102))
* compute: migrated `google_compute_subnetwork` data source to use direct HTTP rather than a client library ([#12095](https://github.com/hashicorp/terraform-provider-google-beta/pull/12095))
* compute: migrated `google_compute_vpn_gateway` data source to use direct HTTP rather than a client library ([#12096](https://github.com/hashicorp/terraform-provider-google-beta/pull/12096))

FEATURES:
* **New Data Source:** `google_artifact_registry_file` ([#12107](https://github.com/hashicorp/terraform-provider-google-beta/pull/12107))
* **New Resource:** `google_ces_app_root_agent_association` ([#12083](https://github.com/hashicorp/terraform-provider-google-beta/pull/12083))
* **New Resource:** `google_ces_evaluation` ([#12114](https://github.com/hashicorp/terraform-provider-google-beta/pull/12114))
* **New Resource:** `google_contact_center_insights_qa_question` ([#12097](https://github.com/hashicorp/terraform-provider-google-beta/pull/12097))
* **New Resource:** `google_contact_center_insights_qa_scorecard_revision` ([#12097](https://github.com/hashicorp/terraform-provider-google-beta/pull/12097))
* **New Resource:** `google_contact_center_insights_qa_scorecard` ([#12097](https://github.com/hashicorp/terraform-provider-google-beta/pull/12097))
* **New Resource:** `google_firebase_app_check_resource_policy` ([#12109](https://github.com/hashicorp/terraform-provider-google-beta/pull/12109))
* **New Resource:** `google_network_security_ull_mirroring_collector_rule` ([#12075](https://github.com/hashicorp/terraform-provider-google-beta/pull/12075))

IMPROVEMENTS:
* ces: switched to the v1beta API for Customer Engagement Suite resources ([#12114](https://github.com/hashicorp/terraform-provider-google-beta/pull/12114))
* clouddeploy: added `default_pool` and `private_pool` fields to `google_clouddeploy_target` resource ([#12110](https://github.com/hashicorp/terraform-provider-google-beta/pull/12110))
* clouddeploy: added `tasks` and `analysis` fields to `google_clouddeploy_delivery_pipeline` resource ([#12110](https://github.com/hashicorp/terraform-provider-google-beta/pull/12110))
* compute: added `params.resource_manager_tags` field to `google_compute_image` ([#12076](https://github.com/hashicorp/terraform-provider-google-beta/pull/12076))
* compute: added `params.resource_manager_tags` field to `google_compute_region_commitment` resource ([#12105](https://github.com/hashicorp/terraform-provider-google-beta/pull/12105))
* compute: added `resource_policies.workload_policy` to `google_compute_region_instance_group_manager` resource ([#12098](https://github.com/hashicorp/terraform-provider-google-beta/pull/12098))
* compute: marked csek disk encryption key fields as sensitive in compute resources ([#12115](https://github.com/hashicorp/terraform-provider-google-beta/pull/12115))
* databasemigrationservice: added `objects_config` field to `google_database_migration_service_migration_job` resource ([#12104](https://github.com/hashicorp/terraform-provider-google-beta/pull/12104))
* dataplex: added `attributes`, `template_reference`, `enable_catalog_basedRules`, and `filter` fields to `google_dataplex_datascan` resource ([#12089](https://github.com/hashicorp/terraform-provider-google-beta/pull/12089))
* firestore: added `search_config` field to `google_firestore_index` resource ([#12077](https://github.com/hashicorp/terraform-provider-google-beta/pull/12077))
* oracle_database: added `pluggable_database_id`, `pluggable_database_name` fields to `google_oracle_database_db_system` resource ([#12087](https://github.com/hashicorp/terraform-provider-google-beta/pull/12087))

BUG FIXES:
* provider: fixed a bad `timeouts` diff across a number of resources that had resource identity support added in `7.29.0` ([#12112](https://github.com/hashicorp/terraform-provider-google-beta/pull/12112))
* bigquery: fixed `ignore_auto_generated_schema` evaluation for `google_bigquery_table` external tables which caused spurious replacement ([#12111](https://github.com/hashicorp/terraform-provider-google-beta/pull/12111))
* cloudscheduler: fixed perpetual diff on `google_cloud_scheduler_job.http_target.headers` when `oidc_token` or `oauth_token` is set ([#12101](https://github.com/hashicorp/terraform-provider-google-beta/pull/12101))
* servicenetworking: fixed a permadiff issue of `reserved_peering_ranges` in `google_service_networking_connection` ([#12091](https://github.com/hashicorp/terraform-provider-google-beta/pull/12091))
* storage: fix inconsistent plan issue for `google_storage_notification.custom_attributes` field ([#12088](https://github.com/hashicorp/terraform-provider-google-beta/pull/12088))

## 7.30.0 (Apr 28, 2026)

BREAKING CHANGES:
* apigee: fixed `google_apigee_env_keystore` to require the `name` field which is mandatory in the Apigee API ([#12006](https://github.com/hashicorp/terraform-provider-google-beta/pull/12006))

FEATURES:
* **New Data Source:** `google_data_lineage_config` ([#12068](https://github.com/hashicorp/terraform-provider-google-beta/pull/12068))
* **New Resource:** `google_artifact_registry_rule` ([#12031](https://github.com/hashicorp/terraform-provider-google-beta/pull/12031))
* **New Resource:** `google_data_lineage_config` ([#12068](https://github.com/hashicorp/terraform-provider-google-beta/pull/12068))
* **New Resource:** `google_document_ai_schema` ([#12072](https://github.com/hashicorp/terraform-provider-google-beta/pull/12072))
* **New Resource:** `google_firebase_remote_config_remote_config` ([#12032](https://github.com/hashicorp/terraform-provider-google-beta/pull/12032))

IMPROVEMENTS:
* provider: added support for `prefer_global_endpoints` and `prefer_regional_endpoints` to the provider configuration. Support for regional endpoints will be rolled out on a per-product level ([#12014](https://github.com/hashicorp/terraform-provider-google-beta/pull/12014))
* artifactregistry: added support for regionalized endpoints ([#12014](https://github.com/hashicorp/terraform-provider-google-beta/pull/12014))
* assuredworkloads: added `SPAIN_DATA_BOUNDARY_BY_TELEFONICA` value to `partner` field on `google_assured_workloads_workload` resource ([#12018](https://github.com/hashicorp/terraform-provider-google-beta/pull/12018))
* bigqueryconnection: added `configuration` block to `google_bigquery_connection` resource to support AlloyDB and other connector types via the BigQuery Connector framework ([#12021](https://github.com/hashicorp/terraform-provider-google-beta/pull/12021))
* bigtable: added support for `tags` to `google_bigtable_instance` ([#12041](https://github.com/hashicorp/terraform-provider-google-beta/pull/12041))
* cloudrunv2: added `DISK` fields to `google_cloud_run_v2_job` resource ([#12034](https://github.com/hashicorp/terraform-provider-google-beta/pull/12034))
* cloudrunv2: added `DISK` fields to `google_cloud_run_v2_worker_pool` resource ([#12030](https://github.com/hashicorp/terraform-provider-google-beta/pull/12030))
* compute: add `params.resourceManagerTags` field to the `google_compute_storage_pool` ([#12033](https://github.com/hashicorp/terraform-provider-google-beta/pull/12033))
* compute: added `erase_windows_vss_signature` field to `google_compute_instance` (beta) ([#12025](https://github.com/hashicorp/terraform-provider-google-beta/pull/12025))
* compute: added `params.resource_manager_tags` field to `google_compute_instant_snapshot` resource ([#12062](https://github.com/hashicorp/terraform-provider-google-beta/pull/12062))
* compute: added `resource_manager_tags` field to `google_compute_machine_image` resource ([#12051](https://github.com/hashicorp/terraform-provider-google-beta/pull/12051))
* container: added `node_config.linux_node_config.accurate_time_config` field to `google_container_node_pool` resource ([#12045](https://github.com/hashicorp/terraform-provider-google-beta/pull/12045))
* container: added `node_pool.node_config.linux_node_config.accurate_time_config` and `node_config.linux_node_config.accurate_time_config` fields to `google_container_cluster` resource ([#12045](https://github.com/hashicorp/terraform-provider-google-beta/pull/12045))
* container: added `node_pool.node_config.linux_node_config.swap_config` field to `google_container_node_pool` resource ([#11992](https://github.com/hashicorp/terraform-provider-google-beta/pull/11992))
* container: increased default timeout for `google_container_cluster` to 90 minutes (from 40/60 depending on operation) and `google_container_node_pool` to 60 minutes (from 30) ([#12071](https://github.com/hashicorp/terraform-provider-google-beta/pull/12071))
* discoveryengine: added `destionation_configs.destionations.port` and `destionation_configs.params` fields to `google_discovery_engine_data_connector ` resource ([#12039](https://github.com/hashicorp/terraform-provider-google-beta/pull/12039))
* dns: added support for IAM conditions to `google_dns_managed_zone` resource ([#12010](https://github.com/hashicorp/terraform-provider-google-beta/pull/12010))
* datastream: added `deletion_policy` field to control whether child routes are force-deleted to `google_datastream_private_connection` ([#12024](https://github.com/hashicorp/terraform-provider-google-beta/pull/12024))
* networkconnectivity: added support for IAM conditions to `google_network_connectivity_hub` resource ([#12005](https://github.com/hashicorp/terraform-provider-google-beta/pull/12005))
* networksecurity: added `parent` field to `google_network_security_address_groups` data source ([#12057](https://github.com/hashicorp/terraform-provider-google-beta/pull/12057))
* workbench: added support for new disk types and accelerators to `google_workbench_instance` ([#12042](https://github.com/hashicorp/terraform-provider-google-beta/pull/12042))

BUG FIXES:
* alloydb: fixed `google_alloydb_cluster` so that `maintenance_update_policy.maintenance_windows.start_time.hours` can be set to `0` (midnight) ([#11991](https://github.com/hashicorp/terraform-provider-google-beta/pull/11991))
* ces: fixed type mismatch in `google_ces_app` variable default value ([#12059](https://github.com/hashicorp/terraform-provider-google-beta/pull/12059))
* compute: fixed an issue where an erroneous error could occur for having an unset `zone` field in `google_compute_instance_template` ([#12052](https://github.com/hashicorp/terraform-provider-google-beta/pull/12052))
* compute: fixed permadiff for `iap.oauth2_client_id` in `google_compute_backend_service` and `google_compute_region_backend_service` when the API returns a single space ([#11990](https://github.com/hashicorp/terraform-provider-google-beta/pull/11990))
* container: fixed a permadiff in `google_container_cluster` where `database_encryption.state` returning `ALL_OBJECTS_ENCRYPTION_ENABLED` instead of the configured `ENCRYPTED` caused unintended reapplies ([#12029](https://github.com/hashicorp/terraform-provider-google-beta/pull/12029))
* dataplex: fixed acceptance test failure for one time scans ([#12065](https://github.com/hashicorp/terraform-provider-google-beta/pull/12065))
* dialogflowcx: fixed a perma-diff in `google_dialogflow_cx_test_case` when `session_parameters` was omitted from the configuration ([#11995](https://github.com/hashicorp/terraform-provider-google-beta/pull/11995))
* hypercomputecluster: fixed a permadiff in `google_hypercomputecluster_cluster` when `count`, `static_node_count`, or `max_dynamic_node_count` were explicitly set to `0`. ([#12049](https://github.com/hashicorp/terraform-provider-google-beta/pull/12049))
* identityplatform: fixed a premadiff on `multi_tenant` in `google_identity_platform_config` resource. Removing the value from config will now preserve the existing settings instead of removing them. ([#11996](https://github.com/hashicorp/terraform-provider-google-beta/pull/11996))
* memorystore: fixed an issue preventing updating multiple properties at once for `google_redis_cluster` ([#12053](https://github.com/hashicorp/terraform-provider-google-beta/pull/12053))

NOTES:
* compute: Migrate `resource_compute_instance_group.go.tmpl` resource to use direct HTTP rather then a client library ([#12056](https://github.com/hashicorp/terraform-provider-google-beta/pull/12056))
* compute: migrated `compute-operation` resource to use direct HTTP rather then a client library ([#12035](https://github.com/hashicorp/terraform-provider-google-beta/pull/12035))
* compute: migrated `compute_backend_bucket_security_policy` resource to use direct HTTP rather than a client library ([#12012](https://github.com/hashicorp/terraform-provider-google-beta/pull/12012))
* compute: migrated `compute_instance_network_interface_helpers` resource to use direct HTTP rather than a client library ([#12074](https://github.com/hashicorp/terraform-provider-google-beta/pull/12074))
* compute: migrated `data_source_google_compute_addresses.go.tmpl` data source to use direct HTTP rather then a client library ([#12015](https://github.com/hashicorp/terraform-provider-google-beta/pull/12015))
* compute: migrated `data_source_google_compute_machine_types` datasource to use direct HTTP rather than a client library ([#12016](https://github.com/hashicorp/terraform-provider-google-beta/pull/12016))
* compute: migrated `google_disk_test` to use direct HTTP rather than a client library ([#12055](https://github.com/hashicorp/terraform-provider-google-beta/pull/12055))
* compute: migrated `resource_compute_disk_async_replication` resource to use direct HTTP rather then a client library ([#12019](https://github.com/hashicorp/terraform-provider-google-beta/pull/12019))
* compute: migrated `resource_compute_http_health_check_test.go.tmpl` resource to use direct HTTP rather then a client library ([#12038](https://github.com/hashicorp/terraform-provider-google-beta/pull/12038))

## 7.29.0 (Apr 21, 2026)

NOTES:
* provider: List resources are now supported in both google and google-beta providers with the introduction of `google_service_account` list resource - more info can be found [here](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/using_list_resources_with_terraform_query) ([#11967](https://github.com/hashicorp/terraform-provider-google-beta/pull/11967))

FEATURES:
* **New Data Source:** `google_firebase_admin_sdk_config` ([#11954](https://github.com/hashicorp/terraform-provider-google-beta/pull/11954))
* **New Resource:** `google_firebase_storage_default_bucket` ([#11980](https://github.com/hashicorp/terraform-provider-google-beta/pull/11980))

IMPROVEMENTS:
* alloydb: added `track_client_address` field to `google_alloydb_instance` resource ([#11979](https://github.com/hashicorp/terraform-provider-google-beta/pull/11979))
* clouddeploy: added `tasks` field to `google_clouddeploy_custom_target_type` resource ([#11970](https://github.com/hashicorp/terraform-provider-google-beta/pull/11970))
* compute: added `header_action` and `redirect_options` fields to `google_compute_organization_security_policy_rule` resource ([#11971](https://github.com/hashicorp/terraform-provider-google-beta/pull/11971))
* dataplex: added `execution_identity` field to `google_dataplex_datascan` resource ([#11961](https://github.com/hashicorp/terraform-provider-google-beta/pull/11961))
* dataproc: added `cluster_config.engine` field to `google_dataproc_cluster` resource ([#11978](https://github.com/hashicorp/terraform-provider-google-beta/pull/11978))
* gkehub: added `ignored_clusters_selector` field to `google_gke_hub_rollout_sequence` resource ([#11963](https://github.com/hashicorp/terraform-provider-google-beta/pull/11963))
* iambeta: added `trust_default_shared_ca` field to `google_iam_workload_identity_pool` resource ([#11989](https://github.com/hashicorp/terraform-provider-google-beta/pull/11989))
* netapp: added `kms_config`, `encryption_state` and `backups_crypto_key_version` fields to `google_netapp_backup_vault` resource ([#11968](https://github.com/hashicorp/terraform-provider-google-beta/pull/11968))
* resourcemanager: add resource-identity support to `google_service_account` resource ([#11967](https://github.com/hashicorp/terraform-provider-google-beta/pull/11967))
* sql: added `entraid_config` field to `google_sql_database_instance` resource ([#11958](https://github.com/hashicorp/terraform-provider-google-beta/pull/11958))
* vectorsearch: added `encryption_spec` field to `google_vector_search_collection` resource ([#11987](https://github.com/hashicorp/terraform-provider-google-beta/pull/11987))

BUG FIXES:
* apigee: fixed ignoring `is_enabled = false` on create and update in `google_apigee_target_server` resource ([#11938](https://github.com/hashicorp/terraform-provider-google-beta/pull/11938))
* bigquery: fixed inability to set `default_collation` to empty string in `google_bigquery_dataset` ([#11962](https://github.com/hashicorp/terraform-provider-google-beta/pull/11962))
* ces: fixed a diff on `logging_settings` when unspecified in `google_ces_app`. Removing the value from config will now preserve the existing settings instead of removing them. ([#11952](https://github.com/hashicorp/terraform-provider-google-beta/pull/11952))
* compute: fixed a permadiff on `iap.oauth2_client_id` in `google_compute_backend_service` and `google_compute_region_backend_service` when the API returns a single space ([#11990](https://github.com/hashicorp/terraform-provider-google-beta/pull/11990))
* container: fixed a bug in `google_container_cluster` where setting multiple fields in `dns_endpoint_config` failed to apply all changes ([#11983](https://github.com/hashicorp/terraform-provider-google-beta/pull/11983))
* workstations: fixed a permadiff on `persistent_directories.gce_pd.reclaim_policy` in `google_workstations_workstation_config` resource ([#11986](https://github.com/hashicorp/terraform-provider-google-beta/pull/11986))




## 7.28.0 (Apr 14, 2026)

NOTES:
* compute: migrated `data_source_google_compute_instance_template` datasource to use direct HTTP rather then a client library ([#11901](https://github.com/hashicorp/terraform-provider-google-beta/pull/11901))
* compute: migrated `google_compute_instance_guest_attributes` datasource to use direct HTTP rather then a client library ([#11897](https://github.com/hashicorp/terraform-provider-google-beta/pull/11897))
* provider: added provider-wide `Identity()` schema support, allowing imports with MMv1 resources to occur using the identity block instead of id field ([#11874](https://github.com/hashicorp/terraform-provider-google-beta/pull/11874))

FEATURES:
* **New Data Source:** `google_vertex_ai_reasoning_engine_query` ([#11878](https://github.com/hashicorp/terraform-provider-google-beta/pull/11878))
* **New Resource:** `google_apigee_space` ([#11922](https://github.com/hashicorp/terraform-provider-google-beta/pull/11922))
* **New Resource:** `google_iap_agent_registry_iam_binding` ([#11905](https://github.com/hashicorp/terraform-provider-google-beta/pull/11905))
* **New Resource:** `google_iap_agent_registry_iam_member` ([#11905](https://github.com/hashicorp/terraform-provider-google-beta/pull/11905))
* **New Resource:** `google_iap_agent_registry_iam_policy` ([#11905](https://github.com/hashicorp/terraform-provider-google-beta/pull/11905))
* **New Resource:** `google_iap_location_web_iam_binding` ([#11905](https://github.com/hashicorp/terraform-provider-google-beta/pull/11905))
* **New Resource:** `google_iap_location_web_iam_member` ([#11905](https://github.com/hashicorp/terraform-provider-google-beta/pull/11905))
* **New Resource:** `google_iap_location_web_iam_policy` ([#11905](https://github.com/hashicorp/terraform-provider-google-beta/pull/11905))
* **New Resource:** `google_vertex_ai_reasoning_engine_iam_binding` ([#11876](https://github.com/hashicorp/terraform-provider-google-beta/pull/11876))
* **New Resource:** `google_vertex_ai_reasoning_engine_iam_member` ([#11876](https://github.com/hashicorp/terraform-provider-google-beta/pull/11876))
* **New Resource:** `google_vertex_ai_reasoning_engine_iam_policy` ([#11876](https://github.com/hashicorp/terraform-provider-google-beta/pull/11876))
* **New Resource:** `google_workload_identity_service_agent` ([#11871](https://github.com/hashicorp/terraform-provider-google-beta/pull/11871))

IMPROVEMENTS:
* composer: increased `google_composer_environment` default delete timeout to 120m from 30m ([#11917](https://github.com/hashicorp/terraform-provider-google-beta/pull/11917))
* compute: added `target_size_policy` field to `google_compute_instance_group_manager` and `google_compute_region_instance_group_manager` resources ([#11915](https://github.com/hashicorp/terraform-provider-google-beta/pull/11915))
* compute: increased `google_compute_security_policy` default timeout to 60m from 30m ([#11916](https://github.com/hashicorp/terraform-provider-google-beta/pull/11916))
* compute: supported simultaneous updates for Hyperdisk IOPS and throughput in `google_compute_disk` and `google_compute_region_disk` resources ([#11889](https://github.com/hashicorp/terraform-provider-google-beta/pull/11889))
* container: added `autopilot_cluster_policy_config` field to `google_container_cluster` resource ([#11893](https://github.com/hashicorp/terraform-provider-google-beta/pull/11893))
* container: added `disable_multi_nic` field to `lustre_csi_driver_config` in `google_container_cluster` resource ([#11864](https://github.com/hashicorp/terraform-provider-google-beta/pull/11864))
* container: added `managed_machine_learning_diagnostics_config` field to `google_container_cluster` resource ([#11927](https://github.com/hashicorp/terraform-provider-google-beta/pull/11927))
* dataform: added `disabled` field to `google_dataform_repository_release_config` resource ([#11921](https://github.com/hashicorp/terraform-provider-google-beta/pull/11921))
* developerconnect: added `custom_oauth_config`, `etag`, and `proxy_config` fields to `google_developer_connect_account_connector` resource ([#11858](https://github.com/hashicorp/terraform-provider-google-beta/pull/11858))
* netapp: added `scale_type` field to `google_netapp_storage_pool` resource ([#11892](https://github.com/hashicorp/terraform-provider-google-beta/pull/11892))
* networkservices: added `all_ports` field to `google_network_services_gateway` resource ([#11884](https://github.com/hashicorp/terraform-provider-google-beta/pull/11884))
* sql: added `SQLSERVER_2025` value to `database_version` field in `database_instance` resource ([#11911](https://github.com/hashicorp/terraform-provider-google-beta/pull/11911))
* vertexai: add `labels` field to `google_vertex_ai_reasoning_engine` resource ([#11896](https://github.com/hashicorp/terraform-provider-google-beta/pull/11896))
* vertexai: added `spec.source_code_spec.image_spec` field to `google_vertex_ai_reasoning_engine` resource ([#11881](https://github.com/hashicorp/terraform-provider-google-beta/pull/11881))
* vertexai: added `container_spec` field to `google_vertex_ai_reasoning_engine` resource ([#11888](https://github.com/hashicorp/terraform-provider-google-beta/pull/11888))

BUG FIXES:
* apigee: fixed a crash in `google_apigee_environment_addons_config` resource when analytics are not configured ([#11886](https://github.com/hashicorp/terraform-provider-google-beta/pull/11886))
* apigee: fixed overly restrictive validation of `name` field in `google_apigee_api_product` that rejected uppercase letters, aligning provider behavior with the Apigee API ([#11863](https://github.com/hashicorp/terraform-provider-google-beta/pull/11863))
* bigquery: fixed crash when `hive_partitioning_options` is defined with all null values in `google_bigquery_table` resource ([#11912](https://github.com/hashicorp/terraform-provider-google-beta/pull/11912))
* firebaseailogic: fixed permadiff on `traffic_filter` field in `google_firebase_ai_logic_config` resource ([#11857](https://github.com/hashicorp/terraform-provider-google-beta/pull/11857))
* networksecurity: fixed permadiff on `policy_profile` field in `google_network_security_authz_policy` resource ([#11929](https://github.com/hashicorp/terraform-provider-google-beta/pull/11929))
* vertexai: added 10-second wait before reading the updated resource in `google_vertex_ai_reasoning_engine`, preventing stale values getting written to state ([#11918](https://github.com/hashicorp/terraform-provider-google-beta/pull/11918))

## 7.27.0 (Apr 07, 2026)

BREAKING CHANGES:
* lustre: marked `maintenance_policy.weekly_maintenance_windows` field required in `google_lustre_instance` resource. Configuring `maintenance_policy` without `weekly_maintenance_windows` will cause an API error. ([#11855](https://github.com/hashicorp/terraform-provider-google-beta/pull/11855))

FEATURES:
* **New Data Source:** `google_discovery_engine_data_store` ([#11821](https://github.com/hashicorp/terraform-provider-google-beta/pull/11821))
* **New Data Source:** `google_discovery_engine_data_stores` ([#11821](https://github.com/hashicorp/terraform-provider-google-beta/pull/11821))
* **New Data Source:** `google_dns_record_sets` ([#11850](https://github.com/hashicorp/terraform-provider-google-beta/pull/11850))
* **New Resource:** `google_chronicle_dashboard_chart` ([#11841](https://github.com/hashicorp/terraform-provider-google-beta/pull/11841))
* **New Resource:** `google_chronicle_feed` ([#11856](https://github.com/hashicorp/terraform-provider-google-beta/pull/11856))
* **New Resource:** `google_network_connectivity_transport` ([#11810](https://github.com/hashicorp/terraform-provider-google-beta/pull/11810))
* **New Resource:** `google_organization_service_identity` ([#11844](https://github.com/hashicorp/terraform-provider-google-beta/pull/11844))

IMPROVEMENTS:
* compute: added `SEV_LIVE_MIGRATABLE_V2` to `guest_os_features` enum for `google_compute_region_disk` resource ([#11849](https://github.com/hashicorp/terraform-provider-google-beta/pull/11849))
* compute: added `SNP_SVSM_CAPABLE` to `guest_os_features` enum for `google_compute_image` and `google_compute_region_disk` resources ([#11849](https://github.com/hashicorp/terraform-provider-google-beta/pull/11849))
* compute: added `cache_policy` field to `google_compute_url_map` (beta) ([#11842](https://github.com/hashicorp/terraform-provider-google-beta/pull/11842))
* compute: added `excluded_folders` and `excluded_projects` fields  to `google_compute_organization_security_policy_association` resource ([#11835](https://github.com/hashicorp/terraform-provider-google-beta/pull/11835))
* compute: supported in-place update for `secondary_ip_range` field in `google_compute_subnetwork` resource ([#11830](https://github.com/hashicorp/terraform-provider-google-beta/pull/11830))
* container: added `autopilot_privileged_admission` field to `google_container_cluster` for Customer-Driven Allowlisting ([#11826](https://github.com/hashicorp/terraform-provider-google-beta/pull/11826))
* dataplex: added `aspects` field to `google_dataplex_entry_link` resource ([#11824](https://github.com/hashicorp/terraform-provider-google-beta/pull/11824))
* dataplex: supported in-place update for `aspects` field in `google_dataplex_entry_link` resource ([#11837](https://github.com/hashicorp/terraform-provider-google-beta/pull/11837))
* dataproc: added `boot_disk_provisioned_iops` and `boot_disk_provisioned_throughput` fields  to `cluster_config.worker_config.disk_config` in `google_dataproc_cluster` resource  ([#11832](https://github.com/hashicorp/terraform-provider-google-beta/pull/11832))
* dataproc: added value `AUTO` to `runtime_config.autotuning_config.scenarios` field in `google_dataproc_batch` resource ([#11818](https://github.com/hashicorp/terraform-provider-google-beta/pull/11818))
* iambeta: added `attestation_rules` field to `google_iam_workload_identity_pool` resource ([#11840](https://github.com/hashicorp/terraform-provider-google-beta/pull/11840))
* lustre: added `dynamic_tier_options` field to `google_lustre_instance` resource ([#11855](https://github.com/hashicorp/terraform-provider-google-beta/pull/11855))
* migrationcenter: added `virtual_machine_preferences.compute_engine_preferences.persistent_disk_type` field to `google_migration_center_preference_set` resource ([#11834](https://github.com/hashicorp/terraform-provider-google-beta/pull/11834))
* networkconnectivity: added `exclude_import_ranges`, `include_export_ranges`, `exclude_export_ranges` fields to `google_network_connectivity_spoke` resource ([#11845](https://github.com/hashicorp/terraform-provider-google-beta/pull/11845))
* pubsub: added `ai_inference` field to `google_pubsub_topic` and `google_pubsub_subscription` resources ([#11852](https://github.com/hashicorp/terraform-provider-google-beta/pull/11852))
* vertexai: added `context_spec` field to `google_vertex_ai_reasoning_engine` resource ([#11827](https://github.com/hashicorp/terraform-provider-google-beta/pull/11827))

BUG FIXES:
* compute: fixed a permadiff on the `adaptive_protection_config` field in `google_compute_security_policy` resource ([#11833](https://github.com/hashicorp/terraform-provider-google-beta/pull/11833))
* compute: fixed panic when setting `google_compute_project_metadata` on a project with no existing metadata ([#11812](https://github.com/hashicorp/terraform-provider-google-beta/pull/11812))
* biglakeiceberg: changed the `primary-location` parameter to `primary_location` in the create URL of `google_biglake_iceberg_catalog` resource ([#11836](https://github.com/hashicorp/terraform-provider-google-beta/pull/11836))
* securityposture: always sent value of `enforce` in `policies.constraint.org_policy_constraint.policy_rules` to the api  in `google_securityposture_posture` resource ([#11817](https://github.com/hashicorp/terraform-provider-google-beta/pull/11817))
* vertexai: fixed missing Private Service Connect service attachment for `service_attachment` field in `google_vertex_ai_endpoint_with_model_garden_deployment` resource ([#11831](https://github.com/hashicorp/terraform-provider-google-beta/pull/11831))
* workstations: fixed update of `private_cluster_config.allowed_projects` in `google_workstations_workstation_cluster` ([#11839](https://github.com/hashicorp/terraform-provider-google-beta/pull/11839))
1839))

## 7.26.0 (Mar 31, 2026)
FEATURES:
* **New Data Source:** `google_network_security_address_groups` ([#11773](https://github.com/hashicorp/terraform-provider-google-beta/pull/11773))
* **New Resource:** `google_bigqueryreservation_reservation_group` ([#11771](https://github.com/hashicorp/terraform-provider-google-beta/pull/11771))
* **New Resource:** `google_chronicle_datatable_row` ([#11769](https://github.com/hashicorp/terraform-provider-google-beta/pull/11769))
* **New Resource:** `google_chronicle_native_dashboard` ([#11806](https://github.com/hashicorp/terraform-provider-google-beta/pull/11806))
* **New Resource:** `google_contact_center_insights_assessment_rule` ([#11764](https://github.com/hashicorp/terraform-provider-google-beta/pull/11764))

IMPROVEMENTS:
* bigqueryreservation: added `reservation_group` field to `reservation` resource ([#11771](https://github.com/hashicorp/terraform-provider-google-beta/pull/11771))
* ces: added `remote_dialogflow_agent.respect_response_interruption_settings` field to `google_ces_agent` resource ([#11786](https://github.com/hashicorp/terraform-provider-google-beta/pull/11786))
* clusterdirector: made `boot_disk.size_gb` and `boot_disk.type` editable within nodesets and login nodes in `google_hypercomputecluster_cluster` ([#11808](https://github.com/hashicorp/terraform-provider-google-beta/pull/11808))
* colab: added `colab_image` field to `google_colab_runtime_template` resource ([#11790](https://github.com/hashicorp/terraform-provider-google-beta/pull/11790))
* colab: made `google_colab_runtime_template` resource updatable ([#11790](https://github.com/hashicorp/terraform-provider-google-beta/pull/11790))
* compute: added `hyperdisk-balanced` as an option for `disk_type` field in `google_container_cluster` resource ([#11789](https://github.com/hashicorp/terraform-provider-google-beta/pull/11789))
* compute: added `priority` field to `google_compute_region_network_firewall_policy_association` resource ([#11770](https://github.com/hashicorp/terraform-provider-google-beta/pull/11770))
* compute: added `load_balancing_scheme` field to `google_compute_target_tcp_proxy` resource ([#11759](https://github.com/hashicorp/terraform-provider-google-beta/pull/11759))
* compute: made `backend_service` field optional for `google_compute_target_tcp_proxy` resource ([#11759](https://github.com/hashicorp/terraform-provider-google-beta/pull/11759))
* iambeta: added `SYSTEM_TRUST_DOMAIN` as an additional enum value to `mode` field of `google_iam_workload_identity_pool` resource ([#11779](https://github.com/hashicorp/terraform-provider-google-beta/pull/11779))
* iambeta: added `use_default_shared_ca` field to `google_iam_workload_identity_pool` resource ([#11791](https://github.com/hashicorp/terraform-provider-google-beta/pull/11791))
* spanner: added autoscaling config for instance partition and missing asymmetric autoscaling override fields to `google_spanner_instance` resource ([#11785](https://github.com/hashicorp/terraform-provider-google-beta/pull/11785))
* sql: added `server_certificate_rotation_mode` field to `google_sql_database_instance` resource ([#11783](https://github.com/hashicorp/terraform-provider-google-beta/pull/11783))
* storage: added `google_managed_encryption_enforcement_config`, `customer_managed_encryption_enforcement_config` and `customer_supplied_encryption_enforcement_config` to `google_storage_bucket` resource ([#11763](https://github.com/hashicorp/terraform-provider-google-beta/pull/11763))

BUG FIXES:
* alloydb: fixed an issue where `password_wo` and `password_wo_version` fields were not functioning properly during update requests in `google_alloydb_user` resource ([#11782](https://github.com/hashicorp/terraform-provider-google-beta/pull/11782))
* biglake: fixed erroneous diff for the `properties` field in the `google_biglake_iceberg_table` and `google_biglake_iceberg_namespace` resources ([#11799](https://github.com/hashicorp/terraform-provider-google-beta/pull/11799))
* cloudfunctionsv2: fixed validation to only allow one of `direct_vpc_network_interface` or `vpc_connector` on `google_cloudfunctions2_function` resource ([#11778](https://github.com/hashicorp/terraform-provider-google-beta/pull/11778))
* cloudrunv2: fixed validation to only allow one of `network_interfaces` or `connector` on `google_cloud_run_v2_service` and `google_cloud_run_v2_job` resources ([#11778](https://github.com/hashicorp/terraform-provider-google-beta/pull/11778))
* compute: fixed invalid update mask used for `rate_limit_options` field in `google_compute_region_security_policy_rule` resource ([#11761](https://github.com/hashicorp/terraform-provider-google-beta/pull/11761))
* compute: fixed invalid update mask used for `rate_limit_options` field in `google_compute_security_policy` and `google_compute_security_policy_rule` resources ([#11760](https://github.com/hashicorp/terraform-provider-google-beta/pull/11760))
* iambeta: fixed a perma-diff on `mode` field for `google_iam_workload_identity_pool` resource ([#11805](https://github.com/hashicorp/terraform-provider-google-beta/pull/11805))
* provider: fixed an issue when custom endpoints use `http://` ([#11804](https://github.com/hashicorp/terraform-provider-google-beta/pull/11804))
* vertexai: fixed operation calls in `google_vertex_ai_` resources not respecting `universe_domain` and `vertex_custom_endpoint` ([#11768](https://github.com/hashicorp/terraform-provider-google-beta/pull/11768))

## 7.25.0 (Mar 24, 2026)

FEATURES:
* **New Data Source:** `google_compute_network_endpoint_groups` ([#11756](https://github.com/hashicorp/terraform-provider-google-beta/pull/11756))
* **New Resource:** `google_dialogflow_environment` ([#11737](https://github.com/hashicorp/terraform-provider-google-beta/pull/11737))
* **New Resource:** `google_kms_project_autokey_config` ([#11746](https://github.com/hashicorp/terraform-provider-google-beta/pull/11746))

IMPROVEMENTS:
* backupdr: added `disk_backup_plan_properties` field to `google_backup_dr_backup_plan` resource ([#11742](https://github.com/hashicorp/terraform-provider-google-beta/pull/11742))
* backupdr: made `backup_rules` optional in `google_backup_dr_backup_plan` resource ([#11741](https://github.com/hashicorp/terraform-provider-google-beta/pull/11741))
* blockchainnodeengine: added `ethereum_details.validator_config.beacon_fee_recipient` field to `google_blockchain_node_engine_blockchain_nodes` resource ([#11744](https://github.com/hashicorp/terraform-provider-google-beta/pull/11744))
* ces: added `custom_headers` field to MCP toolset in CES `google_ces_toolset` resource ([#11727](https://github.com/hashicorp/terraform-provider-google-beta/pull/11727))
* compute: added `forward_proxy` field to `google_compute_region_backend_service` resource ([#11723](https://github.com/hashicorp/terraform-provider-google-beta/pull/11723))
* compute: added `expr` field  to `google_compute_organization_security_policy_rule` resource ([#11749](https://github.com/hashicorp/terraform-provider-google-beta/pull/11749))
* compute: added `load_balancing_scheme` field to `google_compute_target_tcp_proxy` resource ([#11759](https://github.com/hashicorp/terraform-provider-google-beta/pull/11759))
* compute: added `load_balancing_scheme` field to `google_compute_region_target_tcp_proxy` resource ([#11740](https://github.com/hashicorp/terraform-provider-google-beta/pull/11740))
* compute: added `location` field to `google_network_services_tls_route` resource ([#11755](https://github.com/hashicorp/terraform-provider-google-beta/pull/11755))
* compute: added `target_proxies` field to `google_network_services_tls_route` resource ([#11757](https://github.com/hashicorp/terraform-provider-google-beta/pull/11757))
* compute: made `backend_service` field optional for `google_compute_target_tcp_proxy` resource ([#11759](https://github.com/hashicorp/terraform-provider-google-beta/pull/11759))
* compute: made `backend_service` field optional for `google_compute_region_target_tcp_proxy` resource ([#11740](https://github.com/hashicorp/terraform-provider-google-beta/pull/11740))
* dialogflow: added `use_bidi_streaming` field to `google_dialogflow_conversation_profile` resource ([#11736](https://github.com/hashicorp/terraform-provider-google-beta/pull/11736))
* firebaseailogic: added `traffic_filter` field to `google_firebase_ai_logic_config` resource ([#11738](https://github.com/hashicorp/terraform-provider-google-beta/pull/11738))
* iamworkforcepool: added `detailed_audit_logging` field to `google_iam_workforce_pool_provider` resource ([#11745](https://github.com/hashicorp/terraform-provider-google-beta/pull/11745))
* kms: added `key_project_resolution_mode` field to `google_kms_autokey_config` resource ([#11746](https://github.com/hashicorp/terraform-provider-google-beta/pull/11746))
* lustre: added `maintenance_policy` field to `google_lustre_instance` resource ([#11753](https://github.com/hashicorp/terraform-provider-google-beta/pull/11753))
* securesourcemanager: added `private_config.custom_host_config` field to `google_secure_source_manager_instance` resource ([#11747](https://github.com/hashicorp/terraform-provider-google-beta/pull/11747))
* sql: added `point_in_time_restore_context.region` field to `google_sql_database_instance` resource ([#11751](https://github.com/hashicorp/terraform-provider-google-beta/pull/11751))
* vertexai: added `deletion_policy` field to `resource_vertex_ai_reasoning_engine` resource ([#11758](https://github.com/hashicorp/terraform-provider-google-beta/pull/11758))
* vertexai: added `spec.identity_type` and `spec.effective_identity` fields to `google_vertex_ai_reasoning_engine` resource ([#11724](https://github.com/hashicorp/terraform-provider-google-beta/pull/11724))

BUG FIXES:
* vertexai: fixed permadiff on `spec` field in `google_vertex_ai_reasoning_engine` resource ([#11724](https://github.com/hashicorp/terraform-provider-google-beta/pull/11724))

## 7.24.0 (Mar 17, 2026)

DEPRECATIONS:
* iamworkforcepool: deprecated `extended_attributes_oauth2_client` on `google_iam_workforce_pool_provider`. Use `scim_usage` instead. ([#11688](https://github.com/hashicorp/terraform-provider-google-beta/pull/11688))

FEATURES:
* **New Resource:** `google_biglake_iceberg_table` ([#11689](https://github.com/hashicorp/terraform-provider-google-beta/pull/11689))
* **New Resource:** `google_chronicle_datatable` ([#11708](https://github.com/hashicorp/terraform-provider-google-beta/pull/11708))
* **New Resource:** `google_contact_center_insights_auto_labeling_rule` ([#11707](https://github.com/hashicorp/terraform-provider-google-beta/pull/11707))
* **New Resource:** `google_network_security_ull_mirroring_collector` ([#11687](https://github.com/hashicorp/terraform-provider-google-beta/pull/11687))
* **New Resource:** `google_network_security_ull_mirroring_engine` ([#11687](https://github.com/hashicorp/terraform-provider-google-beta/pull/11687))
* **New Resource:** `google_sql_provision_script` ([#11711](https://github.com/hashicorp/terraform-provider-google-beta/pull/11711))

IMPROVEMENTS:
* ces: added Service Account OAuth `scopes` fields to `google_ces_toolset` resource ([#11679](https://github.com/hashicorp/terraform-provider-google-beta/pull/11679))
* cloudrunv2: added `DISK` fields to `google_cloud_run_v2_service` resource ([#11705](https://github.com/hashicorp/terraform-provider-google-beta/pull/11705))
* cloudsql: added `max_custom_on_demand_retention_days` field to `sqladmin` resource ([#11700](https://github.com/hashicorp/terraform-provider-google-beta/pull/11700))
* compute: added `forward_proxy` field in `google_compute_region_backend_service` resource ([#11723](https://github.com/hashicorp/terraform-provider-google-beta/pull/11723))
* compute: added `accelerator_topology_mode` field to `google_compute_resource_policy` resource ([#11683](https://github.com/hashicorp/terraform-provider-google-beta/pull/11683))
* compute: added `target_type` and `target_forwarding_rules` on `google_compute_region_network_firewall_policy_rule` resource ([#11680](https://github.com/hashicorp/terraform-provider-google-beta/pull/11680))
* container: marked `subnetwork` as settable in `google_container_node_pool` ([#11703](https://github.com/hashicorp/terraform-provider-google-beta/pull/11703))
* container: added `disruption_budget` field to `google_container_cluster` resource ([#11706](https://github.com/hashicorp/terraform-provider-google-beta/pull/11706))
* discoveryengine: added `search_engine_config.required_subscription_tier ` field to `google_discovery_engine_search_engine` resource ([#11692](https://github.com/hashicorp/terraform-provider-google-beta/pull/11692))
* discoveryengine: marked `content_config` as optional field in `google_discovery_engine_data_store` ([#11692](https://github.com/hashicorp/terraform-provider-google-beta/pull/11692))
* memorystore: added `server_ca_mode` and `server_ca_pool` fields to `google_memorystore_instance` resource ([#11716](https://github.com/hashicorp/terraform-provider-google-beta/pull/11716))
* networkservices: relaxed `authority` validation in `google_network_services_authz_extension` for different target types ([#11686](https://github.com/hashicorp/terraform-provider-google-beta/pull/11686))
* redis: added `server_ca_mode` and `server_ca_pool` fields to `google_redis_cluster` resource ([#11716](https://github.com/hashicorp/terraform-provider-google-beta/pull/11716))
* sql: added `clone_context.source_project` field to `google_sql_database_instance` resource to support cross project clone (beta) ([#11684](https://github.com/hashicorp/terraform-provider-google-beta/pull/11684))
* transport: added automatic retry for GCE 403 errors with reason `CONCURRENT_OPERATIONS_QUOTA_EXCEEDED` ([#11704](https://github.com/hashicorp/terraform-provider-google-beta/pull/11704))

BUG FIXES:
* compute: fixed perpetual diff for `oauth2_client_id` in `iap` block of `google_compute_backend_service` and `google_compute_region_backend_service` when disabling IAP ([#11685](https://github.com/hashicorp/terraform-provider-google-beta/pull/11685))
* datastream: fixed an issue in `google_datastream_stream` where `source_config.mysql_source_config.binary_log_position` would show a diff when unset ([#11714](https://github.com/hashicorp/terraform-provider-google-beta/pull/11714))
* workbench: marked `install-nvidia-driver` metadata key as settable for `google_workbench_instance` ([#11696](https://github.com/hashicorp/terraform-provider-google-beta/pull/11696))

## 7.23.0 (Mar 10, 2026)

DEPRECATIONS:
* notebooks: `google_notebooks_environment` is deprecated and will be removed in a future major release. Use `google_workbench_instance` instead ([#11646](https://github.com/hashicorp/terraform-provider-google-beta/pull/11646))
* provider: `google_*_iam_*` resources and datasources will now show deprecation messages when their parent resource has been deprecated ([#11646](https://github.com/hashicorp/terraform-provider-google-beta/pull/11646))

FEATURES:
* **New Data Source:** `google_observability_folder_settings` ([#11668](https://github.com/hashicorp/terraform-provider-google-beta/pull/11668))
* **New Data Source:** `google_observability_organization_settings` ([#11649](https://github.com/hashicorp/terraform-provider-google-beta/pull/11649))
* **New Data Source:** `google_observability_project_settings` ([#11663](https://github.com/hashicorp/terraform-provider-google-beta/pull/11663))
* **New Data Source:** `google_oracle_database_odb_network` ([#11648](https://github.com/hashicorp/terraform-provider-google-beta/pull/11648))
* **New Data Source:** `google_oracle_database_odb_subnet` ([#11648](https://github.com/hashicorp/terraform-provider-google-beta/pull/11648))
* **New Resource:** `google_compute_router_named_set` ([#11667](https://github.com/hashicorp/terraform-provider-google-beta/pull/11667))
* **New Resource:** `google_dataform_config` ([#11662](https://github.com/hashicorp/terraform-provider-google-beta/pull/11662))
* **New Resource:** `google_observability_folder_settings` ([#11668](https://github.com/hashicorp/terraform-provider-google-beta/pull/11668))
* **New Resource:** `google_observability_organization_settings` ([#11649](https://github.com/hashicorp/terraform-provider-google-beta/pull/11649))
* **New Resource:** `google_observability_project_settings` ([#11663](https://github.com/hashicorp/terraform-provider-google-beta/pull/11663))

IMPROVEMENTS:
* alloydb: added `dataplex_config` field to `google_alloydb_cluster` resource ([#11651](https://github.com/hashicorp/terraform-provider-google-beta/pull/11651))
* biglake: added `primary_location` to `google_biglake_iceberg_catalog` resource ([#11654](https://github.com/hashicorp/terraform-provider-google-beta/pull/11654))
* compute: added `preemption_notice_duration` field to `google_compute_instance`, `google_compute_instance_template`, and `google_compute_instance_from_machine_image` resources ([#11655](https://github.com/hashicorp/terraform-provider-google-beta/pull/11655))
* compute: added `storage_pool` support to `google_compute_instance_template` and `google_compute_region_instance_template` disks ([#11671](https://github.com/hashicorp/terraform-provider-google-beta/pull/11671))
* container: added `control_plane_disk_encryption_key_versions` field to `user_managed_keys_config` in `google_container_cluster` resource ([#11647](https://github.com/hashicorp/terraform-provider-google-beta/pull/11647))
* dataproc: added `cluster_type` to `google_dataproc_cluster` resource ([#11673](https://github.com/hashicorp/terraform-provider-google-beta/pull/11673))
* discoveryengine: added `action_config`, `bap_config`, `destination_configs`, and `data_source_version` to `google_discovery_engine_data_connector` resource ([#11645](https://github.com/hashicorp/terraform-provider-google-beta/pull/11645))
* dlp: added `actions.publish_to_scc`, `actions.publish_to_chronicle`, `actions.export_data.sample_findings_table` and `targets.big_query_target.filter.table_reference.project_id` fields to `google_data_loss_prevention_discovery_config` resource ([#11639](https://github.com/hashicorp/terraform-provider-google-beta/pull/11639))
* gkebackup: added `protected_namespace_count` field to `google_gke_backup_backup_plan` resource ([#11641](https://github.com/hashicorp/terraform-provider-google-beta/pull/11641))
* netapp: added `mode` field to `google_netapp_storage_pool` resource ([#11656](https://github.com/hashicorp/terraform-provider-google-beta/pull/11656))
* osconfig: added `patch_config.skip_unpatchable_vms` field to `google_os_config_patch_deployment` resource ([#11640](https://github.com/hashicorp/terraform-provider-google-beta/pull/11640))
* pubsub: added `text_config` field to `google_pubsub_subscription` resource ([#11660](https://github.com/hashicorp/terraform-provider-google-beta/pull/11660))

BUG FIXES:
* tags: fixed iam read-after-write consistency issue with conditions in `google_tags_tag_key_iam_member` resource ([#11661](https://github.com/hashicorp/terraform-provider-google-beta/pull/11661))

## 7.22.0 (Mar 3, 2026)

DEPRECATIONS:
* dataplex: deprecated `google_dataplex_data_asset`. Use `google_dataplex_data_product_data_asset` instead. ([#26256](https://github.com/hashicorp/terraform-provider-google/pull/26256))

FEATURES:
* **New Resource:** `google_hypercomputecluster_cluster` ([#26180](https://github.com/hashicorp/terraform-provider-google/pull/26180))

IMPROVEMENTS:
* compute: `initialize_params.size` is now updatable in-place in the `google_compute_instance` resource ([#26195](https://github.com/hashicorp/terraform-provider-google/pull/26195))
* compute: added `dest_network_context`, `src_network_context` and `src_networks` fields to `google_compute_firewall_policy_rule` resource ([#26227](https://github.com/hashicorp/terraform-provider-google/pull/26227))
* compute: added `dest_network_context`, `src_network_context` and `src_networks` fields to `google_compute_network_firewall_policy_rule` resource ([#26227](https://github.com/hashicorp/terraform-provider-google/pull/26227))
* compute: added `dest_network_context`, `src_network_context` and `src_networks` fields to `google_compute_region_network_firewall_policy_rule` resource ([#26227](https://github.com/hashicorp/terraform-provider-google/pull/26227))
* compute: added `consumer_accept_lists.endpoint_url` to `google_compute_service_attachment` resource ([#26234](https://github.com/hashicorp/terraform-provider-google/pull/26234))
* compute: added `erase_windows_vss_signature` field to `google_compute_disk` resource ([#26249](https://github.com/hashicorp/terraform-provider-google/pull/26249))
* compute: added `erase_windows_vss_signature` field to `google_compute_region_disk` resource ([#26249](https://github.com/hashicorp/terraform-provider-google/pull/26249))
* developerconnect: added `http_config` field to `google_developer_connect_connection` resource ([#26232](https://github.com/hashicorp/terraform-provider-google/pull/26232))
* filestore: added `source_backupdr_backup` field to `google_filestore_instance` resource ([#26238](https://github.com/hashicorp/terraform-provider-google/pull/26238))
* gkehub2: added field `spec.workloadidentity` to resource `google_gke_hub_feature` ([#26259](https://github.com/hashicorp/terraform-provider-google/pull/26259))
* iam: added AZURE_AD_GROUPS_DISPLAY_NAME enum value to `extra_attributes_oauth2_client.attribute-type` field in `google_iam_workforce_pool_provider` resource ([#26226](https://github.com/hashicorp/terraform-provider-google/pull/26226))
* kms: added a KMS AutokeyConfig-specific 10s post-create/post-update ([#26236](https://github.com/hashicorp/terraform-provider-google/pull/26236))
* netapp: added `large_capacity_config` block with `constituent_count` field to `google_netapp_volume` resource ([#26243](https://github.com/hashicorp/terraform-provider-google/pull/26243))
* netapp: added `scale_tier` field to `netapp_storage_pool` resource ([#26243](https://github.com/hashicorp/terraform-provider-google/pull/26243))
* networkservices: added support for use of multiple `ports` for `google_network_services_gateway` resources of type `SECURE_WEB_GATEWAY` ([#26265](https://github.com/hashicorp/terraform-provider-google/pull/26265))
* sql: added `auto_upgrade_enabled` field to `google_sql_database_instance` resource. ([#26205](https://github.com/hashicorp/terraform-provider-google/pull/26205))
* sql: added `data_api_access` field to `google_sql_database_instance` resource ([#26217](https://github.com/hashicorp/terraform-provider-google/pull/26217))
* sql: added `enhanced_query_insights_enabled` field to `google_sql_database_instance` resource ([#26244](https://github.com/hashicorp/terraform-provider-google/pull/26244))

BUG FIXES:
* datastream: fixed permadiff where `google_datastream_connection_profile.salesforce_profile.oauth2_client_credentials.client_id` is not read properly from the API ([#26201](https://github.com/hashicorp/terraform-provider-google/pull/26201))
* servicenetworking: added retry when creating `google_service_networking_connection` if it looks like the service account permissions haven't yet propagated ([#26220](https://github.com/hashicorp/terraform-provider-google/pull/26220))

## 7.21.0 (Feb 24, 2026)

FEATURES:
* **New Data Source:** `google_vmwareengine_announcements` ([#11577](https://github.com/hashicorp/terraform-provider-google-beta/pull/11577))
* **New Data Source:** `google_vmwareengine_upgrades` ([#11591](https://github.com/hashicorp/terraform-provider-google-beta/pull/11591))
* **New Resource:** `google_compute_region_backend_bucket` ([#11576](https://github.com/hashicorp/terraform-provider-google-beta/pull/11576))
* **New Resource:** `google_dataform_team_folder` (beta) ([#11573](https://github.com/hashicorp/terraform-provider-google-beta/pull/11573))
* **New Resource:** `google_hypercomputecluster_cluster` ([#11596](https://github.com/hashicorp/terraform-provider-google-beta/pull/11596))
* **New Resource:** `google_network_services_agent_gateway` (beta) ([#11572](https://github.com/hashicorp/terraform-provider-google-beta/pull/11572))

IMPROVEMENTS:
* beyondcorp: added `logging` field to `google_beyondcorp_security_gateway` resource ([#11581](https://github.com/hashicorp/terraform-provider-google-beta/pull/11581))
* dataproc: added `wait_for_completion` field to `google_dataproc_job` resource ([#11593](https://github.com/hashicorp/terraform-provider-google-beta/pull/11593))
* discoveryengine: added `disable_analytics` field to `google_discovery_engine_search_engine` resource ([#11589](https://github.com/hashicorp/terraform-provider-google-beta/pull/11589))
* dlp: added `targets.cloud_storage_target.filter.collection.include_tags` block to `google_data_loss_prevention_discovery_config` resource ([#11594](https://github.com/hashicorp/terraform-provider-google-beta/pull/11594))
* iap: added `client_id`, `client_secret`, and `client_secret_sha256` fields to `google_iap_settings` resource ([#11588](https://github.com/hashicorp/terraform-provider-google-beta/pull/11588))

BUG FIXES:
* cloudrun: fixed perma-diff on `http_target.uri_override.query_override` in `google_cloud_tasks_queue` ([#11590](https://github.com/hashicorp/terraform-provider-google-beta/pull/11590))
* storage: fixed a bug in `google_storage_bucket` where `force_destroy = true` would fail to delete buckets with large number of objects due to missing pagination ([#11586](https://github.com/hashicorp/terraform-provider-google-beta/pull/11586))

## 7.20.0 (Feb 17, 2026)

FEATURES:
* **New Data Source:** `google_access_context_manager_supported_service` ([#11561](https://github.com/hashicorp/terraform-provider-google-beta/pull/11561))
* **New Data Source:** `google_access_context_manager_supported_services` ([#11561](https://github.com/hashicorp/terraform-provider-google-beta/pull/11561))
* **New Data Source:** `google_backup_dr_data_sources` ([#11551](https://github.com/hashicorp/terraform-provider-google-beta/pull/11551))
* **New Data Source:** `google_storage_bucket_object_contents` ([#11529](https://github.com/hashicorp/terraform-provider-google-beta/pull/11529))
* **New Resource:** `google_biglake_iceberg_namespace` ([#11547](https://github.com/hashicorp/terraform-provider-google-beta/pull/11547))
* **New Resource:** `google_cloud_identity_policy` ([#11536](https://github.com/hashicorp/terraform-provider-google-beta/pull/11536))
* **New Resource:** `google_compute_rollout_plan` ([#11562](https://github.com/hashicorp/terraform-provider-google-beta/pull/11562))
* **New Resource:** `google_dataform_folder` ([#11542](https://github.com/hashicorp/terraform-provider-google-beta/pull/11542))
* **New Resource:** `google_network_services_agent_gateway` (beta) ([#11571](https://github.com/hashicorp/terraform-provider-google-beta/pull/11571))
* **New Resource:** `google_oracle_database_exadb_vm_cluster` ([#11508](https://github.com/hashicorp/terraform-provider-google-beta/pull/11508))
* **New Resource:** `google_vector_search_collection` ([#11566](https://github.com/hashicorp/terraform-provider-google-beta/pull/11566))

IMPROVEMENTS:
* alloydb: added write-only support for `initial_user.password_wo` to `google_alloydb_cluster` ([#11546](https://github.com/hashicorp/terraform-provider-google-beta/pull/11546))
* ces: added `mcp_toolset` field to `google_ces_toolset` resource ([#11513](https://github.com/hashicorp/terraform-provider-google-beta/pull/11513))
* compute: added write-only support for `private_key` to `google_compute_region_ssl_certificate` resource ([#11544](https://github.com/hashicorp/terraform-provider-google-beta/pull/11544))
* compute: added write-only support for `private_key` to `google_compute_ssl_certificate` resource ([#11544](https://github.com/hashicorp/terraform-provider-google-beta/pull/11544))
* compute: added `enable` field to `google_compute_packet_mirroring` resource ([#11537](https://github.com/hashicorp/terraform-provider-google-beta/pull/11537))
* compute: added `params` field to `google_compute_external_vpn_gateway` resource ([#11558](https://github.com/hashicorp/terraform-provider-google-beta/pull/11558))
* compute: added `params` field to `google_compute_ha_vpn_gateway` resource ([#11558](https://github.com/hashicorp/terraform-provider-google-beta/pull/11558))
* compute: added `params` field to `google_compute_interconnect_attachment` resource ([#11518](https://github.com/hashicorp/terraform-provider-google-beta/pull/11518))
* compute: added `params` field to `google_compute_vpn_gateway` resource ([#11558](https://github.com/hashicorp/terraform-provider-google-beta/pull/11558))
* compute: added `params` field to `google_compute_vpn_tunnel` resource ([#11558](https://github.com/hashicorp/terraform-provider-google-beta/pull/11558))
* compute: added `slice_controller_config` field to `google_container_cluster` resource ([#11510](https://github.com/hashicorp/terraform-provider-google-beta/pull/11510))
* container: added `additional_ip_ranges_config.status` to `google_container_cluster` resource ([#11534](https://github.com/hashicorp/terraform-provider-google-beta/pull/11534))
* dataproc: added `instance_flexibility_policy` to `master_config` and `worker_config` in `google_dataproc_cluster` resource ([#11532](https://github.com/hashicorp/terraform-provider-google-beta/pull/11532))
* developerconnect: added `target_projects` field to `google_developer_connect_insights_config` resource ([#11545](https://github.com/hashicorp/terraform-provider-google-beta/pull/11545))
* filestore: added `replica_action` to `google_filestore_instance` resource ([#11553](https://github.com/hashicorp/terraform-provider-google-beta/pull/11553))
* networksecurity: added `policy_profile`, `http_rules.0.to.0.operations.0.mcp` to `google_network_security_authz_policy` ([#11559](https://github.com/hashicorp/terraform-provider-google-beta/pull/11559))
* networkservices: added `ull_multicast_domain` field to `google_network_services_multicast_domain` resource ([#11543](https://github.com/hashicorp/terraform-provider-google-beta/pull/11543))
* networkservices: relaxed `load_balancing_scheme` validation to support non-Backend Service targets in `google_network_services_authz_extension` ([#11559](https://github.com/hashicorp/terraform-provider-google-beta/pull/11559))
* spanner: added support for `user_project_override` in `google_spanner_database_iam` and `google_spanner_instance_iam` resources ([#11527](https://github.com/hashicorp/terraform-provider-google-beta/pull/11527))
* vmwareengine: added `datastore_mount_config` field to `google_vmwareengine_cluster` resource ([#11554](https://github.com/hashicorp/terraform-provider-google-beta/pull/11554))

BUG FIXES:
* bigquery: fixed permadiff with the `collation` field in `google_bigquery_table.schema` when it inherits the value from `google_bigquery_dataset.default_collation` ([#11538](https://github.com/hashicorp/terraform-provider-google-beta/pull/11538))
* bigqueryanalyticshub: fixed update failure for `replica_locations` in `google_bigquery_analytics_hub_listing` (beta) ([#11523](https://github.com/hashicorp/terraform-provider-google-beta/pull/11523))
* iam: fixed an issue where iam resources not retry on error 409 concurrent policy changes ([#11564](https://github.com/hashicorp/terraform-provider-google-beta/pull/11564))
* publicca: fixed `mac_key` fields not being properly set in `google_public_ca_external_account_key` ([#11567](https://github.com/hashicorp/terraform-provider-google-beta/pull/11567))

## 7.19.0 (Feb 10, 2026)

DEPRECATIONS:
* backupdr: `google_backupdr_restore_workload.name` is deprecated and will be removed in a future major release. The backup is identified by the parameters (location, backup_vault_id, data_source_id, backup_id). ([#11490](https://github.com/hashicorp/terraform-provider-google-beta/pull/11490))
* publicca: `google_public_ca_external_account_key.b64url_mac_key` is deprecated and will be removed in a future major release. Use `mac_key` instead. ([#11482](https://github.com/hashicorp/terraform-provider-google-beta/pull/11482))

IMPROVEMENTS:
* compute: added `RDMA_FALCON_POLICY` and `ULL_POLICY` values to `policy_type` field in `google_compute_region_network_firewall_policy`, `google_compute_region_network_firewall_policy_with_rules` ([#11489](https://github.com/hashicorp/terraform-provider-google-beta/pull/11489))
* compute: added support for `network_interface.network_attachment` to `google_compute_instance_template` (GA) ([#11496](https://github.com/hashicorp/terraform-provider-google-beta/pull/11496))
* compute: added support for `network_interface.network_attachment` to `google_compute_region_instance_template` ([#11496](https://github.com/hashicorp/terraform-provider-google-beta/pull/11496))
* compute: added support for `network_interface.vlan` to `google_compute_instance_template`, enabling dynamic NIC ([#11496](https://github.com/hashicorp/terraform-provider-google-beta/pull/11496))
* compute: added support for `network_interface.vlan` to `google_compute_instance`, enabling dynamic NIC. Creating and deleting from an existing instance is not yet supported. ([#11496](https://github.com/hashicorp/terraform-provider-google-beta/pull/11496))
* compute: added support for `network_interface.vlan` to `google_compute_region_instance_template`, enabling dynamic NIC ([#11496](https://github.com/hashicorp/terraform-provider-google-beta/pull/11496))
* discoveryengine: added `knowledge_graph_config` field to `google_discovery_engine_search_engine` resource ([#11484](https://github.com/hashicorp/terraform-provider-google-beta/pull/11484))
* firestore: added `firestore_data_access_mode`, `mongodb_compatible_data_acess_mode`, and `realtime_updates_mode` fields to the `google_firestore_database` resource ([#11497](https://github.com/hashicorp/terraform-provider-google-beta/pull/11497))
* firestore: added `deletion_policy` virtual field to `google_firestore_index` resource ([#11488](https://github.com/hashicorp/terraform-provider-google-beta/pull/11488))
* monitoring: added write-only variants (`auth_token_wo` + `auth_token_wo_version`, `password_wo` + `password_wo_version`, `service_key_wo` + `service_key_wo_version`) for `google_monitoring_notification_channel.sensitive_labels` ([#11487](https://github.com/hashicorp/terraform-provider-google-beta/pull/11487))
* networkconnectivity: added support for update operation on `google_network_connectivity_gateway_advertised_route` resource ([#11481](https://github.com/hashicorp/terraform-provider-google-beta/pull/11481))
* provider: added a configurable `poll_interval` field to the provider for rare cases where it is being used in latency-sensitive situations. This can be set to a custom duration to change operation polling intervals. The default is unchanged, at `10s`. ([#11500](https://github.com/hashicorp/terraform-provider-google-beta/pull/11500))
* publicca: added `mac_key` to `google_public_ca_external_account_key` ([#11482](https://github.com/hashicorp/terraform-provider-google-beta/pull/11482))
* run: added `readiness_probe` field to `google_cloud_run_v2_service` resource ([#11498](https://github.com/hashicorp/terraform-provider-google-beta/pull/11498))
* vertexai: added support for `developer_connect_source` to `spec.source_code_spec` in `google_vertex_ai_reasoning_engine` ([#11502](https://github.com/hashicorp/terraform-provider-google-beta/pull/11502))

BUG FIXES:
* compute: fixed issue where it wasn't possible to set both `ssl_certificates` and `certificate_map` in `google_compute_target_ssl_proxy` ([#11503](https://github.com/hashicorp/terraform-provider-google-beta/pull/11503))
* container: fixed an issue when toggling `default_compute_class_enabled` in `google_container_cluster` with Autopilot enabled ([#11483](https://github.com/hashicorp/terraform-provider-google-beta/pull/11483))
* firebaseailogic: fixed bug in `google_firebase_ai_logic_config.generative_language_config.api_key_wo` where the value set wouldn't be sent to the API. ([#11487](https://github.com/hashicorp/terraform-provider-google-beta/pull/11487))
* publicca: fixed `b64url_mac_key` sometimes being empty in `google_public_ca_external_account_key` ([#11482](https://github.com/hashicorp/terraform-provider-google-beta/pull/11482))

## 7.18.0 (Feb 3, 2026)

BREAKING CHANGES:
* alloydb: removed the incorrect top-level field `last_successful_backup_consistency_time` from `google_backup_dr_backup_plan_association`. No value has been present in this output-only field. ([#11473](https://github.com/hashicorp/terraform-provider-google-beta/pull/11473))

FEATURES:
* **New Resource:** `google_dataplex_data_asset` ([#11467](https://github.com/hashicorp/terraform-provider-google-beta/pull/11467))
* **New Resource:** `google_logging_saved_query` ([#11466](https://github.com/hashicorp/terraform-provider-google-beta/pull/11466))

IMPROVEMENTS:
* alloydb: added `restore_backupdr_backup_source`, `restore_backupdr_pitr_source`, and `backupdr_backup_source` in `google_alloydb_cluster` ([#11473](https://github.com/hashicorp/terraform-provider-google-beta/pull/11473))
* alloydb: added `rules_config_info.last_successful_backup_consistency_time` to `google_backup_dr_backup_plan_association` ([#11473](https://github.com/hashicorp/terraform-provider-google-beta/pull/11473))
* compute: updated `target_service` field to be mutable in `google_compute_service_attachment` resource ([#11469](https://github.com/hashicorp/terraform-provider-google-beta/pull/11469))
* datafusion: added `patch_revision` field to `google_data_fusion_instance` resource ([#11468](https://github.com/hashicorp/terraform-provider-google-beta/pull/11468))
* firestore: added `skip_wait` field to `google_firestore_index` resource, skipping the wait for index creation ([#11476](https://github.com/hashicorp/terraform-provider-google-beta/pull/11476))
* gkeonprem: added `skip_validations` field to `google_gkeonprem_vmware_cluster` resource ([#11465](https://github.com/hashicorp/terraform-provider-google-beta/pull/11465))
* sql: added `database_role` field and `iam_email` field to `google_sql_user` resource to support managing Cloud SQL users with database roles. ([#11471](https://github.com/hashicorp/terraform-provider-google-beta/pull/11471))

BUG FIXES:
* cloudbuild: fixed `google_cloudbuild_trigger` to allow creation without source configuration for manual triggers ([#11470](https://github.com/hashicorp/terraform-provider-google-beta/pull/11470))
* cloudrunv2: fix permadiff on `scaling.scaling_mode` in `google_cloud_run_v2_worker_pool` ([#11472](https://github.com/hashicorp/terraform-provider-google-beta/pull/11472))
* compute: resolved issues where `show_nat_ips` and `nat_ips` in `google_compute_service_attachment` were causing test failures due to an underlying API problem. These fields are now temporarily non-functional and will be ignored. ([#11460](https://github.com/hashicorp/terraform-provider-google-beta/pull/11460))
* container: fixed a bug in `google_container_node_pool` that prevented creation when `blue_green_settings` was specified ([#11464](https://github.com/hashicorp/terraform-provider-google-beta/pull/11464))
* container: fixed perma-diff in `google_container_cluster` when setting `resource_limits` with disabled node autoprovisioning ([#11474](https://github.com/hashicorp/terraform-provider-google-beta/pull/11474))

## 7.17.0 (January 27, 2026)

BREAKING CHANGES:
* networkconnectivity: changed `services` in `google_network_connectivity_multicloud_data_transfer_config` from TypeList to TypeSet. The order of or value of interpolations referencing the field may change. ([#11393](https://github.com/hashicorp/terraform-provider-google-beta/pull/11393))

FEATURES:
* **New Data Source:** `google_vmwareengine_datastore` ([#11428](https://github.com/hashicorp/terraform-provider-google-beta/pull/11428))
* **New Resource:** `google_compute_region_composite_health_check` (beta) ([#11436](https://github.com/hashicorp/terraform-provider-google-beta/pull/11436))
* **New Resource:** `google_dataplex_data_product` ([#11427](https://github.com/hashicorp/terraform-provider-google-beta/pull/11427))
* **New Resource:** `google_dialogflow_cx_tool_version` ([#11421](https://github.com/hashicorp/terraform-provider-google-beta/pull/11421))
* **New Resource:** `google_firebase_ai_logic_config` ([#11429](https://github.com/hashicorp/terraform-provider-google-beta/pull/11429))
* **New Resource:** `google_firebase_ai_logic_prompt_template` ([#11439](https://github.com/hashicorp/terraform-provider-google-beta/pull/11439))
* **New Resource:** `google_firebase_ai_logic_prompt_template_lock` ([#11448](https://github.com/hashicorp/terraform-provider-google-beta/pull/11448))
* **New Resource:** `google_gke_hub_rollout_sequence` (beta) ([#11423](https://github.com/hashicorp/terraform-provider-google-beta/pull/11423))
* **New Resource:** `google_saas_runtime_unit_operation` ([#11388](https://github.com/hashicorp/terraform-provider-google-beta/pull/11388))
* **New Resource:** `google_vmwareengine_datastore` ([#11428](https://github.com/hashicorp/terraform-provider-google-beta/pull/11428))

IMPROVEMENTS:
* backupdr: added support for restore compute instance and disk ([#11373](https://github.com/hashicorp/terraform-provider-google-beta/pull/11373))
* bigquery: added `source_column_match` field to `csv_options` in `google_bigquery_table` resource ([#11441](https://github.com/hashicorp/terraform-provider-google-beta/pull/11441))
* bigqueryanalyticshub: added `replica_locations` and `effective_replicas` fields to `google_bigquery_analytics_hub_listing` and `replica_locations` field to `google_bigquery_analytics_hub_listing_subscription` (beta) ([#11443](https://github.com/hashicorp/terraform-provider-google-beta/pull/11443))
* compute: added `FIPS_202205` enum to `PROFILE` field in `SSL_POLICY` and `REGION_SSL_POLICY` resources, and added `TLS_1_3` enum to `MIN_TLS_VERSION` field in `SSL_POLICY` and `REGION_SSL_POLICY` resources. ([#11398](https://github.com/hashicorp/terraform-provider-google-beta/pull/11398))
* compute: added `attachments` field to `google_compute_interconnect_attachment_group.logicalStructure.regions.metros.facilities.zones` and deprecated `attachment` field ([#11426](https://github.com/hashicorp/terraform-provider-google-beta/pull/11426))
* compute: added `enable_enhanced_ipv4_allocation` field to `google_compute_public_delegated_prefix` resource ([#11378](https://github.com/hashicorp/terraform-provider-google-beta/pull/11378))
* compute: added `ip_collection` field to `google_compute_address` resource ([#11378](https://github.com/hashicorp/terraform-provider-google-beta/pull/11378))
* compute: added `mac_address` field to `google_compute_instance` (beta) ([#11424](https://github.com/hashicorp/terraform-provider-google-beta/pull/11424))
* compute: added `source_instant_snapshot` field to `google_compute_snapshot` resource ([#11401](https://github.com/hashicorp/terraform-provider-google-beta/pull/11401))
* compute: added support for "IF_L2_FORWARDING" as a value for the `availableFeatures` field of the `google_compute_interconnect` resource ([#11384](https://github.com/hashicorp/terraform-provider-google-beta/pull/11384))
* compute: added support for "IF_L2_FORWARDING" as a value for the `requestedFeatures` field  of the `google_compute_interconnect` resource ([#11384](https://github.com/hashicorp/terraform-provider-google-beta/pull/11384))
* compute: added support for "L2_DEDICATED" as a value for the `type` field of the `google_compute_interconnect_attachment` resource. ([#11384](https://github.com/hashicorp/terraform-provider-google-beta/pull/11384))
* compute: added support for `igmp_query` field in `google_compute_instance`, `google_compute_instance_template`, and related instance resources. ([#11385](https://github.com/hashicorp/terraform-provider-google-beta/pull/11385))
* compute: added support for the `l2Forwarding` field to `google_compute_interconnect_attachment` ([#11384](https://github.com/hashicorp/terraform-provider-google-beta/pull/11384))
* container: added `accelerator_network_config` field to `node_pool` resource ([#11435](https://github.com/hashicorp/terraform-provider-google-beta/pull/11435))
* container: added `managed_opentelemetry_config` to `google_container_cluster` resource ([#11438](https://github.com/hashicorp/terraform-provider-google-beta/pull/11438))
* container: added `node_drain_config` field to `google_container_node_pool` resources ([#11409](https://github.com/hashicorp/terraform-provider-google-beta/pull/11409))
* container: added `pod_snapshot_config` field to `google_container_cluster` resource (beta) ([#11394](https://github.com/hashicorp/terraform-provider-google-beta/pull/11394))
* datastream: added `backfill_all.spanner_excluded_objects` and  `source_config.spanner_source_config` fields to `google_datastream_stream` ([#11416](https://github.com/hashicorp/terraform-provider-google-beta/pull/11416))
* datastream: added `spanner_profile` field to `google_datastream_connection_profile` ([#11416](https://github.com/hashicorp/terraform-provider-google-beta/pull/11416))
* dialogflowcx: added `serviceAccountAuthConfig ` field to `google_dialogflow_cx_webhook` resource ([#11402](https://github.com/hashicorp/terraform-provider-google-beta/pull/11402))
* oracledatabase: added `peerAutonomousDatabases`, `disasterRecoverySupportedLocations`, `sourceConfig` fields to Autonomous database resource. ([#11437](https://github.com/hashicorp/terraform-provider-google-beta/pull/11437))
* tags: added `allowed_values_regex` field to `google_tags_tag_key` resource ([#11442](https://github.com/hashicorp/terraform-provider-google-beta/pull/11442))
* tags: added support for dynamic tag keys in `google_tags_tag_binding` and `google_tags_location_tag_binding`  resources ([#11446](https://github.com/hashicorp/terraform-provider-google-beta/pull/11446))
* vertex_ai: added `deployment_spec.psc_interface_config` to `google_vertex_ai_reasoning_engine` ([#11391](https://github.com/hashicorp/terraform-provider-google-beta/pull/11391))

BUG FIXES:
* bigquery: fixed permadiff with the `collation` field in `google_bigquery_table.schema` ([#11389](https://github.com/hashicorp/terraform-provider-google-beta/pull/11389))
* cloudasset: fixed bug in `google_cloud_asset_folder_feed` where `folder_id` was always empty ([#11415](https://github.com/hashicorp/terraform-provider-google-beta/pull/11415))
* cloudbuild: fixed permadiff on `google_cloudbuild_trigger.pubsub_config.service_account_email` ([#11410](https://github.com/hashicorp/terraform-provider-google-beta/pull/11410))
* compute: fix crash when specifying an empty `instance_flexibility_policy` block on the `google_compute_region_instance_group_manager` resource ([#11377](https://github.com/hashicorp/terraform-provider-google-beta/pull/11377))
* compute: fixed a permadiff that could occur when using mixed short and long form IPv6 addresses in the `source_ranges` field of `google_compute_firewall` ([#11440](https://github.com/hashicorp/terraform-provider-google-beta/pull/11440))
* iambeta: fixed a permadiff that could occur in the `jwks_json` field for `google_iam_workload_identity_pool_provider` resource ([#11430](https://github.com/hashicorp/terraform-provider-google-beta/pull/11430))
* netapp: fixed export_policy update bug with squash_mode in netapp volume ([#11397](https://github.com/hashicorp/terraform-provider-google-beta/pull/11397))
* networkconnectivity: fixed a diff on `services` in `google_network_connectivity_multicloud_data_transfer_config` reordering elements ([#11393](https://github.com/hashicorp/terraform-provider-google-beta/pull/11393))
* sql: fixed an issue where transient server errors caused false failures for SQL operations that eventually completed successfully ([#11381](https://github.com/hashicorp/terraform-provider-google-beta/pull/11381))
* workbench: made `enable-jupyterlab4` metadata key settable for `google_workbench_instance` ([#11395](https://github.com/hashicorp/terraform-provider-google-beta/pull/11395))

## 7.16.0 (January 13, 2026)

DEPRECATIONS:
* cloudrunv2: deprecated `custom_audience` field in the `google_cloud_run_v2_worker_pool` resource, as this field is not applicable to the WorkerPools resource ([#11348](https://github.com/hashicorp/terraform-provider-google-beta/pull/11348))

FEATURES:
* **New Data Source:** `google_compute_routers` ([#11367](https://github.com/hashicorp/terraform-provider-google-beta/pull/11367))
* **New Resource:** `google_backup_dr_restore_workload` ([#11373](https://github.com/hashicorp/terraform-provider-google-beta/pull/11373))

IMPROVEMENTS:
* backupdr: added `max_custom_on_demand_retention_days` field to `google_backup_dr_backup_plan` resource ([#11359](https://github.com/hashicorp/terraform-provider-google-beta/pull/11359))
* bigquery: added support for merge and update operations for dataPolicies in `schema` field in `google_bigquery_table` resource when `ignore_schema_changes` is defined ([#11371](https://github.com/hashicorp/terraform-provider-google-beta/pull/11371))
* bigtable: added `etag` field to `google_bigtable_schema_bundle` resource ([#11346](https://github.com/hashicorp/terraform-provider-google-beta/pull/11346))
* compute: added `BPS_400G` enum value to `bandwidth` field in `google_compute_interconnect_attachment` resource ([#11366](https://github.com/hashicorp/terraform-provider-google-beta/pull/11366))
* container: added `registry_hosts` field to `containerd_config` in `google_container_cluster` and `google_container_node_pool` resources ([#11360](https://github.com/hashicorp/terraform-provider-google-beta/pull/11360))
* dataplex: added `one_time` field to `google_dataplex_datascan` resource ([#11351](https://github.com/hashicorp/terraform-provider-google-beta/pull/11351))
* datastream: added `postgresql_profile.ssl_config` to `google_datastream_connection_profile` resource ([#11344](https://github.com/hashicorp/terraform-provider-google-beta/pull/11344))
* networkservices: added `EXT_AUTHZ_GRPC` enum value to `wire_format` field in `google_network_services_authz_extension` resource ([#11361](https://github.com/hashicorp/terraform-provider-google-beta/pull/11361))
* networkservices: added `disable_placement_policy` field to `google_network_services_multicast_domain_activation` resource ([#11370](https://github.com/hashicorp/terraform-provider-google-beta/pull/11370))
* networkservices: added `metadata`, `supported_events`, `request_body_send_mode`, and `observability_mode` fields to `google_network_services_lb_route_extension` resource ([#11356](https://github.com/hashicorp/terraform-provider-google-beta/pull/11356))
* securitycenterv2: added support for supplying `location` values other than "GLOBAL" to `google_scc_v2_project_notification_config` resource ([#11354](https://github.com/hashicorp/terraform-provider-google-beta/pull/11354))
* storageinsights: added `activity_data_retention_period_days` field to `google_storage_insights_dataset_config` resource ([#11357](https://github.com/hashicorp/terraform-provider-google-beta/pull/11357))
* workbench: added support to set post-startup script metadata keys with managed EUC in `google_workbench_instance` resource ([#11369](https://github.com/hashicorp/terraform-provider-google-beta/pull/11369))

## 7.15.0 (January 6, 2026)

NOTES:
* lustre: increased delete and update operation timeouts from 20 minutes to 60 minutes for `google_lustre_instance` resource ([#11339](https://github.com/hashicorp/terraform-provider-google-beta/pull/11339))

BREAKING CHANGES:
* compute: changed `cipher_suite` fields in the `google_compute_vpn_tunnel` resource to track order ([#11334](https://github.com/hashicorp/terraform-provider-google-beta/pull/11334))

FEATURES:
* **New Resource:** `google_apigee_security_feedback` ([#11299](https://github.com/hashicorp/terraform-provider-google-beta/pull/11299))
* **New Resource:** `google_apphub_boundary` ([#11324](https://github.com/hashicorp/terraform-provider-google-beta/pull/11324))
* **New Resource:** `google_biglake_iceberg_catalog_iam_binding` ([#11322](https://github.com/hashicorp/terraform-provider-google-beta/pull/11322))
* **New Resource:** `google_biglake_iceberg_catalog_iam_member` ([#11322](https://github.com/hashicorp/terraform-provider-google-beta/pull/11322))
* **New Resource:** `google_biglake_iceberg_catalog_iam_policy` ([#11322](https://github.com/hashicorp/terraform-provider-google-beta/pull/11322))
* **New Resource:** `google_biglake_iceberg_catalog` ([#11273](https://github.com/hashicorp/terraform-provider-google-beta/pull/11273))
* **New Resource:** `google_network_connectivity_destination` ([#11340](https://github.com/hashicorp/terraform-provider-google-beta/pull/11340))
* **New Resource:** `google_network_connectivity_multicloud_data_transfer_config` ([#11307](https://github.com/hashicorp/terraform-provider-google-beta/pull/11307))

IMPROVEMENTS:
* backupdr: added ignore_read to `encryption_config` field in `google_backup_dr_backup_vault` resource ([#11345](https://github.com/hashicorp/terraform-provider-google-beta/pull/11345))
* biglakeiceberg: made `google_biglake_iceberg_catalog` use the resource project as the quota project when `user_project_override` is `true` ([#11322](https://github.com/hashicorp/terraform-provider-google-beta/pull/11322))
* cloudrunv2: added `source_code` field to `template.container` in `google_cloud_run_v2_service` resource ([#11336](https://github.com/hashicorp/terraform-provider-google-beta/pull/11336))
* composer: added new enum `ENVIRONMENT_SIZE_EXTRA_LARGE` to `environment_size` field to `google_composer_environment` resource ([#11276](https://github.com/hashicorp/terraform-provider-google-beta/pull/11276))
* compute: added `guest_flush` field to `google_compute_snapshot` resource ([#11281](https://github.com/hashicorp/terraform-provider-google-beta/pull/11281))
* compute: added `prefix_length` field to `google_compute_addresses` data source ([#11331](https://github.com/hashicorp/terraform-provider-google-beta/pull/11331))
* iamworkforcepool: added `hard_delete` support in `google_iam_workforce_pool_provider_scim_tenant` resource ([#11333](https://github.com/hashicorp/terraform-provider-google-beta/pull/11333))
* looker: added `periodic_export_config` field to `google_looker_instance` resource ([#11308](https://github.com/hashicorp/terraform-provider-google-beta/pull/11308))
* lustre: added `access_rules_options` field to `google_lustre_instance` resource to support root squashing and IP-based access control configuration ([#11311](https://github.com/hashicorp/terraform-provider-google-beta/pull/11311))
* managedkafka: replaced `disk_size_gb` with `disk_size_gib` in `broker_capacity_config` within the `google_managed_kafka_cluster` resource ([#11310](https://github.com/hashicorp/terraform-provider-google-beta/pull/11310))
* networkservices: added `state` field to `google_network_services_multicast_domain` resource ([#11277](https://github.com/hashicorp/terraform-provider-google-beta/pull/11277))
* redis: added support for `labels` to `google_redis_cluster` ([#11323](https://github.com/hashicorp/terraform-provider-google-beta/pull/11323))
* sql: marked `replication_cluster.psa_write_endpoint` field as Computed in `google_sql_database_instance` resource ([#11289](https://github.com/hashicorp/terraform-provider-google-beta/pull/11289))
* sql: set `replication_cluster` when update `google_sql_database_instance` resource if there is a disaster recovery(DR) replica set or there is a PSA write endpoint ([#11289](https://github.com/hashicorp/terraform-provider-google-beta/pull/11289))
* storage: updated datasource `google_storage_object_signed_url.signed_url` to use virtual style hosted url ([#11284](https://github.com/hashicorp/terraform-provider-google-beta/pull/11284))
* vertexai: added `bigtable`, `zone`, `encryption_spec`, and `bigtable_options` fields to `google_vertex_ai_feature_online_store` resource ([#11304](https://github.com/hashicorp/terraform-provider-google-beta/pull/11304))
* vertexai: added `psc_automation_configs` field to `google_vertex_ai_index_endpoint` resource ([#11286](https://github.com/hashicorp/terraform-provider-google-beta/pull/11286))

BUG FIXES:
* provider: fixed an issue where error type 409 and 412 were not being correctly retried. This commonly shows up in IAM resources, but can appear in other resources as well ([#11302](https://github.com/hashicorp/terraform-provider-google-beta/pull/11302))
* alloydb: fixed an issue where boolean fields were ignored when set to `false` for `google_alloydb_cluster` and `google_alloydb_instance` ([#11278](https://github.com/hashicorp/terraform-provider-google-beta/pull/11278))
* cloudrunv2: fixed a permadiff when default values of the `scaling` block were explicitly declared on the `google_cloud_run_v2_service` resource ([#11285](https://github.com/hashicorp/terraform-provider-google-beta/pull/11285))
* compute: fixed a crash in `google_compute_disk`/`google_compute_region_disk` when deleting a disk attached to an instance that had any scratch disks attached ([#11325](https://github.com/hashicorp/terraform-provider-google-beta/pull/11325))
* compute: fixed issue where `endpoints.interconnects.vlan_tags` wouldn't be read correctly from the API  in `google_compute_wire_group` resource ([#11305](https://github.com/hashicorp/terraform-provider-google-beta/pull/11305))
* compute: fixed update logic that causes empty instance being sent for `google_compute_network_endpoints` ([#11314](https://github.com/hashicorp/terraform-provider-google-beta/pull/11314))
* datacatalog: fixed issue where `fields.display_name` wouldn't be read correctly from the API in `google_data_catalog_tag` resource ([#11305](https://github.com/hashicorp/terraform-provider-google-beta/pull/11305))
* discoveryengine: marked `cmek_config_id` field in `google_discovery_engine_cmek_config` resource as required ([#11272](https://github.com/hashicorp/terraform-provider-google-beta/pull/11272))
* securitygateway: allowed empty field for `service_discovery` in `google_beyondcorp_security_gateway` ([#11330](https://github.com/hashicorp/terraform-provider-google-beta/pull/11330))
* securitygateway: allowed empty fields for `user_info`, `group_info` and `device_info` in `google_beyondcorp_security_gateway_application` ([#11330](https://github.com/hashicorp/terraform-provider-google-beta/pull/11330))
* servicedirectory: fixed an issue where `google_service_directory_endpoint` or `google_service_directory_service` without `metadata` specified would have other fields removed ([#11297](https://github.com/hashicorp/terraform-provider-google-beta/pull/11297))
* storage: fixed the behavior in `google_storage_bucket` resource when `force_destroy` is set to `true`. Previously, failing to list anywhere caches would prevent destroying objects on the bucket. Now, both objects and caches are processed independently. ([#11332](https://github.com/hashicorp/terraform-provider-google-beta/pull/11332))

## 7.14.0 (December 16, 2025)

DEPRECATIONS:
* managedkafka: added deprecation warning for `google_managed_kafka_connect_cluster` `additional_subnets` field ([#11253](https://github.com/hashicorp/terraform-provider-google-beta/pull/11253))

FEATURES:
* **New Data Source:** `google_artifact_registry_versions` ([#11266](https://github.com/hashicorp/terraform-provider-google-beta/pull/11266))
* **New Data Source:** `google_cloud_identity_policies` ([#11267](https://github.com/hashicorp/terraform-provider-google-beta/pull/11267))
* **New Data Source:** `google_compute_region_security_policy` ([#11254](https://github.com/hashicorp/terraform-provider-google-beta/pull/11254))
* **New Data Source:** `google_compute_storage_pool` ([#11251](https://github.com/hashicorp/terraform-provider-google-beta/pull/11251))
* **New Resource:** `google_compute_region_health_source` ([#11271](https://github.com/hashicorp/terraform-provider-google-beta/pull/11271))
* **New Resource:** `google_network_services_multicast_group_consumer_activation` ([#11269](https://github.com/hashicorp/terraform-provider-google-beta/pull/11269))
* **New Resource:** `google_network_services_multicast_group_producer_activation` ([#11243](https://github.com/hashicorp/terraform-provider-google-beta/pull/11243))

IMPROVEMENTS:
* alloydb: added `connection_pool_config`, `connection_pool_config.enabled` and `connection_pool_config.flags` in `google_alloydb_instance` resource ([#11250](https://github.com/hashicorp/terraform-provider-google-beta/pull/11250))
* colab: added `software_config.post_startup_script_config` field to `google_colab_runtime_template` ([#11264](https://github.com/hashicorp/terraform-provider-google-beta/pull/11264))
* compute: added new field `instance_flexibility_policy.instance_selection.min_cpu_platform` & `instance_flexibility_policy.instance_selection.disks` to `google_compute_region_instance_group_manager` ([#11231](https://github.com/hashicorp/terraform-provider-google-beta/pull/11231))
* dataplex: removed the need for import in `google_dataplex_entry` when using first party source systems ([#11262](https://github.com/hashicorp/terraform-provider-google-beta/pull/11262))
* dataproc: added `auto_stop_time` and  `idle_stop_ttl` to `google_dataproc_cluster` resource ([#11236](https://github.com/hashicorp/terraform-provider-google-beta/pull/11236))
* eventarc: added `retry_policy` field to `google_eventarc_trigger` resource ([#11239](https://github.com/hashicorp/terraform-provider-google-beta/pull/11239))
* networksecurity: enabled in-place update for `custom_mirroring_profile.mirroring_deployment_groups` on `google_network_security_security_profile` ([#11263](https://github.com/hashicorp/terraform-provider-google-beta/pull/11263))
* spanner: added `autoscaling_config.autoscaling_targets.total_cpu_utilization_percent` field to  `google_spanner_instance` resource ([#11257](https://github.com/hashicorp/terraform-provider-google-beta/pull/11257))
* sql: added changes to ignore changes in backup configuration's fields like `enabled`, `binary_log_enabled`, `start_time`, `point_in_time_recovery_enabled`, `transaction_log_retention_days` and `backup_retention_settings.retained_backups` in `google_sql_database_instance` if the instance is managed by Google Cloud Backup and Disaster (DR) Recovery Service. ([#11270](https://github.com/hashicorp/terraform-provider-google-beta/pull/11270))

BUG FIXES:
* compute: fixed `google_compute_network` in-place update to enable `enable_ula_internal_ipv6`. ([#11240](https://github.com/hashicorp/terraform-provider-google-beta/pull/11240))
* iam: fixed error 409 concurrency policy changes by correctly detecting the error type. ([#11244](https://github.com/hashicorp/terraform-provider-google-beta/pull/11244))
* sql: fixed an issue where the computed `psc_service_attachment_link` attribute was not being exported properly in `google_sql_database_instance` resource and datasources ([#11265](https://github.com/hashicorp/terraform-provider-google-beta/pull/11265))

## 7.13.0 (December 9, 2025)

NOTES:
* alloydb: reverted requiring `initial_user.password` as required on create for new `google_alloydb_cluster` resources, instead `initial_user.password` or `initial_user.user` must be set if `initial_user` is specified for `google_alloydb_cluster` resources ([#11158](https://github.com/hashicorp/terraform-provider-google-beta/pull/11158))
* privateca: modified `encryption_spec` field from `google_privateca_ca_pool` resource to be mutable and allow cmek key rotation ([#11110](https://github.com/hashicorp/terraform-provider-google-beta/pull/11110))

DEPRECATIONS:
* cloudquotas: deprecated `effective_container` and `effective_enablement` fields in the `google_cloud_quotas_quota_adjuster_settings` resource ([#11230](https://github.com/hashicorp/terraform-provider-google-beta/pull/11230))
* dlp: deprecated `publish_findings_to_cloud_data_catalog` field in `google_data_loss_prevention_job_trigger` resource. Use `publish_findings_to_dataplex_catalog` field instead. ([#11096](https://github.com/hashicorp/terraform-provider-google-beta/pull/11096))
* networkservices: removed `google_service_binding` resource due to service binding support being disabled ([#11159](https://github.com/hashicorp/terraform-provider-google-beta/pull/11159))

FEATURES:
* **New Resource:** `google_ces_app_version` ([#11128](https://github.com/hashicorp/terraform-provider-google-beta/pull/11128))
* **New Resource:** `google_dialogflow_generator` ([#11144](https://github.com/hashicorp/terraform-provider-google-beta/pull/11144))
* **New Resource:** `google_dialogflow_version` ([#11079](https://github.com/hashicorp/terraform-provider-google-beta/pull/11079))
* **New Resource:** `google_discovery_engine_widget_config` ([#11166](https://github.com/hashicorp/terraform-provider-google-beta/pull/11166))
* **New Resource:** `google_iam_workforce_pool_provider_scim_token` ([#11113](https://github.com/hashicorp/terraform-provider-google-beta/pull/11113))
* **New Resource:** `google_network_connectivity_gateway_advertised_route` ([#11156](https://github.com/hashicorp/terraform-provider-google-beta/pull/11156))
* **New Resource:** `google_network_security_sac_attachment` ([#11146](https://github.com/hashicorp/terraform-provider-google-beta/pull/11146))
* **New Resource:** `google_network_security_sac_realm` ([#11084](https://github.com/hashicorp/terraform-provider-google-beta/pull/11084))
* **New Resource:** `google_network_services_lb_edge_extension` ([#11129](https://github.com/hashicorp/terraform-provider-google-beta/pull/11129))
* **New Resource:** `google_network_services_multicast_consumer_association` ([#11141](https://github.com/hashicorp/terraform-provider-google-beta/pull/11141))
* **New Resource:** `google_network_services_multicast_group_range_activation` ([#11174](https://github.com/hashicorp/terraform-provider-google-beta/pull/11174))
* **New Resource:** `google_network_services_multicast_group_range` ([#11151](https://github.com/hashicorp/terraform-provider-google-beta/pull/11151))
* **New Resource:** `google_network_services_multicast_producer_association` ([#11123](https://github.com/hashicorp/terraform-provider-google-beta/pull/11123))

IMPROVEMENTS:
* alloydb: added `password_wo` and `password_wo_version` fields to `google_alloydb_user` resource ([#11109](https://github.com/hashicorp/terraform-provider-google-beta/pull/11109))
* apphub: added `identity` field to `google_apphub_service` and `google_apphub_workload` resources ([#11155](https://github.com/hashicorp/terraform-provider-google-beta/pull/11155))
* backupdr: added `encryption_config` field to `google_backup_dr_backup_vault` resource ([#11086](https://github.com/hashicorp/terraform-provider-google-beta/pull/11086))
* bigquery: added `external_runtime_options` and `python_options` fields to `bigquery_routine` resource ([#11125](https://github.com/hashicorp/terraform-provider-google-beta/pull/11125))
* ces: added `client_function.parameters.max_items`, `client_function.parameters.min_items`, `client_function.parameters.maximum`, `client_function.parameters.minimum`, `client_function.parameters.title`, `client_function.response.max_items`, `client_function.response.min_items`, `client_function.response.maximum`, `client_function.response.minimum`, and `client_function.response.title` fields to `google_ces_tool` resource ([#11131](https://github.com/hashicorp/terraform-provider-google-beta/pull/11131))
* ces: added `entry_agent` field to `google_ces_example` resource ([#11081](https://github.com/hashicorp/terraform-provider-google-beta/pull/11081))
* ces: added `google_search_tool.context_urls`, `google_search_tool.preferred_domains`, and `open_api_tool.api_authentication.bearer_token_config` fields to `google_ces_tool` resource ([#11131](https://github.com/hashicorp/terraform-provider-google-beta/pull/11131))
* ces: added `message.chunk.tool_response` and `message.chunk.tool_call` fields to `google_ces_example` resource ([#11081](https://github.com/hashicorp/terraform-provider-google-beta/pull/11081))
* ces: added `pinned` and `variable_declarations.schema.title` fields to `google_ces_app` resource ([#11090](https://github.com/hashicorp/terraform-provider-google-beta/pull/11090))
* cloudsecuritycompliance: added `cloud_control_details.parameters.parameter_value.oneof_value` fields to `google_cloud_security_compliance_framework_deployment` resource ([#11170](https://github.com/hashicorp/terraform-provider-google-beta/pull/11170))
* cloudsecuritycompliance: added `cloud_control_details.parameters.parameter_value.oneof_value` fields to `google_cloud_security_compliance_framework` resource ([#11170](https://github.com/hashicorp/terraform-provider-google-beta/pull/11170))
* cloudsecuritycompliance: added `parameter_spec.default_value.oneof_value` and `validation.allowed_values.values.oneof_value` fields to `google_cloud_security_compliance_cloud_control ` resource ([#11228](https://github.com/hashicorp/terraform-provider-google-beta/pull/11228))
* cloudsecuritycompliance: added `sub_parameters` field to `google_cloud_security_compliance_cloud_control ` resource ([#11228](https://github.com/hashicorp/terraform-provider-google-beta/pull/11228))
* colab: added `custom_environment_spec` field to `google_colab_notebook_execution` resource ([#11167](https://github.com/hashicorp/terraform-provider-google-beta/pull/11167))
* compute: added `network_pass_through_lb_traffic_policy` field to `google_compute_region_backend_service` resource. ([#11088](https://github.com/hashicorp/terraform-provider-google-beta/pull/11088))
* compute: added `params` field to `google_compute_interconnect` resource ([#11149](https://github.com/hashicorp/terraform-provider-google-beta/pull/11149))
* compute: added `show_nat_ips` and `nat_ips` fields to `google_compute_service_attachment` ([#11127](https://github.com/hashicorp/terraform-provider-google-beta/pull/11127))
* compute: added `snapshot_type` field to `google_compute_snapshot` resource ([#11147](https://github.com/hashicorp/terraform-provider-google-beta/pull/11147))
* compute: added new field `instance_flexibility_policy.instance_selection.min_cpu_platform` & `instance_flexibility_policy.instance_selection.disks` to `google_compute_region_instance_group_manager` ([#11231](https://github.com/hashicorp/terraform-provider-google-beta/pull/11231))
* container: added `autoscaled_rollout_policy` field to `google_container_node_pool` resource (beta) ([#11154](https://github.com/hashicorp/terraform-provider-google-beta/pull/11154))
* container: added `node_kernel_module_loading.policy` field to `google_container_node_pool` and `google_container_cluster` resources ([#11171](https://github.com/hashicorp/terraform-provider-google-beta/pull/11171))
* filestore: added support for updating `directory_services` fields in place in `google_filestore_instance` ([#11137](https://github.com/hashicorp/terraform-provider-google-beta/pull/11137))
* iamworkforcepool: added `claim_mapping`, `purge_time`, and `service_agent` fields to `google_iam_workforce_pool_provider_scim_tenant` resource ([#11113](https://github.com/hashicorp/terraform-provider-google-beta/pull/11113))
* looker: added `controlled_egress_enabled` and `controlled_egress_config` fields to `google_looker_instance` resource ([#11082](https://github.com/hashicorp/terraform-provider-google-beta/pull/11082))
* lustre: added `kms_key` field to `google_lustre_instance` resource ([#11104](https://github.com/hashicorp/terraform-provider-google-beta/pull/11104))
* modelarmor: added `google_mcp_server_floor_setting` field to `google_model_armor_floorsetting ` resource ([#11135](https://github.com/hashicorp/terraform-provider-google-beta/pull/11135))
* monitoring: fixes an issue with `google_monitoring_alert_policy` where it ignores the resource project during Import ([#11120](https://github.com/hashicorp/terraform-provider-google-beta/pull/11120))
* netapp: added public docs link for `google_netapp_host_group` resource ([#11160](https://github.com/hashicorp/terraform-provider-google-beta/pull/11160))
* netapp: added 'nfsv4' to custom update export_policy object in `google_netapp_volume` resource ([#11229](https://github.com/hashicorp/terraform-provider-google-beta/pull/11229))
* networksecurity: added field `custom_mirroring_profile.mirroring_deployment_groups` and output field `custom_mirroring_profile.mirroring_endpoint_group_type` to `google_network_security_security_profile` resource (beta) ([#11095](https://github.com/hashicorp/terraform-provider-google-beta/pull/11095))
* oracledatabase: added `properties.cpu_core_count`, `properties.secret_id`, and `properties.vault_id` fields to `google_oracle_database_autonomous` resource ([#11107](https://github.com/hashicorp/terraform-provider-google-beta/pull/11107))
* oracledatabase: added `properties.time_zone.version` field to `google_oracle_database_cloud_vm_cluster` resource ([#11107](https://github.com/hashicorp/terraform-provider-google-beta/pull/11107))
* servicedirectory: replaced `metadata` KeyValuePair with `annotations` KeyValueAnnotations in  `google_service_directory_service`, and `google_service_directory_endpoint` resources ([#11078](https://github.com/hashicorp/terraform-provider-google-beta/pull/11078))
* sql: added write-only argument for `root_password` in `google_sql_database_instance` resource ([#11098](https://github.com/hashicorp/terraform-provider-google-beta/pull/11098))
* storage: added `contexts` for resource `google_storage_bucket_object` ([#11145](https://github.com/hashicorp/terraform-provider-google-beta/pull/11145))
* vertex_ai: added `resourceLimits`, `minInstances`, `maxInstances`, `containerConcurrency` and `sourceCodeSpec` fields to `google_vertex_ai_reasoning_engine` resource ([#11148](https://github.com/hashicorp/terraform-provider-google-beta/pull/11148))

BUG FIXES:
* bigquery: fixed the permadiff when email field values contain non-lower-case characters in `access` in `google_bigquery_dataset` ([#11139](https://github.com/hashicorp/terraform-provider-google-beta/pull/11139))
* bigquery: fixed the permadiff when table schema is unchanged for a `google_bigquery_table` with row access policies ([#11102](https://github.com/hashicorp/terraform-provider-google-beta/pull/11102))
* cloudrunv2: fixed permadiff if `scaling` field is unset on resource `google_cloud_run_v2_service` ([#11132](https://github.com/hashicorp/terraform-provider-google-beta/pull/11132))
* compute: fixed an issue where the `bgp_always_compare_med` field could not be unset in in `google_compute_network`. It can now be unset by configuring the new field `delete_bgp_always_compare_med` to a value of `true`. ([#11121](https://github.com/hashicorp/terraform-provider-google-beta/pull/11121))
* compute: fixed crashes when no `network_endpoints` block specified in `google_compute_network_endpoints` resource or no network endpoints exist ([#11085](https://github.com/hashicorp/terraform-provider-google-beta/pull/11085))
* compute: fixed the `terms` field in `google_compute_router_route_policy` to be updatable without forcing resource recreation ([#11122](https://github.com/hashicorp/terraform-provider-google-beta/pull/11122))
* container: fixed a perpetual diff in `google_container_cluster` resource when `enable_l4_ilb_subsetting` is enabled by the GKE control plane and not explicitly set in the configuration ([#11143](https://github.com/hashicorp/terraform-provider-google-beta/pull/11143))
* dialogflowcx: fixed update_mask in `google_dialogflow_cx_playbook` where a granular update mask is required. ([#11100](https://github.com/hashicorp/terraform-provider-google-beta/pull/11100))
* discoveryengine: fixed a permadiff on `advanced_site_search_config` in `google_discovery_engine_data_store` resource ([#11175](https://github.com/hashicorp/terraform-provider-google-beta/pull/11175))
* iamworkforcepool: fixed bug in `google_iam_workforce_pool_provider_scim_token` where `base_uri` wasn't set correctly from the API ([#11113](https://github.com/hashicorp/terraform-provider-google-beta/pull/11113))
* logging: fixed an issue with `google_logging_*_sink.include_children` fields not being updatable to true ([#11094](https://github.com/hashicorp/terraform-provider-google-beta/pull/11094))
* memorystore: fixed an issue where a permadiff on `desired_auto_created_endpoints` caused the `google_memorystore_instance` resource to recreated. ([#11119](https://github.com/hashicorp/terraform-provider-google-beta/pull/11119))
* spanner: prevented recreation when `kms_key_name` and `kms_key_names` are same for `google_spanner_database` ([#11083](https://github.com/hashicorp/terraform-provider-google-beta/pull/11083))

## 7.12.0 (November 18, 2025)

DEPRECATIONS:
* backupdr: deprecated `required_type` in `google_backup_dr_backup_plan_associations` and `google_backup_dr_data_source_references`. Both resources no longer have functionality, and will be removed in the next major release. ([#25107](https://github.com/hashicorp/terraform-provider-google/pull/25107))

FEATURES:
* **New Resource:** `google_ces_agent` ([#11033](https://github.com/hashicorp/terraform-provider-google-beta/pull/11033))
* **New Resource:** `google_ces_guardrail` ([#11037](https://github.com/hashicorp/terraform-provider-google-beta/pull/11037))
* **New Resource:** `google_ces_tool` ([#11038](https://github.com/hashicorp/terraform-provider-google-beta/pull/11038))
* **New Resource:** `google_cloud_security_compliance_cloud_control` ([#11056](https://github.com/hashicorp/terraform-provider-google-beta/pull/11056))
* **New Resource:** `google_cloud_security_compliance_framework_deployment` ([#11057](https://github.com/hashicorp/terraform-provider-google-beta/pull/11057))
* **New Resource:** `google_cloud_security_compliance_framework` ([#11036](https://github.com/hashicorp/terraform-provider-google-beta/pull/11036))
* **New Resource:** `google_discovery_engine_serving_config` ([#11032](https://github.com/hashicorp/terraform-provider-google-beta/pull/11032))
* **New Resource:** `google_oracle_database_exascale_db_storage_vault` ([#11048](https://github.com/hashicorp/terraform-provider-google-beta/pull/11048))

IMPROVEMENTS:
* apphub: added `functional_type`, `registration_type`, and `extended_metadata` fields to `google_apphub_service` and `google_apphub_workload` resources ([#11062](https://github.com/hashicorp/terraform-provider-google-beta/pull/11062))
* ces: added `bearer_token_config` field to `google_ces_toolset` ([#11044](https://github.com/hashicorp/terraform-provider-google-beta/pull/11044))
* ces: added `client_certificate_settings` field to `google_ces_app` resource ([#11042](https://github.com/hashicorp/terraform-provider-google-beta/pull/11042))
* compute: added `block_names` field to `google_compute_reservation` resource ([#11046](https://github.com/hashicorp/terraform-provider-google-beta/pull/11046))
* compute: added `resolve_subnetwork_mask` field to the `google_compute_subnetwork` ([#11051](https://github.com/hashicorp/terraform-provider-google-beta/pull/11051))
* compute: added `sub_block_names` field to `google_compute_reservation_block` data source ([#11046](https://github.com/hashicorp/terraform-provider-google-beta/pull/11046))
* compute: added `tls_settings` field to `google_compute_regional_backend_service` resource ([#11025](https://github.com/hashicorp/terraform-provider-google-beta/pull/11025))
* container: added `end_time_behavior` field to `google_container_cluster` resource ([#11045](https://github.com/hashicorp/terraform-provider-google-beta/pull/11045))
* container: added `secret_sync_config` field to `google_container_cluster` resource (beta) ([#11049](https://github.com/hashicorp/terraform-provider-google-beta/pull/11049))
* container: added `writable_cgroups` field to `node_config.defaults.containerd_config` in `google_container_cluster` resource ([#25140](https://github.com/hashicorp/terraform-provider-google/pull/25140))
* dataplex: added `catalog_publishing_enabled` field to `data_profile_spec` in `google_dataplex_datascan` resource ([#11060](https://github.com/hashicorp/terraform-provider-google-beta/pull/11060))
* dns: added `forwarding_config.target_name_servers.ipv6_address` argument to `google_dns_managed_zone` resource ([#11050](https://github.com/hashicorp/terraform-provider-google-beta/pull/11050))
* gkeonprem: added `advanced_networking`, `multiple_network_interfaces_config` and `bgp_lb_config` fields to `google_gkeonprem_bare_metal_cluster` resource ([#11055](https://github.com/hashicorp/terraform-provider-google-beta/pull/11055))
* managedkafka: added `broker_capacity_config` field to `google_managed_kafka_cluster` resource ([#11027](https://github.com/hashicorp/terraform-provider-google-beta/pull/11027))
* networksecurity: added `endpoint_settings.jumbo_frames_enabled` field to `google_network_security_firewall_endpoint` resource ([#11026](https://github.com/hashicorp/terraform-provider-google-beta/pull/11026))
* run: added `readiness_probe` field to `cloud_run_service` resource ([#11039](https://github.com/hashicorp/terraform-provider-google-beta/pull/11039))

BUG FIXES:
* backupdr: updated `google_backup_dr_backup_plan_associations` and `google_backup_dr_data_source_references` to use LIST APIs, and require the correct List permissions ([#11034](https://github.com/hashicorp/terraform-provider-google-beta/pull/11034))
* provider: resolved an issue preventing X.509 certificates from being used for authentication when supplied as Application Default Credentials ([#11061](https://github.com/hashicorp/terraform-provider-google-beta/pull/11061))

## 7.11.0 (November 11, 2025)

DEPRECATIONS:
* `pubsublite`: `google_pubsub_lite_reservation` will be turned down effective March 18, 2026. Use `google_pubsub_reservation` instead. ([#11022](https://github.com/hashicorp/terraform-provider-google-beta/pull/11022))
* `pubsublite`: `google_pubsub_lite_subscription` will be turned down effective March 18, 2026. Use `google_pubsub_subscription` instead. ([#11022](https://github.com/hashicorp/terraform-provider-google-beta/pull/11022))
* `pubsublite`: `google_pubsub_lite_topic` will be turned down effective March 18, 2026. Use `google_pubsub_topic` instead. ([#11022](https://github.com/hashicorp/terraform-provider-google-beta/pull/11022))

BREAKING CHANGES:
* netapp: made `google_netapp_volume.export_policy.rules.squash_mode` not preserve values returned by the API. Without this change, unsetting `squash_mode` in the provider can cause an API error. ([#11023](https://github.com/hashicorp/terraform-provider-google-beta/pull/11023))

FEATURES:
* **New Data Source:** `google_artifact_registry_python_packages` ([#11017](https://github.com/hashicorp/terraform-provider-google-beta/pull/11017))
* **New Data Source:** `google_cloud_identity_policy` ([#10991](https://github.com/hashicorp/terraform-provider-google-beta/pull/10991))
* **New Data Source:** `google_compute_reservation_block` ([#11005](https://github.com/hashicorp/terraform-provider-google-beta/pull/11005))
* **New Data Source:** `google_compute_reservation_sub_block` ([#11005](https://github.com/hashicorp/terraform-provider-google-beta/pull/11005))
* **New Resource:** `google_ces_deployment` ([#10990](https://github.com/hashicorp/terraform-provider-google-beta/pull/10990))
* **New Resource:** `google_ces_example` ([#11020](https://github.com/hashicorp/terraform-provider-google-beta/pull/11020))
* **New Resource:** `google_compute_region_health_aggregation_policy` ([#10995](https://github.com/hashicorp/terraform-provider-google-beta/pull/10995))
* **New Resource:** `google_discovery_engine_user_store` ([#11018](https://github.com/hashicorp/terraform-provider-google-beta/pull/11018))

IMPROVEMENTS:
* bigquery: added `external_data_configuration.decimal_target_types` to `google_bigquery_table` ([#10983](https://github.com/hashicorp/terraform-provider-google-beta/pull/10983))
* compute: added `internal_ipv6_prefix` field to the `google_compute_subnetwork` resource ([#11007](https://github.com/hashicorp/terraform-provider-google-beta/pull/11007))
* compute: added `ipv6_access_type` field and `INTERNAL_IPV6_SUBNETWORK_CREATION` as a supported value for the `mode` field in `google_compute_public_delegated_prefix` resource ([#10985](https://github.com/hashicorp/terraform-provider-google-beta/pull/10985))
* compute: added `ipv6_access_type` field to `google_compute_public_advertised_prefix` resource ([#10978](https://github.com/hashicorp/terraform-provider-google-beta/pull/10978))
* dataplex: added `data_documentation_spec` field to `google_dataplex_datascan` resource to support the `DATA_DOCUMENTATION` scan type ([#11014](https://github.com/hashicorp/terraform-provider-google-beta/pull/11014))
* dataproc: added `resource_manager_tags` to `google_dataproc_cluster` resource ([#11021](https://github.com/hashicorp/terraform-provider-google-beta/pull/11021))
* lustre: added `placement_policy` field to `google_lustre_instance` resource ([#11011](https://github.com/hashicorp/terraform-provider-google-beta/pull/11011))
* netapp: added `cache_parameters` field to `google_netapp_volume` resource ([#10976](https://github.com/hashicorp/terraform-provider-google-beta/pull/10976))
* secretmanager: added project and short name support for `secret` on `google_secret_manager_secret_version` ([#11015](https://github.com/hashicorp/terraform-provider-google-beta/pull/11015))
* secretmanager: added project and short name support for `secret` on ephemeral `google_secret_manager_secret_version` ([#11015](https://github.com/hashicorp/terraform-provider-google-beta/pull/11015))

BUG FIXES:
* alloydb: fixed issue with creation when `initial_user.password` was set to a computed value in `google_alloydb_cluster` ([#11006](https://github.com/hashicorp/terraform-provider-google-beta/pull/11006))
* bigquery: fixed extraneous diffs in `google_bigquery_table.external_data_configuration.schema` ([#10983](https://github.com/hashicorp/terraform-provider-google-beta/pull/10983))
* compute: fixed a breaking change in `google_compute_instance` introduced in 7.9.0 where a destroy-diff is prompted for instances with preset GPUs ([#10998](https://github.com/hashicorp/terraform-provider-google-beta/pull/10998))
* container: added `KUBE_DNS` as an accepted value for `cluster_dns` field on `google_container_cluster` ([#10997](https://github.com/hashicorp/terraform-provider-google-beta/pull/10997))
* netapp: fixed bug where unsetting `export_policy.rules.squash_mode` on `google_netapp_volume` can cause an API error ([#11023](https://github.com/hashicorp/terraform-provider-google-beta/pull/11023))
* pubsub: fixed bug where `google_pubsub_subscription` could only be updated if `bigquery_config` was modified ([#10996](https://github.com/hashicorp/terraform-provider-google-beta/pull/10996))
* sql: fixed bug where `final_backup_description` in `google_sql_database_instance` resource wasn't set on the final backup on delete ([#11019](https://github.com/hashicorp/terraform-provider-google-beta/pull/11019))
* storage: fixed bug where certain changes to `google_storage_bucket_acl.role_entity` were ignored ([#10994](https://github.com/hashicorp/terraform-provider-google-beta/pull/10994))
* workstations: fixed bug in `google_workstations_workstation` where setting `source_workstation` caused a permadiff that forced recreation ([#10986](https://github.com/hashicorp/terraform-provider-google-beta/pull/10986))
* vmwareengine: made deletion of `google_vmwareengine_private_cloud` wait until the deletion completes ([#11010](https://github.com/hashicorp/terraform-provider-google-beta/pull/11010))

## 7.10.0 (November 4, 2025)

BREAKING CHANGES:
* alloydb: marked `initial_user.password` as required on create of new `google_alloydb_cluster` resources. This change aligns the provider with existing API constraints to surface errors earlier. ([#10999](https://github.com/hashicorp/terraform-provider-google-beta/pull/10999))

FEATURES:
* **New Resource:** `google_ces_app` ([#10950](https://github.com/hashicorp/terraform-provider-google-beta/pull/10950))
* **New Resource:** `google_ces_toolset` ([#10967](https://github.com/hashicorp/terraform-provider-google-beta/pull/10967))
* **New Resource:** `google_discovery_engine_control` ([#10966](https://github.com/hashicorp/terraform-provider-google-beta/pull/10966))
* **New Resource:** `google_netapp_host_group` ([#10959](https://github.com/hashicorp/terraform-provider-google-beta/pull/10959))
* **New Resource:** `google_network_security_mirroring_endpoint` (beta) ([#10941](https://github.com/hashicorp/terraform-provider-google-beta/pull/10941))
* **New Resource:** `google_network_services_multicast_domain` ([#10952](https://github.com/hashicorp/terraform-provider-google-beta/pull/10952))
* **New Resource:** `google_privileged_access_manager_settings` ([#10961](https://github.com/hashicorp/terraform-provider-google-beta/pull/10961))
* **New Ephemeral Resource:** `google_client_config` ([#10975](https://github.com/hashicorp/terraform-provider-google-beta/pull/10975))

IMPROVEMENTS:
* cloudfunctions2: added `direct_vpc_network_interface` and `direct_vpc_egress` field to `google_cloudfunctions2_function` resource ([#10971](https://github.com/hashicorp/terraform-provider-google-beta/pull/10971))
* cloudrunv2: added `template.container.depends_on` field to `google_cloud_run_v2_worker_pool` resource ([#10970](https://github.com/hashicorp/terraform-provider-google-beta/pull/10970))
* container: added `network_tier_config` to `google_container_cluster` resource. ([#10960](https://github.com/hashicorp/terraform-provider-google-beta/pull/10960))
* eventarc: added `labels` field to `google_eventarc_channel` resource ([#10944](https://github.com/hashicorp/terraform-provider-google-beta/pull/10944))
* netapp: added `block_devices` field and `ISCSI` protocol support to `goolge_netapp_volume` resource, and increased timeouts on its operations ([#10974](https://github.com/hashicorp/terraform-provider-google-beta/pull/10974))
* netapp: added additional field `type` in `google_netapp_storage_pool` resource ([#10953](https://github.com/hashicorp/terraform-provider-google-beta/pull/10953))
* vertexai: added `psc_automation_configs` field to `google_vertex_ai_endpoint` resource ([#10956](https://github.com/hashicorp/terraform-provider-google-beta/pull/10956))
* vertexai: added `sync_config.continuous` field to `google_vertex_ai_feature_online_store_featureview` ([#10964](https://github.com/hashicorp/terraform-provider-google-beta/pull/10964))

BUG FIXES:
* accesscontextmanager: fixed issue where `google_access_context_manager_service_perimeter_[dry_run_][egress|ingress]_policy` caused the provider to crash when a provided identity casing was invalid. ([#10968](https://github.com/hashicorp/terraform-provider-google-beta/pull/10968))
* apigee: fixed issue where `credentials` block was not populated in the Terraform state in `google_apigee_developer_app` resource ([#10963](https://github.com/hashicorp/terraform-provider-google-beta/pull/10963))
* compute: fixed `google_compute_network_firewall_policy_rule` staying disabled after apply with `disabled = false` ([#10962](https://github.com/hashicorp/terraform-provider-google-beta/pull/10962))
* compute: fixed a breaking change in `google_compute_instance` introduced in 7.9.0 where a destroy-diff is prompted for instances with preset GPUs ([#10998](https://github.com/hashicorp/terraform-provider-google-beta/pull/10998))
* compute: resolve permadiff for `display_name` in new deployments of `google_compute_organization_security_policy` ([#10965](https://github.com/hashicorp/terraform-provider-google-beta/pull/10965))
* storage: fixed a conversion error in `google_storage_bucket` state migration. This bug impacted Pulumi users. ([#10943](https://github.com/hashicorp/terraform-provider-google-beta/pull/10943))

## 7.9.0 (October 28, 2025)

BREAKING CHANGES:
* beyondcorp: made the `ports` field in `endpoint_matchers` required in response to a change in the API surface. ([#10909](https://github.com/hashicorp/terraform-provider-google-beta/pull/10909))

FEATURES:
* **New Resource:** `google_firestore_user_creds` ([#10922](https://github.com/hashicorp/terraform-provider-google-beta/pull/10922))
* **New Resource:** `google_network_security_dns_threat_detector` ([#10898](https://github.com/hashicorp/terraform-provider-google-beta/pull/10898))

IMPROVEMENTS:
* appengine: added `ssl_policy` to `application` on `google_app_engine_application` resource ([#10915](https://github.com/hashicorp/terraform-provider-google-beta/pull/10915))
* bigquery: added support for IAM conditions in `google_bigquery_dataset_iam_*` resources ([#10913](https://github.com/hashicorp/terraform-provider-google-beta/pull/10913))
* container: added `dns_endpoint_confg.enable_k8s_tokens_via_dns` and  `dns_endpoint_config.enable_k8s_certs_via_dns` fields to `google_container_cluster` resource ([#10910](https://github.com/hashicorp/terraform-provider-google-beta/pull/10910))
* container: added `fleet.membership_type` field to `google_container_cluster` resource ([#10901](https://github.com/hashicorp/terraform-provider-google-beta/pull/10901))
* dataplex: added `data_classification` field to `google_dataplex_aspect_type` resource ([#10929](https://github.com/hashicorp/terraform-provider-google-beta/pull/10929))
* iamworkforcepool: added `scim_usage` field to `workforce_pool_provider` resource ([#10916](https://github.com/hashicorp/terraform-provider-google-beta/pull/10916))
* memorystore: added `available_maintenance_versions` field to `google_memorystore_instance` resource ([#10899](https://github.com/hashicorp/terraform-provider-google-beta/pull/10899))
* memorystore: added `maintenance_version` field to `google_memorystore_instance` resource ([#10896](https://github.com/hashicorp/terraform-provider-google-beta/pull/10896))
* redis: added `available_maintenance_versions` field to `google_redis_cluster` resource ([#10899](https://github.com/hashicorp/terraform-provider-google-beta/pull/10899))
* redis: added `maintenance_version` field to `google_redis_cluster` resource ([#10896](https://github.com/hashicorp/terraform-provider-google-beta/pull/10896))
* storagetransfer: added `transfer_manifest` field to `google_storage_transfer_job` resource ([#10907](https://github.com/hashicorp/terraform-provider-google-beta/pull/10907))

BUG FIXES:
* bigquery: added validation for `target_types` in `google_bigquery_dataset_access` ([#10932](https://github.com/hashicorp/terraform-provider-google-beta/pull/10932))
* cloudquotas: resolved permadiff for `preferred_value` in `google_cloud_quotas_quota_preference` ([#10911](https://github.com/hashicorp/terraform-provider-google-beta/pull/10911))
* compute: fixed scenario where `google_compute_instance` would not be staged for recreation if `guest_accelerator.count` was updated to 0 from non-zero value ([#10902](https://github.com/hashicorp/terraform-provider-google-beta/pull/10902))
* sql: fixed an issue where `dataDiskSize` was unintentionally null instead of set to the current value in API requests, triggering unrelated errors ([#10919](https://github.com/hashicorp/terraform-provider-google-beta/pull/10919))

## 7.8.0 (October 21, 2025)

FEATURES:
* **New Data Source:** `google_artifact_registry_packages` ([#10869](https://github.com/hashicorp/terraform-provider-google-beta/pull/10869))
* **New Data Source:** `google_network_management_connectivity_tests` ([#10856](https://github.com/hashicorp/terraform-provider-google-beta/pull/10856))
* **New Resource:** `google_apigee_environment_api_revision_deployment` ([#10862](https://github.com/hashicorp/terraform-provider-google-beta/pull/10862))
* **New Resource:** `google_dataplex_entry_link` ([#10893](https://github.com/hashicorp/terraform-provider-google-beta/pull/10893))
* **New Resource:** `google_discovery_engine_assistant` ([#10885](https://github.com/hashicorp/terraform-provider-google-beta/pull/10885))
* **New Resource:** `google_observability_trace_scope` ([#10873](https://github.com/hashicorp/terraform-provider-google-beta/pull/10873))
* **New Resource:** `google_oracle_database_db_system` ([#10889](https://github.com/hashicorp/terraform-provider-google-beta/pull/10889))
* **New Resource:** `google_saas_runtime_unit` ([#10865](https://github.com/hashicorp/terraform-provider-google-beta/pull/10865))

IMPROVEMENTS:
* compute: added `IN_FLIGHT` to `balancing_mode` on `google_compute_backend_service` resource ([#10875](https://github.com/hashicorp/terraform-provider-google-beta/pull/10875))
* compute: added `ncc_gateway` field to `google_compute_router` resource (beta) ([#10894](https://github.com/hashicorp/terraform-provider-google-beta/pull/10894))
* compute: added new field `instance_lifecycle_policy.on_repair.allow_changing_zone` to `google_compute_region_instance_group_manager` & `google_compute_instance_group_manager` ([#10871](https://github.com/hashicorp/terraform-provider-google-beta/pull/10871))
* compute: promoted `security_policy` in `compute_region_backend_service` resource to GA ([#10866](https://github.com/hashicorp/terraform-provider-google-beta/pull/10866))
* compute: promoted the `google_compute_preview_feature` resource to GA. ([#10886](https://github.com/hashicorp/terraform-provider-google-beta/pull/10886))
* compute: the `activation_status` attribute within the `google_compute_preview_feature` resource now uses the `ACTIVATION_STATE_UNSPECIFIED` value instead of `DISABLED`. Support for `DISABLED` will be added in a future release. ([#10886](https://github.com/hashicorp/terraform-provider-google-beta/pull/10886))
* datastream: added `backfill_all.mongodb_excluded_objects` and  `source_config.mongodb_source_config` fields to `google_datastream_stream` ([#10888](https://github.com/hashicorp/terraform-provider-google-beta/pull/10888))
* datastream: added `mongodb_profile` field to `google_datastream_connection_profile` ([#10888](https://github.com/hashicorp/terraform-provider-google-beta/pull/10888))
* discoveryengine: added `connector_modes`, `sync_mode`, `incremental_refresh_interval`, `auto_run_disabled`, and `incremental_sync_disabled` fields to `google_discovery_engine_data_connector` resource ([#10863](https://github.com/hashicorp/terraform-provider-google-beta/pull/10863))
* discoveryengine: added `kms_key_name` field to `google_discovery_engine_search_engine` resource ([#10863](https://github.com/hashicorp/terraform-provider-google-beta/pull/10863))
* discoveryengine: added in-place update support for `entities.params` and `entities.key_property_mappings` in `google_discovery_engine_data_connector` ([#10895](https://github.com/hashicorp/terraform-provider-google-beta/pull/10895))
* dlp: added `publish_findings_to_dataplex_catalog` field to `google_data_loss_prevention_job_trigger ` ([#10883](https://github.com/hashicorp/terraform-provider-google-beta/pull/10883))
* iambeta: allowed GKE workload identity pool pattern in `workload_identity_pool_id` field of `google_iam_workload_identity_pool` resource. ([#10861](https://github.com/hashicorp/terraform-provider-google-beta/pull/10861))
* memorystore: added `maintenance_version` field to `google_memorystore_instance` resource ([#10896](https://github.com/hashicorp/terraform-provider-google-beta/pull/10896))
* memorystore: added `available_maintenance_versions` field to `google_memorystore_instance` resource ([#10899](https://github.com/hashicorp/terraform-provider-google-beta/pull/10899))
* networkconnectivity: added `HYBRID_INSPECTION` enum value to `preset_topology` field in `google_network_connectivity_hub` resource ([#10894](https://github.com/hashicorp/terraform-provider-google-beta/pull/10894))
* networkconnectivity: added `gateway` field to `google_network_connectivity_spoke` resource (beta) ([#10894](https://github.com/hashicorp/terraform-provider-google-beta/pull/10894))
* networkservices: added `isolationConfig` on `google_network_services_service_lb_policies` resource ([#10858](https://github.com/hashicorp/terraform-provider-google-beta/pull/10858))
* redis: added `deletion_protection` field to `redis_instance` to make deleting them require an explicit intent. `redis_instance` resources now cannot be destroyed unless `deletion_protection = false` is set for the resource. ([#10860](https://github.com/hashicorp/terraform-provider-google-beta/pull/10860))
* redis: added `maintenance_version` field to `google_redis_cluster` resource ([#10896](https://github.com/hashicorp/terraform-provider-google-beta/pull/10896))
* redis: added `available_maintenance_versions` field to `google_redis_cluster` resource ([#10899](https://github.com/hashicorp/terraform-provider-google-beta/pull/10899))
* saas_runtime: added `default_release` field to `google_saas_runtime_unit_kind` resource ([#10887](https://github.com/hashicorp/terraform-provider-google-beta/pull/10887))
* sql: added `read_pool_auto_scale_config` support to `sql_database_instance` resource ([#10884](https://github.com/hashicorp/terraform-provider-google-beta/pull/10884))

BUG FIXES:
* bigquery: fixed the issue where `google_bigquery_table` detected an incorrect `schema` diff on tables with row access policies when the schema was unchanged. ([#10876](https://github.com/hashicorp/terraform-provider-google-beta/pull/10876))
* compute: allow `requested_link_count` to be updated in-place in `google_compute_interconnect` resource ([#10870](https://github.com/hashicorp/terraform-provider-google-beta/pull/10870))

## 7.7.0 (October 14, 2025)
BREAKING CHANGES:
* discoveryengine: changed type of `google_discovery_engine_data_connector.entities.params`. Previously, it was a map of string keys to string values; now, it must be a [JSON-encoded](https://developer.hashicorp.com/terraform/language/functions/jsonencode) string containing an object. This change is being made in a minor release because the field wasn't usable as intended – specifically, all current valid uses require mapping strings to _lists_ of strings. ([#10863](https://github.com/hashicorp/terraform-provider-google-beta/pull/10863))

FEATURES:
* **New Data Source:** `google_network_management_connectivity_tests` ([#10856](https://github.com/hashicorp/terraform-provider-google-beta/pull/10856))
* **New Resource:** `google_apigee_developer_app` ([#10851](https://github.com/hashicorp/terraform-provider-google-beta/pull/10851))
* **New Resource:** `google_discovery_engine_license_config` ([#10848](https://github.com/hashicorp/terraform-provider-google-beta/pull/10848))
* **New Resource:** `google_iam_workforce_pool_provider_scim_tenant` ([#10834](https://github.com/hashicorp/terraform-provider-google-beta/pull/10834))
* **New Resource:** `google_kms_project_kaj_policy_config` ([#10850](https://github.com/hashicorp/terraform-provider-google-beta/pull/10850))
* **New Resource:** `google_saas_runtime_tenant` ([#10841](https://github.com/hashicorp/terraform-provider-google-beta/pull/10841))

IMPROVEMENTS:
* apigee: updated the `scopes` argument in `google_apigee_api_product` resource to be order-insensitive. ([#10851](https://github.com/hashicorp/terraform-provider-google-beta/pull/10851))
* beyondcorp: added `proxy_protocol_config` and `service_discovery` fields to `google_beyondcorp_security_gateway` resource ([#10842](https://github.com/hashicorp/terraform-provider-google-beta/pull/10842))
* cloudrunv2: added `health_check_disabled` field to `google_cloud_run_v2_service` resource. ([#10839](https://github.com/hashicorp/terraform-provider-google-beta/pull/10839))
* compute: added `params` field to `google_compute_router` resource (GA) ([#10844](https://github.com/hashicorp/terraform-provider-google-beta/pull/10844))
* discoveryengine: added `connector_modes`, `sync_mode`, `incremental_refresh_interval`, `auto_run_disabled`, and `incremental_sync_disabled` fields to `google_discovery_engine_data_connector` resource ([#10863](https://github.com/hashicorp/terraform-provider-google-beta/pull/10863))
* discoveryengine: added `kms_key_name` field to `google_discovery_engine_search_engine` resource ([#10863](https://github.com/hashicorp/terraform-provider-google-beta/pull/10863))
* dlp: added `publish_to_dataplex_catalog` field to `discovery_config` resource ([#10849](https://github.com/hashicorp/terraform-provider-google-beta/pull/10849))
* gkeonprem: made it possible to set the `on_prem_version` field on `google_gkeonprem_vmware_node_pool` (previously output-only) ([#10847](https://github.com/hashicorp/terraform-provider-google-beta/pull/10847))
* memcache: added `deletion_protection` field to `memcache_instance` to make deleting them require an explicit intent. `memcache_instance` resources now cannot be destroyed unless `deletion_protection = false` is set for the resource. ([#10846](https://github.com/hashicorp/terraform-provider-google-beta/pull/10846))
* metastore: added `tags` field to `google_dataproc_metastore_service` and 'google_dataproc_metastore_federation' resources to allow setting tags for services and federation at creation time ([#10854](https://github.com/hashicorp/terraform-provider-google-beta/pull/10854))
* networksecurity: added `URL_FILTERING` option to enum field `type` for `google_network_security_security_profile` resource ([#10829](https://github.com/hashicorp/terraform-provider-google-beta/pull/10829))
* networksecurity: added `url_filtering_profile` field  to `google_network_security_security_profile_group` resource ([#10829](https://github.com/hashicorp/terraform-provider-google-beta/pull/10829))
* networksecurity: added `url_filtering_profile` field  to `google_network_security_security_profile` resource ([#10829](https://github.com/hashicorp/terraform-provider-google-beta/pull/10829))
* sql: added `source_instance_deletion_time` field to `google_sql_database_instance_latest_recovery_time` data source ([#10827](https://github.com/hashicorp/terraform-provider-google-beta/pull/10827))
* sql: added `source_instance_deletion_time` field to `google_sql_database_instance` resource ([#10827](https://github.com/hashicorp/terraform-provider-google-beta/pull/10827))

BUG FIXES:
* bigqueryanalyticshub: fixed `google_bigquery_analytics_hub_listing_subscription` import ([#10855](https://github.com/hashicorp/terraform-provider-google-beta/pull/10855))
* discoveryengine: fixed bug where it wasn't possible to specify values for `knowledgeBaseSysId` or `catalogSysId` in `google_discovery_engine_data_connector.entities.params`. ([#10863](https://github.com/hashicorp/terraform-provider-google-beta/pull/10863))

## 7.6.0 (October 7, 2025)
DEPRECATIONS:
* networksecurity: deprecated `ignore_case`, `exact`, `prefix`, `suffix` and `contains` fields in `http_rules.from.not_sources.principals` and `http_rules.from.sources.principals` blocks in `google_network_security_authz_policy` resource. Use the equivalent fields in `http_rules.from.not_sources.principals.principal` or `http_rules.from.sources.principals.principal` instead. ([#10809](https://github.com/hashicorp/terraform-provider-google-beta/pull/10809))

BREAKING CHANGES:
* container: `node_config` blocks that had set `kubelet_config` without explicitly setting `cpu_cfs_quota` implicitly set `cfu_cfs_quota` to `false` when unset. From this version onwards, an unset `cpu_cfs_quota` will instead match the API default of true `true`. Resources that are recreated will receive the new value; old resources are unaffected, and may change values by explicitly setting the intended one. ([#10823](https://github.com/hashicorp/terraform-provider-google-beta/pull/10823))
* storageinsights: removed `activity_data_retention_period_days` field from `google_storage_insights_dataset_config` resource due to a delayed launch. It will be readded when the feature launches. ([#10824](https://github.com/hashicorp/terraform-provider-google-beta/pull/10824))

FEATURES:
* **New Resource:** `google_kms_folder_kaj_policy_config` ([#10798](https://github.com/hashicorp/terraform-provider-google-beta/pull/10798))
* **New Resource:** `google_vertex_ai_cache_config` ([#10807](https://github.com/hashicorp/terraform-provider-google-beta/pull/10807))
* **New Resource:** `google_vertex_ai_reasoning_engine` ([#10797](https://github.com/hashicorp/terraform-provider-google-beta/pull/10797))

IMPROVEMENTS:
* backupdr: added `data_source` and `rules_config_info` fields to `google_backup_dr_backup_plan_associations` datasource ([#10802](https://github.com/hashicorp/terraform-provider-google-beta/pull/10802))
* beyondcorp: added `external`, `proxy_protocol`, and `schema` fields to `google_beyondcorp_security_gateway_application` resource ([#10808](https://github.com/hashicorp/terraform-provider-google-beta/pull/10808))
* beyondcorp: changed `endpoint_matchers` field to not be required anymore in the `google_beyondcorp_security_gateway_application` resource ([#10808](https://github.com/hashicorp/terraform-provider-google-beta/pull/10808))
* bigquery: added `reservation` field to `google_bigquery_job` resource ([#10796](https://github.com/hashicorp/terraform-provider-google-beta/pull/10796))
* compute: added `backend.max_in_flight_requests`, `backend.max_in_flight_requests_per_instance`, `backend.max_in_flight_requests_per_endpoint` and `backend.traffic_duration` fields to `google_compute_backend_service` resource ([#10799](https://github.com/hashicorp/terraform-provider-google-beta/pull/10799))
* compute: added `shared_secret_wo` and `shared_secret_wo_version` fields to `google_compute_vpn_tunnel` resource, enabling write-only management of the shared secret. ([#10788](https://github.com/hashicorp/terraform-provider-google-beta/pull/10788))
* dlp: added `SENSITIVITY_UNKNOWN` as possible enum value for `actions.tag_resources.tag_conditions.sensitivity_score.score` in `google_data_loss_prevention_discovery_config` resource ([#10820](https://github.com/hashicorp/terraform-provider-google-beta/pull/10820))
* dlp: added `actions.save_findings.output_config.storage_path` field to `google_data_loss_prevention_job_trigger` resource ([#10816](https://github.com/hashicorp/terraform-provider-google-beta/pull/10816))
* lustre: increased creation timeout from 20min to 40min for `google_lustre_instance` resource ([#10817](https://github.com/hashicorp/terraform-provider-google-beta/pull/10817))
* netapp: added `hybrid_replication_user_commands` field  with subfield `commands` to `google_netapp_volume_replication` resource ([#10813](https://github.com/hashicorp/terraform-provider-google-beta/pull/10813))
* netapp: added `replication_schedule`, `hybrid_replication_type`, `large_volume_constituent_count` fields to `hybrid_replication_parameters` field in `google_netapp_volume` resource ([#10813](https://github.com/hashicorp/terraform-provider-google-beta/pull/10813))
* networksecurity: added `ip_blocks` field to `google_network_security_authz_policy` resource ([#10809](https://github.com/hashicorp/terraform-provider-google-beta/pull/10809))
* secretmanager: added ephemeral support for `google_secret_manager_secret_version` resource ([#10821](https://github.com/hashicorp/terraform-provider-google-beta/pull/10821))
* sql: added `source_instance_deletion_time` field to `google_sql_database_instance_latest_recovery_time` data source ([#10827](https://github.com/hashicorp/terraform-provider-google-beta/pull/10827))
* sql: added `source_instance_deletion_time` field to `google_sql_database_instance` resource ([#10827](https://github.com/hashicorp/terraform-provider-google-beta/pull/10827))
* storagetransfer: added `user_project_override` and `billing_project` fields to  `google_storage_transfer_job` resource ([#10790](https://github.com/hashicorp/terraform-provider-google-beta/pull/10790))

BUG FIXES:
* container: fixed the default for `node_config.kubelet_config.cpu_cfs_quota` on `google_container_cluster`, `google_container_node_pool`, `google_container_cluster.node_pool` to align with the API. Terraform will now send a `true` value when the field is unset on creation, and preserve any previously set value when unset. Explicitly set values will work as defined in configuration. ([#10823](https://github.com/hashicorp/terraform-provider-google-beta/pull/10823))



## 7.5.0 (September 30, 2025)

BREAKING CHANGES:
* netapp: changed `peer_ip_addresses` field type from String to Array in `google_netapp_volume` resource, as it was unusable otherwise ([#10757](https://github.com/hashicorp/terraform-provider-google-beta/pull/10757))

FEATURES:
* **New Data Source:** `google_artifact_registry_maven_artifacts` ([#10785](https://github.com/hashicorp/terraform-provider-google-beta/pull/10785))
* **New Data Source:** `google_artifact_registry_npm_packages` ([#10784](https://github.com/hashicorp/terraform-provider-google-beta/pull/10784))
* **New Resource:** `google_apigee_api_deployment` ([#10776](https://github.com/hashicorp/terraform-provider-google-beta/pull/10776))
* **New Resource:** `google_discovery_engine_data_connector` ([#10778](https://github.com/hashicorp/terraform-provider-google-beta/pull/10778))
* **New Resource:** `google_kms_organization_kaj_policy_config` ([#10777](https://github.com/hashicorp/terraform-provider-google-beta/pull/10777))
* **New Resource:** `google_saas_runtime_rollout_kind` ([#10764](https://github.com/hashicorp/terraform-provider-google-beta/pull/10764))

IMPROVEMENTS:
* cloudrunv2: added `startup_probe` and `liveness_probe` to `google_cloud_run_v2_worker_pool` resource ([#10749](https://github.com/hashicorp/terraform-provider-google-beta/pull/10749))
* compute: added `bandwidth_allocation` field to `google_compute_wire_group` resource ([#10770](https://github.com/hashicorp/terraform-provider-google-beta/pull/10770))
* compute: added `shared_secret_wo` and `shared_secret_wo_version` fields for `google_compute_vpn_tunnel` resource, enabling write-only management of the shared secret. ([#10788](https://github.com/hashicorp/terraform-provider-google-beta/pull/10788))
* dialogflow:  added `new_recognition_result_notification_config` field to `google_dialogflow_conversation_profile ` resource ([#10775](https://github.com/hashicorp/terraform-provider-google-beta/pull/10775))
* discoveryengine: added `features` field to `google_discovery_engine_search_engine` resource ([#10762](https://github.com/hashicorp/terraform-provider-google-beta/pull/10762))
* dlp: added `other_cloud_target` and `other_cloud_starting_location` to `google_data_loss_prevention_discovery_config` ([#10773](https://github.com/hashicorp/terraform-provider-google-beta/pull/10773))
* gkebackup: added `backup_config.selected_namespace_labels` field to `google_gke_backup_backup_plan` resource ([#10756](https://github.com/hashicorp/terraform-provider-google-beta/pull/10756))
* looker: added `gemini_enabled` field to `google_looker_instance` resource ([#10771](https://github.com/hashicorp/terraform-provider-google-beta/pull/10771))
* netapp: added `hot_tier_size_used_gib` fields to `google_netapp_volume` ([#10766](https://github.com/hashicorp/terraform-provider-google-beta/pull/10766))
* netapp: added `cold_tier_size_used_gib` and `hot_tier_size_used_gib` fields to `google_netapp_storage_pool` ([#10766](https://github.com/hashicorp/terraform-provider-google-beta/pull/10766))
* networksecurity: added `type` and `mirroring_deployment_groups` fields to `google_network_security_mirroring_endpoint_group` resource ([#10783](https://github.com/hashicorp/terraform-provider-google-beta/pull/10783))
* oracledatabase: added `gcp_oracle_zone` field to `google_oracle_database_odb_network` resource ([#10767](https://github.com/hashicorp/terraform-provider-google-beta/pull/10767))
* privilegedaccessmanager: added `approval_workflow.steps.id` field to `google_privileged_access_manager_entitlement` resource ([#10750](https://github.com/hashicorp/terraform-provider-google-beta/pull/10750))
* pubsub: added support for `tags` field to `google_pubsub_topic` and `google_pubsub_subscription` resources ([#10760](https://github.com/hashicorp/terraform-provider-google-beta/pull/10760))
* sql: added `point_in_time_restore_context` field to `google_sql_database_instance` ([#10786](https://github.com/hashicorp/terraform-provider-google-beta/pull/10786))
* storage: added `force_destroy` field to `google_storage_insights_report_config` resource ([#10772](https://github.com/hashicorp/terraform-provider-google-beta/pull/10772))
* storageinsights: added `activity_data_retention_period_days` field to `google_storage_insights_dataset_config` resource ([#10769](https://github.com/hashicorp/terraform-provider-google-beta/pull/10769))
* vertexai: added `endpoint_config.private_service_connect_config` block to `google_vertex_ai_endpoint_with_model_garden_deployment` resource ([#10754](https://github.com/hashicorp/terraform-provider-google-beta/pull/10754))
* vertexai: added `encryption_spec.kms_key_name` field to `google_vertex_ai_index_endpoint ` resource ([#10787](https://github.com/hashicorp/terraform-provider-google-beta/pull/10787))
* vertexai: added `encryption_spec.kms_key_name` field to `google_vertex_ai_index` resource ([#10759](https://github.com/hashicorp/terraform-provider-google-beta/pull/10759))

BUG FIXES:
* apihub: fixed a permadiff on `config_template` in `google_apihub_plugin` resource ([#10758](https://github.com/hashicorp/terraform-provider-google-beta/pull/10758))
* storage: fixed a panic caused by empty `cors` blocks `google_storage_bucket` resource ([#10781](https://github.com/hashicorp/terraform-provider-google-beta/pull/10781))

## 7.4.0 (September 23, 2025)

DEPRECATIONS:
* compute: deprecated the option to deploy a container during VM creation using the container startup agent in `google_compute_instance`. Use alternative services to run containers on your VMs. Learn more at https://cloud.google.com/compute/docs/containers/migrate-containers. ([#10725](https://github.com/hashicorp/terraform-provider-google-beta/pull/10725))

FEATURES:
* **New Data Source:** `google_artifact_registry_maven_artifact` ([#10718](https://github.com/hashicorp/terraform-provider-google-beta/pull/10718))
* **New Data Source:** `google_compute_interconnect_location` ([#10727](https://github.com/hashicorp/terraform-provider-google-beta/pull/10727))
* **New Resource:** `google_network_services_wasm_plugin` ([#10742](https://github.com/hashicorp/terraform-provider-google-beta/pull/10742))

IMPROVEMENTS:
* compute: added `scheduling.0.skip_guest_os_shutdown` field to `google_compute_instance_template` resource ([#10729](https://github.com/hashicorp/terraform-provider-google-beta/pull/10729))
* compute: added `scheduling.0.skip_guest_os_shutdown` field to `google_compute_instance` resource ([#10729](https://github.com/hashicorp/terraform-provider-google-beta/pull/10729))
* compute: added `scheduling.0.skip_guest_os_shutdown` field to `google_compute_region_instance_template` resource ([#10729](https://github.com/hashicorp/terraform-provider-google-beta/pull/10729))
* compute: added `tunneling_config` field to `google_compute_service_attachment` resource ([#10730](https://github.com/hashicorp/terraform-provider-google-beta/pull/10730))
* container: added `auto_ipam_config` to `google_container_cluster` resource. ([#10737](https://github.com/hashicorp/terraform-provider-google-beta/pull/10737))
* privilegedaccessmanager: added `privileged_access.gcp_iam_access.role_bindings.id` field to `google_privileged_access_manager_entitlement` resource ([#10743](https://github.com/hashicorp/terraform-provider-google-beta/pull/10743))
* storage: added support for `timeouts`  to  `google_storage_bucket_iam_binding`, `google_storage_bucket_iam_member`, `google_storage_bucket_iam_policy` resources ([#10726](https://github.com/hashicorp/terraform-provider-google-beta/pull/10726))

BUG FIXES:
* bigtable: fixed `node_scaling_factor` forcing new instance on `google_bigtable_instance` when adding new cluster ([#10744](https://github.com/hashicorp/terraform-provider-google-beta/pull/10744))
* cloudscheduler: fixed a type assertion panic in `google_cloud_scheduler_job` when processing HTTP headers with nil or unexpected data types ([#10720](https://github.com/hashicorp/terraform-provider-google-beta/pull/10720))
* compute: fixed the `Network field cannot be modified` issue in `google_compute_region_backend_service`. Now updating the `network` field will force the resource to be recreated ([#10738](https://github.com/hashicorp/terraform-provider-google-beta/pull/10738))
* netapp: fixed incorrect default value handling in `google_netapp_volume` for `export_policy.rules` attributes `has_root_access` and `squash_mode`. When not specified, these fields will now take on the API default value with no diff. ([#10736](https://github.com/hashicorp/terraform-provider-google-beta/pull/10736))
* netapp: updated `google_netapp_storage_pool` to source the default value for the `qos_type` field from the API. If not specified in the configuration, `qos_type` will now default to the value provided by the NetApp Volumes API. ([#10735](https://github.com/hashicorp/terraform-provider-google-beta/pull/10735))
* sql: fixed the permadiffs on `disk_size ` when `disk_autoresize` is enabled in `google_sql_database_instance` ([#10739](https://github.com/hashicorp/terraform-provider-google-beta/pull/10739))
* workbench: added retry for `unable to queue the operation` 409 errors in `google_workbench_instance` resource. ([#10733](https://github.com/hashicorp/terraform-provider-google-beta/pull/10733))

## 7.3.0 (September 16, 2025)

FEATURES:
* **New Data Source:** `google_backup_dr_data_source_reference` ([#10707](https://github.com/hashicorp/terraform-provider-google-beta/pull/10707))
* **New Resource:** `google_bigquery_datapolicyv2_data_policy` ([#10693](https://github.com/hashicorp/terraform-provider-google-beta/pull/10693))
* **New Resource:** `google_saas_runtime_release` ([#10685](https://github.com/hashicorp/terraform-provider-google-beta/pull/10685))
* **New Resource:** `google_secure_source_manager_hook` ([#10706](https://github.com/hashicorp/terraform-provider-google-beta/pull/10706))

IMPROVEMENTS:
* cloudrun: added `sub_path` field to `google_cloud_run_service` resource. ([#10705](https://github.com/hashicorp/terraform-provider-google-beta/pull/10705))
* cloudrunv2: added `sub_path` field to `google_cloud_run_v2_service` `google_cloud_run_v2_job` and `google_cloud_run_v2_worker_pool` resource. ([#10705](https://github.com/hashicorp/terraform-provider-google-beta/pull/10705))
* compute: added `labels` and `label_fingerprint` fields to `google_compute_security_policy` resource ([#10696](https://github.com/hashicorp/terraform-provider-google-beta/pull/10696))
* compute: `labels` under `initialize_params` are now updatable on `google_compute_instance` ([#10710](https://github.com/hashicorp/terraform-provider-google-beta/pull/10710))
* container: added new fields `memory_manager` and `topology_manager` to `node_kubelet_config` block ([#10681](https://github.com/hashicorp/terraform-provider-google-beta/pull/10681))
* datastream: added `destination_config.bigquery_destination_config.source_hierarchy_datasets.project_id` field to `google_datastream_stream` resource ([#10704](https://github.com/hashicorp/terraform-provider-google-beta/pull/10704))
* discoveryengine: added `app_type` field to `google_discovery_engine_search_engine` resource ([#10694](https://github.com/hashicorp/terraform-provider-google-beta/pull/10694))
* gkeonprem: added `proxy` field to `google_gkeonprem_vmware_admin_cluster` resource ([#10702](https://github.com/hashicorp/terraform-provider-google-beta/pull/10702))
* healthcare: added `validation_config` to `google_healthcare_fhir_store` resource ([#10700](https://github.com/hashicorp/terraform-provider-google-beta/pull/10700))
* iamworkforcepool: added `extended_attributes` field to `workforce_pool_provider` resource ([#10688](https://github.com/hashicorp/terraform-provider-google-beta/pull/10688))
* netapp: added `export_policy.rules.squash_mode` field to `google_netapp_volume` resource. ([#10711](https://github.com/hashicorp/terraform-provider-google-beta/pull/10711))
* privateca: added `encryption_spec` field to `google_privateca_ca_pool` resource ([#10699](https://github.com/hashicorp/terraform-provider-google-beta/pull/10699))
* run: added `connector`  to `vpcAccess` on `google_cloud_run_v2_worker_pool` resource ([#10701](https://github.com/hashicorp/terraform-provider-google-beta/pull/10701))
* tags: added the `DATA_GOVERNANCE` value to `google_tags_tag_key.purpose` ([#10687](https://github.com/hashicorp/terraform-provider-google-beta/pull/10687))

BUG FIXES:
* bigquery: updated the schema change detection for `google_bigquery_table` to take into account presence of row access policy ([#10683](https://github.com/hashicorp/terraform-provider-google-beta/pull/10683))
* compute: fixed `allow_global_access` to correctly be immutable for `google_compute_forwarding_rule` resources with load balancing scheme of INTERNAL_MANAGED ([#10692](https://github.com/hashicorp/terraform-provider-google-beta/pull/10692))
* compute: fixed a crash in `google_compute_security_policy` due to a changed API response for empty `match.0.expr_options` blocks ([#10715](https://github.com/hashicorp/terraform-provider-google-beta/pull/10715))
* dialogflow: added support for non-global endpoints for `google_dialogflow_conversation_profile` ([#10712](https://github.com/hashicorp/terraform-provider-google-beta/pull/10712))
* publicca: use `RawURLEncoding` instead of `URLEncoding` for unpadded base64 encoding ([#10682](https://github.com/hashicorp/terraform-provider-google-beta/pull/10682))
* secretmanager: fixed a panic in `google_secret_manager_secret_version` in a `secret_manager` ([#10698](https://github.com/hashicorp/terraform-provider-google-beta/pull/10698))
* workbench: fixed issue that resource creation with computed `labels` field fails in `google_workbench_instance` resource ([#10691](https://github.com/hashicorp/terraform-provider-google-beta/pull/10691))
* workbench: made `report-notebook-metrics` metadata key settable for `google_workbench_instance` ([#10690](https://github.com/hashicorp/terraform-provider-google-beta/pull/10690))

## 7.2.0 (September 9, 2025)

FEATURES:
* **New Data Source:** `google_artifact_registry_python_package` ([#10671](https://github.com/hashicorp/terraform-provider-google-beta/pull/10671))
* **New Data Source:** `google_backup_dr_data_source_references` ([#10672](https://github.com/hashicorp/terraform-provider-google-beta/pull/10672))
* **New Resource:** `google_discovery_engine_acl_config` ([#10680](https://github.com/hashicorp/terraform-provider-google-beta/pull/10680))
* **New Resource:** `google_saas_runtime_unit_kind` ([#10652](https://github.com/hashicorp/terraform-provider-google-beta/pull/10652))

IMPROVEMENTS:
* chronicle: made the `scope_info` field in `google_chronicle_reference_list` configurable ([#10663](https://github.com/hashicorp/terraform-provider-google-beta/pull/10663))
* compute: added `header_action` to `path_matcher` and `default_service` level on `google_compute_region_url_map` resource ([#10665](https://github.com/hashicorp/terraform-provider-google-beta/pull/10665))
* container: added `secret_manager_config.rotation_config` field to `google_container_cluster` resource ([#10659](https://github.com/hashicorp/terraform-provider-google-beta/pull/10659))
* container: added new fields `memory_manager` and `topology_manager` to `google_container_cluster.node_config.kubelet_config` and `google_container_node_pool.node_config.kubelet_config` ([#10681](https://github.com/hashicorp/terraform-provider-google-beta/pull/10681))
* healthcare: added `consent_config` field to `google_healthcare_fhir_store` resource ([#10666](https://github.com/hashicorp/terraform-provider-google-beta/pull/10666))
**New Resource:** `google_network_management_organization_vpc_flow_logs_config` ([#10660](https://github.com/hashicorp/terraform-provider-google-beta/pull/10660))
* sql: added `final_backup_description` and `final_backup_config` fields to `google_sql_database_instance` resource ([#10678](https://github.com/hashicorp/terraform-provider-google-beta/pull/10678))
* storage: added `aws_s3_compatible_data_source` to `google_storage_transfer_job` resource ([#10656](https://github.com/hashicorp/terraform-provider-google-beta/pull/10656))

BUG FIXES:
* provider: fixed an issue with `universe_domain` where the provider tried to connect to "googleapis.com" for user email logging when `universe_domain` was set ([#10654](https://github.com/hashicorp/terraform-provider-google-beta/pull/10654))
* container: fixed a faulty diff for arrays on `user_managed_keys_config` that caused faulty cluster updates to be triggered in `google_container_cluster` ([#10668](https://github.com/hashicorp/terraform-provider-google-beta/pull/10668))
* osconfig: fixed permadiff in `google_osconfig_patch_deployment` where `patch_config.yum.minimal` doesn't send `false` for empty values ([#10661](https://github.com/hashicorp/terraform-provider-google-beta/pull/10661))

## 7.1.1 (September 4, 2025)

BUG FIXES:
* bigtable: fixed an error encountered when applying `google_bigtable_table_iam_*` resources after upgrading to 7.x and replacing `instance` with `instance_name` ([#10667](https://github.com/hashicorp/terraform-provider-google-beta/pull/10667))

## 7.1.0 (September 2, 2025)

DEPRECATIONS:
* container: deprecated `enterprise_config` field in `google_container_cluster` resource. GKE Enterprise features are now available without an Enterprise tier. ([#10646](https://github.com/hashicorp/terraform-provider-google-beta/pull/10646))
* storage: removed deprecated status for field to `detect_md5hash`  in  `google_storage_bucket_object` resource ([#10605](https://github.com/hashicorp/terraform-provider-google-beta/pull/10605))

FEATURES:
* **New Data Source:** `google_iap_web_forwarding_rule_service_iam_policy` ([#10621](https://github.com/hashicorp/terraform-provider-google-beta/pull/10621))
* **New Resource:** `google_iap_web_forwarding_rule_service_iam_binding` ([#10621](https://github.com/hashicorp/terraform-provider-google-beta/pull/10621))
* **New Resource:** `google_iap_web_forwarding_rule_service_iam_member` ([#10621](https://github.com/hashicorp/terraform-provider-google-beta/pull/10621))
* **New Resource:** `google_iap_web_forwarding_rule_service_iam_policy` ([#10621](https://github.com/hashicorp/terraform-provider-google-beta/pull/10621))

IMPROVEMENTS:
* artifactregistry: added `registry_uri` as attribute to `google_artifact_registry_repository` ([#10618](https://github.com/hashicorp/terraform-provider-google-beta/pull/10618))
* backupdr: added `create_time` field to `google_backup_dr_backup` data source ([#10626](https://github.com/hashicorp/terraform-provider-google-beta/pull/10626))
* cloudbuild: added `worker_config.enable_nested_virtualization` field to `google_cloudbuild_worker_pool` resource ([#10619](https://github.com/hashicorp/terraform-provider-google-beta/pull/10619))
* cloudrunv2: added support for `multi_region_settings` field to `google_cloud_run_v2_service` resource ([#10607](https://github.com/hashicorp/terraform-provider-google-beta/pull/10607))
* compute: add `params.resource_manager_tags` field to the `google_compute_region_backend_service` ([#10634](https://github.com/hashicorp/terraform-provider-google-beta/pull/10634))
* compute: added `public_delegated_sub_prefixs` field to resource `google_compute_public_delegated_prefix` ([#10638](https://github.com/hashicorp/terraform-provider-google-beta/pull/10638))
* compute: added `update_strategy`  field to `google_compute_network_peering ` resource ([#10623](https://github.com/hashicorp/terraform-provider-google-beta/pull/10623))
* firestore: added `unique` field to `google_firestore_index` resource ([#10617](https://github.com/hashicorp/terraform-provider-google-beta/pull/10617))
* netapp: added  `qos_type` and `available_throughput_mibps` fields to `google_netapp_storage_pool` resource ([#10615](https://github.com/hashicorp/terraform-provider-google-beta/pull/10615))
* netapp: added  `throughput_mibps` field to `google_netapp_volume` resource ([#10615](https://github.com/hashicorp/terraform-provider-google-beta/pull/10615))
* networkservices: allowed `EXPLICIT_ROUTING_MODE` for `routing_mode` on `google_network_services_gateway` resource ([#10608](https://github.com/hashicorp/terraform-provider-google-beta/pull/10608))
* sql: added `consumer_network_status`, `ip_address`, and `status` fields to `psc_auto_connections` field on `google_sql_database_instance` resource ([#10637](https://github.com/hashicorp/terraform-provider-google-beta/pull/10637))
* storagetransfer: added `service_account` field to `google_storage_transfer_job` resource ([#10635](https://github.com/hashicorp/terraform-provider-google-beta/pull/10635))
* storagetransfer: added `transfer_spec.aws_s3_data_source.credentials_secret` to `google_storage_transfer_job` resource ([#10609](https://github.com/hashicorp/terraform-provider-google-beta/pull/10609))

BUG FIXES:
* compute: fixed certain spurious diffs for `google_compute_region_backend_service.backend.group` ([#10611](https://github.com/hashicorp/terraform-provider-google-beta/pull/10611))
* compute: fixed permadiff on `google_compute_region_network_endpoint_group` when no `network` is specified ([#10625](https://github.com/hashicorp/terraform-provider-google-beta/pull/10625))
* memorystore: fixed permadiffs that cause destroy+recreate on new `google_memorystore_instance` when `desired_psc_auto_connections` is set ([#10648](https://github.com/hashicorp/terraform-provider-google-beta/pull/10648))
* netapp: fixed a permadiff on `total_iops` in `google_netapp_storage_pool` resource ([#10643](https://github.com/hashicorp/terraform-provider-google-beta/pull/10643))
* oracledatabase: fixed permadiffs on `google_oracle_database_autonomous_database` resource for the `odb_network` and `odb_subnet` fields ([#10627](https://github.com/hashicorp/terraform-provider-google-beta/pull/10627))

## 7.0.1 (August 27, 2025)

BUG FIXES:
* storage: fixed a conversion crash in `google_storage_bucket` state migration [#10629](https://github.com/hashicorp/terraform-provider-google-beta/pull/10629)

## 7.0.0 (August 26, 2025)

[Terraform Google Provider 7.0.0 Upgrade Guide](https://registry.terraform.io/providers/hashicorp/google-beta/latest/docs/guides/version_7_upgrade)

BREAKING RESOURCE REMOVALS:
* beyondcorp: removed `google_beyondcorp_application`, its associated IAM resources `google_beyondcorp_application_iam_binding`, `google_beyondcorp_application_iam_member`, and `google_beyondcorp_application_iam_policy`, and the `google_beyondcorp_application_iam_policy` datasource. Use `google_beyondcorp_security_gateway_application` instead. [#10536](https://github.com/hashicorp/terraform-provider-google-beta/pull/10536)
* notebooks: removed `google_notebooks_location` [#10350](https://github.com/hashicorp/terraform-provider-google-beta/pull/10350)
* tpu: removed`google_tpu_node`. Use `google_tpu_v2_vm` instead. [#10516](https://github.com/hashicorp/terraform-provider-google-beta/pull/10516)

BREAKING FIELD REMOVALS:
* cloudrunv2: removed `template.containers.depends_on` within `resource google_cloud_run_v2_worker_pool` [#10444](https://github.com/hashicorp/terraform-provider-google-beta/pull/10444)
* colab: removed `post_startup_script_config` field from from `google_colab_runtime_template` resource [#10555](https://github.com/hashicorp/terraform-provider-google-beta/pull/10555)
* compute: removed field `enable_flow_logs` from `google_compute_subnetwork` [#10398](https://github.com/hashicorp/terraform-provider-google-beta/pull/10398)
* gkehub: removed `configmanagement.binauthz` field in `google_gke_hub_feature_membership` [#10585](https://github.com/hashicorp/terraform-provider-google-beta/pull/10585)
* gkehub: removed `description` field in `google_gke_hub_membership` [#10344](https://github.com/hashicorp/terraform-provider-google-beta/pull/10344)
* memorystore: removed `allow_fewer_zones_deployment` field from `google_memorystore_instance` resource because it isn't user-configurable [#10588](https://github.com/hashicorp/terraform-provider-google-beta/pull/10588)
* redis: removed `allow_fewer_zones_deployment` field from `google_redis_cluster` resource because it isn't user-configurable [#10588](https://github.com/hashicorp/terraform-provider-google-beta/pull/10588)
* resourcemanager: removed non-functional `project` field from `google_service_account_key` datasource [#10537](https://github.com/hashicorp/terraform-provider-google-beta/pull/10537)

BREAKING INCREASED VALIDATION:
* cloudfunctions2: made `event_type` a required field for `event_trigger` in `google_cloudfunctions2_function` [#10501](https://github.com/hashicorp/terraform-provider-google-beta/pull/10501)
* networkservices: made `load_balancing_scheme` required in `google_network_services_lb_traffic_extension` [#10419](https://github.com/hashicorp/terraform-provider-google-beta/pull/10419)
* sql: made `password_wo_version` required when `password_wo` is set in `google_sql_user` [#10591](https://github.com/hashicorp/terraform-provider-google-beta/pull/10591)
* storage: added validation requiring the `topic` field to be in the form "projects/<project>/topics/<topic>" in `google_storage_notification` [#10602](https://github.com/hashicorp/terraform-provider-google-beta/pull/10602)
* storagetransfer: added path validation for GCS path source and sink in `google_storage_transfer_job` [#10297](https://github.com/hashicorp/terraform-provider-google-beta/pull/10297)
* vertexai: made `metadata`, and `metadata.config` required in `google_vertex_ai_index`. Resource creation would fail without these attributes already, so no change is necessary to existing configurations. [#10520](https://github.com/hashicorp/terraform-provider-google-beta/pull/10520)

OTHER BREAKING CHANGES:
* provider: fixed many import functions throughout the provider that erroneously matched a subset of the provided input, leading to unclear error messages when using `terraform input` with invalid resource IDs. [#10545](https://github.com/hashicorp/terraform-provider-google-beta/pull/10545)
* alloydb: added `deletion_protection` field with a default value of `true` to `google_alloydb_cluster` resource [#10553](https://github.com/hashicorp/terraform-provider-google-beta/pull/10553)
* apigee: changed `certs_info` field in `google_apigee_keystores_aliases_key_cert_file` to be output-only [#10602](https://github.com/hashicorp/terraform-provider-google-beta/pull/10602)
* apigee: migrated `google_apigee_keystores_aliases_key_cert_file` to the plugin framework [#10602](https://github.com/hashicorp/terraform-provider-google-beta/pull/10602)
* artifactregistry: removed the default values for `public_repository` fields in `google_artifact_registry_repository`. If your state is reliant on them, they will now need to be manually included in your configuration. [#10519](https://github.com/hashicorp/terraform-provider-google-beta/pull/10519)
* bigquery: removed the default value of `view.use_legacy_sql` in `google_bigquery_table` [#10578](https://github.com/hashicorp/terraform-provider-google-beta/pull/10578)
* bigtable: renamed instance to `instance_name` for bigtable_table_iam objects [#10248](https://github.com/hashicorp/terraform-provider-google-beta/pull/10248)
* billing: made `budget_filter.credit types` and `budget_filter.subaccounts` no longer optional+computed, only optional, in `google_billing_budget` resource [#10587](https://github.com/hashicorp/terraform-provider-google-beta/pull/10587)
* cloudfunctions2: changed `service_config.service` field in `google_cloudfunctions2_function` resource to be output-only [#10432](https://github.com/hashicorp/terraform-provider-google-beta/pull/10432)
* compute: `subnetworks` and `instances` fields in `google_compute_packet_mirroring` have been converted from arrays to sets [#10550](https://github.com/hashicorp/terraform-provider-google-beta/pull/10550)
* compute: `advertised_ip_ranges` field group in `google_compute_router` has been converted from a list to a set [#10557](https://github.com/hashicorp/terraform-provider-google-beta/pull/10557)
* compute: `disk.type`, `disk.mode` and `disk.interface` no longer use provider configured default values and instead will be set by the API in `google_compute_instance_template` and `google_compute_region_instance_template` resources [#10569](https://github.com/hashicorp/terraform-provider-google-beta/pull/10569)
* gkehub: updated beta api endpoint from v1beta1 to v1beta [#10344](https://github.com/hashicorp/terraform-provider-google-beta/pull/10344)
* resourcemanager: changed `disable_on_destroy` default value to `false` in `google_project_service` [#10508](https://github.com/hashicorp/terraform-provider-google-beta/pull/10508)
* securesourcemanager: changed `deletion_policy` default value from `DELETE` to `PREVENT` [#10515](https://github.com/hashicorp/terraform-provider-google-beta/pull/10515)
* storage: changed `retention_period` to `string` data type in resource `google_storage_bucket` [#10311](https://github.com/hashicorp/terraform-provider-google-beta/pull/10311)
* storage: migrated `google_storage_notification` to the plugin framework [#10602](https://github.com/hashicorp/terraform-provider-google-beta/pull/10602)

FEATURES:
* **New Data Source:** `google_artifact_registry_npm_package` ([#10582](https://github.com/hashicorp/terraform-provider-google-beta/pull/10582))
* **New Data Source:** `google_certificate_manager_dns_authorization` ([#10544](https://github.com/hashicorp/terraform-provider-google-beta/pull/10544))
* **New Resource:** `google_iap_web_region_forwarding_rule_service_iam_binding` ([#10561](https://github.com/hashicorp/terraform-provider-google-beta/pull/10561))
* **New Resource:** `google_iap_web_region_forwarding_rule_service_iam_member` ([#10561](https://github.com/hashicorp/terraform-provider-google-beta/pull/10561))
* **New Resource:** `google_iap_web_region_forwarding_rule_service_iam_policy` ([#10561](https://github.com/hashicorp/terraform-provider-google-beta/pull/10561))
* **New Resource:** `google_saas_runtime_saas` ([#10556](https://github.com/hashicorp/terraform-provider-google-beta/pull/10556))

IMPROVEMENTS:
* bigquery: added support for "connection_properties" for bigquery to `google_bigquery_job` (beta) ([#10554](https://github.com/hashicorp/terraform-provider-google-beta/pull/10554))
* cloudbuild: added `developer_connect_event_config` field to `google_cloudbuild_trigger` resource ([#10563](https://github.com/hashicorp/terraform-provider-google-beta/pull/10563))
* cloudtasks: added `desired_state` field to `google_cloud_tasks_queue ` resource ([#10567](https://github.com/hashicorp/terraform-provider-google-beta/pull/10567))
* cloudrunv2: added `max_instance_count` field to `google_cloud_run_v2_service` resource. ([#10558](https://github.com/hashicorp/terraform-provider-google-beta/pull/10558))
* compute: added `params.resourceManagerTags` field to the `google_compute_backend_service` ([#10575](https://github.com/hashicorp/terraform-provider-google-beta/pull/10575))
* compute: added `params.resource_manager_tags` field to `google_compute_backend_bucket` ([#10581](https://github.com/hashicorp/terraform-provider-google-beta/pull/10581))
* compute: added `short_name` field to `google_compute_organization_security_policy` resource ([#10572](https://github.com/hashicorp/terraform-provider-google-beta/pull/10572))
* container: added `cluster_autoscaling.default_compute_class_enabled` field to `google_container_cluster` resource ([#10552](https://github.com/hashicorp/terraform-provider-google-beta/pull/10552))
* dialogflowcx: added `enableMultiLanguageTraining`, `locked`, `answerFeedbackSettings`, `personalizationSettings`, `clientCertificateSettings`, `startPlaybook`, `satisfiesPzs`, and `satisfiesPzi` to `google_dialogflow_cx_agent` resource. ([#10543](https://github.com/hashicorp/terraform-provider-google-beta/pull/10543))
* lustre: increased `google_lustre_instance` resource create timeout to 120m from 20m ([#10570](https://github.com/hashicorp/terraform-provider-google-beta/pull/10570))
* oracledatabase: enabled default_from_api flag for ODB Network related fields in `google_oracle_database_cloud_vm_cluster` resource ([#10564](https://github.com/hashicorp/terraform-provider-google-beta/pull/10564))
* sql: added feature to restore `google_sql_database_instance` using `backupdr_backup` ([#10579](https://github.com/hashicorp/terraform-provider-google-beta/pull/10579))
* ssm: made `ca_pool` argument optional for private instances that use Google-managed trusted certificates.` to `secure_source_manager` resource ([#10559](https://github.com/hashicorp/terraform-provider-google-beta/pull/10559))

BUG FIXES:
* container: fixed issue where a failed creation on `google_container_node_pool` would result in an unrecoverable tainted state ([#10586](https://github.com/hashicorp/terraform-provider-google-beta/pull/10586))
* gkeonprem: set `default_from_api` in image field in `google_vmware_node_pool` ([#10551](https://github.com/hashicorp/terraform-provider-google-beta/pull/10551))
* workbench: made `install-monitoring-agent` metadata key settable for `google_workbench_instance` ([#10589](https://github.com/hashicorp/terraform-provider-google-beta/pull/10589))
