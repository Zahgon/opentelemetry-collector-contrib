// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ecsobserver // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecsobserver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

// NewFactory creates a factory for ECSObserver extension.
func NewFactory() extension.Factory { _ = "STUB: not implemented"; return *new(extension.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createExtension(ctx context.Context, params extension.Settings, cfg component.Config) (extension.Extension, error) {
	_ = "STUB: not implemented"
	return *new(extension.Extension), nil
}

// fetcher is mock in unit test or AWS API client
func createExtensionWithFetcher(params extension.Settings, sdCfg *Config, fetcher *taskFetcher) (extension.Extension, error) {
	_ = "STUB: not implemented"
	return *new(extension.Extension), nil
}
