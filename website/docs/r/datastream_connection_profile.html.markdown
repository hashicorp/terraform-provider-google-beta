---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/datastream/ConnectionProfile.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Datastream"
description: |-
  A set of reusable connection configurations to be used as a source or destination for a stream.
---

# google_datastream_connection_profile

A set of reusable connection configurations to be used as a source or destination for a stream.


To get more information about ConnectionProfile, see:

* [API documentation](https://cloud.google.com/datastream/docs/reference/rest/v1/projects.locations.connectionProfiles)
* How-to Guides
    * [Official Documentation](https://cloud.google.com/datastream/docs/create-connection-profiles)

~> **Warning:** All arguments including the following potentially sensitive
values will be stored in the raw state as plain text: `oracle_profile.password`, `mysql_profile.password`, `mysql_profile.ssl_config.client_key`, `mysql_profile.ssl_config.client_certificate`, `mysql_profile.ssl_config.ca_certificate`, `postgresql_profile.password`, `sql_server_profile.password`, `forward_ssh_connectivity.password`, `forward_ssh_connectivity.private_key`.
[Read more about sensitive data in state](https://www.terraform.io/language/state/sensitive-data).

<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=datastream_connection_profile_basic&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Datastream Connection Profile Basic


```hcl
resource "google_datastream_connection_profile" "default" {
	display_name          = "Connection profile"
	location              = "us-central1"
	connection_profile_id = "my-profile"

	gcs_profile {
		bucket    = "my-bucket"
		root_path = "/path"
	}
}
```
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=datastream_connection_profile_postgresql_private_connection&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Datastream Connection Profile Postgresql Private Connection


```hcl
resource "google_compute_network" "default" {
    name = "my-network"
    auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  name          = "my-subnetwork"
  ip_cidr_range = "10.1.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.default.id
}

resource "google_datastream_private_connection" "private_connection" {
    display_name          = "Private connection"
    location              = "us-central1"
    private_connection_id = "my-connection"

    vpc_peering_config {
        vpc = google_compute_network.default.id
        subnet = "10.0.0.0/29"
    }
}

resource "google_sql_database_instance" "instance" {
    name             = "my-instance"
    database_version = "POSTGRES_14"
    region           = "us-central1"
    settings {
        tier = "db-f1-micro"
        ip_configuration {
            authorized_networks {
                value = google_compute_address.nat_vm_ip.address
            }
        }
    }

    deletion_protection  = true
}

resource "google_sql_database" "db" {
    instance = google_sql_database_instance.instance.name
    name     = "db"
}

resource "random_password" "pwd" {
    length = 16
    special = false
}

resource "google_sql_user" "user" {
    name = "user"
    instance = google_sql_database_instance.instance.name
    password = random_password.pwd.result
}

resource "google_compute_address" "nat_vm_ip" {
  name         = "nat-vm-ip"
}

resource "google_compute_instance" "nat_vm" {
  name           = "nat-vm"
  machine_type   = "e2-medium"
  zone           = "us-central1-a"
  desired_status  = "RUNNING"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-12"
    }
  }

  network_interface {
    network     = google_datastream_private_connection.private_connection.vpc_peering_config.0.vpc
    subnetwork  = google_compute_subnetwork.default.self_link
    access_config {
        nat_ip = google_compute_address.nat_vm_ip.address
    }
  }

  metadata_startup_script = <<EOT
#! /bin/bash
# See https://cloud.google.com/datastream/docs/private-connectivity#set-up-reverse-proxy
export DB_ADDR=${google_sql_database_instance.instance.public_ip_address}
export DB_PORT=5432
echo 1 > /proc/sys/net/ipv4/ip_forward
md_url_prefix="http://169.254.169.254/computeMetadata/v1/instance"
vm_nic_ip="$(curl -H "Metadata-Flavor: Google" $${md_url_prefix}/network-interfaces/0/ip)"
iptables -t nat -F
iptables -t nat -A PREROUTING \
     -p tcp --dport $DB_PORT \
     -j DNAT \
     --to-destination $DB_ADDR
iptables -t nat -A POSTROUTING \
     -p tcp --dport $DB_PORT \
     -j SNAT \
     --to-source $vm_nic_ip
iptables-save
EOT
}

resource "google_compute_firewall" "rules" {
  name        = "ingress-rule"
  network     = google_datastream_private_connection.private_connection.vpc_peering_config.0.vpc
  description = "Allow traffic into NAT VM"
  direction   = "INGRESS"

  allow {
    protocol = "tcp"
    ports    = ["5432"]
  }

  source_ranges = [google_datastream_private_connection.private_connection.vpc_peering_config.0.subnet]
}

resource "google_datastream_connection_profile" "default" {
    display_name          = "Connection profile"
    location              = "us-central1"
    connection_profile_id = "my-profile"

    postgresql_profile {
        hostname = google_compute_instance.nat_vm.network_interface.0.network_ip
        username = google_sql_user.user.name
        password = google_sql_user.user.password
        database = google_sql_database.db.name
        port = 5432
    }

    private_connectivity {
        private_connection = google_datastream_private_connection.private_connection.id
    }
}
```
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=datastream_connection_profile_full&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Datastream Connection Profile Full


```hcl
resource "google_datastream_connection_profile" "default" {
	display_name          = "Connection profile"
	location              = "us-central1"
	connection_profile_id = "my-profile"

	gcs_profile {
		bucket    = "my-bucket"
		root_path = "/path"
	}

	forward_ssh_connectivity {
		hostname = "google.com"
		username = "my-user"
		port     = 8022
		password = "swordfish"
	}
	labels = {
		key = "value"
	}
}
```
## Example Usage - Datastream Connection Profile Postgres


```hcl
resource "google_sql_database_instance" "instance" {
    name             = "my-instance"
    database_version = "POSTGRES_14"
    region           = "us-central1"
    settings {
        tier = "db-f1-micro"

        ip_configuration {

            // Datastream IPs will vary by region.
            authorized_networks {
                value = "34.71.242.81"
            }

            authorized_networks {
                value = "34.72.28.29"
            }

            authorized_networks {
                value = "34.67.6.157"
            }

            authorized_networks {
                value = "34.67.234.134"
            }

            authorized_networks {
                value = "34.72.239.218"
            }
        }
    }

    deletion_protection  = true
}

resource "google_sql_database" "db" {
    instance = google_sql_database_instance.instance.name
    name     = "db"
}

resource "random_password" "pwd" {
    length = 16
    special = false
}

resource "google_sql_user" "user" {
    name = "user"
    instance = google_sql_database_instance.instance.name
    password = random_password.pwd.result
}

resource "google_datastream_connection_profile" "default" {
    display_name          = "Connection profile"
    location              = "us-central1"
    connection_profile_id = "my-profile"

    postgresql_profile {
        hostname = google_sql_database_instance.instance.public_ip_address
        username = google_sql_user.user.name
        password = google_sql_user.user.password
        database = google_sql_database.db.name
    }
}
```
## Example Usage - Datastream Connection Profile Sql Server


```hcl
resource "google_sql_database_instance" "instance" {
    name                = "sql-server"
    database_version    = "SQLSERVER_2019_STANDARD"
    region              = "us-central1"
    root_password       = "root-password"
    deletion_protection = true

    settings {
        tier = "db-custom-2-4096"
        ip_configuration {
            // Datastream IPs will vary by region.
            // https://cloud.google.com/datastream/docs/ip-allowlists-and-regions
            authorized_networks {
                value = "34.71.242.81"
            }

            authorized_networks {
                value = "34.72.28.29"
            }

            authorized_networks {
                value = "34.67.6.157"
            }

            authorized_networks {
                value = "34.67.234.134"
            }

            authorized_networks {
                value = "34.72.239.218"
            }
        }
    }
}

resource "google_sql_database" "db" {
    name       = "db"
    instance   = google_sql_database_instance.instance.name
}

resource "google_sql_user" "user" {
    name     = "user"
    instance = google_sql_database_instance.instance.name
    password = "password"
}

resource "google_datastream_connection_profile" "default" {
    display_name          = "SQL Server Source"
    location              = "us-central1"
    connection_profile_id = "source-profile"

    sql_server_profile {
        hostname = google_sql_database_instance.instance.public_ip_address
        port     = 1433
        username = google_sql_user.user.name
        password = google_sql_user.user.password
        database = google_sql_database.db.name
    }
}
```
## Example Usage - Datastream Connection Profile Salesforce


```hcl
resource "google_datastream_connection_profile" "default" {
    display_name          = "Salesforce Source"
    location              = "us-central1"
    connection_profile_id = "source-profile"
    create_without_validation = true
    provider = google-beta

    salesforce_profile {
        domain = "fake-domain.my.salesforce.com"
        user_credentials {
          username = "fake-username"
          secret_manager_stored_password = "fake-password"
          secret_manager_stored_security_token = "fake-token"
        }
    }
}
```
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=datastream_connection_profile_postgres_secret_manager&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Datastream Connection Profile Postgres Secret Manager


```hcl
resource "google_datastream_connection_profile" "default" {
    display_name              = "Postgres Source With Secret Manager"
    location                  = "us-central1"
    connection_profile_id     = "source-profile"
    create_without_validation = true


    postgresql_profile {
        hostname = "fake-hostname"
        port = 3306
        username = "fake-username"
        secret_manager_stored_password = "projects/fake-project/secrets/fake-secret/versions/1"
        database = "fake-database"
    }
}
```

## Argument Reference

The following arguments are supported:


* `display_name` -
  (Required)
  Display name.

* `connection_profile_id` -
  (Required)
  The connection profile identifier.

* `location` -
  (Required)
  The name of the location this connection profile is located in.


* `labels` -
  (Optional)
  Labels.
  **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
  Please refer to the field `effective_labels` for all of the labels present on the resource.

* `oracle_profile` -
  (Optional)
  Oracle database profile.
  Structure is [documented below](#nested_oracle_profile).

* `gcs_profile` -
  (Optional)
  Cloud Storage bucket profile.
  Structure is [documented below](#nested_gcs_profile).

* `mysql_profile` -
  (Optional)
  MySQL database profile.
  Structure is [documented below](#nested_mysql_profile).

* `bigquery_profile` -
  (Optional)
  BigQuery warehouse profile.

* `postgresql_profile` -
  (Optional)
  PostgreSQL database profile.
  Structure is [documented below](#nested_postgresql_profile).

* `salesforce_profile` -
  (Optional, [Beta](https://terraform.io/docs/providers/google/guides/provider_versions.html))
  Salesforce profile.
  Structure is [documented below](#nested_salesforce_profile).

* `sql_server_profile` -
  (Optional)
  SQL Server database profile.
  Structure is [documented below](#nested_sql_server_profile).

* `forward_ssh_connectivity` -
  (Optional)
  Forward SSH tunnel connectivity.
  Structure is [documented below](#nested_forward_ssh_connectivity).

* `private_connectivity` -
  (Optional)
  Private connectivity.
  Structure is [documented below](#nested_private_connectivity).

* `create_without_validation` -
  (Optional)
  Create the connection profile without validating it.

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.



<a name="nested_oracle_profile"></a>The `oracle_profile` block supports:

* `hostname` -
  (Required)
  Hostname for the Oracle connection.

* `port` -
  (Optional)
  Port for the Oracle connection.

* `username` -
  (Required)
  Username for the Oracle connection.

* `password` -
  (Optional)
  Password for the Oracle connection.
  **Note**: This property is sensitive and will not be displayed in the plan.

* `secret_manager_stored_password` -
  (Optional)
  A reference to a Secret Manager resource name storing the user's password.

* `database_service` -
  (Required)
  Database for the Oracle connection.

* `connection_attributes` -
  (Optional)
  Connection string attributes

<a name="nested_gcs_profile"></a>The `gcs_profile` block supports:

* `bucket` -
  (Required)
  The Cloud Storage bucket name.

* `root_path` -
  (Optional)
  The root path inside the Cloud Storage bucket.

<a name="nested_mysql_profile"></a>The `mysql_profile` block supports:

* `hostname` -
  (Required)
  Hostname for the MySQL connection.

* `port` -
  (Optional)
  Port for the MySQL connection.

* `username` -
  (Required)
  Username for the MySQL connection.

* `password` -
  (Optional)
  Password for the MySQL connection.
  **Note**: This property is sensitive and will not be displayed in the plan.

* `secret_manager_stored_password` -
  (Optional)
  A reference to a Secret Manager resource name storing the user's password.

* `ssl_config` -
  (Optional)
  SSL configuration for the MySQL connection.
  Structure is [documented below](#nested_mysql_profile_ssl_config).


<a name="nested_mysql_profile_ssl_config"></a>The `ssl_config` block supports:

* `client_key` -
  (Optional)
  PEM-encoded private key associated with the Client Certificate.
  If this field is used then the 'client_certificate' and the
  'ca_certificate' fields are mandatory.
  **Note**: This property is sensitive and will not be displayed in the plan.

* `client_key_set` -
  (Output)
  Indicates whether the clientKey field is set.

* `client_certificate` -
  (Optional)
  PEM-encoded certificate that will be used by the replica to
  authenticate against the source database server. If this field
  is used then the 'clientKey' and the 'caCertificate' fields are
  mandatory.
  **Note**: This property is sensitive and will not be displayed in the plan.

* `client_certificate_set` -
  (Output)
  Indicates whether the clientCertificate field is set.

* `ca_certificate` -
  (Optional)
  PEM-encoded certificate of the CA that signed the source database
  server's certificate.
  **Note**: This property is sensitive and will not be displayed in the plan.

* `ca_certificate_set` -
  (Output)
  Indicates whether the clientKey field is set.

<a name="nested_postgresql_profile"></a>The `postgresql_profile` block supports:

* `hostname` -
  (Required)
  Hostname for the PostgreSQL connection.

* `port` -
  (Optional)
  Port for the PostgreSQL connection.

* `username` -
  (Required)
  Username for the PostgreSQL connection.

* `password` -
  (Optional)
  Password for the PostgreSQL connection.
  **Note**: This property is sensitive and will not be displayed in the plan.

* `secret_manager_stored_password` -
  (Optional)
  A reference to a Secret Manager resource name storing the user's password.

* `database` -
  (Required)
  Database for the PostgreSQL connection.

<a name="nested_salesforce_profile"></a>The `salesforce_profile` block supports:

* `domain` -
  (Required)
  Domain for the Salesforce Org.

* `user_credentials` -
  (Optional)
  User credentials to use for Salesforce authentication.
  Structure is [documented below](#nested_salesforce_profile_user_credentials).

* `oauth2_client_credentials` -
  (Optional)
  OAuth credentials to use for Salesforce authentication.
  Structure is [documented below](#nested_salesforce_profile_oauth2_client_credentials).


<a name="nested_salesforce_profile_user_credentials"></a>The `user_credentials` block supports:

* `username` -
  (Optional)
  Username to use for authentication.

* `password` -
  (Optional)
  Password of the user.

* `security_token` -
  (Optional)
  Security token of the user.

* `secret_manager_stored_password` -
  (Optional)
  A reference to a Secret Manager resource name storing the user's password.

* `secret_manager_stored_security_token` -
  (Optional)
  A reference to a Secret Manager resource name storing the user's security token.

<a name="nested_salesforce_profile_oauth2_client_credentials"></a>The `oauth2_client_credentials` block supports:

* `client_id` -
  (Optional)
  Client ID to use for authentication.

* `client_secret` -
  (Optional)
  Client secret to use for authentication.

* `secret_manager_stored_client_secret` -
  (Optional)
  A reference to a Secret Manager resource name storing the client secret.

<a name="nested_sql_server_profile"></a>The `sql_server_profile` block supports:

* `hostname` -
  (Required)
  Hostname for the SQL Server connection.

* `port` -
  (Optional)
  Port for the SQL Server connection.

* `username` -
  (Required)
  Username for the SQL Server connection.

* `password` -
  (Optional)
  Password for the SQL Server connection.
  **Note**: This property is sensitive and will not be displayed in the plan.

* `secret_manager_stored_password` -
  (Optional)
  A reference to a Secret Manager resource name storing the user's password.

* `database` -
  (Required)
  Database for the SQL Server connection.

<a name="nested_forward_ssh_connectivity"></a>The `forward_ssh_connectivity` block supports:

* `hostname` -
  (Required)
  Hostname for the SSH tunnel.

* `username` -
  (Required)
  Username for the SSH tunnel.

* `port` -
  (Optional)
  Port for the SSH tunnel.

* `password` -
  (Optional)
  SSH password.
  **Note**: This property is sensitive and will not be displayed in the plan.

* `private_key` -
  (Optional)
  SSH private key.
  **Note**: This property is sensitive and will not be displayed in the plan.

<a name="nested_private_connectivity"></a>The `private_connectivity` block supports:

* `private_connection` -
  (Required)
  A reference to a private connection resource. Format: `projects/{project}/locations/{location}/privateConnections/{name}`

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/{{location}}/connectionProfiles/{{connection_profile_id}}`

* `name` -
  The resource's name.

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


ConnectionProfile can be imported using any of these accepted formats:

* `projects/{{project}}/locations/{{location}}/connectionProfiles/{{connection_profile_id}}`
* `{{project}}/{{location}}/{{connection_profile_id}}`
* `{{location}}/{{connection_profile_id}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import ConnectionProfile using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/locations/{{location}}/connectionProfiles/{{connection_profile_id}}"
  to = google_datastream_connection_profile.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), ConnectionProfile can be imported using one of the formats above. For example:

```
$ terraform import google_datastream_connection_profile.default projects/{{project}}/locations/{{location}}/connectionProfiles/{{connection_profile_id}}
$ terraform import google_datastream_connection_profile.default {{project}}/{{location}}/{{connection_profile_id}}
$ terraform import google_datastream_connection_profile.default {{location}}/{{connection_profile_id}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
