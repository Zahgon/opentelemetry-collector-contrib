// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package yanggrpcreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/yanggrpcreceiver"

import (
	"go.opentelemetry.io/collector/receiver"
)

// createValidTestConfig creates a config suitable for testing
func createValidTestConfig() *Config { _ = "STUB: not implemented"; return nil }

// createTestSettings creates proper receiver settings for testing
func createTestSettings() receiver.Settings {
	_ = "STUB: not implemented"
	return *new(receiver.Settings)
}
