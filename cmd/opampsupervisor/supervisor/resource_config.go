// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package supervisor

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pcommon"
	telemetryconfig "go.opentelemetry.io/contrib/otelconf/v0.3.0"

	"github.com/open-telemetry/opentelemetry-collector-contrib/cmd/opampsupervisor/supervisor/config"
)

func buildSupervisorResourceConfig(cfg *config.ResourceConfig) (*telemetryconfig.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func resourceConfigToPcommon(ctx context.Context, resourceCfg *telemetryconfig.Resource) (pcommon.Resource, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource), nil
}
