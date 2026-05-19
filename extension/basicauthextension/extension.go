// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package basicauthextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/basicauthextension"

import (
	"context"
	"net/http"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/extension/extensionauth"
	"go.uber.org/zap"
	creds "google.golang.org/grpc/credentials"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/internal/basicauth"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/internal/credentialsfile"
)

var (
	errNoAuth              = basicauth.ErrNoAuth
	errInvalidCredentials  = basicauth.ErrInvalidCredentials
	errInvalidSchemePrefix = basicauth.ErrInvalidSchemePrefix
	errInvalidFormat       = basicauth.ErrInvalidFormat
)

func newClientAuthExtension(cfg *Config) *basicAuthClient { _ = "STUB: not implemented"; return nil }

func newServerAuthExtension(cfg *Config) (*basicAuthServer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var (
	_ extension.Extension  = (*basicAuthServer)(nil)
	_ extensionauth.Server = (*basicAuthServer)(nil)
)

type basicAuthServer struct {
	htpasswd  *HtpasswdSettings
	matchFunc func(username, password string) bool
	component.ShutdownFunc
}

func (ba *basicAuthServer) Start(_ context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Ensure that the inline content is read the last.
// This way the inline content will override the content from file.

func (ba *basicAuthServer) Authenticate(ctx context.Context, headers map[string][]string) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

var (
	_ extension.Extension      = (*basicAuthClient)(nil)
	_ extensionauth.HTTPClient = (*basicAuthClient)(nil)
	_ extensionauth.GRPCClient = (*basicAuthClient)(nil)
)

type basicAuthClient struct {
	clientAuth       *ClientAuthSettings
	logger           *zap.Logger
	usernameResolver credentialsfile.ValueResolver
	passwordResolver credentialsfile.ValueResolver
}

func (ba *basicAuthClient) Start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (ba *basicAuthClient) Shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (ba *basicAuthClient) Username() string { _ = "STUB: not implemented"; return "" }

func (ba *basicAuthClient) Password() string { _ = "STUB: not implemented"; return "" }

func (ba *basicAuthClient) RoundTripper(base http.RoundTripper) (http.RoundTripper, error) {
	_ = "STUB: not implemented"
	return *new(http.RoundTripper), nil
}

func (ba *basicAuthClient) PerRPCCredentials() (creds.PerRPCCredentials, error) {
	_ = "STUB: not implemented"
	return *new(creds.PerRPCCredentials), nil
}
