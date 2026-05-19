// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package stefexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/stefexporter"

import (
	"context"
	"time"

	stefpkg "github.com/splunk/stef/go/pkg"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"google.golang.org/grpc"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/stefexporter/internal"
)

// stefExporter implements sending metrics over STEF/gRPC stream.
//
// The exporter uses a single stream and accepts concurrent exportMetrics calls,
// sequencing the metric data as needed over a single stream.
//
// The exporter will block exportMetrics call until an acknowledgement is
// received from destination.
//
// The exporter relies on a preceding Retry helper to retry sending data that is
// not acknowledged or otherwise fails to be sent. The exporter will not retry
// sending the data itself.
type stefExporter struct {
	set         component.TelemetrySettings
	cfg         Config
	compression stefpkg.Compression
	started     bool

	grpcConn *grpc.ClientConn

	connMan    *internal.ConnManager
	sync2Async *internal.Sync2Async
}

const (
	flushPeriod     = 100 * time.Millisecond
	reconnectPeriod = 10 * time.Minute
)

// TODO: make connection count configurable.
const connCount = 1

func newStefExporter(set component.TelemetrySettings, cfg *Config) *stefExporter {
	_ = "STUB: not implemented"
	return nil
}

// Disable built-in grpc compression. STEF has its own zstd compression support.

func (s *stefExporter) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	// Prepare gRPC connection.
	return nil
}

// Create a connection creator and manager to take care of the connections.

// Wrap async implementation of sendMetricsAsync into a sync-callable API.

// Begin connection attempt in a goroutine to avoid blocking Start().

// Acquire() triggers a connection attempt and blocks until it succeeds or fails.

// This is not a fatal error. Next sending attempt will try to
// connect again as needed.

// Connection is established. Return it, this is all we needed for now.

func (s *stefExporter) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *stefExporter) exportMetrics(ctx context.Context, data pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

// sendMetricsAsync is an async implementation of sending metric data.
// The result of sending will be reported via resultChan.
func (s *stefExporter) sendMetricsAsync(
	ctx context.Context,
	data any,
	resultChan internal.ResultChan,
) (internal.DataID, error) {
	_ = "STUB: not implemented"
	// Acquire a connection to send the data over.
	return *new(internal.DataID), nil
}

// It must be a StefConn with a Writer.

// Convert and write the data to the Writer.

// Error to write to STEF stream typically indicates either:
// 1) A problem with the connection. We need to reconnect.
// 2) Encoding failure, possibly due to encoder bug. In this case
//    we need to reconnect too, to make sure encoders start from
//    initial state, which is our best chance to succeed next time.
//
// We need to reconnect. Disconnect here and the next exportMetrics()
// call will connect again.

// TODO: check if err is because STEF encoding failed. If so we must not
// try to re-encode the same data. Return consumererror.NewPermanent(err)
// to the caller. This requires changes in STEF Go library.

// Return an error to retry sending these metrics again next time.

// According to STEF gRPC spec the destination ack IDs match written record number.
// When the data we have just written is received by destination it will send us
// back an ack ID that numerically matches the last written record number.

// Register to be notified via resultChan when the ack of the
// written record is received.

// We are done with the connection.
