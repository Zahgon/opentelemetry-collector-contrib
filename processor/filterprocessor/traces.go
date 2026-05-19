// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package filterprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/expr"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlresource"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspan"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspanevent"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor/internal/condition"
)

type filterSpanProcessor struct {
	consumers         []condition.TracesConsumer
	skipResourceExpr  expr.BoolExpr[*ottlresource.TransformContext]
	skipSpanExpr      expr.BoolExpr[*ottlspan.TransformContext]
	skipSpanEventExpr expr.BoolExpr[*ottlspanevent.TransformContext]
	telemetry         *filterTelemetry
	logger            *zap.Logger
}

func newFilterSpansProcessor(set processor.Settings, cfg *Config) (*filterSpanProcessor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// processTraces filters the given spans of a traces based off the filterSpanProcessor's filters.
func (fsp *filterSpanProcessor) processTraces(ctx context.Context, td ptrace.Traces) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

func (fsp *filterSpanProcessor) processSkipExpression(ctx context.Context, td ptrace.Traces) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

func (fsp *filterSpanProcessor) processConditions(ctx context.Context, td ptrace.Traces) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}
