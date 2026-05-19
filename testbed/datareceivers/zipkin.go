// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datareceivers // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/datareceivers"

import (
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"
)

// zipkinDataReceiver implements Zipkin format receiver.
type zipkinDataReceiver struct {
	testbed.DataReceiverBase
	receiver receiver.Traces
}

// NewZipkinDataReceiver creates a new Zipkin DataReceiver that will listen on the specified port after Start
// is called.
func NewZipkinDataReceiver(port int) testbed.DataReceiver {
	_ = "STUB: not implemented"
	return *new(testbed.DataReceiver)
}

func (zr *zipkinDataReceiver) Start(tc consumer.Traces, _ consumer.Metrics, _ consumer.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (zr *zipkinDataReceiver) Stop() error { _ = "STUB: not implemented"; return nil }

func (zr *zipkinDataReceiver) GenConfigYAMLStr() string {
	_ = "STUB: not implemented"
	// Note that this generates an exporter config for agent.
	return ""
}

func (*zipkinDataReceiver) ProtocolName() string { _ = "STUB: not implemented"; return "" }
