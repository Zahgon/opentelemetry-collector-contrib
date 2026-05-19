// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package oauth2clientauthextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/oauth2clientauthextension"

import (
	"context"
	"net/url"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
)

const (
	grantTypeJWTBearer = "urn:ietf:params:oauth:grant-type:jwt-bearer" //nolint:gosec // false positive, this is the grant-type name
)

func newJwtGrantTypeConfig(cfg *Config) (*jwtGrantTypeConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Config is the configuration for using JWT to fetch tokens,
// commonly known as "two-legged OAuth 2.0".
type jwtGrantTypeConfig struct {
	// Iss is the OAuth client identifier used when communicating with
	// the configured OAuth provider.
	Iss string

	// PrivateKey contains the contents of an RSA private key or the
	// contents of a PEM file that contains a private key. The provided
	// private key is used to sign JWT payloads.
	// PEM containers with a passphrase are not supported.
	// Use the following command to convert a PKCS 12 file into a PEM.
	//
	//    $ openssl pkcs12 -in key.p12 -out key.pem -nodes
	//
	PrivateKey []byte

	// SigningAlgorithm is the RSA algorithm used to sign JWT payloads
	SigningAlgorithm *jwt.SigningMethodRSA

	// PrivateKeyID contains an optional hint indicating which key is being
	// used.
	PrivateKeyID string

	// Subject is the optional user to impersonate.
	Subject string

	// Scopes optionally specifies a list of requested permission scopes.
	Scopes []string

	// TokenURL is the endpoint required to complete the 2-legged JWT flow.
	TokenURL string

	// EndpointParams specifies additional parameters for requests to the token endpoint.
	EndpointParams url.Values

	// Expires optionally specifies how long the token is valid for.
	Expires time.Duration

	// Audience optionally specifies the intended audience of the
	// request.  If empty, the value of TokenURL is used as the
	// intended audience.
	Audience string

	// PrivateClaims optionally specifies custom private claims in the JWT.
	// See http://tools.ietf.org/html/draft-jones-json-web-token-10#section-4.3
	PrivateClaims map[string]any
}

// TokenSource returns a JWT TokenSource using the configuration
// in c and the HTTP client from the provided context.
func (c *jwtGrantTypeConfig) TokenSource(ctx context.Context) oauth2.TokenSource {
	_ = "STUB: not implemented"
	return *new(oauth2.TokenSource)
}

func (c *jwtGrantTypeConfig) TokenEndpoint() string {
	_ = "STUB: not implemented"

	// jwtSource implements TokenSource
	return ""
}

var _ oauth2.TokenSource = (*jwtSource)(nil)

// jwtSource is a source that always does a signed JWT request for a token.
// It should typically be wrapped with a reuseTokenSource.
type jwtSource struct {
	ctx  context.Context
	conf *jwtGrantTypeConfig
}

func (js jwtSource) Token() (*oauth2.Token, error) { _ = "STUB: not implemented"; return nil, nil }

// Allow grant_type to be overridden to allow interoperability with
// non-compliant implementations.

// tokenRes is the JSON response body.
