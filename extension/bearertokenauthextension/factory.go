// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package bearertokenauthextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/bearertokenauthextension"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

const (
	defaultHeader = "Authorization"
	defaultScheme = "Bearer"
)

// NewFactory creates a factory for the static bearer token Authenticator extension.
func NewFactory() extension.Factory { _ = "STUB: not implemented"; return *new(extension.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createExtension(_ context.Context, set extension.Settings, cfg component.Config) (extension.Extension, error) {
	_ = "STUB: not implemented"
	return *new(extension.Extension), nil
}
