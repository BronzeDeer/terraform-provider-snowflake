package invokeactionassert

import (
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/v2/pkg/acceptance/bettertestspoc/assert"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/v2/pkg/acceptance/helpers"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/v2/pkg/sdk"
)

func SecurityIntegrationDoesNotExist(t *testing.T, id sdk.AccountObjectIdentifier) assert.TestCheckFuncProvider {
	t.Helper()
	return newNonExistenceCheck(
		sdk.ObjectTypeSecurityIntegration,
		id,
		func(testClient *helpers.TestClient) showByIDFunc[*sdk.SecurityIntegration, sdk.AccountObjectIdentifier] {
			return testClient.SecurityIntegration.Show
		},
	)
}
