// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package statsreader // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/statsreader"

import (
	"time"

	"cloud.google.com/go/spanner"
)

const (
	topMetricsQueryLimitParameterName = "topMetricsQueryMaxRows"
	topMetricsQueryLimitCondition     = " LIMIT @" + topMetricsQueryLimitParameterName

	pullTimestampParameterName = "pullTimestamp"
)

type statementArgs struct {
	query                  string
	topMetricsQueryMaxRows int
	pullTimestamp          time.Time
	stalenessRead          bool
}

type statsStatement struct {
	statement     spanner.Statement
	stalenessRead bool
}

func currentStatsStatement(args statementArgs) statsStatement {
	_ = "STUB: not implemented"
	return *new(statsStatement)
}

func intervalStatsStatement(args statementArgs) statsStatement {
	_ = "STUB: not implemented"
	return *new(statsStatement)
}
