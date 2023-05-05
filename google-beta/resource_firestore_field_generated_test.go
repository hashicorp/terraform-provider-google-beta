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

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccFirestoreField_firestoreFieldBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    acctest.GetTestFirestoreProjectFromEnv(t),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckFirestoreFieldDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirestoreField_firestoreFieldBasicExample(context),
			},
			{
				ResourceName:            "google_firestore_field.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"database", "collection", "field"},
			},
		},
	})
}

func testAccFirestoreField_firestoreFieldBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_firestore_field" "basic" {
  project = "%{project_id}"
  database = "(default)"
  collection = "chatrooms_%{random_suffix}"
  field = "basic"

  index_config {
    indexes {
        order = "ASCENDING"
        query_scope = "COLLECTION_GROUP"
    }
    indexes {
        array_config = "CONTAINS"
    }
  }

  ttl_config {}
}
`, context)
}

func TestAccFirestoreField_firestoreFieldTimestampExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    acctest.GetTestFirestoreProjectFromEnv(t),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckFirestoreFieldDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirestoreField_firestoreFieldTimestampExample(context),
			},
			{
				ResourceName:            "google_firestore_field.timestamp",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"database", "collection", "field"},
			},
		},
	})
}

func testAccFirestoreField_firestoreFieldTimestampExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_firestore_field" "timestamp" {
  project = "%{project_id}"
  collection = "chatrooms_%{random_suffix}"
  field = "timestamp"

  // Disable all single field indexes for the timestamp property.
  index_config {}
  ttl_config {}
}
`, context)
}

func TestAccFirestoreField_firestoreFieldMatchOverrideExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    acctest.GetTestFirestoreProjectFromEnv(t),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckFirestoreFieldDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirestoreField_firestoreFieldMatchOverrideExample(context),
			},
			{
				ResourceName:            "google_firestore_field.match_override",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"database", "collection", "field"},
			},
		},
	})
}

func testAccFirestoreField_firestoreFieldMatchOverrideExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_firestore_field" "match_override" {
  project = "%{project_id}"
  collection = "chatrooms_%{random_suffix}"
  field = "field_with_same_configuration_as_ancestor"

  index_config {
    indexes {
        order = "ASCENDING"
    }
    indexes {
        order = "DESCENDING"
    }
    indexes {
        array_config = "CONTAINS"
    }
  }
}
`, context)
}

func testAccCheckFirestoreFieldDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_firestore_field" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			// Firestore fields are not deletable. We consider the field deleted if:
			// 1) the index configuration has no overrides and matches the ancestor configuration.
			// 2) the ttl configuration is unset.

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{FirestoreBasePath}}projects/{{project}}/databases/{{database}}/collectionGroups/{{collection}}/fields/{{field}}")
			if err != nil {
				return err
			}

			res, err := SendRequest(config, "GET", "", url, config.UserAgent, nil)
			if err != nil {
				return err
			}

			if v := res["indexConfig"]; v != nil {
				indexConfig := v.(map[string]interface{})

				usesAncestorConfig, ok := indexConfig["usesAncestorConfig"].(bool)

				if !ok || !usesAncestorConfig {
					return fmt.Errorf("Index configuration is not using the ancestor config %s.", url)
				}
			}

			if res["ttlConfig"] != nil {
				return fmt.Errorf("TTL configuration was not deleted at %s.", url)
			}
		}

		return nil
	}
}
