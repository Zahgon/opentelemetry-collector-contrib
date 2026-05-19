// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package zipkinexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/zipkinexporter"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
)

const (
	defaultTimeout = time.Second * 5

	defaultFormat = "json"

	defaultServiceName string = "<missing service name>"
)

// NewFactory creates a factory for Zipkin exporter.
func NewFactory() exporter.Factory { _ = "STUB: not implemented"; return *new(exporter.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createTracesExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg component.Config,
) (exporter.Traces, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Traces), nil
}

// explicitly disable since we rely on http.Client timeout logic.
