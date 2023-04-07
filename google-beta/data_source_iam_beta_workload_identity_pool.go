package google

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIAMBetaWorkloadIdentityPool() *schema.Resource {

	dsSchema := datasourceSchemaFromResourceSchema(ResourceIAMBetaWorkloadIdentityPool().Schema)
	addRequiredFieldsToSchema(dsSchema, "workload_identity_pool_id")
	addOptionalFieldsToSchema(dsSchema, "project")

	return &schema.Resource{
		Read:   dataSourceIAMBetaWorkloadIdentityPoolRead,
		Schema: dsSchema,
	}
}

func dataSourceIAMBetaWorkloadIdentityPoolRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	id, err := ReplaceVars(d, config, "projects/{{project}}/locations/global/workloadIdentityPools/{{workload_identity_pool_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)
	return resourceIAMBetaWorkloadIdentityPoolRead(d, meta)

}
