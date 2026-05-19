// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package matcher // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/matcher"

import (
	"regexp"
	"time"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/matcher/internal/filter"
)

const (
	sortTypeNumeric      = "numeric"
	sortTypeTimestamp    = "timestamp"
	sortTypeAlphabetical = "alphabetical"
	sortTypeMtime        = "mtime"
)

const (
	defaultOrderingCriteriaTopN = 1
)

type Criteria struct {
	Include []string `mapstructure:"include,omitempty"`
	Exclude []string `mapstructure:"exclude,omitempty"`

	// ExcludeOlderThan allows excluding files whose modification time is older
	// than the specified age.
	ExcludeOlderThan time.Duration    `mapstructure:"exclude_older_than"`
	OrderingCriteria OrderingCriteria `mapstructure:"ordering_criteria,omitempty"`
}

type OrderingCriteria struct {
	Regex   string `mapstructure:"regex,omitempty"`
	TopN    int    `mapstructure:"top_n,omitempty"`
	SortBy  []Sort `mapstructure:"sort_by,omitempty"`
	GroupBy string `mapstructure:"group_by,omitempty"`
}

type Sort struct {
	SortType  string `mapstructure:"sort_type,omitempty"`
	RegexKey  string `mapstructure:"regex_key,omitempty"`
	Ascending bool   `mapstructure:"ascending,omitempty"`

	// Timestamp only
	Layout   string `mapstructure:"layout,omitempty"`
	Location string `mapstructure:"location,omitempty"`
}

func New(c Criteria) (*Matcher, error) { _ = "STUB: not implemented"; return nil, nil }

// orderingCriteriaNeedsRegex returns true if any of the sort options require a regex to be set.
func orderingCriteriaNeedsRegex(sorts []Sort) bool { _ = "STUB: not implemented"; return false }

type Matcher struct {
	include    []string
	exclude    []string
	regex      *regexp.Regexp
	filterOpts []filter.Option
	groupBy    *regexp.Regexp
}

// MatchFiles gets a list of paths given an array of glob patterns to include and exclude
func (m Matcher) MatchFiles() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }
