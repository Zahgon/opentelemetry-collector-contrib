// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package filter // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/matcher/internal/filter"

import (
	"regexp"
)

type Option interface {
	// Returned error is for explanatory purposes only.
	// All options will be called regardless of error.
	apply([]*item) ([]*item, error)
}

func Filter(values []string, regex *regexp.Regexp, opts ...Option) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type item struct {
	value    string
	captures map[string]string

	// Used when an Option is unable to interpret the value.
	// For example, a numeric sort may fail to parse the value as a number.
	err error
}

func newItem(value string, regex *regexp.Regexp) (*item, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
