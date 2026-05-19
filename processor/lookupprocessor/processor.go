// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package lookupprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/lookupprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottllog"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/lookupprocessor/lookupsource"
)

type lookupProcessor struct {
	source  lookupsource.Source
	lookups []parsedLookup
	logger  *zap.Logger
}

func newLookupProcessor(source lookupsource.Source, lookups []parsedLookup, logger *zap.Logger) *lookupProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (p *lookupProcessor) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *lookupProcessor) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *lookupProcessor) processLogs(ctx context.Context, ld plog.Logs) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

func (p *lookupProcessor) processLookup(
	ctx context.Context,
	tCtx *ottllog.TransformContext,
	lookup *parsedLookup,
	logAttrs pcommon.Map,
	resourceAttrs pcommon.Map,
) {
	_ = "STUB: not implemented"
	return
}

// extractValue gets the value to write for a given attribute mapping.
// Returns (value, shouldWrite).
func extractValue(result any, found bool, attr *AttributeMapping) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// 1:1 scalar lookup — use entire result

// 1:N map lookup — extract the named field

func anyToString(v any) string { _ = "STUB: not implemented"; return "" }
