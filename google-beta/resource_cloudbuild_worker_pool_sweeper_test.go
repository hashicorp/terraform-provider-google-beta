package google

import (
	"context"
	"log"
	"testing"

	cloudbuild "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbuild/beta"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func init() {
	resource.AddTestSweepers("CloudBuildWorker_pool", &resource.Sweeper{
		Name: "CloudBuildWorker_pool",
		F:    testSweepCloudBuildWorker_pool,
	})
}

func testSweepCloudBuildWorker_pool(region string) error {
	resourceName := "CloudBuildWorker_pool"
	log.Printf("[INFO][SWEEPER_LOG] Starting sweeper for %s", resourceName)

	config, err := sharedConfigForRegion(region)
	if err != nil {
		log.Printf("[INFO][SWEEPER_LOG] error getting shared config for region: %s", err)
		return err
	}

	err = config.LoadAndValidate(context.Background())
	if err != nil {
		log.Printf("[INFO][SWEEPER_LOG] error loading: %s", err)
		return err
	}

	t := &testing.T{}
	billingId := getTestBillingAccountFromEnv(t)

	// Setup variables to be used for Delete arguments.
	d := map[string]string{
		"project":         config.Project,
		"region":          region,
		"location":        region,
		"zone":            "-",
		"billing_account": billingId,
	}

	client := NewDCLCloudBuildClient(config, config.userAgent, "")
	err = client.DeleteAllWorkerPool(context.Background(), d["project"], d["location"], isDeletableCloudBuildWorker_pool)
	if err != nil {
		return err
	}
	return nil
}

func isDeletableCloudBuildWorker_pool(r *cloudbuild.WorkerPool) bool {
	return isSweepableTestResource(*r.Name)
}
