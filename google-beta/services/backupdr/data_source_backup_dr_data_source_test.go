// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package backupdr_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

// this test cannot be ran locally without seeding your environment with a backup vault and scheduling the backup of a compute instance
// to generate a unique datasourceId, that this test would need to be modified to use. the values in this test correspond to those used
// in our testing processes.
func TestAccDataSourceGoogleCloudBackupDRDataSource_basic(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	stepChecks := func(wantName string, wantState string) []resource.TestCheckFunc {
		stepCheck := []resource.TestCheckFunc{
			resource.TestCheckResourceAttr("data.google_backup_dr_data_source.foo", "name", wantName),
			resource.TestCheckResourceAttr("data.google_backup_dr_data_source.foo", "state", wantState),
		}
		return stepCheck
	}
	project := envvar.GetTestProjectFromEnv()
	expectedName := fmt.Sprintf("projects/%s/locations/us-central1/backupVaults/bv-test/dataSources/ds-test", project)
	expectedState := "ACTIVE"

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGoogleCloudBackupDRDataSource_basic(context),
				Check:  resource.ComposeTestCheckFunc(stepChecks(expectedName, expectedState)...),
			},
		},
	})
}

func testAccDataSourceGoogleCloudBackupDRDataSource_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_project" "project" {
}

data "google_backup_dr_data_source" "foo" {
  project = data.google_project.project.project_id
  location      = "us-central1"
  backup_vault_id = "bv-test"
  data_source_id = "ds-test"
  
}

`, context)
}
