// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package logs // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/otelarrowreceiver/internal/logs"

import (
	"context"

	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/plog/plogotlp"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/otelarrow/admission2"
)

const dataFormatProtobuf = "protobuf"

// Receiver is the type used to handle logs from OpenTelemetry exporters.
type Receiver struct {
	plogotlp.UnimplementedGRPCServer
	nextConsumer consumer.Logs
	obsrecv      *receiverhelper.ObsReport
	boundedQueue admission2.Queue
	sizer        *plog.ProtoMarshaler
	logger       *zap.Logger
}

// New creates a new Receiver reference.
func New(logger *zap.Logger, nextConsumer consumer.Logs, obsrecv *receiverhelper.ObsReport, bq admission2.Queue) *Receiver {
	_ = "STUB: not implemented"
	return nil
}

// Export implements the service Export logs func.
func (r *Receiver) Export(ctx context.Context, req plogotlp.ExportRequest) (plogotlp.ExportResponse, error) {
	_ = "STUB: not implemented"
	return *new(plogotlp.ExportResponse), nil
}

// immediate release

// Use appropriate status codes for permanent/non-permanent errors.
// If we return the error straightaway, then the grpc implementation will
// set status code to Unknown, which is not retryable.
// See: https://github.com/grpc/grpc-go/blob/v1.59.0/server.go#L1345

func (r *Receiver) Consumer() consumer.Logs { _ = "STUB: not implemented"; return *new(consumer.Logs) }
