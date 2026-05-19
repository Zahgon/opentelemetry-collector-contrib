// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datareceivers // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/datareceivers"

import (
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"
)

// StefDataReceiver implements SignalFx format receiver.
type StefDataReceiver struct {
	testbed.DataReceiverBase
	receiver receiver.Metrics
	Logger   *zap.Logger
}

// Ensure StefDataReceiver implements MetricDataSender.
var _ testbed.DataReceiver = (*StefDataReceiver)(nil)

// NewStefDataReceiver creates a new StefDataReceiver that will listen on the
// specified port after Start is called.
func NewStefDataReceiver(port int) *StefDataReceiver { _ = "STUB: not implemented"; return nil }

// Start the receiver.
func (sr *StefDataReceiver) Start(_ consumer.Traces, mc consumer.Metrics, _ consumer.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

// Stop the receiver.
func (sr *StefDataReceiver) Stop() error { _ = "STUB: not implemented"; return nil }

// GenConfigYAMLStr returns exporter config for the agent.
func (sr *StefDataReceiver) GenConfigYAMLStr() string {
	_ = "STUB: not implemented"
	// Note that this generates an exporter config for agent.
	return ""
}

// ProtocolName returns protocol name as it is specified in Collector config.
func (*StefDataReceiver) ProtocolName() string { _ = "STUB: not implemented"; return "" }
