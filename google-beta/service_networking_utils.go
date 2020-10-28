package google

import (
	"fmt"
	"log"
	"regexp"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"
)

func retrieveProjectById(config *Config, pid, userAgent string) (*cloudresourcemanager.Project, error) {
	log.Printf("[DEBUG] Retrieving project number by doing a GET with the project id, as required by service networking")
	project, err := config.NewResourceManagerClient(userAgent).Projects.Get(pid).Do()
	if err != nil {
		// note: returning a wrapped error is part of this method's contract!
		// https://blog.golang.org/go1.13-errors
		return nil, fmt.Errorf("Failed to retrieve project, pid: %s, err: %s", pid, err)
	}

	return project, nil
}

// NOTE(craigatgoogle): An out of band aspect of this API is that it uses a unique formatting of network
// different from the standard self_link URI. It requires a call to the resource manager to get the project
// number for the current project.
func retrieveServiceNetworkingNetworkName(d *schema.ResourceData, config *Config, network, userAgent string) (string, error) {
	networkFieldValue, err := ParseNetworkFieldValue(network, d, config)
	if err != nil {
		return "", errwrap.Wrapf("Failed to retrieve network field value, err: {{err}}", err)
	}

	pid := networkFieldValue.Project
	if pid == "" {
		return "", fmt.Errorf("Could not determine project")
	}

	project, err := retrieveProjectById(config, pid, userAgent)
	if err != nil {
		return "", err
	}

	networkName := networkFieldValue.Name
	if networkName == "" {
		return "", fmt.Errorf("Failed to parse network")
	}

	// return the network name formatting unique to this API
	return fmt.Sprintf("projects/%v/global/networks/%v", project.ProjectNumber, networkName), nil

}

func retrieveServiceNetworkingNetworkProject(d *schema.ResourceData, config *Config, network, userAgent string) (*cloudresourcemanager.Project, error) {
	networkFieldValue, err := ParseNetworkFieldValue(network, d, config)
	if err != nil {
		return nil, errwrap.Wrapf("Failed to retrieve network field value, err: {{err}}", err)
	}

	pid := networkFieldValue.Project
	if pid == "" {
		return nil, fmt.Errorf("Could not determine project")
	}

	project, err := retrieveProjectById(config, pid, userAgent)
	if err != nil {
		return nil, err
	}

	return project, nil
}

const parentServicePattern = "^services/.+$"

// NOTE(craigatgoogle): An out of band aspect of this API is that it requires the service name to be
// formatted as "services/<serviceName>"
func formatParentService(service string) string {
	r := regexp.MustCompile(parentServicePattern)
	if !r.MatchString(service) {
		return fmt.Sprintf("services/%s", service)
	} else {
		return service
	}
}
