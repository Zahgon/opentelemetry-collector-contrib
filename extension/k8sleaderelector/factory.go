// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8sleaderelector // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/k8sleaderelector"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

const (
	defaultLeaseDuration = 15 * time.Second
	defaultRenewDeadline = 10 * time.Second
	defaultRetryPeriod   = 2 * time.Second
)

// createDefaultConfig returns the default configuration for the extension.
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// createExtension creates the extension instance based on the configuration.
func createExtension(
	_ context.Context,
	set extension.Settings,
	cfg component.Config,
) (extension.Extension, error) {
	_ = "STUB: not implemented"
	return *new(extension.Extension), nil
}

// Initialize k8s client in factory as doing it in extension.Start()
// should cause race condition as http Proxy gets shared.

// NewFactory creates a new factory for your extension.
func NewFactory() extension.Factory { _ = "STUB: not implemented"; return *new(extension.Factory) }
