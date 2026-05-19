// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package commander

import (
	"context"
	"os/exec"
	"sync/atomic"

	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/cmd/opampsupervisor/supervisor/config"
)

// Commander can start/stop/restart the Agent executable and also watch for a signal
// for the Agent process to finish.
type Commander struct {
	logger  *zap.Logger
	cfg     config.Agent
	logsDir string
	args    []string
	cmd     *exec.Cmd
	doneCh  chan struct{}
	exitCh  chan struct{}
	running *atomic.Int64
}

func NewCommander(logger *zap.Logger, logsDir string, cfg config.Agent, args ...string) (*Commander, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Buffer channels so we can send messages without blocking on listeners.

// Start the Agent and begin watching the process.
// Agent's stdout and stderr are written to a file.
// Calling this method when a command is already running
// is a no-op.
func (c *Commander) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Already started, nothing to do

// Drain channels in case there are no listeners that
// drained messages from previous runs.

// #nosec G204

// PassthroughLogging changes how collector start up happens

func (c *Commander) Restart(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *Commander) ReloadConfigFile() error { _ = "STUB: not implemented"; return nil }

// ValidateConfig runs the collector's validate command on the specified configuration file
// to check if the configuration is valid without starting the collector.
// Returns an error if the configuration is invalid or if the validation process fails.
func (c *Commander) ValidateConfig(ctx context.Context, configPath string) error {
	_ = "STUB: not implemented"
	return nil
}

// #nosec G204

func (c *Commander) startNormal() error { _ = "STUB: not implemented"; return nil }

// Capture standard output and standard error.
// https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/21072

func (c *Commander) startWithPassthroughLogging() error {
	_ = "STUB: not implemented"
	// grab cmd pipes
	return nil
}

// start agent

// capture agent output

// Trim and log the last line if it exists

// Trim and log the last line if it exists

func (c *Commander) watch() {
	_ = "STUB: not implemented"

	// cmd.Wait returns an exec.ExitError when the Collector exits unsuccessfully or stops
	// after receiving a signal. The Commander caller will handle these cases, so we filter
	// them out here.
	return
}

// StartOneShot starts the Collector with the expectation that it will immediately
// exit after it finishes a quick operation. This is useful for situations like reading stdout/sterr
// to e.g. check the feature gate the Collector supports.
func (c *Commander) StartOneShot() ([]byte, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// #nosec G204

// grab cmd pipes

// start agent

// capture agent output

// Trim and append the last line if it exists
// Normalize line endings to \n

// Normalize line endings to \n

// Trim and append the last line if it exists
// Normalize line endings to \n

// Normalize line endings to \n

// Gracefully signal process to stop.

// Setup a goroutine to wait a while for process to finish and send kill signal
// to the process if it doesn't finish.

// Time is out. Kill the process.

// Exited returns a channel that will send a signal when the Agent process exits.
func (c *Commander) Exited() <-chan struct{} {
	_ = "STUB: not implemented"

	// Pid returns Agent process PID if it is started or 0 if it is not.
	return nil
}

func (c *Commander) Pid() int { _ = "STUB: not implemented"; return 0 }

// ExitCode returns Agent process exit code if it exited or 0 if it is not.
func (c *Commander) ExitCode() int { _ = "STUB: not implemented"; return 0 }

func (c *Commander) IsRunning() bool { _ = "STUB: not implemented"; return false }

// Stop the Agent process. Sends SIGTERM to the process and wait for up 10 seconds
// and if the process does not finish kills it forcedly by sending SIGKILL.
// Returns after the process is terminated.
func (c *Commander) Stop(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Not started, nothing to do.

// Gracefully signal process to stop.

// Setup a goroutine to wait a while for process to finish and send kill signal
// to the process if it doesn't finish.

// Time is out. Kill the process.

// Wait for process to terminate

// Let goroutine know process is finished.

func envVarMapToEnvMapSlice(m map[string]string) []string {
	_ = "STUB: not implemented"
	// let the command initialize the env itself
	return nil
}
