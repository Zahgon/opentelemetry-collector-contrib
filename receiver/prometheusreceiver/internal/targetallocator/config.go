// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package targetallocator // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusreceiver/internal/targetallocator"

import (
	"time"

	commonconfig "github.com/prometheus/common/config"
	promHTTP "github.com/prometheus/prometheus/discovery/http"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/confmap"
)

type Config struct {
	confighttp.ClientConfig `mapstructure:",squash"`
	Interval                time.Duration         `mapstructure:"interval"`
	CollectorID             string                `mapstructure:"collector_id"`
	HTTPSDConfig            *PromHTTPSDConfig     `mapstructure:"http_sd_config"`
	HTTPScrapeConfig        *PromHTTPClientConfig `mapstructure:"http_scrape_config"`
}

// PromHTTPSDConfig is a redeclaration of promHTTP.SDConfig because we need custom unmarshaling
// as prometheus "config" uses `yaml` tags.
type PromHTTPSDConfig promHTTP.SDConfig

func (cfg *Config) Validate() error {
	_ = "STUB: not implemented"
	// ensure valid endpoint
	return nil
}

// ensure valid collectorID without variables

// ensure valid interval

var _ confmap.Unmarshaler = (*PromHTTPSDConfig)(nil)

func (cfg *PromHTTPSDConfig) Unmarshal(componentParser *confmap.Conf) error {
	_ = "STUB: not implemented"
	return nil
}

// we have to set it as else marshaling will fail

type PromHTTPClientConfig commonconfig.HTTPClientConfig

var _ confmap.Unmarshaler = (*PromHTTPClientConfig)(nil)

func (cfg *PromHTTPClientConfig) Unmarshal(componentParser *confmap.Conf) error {
	_ = "STUB: not implemented"
	return nil
}

func (cfg *PromHTTPClientConfig) Validate() error { _ = "STUB: not implemented"; return nil }

// Prometheus UnmarshalYaml implementation by default calls Validate,
// but it is safer to do it here as well.

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

// unmarshalConf unmarshals conf to out.
// YAML is used as an intermediary format, because confmap.Conf.Unmarshal is incompatible with Prometheus types,
// as the former uses mapstructure tags .
// If cb is not nil, it's called to mutate the map representation of conf.
func unmarshalConf(conf *confmap.Conf, cb func(map[string]any), out any) error {
	_ = "STUB: not implemented"
	return nil
}

// convertTLSVersion converts a string TLS version to the corresponding config.TLSVersion value in prometheus common.
func convertTLSVersion(version string) (commonconfig.TLSVersion, error) {
	_ = "STUB: not implemented"
	return *new(commonconfig.TLSVersion), nil
}

// configureSDHTTPClientConfigFromTA configures the http client for the service discovery manager
// based on the provided TargetAllocator configuration.
func configureSDHTTPClientConfigFromTA(httpSD *promHTTP.SDConfig, allocConf *Config) error {
	_ = "STUB: not implemented"
	return nil
}
