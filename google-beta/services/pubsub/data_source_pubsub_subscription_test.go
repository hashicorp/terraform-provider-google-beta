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
//	Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/services/pubsub/data_source_pubsub_subscription_test.go
//
//	DO NOT EDIT this file directly. Any changes made to this file will be
//	overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
package pubsub_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccDataSourceGooglePubsubSubscription_basic(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckPubsubSubscriptionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGooglePubsubSubscription_basic(context),
				Check: resource.ComposeTestCheckFunc(
					acctest.CheckDataSourceStateMatchesResourceState("data.google_pubsub_subscription.foo", "google_pubsub_subscription.foo"),
				),
			},
		},
	})
}

func TestAccDataSourceGooglePubsubSubscription_optionalProject(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckPubsubSubscriptionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGooglePubsubSubscription_optionalProject(context),
				Check: resource.ComposeTestCheckFunc(
					acctest.CheckDataSourceStateMatchesResourceState("data.google_pubsub_subscription.foo", "google_pubsub_subscription.foo"),
				),
			},
		},
	})
}

func testAccDataSourceGooglePubsubSubscription_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_pubsub_topic" "foo" {
  name     = "tf-test-pubsub-%{random_suffix}"
}

resource "google_pubsub_subscription" "foo" {
  name     = "tf-test-pubsub-subscription-%{random_suffix}"
  topic    = google_pubsub_topic.foo.name
  labels   = {
    my-label = "my-label-value"
  }
}

data "google_pubsub_subscription" "foo" {
  name     = google_pubsub_subscription.foo.name
  project  = google_pubsub_subscription.foo.project
}
`, context)
}

func testAccDataSourceGooglePubsubSubscription_optionalProject(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_pubsub_topic" "foo" {
  name     = "tf-test-pubsub-%{random_suffix}"
}

resource "google_pubsub_subscription" "foo" {
  name     = "tf-test-pubsub-subscription-%{random_suffix}"
  topic    = google_pubsub_topic.foo.name
}

data "google_pubsub_subscription" "foo" {
  name     = google_pubsub_subscription.foo.name
}
`, context)
}
