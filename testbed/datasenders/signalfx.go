// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datasenders // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/datasenders"

import (
	"go.opentelemetry.io/collector/consumer"

	"github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"
)

// SFxMetricsDataSender implements MetricDataSender for SignalFx metrics protocol.
type SFxMetricsDataSender struct {
	testbed.DataSenderBase
	consumer.Metrics
}

// Ensure SFxMetricsDataSender implements MetricDataSenderOld.
var _ testbed.MetricDataSender = (*SFxMetricsDataSender)(nil)

// NewSFxMetricDataSender creates a new SignalFx metric protocol sender that will send
// to the specified port after Start is called.
func NewSFxMetricDataSender(port int) *SFxMetricsDataSender { _ = "STUB: not implemented"; return nil }

// Start the sender.
func (sf *SFxMetricsDataSender) Start() error { _ = "STUB: not implemented"; return nil }

// GenConfigYAMLStr returns receiver config for the agent.
func (sf *SFxMetricsDataSender) GenConfigYAMLStr() string {
	_ = "STUB: not implemented"
	// Note that this generates a receiver config for agent.
	return ""
}

// ProtocolName returns protocol name as it is specified in Collector config.
func (*SFxMetricsDataSender) ProtocolName() string { _ = "STUB: not implemented"; return "" }
