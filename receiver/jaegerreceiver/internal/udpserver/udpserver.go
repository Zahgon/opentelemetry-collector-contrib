// Copyright The OpenTelemetry Authors
// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package udpserver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jaegerreceiver/internal/udpserver"

import (
	"bytes"
	"io"
	"sync"
	"sync/atomic"
)

// TBufferedServer is an alias for UDPServer, for backwards compatibility.
type TBufferedServer = UDPServer

// NewTBufferedServer is an alias for NewUDPServer, for backwards compatibility.
var NewTBufferedServer = NewUDPServer

// UDPConn is a an abstraction of *net.UDPConn, for easier mocking.
type UDPConn interface {
	io.Reader
	io.Closer
}

// UDPServer reads packets from a UDP connection into bytes.Buffer and places
// each buffer into a bounded channel to be consumed by the receiver.
// After consuming the buffer, the receiver SHOULD call DataRecd() to signal
// that the buffer is no longer in use and to return it to the pool.
type UDPServer struct {
	queueSize     atomic.Int64
	dataChan      chan *bytes.Buffer
	maxPacketSize int
	maxQueueSize  int
	serving       uint32
	transport     UDPConn
	readBufPool   sync.Pool
}

// state values for TBufferedServer.serving
//
// init -> serving -> stopped
// init -> stopped (might happen in unit tests)
const (
	stateStopped = iota
	stateServing
	stateInit
)

// NewUDPServer creates a UDPServer
func NewUDPServer(
	transport UDPConn,
	maxQueueSize int,
	maxPacketSize int,
) *UDPServer {
	_ = "STUB: not implemented"
	return nil
}

// packetReader is a helper for reading a single packet no larger than maxPacketSize
// from the underlying reader. Without it the ReadFrom() method of bytes.Buffer would
// read multiple packets and won't even stop at maxPacketSize.
type packetReader struct {
	maxPacketSize int
	reader        io.LimitedReader
	attempt       int
}

func (r *packetReader) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *packetReader) readPacket(buf *bytes.Buffer) (int, error) {
	_ = "STUB: not implemented"
	// reset the readers since we're reusing them to avoid allocations
	return 0, nil
}

// prepare the buffer for expected packet size

// use Buffer's ReadFrom() as otherwise it's hard to get it into the right state

// Serve initiates the readers and starts serving traffic
func (s *UDPServer) Serve() { _ = "STUB: not implemented"; return }

// Stop already called

func (s *UDPServer) updateQueueSize(delta int64) { _ = "STUB: not implemented"; return }

// IsServing indicates whether the server is currently serving traffic
func (s *UDPServer) IsServing() bool { _ = "STUB: not implemented"; return false }

// Stop stops the serving of traffic and waits until the queue is
// emptied by the readers
func (s *UDPServer) Stop() { _ = "STUB: not implemented"; return }

// DataChan returns the data chan of the buffered server
func (s *UDPServer) DataChan() chan *bytes.Buffer {
	_ = "STUB: not implemented"

	// DataRecd is called by the consumers every time they read a data item from DataChan
	return nil
}

func (s *UDPServer) DataRecd(buf *bytes.Buffer) { _ = "STUB: not implemented"; return }
