// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package common // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor/internal/common"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/expr"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottldatapoint"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlmetric"
)

type MetricsConsumer interface {
	Context() ContextID
	ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error
}

type metricStatements struct {
	ottl.StatementSequence[*ottlmetric.TransformContext]
	expr.BoolExpr[*ottlmetric.TransformContext]
}

func (metricStatements) Context() ContextID { _ = "STUB: not implemented"; return *new(ContextID) }

func (m metricStatements) ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

type dataPointStatements struct {
	ottl.StatementSequence[*ottldatapoint.TransformContext]
	expr.BoolExpr[*ottldatapoint.TransformContext]
}

func (dataPointStatements) Context() ContextID { _ = "STUB: not implemented"; return *new(ContextID) }

func (d dataPointStatements) ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

//exhaustive:enforce

func (d dataPointStatements) handleNumberDataPoints(ctx context.Context, resourceMetrics pmetric.ResourceMetrics, scopeMetrics pmetric.ScopeMetrics, metric pmetric.Metric, dps pmetric.NumberDataPointSlice) error {
	_ = "STUB: not implemented"
	return nil
}

func (d dataPointStatements) handleHistogramDataPoints(ctx context.Context, resourceMetrics pmetric.ResourceMetrics, scopeMetrics pmetric.ScopeMetrics, metric pmetric.Metric, dps pmetric.HistogramDataPointSlice) error {
	_ = "STUB: not implemented"
	return nil
}

func (d dataPointStatements) handleExponentialHistogramDataPoints(ctx context.Context, resourceMetrics pmetric.ResourceMetrics, scopeMetrics pmetric.ScopeMetrics, metric pmetric.Metric, dps pmetric.ExponentialHistogramDataPointSlice) error {
	_ = "STUB: not implemented"
	return nil
}

func (d dataPointStatements) handleSummaryDataPoints(ctx context.Context, resourceMetrics pmetric.ResourceMetrics, scopeMetrics pmetric.ScopeMetrics, metric pmetric.Metric, dps pmetric.SummaryDataPointSlice) error {
	_ = "STUB: not implemented"
	return nil
}

type MetricParserCollection ottl.ParserCollection[MetricsConsumer]

type MetricParserCollectionOption ottl.ParserCollectionOption[MetricsConsumer]

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

func NewMetricParserCollection(settings component.TelemetrySettings, options ...MetricParserCollectionOption) (*MetricParserCollection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertMetricStatements(pc *ottl.ParserCollection[MetricsConsumer], statements ottl.StatementsGetter, parsedStatements []*ottl.Statement[*ottlmetric.TransformContext]) (MetricsConsumer, error) {
	_ = "STUB: not implemented"
	return *new(MetricsConsumer), nil
}

func convertDataPointStatements(pc *ottl.ParserCollection[MetricsConsumer], statements ottl.StatementsGetter, parsedStatements []*ottl.Statement[*ottldatapoint.TransformContext]) (MetricsConsumer, error) {
	_ = "STUB: not implemented"
	return *new(MetricsConsumer), nil
}

func (mpc *MetricParserCollection) ParseContextStatements(contextStatements ContextStatements) (MetricsConsumer, error) {
	_ = "STUB: not implemented"
	return *new(MetricsConsumer), nil
}
