package resourceassert

import (
	"github.com/Snowflake-Labs/terraform-provider-snowflake/v2/pkg/internal/collections"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/v2/pkg/sdk"
)

func (s *StorageIntegrationAwsResourceAssert) HasStorageAllowedLocationsStorageLocation(expected ...sdk.StorageLocation) *StorageIntegrationAwsResourceAssert {
	return s.HasStorageAllowedLocations(collections.Map(expected, func(v sdk.StorageLocation) string {
		return v.Path
	})...)
}

func (s *StorageIntegrationAwsResourceAssert) HasStorageBlockedLocationsStorageLocation(expected ...sdk.StorageLocation) *StorageIntegrationAwsResourceAssert {
	return s.HasStorageBlockedLocations(collections.Map(expected, func(v sdk.StorageLocation) string {
		return v.Path
	})...)
}
