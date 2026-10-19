package resourceshowoutputassert

import (
	"github.com/Snowflake-Labs/terraform-provider-snowflake/v2/pkg/acceptance/bettertestspoc/assert"
)

func (p *PostgresInstanceDescribeOutputAssert) HasCreatedOnNotEmpty() *PostgresInstanceDescribeOutputAssert {
	p.AddAssertion(assert.ValuePresent("created_on"))
	return p
}
