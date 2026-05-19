// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sumconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/sumconnector"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

var noAttributes = [16]byte{}

func newSummer[K any](metricDefs map[string]metricDef[K]) *summer[K] {
	_ = "STUB: not implemented"
	return nil
}

type summer[K any] struct {
	metricDefs map[string]metricDef[K]
	sums       map[string]map[[16]byte]*attrSummer
	timestamp  time.Time
}

type attrSummer struct {
	attrs pcommon.Map
	sum   float64
}

func (c *summer[K]) update(ctx context.Context, attrs pcommon.Map, tCtx K) error {
	_ = "STUB: not implemented"
	return nil
}

// Get source attribute value

// Get attribute values to include otherwise use default value

// Missing necessary attributes

// Perform condition matching or not

func (c *summer[K]) increment(metricName string, sumVal float64, attrs pcommon.Map) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *summer[K]) appendMetricsTo(metricSlice pmetric.MetricSlice) {
	_ = "STUB: not implemented"
	return
}

// The delta value is always positive, so a value accumulated downstream is monotonic
