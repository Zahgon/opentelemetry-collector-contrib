// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package geoipprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/geoipprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/plog"
)

func (g *geoIPProcessor) processLogs(ctx context.Context, ls plog.Logs) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}
