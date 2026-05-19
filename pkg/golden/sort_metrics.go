// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package golden // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// sorts all Resource Metrics attributes and Datapoint Slice metric attributes and all Resource, Scope, and Datapoint Slices
func sortMetrics(ms pmetric.Metrics) { _ = "STUB: not implemented"; return }

//exhaustive:enforce

// sortAttributeMap sorts the attributes of a pcommon.Map according to the alphanumeric ordering of the keys
func sortAttributeMap(mp pcommon.Map) { _ = "STUB: not implemented"; return }

// sortMetricDataPointSlices sorts the datapoint slice of a pmetric.Metrics according to the alphanumeric ordering of map key
func sortMetricDataPointSlices(ms pmetric.Metrics) { _ = "STUB: not implemented"; return }

//exhaustive:enforce

func sortResources(ms pmetric.Metrics) { _ = "STUB: not implemented"; return }

func sortScopes(ms pmetric.Metrics) { _ = "STUB: not implemented"; return }

func sortNumberDataPointSlice(ndps pmetric.NumberDataPointSlice) { _ = "STUB: not implemented"; return }

func sortSummaryDataPointSlice(sdps pmetric.SummaryDataPointSlice) {
	_ = "STUB: not implemented"
	return
}

func sortHistogramDataPointSlice(hdps pmetric.HistogramDataPointSlice) {
	_ = "STUB: not implemented"
	return
}

func sortExponentialHistogramDataPointSlice(ehdps pmetric.ExponentialHistogramDataPointSlice) {
	_ = "STUB: not implemented"
	return
}

func compareMaps(a, b pcommon.Map) int { _ = "STUB: not implemented"; return 0 }
