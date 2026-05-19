// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package finder // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/matcher/internal/finder"

import (
	"github.com/bmatcuk/doublestar/v4"
)

func getDefaultDoublestarOptions() []doublestar.GlobOption {
	_ = "STUB: not implemented"
	// On Windows, filepaths are case-insensitive by default. As a result,
	// we want our globs to be case-insensitive.
	return nil
}

// This is currently guarded by a featuregate, which will eventually become
// the default.

func pathExcluded(excludes []string, path string) bool {
	_ = "STUB: not implemented"
	// To allow case-insensitive matching, the path and exclude
	// are unified to lowercase before matching.
	return false
}

// fixUNCPath corrects UNC path corruption that occurs when doublestar's path.Join
// collapses // to /. If the pattern starts with \\ (UNC) but the match only has \,
// we restore the UNC prefix.
func fixUNCPath(pattern, match string) string {
	_ = "STUB: not implemented"
	// Check if pattern is a UNC path (starts with \\)
	return ""
}

// Check if match is corrupted (starts with single \ instead of \\)

// Restore the missing backslash
