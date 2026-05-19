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
	// Max duration of data backfilling(if enabled).
	// Different backends can support or not support such option.
	// Since, the initial intent was to work mainly with Prometheus backend,
	// this constant was set to 1 hour - max allowed interval by Prometheus.
	backfillIntervalDuration = time.Hour
	topLockStatsMetricName   = "top minute lock stats"
	topQueryStatsMetricName  = "top minute query stats"
	maxLengthTruncateText    = 1024
)

type intervalStatsReader struct {
	currentStatsReader
	timestampsGenerator               *timestampsGenerator
	lastPullTimestamp                 time.Time
	hideTopnLockstatsRowrangestartkey bool
	truncateText                      bool
}

func newIntervalStatsReader(
	logger *zap.Logger,
	database *datasource.Database,
	metricsMetadata *metadata.MetricsMetadata,
	config ReaderConfig,
) *intervalStatsReader {
	_ = "STUB: not implemented"
	return nil
}

func (reader *intervalStatsReader) Read(ctx context.Context) ([]*metadata.MetricsDataPoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Generating pull timestamps

// Pulling metrics for each generated pull timestamp

// Latest timestamp for backfilling must be read from actual data(not stale)

func (reader *intervalStatsReader) newPullStatement(pullTimestamp time.Time) statsStatement {
	_ = "STUB: not implemented"
	return *new(statsStatement)
}

func (reader *intervalStatsReader) isBackfillExecution() bool {
	_ = "STUB: not implemented"
	return false
}
