// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azureauthextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/azureauthextension"

import (
	"context"
	"crypto"
	"crypto/x509"
	"errors"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/coreos/go-oidc/v3/oidc"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/extension/extensionauth"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

var errServerAuthConfigRequired = errors.New("azure_auth server authentication requires server.issuer_url and server.audience")

type configServer struct {
	issuerURL string
	audience  string
}

type tokenVerifier interface {
	Verify(context.Context, string) error
}

type oidcTokenVerifier struct {
	verifier *oidc.IDTokenVerifier
}

func (v *oidcTokenVerifier) Verify(ctx context.Context, rawToken string) error {
	_ = "STUB: not implemented"
	return nil
}

type authenticator struct {
	credential azcore.TokenCredential
	logger     *zap.Logger
	scopes     []string
	server     configServer
	verifier   tokenVerifier
}

type tokenSource interface {
	Token(context.Context) (*oauth2.Token, error)
}

var (
	_ extension.Extension      = (*authenticator)(nil)
	_ extensionauth.HTTPClient = (*authenticator)(nil)
	_ extensionauth.Server     = (*authenticator)(nil)
	_ azcore.TokenCredential   = (*authenticator)(nil)
	_ tokenSource              = (*authenticator)(nil)
)

func newAzureAuthenticator(cfg *Config, logger *zap.Logger) (*authenticator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getCertificateAndKey from the file
func getCertificateAndKey(filename string) (*x509.Certificate, crypto.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, *new(crypto.PrivateKey), nil
}

func (a *authenticator) Start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (*authenticator) Shutdown(context.Context) error {
	_ = "STUB: not implemented"

	// GetToken returns an access token with a valid token for authorization.
	// Implements tokenSource interface.
	return nil
}

func (a *authenticator) GetToken(ctx context.Context, options policy.TokenRequestOptions) (azcore.AccessToken, error) {
	_ = "STUB: not implemented"
	return *

	// This is not expected, since creating a new authenticator
	// instance returns error if the supported credentials fail
	// to initialize, and any unexpected ones should be prevented
	// from validating the config.
	new(azcore.AccessToken), nil
}

// Token returns an access token with a valid token for authorization.
// Implements oauth2.TokenSource interface.
func (a *authenticator) Token(ctx context.Context) (*oauth2.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getHeaderValue(header string, headers map[string][]string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// getTokenForHost will request an access token based on a scope
// computed from the host value. It will return the token value
// or an error if request failed.
func (a *authenticator) getTokenForHost(ctx context.Context, host string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Example: if host is "management.azure.com", then the scope to get the
// token will be "https://management.azure.com/.default".
// See default scope: https://learn.microsoft.com/en-us/entra/identity-platform/scopes-oidc#the-default-scope.

func (a *authenticator) Authenticate(ctx context.Context, headers map[string][]string) (context.Context, error) {
	_ = "STUB: not implemented"
	// See request header: https://learn.microsoft.com/en-us/rest/api/azure/#request-header
	return *new(context.Context), nil
}

func (a *authenticator) RoundTripper(base http.RoundTripper) (http.RoundTripper, error) {
	_ = "STUB: not implemented"
	return *new(http.RoundTripper), nil
}

type roundTripper struct {
	base http.RoundTripper
	auth *authenticator
}

var _ http.RoundTripper = (*roundTripper)(nil)

func (r *roundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
