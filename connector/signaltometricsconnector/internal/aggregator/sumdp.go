// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package aggregator // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/signaltometricsconnector/internal/aggregator"

import (
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// sumDP counts the number of events (supports all event types)
type sumDP struct {
	attrs pcommon.Map

	isDbl  bool
	intVal int64
	dblVal float64
}

func newSumDP(attrs pcommon.Map, isDbl bool) *sumDP { _ = "STUB: not implemented"; return nil }

func (dp *sumDP) AggregateInt(v int64) { _ = "STUB: not implemented"; return }

func (dp *sumDP) AggregateDouble(v float64) { _ = "STUB: not implemented"; return }

func (dp *sumDP) Copy(
	timestamp time.Time,
	dest pmetric.NumberDataPoint,
) {
	_ = "STUB: not implemented"
	return
}

// TODO determine appropriate start time
