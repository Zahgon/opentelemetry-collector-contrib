// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package opampextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/opampextension"

import (
	"net/http"

	"go.opentelemetry.io/collector/component"
	"go.uber.org/zap"
)

// headerCaptureRoundTripper is a RoundTripper that captures the headers of the request
// that passes through it.
type headerCaptureRoundTripper struct {
	lastHeader http.Header
}

func (h *headerCaptureRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Dummy response is recorded here

func makeHeadersFunc(logger *zap.Logger, serverCfg *OpAMPServer, host component.Host) (func(http.Header) http.Header, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This is a workaround while websocket authentication is being worked on.
// Currently, we are waiting on the auth module to be stabilized.
// See for more info: https://github.com/open-telemetry/opentelemetry-collector/issues/10864
