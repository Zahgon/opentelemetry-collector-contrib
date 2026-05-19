// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package coralogixexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/coralogixexporter"

import (
	"net/http"
	"time"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// shouldRetry returns true if the error should be retried.
// The second return value indicates whether the error should trigger a stop in retries by flagging
// the rate limiting mechanism, since these errors (like authentication or quota failures) indicate a problem
// that won't be fixed just by retrying.
func shouldRetry(code codes.Code, retryInfo *errdetails.RetryInfo) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

// Retry only if RetryInfo was supplied by the server.
// This indicates that the server can still recover from resource exhaustion.

func getRetryInfo(status *status.Status) *errdetails.RetryInfo {
	_ = "STUB: not implemented"
	return nil
}

func getThrottleDuration(t *errdetails.RetryInfo) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// shouldRetryHTTP returns true if the HTTP status code should be retried.
// The second return value indicates whether the error should trigger a stop in retries by flagging
// the rate limiting mechanism.
func shouldRetryHTTP(statusCode int) (bool, bool) { _ = "STUB: not implemented"; return false, false }

// Permanent errors - don't retry

// Don't retry, but flag for rate limiting

// Retryable errors

// Retry with throttle

// getHTTPThrottleDuration extracts retry delay from Retry-After header for HTTP 429 responses
func getHTTPThrottleDuration(statusCode int, headers http.Header) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// Try parsing as seconds (integer)

// Try parsing as HTTP date
