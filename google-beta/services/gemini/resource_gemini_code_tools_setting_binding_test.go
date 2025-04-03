// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package gemini_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccGeminiCodeToolsSettingBinding_update(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"code_tools_setting_id": fmt.Sprintf("tf-test-ls-%s", acctest.RandString(t, 10)),
		"setting_binding_id":    fmt.Sprintf("tf-test-lsb-%s", acctest.RandString(t, 10)),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGeminiCodeToolsSettingBinding_basic(context),
			},
			{
				ResourceName:            "google_gemini_code_tools_setting_binding.basic_binding",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "location", "code_tools_setting_id", "terraform_labels"},
			},
			{
				Config: testAccGeminiCodeToolsSettingBinding_update(context),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("google_gemini_code_tools_setting_binding.basic_binding", plancheck.ResourceActionUpdate),
					},
				},
			},
			{
				ResourceName:            "google_gemini_code_tools_setting_binding.basic_binding",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "location", "code_tools_setting_id", "terraform_labels"},
			},
		},
	})
}

func testAccGeminiCodeToolsSettingBinding_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_project" "project" {
}

resource "google_gemini_code_tools_setting" "basic" {
    code_tools_setting_id = "%{code_tools_setting_id}"
    location = "global"
    labels = {"my_key" = "my_value"}
    enabled_tool {
        handle = "my_handle"
        tool = "my_tool"
        account_connector = "my_con"
        config {
            key = "my_key"
            value = "my_value"
        }
        uri_override = "my_uri_override"
    }
}

resource "google_gemini_code_tools_setting_binding" "basic_binding" {
    code_tools_setting_id = google_gemini_code_tools_setting.basic.code_tools_setting_id
    setting_binding_id = "%{setting_binding_id}"
    location = "global"
    target = "projects/${data.google_project.project.number}"
}
`, context)
}

func testAccGeminiCodeToolsSettingBinding_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_project" "project" {
}

resource "google_gemini_code_tools_setting" "basic" {
    code_tools_setting_id = "%{code_tools_setting_id}"
    location = "global"
    labels = {"my_key" = "my_value"}
    enabled_tool {
        handle = "my_handle"
        tool = "my_tool"
        account_connector = "my_con"
        config {
            key = "my_key"
            value = "my_value"
        }
        uri_override = "my_uri_override"
    }
}

resource "google_gemini_code_tools_setting_binding" "basic_binding" {
    code_tools_setting_id = google_gemini_code_tools_setting.basic.code_tools_setting_id
    setting_binding_id = "%{setting_binding_id}"
    location = "global"
    target = "projects/${data.google_project.project.number}"
    labels = {"my_key" = "my_value"}
	product = "GEMINI_CODE_ASSIST"
}
`, context)
}
