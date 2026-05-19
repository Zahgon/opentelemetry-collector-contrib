// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package remotetapextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/remotetapextension"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
)

type Config struct {
	confighttp.ServerConfig `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct
}

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}
