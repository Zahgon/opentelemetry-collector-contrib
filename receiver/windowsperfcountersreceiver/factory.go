// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package windowsperfcountersreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/windowsperfcountersreceiver"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/receiver"
)

// This file implements Factory for WindowsPerfCounters receiver.

// NewFactory creates a new factory for windows perf counters receiver.
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

// createDefaultConfig creates the default configuration for receiver.
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}
