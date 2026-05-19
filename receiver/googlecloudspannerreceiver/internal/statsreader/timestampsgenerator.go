// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package statsreader // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/statsreader"

import "time"

type timestampsGenerator struct {
	backfillEnabled bool
	difference      time.Duration
}

// This slice will always contain at least one value - now shifted to the start of minute(upper bound).
// In case lastPullTimestamp is greater than now argument slice will contain only one value - now shifted to the start of minute(upper bound).
func (g *timestampsGenerator) pullTimestamps(lastPullTimestamp, now time.Time) []time.Time {
	_ = "STUB: not implemented"
	return nil
}

// lastPullTimestamp is already set to start of minute

// This slice will always contain at least one value(upper bound).
// Difference between each two points is 1 minute.
func pullTimestampsWithDifference(lowerBound, upperBound time.Time, difference time.Duration) []time.Time {
	_ = "STUB: not implemented"
	return nil
}

// To ensure that we did not miss upper bound and timestamps slice will contain at least one value

func shiftToStartOfMinute(now time.Time) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (g *timestampsGenerator) isBackfillExecution(lastPullTimestamp time.Time) bool {
	_ = "STUB: not implemented"
	return false
}
