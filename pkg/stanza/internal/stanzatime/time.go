// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package stanzatime // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/internal/stanzatime"

import (
	"time"

	"github.com/jonboulle/clockwork"
)

var (
	Now   = time.Now
	Since = time.Since
)

// Clock where Now() always returns a greater value than the previous return value
type AlwaysIncreasingClock struct {
	*clockwork.FakeClock
}

func NewAlwaysIncreasingClock() AlwaysIncreasingClock {
	_ = "STUB: not implemented"
	return *new(AlwaysIncreasingClock)
}

func (c AlwaysIncreasingClock) Now() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (c AlwaysIncreasingClock) Since(t time.Time) time.Duration {
	_ = "STUB: not implemented"
	// ensure that internal c.FakeClock.Now() will return a greater value
	return *new(time.Duration)
}

func (c AlwaysIncreasingClock) Advance(d time.Duration) { _ = "STUB: not implemented"; return }
