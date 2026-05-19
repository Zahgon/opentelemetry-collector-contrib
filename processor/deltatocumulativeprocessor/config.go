// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package deltatocumulativeprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor"

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap/xconfmap"

	telemetry "github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor/internal/telemetry"
)

var _ xconfmap.Validator = (*Config)(nil)

type Config struct {
	MaxStale   time.Duration `mapstructure:"max_stale"`
	MaxStreams int           `mapstructure:"max_streams"`
}

func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// TODO: find good default
// https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/31603

func (c Config) Metrics(tel telemetry.Metrics) { _ = "STUB: not implemented"; return }
