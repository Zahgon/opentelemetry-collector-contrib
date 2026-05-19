// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datareceivers // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/datareceivers"

import (
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"
)

// SyslogDataReceiver implements Syslog format receiver.
type SyslogDataReceiver struct {
	testbed.DataReceiverBase
	receiver receiver.Logs
	protocol string
}

// Ensure SyslogDataReceiver implements LogDataReceiver.
var _ testbed.DataReceiver = (*SyslogDataReceiver)(nil)

// NewSyslogDataReceiver creates a new SyslogDataReceiver that will listen on the
// specified port after Start is called.
func NewSyslogDataReceiver(protocol string, port int) *SyslogDataReceiver {
	_ = "STUB: not implemented"
	return nil
}

// Start the receiver.
func (cr *SyslogDataReceiver) Start(_ consumer.Traces, _ consumer.Metrics, lc consumer.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

// Stop the receiver.
func (cr *SyslogDataReceiver) Stop() error { _ = "STUB: not implemented"; return nil }

// GenConfigYAMLStr returns receiver config for the agent.
func (cr *SyslogDataReceiver) GenConfigYAMLStr() string {
	_ = "STUB: not implemented"
	// Note that this generates an receiver config for agent.
	return ""
}

// ProtocolName returns protocol name as it is specified in Collector config.
func (*SyslogDataReceiver) ProtocolName() string { _ = "STUB: not implemented"; return "" }
