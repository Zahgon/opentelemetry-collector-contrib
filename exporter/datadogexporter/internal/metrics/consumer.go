// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter/internal/metrics"

import (
	"context"

	"github.com/DataDog/datadog-agent/pkg/opentelemetry-mapping-go/otlp/attributes"
	"github.com/DataDog/datadog-agent/pkg/opentelemetry-mapping-go/otlp/metrics"
	"github.com/DataDog/datadog-agent/pkg/util/quantile"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter/internal/metrics/sketches"
)

var (
	_ metrics.Consumer     = (*Consumer)(nil)
	_ metrics.HostConsumer = (*Consumer)(nil)
	_ metrics.TagsConsumer = (*Consumer)(nil)
)

// Consumer implements metrics.Consumer. It records consumed metrics, sketches and
// APM stats payloads. It provides them to the caller using the All method.
type Consumer struct {
	ms           []datadogV2.MetricSeries
	sl           sketches.SketchSeriesList
	seenHosts    map[string]struct{}
	seenTags     map[string]struct{}
	gatewayUsage *attributes.GatewayUsage
}

// NewConsumer creates a new Datadog consumer. It implements metrics.Consumer.
func NewConsumer(gatewayUsage *attributes.GatewayUsage) *Consumer {
	_ = "STUB: not implemented"
	return nil
}

// toDataType maps translator datatypes to DatadogV2's datatypes.
func (*Consumer) toDataType(dt metrics.DataType) (out datadogV2.MetricIntakeType) {
	_ = "STUB: not implemented"
	return *new(datadogV2.MetricIntakeType)
}

// runningMetrics gets the running metrics for the exporter.
func (c *Consumer) runningMetrics(timestamp uint64, buildInfo component.BuildInfo, metadata metrics.Metadata) (series []datadogV2.MetricSeries) {
	_ = "STUB: not implemented"
	return nil
}

// Report the host as running

//nolint:gocritic

// All gets all metrics (consumed metrics and running metrics).
func (c *Consumer) All(timestamp uint64, buildInfo component.BuildInfo, tags []string, metadata metrics.Metadata) ([]datadogV2.MetricSeries, sketches.SketchSeriesList) {
	_ = "STUB: not implemented"
	return nil, *new(sketches.SketchSeriesList)
}

// ConsumeTimeSeries implements the metrics.Consumer interface.
func (c *Consumer) ConsumeTimeSeries(
	_ context.Context,
	dims *metrics.Dimensions,
	typ metrics.DataType,
	timestamp uint64,
	interval int64,
	value float64,
) {
	_ = "STUB: not implemented"
	return
}

// ConsumeSketch implements the metrics.Consumer interface.
func (c *Consumer) ConsumeSketch(
	_ context.Context,
	dims *metrics.Dimensions,
	timestamp uint64,
	interval int64,
	sketch *quantile.Sketch,
) {
	_ = "STUB: not implemented"
	return
}

// ConsumeHost implements the metrics.HostConsumer interface.
func (c *Consumer) ConsumeHost(host string) { _ = "STUB: not implemented"; return }

// ConsumeTag implements the metrics.TagsConsumer interface.
func (c *Consumer) ConsumeTag(tag string) { _ = "STUB: not implemented"; return }

// ConsumeExplicitBoundHistogram implements the metrics.ExplicitBoundHistogramConsumer interface.
// This is a no-op implementation as we use sketch-based histograms.
func (*Consumer) ConsumeExplicitBoundHistogram(
	_ context.Context,
	_ *metrics.Dimensions,
	_ pmetric.HistogramDataPointSlice,
) {
	_ = "STUB: not implemented"
	// No-op: we use sketch-based histograms
	return
}

// ConsumeExponentialHistogram implements the metrics.ExponentialHistogramConsumer interface.
// This is a no-op implementation as we use sketch-based histograms.
func (*Consumer) ConsumeExponentialHistogram(
	_ context.Context,
	_ *metrics.Dimensions,
	_ pmetric.ExponentialHistogramDataPointSlice,
) {
	_ = "STUB: not implemented"
	// No-op: we use sketch-based histograms
	return
}
