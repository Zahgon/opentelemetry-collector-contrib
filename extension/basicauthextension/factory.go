// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package basicauthextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/basicauthextension"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

// NewFactory creates a factory for the static bearer token Authenticator extension.
func NewFactory() extension.Factory { _ = "STUB: not implemented"; return *new(extension.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createExtension(_ context.Context, set extension.Settings, cfg component.Config) (extension.Extension, error) {
	_ = "STUB: not implemented"
	// check if config is a server auth(Htpasswd should be set)
	return *new(extension.Extension), nil
}
