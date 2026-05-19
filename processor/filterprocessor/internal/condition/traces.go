// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package condition // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor/internal/condition"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/expr"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlresource"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlscope"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspan"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspanevent"
)

type TracesConsumer struct {
	resourceExpr  expr.BoolExpr[*ottlresource.TransformContext]
	scopeExpr     expr.BoolExpr[*ottlscope.TransformContext]
	spanExpr      expr.BoolExpr[*ottlspan.TransformContext]
	spanEventExpr expr.BoolExpr[*ottlspanevent.TransformContext]
}

// parsedTraceConditions is the type R for ParserCollection[R] that holds parsed OTTL conditions
type parsedTraceConditions struct {
	resourceConditions  []*ottl.Condition[*ottlresource.TransformContext]
	scopeConditions     []*ottl.Condition[*ottlscope.TransformContext]
	spanConditions      []*ottl.Condition[*ottlspan.TransformContext]
	spanEventConditions []*ottl.Condition[*ottlspanevent.TransformContext]
	telemetrySettings   component.TelemetrySettings
	errorMode           ottl.ErrorMode
}

func (tc TracesConsumer) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func newTraceConditionsFromResource(rc []*ottl.Condition[*ottlresource.TransformContext], telemetrySettings component.TelemetrySettings, errorMode ottl.ErrorMode) parsedTraceConditions {
	_ = "STUB: not implemented"
	return *new(parsedTraceConditions)
}

func newTraceConditionsFromScope(sc []*ottl.Condition[*ottlscope.TransformContext], telemetrySettings component.TelemetrySettings, errorMode ottl.ErrorMode) parsedTraceConditions {
	_ = "STUB: not implemented"
	return *new(parsedTraceConditions)
}

func newTracesConsumer(tc *parsedTraceConditions) TracesConsumer {
	_ = "STUB: not implemented"
	return *new(TracesConsumer)
}

type TraceParserCollection ottl.ParserCollection[parsedTraceConditions]

type TraceParserCollectionOption ottl.ParserCollectionOption[parsedTraceConditions]

func WithSpanParser(functions map[string]ottl.Factory[*ottlspan.TransformContext]) TraceParserCollectionOption {
	_ = "STUB: not implemented"
	return *new(TraceParserCollectionOption)
}

func WithSpanEventParser(functions map[string]ottl.Factory[*ottlspanevent.TransformContext]) TraceParserCollectionOption {
	_ = "STUB: not implemented"
	return *new(TraceParserCollectionOption)
}

func WithTraceErrorMode(errorMode ottl.ErrorMode) TraceParserCollectionOption {
	_ = "STUB: not implemented"
	return *new(TraceParserCollectionOption)
}

func WithTraceCommonParsers(functions map[string]ottl.Factory[*ottlresource.TransformContext]) TraceParserCollectionOption {
	_ = "STUB: not implemented"
	return *new(TraceParserCollectionOption)
}

func NewTraceParserCollection(settings component.TelemetrySettings, options ...TraceParserCollectionOption) (*TraceParserCollection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertSpanConditions(pc *ottl.ParserCollection[parsedTraceConditions], conditions ottl.ConditionsGetter, parsedConditions []*ottl.Condition[*ottlspan.TransformContext]) (parsedTraceConditions, error) {
	_ = "STUB: not implemented"
	return *new(parsedTraceConditions), nil
}

func convertSpanEventConditions(pc *ottl.ParserCollection[parsedTraceConditions], conditions ottl.ConditionsGetter, parsedConditions []*ottl.Condition[*ottlspanevent.TransformContext]) (parsedTraceConditions, error) {
	_ = "STUB: not implemented"
	return *new(parsedTraceConditions), nil
}

func (tpc *TraceParserCollection) ParseContextConditions(contextConditions ContextConditions) (TracesConsumer, error) {
	_ = "STUB: not implemented"
	return *new(TracesConsumer), nil
}
