// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package correlation // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/correlation"

import (
	"time"

	"go.opentelemetry.io/collector/config/confighttp"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/apm/correlations"
)

// DefaultConfig returns default configuration correlation values.
func DefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

// Config defines configuration for correlation via traces.
type Config struct {
	confighttp.ClientConfig `mapstructure:",squash"`
	correlations.Config     `mapstructure:",squash"`

	// How long to wait after a trace span's service name is last seen before
	// uncorrelating that service.
	StaleServiceTimeout time.Duration `mapstructure:"stale_service_timeout"`
	// SyncAttributes is a key of the span attribute name to sync to the dimension as the value.
	SyncAttributes map[string]string `mapstructure:"sync_attributes"`
}

func (c *Config) validate() error { _ = "STUB: not implemented"; return nil }
