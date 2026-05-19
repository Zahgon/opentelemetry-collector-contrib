// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package cfgardenobserver // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/cfgardenobserver"

import (
	"time"
)

// Config defines configuration for CF Garden observer.
type Config struct {
	// CloudFoundry API Configuration
	CloudFoundry CfConfig `mapstructure:"cloud_foundry"`

	// Garden API Configuration
	Garden GardenConfig `mapstructure:"garden"`

	// RefreshInterval determines the frequency at which the observer
	// needs to poll for collecting information about new processes.
	// Default: "1m"
	RefreshInterval time.Duration `mapstructure:"refresh_interval"`

	// The time to wait before resyncing app information on cached containers
	// using the CloudFoundry API.
	// Default: "5m"
	CacheSyncInterval time.Duration `mapstructure:"cache_sync_interval"`

	// Determines whether or not Application labels get added to the Endpoint labels.
	// This requires cloud_foundry to be configured, such that API calls can be made
	// Default: false
	IncludeAppLabels bool `mapstructure:"include_app_labels"`
}

// Validate overrides the embedded noop validation so that load config can trigger
// our own validation logic.
func (config *Config) Validate() error { _ = "STUB: not implemented"; return nil }

func fieldError(authType authType, param string) error { _ = "STUB: not implemented"; return nil }

type GardenConfig struct {
	// The URL of the CF Garden api. Default is "/var/vcap/data/garden/garden.sock"
	Endpoint string `mapstructure:"endpoint"`

	// prevent unkeyed literal initialization
	_ struct{}
}

type CfConfig struct {
	// The URL of the CloudFoundry API
	Endpoint string `mapstructure:"endpoint"`

	// Authentication details
	Auth CfAuth `mapstructure:"auth"`

	// prevent unkeyed literal initialization
	_ struct{}
}

type CfAuth struct {
	// Authentication method, there are 3 options
	Type authType `mapstructure:"type"`

	// Used for user_pass authentication method
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`

	// Used for token authentication method
	AccessToken  string `mapstructure:"access_token"`
	RefreshToken string `mapstructure:"refresh_token"`

	// Used for client_credentials authentication method
	ClientID     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
}

// authType describes the type of authentication to use for the CloudFoundry API
type authType string

const (
	// authTypeClientCredentials uses a client ID and client secret to authenticate
	authTypeClientCredentials authType = "client_credentials"
	// authTypeUserPass uses username and password to authenticate
	authTypeUserPass authType = "user_pass"
	// authTypeToken uses access token and refresh token to authenticate
	authTypeToken authType = "token"
)
