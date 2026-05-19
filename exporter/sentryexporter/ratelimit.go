// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sentryexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sentryexporter"

import (
	"errors"
	"net/http"
	"sync"
	"time"

	internalrl "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sentryexporter/internal/ratelimit"
)

var errRateLimited = errors.New("error sending data, rate-limited")

type dataCategory string

const (
	dataCategoryTrace dataCategory = "trace"
	dataCategoryLog   dataCategory = "log"
)

// rateLimiter keeps track of rate limit maps per DSN.
type rateLimiter struct {
	mu        sync.Mutex
	dsnLimits map[string]internalrl.Map
}

func newRateLimiter() *rateLimiter { _ = "STUB: not implemented"; return nil }

// isRateLimited returns the remaining backoff for the given category.
func (r *rateLimiter) isRateLimited(dsn string, category dataCategory, now time.Time) (time.Duration, bool) {
	_ = "STUB: not implemented"
	return *new(time.Duration), false
}

// cleanup expired map entries

// updateFromResponse updates rate limits using Sentry response headers.
func (r *rateLimiter) updateFromResponse(dsn string, resp *http.Response, now time.Time) {
	_ = "STUB: not implemented"
	return
}

func mapCategory(category dataCategory) internalrl.Category {
	_ = "STUB: not implemented"
	return *new(internalrl.Category)
}

// parseXSentryRateLimitReset parses the X-Sentry-Rate-Limit-Reset header from the Sentry API.
func parseXSentryRateLimitReset(reset string) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
