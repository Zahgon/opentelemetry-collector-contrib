// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor/internal/metrics"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/exp/metrics/identity"
)

type Ident = identity.Metric

type Metric struct {
	res   pcommon.Resource
	scope pcommon.InstrumentationScope
	pmetric.Metric
}

func (m *Metric) Ident() Ident { _ = "STUB: not implemented"; return *new(Ident) }

func (m *Metric) Resource() pcommon.Resource {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource)
}

func (m *Metric) Scope() pcommon.InstrumentationScope {
	_ = "STUB: not implemented"
	return *new(pcommon.InstrumentationScope)
}

func From(res pcommon.Resource, scope pcommon.InstrumentationScope, metric pmetric.Metric) Metric {
	_ = "STUB: not implemented"
	return *new(Metric)
}

func (m Metric) AggregationTemporality() pmetric.AggregationTemporality {
	_ = "STUB: not implemented"
	return *new(pmetric.AggregationTemporality)
}

func (m Metric) Typed() Any {
	_ = "STUB: not implemented"
	//exhaustive:enforce
	return *new(Any)
}

var (
	_ Any = Sum{}
	_ Any = Gauge{}
	_ Any = ExpHistogram{}
	_ Any = Histogram{}
	_ Any = Summary{}
)

type Any interface {
	Len() int
	Ident() identity.Metric
	SetAggregationTemporality(pmetric.AggregationTemporality)
}

func (m Metric) Filter(ok func(id identity.Stream, dp any) bool) { _ = "STUB: not implemented"; return }
