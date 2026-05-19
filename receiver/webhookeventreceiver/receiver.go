// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package webhookeventreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/webhookeventreceiver"

import (
	"context"
	"errors"
	"net/http"
	"regexp"
	"sync"

	"github.com/julienschmidt/httprouter"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
)

var (
	errNilLogsConsumer       = errors.New("missing a logs consumer")
	errInvalidRequestMethod  = errors.New("invalid method. Valid method is POST")
	errInvalidEncodingType   = errors.New("invalid encoding type")
	errEmptyResponseBody     = errors.New("request body content length is zero")
	errMissingRequiredHeader = errors.New("request was missing required header or incorrect header value")
)

const healthyResponse = `{"text": "Webhookevent receiver is healthy"}`

type eventReceiver struct {
	settings            receiver.Settings
	cfg                 *Config
	logConsumer         consumer.Logs
	server              *http.Server
	shutdownWG          sync.WaitGroup
	obsrecv             *receiverhelper.ObsReport
	gzipPool            *sync.Pool
	includeHeadersRegex *regexp.Regexp
	maxRequestBodySize  int // Computed max token size for scanner (minimum 64KB)
}

func newLogsReceiver(params receiver.Settings, cfg Config, consumer consumer.Logs) (receiver.Logs, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Logs), nil
}

// Valdiate() call above has already ensured this will compile

// create eventReceiver instance

// Start function manages receiver startup tasks. part of the receiver.Logs interface.
func (er *eventReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	// noop if not nil. if start has not been called before these values should be nil.
	return nil
}

// create listener from config

// set up router.

// webhook server standup and configuration

// set timeouts

// shutdown

// Shutdown function manages receiver shutdown tasks. part of the receiver.Logs interface.
func (er *eventReceiver) Shutdown(context.Context) error {
	_ = "STUB: not implemented"
	// server must exist to be closed.
	return nil
}

// handleReq handles incoming request from webhook. On success returns a 200 response code to the webhook
func (er *eventReceiver) handleReq(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_ = "STUB: not implemented"
	return
}

// only support gzip if encoding header is set.

// gzip encoded case

// send body into a scanner and then convert the request body into a log

// Simple healthcheck endpoint.
func (*eventReceiver) handleHealthCheck(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	_ = "STUB: not implemented"
	return
}

// write response on a failed/bad request. Generates a small json body based on the thrown by
// the handle func and the appropriate http status code. many webhooks will either log these responses or
// notify webhook users should a none 2xx code be detected.
func (er *eventReceiver) failBadReq(_ context.Context,
	w http.ResponseWriter,
	httpStatusCode int,
	err error,
) {
	_ = "STUB: not implemented"
	return
}

// write response to webhook

// log bad webhook request if debug is enabled
