// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package deltatocumulativeprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor"

import (
	"context"
	"sync"
	"time"

	"github.com/puzpuzpuz/xsync/v4"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/processor"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/exp/metrics/identity"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor/internal/data"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor/internal/maps"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor/internal/telemetry"
)

var _ processor.Metrics = (*deltaToCumulativeProcessor)(nil)

type deltaToCumulativeProcessor struct {
	next consumer.Metrics
	cfg  Config

	last state
	aggr data.Aggregator

	ctx    context.Context
	cancel context.CancelFunc

	stale *xsync.Map[identity.Stream, time.Time]
	tel   telemetry.Metrics
}

func newProcessor(cfg *Config, tel telemetry.Metrics, next consumer.Metrics) *deltaToCumulativeProcessor {
	_ = "STUB: not implemented"
	return nil
}

type vals struct {
	nums *mutex[pmetric.NumberDataPoint]
	hist *mutex[pmetric.HistogramDataPoint]
	expo *mutex[pmetric.ExponentialHistogramDataPoint]
}

func (p *deltaToCumulativeProcessor) ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

// aggregate the datapoints.
// using filter here, as the pmetric.*DataPoint are reference types so
// we can modify them using their "value".

// count the processed datatype.
// uses whatever value of attrs has at return-time

// state is full, reject stream

// stream is ok and active, update stale tracker

// cached zero was stored, alloc new one

// state is full, reject stream

// stream is ok and active, update stale tracker

// cached zero was stored, alloc new one

// state is full, reject stream

// stream is ok and active, update stale tracker

// cached zero was stored, alloc new one

// all remaining datapoints of this metric are now cumulative

// if no datapoints remain, drop empty metric

// no need to continue pipeline if we dropped all metrics

func (p *deltaToCumulativeProcessor) Start(_ context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil

	// delete stale streams once per minute
}

func (p *deltaToCumulativeProcessor) Shutdown(_ context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (*deltaToCumulativeProcessor) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

// state keeps a cumulative value, aggregated over time, per stream
type state struct {
	ctx  maps.Context
	nums *maps.Parallel[identity.Stream, *mutex[pmetric.NumberDataPoint]]
	hist *maps.Parallel[identity.Stream, *mutex[pmetric.HistogramDataPoint]]
	expo *maps.Parallel[identity.Stream, *mutex[pmetric.ExponentialHistogramDataPoint]]
}

func (s state) Size() int { _ = "STUB: not implemented"; return 0 }

type mutex[T any] struct {
	mtx sync.Mutex
	v   T
}

func (mtx *mutex[T]) use(do func(T)) { _ = "STUB: not implemented"; return }

func guard[T any](v T) *mutex[T] { _ = "STUB: not implemented"; return nil }
