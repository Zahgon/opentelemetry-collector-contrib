// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package webhookeventreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/webhookeventreceiver"

import (
	"bufio"
	"errors"
	"net/http"
	"net/url"
	"regexp"

	"go.opentelemetry.io/collector/pdata/plog"
)

const (
	headerNamespace = "header"
)

var errRequestBodyTooLarge = errors.New("request body exceeds maximum allowed size")

func (er *eventReceiver) reqToLog(sc *bufio.Scanner,
	headers http.Header,
	query url.Values,
) (plog.Logs, int, error) {
	_ = "STUB: not implemented"
	// we simply dont split the data passed into scan (i.e. scan the whole thing)
	// the downside to this approach is that only 1 log per request can be handled.
	// NOTE: logs will contain these newline characters which could have formatting
	// consequences downstream.
	return *new(plog.Logs), 0, nil
}

// Increase max token size from default 64KB to configured value

// append query parameters and webhook source as resource attributes
func appendMetadata(resourceLog plog.ResourceLogs, query url.Values) {
	_ = "STUB: not implemented"
	return
}

// append headers as logRecord attributes if they match supplied regex
func appendHeaders(h http.Header, l plog.LogRecord, r *regexp.Regexp) {
	_ = "STUB: not implemented"

	// Skip the required header used for authentication
	return
}

// prepend the header key with the "header." namespace
func headerAttributeKey(header string) string { _ = "STUB: not implemented"; return "" }

func splitJSONObjects(data string) []string { _ = "STUB: not implemented"; return nil }

// If we hit EOF or any other error, we're done

// If no valid JSON objects were found, return the original data
