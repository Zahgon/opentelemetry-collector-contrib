// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datasenders // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/datasenders"

import (
	"go.opentelemetry.io/collector/consumer"

	"github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"
)

// datadogDataSender implements TraceDataSender for Zipkin http exporter.
type datadogDataSender struct {
	testbed.DataSenderBase
	consumer.Traces
}

// NewDatadogDataSender creates a new Zipkin exporter sender that will send
// to the specified port after Start is called.
func NewDatadogDataSender() testbed.TraceDataSender {
	_ = "STUB: not implemented"
	return *new(testbed.TraceDataSender)
}

func (dd *datadogDataSender) Start() error { _ = "STUB: not implemented"; return nil }

func (dd *datadogDataSender) GenConfigYAMLStr() string { _ = "STUB: not implemented"; return "" }

func (*datadogDataSender) ProtocolName() string { _ = "STUB: not implemented"; return "" }
