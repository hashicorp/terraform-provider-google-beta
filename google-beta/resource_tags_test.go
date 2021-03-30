package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// Tags tests cannot be run in parallel without running into Error Code 10: ABORTED
// See https://github.com/hashicorp/terraform-provider-google/issues/8637

func TestAccTags(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"tagKeyBasic":    testAccTagsTagKey_tagKeyBasic,
		"tagKeyUpdate":   testAccTagsTagKey_tagKeyUpdate,
		"tagValueBasic":  testAccTagsTagValue_tagValueBasic,
		"tagValueUpdate": testAccTagsTagValue_tagValueUpdate,
	}

	for name, tc := range testCases {
		// shadow the tc variable into scope so that when
		// the loop continues, if t.Run hasn't executed tc(t)
		// yet, we don't have a race condition
		// see https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		tc := tc
		t.Run(name, func(t *testing.T) {
			tc(t)
		})
	}
}

func testAccTagsTagKey_tagKeyBasic(t *testing.T) {
	context := map[string]interface{}{
		"org_id":        getTestOrgFromEnv(t),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckTagsTagKeyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccTagsTagKey_tagKeyBasicExample(context),
			},
		},
	})
}

func testAccTagsTagKey_tagKeyBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_tags_tag_key" "key" {
  provider = google-beta

  parent = "organizations/%{org_id}"
  short_name = "foo%{random_suffix}"
  description = "For foo%{random_suffix} resources."
}
`, context)
}

func testAccTagsTagKey_tagKeyUpdate(t *testing.T) {
	context := map[string]interface{}{
		"org_id":        getTestOrgFromEnv(t),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckTagsTagKeyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccTagsTagKey_basic(context),
			},
			{
				ResourceName:      "google_tags_tag_key.key",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccTagsTagKey_basicUpdated(context),
			},
			{
				ResourceName:      "google_tags_tag_key.key",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccTagsTagKey_basic(context map[string]interface{}) string {
	return Nprintf(`
resource "google_tags_tag_key" "key" {
  provider = google-beta

  parent = "organizations/%{org_id}"
  short_name = "foo%{random_suffix}"
  description = "For foo%{random_suffix} resources."
}
`, context)
}

func testAccTagsTagKey_basicUpdated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_tags_tag_key" "key" {
  provider = google-beta

  parent = "organizations/%{org_id}"
  short_name = "foo%{random_suffix}"
  description = "Anything related to foo%{random_suffix}"
}
`, context)
}

func testAccCheckTagsTagKeyDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_tags_tag_key" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{TagsBasePath}}tagKeys/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("TagsTagKey still exists at %s", url)
			}
		}

		return nil
	}
}

func testAccTagsTagValue_tagValueBasic(t *testing.T) {
	context := map[string]interface{}{
		"org_id":        getTestOrgFromEnv(t),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckTagsTagValueDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccTagsTagValue_tagValueBasicExample(context),
			},
		},
	})
}

func testAccTagsTagValue_tagValueBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_tags_tag_key" "key" {
  provider = google-beta

  parent = "organizations/%{org_id}"
  short_name = "foobarbaz%{random_suffix}"
  description = "For foo/bar/baz resources."
}

resource "google_tags_tag_value" "value" {
  provider = google-beta

  parent = "tagKeys/${google_tags_tag_key.key.name}"
  short_name = "foo%{random_suffix}"
  description = "For foo resources."
}
`, context)
}

func testAccTagsTagValue_tagValueUpdate(t *testing.T) {
	context := map[string]interface{}{
		"org_id":        getTestOrgFromEnv(t),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckTagsTagValueDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccTagsTagValue_basic(context),
			},
			{
				ResourceName:      "google_tags_tag_key.key",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccTagsTagValue_basicUpdated(context),
			},
			{
				ResourceName:      "google_tags_tag_key.key",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccTagsTagValue_basic(context map[string]interface{}) string {
	return Nprintf(`
resource "google_tags_tag_key" "key" {
  provider = google-beta

  parent = "organizations/%{org_id}"
  short_name = "foobarbaz%{random_suffix}"
  description = "For foo/bar/baz resources."
}

resource "google_tags_tag_value" "value" {
  provider = google-beta

  parent = "tagKeys/${google_tags_tag_key.key.name}"
  short_name = "foo%{random_suffix}"
  description = "For foo resources."
}
`, context)
}

func testAccTagsTagValue_basicUpdated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_tags_tag_key" "key" {
  provider = google-beta

  parent = "organizations/%{org_id}"
  short_name = "foobarbaz%{random_suffix}"
  description = "For foo/bar/baz resources."
}

resource "google_tags_tag_value" "value" {
  provider = google-beta

  parent = "tagKeys/${google_tags_tag_key.key.name}"
  short_name = "foo%{random_suffix}"
  description = "For any foo resources."
}
`, context)
}

func testAccCheckTagsTagValueDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_tags_tag_key" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{TagsBasePath}}tagValues/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("TagsTagValue still exists at %s", url)
			}
		}

		return nil
	}
}
