// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package statsreader // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/statsreader"

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/datasource"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/metadata"
)

const (
	dataStalenessSeconds = 15
	dataStalenessPeriod  = dataStalenessSeconds * time.Second

	// Data reads for backfilling are always performed from stale replica nodes and not from the master.
	// For current/fresh data reads main node is used. But in case collector started for example at mm:30+ it is safe
	// to read data from stale replica nodes(at this time replica node will contain required data), since we are
	// requesting data for mm:00. Also stale reads are faster than reads from main node.
	dataStalenessSafeThresholdSeconds = 2 * dataStalenessSeconds
)

type currentStatsReader struct {
	logger                 *zap.Logger
	database               *datasource.Database
	metricsMetadata        *metadata.MetricsMetadata
	topMetricsQueryMaxRows int
	statement              func(args statementArgs) statsStatement
}

func newCurrentStatsReader(
	logger *zap.Logger,
	database *datasource.Database,
	metricsMetadata *metadata.MetricsMetadata,
	config ReaderConfig,
) *currentStatsReader {
	_ = "STUB: not implemented"
	return nil
}

func (reader *currentStatsReader) Name() string { _ = "STUB: not implemented"; return "" }

func (reader *currentStatsReader) Read(ctx context.Context) ([]*metadata.MetricsDataPoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (reader *currentStatsReader) newPullStatement() statsStatement {
	_ = "STUB: not implemented"
	return *new(statsStatement)
}

func (reader *currentStatsReader) pull(ctx context.Context, stmt statsStatement) ([]*metadata.MetricsDataPoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isSafeToUseStaleRead(readTimestamp time.Time) bool { _ = "STUB: not implemented"; return false }
