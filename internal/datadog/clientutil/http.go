// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package clientutil // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/datadog/clientutil"

import (
	"net/http"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
)

var (
	// JSONHeaders headers for JSON requests.
	JSONHeaders = map[string]string{
		"Content-Type":     "application/json",
		"Content-Encoding": "gzip",
	}
	// ProtobufHeaders headers for protobuf requests.
	ProtobufHeaders = map[string]string{
		"Content-Type":     "application/x-protobuf",
		"Content-Encoding": "identity",
	}
)

// NewHTTPClient returns a http.Client configured with a subset of the confighttp.ClientConfig options.
func NewHTTPClient(hcs confighttp.ClientConfig) *http.Client { _ = "STUB: not implemented"; return nil }

// NewHTTPTransport returns a http.Transport configured with a subset of the confighttp.ClientConfig options.
func NewHTTPTransport(hcs confighttp.ClientConfig) *http.Transport {
	_ = "STUB: not implemented"
	// If the ProxyURL field in the configuration is set, the HTTP client will use the proxy.
	// Otherwise, the HTTP client will use the system's proxy settings.
	return nil
}

// Default values consistent with https://github.com/DataDog/datadog-agent/blob/f9ae7f4b842f83b23b2dfe3f15d31f9e6b12e857/pkg/util/http/transport.go#L91-L106

// Enables TCP keepalives to detect broken connections

// Disable RFC 6555 Fast Fallback ("Happy Eyeballs")

// This parameter is set to avoid connections sitting idle in the pool indefinitely

// Not supported by intake

// SetExtraHeaders appends a header map to HTTP headers.
func SetExtraHeaders(h http.Header, extras map[string]string) { _ = "STUB: not implemented"; return }

func UserAgent(buildInfo component.BuildInfo) string { _ = "STUB: not implemented"; return "" }

// SetDDHeaders sets the Datadog-specific headers
func SetDDHeaders(reqHeader http.Header, buildInfo component.BuildInfo, apiKey string) {
	_ = "STUB: not implemented"
	return
}
