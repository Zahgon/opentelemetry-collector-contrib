// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal"

import (
	"context"

	"github.com/shirou/gopsutil/v4/common"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/scraper"
)

// Config is the configuration of a scraper.
type Config interface {
	SetRootPath(rootPath string)
}

func NewEnvVarFactory(delegate scraper.Factory, envMap common.EnvMap) scraper.Factory {
	_ = "STUB: not implemented"
	return *new(scraper.Factory)
}

type envVarScraper struct {
	delegate scraper.Metrics
	envMap   common.EnvMap
}

func (evs *envVarScraper) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (evs *envVarScraper) ScrapeMetrics(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

func (evs *envVarScraper) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
