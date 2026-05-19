// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate make mdatagen

//go:build !aix

package datadogconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/datadogconnector"

import (
	"go.opentelemetry.io/collector/connector"
)

// NewFactory creates a factory for datadog connector.
func NewFactory() connector.Factory {
	_ = "STUB: not implemented"
	// OTel connector factory to make a factory for connectors
	return *new(connector.Factory)
}
