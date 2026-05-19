// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package unrollprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/unrollprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
)

type unrollProcessor struct {
	cfg *Config
}

// newUnrollProcessor returns a new unrollProcessor.
func newUnrollProcessor(config *Config) (*unrollProcessor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ProcessLogs implements the processor interface
func (p *unrollProcessor) ProcessLogs(_ context.Context, ld plog.Logs) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

// setBody will set the body of the log record to the provided value
func setBody(newLogRecord plog.LogRecord, expansion pcommon.Value) error {
	_ = "STUB: not implemented"
	return nil
}
