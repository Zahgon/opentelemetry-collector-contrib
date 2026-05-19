// Copyright The OpenTelemetry Authors
// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package thriftudp // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jaegerreceiver/internal/udpserver/thriftudp"

import (
	"bytes"
	"context"
	"errors"
	"net"

	"github.com/apache/thrift/lib/go/thrift"
)

// MaxLength of UDP packet
const (
	MaxLength = 65000
)

var errConnAlreadyClosed = errors.New("connection already closed")

// TUDPTransport does UDP as a thrift.TTransport
type TUDPTransport struct {
	conn     *net.UDPConn
	addr     net.Addr
	writeBuf bytes.Buffer
	closed   uint32 // atomic flag
}

var _ thrift.TTransport = (*TUDPTransport)(nil)

// NewTUDPClientTransport creates a net.UDPConn-backed TTransport for Thrift clients
// All writes are buffered and flushed in one UDP packet. If locHostPort is not "", it
// will be used as the local address for the connection
// Example:
//
//	trans, err := thriftudp.NewTUDPClientTransport("192.168.1.1:9090", "")
func NewTUDPClientTransport(destHostPort, locHostPort string) (*TUDPTransport, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createClient(destAddr, locAddr *net.UDPAddr) (*TUDPTransport, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewTUDPServerTransport creates a net.UDPConn-backed TTransport for Thrift servers
// It will listen for incoming udp packets on the specified host/port
// Example:
//
//	trans, err := thriftudp.NewTUDPClientTransport("localhost:9001")
func NewTUDPServerTransport(hostPort string) (*TUDPTransport, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Open does nothing as connection is opened on creation
// Required to maintain thrift.TTransport interface
func (*TUDPTransport) Open() error {
	_ = "STUB: not implemented"

	// Conn retrieves the underlying net.UDPConn
	return nil
}

func (p *TUDPTransport) Conn() *net.UDPConn {
	_ = "STUB: not implemented"

	// IsOpen returns true if the connection is open
	return nil
}

func (p *TUDPTransport) IsOpen() bool { _ = "STUB: not implemented"; return false }

// Close closes the connection
func (p *TUDPTransport) Close() error { _ = "STUB: not implemented"; return nil }

// Addr returns the address that the transport is listening on or writing to
func (p *TUDPTransport) Addr() net.Addr {
	_ = "STUB: not implemented"

	// Read reads one UDP packet and puts it in the specified buf
	return *new(net.Addr)
}

func (p *TUDPTransport) Read(buf []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// RemainingBytes returns the max number of bytes (same as Thrift's StreamTransport) as we
// do not know how many bytes we have left.
func (*TUDPTransport) RemainingBytes() uint64 { _ = "STUB: not implemented"; return 0 }

// Write writes specified buf to the write buffer
func (p *TUDPTransport) Write(buf []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Flush flushes the write buffer as one udp packet
func (p *TUDPTransport) Flush(_ context.Context) error { _ = "STUB: not implemented"; return nil }

// always reset the buffer, even in case of an error

// SetSocketBufferSize sets udp buffer size
func (p *TUDPTransport) SetSocketBufferSize(bufferSize int) error {
	_ = "STUB: not implemented"
	return nil
}
