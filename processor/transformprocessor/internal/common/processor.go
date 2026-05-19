// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package common // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor/internal/common"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/pprofile"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/expr"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlresource"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlscope"
)

var _ baseContext = &resourceStatements{}

type resourceStatements struct {
	ottl.StatementSequence[*ottlresource.TransformContext]
	expr.BoolExpr[*ottlresource.TransformContext]
}

func (resourceStatements) Context() ContextID { _ = "STUB: not implemented"; return *new(ContextID) }

func (r resourceStatements) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func (r resourceStatements) ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (r resourceStatements) ConsumeLogs(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (r resourceStatements) ConsumeProfiles(ctx context.Context, ld pprofile.Profiles) error {
	_ = "STUB: not implemented"
	return nil
}

var _ baseContext = &scopeStatements{}

type scopeStatements struct {
	ottl.StatementSequence[*ottlscope.TransformContext]
	expr.BoolExpr[*ottlscope.TransformContext]
}

func (scopeStatements) Context() ContextID { _ = "STUB: not implemented"; return *new(ContextID) }

func (s scopeStatements) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func (s scopeStatements) ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (s scopeStatements) ConsumeLogs(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (s scopeStatements) ConsumeProfiles(ctx context.Context, ld pprofile.Profiles) error {
	_ = "STUB: not implemented"
	return nil
}

type baseContext interface {
	TracesConsumer
	MetricsConsumer
	LogsConsumer
	ProfilesConsumer
}

func withCommonContextParsers[R any]() ottl.ParserCollectionOption[R] {
	_ = "STUB: not implemented"
	return nil
}

func parseResourceContextStatements[R any](
	pc *ottl.ParserCollection[R],
	statements ottl.StatementsGetter,
	parsedStatements []*ottl.Statement[*ottlresource.TransformContext],
) (R, error) {
	_ = "STUB: not implemented"
	return *new(R), nil
}

func parseScopeContextStatements[R any](
	pc *ottl.ParserCollection[R],
	statements ottl.StatementsGetter,
	parsedStatements []*ottl.Statement[*ottlscope.TransformContext],
) (R, error) {
	_ = "STUB: not implemented"
	return *new(R), nil
}

func parseGlobalExpr[K, O any](
	boolExprFunc func([]string, map[string]ottl.Factory[K], ottl.ErrorMode, component.TelemetrySettings, []O) (*ottl.ConditionSequence[K], error),
	conditions []string,
	errorMode ottl.ErrorMode,
	settings component.TelemetrySettings,
	standardFuncs map[string]ottl.Factory[K],
	parserOptions []O,
) (expr.BoolExpr[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// By default, set the global expression to always true unless conditions are specified.
