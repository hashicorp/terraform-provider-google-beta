// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"

	compute "google.golang.org/api/compute/v0.beta"
)

func ResourceComputeInstanceFromMachineImage() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeInstanceFromMachineImageCreate,
		Read:   resourceComputeInstanceRead,
		Update: resourceComputeInstanceUpdate,
		Delete: resourceComputeInstanceDelete,

		// Import doesn't really make sense, because you could just import
		// as a google_compute_instance.

		Timeouts: ResourceComputeInstance().Timeouts,

		Schema:        computeInstanceFromMachineImageSchema(),
		CustomizeDiff: ResourceComputeInstance().CustomizeDiff,
		UseJSONNumber: true,
	}
}

func computeInstanceFromMachineImageSchema() map[string]*schema.Schema {
	s := ResourceComputeInstance().Schema

	for _, field := range []string{"boot_disk", "machine_type", "network_interface"} {
		// The user can set these fields as an override, but doesn't need to -
		// the machine image values will be used if they're unset.
		s[field].Required = false
		s[field].Optional = true
	}

	recurseOnSchema(s, func(field *schema.Schema) {
		// We don't want to accidentally use default values to override the instance
		// machine image, so remove defaults.
		field.Default = nil

		// Make non-required fields computed since they'll be set by the template.
		// Leave deprecated and removed fields alone because we don't set them.
		if !field.Required && !(field.Deprecated != "") {
			field.Computed = true
		}
	})

	s["source_machine_image"] = &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		ForceNew:    true,
		Description: `Name or self link of a machine image to create the instance from on.`,
	}

	s["source_machine_image_encryption_key"] = &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: `Encryption key for the source machine image.`,
		ForceNew:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"raw_key": {
					Type:      schema.TypeString,
					Optional:  true,
					Sensitive: true,
				},
				"rsa_encrypted_key": {
					Type:      schema.TypeString,
					Optional:  true,
					Sensitive: true,
					ForceNew:  true,
				},
				"kms_key_name": {
					Type:     schema.TypeString,
					Optional: true,
					ForceNew: true,
				},
				"kms_key_service_account": {
					Type:     schema.TypeString,
					Optional: true,
					ForceNew: true,
				},
				"sha256": {
					Type:     schema.TypeString,
					Computed: true,
				},
			},
		},
	}

	// Modifying the schema to disable disk overrides, due to an API bug (b/170964971)
	// TODO: (camthornton) Remove this when disk override functionality in the API is restored
	for _, field := range []string{"boot_disk", "attached_disk", "scratch_disk"} {
		s[field].Required = false
		s[field].Optional = false
		s[field].Computed = true
		s[field].MaxItems = 0
	}
	bootDiskSchema := s["boot_disk"].Elem.(*schema.Resource).Schema
	for _, field := range bootDiskSchema {
		field.AtLeastOneOf = []string{}
		field.ConflictsWith = []string{}
	}
	initializeParamsSchema := bootDiskSchema["initialize_params"].Elem.(*schema.Resource).Schema
	for _, field := range initializeParamsSchema {
		field.AtLeastOneOf = []string{}
	}
	// End disk schema modifications

	return s
}

func resourceComputeInstanceFromMachineImageCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return err
	}

	// Get the zone
	z, err := tpgresource.GetZone(d, config)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] Loading zone: %s", z)
	zone, err := config.NewComputeClient(userAgent).Zones.Get(project, z).Do()
	if err != nil {
		return fmt.Errorf("Error loading zone '%s': %s", z, err)
	}

	instance, err := expandComputeInstance(project, d, config)
	if err != nil {
		return err
	}

	sa := d.Get("service_account").([]interface{})
	if len(sa) == 0 {
		// ServiceAccount is required when the image is from a different project
		accounts := make([]*compute.ServiceAccount, 1)
		accounts[0] = &compute.ServiceAccount{
			Email:  "default",
			Scopes: nil,
		}
		instance.ServiceAccounts = accounts
	}

	src := d.Get("source_machine_image").(string)
	instance.SourceMachineImage = src

	// Add source machine image encryption key if specified
	if encryptionKey := expandSourceMachineImageEncryptionKey(d); encryptionKey != nil {
		instance.SourceMachineImageEncryptionKey = encryptionKey
	}

	tpl, err := tpgresource.ParseMachineImageFieldValue(src, d, config)
	if err != nil {
		return err
	}

	// obtain the project where the image resides (could be different from the default)
	tmp := strings.Split(src, "projects/")
	mi_project := strings.Split(tmp[len(tmp)-1], "/")[0]

	mi, err := config.NewComputeClient(userAgent).MachineImages.Get(mi_project, tpl.Name).Do()
	if err != nil {
		return err
	}

	instance.Disks, err = adjustInstanceFromMachineImageDisks(d, config, mi, zone, project)
	if err != nil {
		return err
	}

	// when we make the original call to expandComputeInstance expandScheduling is called, which sets default values.
	// However, we want the values to be read from the machine image instead.
	if _, hasSchedule := d.GetOk("scheduling"); !hasSchedule {
		instance.Scheduling = mi.SourceInstanceProperties.Scheduling
	}

	// Force send all top-level fields that have been set in case they're overridden to zero values.
	// Initialize ForceSendFields to empty so we don't get things that the instance resource
	// always force-sends.
	instance.ForceSendFields = []string{}
	for f, s := range computeInstanceFromMachineImageSchema() {
		// It seems that GetOkExists always returns true for sets.
		// TODO: confirm this and file issue against Terraform core.
		// In the meantime, don't force send sets.
		if s.Type == schema.TypeSet {
			continue
		}

		if _, exists := d.GetOkExists(f); exists {
			// Assume for now that all fields are exact snake_case versions of the API fields.
			// This won't necessarily always be true, but it serves as a good approximation and
			// can be adjusted later as we discover issues.
			instance.ForceSendFields = append(instance.ForceSendFields, tpgresource.SnakeToPascalCase(f))
		}
	}

	log.Printf("[INFO] Requesting instance creation")
	op, err := config.NewComputeClient(userAgent).Instances.Insert(project, zone.Name, instance).Do()
	if err != nil {
		return fmt.Errorf("Error creating instance: %s", err)
	}

	// Store the ID now
	d.SetId(fmt.Sprintf("projects/%s/zones/%s/instances/%s", project, z, instance.Name))

	// Wait for the operation to complete
	waitErr := ComputeOperationWaitTime(config, op, project,
		"instance to create", userAgent, d.Timeout(schema.TimeoutCreate))
	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return waitErr
	}

	return resourceComputeInstanceRead(d, meta)
}

