// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package mcp // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/mcp"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

// NewFactory creates a factory for the mcp extension.
func NewFactory() extension.Factory { _ = "STUB: not implemented"; return *new(extension.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createExtension(_ context.Context, set extension.Settings, cfg component.Config) (extension.Extension, error) {
	_ = "STUB: not implemented"
	return *new(extension.Extension), nil
}
