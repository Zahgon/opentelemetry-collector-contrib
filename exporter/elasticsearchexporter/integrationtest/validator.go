// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package integrationtest // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/integrationtest"

import (
	"testing"

	"github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"
)

// countValidator provides a testbed validator that only asserts for counts.
type countValidator struct {
	t            testing.TB
	dataProvider testbed.DataProvider
}

// newCountValidator creates a new instance of the CountValidator.
func newCountValidator(tb testing.TB, provider testbed.DataProvider) *countValidator {
	_ = "STUB: not implemented"
	return nil
}

func (v *countValidator) Validate(tc *testbed.TestCase) { _ = "STUB: not implemented"; return }

func (*countValidator) RecordResults(*testbed.TestCase) { _ = "STUB: not implemented"; return }
