// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package countconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/countconnector"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

var noAttributes = [16]byte{}

func newCounter[K any](metricDefs map[string]metricDef[K]) *counter[K] {
	_ = "STUB: not implemented"
	return nil
}

type counter[K any] struct {
	metricDefs map[string]metricDef[K]
	counts     map[string]map[[16]byte]*attrCounter
	startTime  pcommon.Timestamp
	endTime    pcommon.Timestamp
}

type attrCounter struct {
	attrs pcommon.Map
	count uint64
}

func (c *counter[K]) update(ctx context.Context, attrs, scopeAttrs, resourceAttrs pcommon.Map, tCtx K) error {
	_ = "STUB: not implemented"
	return nil
}

// Missing necessary attributes to be counted

// No conditions, so match all.

// updateTimestamp updates the start and end timestamps based on the provided timestamp
func (c *counter[K]) updateTimestamp(timestamp pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

// getTimestamps either gets the valid start and end timestamps or returns the current time
func (c *counter[K]) getTimestamps() (pcommon.Timestamp, pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp), *new(pcommon.Timestamp)
}

func (c *counter[K]) increment(metricName string, attrs pcommon.Map) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *counter[K]) appendMetricsTo(metricSlice pmetric.MetricSlice) {
	_ = "STUB: not implemented"
	return
}

// The delta value is always positive, so a value accumulated downstream is monotonic
