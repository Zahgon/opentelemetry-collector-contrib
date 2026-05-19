// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package tracker // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/tracker"

// On windows, we close files immediately after reading because they cannot be moved while open.
func (t *fileTracker) EndConsume() (filesClosed int) {
	_ = "STUB: not implemented"
	// t.currentPollFiles -> t.previousPollFiles
	return 0
}
