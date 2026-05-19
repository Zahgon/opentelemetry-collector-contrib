// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package healthcheckextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckextension"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

const defaultPort = 13133

// NewFactory creates a factory for HealthCheck extension.
func NewFactory() extension.Factory { _ = "STUB: not implemented"; return *new(extension.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createExtension(_ context.Context, set extension.Settings, cfg component.Config) (extension.Extension, error) {
	_ = "STUB: not implemented"
	return *new(extension.Extension), nil
}

// When feature gate is enabled, use v2 implementation directly.
// The feature gate controls behavior, not the presence of v2 config fields.

// If no v2 config is set, create HTTP config from legacy settings for backward compatibility

// Feature gate disabled: use legacy implementation.
