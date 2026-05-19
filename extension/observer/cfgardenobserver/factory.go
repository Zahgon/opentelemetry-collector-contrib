// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package cfgardenobserver // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/cfgardenobserver"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

const (
	defaultCollectionInterval = 1 * time.Minute
	defaultCacheSyncInterval  = 5 * time.Minute
	defaultEndpoint           = "/var/vcap/data/garden/garden.sock"
)

// NewFactory creates a factory for CfGardenObserver extension.
func NewFactory() extension.Factory { _ = "STUB: not implemented"; return *new(extension.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createExtension(
	_ context.Context,
	settings extension.Settings,
	cfg component.Config,
) (extension.Extension, error) {
	_ = "STUB: not implemented"
	return *new(extension.Extension), nil
}
