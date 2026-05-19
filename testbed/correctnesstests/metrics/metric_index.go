// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/correctnesstests/metrics"

import (
	"go.opentelemetry.io/collector/pdata/pmetric"
)

type metricReceived struct {
	pdm      pmetric.Metrics
	received bool
}

type metricsReceivedIndex struct {
	m map[string]*metricReceived
}

func newMetricsReceivedIndex(pdms []pmetric.Metrics) *metricsReceivedIndex {
	_ = "STUB: not implemented"
	return nil
}

func (mi *metricsReceivedIndex) lookup(name string) (*metricReceived, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (mi *metricsReceivedIndex) allReceived() bool { _ = "STUB: not implemented"; return false }
