// Copyright The OpenTelemetry Authors
// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package udpserver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jaegerreceiver/internal/udpserver"

import (
	"bytes"
	"context"

	"github.com/apache/thrift/lib/go/thrift"
)

// TBufferedReadTransport is a thrift.TTransport that reads from a buffer
type TBufferedReadTransport struct {
	readBuf *bytes.Buffer
}

var _ thrift.TTransport = (*TBufferedReadTransport)(nil)

// NewTBufferedReadTransport creates a buffer backed TTransport
func NewTBufferedReadTransport(readBuf *bytes.Buffer) (*TBufferedReadTransport, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IsOpen does nothing as transport is not maintaining the connection
// Required to maintain thrift.TTransport interface
func (*TBufferedReadTransport) IsOpen() bool {
	_ = "STUB: not implemented"

	// Open does nothing as transport is not maintaining the connection
	// Required to maintain thrift.TTransport interface
	return false
}

func (*TBufferedReadTransport) Open() error {
	_ = "STUB: not implemented"

	// Close does nothing as transport is not maintaining the connection
	// Required to maintain thrift.TTransport interface
	return nil
}

func (*TBufferedReadTransport) Close() error {
	_ = "STUB: not implemented"

	// Read reads bytes from the local buffer and puts them in the specified buf
	return nil
}

func (p *TBufferedReadTransport) Read(buf []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// RemainingBytes returns the number of bytes left to be read from the readBuf
func (p *TBufferedReadTransport) RemainingBytes() uint64 { _ = "STUB: not implemented"; return 0 }

// Write writes bytes into the read buffer
// Required to maintain thrift.TTransport interface
func (p *TBufferedReadTransport) Write(buf []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Flush does nothing as udp server does not write responses back
// Required to maintain thrift.TTransport interface
func (*TBufferedReadTransport) Flush(_ context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
