package resourceassert

import (
	"github.com/Snowflake-Labs/terraform-provider-snowflake/v2/pkg/internal/collections"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/v2/pkg/sdk"
)

// TODO [SNOW-3113138]: extract common conversions (like Name/FullyQualifiedName to string)
func (c *PrimaryConnectionResourceAssert) HasExactlyFailoverToAccountsInOrder(expected ...sdk.AccountIdentifier) *PrimaryConnectionResourceAssert {
	return c.HasEnableFailoverToAccounts(collections.Map(expected, func(v sdk.AccountIdentifier) string {
		return v.Name()
	})...)
}
