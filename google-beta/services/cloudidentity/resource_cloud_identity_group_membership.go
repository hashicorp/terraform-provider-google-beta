// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/cloudidentity/GroupMembership.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package cloudidentity

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceCloudIdentityGroupMembership() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudIdentityGroupMembershipCreate,
		Read:   resourceCloudIdentityGroupMembershipRead,
		Update: resourceCloudIdentityGroupMembershipUpdate,
		Delete: resourceCloudIdentityGroupMembershipDelete,

		Importer: &schema.ResourceImporter{
			State: resourceCloudIdentityGroupMembershipImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"group": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The name of the Group to create this membership in.`,
			},
			"roles": {
				Type:     schema.TypeSet,
				Required: true,
				Description: `The MembershipRoles that apply to the Membership.
Must not contain duplicate MembershipRoles with the same name.`,
				Elem: cloudidentityGroupMembershipRolesSchema(),
				// Default schema.HashSchema is used.
			},
			"member_key": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `EntityKey of the member.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
							Description: `The ID of the entity.

For Google-managed entities, the id must be the email address of an existing
group or user.

For external-identity-mapped entities, the id must be a string conforming
to the Identity Source's requirements.

Must be unique within a namespace.`,
						},
						"namespace": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `The namespace in which the entity exists.

If not specified, the EntityKey represents a Google-managed entity
such as a Google user or a Google Group.

If specified, the EntityKey represents an external-identity-mapped group.
The namespace must correspond to an identity source created in Admin Console
and must be in the form of 'identitysources/{identity_source_id}'.`,
						},
					},
				},
				ExactlyOneOf: []string{"member_key", "preferred_member_key"},
			},
			"preferred_member_key": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `EntityKey of the member.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
							Description: `The ID of the entity.

For Google-managed entities, the id must be the email address of an existing
group or user.

For external-identity-mapped entities, the id must be a string conforming
to the Identity Source's requirements.

Must be unique within a namespace.`,
						},
						"namespace": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `The namespace in which the entity exists.

If not specified, the EntityKey represents a Google-managed entity
such as a Google user or a Google Group.

If specified, the EntityKey represents an external-identity-mapped group.
The namespace must correspond to an identity source created in Admin Console
and must be in the form of 'identitysources/{identity_source_id}'.`,
						},
					},
				},
				ExactlyOneOf: []string{"member_key", "preferred_member_key"},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time when the Membership was created.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The resource name of the Membership, of the form groups/{group_id}/memberships/{membership_id}.`,
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The type of the membership.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time when the Membership was last updated.`,
			},
		},
		UseJSONNumber: true,
	}
}

func cloudidentityGroupMembershipRolesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: verify.ValidateEnum([]string{"OWNER", "MANAGER", "MEMBER"}),
				Description:  `The name of the MembershipRole. Must be one of OWNER, MANAGER, MEMBER. Possible values: ["OWNER", "MANAGER", "MEMBER"]`,
			},
			"expiry_detail": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `The MembershipRole expiry details, only supported for MEMBER role.
Other roles cannot be accompanied with MEMBER role having expiry.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"expire_time": {
							Type:     schema.TypeString,
							Required: true,
							Description: `The time at which the MembershipRole will expire.

A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
resolution and up to nine fractional digits.

Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".`,
						},
					},
				},
			},
		},
	}
}

func resourceCloudIdentityGroupMembershipCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	memberKeyProp, err := expandCloudIdentityGroupMembershipMemberKey(d.Get("member_key"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("member_key"); !tpgresource.IsEmptyValue(reflect.ValueOf(memberKeyProp)) && (ok || !reflect.DeepEqual(v, memberKeyProp)) {
		obj["memberKey"] = memberKeyProp
	}
	preferredMemberKeyProp, err := expandCloudIdentityGroupMembershipPreferredMemberKey(d.Get("preferred_member_key"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("preferred_member_key"); !tpgresource.IsEmptyValue(reflect.ValueOf(preferredMemberKeyProp)) && (ok || !reflect.DeepEqual(v, preferredMemberKeyProp)) {
		obj["preferredMemberKey"] = preferredMemberKeyProp
	}
	rolesProp, err := expandCloudIdentityGroupMembershipRoles(d.Get("roles"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("roles"); !tpgresource.IsEmptyValue(reflect.ValueOf(rolesProp)) && (ok || !reflect.DeepEqual(v, rolesProp)) {
		obj["roles"] = rolesProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{CloudIdentityBasePath}}{{group}}/memberships")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new GroupMembership: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating GroupMembership: %s", err)
	}
	// Set computed resource properties from create API response so that they're available on the subsequent Read
	// call.
	err = resourceCloudIdentityGroupMembershipPostCreateSetComputedFields(d, meta, res)
	if err != nil {
		return fmt.Errorf("setting computed ID format fields: %w", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// `name` is autogenerated from the api so needs to be set post-create
	name, ok := res["name"]
	if !ok {
		respBody, ok := res["response"]
		if !ok {
			return fmt.Errorf("Create response didn't contain critical fields. Create may not have succeeded.")
		}

		name, ok = respBody.(map[string]interface{})["name"]
		if !ok {
			return fmt.Errorf("Create response didn't contain critical fields. Create may not have succeeded.")
		}
	}
	if err := d.Set("name", name.(string)); err != nil {
		return fmt.Errorf("Error setting name: %s", err)
	}
	d.SetId(name.(string))

	log.Printf("[DEBUG] Finished creating GroupMembership %q: %#v", d.Id(), res)

	return resourceCloudIdentityGroupMembershipRead(d, meta)
}

func resourceCloudIdentityGroupMembershipRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{CloudIdentityBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(transformCloudIdentityGroupMembershipReadError(err), d, fmt.Sprintf("CloudIdentityGroupMembership %q", d.Id()))
	}

	if err := d.Set("name", flattenCloudIdentityGroupMembershipName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading GroupMembership: %s", err)
	}
	if err := d.Set("member_key", flattenCloudIdentityGroupMembershipMemberKey(res["memberKey"], d, config)); err != nil {
		return fmt.Errorf("Error reading GroupMembership: %s", err)
	}
	if err := d.Set("preferred_member_key", flattenCloudIdentityGroupMembershipPreferredMemberKey(res["preferredMemberKey"], d, config)); err != nil {
		return fmt.Errorf("Error reading GroupMembership: %s", err)
	}
	if err := d.Set("create_time", flattenCloudIdentityGroupMembershipCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading GroupMembership: %s", err)
	}
	if err := d.Set("update_time", flattenCloudIdentityGroupMembershipUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading GroupMembership: %s", err)
	}
	if err := d.Set("roles", flattenCloudIdentityGroupMembershipRoles(res["roles"], d, config)); err != nil {
		return fmt.Errorf("Error reading GroupMembership: %s", err)
	}
	if err := d.Set("type", flattenCloudIdentityGroupMembershipType(res["type"], d, config)); err != nil {
		return fmt.Errorf("Error reading GroupMembership: %s", err)
	}

	return nil
}

func resourceCloudIdentityGroupMembershipUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	d.Partial(true)

	if d.HasChange("roles") {
		url, err := tpgresource.ReplaceVars(d, config, "{{CloudIdentityBasePath}}{{name}}:modifyMembershipRoles")
		if err != nil {
			return err
		}

		// err == nil indicates that the billing_project value was found
		if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
			billingProject = bp
		}

		// Return object for modifyMembershipRoles (we build request object from scratch, without using `obj`)
		b, a := d.GetChange("roles")
		before := b.(*schema.Set)
		after := a.(*schema.Set)

		ignoreUpdateR := make(map[string]struct{})
		addRoleList := after.Difference(before).List()
		removeRoleList := before.Difference(after).List()

		var updateRolesParams []map[string]interface{}
		for _, addR := range addRoleList {
			ar := addR.(map[string]interface{})["name"].(string)
			ae := addR.(map[string]interface{})["expiry_detail"].([]interface{})
			for _, removeR := range removeRoleList {
				if ar == removeR.(map[string]interface{})["name"].(string) {
					ignoreUpdateR[ar] = struct{}{}
					var updateR map[string]interface{}
					if len(ae) == 0 {
						updateR = map[string]interface{}{"name": ar}
					} else {
						updateR = map[string]interface{}{"name": ar, "expiry_detail": ae[0]}
					}
					updateP := map[string]interface{}{"field_mask": "expiryDetail.expire_time", "membership_role": updateR}
					updateRolesParams = append(updateRolesParams, updateP)
				}
			}
		}

		var addRoles []map[string]interface{}
		for _, r := range addRoleList {
			name := r.(map[string]interface{})["name"].(string)
			if _, ignore := ignoreUpdateR[name]; ignore {
				continue
			}
			expiryDetail := r.(map[string]interface{})["expiry_detail"].([]interface{})
			if len(expiryDetail) == 0 {
				addRoles = append(addRoles, map[string]interface{}{"name": name})
			} else {
				addRoles = append(addRoles, map[string]interface{}{"name": name, "expiry_detail": expiryDetail[0]})
			}
		}
		var removeRoles []string
		for _, r := range removeRoleList {
			name := r.(map[string]interface{})["name"].(string)
			if _, ignore := ignoreUpdateR[name]; ignore {
				continue
			}
			removeRoles = append(removeRoles, name)
		}

		// ref: https://cloud.google.com/identity/docs/reference/rest/v1/groups.memberships/modifyMembershipRoles#request-body
		// Only single operation per request is allowed.
		if len(removeRoles) > 0 {
			res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "POST",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: userAgent,
				Body:      map[string]interface{}{"removeRoles": removeRoles},
				Timeout:   d.Timeout(schema.TimeoutUpdate),
			})
			if err != nil {
				return fmt.Errorf("Error removing GroupMembership %q: %s", d.Id(), err)
			} else {
				log.Printf("[DEBUG] Finished removing GroupMembership %q: %#v", d.Id(), res)
			}
		}
		if len(updateRolesParams) > 0 {
			res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "POST",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: userAgent,
				Body:      map[string]interface{}{"updateRolesParams": updateRolesParams},
				Timeout:   d.Timeout(schema.TimeoutUpdate),
			})
			if err != nil {
				return fmt.Errorf("Error updating GroupMembership %q: %s", d.Id(), err)
			} else {
				log.Printf("[DEBUG] Finished updating GroupMembership %q: %#v", d.Id(), res)
			}
		}
		if len(addRoles) > 0 {
			res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "POST",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: userAgent,
				Body:      map[string]interface{}{"addRoles": addRoles},
				Timeout:   d.Timeout(schema.TimeoutUpdate),
			})
			if err != nil {
				return fmt.Errorf("Error adding GroupMembership %q: %s", d.Id(), err)
			} else {
				log.Printf("[DEBUG] Finished adding GroupMembership %q: %#v", d.Id(), res)
			}
		}
	}

	d.Partial(false)

	return resourceCloudIdentityGroupMembershipRead(d, meta)
}

func resourceCloudIdentityGroupMembershipDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{CloudIdentityBasePath}}{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting GroupMembership %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "GroupMembership")
	}

	log.Printf("[DEBUG] Finished deleting GroupMembership %q: %#v", d.Id(), res)
	return nil
}

func resourceCloudIdentityGroupMembershipImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^(?P<name>.+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Configure "group" property, which does not appear in the response body.
	group := regexp.MustCompile(`groups/[^/]+`).FindString(id)
	if err := d.Set("group", group); err != nil {
		return nil, fmt.Errorf("Error setting group property: %s", err)
	}

	return []*schema.ResourceData{d}, nil
}

func flattenCloudIdentityGroupMembershipName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudIdentityGroupMembershipMemberKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["id"] =
		flattenCloudIdentityGroupMembershipMemberKeyId(original["id"], d, config)
	transformed["namespace"] =
		flattenCloudIdentityGroupMembershipMemberKeyNamespace(original["namespace"], d, config)
	return []interface{}{transformed}
}
func flattenCloudIdentityGroupMembershipMemberKeyId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudIdentityGroupMembershipMemberKeyNamespace(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudIdentityGroupMembershipPreferredMemberKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["id"] =
		flattenCloudIdentityGroupMembershipPreferredMemberKeyId(original["id"], d, config)
	transformed["namespace"] =
		flattenCloudIdentityGroupMembershipPreferredMemberKeyNamespace(original["namespace"], d, config)
	return []interface{}{transformed}
}
func flattenCloudIdentityGroupMembershipPreferredMemberKeyId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudIdentityGroupMembershipPreferredMemberKeyNamespace(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudIdentityGroupMembershipCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudIdentityGroupMembershipUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudIdentityGroupMembershipRoles(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(schema.HashResource(cloudidentityGroupMembershipRolesSchema()), []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"name":          flattenCloudIdentityGroupMembershipRolesName(original["name"], d, config),
			"expiry_detail": flattenCloudIdentityGroupMembershipRolesExpiryDetail(original["expiryDetail"], d, config),
		})
	}
	return transformed
}
func flattenCloudIdentityGroupMembershipRolesName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudIdentityGroupMembershipRolesExpiryDetail(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["expire_time"] =
		flattenCloudIdentityGroupMembershipRolesExpiryDetailExpireTime(original["expireTime"], d, config)
	return []interface{}{transformed}
}
func flattenCloudIdentityGroupMembershipRolesExpiryDetailExpireTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudIdentityGroupMembershipType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandCloudIdentityGroupMembershipMemberKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedId, err := expandCloudIdentityGroupMembershipMemberKeyId(original["id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["id"] = transformedId
	}

	transformedNamespace, err := expandCloudIdentityGroupMembershipMemberKeyNamespace(original["namespace"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNamespace); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["namespace"] = transformedNamespace
	}

	return transformed, nil
}

