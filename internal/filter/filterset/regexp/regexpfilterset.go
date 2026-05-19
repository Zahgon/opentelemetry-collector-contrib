// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package regexp // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filterset/regexp"

import (
	"regexp"

	lru "github.com/hashicorp/golang-lru/v2"
)

// FilterSet encapsulates a set of filters and caches match results.
// Filters are re2 regex strings.
// FilterSet is exported for convenience, but has unexported fields and should be constructed through NewFilterSet.
//
// FilterSet satisfies the FilterSet interface from
// "go.opentelemetry.io/collector/internal/processor/filterset"
type FilterSet struct {
	regexes []*regexp.Regexp
	cache   *lru.Cache[string, bool]
}

// NewFilterSet constructs a FilterSet of re2 regex strings.
// If any of the given filters fail to compile into re2, an error is returned.
func NewFilterSet(filters []string, cfg *Config) (*FilterSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Because of legacy behavior, CacheMaxNumEntries == 0 means unbounded cache.

// Matches returns true if the given string matches any of the FilterSet's filters.
// The given string must be fully matched by at least one filter's re2 regex.
func (rfs *FilterSet) Matches(toMatch string) bool { _ = "STUB: not implemented"; return false }

// addFilters compiles all the given filters and stores them as regexes.
func (rfs *FilterSet) addFilters(filters []string) error { _ = "STUB: not implemented"; return nil }
