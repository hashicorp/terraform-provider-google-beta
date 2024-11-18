// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package backupdr_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccBackupDRBackupVault_backupDrBackupVaultFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckBackupDRBackupVaultDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBackupDRBackupVault_backupDrBackupVaultFullExample(context),
			},
			{
				ResourceName:            "google_backup_dr_backup_vault.backup-vault-test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"allow_missing", "annotations", "backup_vault_id", "force_delete", "force_update", "ignore_backup_plan_references", "ignore_inactive_datasources", "labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccBackupDRBackupVault_backupDrBackupVaultFullExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_backup_dr_backup_vault" "backup-vault-test" {
  location = "us-central1"
  backup_vault_id    = "tf-test-backup-vault-test%{random_suffix}"
  description = "This is a second backup vault built by Terraform."
  backup_minimum_enforced_retention_duration = "100000s"
  annotations = {
    annotations1 = "bar1"
    annotations2 = "baz1"
  }
  labels = {
    foo = "bar1"
    bar = "baz1"
  }
  force_update = "true"
  access_restriction = "WITHIN_ORGANIZATION"
  ignore_inactive_datasources = "true"
  ignore_backup_plan_references = "true"
  allow_missing = "true"
}
`, context)
}

func testAccCheckBackupDRBackupVaultDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_backup_dr_backup_vault" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{BackupDRBasePath}}projects/{{project}}/locations/{{location}}/backupVaults/{{backup_vault_id}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("BackupDRBackupVault still exists at %s", url)
			}
		}

		return nil
	}
}
