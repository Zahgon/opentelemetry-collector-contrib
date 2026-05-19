// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build linux

package journald // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/journald"

import (
	"context"
	"io"
	"sync"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

// Input is an operator that process logs using journald
type Input struct {
	helper.InputOperator

	newCmd func(ctx context.Context, cursor []byte) cmd

	persister           operator.Persister
	convertMessageBytes bool
	cancel              context.CancelFunc
	wg                  sync.WaitGroup
	errChan             chan error
}

type cmd interface {
	StdoutPipe() (io.ReadCloser, error)
	StderrPipe() (io.ReadCloser, error)
	Start() error
	Wait() error
}

type journalctl struct {
	cmd    cmd
	stdout io.ReadCloser
	stderr io.ReadCloser
}

var lastReadCursorKey = "lastReadCursor"

// Start will start generating log entries.
func (operator *Input) Start(persister operator.Persister) error {
	_ = "STUB: not implemented"
	return nil
}

// run starts the journalctl process and monitor it.
// If there is an error in operator.newJournalctl, the error will be sent to operator.errChan.
// If the journalctl process started successfully, but there is an error in operator.runJournalctl,
// The error will be logged and the journalctl process will be restarted.
func (operator *Input) run(ctx context.Context) { _ = "STUB: not implemented"; return }

// If we can't start journalctl, there is nothing we can do but logging the error and return.

// Backoff before restart.

// newJournalctl creates a new journalctl command.
func (operator *Input) newJournalctl(ctx context.Context) (*journalctl, error) {
	_ = "STUB: not implemented"
	// Start from a cursor if there is a saved offset
	return nil, nil
}

// runJournalctl runs the journalctl command. This is a blocking call that returns
// when the command exits.
func (operator *Input) runJournalctl(ctx context.Context, jctl *journalctl) error {
	_ = "STUB: not implemented"
	// Start the stderr reader goroutine.
	// This goroutine reads the stderr from the journalctl process. If the
	// process exits for any reason, then the stderr will be closed, this
	// goroutine will get an EOF error and exit.
	return nil
}

// Start the reader goroutine.
// This goroutine reads the stdout from the journalctl process, parses
// the data, and writes to output. If the journalctl process exits for
// any reason, then the stdout will be closed, this goroutine will get
// an EOF error and exits.

// we wait for the reader goroutines to exit before calling Cmd.Wait().
// As per documentation states, "It is thus incorrect to call Wait before all reads from the pipe have completed".

func (operator *Input) parseJournalEntry(line []byte) (*entry.Entry, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

// Convert the message bytes to string if given as a byte array

// in microseconds

// Stop will stop generating logs.
func (operator *Input) Stop() error { _ = "STUB: not implemented"; return nil }
