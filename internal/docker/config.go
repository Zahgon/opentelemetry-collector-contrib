// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package docker // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/docker"

import (
	"time"

	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/config/configtls"
	"go.opentelemetry.io/collector/confmap"
)

type Config struct {
	// The URL of the docker server. Default is "unix:///var/run/docker.sock"
	// on non-Windows and "npipe:////./pipe/docker_engine" on Windows
	Endpoint string `mapstructure:"endpoint"`

	// The maximum amount of time to wait for docker API responses. Default is 5s
	Timeout time.Duration `mapstructure:"timeout"`

	// A list of filters whose matching images are to be excluded. Supports literals, globs, and regex.
	ExcludedImages []string `mapstructure:"excluded_images"`

	// Docker client API version. If empty, the client will auto-negotiate
	// the API version with the Docker daemon using version negotiation.
	DockerAPIVersion string `mapstructure:"api_version"`

	// TLS holds optional TLS client configuration for connecting to the Docker daemon
	// over HTTPS. When nil (the default), the connection uses no custom TLS — suitable
	// for Unix sockets and plain HTTP endpoints.
	TLS configoptional.Optional[configtls.ClientConfig] `mapstructure:"tls,omitempty"`
	// StreamStats enables a persistent streaming connection per container to collect stats.
	// When true, each container maintains an open Docker stats stream and the scraper reads
	// from the cached latest value, which reduces connection overhead.  When false (default),
	// a new connection is opened and closed on every scrape cycle, matching the original behavior.
	StreamStats bool `mapstructure:"stream_stats"`
}

func (config *Config) Unmarshal(conf *confmap.Conf) error {
	_ = "STUB: not implemented"
	// WithIgonreUnused needed because this configuration is embedded inside other configurations
	return nil
}

func (config Config) Validate() error { _ = "STUB: not implemented"; return nil }

// NewConfig creates a new config to be used when creating
// a docker client
func NewConfig(endpoint string, timeout time.Duration, excludedImages []string, apiVersion string) *Config {
	_ = "STUB: not implemented"
	return nil
}

// NewDefaultConfig creates a new config with default values
// to be used when creating a docker client
// DockerAPIVersion is intentionally left empty for auto-negotiation.
func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type apiVersion struct {
	major int
	minor int
}

func NewAPIVersion(version string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// MustNewAPIVersion evaluates version as a client api version and panics if invalid.
func MustNewAPIVersion(version string) string { _ = "STUB: not implemented"; return "" }

// VersionIsValidAndGTE evalutes version as a client api version and returns an error if invalid or less than gte.
// gte is assumed to be valid (easiest if result of MustNewAPIVersion on initialization)
func VersionIsValidAndGTE(version, gte string) error { _ = "STUB: not implemented"; return nil }
