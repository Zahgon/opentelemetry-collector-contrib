// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package proxy provides an http server to act as a signing proxy for SDKs calling AWS X-Ray APIs
package proxy // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/proxy"

import (
	"context"
	"io"
	"net/http"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
)

// Server represents HTTP server.
type Server interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

var _ Server = (*server)(nil)

type server struct {
	server       *http.Server
	serverConfig *confighttp.ServerConfig
}

func (s *server) ListenAndServe() error { _ = "STUB: not implemented"; return nil }

func (s *server) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// NewServer returns a local TCP server that proxies requests to AWS
// backend using the given credentials.
func NewServer(cfg *Config, host component.Host, settings component.TelemetrySettings) (Server, error) {
	_ = "STUB: not implemented"
	return *new(Server), nil
}

// Parse url from endpoint

// Reverse proxy handler

// Set outbound request URL to the AWS endpoint

// Consume body and calculate payload hash for signing

// Restore body for the request

// Retrieve credentials for signing

// Sign request using v4 signer

// consumeBody reads the body, calculates SHA-256 hash, and returns the body bytes and hash.
// v4.Signer requires a payload hash for signing.
func consumeBody(body io.ReadCloser) ([]byte, string, error) {
	_ = "STUB: not implemented"
	// Return empty hash if body is nil
	return nil, "", nil
}

// SHA-256 of empty string

// Consume body

// Calculate SHA-256 hash
