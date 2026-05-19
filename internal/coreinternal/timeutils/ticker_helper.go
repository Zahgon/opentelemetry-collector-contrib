// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package timeutils // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/timeutils"

import "time"

// TTicker interface allows easier testing of Ticker related functionality
type TTicker interface {
	// start sets the frequency of the Ticker and starts the periodic calls to OnTick.
	Start(d time.Duration)
	// OnTick is called when the Ticker fires.
	OnTick()
	// Stop firing the Ticker.
	Stop()
}

// Implements TTicker and abstracts underlying time ticker's functionality to make usage
// simpler.
type PolicyTicker struct {
	Ticker     *time.Ticker
	OnTickFunc func()
	StopCh     chan struct{}
}

// Ensure PolicyTicker implements TTicker interface
var _ TTicker = (*PolicyTicker)(nil)

func (pt *PolicyTicker) Start(d time.Duration) { _ = "STUB: not implemented"; return }

func (pt *PolicyTicker) OnTick() { _ = "STUB: not implemented"; return }

func (pt *PolicyTicker) Stop() { _ = "STUB: not implemented"; return }
