// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package netstats // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/otelarrow/netstats"

import (
	"context"

	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

const (
	// ExporterKey is an attribute name that identifies an
	// exporter component that produces internal metrics, logs,
	// and traces.
	ExporterKey = "exporter"

	// ReceiverKey is an attribute name that identifies an
	// receiver component that produces internal metrics, logs,
	// and traces.
	ReceiverKey = "receiver"

	// SentBytes is used to track bytes sent by exporters and receivers.
	SentBytes = "sent"

	// SentWireBytes is used to track bytes sent on the wire
	// (includes compression) by exporters and receivers.
	SentWireBytes = "sent_wire"

	// RecvBytes is used to track bytes received by exporters and receivers.
	RecvBytes = "recv"

	// RecvWireBytes is used to track bytes received on the wire
	// (includes compression) by exporters and receivers.
	RecvWireBytes = "recv_wire"

	// CompSize is used for compressed size histogram metrics.
	CompSize = "compressed_size"

	scopeName = "github.com/open-telemetry/opentelemetry-collector-contrib/internal/otelarrow/netstats"
)

// NetworkReporter is a helper to add network-level observability to
// an exporter or receiver.
type NetworkReporter struct {
	isExporter    bool
	staticAttr    attribute.KeyValue
	sentBytes     metric.Int64Counter
	sentWireBytes metric.Int64Counter
	recvBytes     metric.Int64Counter
	recvWireBytes metric.Int64Counter
	compSizeHisto metric.Int64Histogram
}

var _ Interface = &NetworkReporter{}

// SizesStruct is used to pass uncompressed on-wire message lengths to
// the CountSend() and CountReceive() methods.
type SizesStruct struct {
	// Method refers to the gRPC method name
	Method string
	// Length is the uncompressed size
	Length int64
	// WireLength is compressed size
	WireLength int64
}

// Interface describes a *NetworkReporter or a Noop.
type Interface interface {
	// CountSend reports outbound bytes.
	CountSend(ctx context.Context, ss SizesStruct)

	// CountSend reports inbound bytes.
	CountReceive(ctx context.Context, ss SizesStruct)
}

// Noop is a no-op implementation of Interface.
type Noop struct{}

var _ Interface = Noop{}

func (Noop) CountSend(context.Context, SizesStruct)    { _ = "STUB: not implemented"; return }
func (Noop) CountReceive(context.Context, SizesStruct) { _ = "STUB: not implemented"; return }

const (
	bytesUnit           = "bytes"
	sentDescription     = "Number of bytes sent by the component."
	sentWireDescription = "Number of bytes sent on the wire by the component."
	recvDescription     = "Number of bytes received by the component."
	recvWireDescription = "Number of bytes received on the wire by the component."
	compSizeDescription = "Size of compressed payload"
)

// makeSentMetrics builds the sent and sent-wire metric instruments
// for an exporter or receiver using the corresponding `prefix`.
// major` indicates the major direction of the pipeline,
// which is true when sending for exporters, receiving for receivers.
func makeSentMetrics(prefix string, meter metric.Meter, major bool) (sent, sentWire metric.Int64Counter, _ error) {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter), *new(metric.Int64Counter), nil
}

// makeRecvMetrics builds the received and received-wire metric
// instruments for an exporter or receiver using the corresponding
// `prefix`.  `major` indicates the major direction of the pipeline,
// which is true when sending for exporters, receiving for receivers.
func makeRecvMetrics(prefix string, meter metric.Meter, major bool) (recv, recvWire metric.Int64Counter, _ error) {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter), *new(metric.Int64Counter), nil
}

// NewExporterNetworkReporter creates a new NetworkReporter configured for an exporter.
func NewExporterNetworkReporter(settings exporter.Settings) (*NetworkReporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Normally, an exporter counts sent bytes, and skips received
// bytes.  LevelDetailed will reveal exporter-received bytes.

// NewReceiverNetworkReporter creates a new NetworkReporter configured for an exporter.
func NewReceiverNetworkReporter(settings receiver.Settings) (*NetworkReporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Normally, a receiver counts received bytes, and skips sent
// bytes.  LevelDetailed will reveal receiver-sent bytes.

// CountSend is used to report a message sent by the component.  For
// exporters, SizesStruct indicates the size of a request.  For
// receivers, SizesStruct indicates the size of a response.
func (rep *NetworkReporter) CountSend(ctx context.Context, ss SizesStruct) {
	_ = "STUB: not implemented"
	// Indicates basic level telemetry, not counting bytes.
	return
}

// CountReceive is used to report a message received by the component.  For
// exporters, SizesStruct indicates the size of a response.  For
// receivers, SizesStruct indicates the size of a request.
func (rep *NetworkReporter) CountReceive(ctx context.Context, ss SizesStruct) {
	_ = "STUB: not implemented"
	// Indicates basic level telemetry, not counting bytes.
	return
}
