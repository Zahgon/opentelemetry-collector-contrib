// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package logging contains utility functions for logging.
package logging // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/logging"

import (
	"time"

	"go.uber.org/zap"
)

// WithRateLimit returns a zap.Option which rate limits messages
// with approximately the given frequency.
// interval <= 0 disables any rate limiting.
func WithRateLimit(interval time.Duration) zap.Option {
	_ = "STUB: not implemented"
	return *new(zap.Option)
}
