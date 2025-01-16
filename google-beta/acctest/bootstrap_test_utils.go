// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package acctest

import (
	"context"
	"fmt"
	"log"
	"maps"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	// For beta tests only
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/kms"
	tpgservicusage "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/serviceusage"
	resourceManagerV3 "google.golang.org/api/cloudresourcemanager/v3"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	tpgcompute "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/compute"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/privateca"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/resourcemanager"
	tpgservicenetworking "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/servicenetworking"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/sql"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgiamresource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"

	"google.golang.org/api/cloudbilling/v1"
	cloudkms "google.golang.org/api/cloudkms/v1"
	cloudresourcemanager "google.golang.org/api/cloudresourcemanager/v1"
	iam "google.golang.org/api/iam/v1"
	"google.golang.org/api/iamcredentials/v1"
	"google.golang.org/api/servicenetworking/v1"
	"google.golang.org/api/serviceusage/v1"
	sqladmin "google.golang.org/api/sqladmin/v1beta4"
)

var SharedKeyRing = "tftest-shared-keyring-1"

var DefaultKeyHandleName = "eed58b7b-20ad-4da8-ad85-ba78a0d5ab87"
var DefaultKeyHandleResourceType = "compute.googleapis.com/Disk"
var CloudKmsSrviceName = "cloudkms.googleapis.com"

var SharedCryptoKey = map[string]string{
	"ENCRYPT_DECRYPT":    "tftest-shared-key-1",
	"ASYMMETRIC_SIGN":    "tftest-shared-sign-key-1",
	"ASYMMETRIC_DECRYPT": "tftest-shared-decrypt-key-1",
}

type BootstrappedKMS struct {
	*cloudkms.KeyRing
	*cloudkms.CryptoKey
	CryptoKeyVersions []*cloudkms.CryptoKeyVersion
}

func BootstrapKMSKey(t *testing.T) BootstrappedKMS {
	return BootstrapKMSKeyInLocation(t, "global")
}

func BootstrapKMSKeyInLocation(t *testing.T, locationID string) BootstrappedKMS {
	return BootstrapKMSKeyWithPurposeInLocation(t, "ENCRYPT_DECRYPT", locationID)
}

// BootstrapKMSKeyWithPurpose returns a KMS key in the "global" location.
// See BootstrapKMSKeyWithPurposeInLocation.
func BootstrapKMSKeyWithPurpose(t *testing.T, purpose string) BootstrappedKMS {
	return BootstrapKMSKeyWithPurposeInLocation(t, purpose, "global")
}

/**
* BootstrapKMSKeyWithPurposeInLocation will return a KMS key in a
* particular location with the given purpose that can be used
* in tests that are testing KMS integration with other resources.
*
* This will either return an existing key or create one if it hasn't been created
* in the project yet. The motivation is because keyrings don't get deleted and we
* don't want a linear growth of disabled keyrings in a project. We also don't want
* to incur the overhead of creating a new project for each test that needs to use
* a KMS key.
**/
func BootstrapKMSKeyWithPurposeInLocation(t *testing.T, purpose, locationID string) BootstrappedKMS {
	return BootstrapKMSKeyWithPurposeInLocationAndName(t, purpose, locationID, SharedCryptoKey[purpose])
}

type BootstrappedKMSAutokey struct {
	*cloudkms.AutokeyConfig
	*cloudkms.KeyHandle
}

func BootstrapKMSAutokeyKeyHandle(t *testing.T) BootstrappedKMSAutokey {
	return BootstrapKMSAutokeyKeyHandleWithLocation(t, "global")
}

func BootstrapKMSAutokeyKeyHandleWithLocation(t *testing.T, locationID string) BootstrappedKMSAutokey {
	config := BootstrapConfig(t)
	if config == nil {
		return BootstrappedKMSAutokey{
			&cloudkms.AutokeyConfig{},
			&cloudkms.KeyHandle{},
		}
	}

	autokeyFolder, kmsProject, resourceProject := setupAutokeyTestResources(t, config)

	// Enable autokey on autokey test folder
	kmsClient := config.NewKmsClient(config.UserAgent)
	autokeyConfigID := fmt.Sprintf("%s/autokeyConfig", autokeyFolder.Name)
	autokeyConfig, err := kmsClient.Folders.UpdateAutokeyConfig(autokeyConfigID, &cloudkms.AutokeyConfig{
		KeyProject: fmt.Sprintf("projects/%s", kmsProject.ProjectId),
	}).UpdateMask("keyProject").Do()
	if err != nil {
		t.Errorf("unable to bootstrap KMS keyHandle. Cannot enable autokey on autokey test folder: %s", err)
	}

	keyHandleParent := fmt.Sprintf("projects/%s/locations/%s", resourceProject.ProjectId, locationID)
	keyHandleName := fmt.Sprintf("%s/keyHandles/%s", keyHandleParent, DefaultKeyHandleName)

	// Get or Create the hard coded keyHandle for testing
	keyHandle, err := kmsClient.Projects.Locations.KeyHandles.Get(keyHandleName).Do()

	if err != nil {
		if transport_tpg.IsGoogleApiErrorWithCode(err, 404) {
			newKeyHandle := cloudkms.KeyHandle{
				ResourceTypeSelector: DefaultKeyHandleResourceType,
			}

			keyHandleOp, err := kmsClient.Projects.Locations.KeyHandles.Create(keyHandleParent, &newKeyHandle).KeyHandleId(DefaultKeyHandleName).Do()
			if err != nil {
				t.Errorf("unable to bootstrap KMS keyHandle. Cannot create new KeyHandle: %s", err)
			}

			opAsMap, err := tpgresource.ConvertToMap(keyHandleOp)
			if err != nil {
				t.Errorf("unable to bootstrap KMS keyHandle. Cannot get operation map: %s", err)
			}

			var response map[string]interface{}
			err = kms.KMSOperationWaitTimeWithResponse(config, opAsMap, &response, resourceProject.ProjectId, "creating keyHandle", config.UserAgent, time.Duration(5)*time.Minute)
			if err != nil {
				t.Errorf("unable to bootstrap KMS keyHandle. Cannot wait for create keyhandle operation: %s", err)
			}
			keyHandle = &cloudkms.KeyHandle{
				Name:                 response["name"].(string),
				KmsKey:               response["kmsKey"].(string),
				ResourceTypeSelector: response["resourceTypeSelector"].(string),
			}
		} else {
			t.Errorf("unable to bootstrap KMS keyHandle. Cannot call KeyHandle service: %s", err)
		}
	}

	if keyHandle == nil {
		t.Fatalf("unable to bootstrap KMS keyHandle. KeyHandle is nil!")
	}

	return BootstrappedKMSAutokey{
		autokeyConfig,
		keyHandle,
	}
}

func setupAutokeyTestResources(t *testing.T, config *transport_tpg.Config) (*resourceManagerV3.Folder, *cloudresourcemanager.Project, *cloudresourcemanager.Project) {
	projectIDSuffix := strings.Replace(envvar.GetTestProjectFromEnv(), "ci-test-project-", "", 1)
	defaultAutokeyTestFolderName := fmt.Sprintf("autokeytest-%s-fd", projectIDSuffix)
	defaultAutokeyTestKmsProject := fmt.Sprintf("test-kms-%s-prj", projectIDSuffix)
	defaultAutokeyTestResourceProject := fmt.Sprintf("test-res-%s-prj", projectIDSuffix)

	curUserEmail, err := transport_tpg.GetCurrentUserEmail(config, config.UserAgent)
	if err != nil {
		t.Errorf("unable to bootstrap KMS keyHandle. Cannot get current usr: %s", err)
	}
	// create a folder to configure autokey config and resource folder
	autokeyFolder := BootstrapFolder(t, defaultAutokeyTestFolderName)
	parent := &cloudresourcemanager.ResourceId{
		Type: "folder",
		Id:   strings.Split(autokeyFolder.Name, "/")[1],
	}
	// create and setup kms project for hosting keyring and keys for autokey
	kmsProject := BootstrapProjectWithParent(t, defaultAutokeyTestKmsProject, envvar.GetTestBillingAccountFromEnv(t), parent, []string{CloudKmsSrviceName})
	kmsProjectID := fmt.Sprintf("projects/%s", kmsProject.ProjectId)
	kmsSAEmail, err := GenerateCloudKmsServiceIdentity(config, fmt.Sprintf("%v", kmsProject.ProjectNumber))
	if err != nil {
		t.Errorf("unable to bootstrap KMS keyHandle. Cannot create cloudkms service identity: %s", err)
	}
	err = addFolderBinding2(config.NewResourceManagerV3Client(config.UserAgent), autokeyFolder.Name, fmt.Sprintf("user:%s", curUserEmail), []string{"roles/cloudkms.admin"})
	if err != nil {
		t.Errorf("unable to bootstrap KMS keyHandle. Cannot assign cloudkms.admin role to current user on autokey test folder: %s", err)
	}
	err = addProjectBinding(config.NewResourceManagerV3Client(config.UserAgent), kmsProjectID, fmt.Sprintf("user:%s", curUserEmail), []string{"roles/resourcemanager.projectIamAdmin", "roles/cloudkms.admin"})
	if err != nil {
		t.Errorf("unable to bootstrap KMS keyHandle. Cannot assign cloudkms.admin and projectIamAdmin role to current user on kms project: %s", err)
	}
	err = addProjectBinding(config.NewResourceManagerV3Client(config.UserAgent), kmsProjectID, fmt.Sprintf("serviceAccount:%s", kmsSAEmail), []string{"roles/cloudkms.admin"})
	if err != nil {
		t.Errorf("unable to bootstrap KMS keyHandle. Cannot assign cloudkms.admin role to cloudkms service identity on kms project: %s", err)
	}

	// create and setup resource folder to host keyhandle
	resourceProject := BootstrapProjectWithParent(t, defaultAutokeyTestResourceProject, envvar.GetTestBillingAccountFromEnv(t), parent, []string{})
	return autokeyFolder, kmsProject, resourceProject
}

