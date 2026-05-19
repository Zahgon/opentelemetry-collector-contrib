// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package windowsservicereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/windowsservicereceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/windowsservicereceiver/internal/metadata"
)

type windowsServiceScraper struct {
	logger     *zap.Logger
	cfg        *Config
	mb         *metadata.MetricsBuilder
	mgr        serviceManager
	includeSet map[string]struct{}
	excludeSet map[string]struct{}

	disabled bool
}

func newWindowsServiceScraper(settings receiver.Settings, cfg *Config, mb *metadata.MetricsBuilder) *windowsServiceScraper {
	_ = "STUB: not implemented"
	return nil
}

func mapStartTypeToAttr(st StartType) metadata.AttributeStartupMode {
	_ = "STUB: not implemented"
	return *new(metadata.AttributeStartupMode)
}

func (ws *windowsServiceScraper) start(_ context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (ws *windowsServiceScraper) shutdown(_ context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (ws *windowsServiceScraper) allowed(name string) bool { _ = "STUB: not implemented"; return false }

func (ws *windowsServiceScraper) scrape(_ context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}
