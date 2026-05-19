// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package dockerobserver // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/dockerobserver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

// NewFactory should be called to create a factory with default values.
func NewFactory() extension.Factory { _ = "STUB: not implemented"; return *new(extension.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createExtension(
	_ context.Context,
	settings extension.Settings,
	cfg component.Config,
) (extension.Extension, error) {
	_ = "STUB: not implemented"
	return *new(extension.Extension), nil
}
