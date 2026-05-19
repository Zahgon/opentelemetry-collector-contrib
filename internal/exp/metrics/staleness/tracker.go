// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package staleness // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/exp/metrics/staleness"

import (
	"time"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/exp/metrics/identity"
)

type Tracker struct {
	pq PriorityQueue
}

func NewTracker() Tracker { _ = "STUB: not implemented"; return *new(Tracker) }

func (tr Tracker) Refresh(ts time.Time, ids ...identity.Stream) { _ = "STUB: not implemented"; return }

func (tr Tracker) Collect(maxDuration time.Duration) []identity.Stream {
	_ = "STUB: not implemented"
	return nil
}
