// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package producer // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awskinesisexporter/internal/producer"

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awskinesisexporter/internal/batch"
)

type batcher struct {
	stream *string

	client Kinesis
	log    *zap.Logger
}

var _ Batcher = (*batcher)(nil)

var (
	permanentErrResourceNotFound = new(*types.ResourceNotFoundException)
	permanentErrInvalidArgument  = new(*types.InvalidArgumentException)
)

func NewBatcher(kinesisAPI Kinesis, stream string, opts ...BatcherOptions) (Batcher, error) {
	_ = "STUB: not implemented"
	return *new(Batcher), nil
}

func (b *batcher) Put(ctx context.Context, bt *batch.Batch) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *batcher) Ready(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
