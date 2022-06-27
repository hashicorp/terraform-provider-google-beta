## 4.28.0 (Unreleased)

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
