// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package httpforwarderextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/httpforwarderextension"

import (
	"context"
	"net/http"
	"net/url"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

type httpForwarder struct {
	forwardTo  *url.URL
	httpClient *http.Client
	server     *http.Server
	settings   component.TelemetrySettings
	config     *Config
	shutdownWG sync.WaitGroup
}

var _ extension.Extension = (*httpForwarder)(nil)

func (h *httpForwarder) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *httpForwarder) Shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (h *httpForwarder) forwardRequest(writer http.ResponseWriter, request *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Clear RequestURI to avoid getting "http: Request.RequestURI can't be set in client requests" error.

// Add additional headers.

// Add "Via" header for tracking purposes on both the outgoing requests and responses.
// See https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Via.

// Copy over response from the final destination.

func addViaHeader(header http.Header, protocol, host string) { _ = "STUB: not implemented"; return }

func newHTTPForwarder(config *Config, settings component.TelemetrySettings) (extension.Extension, error) {
	_ = "STUB: not implemented"
	return *new(extension.Extension), nil
}
