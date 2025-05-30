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
//	Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/acctest/test_utils.go.tmpl
//
//	DO NOT EDIT this file directly. Any changes made to this file will be
//	overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
package acctest

import (
	"archive/zip"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-mux/tf5muxserver"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
	googleoauth "golang.org/x/oauth2/google"
)

func CheckDataSourceStateMatchesResourceState(dataSourceName, resourceName string) func(*terraform.State) error {
	return CheckDataSourceStateMatchesResourceStateWithIgnores(dataSourceName, resourceName, map[string]struct{}{})
}

func CheckDataSourceStateMatchesResourceStateWithIgnores(dataSourceName, resourceName string, ignoreFields map[string]struct{}) func(*terraform.State) error {
	return func(s *terraform.State) error {
		ds, ok := s.RootModule().Resources[dataSourceName]
		if !ok {
			return fmt.Errorf("can't find %s in state", dataSourceName)
		}

		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("can't find %s in state", resourceName)
		}

		dsAttr := ds.Primary.Attributes
		rsAttr := rs.Primary.Attributes

		errMsg := ""
		// Data sources are often derived from resources, so iterate over the resource fields to
		// make sure all fields are accounted for in the data source.
		// If a field exists in the data source but not in the resource, its expected value should
		// be checked separately.
		for k := range rsAttr {
			if _, ok := ignoreFields[k]; ok {
				continue
			}
			if strings.HasPrefix(k, "labels.") || strings.HasPrefix(k, "terraform_labels.") || strings.HasPrefix(k, "effective_labels.") {
				continue
			}
			if k == "%" {
				continue
			}
			if dsAttr[k] != rsAttr[k] {
				// ignore data sources where an empty list is being compared against a null list.
				if k[len(k)-1:] == "#" && (dsAttr[k] == "" || dsAttr[k] == "0") && (rsAttr[k] == "" || rsAttr[k] == "0") {
					continue
				}
				errMsg += fmt.Sprintf("%s is %s; want %s\n", k, dsAttr[k], rsAttr[k])
			}
		}

		if errMsg != "" {
			return errors.New(errMsg)
		}

		return nil
	}
}

// General test utils

// MuxedProviders returns the correct test provider (between the sdk version or the framework version)
func MuxedProviders(testName string) (func() tfprotov5.ProviderServer, error) {
	ctx := context.Background()

	// primary is the SDKv2 implementation of the provider
	// If tests are run in VCR mode, the provider will use a cached config specific to the test name
	primary := GetSDKProvider(testName)

	providers := []func() tfprotov5.ProviderServer{
		primary.GRPCProvider, // sdk provider
		providerserver.NewProtocol5(NewFrameworkTestProvider(testName, primary)), // framework provider
	}

	muxServer, err := tf5muxserver.NewMuxServer(ctx, providers...)

	if err != nil {
		return nil, err
	}

	return muxServer.ProviderServer, nil
}

func RandString(t *testing.T, length int) string {
	if !IsVcrEnabled() {
		return acctest.RandString(length)
	}
	envPath := os.Getenv("VCR_PATH")
	vcrMode := os.Getenv("VCR_MODE")
	s, err := vcrSource(t, envPath, vcrMode)
	if err != nil {
		// At this point we haven't created any resources, so fail fast
		t.Fatal(err)
	}

	r := rand.New(s.source)
	result := make([]byte, length)
	set := "abcdefghijklmnopqrstuvwxyz012346789"
	for i := 0; i < length; i++ {
		result[i] = set[r.Intn(len(set))]
	}
	return string(result)
}

func RandInt(t *testing.T) int {
	if !IsVcrEnabled() {
		return acctest.RandInt()
	}
	envPath := os.Getenv("VCR_PATH")
	vcrMode := os.Getenv("VCR_MODE")
	s, err := vcrSource(t, envPath, vcrMode)
	if err != nil {
		// At this point we haven't created any resources, so fail fast
		t.Fatal(err)
	}

	return rand.New(s.source).Int()
}

func RandIntRange(t *testing.T, minInt int, maxInt int) int {
	if !IsVcrEnabled() {
		return acctest.RandIntRange(minInt, maxInt)
	}
	envPath := os.Getenv("VCR_PATH")
	vcrMode := os.Getenv("VCR_MODE")
	s, err := vcrSource(t, envPath, vcrMode)
	if err != nil {
		// At this point we haven't created any resources, so fail fast
		t.Fatal(err)
	}

	return rand.New(s.source).Intn(maxInt-minInt) + minInt
}

// ProtoV5ProviderFactories returns a muxed ProviderServer that uses the provider code from this repo (SDK and plugin-framework).
// Used to set ProtoV5ProviderFactories in a resource.TestStep within an acceptance test.
func ProtoV5ProviderFactories(t *testing.T) map[string]func() (tfprotov5.ProviderServer, error) {
	return map[string]func() (tfprotov5.ProviderServer, error){
		"google": func() (tfprotov5.ProviderServer, error) {
			provider, err := MuxedProviders(t.Name())
			return provider(), err
		},
	}
}

