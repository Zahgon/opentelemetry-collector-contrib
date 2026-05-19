// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package googlecloudspannerreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver"

import (
	"context"
	_ "embed"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/metadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/statsreader"
)

//go:embed "internal/metadataconfig/metrics.yaml"
var metadataYaml []byte

var _ receiver.Metrics = (*googleCloudSpannerReceiver)(nil)

type googleCloudSpannerReceiver struct {
	logger         *zap.Logger
	config         *Config
	cancel         context.CancelFunc
	projectReaders []statsreader.CompositeReader
	metricsBuilder metadata.MetricsBuilder
}

func newGoogleCloudSpannerReceiver(logger *zap.Logger, config *Config) *googleCloudSpannerReceiver {
	_ = "STUB: not implemented"
	return nil
}

func (r *googleCloudSpannerReceiver) Scrape(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

func (r *googleCloudSpannerReceiver) Start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *googleCloudSpannerReceiver) Shutdown(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *googleCloudSpannerReceiver) initialize(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *googleCloudSpannerReceiver) initializeProjectReaders(ctx context.Context,
	parsedMetadata []*metadata.MetricsMetadata,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *googleCloudSpannerReceiver) initializeMetricsBuilder(parsedMetadata []*metadata.MetricsMetadata) error {
	_ = "STUB: not implemented"
	return nil
}

func newProjectReader(ctx context.Context, logger *zap.Logger, project Project, parsedMetadata []*metadata.MetricsMetadata,
	readerConfig statsreader.ReaderConfig,
) (*statsreader.ProjectReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
