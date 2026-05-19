// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sigv4authextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/sigv4authextension"

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"go.opentelemetry.io/collector/component"
)

// Config stores the configuration for the Sigv4 Authenticator
type Config struct {
	Region        string     `mapstructure:"region,omitempty"`
	Service       string     `mapstructure:"service,omitempty"`
	AssumeRole    AssumeRole `mapstructure:"assume_role"`
	credsProvider *aws.CredentialsProvider
}

// AssumeRole holds the configuration needed to assume a role
type AssumeRole struct {
	ARN                  string `mapstructure:"arn,omitempty"`
	SessionName          string `mapstructure:"session_name,omitempty"`
	STSRegion            string `mapstructure:"sts_region,omitempty"`
	WebIdentityTokenFile string `mapstructure:"web_identity_token_file,omitempty"`
	ExternalID           string `mapstructure:"external_id,omitempty"`
}

// compile time check that the Config struct satisfies the component.Config interface
var _ component.Config = (*Config)(nil)

// Validate checks that the configuration is valid.
// We aim to catch most errors here to ensure that we
// fail early and to avoid revalidating static data.
func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }
