// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metricgroup // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/metricgroup"

import (
	"hash"

	"go.opentelemetry.io/collector/pdata/pcommon"
)

// mapHashSortedExcludeReservedAttrs is mapHash but ignoring some reserved attributes and is independent of order in Map.
// e.g. index is already considered during routing and DS attributes do not need to be considered in hashing
//
// TODO(carsonip): https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/39377
// Use opentelemetry-collector-contrib/pkg/pdatautil/hash.go when it can optionally exclude attributes
// We could have used it now but it'll involve creating a new Map and copying things over.
func mapHashSortedExcludeReservedAttrs(hasher hash.Hash, m pcommon.Map, extraExcludes ...string) {
	_ = "STUB: not implemented"
	return
}

func mapHash(hasher hash.Hash, m pcommon.Map) { _ = "STUB: not implemented"; return }

func valueHash(h hash.Hash, v pcommon.Value) { _ = "STUB: not implemented"; return }

func sliceHash(h hash.Hash, s pcommon.Slice) { _ = "STUB: not implemented"; return }
