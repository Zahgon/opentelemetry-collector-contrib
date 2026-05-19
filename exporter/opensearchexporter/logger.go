// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package opensearchexporter contains an opentelemetry-collector exporter
// for OpenSearch.
package opensearchexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/opensearchexporter"

import (
	"net/http"
	"time"

	"github.com/opensearch-project/opensearch-go/v4/opensearchtransport"
	"go.uber.org/zap"
)

type clientLogger struct {
	zapLogger *zap.Logger
}

func newClientLogger(zl *zap.Logger) opensearchtransport.Logger {
	_ = "STUB: not implemented"
	return *

	// LogRoundTrip should not modify the request or response, except for consuming and closing the body.
	// Implementations have to check for nil values in request and response.
	new(opensearchtransport.Logger)
}

func (cl *clientLogger) LogRoundTrip(requ *http.Request, resp *http.Response, err error, _ time.Time, dur time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// RequestBodyEnabled makes the client pass a copy of request body to the logger.
func (*clientLogger) RequestBodyEnabled() bool {
	_ = "STUB: not implemented"
	// TODO: introduce setting log the bodies for more detailed debug logs
	return false
}

// ResponseBodyEnabled makes the client pass a copy of response body to the logger.
func (*clientLogger) ResponseBodyEnabled() bool {
	_ = "STUB: not implemented"
	// TODO: introduce setting log the bodies for more detailed debug logs
	return false
}
