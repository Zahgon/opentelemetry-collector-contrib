// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metadata // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/metadata"

import (
	"cloud.google.com/go/spanner"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

type newMetricValueFunction func(m MetricValueMetadata, value any) MetricValue

type MetricValueMetadata interface {
	ValueMetadata
	ValueType() ValueType
	DataType() MetricType
	Unit() string
	NewMetricValue(value any) MetricValue
}

type MetricValue interface {
	Metadata() MetricValueMetadata
	Value() any
	SetValueTo(ndp pmetric.NumberDataPoint)
}

type queryMetricValueMetadata struct {
	name               string
	columnName         string
	dataType           MetricType
	unit               string
	valueType          ValueType
	newMetricValueFunc newMetricValueFunction
	valueHolderFunc    valueHolderFunction
}

type int64MetricValue struct {
	metadata MetricValueMetadata
	value    int64
}

type float64MetricValue struct {
	metadata MetricValueMetadata
	value    float64
}

type nullFloat64MetricValue struct {
	metadata MetricValueMetadata
	value    spanner.NullFloat64
}

func (m queryMetricValueMetadata) ValueHolder() any { _ = "STUB: not implemented"; return *new(any) }

func (m queryMetricValueMetadata) NewMetricValue(value any) MetricValue {
	_ = "STUB: not implemented"
	return *new(MetricValue)
}

func (m queryMetricValueMetadata) Name() string { _ = "STUB: not implemented"; return "" }

func (m queryMetricValueMetadata) ColumnName() string { _ = "STUB: not implemented"; return "" }

func (m queryMetricValueMetadata) ValueType() ValueType {
	_ = "STUB: not implemented"
	return *new(ValueType)
}

func (m queryMetricValueMetadata) DataType() MetricType {
	_ = "STUB: not implemented"
	return *new(MetricType)
}

func (m queryMetricValueMetadata) Unit() string { _ = "STUB: not implemented"; return "" }

func (v int64MetricValue) Metadata() MetricValueMetadata {
	_ = "STUB: not implemented"
	return *new(MetricValueMetadata)
}

func (v float64MetricValue) Metadata() MetricValueMetadata {
	_ = "STUB: not implemented"
	return *new(MetricValueMetadata)
}

func (v nullFloat64MetricValue) Metadata() MetricValueMetadata {
	_ = "STUB: not implemented"
	return *new(MetricValueMetadata)
}

func (v int64MetricValue) Value() any { _ = "STUB: not implemented"; return *new(any) }

func (v float64MetricValue) Value() any { _ = "STUB: not implemented"; return *new(any) }

func (v nullFloat64MetricValue) Value() any { _ = "STUB: not implemented"; return *new(any) }

func (v int64MetricValue) SetValueTo(point pmetric.NumberDataPoint) {
	_ = "STUB: not implemented"
	return
}

func (v float64MetricValue) SetValueTo(point pmetric.NumberDataPoint) {
	_ = "STUB: not implemented"
	return
}

func (v nullFloat64MetricValue) SetValueTo(point pmetric.NumberDataPoint) {
	_ = "STUB: not implemented"
	return
}

func newInt64MetricValue(metadata MetricValueMetadata, valueHolder any) MetricValue {
	_ = "STUB: not implemented"
	return *new(MetricValue)
}

func newFloat64MetricValue(metadata MetricValueMetadata, valueHolder any) MetricValue {
	_ = "STUB: not implemented"
	return *new(MetricValue)
}

func newNullFloat64MetricValue(metadata MetricValueMetadata, valueHolder any) MetricValue {
	_ = "STUB: not implemented"
	return *new(MetricValue)
}

func NewMetricValueMetadata(name, columnName string, dataType MetricType, unit string,
	valueType ValueType,
) (MetricValueMetadata, error) {
	_ = "STUB: not implemented"
	return *new(MetricValueMetadata), nil
}
