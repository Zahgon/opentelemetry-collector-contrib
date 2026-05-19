// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package dataconnectors // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/dataconnectors"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"
)

type RoutingDataConnector struct {
	testbed.DataConnectorBase
}

var _ testbed.DataConnector = (*RoutingDataConnector)(nil)

func NewRoutingDataConnector(receiverDataType string) *RoutingDataConnector {
	_ = "STUB: not implemented"
	return nil
}

func (*RoutingDataConnector) GenConfigYAMLStr() string {
	_ = "STUB: not implemented"
	// Note that this generates an exporter config for agent.
	return ""
}

// ProtocolName returns protocol name as it is specified in Collector config.
func (*RoutingDataConnector) ProtocolName() string { _ = "STUB: not implemented"; return "" }

func (rc *RoutingDataConnector) GetReceiverType() string { _ = "STUB: not implemented"; return "" }
