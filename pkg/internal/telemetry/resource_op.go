package telemetry

import (
	"context"
	"time"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/v2/internal/tracking"
	internalprovider "github.com/Snowflake-Labs/terraform-provider-snowflake/v2/pkg/internal/provider"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/v2/pkg/provider/resources"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func EmitResourceOp(ctx context.Context, meta any, resourceName resources.Resource, operation tracking.Operation, duration time.Duration, diags diag.Diagnostics, recovered any) {
	providerCtx, ok := meta.(*internalprovider.Context)
	if !ok || providerCtx == nil {
		return
	}
	emit(ctx, providerCtx.Client, typeResourceOp, providerCtx.SpanID, resourceOpFields(resourceName, operation, duration, diags, recovered))
}

func resourceOpFields(resourceName resources.Resource, operation tracking.Operation, duration time.Duration, diags diag.Diagnostics, recovered any) map[string]string {
	return map[string]string{
		"resource_name":  resourceName.String(),
		"operation_type": string(operation),
		"status":         string(getOperationStatus(diags, recovered)),
		"duration_s":     formatSeconds(duration),
	}
}
