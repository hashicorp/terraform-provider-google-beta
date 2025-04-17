// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package firebase_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccDataSourceGoogleFirebaseWebAppConfig(t *testing.T) {
	t.Parallel()
	context := map[string]interface{}{
		"project_id":   envvar.GetTestProjectFromEnv(),
		"display_name": "tf_test Display Name WebApp DataSource",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:     func() { acctest.AccTestPreCheck(t) },
		CheckDestroy: testAccCheckFirebaseWebAppDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
				Config:                   testAccDataSourceGoogleFirebaseWebAppConfig(context),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.google_firebase_web_app_config.my_app_config", "project", context["project_id"].(string)),
					resource.TestCheckResourceAttrSet("data.google_firebase_web_app_config.my_app_config", "api_key"),
					resource.TestCheckResourceAttrSet("data.google_firebase_web_app_config.my_app_config", "auth_domain"),
					resource.TestCheckResourceAttrSet("data.google_firebase_web_app_config.my_app_config", "storage_bucket"),
				),
			},
		},
	})
}

func testAccDataSourceGoogleFirebaseWebAppConfig(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_firebase_web_app" "my_app" {
  project = "%{project_id}"
  display_name = "%{display_name}"
}

data "google_firebase_web_app_config" "my_app_config" {
  web_app_id = google_firebase_web_app.my_app.app_id
}
`, context)
}
