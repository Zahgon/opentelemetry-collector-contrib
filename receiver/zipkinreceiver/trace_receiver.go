// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package zipkinreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/zipkinreceiver"

import (
	"context"
	"io"
	"net/http"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
)

const (
	receiverTransportV1Thrift = "http_v1_thrift"
	receiverTransportV1JSON   = "http_v1_json"
	receiverTransportV2JSON   = "http_v2_json"
	receiverTransportV2PROTO  = "http_v2_proto"
)

var (
	errNextConsumerRespBody = []byte(`"Internal Server Error"`)
	errBadRequestRespBody   = []byte(`"Bad Request"`)
)

// zipkinReceiver type is used to handle spans received in the Zipkin format.
type zipkinReceiver struct {
	nextConsumer consumer.Traces

	shutdownWG sync.WaitGroup
	server     *http.Server
	config     *Config

	v1ThriftUnmarshaler      ptrace.Unmarshaler
	v1JSONUnmarshaler        ptrace.Unmarshaler
	jsonUnmarshaler          ptrace.Unmarshaler
	protobufUnmarshaler      ptrace.Unmarshaler
	protobufDebugUnmarshaler ptrace.Unmarshaler

	settings  receiver.Settings
	obsrecvrs map[string]*receiverhelper.ObsReport
}

var _ http.Handler = (*zipkinReceiver)(nil)

// newReceiver creates a new zipkinReceiver reference.
func newReceiver(config *Config, nextConsumer consumer.Traces, settings receiver.Settings) (*zipkinReceiver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Start spins up the receiver's HTTP server and makes the receiver start its processing.
func (zr *zipkinReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// v1ToTraceSpans parses Zipkin v1 JSON traces and converts them to OpenCensus Proto spans.
func (zr *zipkinReceiver) v1ToTraceSpans(blob []byte, hdr http.Header) (reqs ptrace.Traces, err error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

// v2ToTraceSpans parses Zipkin v2 JSON or Protobuf traces and converts them to OpenCensus Proto spans.
func (zr *zipkinReceiver) v2ToTraceSpans(blob []byte, hdr http.Header) (reqs ptrace.Traces, err error) {
	_ = "STUB: not implemented"
	// This flag's reference is from:
	//      https://github.com/openzipkin/zipkin-go/blob/3793c981d4f621c0e3eb1457acffa2c1cc591384/proto/v2/zipkin.proto#L154
	return *new(ptrace.Traces), nil
}

// By default, we'll assume using JSON

// Zipkin can send protobuf via http

// TODO: (@odeke-em) record the unique types of Content-Type uploads

// Shutdown tells the receiver that should stop reception,
// giving it a chance to perform any necessary clean-up and shutting down
// its HTTP server.
func (zr *zipkinReceiver) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

// processBodyIfNecessary checks the "Content-Encoding" HTTP header and if
// a compression such as "gzip", "deflate", "zlib", is found, the body will
// be uncompressed accordingly or return the body untouched if otherwise.
// Clients such as Zipkin-Java do this behavior e.g.
//
//	send "Content-Encoding":"gzip" of the JSON content.
func processBodyIfNecessary(req *http.Request) io.Reader {
	_ = "STUB: not implemented"
	return *new(io.Reader)
}

func gunzippedBodyIfPossible(r io.Reader) io.Reader {
	_ = "STUB: not implemented"
	return *new(io.Reader)
}

// Just return the old body as was

func zlibUncompressedbody(r io.Reader) io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

// Just return the old body as was

const (
	zipkinV1TagValue = "zipkinV1"
	zipkinV2TagValue = "zipkinV2"
)

// The zipkinReceiver receives spans from endpoint /api/v2 as JSON,
// unmarshalls them and sends them along to the nextConsumer.
func (zr *zipkinReceiver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"

	// Now deserialize and process the spans.
	return
}

// Send back the response "Accepted" as
// required at https://zipkin.io/zipkin-api/#/default/post_spans

// Transient error, due to some internal condition.

func transportType(r *http.Request, asZipkinv1 bool) string { _ = "STUB: not implemented"; return "" }
