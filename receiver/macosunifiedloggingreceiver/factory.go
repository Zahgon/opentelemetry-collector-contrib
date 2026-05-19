// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package macosunifiedloggingreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/macosunifiedloggingreceiver"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/receiver"
)

// NewFactory creates a factory for the macOS unified logging receiver
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

// createDefaultConfig creates a config with default values
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	// Default to live mode
	return *new(component.Config)
}
