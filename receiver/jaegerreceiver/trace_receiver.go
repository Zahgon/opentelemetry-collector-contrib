// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package jaegerreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jaegerreceiver"

import (
	"context"
	"net/http"
	"sync"

	apacheThrift "github.com/apache/thrift/lib/go/thrift"
	"github.com/jaegertracing/jaeger-idl/proto-gen/api_v2"
	"github.com/jaegertracing/jaeger-idl/thrift-gen/agent"
	"github.com/jaegertracing/jaeger-idl/thrift-gen/jaeger"
	"github.com/jaegertracing/jaeger-idl/thrift-gen/zipkincore"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
	"google.golang.org/grpc"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jaegerreceiver/internal/udpserver"
)

// Receiver type is used to receive spans that were originally intended to be sent to Jaeger.
// This receiver is basically a Jaeger collector.
type jReceiver struct {
	nextConsumer consumer.Traces
	id           component.ID

	config Protocols

	grpc            *grpc.Server
	collectorServer *http.Server

	agentProcessors []*udpserver.ThriftProcessor

	goroutines sync.WaitGroup

	settings receiver.Settings

	grpcObsrecv *receiverhelper.ObsReport
	httpObsrecv *receiverhelper.ObsReport
}

const (
	agentTransportBinary   = "udp_thrift_binary"
	agentTransportCompact  = "udp_thrift_compact"
	collectorHTTPTransport = "collector_http"
	grpcTransport          = "grpc"

	thriftFormat   = "thrift"
	protobufFormat = "protobuf"
)

var acceptedThriftFormats = map[string]struct{}{
	"application/x-thrift":                 {},
	"application/vnd.apache.thrift.binary": {},
}

// newJaegerReceiver creates a TracesReceiver that receives traffic as a Jaeger collector, and
// also as a Jaeger agent.
func newJaegerReceiver(
	id component.ID,
	config Protocols,
	nextConsumer consumer.Traces,
	set receiver.Settings,
) (*jReceiver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (jr *jReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (jr *jReceiver) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func consumeTraces(ctx context.Context, batch *jaeger.Batch, consumer consumer.Traces) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

var (
	_ agent.Agent                   = (*agentHandler)(nil)
	_ api_v2.CollectorServiceServer = (*jReceiver)(nil)
)

type agentHandler struct {
	nextConsumer consumer.Traces
	obsrecv      *receiverhelper.ObsReport
}

// EmitZipkinBatch is unsupported agent's
func (*agentHandler) EmitZipkinBatch(context.Context, []*zipkincore.Span) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// EmitBatch implements thrift-gen/agent/Agent and it forwards
// Jaeger spans received by the Jaeger agent processor.
func (h *agentHandler) EmitBatch(ctx context.Context, batch *jaeger.Batch) error {
	_ = "STUB: not implemented"
	return nil
}

func (jr *jReceiver) PostSpans(ctx context.Context, r *api_v2.PostSpansRequest) (*api_v2.PostSpansResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (jr *jReceiver) startAgent() error { _ = "STUB: not implemented"; return nil }

func (jr *jReceiver) buildProcessor(address string, cfg ServerConfigUDP, factory apacheThrift.TProtocolFactory, a agent.Agent) (*udpserver.ThriftProcessor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*jReceiver) decodeThriftHTTPBody(r *http.Request) (*jaeger.Batch, *httpError) {
	_ = "STUB: not implemented"
	return nil, nil
}

// HandleThriftHTTPBatch implements Jaeger HTTP Thrift handler.
func (jr *jReceiver) HandleThriftHTTPBatch(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (jr *jReceiver) startCollector(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}
