// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package signalfx // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/signalfx"

import (
	"math"

	sfxpb "github.com/signalfx/com_signalfx_metrics_protobuf/model"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// Some fields on SignalFx protobuf are pointers, in order to reduce
// allocations create the most used ones.
var (
	// SignalFx metric types used in the conversions.
	sfxMetricTypeGauge             = sfxpb.MetricType_GAUGE
	sfxMetricTypeCumulativeCounter = sfxpb.MetricType_CUMULATIVE_COUNTER
	sfxMetricTypeCounter           = sfxpb.MetricType_COUNTER

	// infinity bound dimension value is used on all histograms.
	infinityBoundSFxDimValue = float64ToDimValue(math.Inf(1))
)

const (
	// prometheus compatible dimension key for histogram buckets.
	bucketDimensionKey = "le"

	// quantile dimension key for summary quantiles.
	quantileDimensionKey = "quantile"
)

// FromTranslator converts from pdata to SignalFx proto data model.
type FromTranslator struct{}

// FromMetrics converts pmetric.Metrics to SignalFx proto data points.
func (ft *FromTranslator) FromMetrics(md pmetric.Metrics, dropHistogramBuckets, processHistograms bool) ([]*sfxpb.DataPoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FromMetric converts pmetric.Metric to SignalFx proto data points.
// TODO: Remove this and change signalfxexporter to us FromMetrics.
func (*FromTranslator) FromMetric(m pmetric.Metric, extraDimensions []*sfxpb.Dimension, dropHistogramBuckets, processHistograms bool) []*sfxpb.DataPoint {
	_ = "STUB: not implemented"
	return nil
}

func fromMetricTypeToMetricType(metric pmetric.Metric) *sfxpb.MetricType {
	_ = "STUB: not implemented"
	return nil
}

func convertNumberDataPoints(in pmetric.NumberDataPointSlice, name string, mt *sfxpb.MetricType, extraDims []*sfxpb.Dimension) []*sfxpb.DataPoint {
	_ = "STUB: not implemented"
	return nil
}

func convertHistogram(in pmetric.HistogramDataPointSlice, name string, mt *sfxpb.MetricType, extraDims []*sfxpb.Dimension, dropHistogramBuckets bool) []*sfxpb.DataPoint {
	_ = "STUB: not implemented"
	return nil
}

// Min is always a gauge.

// Max is always a gauge.

// Drop Histogram Buckets if flag is set.

// Spec says counts is optional but if present it must have one more
// element than the bounds array.

func convertSummaryDataPoints(in pmetric.SummaryDataPointSlice, name string, extraDims []*sfxpb.Dimension) []*sfxpb.DataPoint {
	_ = "STUB: not implemented"
	return nil
}

func attributesToDimensions(attributes pcommon.Map, extraDims []*sfxpb.Dimension) []*sfxpb.Dimension {
	_ = "STUB: not implemented"
	return nil
}

type dpsBuilder struct {
	baseOut []sfxpb.DataPoint
	out     []*sfxpb.DataPoint
	pos     int
}

func newDpsBuilder(capacity int) dpsBuilder { _ = "STUB: not implemented"; return *new(dpsBuilder) }

func (dp *dpsBuilder) appendPoint(name string, mt *sfxpb.MetricType, ts int64, dims []*sfxpb.Dimension) *sfxpb.DataPoint {
	_ = "STUB: not implemented"
	return nil
}

// Is equivalent to strconv.FormatFloat(f, 'g', -1, 64), but hard-codes a few common cases for increased efficiency.
func float64ToDimValue(f float64) string {
	_ = "STUB: not implemented"
	// Parameters below are the same used by Prometheus
	// see https://github.com/prometheus/common/blob/b5fe7d854c42dc7842e48d1ca58f60feae09d77b/expfmt/text_create.go#L450
	// SignalFx agent uses a different pattern
	// https://github.com/signalfx/signalfx-agent/blob/5779a3de0c9861fa07316fd11b3c4ff38c0d78f0/internal/monitors/prometheusexporter/conversion.go#L77
	// The important issue here is consistency with the exporter, opting for the more common one used by Prometheus.
	return ""
}
