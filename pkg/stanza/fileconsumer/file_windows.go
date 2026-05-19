// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package fileconsumer // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer"

import (
	"context"
	"os"
)

// Noop on windows because we close files immediately after reading.
func (*Manager) readLostFiles(context.Context) {
	_ = "STUB: not implemented"

	// normalizePath ensures Windows UNC paths are properly formatted for os.Open().
	// It converts UNC paths to extended-length format (\\?\UNC\server\share\path)
	// for reliable file access on Windows. Extended-length paths bypass path parsing
	// and provide more reliable access to network shares.
	// Returns the normalized path and a boolean (always false, kept for API compatibility).
	return
}

func normalizePath(path string) (string, bool) { _ = "STUB: not implemented"; return "", false }

// Already in extended-length format

// Convert proper UNC paths (\\server\share) to extended-length format
// Note: We check for exactly 2 leading backslashes to identify UNC paths

// Extract the path after the \\ prefix and clean it

// For non-UNC paths (including paths starting with single backslash which are
// valid local paths on the current drive), just clean normally

// openFile opens a file on Windows with FILE_SHARE_DELETE flag to allow
// other processes to delete or rename the file while it's open.
func openFile(path string) (*os.File, error) { _ = "STUB: not implemented"; return nil, nil }
