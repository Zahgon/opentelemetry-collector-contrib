// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package osqueryreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/osqueryreceiver"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/scraper/scraperhelper"
)

const (
	defaultSocket = "/var/osquery/osquery.em"
)

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

type Config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`
	ExtensionsSocket               string   `mapstructure:"extensions_socket"`
	Queries                        []string `mapstructure:"queries"`
}

func (c Config) Validate() error { _ = "STUB: not implemented"; return nil }
