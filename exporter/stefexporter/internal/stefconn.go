// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/stefexporter/internal"

import (
	"context"
	"sync"

	stefgrpc "github.com/splunk/stef/go/grpc"
	"github.com/splunk/stef/go/otel/otelstef"
	"github.com/splunk/stef/go/pkg"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// StefConnCreator implements ConnCreator interface for STEF/gRPC connections.
type StefConnCreator struct {
	logger      *zap.Logger
	grpcConn    *grpc.ClientConn
	compression pkg.Compression
}

// StefConn implements Conn interface for STEF/gRPC connections.
type StefConn struct {
	client *stefgrpc.Client
	writer *otelstef.MetricsWriter

	// cancel func for the context that was used to create the gRPC stream.
	// Calling it will cancel stream operations.
	cancel context.CancelFunc

	// pendingAcks is a map of channels that are registered and awaiting
	// acknowledgments of the particular DataID.
	pendingAcks map[DataID]chan<- AsyncResult
	// mux protects the pendingAcks.
	mux sync.RWMutex

	flushReqCh chan struct{}
	flushResCh chan error
}

func NewStefConnCreator(logger *zap.Logger, grpcConn *grpc.ClientConn, compression pkg.Compression) *StefConnCreator {
	_ = "STUB: not implemented"
	return nil
}

// Create a new connection. May be called concurrently.
// The attempt to create the connection should be cancelled if ctx is done.
func (s *StefConnCreator) Create(ctx context.Context) (Conn, error) {
	_ = "STUB: not implemented"
	// Prepare to open a STEF/gRPC stream to the server.
	return *new(Conn), nil
}

// Let server know about our schema.

// Start a goroutine that waits for success, failure or cancellation of
// the connection attempt.

// Wait for either connection attempt to be done or for the caller
// of Create() to give up.

// The caller of Create() cancelled while we are waiting
// for connection to be established. We have to cancel the
// connection attempt (and the whole connection if it raced us and
// managed to connect - we will reconnect later again in that case).

// Connection attempt already finished, nothing to do. This can happen
// if <-ctx.Done() above selects sooner than <-connectionAttemptDone.
// That is ok, we are done.

// Connection attempt finished (successfully or no). No need to wait for the
// previous case, calling connCancel() is not needed anymore now. It will be
// called later, when disconnecting.
// From this moment we are essentially detaching from the Context
// that passed to Create() since we wanted to honor it only
// for the duration of the connection attempt, but not for the duration
// of the entire existence of the connection.

// Create STEF record writer over gRPC.

// We need to call the cancel func when this connection is over so that we don't
// leak the Context we just created. This will be done in disconnect().

// Run flusher in a separate goroutine.

// Writer returns the metrics writer that exists over this connection.
func (s *StefConn) Writer() *otelstef.MetricsWriter {
	_ = "STUB: not implemented"

	// OnAck registers to notify via ackCh when the acknowledgment with
	// the given ackID is received over this connection. When acknowledgment
	// with the specified ackID is received, the AsyncResult with ackID
	// will send to ackCh.
	return nil
}

func (s *StefConn) OnAck(ackID uint64, ackCh chan<- AsyncResult) { _ = "STUB: not implemented"; return }

// onGrpcAck is called by stefgrpc.Client when an acknowledgment is received.
func (s *StefConn) onGrpcAck(ackID uint64) error {
	_ = "STUB: not implemented"

	// Notify all pending acks that have ackID smaller or equal to the received ackID.
	return nil
}

// Close the connection.
func (s *StefConn) Close(ctx context.Context) error {
	_ = "STUB: not implemented"

	// Stop flusher goroutine
	return nil
}

func (s *StefConn) flusher() { _ = "STUB: not implemented"; return }

// Flush any pending data over the connection.
func (s *StefConn) Flush(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Request a flush.
	return nil
}

// Wait until flushed.

type loggerWrapper struct {
	logger *zap.Logger
}

func (w *loggerWrapper) Debugf(_ context.Context, format string, v ...any) {
	_ = "STUB: not implemented"
	return
}

func (w *loggerWrapper) Errorf(_ context.Context, format string, v ...any) {
	_ = "STUB: not implemented"
	return
}
