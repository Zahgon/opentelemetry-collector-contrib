// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package fluentforwardreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/fluentforwardreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/fluentforwardreceiver/internal/metadata"
)

// collector acts as an aggregator of LogRecords so that we don't have to
// generate as many plog.Logs instances...we can pre-batch the LogRecord
// instances from several Forward events into one to hopefully reduce
// allocations and GC overhead.
type collector struct {
	nextConsumer     consumer.Logs
	eventCh          <-chan event
	logger           *zap.Logger
	obsrecv          *receiverhelper.ObsReport
	telemetryBuilder *metadata.TelemetryBuilder
}

func newCollector(eventCh <-chan event, next consumer.Logs, logger *zap.Logger, obsrecv *receiverhelper.ObsReport, telemetryBuilder *metadata.TelemetryBuilder) *collector {
	_ = "STUB: not implemented"
	return nil
}

func (c *collector) Start(ctx context.Context) { _ = "STUB: not implemented"; return }

func (c *collector) processEvents(ctx context.Context) { _ = "STUB: not implemented"; return }

// Pull out anything waiting on the eventCh to get better
// efficiency on LogResource allocations.

func (c *collector) fillBufferUntilChanEmpty(dest plog.LogRecordSlice) {
	_ = "STUB: not implemented"
	return
}
