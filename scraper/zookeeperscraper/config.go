// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package zookeeperscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/scraper/zookeeperscraper"

import (
	"go.opentelemetry.io/collector/config/confignet"

	"github.com/open-telemetry/opentelemetry-collector-contrib/scraper/zookeeperscraper/internal/metadata"
)

type Config struct {
	confignet.TCPAddrConfig       `mapstructure:",squash"`
	metadata.MetricsBuilderConfig `mapstructure:",squash"`
}

func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }
