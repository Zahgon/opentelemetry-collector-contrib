// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/correctnesstests/metrics"

import (
	"os"
)

type results struct {
	resultsFile *os.File
}

type result struct {
	testName   string
	testResult string
	numDiffs   int
}

func (r *results) Init(resultsDir string) { _ = "STUB: not implemented"; return }

func (r *results) Add(_ string, rslt any) { _ = "STUB: not implemented"; return }

func (r *results) writeString(s string) { _ = "STUB: not implemented"; return }

func (r *results) Save() { _ = "STUB: not implemented"; return }
