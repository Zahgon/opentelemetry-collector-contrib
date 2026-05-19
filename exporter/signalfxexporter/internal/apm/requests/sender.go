// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
// Originally copied from https://github.com/signalfx/signalfx-agent/blob/fbc24b0fdd3884bd0bbfbd69fe3c83f49d4c0b77/pkg/apm/requests/sender.go

package requests // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/apm/requests"

import (
	"context"
	"net/http"
)

type ReqSender struct {
	client      *http.Client
	requests    chan *http.Request
	workerCount uint32
	ctx         context.Context
	clientName  string

	RunningWorkers         int64
	TotalRequestsStarted   int64
	TotalRequestsCompleted int64
	TotalRequestsFailed    int64
}

func NewReqSender(ctx context.Context, client *http.Client, workerCount uint, clientName string) *ReqSender {
	_ = "STUB: not implemented"
	return nil
}

// Unbuffered so that it blocks clients

func (rs *ReqSender) Send(req *http.Request) {
	_ = "STUB: not implemented"
	// Slight optimization to avoid spinning up unnecessary workers if there
	// aren't ever that many dim updates. Once workers start, they remain for the
	// duration of the agent.
	return
}

// Block until we can get the request through, or until the context is cancelled. The request processor
// shuts down when the context has been cancelled, so there's no value added to keep blocking. Blocking
// forever results in Shutdown never completing.

func (rs *ReqSender) processRequests() { _ = "STUB: not implemented"; return }

func (rs *ReqSender) sendRequest(req *http.Request) error { _ = "STUB: not implemented"; return nil }

// If it was successful there is nothing else to do.

type key int

const (
	RequestFailedCallbackKey  key = 1
	RequestSuccessCallbackKey key = 2
)

type (
	RequestFailedCallback  func(body []byte, statusCode int, err error)
	RequestSuccessCallback func([]byte)
)

func onRequestSuccess(req *http.Request, body []byte) { _ = "STUB: not implemented"; return }

func onRequestFailed(req *http.Request, body []byte, statusCode int, err error) {
	_ = "STUB: not implemented"
	return
}

func sendRequest(client *http.Client, req *http.Request) ([]byte, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}
