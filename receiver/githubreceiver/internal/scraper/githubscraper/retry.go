// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package githubscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/githubreceiver/internal/scraper/githubscraper"

import (
	"net/http"

	"go.uber.org/zap"
)

// retryRoundTripper wraps an http.RoundTripper and retries on transient GitHub
// API errors (429, 502, 503, 504) and secondary rate limits (403 + Retry-After).
// Retries use exponential backoff with jitter and are bounded by MaxRetries,
// MaxElapsedTime, and the request context (cancelled when the scrape cycle ends).
type retryRoundTripper struct {
	base   http.RoundTripper
	cfg    RetryConfig
	logger *zap.Logger
}

func (rt *retryRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Honor Retry-After header from GitHub rate limits. The value is
// used as-is (not capped) because retrying before the server's
// requested delay just wastes an attempt. Context cancellation
// and MaxRetries already bound total retry behavior.

// Drain and close the response body to reuse the TCP connection.

// Wait for backoff or context cancellation.

// Reset request body for retry (genqlient uses bytes.NewReader which
// auto-sets GetBody, making POST bodies replayable).

// isRetryable returns true for HTTP status codes that indicate a transient
// GitHub API error worth retrying.
func isRetryable(resp *http.Response) bool { _ = "STUB: not implemented"; return false }

// 429
// 502
// 503
// 504

// 403 -- only with Retry-After (secondary rate limit)

// parseRetryAfter extracts the delay in seconds from a Retry-After header.
// Returns 0 if the header is absent or not a valid integer.
func parseRetryAfter(h http.Header) int { _ = "STUB: not implemented"; return 0 }
