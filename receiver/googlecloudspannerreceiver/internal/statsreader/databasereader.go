// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package statsreader // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/statsreader"

import (
	"context"

	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/datasource"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/metadata"
)

type DatabaseReader struct {
	database *datasource.Database
	logger   *zap.Logger
	readers  []Reader
}

func NewDatabaseReader(ctx context.Context,
	parsedMetadata []*metadata.MetricsMetadata,
	databaseID *datasource.DatabaseID,
	serviceAccountPath string,
	readerConfig ReaderConfig,
	logger *zap.Logger,
) (*DatabaseReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func initializeReaders(logger *zap.Logger, parsedMetadata []*metadata.MetricsMetadata,
	database *datasource.Database, readerConfig ReaderConfig,
) []Reader {
	_ = "STUB: not implemented"
	return nil
}

func (databaseReader *DatabaseReader) Name() string { _ = "STUB: not implemented"; return "" }

func (databaseReader *DatabaseReader) Shutdown() { _ = "STUB: not implemented"; return }

func (databaseReader *DatabaseReader) Read(ctx context.Context) ([]*metadata.MetricsDataPoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
