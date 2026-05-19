// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package testbed // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"

import (
	"os/exec"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/process"
)

// childProcessCollector implements the OtelcolRunner interface as a child process on the same machine executing
// the test. The process can be monitored and the output of which will be written to a log file.
type childProcessCollector struct {
	// Path to agent executable. If unset the default executable in
	// bin/otelcol_{{.GOOS}}_{{.GOARCH}} will be used.
	// Can be set for example to use the unstable executable for a specific test.
	agentExePath string

	// Descriptive name of the process
	name string

	// Config file name
	configFileName string

	// Command to execute
	cmd *exec.Cmd

	// additional env vars (os.Environ() populated by default)
	additionalEnv map[string]string

	// Various starting/stopping flags
	isStarted  bool
	stopOnce   sync.Once
	isStopped  bool
	doneSignal chan struct{}

	// Resource specification that must be monitored for.
	resourceSpec *ResourceSpec

	// Process monitoring data.
	processMon *process.Process

	// Time when process was started.
	startTime time.Time

	// Last tick time we monitored the process.
	lastElapsedTime time.Time

	// Process times that were fetched on last monitoring tick.
	lastProcessTimes *cpu.TimesStat

	// Current RAM RSS in MiBs
	ramMiBCur atomic.Uint32

	// Current CPU percentage times 1000 (we use scaling since we have to use int for atomic operations).
	cpuPercentX1000Cur atomic.Uint32

	// Maximum CPU seen
	cpuPercentMax float64

	// Number of memory measurements
	memProbeCount int

	// Cumulative RAM RSS in MiBs
	ramMiBTotal uint64

	// Maximum RAM seen
	ramMiBMax uint32
}

type ChildProcessOption func(*childProcessCollector)

// NewChildProcessCollector creates a new OtelcolRunner as a child process on the same machine executing the test.
func NewChildProcessCollector(options ...ChildProcessOption) OtelcolRunner {
	_ = "STUB: not implemented"
	return *new(OtelcolRunner)
}

// WithAgentExePath sets the path of the Collector executable
func WithAgentExePath(exePath string) ChildProcessOption {
	_ = "STUB: not implemented"
	return *new(ChildProcessOption)
}

// WithEnvVar sets an additional environment variable for the process
func WithEnvVar(k, v string) ChildProcessOption {
	_ = "STUB: not implemented"
	return *new(ChildProcessOption)
}

func (cp *childProcessCollector) PrepareConfig(t *testing.T, configStr string) (configCleanup func(), err error) {
	_ = "STUB: not implemented"
	return nil,

		// NoOp
		nil
}

func expandExeFileName(exeName string) string { _ = "STUB: not implemented"; return "" }

// Start a child process.
//
// cp.AgentExePath defines the executable to run. If unspecified
// "../../bin/otelcol_{{.GOOS}}_{{.GOARCH}}" will be used.
// {{.GOOS}} and {{.GOARCH}} will be expanded to the current OS and ARCH correspondingly.
//
// Parameters:
// name is the human readable name of the process (e.g. "Agent"), used for logging.
// logFilePath is the file path to write the standard output and standard error of
// the process to.
// cmdArgs is the command line arguments to pass to the process.
func (cp *childProcessCollector) Start(params StartParams) error {
	_ = "STUB: not implemented"
	return nil
}

// Prepare log file

// Prepare to start the process.
// #nosec

// #nosec

// update env deterministically

// Capture standard output and standard error.

// Start the process.

func (cp *childProcessCollector) Stop() (stopped bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Process wasn't started, nothing to stop.

// Notify resource monitor to stop.

// Gracefully signal process to stop.

// Setup a goroutine to wait a while for process to finish and send kill signal
// to the process if it doesn't finish.

// Wait 10 seconds.

// Time is out. Kill the process.

// Process is successfully finished.

// Wait for process to terminate

// Let goroutine know process is finished.

// Set resource consumption stats to 0

func (cp *childProcessCollector) WatchResourceConsumption() error {
	_ = "STUB: not implemented"
	return nil
}

// Resource monitoring is not enabled.

// Begin measuring elapsed and process CPU times.

// Measure every ResourceCheckPeriod.

// on first start must be under the cpu and ram max usage add a max minute delay

func (cp *childProcessCollector) GetProcessMon() *process.Process {
	_ = "STUB: not implemented"
	return nil
}

func (cp *childProcessCollector) fetchRAMUsage() {
	_ = "STUB: not implemented"
	// Get process memory and CPU times
	return
}

// Calculate RSS in MiBs.

// Calculate aggregates.

// Store current usage.

func (cp *childProcessCollector) fetchCPUUsage() { _ = "STUB: not implemented"; return }

// Calculate elapsed and process CPU time deltas in seconds

// We sometimes get negative difference when the process is terminated.

// Calculate CPU usage percentage in elapsed period.

// Store current usage.

func (cp *childProcessCollector) checkAllowedResourceUsage() error {
	_ = "STUB: not implemented"
	// Check if current CPU usage exceeds expected.
	return nil
}

// Check if current RAM usage exceeds expected.

// GetResourceConsumption returns resource consumption as a string
func (cp *childProcessCollector) GetResourceConsumption() string {
	_ = "STUB: not implemented"
	return ""
}

// Monitoring is not enabled.

// GetTotalConsumption returns total resource consumption since start of process
func (cp *childProcessCollector) GetTotalConsumption() *ResourceConsumption {
	_ = "STUB: not implemented"
	return nil
}

// Get total elapsed time since process start

// Calculate average CPU usage since start of process

// Calculate average RAM usage by averaging all RAM measurements

func containsConfig(s []string) bool { _ = "STUB: not implemented"; return false }

// Copied from cpu.TimesStat.Total(), since that func is deprecated.
func totalCPU(c *cpu.TimesStat) float64 { _ = "STUB: not implemented"; return 0 }
