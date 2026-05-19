// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package aggregator // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/signaltometricsconnector/internal/aggregator"
import (
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/connector/signaltometricsconnector/internal/model"
)

// valueCountDP is a wrapper DP to aggregate all datapoints that record
// value and count.
type valueCountDP struct {
	expHistogramDP      *exponentialHistogramDP
	explicitHistogramDP *explicitHistogramDP
}

func newValueCountDP[K any](
	md model.MetricDef[K],
	attrs pcommon.Map,
) *valueCountDP {
	_ = "STUB: not implemented"
	return nil
}

func (dp *valueCountDP) Aggregate(value float64, count int64) { _ = "STUB: not implemented"; return }

func (dp *valueCountDP) Copy(
	timestamp time.Time,
	destExpHist pmetric.ExponentialHistogram,
	destExplicitHist pmetric.Histogram,
) {
	_ = "STUB: not implemented"
	return
}
