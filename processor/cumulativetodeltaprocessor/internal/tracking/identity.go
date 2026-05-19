// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package tracking // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/cumulativetodeltaprocessor/internal/tracking"

import (
	"bytes"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

type MetricIdentity struct {
	Resource               pcommon.Resource
	InstrumentationLibrary pcommon.InstrumentationScope
	MetricType             pmetric.MetricType
	MetricIsMonotonic      bool
	MetricName             string
	MetricUnit             string
	StartTimestamp         pcommon.Timestamp
	Attributes             pcommon.Map
	MetricValueType        pmetric.NumberDataPointValueType
}

const (
	A      = int32('A')
	SEP    = byte(0x1E)
	SEPSTR = string(SEP)
)

func (mi *MetricIdentity) Write(b *bytes.Buffer) { _ = "STUB: not implemented"; return }

func (mi *MetricIdentity) IsFloatVal() bool { _ = "STUB: not implemented"; return false }

func (mi *MetricIdentity) IsSupportedMetricType() bool { _ = "STUB: not implemented"; return false }
