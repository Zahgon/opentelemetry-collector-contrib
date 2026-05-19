// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package common // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor/internal/common"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/expr"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottllog"
)

type LogsConsumer interface {
	Context() ContextID
	ConsumeLogs(ctx context.Context, ld plog.Logs) error
}

type logStatements struct {
	ottl.StatementSequence[*ottllog.TransformContext]
	expr.BoolExpr[*ottllog.TransformContext]
}

func (logStatements) Context() ContextID { _ = "STUB: not implemented"; return *new(ContextID) }

func (l logStatements) ConsumeLogs(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

type LogParserCollection ottl.ParserCollection[LogsConsumer]

type LogParserCollectionOption ottl.ParserCollectionOption[LogsConsumer]

func WithLogParser(functions map[string]ottl.Factory[*ottllog.TransformContext]) LogParserCollectionOption {
	_ = "STUB: not implemented"
	return *new(LogParserCollectionOption)
}

func WithLogErrorMode(errorMode ottl.ErrorMode) LogParserCollectionOption {
	_ = "STUB: not implemented"
	return *new(LogParserCollectionOption)
}

func NewLogParserCollection(settings component.TelemetrySettings, options ...LogParserCollectionOption) (*LogParserCollection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertLogStatements(pc *ottl.ParserCollection[LogsConsumer], statements ottl.StatementsGetter, parsedStatements []*ottl.Statement[*ottllog.TransformContext]) (LogsConsumer, error) {
	_ = "STUB: not implemented"
	return *new(LogsConsumer), nil
}

func (lpc *LogParserCollection) ParseContextStatements(contextStatements ContextStatements) (LogsConsumer, error) {
	_ = "STUB: not implemented"
	return *new(LogsConsumer), nil
}
