// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package headerssetterextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/headerssetterextension"

import (
	"context"
	"net/http"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/extension/extensionauth"
	"go.opentelemetry.io/collector/extension/extensioncapabilities"
	"go.uber.org/zap"
	"google.golang.org/grpc/credentials"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/headerssetterextension/internal/action"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/headerssetterextension/internal/source"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/internal/credentialsfile"
)

type header struct {
	action action.Action
	source source.Source
}

var (
	_ extension.Extension             = (*headerSetterExtension)(nil)
	_ extensionauth.HTTPClient        = (*headerSetterExtension)(nil)
	_ extensionauth.GRPCClient        = (*headerSetterExtension)(nil)
	_ extensioncapabilities.Dependent = (*headerSetterExtension)(nil)
)

type headerSetterExtension struct {
	component.StartFunc
	component.ShutdownFunc

	headers        []header
	additionalAuth *component.ID
	host           component.Host
	resolvers      []credentialsfile.ValueResolver
	logger         *zap.Logger
}

// Dependencies implements extensioncapabilities.Dependent.
func (h *headerSetterExtension) Dependencies() []component.ID {
	_ = "STUB: not implemented"
	return nil
}

// Start stores the host for later use in getting the additional auth extension.
func (h *headerSetterExtension) start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"

	// Start all file resolvers
	return nil
}

// Shutdown stops all file resolvers
func (h *headerSetterExtension) shutdown(_ context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// getAdditionalAuthExtension retrieves the configured additional auth extension if present.
// Returns nil if no additional auth is configured.
func (h *headerSetterExtension) getAdditionalAuthExtension() (component.Component, error) {
	_ = "STUB: not implemented"
	return *new(component.Component), nil
}

// PerRPCCredentials implements extensionauth.GRPCClient.
func (h *headerSetterExtension) PerRPCCredentials() (credentials.PerRPCCredentials, error) {
	_ = "STUB: not implemented"
	return *new(credentials.PerRPCCredentials), nil
}

// If additional_auth is configured, chain with it first

// RoundTripper implements extensionauth.HTTPClient.
func (h *headerSetterExtension) RoundTripper(base http.RoundTripper) (http.RoundTripper, error) {
	_ = "STUB: not implemented"
	// If additional_auth is configured, chain with it first
	return *new(http.RoundTripper), nil
}

// Check if it implements HTTPClient

// Now wrap with our headers

func newHeadersSetterExtension(cfg *Config, logger *zap.Logger) (*headerSetterExtension, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Enable Start/Shutdown methods if additional_auth is configured or if we have file resolvers

// headersPerRPC is a gRPC credentials.PerRPCCredentials implementation sets
// headers with values extracted from provided sources.
type headersPerRPC struct {
	headers         []header
	baseCredentials credentials.PerRPCCredentials
}

// GetRequestMetadata returns the request metadata to be used with the RPC.
func (h *headersPerRPC) GetRequestMetadata(
	ctx context.Context,
	uri ...string,
) (map[string]string, error) {
	_ = "STUB: not implemented"
	// Start with base credentials if available
	return nil, nil
}

// Copy base metadata

// Now apply our headers on top

// RequireTransportSecurity returns whether transport security is required.
// If chained with another auth extension, delegate to it.
func (h *headersPerRPC) RequireTransportSecurity() bool { _ = "STUB: not implemented"; return false }

// headersRoundTripper intercepts downstream requests and sets headers with
// values extracted from configured sources.
type headersRoundTripper struct {
	base    http.RoundTripper
	headers []header
}

// RoundTrip copies the original request and sets headers of the new requests
// with values extracted from configured sources.
func (h *headersRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
