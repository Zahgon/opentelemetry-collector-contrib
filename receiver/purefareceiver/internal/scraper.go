// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/purefareceiver/internal"

import (
	"context"
	"time"

	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/config"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configtls"
	"go.opentelemetry.io/collector/receiver"
)

type Scraper interface {
	ToPrometheusReceiverConfig(host component.Host, fact receiver.Factory) ([]*config.ScrapeConfig, error)
}

type ScraperType string

const (
	ScraperTypeArray       ScraperType = "array"
	ScraperTypeHosts       ScraperType = "hosts"
	ScraperTypeDirectories ScraperType = "directories"
	ScraperTypePods        ScraperType = "pods"
	ScraperTypeVolumes     ScraperType = "volumes"
)

type scraper struct {
	scraperType    ScraperType
	endpoint       string
	namespace      string
	tlsSettings    configtls.ClientConfig
	configs        []ScraperConfig
	scrapeInterval time.Duration
	labels         model.LabelSet
}

func NewScraper(_ context.Context,
	scraperType ScraperType,
	endpoint string,
	namespace string,
	tlsSettings configtls.ClientConfig,
	configs []ScraperConfig,
	scrapeInterval time.Duration,
	labels model.LabelSet,
) Scraper {
	_ = "STUB: not implemented"
	return *new(Scraper)
}

func (h *scraper) ToPrometheusReceiverConfig(host component.Host, _ receiver.Factory) ([]*config.ScrapeConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
