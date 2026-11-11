package gen

import "github.com/Snowflake-Labs/terraform-provider-snowflake/v2/pkg/internal/ssot"

type (
	Enum                    = ssot.Enum
	EnumValueRepresentation = ssot.EnumValueRepresentation
)

func NewEnum(name, namePlural string, values ...string) *Enum {
	return ssot.NewEnum(name, namePlural, values...)
}
