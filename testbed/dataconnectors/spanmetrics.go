// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package dataconnectors // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/dataconnectors"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"
)

type SpanMetricDataConnector struct {
	testbed.DataConnectorBase
}

var _ testbed.DataConnector = (*SpanMetricDataConnector)(nil)

func NewSpanMetricDataConnector(receiverDataType string) *SpanMetricDataConnector {
	_ = "STUB: not implemented"
	return nil
}

func (*SpanMetricDataConnector) GenConfigYAMLStr() string {
	_ = "STUB: not implemented"
	// Note that this generates an exporter config for agent.
	return ""
}

// ProtocolName returns protocol name as it is specified in Collector config.
func (*SpanMetricDataConnector) ProtocolName() string { _ = "STUB: not implemented"; return "" }

func (smc *SpanMetricDataConnector) GetReceiverType() string { _ = "STUB: not implemented"; return "" }
