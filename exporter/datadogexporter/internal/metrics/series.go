// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter/internal/metrics"

import (
	"github.com/DataDog/datadog-agent/pkg/opentelemetry-mapping-go/otlp/attributes"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"go.opentelemetry.io/collector/component"
)

// newMetricSeries creates a new Datadog metric series given a name, a Unix nanoseconds timestamp
// a value and a slice of tags
func newMetricSeries(name string, ts uint64, value float64, tags []string) datadogV2.MetricSeries {
	_ = "STUB: not implemented"
	// Transform UnixNano timestamp into Unix timestamp
	// 1 second = 1e9 ns
	return *new(datadogV2.MetricSeries)
}

// NewMetric creates a new DatadogV2 metric given a name, a type, a Unix nanoseconds timestamp
// a value and a slice of tags
func NewMetric(name string, dt datadogV2.MetricIntakeType, ts uint64, interval int64, value float64, tags []string) datadogV2.MetricSeries {
	_ = "STUB: not implemented"
	return *new(datadogV2.MetricSeries)
}

// NewGauge creates a new DatadogV2 Gauge metric given a name, a Unix nanoseconds timestamp
// a value and a slice of tags
func NewGauge(name string, ts uint64, interval int64, value float64, tags []string) datadogV2.MetricSeries {
	_ = "STUB: not implemented"
	return *new(datadogV2.MetricSeries)
}

// NewCount creates a new DatadogV2 count metric given a name, a Unix nanoseconds timestamp
// a value and a slice of tags
func NewCount(name string, ts uint64, interval int64, value float64, tags []string) datadogV2.MetricSeries {
	_ = "STUB: not implemented"
	return *new(datadogV2.MetricSeries)
}

// DefaultMetrics creates built-in metrics to report that an exporter is running
func DefaultMetrics(exporterType, hostname string, timestamp uint64, tags []string) []datadogV2.MetricSeries {
	_ = "STUB: not implemented"
	return nil
}

// GatewayUsageGauge creates a gauge metric to report if there is a gateway
func GatewayUsageGauge(timestamp uint64, hostname string, tags []string, gatewayUsage *attributes.GatewayUsage) datadogV2.MetricSeries {
	_ = "STUB: not implemented"
	return *new(datadogV2.MetricSeries)
}

// TagsFromBuildInfo returns a list of tags derived from buildInfo to be used when creating metrics
func TagsFromBuildInfo(buildInfo component.BuildInfo) []string {
	_ = "STUB: not implemented"
	return nil
}
