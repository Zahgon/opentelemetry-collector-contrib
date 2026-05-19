// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package githubreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/githubreceiver"

import (
	"context"
	"errors"
	"net/http"
	"sync"

	"github.com/google/go-github/v86/github"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
	"go.uber.org/zap"
)

var errMissingEndpoint = errors.New("missing a receiver endpoint")

const healthyResponse = `{"text": "GitHub receiver webhook is healthy"}`

type githubTracesReceiver struct {
	traceConsumer consumer.Traces
	cfg           *Config
	server        *http.Server
	shutdownWG    sync.WaitGroup
	settings      receiver.Settings
	logger        *zap.Logger
	obsrecv       *receiverhelper.ObsReport
	ghClient      *github.Client
}

func newTracesReceiver(
	params receiver.Settings,
	config *Config,
	traceConsumer consumer.Traces,
) (*githubTracesReceiver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (gtr *githubTracesReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// noop if not nil. if start has not been called before these values should be nil.

// create listener from config

// use gorilla mux to set up a router

// setup health route

// setup webhook route for traces

// webhook server standup and configuration

func (gtr *githubTracesReceiver) Shutdown(_ context.Context) error {
	_ = "STUB: not implemented"
	// server must exist to be closed.
	return nil
}

// handleReq handles incoming request sent to the webhook endpoint. On success
// returns a 200 response code.
func (gtr *githubTracesReceiver) handleReq(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Simple healthcheck endpoint.
func (*githubTracesReceiver) handleHealthCheck(w http.ResponseWriter, _ *http.Request) {
	_ = "STUB: not implemented"
	return
}
