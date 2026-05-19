// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datasenders // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/datasenders"

import (
	"go.opentelemetry.io/collector/consumer"

	"github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"
)

// zipkinDataSender implements TraceDataSender for Zipkin http exporter.
type zipkinDataSender struct {
	testbed.DataSenderBase
	consumer.Traces
}

// NewZipkinDataSender creates a new Zipkin exporter sender that will send
// to the specified port after Start is called.
func NewZipkinDataSender(host string, port int) testbed.TraceDataSender {
	_ = "STUB: not implemented"
	return *new(testbed.TraceDataSender)
}

func (zs *zipkinDataSender) Start() error { _ = "STUB: not implemented"; return nil }

// Disable retries, we should push data and if error just log it.

// Disable sending queue, we should push data from the caller goroutine.

func (zs *zipkinDataSender) GenConfigYAMLStr() string { _ = "STUB: not implemented"; return "" }

func (*zipkinDataSender) ProtocolName() string { _ = "STUB: not implemented"; return "" }
