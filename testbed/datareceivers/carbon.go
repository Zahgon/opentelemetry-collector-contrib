// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datareceivers // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/datareceivers"

import (
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"
)

// CarbonDataReceiver implements Carbon format receiver.
type CarbonDataReceiver struct {
	testbed.DataReceiverBase
	receiver receiver.Metrics
}

// Ensure CarbonDataReceiver implements DataReceiver
var _ testbed.DataReceiver = (*CarbonDataReceiver)(nil)

// NewCarbonDataReceiver creates a new CarbonDataReceiver that will listen on the
// specified port after Start is called.
func NewCarbonDataReceiver(port int) *CarbonDataReceiver { _ = "STUB: not implemented"; return nil }

// Start the receiver.
func (cr *CarbonDataReceiver) Start(_ consumer.Traces, mc consumer.Metrics, _ consumer.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

// Stop the receiver.
func (cr *CarbonDataReceiver) Stop() error { _ = "STUB: not implemented"; return nil }

// GenConfigYAMLStr returns exporter config for the agent.
func (cr *CarbonDataReceiver) GenConfigYAMLStr() string {
	_ = "STUB: not implemented"
	// Note that this generates an exporter config for agent.
	return ""
}

// ProtocolName returns protocol name as it is specified in Collector config.
func (*CarbonDataReceiver) ProtocolName() string { _ = "STUB: not implemented"; return "" }
