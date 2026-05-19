// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package filetest // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/internal/filetest"

import (
	"os"
	"testing"
)

func OpenFile(tb testing.TB, path string) *os.File { _ = "STUB: not implemented"; return nil }

func OpenTemp(tb testing.TB, tempDir string) *os.File { _ = "STUB: not implemented"; return nil }

func ReopenTemp(tb testing.TB, name string) *os.File { _ = "STUB: not implemented"; return nil }

func OpenTempWithPattern(tb testing.TB, tempDir, pattern string) *os.File {
	_ = "STUB: not implemented"
	return nil
}

func WriteString(tb testing.TB, file *os.File, s string) { _ = "STUB: not implemented"; return }

func TokenWithLength(length int) []byte { _ = "STUB: not implemented"; return nil }
