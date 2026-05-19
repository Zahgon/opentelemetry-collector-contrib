// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metadata // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/metadata"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type newLabelValueFunction func(m LabelValueMetadata, value any) LabelValue

type LabelValueMetadata interface {
	ValueMetadata
	ValueType() ValueType
	NewLabelValue(value any) LabelValue
}

type LabelValue interface {
	Metadata() LabelValueMetadata
	Value() any
	SetValueTo(attributes pcommon.Map)
}

type queryLabelValueMetadata struct {
	name              string
	columnName        string
	valueType         ValueType
	newLabelValueFunc newLabelValueFunction
	valueHolderFunc   valueHolderFunction
}

func (m queryLabelValueMetadata) ValueHolder() any { _ = "STUB: not implemented"; return *new(any) }

func (m queryLabelValueMetadata) NewLabelValue(value any) LabelValue {
	_ = "STUB: not implemented"
	return *new(LabelValue)
}

func (m queryLabelValueMetadata) ValueType() ValueType {
	_ = "STUB: not implemented"
	return *new(ValueType)
}

type stringLabelValue struct {
	metadata LabelValueMetadata
	value    string
}

type int64LabelValue struct {
	metadata LabelValueMetadata
	value    int64
}

type boolLabelValue struct {
	metadata LabelValueMetadata
	value    bool
}

type stringSliceLabelValue struct {
	metadata LabelValueMetadata
	value    string
}

type byteSliceLabelValue struct {
	metadata LabelValueMetadata
	value    string
}

type lockRequestSliceLabelValue struct {
	metadata LabelValueMetadata
	value    string
}

func (m queryLabelValueMetadata) Name() string { _ = "STUB: not implemented"; return "" }

func (m queryLabelValueMetadata) ColumnName() string { _ = "STUB: not implemented"; return "" }

func (v stringLabelValue) Metadata() LabelValueMetadata {
	_ = "STUB: not implemented"
	return *new(LabelValueMetadata)
}

func (v stringLabelValue) Value() any { _ = "STUB: not implemented"; return *new(any) }

func (v stringLabelValue) SetValueTo(attributes pcommon.Map) { _ = "STUB: not implemented"; return }

func newStringLabelValue(metadata LabelValueMetadata, valueHolder any) LabelValue {
	_ = "STUB: not implemented"
	return *new(LabelValue)
}

func (v int64LabelValue) Metadata() LabelValueMetadata {
	_ = "STUB: not implemented"
	return *new(LabelValueMetadata)
}

func (v int64LabelValue) Value() any { _ = "STUB: not implemented"; return *new(any) }

func (v int64LabelValue) SetValueTo(attributes pcommon.Map) { _ = "STUB: not implemented"; return }

func newInt64LabelValue(metadata LabelValueMetadata, valueHolder any) LabelValue {
	_ = "STUB: not implemented"
	return *new(LabelValue)
}

func (v boolLabelValue) Metadata() LabelValueMetadata {
	_ = "STUB: not implemented"
	return *new(LabelValueMetadata)
}

func (v boolLabelValue) Value() any { _ = "STUB: not implemented"; return *new(any) }

func (v boolLabelValue) SetValueTo(attributes pcommon.Map) { _ = "STUB: not implemented"; return }

func newBoolLabelValue(metadata LabelValueMetadata, valueHolder any) LabelValue {
	_ = "STUB: not implemented"
	return *new(LabelValue)
}

func (v stringSliceLabelValue) Metadata() LabelValueMetadata {
	_ = "STUB: not implemented"
	return *new(LabelValueMetadata)
}

func (v stringSliceLabelValue) Value() any { _ = "STUB: not implemented"; return *new(any) }

func (v stringSliceLabelValue) SetValueTo(attributes pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func newStringSliceLabelValue(metadata LabelValueMetadata, valueHolder any) LabelValue {
	_ = "STUB: not implemented"
	return *new(LabelValue)
}

func (v byteSliceLabelValue) Metadata() LabelValueMetadata {
	_ = "STUB: not implemented"
	return *new(LabelValueMetadata)
}

func (v byteSliceLabelValue) Value() any { _ = "STUB: not implemented"; return *new(any) }

func (v byteSliceLabelValue) SetValueTo(attributes pcommon.Map) { _ = "STUB: not implemented"; return }

func (v *byteSliceLabelValue) ModifyValue(s string) { _ = "STUB: not implemented"; return }

func (v *stringSliceLabelValue) ModifyValue(s string) { _ = "STUB: not implemented"; return }

func (v *stringLabelValue) ModifyValue(s string) { _ = "STUB: not implemented"; return }

func newByteSliceLabelValue(metadata LabelValueMetadata, valueHolder any) LabelValue {
	_ = "STUB: not implemented"
	return *new(LabelValue)
}

func (v lockRequestSliceLabelValue) Metadata() LabelValueMetadata {
	_ = "STUB: not implemented"
	return *new(LabelValueMetadata)
}

func (v lockRequestSliceLabelValue) Value() any { _ = "STUB: not implemented"; return *new(any) }

func (v lockRequestSliceLabelValue) SetValueTo(attributes pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

type lockRequest struct {
	LockMode       string `spanner:"lock_mode"`
	Column         string `spanner:"column"`
	TransactionTag string `spanner:"transaction_tag"`
}

func newLockRequestSliceLabelValue(metadata LabelValueMetadata, valueHolder any) LabelValue {
	_ = "STUB: not implemented"
	return *new(LabelValue)
}

// During the specifics of this label we need to take into account only distinct values

func NewLabelValueMetadata(name, columnName string, valueType ValueType) (LabelValueMetadata, error) {
	_ = "STUB: not implemented"
	return *new(LabelValueMetadata), nil
}
