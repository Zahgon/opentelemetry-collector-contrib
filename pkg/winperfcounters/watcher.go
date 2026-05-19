// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package winperfcounters // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/winperfcounters"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/winperfcounters/internal/third_party/telegraf/win_perf_counters"
)

const totalInstanceName = "_Total"

var _ PerfCounterWatcher = (*perfCounter)(nil)

// PerfCounterWatcher represents how to scrape data
type PerfCounterWatcher interface {
	// Path returns the counter path
	Path() string
	// ScrapeData collects a measurement and returns the value(s).
	ScrapeData() ([]CounterValue, error)
	// ScrapeRawValue collects a measurement and returns the raw value.
	ScrapeRawValue(rawValue *int64) (bool, error)
	// ScrapeRawValues collects a measurement and returns the raw value(s) for all instances.
	ScrapeRawValues() ([]RawCounterValue, error)
	// Resets the perfcounter query.
	Reset() error
	// Close all counters/handles related to the query and free all associated memory.
	Close() error
}

type (
	CounterValue    = win_perf_counters.CounterValue
	RawCounterValue = win_perf_counters.RawCounterValue
)

type perfCounter struct {
	path   string
	query  win_perf_counters.PerformanceQuery
	handle win_perf_counters.PDH_HCOUNTER
}

// NewWatcher creates new PerfCounterWatcher by provided parts of its path.
func NewWatcher(object, instance, counterName string) (PerfCounterWatcher, error) {
	_ = "STUB: not implemented"
	return *new(PerfCounterWatcher), nil
}

// NewWatcherFromPath creates new PerfCounterWatcher by provided path.
func NewWatcherFromPath(path string) (PerfCounterWatcher, error) {
	_ = "STUB: not implemented"
	return *new(PerfCounterWatcher), nil
}

func counterPath(object, instance, counterName string) string { _ = "STUB: not implemented"; return "" }

// newPerfCounter returns a new performance counter for the specified descriptor.
func newPerfCounter(counterPath string, collectOnStartup bool) (*perfCounter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func initQuery(counterPath string, collectOnStartup bool) (*win_perf_counters.PerformanceQueryImpl, *win_perf_counters.PDH_HCOUNTER, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Some perf counters (e.g. cpu) return the usage stats since the last measure.
// We collect data on startup to avoid an invalid initial reading

// Ignore PDH_NO_DATA error, it is expected when there are no
// matching instances.

// Reset re-creates the PerformanceCounter query and if the operation succeeds, closes the previous query.
// This is useful when scraping wildcard counters.
func (pc *perfCounter) Reset() error { _ = "STUB: not implemented"; return nil }

func (pc *perfCounter) Close() error { _ = "STUB: not implemented"; return nil }

func (pc *perfCounter) Path() string { _ = "STUB: not implemented"; return "" }

func (pc *perfCounter) ScrapeData() ([]CounterValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pc *perfCounter) ScrapeRawValues() ([]RawCounterValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pc *perfCounter) ScrapeRawValue(rawValue *int64) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// ExpandWildCardPath examines the local computer and returns those counter paths that match the given counter path which contains wildcard characters.
func ExpandWildCardPath(counterPath string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getInstanceName(ctr any) string { _ = "STUB: not implemented"; return "" }

func setInstanceName(ctr any, name string) { _ = "STUB: not implemented"; return }

// cleanupScrapedValues handles instance name collisions and standardizes names.
// It cleans up the list in-place to avoid unnecessary copies.
func cleanupScrapedValues[C CounterValue | RawCounterValue](vals []C) []C {
	_ = "STUB: not implemented"
	return nil
}

// If there is only one "_Total" instance, clear the instance name.

// Remember if a "_Total" instance was present.

// Append indices to duplicate instance names.

// Remove the "_Total" instance, as it can be computed with a sum aggregation.

func removeItemAt[C CounterValue | RawCounterValue](vals []C, idx int) []C {
	_ = "STUB: not implemented"
	return nil
}

func (pc *perfCounter) collectDataForScrape() (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// No data is available for the counter, so no error but also no data

// A counter rolled over, so the value is invalid
// See https://support.microfocus.com/kb/doc.php?id=7010545
// Wait one second and retry once
