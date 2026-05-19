// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package scrub contains a Scrubber that scrubs error from sensitive details
package scrub // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/datadog/scrub"

import (
	"regexp"
)

// Scrubber scrubs error from sensitive details.
type Scrubber interface {
	// Scrub sensitive data from an error.
	Scrub(error) error
}

// replacer structure to store regex matching and replacement functions.
type replacer struct {
	Regex *regexp.Regexp
	Repl  string
}

var _ error = (*scrubbedError)(nil)

// scrubbedError wraps an error and scrubs its `Error()` output.
type scrubbedError struct {
	err      error
	scrubbed string
}

func (s *scrubbedError) Error() string { _ = "STUB: not implemented"; return "" }

func (s *scrubbedError) Unwrap() error { _ = "STUB: not implemented"; return nil }

var _ Scrubber = (*scrubber)(nil)

// scrubber scrubs sensitive information from logs
type scrubber struct {
	replacers []replacer
}

func NewScrubber() Scrubber { _ = "STUB: not implemented"; return *new(Scrubber) }

// API key as URL parameter (api_key=<API KEY> or apikey=<API KEY>).
// Any alphanumeric string gets censored, even if not 32 characters long.

// Application key as URL parameter (api_key=<API KEY> or apikey=<API KEY>).
// Any alphanumeric string gets censored, even if not 40 characters long.

// API key in any place (32 character long alphanumeric ASCII string).

// Application key in any place (40 character long alphanumeric ASCII string).

func (s *scrubber) Scrub(err error) error { _ = "STUB: not implemented"; return nil }

// Scrub sensitive details from a string.
func (s *scrubber) scrubStr(data string) string { _ = "STUB: not implemented"; return "" }
