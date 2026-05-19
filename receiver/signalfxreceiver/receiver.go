// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package signalfxreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/signalfxreceiver"

import (
	"context"
	"net/http"
	"sync"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/signalfx"
)

const (
	defaultServerTimeout = 20 * time.Second

	responseOK                       = "OK"
	responseInvalidMethod            = "Only \"POST\" method is supported"
	responseEventsInvalidContentType = "\"Content-Type\" must be \"application/x-protobuf\""

	responseInvalidContentType      = "\"Content-Type\" must be either \"application/x-protobuf\" or \"application/x-protobuf;format=otlp\""
	responseInvalidEncoding         = "\"Content-Encoding\" must be \"gzip\" or empty"
	responseErrGzipReader           = "Error on gzip body"
	responseErrReadBody             = "Failed to read message body"
	responseErrUnmarshalBody        = "Failed to unmarshal message body"
	responseErrNextConsumer         = "Internal Server Error"
	responseErrLogsNotConfigured    = "Log pipeline has not been configured to handle events"
	responseErrMetricsNotConfigured = "Metric pipeline has not been configured to handle datapoints"

	// Centralizing some HTTP and related string constants.
	protobufContentType       = "application/x-protobuf"
	otlpProtobufContentType   = "application/x-protobuf;format=otlp"
	gzipEncoding              = "gzip"
	httpContentTypeHeader     = "Content-Type"
	httpContentEncodingHeader = "Content-Encoding"
)

var (
	okRespBody                   = initJSONResponse(responseOK)
	invalidMethodRespBody        = initJSONResponse(responseInvalidMethod)
	invalidContentRespBody       = initJSONResponse(responseInvalidContentType)
	invalidEventsContentRespBody = initJSONResponse(responseEventsInvalidContentType)
	invalidEncodingRespBody      = initJSONResponse(responseInvalidEncoding)
	errGzipReaderRespBody        = initJSONResponse(responseErrGzipReader)
	errReadBodyRespBody          = initJSONResponse(responseErrReadBody)
	errUnmarshalBodyRespBody     = initJSONResponse(responseErrUnmarshalBody)
	errNextConsumerRespBody      = initJSONResponse(responseErrNextConsumer)
	errLogsNotConfigured         = initJSONResponse(responseErrLogsNotConfigured)
	errMetricsNotConfigured      = initJSONResponse(responseErrMetricsNotConfigured)

	translator = &signalfx.ToTranslator{}
)

// sfxReceiver implements the receiver.Metrics for SignalFx metric protocol.
type sfxReceiver struct {
	settings        receiver.Settings
	config          *Config
	metricsConsumer consumer.Metrics
	logsConsumer    consumer.Logs
	server          *http.Server
	shutdownWG      sync.WaitGroup
	obsrecv         *receiverhelper.ObsReport
}

var _ receiver.Metrics = (*sfxReceiver)(nil)

// New creates the SignalFx receiver with the given configuration.
func newReceiver(
	settings receiver.Settings,
	config Config,
) (*sfxReceiver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *sfxReceiver) RegisterMetricsConsumer(mc consumer.Metrics) {
	_ = "STUB: not implemented"
	return
}

func (r *sfxReceiver) RegisterLogsConsumer(lc consumer.Logs) { _ = "STUB: not implemented"; return }

// Start tells the receiver to start its processing.
// By convention the consumer of the received data is set when the receiver
// instance is created.
func (r *sfxReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// set up the listener

// TODO: Evaluate what properties should be configurable, for now
//		set some hard-coded values.

// Shutdown tells the receiver that should stop reception,
// giving it a chance to perform any necessary clean-up.
func (r *sfxReceiver) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

func (r *sfxReceiver) readBody(ctx context.Context, resp http.ResponseWriter, req *http.Request) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (r *sfxReceiver) writeResponse(ctx context.Context, resp http.ResponseWriter, err error) {
	_ = "STUB: not implemented"
	return
}

func (r *sfxReceiver) handleDatapointReq(resp http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (r *sfxReceiver) handleEventReq(resp http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (r *sfxReceiver) failRequest(
	ctx context.Context,
	resp http.ResponseWriter,
	httpStatusCode int,
	jsonResponse []byte,
	err error,
) {
	_ = "STUB: not implemented"
	return
}

// The response needs to be written as a JSON string.

// Use the same pattern as strings.Builder String().

// It handles nil error

func initJSONResponse(s string) []byte { _ = "STUB: not implemented"; return nil }

// This is to be used in initialization so panic here is fine.
