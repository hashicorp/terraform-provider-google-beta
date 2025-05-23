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
//	Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/services/kms/data_source_google_kms_secret.go
//
//	DO NOT EDIT this file directly. Any changes made to this file will be
//	overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
package kms

import (
	"google.golang.org/api/cloudkms/v1"

	"encoding/base64"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceGoogleKmsSecret() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGoogleKmsSecretRead,
		Schema: map[string]*schema.Schema{
			"crypto_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ciphertext": {
				Type:     schema.TypeString,
				Required: true,
			},
			"plaintext": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"additional_authenticated_data": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceGoogleKmsSecretRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	cryptoKeyId, err := ParseKmsCryptoKeyId(d.Get("crypto_key").(string), config)

	if err != nil {
		return err
	}

	ciphertext := d.Get("ciphertext").(string)

	kmsDecryptRequest := &cloudkms.DecryptRequest{
		Ciphertext: ciphertext,
	}

	if aad, ok := d.GetOk("additional_authenticated_data"); ok {
		kmsDecryptRequest.AdditionalAuthenticatedData = aad.(string)
	}

	decryptResponse, err := config.NewKmsClient(userAgent).Projects.Locations.KeyRings.CryptoKeys.Decrypt(cryptoKeyId.CryptoKeyId(), kmsDecryptRequest).Do()

	if err != nil {
		return fmt.Errorf("Error decrypting ciphertext: %s", err)
	}

	plaintext, err := base64.StdEncoding.DecodeString(decryptResponse.Plaintext)

	if err != nil {
		return fmt.Errorf("Error decoding base64 response: %s", err)
	}

	log.Printf("[INFO] Successfully decrypted ciphertext: %s", ciphertext)

	if err := d.Set("plaintext", string(plaintext[:])); err != nil {
		return fmt.Errorf("Error setting plaintext: %s", err)
	}
	d.SetId(fmt.Sprintf("%s:%s", d.Get("crypto_key").(string), ciphertext))

	return nil
}