// ProtoV5ProviderBetaFactories returns the same as ProtoV5ProviderFactories only the provider is mapped with
// "google-beta" to ensure that registry examples use `google-beta` if the example is versioned as beta;
// normal beta tests should continue to use ProtoV5ProviderFactories
func ProtoV5ProviderBetaFactories(t *testing.T) map[string]func() (tfprotov5.ProviderServer, error) {
	return map[string]func() (tfprotov5.ProviderServer, error){
		"google-beta": func() (tfprotov5.ProviderServer, error) {
			provider, err := MuxedProviders(t.Name())
			return provider(), err
		},
	}
}

// This is a Printf sibling (Nprintf; Named Printf), which handles strings like
// Nprintf("Hello %{target}!", map[string]interface{}{"target":"world"}) == "Hello world!".
// This is particularly useful for generated tests, where we don't want to use Printf,
// since that would require us to generate a very particular ordering of arguments.
func Nprintf(format string, params map[string]interface{}) string {
	for key, val := range params {
		format = strings.Replace(format, "%{"+key+"}", fmt.Sprintf("%v", val), -1)
	}
	return format
}

func TestBucketName(t *testing.T) string {
	return fmt.Sprintf("%s-%d", "tf-test-bucket", RandInt(t))
}

func CreateZIPArchiveForCloudFunctionSource(t *testing.T, sourcePath string) string {
	source, err := ioutil.ReadFile(sourcePath)
	if err != nil {
		t.Fatal(err.Error())
	}
	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)

	// Create a new zip archive.
	w := zip.NewWriter(buf)

	f, err := w.Create("index.js")
	if err != nil {
		t.Fatal(err.Error())
	}
	_, err = f.Write(source)
	if err != nil {
		t.Fatal(err.Error())
	}

	// Make sure to check the error on Close.
	err = w.Close()
	if err != nil {
		t.Fatal(err.Error())
	}
	// Create temp file to write zip to
	tmpfile, err := ioutil.TempFile("", "sourceArchivePrefix")
	if err != nil {
		t.Fatal(err.Error())
	}

	if _, err := tmpfile.Write(buf.Bytes()); err != nil {
		t.Fatal(err.Error())
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err.Error())
	}
	return tmpfile.Name()
}

func GetAccessTokenFromTestCredsFromEnv(t *testing.T) string {
	credentials := envvar.GetTestCredsFromEnv()

	if credentials == "" {
		t.Fatalf("no creds provided in test environment: one of %s must be set for acceptance tests", strings.Join(envvar.CredsEnvVars, ", "))
	}

	// Environment might return a path or a JSON
	contents, _, err := verify.PathOrContents(credentials)
	if err != nil {
		t.Fatalf("error determining if creds in test environment are a path or contents: %s", err)
	}
	// Get googleoauth.Credentials
	c, err := googleoauth.CredentialsFromJSON(context.Background(), []byte(contents), transport_tpg.DefaultClientScopes...)
	if err != nil {
		t.Fatalf("invalid test credentials: %s", err)
	}
	// Get value for access_token
	token, err := c.TokenSource.Token()
	if err != nil {
		t.Fatalf("Unable to generate test access token: %s", err)
	}
	return token.AccessToken
}

// providerConfigEnvNames returns a list of all the environment variables that could be set by a user to configure the provider
func providerConfigEnvNames() []string {

	envs := []string{}

	// Use existing collections of ENV names
	envVarsSets := [][]string{
		envvar.CredsEnvVars,   // credentials field
		envvar.ProjectEnvVars, // project field
		envvar.RegionEnvVars,  // region field
		envvar.ZoneEnvVars,    // zone field
	}
	for _, set := range envVarsSets {
		envs = append(envs, set...)
	}

	// Add remaining ENVs
	envs = append(envs, "GOOGLE_OAUTH_ACCESS_TOKEN")          // access_token field
	envs = append(envs, "GOOGLE_BILLING_PROJECT")             // billing_project field
	envs = append(envs, "GOOGLE_IMPERSONATE_SERVICE_ACCOUNT") // impersonate_service_account field
	envs = append(envs, "USER_PROJECT_OVERRIDE")              // user_project_override field
	envs = append(envs, "CLOUDSDK_CORE_REQUEST_REASON")       // request_reason field

	envs = append(envs, "GOOGLE_APPLICATION_CREDENTIALS") // ADC used to configure clients when provider lacks credentials and access_token

	return envs
}

// UnsetProviderConfigEnvs unsets any ENVs in the test environment that
// configure the provider.
// The testing package will restore the original values after the test
func UnsetTestProviderConfigEnvs(t *testing.T) {
	envs := providerConfigEnvNames()
	if len(envs) > 0 {
		for _, k := range envs {
			t.Setenv(k, "")
		}
	}
}

func SetupTestEnvs(t *testing.T, envValues map[string]string) {
	// Set ENVs
	if len(envValues) > 0 {
		for k, v := range envValues {
			t.Setenv(k, v)
		}
	}
}

// Returns a fake credentials JSON string with the client_email set to a test-specific value
func GenerateFakeCredentialsJson(testId string) string {
	json := fmt.Sprintf(`{"private_key_id": "foo","private_key": "bar","client_email": "%s@example.com","client_id": "id@foo.com","type": "service_account"}`, testId)
	return json
}
