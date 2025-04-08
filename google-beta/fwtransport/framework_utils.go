// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package fwtransport

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/fwmodels"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/fwresource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"google.golang.org/api/googleapi"
)

const uaEnvVar = "TF_APPEND_USER_AGENT"

func CompileUserAgentString(ctx context.Context, name, tfVersion, provVersion string) string {
	ua := fmt.Sprintf("Terraform/%s (+https://www.terraform.io) Terraform-Plugin-SDK/%s %s/%s", tfVersion, "terraform-plugin-framework", name, provVersion)

	if add := os.Getenv(uaEnvVar); add != "" {
		add = strings.TrimSpace(add)
		if len(add) > 0 {
			ua += " " + add
			tflog.Debug(ctx, fmt.Sprintf("Using modified User-Agent: %s", ua))
		}
	}

	return ua
}

func GenerateFrameworkUserAgentString(metaData *fwmodels.ProviderMetaModel, currUserAgent string) string {
	if metaData != nil && !metaData.ModuleName.IsNull() && metaData.ModuleName.ValueString() != "" {
		return strings.Join([]string{currUserAgent, metaData.ModuleName.ValueString()}, " ")
	}

	return currUserAgent
}

func HandleDatasourceNotFoundError(ctx context.Context, err error, state *tfsdk.State, resource string, diags *diag.Diagnostics) {
	if transport_tpg.IsGoogleApiErrorWithCode(err, 404) {
		tflog.Warn(ctx, fmt.Sprintf("Removing %s because it's gone", resource))
		// The resource doesn't exist anymore
		state.RemoveResource(ctx)
	}

	diags.AddError(fmt.Sprintf("Error when reading or editing %s", resource), err.Error())
}

var DefaultRequestTimeout = 5 * time.Minute

type SendRequestOptions struct {
	Config               *transport_tpg.Config
	Method               string
	Project              string
	RawURL               string
	UserAgent            string
	Body                 map[string]any
	Timeout              time.Duration
	Headers              http.Header
	ErrorRetryPredicates []transport_tpg.RetryErrorPredicateFunc
	ErrorAbortPredicates []transport_tpg.RetryErrorPredicateFunc
}

func SendRequest(opt SendRequestOptions, diags *diag.Diagnostics) map[string]interface{} {
	reqHeaders := opt.Headers
	if reqHeaders == nil {
		reqHeaders = make(http.Header)
	}
	reqHeaders.Set("User-Agent", opt.UserAgent)
	reqHeaders.Set("Content-Type", "application/json")

	if opt.Config.UserProjectOverride && opt.Project != "" {
		// When opt.Project is "NO_BILLING_PROJECT_OVERRIDE" in the function GetCurrentUserEmail,
		// set the header X-Goog-User-Project to be empty string.
		if opt.Project == "NO_BILLING_PROJECT_OVERRIDE" {
			reqHeaders.Set("X-Goog-User-Project", "")
		} else {
			// Pass the project into this fn instead of parsing it from the URL because
			// both project names and URLs can have colons in them.
			reqHeaders.Set("X-Goog-User-Project", opt.Project)
		}
	}

	if opt.Timeout == 0 {
		opt.Timeout = DefaultRequestTimeout
	}

	var res *http.Response
	err := transport_tpg.Retry(transport_tpg.RetryOptions{
		RetryFunc: func() error {
			var buf bytes.Buffer
			if opt.Body != nil {
				err := json.NewEncoder(&buf).Encode(opt.Body)
				if err != nil {
					return err
				}
			}

			u, err := transport_tpg.AddQueryParams(opt.RawURL, map[string]string{"alt": "json"})
			if err != nil {
				return err
			}
			req, err := http.NewRequest(opt.Method, u, &buf)
			if err != nil {
				return err
			}

			req.Header = reqHeaders
			res, err = opt.Config.Client.Do(req)
			if err != nil {
				return err
			}

			if err := googleapi.CheckResponse(res); err != nil {
				googleapi.CloseBody(res)
				return err
			}

			return nil
		},
		Timeout:              opt.Timeout,
		ErrorRetryPredicates: opt.ErrorRetryPredicates,
		ErrorAbortPredicates: opt.ErrorAbortPredicates,
	})
	if err != nil {
		diags.AddError("Error when sending HTTP request: ", err.Error())
		return nil
	}

	if res == nil {
		diags.AddError("Unable to parse server response. This is most likely a terraform problem, please file a bug at https://github.com/hashicorp/terraform-provider-google/issues.", "")
		return nil
	}

	// The defer call must be made outside of the retryFunc otherwise it's closed too soon.
	defer googleapi.CloseBody(res)

	// 204 responses will have no body, so we're going to error with "EOF" if we
	// try to parse it. Instead, we can just return nil.
	if res.StatusCode == 204 {
		return nil
	}
	result := make(map[string]interface{})
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		diags.AddError("Error when sending HTTP request: ", err.Error())
		return nil
	}

	return result
}

