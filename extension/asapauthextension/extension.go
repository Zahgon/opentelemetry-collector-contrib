// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package asapauthextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/asapauthextension"

import (
	"context"
	"net/http"

	"bitbucket.org/atlassian/go-asap/v2"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/extension/extensionauth"
	"google.golang.org/grpc/credentials"
)

var (
	_ extension.Extension      = (*asapAuthExtension)(nil)
	_ extensionauth.HTTPClient = (*asapAuthExtension)(nil)
	_ extensionauth.GRPCClient = (*asapAuthExtension)(nil)
)

type asapAuthExtension struct {
	component.StartFunc
	component.ShutdownFunc

	provisioner asap.Provisioner
	privateKey  any
}

// PerRPCCredentials returns extensionauth.GRPCClient.
func (e *asapAuthExtension) PerRPCCredentials() (credentials.PerRPCCredentials, error) {
	_ = "STUB: not implemented"
	return *new(credentials.PerRPCCredentials), nil
}

// RoundTripper implements extensionauth.HTTPClient.
func (e *asapAuthExtension) RoundTripper(base http.RoundTripper) (http.RoundTripper, error) {
	_ = "STUB: not implemented"
	return *new(http.RoundTripper), nil
}

func createASAPClientAuthenticator(cfg *Config) (*asapAuthExtension, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Caching provisioner will only issue a new token after the current token's expiry (determined by TTL).

// perRPCAuth is a gRPC credentials.PerRPCCredentials implementation that returns an 'authorization' header.
type perRPCAuth struct {
	provisioner asap.Provisioner
	privateKey  any
}

// GetRequestMetadata returns the request metadata to be used with the RPC.
func (c *perRPCAuth) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RequireTransportSecurity always returns true for this implementation.
func (*perRPCAuth) RequireTransportSecurity() bool { _ = "STUB: not implemented"; return false }