// Instances have disks spread across multiple schema properties. This function
// ensures that overriding one of these properties does not override the others.
func adjustInstanceFromMachineImageDisks(d *schema.ResourceData, config *transport_tpg.Config, mi *compute.MachineImage, zone *compute.Zone, project string) ([]*compute.AttachedDisk, error) {
	disks := []*compute.AttachedDisk{}
	if _, hasBootDisk := d.GetOk("boot_disk"); hasBootDisk {
		bootDisk, err := expandBootDisk(d, config, project)
		if err != nil {
			return nil, err
		}
		disks = append(disks, bootDisk)
	} else {
		// boot disk was not overridden, so use the one from the machine image
		for _, disk := range mi.SourceInstanceProperties.Disks {
			if disk.Boot {
				newdisk := &compute.AttachedDisk{
					AutoDelete: disk.AutoDelete,
					Type:       disk.Type,
					DeviceName: disk.DeviceName,
				}
				disks = append(disks, newdisk)
				break
			}
		}
	}

	if _, hasScratchDisk := d.GetOk("scratch_disk"); hasScratchDisk {
		scratchDisks, err := expandScratchDisks(d, config, project)
		if err != nil {
			return nil, err
		}
		disks = append(disks, scratchDisks...)
	} else {
		// scratch disks were not overridden, so use the ones from the machine image
		for _, disk := range mi.SourceInstanceProperties.Disks {
			if disk.Type == "SCRATCH" {
				newdisk := &compute.AttachedDisk{
					AutoDelete: disk.AutoDelete,
					Type:       disk.Type,
					DeviceName: disk.DeviceName,
				}
				disks = append(disks, newdisk)
			}
		}
	}

	attachedDisksCount := d.Get("attached_disk.#").(int)
	if attachedDisksCount > 0 {
		for i := 0; i < attachedDisksCount; i++ {
			diskConfig := d.Get(fmt.Sprintf("attached_disk.%d", i)).(map[string]interface{})
			disk, err := expandAttachedDisk(diskConfig, d, config)
			if err != nil {
				return nil, err
			}

			disks = append(disks, disk)
		}
	} else {
		// attached disks were not overridden, so use the ones from the machine image
		for _, disk := range mi.SourceInstanceProperties.Disks {
			if !disk.Boot && disk.Type != "SCRATCH" {
				newdisk := &compute.AttachedDisk{
					AutoDelete: disk.AutoDelete,
					Type:       disk.Type,
					DeviceName: disk.DeviceName,
				}
				disks = append(disks, newdisk)
			}
		}
	}

	return disks, nil
}

func expandSourceMachineImageEncryptionKey(d tpgresource.TerraformResourceData) *compute.CustomerEncryptionKey {
	if v, ok := d.GetOk("source_machine_image_encryption_key"); !ok || v == nil {
		return nil
	}

	encryptionKeyList := d.Get("source_machine_image_encryption_key").([]interface{})
	if len(encryptionKeyList) == 0 || encryptionKeyList[0] == nil {
		return nil
	}

	encryptionKeyMap := encryptionKeyList[0].(map[string]interface{})
	encryptionKey := &compute.CustomerEncryptionKey{}

	if rawKey, ok := encryptionKeyMap["raw_key"]; ok && rawKey != "" {
		encryptionKey.RawKey = rawKey.(string)
	}

	if rsaKey, ok := encryptionKeyMap["rsa_encrypted_key"]; ok && rsaKey != "" {
		encryptionKey.RsaEncryptedKey = rsaKey.(string)
	}

	if kmsKey, ok := encryptionKeyMap["kms_key_name"]; ok && kmsKey != "" {
		encryptionKey.KmsKeyName = kmsKey.(string)
	}

	if kmsKeyServiceAccount, ok := encryptionKeyMap["kms_key_service_account"]; ok && kmsKeyServiceAccount != "" {
		encryptionKey.KmsKeyServiceAccount = kmsKeyServiceAccount.(string)
	}

	return encryptionKey
}

func flattenSourceMachineImageEncryptionKey(key *compute.CustomerEncryptionKey) []map[string]interface{} {
	if key == nil {
		return nil
	}

	m := map[string]interface{}{}

	if key.RawKey != "" {
		m["raw_key"] = key.RawKey
	}

	if key.RsaEncryptedKey != "" {
		m["rsa_encrypted_key"] = key.RsaEncryptedKey
	}

	if key.KmsKeyName != "" {
		m["kms_key_name"] = key.KmsKeyName
	}

	if key.KmsKeyServiceAccount != "" {
		m["kms_key_service_account"] = key.KmsKeyServiceAccount
	}

	if key.Sha256 != "" {
		m["sha256"] = key.Sha256
	}

	return []map[string]interface{}{m}
}
