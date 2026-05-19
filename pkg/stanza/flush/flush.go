// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package flush // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/flush"

import (
	"bufio"
	"time"
)

type State struct {
	LastDataChange time.Time
	LastDataLength int
}

// Func wraps a bufio.SplitFunc with a timer.
// When the timer expires, an incomplete token may be returned.
// The timer will reset any time the data parameter changes.
func (s *State) Func(splitFunc bufio.SplitFunc, period time.Duration) bufio.SplitFunc {
	_ = "STUB: not implemented"
	return *new(bufio.SplitFunc)
}

// Don't interfere with errors

// If there's a token, return it

// Can't flush something from nothing

// We're seeing new data so postpone the next flush

// Flush timed out

// Ask for more data
