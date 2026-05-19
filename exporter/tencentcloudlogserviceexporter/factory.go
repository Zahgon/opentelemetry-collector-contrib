// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate make mdatagen

package tencentcloudlogserviceexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/tencentcloudlogserviceexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
)

// NewFactory creates a factory for tencentcloud LogService exporter.
func NewFactory() exporter.Factory { _ = "STUB: not implemented"; return *new(exporter.Factory) }

// CreateDefaultConfig creates the default configuration for exporter.
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createLogsExporter(
	_ context.Context,
	set exporter.Settings,
	cfg component.Config,
) (exp exporter.Logs, err error) {
	_ = "STUB: not implemented"
	return *new(exporter.Logs), nil
}
