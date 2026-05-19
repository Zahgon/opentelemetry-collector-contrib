// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package splunk // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk"

import (
	"net/http"
)

const HeaderRetryAfter = "Retry-After"

// HandleHTTPCode handles an http response and returns the right type of error in case of a failure.
func HandleHTTPCode(resp *http.Response) error {
	_ = "STUB: not implemented"
	// Splunk accepts all 2XX codes.
	return nil
}

// Check for responses that may include "Retry-After" header.

// Fallback to 0 if the Retry-After header is not present. This will trigger the
// default backoff policy by our caller (retry handler).

// Indicate to our caller to pause for the specified number of seconds.

// Check for permanent errors.
