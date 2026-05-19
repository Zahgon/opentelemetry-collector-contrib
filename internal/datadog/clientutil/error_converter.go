// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package clientutil // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/datadog/clientutil"

import (
	"net/http"
)

// WrapError wraps an error to a permanent consumer error that won't be retried if the http response code is non-retriable.
func WrapError(err error, resp *http.Response) error { _ = "STUB: not implemented"; return nil }

func isNonRetriable(resp *http.Response) bool { _ = "STUB: not implemented"; return false }
