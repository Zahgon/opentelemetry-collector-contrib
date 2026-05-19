// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package splunkhecreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/splunkhecreceiver"

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/ackextension"
)

const (
	defaultServerTimeout = 20 * time.Second

	ackResponse                       = `{"acks": %s}`
	responseOK                        = `{"text": "Success", "code": 0}`
	responseOKWithAckID               = `{"text": "Success", "code": 0, "ackId": %d}`
	responseHecHealthy                = `{"text": "HEC is healthy", "code": 17}`
	responseInvalidMethodPostOnly     = `"Only \"POST\" method is supported"`
	responseInvalidEncoding           = `"\"Content-Encoding\" must be \"gzip\" or empty"`
	responseInvalidDataFormat         = `{"text":"Invalid data format","code":6}`
	responseErrEventRequired          = `{"text":"Event field is required","code":12}`
	responseErrEventBlank             = `{"text":"Event field cannot be blank","code":13}`
	responseErrGzipReader             = `"Error on gzip body"`
	responseErrUnmarshalBody          = `"Failed to unmarshal message body"`
	responseErrInternalServerError    = `"Internal Server Error"`
	responseErrUnsupportedMetricEvent = `"Unsupported metric event"`
	responseErrUnsupportedLogEvent    = `"Unsupported log event"`
	responseErrHandlingIndexedFields  = `{"text":"Error in handling indexed fields","code":15,"invalid-event-number":%d}`
	responseErrDataChannelMissing     = `{"text": "Data channel is missing","code":10}`
	responseErrInvalidDataChannel     = `{"text": "Invalid data channel", "code": 11}`
	responseNoData                    = `{"text":"No data","code":5}`
	// Centralizing some HTTP and related string constants.
	gzipEncoding              = "gzip"
	httpContentEncodingHeader = "Content-Encoding"
	httpContentTypeHeader     = "Content-Type"
	httpJSONTypeHeader        = "application/json"
)

var (
	errEmptyEndpoint    = errors.New("empty endpoint")
	errInvalidMethod    = errors.New("invalid http method")
	errInvalidEncoding  = errors.New("invalid encoding")
	errExtensionMissing = errors.New("ack extension not found")

	okRespBody                    = []byte(responseOK)
	eventRequiredRespBody         = []byte(responseErrEventRequired)
	eventBlankRespBody            = []byte(responseErrEventBlank)
	requiredDataChannelHeader     = []byte(responseErrDataChannelMissing)
	invalidEncodingRespBody       = []byte(responseInvalidEncoding)
	invalidFormatRespBody         = []byte(responseInvalidDataFormat)
	invalidMethodRespBodyPostOnly = []byte(responseInvalidMethodPostOnly)
	errGzipReaderRespBody         = []byte(responseErrGzipReader)
	errUnmarshalBodyRespBody      = []byte(responseErrUnmarshalBody)
	errInternalServerError        = []byte(responseErrInternalServerError)
	errUnsupportedMetricEvent     = []byte(responseErrUnsupportedMetricEvent)
	errUnsupportedLogEvent        = []byte(responseErrUnsupportedLogEvent)
	noDataRespBody                = []byte(responseNoData)
)

// splunkReceiver implements the receiver.Metrics for Splunk HEC metric protocol.
type splunkReceiver struct {
	settings        receiver.Settings
	config          *Config
	logsConsumer    consumer.Logs
	metricsConsumer consumer.Metrics
	server          *http.Server
	shutdownWG      sync.WaitGroup
	obsrecv         *receiverhelper.ObsReport
	gzipReaderPool  *sync.Pool
	ackExt          ackextension.AckExtension
}

var (
	_ receiver.Metrics = (*splunkReceiver)(nil)
	_ receiver.Logs    = (*splunkReceiver)(nil)
)

// newReceiver creates the Splunk HEC receiver with the given configuration.
func newReceiver(settings receiver.Settings, config Config) (*splunkReceiver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Start tells the receiver to start its processing.
// By convention the consumer of the received data is set when the receiver
// instance is created.
func (r *splunkReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	// server.Handler will be nil on initial call, otherwise noop.
	return nil
}

// set up the ack API handler if the ack extension is present

// set up the listener

// TODO: Evaluate what properties should be configurable, for now
//		set some hard-coded values.

// Shutdown tells the receiver that should stop reception,
// giving it a chance to perform any necessary clean-up.
func (r *splunkReceiver) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

func (r *splunkReceiver) processSuccessResponseWithAck(resp http.ResponseWriter, channelID string) error {
	_ = "STUB: not implemented"
	return nil
}

func (*splunkReceiver) processSuccessResponse(resp http.ResponseWriter, bodyContent []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *splunkReceiver) handleAck(resp http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

// shouldn't run into this case since we only enable this handler IF ackExt exists. But we have this check just in case

func (r *splunkReceiver) handleRawReq(resp http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (*splunkReceiver) extractChannel(req *http.Request) (string, bool) {
	_ = "STUB: not implemented"
	// check header
	return "", false
}

// check query param

func (*splunkReceiver) validateChannelHeader(channelID string) error {
	_ = "STUB: not implemented"
	return nil
}

// channel id must be a valid uuid
// https://docs.splunk.com/Documentation/Splunk/9.2.1/Data/AboutHECIDXAck#:~:text=close%20the%20file.-,About%20channels%20and%20sending%20data,-Sending%20events%20to

func (r *splunkReceiver) handleReq(resp http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Array of events

// If event is JSON array, there should be no more tokens after the first one

// Single event

func (r *splunkReceiver) createResourceCustomizer(req *http.Request) func(resource pcommon.Resource) {
	_ = "STUB: not implemented"
	return nil
}

func (r *splunkReceiver) failRequest(
	resp http.ResponseWriter,
	httpStatusCode int,
	jsonResponse []byte,
	err error,
) {
	_ = "STUB: not implemented"
	return
}

// The response needs to be written as a JSON string.

// It handles nil error

func (*splunkReceiver) handleHealthReq(writer http.ResponseWriter, _ *http.Request) {
	_ = "STUB: not implemented"
	return
}

func isFlatJSONField(field any) bool { _ = "STUB: not implemented"; return false }
