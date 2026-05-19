// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package fileexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/fileexporter"

import (
	"io"
	"sync"
	"time"
)

// exportFunc defines how to export encoded telemetry data.
type exportFunc func(e *fileWriter, buf []byte) error

type fileWriter struct {
	path  string
	file  io.WriteCloser
	mutex sync.Mutex

	exporter exportFunc

	flushInterval time.Duration
	flushTicker   *time.Ticker
	stopTicker    chan struct{}
}

func exportMessageAsLine(w *fileWriter, buf []byte) error {
	_ = "STUB: not implemented"
	// Ensure only one write operation happens at a time.
	return nil
}

func exportMessageAsBuffer(w *fileWriter, buf []byte) error {
	_ = "STUB: not implemented"
	// Ensure only one write operation happens at a time.
	return nil
}

// write the size of each message before writing the message itself.  https://developers.google.com/protocol-buffers/docs/techniques
// each encoded object is preceded by 4 bytes (an unsigned 32 bit integer)

func (w *fileWriter) export(buf []byte) error { _ = "STUB: not implemented"; return nil }

// startFlusher starts the flusher.
// It does not check the flushInterval
func (w *fileWriter) startFlusher() { _ = "STUB: not implemented"; return }

// Just in case.

// Create the stop channel.

// Start the ticker.

// Start starts the flush timer if set.
func (w *fileWriter) start() { _ = "STUB: not implemented"; return }

// Shutdown stops the exporter and is invoked during shutdown.
// It stops the flush ticker if set.
func (w *fileWriter) shutdown() error {
	_ = "STUB: not implemented"
	// Stop the flush ticker.
	return nil
}

// Stop the go routine.

func buildExportFunc(cfg *Config) func(w *fileWriter, buf []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// if the data format is JSON and needs to be compressed, telemetry data can't be written to file in JSON format.
