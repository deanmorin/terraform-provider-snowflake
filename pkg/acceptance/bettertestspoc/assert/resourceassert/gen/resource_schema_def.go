package gen

import (
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/internal/genhelpers"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/resources"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ResourceSchemaDef struct {
	name   string
	schema map[string]*schema.Schema
}

func GetResourceSchemaDetails() []genhelpers.ResourceSchemaDetails {
	allResourceSchemas := allResourceSchemaDefs
	allResourceSchemasDetails := make([]genhelpers.ResourceSchemaDetails, len(allResourceSchemas))
	for idx, s := range allResourceSchemas {
		allResourceSchemasDetails[idx] = genhelpers.ExtractResourceSchemaDetails(s.name, s.schema)
	}
	return allResourceSchemasDetails
}

var allResourceSchemaDefs = []ResourceSchemaDef{
	{
		name:   "Warehouse",
		schema: resources.Warehouse().Schema,
	},
	{
		name:   "User",
		schema: resources.User().Schema,
	},
	{
		name:   "View",
		schema: resources.View().Schema,
	},
	{
		name:   "DatabaseRole",
		schema: resources.DatabaseRole().Schema,
	},
	{
		name:   "ResourceMonitor",
		schema: resources.ResourceMonitor().Schema,
	},
	{
		name:   "RowAccessPolicy",
		schema: resources.RowAccessPolicy().Schema,
	},
}
