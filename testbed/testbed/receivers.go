// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package testbed // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"

import (
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

// DataReceiver allows to receive traces or metrics. This is an interface that must
// be implemented by all protocols that want to be used in MockBackend.
// Note the terminology: DataReceiver is something that can listen and receive data
// from Collector and the corresponding entity in the Collector that sends this data is
// an exporter.
type DataReceiver interface {
	Start(tc consumer.Traces, mc consumer.Metrics, lc consumer.Logs) error
	Stop() error

	// GenConfigYAMLStr generates a config string to place in exporter part of collector config
	// so that it can send data to this receiver.
	GenConfigYAMLStr() string

	// ProtocolName returns exporterType name to use in collector config pipeline.
	ProtocolName() string
}

// DataReceiverBase implement basic functions needed by all receivers.
type DataReceiverBase struct {
	// Port on which to listen.
	Port int
	// prevent unkeyed literal initialization
	_ struct{}
}

// TODO: Move these constants.
const (
	DefaultHost     = "127.0.0.1"
	DefaultOTLPPort = 4317
)

// BaseOTLPDataReceiver implements the OTLP format receiver.
type BaseOTLPDataReceiver struct {
	DataReceiverBase
	// One of the "otlp_grpc" for OTLP over gRPC or "otlp_http" for OTLP over HTTP.
	exporterType    string
	traceReceiver   receiver.Traces
	metricsReceiver receiver.Metrics
	logReceiver     receiver.Logs
	compression     string
	retry           string
	sendingQueue    string
	timeout         string
	batcher         string
}

// InsertDefault is a helper function to insert a default value for a configoptional.Optional type.
func InsertDefault[T any](opt *configoptional.Optional[T]) error {
	_ = "STUB: not implemented"
	return nil
}

func (bor *BaseOTLPDataReceiver) Start(tc consumer.Traces, mc consumer.Metrics, lc consumer.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

// we reuse the receiver across signals. Starting the log receiver starts the metrics and traces receiver.

func (bor *BaseOTLPDataReceiver) WithCompression(compression string) *BaseOTLPDataReceiver {
	_ = "STUB: not implemented"
	return nil
}

func (bor *BaseOTLPDataReceiver) WithRetry(retry string) *BaseOTLPDataReceiver {
	_ = "STUB: not implemented"
	return nil
}

func (bor *BaseOTLPDataReceiver) WithQueue(sendingQueue string) *BaseOTLPDataReceiver {
	_ = "STUB: not implemented"
	return nil
}

func (bor *BaseOTLPDataReceiver) WithTimeout(timeout string) *BaseOTLPDataReceiver {
	_ = "STUB: not implemented"
	return nil
}

func (bor *BaseOTLPDataReceiver) WithBatcher(batcher string) *BaseOTLPDataReceiver {
	_ = "STUB: not implemented"
	return nil
}

func (bor *BaseOTLPDataReceiver) Stop() error {
	_ = "STUB: not implemented"
	// we reuse the receiver across signals. Shutting down the log receiver shuts down the metrics and traces receiver.
	return nil
}

func (bor *BaseOTLPDataReceiver) ProtocolName() string { _ = "STUB: not implemented"; return "" }

func (bor *BaseOTLPDataReceiver) GenConfigYAMLStr() string { _ = "STUB: not implemented"; return "" }

// Note that this generates an exporter config for agent.

var _ DataReceiver = (*BaseOTLPDataReceiver)(nil)

// NewOTLPDataReceiver creates a new OTLP DataReceiver that will listen on the specified port after Start
// is called.
func NewOTLPDataReceiver(port int) *BaseOTLPDataReceiver { _ = "STUB: not implemented"; return nil }

// NewOTLPHTTPDataReceiver creates a new OTLP/HTTP DataReceiver that will listen on the specified port after Start
// is called.
func NewOTLPHTTPDataReceiver(port int) *BaseOTLPDataReceiver { _ = "STUB: not implemented"; return nil }
