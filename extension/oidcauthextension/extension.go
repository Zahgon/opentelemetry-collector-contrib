// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package oidcauthextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/oidcauthextension"

import (
	"context"
	"crypto/x509"
	"errors"
	"net/http"
	"sync"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/fsnotify/fsnotify"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/extension/extensionauth"
	"go.uber.org/zap"
)

var (
	_ extension.Extension  = (*oidcExtension)(nil)
	_ extensionauth.Server = (*oidcExtension)(nil)
)

type providerContainer struct {
	providerCfg ProviderCfg
	provider    *oidc.Provider
	verifier    *oidc.IDTokenVerifier
	verifierMu  sync.RWMutex // protects verifier for hot-reload
	verifierCfg *oidc.Config // stored for recreating verifier on reload
	client      *http.Client
	transport   *http.Transport
}

func (pc *providerContainer) Verify(ctx context.Context, raw string) (*oidc.IDToken, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// No need to grab the lock, we aren't hot-reloading from a file.

func (pc *providerContainer) refreshPublicKeysFile() error { _ = "STUB: not implemented"; return nil }

// jose.SignatureAlgorithm actually is a string.

type oidcExtension struct {
	cfg *Config

	providerContainers map[string]*providerContainer
	logger             *zap.Logger
	shutdownCH         chan struct{}
	shutdownOnce       sync.Once
}

var (
	errNoAudienceProvided                = errors.New("no Audience provided for the OIDC configuration")
	errNoIssuerURL                       = errors.New("no IssuerURL provided for the OIDC configuration")
	errInvalidAuthenticationHeaderFormat = errors.New("invalid authorization header format")
	errFailedToObtainClaimsFromToken     = errors.New("failed to get the subject from the token issued by the OIDC provider")
	errClaimNotFound                     = errors.New("username claim from the OIDC configuration not found on the token returned by the OIDC provider")
	errUsernameNotString                 = errors.New("the username returned by the OIDC provider isn't a regular string")
	errGroupsClaimNotFound               = errors.New("groups claim from the OIDC configuration not found on the token returned by the OIDC provider")
	errNotAuthenticated                  = errors.New("authentication didn't succeed")
	errNoSupportedKeys                   = errors.New("file contains no supported keys (supported types are RSA, ECDSA, and ED25519)")
)

func newExtension(cfg *Config, logger *zap.Logger) extension.Extension {
	_ = "STUB: not implemented"
	return *new(extension.Extension)
}

func (e *oidcExtension) Start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Start file watchers for providers with JWKS files

func (e *oidcExtension) Shutdown(context.Context) error {
	_ = "STUB: not implemented"
	// Signal file watchers to stop
	return nil
}

// authenticate checks whether the given context contains valid auth data. Successfully authenticated calls will always return a nil error and a context with the auth data.
func (e *oidcExtension) Authenticate(ctx context.Context, headers map[string][]string) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

// we only use the first header, if multiple values exist

// currently, this isn't a valid condition, the Verify call a few lines above
// will already attempt to parse the payload as a json and set it as the claims
// for the token. As we are using a map to hold the claims, there's no way to fail
// to read the claims. It could fail if we were using a custom struct. Instead of
// swallowing the error, it's better to make this future-proof, in case the underlying
// code changes

func (e *oidcExtension) resolveProvider(issuer string) (*providerContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *oidcExtension) processProviderConfig(ctx context.Context, p ProviderCfg) error {
	_ = "STUB: not implemented"
	return nil
}

// the errors from this path have enough context already

func getSubjectFromClaims(claims map[string]any, usernameClaim, fallback string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getGroupsFromClaims(claims map[string]any, groupsClaim string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getIssuerCACertFromPath(path string) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *oidcExtension) startFileWatchers(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *oidcExtension) watchFiles(ctx context.Context, watcher *fsnotify.Watcher, filesToWatch map[string][]*providerContainer) {
	_ = "STUB: not implemented"
	return
}

// This struct is kept minimal to avoid unnecessary allocations, especially
// if the JWT is malicious and contains too many claims.
type idToken struct {
	Issuer string `json:"iss"`
}

// Get the issuer from the raw ID token.
// This function is unsafe because it does not verify the token's signature.
// It should only be used to determine which verifier to use for the token.
func getIssuerFromUnverifiedJWT(rawIDToken string) (string, error) {
	_ = "STUB: not implemented"
	// TODO: it would be nice if we didn't have to parse the JWT here and then again in the verifier...
	return "", nil
}
