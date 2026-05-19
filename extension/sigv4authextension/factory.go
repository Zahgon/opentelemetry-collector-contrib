// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sigv4authextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/sigv4authextension"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

// NewFactory creates a factory for the Sigv4 Authenticator extension.
func NewFactory() extension.Factory { _ = "STUB: not implemented"; return *new(extension.Factory) }

// createDefaultConfig() creates a Config struct with default values.
// We only set the ID here.
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"

	// createExtension() calls newSigv4Extension() in extension.go to create the extension.
	return *new(component.Config)
}

func createExtension(_ context.Context, set extension.Settings, cfg component.Config) (extension.Extension, error) {
	_ = "STUB: not implemented"
	return *new(extension.Extension), nil
}