// GenerateCloudKmsServiceIdentity generates cloud kms service identity within a project
func GenerateCloudKmsServiceIdentity(config *transport_tpg.Config, projectNum string) (string, error) {
	url := fmt.Sprintf("https://serviceusage.googleapis.com/v1beta1/projects/%s/services/%s:generateServiceIdentity", projectNum, CloudKmsSrviceName)

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   projectNum,
		RawURL:    url,
		UserAgent: config.UserAgent,
		Timeout:   time.Minute * 4,
	})
	if err != nil {
		return "", fmt.Errorf("error creating cloudkms service identity: %s", err)
	}

	var opRes map[string]interface{}
	err = tpgservicusage.ServiceUsageOperationWaitTimeWithResponse(
		config, res, &opRes, projectNum, "Creating cloudkms service identity", config.UserAgent,
		time.Minute*4)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("service-%s@gcp-sa-cloudkms.iam.gserviceaccount.com", projectNum), nil
}

func addProjectBinding(crmService *resourceManagerV3.Service, projectID string, member string, roles []string) error {
	return addBinding(crmService, "project", projectID, member, roles)
}

func addFolderBinding2(crmService *resourceManagerV3.Service, folderID string, member string, roles []string) error {
	return addBinding(crmService, "folder", folderID, member, roles)
}

// addBinding adds the member to the project's IAM policy
func addBinding(crmService *resourceManagerV3.Service, resourceType string, resourceID string, member string, roles []string) error {

	policy, err := getPolicy(crmService, resourceType, resourceID)
	if err != nil {
		return err
	}

	// Find the policy binding for role. Only one binding can have the role.
	var binding *resourceManagerV3.Binding
	for _, role := range roles {
		for _, b := range policy.Bindings {
			if b.Role == role {
				binding = b
				break
			}
		}

		if binding != nil {
			// If the binding exists, adds the member to the binding
			binding.Members = append(binding.Members, member)
		} else {
			// If the binding does not exist, adds a new binding to the policy
			binding = &resourceManagerV3.Binding{
				Role:    role,
				Members: []string{member},
			}
			policy.Bindings = append(policy.Bindings, binding)
		}
	}
	setPolicy(crmService, resourceType, resourceID, policy)
	return nil
}

// getPolicy gets the IAM policy on input resourceID
// resourceType can be "project" or "folder"
func getPolicy(crmService *resourceManagerV3.Service, resourceType string, resourceID string) (*resourceManagerV3.Policy, error) {

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	request := new(resourceManagerV3.GetIamPolicyRequest)
	var policy *resourceManagerV3.Policy
	var err error
	if resourceType == "project" {
		policy, err = crmService.Projects.GetIamPolicy(resourceID, request).Do()
	} else if resourceType == "folder" {
		policy, err = crmService.Folders.GetIamPolicy(resourceID, request).Do()
	} else {
		return nil, fmt.Errorf("invalid resourceType, supported values: project or folder")
	}
	if err != nil {
		return nil, fmt.Errorf("error getting iam policy: %s", err)
	}
	return policy, nil
}

// setPolicy sets the IAM policy on input resourceID
// resourceType can be "project" or "folder"
func setPolicy(crmService *resourceManagerV3.Service, resourceType string, resourceID string, policy *resourceManagerV3.Policy) error {

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	request := new(resourceManagerV3.SetIamPolicyRequest)
	request.Policy = policy
	var err error
	if resourceType == "project" {
		_, err = crmService.Projects.SetIamPolicy(resourceID, request).Do()
	} else if resourceType == "folder" {
		_, err = crmService.Folders.SetIamPolicy(resourceID, request).Do()
	} else {
		return fmt.Errorf("invalid resourceType, supported values: project or folder")
	}
	if err != nil {
		return fmt.Errorf("error setting iam policy: %s", err)
	}
	return nil
}

func BootstrapKMSKeyWithPurposeInLocationAndName(t *testing.T, purpose, locationID, keyShortName string) BootstrappedKMS {
	config := BootstrapConfig(t)
	if config == nil {
		return BootstrappedKMS{
			&cloudkms.KeyRing{},
			&cloudkms.CryptoKey{},
			nil,
		}
	}

	projectID := envvar.GetTestProjectFromEnv()
	keyRingParent := fmt.Sprintf("projects/%s/locations/%s", projectID, locationID)
	keyRingName := fmt.Sprintf("%s/keyRings/%s", keyRingParent, SharedKeyRing)
	keyParent := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", projectID, locationID, SharedKeyRing)
	keyName := fmt.Sprintf("%s/cryptoKeys/%s", keyParent, keyShortName)

	// Get or Create the hard coded shared keyring for testing
	kmsClient := config.NewKmsClient(config.UserAgent)
	keyRing, err := kmsClient.Projects.Locations.KeyRings.Get(keyRingName).Do()
	if err != nil {
		if transport_tpg.IsGoogleApiErrorWithCode(err, 404) {
			keyRing, err = kmsClient.Projects.Locations.KeyRings.Create(keyRingParent, &cloudkms.KeyRing{}).
				KeyRingId(SharedKeyRing).Do()
			if err != nil {
				t.Errorf("Unable to bootstrap KMS key. Cannot create keyRing: %s", err)
			}
		} else {
			t.Errorf("Unable to bootstrap KMS key. Cannot retrieve keyRing: %s", err)
		}
	}

	if keyRing == nil {
		t.Fatalf("Unable to bootstrap KMS key. keyRing is nil!")
	}

	// Get or Create the hard coded, shared crypto key for testing
	cryptoKey, err := kmsClient.Projects.Locations.KeyRings.CryptoKeys.Get(keyName).Do()
	if err != nil {
		if transport_tpg.IsGoogleApiErrorWithCode(err, 404) {
			algos := map[string]string{
				"ENCRYPT_DECRYPT":    "GOOGLE_SYMMETRIC_ENCRYPTION",
				"ASYMMETRIC_SIGN":    "RSA_SIGN_PKCS1_4096_SHA256",
				"ASYMMETRIC_DECRYPT": "RSA_DECRYPT_OAEP_4096_SHA256",
			}
			template := cloudkms.CryptoKeyVersionTemplate{
				Algorithm: algos[purpose],
			}

			newKey := cloudkms.CryptoKey{
				Purpose:         purpose,
				VersionTemplate: &template,
			}

			cryptoKey, err = kmsClient.Projects.Locations.KeyRings.CryptoKeys.Create(keyParent, &newKey).
				CryptoKeyId(keyShortName).Do()
			if err != nil {
				t.Errorf("Unable to bootstrap KMS key. Cannot create new CryptoKey: %s", err)
			}

		} else {
			t.Errorf("Unable to bootstrap KMS key. Cannot call CryptoKey service: %s", err)
		}
	}

	if cryptoKey == nil {
		t.Fatalf("Unable to bootstrap KMS key. CryptoKey is nil!")
	}

	// TODO(b/372305432): Use the pagination properly.
	ckvResp, err := kmsClient.Projects.Locations.KeyRings.CryptoKeys.CryptoKeyVersions.List(keyName).Do()
	if err != nil {
		t.Fatalf("Unable to list cryptoKeyVersions: %v", err)
	}

	return BootstrappedKMS{
		keyRing,
		cryptoKey,
		ckvResp.CryptoKeyVersions,
	}
}

var serviceAccountPrefix = "tf-bootstrap-sa-"
var serviceAccountDisplay = "Bootstrapped Service Account for Terraform tests"

// Some tests need a second service account, other than the test runner, to assert functionality on.
// This provides a well-known service account that can be used when dynamically creating a service
// account isn't an option.
func getOrCreateServiceAccount(config *transport_tpg.Config, project, serviceAccountEmail string) (*iam.ServiceAccount, error) {
	name := fmt.Sprintf("projects/%s/serviceAccounts/%s@%s.iam.gserviceaccount.com", project, serviceAccountEmail, project)
	log.Printf("[DEBUG] Verifying %s as bootstrapped service account.\n", name)

	sa, err := config.NewIamClient(config.UserAgent).Projects.ServiceAccounts.Get(name).Do()
	if err != nil && !transport_tpg.IsGoogleApiErrorWithCode(err, 404) {
		return nil, fmt.Errorf("encountered a non-404 error when looking for bootstrapped service account %s: %w", name, err)
	}

	if sa == nil {
		log.Printf("[DEBUG] Account missing. Creating %s as bootstrapped service account.\n", name)
		sa = &iam.ServiceAccount{
			DisplayName: serviceAccountDisplay,
		}

		r := &iam.CreateServiceAccountRequest{
			AccountId:      serviceAccountEmail,
			ServiceAccount: sa,
		}
		sa, err = config.NewIamClient(config.UserAgent).Projects.ServiceAccounts.Create("projects/"+project, r).Do()
		if err != nil {
			return nil, fmt.Errorf("error when creating bootstrapped service account %s: %w", name, err)
		}
	}

	return sa, nil
}

// In order to test impersonation we need to grant the testRunner's account the ability to grant tokens
// on a different service account. Granting permissions takes time and there is no operation to wait on
// so instead this creates a single service account once per test-suite with the correct permissions.
// The first time this test is run it will fail, but subsequent runs will succeed.
func impersonationServiceAccountPermissions(config *transport_tpg.Config, sa *iam.ServiceAccount, testRunner string) error {
	log.Printf("[DEBUG] Setting service account permissions.\n")
	policy := iam.Policy{
		Bindings: []*iam.Binding{},
	}

	binding := &iam.Binding{
		Role:    "roles/iam.serviceAccountTokenCreator",
		Members: []string{"serviceAccount:" + sa.Email, "serviceAccount:" + testRunner},
	}
	policy.Bindings = append(policy.Bindings, binding)

	// Overwrite the roles each time on this service account. This is because this account is
	// only created for the test suite and will stop snowflaking of permissions to get tests
	// to run. Overwriting permissions on 1 service account shouldn't affect others.
	_, err := config.NewIamClient(config.UserAgent).Projects.ServiceAccounts.SetIamPolicy(sa.Name, &iam.SetIamPolicyRequest{
		Policy: &policy,
	}).Do()
	if err != nil {
		return err
	}

	return nil
}

// A separate testId should be used for each test, to create separate service accounts for each,
// and avoid race conditions where the policy of the same service account is being modified by 2
// tests at once. This is needed as long as the function overwrites the policy on every run.
func BootstrapServiceAccount(t *testing.T, testId, testRunner string) string {
	project := envvar.GetTestProjectFromEnv()
	serviceAccountEmail := serviceAccountPrefix + testId

	config := BootstrapConfig(t)
	if config == nil {
		return ""
	}

	sa, err := getOrCreateServiceAccount(config, project, serviceAccountEmail)
	if err != nil {
		t.Fatalf("Bootstrapping failed. Cannot retrieve service account, %s", err)
	}

	err = impersonationServiceAccountPermissions(config, sa, testRunner)
	if err != nil {
		t.Fatalf("Bootstrapping failed. Cannot set service account permissions, %s", err)
	}

	return sa.Email
}

