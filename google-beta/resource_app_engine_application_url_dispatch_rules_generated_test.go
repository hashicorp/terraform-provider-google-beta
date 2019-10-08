// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccAppEngineApplicationUrlDispatchRules_appEngineApplicationUrlDispatchRulesBasicExampleExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAppEngineApplicationUrlDispatchRulesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppEngineApplicationUrlDispatchRules_appEngineApplicationUrlDispatchRulesBasicExampleExample(context),
			},
			{
				ResourceName:      "google_app_engine_application_url_dispatch_rules.",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccAppEngineApplicationUrlDispatchRules_appEngineApplicationUrlDispatchRulesBasicExampleExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "bucket" {
	name = ""
}

resource "google_storage_bucket_object" "object" {
	name   = "hello-world.zip"
	bucket = "${google_storage_bucket.bucket.name}"
	source = "./test-fixtures/appengine/hello-world.zip"
}

resource "google_app_engine_standard_app_version" "myapp_v1" {
  version_id = "v1"
  service = "myapp"
  runtime = "nodejs10"
  noop_on_destroy = true
  entrypoint {
    shell = "node ./app.js"
  }
  deployment {
    zip {
      source_url = "https://storage.googleapis.com/${google_storage_bucket.bucket.name}/hello-world.zip"
    }  
  }
  env_variables = {
    port = "8080"
  } 
  depends_on = ["google_storage_bucket_object.object"]

}

resource "google_app_engine_application_url_dispatch_rules" "service_rules" {
  # project = "my-project"
  dispatch_rules {
    domain = "*"
    path = "/default/*"
    service = "default"
  }
  dispatch_rules {
    domain = "*"
    path = "/myapp/*"
    service = "myapp"
  }
}
`, context)
}

func testAccCheckAppEngineApplicationUrlDispatchRulesDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_app_engine_application_url_dispatch_rules" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(config, rs, "{{AppEngineBasePath}}/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", "", url, nil)
		if err == nil {
			return fmt.Errorf("AppEngineApplicationUrlDispatchRules still exists at %s", url)
		}
	}

	return nil
}
