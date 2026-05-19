// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics

import (
	"sync"
	"time"

	"go.opentelemetry.io/otel/attribute"
)

/*
timeBox allows for only one metric timeseries per second. It'll generate additional timeseries if the attribute is
requested within a one-second window
*/
type timeBox struct {
	// enforceUnique is set to true if the attribute should be unique
	// for each second. If false, the attribute will always be 0.
	enforceUnique bool

	// locks attribute creation for a moment once a second
	mutex *sync.Mutex

	// used to gracefully shut down the timer
	stop chan struct{}

	// The attribute value last used
	offset int64
}

const timeBoxAttributeName = "timebox"

func newTimeBox(enforceUnique bool, uniqueTimeLimit time.Duration) *timeBox {
	_ = "STUB: not implemented"
	return nil
}

func (tb *timeBox) getAttribute() attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

// resetTimer resets the timer to 1 second every time the timer is stopped.
func (tb *timeBox) resetTimerLoop(t *time.Timer) {
	_ = "STUB: not implemented"
	// no-op when enforceUnique is false
	return
}

func (tb *timeBox) shutdown() { _ = "STUB: not implemented"; return }