const SharedTestADDomainPrefix = "tf-bootstrap-ad"

func BootstrapSharedTestADDomain(t *testing.T, testId string, networkName string) string {
	project := envvar.GetTestProjectFromEnv()
	sharedADDomain := fmt.Sprintf("%s.%s.com", SharedTestADDomainPrefix, testId)
	adDomainName := fmt.Sprintf("projects/%s/locations/global/domains/%s", project, sharedADDomain)

	config := BootstrapConfig(t)
	if config == nil {
		return ""
	}

	log.Printf("[DEBUG] Getting shared test active directory domain %q", adDomainName)
	getURL := fmt.Sprintf("%s%s", config.ActiveDirectoryBasePath, adDomainName)
	_, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   project,
		RawURL:    getURL,
		UserAgent: config.UserAgent,
		Timeout:   4 * time.Minute,
	})
	if err != nil && transport_tpg.IsGoogleApiErrorWithCode(err, 404) {
		log.Printf("[DEBUG] AD domain %q not found, bootstrapping", sharedADDomain)
		postURL := fmt.Sprintf("%sprojects/%s/locations/global/domains?domainName=%s", config.ActiveDirectoryBasePath, project, sharedADDomain)
		domainObj := map[string]interface{}{
			"locations":          []string{"us-central1"},
			"reservedIpRange":    "10.0.1.0/24",
			"authorizedNetworks": []string{fmt.Sprintf("projects/%s/global/networks/%s", project, networkName)},
		}

		_, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "POST",
			Project:   project,
			RawURL:    postURL,
			UserAgent: config.UserAgent,
			Body:      domainObj,
			Timeout:   60 * time.Minute,
		})
		if err != nil {
			t.Fatalf("Error bootstrapping shared active directory domain %q: %s", adDomainName, err)
		}

		log.Printf("[DEBUG] Waiting for active directory domain creation to finish")
	}

	_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   project,
		RawURL:    getURL,
		UserAgent: config.UserAgent,
		Timeout:   4 * time.Minute,
	})

	if err != nil {
		t.Fatalf("Error getting shared active directory domain %q: %s", adDomainName, err)
	}

	return sharedADDomain
}

const SharedTestNetworkPrefix = "tf-bootstrap-net-"

// BootstrapSharedTestNetwork will return a persistent compute network for a
// test or set of tests.
//
// Usage 1
// Resources like service_networking_connection use a consumer network and
// create a complementing tenant network which we don't control. These tenant
// networks never get cleaned up and they can accumulate to the point where a
// limit is reached for the organization. By reusing a consumer network across
// test runs, we can reduce the number of tenant networks that are needed.
// See b/146351146 for more context.
//
// Usage 2
// Bootstrap networks used in tests (gke clusters, dataproc clusters...)
// to avoid traffic to the default network
//
// testId specifies the test for which a shared network is used/initialized.
// Note that if the network is being used for a service_networking_connection,
// the same testId should generally not be used across tests, to avoid race
// conditions where multiple tests attempt to modify the connection at once.
//
// Returns the name of a network, creating it if it hasn't been created in the
// test project.
func BootstrapSharedTestNetwork(t *testing.T, testId string) string {
	project := envvar.GetTestProjectFromEnv()
	networkName := SharedTestNetworkPrefix + testId

	config := BootstrapConfig(t)
	if config == nil {
		return ""
	}

	log.Printf("[DEBUG] Getting shared test network %q", networkName)
	_, err := config.NewComputeClient(config.UserAgent).Networks.Get(project, networkName).Do()
	if err != nil && transport_tpg.IsGoogleApiErrorWithCode(err, 404) {
		log.Printf("[DEBUG] Network %q not found, bootstrapping", networkName)
		url := fmt.Sprintf("%sprojects/%s/global/networks", config.ComputeBasePath, project)
		netObj := map[string]interface{}{
			"name":                  networkName,
			"autoCreateSubnetworks": false,
		}

		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "POST",
			Project:   project,
			RawURL:    url,
			UserAgent: config.UserAgent,
			Body:      netObj,
			Timeout:   4 * time.Minute,
		})
		if err != nil {
			t.Fatalf("Error bootstrapping shared test network %q: %s", networkName, err)
		}

		log.Printf("[DEBUG] Waiting for network creation to finish")
		err = tpgcompute.ComputeOperationWaitTime(config, res, project, "Error bootstrapping shared test network", config.UserAgent, 4*time.Minute)
		if err != nil {
			t.Fatalf("Error bootstrapping shared test network %q: %s", networkName, err)
		}
	}

	network, err := config.NewComputeClient(config.UserAgent).Networks.Get(project, networkName).Do()
	if err != nil {
		t.Errorf("Error getting shared test network %q: %s", networkName, err)
	}
	if network == nil {
		t.Fatalf("Error getting shared test network %q: is nil", networkName)
	}
	return network.Name
}

type AddressSettings struct {
	PrefixLength int
}

func AddressWithPrefixLength(prefixLength int) func(*AddressSettings) {
	return func(settings *AddressSettings) {
		settings.PrefixLength = prefixLength
	}
}

func NewAddressSettings(options ...func(*AddressSettings)) *AddressSettings {
	settings := &AddressSettings{
		PrefixLength: 16, // default prefix length
	}

	for _, o := range options {
		o(settings)
	}
	return settings
}

const SharedTestGlobalAddressPrefix = "tf-bootstrap-addr-"

// params are the functions to set compute global address
func BootstrapSharedTestGlobalAddress(t *testing.T, testId string, params ...func(*AddressSettings)) string {
	project := envvar.GetTestProjectFromEnv()
	addressName := SharedTestGlobalAddressPrefix + testId
	networkName := BootstrapSharedTestNetwork(t, testId)
	networkId := fmt.Sprintf("projects/%v/global/networks/%v", project, networkName)

	config := BootstrapConfig(t)
	if config == nil {
		return ""
	}

	log.Printf("[DEBUG] Getting shared test global address %q", addressName)
	_, err := config.NewComputeClient(config.UserAgent).GlobalAddresses.Get(project, addressName).Do()
	if err != nil && transport_tpg.IsGoogleApiErrorWithCode(err, 404) {
		log.Printf("[DEBUG] Global address %q not found, bootstrapping", addressName)
		url := fmt.Sprintf("%sprojects/%s/global/addresses", config.ComputeBasePath, project)

		settings := NewAddressSettings(params...)

		netObj := map[string]interface{}{
			"name":          addressName,
			"address_type":  "INTERNAL",
			"purpose":       "VPC_PEERING",
			"prefix_length": settings.PrefixLength,
			"network":       networkId,
		}

		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "POST",
			Project:   project,
			RawURL:    url,
			UserAgent: config.UserAgent,
			Body:      netObj,
			Timeout:   4 * time.Minute,
		})
		if err != nil {
			t.Fatalf("Error bootstrapping shared test global address %q: %s", addressName, err)
		}

		log.Printf("[DEBUG] Waiting for global address creation to finish")
		err = tpgcompute.ComputeOperationWaitTime(config, res, project, "Error bootstrapping shared test global address", config.UserAgent, 4*time.Minute)
		if err != nil {
			t.Fatalf("Error bootstrapping shared test global address %q: %s", addressName, err)
		}
	}

	address, err := config.NewComputeClient(config.UserAgent).GlobalAddresses.Get(project, addressName).Do()
	if err != nil {
		t.Errorf("Error getting shared test global address %q: %s", addressName, err)
	}
	if address == nil {
		t.Fatalf("Error getting shared test global address %q: is nil", addressName)
	}
	return address.Name
}

type ServiceNetworkSettings struct {
	PrefixLength  int
	ParentService string
}

func ServiceNetworkWithPrefixLength(prefixLength int) func(*ServiceNetworkSettings) {
	return func(settings *ServiceNetworkSettings) {
		settings.PrefixLength = prefixLength
	}
}

func ServiceNetworkWithParentService(parentService string) func(*ServiceNetworkSettings) {
	return func(settings *ServiceNetworkSettings) {
		settings.ParentService = parentService
	}
}

func NewServiceNetworkSettings(options ...func(*ServiceNetworkSettings)) *ServiceNetworkSettings {
	settings := &ServiceNetworkSettings{
		PrefixLength:  16,                                 // default prefix length
		ParentService: "servicenetworking.googleapis.com", // default parent service
	}

	for _, o := range options {
		o(settings)
	}
	return settings
}

// BootstrapSharedServiceNetworkingConnection will create a shared network
// if it hasn't been created in the test project, a global address
// if it hasn't been created in the test project, and a service networking connection
// if it hasn't been created in the test project.
//
// params are the functions to set compute global address
//
// BootstrapSharedServiceNetworkingConnection returns a persistent compute network name
// for a test or set of tests.
//
// To delete a service networking connection, all of the service instances that use that connection
// must be deleted first. After the service instances are deleted, some service producers delay the deletion
// utnil a waiting period has passed. For example, after four days that you delete a SQL instance,
// the service networking connection can be deleted.
// That is the reason to use the shared service networking connection for test resources.
// https://cloud.google.com/vpc/docs/configure-private-services-access#removing-connection
//
// testId specifies the test for which a shared network and a global address are used/initialized.
func BootstrapSharedServiceNetworkingConnection(t *testing.T, testId string, params ...func(*ServiceNetworkSettings)) string {
	settings := NewServiceNetworkSettings(params...)
	parentService := "services/" + settings.ParentService
	projectId := envvar.GetTestProjectFromEnv()

	config := BootstrapConfig(t)
	if config == nil {
		return ""
	}

	// Get project number by calling the API
	crmClient := config.NewResourceManagerClient(config.UserAgent)
	project, err := crmClient.Projects.Get(projectId).Do()
	if err != nil {
		t.Fatalf("Error getting project: %s", err)
	}

	networkName := SharedTestNetworkPrefix + testId
	networkId := fmt.Sprintf("projects/%v/global/networks/%v", project.ProjectNumber, networkName)
	globalAddressName := BootstrapSharedTestGlobalAddress(t, testId, AddressWithPrefixLength(settings.PrefixLength))

	readCall := config.NewServiceNetworkingClient(config.UserAgent).Services.Connections.List(parentService).Network(networkId)
	if config.UserProjectOverride {
		readCall.Header().Add("X-Goog-User-Project", projectId)
	}
	response, err := readCall.Do()
	if err != nil {
		t.Errorf("Error getting shared test service networking connection: %s", err)
	}

	var connection *servicenetworking.Connection
	for _, c := range response.Connections {
		if c.Network == networkId {
			connection = c
			break
		}
	}

	if connection == nil {
		log.Printf("[DEBUG] Service networking connection not found, bootstrapping")

		connection := &servicenetworking.Connection{
			Network:               networkId,
			ReservedPeeringRanges: []string{globalAddressName},
		}

		createCall := config.NewServiceNetworkingClient(config.UserAgent).Services.Connections.Create(parentService, connection)
		if config.UserProjectOverride {
			createCall.Header().Add("X-Goog-User-Project", projectId)
		}
		op, err := createCall.Do()
		if err != nil {
			t.Fatalf("Error bootstrapping shared test service networking connection: %s", err)
		}

		log.Printf("[DEBUG] Waiting for service networking connection creation to finish")
		if err := tpgservicenetworking.ServiceNetworkingOperationWaitTimeHW(config, op, "Create Service Networking Connection", config.UserAgent, projectId, 4*time.Minute); err != nil {
			t.Fatalf("Error bootstrapping shared test service networking connection: %s", err)
		}
	}

	log.Printf("[DEBUG] Getting shared test service networking connection")

	return networkName
}

