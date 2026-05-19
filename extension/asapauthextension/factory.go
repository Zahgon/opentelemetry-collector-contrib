// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package asapauthextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/asapauthextension"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

// NewFactory creates a factory for asapauthextension.
func NewFactory() extension.Factory { _ = "STUB: not implemented"; return *new(extension.Factory) }

func createExtension(_ context.Context, _ extension.Settings, cfg component.Config) (extension.Extension, error) {
	_ = "STUB: not implemented"
	return *new(extension.Extension), nil
}

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}