type DefaultVars struct {
	BillingProject types.String
	Project        types.String
	Region         types.String
	Zone           types.String
}

func ReplaceVars(ctx context.Context, req interface{}, diags *diag.Diagnostics, data DefaultVars, config *transport_tpg.Config, linkTmpl string) string {
	return ReplaceVarsRecursive(ctx, req, diags, data, config, linkTmpl, false, 0)
}

// relaceVarsForId shortens variables by running them through GetResourceNameFromSelfLink
// this allows us to use long forms of variables from configs without needing
// custom id formats. For instance:
// accessPolicies/{{access_policy}}/accessLevels/{{access_level}}
// with values:
// access_policy: accessPolicies/foo
// access_level: accessPolicies/foo/accessLevels/bar
// becomes accessPolicies/foo/accessLevels/bar
func ReplaceVarsForId(ctx context.Context, req interface{}, diags *diag.Diagnostics, data DefaultVars, config *transport_tpg.Config, linkTmpl string) string {
	return ReplaceVarsRecursive(ctx, req, diags, data, config, linkTmpl, true, 0)
}

// ReplaceVars must be done recursively because there are baseUrls that can contain references to regions
// (eg cloudrun service) there aren't any cases known for 2+ recursion but we will track a run away
// substitution as 10+ calls to allow for future use cases.
func ReplaceVarsRecursive(ctx context.Context, req interface{}, diags *diag.Diagnostics, data DefaultVars, config *transport_tpg.Config, linkTmpl string, shorten bool, depth int) string {
	if depth > 10 {
		diags.AddError("url building error", "Recursive substitution detected.")
	}

	// https://github.com/google/re2/wiki/Syntax
	re := regexp.MustCompile("{{([%[:word:]]+)}}")
	f := BuildReplacementFunc(ctx, re, req, diags, data, config, linkTmpl, shorten)
	if diags.HasError() {
		return ""
	}
	final := re.ReplaceAllStringFunc(linkTmpl, f)

	if re.Match([]byte(final)) {
		return ReplaceVarsRecursive(ctx, req, diags, data, config, final, shorten, depth+1)
	}

	return final
}