var SharedServicePerimeterProjectPrefix = "tf-bootstrap-sp-"

func BootstrapServicePerimeterProjects(t *testing.T, desiredProjects int) []*cloudresourcemanager.Project {
	config := BootstrapConfig(t)
	if config == nil {
		return nil
	}

	org := envvar.GetTestOrgFromEnv(t)

	// The filter endpoint works differently if you provide both the parent id and parent type, and
	// doesn't seem to allow for prefix matching. Don't change this to include the parent type unless
	// that API behavior changes.
	prefixFilter := fmt.Sprintf("id:%s* parent.id:%s", SharedServicePerimeterProjectPrefix, org)
	res, err := config.NewResourceManagerClient(config.UserAgent).Projects.List().Filter(prefixFilter).Do()
	if err != nil {
		t.Fatalf("Error getting shared test projects: %s", err)
	}

	projects := res.Projects
	for len(projects) < desiredProjects {
		pid := SharedServicePerimeterProjectPrefix + RandString(t, 10)
		project := &cloudresourcemanager.Project{
			ProjectId: pid,
			Name:      "TF Service Perimeter Test",
			Parent: &cloudresourcemanager.ResourceId{
				Type: "organization",
				Id:   org,
			},
		}
		op, err := config.NewResourceManagerClient(config.UserAgent).Projects.Create(project).Do()
		if err != nil {
			t.Fatalf("Error bootstrapping shared test project: %s", err)
		}

		opAsMap, err := tpgresource.ConvertToMap(op)
		if err != nil {
			t.Fatalf("Error bootstrapping shared test project: %s", err)
		}

		err = resourcemanager.ResourceManagerOperationWaitTime(config, opAsMap, "creating project", config.UserAgent, 4)
		if err != nil {
			t.Fatalf("Error bootstrapping shared test project: %s", err)
		}

		p, err := config.NewResourceManagerClient(config.UserAgent).Projects.Get(pid).Do()
		if err != nil {
			t.Fatalf("Error getting shared test project: %s", err)
		}
		projects = append(projects, p)
	}

	return projects
}

// BootstrapFolder creates or get a folder having a input folderDisplayName within a TestOrgEnv
func BootstrapFolder(t *testing.T, folderDisplayName string) *resourceManagerV3.Folder {
	config := BootstrapConfig(t)
	if config == nil {
		return nil
	}

	crmClient := config.NewResourceManagerV3Client(config.UserAgent)
	searchQuery := fmt.Sprintf("displayName=%s", folderDisplayName)
	folderSearchResp, err := crmClient.Folders.Search().Query(searchQuery).Do()
	if err != nil {
		t.Fatalf("error searching for folder with displayName: %s", folderDisplayName)
	}
	var folder *resourceManagerV3.Folder
	if len(folderSearchResp.Folders) == 0 {
		op, err := crmClient.Folders.Create(&resourceManagerV3.Folder{
			DisplayName: folderDisplayName,
			Parent:      fmt.Sprintf("organizations/%s", envvar.GetTestOrgFromEnv(t)),
		}).Do()
		if err != nil {
			t.Fatalf("error bootstrapping test folder: %s", err)
		}

		opAsMap, err := tpgresource.ConvertToMap(op)
		if err != nil {
			t.Fatalf("error converting folder operation map: %s", err)
		}
		var responseMap map[string]interface{}
		err = resourcemanager.ResourceManagerOperationWaitTimeWithResponse(config, opAsMap, &responseMap, "creating folder", config.UserAgent, 4*time.Minute)
		if err != nil {
			t.Fatalf("error waiting for create folder operation: %s", err)
		}
		folder, err = crmClient.Folders.Get(responseMap["name"].(string)).Do()
		if err != nil {
			t.Fatalf("error getting folder: %s", err)
		}
	} else {
		folder = folderSearchResp.Folders[0]
	}

	if folder.State == "DELETE_REQUESTED" {
		_, err := crmClient.Folders.Undelete(folder.Name, &resourceManagerV3.UndeleteFolderRequest{}).Do()
		if err != nil {
			t.Fatalf("error undeleting folder: %s", err)
		}
	}
	return folder
}

// BootstrapProject will create or get a project named
// "<projectIDPrefix><projectIDSuffix>" that will persist across test runs,
// where projectIDSuffix is based off of getTestProjectFromEnv(). The reason
// for the naming is to isolate bootstrapped projects by test environment.
// Given the existing projects being used by our team, the prefix provided to
// this function can be no longer than 18 characters.
func BootstrapProject(t *testing.T, projectIDPrefix, billingAccount string, services []string) *cloudresourcemanager.Project {
	org := envvar.GetTestOrgFromEnv(t)
	parent := &cloudresourcemanager.ResourceId{
		Type: "organization",
		Id:   org,
	}
	projectIDSuffix := strings.Replace(envvar.GetTestProjectFromEnv(), "ci-test-project-", "", 1)
	projectID := projectIDPrefix + projectIDSuffix

	return BootstrapProjectWithParent(t, projectID, billingAccount, parent, services)
}

func BootstrapProjectWithParent(t *testing.T, projectID string, billingAccount string, parent *cloudresourcemanager.ResourceId, services []string) *cloudresourcemanager.Project {
	config := BootstrapConfig(t)
	if config == nil {
		return nil
	}
	crmClient := config.NewResourceManagerClient(config.UserAgent)

	project, err := crmClient.Projects.Get(projectID).Do()
	if err != nil {
		if !transport_tpg.IsGoogleApiErrorWithCode(err, 403) {
			t.Fatalf("Error getting bootstrapped project: %s", err)
		}
		op, err := crmClient.Projects.Create(&cloudresourcemanager.Project{
			ProjectId: projectID,
			Name:      "Bootstrapped Test Project",
			Parent:    parent,
		}).Do()
		if err != nil {
			t.Fatalf("Error creating bootstrapped test project: %s", err)
		}

		opAsMap, err := tpgresource.ConvertToMap(op)
		if err != nil {
			t.Fatalf("Error converting create project operation to map: %s", err)
		}

		err = resourcemanager.ResourceManagerOperationWaitTime(config, opAsMap, "creating project", config.UserAgent, 4*time.Minute)
		if err != nil {
			t.Fatalf("Error waiting for create project operation: %s", err)
		}

		project, err = crmClient.Projects.Get(projectID).Do()
		if err != nil {
			t.Fatalf("Error getting bootstrapped project: %s", err)
		}

	}

	if project.LifecycleState == "DELETE_REQUESTED" {
		_, err := crmClient.Projects.Undelete(projectID, &cloudresourcemanager.UndeleteProjectRequest{}).Do()
		if err != nil {
			t.Fatalf("Error undeleting bootstrapped project: %s", err)
		}
	}

	if billingAccount != "" {
		billingClient := config.NewBillingClient(config.UserAgent)
		var pbi *cloudbilling.ProjectBillingInfo
		err = transport_tpg.Retry(transport_tpg.RetryOptions{
			RetryFunc: func() error {
				var reqErr error
				pbi, reqErr = billingClient.Projects.GetBillingInfo(resourcemanager.PrefixedProject(projectID)).Do()
				return reqErr
			},
			Timeout: 30 * time.Second,
		})
		if err != nil {
			t.Fatalf("Error getting billing info for project %q: %v", projectID, err)
		}
		if strings.TrimPrefix(pbi.BillingAccountName, "billingAccounts/") != billingAccount {
			pbi.BillingAccountName = "billingAccounts/" + billingAccount
			err := transport_tpg.Retry(transport_tpg.RetryOptions{
				RetryFunc: func() error {
					_, err := config.NewBillingClient(config.UserAgent).Projects.UpdateBillingInfo(resourcemanager.PrefixedProject(projectID), pbi).Do()
					return err
				},
				Timeout: 2 * time.Minute,
			})
			if err != nil {
				t.Fatalf("Error setting billing account for project %q to %q: %s", projectID, billingAccount, err)
			}
		}
	}

	if len(services) > 0 {

		enabledServices, err := resourcemanager.ListCurrentlyEnabledServices(projectID, "", config.UserAgent, config, 1*time.Minute)
		if err != nil {
			t.Fatalf("Error listing services for project %q: %s", projectID, err)
		}

		servicesToEnable := make([]string, 0, len(services))
		for _, service := range services {
			if _, ok := enabledServices[service]; !ok {
				servicesToEnable = append(servicesToEnable, service)
			}
		}

		if len(servicesToEnable) > 0 {
			if err := resourcemanager.EnableServiceUsageProjectServices(servicesToEnable, projectID, "", config.UserAgent, config, 10*time.Minute); err != nil {
				t.Fatalf("Error enabling services for project %q: %s", projectID, err)
			}
		}
	}

	return project
}

