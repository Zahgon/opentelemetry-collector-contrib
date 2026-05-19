// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package filterprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/expr"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filterconfig"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filtermatcher"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottldatapoint"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlmetric"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlresource"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor/internal/condition"
)

type filterMetricProcessor struct {
	consumers         []condition.MetricsConsumer
	skipResourceExpr  expr.BoolExpr[*ottlresource.TransformContext]
	skipMetricExpr    expr.BoolExpr[*ottlmetric.TransformContext]
	skipDataPointExpr expr.BoolExpr[*ottldatapoint.TransformContext]
	telemetry         *filterTelemetry
	logger            *zap.Logger
}

func newFilterMetricProcessor(set processor.Settings, cfg *Config) (*filterMetricProcessor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// processMetrics filters the given metrics based off the filterMetricProcessor's filters.
func (fmp *filterMetricProcessor) processMetrics(ctx context.Context, md pmetric.Metrics) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

func (fmp *filterMetricProcessor) processConditions(ctx context.Context, md pmetric.Metrics) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

func (fmp *filterMetricProcessor) processSkipExpression(ctx context.Context, md pmetric.Metrics) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

//exhaustive:enforce

func newSkipResExpr(include, exclude *filterconfig.MetricMatchProperties) (expr.BoolExpr[*ottlresource.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type resExpr filtermatcher.AttributesMatcher

func (r resExpr) Eval(_ context.Context, tCtx *ottlresource.TransformContext) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func newResExpr(mp *filterconfig.MetricMatchProperties) (expr.BoolExpr[*ottlresource.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (fmp *filterMetricProcessor) handleNumberDataPoints(ctx context.Context, rm pmetric.ResourceMetrics, sm pmetric.ScopeMetrics, m pmetric.Metric, dps pmetric.NumberDataPointSlice) error {
	_ = "STUB: not implemented"
	return nil
}

func (fmp *filterMetricProcessor) handleHistogramDataPoints(ctx context.Context, rm pmetric.ResourceMetrics, sm pmetric.ScopeMetrics, m pmetric.Metric, dps pmetric.HistogramDataPointSlice) error {
	_ = "STUB: not implemented"
	return nil
}

func (fmp *filterMetricProcessor) handleExponentialHistogramDataPoints(ctx context.Context, rm pmetric.ResourceMetrics, sm pmetric.ScopeMetrics, m pmetric.Metric, dps pmetric.ExponentialHistogramDataPointSlice) error {
	_ = "STUB: not implemented"
	return nil
}

func (fmp *filterMetricProcessor) handleSummaryDataPoints(ctx context.Context, rm pmetric.ResourceMetrics, sm pmetric.ScopeMetrics, m pmetric.Metric, dps pmetric.SummaryDataPointSlice) error {
	_ = "STUB: not implemented"
	return nil
}
