// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/exp/metrics"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/exp/metrics/identity"
)

// Merge will merge the metrics data in mdB into mdA, then return mdA.
// mdB will not be modified. The function will attempt to merge the data in mdB into
// existing ResourceMetrics / ScopeMetrics / Metrics in mdA if possible. If they don't
// exist, new entries will be created as needed.
//
// NOTE: Any "unnecessary" duplicate entries in mdA will *not* be combined. For example if
// mdA contains two ResourcMetric entries with identical Resource values, they will not be
// combined. If you wish to have this behavior, you could call this function twice:
//
//	cleanedMetrics := Merge(pmetric.NewMetrics(), mdA)
//	Merge(cleanedMetrics, mdB)
//
// That said, this will do a large amount of memory copying
func Merge(mdA, mdB pmetric.Metrics) pmetric.Metrics {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics)
}

// We didn't find a match
// Add it to mdA

func mergeResourceMetrics(resourceID identity.Resource, rmA, rmB pmetric.ResourceMetrics) pmetric.ResourceMetrics {
	_ = "STUB: not implemented"
	return *new(pmetric.ResourceMetrics)
}

// We didn't find a match
// Add it to rmA

func mergeScopeMetrics(scopeID identity.Scope, smA, smB pmetric.ScopeMetrics) pmetric.ScopeMetrics {
	_ = "STUB: not implemented"
	return *new(pmetric.ScopeMetrics)
}

//exhaustive:enforce

// We didn't find a match
// Add it to smA

func mergeDataPoints[DPS dataPointSlice[DP], DP dataPoint[DP]](dataPointsA, dataPointsB DPS) DPS {
	_ = "STUB: not implemented"
	// Append all the datapoints from B to A
	return *new(DPS)
}

type dataPointSlice[DP dataPoint[DP]] interface {
	Len() int
	At(i int) DP
	AppendEmpty() DP
}

type dataPoint[Self any] interface {
	Timestamp() pcommon.Timestamp
	Attributes() pcommon.Map
	CopyTo(dest Self)
}
