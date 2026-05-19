// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package testbed // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"

import (
	"sync"
	"testing"
	"time"
)

// TestCase defines a running test case.
type TestCase struct {
	t *testing.T

	// Directory where test case results and logs will be written.
	resultDir string

	// does not write out results when set to true
	skipResults bool

	// Resource spec for agent.
	resourceSpec ResourceSpec

	// Agent process.
	agentProc OtelcolRunner

	receiver DataReceiver

	LoadGenerator LoadGenerator
	MockBackend   *MockBackend
	validator     TestCaseValidator

	startTime time.Time

	// errorSignal indicates an error in the test case execution, e.g. process execution
	// failure or exceeding resource consumption, etc. The actual error message is already
	// logged, this is only an indicator on which you can wait to be informed.
	errorSignal       chan struct{}
	errorSignalCloser *sync.Once
	// Duration is the requested duration of the tests. Configured via TESTBED_DURATION
	// env variable and defaults to 15 seconds if env variable is unspecified.
	Duration       time.Duration
	doneSignal     chan struct{}
	errorCause     string
	resultsSummary TestResultsSummary

	// decision makes mockbackend return permanent/non-permament errors at random basis
	decision decisionFunc
}

const (
	mibibyte            = 1024 * 1024
	testcaseDurationVar = "TESTCASE_DURATION"
)

// NewTestCase creates a new TestCase. It expects agent-config.yaml in the specified directory.
func NewTestCase(
	t *testing.T,
	dataProvider DataProvider,
	sender DataSender,
	receiver DataReceiver,
	agentProc OtelcolRunner,
	validator TestCaseValidator,
	resultsSummary TestResultsSummary,
	opts ...TestCaseOption,
) *TestCase {
	_ = "STUB: not implemented"
	return nil
}

func NewLoadGeneratorTestCase(t *testing.T, loadGenerator LoadGenerator, receiver DataReceiver, agentProc OtelcolRunner, validator TestCaseValidator, resultsSummary TestResultsSummary, opts ...TestCaseOption) *TestCase {
	_ = "STUB: not implemented"
	return nil
}

// Get requested test case duration from env variable.

// Apply all provided options.

// Prepare directory for results.

// Set default resource check period.

// Resource check period should not be longer than entire test duration.

func (tc *TestCase) ComposeTestResultFileName(fileName string) string {
	_ = "STUB: not implemented"
	return ""
}

// StartAgent starts the agent and redirects its standard output and standard error
// to "agent.log" file located in the test directory.
func (tc *TestCase) StartAgent(args ...string) { _ = "STUB: not implemented"; return }

// Start watching resource consumption.

// StopAgent stops agent process.
func (tc *TestCase) StopAgent() { _ = "STUB: not implemented"; return }

// StartLoad starts the load generator and redirects its standard output and standard error
// to "load-generator.log" file located in the test directory.
func (tc *TestCase) StartLoad(options LoadOptions) { _ = "STUB: not implemented"; return }

// StopLoad stops load generator.
func (tc *TestCase) StopLoad() { _ = "STUB: not implemented"; return }

// StartBackend starts the specified backend type.
func (tc *TestCase) StartBackend() { _ = "STUB: not implemented"; return }

// StopBackend stops the backend.
func (tc *TestCase) StopBackend() { _ = "STUB: not implemented"; return }

// EnableRecording enables recording of all data received by MockBackend.
func (tc *TestCase) EnableRecording() { _ = "STUB: not implemented"; return }

// AgentMemoryInfo returns raw memory info struct about the agent
// as returned by github.com/shirou/gopsutil/process
func (tc *TestCase) AgentMemoryInfo() (uint32, uint32, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// Stop stops the load generator, the agent and the backend.
func (tc *TestCase) Stop() {
	_ = "STUB: not implemented"
	// Stop monitoring the agent
	return
}

// Stop all components

// Report test results

// ValidateData validates data received by mock backend against what was generated and sent to the collector
// instance(s) under test by the ProviderSender.
func (tc *TestCase) ValidateData() { _ = "STUB: not implemented"; return }

// Error is already signaled and recorded. Validating data is pointless.

// Sleep for specified duration or until error is signaled.
func (tc *TestCase) Sleep(d time.Duration) { _ = "STUB: not implemented"; return }

// WaitForN the specific condition for up to a specified duration. Records a test error
// if time is out and condition does not become true. If error is signaled
// while waiting the function will return false, but will not record additional
// test error (we assume that signaled error is already recorded in indicateError()).
func (tc *TestCase) WaitForN(cond func() bool, duration time.Duration, errMsg any) bool {
	_ = "STUB: not implemented"
	return false

	// Start with 5 ms waiting interval between condition re-evaluation.
}

// Increase waiting interval exponentially up to 500 ms.

// Waited too long

// WaitFor is like WaitForN but with a fixed duration of 10 seconds
func (tc *TestCase) WaitFor(cond func() bool, errMsg any) bool {
	_ = "STUB: not implemented"
	return false
}

func (tc *TestCase) indicateError(err error) {
	_ = "STUB: not implemented"
	// Print for visibility but only set test error on first pass
	return
}

// Signal the error via channel

func (tc *TestCase) logStats() { _ = "STUB: not implemented"; return }

func (tc *TestCase) logStatsOnce() { _ = "STUB: not implemented"; return }

// Used to search for text in agent.log
// It can be used to verify if we've hit QueuedRetry sender or memory limiter
func (tc *TestCase) AgentLogsContains(text string) bool { _ = "STUB: not implemented"; return false }
