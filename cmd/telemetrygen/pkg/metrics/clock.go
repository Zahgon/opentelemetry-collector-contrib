// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics

import (
	"time"
)

type Clock interface {
	Now() time.Time
}

type realClock struct{}

func (*realClock) Now() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

type mockClock struct {
	now time.Time
}

func (c *mockClock) Now() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// Add 100ms to the mock clock to avoid timestamp collisions
