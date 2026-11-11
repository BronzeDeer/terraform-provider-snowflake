package model

import (
	"github.com/Snowflake-Labs/terraform-provider-snowflake/v2/pkg/sdk"
)

func SequenceWithId(
	resourceName string,
	id sdk.SchemaObjectIdentifier,
) *SequenceModel {
	return Sequence(resourceName, id.DatabaseName(), id.SchemaName(), id.Name())
}
