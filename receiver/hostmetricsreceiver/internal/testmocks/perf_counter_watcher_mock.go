// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package testmocks // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/testmocks"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/winperfcounters"
)

type PerfCounterWatcherMock struct {
	Val       int64
	ScrapeErr error
	CloseErr  error
	closed    bool
}

// ScrapeRawValue implements winperfcounters.PerfCounterWatcher.
func (w *PerfCounterWatcherMock) ScrapeRawValue(rawValue *int64) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// ScrapeRawValues implements winperfcounters.PerfCounterWatcher.
func (w *PerfCounterWatcherMock) ScrapeRawValues() ([]winperfcounters.RawCounterValue, error) {
	_ = "STUB: not implemented"
	return nil,

		// ScrapeData returns scrapeErr if it's set, otherwise it returns a single countervalue with the mock's val
		nil
}

func (PerfCounterWatcherMock) ScrapeData() ([]winperfcounters.CounterValue, error) {
	_ = "STUB: not implemented"
	return nil,

		// Reset panics; it should not be called
		nil
}

func (PerfCounterWatcherMock) Reset() error { _ = "STUB: not implemented"; return nil }

// Path panics; It should not be called
func (PerfCounterWatcherMock) Path() string { _ = "STUB: not implemented"; return "" }

// Close all counters/handles related to the query and free all associated memory.
func (w *PerfCounterWatcherMock) Close() error { _ = "STUB: not implemented"; return nil }
