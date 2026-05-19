// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !windows

package fileconsumer // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer"

import (
	"context"
	"os"
)

// Take care of files which disappeared from the pattern since the last poll cycle
// this can mean either files which were removed, or rotated into a name not matching the pattern
// we do this before reading existing files to ensure we emit older log lines before newer ones
func (m *Manager) readLostFiles(ctx context.Context) { _ = "STUB: not implemented"; return }

// Lost files are not expected when delete_at_eof is enabled
// since we are deleting the files before they can become lost.

// At this point, we know that the file has been rotated out of the matching pattern.
// However, we do not know if it was moved or truncated.
// If truncated, then both handles point to the same file, in which case
// we should only read from it using the new reader. We can use
// the Validate method to ensure that the file has not been truncated.

// oldreader points to the rotated file after the move/rename. We can still read from it.

// normalizePath cleans the path on non-Windows systems.
// Returns the cleaned path and false (no corruption detection on non-Windows).
func normalizePath(path string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func openFile(path string) (*os.File, error) {
	_ = "STUB: not implemented"
	// #nosec - operator must read in files defined by user
	return nil, nil
}
