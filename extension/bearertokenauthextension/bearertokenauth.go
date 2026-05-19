// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package bearertokenauthextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/bearertokenauthextension"

import (
	"context"
	"net/http"
	"sync/atomic"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/extension/extensionauth"
	"go.uber.org/zap"
	"google.golang.org/grpc/credentials"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/internal/credentialsfile"
)

var _ credentials.PerRPCCredentials = (*perRPCAuth)(nil)

// PerRPCAuth is a gRPC credentials.PerRPCCredentials implementation that returns an 'authorization' header.
type perRPCAuth struct {
	auth *bearerTokenAuth
}

// GetRequestMetadata returns the request metadata to be used with the RPC.
func (c *perRPCAuth) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RequireTransportSecurity always returns true for this implementation. Passing bearer tokens in plain-text connections is a bad idea.
func (*perRPCAuth) RequireTransportSecurity() bool { _ = "STUB: not implemented"; return false }

var (
	_ extension.Extension      = (*bearerTokenAuth)(nil)
	_ extensionauth.Server     = (*bearerTokenAuth)(nil)
	_ extensionauth.HTTPClient = (*bearerTokenAuth)(nil)
	_ extensionauth.GRPCClient = (*bearerTokenAuth)(nil)
)

// BearerTokenAuth is an implementation of extensionauth interfaces. It embeds a static authorization "bearer" token in every rpc call.
type bearerTokenAuth struct {
	header                    string
	scheme                    string
	authorizationValuesAtomic atomic.Value

	tokenResolver credentialsfile.ValueResolver
	logger        *zap.Logger
}

func newBearerTokenAuth(cfg *Config, logger *zap.Logger) *bearerTokenAuth {
	_ = "STUB: not implemented"
	return nil
}

// Create token resolver for single token (inline or file)

// Initialize token values

// Start of BearerTokenAuth does nothing and returns nil if no filename
// is specified. Otherwise a routine is started to monitor the file containing
// the token to be transferred.
func (b *bearerTokenAuth) Start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *bearerTokenAuth) updateAuthorizationValues() { _ = "STUB: not implemented"; return }

// Split by whitespace (spaces, tabs, etc.)
// strings.Fields handles leading/trailing whitespace and multiple spaces automatically.

// If the line has at least one part, the first part is the token.
// Everything else is treated as a comment/ignored.

func (b *bearerTokenAuth) setAuthorizationValues(tokens []string) {
	_ = "STUB: not implemented"
	return
}

// authorizationValues returns the Authorization header/metadata values
// to set for client auth, and expected values for server auth.
func (b *bearerTokenAuth) authorizationValues() []string { _ = "STUB: not implemented"; return nil }

// authorizationValue returns the first Authorization header/metadata value
// to set for client auth, and expected value for server auth.
func (b *bearerTokenAuth) authorizationValue() string { _ = "STUB: not implemented"; return "" }

// Return the first token

// Shutdown of BearerTokenAuth does nothing and returns nil
func (b *bearerTokenAuth) Shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }

// PerRPCCredentials returns PerRPCAuth an implementation of credentials.PerRPCCredentials that
func (b *bearerTokenAuth) PerRPCCredentials() (credentials.PerRPCCredentials, error) {
	_ = "STUB: not implemented"
	return *new(credentials.PerRPCCredentials), nil
}

// RoundTripper is not implemented by BearerTokenAuth
func (b *bearerTokenAuth) RoundTripper(base http.RoundTripper) (http.RoundTripper, error) {
	_ = "STUB: not implemented"
	return *new(http.RoundTripper), nil
}

// Authenticate checks whether the given context contains valid auth data. Validates tokens from clients trying to access the service (incoming requests)
func (b *bearerTokenAuth) Authenticate(ctx context.Context, headers map[string][]string) (context.Context, error) {
	_ = "STUB: not implemented"
	// Use canonical header key to match how Go's HTTP server stores headers
	return *new(context.Context), nil
}

// Also check lower-case header key to support gRPC metadata format

// Extract token from authorization header

// Authentication successful, token is valid

// Token is invalid

// BearerAuthRoundTripper intercepts and adds Bearer token Authorization headers to each http request.
type bearerAuthRoundTripper struct {
	header        string
	baseTransport http.RoundTripper
	auth          *bearerTokenAuth
}

// RoundTrip modifies the original request and adds Bearer token Authorization headers. Incoming requests support multiple tokens, but outgoing requests only use one.
func (interceptor *bearerAuthRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
