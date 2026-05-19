// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package correctnesstests // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/correctnesstests"

import (
	"testing"

	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"
)

type ProcessorNameAndConfigBody struct {
	Name string
	Body string
}

// CreateConfigYaml creates a yaml config for an otel collector given a testbed sender, testbed receiver, any
// processors, and a pipeline type. A collector created from the resulting yaml string should be able to talk
// the specified sender and receiver.
func CreateConfigYaml(
	tb testing.TB,
	sender testbed.DataSender,
	receiver testbed.DataReceiver,
	connector testbed.DataConnector,
	processors []ProcessorNameAndConfigBody,
) string {
	_ = "STUB: not implemented"
	return ""
}

// Avoid picking a telemetry port that is already planned for the sender.
// This prevents port collisions between the collector's Prometheus telemetry
// and the collector's receivers (sender endpoint), which can happen if CreateConfigYaml
// is called before the receivers are started (bound).

// Prepare extra processor config section and comma-separated list of extra processor
// names to use in corresponding "processors" settings.

// PipelineDef holds the information necessary to run a single testbed configuration.
type PipelineDef struct {
	Receiver      string
	Exporter      string
	Connector     string
	TestName      string
	DataSender    testbed.DataSender
	DataReceiver  testbed.DataReceiver
	DataConnector testbed.DataConnector
	ResourceSpec  testbed.ResourceSpec
}

// LoadPictOutputPipelineDefs generates a slice of PipelineDefs from the passed-in generated PICT file. The
// result should be a set of PipelineDefs that covers all possible pipeline configurations.
func LoadPictOutputPipelineDefs(fileName string) ([]PipelineDef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newDbgLogger() *zap.Logger { _ = "STUB: not implemented"; return nil }

// ConstructTraceSender creates a testbed trace sender from the passed-in trace sender identifier.
func ConstructTraceSender(t *testing.T, receiver string) testbed.DataSender {
	_ = "STUB: not implemented"
	return *new(testbed.DataSender)
}

// ConstructMetricsSender creates a testbed metrics sender from the passed-in metrics sender identifier.
func ConstructMetricsSender(t *testing.T, receiver string) testbed.MetricDataSender {
	_ = "STUB: not implemented"
	return *new(testbed.MetricDataSender)
}

// ConstructReceiver creates a testbed receiver from the passed-in recevier identifier.
func ConstructReceiver(t *testing.T, exporter string) testbed.DataReceiver {
	_ = "STUB: not implemented"
	return *new(testbed.DataReceiver)
}

func ConstructConnector(t *testing.T, connector, receiverType string) testbed.DataConnector {
	_ = "STUB: not implemented"
	return *new(testbed.DataConnector)
}
