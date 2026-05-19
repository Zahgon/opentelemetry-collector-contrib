// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package hostobserver // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/hostobserver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

const (
	defaultCollectionInterval = 10
)

// NewFactory creates a factory for HostObserver extension.
func NewFactory() extension.Factory { _ = "STUB: not implemented"; return *new(extension.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createExtension(
	_ context.Context,
	params extension.Settings,
	cfg component.Config,
) (extension.Extension, error) {
	_ = "STUB: not implemented"
	return *new(extension.Extension), nil
}
