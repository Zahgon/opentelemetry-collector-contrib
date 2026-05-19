// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package aggregator // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/signaltometricsconnector/internal/aggregator"

import (
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// gaugeDP is a data point for gauge metrics.
type gaugeDP struct {
	attrs pcommon.Map
	val   any
}

func newGaugeDP(attrs pcommon.Map) *gaugeDP { _ = "STUB: not implemented"; return nil }

func (dp *gaugeDP) Aggregate(v any) { _ = "STUB: not implemented"; return }

// Copy copies the gauge data point to the destination number data point
func (dp *gaugeDP) Copy(
	timestamp time.Time,
	dest pmetric.NumberDataPoint,
) {
	_ = "STUB: not implemented"
	return
}

// TODO determine appropriate start time
