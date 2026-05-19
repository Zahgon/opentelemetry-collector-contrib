// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package intervalprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/intervalprocessor"

import (
	"context"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/exp/metrics/identity"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/intervalprocessor/internal/metrics"
)

var _ processor.Metrics = (*intervalProcessor)(nil)

type intervalProcessor struct {
	ctx    context.Context
	cancel context.CancelFunc
	logger *zap.Logger

	wg        sync.WaitGroup
	stateLock sync.Mutex

	md                 pmetric.Metrics
	rmLookup           map[identity.Resource]pmetric.ResourceMetrics
	smLookup           map[identity.Scope]pmetric.ScopeMetrics
	mLookup            map[identity.Metric]pmetric.Metric
	numberLookup       map[identity.Stream]pmetric.NumberDataPoint
	histogramLookup    map[identity.Stream]pmetric.HistogramDataPoint
	expHistogramLookup map[identity.Stream]pmetric.ExponentialHistogramDataPoint
	summaryLookup      map[identity.Stream]pmetric.SummaryDataPoint

	config *Config

	nextConsumer consumer.Metrics
}

func newProcessor(config *Config, log *zap.Logger, nextConsumer consumer.Metrics) *intervalProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (p *intervalProcessor) Start(_ context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Flush remaining buffered metrics before exiting.
// Use context.Background() since p.ctx is already cancelled.

func (p *intervalProcessor) Shutdown(_ context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (*intervalProcessor) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (p *intervalProcessor) ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if we care about this value

func aggregateDataPoints[DPS metrics.DataPointSlice[DP], DP metrics.DataPoint[DP]](dataPoints, mCloneDataPoints DPS, metricID identity.Metric, dpLookup map[identity.Stream]DP) {
	_ = "STUB: not implemented"
	return
}

// Check if the datapoint is newer

// Otherwise, we leave existing as-is

func (p *intervalProcessor) exportMetrics(ctx context.Context) { _ = "STUB: not implemented"; return }

// ConsumeMetrics() has prepared our own pmetric.Metrics instance ready for us to use
// Take it and clear replace it with a new empty one

// Clear all the lookup references

func (p *intervalProcessor) getOrCloneMetric(rm pmetric.ResourceMetrics, sm pmetric.ScopeMetrics, m pmetric.Metric) (pmetric.Metric, identity.Metric) {
	_ = "STUB: not implemented"
	// Find the ResourceMetrics
	return *new(pmetric.Metric), *new(identity.Metric)
}

// We need to clone it *without* the ScopeMetricsSlice data

// Find the ScopeMetrics

// We need to clone it *without* the MetricSlice data

// Find the Metric

// We need to clone it *without* the datapoint data
