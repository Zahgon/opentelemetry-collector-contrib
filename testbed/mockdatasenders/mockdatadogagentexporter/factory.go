// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package mockdatadogagentexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/mockdatasenders/mockdatadogagentexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
)

// This file implements factory for awsxray receiver.

// The value of "type" key in configuration.
var compType = component.MustNewType("datadog")

func NewFactory() exporter.Factory { _ = "STUB: not implemented"; return *new(exporter.Factory) }

// CreateDefaultConfig creates the default configuration for DDAPM Exporter
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func CreateTracesExporter(
	_ context.Context,
	set exporter.Settings,
	cfg component.Config,
) (exporter.Traces, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Traces), nil
}

// TODO https://github.com/open-telemetry/opentelemetry-collector/issues/215

// explicitly disable since we rely on http.Client timeout logic.