// BootstrapConfig returns a Config pulled from the environment.
func BootstrapConfig(t *testing.T) *transport_tpg.Config {
	if v := os.Getenv("TF_ACC"); v == "" {
		t.Skip("Acceptance tests and bootstrapping skipped unless env 'TF_ACC' set")
		return nil
	}

	config := &transport_tpg.Config{
		Credentials:               envvar.GetTestCredsFromEnv(),
		ImpersonateServiceAccount: envvar.GetTestImpersonateServiceAccountFromEnv(),
		Project:                   envvar.GetTestProjectFromEnv(),
		Region:                    envvar.GetTestRegionFromEnv(),
		Zone:                      envvar.GetTestZoneFromEnv(),
	}

	transport_tpg.ConfigureBasePaths(config)

	if err := config.LoadAndValidate(context.Background()); err != nil {
		t.Fatalf("Bootstrapping failed. Unable to load test config: %s", err)
	}
	return config
}

// SQL Instance names are not reusable for a week after deletion
const SharedTestSQLInstanceNamePrefix = "tf-bootstrap-"

// BootstrapSharedSQLInstanceBackupRun will return a shared SQL db instance that
// has a backup created for it.
func BootstrapSharedSQLInstanceBackupRun(t *testing.T) string {
	project := envvar.GetTestProjectFromEnv()

	config := BootstrapConfig(t)
	if config == nil {
		return ""
	}

	log.Printf("[DEBUG] Getting list of existing sql instances")

	instances, err := config.NewSqlAdminClient(config.UserAgent).Instances.List(project).Do()
	if err != nil {
		t.Fatalf("Unable to bootstrap SQL Instance. Cannot retrieve instance list: %s", err)
	}

	var bootstrapInstance *sqladmin.DatabaseInstance

	// Look for any existing bootstrap instances
	for _, i := range instances.Items {
		if strings.HasPrefix(i.Name, SharedTestSQLInstanceNamePrefix) {
			bootstrapInstance = i
			break
		}
	}

	if bootstrapInstance == nil {
		bootstrapInstanceName := SharedTestSQLInstanceNamePrefix + RandString(t, 10)
		log.Printf("[DEBUG] Bootstrap SQL Instance not found, bootstrapping new instance %s", bootstrapInstanceName)

		backupConfig := &sqladmin.BackupConfiguration{
			Enabled:                    true,
			PointInTimeRecoveryEnabled: true,
		}
		settings := &sqladmin.Settings{
			Tier:                "db-f1-micro",
			BackupConfiguration: backupConfig,
		}
		bootstrapInstance = &sqladmin.DatabaseInstance{
			Name:            bootstrapInstanceName,
			Region:          "us-central1",
			Settings:        settings,
			DatabaseVersion: "POSTGRES_11",
		}

		var op *sqladmin.Operation
		err = transport_tpg.Retry(transport_tpg.RetryOptions{
			RetryFunc: func() (operr error) {
				op, operr = config.NewSqlAdminClient(config.UserAgent).Instances.Insert(project, bootstrapInstance).Do()
				return operr
			},
			Timeout:              20 * time.Minute,
			ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.IsSqlOperationInProgressError},
		})
		if err != nil {
			t.Fatalf("Error, failed to create instance %s: %s", bootstrapInstance.Name, err)
		}
		err = sql.SqlAdminOperationWaitTime(config, op, project, "Create Instance", config.UserAgent, 40*time.Minute)
		if err != nil {
			t.Fatalf("Error, failed to create instance %s: %s", bootstrapInstance.Name, err)
		}
	}

	// Look for backups in bootstrap instance
	res, err := config.NewSqlAdminClient(config.UserAgent).BackupRuns.List(project, bootstrapInstance.Name).Do()
	if err != nil {
		t.Fatalf("Unable to bootstrap SQL Instance. Cannot retrieve backup list: %s", err)
	}
	backupsList := res.Items
	if len(backupsList) == 0 {
		log.Printf("[DEBUG] No backups found for %s, creating backup", bootstrapInstance.Name)
		backupRun := &sqladmin.BackupRun{
			Instance: bootstrapInstance.Name,
		}

		var op *sqladmin.Operation
		err = transport_tpg.Retry(transport_tpg.RetryOptions{
			RetryFunc: func() (operr error) {
				op, operr = config.NewSqlAdminClient(config.UserAgent).BackupRuns.Insert(project, bootstrapInstance.Name, backupRun).Do()
				return operr
			},
			Timeout:              20 * time.Minute,
			ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.IsSqlOperationInProgressError},
		})
		if err != nil {
			t.Fatalf("Error, failed to create instance backup: %s", err)
		}
		err = sql.SqlAdminOperationWaitTime(config, op, project, "Backup Instance", config.UserAgent, 20*time.Minute)
		if err != nil {
			t.Fatalf("Error, failed to create instance backup: %s", err)
		}
	}

	return bootstrapInstance.Name
}

func BootstrapSharedCaPoolInLocation(t *testing.T, location string) string {
	project := envvar.GetTestProjectFromEnv()
	poolName := "static-ca-pool"

	config := BootstrapConfig(t)
	if config == nil {
		return ""
	}

	log.Printf("[DEBUG] Getting shared CA pool %q", poolName)
	url := fmt.Sprintf("%sprojects/%s/locations/%s/caPools/%s", config.PrivatecaBasePath, project, location, poolName)
	_, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   project,
		RawURL:    url,
		UserAgent: config.UserAgent,
	})
	if err != nil {
		log.Printf("[DEBUG] CA pool %q not found, bootstrapping", poolName)
		poolObj := map[string]interface{}{
			"tier": "ENTERPRISE",
		}
		createUrl := fmt.Sprintf("%sprojects/%s/locations/%s/caPools?caPoolId=%s", config.PrivatecaBasePath, project, location, poolName)
		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "POST",
			Project:   project,
			RawURL:    createUrl,
			UserAgent: config.UserAgent,
			Body:      poolObj,
			Timeout:   4 * time.Minute,
		})
		if err != nil {
			t.Fatalf("Error bootstrapping shared CA pool %q: %s", poolName, err)
		}

		log.Printf("[DEBUG] Waiting for CA pool creation to finish")
		var opRes map[string]interface{}
		err = privateca.PrivatecaOperationWaitTimeWithResponse(
			config, res, &opRes, project, "Creating CA pool", config.UserAgent,
			4*time.Minute)
		if err != nil {
			t.Errorf("Error getting shared CA pool %q: %s", poolName, err)
		}
		_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "GET",
			Project:   project,
			RawURL:    url,
			UserAgent: config.UserAgent,
		})
		if err != nil {
			t.Errorf("Error getting shared CA pool %q: %s", poolName, err)
		}
	}
	return poolName
}

func BootstrapSubnetForDataprocBatches(t *testing.T, subnetName string, networkName string) string {
	subnetOptions := map[string]interface{}{
		"privateIpGoogleAccess": true,
	}
	return BootstrapSubnetWithOverrides(t, subnetName, networkName, subnetOptions)
}

func BootstrapSubnet(t *testing.T, subnetName string, networkName string) string {
	return BootstrapSubnetWithOverrides(t, subnetName, networkName, make(map[string]interface{}))
}

func BootstrapSubnetWithFirewallForDataprocBatches(t *testing.T, testId string, subnetName string) string {
	networkName := BootstrapSharedTestNetwork(t, testId)
	subnetworkName := BootstrapSubnetForDataprocBatches(t, subnetName, networkName)
	BootstrapFirewallForDataprocSharedNetwork(t, subnetName, networkName)
	return subnetworkName
}

func BootstrapSubnetWithOverrides(t *testing.T, subnetName string, networkName string, subnetOptions map[string]interface{}) string {
	projectID := envvar.GetTestProjectFromEnv()
	region := envvar.GetTestRegionFromEnv()

	config := BootstrapConfig(t)
	if config == nil {
		t.Fatal("Could not bootstrap config.")
	}

	computeService := config.NewComputeClient(config.UserAgent)
	if computeService == nil {
		t.Fatal("Could not create compute client.")
	}

	// In order to create a networkAttachment we need to bootstrap a subnet.
	_, err := computeService.Subnetworks.Get(projectID, region, subnetName).Do()
	if err != nil && transport_tpg.IsGoogleApiErrorWithCode(err, 404) {
		log.Printf("[DEBUG] Subnet %q not found, bootstrapping", subnetName)

		networkUrl := fmt.Sprintf("%sprojects/%s/global/networks/%s", config.ComputeBasePath, projectID, networkName)
		url := fmt.Sprintf("%sprojects/%s/regions/%s/subnetworks", config.ComputeBasePath, projectID, region)

		defaultSubnetObj := map[string]interface{}{
			"name":        subnetName,
			"region ":     region,
			"network":     networkUrl,
			"ipCidrRange": "10.77.0.0/20",
		}

		if len(subnetOptions) != 0 {
			maps.Copy(defaultSubnetObj, subnetOptions)
		}

		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "POST",
			Project:   projectID,
			RawURL:    url,
			UserAgent: config.UserAgent,
			Body:      defaultSubnetObj,
			Timeout:   4 * time.Minute,
		})

		log.Printf("Response is, %s", res)
		if err != nil {
			t.Fatalf("Error bootstrapping test subnet %s: %s", subnetName, err)
		}

		log.Printf("[DEBUG] Waiting for network creation to finish")
		err = tpgcompute.ComputeOperationWaitTime(config, res, projectID, "Error bootstrapping test subnet", config.UserAgent, 4*time.Minute)
		if err != nil {
			t.Fatalf("Error bootstrapping test subnet %s: %s", subnetName, err)
		}
	}

	subnet, err := computeService.Subnetworks.Get(projectID, region, subnetName).Do()

	if subnet == nil {
		t.Fatalf("Error getting test subnet %s: is nil", subnetName)
	}

	if err != nil {
		t.Fatalf("Error getting test subnet %s: %s", subnetName, err)
	}
	return subnet.Name
}

