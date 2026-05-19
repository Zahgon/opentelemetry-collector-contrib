// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/correctnesstests/metrics"

type resultsDir struct {
	dir string
}

func newResultsDir(dirName string) (*resultsDir, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *resultsDir) mkDir() error { _ = "STUB: not implemented"; return nil }

func (d *resultsDir) fullPath(name string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
