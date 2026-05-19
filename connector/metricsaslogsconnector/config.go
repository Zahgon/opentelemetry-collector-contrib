// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metricsaslogsconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/metricsaslogsconnector"

import (
	"go.opentelemetry.io/collector/component"
)

type Config struct {
	IncludeResourceAttributes bool `mapstructure:"include_resource_attributes"`

	IncludeScopeInfo bool `mapstructure:"include_scope_info"`

	_ struct{}
}

func (*Config) Validate() error { _ = "STUB: not implemented"; return nil }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}
