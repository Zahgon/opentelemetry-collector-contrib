// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azuredataexplorerexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azuredataexplorerexporter"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

/*
A converter package that converts and marshals data to be written to ADX metrics tables
*/
const (
	hostkey = "host.name"
	// Indicates the sum that is used in both summary and in histogram
	sumsuffix = "sum"
	// Count used in summary , histogram and also in exponential histogram
	countsuffix = "count"
	// Indicates the sum that is used in both summary and in histogram
	sumdescription = "(Sum total of samples)"
	// Count used in summary , histogram and also in exponential histogram
	countdescription = "(Count of samples)"
)

// This is derived from the specification https://opentelemetry.io/docs/reference/specification/metrics/datamodel/
type adxMetric struct {
	Timestamp string // The timestamp of the occurrence. A metric is measured at a point of time. Formatted into string as RFC3339Nano
	// Including name, the Metric object is defined by the following properties:
	MetricName        string         // Name of the metric field
	MetricType        string         // The data point type (e.g. Sum, Gauge, Histogram ExponentialHistogram, Summary)
	MetricUnit        string         // The metric stream’s unit
	MetricDescription string         // The metric stream’s description
	MetricValue       float64        // the value of the metric
	MetricAttributes  map[string]any // JSON attributes that can then be parsed. Extrinsic properties
	// Additional properties
	Host               string         // The hostname for analysis of the metric. Extracted from https://opentelemetry.io/docs/reference/specification/resource/semantic_conventions/host/
	ResourceAttributes map[string]any // The originating Resource attributes. Refer https://opentelemetry.io/docs/reference/specification/resource/sdk/
}

/*
	Convert the pMetric to the type ADXMetric , this matches the scheme in the OTELMetric table in the database
*/

func mapToAdxMetric(res pcommon.Resource, md pmetric.Metric, scopeattrs map[string]any, logger *zap.Logger) []*adxMetric {
	_ = "STUB: not implemented"
	return nil
}

// default to collectors host name. Ignore the error here. This should not cause the failure of the process

//exhaustive:enforce

// first, add one event for sum, and one for count

// Change int to float. The value is a float64 in the table

// Spec says counts is optional but if present it must have one more
// element than the bounds array.

// now create buckets for each bound.

//nolint:errcheck

// Change int to float. The value is a float64 in the table

// add an upper bound for +Inf

// Add the LE field for the bucket's bound

//nolint:errcheck

// Change int to float. The value is a float64 in the table

// first, add one event for sum, and one for count

// counts

// now create values for each quantile.

//nolint:errcheck

// Given all the metrics , transform that to the representative structure
func rawMetricsToAdxMetrics(_ context.Context, metrics pmetric.Metrics, logger *zap.Logger) []*adxMetric {
	_ = "STUB: not implemented"
	return nil
}

// get details of the scope from the scope metric

func float64ToDimValue(f float64) string { _ = "STUB: not implemented"; return "" }

func sanitizeFloat(value float64) any { _ = "STUB: not implemented"; return *new(any) }
