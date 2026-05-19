// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package filter // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/matcher/internal/filter"

import (
	"time"
)

type parseFunc func(string) (any, error)

type compareFunc func(a, b any) bool

type regexSortOption struct {
	regexKey string
	parseFunc
	compareFunc
}

func newRegexSortOption(regexKey string, parseFunc parseFunc, compareFunc compareFunc) (Option, error) {
	_ = "STUB: not implemented"
	return *new(Option), nil
}

func (o regexSortOption) apply(items []*item) ([]*item, error) {
	_ = "STUB: not implemented"
	// Special case where sort.Slice will not run the 'less' func.
	// We still need to ensure it parses in order to ensure the file should be included.
	return nil, nil
}

// Parse both values before checking for errors

// Sort i to the top of the slice

// Sort i to top of the slice

// Sort j to top of the slice

// If there were errors, they are at the top of the slice.

// No more errors, return the good items

// All items errored, clear the slice

func SortNumeric(regexKey string, ascending bool) (Option, error) {
	_ = "STUB: not implemented"
	return *new(Option), nil
}

func SortAlphabetical(regexKey string, ascending bool) (Option, error) {
	_ = "STUB: not implemented"
	return *new(Option), nil
}

func SortTemporal(regexKey string, ascending bool, layout, location string) (Option, error) {
	_ = "STUB: not implemented"
	return *new(Option), nil
}

type TopNOption int

//nolint:unparam
func (t TopNOption) apply(items []*item) ([]*item, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type mtimeSortOption struct {
	ascending bool
}

type mtimeItem struct {
	mtime time.Time
	path  string
	item  *item
}

func (m mtimeSortOption) apply(items []*item) ([]*item, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This checks if item i < j

// This checks if item i > j, in order to reverse the sort (most recently modified file is first in the list)

func SortMtime(ascending bool) Option { _ = "STUB: not implemented"; return *new(Option) }
