// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awsfirehosereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver"

import (
	"context"
	"errors"
	"net/http"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/receiver"
)

const (
	headerFirehoseRequestID        = "X-Amz-Firehose-Request-Id"
	headerFirehoseAccessKey        = "X-Amz-Firehose-Access-Key"
	headerFirehoseCommonAttributes = "X-Amz-Firehose-Common-Attributes"
	headerContentType              = "Content-Type"
	headerContentLength            = "Content-Length"
)

var (
	errInvalidAccessKey         = errors.New("invalid firehose access key")
	errInHeaderMissingRequestID = errors.New("missing request id in header")
	errInBodyMissingRequestID   = errors.New("missing request id in body")
	errInBodyDiffRequestID      = errors.New("different request id in body")
)

// The firehoseConsumer is responsible for using the unmarshaler and the consumer.
type firehoseConsumer interface {
	Start(context.Context, component.Host) error

	// Consume unmarshals and consumes the records returned by f.
	Consume(ctx context.Context, f nextRecordFunc, commonAttributes map[string]string) (int, error)
}

// nextRecordFunc is a function provided to consumers for obtaining the
// next record to consume. The function returns (nil, io.EOF) when there
// are no more records.
type nextRecordFunc func() ([]byte, error)

// firehoseReceiver
type firehoseReceiver struct {
	// settings is the base receiver settings.
	settings receiver.Settings
	// config is the configuration for the receiver.
	config *Config
	// server is the HTTP/HTTPS server set up to listen
	// for requests.
	server *http.Server
	// shutdownWG is the WaitGroup that is used to wait until
	// the server shutdown has completed.
	shutdownWG sync.WaitGroup
	// consumer is the firehoseConsumer to use to process/send
	// the records in each request.
	consumer firehoseConsumer
}

// The firehoseRequest is the format of the received request body.
type firehoseRequest struct {
	// RequestID is a GUID that should be the same value as
	// the one in the header.
	RequestID string `json:"requestId"`
	// Timestamp is the milliseconds since epoch for when the
	// request was generated.
	Timestamp int64 `json:"timestamp"`
	// Records contains the data.
	Records []firehoseRecord `json:"records"`
}

// The firehoseRecord is an individual record within the firehoseRequest.
type firehoseRecord struct {
	// Data is a base64 encoded string. Can be empty.
	Data string `json:"data"`
}

// The firehoseResponse is the expected body for the response back to
// the delivery stream.
type firehoseResponse struct {
	// RequestID is the same GUID that was received in
	// the request.
	RequestID string `json:"requestId"`
	// Timestamp is the milliseconds since epoch for when the
	// request finished being processed.
	Timestamp int64 `json:"timestamp"`
	// ErrorMessage is the error to report. Empty if request
	// was successfully processed.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// The firehoseCommonAttributes is the format for the common attributes
// found in the header of requests.
type firehoseCommonAttributes struct {
	// CommonAttributes can be set when creating the delivery stream.
	// These will be passed to the firehoseConsumer, which should
	// attach the attributes.
	CommonAttributes map[string]string `json:"commonAttributes"`
}

var (
	_ receiver.Metrics = (*firehoseReceiver)(nil)
	_ http.Handler     = (*firehoseReceiver)(nil)
)

// Start spins up the receiver's HTTP server and makes the receiver start
// its processing.
func (fmr *firehoseReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Shutdown tells the receiver that should stop reception,
// giving it a chance to perform any necessary clean-up and
// shutting down its HTTP server.
func (fmr *firehoseReceiver) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

// ServeHTTP receives Firehose requests, unmarshalls them, and sends them along to the firehoseConsumer,
// which is responsible for unmarshalling the records and sending them to the next consumer.
func (fmr *firehoseReceiver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// validate checks the Firehose access key in the header against
// the one passed into the Config
func (fmr *firehoseReceiver) validate(r *http.Request) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// No access key is configured - accept all requests.

// getCommonAttributes unmarshalls the common attributes from the request header
func (*firehoseReceiver) getCommonAttributes(r *http.Request) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// sendResponse writes a response to Firehose in the expected format.
func (fmr *firehoseReceiver) sendResponse(w http.ResponseWriter, requestID string, statusCode int, err error) {
	_ = "STUB: not implemented"
	return
}

func gunzipRecordIfNeeded(record []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// loadEncodingExtension tries to load an available extension for the given encoding.
func loadEncodingExtension[T any](host component.Host, encoding, signalType string) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

// encodingToComponentID attempts to parse the encoding string as a component ID.
func encodingToComponentID(encoding string) (*component.ID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
