// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datasenders // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/datasenders"

import (
	"go.opentelemetry.io/collector/consumer"

	"github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"
)

type otelarrowDataSender struct {
	testbed.DataSenderBase
	consumer.Metrics
}

// NewOtelarrowDataSender creates a new sender that will send
// to the specified port after Start is called.
func NewOtelarrowDataSender(host string, port int) testbed.MetricDataSender {
	_ = "STUB: not implemented"
	return *new(testbed.MetricDataSender)
}

func (ds *otelarrowDataSender) Start() error { _ = "STUB: not implemented"; return nil }

func (ds *otelarrowDataSender) GenConfigYAMLStr() string { _ = "STUB: not implemented"; return "" }

func (*otelarrowDataSender) ProtocolName() string { _ = "STUB: not implemented"; return "" }
