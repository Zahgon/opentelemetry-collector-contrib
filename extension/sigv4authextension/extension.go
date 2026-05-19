// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sigv4authextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/sigv4authextension"

import (
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/extension/extensionauth"
	"go.uber.org/zap"
)

// sigv4Auth is a struct that implements the extensionauth.HTTPClient interface.
// It provides the implementation for providing Sigv4 authentication for HTTP requests only.
type sigv4Auth struct {
	cfg                    *Config
	logger                 *zap.Logger
	awsSDKInfo             string
	component.StartFunc    // embedded default behavior to do nothing with Start()
	component.ShutdownFunc // embedded default behavior to do nothing with Shutdown()
}

// compile time check that the sigv4Auth struct satisfies the extensionauth.HTTPClient interface
var (
	_ extension.Extension      = (*sigv4Auth)(nil)
	_ extensionauth.HTTPClient = (*sigv4Auth)(nil)
)

// RoundTripper() returns a custom signingRoundTripper.
func (sa *sigv4Auth) RoundTripper(base http.RoundTripper) (http.RoundTripper, error) {
	_ = "STUB: not implemented"
	return *new(http.RoundTripper), nil
}

// Create the signingRoundTripper struct

// newSigv4Extension() is called by createExtension() in factory.go and
// returns a new sigv4Auth struct.
func newSigv4Extension(cfg *Config, awsSDKInfo string, logger *zap.Logger) *sigv4Auth {
	_ = "STUB: not implemented"
	return nil
}

// getCredsProviderFromConfig() is a helper function that gets AWS credentials
// from the Config.
func getCredsProviderFromConfig(cfg *Config) (*aws.CredentialsProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getCredsProviderFromWebIdentityConfig(cfg *Config) (*aws.CredentialsProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func assumeRoleOptions(cfg *Config) func(*stscreds.AssumeRoleOptions) {
	_ = "STUB: not implemented"
	return nil
}
