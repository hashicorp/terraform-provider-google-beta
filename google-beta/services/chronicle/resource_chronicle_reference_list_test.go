// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package chronicle_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccChronicleReferenceList_chronicleReferencelistBasicExample_update(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"chronicle_id":  envvar.GetTestChronicleInstanceIdFromEnv(t),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccChronicleReferenceList_chronicleReferencelistBasicExample_basic(context),
			},
			{
				ResourceName:            "google_chronicle_reference_list.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"instance", "location", "reference_list_id"},
			},
			{
				Config: testAccChronicleReferenceList_chronicleReferencelistBasicExample_update(context),
			},
			{
				ResourceName:            "google_chronicle_reference_list.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"instance", "location", "reference_list_id"},
			},
		},
	})
}

func testAccChronicleReferenceList_chronicleReferencelistBasicExample_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_chronicle_reference_list" "example" {
 location = "us"
 instance = "%{chronicle_id}"
 reference_list_id = "tf_test_reference_list_id%{random_suffix}"
 description = "referencelist-description"
 entries {
  value = "referencelist-entry-value"
 }
 syntax_type = "REFERENCE_LIST_SYNTAX_TYPE_PLAIN_TEXT_STRING"
}
`, context)
}

func testAccChronicleReferenceList_chronicleReferencelistBasicExample_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_chronicle_reference_list" "example" {
 location = "us"
 instance = "%{chronicle_id}"
 reference_list_id = "tf_test_reference_list_id%{random_suffix}"
 description = "referencelist-description-updated"
 entries {
  value = "referencelist-entry-value-updated"
 }
 syntax_type = "REFERENCE_LIST_SYNTAX_TYPE_REGEX"
}
`, context)
}