// This function replaces references to Terraform properties (in the form of {{var}}) with their value in Terraform
// It also replaces {{project}}, {{project_id_or_project}}, {{region}}, and {{zone}} with their appropriate values
// This function supports URL-encoding the result by prepending '%' to the field name e.g. {{%var}}
func BuildReplacementFunc(ctx context.Context, re *regexp.Regexp, req interface{}, diags *diag.Diagnostics, data DefaultVars, config *transport_tpg.Config, linkTmpl string, shorten bool) func(string) string {
	var project, region, zone string
	var projectID types.String

	if strings.Contains(linkTmpl, "{{project}}") {
		project = fwresource.GetProjectFramework(data.Project, types.StringValue(config.Project), diags).ValueString()
		if diags.HasError() {
			return nil
		}
		if shorten {
			project = strings.TrimPrefix(project, "projects/")
		}
	}

	if strings.Contains(linkTmpl, "{{project_id_or_project}}") {
		var diagInfo diag.Diagnostics
		switch req.(type) {
		case resource.CreateRequest:
			pReq := req.(resource.CreateRequest)
			diagInfo = pReq.Plan.GetAttribute(ctx, path.Root("project_id"), &projectID)
		case resource.UpdateRequest:
			pReq := req.(resource.UpdateRequest)
			diagInfo = pReq.Plan.GetAttribute(ctx, path.Root("project_id"), &projectID)
		case resource.ReadRequest:
			sReq := req.(resource.ReadRequest)
			diagInfo = sReq.State.GetAttribute(ctx, path.Root("project_id"), &projectID)
		case resource.DeleteRequest:
			sReq := req.(resource.DeleteRequest)
			diagInfo = sReq.State.GetAttribute(ctx, path.Root("project_id"), &projectID)
		}
		diags.Append(diagInfo...)
		if diags.HasError() {
			return nil
		}
		if projectID.ValueString() != "" {
			project = fwresource.GetProjectFramework(data.Project, types.StringValue(config.Project), diags).ValueString()
			if diags.HasError() {
				return nil
			}
		}
		if shorten {
			project = strings.TrimPrefix(project, "projects/")
			projectID = types.StringValue(strings.TrimPrefix(projectID.ValueString(), "projects/"))
		}
	}

	if strings.Contains(linkTmpl, "{{region}}") {
		region = fwresource.GetRegionFramework(data.Region, types.StringValue(config.Region), diags).ValueString()
		if diags.HasError() {
			return nil
		}
		if shorten {
			region = strings.TrimPrefix(region, "regions/")
		}
	}

	if strings.Contains(linkTmpl, "{{zone}}") {
		zone = fwresource.GetRegionFramework(data.Zone, types.StringValue(config.Zone), diags).ValueString()
		if diags.HasError() {
			return nil
		}
		if shorten {
			zone = strings.TrimPrefix(region, "zones/")
		}
	}

	f := func(s string) string {

		m := re.FindStringSubmatch(s)[1]
		if m == "project" {
			return project
		}
		if m == "project_id_or_project" {
			if projectID.ValueString() != "" {
				return projectID.ValueString()
			}
			return project
		}
		if m == "region" {
			return region
		}
		if m == "zone" {
			return zone
		}
		if string(m[0]) == "%" {
			var v types.String
			var diagInfo diag.Diagnostics
			switch req.(type) {
			case resource.CreateRequest:
				pReq := req.(resource.CreateRequest)
				diagInfo = pReq.Plan.GetAttribute(ctx, path.Root("m[1:]"), &v)
			case resource.UpdateRequest:
				pReq := req.(resource.UpdateRequest)
				diagInfo = pReq.Plan.GetAttribute(ctx, path.Root("m[1:]"), &v)
			case resource.ReadRequest:
				sReq := req.(resource.ReadRequest)
				diagInfo = sReq.State.GetAttribute(ctx, path.Root("m[1:]"), &v)
			case resource.DeleteRequest:
				sReq := req.(resource.DeleteRequest)
				diagInfo = sReq.State.GetAttribute(ctx, path.Root("m[1:]"), &v)
			}
			diags.Append(diagInfo...)
			if !diags.HasError() {
				if v.ValueString() != "" {
					if shorten {
						return tpgresource.GetResourceNameFromSelfLink(fmt.Sprintf("%v", v.ValueString()))
					} else {
						return fmt.Sprintf("%v", v.ValueString())
					}
				}
			}
		} else {
			var v types.String
			var diagInfo diag.Diagnostics
			switch req.(type) {
			case resource.CreateRequest:
				pReq := req.(resource.CreateRequest)
				diagInfo = pReq.Plan.GetAttribute(ctx, path.Root("m"), &v)
			case resource.UpdateRequest:
				pReq := req.(resource.UpdateRequest)
				diagInfo = pReq.Plan.GetAttribute(ctx, path.Root("m"), &v)
			case resource.ReadRequest:
				sReq := req.(resource.ReadRequest)
				diagInfo = sReq.State.GetAttribute(ctx, path.Root("m"), &v)
			case resource.DeleteRequest:
				sReq := req.(resource.DeleteRequest)
				diagInfo = sReq.State.GetAttribute(ctx, path.Root("m"), &v)
			}
			diags.Append(diagInfo...)
			if !diags.HasError() {
				if v.ValueString() != "" {
					if shorten {
						return tpgresource.GetResourceNameFromSelfLink(fmt.Sprintf("%v", v.ValueString()))
					} else {
						return fmt.Sprintf("%v", v.ValueString())
					}
				}
			}
		}

		// terraform-google-conversion doesn't provide a provider config in tests.
		if config != nil {
			// Attempt to draw values from the provider config if it's present.
			if f := reflect.Indirect(reflect.ValueOf(config)).FieldByName(m); f.IsValid() {
				return f.String()
			}
		}
		return ""
	}

	return f
}
