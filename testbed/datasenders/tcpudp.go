// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datasenders // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/datasenders"

import (
	"context"
	"net"

	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"

	"github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"
)

type TCPUDPWriter struct {
	testbed.DataSenderBase
	conn    net.Conn
	buf     []string
	bufSize int
	network string
}

var _ testbed.LogDataSender = (*TCPUDPWriter)(nil)

func NewTCPUDPWriter(network, host string, port, batchSize int) *TCPUDPWriter {
	_ = "STUB: not implemented"
	return nil
}

func (f *TCPUDPWriter) GetEndpoint() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (*TCPUDPWriter) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (f *TCPUDPWriter) Start() (err error) { _ = "STUB: not implemented"; return nil }

// udp not ack, can't use net.Dial to check udp server is ready, use sleep 1 second to wait udp server start

func (f *TCPUDPWriter) ConsumeLogs(_ context.Context, logs plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *TCPUDPWriter) GenConfigYAMLStr() string { _ = "STUB: not implemented"; return "" }

func (f *TCPUDPWriter) Send(lr plog.LogRecord) error { _ = "STUB: not implemented"; return nil }

func (f *TCPUDPWriter) SendCheck() error { _ = "STUB: not implemented"; return nil }

func (*TCPUDPWriter) Flush() { _ = "STUB: not implemented"; return }

func (f *TCPUDPWriter) ProtocolName() string { _ = "STUB: not implemented"; return "" }
