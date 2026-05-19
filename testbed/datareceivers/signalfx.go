// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datareceivers // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/datareceivers"

import (
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"
)

// SFxMetricsDataReceiver implements SignalFx format receiver.
type SFxMetricsDataReceiver struct {
	testbed.DataReceiverBase
	receiver receiver.Metrics
}

// Ensure SFxMetricsDataReceiver implements MetricDataSender.
var _ testbed.DataReceiver = (*SFxMetricsDataReceiver)(nil)

// NewSFxMetricsDataReceiver creates a new SFxMetricsDataReceiver that will listen on the
// specified port after Start is called.
func NewSFxMetricsDataReceiver(port int) *SFxMetricsDataReceiver {
	_ = "STUB: not implemented"
	return nil
}

// Start the receiver.
func (sr *SFxMetricsDataReceiver) Start(_ consumer.Traces, mc consumer.Metrics, _ consumer.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

// Stop the receiver.
func (sr *SFxMetricsDataReceiver) Stop() error { _ = "STUB: not implemented"; return nil }

// GenConfigYAMLStr returns exporter config for the agent.
func (sr *SFxMetricsDataReceiver) GenConfigYAMLStr() string {
	_ = "STUB: not implemented"
	// Note that this generates an exporter config for agent.
	return ""
}

// ProtocolName returns protocol name as it is specified in Collector config.
func (*SFxMetricsDataReceiver) ProtocolName() string { _ = "STUB: not implemented"; return "" }
