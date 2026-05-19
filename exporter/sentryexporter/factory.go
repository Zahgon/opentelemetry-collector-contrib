// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate make mdatagen

package sentryexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sentryexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/sharedcomponent"
)

// NewFactory creates a factory for Sentry exporter.
func NewFactory() exporter.Factory { _ = "STUB: not implemented"; return *new(exporter.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createTracesExporter(
	ctx context.Context,
	set exporter.Settings,
	config component.Config,
) (exporter.Traces, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Traces), nil
}

func createLogsExporter(
	ctx context.Context,
	set exporter.Settings,
	config component.Config,
) (exporter.Logs, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Logs), nil
}

// getOrCreateEndpointState creates an endpointState and caches it for a particular configuration.
func getOrCreateEndpointState(cfg component.Config, set exporter.Settings) (*sharedcomponent.SharedComponent, *endpointState, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

var states = sharedcomponent.NewSharedComponents()
