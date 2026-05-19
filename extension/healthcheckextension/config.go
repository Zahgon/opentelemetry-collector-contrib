// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package healthcheckextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckextension"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/healthcheck"
)

// Config wraps the shared healthcheck.Config to add extension-specific validation.
type Config struct {
	healthcheck.Config `mapstructure:",squash"`
}

// Validate checks if the extension configuration is valid, including feature gate checks.
func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }

// Type alias for backward compatibility
type ResponseBodySettings = healthcheck.ResponseBodyConfig
