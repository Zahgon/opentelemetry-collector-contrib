// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/otelarrowreceiver/internal/metrics"

import (
	"context"

	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/pmetric/pmetricotlp"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/otelarrow/admission2"
)

const dataFormatProtobuf = "protobuf"

// Receiver is the type used to handle metrics from OpenTelemetry exporters.
type Receiver struct {
	pmetricotlp.UnimplementedGRPCServer
	nextConsumer consumer.Metrics
	obsrecv      *receiverhelper.ObsReport
	boundedQueue admission2.Queue
	sizer        *pmetric.ProtoMarshaler
	logger       *zap.Logger
}

// New creates a new Receiver reference.
func New(logger *zap.Logger, nextConsumer consumer.Metrics, obsrecv *receiverhelper.ObsReport, bq admission2.Queue) *Receiver {
	_ = "STUB: not implemented"
	return nil
}

// Export implements the service Export metrics func.
func (r *Receiver) Export(ctx context.Context, req pmetricotlp.ExportRequest) (pmetricotlp.ExportResponse, error) {
	_ = "STUB: not implemented"
	return *new(pmetricotlp.ExportResponse), nil
}

// immediate release

// Use appropriate status codes for permanent/non-permanent errors.
// If we return the error straightaway, then the grpc implementation will
// set status code to Unknown, which is not retryable.
// See: https://github.com/grpc/grpc-go/blob/v1.59.0/server.go#L1345

func (r *Receiver) Consumer() consumer.Metrics {
	_ = "STUB: not implemented"
	return *new(consumer.Metrics)
}
