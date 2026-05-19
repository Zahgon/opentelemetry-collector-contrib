// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
//
//go:generate make mdatagen

package sumologicextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/sumologicextension"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

const (
	// The value of extension "type" in configuration.
	DefaultAPIBaseURL = "https://open-collectors.sumologic.com"
)

// NewFactory creates a factory for Sumo Logic extension.
func NewFactory() extension.Factory { _ = "STUB: not implemented"; return *new(extension.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createExtension(_ context.Context, params extension.Settings, cfg component.Config) (extension.Extension, error) {
	_ = "STUB: not implemented"
	return *new(extension.Extension), nil
}
