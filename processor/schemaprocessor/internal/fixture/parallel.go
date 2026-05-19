// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package fixture // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor/internal/fixture"

import (
	"testing"
)

// ParallelRaceCompute starts `count` number of go routines that calls the provided function `fn`
// at the same to allow the race detector greater opportunity to capture known race conditions.
// This method blocks until each count number of fn has completed, any returned errors is considered
// a failing test method.
// If the race detector is not enabled, the function then skips with an notice.
// This is intended to show that a test was intentionally skipped instead of just missing.
func ParallelRaceCompute(tb testing.TB, count int, fn func() error) {
	_ = "STUB: not implemented"
	return
}
