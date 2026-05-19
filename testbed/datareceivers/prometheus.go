// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datareceivers // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/datareceivers"

import (
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"
)

type prometheusDataReceiver struct {
	testbed.DataReceiverBase
	receiver receiver.Metrics
}

func NewPrometheusDataReceiver(port int) testbed.DataReceiver {
	_ = "STUB: not implemented"
	return *new(testbed.DataReceiver)
}

func (dr *prometheusDataReceiver) Start(_ consumer.Traces, mc consumer.Metrics, _ consumer.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (dr *prometheusDataReceiver) Stop() error { _ = "STUB: not implemented"; return nil }

func (dr *prometheusDataReceiver) GenConfigYAMLStr() string { _ = "STUB: not implemented"; return "" }

func (*prometheusDataReceiver) ProtocolName() string { _ = "STUB: not implemented"; return "" }