func BootstrapNetworkAttachment(t *testing.T, networkAttachmentName string, subnetName string) string {
	projectID := envvar.GetTestProjectFromEnv()
	region := envvar.GetTestRegionFromEnv()

	config := BootstrapConfig(t)
	if config == nil {
		return ""
	}

	computeService := config.NewComputeClient(config.UserAgent)
	if computeService == nil {
		return ""
	}

	networkAttachment, err := computeService.NetworkAttachments.Get(projectID, region, networkAttachmentName).Do()
	if err != nil && transport_tpg.IsGoogleApiErrorWithCode(err, 404) {
		// Create Network Attachment Here.
		log.Printf("[DEBUG] Network Attachment %s not found, bootstrapping", networkAttachmentName)
		url := fmt.Sprintf("%sprojects/%s/regions/%s/networkAttachments", config.ComputeBasePath, projectID, region)

		subnetURL := fmt.Sprintf("%sprojects/%s/regions/%s/subnetworks/%s", config.ComputeBasePath, projectID, region, subnetName)
		networkAttachmentObj := map[string]interface{}{
			"name":                 networkAttachmentName,
			"region":               region,
			"subnetworks":          []string{subnetURL},
			"connectionPreference": "ACCEPT_AUTOMATIC",
		}

		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "POST",
			Project:   projectID,
			RawURL:    url,
			UserAgent: config.UserAgent,
			Body:      networkAttachmentObj,
			Timeout:   4 * time.Minute,
		})
		if err != nil {
			t.Fatalf("Error bootstrapping test Network Attachment %s: %s", networkAttachmentName, err)
		}

		log.Printf("[DEBUG] Waiting for network creation to finish")
		err = tpgcompute.ComputeOperationWaitTime(config, res, projectID, "Error bootstrapping shared test subnet", config.UserAgent, 4*time.Minute)
		if err != nil {
			t.Fatalf("Error bootstrapping test Network Attachment %s: %s", networkAttachmentName, err)
		}
	}

	networkAttachment, err = computeService.NetworkAttachments.Get(projectID, region, networkAttachmentName).Do()

	if networkAttachment == nil {
		t.Fatalf("Error getting test network attachment %s: is nil", networkAttachmentName)
	}

	if err != nil {
		t.Fatalf("Error getting test Network Attachment %s: %s", networkAttachmentName, err)
	}

	return networkAttachment.Name
}

// The default network within GCP already comes pre configured with
// certain firewall rules open to allow internal communication. As we
// are boostrapping a network for dataproc tests, we need to additionally
// open up similar rules to allow the nodes to talk to each other
// internally as part of their configuration or this will just hang.
const SharedTestFirewallPrefix = "tf-bootstrap-firewall-"

func BootstrapFirewallForDataprocSharedNetwork(t *testing.T, firewallName string, networkName string) string {
	project := envvar.GetTestProjectFromEnv()
	firewallName = SharedTestFirewallPrefix + firewallName

	config := BootstrapConfig(t)
	if config == nil {
		return ""
	}

	log.Printf("[DEBUG] Getting Firewall %q for Network %q", firewallName, networkName)
	_, err := config.NewComputeClient(config.UserAgent).Firewalls.Get(project, firewallName).Do()
	if err != nil && transport_tpg.IsGoogleApiErrorWithCode(err, 404) {
		log.Printf("[DEBUG] firewallName %q not found, bootstrapping", firewallName)
		url := fmt.Sprintf("%sprojects/%s/global/firewalls", config.ComputeBasePath, project)

		networkId := fmt.Sprintf("projects/%s/global/networks/%s", project, networkName)
		allowObj := []interface{}{
			map[string]interface{}{
				"IPProtocol": "icmp",
			},
			map[string]interface{}{
				"IPProtocol": "tcp",
				"ports":      []string{"0-65535"},
			},
			map[string]interface{}{
				"IPProtocol": "udp",
				"ports":      []string{"0-65535"},
			},
		}

		firewallObj := map[string]interface{}{
			"name":    firewallName,
			"network": networkId,
			"allowed": allowObj,
		}

		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "POST",
			Project:   project,
			RawURL:    url,
			UserAgent: config.UserAgent,
			Body:      firewallObj,
			Timeout:   4 * time.Minute,
		})
		if err != nil {
			t.Fatalf("Error bootstrapping Firewall %q for Network %q: %s", firewallName, networkName, err)
		}

		log.Printf("[DEBUG] Waiting for Firewall creation to finish")
		err = tpgcompute.ComputeOperationWaitTime(config, res, project, "Error bootstrapping Firewall", config.UserAgent, 4*time.Minute)
		if err != nil {
			t.Fatalf("Error bootstrapping Firewall %q: %s", firewallName, err)
		}
	}

	firewall, err := config.NewComputeClient(config.UserAgent).Firewalls.Get(project, firewallName).Do()
	if err != nil {
		t.Errorf("Error getting Firewall %q: %s", firewallName, err)
	}
	if firewall == nil {
		t.Fatalf("Error getting Firewall %q: is nil", firewallName)
	}
	return firewall.Name
}

const SharedStoragePoolPrefix = "tf-bootstrap-storage-pool-"

func BootstrapComputeStoragePool(t *testing.T, storagePoolName, storagePoolType string) string {
	projectID := envvar.GetTestProjectFromEnv()
	zone := envvar.GetTestZoneFromEnv()

	storagePoolName = SharedStoragePoolPrefix + storagePoolType + "-" + storagePoolName

	config := BootstrapConfig(t)
	if config == nil {
		t.Fatal("Could not bootstrap config.")
	}

	computeService := config.NewComputeClient(config.UserAgent)
	if computeService == nil {
		t.Fatal("Could not create compute client.")
	}

	_, err := computeService.StoragePools.Get(projectID, zone, storagePoolName).Do()
	if err != nil && transport_tpg.IsGoogleApiErrorWithCode(err, 404) {
		log.Printf("[DEBUG] Storage pool %q not found, bootstrapping", storagePoolName)

		url := fmt.Sprintf("%sprojects/%s/zones/%s/storagePools", config.ComputeBasePath, projectID, zone)
		storagePoolTypeUrl := fmt.Sprintf("/projects/%s/zones/%s/storagePoolTypes/%s", projectID, zone, storagePoolType)

		storagePoolObj := map[string]interface{}{
			"name":                      storagePoolName,
			"poolProvisionedCapacityGb": 10240,
			"poolProvisionedThroughput": 180,
			"storagePoolType":           storagePoolTypeUrl,
			"capacityProvisioningType":  "ADVANCED",
		}

		if storagePoolType == "hyperdisk-balanced" {
			storagePoolObj["poolProvisionedIops"] = 10000
			storagePoolObj["poolProvisionedThroughput"] = 1024
		}

		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "POST",
			Project:   projectID,
			RawURL:    url,
			UserAgent: config.UserAgent,
			Body:      storagePoolObj,
			Timeout:   20 * time.Minute,
		})

		log.Printf("Response is, %s", res)
		if err != nil {
			t.Fatalf("Error bootstrapping storage pool %s: %s", storagePoolName, err)
		}

		log.Printf("[DEBUG] Waiting for storage pool creation to finish")
		err = tpgcompute.ComputeOperationWaitTime(config, res, projectID, "Error bootstrapping storage pool", config.UserAgent, 4*time.Minute)
		if err != nil {
			t.Fatalf("Error bootstrapping test storage pool %s: %s", storagePoolName, err)
		}
	}

	storagePool, err := computeService.StoragePools.Get(projectID, zone, storagePoolName).Do()

	if storagePool == nil {
		t.Fatalf("Error getting storage pool %s: is nil", storagePoolName)
	}

	if err != nil {
		t.Fatalf("Error getting storage pool %s: %s", storagePoolName, err)
	}

	storagePoolResourceName, err := tpgresource.GetRelativePath(storagePool.SelfLink)

	if err != nil {
		t.Fatal("Failed to extract Storage Pool resource name from URL.")
	}

	return storagePoolResourceName
}

