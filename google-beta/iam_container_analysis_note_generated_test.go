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

package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccContainerAnalysisNoteIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/containeranalysis.notes.occurrences.viewer",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccContainerAnalysisNoteIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_container_analysis_note_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/notes/%s roles/containeranalysis.notes.occurrences.viewer", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-attestor-note%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccContainerAnalysisNoteIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_container_analysis_note_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/notes/%s roles/containeranalysis.notes.occurrences.viewer", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-attestor-note%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccContainerAnalysisNoteIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/containeranalysis.notes.occurrences.viewer",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccContainerAnalysisNoteIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_container_analysis_note_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/notes/%s roles/containeranalysis.notes.occurrences.viewer user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-attestor-note%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccContainerAnalysisNoteIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/containeranalysis.notes.occurrences.viewer",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccContainerAnalysisNoteIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_container_analysis_note_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_container_analysis_note_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/notes/%s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-attestor-note%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccContainerAnalysisNoteIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_container_analysis_note_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/notes/%s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-attestor-note%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccContainerAnalysisNoteIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_container_analysis_note" "note" {
  name = "tf-test-attestor-note%{random_suffix}"
  attestation_authority {
    hint {
      human_readable_name = "Attestor Note"
    }
  }
}

resource "google_container_analysis_note_iam_member" "foo" {
  project = google_container_analysis_note.note.project
  note = google_container_analysis_note.note.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccContainerAnalysisNoteIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_container_analysis_note" "note" {
  name = "tf-test-attestor-note%{random_suffix}"
  attestation_authority {
    hint {
      human_readable_name = "Attestor Note"
    }
  }
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_container_analysis_note_iam_policy" "foo" {
  project = google_container_analysis_note.note.project
  note = google_container_analysis_note.note.name
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_container_analysis_note_iam_policy" "foo" {
  project = google_container_analysis_note.note.project
  note = google_container_analysis_note.note.name
  depends_on = [
    google_container_analysis_note_iam_policy.foo
  ]
}
`, context)
}

func testAccContainerAnalysisNoteIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_container_analysis_note" "note" {
  name = "tf-test-attestor-note%{random_suffix}"
  attestation_authority {
    hint {
      human_readable_name = "Attestor Note"
    }
  }
}

data "google_iam_policy" "foo" {
}

resource "google_container_analysis_note_iam_policy" "foo" {
  project = google_container_analysis_note.note.project
  note = google_container_analysis_note.note.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccContainerAnalysisNoteIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_container_analysis_note" "note" {
  name = "tf-test-attestor-note%{random_suffix}"
  attestation_authority {
    hint {
      human_readable_name = "Attestor Note"
    }
  }
}

resource "google_container_analysis_note_iam_binding" "foo" {
  project = google_container_analysis_note.note.project
  note = google_container_analysis_note.note.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccContainerAnalysisNoteIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_container_analysis_note" "note" {
  name = "tf-test-attestor-note%{random_suffix}"
  attestation_authority {
    hint {
      human_readable_name = "Attestor Note"
    }
  }
}

resource "google_container_analysis_note_iam_binding" "foo" {
  project = google_container_analysis_note.note.project
  note = google_container_analysis_note.note.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
