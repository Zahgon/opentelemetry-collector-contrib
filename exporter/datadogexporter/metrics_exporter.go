// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !aix

package datadogexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter"

import (
	"context"
	"sync"

	"github.com/DataDog/datadog-agent/pkg/opentelemetry-mapping-go/inframetadata"
	"github.com/DataDog/datadog-agent/pkg/opentelemetry-mapping-go/otlp/attributes"
	"github.com/DataDog/datadog-agent/pkg/opentelemetry-mapping-go/otlp/attributes/source"
	otlpmetrics "github.com/DataDog/datadog-agent/pkg/opentelemetry-mapping-go/otlp/metrics"
	"github.com/DataDog/datadog-agent/pkg/trace/config"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter/internal/metrics/sketches"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/datadog/clientutil"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/datadog/scrub"
	datadogconfig "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/datadog/config"
)

type metricsExporter struct {
	params           exporter.Settings
	cfg              *datadogconfig.Config
	agntConfig       *config.AgentConfig
	ctx              context.Context
	metricsAPI       *datadogV2.MetricsApi
	tr               otlpmetrics.Provider
	scrubber         scrub.Scrubber
	retrier          *clientutil.Retrier
	onceMetadata     *sync.Once
	sourceProvider   source.Provider
	metadataReporter *inframetadata.Reporter
	// getPushTime returns a Unix time in nanoseconds, representing the time pushing metrics.
	// It will be overwritten in tests.
	getPushTime func() uint64

	gatewayUsage *attributes.GatewayUsage
}

func newMetricsExporter(
	ctx context.Context,
	params exporter.Settings,
	cfg *datadogconfig.Config,
	agntConfig *config.AgentConfig,
	onceMetadata *sync.Once,
	attrsTranslator *attributes.Translator,
	sourceProvider source.Provider,
	metadataReporter *inframetadata.Reporter,
	statsOut chan []byte,
	gatewayUsage *attributes.GatewayUsage,
) (*metricsExporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Do nothing.

func (exp *metricsExporter) pushSketches(ctx context.Context, sl sketches.SketchSeriesList) error {
	_ = "STUB: not implemented"
	return nil
}

// We must read the full response body from the http request to ensure that connections can be
// properly re-used. https://pkg.go.dev/net/http#Client.Do

func (exp *metricsExporter) PushMetricsDataScrubbed(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (exp *metricsExporter) PushMetricsData(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

// Start host metadata with resource attributes from
// the first payload.

// Consume resources for host metadata
