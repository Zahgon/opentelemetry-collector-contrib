// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datareceivers // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/datareceivers"

import (
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"
)

// jaegerDataReceiver implements Jaeger format receiver.
type jaegerDataReceiver struct {
	testbed.DataReceiverBase
	receiver receiver.Traces
}

// NewJaegerDataReceiver creates a new Jaeger DataReceiver that will listen on the specified port after Start
// is called.
func NewJaegerDataReceiver(port int) testbed.DataReceiver {
	_ = "STUB: not implemented"
	return *new(testbed.DataReceiver)
}

func (jr *jaegerDataReceiver) Start(tc consumer.Traces, _ consumer.Metrics, _ consumer.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (jr *jaegerDataReceiver) Stop() error { _ = "STUB: not implemented"; return nil }

func (jr *jaegerDataReceiver) GenConfigYAMLStr() string {
	_ = "STUB: not implemented"
	// Note that this generates an exporter config for agent.
	// The Jaeger exporter is no longer supported, therefore
	// we export data using OTLP instead
	return ""
}

func (*jaegerDataReceiver) ProtocolName() string { _ = "STUB: not implemented"; return "" }
