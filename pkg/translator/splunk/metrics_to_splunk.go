// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package splunk // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/splunk"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

const (
	// unknownHostName is the default host name when no hostname label is passed.
	unknownHostName = "unknown"
	// splunkMetricTypeKey is the key which maps to the type of the metric.
	splunkMetricTypeKey = "metric_type"
	// splunkMetricValue is the splunk metric value prefix.
	splunkMetricValue = "metric_name"
	// countSuffix is the count metric value suffix.
	countSuffix = "_count"
	// sumSuffix is the sum metric value suffix.
	sumSuffix = "_sum"
	// bucketSuffix is the bucket metric value suffix.
	bucketSuffix = "_bucket"
	// nanValue is the string representation of a NaN value in HEC events
	nanValue = "NaN"
	// plusInfValue is the string representation of a +Inf value in HEC events
	plusInfValue = "+Inf"
	// minusInfValue is the string representation of a -Inf value in HEC events
	minusInfValue = "-Inf"
)

func sanitizeFloat(value float64) any { _ = "STUB: not implemented"; return *new(any) }

func MetricToSplunkEvent(res pcommon.Resource, m pmetric.Metric, logger *zap.Logger, mapping HecToOtelAttrs, source, sourceType, index string) []*Event {
	_ = "STUB: not implemented"
	return nil
}

// ignore

//exhaustive:enforce

// first, add one event for sum, and one for count

// Spec says counts is optional but if present it must have one more
// element than the bounds array.

// now create buckets for each bound.

// add an upper bound for +Inf

// first, add one event for sum, and one for count

// now create values for each quantile.

func createEvent(timestamp pcommon.Timestamp, host, source, sourceType, index string, fields map[string]any) *Event {
	_ = "STUB: not implemented"
	return nil
}

func copyEventWithoutValues(event *Event) *Event { _ = "STUB: not implemented"; return nil }

func populateAttributes(fields map[string]any, attributeMap pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func cloneMap(fields map[string]any) map[string]any { _ = "STUB: not implemented"; return nil }

func cloneMapWithSelector(fields map[string]any, selector func(string) bool) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func float64ToDimValue(f float64) string { _ = "STUB: not implemented"; return "" }

// merge metric events to adhere to the multimetric format event.
func MergeEventsToMultiMetricFormat(events []*Event) ([]*Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
