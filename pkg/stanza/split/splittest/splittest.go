// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package splittest // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/split/splittest"

import (
	"bufio"
	"testing"
	"time"
)

type Step struct {
	tick     time.Duration
	timeout  time.Duration
	validate func(t *testing.T, advance int, token []byte, err error)
}

func ExpectReadMore() Step { _ = "STUB: not implemented"; return *new(Step) }

func ExpectToken(expectToken string) Step { _ = "STUB: not implemented"; return *new(Step) }

func ExpectAdvanceToken(expectAdvance int, expectToken string) Step {
	_ = "STUB: not implemented"
	return *new(Step)
}

func ExpectAdvanceNil(expectAdvance int) Step { _ = "STUB: not implemented"; return *new(Step) }

func ExpectError(expectErr string) Step { _ = "STUB: not implemented"; return *new(Step) }

func Eventually(step Step, maxTime, tick time.Duration) Step {
	_ = "STUB: not implemented"
	return *new(Step)
}

func New(splitFunc bufio.SplitFunc, input []byte, steps ...Step) func(*testing.T) {
	_ = "STUB: not implemented"
	return nil
}

// Split funcs do not have control over the size of the
// buffer so must be able to ask for more data as needed.
// Start with a tiny buffer and grow it slowly to ensure
// the split func is capable of asking appropriately.

// Grow the buffer at a slow pace to ensure that we're
// exercising the split func's ability to ask for more data.

// t.Errorf("\nbuffer: %d, advance: %d, token: %q, err: %v", bufferSize, advance, token, err)

func needMoreData(advance int, token []byte, err error) bool {
	_ = "STUB: not implemented"
	return false
}

// ScanLinesStrict behaves like bufio.ScanLines except EOF is not considered a line ending.
func ScanLinesStrict(data []byte, atEOF bool) (advance int, token []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func GenerateBytes(length int) []byte { _ = "STUB: not implemented"; return nil }
