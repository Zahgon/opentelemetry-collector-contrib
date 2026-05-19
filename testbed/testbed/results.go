// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package testbed // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"

import (
	"os"
	"time"
)

// TestResultsSummary defines the interface to record results of one category of testing.
type TestResultsSummary interface {
	// Init creates and open the file and write headers.
	Init(resultsDir string)
	// Add results for one test.
	Add(testName string, result any)
	// Save the total results and close the file.
	Save()
}

// benchmarkResult holds the results of a benchmark to be stored by benchmark-action. See
// https://github.com/benchmark-action/github-action-benchmark#examples for more details on the
// format
type benchmarkResult struct {
	Name  string  `json:"name"`
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
	Range string  `json:"range,omitempty"`
	Extra string  `json:"extra,omitempty"`
}

type LogPresentResults struct {
	testName string
	result   string
	duration time.Duration
}

// PerformanceResults implements the TestResultsSummary interface with fields suitable for reporting
// performance test results.
type PerformanceResults struct {
	resultsDir       string
	resultsFile      *os.File
	perTestResults   []*PerformanceTestResult
	benchmarkResults []*benchmarkResult
	totalDuration    time.Duration
}

// PerformanceTestResult reports the results of a single performance test.
type PerformanceTestResult struct {
	testName           string
	result             string
	duration           time.Duration
	cpuPercentageAvg   float64
	cpuPercentageMax   float64
	cpuPercentageLimit float64
	ramMibAvg          uint32
	ramMibMax          uint32
	ramMibLimit        uint32
	sentSpanCount      uint64
	receivedSpanCount  uint64
	errorCause         string
}

func (r *PerformanceResults) Init(resultsDir string) { _ = "STUB: not implemented"; return }

// Create resultsSummary file

// Write the header

// Save the total results and close the file.
func (r *PerformanceResults) Save() { _ = "STUB: not implemented"; return }

// Add results for one test.
func (r *PerformanceResults) Add(_ string, result any) { _ = "STUB: not implemented"; return }

// individual benchmark results

// saveBenchmarks writes benchmarks to file as json to be stored by
// benchmark-action
func (r *PerformanceResults) saveBenchmarks() { _ = "STUB: not implemented"; return }

// CorrectnessResults implements the TestResultsSummary interface with fields suitable for reporting data translation
// correctness test results.
type CorrectnessResults struct {
	resultsDir             string
	resultsFile            *os.File
	perTestResults         []*CorrectnessTestResult
	totalAssertionFailures uint64
	totalDuration          time.Duration
}

// CorrectnessTestResult reports the results of a single correctness test.
type CorrectnessTestResult struct {
	testName                   string
	result                     string
	duration                   time.Duration
	sentSpanCount              uint64
	receivedSpanCount          uint64
	traceAssertionFailureCount uint64
	traceAssertionFailures     []*TraceAssertionFailure
}

type TraceAssertionFailure struct {
	typeName      string
	dataComboName string
	fieldPath     string
	expectedValue any
	actualValue   any
	sumCount      int
}

func (af TraceAssertionFailure) String() string { _ = "STUB: not implemented"; return "" }

func (r *CorrectnessResults) Init(resultsDir string) { _ = "STUB: not implemented"; return }

// Create resultsSummary file

// Write the header

func (r *CorrectnessResults) Add(_ string, result any) { _ = "STUB: not implemented"; return }

func (r *CorrectnessResults) Save() { _ = "STUB: not implemented"; return }

func consolidateAssertionFailures(failures []*TraceAssertionFailure) map[string]*TraceAssertionFailure {
	_ = "STUB: not implemented"
	return nil
}
