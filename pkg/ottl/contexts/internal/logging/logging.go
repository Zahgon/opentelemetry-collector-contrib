// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package logging // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/logging"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap/zapcore"
)

type Slice pcommon.Slice

func (s Slice) MarshalLogArray(encoder zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type Map pcommon.Map

func (m Map) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type Resource pcommon.Resource

func (r Resource) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type InstrumentationScope pcommon.InstrumentationScope

func (i InstrumentationScope) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type Span ptrace.Span

func (s Span) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type SpanEventSlice ptrace.SpanEventSlice

func (s SpanEventSlice) MarshalLogArray(encoder zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type SpanEvent ptrace.SpanEvent

func (s SpanEvent) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type SpanLinkSlice ptrace.SpanLinkSlice

func (s SpanLinkSlice) MarshalLogArray(encoder zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type SpanLink ptrace.SpanLink

func (s SpanLink) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type Metric pmetric.Metric

func (m Metric) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type NumberDataPointSlice pmetric.NumberDataPointSlice

func (n NumberDataPointSlice) MarshalLogArray(encoder zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type NumberDataPoint pmetric.NumberDataPoint

func (n NumberDataPoint) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type HistogramDataPointSlice pmetric.HistogramDataPointSlice

func (h HistogramDataPointSlice) MarshalLogArray(encoder zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type HistogramDataPoint pmetric.HistogramDataPoint

func (h HistogramDataPoint) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type ExponentialHistogramDataPointSlice pmetric.ExponentialHistogramDataPointSlice

func (e ExponentialHistogramDataPointSlice) MarshalLogArray(encoder zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type ExponentialHistogramDataPoint pmetric.ExponentialHistogramDataPoint

func (e ExponentialHistogramDataPoint) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type ExponentialHistogramDataPointBuckets pmetric.ExponentialHistogramDataPointBuckets

func (e ExponentialHistogramDataPointBuckets) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type SummaryDataPointSlice pmetric.SummaryDataPointSlice

func (s SummaryDataPointSlice) MarshalLogArray(encoder zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type SummaryDataPoint pmetric.SummaryDataPoint

func (s SummaryDataPoint) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type SummaryDataPointValueAtQuantileSlice pmetric.SummaryDataPointValueAtQuantileSlice

func (s SummaryDataPointValueAtQuantileSlice) MarshalLogArray(encoder zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type SummaryDataPointValueAtQuantile pmetric.SummaryDataPointValueAtQuantile

func (s SummaryDataPointValueAtQuantile) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type UInt64Slice pcommon.UInt64Slice

func (u UInt64Slice) MarshalLogArray(encoder zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type Float64Slice pcommon.Float64Slice

func (f Float64Slice) MarshalLogArray(encoder zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type ExemplarSlice pmetric.ExemplarSlice

func (e ExemplarSlice) MarshalLogArray(encoder zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type Exemplar pmetric.Exemplar

func (e Exemplar) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}
