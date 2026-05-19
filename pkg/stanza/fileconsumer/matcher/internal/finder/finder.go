// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package finder // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/matcher/internal/finder"

func Validate(globs []string) error { _ = "STUB: not implemented"; return nil }

// FindFiles gets a list of paths given an array of glob patterns to include and exclude
func FindFiles(includes, excludes []string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// the same pattern could cause an IO error due to one file or directory,
// but also could still find files without `doublestar.WithFailOnIOErrors()`.

// Fix UNC path corruption on Windows: if the include pattern starts with \\
// but the match only has \, restore the UNC prefix
