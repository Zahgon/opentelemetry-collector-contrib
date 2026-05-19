// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package cumulativetodeltaprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/cumulativetodeltaprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filterset"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/cumulativetodeltaprocessor/internal/tracking"
)

type cumulativeToDeltaProcessor struct {
	includeFS          filterset.FilterSet
	excludeFS          filterset.FilterSet
	includeMetricTypes map[pmetric.MetricType]bool
	excludeMetricTypes map[pmetric.MetricType]bool
	logger             *zap.Logger
	deltaCalculator    *tracking.MetricTracker
	cancelFunc         context.CancelFunc
}

func newCumulativeToDeltaProcessor(config *Config, logger *zap.Logger) (*cumulativeToDeltaProcessor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getMetricTypeFilter(types []string) (map[pmetric.MetricType]bool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// processMetrics implements the ProcessMetricsFunc type.
func (ctdp *cumulativeToDeltaProcessor) processMetrics(_ context.Context, md pmetric.Metrics) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// Ignore any metrics that aren't monotonic

func (ctdp *cumulativeToDeltaProcessor) shutdown(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (ctdp *cumulativeToDeltaProcessor) shouldConvertMetric(metric pmetric.Metric) bool {
	_ = "STUB: not implemented"
	return false
}

func (ctdp *cumulativeToDeltaProcessor) convertNumberDataPoints(dps pmetric.NumberDataPointSlice, baseIdentity tracking.MetricIdentity) {
	_ = "STUB: not implemented"
	return
}

// drop points with no value

// Do not attempt to transform NaN values

func (ctdp *cumulativeToDeltaProcessor) convertHistogramDataPoints(dps pmetric.HistogramDataPointSlice, baseIdentity tracking.MetricIdentity) {
	_ = "STUB: not implemented"
	return
}

// drop points with no value

func (ctdp *cumulativeToDeltaProcessor) convertExponentialHistogramDataPoints(dps pmetric.ExponentialHistogramDataPointSlice, baseIdentity tracking.MetricIdentity) {
	_ = "STUB: not implemented"
	return
}

// drop points with no value

// Scale and ZeroThreshold are unchanged

// Cannot consistently compute min/max
