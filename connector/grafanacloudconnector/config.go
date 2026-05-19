// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package grafanacloudconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/grafanacloudconnector"

import (
	"time"

	"go.opentelemetry.io/collector/confmap/xconfmap"
)

// Config defines the configuration options for the Grafana Cloud connector.
type Config struct {
	// HostIdentifiers defines the list of resource attributes used to derive
	// a unique `grafana.host.id` value. In most cases, this should be [ "host.id" ]
	HostIdentifiers      []string      `mapstructure:"host_identifiers"`
	MetricsFlushInterval time.Duration `mapstructure:"metrics_flush_interval"`
	// prevent unkeyed literal initialization
	_ struct{}
}

var _ xconfmap.Validator = (*Config)(nil)

// Validate checks if the configuration is valid
func (c Config) Validate() error { _ = "STUB: not implemented"; return nil }
