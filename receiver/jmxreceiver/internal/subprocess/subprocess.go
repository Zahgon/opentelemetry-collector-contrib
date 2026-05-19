// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package subprocess // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jmxreceiver/internal/subprocess"

import (
	"bufio"
	"context"
	"io"
	"os/exec"
	"sync"
	"sync/atomic"
	"time"

	"go.uber.org/zap"
)

const (
	defaultRestartDelay    = 5 * time.Second
	defaultShutdownTimeout = 5 * time.Second
	noPid                  = -1
)

// Config exported to be used by jmx metric receiver.
type Config struct {
	ExecutablePath       string            `mapstructure:"executable_path"`
	Args                 []string          `mapstructure:"args"`
	EnvironmentVariables map[string]string `mapstructure:"environment_variables"`
	StdInContents        string            `mapstructure:"stdin_contents"`
	RestartOnError       bool              `mapstructure:"restart_on_error"`
	RestartDelay         *time.Duration    `mapstructure:"restart_delay"`
	ShutdownTimeout      *time.Duration    `mapstructure:"shutdown_timeout"`
}

// Subprocess exported to be used by jmx metric receiver.
type Subprocess struct {
	Stdout         chan string
	cancel         context.CancelFunc
	config         *Config
	envVars        []string
	logger         *zap.Logger
	pid            pid
	shutdownSignal chan struct{}
	// configurable for testing purposes
	sendToStdIn func(string, io.Writer) error
}

type pid struct {
	pid     int
	pidLock sync.Mutex
}

func (p *pid) setPid(pid int) { _ = "STUB: not implemented"; return }

func (p *pid) getPid() int { _ = "STUB: not implemented"; return 0 }

func (subprocess *Subprocess) Pid() int { _ = "STUB: not implemented"; return 0 }

// NewSubprocess exported to be used by jmx metric receiver.
func NewSubprocess(conf *Config, logger *zap.Logger) *Subprocess {
	_ = "STUB: not implemented"
	return nil
}

const (
	starting     = "Starting"
	running      = "Running"
	shuttingDown = "ShuttingDown"
	stopped      = "Stopped"
	restarting   = "Restarting"
	errored      = "Errored"
)

func (subprocess *Subprocess) Start(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// will block for lifetime of process

// Shutdown is invoked during service shutdown.
func (subprocess *Subprocess) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Wait for the subprocess to exit or the timeout period to elapse

// A synchronization helper to ensure that signalWhenProcessReturned
// doesn't write to a closed channel
type processReturned struct {
	ReturnedChan chan error
	isOpen       *atomic.Bool
	lock         *sync.Mutex
}

func newProcessReturned() *processReturned { _ = "STUB: not implemented"; return nil }

func (pr *processReturned) signal(err error) { _ = "STUB: not implemented"; return }

func (pr *processReturned) close() { _ = "STUB: not implemented"; return }

// Core event loop
func (subprocess *Subprocess) run(ctx context.Context) { _ = "STUB: not implemented"; return }

// writer is signalWhenProcessReturned() and closer is this loop, so we need synchronization

// We aren't supposed to shutdown yet so this is an error state.

// We must close this channel or can wait indefinitely at shuttingDown

// context-based cancel.

// We must close this channel or can wait indefinitely at shuttingDown

func signalWhenProcessReturned(cmd *exec.Cmd, pr *processReturned) {
	_ = "STUB: not implemented"
	return
}

func collectStdout(stdoutScanner *bufio.Scanner, stdoutChan chan<- string, logger *zap.Logger) {
	_ = "STUB: not implemented"
	return
}

// Returns when stdout is closed when the process ends

func sendToStdIn(contents string, writer io.Writer) error { _ = "STUB: not implemented"; return nil }

func createCommand(execPath string, args, envVars []string) (*exec.Cmd, io.WriteCloser, io.ReadCloser) {
	_ = "STUB: not implemented"
	return nil, *new(io.WriteCloser), *new(io.ReadCloser)
}
