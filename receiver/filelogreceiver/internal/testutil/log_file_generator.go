// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package testutil // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/filelogreceiver/internal/testutil"

import (
	"testing"
)

type LogFileGenerator struct {
	tb       testing.TB
	charset  []byte
	logLines [][]byte
}

func NewLogFileGenerator(tb testing.TB) *LogFileGenerator { _ = "STUB: not implemented"; return nil }

func (g *LogFileGenerator) generateLogLines(numLines, lineLength int) (logLines [][]byte) {
	_ = "STUB: not implemented"
	return nil
}

func (g *LogFileGenerator) GenerateLogFile(numLines int) (logFilePath string) {
	_ = "STUB: not implemented"
	return ""
}
