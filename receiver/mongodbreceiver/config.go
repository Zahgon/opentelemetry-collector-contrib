// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package mongodbreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbreceiver"

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.opentelemetry.io/collector/config/confignet"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/config/configtls"
	"go.opentelemetry.io/collector/scraper/scraperhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbreceiver/internal/metadata"
)

type Config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`
	configtls.ClientConfig         `mapstructure:"tls,omitempty"`
	// MetricsBuilderConfig defines which metrics/attributes to enable for the scraper
	metadata.MetricsBuilderConfig `mapstructure:",squash"`
	// Deprecated - Transport option will be removed in v0.102.0
	Hosts                   []confignet.TCPAddrConfig `mapstructure:"hosts"`
	Scheme                  string                    `mapstructure:"scheme"`
	Username                string                    `mapstructure:"username"`
	Password                configopaque.String       `mapstructure:"password"`
	AuthMechanism           string                    `mapstructure:"auth_mechanism,omitempty"`
	AuthSource              string                    `mapstructure:"auth_source,omitempty"`
	AuthMechanismProperties map[string]string         `mapstructure:"auth_mechanism_properties,omitempty"`
	ReplicaSet              string                    `mapstructure:"replica_set,omitempty"`
	Timeout                 time.Duration             `mapstructure:"timeout"`
	DirectConnection        bool                      `mapstructure:"direct_connection"`
}

func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }

func (c *Config) ClientOptions(secondary bool) *options.ClientOptions {
	_ = "STUB: not implemented"

	// For secondary nodes, create a direct connection
	return nil
}

// Set up authentication if username/password are provided or if an auth mechanism is specified
// Some mechanisms (e.g., MONGODB-X509, MONGODB-AWS with IAM) don't require username/password

// Set up authentication if username/password are provided or if an auth mechanism is specified
// Some mechanisms (e.g., MONGODB-X509, MONGODB-AWS with IAM) don't require username/password

func (c *Config) buildCredential() options.Credential {
	_ = "STUB: not implemented"
	return *new(options.Credential)
}

// PasswordSet is required for GSSAPI (Kerberos) when a password is explicitly provided.
// For other mechanisms, this field is ignored by the driver.

func (c *Config) hostlist() []string { _ = "STUB: not implemented"; return nil }
