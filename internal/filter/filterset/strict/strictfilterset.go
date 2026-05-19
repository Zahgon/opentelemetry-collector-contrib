// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package strict // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filterset/strict"

// FilterSet encapsulates a set of exact string match filters.
// FilterSet is exported for convenience, but has unexported fields and should be constructed through NewFilterSet.
//
// regexpFilterSet satisfies the FilterSet interface from
// "go.opentelemetry.io/collector/internal/processor/filterset"
type FilterSet struct {
	filters map[string]struct{}
}

// NewFilterSet constructs a FilterSet of exact string matches.
func NewFilterSet(filters []string) *FilterSet { _ = "STUB: not implemented"; return nil }

// Matches returns true if the given string matches any of the FilterSet's filters.
func (sfs *FilterSet) Matches(toMatch string) bool { _ = "STUB: not implemented"; return false }
