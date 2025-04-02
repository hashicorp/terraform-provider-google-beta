## 5.45.2 (Feb 10, 2025)

NOTES:
* 5.45.2 is a backport release. The changes in this release will be available in 6.20.0 and users upgrading to 6.X should upgrade to that version or higher.

BUG FIXES:
* dns: fixed a bug where `google_dns_managed_zone` is unable to update with `service_directory_config` specified ([#9171](https://github.com/hashicorp/terraform-provider-google-beta/pull/9239))

## 5.45.1 (January 29, 2025)

NOTES:
* 5.45.1 is a backport release, responding to a new GKE label being applied that can cause unwanted diffs in node pools. The changes in this release will be available in 6.18.1 and users upgrading to 6.X should upgrade to that version or higher.

BUG FIXES:
* container: fixed a diff caused by server-side set values for `node_config.resource_labels` ([#9171](https://github.com/hashicorp/terraform-provider-google-beta/pull/9171))

## 5.45.0 (November 11, 2024)

NOTES:
* 5.45.0 is a backport release, responding to a new Spanner feature that may result in creation of unwanted backups for users. The changes in this release will be available in 6.11.0 and users upgrading to 6.X should upgrade to that version or higher.

IMPROVEMENTS:
* spanner: added `default_backup_schedule_type` field to  `google_spanner_instance` ([#8644](https://github.com/hashicorp/terraform-provider-google-beta/pull/8644))

## 5.44.2 (October 14, 2024)

Notes:
* 5.44.2 is a backport release, responding to a GKE rollout that created permadiffs for many users. The changes in this release will be available in 6.7.0 and users upgrading to 6.X should upgrade to that version or higher.

IMPROVEMENTS:
* container: `google_container_cluster` will now accept server-specified values for `node_pool_auto_config.0.node_kubelet_config` when it is not defined in configuration and will not detect drift. Note that this means that removing the value from configuration will now preserve old settings instead of reverting the old settings. ([#8385](https://github.com/hashicorp/terraform-provider-google-beta/pull/8385))

BUG FIXES:
* container: fixed a diff triggered by a new API-side default value for `node_config.0.kubelet_config.0.insecure_kubelet_readonly_port_enabled`. Terraform will now accept server-specified values for `node_config.0.kubelet_config` when it is not defined in configuration and will not detect drift. Note that this means that removing the value from configuration will now preserve old settings instead of reverting the old settings. ([#8385](https://github.com/hashicorp/terraform-provider-google-beta/pull/8385))

## 5.44.1 (September 23, 2024)

NOTES:
* 5.44.1 is a backport release, intended to pull in critical container improvements and fixes for issues introduced in 5.44.0

IMPROVEMENTS:
* container: added in-place update support for `gcfs_config` in in `google_container_cluster` and `google_container_node_pool` ([#8101](https://github.com/hashicorp/terraform-provider-google-beta/pull/8101)) ([#8207](https://github.com/hashicorp/terraform-provider-google-beta/pull/8207))

BUG FIXES:
* container: fixed a permadiff on `gcfs_config` in `google_container_cluster` and `google_container_node_pool` ([#8207](https://github.com/hashicorp/terraform-provider-google-beta/pull/8207))
* container: fixed a bug where specifying `node_pool_defaults.node_config_defaults` with `enable_autopilot = true` will cause `google_container_cluster` resource creation failure. ([#8223](https://github.com/hashicorp/terraform-provider-google-beta/pull/8223))

## 5.44.0 (September 9, 2024)

NOTES:
* 5.44.0 is a backport release, intended to pull in critical container improvements from 6.2.0

IMPROVEMENTS:
* container: added `insecure_kubelet_readonly_port_enabled` to `node_pool.node_config.kubelet_config` and `node_config.kubelet_config` in `google_container_node_pool` resource. ([#8071](https://github.com/hashicorp/terraform-provider-google-beta/pull/8071))
* container: added `insecure_kubelet_readonly_port_enabled` to `node_pool_defaults.node_config_defaults`, `node_pool.node_config.kubelet_config`, and `node_config.kubelet_config` in `google_container_cluster` resource. ([#8071](https://github.com/hashicorp/terraform-provider-google-beta/pull/8071))
* container: added `node_pool_auto_config.node_kublet_config.insecure_kubelet_readonly_port_enabled` field to `google_container_cluster`. ([#8076](https://github.com/hashicorp/terraform-provider-google-beta/pull/8076))

## 5.43.1 (August 30, 2024)

NOTES:
* 5.43.1 is a backport release, and some changes will not appear in 6.X series releases until 6.1.0

BUG FIXES:
* pubsub: fixed a validation bug that didn't allow empty filter definitions for `google_pubsub_subscription` resources ([#8055](https://github.com/hashicorp/terraform-provider-google-beta/pull/8055))

## 5.43.0 (August 26, 2024)

DEPRECATIONS:
* storage: deprecated `lifecycle_rule.condition.no_age` field in `google_storage_bucket`. Use the new `lifecycle_rule.condition.send_age_if_zero` field instead. ([#7994](https://github.com/hashicorp/terraform-provider-google-beta/pull/7994))

FEATURES:
* **New Resource:** `google_kms_ekm_connection_iam_binding` ([#7969](https://github.com/hashicorp/terraform-provider-google-beta/pull/7969))
* **New Resource:** `google_kms_ekm_connection_iam_member` ([#7969](https://github.com/hashicorp/terraform-provider-google-beta/pull/7969))
* **New Resource:** `google_kms_ekm_connection_iam_policy` ([#7969](https://github.com/hashicorp/terraform-provider-google-beta/pull/7969))
* **New Resource:** `google_scc_v2_organization_scc_big_query_exports` ([#8002](https://github.com/hashicorp/terraform-provider-google-beta/pull/8002))

IMPROVEMENTS:
* compute: exposed service side id as new output field `forwarding_rule_id` on resource `google_compute_forwarding_rule` ([#7972](https://github.com/hashicorp/terraform-provider-google-beta/pull/7972))
* container: added EXTENDED as a valid option for `release_channel` field in `google_container_cluster` resource ([#7973](https://github.com/hashicorp/terraform-provider-google-beta/pull/7973))
* logging: changed `enable_analytics` parsing to "no preference" in analytics if omitted, instead of explicitly disabling analytics in `google_logging_project_bucket_config`. ([#7964](https://github.com/hashicorp/terraform-provider-google-beta/pull/7964))
* networkservices: added `idle_timeout` field to the `google_network_services_tcp_route` resource ([#7996](https://github.com/hashicorp/terraform-provider-google-beta/pull/7996))
* pusbub: added validation to `filter` field in resource `google_pubsub_subscription` ([#7968](https://github.com/hashicorp/terraform-provider-google-beta/pull/7968))
* resourcemanager: added `default_labels` field to `google_client_config` data source ([#7992](https://github.com/hashicorp/terraform-provider-google-beta/pull/7992))
* vmwareengine: added PC undelete support in `google_vmwareengine_private_cloud` ([#8005](https://github.com/hashicorp/terraform-provider-google-beta/pull/8005))

BUG FIXES:
* alloydb: fixed a permadiff on `psc_instance_config` in `google_alloydb_instance` resource ([#7975](https://github.com/hashicorp/terraform-provider-google-beta/pull/7975))
* compute: fixed a malformed URL that affected updating the `server_tls_policy` property on `google_compute_target_https_proxy` resources ([#7988](https://github.com/hashicorp/terraform-provider-google-beta/pull/7988))
* compute: fixed force diff replacement logic for `network_ip` on resource `google_compute_instance` ([#7971](https://github.com/hashicorp/terraform-provider-google-beta/pull/7971))

## 5.42.0 (August 19, 2024)
DEPRECATIONS:
* compute: setting `google_compute_subnetwork.secondary_ip_range = []` to explicitly set a list of empty objects is deprecated and will produce an error in the upcoming major release. Use `send_secondary_ip_range_if_empty` while removing `secondary_ip_range` from config instead. ([#7961](https://github.com/hashicorp/terraform-provider-google-beta/pull/7961))


FEATURES:
* **New Data Source:** `google_artifact_registry_locations` ([#7922](https://github.com/hashicorp/terraform-provider-google-beta/pull/7922))
* **New Data Source:** `google_cloud_identity_transitive_group_memberships` ([#7917](https://github.com/hashicorp/terraform-provider-google-beta/pull/7917))
* **New Resource:** `google_discovery_engine_schema` ([#7963](https://github.com/hashicorp/terraform-provider-google-beta/pull/7963))
* **New Resource:** `google_scc_folder_notification_config` ([#7928](https://github.com/hashicorp/terraform-provider-google-beta/pull/7928))
* **New Resource:** `google_scc_v2_folder_notification_config` ([#7927](https://github.com/hashicorp/terraform-provider-google-beta/pull/7927))
* **New Resource:** `google_vertex_ai_index_endpoint_deployed_index` ([#7931](https://github.com/hashicorp/terraform-provider-google-beta/pull/7931))

IMPROVEMENTS:
* clouddeploy: added `serial_pipeline.stages.strategy.canary.runtime_config.kubernetes.gateway_service_mesh.pod_selector_label` and `serial_pipeline.stages.strategy.canary.runtime_config.kubernetes.service_networking.pod_selector_label` fields to `google_clouddeploy_delivery_pipeline` resource ([#7945](https://github.com/hashicorp/terraform-provider-google-beta/pull/7945))
* compute: added `TDX` instance option to `confidential_instance_type` instance in `google_compute_instance` ([#7913](https://github.com/hashicorp/terraform-provider-google-beta/pull/7913))
* compute: added `send_secondary_ip_range_if_empty` to `google_compute_subnetwork` ([#7961](https://github.com/hashicorp/terraform-provider-google-beta/pull/7961))
* discoveryengine: added `skip_default_schema_creation` field to `google_data_store` resource ([#7900](https://github.com/hashicorp/terraform-provider-google-beta/pull/7900))
* dns: changed `load_balancer_type` field from required to optional in `google_dns_record_set` ([#7925](https://github.com/hashicorp/terraform-provider-google-beta/pull/7925))
* parallelstore: added `file_stripe_level`, `directory_stripe_level` fields to `google_parallelstore_instance` resource ([#7942](https://github.com/hashicorp/terraform-provider-google-beta/pull/7942))
* servicenetworking: added `update_on_creation_fail` field to `google_service_networking_connection` resource. When it is set to true, enforce an update of the reserved peering ranges on the existing service networking connection in case of a new connection creation failure. ([#7915](https://github.com/hashicorp/terraform-provider-google-beta/pull/7915))
* sql: added `server_ca_mode` field to `google_sql_database_instance` resource ([#7886](https://github.com/hashicorp/terraform-provider-google-beta/pull/7886))

BUG FIXES:
* bigquery: made `google_bigquery_dataset_iam_member` non-authoritative. To remove a bigquery dataset iam member, use an authoritative resource like `google_bigquery_dataset_iam_policy` ([#7960](https://github.com/hashicorp/terraform-provider-google-beta/pull/7960))
* cloudfunctions2: fixed a "Provider produced inconsistent final plan" bug affecting the `service_config.environment_variables` field in `google_cloudfunctions2_function` resource ([#7905](https://github.com/hashicorp/terraform-provider-google-beta/pull/7905))
* cloudfunctions2: fixed a permadiff on `storage_source.generation` in `google_cloudfunctions2_function` resource ([#7912](https://github.com/hashicorp/terraform-provider-google-beta/pull/7912))
* compute: fixed issue where sub-resources managed by `google_compute_forwarding_rule` prevented resource deletion ([#7958](https://github.com/hashicorp/terraform-provider-google-beta/pull/7958))
* logging: changed `google_logging_project_bucket_config.enable_analytics` behavior to set "no preference" in analytics if omitted, instead of explicitly disabling analytics. ([#19126](https://github.com/hashicorp/terraform-provider-google/pull/19126))
* workbench: fixed a bug with `google_workbench_instance`  metadata drifting when using custom containers. ([#7959](https://github.com/hashicorp/terraform-provider-google-beta/pull/7959))

## 5.41.0 (August 13, 2024)

DEPRECATIONS:
* resourcemanager: deprecated `skip_delete` field in the `google_project` resource. Use `deletion_policy` instead. ([#7817](https://github.com/hashicorp/terraform-provider-google-beta/pull/7817))

FEATURES:
* **New Data Source:** `google_scc_v2_organization_source_iam_policy` ([#7888](https://github.com/hashicorp/terraform-provider-google-beta/pull/7888))
* **New Resource:** `google_access_context_manager_service_perimeter_dry_run_egress_policy` ([#7882](https://github.com/hashicorp/terraform-provider-google-beta/pull/7882))
* **New Resource:** `google_access_context_manager_service_perimeter_dry_run_ingress_policy` ([#7882](https://github.com/hashicorp/terraform-provider-google-beta/pull/7882))
* **New Resource:** `google_scc_v2_folder_mute_config` ([#7846](https://github.com/hashicorp/terraform-provider-google-beta/pull/7846))
* **New Resource:** `google_scc_v2_project_mute_config` ([#7881](https://github.com/hashicorp/terraform-provider-google-beta/pull/7881))
* **New Resource:** `google_scc_v2_project_notification_config` ([#7892](https://github.com/hashicorp/terraform-provider-google-beta/pull/7892))
* **New Resource:** `google_scc_v2_organization_source` ([#7888](https://github.com/hashicorp/terraform-provider-google-beta/pull/7888))
* **New Resource:** `google_scc_v2_organization_source_iam_binding` ([#7888](https://github.com/hashicorp/terraform-provider-google-beta/pull/7888))
* **New Resource:** `google_scc_v2_organization_source_iam_member` ([#7888](https://github.com/hashicorp/terraform-provider-google-beta/pull/7888))
* **New Resource:** `google_scc_v2_organization_source_iam_policy` ([#7888](https://github.com/hashicorp/terraform-provider-google-beta/pull/7888))

IMPROVEMENTS:
* clouddeploy: added `gke.proxy_url` field to `google_clouddeploy_target` ([#7899](https://github.com/hashicorp/terraform-provider-google-beta/pull/7899))
* cloudrunv2: added field `binary_authorization.policy` to resource `google_cloud_run_v2_job` and resource `google_cloud_run_v2_service` to support named binary authorization policy. ([#7883](https://github.com/hashicorp/terraform-provider-google-beta/pull/7883))
* compute: added update-in-place support for the `google_compute_target_https_proxy.server_tls_policy` field ([#7884](https://github.com/hashicorp/terraform-provider-google-beta/pull/7884))
* compute: added update-in-place support for the `google_compute_region_target_https_proxy.server_tls_policy` field ([#7891](https://github.com/hashicorp/terraform-provider-google-beta/pull/7891))
* container: added `auto_provisioning_locations` field to `google_container_cluster` ([#7849](https://github.com/hashicorp/terraform-provider-google-beta/pull/7849))
* dataform: added `kms_key_name` field to `google_dataform_repository` resource ([#7855](https://github.com/hashicorp/terraform-provider-google-beta/pull/7855))
* discoveryengine: added `skip_default_schema_creation` field to `google_discovery_engine_data_store` resource ([#7900](https://github.com/hashicorp/terraform-provider-google-beta/pull/7900))
* gkehub: added `configmanagement.management` and `configmanagement.config_sync.enabled` fields to `google_gkehub_feature_membership` ([#7899](https://github.com/hashicorp/terraform-provider-google-beta/pull/7899))
* gkehub: added `management` field to `google_gke_hub_feature.fleet_default_member_config.configmanagement` ([#7862](https://github.com/hashicorp/terraform-provider-google-beta/pull/7862))
* resourcemanager: added `deletion_policy` field to the `google_project` resource. Setting `deletion_policy` to `PREVENT` will protect the project against any destroy actions caused by a terraform apply or terraform destroy. Setting `deletion_policy` to `ABANDON` allows the resource to be abandoned rather than deleted and it behaves the same with `skip_delete = true`. Default value is `DELETE`. `skip_delete = true` takes precedence over `deletion_policy = "DELETE"`.
* storage: added `force_destroy` field to `google_storage_managed_folder` resource ([#7867](https://github.com/hashicorp/terraform-provider-google-beta/pull/7867))
* storage: added `generation` field to `google_storage_bucket_object` resource ([#7866](https://github.com/hashicorp/terraform-provider-google-beta/pull/7866))

BUG FIXES:
* compute: fixed `google_compute_instance.alias_ip_range` update behavior to avoid temporarily deleting unchanged alias IP ranges ([#7898](https://github.com/hashicorp/terraform-provider-google-beta/pull/7898))
* compute: fixed the bug that creation of PSC forwarding rules fails in `google_compute_forwarding_rule` resource when provider default labels are set ([#7873](https://github.com/hashicorp/terraform-provider-google-beta/pull/7873))
* sql: fixed a perma-diff in `settings.insights_config` in `google_sql_database_instance` ([#7861](https://github.com/hashicorp/terraform-provider-google-beta/pull/7861))




## 5.40.0 (August 5, 2024)

IMPROVEMENTS:
* bigquery: added support for value `DELTA_LAKE` to `source_format` in `google_bigquery_table` resource ([#7841](https://github.com/hashicorp/terraform-provider-google-beta/pull/7841))
* compute: added `access_mode` field to `google_compute_disk` resource ([#7813](https://github.com/hashicorp/terraform-provider-google-beta/pull/7813))
* compute: added `stack_type`, and `gateway_ip_version` fields to `google_compute_router` resource ([#7801](https://github.com/hashicorp/terraform-provider-google-beta/pull/7801))
* container: added field `ray_operator_config` for `resource_container_cluster` ([#7795](https://github.com/hashicorp/terraform-provider-google-beta/pull/7795))
* monitoring: updated `goal` field to accept a max threshold of up to 0.9999 in `google_monitoring_slo` resource to 0.9999 ([#7807](https://github.com/hashicorp/terraform-provider-google-beta/pull/7807))
* networkconnectivity: added `export_psc` field to `google_network_connectivity_hub` resource ([#7816](https://github.com/hashicorp/terraform-provider-google-beta/pull/7816))
* sql: added `enable_dataplex_integration` field to `google_sql_database_instance` resource ([#7810](https://github.com/hashicorp/terraform-provider-google-beta/pull/7810))

BUG FIXES:
* bigquery: fixed a permadiff when handling "assets" in `params` in the `google_bigquery_data_transfer_config` resource ([#7833](https://github.com/hashicorp/terraform-provider-google-beta/pull/7833))
* bigquery: fixed an issue preventing certain keys in `params` from being assigned values in `google_bigquery_data_transfer_config` ([#7828](https://github.com/hashicorp/terraform-provider-google-beta/pull/7828))
* compute: fixed perma-diff in `google_compute_router` ([#7818](https://github.com/hashicorp/terraform-provider-google-beta/pull/7818))
* container: fixed perma-diff on `node_config.guest_accelerator.gpu_driver_installation_config` field in GKE 1.30+ in `google_container_node_pool` resource ([#7799](https://github.com/hashicorp/terraform-provider-google-beta/pull/7799))
* sql: fixed a perma-diff in `settings.insights_config` in `google_sql_database_instance` ([#7861](https://github.com/hashicorp/terraform-provider-google-beta/pull/7861))

## 5.39.1 (July 30, 2024)

BUG FIXES:
* datastream: fixed a breaking change in 5.39.0 `google_datastream_stream` that made one of `destination_config.0.bigquery_destination_config.0.merge` or `destination_config.0.bigquery_destination_config.0.append_only` required ([#7835](https://github.com/hashicorp/terraform-provider-google-beta/pull/7835))

## 5.39.0 (July 29, 2024)

NOTES:
* networkconnectivity: migrated `google_network_connectivity_hub` from DCL to MMv1 ([#7724](https://github.com/hashicorp/terraform-provider-google-beta/pull/7724))
* networkconnectivity: migrated `google_network_connectivity_spoke` from DCL to MMv1 ([#7762](https://github.com/hashicorp/terraform-provider-google-beta/pull/7762))

DEPRECATIONS:
* bigquery: deprecated `allow_resource_tags_on_deletion` in `google_bigquery_table`. ([#7782](https://github.com/hashicorp/terraform-provider-google-beta/pull/7782))
* bigqueryreservation: deprecated `multi_region_auxiliary` on `google_bigquery_reservation`. ([#7778](https://github.com/hashicorp/terraform-provider-google-beta/pull/7778))
* datastore: deprecated the resource `google_datastore_index`. Use the `google_firestore_index` resource instead. ([#7764](https://github.com/hashicorp/terraform-provider-google-beta/pull/7764))

FEATURES:
* **New Resource:** `google_apigee_environment_keyvaluemaps_entries` ([#7717](https://github.com/hashicorp/terraform-provider-google-beta/pull/7717))
* **New Resource:** `google_apigee_environment_keyvaluemaps` ([#7717](https://github.com/hashicorp/terraform-provider-google-beta/pull/7717))
* **New Resource:** `google_compute_resize_request` ([#7725](https://github.com/hashicorp/terraform-provider-google-beta/pull/7725))
* **New Resource:** `google_compute_router_route_policy` ([#7748](https://github.com/hashicorp/terraform-provider-google-beta/pull/7748))
* **New Resource:** `google_scc_v2_organization_mute_config` ([#7744](https://github.com/hashicorp/terraform-provider-google-beta/pull/7744))

IMPROVEMENTS:
* alloydb: added `observability_config` field to `google_alloydb_instance` resource ([#7737](https://github.com/hashicorp/terraform-provider-google-beta/pull/7737))
* bigquery: added `resource_tags` field to `google_bigquery_table` resource ([#7735](https://github.com/hashicorp/terraform-provider-google-beta/pull/7735))
* bigtable: added `data_boost_isolation_read_only` and `data_boost_isolation_read_only.compute_billing_owner` fields to `google_bigtable_app_profile` resource ([#7789](https://github.com/hashicorp/terraform-provider-google-beta/pull/7789))
* cloudfunctions: added `build_service_account` field to `google_cloudfunctions_function` resource ([#7713](https://github.com/hashicorp/terraform-provider-google-beta/pull/7713))
* compute: added `aws_v4_authentication` field to `google_compute_backend_service` resource ([#7775](https://github.com/hashicorp/terraform-provider-google-beta/pull/7775))
* compute: added `custom_learned_ip_ranges` and `custom_learned_route_priority` fields to `google_compute_router_peer` resource ([#7727](https://github.com/hashicorp/terraform-provider-google-beta/pull/7727))
* compute: added `export_policies` and `import_policies` fields  to `google_compute_router_peer` resource ([#7748](https://github.com/hashicorp/terraform-provider-google-beta/pull/7748))
* compute: added `shared_secret` field to `google_compute_public_advertised_prefix` resource ([#7767](https://github.com/hashicorp/terraform-provider-google-beta/pull/7767))
* compute: added `storage_pool` under `boot_disk.initialize_params` to `google_compute_instance` resource ([#7787](https://github.com/hashicorp/terraform-provider-google-beta/pull/7787))
* compute: changed `target_service` field on the `google_compute_service_attachment` resource to accept a `ForwardingRule` or `Gateway` URL. ([#7736](https://github.com/hashicorp/terraform-provider-google-beta/pull/7736))
* container: added field `ray_operator_config` for `google_container_cluster` ([#7795](https://github.com/hashicorp/terraform-provider-google-beta/pull/7795))
* datastream: added `merge` and `append_only` fields to `google_datastream_stream` resource ([#7726](https://github.com/hashicorp/terraform-provider-google-beta/pull/7726))
* dlp: added `cloud_storage_target` field to `google_data_loss_prevention_discovery_config` resource ([#7734](https://github.com/hashicorp/terraform-provider-google-beta/pull/7734))
* resourcemanager: added `check_if_service_has_usage_on_destroy` field to `google_project_service` resource ([#7745](https://github.com/hashicorp/terraform-provider-google-beta/pull/7745))
* resourcemanager: added the `member` property to `google_project_service_identity` ([#7708](https://github.com/hashicorp/terraform-provider-google-beta/pull/7708))
* vmwareengine: added `deletion_delay_hours` field to `google_vmwareengine_private_cloud` resource ([#7710](https://github.com/hashicorp/terraform-provider-google-beta/pull/7710))
* vmwareengine: supported type change from `TIME_LIMITED` to `STANDARD` for multi-node `google_vmwareengine_private_cloud` resource ([#7710](https://github.com/hashicorp/terraform-provider-google-beta/pull/7710))
* workbench: added `access_configs` to `google_workbench_instance` resource ([#7732](https://github.com/hashicorp/terraform-provider-google-beta/pull/7732))

BUG FIXES:
* compute: fixed perma-diff for `interconnect_type` being `DEDICATED` in `google_compute_interconnect` resource ([#7750](https://github.com/hashicorp/terraform-provider-google-beta/pull/7750))
* dialogflowcx: fixed intermittent issues with retrieving resource state soon after creating `google_dialogflow_cx_security_settings` resources ([#7772](https://github.com/hashicorp/terraform-provider-google-beta/pull/7772))
* firestore: fixed missing import of `field` for `google_firestore_field`. ([#7757](https://github.com/hashicorp/terraform-provider-google-beta/pull/7757))
* firestore: fixed bug where fields `database`, `collection`, `document_id`, and `field` could not be updated on `google_firestore_document` and `google_firestore_field` resources. ([#7791](https://github.com/hashicorp/terraform-provider-google-beta/pull/7791))
* netapp: made the `smb_settings` field on the `google_netapp_volume` resource default to the value returned from the API. This solves permadiffs when the field is unset. ([#7770](https://github.com/hashicorp/terraform-provider-google-beta/pull/7770))
* networksecurity: added recreate functionality on update for `client_validation_mode` and `client_validation_trust_config` in  `google_network_security_server_tls_policy` ([#7756](https://github.com/hashicorp/terraform-provider-google-beta/pull/7756))

## 5.38.0 (July 15, 2024)

FEATURES:
* **New Data Source:** `google_gke_hub_membership_binding` ([#7696](https://github.com/hashicorp/terraform-provider-google-beta/pull/7696))
* **New Data Source:** `google_site_verification_token` ([#7704](https://github.com/hashicorp/terraform-provider-google-beta/pull/7704))
* **New Resource:** `google_scc_project_notification_config` ([#7698](https://github.com/hashicorp/terraform-provider-google-beta/pull/7698))

IMPROVEMENTS:
* cloudkms: added `key_access_justifications_policy` field to `google_kms_crypto_key` resource ([#7693](https://github.com/hashicorp/terraform-provider-google-beta/pull/7693))
* compute: made the `google_compute_resource_policy` resource updatable in-place ([#7692](https://github.com/hashicorp/terraform-provider-google-beta/pull/7692))
* vertexai: added `project_number` field to `google_vertex_ai_feature_online_store_featureview` resource ([#7680](https://github.com/hashicorp/terraform-provider-google-beta/pull/7680))

BUG FIXES:
* cloudfunctions2: fixed permadiffs on `service_config.environment_variables` field in `google_cloudfunctions2_function` resource ([#7684](https://github.com/hashicorp/terraform-provider-google-beta/pull/7684))
* networksecurity: fixed permadiffs on `purpose` field in `google_network_security_address_group` resource ([#7687](https://github.com/hashicorp/terraform-provider-google-beta/pull/7687))

## 5.37.0 (July 8, 2024)

FEATURES:
* **New Data Source:** `google_kms_crypto_keys` ([#7656](https://github.com/hashicorp/terraform-provider-google-beta/pull/7656))
* **New Data Source:** `google_kms_key_rings` ([#7662](https://github.com/hashicorp/terraform-provider-google-beta/pull/7662))
* **New Resource:** `google_scc_v2_organization_notification_config` ([#7649](https://github.com/hashicorp/terraform-provider-google-beta/pull/7649))
* **New Resource:** `google_secure_source_manager_repository` ([#7634](https://github.com/hashicorp/terraform-provider-google-beta/pull/7634))
* **New Resource:** `google_storage_managed_folder_iam` ([#7620](https://github.com/hashicorp/terraform-provider-google-beta/pull/7620))
* **New Resource:** `google_storage_managed_folder` ([#7620](https://github.com/hashicorp/terraform-provider-google-beta/pull/7620))

IMPROVEMENTS:
* certificatemanager: added `allowlisted_certificates` field to `google_certificate_manager_trust_config` resource ([#7643](https://github.com/hashicorp/terraform-provider-google-beta/pull/7643))
* compute: added `source_regions` field to `google_compute_healthcheck` resource ([#7647](https://github.com/hashicorp/terraform-provider-google-beta/pull/7647))
* dataplex: added `sql_assertion` field to `google_dataplex_datascan` resource ([#7623](https://github.com/hashicorp/terraform-provider-google-beta/pull/7623))
* gkehub: added `fleet_default_member_config.configmanagement.config_sync.enabled` field to `google_gke_hub_feature` resource ([#7639](https://github.com/hashicorp/terraform-provider-google-beta/pull/7639))
* netapp: added `zone` and `replica_zone` field to `google_netapp_storage_pool` resource ([#7660](https://github.com/hashicorp/terraform-provider-google-beta/pull/7660))
* networksecurity: added `purpose` field to `google_network_security_address_group` resource ([#7677](https://github.com/hashicorp/terraform-provider-google-beta/pull/7677))
* vertexai: added `project_number` field to `google_vertex_ai_feature_online_store_featureview` resource ([#7680](https://github.com/hashicorp/terraform-provider-google-beta/pull/7680))
* workstations: added `host.gce_instance.vm_tags` field to `google_workstations_workstation_config` resource ([#7644](https://github.com/hashicorp/terraform-provider-google-beta/pull/7644))

BUG FIXES:
* compute: fixed a bug preventing the creation of `google_compute_autoscaler` and `google_compute_region_autoscaler` resources if both `autoscaling_policy.max_replicas` and `autoscaling_policy.min_replicas` were configured as zero. ([#7658](https://github.com/hashicorp/terraform-provider-google-beta/pull/7658))
* resourcemanager: mitigated eventual consistency issues by adding a 10s wait after `google_service_account_key` resource creation ([#7629](https://github.com/hashicorp/terraform-provider-google-beta/pull/7629))
* vertexai: fixed issue where updating "metadata" field could fail in `google_vertex_ai_index` resource ([#7675](https://github.com/hashicorp/terraform-provider-google-beta/pull/7675))

## 5.36.0 (July 1, 2024)

FEATURES:
* **New Resource:** `google_storage_managed_folder_iam` ([#7620](https://github.com/hashicorp/terraform-provider-google-beta/pull/7620))
* **New Resource:** `google_storage_managed_folder` ([#7620](https://github.com/hashicorp/terraform-provider-google-beta/pull/7620))

IMPROVEMENTS:
* bigtable: added `ignore_warnings` field to `google_bigtable_gc_policy` resource ([#7571](https://github.com/hashicorp/terraform-provider-google-beta/pull/7571))
* cloudfunctions2: added `build_config.automatic_update_policy` and `build_config.on_deploy_update_policy` to `google_cloudfunctions2_function` resource ([#7608](https://github.com/hashicorp/terraform-provider-google-beta/pull/7608))
* compute: added `tls_early_data` field to `google_compute_target_https_proxy` resource ([#7588](https://github.com/hashicorp/terraform-provider-google-beta/pull/7588))
* compute: added `custom_error_response_policy` and `default_custom_error_response_policy` fields to `google_compute_url_map` resource ([#7587](https://github.com/hashicorp/terraform-provider-google-beta/pull/7587))
* datafusion: added `connection_type` and `private_service_connect_config` fields to `google_data_fusion_instance` resource ([#7598](https://github.com/hashicorp/terraform-provider-google-beta/pull/7598))
* firebasehosting: added support for `google_firebase_hosting_site` resource to be used for an existing site without using import ([#7594](https://github.com/hashicorp/terraform-provider-google-beta/pull/7594))
* healthcare: added `encryption_spec` field to `google_healthcare_dataset` resource ([#7601](https://github.com/hashicorp/terraform-provider-google-beta/pull/7601))
* monitoring: added `links` field to `google_monitoring_alert_policy` resource ([#7616](https://github.com/hashicorp/terraform-provider-google-beta/pull/7616))
* vertexai: added update support for `big_query.entity_id_columns` field on `google_vertex_ai_feature_group` resource ([#7572](https://github.com/hashicorp/terraform-provider-google-beta/pull/7572))

BUG FIXES:
* accesscontextmanager: fixed perma-diff caused by ordering of `service_perimeters` in `google_access_context_manager_service_perimeters` resource ([#7595](https://github.com/hashicorp/terraform-provider-google-beta/pull/7595))
* compute: fixed a crash in `google_compute_reservation` resource when `share_settings` field has changes ([#7577](https://github.com/hashicorp/terraform-provider-google-beta/pull/7577))
* compute: fixed issue in `google_compute_instance` resource where `service_account` is not set when specifying `service_account.email` and no `service_account.scopes` ([#7596](https://github.com/hashicorp/terraform-provider-google-beta/pull/7596))
* gkehub2: fixed `google_gke_hub_feature` resource to allow `fleet_default_member_config` field to be unset ([#7568](https://github.com/hashicorp/terraform-provider-google-beta/pull/7568))
* identityplatform: fixed perma-diff on `google_identity_platform_config` resource when `sms_region_config` is not set ([#7607](https://github.com/hashicorp/terraform-provider-google-beta/pull/7607))
* logging: fixed perma-diff on `index_configs` in `google_logging_organization_bucket_config` resource ([#7579](https://github.com/hashicorp/terraform-provider-google-beta/pull/7579))

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
* **New Resource:** `google_scc_management_project_security_health_analytics_custom_module` ([#7489](https://github.com/hashicorp/terraform-provider-google-beta/pull/7489))
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
