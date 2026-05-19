// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package trace // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/skywalkingreceiver/internal/trace"

import (
	"context"
	"net/http"

	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
	common "skywalking.apache.org/repo/goapi/collect/common/v3"
	agent "skywalking.apache.org/repo/goapi/collect/language/agent/v3"
)

const (
	collectorHTTPTransport = "http"
	grpcTransport          = "grpc"
	failing                = "failing"
)

type Receiver struct {
	nextConsumer consumer.Traces
	grpcObsrecv  *receiverhelper.ObsReport
	httpObsrecv  *receiverhelper.ObsReport
	agent.UnimplementedTraceSegmentReportServiceServer
}

// NewReceiver creates a new Receiver reference.
func NewReceiver(nextConsumer consumer.Traces, set receiver.Settings) (*Receiver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Collect implements the service Collect traces func.
func (r *Receiver) Collect(stream agent.TraceSegmentReportService_CollectServer) error {
	_ = "STUB: not implemented"
	return nil
}

// CollectInSync implements the service CollectInSync traces func.
func (r *Receiver) CollectInSync(ctx context.Context, segments *agent.SegmentCollection) (*common.Commands, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func consumeTraces(ctx context.Context, segment *agent.SegmentObject, consumer consumer.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Receiver) HTTPHandler(rsp http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

type Response struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

func ResponseWithJSON(rsp http.ResponseWriter, response *Response, code int) {
	_ = "STUB: not implemented"
	return
}
