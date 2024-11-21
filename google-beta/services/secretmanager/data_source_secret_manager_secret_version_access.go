// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package secretmanager

import (
	"encoding/base64"
	"fmt"
	"log"
	"regexp"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceSecretManagerSecretVersionAccess() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSecretManagerSecretVersionAccessRead,
		Schema: map[string]*schema.Schema{
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secret": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"secret_data": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"is_secret_data_base64": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func dataSourceSecretManagerSecretVersionAccessRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	dSecret, ok := d.Get("secret").(string)
	if !ok {
		return fmt.Errorf("wrong type for secret field (%T), expected string", d.Get("secret"))
	}

	fv, err := tpgresource.ParseProjectFieldValue("secrets", dSecret, "project", d, config, false)
	if err != nil {
		return err
	}

	project := fv.Project
	if dProject, ok := d.Get("project").(string); !ok {
		return fmt.Errorf("wrong type for project (%T), expected string", d.Get("project"))
	} else if dProject != "" && dProject != project {
		return fmt.Errorf("project field value (%s) does not match project of secret (%s).", dProject, project)
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("error setting project: %s", err)
	}
	if err := d.Set("secret", fv.Name); err != nil {
		return fmt.Errorf("error setting secret: %s", err)
	}

	var url string
	versionNum := d.Get("version")

	if versionNum != "" {
		url, err = tpgresource.ReplaceVars(d, config, "{{SecretManagerBasePath}}projects/{{project}}/secrets/{{secret}}/versions/{{version}}")
		if err != nil {
			return err
		}
	} else {
		url, err = tpgresource.ReplaceVars(d, config, "{{SecretManagerBasePath}}projects/{{project}}/secrets/{{secret}}/versions/latest")
		if err != nil {
			return err
		}
	}

	url = fmt.Sprintf("%s:access", url)
	resp, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   project,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return fmt.Errorf("error retrieving available secret manager secret version access: %s", err.Error())
	}

	nameValue, ok := resp["name"]
	if !ok {
		return fmt.Errorf("read response didn't contain critical fields. Read may not have succeeded.")
	}
	if err := d.Set("name", nameValue.(string)); err != nil {
		return fmt.Errorf("error setting name: %s", err)
	}

	secretVersionRegex := regexp.MustCompile("projects/(.+)/secrets/(.+)/versions/(.+)$")

	parts := secretVersionRegex.FindStringSubmatch(nameValue.(string))
	// should return [full string, project number, secret name, version number]
	if len(parts) != 4 {
		return fmt.Errorf("secret name, %s, does not match format, projects/{{project}}/secrets/{{secret}}/versions/{{version}}", nameValue.(string))
	}

	log.Printf("[DEBUG] Received Google SecretManager Version: %q", parts[3])

	if err := d.Set("version", parts[3]); err != nil {
		return fmt.Errorf("error setting version: %s", err)
	}

	data := resp["payload"].(map[string]interface{})
	var secretData string
	if d.Get("is_secret_data_base64").(bool) {
		secretData = data["data"].(string)
	} else {
		payloadData, err := base64.StdEncoding.DecodeString(data["data"].(string))
		if err != nil {
			return fmt.Errorf("error decoding secret manager secret version data: %s", err.Error())
		}
		secretData = string(payloadData)
	}
	if err := d.Set("secret_data", secretData); err != nil {
		return fmt.Errorf("error setting secret_data: %s", err)
	}

	d.SetId(nameValue.(string))
	return nil
}
