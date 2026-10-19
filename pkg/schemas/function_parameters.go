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
	ShowFunctionParametersSchema = make(map[string]*schema.Schema)
	functionParameters           = defs.ParameterDefsForLevel(parameterdefs.ParameterLevelFunction)
)

func init() {
	for _, def := range functionParameters {
		ShowFunctionParametersSchema[def.FieldName()] = ParameterListSchema
	}
}

func FunctionParametersToSchema(parameters []*sdk.Parameter, providerCtx *provider.Context) map[string]any {
	functionParametersValue := make(map[string]any)
	for _, parameter := range parameters {
		if slices.ContainsFunc(functionParameters, func(def parameterdefs.ParameterDef) bool { return def.SqlName == parameter.Key }) {
			functionParametersValue[strings.ToLower(parameter.Key)] = []map[string]any{ParameterToSchemaReducedOutput(parameter, providerCtx)}
		}
	}
	return functionParametersValue
}
