// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datareceivers // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/datareceivers"

import (
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"
)

// SplunkHECDataReceiver implements Splunk HEC format receiver.
type SplunkHECDataReceiver struct {
	testbed.DataReceiverBase
	receiver receiver.Logs
}

// Ensure SplunkHECDataReceiver implements LogDataSender.
var _ testbed.DataReceiver = (*SplunkHECDataReceiver)(nil)

// NewSplunkHECDataReceiver creates a new SplunkHECDataReceiver that will listen on the
// specified port after Start is called.
func NewSplunkHECDataReceiver(port int) *SplunkHECDataReceiver {
	_ = "STUB: not implemented"
	return nil
}

// Start the receiver.
func (sr *SplunkHECDataReceiver) Start(_ consumer.Traces, _ consumer.Metrics, lc consumer.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

// Stop the receiver.
func (sr *SplunkHECDataReceiver) Stop() error { _ = "STUB: not implemented"; return nil }

// GenConfigYAMLStr returns exporter config for the agent.
func (sr *SplunkHECDataReceiver) GenConfigYAMLStr() string {
	_ = "STUB: not implemented"
	// Note that this generates an exporter config for agent.
	return ""
}

// ProtocolName returns protocol name as it is specified in Collector config.
func (*SplunkHECDataReceiver) ProtocolName() string { _ = "STUB: not implemented"; return "" }
