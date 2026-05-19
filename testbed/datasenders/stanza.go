// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datasenders // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/datasenders"

import (
	"context"
	"net"
	"os"
	"testing"

	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"

	"github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"
)

type FileLogWriter struct {
	file    *os.File
	retry   string
	storage string
}

// Ensure FileLogWriter implements LogDataSender.
var _ testbed.LogDataSender = (*FileLogWriter)(nil)

// NewFileLogWriter creates a new data sender that will write log entries to a
// file, to be tailed by FluentBit and sent to the collector.
func NewFileLogWriter(t *testing.T) *FileLogWriter { _ = "STUB: not implemented"; return nil }

func (f *FileLogWriter) WithRetry(retry string) *FileLogWriter {
	_ = "STUB: not implemented"
	return nil
}

func (f *FileLogWriter) WithStorage(storage string) *FileLogWriter {
	_ = "STUB: not implemented"
	return nil
}

func (*FileLogWriter) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (*FileLogWriter) Start() error { _ = "STUB: not implemented"; return nil }

func (f *FileLogWriter) ConsumeLogs(_ context.Context, logs plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (*FileLogWriter) convertLogToTextLine(lr plog.LogRecord) []byte {
	_ = "STUB: not implemented"
	return nil

	// Timestamp
}

// Severity

func (f *FileLogWriter) Flush() { _ = "STUB: not implemented"; return }

func (f *FileLogWriter) GenConfigYAMLStr() string {
	_ = "STUB: not implemented"
	// Note that this generates a receiver config for agent.
	// We are testing stanza receiver here.
	return ""
}

func (*FileLogWriter) ProtocolName() string { _ = "STUB: not implemented"; return "" }

func (*FileLogWriter) GetEndpoint() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func NewLocalFileStorageExtension(t *testing.T) map[string]string {
	_ = "STUB: not implemented"
	return nil
}
