// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package testbed // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"

import (
	"go.opentelemetry.io/collector/otelcol"
)

// Components returns the set of components for tests
func Components() (
	otelcol.Factories,
	error,
) {
	_ = "STUB: not implemented"
	return *new(otelcol.Factories), nil
}
