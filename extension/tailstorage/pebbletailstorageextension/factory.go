// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pebbletailstorageextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/tailstorage/pebbletailstorageextension"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

// NewFactory creates a factory for the Pebble tail storage extension.
func NewFactory() extension.Factory { _ = "STUB: not implemented"; return *new(extension.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createExtension(_ context.Context, settings extension.Settings, cfg component.Config) (extension.Extension, error) {
	_ = "STUB: not implemented"
	return *new(extension.Extension), nil
}
