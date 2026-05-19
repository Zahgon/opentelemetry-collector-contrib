// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package oauth2clientauthextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/oauth2clientauthextension"

import (
	"context"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	grantTypeClientCredentials = "client_credentials"
)

func newClientCredentialsGrantTypeConfig(cfg *Config) *clientCredentialsConfig {
	_ = "STUB: not implemented"
	return nil
}

// clientCredentialsConfig is a clientcredentials.Config wrapper to allow
// values read from files in the ClientID and ClientSecret fields.
//
// Values from files can be retrieved by populating the ClientIDFile or
// the ClientSecretFile fields with the path to the file.
//
// Priority: File > Raw value
//
// Example - Retrieve secret from file:
//
//	cfg := clientCredentialsConfig{
//		Config: clientcredentials.Config{
//			ClientID:     "clientId",
//			...
//		},
//		ClientSecretFile: "/path/to/client/secret",
//	}
type clientCredentialsConfig struct {
	clientcredentials.Config

	ClientIDFile     string
	ClientSecretFile string
	ExpiryBuffer     time.Duration
}

type clientCredentialsTokenSource struct {
	ctx    context.Context
	config *clientCredentialsConfig
}

// clientCredentialsTokenSource implements TokenSource
var _ oauth2.TokenSource = (*clientCredentialsTokenSource)(nil)

func readCredentialsFile(path string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func getActualValue(value, filepath string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// createConfig creates a proper clientcredentials.Config with values retrieved
// from files, if the user has specified '*_file' values
func (c *clientCredentialsConfig) createConfig() (*clientcredentials.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *clientCredentialsConfig) TokenSource(ctx context.Context) oauth2.TokenSource {
	_ = "STUB: not implemented"
	return *new(oauth2.TokenSource)
}

func (c *clientCredentialsConfig) TokenEndpoint() string { _ = "STUB: not implemented"; return "" }

func (ts clientCredentialsTokenSource) Token() (*oauth2.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
