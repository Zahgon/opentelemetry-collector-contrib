// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package spanpruningprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/spanpruningprocessor"

import (
	"time"

	"go.opentelemetry.io/collector/component"
)

// Config defines the configuration options for the span pruning processor
// and the rules used to identify and aggregate similar spans.
type Config struct {
	// GroupByAttributes lists attribute patterns used to decide which leaf spans
	// belong in the same aggregation group. Spans must share the span name and
	// have identical values for every matched attribute to be grouped. Patterns
	// accept glob syntax, for example:
	//   - "db.*" matches db.operation, db.name, db.statement, etc.
	//   - "http.request.*" matches http.request.method, http.request.header, etc.
	//   - "service" matches only the exact key "service"
	// Examples: ["db.*", "http.method"], ["rpc.*"].
	GroupByAttributes []string `mapstructure:"group_by_attributes"`

	// MinSpansToAggregate is the minimum number of similar spans required before
	// aggregation occurs. Groups smaller than this threshold are preserved.
	// Default: 5
	MinSpansToAggregate int `mapstructure:"min_spans_to_aggregate"`

	// MaxParentDepth bounds how many ancestor levels above the aggregated leaves
	// can also be aggregated. Use 0 to aggregate only leaves, -1 for unlimited
	// depth, or a positive integer to cap traversal.
	// Default: 1
	MaxParentDepth int `mapstructure:"max_parent_depth"`

	// AggregationAttributePrefix prefixes all aggregation-related attributes that
	// are added to summary spans.
	// Default: "aggregation."
	AggregationAttributePrefix string `mapstructure:"aggregation_attribute_prefix"`

	// AggregationHistogramBuckets lists cumulative histogram bucket upper bounds
	// for latency tracking on aggregated spans. Empty slice disables histograms.
	AggregationHistogramBuckets []time.Duration `mapstructure:"aggregation_histogram_buckets"`
}

var _ component.Config = (*Config)(nil)

// Validate checks if the processor configuration is valid
func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }

// Validate AggregationAttributePrefix

// Validate GroupByAttributes glob patterns

// Try to compile the same way processor.go does to catch invalid syntax early

// Validate histogram buckets
