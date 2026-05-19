// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package testbed // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"

import (
	"net"

	"go.opentelemetry.io/collector/config/configcompression"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/exporter/otlpexporter"
	"go.opentelemetry.io/collector/exporter/otlphttpexporter"
)

// DataSender defines the interface that allows sending data. This is an interface
// that must be implemented by all protocols that want to be used in ProviderSender.
// Note the terminology: DataSender is something that sends data to Collector
// and the corresponding entity that receives the data in the Collector is a receiver.
type DataSender interface {
	// Start sender and connect to the configured endpoint. Must be called before
	// sending data.
	Start() error

	// Flush sends any accumulated data.
	Flush()

	// GetEndpoint returns the address to which this sender will send data.
	GetEndpoint() net.Addr

	// GenConfigYAMLStr generates a config string to place in receiver part of collector config
	// so that it can receive data from this sender.
	GenConfigYAMLStr() string

	// ProtocolName returns exporter name to use in collector config pipeline.
	ProtocolName() string
}

// TraceDataSender defines the interface that allows sending trace data. It adds ability
// to send a batch of Spans to the DataSender interface.
type TraceDataSender interface {
	DataSender
	consumer.Traces
}

// MetricDataSender defines the interface that allows sending metric data. It adds ability
// to send a batch of Metrics to the DataSender interface.
type MetricDataSender interface {
	DataSender
	consumer.Metrics
}

// LogDataSender defines the interface that allows sending log data. It adds ability
// to send a batch of Logs to the DataSender interface.
type LogDataSender interface {
	DataSender
	consumer.Logs
}

type DataSenderBase struct {
	Port int
	Host string
}

func (dsb *DataSenderBase) GetEndpoint() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (*DataSenderBase) Flush() {
	_ = "STUB: not implemented"
	// Exporter interface does not support Flush, so nothing to do.
	return
}

type otlpHTTPDataSender struct {
	DataSenderBase
	compression configcompression.Type
}

func (ods *otlpHTTPDataSender) fillConfig(cfg *otlphttpexporter.Config) *otlphttpexporter.Config {
	_ = "STUB: not implemented"
	return nil
}

// Disable retries, we should push data and if error just log it.

// Disable sending queue, we should push data from the caller goroutine.

func (ods *otlpHTTPDataSender) GenConfigYAMLStr() string {
	_ = "STUB: not implemented"
	// Note that this generates a receiver config for agent.
	return ""
}

func (*otlpHTTPDataSender) ProtocolName() string {
	_ = "STUB: not implemented"

	// otlpHTTPTraceDataSender implements TraceDataSender for OTLP/HTTP trace exporter.
	return ""
}

type otlpHTTPTraceDataSender struct {
	otlpHTTPDataSender
	consumer.Traces
}

// NewOTLPHTTPTraceDataSender creates a new TraceDataSender for OTLP/HTTP traces exporter.
func NewOTLPHTTPTraceDataSender(host string, port int, compression configcompression.Type) TraceDataSender {
	_ = "STUB: not implemented"
	return *new(TraceDataSender)
}

func (ote *otlpHTTPTraceDataSender) Start() error { _ = "STUB: not implemented"; return nil }

// otlpHTTPMetricsDataSender implements MetricDataSender for OTLP/HTTP metrics exporter.
type otlpHTTPMetricsDataSender struct {
	otlpHTTPDataSender
	consumer.Metrics
}

// NewOTLPHTTPMetricDataSender creates a new OTLP/HTTP metrics exporter sender that will send
// to the specified port after Start is called.
func NewOTLPHTTPMetricDataSender(host string, port int) MetricDataSender {
	_ = "STUB: not implemented"
	return *new(MetricDataSender)
}

func (ome *otlpHTTPMetricsDataSender) Start() error { _ = "STUB: not implemented"; return nil }

// otlpHTTPLogsDataSender implements LogsDataSender for OTLP/HTTP logs exporter.
type otlpHTTPLogsDataSender struct {
	otlpHTTPDataSender
	consumer.Logs
}

// NewOTLPHTTPLogsDataSender creates a new OTLP/HTTP logs exporter sender that will send
// to the specified port after Start is called.
func NewOTLPHTTPLogsDataSender(host string, port int) LogDataSender {
	_ = "STUB: not implemented"
	return *new(LogDataSender)
}

func (olds *otlpHTTPLogsDataSender) Start() error { _ = "STUB: not implemented"; return nil }

type otlpDataSender struct {
	DataSenderBase
}

func (ods *otlpDataSender) fillConfig(cfg *otlpexporter.Config) *otlpexporter.Config {
	_ = "STUB: not implemented"
	return nil
}

// Disable retries, we should push data and if error just log it.

// Disable sending queue, we should push data from the caller goroutine.

func (ods *otlpDataSender) GenConfigYAMLStr() string {
	_ = "STUB: not implemented"
	// Note that this generates a receiver config for agent.
	return ""
}

func (*otlpDataSender) ProtocolName() string {
	_ = "STUB: not implemented"

	// otlpTraceDataSender implements TraceDataSender for OTLP traces exporter.
	return ""
}

type otlpTraceDataSender struct {
	otlpDataSender
	consumer.Traces
}

// NewOTLPTraceDataSender creates a new TraceDataSender for OTLP traces exporter.
func NewOTLPTraceDataSender(host string, port int) TraceDataSender {
	_ = "STUB: not implemented"
	return *new(TraceDataSender)
}

func (ote *otlpTraceDataSender) Start() error { _ = "STUB: not implemented"; return nil }

// otlpMetricsDataSender implements MetricDataSender for OTLP metrics exporter.
type otlpMetricsDataSender struct {
	otlpDataSender
	consumer.Metrics
}

// NewOTLPMetricDataSender creates a new OTLP metric exporter sender that will send
// to the specified port after Start is called.
func NewOTLPMetricDataSender(host string, port int) MetricDataSender {
	_ = "STUB: not implemented"
	return *new(MetricDataSender)
}

func (ome *otlpMetricsDataSender) Start() error { _ = "STUB: not implemented"; return nil }

// otlpLogsDataSender implements LogsDataSender for OTLP logs exporter.
type otlpLogsDataSender struct {
	otlpDataSender
	consumer.Logs
}

// NewOTLPLogsDataSender creates a new OTLP logs exporter sender that will send
// to the specified port after Start is called.
func NewOTLPLogsDataSender(host string, port int) LogDataSender {
	_ = "STUB: not implemented"
	return *new(LogDataSender)
}

func (olds *otlpLogsDataSender) Start() error { _ = "STUB: not implemented"; return nil }
