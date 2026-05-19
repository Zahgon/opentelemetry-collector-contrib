// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8sobserver // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/k8sobserver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

// NewFactory should be called to create a factory with default values.
func NewFactory() extension.Factory { _ = "STUB: not implemented"; return *new(extension.Factory) }

// CreateDefaultConfig creates the default configuration for the extension.
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// CreateExtension creates the extension based on this config.
func createExtension(
	_ context.Context,
	params extension.Settings,
	cfg component.Config,
) (extension.Extension, error) {
	_ = "STUB: not implemented"
	return *new(extension.Extension), nil
}
