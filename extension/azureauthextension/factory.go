// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azureauthextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/azureauthextension"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

func NewFactory() extension.Factory { _ = "STUB: not implemented"; return *new(extension.Factory) }

func createExtension(_ context.Context, set extension.Settings, cfg component.Config) (extension.Extension, error) {
	_ = "STUB: not implemented"
	return *new(extension.Extension), nil
}

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}
