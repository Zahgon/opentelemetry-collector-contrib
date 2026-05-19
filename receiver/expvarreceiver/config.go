// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package expvarreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/expvarreceiver"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/scraper/scraperhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/expvarreceiver/internal/metadata"
)

type Config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`
	confighttp.ClientConfig        `mapstructure:",squash"`
	metadata.MetricsBuilderConfig  `mapstructure:",squash"`
}

var _ component.Config = (*Config)(nil)

func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }
