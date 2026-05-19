// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package oauth2clientauthextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/oauth2clientauthextension"

import (
	"context"
	"errors"
	"net/http"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/extension/extensionauth"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"google.golang.org/grpc/credentials"
)

var (
	_ extension.Extension      = (*clientAuthenticator)(nil)
	_ extensionauth.HTTPClient = (*clientAuthenticator)(nil)
	_ extensionauth.GRPCClient = (*clientAuthenticator)(nil)
	_ ContextTokenSource       = (*clientAuthenticator)(nil)
)

// errFailedToGetSecurityToken indicates a problem communicating with OAuth2 server.
var errFailedToGetSecurityToken = errors.New("failed to get security token from token endpoint")

type TokenSourceConfiguration interface {
	TokenSource(context.Context) oauth2.TokenSource
	TokenEndpoint() string
}

// ContextTokenSource provides an interface for obtaining an *oauth2.Token,
// with the given context.
type ContextTokenSource interface {
	Token(context.Context) (*oauth2.Token, error)
}

// clientAuthenticator provides implementation for providing client authentication using OAuth2 client credentials
// workflow for both gRPC and HTTP clients.
type clientAuthenticator struct {
	component.StartFunc
	component.ShutdownFunc

	credentials  TokenSourceConfiguration
	logger       *zap.Logger
	client       *http.Client
	expiryBuffer time.Duration

	// sem is a buffered channel of size 1 used as a context-aware mutex
	// to protect token access/refresh.
	sem   chan struct{}
	token *oauth2.Token
}

func newClientAuthenticator(cfg *Config, logger *zap.Logger) (*clientAuthenticator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RoundTripper wraps the provided base http.RoundTripper with an oauth2.Transport
// that injects OAuth2 tokens into outgoing HTTP requests. The returned RoundTripper
// will refresh tokens as needed using the context of the outgoing HTTP request.
func (o *clientAuthenticator) RoundTripper(base http.RoundTripper) (http.RoundTripper, error) {
	_ = "STUB: not implemented"
	return *new(http.RoundTripper), nil
}

type roundTripper struct {
	o    *clientAuthenticator
	base http.RoundTripper
}

func (rt *roundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PerRPCCredentials returns a gRPC PerRPCCredentials that injects OAuth2 tokens into
// outgoing gRPC requests. The returned PerRPCCredentials will refresh tokens as needed
// using the request context.
func (o *clientAuthenticator) PerRPCCredentials() (credentials.PerRPCCredentials, error) {
	_ = "STUB: not implemented"
	return *new(credentials.PerRPCCredentials), nil
}

// Token returns an oauth2.Token, refreshing as needed with the provided context.
// The returned Token must not be modified.
func (o *clientAuthenticator) Token(ctx context.Context) (*oauth2.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// perRPCCredentials is based on google.golang.org/grpc/credentials/oauth.TokenSource,
// but passes through the request context.
type perRPCCredentials struct {
	o *clientAuthenticator
}

// GetRequestMetadata gets the request metadata as a map from a TokenSource.
func (c *perRPCCredentials) GetRequestMetadata(ctx context.Context, _ ...string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RequireTransportSecurity indicates whether the credentials requires transport security.
func (*perRPCCredentials) RequireTransportSecurity() bool { _ = "STUB: not implemented"; return false }
