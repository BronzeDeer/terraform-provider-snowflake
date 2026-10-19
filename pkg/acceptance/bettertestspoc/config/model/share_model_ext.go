package model

import (
	"github.com/Snowflake-Labs/terraform-provider-snowflake/v2/pkg/internal/collections"
	"github.com/hashicorp/terraform-plugin-testing/config"
)

func (s *ShareModel) WithAccounts(accounts ...string) *ShareModel {
	accountStringVariables := collections.Map(accounts, func(account string) config.Variable { return config.StringVariable(account) })
	s.WithAccountsValue(config.ListVariable(accountStringVariables...))
	return s
}
