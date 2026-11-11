package objectassert

import (
	"fmt"
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/v2/pkg/sdk"
)

func (o *OpenflowRuntimeAssert) HasKeyNotEmpty() *OpenflowRuntimeAssert {
	o.AddAssertion(func(t *testing.T, o *sdk.OpenflowRuntime) error {
		t.Helper()
		if o.Key == nil || *o.Key == "" {
			return fmt.Errorf("expected key to be not empty")
		}
		return nil
	})
	return o
}
