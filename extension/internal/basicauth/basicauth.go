// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package basicauth // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/internal/basicauth"

import (
	"context"
	"errors"
	"net/http"

	"go.opentelemetry.io/collector/client"
	creds "google.golang.org/grpc/credentials"
)

var (
	ErrNoAuth              = errors.New("no basic auth provided")
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrInvalidSchemePrefix = errors.New("invalid authorization scheme prefix")
	ErrInvalidFormat       = errors.New("invalid authorization format")
)

// CredentialProvider supplies username and password for client-side basic auth.
type CredentialProvider interface {
	Username() string
	Password() string
}

// AuthData implements client.AuthData for basic auth.
type AuthData struct {
	username string
	password string
	raw      string
}

var _ client.AuthData = (*AuthData)(nil)

func (a *AuthData) GetAttribute(name string) any { _ = "STUB: not implemented"; return *new(any) }

func (*AuthData) GetAttributeNames() []string { _ = "STUB: not implemented"; return nil }

// GetAuthHeader extracts the Authorization header value from a header map,
// handling canonical, lowercase, and case-insensitive lookups.
func GetAuthHeader(h map[string][]string) string { _ = "STUB: not implemented"; return "" }

// ParseBasicAuth parses a "Basic <base64>" authorization header value.
func ParseBasicAuth(auth string) (*AuthData, error) { _ = "STUB: not implemented"; return nil, nil }

// Authenticate performs server-side basic auth validation against the given header map.
// It extracts the Authorization header, parses the credentials, and validates them
// using the provided match function. On success, it stores the auth data in the context.
func Authenticate(ctx context.Context, headers map[string][]string, matchFunc func(username, password string) bool) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

// RoundTripper wraps a base http.RoundTripper to inject basic auth credentials.
type RoundTripper struct {
	Base     http.RoundTripper
	Provider CredentialProvider
}

func (rt *RoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewRoundTripper creates an http.RoundTripper that adds basic auth from the provider.
func NewRoundTripper(base http.RoundTripper, provider CredentialProvider) (http.RoundTripper, error) {
	_ = "STUB: not implemented"
	return *new(http.RoundTripper), nil
}

// PerRPCAuth implements grpc credentials.PerRPCCredentials using a CredentialProvider.
type PerRPCAuth struct {
	Provider CredentialProvider
}

func (p *PerRPCAuth) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*PerRPCAuth) RequireTransportSecurity() bool {
	_ = "STUB: not implemented"

	// NewPerRPCCredentials creates gRPC PerRPCCredentials that add basic auth from the provider.
	return false
}

func NewPerRPCCredentials(provider CredentialProvider) (creds.PerRPCCredentials, error) {
	_ = "STUB: not implemented"
	return *new(creds.PerRPCCredentials), nil
}
