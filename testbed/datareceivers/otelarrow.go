// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datareceivers // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/datareceivers"

import (
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"
)

// OtelarrowDataReceiver implements Otel Arrow format receiver.
type OtelarrowDataReceiver struct {
	testbed.DataReceiverBase
	receiver receiver.Metrics
}

// Ensure OtelarrowDataReceiver implements MetricDataSender.
var _ testbed.DataReceiver = (*OtelarrowDataReceiver)(nil)

// NewOtelarrowDataReceiver creates a new OtelarrowDataReceiver that will listen on the
// specified port after Start is called.
func NewOtelarrowDataReceiver(port int) *OtelarrowDataReceiver {
	_ = "STUB: not implemented"
	return nil
}

// Start the receiver.
func (dr *OtelarrowDataReceiver) Start(_ consumer.Traces, mc consumer.Metrics, _ consumer.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

// Stop the receiver.
func (dr *OtelarrowDataReceiver) Stop() error { _ = "STUB: not implemented"; return nil }

// GenConfigYAMLStr returns exporter config for the agent.
func (dr *OtelarrowDataReceiver) GenConfigYAMLStr() string {
	_ = "STUB: not implemented"
	// Note that this generates an exporter config for agent.
	return ""
}

// ProtocolName returns protocol name as it is specified in Collector config.
func (*OtelarrowDataReceiver) ProtocolName() string { _ = "STUB: not implemented"; return "" }
