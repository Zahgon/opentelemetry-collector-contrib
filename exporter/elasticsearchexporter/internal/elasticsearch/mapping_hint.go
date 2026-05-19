// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package elasticsearch // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/elasticsearch"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

const (
	MappingHintsAttrKey = "elasticsearch.mapping.hints"
)

type MappingHint string

const (
	HintAggregateMetricDouble MappingHint = "aggregate_metric_double"
	HintDocCount              MappingHint = "_doc_count"
	HintHistogramRaw          MappingHint = "histogram:raw"
	// HintNoIndex signals that a document should not be indexed.
	// When present on a signal's attributes, the exporter skips bulk-emission
	// of the resulting document while leaving the signal available to any
	// upstream connector/processor that may have consumed it.
	//
	// Experimental: This hint is experimental and may change or be removed in future releases.
	HintNoIndex MappingHint = "_noindex"
)

type MappingHintGetter struct {
	hints []MappingHint
}

// NewMappingHintGetter creates a new MappingHintGetter
func NewMappingHintGetter(attr pcommon.Map) (g MappingHintGetter) {
	_ = "STUB: not implemented"
	return *new(MappingHintGetter)
}

// HasMappingHint checks whether the getter contains the requested mapping hint
func (g MappingHintGetter) HasMappingHint(hint MappingHint) bool {
	_ = "STUB: not implemented"
	return false
}
