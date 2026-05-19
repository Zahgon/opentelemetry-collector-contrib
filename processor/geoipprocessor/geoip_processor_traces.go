// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package geoipprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/geoipprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/ptrace"
)

func (g *geoIPProcessor) processTraces(ctx context.Context, ts ptrace.Traces) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}