func expandCloudIdentityGroupMembershipMemberKeyId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudIdentityGroupMembershipMemberKeyNamespace(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudIdentityGroupMembershipPreferredMemberKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedId, err := expandCloudIdentityGroupMembershipPreferredMemberKeyId(original["id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["id"] = transformedId
	}

	transformedNamespace, err := expandCloudIdentityGroupMembershipPreferredMemberKeyNamespace(original["namespace"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNamespace); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["namespace"] = transformedNamespace
	}

	return transformed, nil
}

func expandCloudIdentityGroupMembershipPreferredMemberKeyId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudIdentityGroupMembershipPreferredMemberKeyNamespace(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudIdentityGroupMembershipRoles(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandCloudIdentityGroupMembershipRolesName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedExpiryDetail, err := expandCloudIdentityGroupMembershipRolesExpiryDetail(original["expiry_detail"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedExpiryDetail); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["expiryDetail"] = transformedExpiryDetail
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCloudIdentityGroupMembershipRolesName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudIdentityGroupMembershipRolesExpiryDetail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedExpireTime, err := expandCloudIdentityGroupMembershipRolesExpiryDetailExpireTime(original["expire_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExpireTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["expireTime"] = transformedExpireTime
	}

	return transformed, nil
}

func expandCloudIdentityGroupMembershipRolesExpiryDetailExpireTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceCloudIdentityGroupMembershipPostCreateSetComputedFields(d *schema.ResourceData, meta interface{}, res map[string]interface{}) error {
	config := meta.(*transport_tpg.Config)
	if err := d.Set("name", flattenCloudIdentityGroupMembershipName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}
	return nil
}