func SetupProjectsAndGetAccessToken(org, billing, pid, service string, config *transport_tpg.Config) (string, error) {
	// Create project-1 and project-2
	rmService := config.NewResourceManagerClient(config.UserAgent)

	project := &cloudresourcemanager.Project{
		ProjectId: pid,
		Name:      pid,
		Parent: &cloudresourcemanager.ResourceId{
			Id:   org,
			Type: "organization",
		},
	}

	var op *cloudresourcemanager.Operation
	err := transport_tpg.Retry(transport_tpg.RetryOptions{
		RetryFunc: func() (reqErr error) {
			op, reqErr = rmService.Projects.Create(project).Do()
			return reqErr
		},
		Timeout: 5 * time.Minute,
	})
	if err != nil {
		return "", fmt.Errorf("error creating 'project-1' with project id %s: %w", pid, err)
	}

	// Wait for the operation to complete
	opAsMap, err := tpgresource.ConvertToMap(op)
	if err != nil {
		return "", fmt.Errorf("error in ConvertToMap while creating 'project-1' with project id %s: %w", pid, err)
	}

	waitErr := resourcemanager.ResourceManagerOperationWaitTime(config, opAsMap, "creating project", config.UserAgent, 5*time.Minute)
	if waitErr != nil {
		return "", waitErr
	}

	ba := &cloudbilling.ProjectBillingInfo{
		BillingAccountName: fmt.Sprintf("billingAccounts/%s", billing),
	}
	_, err = config.NewBillingClient(config.UserAgent).Projects.UpdateBillingInfo(resourcemanager.PrefixedProject(pid), ba).Do()
	if err != nil {
		return "", fmt.Errorf("error updating billing info for 'project-1' with project id %s: %w", pid, err)
	}

	p2 := fmt.Sprintf("%s-2", pid)
	project.ProjectId = p2
	project.Name = fmt.Sprintf("%s-2", pid)

	err = transport_tpg.Retry(transport_tpg.RetryOptions{
		RetryFunc: func() (reqErr error) {
			op, reqErr = rmService.Projects.Create(project).Do()
			return reqErr
		},
		Timeout: 5 * time.Minute,
	})
	if err != nil {
		return "", fmt.Errorf("error creating 'project-2' with project id %s: %w", p2, err)
	}

	// Wait for the operation to complete
	opAsMap, err = tpgresource.ConvertToMap(op)
	if err != nil {
		return "", err
	}

	waitErr = resourcemanager.ResourceManagerOperationWaitTime(config, opAsMap, "creating project", config.UserAgent, 5*time.Minute)
	if waitErr != nil {
		return "", waitErr
	}

	_, err = config.NewBillingClient(config.UserAgent).Projects.UpdateBillingInfo(resourcemanager.PrefixedProject(p2), ba).Do()
	if err != nil {
		return "", fmt.Errorf("error updating billing info for 'project-2' with project id %s: %w", p2, err)
	}

	// Enable the appropriate service in project-2 only
	suService := config.NewServiceUsageClient(config.UserAgent)

	serviceReq := &serviceusage.BatchEnableServicesRequest{
		ServiceIds: []string{fmt.Sprintf("%s.googleapis.com", service)},
	}

	_, err = suService.Services.BatchEnable(fmt.Sprintf("projects/%s", p2), serviceReq).Do()
	if err != nil {
		return "", fmt.Errorf("error batch enabling services in 'project-2' with project id %s: %w", p2, err)
	}

	// Enable the test runner to create service accounts and get an access token on behalf of
	// the project 1 service account
	curEmail, err := transport_tpg.GetCurrentUserEmail(config, config.UserAgent)
	if err != nil {
		return "", fmt.Errorf("error getting current user email: %w", err)
	}

	proj1SATokenCreator := &cloudresourcemanager.Binding{
		Members: []string{fmt.Sprintf("serviceAccount:%s", curEmail)},
		Role:    "roles/iam.serviceAccountTokenCreator",
	}

	proj1SACreator := &cloudresourcemanager.Binding{
		Members: []string{fmt.Sprintf("serviceAccount:%s", curEmail)},
		Role:    "roles/iam.serviceAccountCreator",
	}

	bindings := tpgiamresource.MergeBindings([]*cloudresourcemanager.Binding{proj1SATokenCreator, proj1SACreator})

	p, err := rmService.Projects.GetIamPolicy(pid,
		&cloudresourcemanager.GetIamPolicyRequest{
			Options: &cloudresourcemanager.GetPolicyOptions{
				RequestedPolicyVersion: tpgiamresource.IamPolicyVersion,
			},
		}).Do()
	if err != nil {
		return "", fmt.Errorf("error getting IAM policy for 'project-1' with project id %s: %w", pid, err)
	}

	p.Bindings = tpgiamresource.MergeBindings(append(p.Bindings, bindings...))
	_, err = config.NewResourceManagerClient(config.UserAgent).Projects.SetIamPolicy(pid,
		&cloudresourcemanager.SetIamPolicyRequest{
			Policy:     p,
			UpdateMask: "bindings,etag,auditConfigs",
		}).Do()
	if err != nil {
		return "", fmt.Errorf("error setting IAM policy for 'project-1' with project id %s: %w", pid, err)
	}

	// Create a service account for project-1
	serviceAccountEmail := serviceAccountPrefix + service
	sa1, err := getOrCreateServiceAccount(config, pid, serviceAccountEmail)
	if err != nil {
		return "", fmt.Errorf("error creating service account %s in 'project-1' with project id %s: %w", serviceAccountEmail, pid, err)
	}
	// Setting IAM policies sometimes fails due to the service account not being created yet
	// Wait a minute to ensure we can use it.
	time.Sleep(1 * time.Minute)

	// Add permissions to service accounts

	// Permission needed for user_project_override
	proj2ServiceUsageBinding := &cloudresourcemanager.Binding{
		Members: []string{fmt.Sprintf("serviceAccount:%s", sa1.Email)},
		Role:    "roles/serviceusage.serviceUsageConsumer",
	}

	// Admin permission for service
	proj2ServiceAdminBinding := &cloudresourcemanager.Binding{
		Members: []string{fmt.Sprintf("serviceAccount:%s", sa1.Email)},
		Role:    fmt.Sprintf("roles/%s.admin", service),
	}

	bindings = tpgiamresource.MergeBindings([]*cloudresourcemanager.Binding{proj2ServiceUsageBinding, proj2ServiceAdminBinding})

	// For KMS test only
	if service == "cloudkms" {
		proj2CryptoKeyBinding := &cloudresourcemanager.Binding{
			Members: []string{fmt.Sprintf("serviceAccount:%s", sa1.Email)},
			Role:    "roles/cloudkms.cryptoKeyEncrypter",
		}

		bindings = tpgiamresource.MergeBindings(append(bindings, proj2CryptoKeyBinding))
	}

	// For Firebase test only
	if service == "firebase" {
		// Additional permissions besides roles/serviceusage.serviceUsageConsumer and roles/firebase.admin are needed
		// https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects/addFirebase
		proj2ServiceUsageBinding := &cloudresourcemanager.Binding{
			Members: []string{fmt.Sprintf("serviceAccount:%s", sa1.Email)},
			Role:    "roles/serviceusage.serviceUsageAdmin",
		}

		bindings = tpgiamresource.MergeBindings(append(bindings, proj2ServiceUsageBinding))
	}

	p, err = rmService.Projects.GetIamPolicy(p2,
		&cloudresourcemanager.GetIamPolicyRequest{
			Options: &cloudresourcemanager.GetPolicyOptions{
				RequestedPolicyVersion: tpgiamresource.IamPolicyVersion,
			},
		}).Do()
	if err != nil {
		return "", fmt.Errorf("error getting IAM policy for 'project-2' with project id %s: %w", p2, err)
	}

	p.Bindings = tpgiamresource.MergeBindings(append(p.Bindings, bindings...))
	_, err = config.NewResourceManagerClient(config.UserAgent).Projects.SetIamPolicy(p2,
		&cloudresourcemanager.SetIamPolicyRequest{
			Policy:     p,
			UpdateMask: "bindings,etag,auditConfigs",
		}).Do()
	if err != nil {
		return "", fmt.Errorf("error setting IAM policy for 'project-2' with project id %s: %w", p2, err)
	}

	// The token creator IAM API call returns success long before the policy is
	// actually usable. Wait a solid 2 minutes to ensure we can use it.
	time.Sleep(2 * time.Minute)

	iamCredsService := config.NewIamCredentialsClient(config.UserAgent)
	tokenRequest := &iamcredentials.GenerateAccessTokenRequest{
		Lifetime: "300s",
		Scope:    []string{"https://www.googleapis.com/auth/cloud-platform"},
	}
	atResp, err := iamCredsService.Projects.ServiceAccounts.GenerateAccessToken(fmt.Sprintf("projects/-/serviceAccounts/%s", sa1.Email), tokenRequest).Do()
	if err != nil {
		return "", fmt.Errorf("error generating access token for service account %s: %w", sa1.Email, err)
	}

	accessToken := atResp.AccessToken

	return accessToken, nil
}

// For bootstrapping Developer Connect git repository link
const SharedGitRepositoryLinkIdPrefix = "tf-bootstrap-git-repository-"

func BootstrapGitRepository(t *testing.T, gitRepositoryLinkId, location, cloneUri, parentConnectionId string) string {
	gitRepositoryLinkId = SharedGitRepositoryLinkIdPrefix + gitRepositoryLinkId

	config := BootstrapConfig(t)
	if config == nil {
		t.Fatal("Could not bootstrap config.")
	}

	log.Printf("[DEBUG] Getting shared git repository link %q", gitRepositoryLinkId)

	getURL := fmt.Sprintf("%sprojects/%s/locations/%s/connections/%s/gitRepositoryLinks/%s",
		config.DeveloperConnectBasePath, config.Project, location, parentConnectionId, gitRepositoryLinkId)

	headers := make(http.Header)
	_, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   config.Project,
		RawURL:    getURL,
		UserAgent: config.UserAgent,
		Headers:   headers,
	})

	if err != nil && transport_tpg.IsGoogleApiErrorWithCode(err, 404) {
		log.Printf("[DEBUG] Git repository link %q not found, bootstrapping", gitRepositoryLinkId)
		obj := map[string]interface{}{
			"clone_uri":   cloneUri,
			"annotations": map[string]string{},
		}

		postURL := fmt.Sprintf("%sprojects/%s/locations/%s/connections/%s/gitRepositoryLinks?gitRepositoryLinkId=%s",
			config.DeveloperConnectBasePath, config.Project, location, parentConnectionId, gitRepositoryLinkId)
		headers := make(http.Header)
		_, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "POST",
			Project:   config.Project,
			RawURL:    postURL,
			UserAgent: config.UserAgent,
			Body:      obj,
			Timeout:   20 * time.Minute,
			Headers:   headers,
		})
		if err != nil {
			t.Fatalf("Error bootstrapping git repository link %q: %s", gitRepositoryLinkId, err)
		}

		_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "GET",
			Project:   config.Project,
			RawURL:    getURL,
			UserAgent: config.UserAgent,
			Timeout:   20 * time.Minute,
			Headers:   headers,
		})
		if err != nil {
			t.Fatalf("Error getting git repository link %q: %s", gitRepositoryLinkId, err)
		}
	}

	return gitRepositoryLinkId
}

const SharedConnectionIdPrefix = "tf-bootstrap-developer-connect-connection-"

// For bootstrapping Developer Connect connection resources
func BootstrapDeveloperConnection(t *testing.T, connectionId, location, tokenResource string, appInstallationId int) string {
	connectionId = SharedConnectionIdPrefix + connectionId

	config := BootstrapConfig(t)
	if config == nil {
		t.Fatal("Could not bootstrap config.")
	}

	log.Printf("[DEBUG] Getting shared developer connection %q", connectionId)

	getURL := fmt.Sprintf("%sprojects/%s/locations/%s/connections/%s",
		config.DeveloperConnectBasePath, config.Project, location, connectionId)

	headers := make(http.Header)
	_, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   config.Project,
		RawURL:    getURL,
		UserAgent: config.UserAgent,
		Headers:   headers,
	})

	if err != nil {
		log.Printf("[DEBUG] Developer connection %q not found, bootstrapping", connectionId)
		authorizerCredential := map[string]string{
			"oauth_token_secret_version": tokenResource,
		}
		githubConfig := map[string]interface{}{
			"github_app":            "DEVELOPER_CONNECT",
			"app_installation_id":   appInstallationId,
			"authorizer_credential": authorizerCredential,
		}
		obj := map[string]interface{}{
			"disabled":      false,
			"github_config": githubConfig,
		}

		postURL := fmt.Sprintf("%sprojects/%s/locations/%s/connections?connectionId=%s",
			config.DeveloperConnectBasePath, config.Project, location, connectionId)
		headers := make(http.Header)
		_, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "POST",
			Project:   config.Project,
			RawURL:    postURL,
			UserAgent: config.UserAgent,
			Body:      obj,
			Timeout:   20 * time.Minute,
			Headers:   headers,
		})
		if err != nil {
			t.Fatalf("Error bootstrapping developer connection %q: %s", connectionId, err)
		}

		_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "GET",
			Project:   config.Project,
			RawURL:    getURL,
			UserAgent: config.UserAgent,
			Timeout:   20 * time.Minute,
			Headers:   headers,
		})
		if err != nil {
			t.Fatalf("Error getting developer connection %q: %s", connectionId, err)
		}
	}

	return connectionId
}

