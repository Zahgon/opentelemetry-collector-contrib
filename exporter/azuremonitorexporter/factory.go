// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate make mdatagen

package azuremonitorexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azuremonitorexporter"

import (
	"context"
	"errors"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/sharedcomponent"
)

var (
	errUnexpectedConfigurationType = errors.New("failed to cast configuration to Azure Monitor Config")
	exporters                      = sharedcomponent.NewSharedComponents()
)

// NewFactory returns a factory for Azure Monitor exporter.
func NewFactory() exporter.Factory { _ = "STUB: not implemented"; return *new(exporter.Factory) }

// Implements the interface from go.opentelemetry.io/collector/exporter/factory.go
type factory struct {
	loggerInitOnce sync.Once
}

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func (f *factory) createTracesExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg component.Config,
) (exporter.Traces, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Traces), nil
}

func (f *factory) createLogsExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg component.Config,
) (exporter.Logs, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Logs), nil
}

func (f *factory) createMetricsExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg component.Config,
) (exporter.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Metrics), nil
}

func getOrCreateAzureMonitorExporter(cfg component.Config, set exporter.Settings) *sharedcomponent.SharedComponent {
	_ = "STUB: not implemented"
	return nil
}

func (f *factory) initLogger(logger *zap.Logger) { _ = "STUB: not implemented"; return }
