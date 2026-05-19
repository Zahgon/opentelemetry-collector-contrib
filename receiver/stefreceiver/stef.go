// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package stefreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/stefreceiver"
import (
	"context"
	"sync/atomic"

	stefgrpc "github.com/splunk/stef/go/grpc"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

type stefReceiver struct {
	cfg        *Config
	serverGRPC *grpc.Server

	nextMetricsConsumer consumer.Metrics
	settings            receiver.Settings

	stopping atomic.Bool
	eg       errgroup.Group
}

// Start runs the STEF gRPC receiver.
func (r *stefReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Shutdown is a method to turn off receiving.
func (r *stefReceiver) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

// Give graceful stop a second to finish.

func (r *stefReceiver) onStream(grpcReader stefgrpc.GrpcReader, stream stefgrpc.STEFStream) error {
	_ = "STUB: not implemented"
	return nil
}

// Create a responder for this stream and run it in a separate goroutine.

// Read, decode, convert the incoming data and push it to the next consumer.

// The receiver is shutting down. Close the connection.

// We had problem sending responses. Can't continue using this connection since
// responding is essential for operation.

// Mark the start of the converted batch.

// Read and convert records. We use ConvertTillEndOfFrame to make sure we are not
// blocked in the middle of a batch indefinitely, with lingering data in memory,
// neither pushed to pipeline, nor acked.

// A regular disconnection case. The client closed the connection.

// Push converted data to the next consumer.

// The next consumer is temporarily unable to process the data.
// Close the stream and indicate to client to try again later.

// This is a permanent error. Let the client know and continue receiving data.

// Successfully received and consumed. Acknowledge it.
