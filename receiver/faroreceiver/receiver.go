// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package faroreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/faroreceiver"

import (
	"context"
	"net/http"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
)

const faroPath = "/"

func newFaroReceiver(cfg *Config, set *receiver.Settings) (*faroReceiver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type faroReceiver struct {
	cfg *Config

	serverHTTP *http.Server
	shutdownWg sync.WaitGroup

	nextTraces consumer.Traces
	nextLogs   consumer.Logs

	obsrepHTTP *receiverhelper.ObsReport

	settings *receiver.Settings
}

func (r *faroReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *faroReceiver) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (r *faroReceiver) RegisterTracesConsumer(tc consumer.Traces) {
	_ = "STUB: not implemented"
	return
}

func (r *faroReceiver) RegisterLogsConsumer(lc consumer.Logs) { _ = "STUB: not implemented"; return }

func (r *faroReceiver) startHTTPServer(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	// Noop if not nil
	return nil
}

func (r *faroReceiver) handleFaroRequest(resp http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	// Preflight request
	return
}

type errorResponse struct {
	Error   string `json:"error"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (r *faroReceiver) errorHandler(w http.ResponseWriter, _ *http.Request, errMsg string, statusCode int) {
	_ = "STUB: not implemented"
	return
}

func (r *faroReceiver) failOnErrorsAndLog(errors []string, payload any, resp http.ResponseWriter, req *http.Request) bool {
	_ = "STUB: not implemented"
	return false
}
