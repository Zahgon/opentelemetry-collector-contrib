// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheusreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusreceiver"

import (
	commonconfig "github.com/prometheus/common/config"
	promconfig "github.com/prometheus/prometheus/config"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/confmap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusreceiver/internal/targetallocator"
)

// Config defines configuration for Prometheus receiver.
type Config struct {
	PrometheusConfig   *PromConfig `mapstructure:"config"`
	TrimMetricSuffixes bool        `mapstructure:"trim_metric_suffixes"`

	TargetAllocator configoptional.Optional[targetallocator.Config] `mapstructure:"target_allocator"`

	//  APIServer has the settings to enable the receiver to host the Prometheus API
	// server in agent mode. This allows the user to call the endpoint to get
	// the config, service discovery, and targets for debugging purposes.
	APIServer APIServer `mapstructure:"api_server"`

	// For testing only.
	ignoreMetadata bool
	skipOffsetting bool
}

// Validate checks the receiver configuration is valid.
func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }

// PromConfig is a redeclaration of promconfig.Config because we need custom unmarshaling
// as prometheus "config" uses `yaml` tags.
type PromConfig promconfig.Config

var _ confmap.Unmarshaler = (*PromConfig)(nil)

// ContainsScrapeConfigs returns true if the Prometheus config contains any scrape configs.
func (cfg *PromConfig) ContainsScrapeConfigs() bool { _ = "STUB: not implemented"; return false }

func (cfg *PromConfig) Reload() error { _ = "STUB: not implemented"; return nil }

func (cfg *PromConfig) Unmarshal(componentParser *confmap.Conf) error {
	_ = "STUB: not implemented"
	return nil
}

func (cfg *PromConfig) Validate() error {
	_ = "STUB: not implemented"
	// Reject features that Prometheus supports but that the receiver doesn't support:
	// See:
	// * https://github.com/open-telemetry/opentelemetry-collector/issues/3863
	// * https://github.com/open-telemetry/wg-prometheus/issues/3
	return nil
}

// Sort the values for deterministic error messages.

// Since Prometheus 3.0, the scrape manager started to fail scrapes that don't have proper
// Content-Type headers, but they provided an extra configuration option to fallback to the
// previous behavior. We need to make sure that this option is set for all scrape configs
// to avoid introducing a breaking change.

// copyStaticConfig copies static service discovery configs from src to dst to assure that labels in StaticConfig are retained.
func copyStaticConfig(dst *PromConfig, src any) {
	_ = "STUB: not implemented"
	// only deal with PromConfig
	return
}

// Job name -> static config list
// The static configs are grouped by job name, so that we can
// copy them over to the destination config.

// Remove all static configs for this job name.

func reloadPromConfig(dst *PromConfig, src any) error { _ = "STUB: not implemented"; return nil }

func validateHTTPClientConfig(cfg *commonconfig.HTTPClientConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func checkFile(fn string) error {
	_ = "STUB: not implemented"
	// Nothing set, nothing to error on.
	return nil
}

func checkTLSConfig(tlsConfig commonconfig.TLSConfig) error { _ = "STUB: not implemented"; return nil }

type APIServer struct {
	Enabled      bool                    `mapstructure:"enabled"`
	ServerConfig confighttp.ServerConfig `mapstructure:"server_config"`
}

func (cfg *APIServer) Validate() error { _ = "STUB: not implemented"; return nil }
