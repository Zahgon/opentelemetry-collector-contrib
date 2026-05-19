// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package common // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor/internal/common"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/expr"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspan"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspanevent"
)

type TracesConsumer interface {
	Context() ContextID
	ConsumeTraces(ctx context.Context, td ptrace.Traces) error
}

type traceStatements struct {
	ottl.StatementSequence[*ottlspan.TransformContext]
	expr.BoolExpr[*ottlspan.TransformContext]
}

func (traceStatements) Context() ContextID { _ = "STUB: not implemented"; return *new(ContextID) }

func (t traceStatements) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

type spanEventStatements struct {
	ottl.StatementSequence[*ottlspanevent.TransformContext]
	expr.BoolExpr[*ottlspanevent.TransformContext]
}

func (spanEventStatements) Context() ContextID { _ = "STUB: not implemented"; return *new(ContextID) }

func (s spanEventStatements) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

type TraceParserCollection ottl.ParserCollection[TracesConsumer]

type TraceParserCollectionOption ottl.ParserCollectionOption[TracesConsumer]

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

func NewTraceParserCollection(settings component.TelemetrySettings, options ...TraceParserCollectionOption) (*TraceParserCollection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertSpanStatements(pc *ottl.ParserCollection[TracesConsumer], statements ottl.StatementsGetter, parsedStatements []*ottl.Statement[*ottlspan.TransformContext]) (TracesConsumer, error) {
	_ = "STUB: not implemented"
	return *new(TracesConsumer), nil
}

func convertSpanEventStatements(pc *ottl.ParserCollection[TracesConsumer], statements ottl.StatementsGetter, parsedStatements []*ottl.Statement[*ottlspanevent.TransformContext]) (TracesConsumer, error) {
	_ = "STUB: not implemented"
	return *new(TracesConsumer), nil
}

func (tpc *TraceParserCollection) ParseContextStatements(contextStatements ContextStatements) (TracesConsumer, error) {
	_ = "STUB: not implemented"
	return *new(TracesConsumer), nil
}
