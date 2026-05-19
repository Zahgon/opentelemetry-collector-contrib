// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package lokireceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/lokireceiver"

import (
	"context"
	"net/http"
	"sync"

	"github.com/grafana/loki/pkg/push"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
	"google.golang.org/grpc"
)

const (
	pbContentType   = "application/x-protobuf"
	jsonContentType = "application/json"
)

const ErrAtLeastOneEntryFailedToProcess = "at least one entry in the push request failed to process"

type lokiReceiver struct {
	conf         *Config
	nextConsumer consumer.Logs
	settings     receiver.Settings
	httpMux      *http.ServeMux
	serverHTTP   *http.Server
	serverGRPC   *grpc.Server
	shutdownWG   sync.WaitGroup
	httpAddr     string
	grpcAddr     string

	obsrepGRPC *receiverhelper.ObsReport
	obsrepHTTP *receiverhelper.ObsReport
}

func newLokiReceiver(conf *Config, nextConsumer consumer.Logs, settings receiver.Settings) (*lokiReceiver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *lokiReceiver) startProtocolsServers(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *lokiReceiver) startHTTPServer(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *lokiReceiver) startGRPCServer(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *lokiReceiver) Push(ctx context.Context, pushRequest *push.PushRequest) (*push.PushResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *lokiReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *lokiReceiver) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func handleUnmatchedMethod(resp http.ResponseWriter) { _ = "STUB: not implemented"; return }

func handleUnmatchedContentType(resp http.ResponseWriter) { _ = "STUB: not implemented"; return }

func writeResponse(w http.ResponseWriter, contentType string, statusCode int, msg []byte) {
	_ = "STUB: not implemented"
	return
}

// Nothing we can do with the error if we cannot write to the response.

func handleLogs(resp http.ResponseWriter, req *http.Request, r *lokiReceiver) {
	_ = "STUB: not implemented"
	return
}
