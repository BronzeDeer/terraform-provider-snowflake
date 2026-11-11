package schemas

import (
	"github.com/Snowflake-Labs/terraform-provider-snowflake/v2/pkg/internal/provider"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/v2/pkg/provider/experimentalfeatures"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/v2/pkg/sdk"
)

// ParameterToSchemaReducedOutput limits the `parameters` output only to the value and level values.
// It utilizes the experimentalfeatures.ParametersReducedOutput experiment.
func ParameterToSchemaReducedOutput(parameter *sdk.Parameter, providerCtx *provider.Context) map[string]any {
	if providerCtx.Experiments.IsEnabled(experimentalfeatures.ParametersReducedOutput) {
		parameterSchema := make(map[string]any)
		parameterSchema["value"] = parameter.Value
		parameterSchema["level"] = string(parameter.Level)
		return parameterSchema
	} else {
		return ParameterToSchema(parameter)
	}
}
