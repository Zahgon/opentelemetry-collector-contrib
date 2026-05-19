// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datasenders // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/datasenders"

import (
	"go.opentelemetry.io/collector/consumer"

	"github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"
)

type prometheusDataSender struct {
	testbed.DataSenderBase
	consumer.Metrics
	namespace string
}

// NewPrometheusDataSender creates a new Prometheus sender that will expose data
// on the specified port after Start is called.
func NewPrometheusDataSender(host string, port int) testbed.MetricDataSender {
	_ = "STUB: not implemented"
	return *new(testbed.MetricDataSender)
}

func (pds *prometheusDataSender) Start() error { _ = "STUB: not implemented"; return nil }

func (pds *prometheusDataSender) GenConfigYAMLStr() string { _ = "STUB: not implemented"; return "" }

func (*prometheusDataSender) ProtocolName() string { _ = "STUB: not implemented"; return "" }
