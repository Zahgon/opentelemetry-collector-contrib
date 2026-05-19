// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ratelimit // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sentryexporter/internal/ratelimit"

import (
	"errors"
	"time"
)

const DefaultRetryAfter = 1 * time.Minute

var errInvalidRetryAfter = errors.New("invalid input")

// parseRetryAfter parses a string s as in the standard Retry-After HTTP header
// and returns a deadline until when requests are rate limited and therefore new
// requests should not be sent. The input may be either a date or a non-negative
// integer number of seconds.
//
// See https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Retry-After
//
// parseRetryAfter always returns a usable deadline, even in case of an error.
//
// This is the original rate limiting mechanism used by Sentry, superseeded by
// the X-Sentry-Rate-Limits response header.
func parseRetryAfter(s string, now time.Time) (Deadline, error) {
	_ = "STUB: not implemented"
	return *new(Deadline), nil
}
