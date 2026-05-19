// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package condition // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor/internal/condition"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/expr"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottldatapoint"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlmetric"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlresource"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlscope"
)

type MetricsConsumer struct {
	resourceExpr  expr.BoolExpr[*ottlresource.TransformContext]
	scopeExpr     expr.BoolExpr[*ottlscope.TransformContext]
	metricExpr    expr.BoolExpr[*ottlmetric.TransformContext]
	dataPointExpr expr.BoolExpr[*ottldatapoint.TransformContext]
}

// parsedMetricConditions is the type R for ParserCollection[R] that holds parsed OTTL conditions
type parsedMetricConditions struct {
	resourceConditions  []*ottl.Condition[*ottlresource.TransformContext]
	scopeConditions     []*ottl.Condition[*ottlscope.TransformContext]
	metricConditions    []*ottl.Condition[*ottlmetric.TransformContext]
	dataPointConditions []*ottl.Condition[*ottldatapoint.TransformContext]
	telemetrySettings   component.TelemetrySettings
	errorMode           ottl.ErrorMode
}

func (mc MetricsConsumer) ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

//exhaustive:enforce

func (mc MetricsConsumer) handleNumberDataPoints(ctx context.Context, rm pmetric.ResourceMetrics, sm pmetric.ScopeMetrics, m pmetric.Metric, dps pmetric.NumberDataPointSlice) error {
	_ = "STUB: not implemented"
	return nil
}

func (mc MetricsConsumer) handleHistogramDataPoints(ctx context.Context, rm pmetric.ResourceMetrics, sm pmetric.ScopeMetrics, m pmetric.Metric, dps pmetric.HistogramDataPointSlice) error {
	_ = "STUB: not implemented"
	return nil
}

func (mc MetricsConsumer) handleExponentialHistogramDataPoints(ctx context.Context, rm pmetric.ResourceMetrics, sm pmetric.ScopeMetrics, m pmetric.Metric, dps pmetric.ExponentialHistogramDataPointSlice) error {
	_ = "STUB: not implemented"
	return nil
}

func (mc MetricsConsumer) handleSummaryDataPoints(ctx context.Context, rm pmetric.ResourceMetrics, sm pmetric.ScopeMetrics, m pmetric.Metric, dps pmetric.SummaryDataPointSlice) error {
	_ = "STUB: not implemented"
	return nil
}

func newMetricConditionsFromResource(rc []*ottl.Condition[*ottlresource.TransformContext], telemetrySettings component.TelemetrySettings, errorMode ottl.ErrorMode) parsedMetricConditions {
	_ = "STUB: not implemented"
	return *new(parsedMetricConditions)
}

func newMetricConditionsFromScope(sc []*ottl.Condition[*ottlscope.TransformContext], telemetrySettings component.TelemetrySettings, errorMode ottl.ErrorMode) parsedMetricConditions {
	_ = "STUB: not implemented"
	return *new(parsedMetricConditions)
}

func newMetricsConsumer(mc *parsedMetricConditions) MetricsConsumer {
	_ = "STUB: not implemented"
	return *new(MetricsConsumer)
}

type MetricParserCollection ottl.ParserCollection[parsedMetricConditions]

type MetricParserCollectionOption ottl.ParserCollectionOption[parsedMetricConditions]

func WithMetricParser(functions map[string]ottl.Factory[*ottlmetric.TransformContext]) MetricParserCollectionOption {
	_ = "STUB: not implemented"
	return *new(MetricParserCollectionOption)
}

func WithDataPointParser(functions map[string]ottl.Factory[*ottldatapoint.TransformContext]) MetricParserCollectionOption {
	_ = "STUB: not implemented"
	return *new(MetricParserCollectionOption)
}

func WithMetricErrorMode(errorMode ottl.ErrorMode) MetricParserCollectionOption {
	_ = "STUB: not implemented"
	return *new(MetricParserCollectionOption)
}

func WithMetricCommonParsers(functions map[string]ottl.Factory[*ottlresource.TransformContext]) MetricParserCollectionOption {
	_ = "STUB: not implemented"
	return *new(MetricParserCollectionOption)
}

func NewMetricParserCollection(settings component.TelemetrySettings, options ...MetricParserCollectionOption) (*MetricParserCollection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertMetricConditions(pc *ottl.ParserCollection[parsedMetricConditions], conditions ottl.ConditionsGetter, parsedConditions []*ottl.Condition[*ottlmetric.TransformContext]) (parsedMetricConditions, error) {
	_ = "STUB: not implemented"
	return *new(parsedMetricConditions), nil
}

func convertDataPointConditions(pc *ottl.ParserCollection[parsedMetricConditions], conditions ottl.ConditionsGetter, parsedConditions []*ottl.Condition[*ottldatapoint.TransformContext]) (parsedMetricConditions, error) {
	_ = "STUB: not implemented"
	return *new(parsedMetricConditions), nil
}

func (mpc *MetricParserCollection) ParseContextConditions(contextConditions ContextConditions) (MetricsConsumer, error) {
	_ = "STUB: not implemented"
	return *new(MetricsConsumer), nil
}
