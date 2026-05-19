// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datareceivers // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/datareceivers"

import (
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"
)

// datadogDataReceiver implements Datadog v3/v4/v5 format receiver.
type datadogDataReceiver struct {
	testbed.DataReceiverBase
	receiver receiver.Traces
}

// NewDataDogDataReceiver creates a new DD DataReceiver that will listen on the specified port after Start
// is called.
func NewDataDogDataReceiver() testbed.DataReceiver {
	_ = "STUB: not implemented"
	return *new(testbed.DataReceiver)
}

func (dd *datadogDataReceiver) Start(tc consumer.Traces, _ consumer.Metrics, _ consumer.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (dd *datadogDataReceiver) Stop() error { _ = "STUB: not implemented"; return nil }

func (*datadogDataReceiver) GenConfigYAMLStr() string {
	_ = "STUB: not implemented"
	// Note that this generates an exporter config for agent.
	return ""
}

func (*datadogDataReceiver) ProtocolName() string { _ = "STUB: not implemented"; return "" }
