package model

import "github.com/Snowflake-Labs/terraform-provider-snowflake/v2/pkg/sdk"

func WarehouseAdaptiveWithId(id sdk.AccountObjectIdentifier) *WarehouseAdaptiveModel {
	return WarehouseAdaptiveWithDefaultMeta(id.Name())
}
