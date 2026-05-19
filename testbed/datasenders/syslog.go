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

type SyslogWriter struct {
	testbed.DataSenderBase
	conn    net.Conn
	buf     []string
	bufSize int
	network string
}

var _ testbed.LogDataSender = (*SyslogWriter)(nil)

func NewSyslogWriter(network, host string, port, batchSize int) *SyslogWriter {
	_ = "STUB: not implemented"
	return nil
}

func (f *SyslogWriter) GetEndpoint() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (*SyslogWriter) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (f *SyslogWriter) Start() (err error) { _ = "STUB: not implemented"; return nil }

// udp not ack, can't use net.Dial to check udp server is ready, use sleep 1 second to wait udp server start

func (f *SyslogWriter) ConsumeLogs(_ context.Context, logs plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *SyslogWriter) GenConfigYAMLStr() string { _ = "STUB: not implemented"; return "" }

func (f *SyslogWriter) Send(lr plog.LogRecord) error { _ = "STUB: not implemented"; return nil }

func (f *SyslogWriter) SendCheck() error { _ = "STUB: not implemented"; return nil }

func (*SyslogWriter) Flush() { _ = "STUB: not implemented"; return }

func (*SyslogWriter) ProtocolName() string { _ = "STUB: not implemented"; return "" }
