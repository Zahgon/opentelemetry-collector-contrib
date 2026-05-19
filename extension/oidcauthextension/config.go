// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package oidcauthextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/oidcauthextension"

import (
	"crypto"
	"maps"
	"slices"

	"github.com/go-jose/go-jose/v4"
)

// Config has the configuration for the OIDC Authenticator extension.
type Config struct {
	// The attribute (header name) to look for auth data. Optional, default value: "authorization".
	Attribute string `mapstructure:"attribute"`

	// Deprecated: use Providers instead.
	// IssuerURL is the base URL for the OIDC provider.
	// Required.
	IssuerURL string `mapstructure:"issuer_url"`

	// Deprecated: use Providers instead.
	// Audience of the token, used during the verification.
	// For example: "https://accounts.google.com" or "https://login.salesforce.com".
	// Required unless IgnoreAudience is true.
	Audience string `mapstructure:"audience"`

	// Deprecated: use Providers instead.
	// When true, this skips validating the audience field.
	// Optional.
	IgnoreAudience bool `mapstructure:"ignore_audience"`

	// Deprecated: use Providers instead.
	// The local path for the issuer CA's TLS server cert.
	// Optional.
	IssuerCAPath string `mapstructure:"issuer_ca_path"`

	// Deprecated: use Providers instead.
	// The claim to use as the username, in case the token's 'sub' isn't the suitable source.
	// Optional.
	UsernameClaim string `mapstructure:"username_claim"`

	// Deprecated: use Providers instead.
	// The claim that holds the subject's group membership information.
	// Optional.
	GroupsClaim string `mapstructure:"groups_claim"`

	// Providers allows configuring multiple OIDC providers.
	// Use the getProviderConfigs() method to get the full list of providers, including the legacy configuration.
	Providers []ProviderCfg `mapstructure:"providers"`
}

func (cfg *Config) getLegacyProviderConfig() *ProviderCfg { _ = "STUB: not implemented"; return nil }

// getProviderConfigs returns a slice of ProviderCfg, including the legacy provider configuration
// if it exists. The legacy configuration is prepended to the list of providers.
// Prefer this function over accessing the Providers field directly.
func (cfg *Config) getProviderConfigs() []ProviderCfg { _ = "STUB: not implemented"; return nil }

func (cfg *Config) Validate() (errs error) { _ = "STUB: not implemented"; return nil }

type ProviderCfg struct {
	// IssuerURL is the base URL for the OIDC provider.
	// Required.
	IssuerURL string `mapstructure:"issuer_url"`

	// Audience of the token, used during the verification.
	// For example: "https://accounts.google.com" or "https://login.salesforce.com".
	// Required unless IgnoreAudience is true.
	Audience string `mapstructure:"audience"`

	// When true, this skips validating the audience field.
	// Optional.
	IgnoreAudience bool `mapstructure:"ignore_audience"`

	// The local path for the issuer CA's TLS server cert.
	// Optional.
	IssuerCAPath string `mapstructure:"issuer_ca_path"`

	// The claim to use as the username, in case the token's 'sub' isn't the suitable source.
	// Optional.
	UsernameClaim string `mapstructure:"username_claim"`

	// The claim that holds the subject's group membership information.
	// Optional.
	GroupsClaim string `mapstructure:"groups_claim"`

	// Path to a local JWKS file containing public keys for token verification.
	// When provided, only keys from this file will be used - no remote key discovery will happen.
	// The file is watched for changes and keys are automatically reloaded on updates.
	// Optional.
	PublicKeysFile string `mapstructure:"public_keys_file"`
}

func (p *ProviderCfg) Validate() error { _ = "STUB: not implemented"; return nil }

var (
	supportedAlgorithms = map[string][]jose.SignatureAlgorithm{
		"rsa": {
			jose.RS256,
			jose.RS384,
			jose.RS512,
			jose.PS256,
			jose.PS384,
			jose.PS512,
		},
		"ecdsa": {
			jose.ES256,
			jose.ES384,
			jose.ES512,
		},
		"ed25519": {
			jose.EdDSA,
		},
	}

	allSupportedAlgorithms = slices.Concat(slices.Collect(maps.Values(supportedAlgorithms))...)
)

// staticPublicKey is a convenience struct used in the return value of
// parseJWKSFile. We cannot use ED25519 public keys as a map key (they
// are byte slices and thus unhashable), so this wrapper struct makes things
// a tiny bit more ergonomic.
type staticPublicKey struct {
	publicKey           crypto.PublicKey
	supportedAlgorithms []jose.SignatureAlgorithm
}

// parseJWKSFile reads and parses a JWKS file, returning a list of
// public keys and the signatures supported for each (based on the key's
// `alg` field if present, or the full list of algorithms possible for
// the key type if the field is not present). At the moment, only RSA,
// ECDSA, and ED25519 keys are supported, which is what go-oidc supports.
func parseJWKSFile(path string) ([]staticPublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
