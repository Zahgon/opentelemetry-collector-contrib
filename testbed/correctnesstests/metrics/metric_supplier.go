// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/correctnesstests/metrics"

import (
	"go.opentelemetry.io/collector/pdata/pmetric"
)

type metricSupplier struct {
	pdms    []pmetric.Metrics
	currIdx int
}

func newMetricSupplier(pdms []pmetric.Metrics) *metricSupplier {
	_ = "STUB: not implemented"
	return nil
}

func (p *metricSupplier) nextMetrics() (pdm pmetric.Metrics, done bool) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), false
}
