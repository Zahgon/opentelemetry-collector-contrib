// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package identity // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/exp/metrics/identity"

import (
	"hash"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

type Metric struct {
	scope

	name string
	unit string
	ty   pmetric.MetricType

	monotonic   bool
	temporality pmetric.AggregationTemporality
}

func (m Metric) Hash() hash.Hash64 { _ = "STUB: not implemented"; return *new(hash.Hash64) }

func (m Metric) Scope() Scope { _ = "STUB: not implemented"; return *new(Scope) }

func OfMetric(scope Scope, m pmetric.Metric) Metric { _ = "STUB: not implemented"; return *new(Metric) }

func (m Metric) String() string { _ = "STUB: not implemented"; return "" }

func OfResourceMetric(res pcommon.Resource, scope pcommon.InstrumentationScope, metric pmetric.Metric) Metric {
	_ = "STUB: not implemented"
	return *new(Metric)
}
