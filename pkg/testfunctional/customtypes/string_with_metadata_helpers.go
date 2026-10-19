package customtypes

import "github.com/Snowflake-Labs/terraform-provider-snowflake/v2/pkg/sdk"

func StringWithMetadataAttributeCreate(v StringWithMetadataValue, createField **string) {
	if !v.IsNull() {
		*createField = sdk.String(v.ValueString())
	}
}
