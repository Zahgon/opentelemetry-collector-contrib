// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/correctnesstests/metrics"

import (
	"context"
	"testing"

	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/metricstestutil"
	"github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"
)

// testHarness listens for datapoints from the receiver to which it is attached
// and when it receives one, it compares it to the datapoint that was previously
// sent out. It then sends the next datapoint, if there is one.
type testHarness struct {
	t                  *testing.T
	metricSupplier     *metricSupplier
	metricIndex        *metricsReceivedIndex
	sender             testbed.MetricDataSender
	diffConsumer       diffConsumer
	outOfMetrics       bool
	allMetricsReceived chan struct{}
}

type diffConsumer interface {
	accept(string, []*metricstestutil.MetricDiff)
}

func newTestHarness(
	t *testing.T,
	s *metricSupplier,
	mi *metricsReceivedIndex,
	ds testbed.MetricDataSender,
	diffConsumer diffConsumer,
) *testHarness {
	_ = "STUB: not implemented"
	return nil
}

func (*testHarness) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (h *testHarness) ConsumeMetrics(_ context.Context, pdm pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *testHarness) compare(pdm pmetric.Metrics) { _ = "STUB: not implemented"; return }

func (h *testHarness) sendNextMetric() { _ = "STUB: not implemented"; return }
