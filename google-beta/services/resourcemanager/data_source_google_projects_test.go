// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
// ----------------------------------------------------------------------------
//
//	***     AUTO GENERATED CODE    ***    Type: Handwritten     ***
//
// ----------------------------------------------------------------------------
//
//	This code is generated by Magic Modules using the following:
//
//	Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/services/resourcemanager/data_source_google_projects_test.go
//
//	DO NOT EDIT this file directly. Any changes made to this file will be
//	overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
package resourcemanager_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccDataSourceGoogleProjects_basic(t *testing.T) {
	t.Parallel()

	project := envvar.GetTestProjectFromEnv()

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGoogleProjectsConfig(project),
				Check: resource.ComposeTestCheckFunc(
					// We can't guarantee no project won't have our project ID as a prefix, so we'll check set-ness rather than correctness
					resource.TestCheckResourceAttrSet("data.google_projects.my-project", "projects.0.project_id"),
					resource.TestCheckResourceAttrSet("data.google_projects.my-project", "projects.0.name"),
					resource.TestCheckResourceAttrSet("data.google_projects.my-project", "projects.0.number"),
					resource.TestCheckResourceAttrSet("data.google_projects.my-project", "projects.0.lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.google_projects.my-project", "projects.0.parent.id"),
					resource.TestCheckResourceAttrSet("data.google_projects.my-project", "projects.0.parent.type"),
					resource.TestCheckResourceAttrSet("data.google_projects.my-project", "projects.0.create_time"),
				),
			},
		},
	})
}

func testAccCheckGoogleProjectsConfig(project string) string {
	return fmt.Sprintf(`
data "google_projects" "my-project" {
  filter = "projectId:%s"
}
`, project)
}
