// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metricgroup // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/metricgroup"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/datapoints"
)

// HashKey is a struct for comparing data point identity.
type HashKey struct {
	resourceHash uint64
	scopeHash    uint64
	dpHash       uint64
}

// DataPointHasher is an interface for hashing data points by their identity,
// for grouping into a single document.
type DataPointHasher interface {
	UpdateResource(pcommon.Resource)
	UpdateScope(pcommon.InstrumentationScope)
	UpdateDataPoint(datapoints.DataPoint)
	HashKey() HashKey
}

type hashableDataPoint interface {
	Timestamp() pcommon.Timestamp
	StartTimestamp() pcommon.Timestamp
	Metric() pmetric.Metric
	Attributes() pcommon.Map
}

// ECSDataPointHasher caches resource and data point, and computes a hash on HashKey
// as data point attributes overwrite resource attributes in ECS mode because they are all stored at root level.
type ECSDataPointHasher struct {
	resource pcommon.Resource
	dp       hashableDataPoint
}

func (h *ECSDataPointHasher) UpdateResource(resource pcommon.Resource) {
	_ = "STUB: not implemented"
	return
}

func (*ECSDataPointHasher) UpdateScope(pcommon.InstrumentationScope) {
	_ = "STUB: not implemented"
	return
}

func (h *ECSDataPointHasher) UpdateDataPoint(dp datapoints.DataPoint) {
	_ = "STUB: not implemented"
	return
}

func (h *ECSDataPointHasher) HashKey() HashKey { _ = "STUB: not implemented"; return *new(HashKey) }

// scope attributes are ignored in ECS mode

// OTelDataPointHasher computes a hash for each of resource, scope and data point on each Update call,
// to avoid wasteful hashing and sorting on data point sharing the same resource and scope.
type OTelDataPointHasher struct {
	resourceHash uint64
	scopeHash    uint64
	dpHash       uint64
}

func (h *OTelDataPointHasher) UpdateResource(resource pcommon.Resource) {
	_ = "STUB: not implemented"
	// We cannot use exp/metrics/identity here because some resource fields e.g. schema url
	// are not dimensions and should not be part of the hash.
	return
}

// There is special handling to merge geo attributes during serialization,
// but we can hash them as if they are separate now.

func (h *OTelDataPointHasher) UpdateScope(scope pcommon.InstrumentationScope) {
	_ = "STUB: not implemented"
	return
}

// There is special handling to merge geo attributes during serialization,
// but we can hash them as if they are separate now.

func (h *OTelDataPointHasher) UpdateDataPoint(dp datapoints.DataPoint) {
	_ = "STUB: not implemented"
	return
}

// There is special handling to merge geo attributes during serialization,
// but we can hash them as if they are separate now.

func (h *OTelDataPointHasher) HashKey() HashKey { _ = "STUB: not implemented"; return *new(HashKey) }
