// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
package coralogixexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/coralogixexporter"

import (
	"sync/atomic"
	"time"
)

type rateError struct {
	rateLimited   atomic.Bool
	enabled       bool
	timestamp     atomic.Pointer[time.Time]
	internalError atomic.Pointer[error]
	errorCount    atomic.Int32
	threshold     int
	duration      time.Duration
}

func (r *rateError) isRateLimited() bool { _ = "STUB: not implemented"; return false }

// canDisableRateLimit checks if we can disable the limiter in the exporter.
func (r *rateError) canDisableRateLimit() bool { _ = "STUB: not implemented"; return false }

func (r *rateError) enableRateLimit() { _ = "STUB: not implemented"; return }

func (r *rateError) disableRateLimit() { _ = "STUB: not implemented"; return }

func (r *rateError) GetError() error { _ = "STUB: not implemented"; return nil }
