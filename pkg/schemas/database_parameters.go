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
	ShowDatabaseParametersSchema = make(map[string]*schema.Schema)
	databaseParameters           = defs.ParameterDefsForLevel(parameterdefs.ParameterLevelDatabase)
)

func init() {
	for _, def := range databaseParameters {
		ShowDatabaseParametersSchema[def.FieldName()] = ParameterListSchema
	}
}

func DatabaseParametersToSchema(parameters []*sdk.Parameter, providerCtx *provider.Context) map[string]any {
	databaseParametersValue := make(map[string]any)
	for _, parameter := range parameters {
		if slices.ContainsFunc(databaseParameters, func(def parameterdefs.ParameterDef) bool { return def.SqlName == parameter.Key }) {
			databaseParametersValue[strings.ToLower(parameter.Key)] = []map[string]any{ParameterToSchemaReducedOutput(parameter, providerCtx)}
		}
	}
	return databaseParametersValue
}
