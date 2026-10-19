package schemas

import "github.com/Snowflake-Labs/terraform-provider-snowflake/v2/pkg/sdk"

func SchemaDetailsListToSchema(details []sdk.SchemaDetails) []map[string]any {
	result := make([]map[string]any, len(details))
	for i := range details {
		result[i] = SchemaDetailsToSchema(&details[i])
	}
	return result
}
