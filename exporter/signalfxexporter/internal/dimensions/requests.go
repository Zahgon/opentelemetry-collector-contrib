// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package dimensions // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/dimensions"

import (
	"context"
	"net/http"
	"sync/atomic"
)

// ReqSender is a direct port of
// https://github.com/signalfx/signalfx-agent/blob/main/pkg/core/writer/requests/sender.go.
type ReqSender struct {
	client               *http.Client
	requests             chan *http.Request
	workerCount          uint
	additionalDimensions map[string]string
	runningWorkers       *atomic.Int64
}

func NewReqSender(client *http.Client,
	workerCount uint, diagnosticDimensions map[string]string,
) *ReqSender {
	_ = "STUB: not implemented"
	return nil
}

// Unbuffered so that it blocks clients

// Send sends the request. Not thread-safe.
func (rs *ReqSender) Send(ctx context.Context, req *http.Request) {
	_ = "STUB: not implemented"
	// Slight optimization to avoid spinning up unnecessary workers if there
	// aren't ever that many dim updates. Once workers start, they remain for the
	// duration of the agent.
	return
}

// Block until we can get through a request, unless context has been cancelled.

func (rs *ReqSender) processRequests(ctx context.Context) { _ = "STUB: not implemented"; return }

func (rs *ReqSender) sendRequest(req *http.Request) error { _ = "STUB: not implemented"; return nil }

// If it was successful there is nothing else to do.

type key int

const (
	RequestFailedCallbackKey  key = 1
	RequestSuccessCallbackKey key = 2
)

type (
	RequestFailedCallback  func(statusCode int, err error)
	RequestSuccessCallback func([]byte)
)

func onRequestSuccess(req *http.Request, body []byte) { _ = "STUB: not implemented"; return }

func onRequestFailed(req *http.Request, statusCode int, err error) {
	_ = "STUB: not implemented"
	return
}

func sendRequest(client *http.Client, req *http.Request) ([]byte, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}
