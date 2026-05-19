// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package condition // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor/internal/condition"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/expr"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottllog"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlresource"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlscope"
)

type LogsConsumer struct {
	resourceExpr expr.BoolExpr[*ottlresource.TransformContext]
	scopeExpr    expr.BoolExpr[*ottlscope.TransformContext]
	logExpr      expr.BoolExpr[*ottllog.TransformContext]
}

// parsedLogConditions is the type R for ParserCollection[R] that holds parsed OTTL conditions
type parsedLogConditions struct {
	resourceConditions []*ottl.Condition[*ottlresource.TransformContext]
	scopeConditions    []*ottl.Condition[*ottlscope.TransformContext]
	logConditions      []*ottl.Condition[*ottllog.TransformContext]
	telemetrySettings  component.TelemetrySettings
	errorMode          ottl.ErrorMode
}

func (lc LogsConsumer) ConsumeLogs(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func newLogConditionsFromResource(rc []*ottl.Condition[*ottlresource.TransformContext], telemetrySettings component.TelemetrySettings, errorMode ottl.ErrorMode) parsedLogConditions {
	_ = "STUB: not implemented"
	return *new(parsedLogConditions)
}

func newLogConditionsFromScope(sc []*ottl.Condition[*ottlscope.TransformContext], telemetrySettings component.TelemetrySettings, errorMode ottl.ErrorMode) parsedLogConditions {
	_ = "STUB: not implemented"
	return *new(parsedLogConditions)
}

func newLogsConsumer(lc *parsedLogConditions) LogsConsumer {
	_ = "STUB: not implemented"
	return *new(LogsConsumer)
}

type LogParserCollection ottl.ParserCollection[parsedLogConditions]

type LogParserCollectionOption ottl.ParserCollectionOption[parsedLogConditions]

func WithLogParser(functions map[string]ottl.Factory[*ottllog.TransformContext]) LogParserCollectionOption {
	_ = "STUB: not implemented"
	return *new(LogParserCollectionOption)
}

func WithLogErrorMode(errorMode ottl.ErrorMode) LogParserCollectionOption {
	_ = "STUB: not implemented"
	return *new(LogParserCollectionOption)
}

func WithLogCommonParsers(functions map[string]ottl.Factory[*ottlresource.TransformContext]) LogParserCollectionOption {
	_ = "STUB: not implemented"
	return *new(LogParserCollectionOption)
}

func NewLogParserCollection(settings component.TelemetrySettings, options ...LogParserCollectionOption) (*LogParserCollection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertLogConditions(pc *ottl.ParserCollection[parsedLogConditions], conditions ottl.ConditionsGetter, parsedConditions []*ottl.Condition[*ottllog.TransformContext]) (parsedLogConditions, error) {
	_ = "STUB: not implemented"
	return *new(parsedLogConditions), nil
}

// ParseContextConditions parses the given ContextConditions and returns a LogsConsumer.
// For undefined context, each condition is parsed independently.
// Conditions are then grouped by their inferred context (resource, scope, log).
// The conditions group's error mode takes precedence over the processor-level error mode.
func (lpc *LogParserCollection) ParseContextConditions(contextConditions ContextConditions) (LogsConsumer, error) {
	_ = "STUB: not implemented"
	return *new(LogsConsumer), nil
}
