// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package trace // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/otelarrowreceiver/internal/trace"

import (
	"context"

	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/pdata/ptrace/ptraceotlp"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/otelarrow/admission2"
)

const dataFormatProtobuf = "protobuf"

// Receiver is the type used to handle spans from OpenTelemetry exporters.
type Receiver struct {
	ptraceotlp.UnimplementedGRPCServer
	nextConsumer consumer.Traces
	obsrecv      *receiverhelper.ObsReport
	boundedQueue admission2.Queue
	sizer        *ptrace.ProtoMarshaler
	logger       *zap.Logger
}

// New creates a new Receiver reference.
func New(logger *zap.Logger, nextConsumer consumer.Traces, obsrecv *receiverhelper.ObsReport, bq admission2.Queue) *Receiver {
	_ = "STUB: not implemented"
	return nil
}

// Export implements the service Export traces func.
func (r *Receiver) Export(ctx context.Context, req ptraceotlp.ExportRequest) (ptraceotlp.ExportResponse, error) {
	_ = "STUB: not implemented"
	return *new(ptraceotlp.ExportResponse), nil
}

// immediate release

// Use appropriate status codes for permanent/non-permanent errors.
// If we return the error straightaway, then the grpc implementation will
// set status code to Unknown, which is not retryable.
// See: https://github.com/grpc/grpc-go/blob/v1.59.0/server.go#L1345

func (r *Receiver) Consumer() consumer.Traces {
	_ = "STUB: not implemented"
	return *new(consumer.Traces)
}
