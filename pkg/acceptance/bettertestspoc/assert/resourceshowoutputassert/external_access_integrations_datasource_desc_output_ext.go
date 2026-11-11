package resourceshowoutputassert

import (
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/v2/pkg/acceptance/bettertestspoc/assert"
)

func ExternalAccessIntegrationsDatasourceDescribeOutput(t *testing.T, name string) *ExternalAccessIntegrationDescribeOutputAssert {
	t.Helper()
	return ExternalAccessIntegrationsDatasourceDescribeOutputOnIdx(t, name, 0)
}

func ExternalAccessIntegrationsDatasourceDescribeOutputOnIdx(t *testing.T, name string, idx int) *ExternalAccessIntegrationDescribeOutputAssert {
	t.Helper()
	return &ExternalAccessIntegrationDescribeOutputAssert{
		ResourceAssert: assert.NewDatasourceDescribeOutputAssert(name, "external_access_integrations", idx),
	}
}
