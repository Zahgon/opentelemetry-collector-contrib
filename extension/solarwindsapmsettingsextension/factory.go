// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package solarwindsapmsettingsextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/solarwindsapmsettingsextension"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

func createExtension(_ context.Context, settings extension.Settings, cfg component.Config) (extension.Extension, error) {
	_ = "STUB: not implemented"
	return *new(extension.Extension), nil
}

func NewFactory() extension.Factory { _ = "STUB: not implemented"; return *new(extension.Factory) }
