// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metricstestutil // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/metricstestutil"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// MetricDiff is intended to support producing human-readable diffs between two MetricData structs during
// testing. Two MetricDatas, when compared, could produce a list of MetricDiffs containing all of their
// differences, which could be used to correct the differences between the expected and actual values.
type MetricDiff struct {
	ExpectedValue any
	ActualValue   any
	Msg           string
}

func (mf MetricDiff) String() string { _ = "STUB: not implemented"; return "" }

func DiffMetrics(diffs []*MetricDiff, expected, actual pmetric.Metrics) []*MetricDiff {
	_ = "STUB: not implemented"
	return nil
}

func diffRMSlices(sent, recd []pmetric.ResourceMetrics) []*MetricDiff {
	_ = "STUB: not implemented"
	return nil
}

func diffRMs(diffs []*MetricDiff, expected, actual pmetric.ResourceMetrics) []*MetricDiff {
	_ = "STUB: not implemented"
	return nil
}

func diffILMSlice(
	diffs []*MetricDiff,
	expected pmetric.ScopeMetricsSlice,
	actual pmetric.ScopeMetricsSlice,
) []*MetricDiff {
	_ = "STUB: not implemented"
	return nil
}

func diffILM(
	diffs []*MetricDiff,
	expected pmetric.ScopeMetrics,
	actual pmetric.ScopeMetrics,
) []*MetricDiff {
	_ = "STUB: not implemented"
	return nil
}

func diffMetrics(diffs []*MetricDiff, expected, actual pmetric.MetricSlice) []*MetricDiff {
	_ = "STUB: not implemented"
	return nil
}

func diffMetricData(expected, actual pmetric.Metrics) []*MetricDiff {
	_ = "STUB: not implemented"
	return nil
}

func toSlice(s pmetric.ResourceMetricsSlice) (out []pmetric.ResourceMetrics) {
	_ = "STUB: not implemented"
	return nil
}

func DiffMetric(diffs []*MetricDiff, expected, actual pmetric.Metric) []*MetricDiff {
	_ = "STUB: not implemented"
	return nil
}

//exhaustive:enforce

// Note: Summary data points are not currently handled

func diffMetricDescriptor(
	diffs []*MetricDiff,
	expected pmetric.Metric,
	actual pmetric.Metric,
) ([]*MetricDiff, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func diffNumberPts(
	diffs []*MetricDiff,
	expected pmetric.NumberDataPointSlice,
	actual pmetric.NumberDataPointSlice,
) []*MetricDiff {
	_ = "STUB: not implemented"
	return nil
}

func diffHistogramPts(
	diffs []*MetricDiff,
	expected pmetric.HistogramDataPointSlice,
	actual pmetric.HistogramDataPointSlice,
) []*MetricDiff {
	_ = "STUB: not implemented"
	return nil
}

func diffHistogramPt(
	diffs []*MetricDiff,
	expected pmetric.HistogramDataPoint,
	actual pmetric.HistogramDataPoint,
) []*MetricDiff {
	_ = "STUB: not implemented"
	return nil
}

// TODO: HasSum, Min, HasMin, Max, HasMax are not covered in tests.

func diffExponentialHistogramPts(
	diffs []*MetricDiff,
	expected pmetric.ExponentialHistogramDataPointSlice,
	actual pmetric.ExponentialHistogramDataPointSlice,
) []*MetricDiff {
	_ = "STUB: not implemented"
	return nil
}

func diffExponentialHistogramPt(
	diffs []*MetricDiff,
	expected pmetric.ExponentialHistogramDataPoint,
	actual pmetric.ExponentialHistogramDataPoint,
) []*MetricDiff {
	_ = "STUB: not implemented"
	return nil
}

func diffExponentialHistogramPtBuckets(
	diffs []*MetricDiff,
	expected pmetric.ExponentialHistogramDataPointBuckets,
	actual pmetric.ExponentialHistogramDataPointBuckets,
) []*MetricDiff {
	_ = "STUB: not implemented"
	return nil
}

func diffExemplars(
	diffs []*MetricDiff,
	expected pmetric.ExemplarSlice,
	actual pmetric.ExemplarSlice,
) []*MetricDiff {
	_ = "STUB: not implemented"
	return nil
}

func diffResource(diffs []*MetricDiff, expected, actual pcommon.Resource) []*MetricDiff {
	_ = "STUB: not implemented"
	return nil
}

func diffResourceAttrs(diffs []*MetricDiff, expected, actual pcommon.Map) []*MetricDiff {
	_ = "STUB: not implemented"
	return nil
}

func diffMetricAttrs(diffs []*MetricDiff, expected, actual pcommon.Map) []*MetricDiff {
	_ = "STUB: not implemented"
	return nil
}

func diff(diffs []*MetricDiff, expected, actual any, msg string) []*MetricDiff {
	_ = "STUB: not implemented"
	return nil
}

func diffValues(
	diffs []*MetricDiff,
	expected any,
	actual any,
	msg string,
) ([]*MetricDiff, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func attrMapToString(m pcommon.Map) string { _ = "STUB: not implemented"; return "" }
