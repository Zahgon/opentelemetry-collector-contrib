// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package faroexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/faroexporter"

import (
	"google.golang.org/grpc/status"
)

func newStatusFromMsgAndHTTPCode(errMsg string, statusCode int) *status.Status {
	_ = "STUB: not implemented"

	// Mapping based on https://github.com/grpc/grpc/blob/master/doc/http-grpc-status-mapping.md
	// 429 mapping to ResourceExhausted and 400 mapping to StatusBadRequest are exceptions.
	return nil
}

func isRetryableStatusCode(code int) bool { _ = "STUB: not implemented"; return false }
