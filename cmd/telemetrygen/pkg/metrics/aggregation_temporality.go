// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics

import (
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
)

type AggregationTemporality metricdata.Temporality

func (t *AggregationTemporality) Set(v string) error { _ = "STUB: not implemented"; return nil }

func (t *AggregationTemporality) String() string { _ = "STUB: not implemented"; return "" }

func (*AggregationTemporality) Type() string { _ = "STUB: not implemented"; return "" }

// AsTemporality converts the AggregationTemporality to metricdata.Temporality
func (t AggregationTemporality) AsTemporality() metricdata.Temporality {
	_ = "STUB: not implemented"
	return *new(metricdata.Temporality)
}