const SharedRepositoryGroupPrefix = "tf-bootstrap-repo-group-"

func BoostrapSharedRepositoryGroup(t *testing.T, repositoryGroupId, location, labels, codeRepositoryIndexId, resource string) string {
	repositoryGroupId = SharedRepositoryGroupPrefix + repositoryGroupId

	config := BootstrapConfig(t)
	if config == nil {
		t.Fatal("Could not bootstrap config.")
	}

	log.Printf("[DEBUG] Getting shared repository group %q", repositoryGroupId)

	getURL := fmt.Sprintf("%sprojects/%s/locations/%s/codeRepositoryIndexes/%s/repositoryGroups/%s",
		config.GeminiBasePath, config.Project, location, codeRepositoryIndexId, repositoryGroupId)

	headers := make(http.Header)
	_, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   config.Project,
		RawURL:    getURL,
		UserAgent: config.UserAgent,
		Headers:   headers,
	})
	if err != nil {
		log.Printf("[DEBUG] Repository group %q not found, bootstrapping", codeRepositoryIndexId)
		repositories := [1]interface{}{map[string]string{
			"resource":       resource,
			"branch_pattern": "main",
		}}
		postURL := fmt.Sprintf("%sprojects/%s/locations/%s/codeRepositoryIndexes/%s/repositoryGroups?repositoryGroupId=%s",
			config.GeminiBasePath, config.Project, location, codeRepositoryIndexId, repositoryGroupId)
		obj := map[string]interface{}{
			"repositories": repositories,
		}
		if labels != "" {
			obj["labels"] = labels
		}

		headers := make(http.Header)
		for {
			_, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "POST",
				Project:   config.Project,
				RawURL:    postURL,
				UserAgent: config.UserAgent,
				Body:      obj,
				Timeout:   20 * time.Minute,
				Headers:   headers,
			})
			if err != nil {
				if transport_tpg.IsGoogleApiErrorWithCode(err, 409) {
					errMsg := fmt.Sprintf("%s", err)
					if strings.Contains(errMsg, "unable to queue the operation") {
						log.Printf("[DEBUG] Waiting for enqueued operation to finish before creating RepositoryGroup: %#v", obj)
						time.Sleep(10 * time.Second)
					} else if strings.Contains(errMsg, "parent resource not in ready state") {
						log.Printf("[DEBUG] Waiting for parent resource to become active before creating RepositoryGroup: %#v", obj)
						time.Sleep(1 * time.Minute)
					} else {
						t.Fatalf("Error creating RepositoryGroup: %s", err)
					}
				} else {
					t.Fatalf("Error creating repository group %q: %s", repositoryGroupId, err)
				}
			} else {
				break
			}
		}

		_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "GET",
			Project:   config.Project,
			RawURL:    getURL,
			UserAgent: config.UserAgent,
			Timeout:   20 * time.Minute,
			Headers:   headers,
		})
		if err != nil {
			t.Errorf("Error getting repository group %q: %s", repositoryGroupId, err)
		}
	}

	return repositoryGroupId
}

// BootstrapSharedCodeRepositoryIndex will create a code repository index
// if it hasn't been created in the test project.
//
// BootstrapSharedCodeRepositoryIndex returns a persistent code repository index
// for a test or set of tests.
//
// Deletion of code repository index takes a few minutes, and creation of it
// currently takes about half an hour.
// That is the reason to use the shared code repository indexes for test resources.
const SharedCodeRepositoryIndexPrefix = "tf-bootstrap-cri-"

func BootstrapSharedCodeRepositoryIndex(t *testing.T, codeRepositoryIndexId, location, kmsKey string, labels map[string]string) string {
	codeRepositoryIndexId = SharedCodeRepositoryIndexPrefix + codeRepositoryIndexId

	config := BootstrapConfig(t)
	if config == nil {
		t.Fatal("Could not bootstrap config.")
	}

	log.Printf("[DEBUG] Getting shared code repository index %q", codeRepositoryIndexId)

	getURL := fmt.Sprintf("%sprojects/%s/locations/%s/codeRepositoryIndexes/%s", config.GeminiBasePath, config.Project, location, codeRepositoryIndexId)

	headers := make(http.Header)
	_, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   config.Project,
		RawURL:    getURL,
		UserAgent: config.UserAgent,
		Timeout:   90 * time.Minute,
		Headers:   headers,
	})

	// CRI not found responds with 404 not found
	if err != nil && transport_tpg.IsGoogleApiErrorWithCode(err, 404) {
		log.Printf("[DEBUG] Code repository index %q not found, bootstrapping", codeRepositoryIndexId)
		postURL := fmt.Sprintf("%sprojects/%s/locations/%s/codeRepositoryIndexes?codeRepositoryIndexId=%s", config.GeminiBasePath, config.Project, location, codeRepositoryIndexId)
		obj := make(map[string]interface{})
		if labels != nil {
			obj["labels"] = labels
		}
		if kmsKey != "" {
			obj["kmsKey"] = kmsKey
		}

		headers := make(http.Header)
		_, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "POST",
			Project:   config.Project,
			RawURL:    postURL,
			UserAgent: config.UserAgent,
			Body:      obj,
			Timeout:   90 * time.Minute,
			Headers:   headers,
		})
		if err != nil {
			t.Fatalf("Error creating code repository index %q: %s", codeRepositoryIndexId, err)
		}

		_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "GET",
			Project:   config.Project,
			RawURL:    getURL,
			UserAgent: config.UserAgent,
			Timeout:   90 * time.Minute,
			Headers:   headers,
		})
		if err != nil {
			t.Fatalf("Error getting code repository index %q: %s", codeRepositoryIndexId, err)
		}
	} else if err != nil {
		t.Fatalf("Error getting code repository index %q: %s", codeRepositoryIndexId, err)
	}

	return codeRepositoryIndexId
}

const sharedTagKeyPrefix = "tf-bootstrap-tagkey"

func BootstrapSharedTestTagKey(t *testing.T, testId string) string {
	org := envvar.GetTestOrgFromEnv(t)
	sharedTagKey := fmt.Sprintf("%s-%s", sharedTagKeyPrefix, testId)
	tagKeyName := fmt.Sprintf("%s/%s", org, sharedTagKey)

	config := BootstrapConfig(t)
	if config == nil {
		return ""
	}

	log.Printf("[DEBUG] Getting shared test tag key %q", sharedTagKey)
	getURL := fmt.Sprintf("%stagKeys/namespaced?name=%s", config.TagsBasePath, tagKeyName)
	_, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   config.Project,
		RawURL:    getURL,
		UserAgent: config.UserAgent,
		Timeout:   2 * time.Minute,
	})
	if err != nil && transport_tpg.IsGoogleApiErrorWithCode(err, 403) {
		log.Printf("[DEBUG] TagKey %q not found, bootstrapping", sharedTagKey)
		tagKeyObj := map[string]interface{}{
			"parent":      "organizations/" + org,
			"shortName":   sharedTagKey,
			"description": "Bootstrapped tag key for Terraform Acceptance testing",
		}

		_, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "POST",
			Project:   config.Project,
			RawURL:    config.TagsBasePath + "tagKeys/",
			UserAgent: config.UserAgent,
			Body:      tagKeyObj,
			Timeout:   10 * time.Minute,
		})
		if err != nil {
			t.Fatalf("Error bootstrapping shared tag key %q: %s", sharedTagKey, err)
		}

		log.Printf("[DEBUG] Waiting for shared tag key creation to finish")
	}

	_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   config.Project,
		RawURL:    getURL,
		UserAgent: config.UserAgent,
		Timeout:   2 * time.Minute,
	})

	if err != nil {
		t.Fatalf("Error getting shared tag key %q: %s", sharedTagKey, err)
	}

	return sharedTagKey
}

const sharedTagValuePrefix = "tf-bootstrap-tagvalue"

func BootstrapSharedTestTagValue(t *testing.T, testId string, tagKey string) string {
	org := envvar.GetTestOrgFromEnv(t)
	sharedTagValue := fmt.Sprintf("%s-%s", sharedTagValuePrefix, testId)
	tagKeyName := fmt.Sprintf("%s/%s", org, tagKey)
	tagValueName := fmt.Sprintf("%s/%s", tagKeyName, sharedTagValue)

	config := BootstrapConfig(t)
	if config == nil {
		return ""
	}

	log.Printf("[DEBUG] Getting shared test tag value %q", sharedTagValue)
	getURL := fmt.Sprintf("%stagValues/namespaced?name=%s", config.TagsBasePath, tagValueName)
	_, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   config.Project,
		RawURL:    getURL,
		UserAgent: config.UserAgent,
		Timeout:   2 * time.Minute,
	})
	if err != nil && transport_tpg.IsGoogleApiErrorWithCode(err, 403) {
		log.Printf("[DEBUG] TagValue %q not found, bootstrapping", sharedTagValue)
		log.Printf("[DEBUG] Fetching permanent id for tagkey %s", tagKeyName)
		tagKeyGetURL := fmt.Sprintf("%stagKeys/namespaced?name=%s", config.TagsBasePath, tagKeyName)
		tagKeyResponse, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "GET",
			Project:   config.Project,
			RawURL:    tagKeyGetURL,
			UserAgent: config.UserAgent,
			Timeout:   2 * time.Minute,
		})
		if err != nil {
			t.Fatalf("Error getting tag key id for %s : %s", tagKeyName, err)
		}
		tagKeyObj := map[string]interface{}{
			"parent":      tagKeyResponse["name"].(string),
			"shortName":   sharedTagValue,
			"description": "Bootstrapped tag value for Terraform Acceptance testing",
		}

		_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "POST",
			Project:   config.Project,
			RawURL:    config.TagsBasePath + "tagValues/",
			UserAgent: config.UserAgent,
			Body:      tagKeyObj,
			Timeout:   10 * time.Minute,
		})
		if err != nil {
			t.Fatalf("Error bootstrapping shared tag value %q: %s", sharedTagValue, err)
		}

		log.Printf("[DEBUG] Waiting for shared tag value creation to finish")
	}

	_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   config.Project,
		RawURL:    getURL,
		UserAgent: config.UserAgent,
		Timeout:   2 * time.Minute,
	})

	if err != nil {
		t.Fatalf("Error getting shared tag value %q: %s", sharedTagValue, err)
	}

	return sharedTagValue
}
