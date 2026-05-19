// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !windows

package tracker // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/tracker"

// On non-windows platforms, we keep files open between poll cycles so that we can detect
// and read "lost" files, which have been moved out of the matching pattern.
func (t *fileTracker) EndConsume() (filesClosed int) { _ = "STUB: not implemented"; return 0 }

// t.currentPollFiles -> t.previousPollFiles
