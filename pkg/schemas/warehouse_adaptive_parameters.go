package schemas

import (
	"slices"
	"strings"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/v2/pkg/internal/provider"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/v2/pkg/sdk"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/v2/pkg/sdk/generator/defs"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/v2/pkg/sdk/generator/parameterdefs"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	ShowWarehouseAdaptiveParametersSchema = make(map[string]*schema.Schema)
	warehouseAdaptiveParameters           = defs.ParameterDefsForLevel(parameterdefs.ParameterLevelWarehouseAdaptive)
)

func init() {
	for _, def := range warehouseAdaptiveParameters {
		ShowWarehouseAdaptiveParametersSchema[def.FieldName()] = ParameterListSchema
	}
}

func WarehouseAdaptiveParametersToSchema(parameters []*sdk.Parameter, providerCtx *provider.Context) map[string]any {
	warehouseParametersValue := make(map[string]any)
	for _, parameter := range parameters {
		if slices.ContainsFunc(warehouseAdaptiveParameters, func(def parameterdefs.ParameterDef) bool { return def.SqlName == parameter.Key }) {
			warehouseParametersValue[strings.ToLower(parameter.Key)] = []map[string]any{ParameterToSchemaReducedOutput(parameter, providerCtx)}
		}
	}
	return warehouseParametersValue
}
