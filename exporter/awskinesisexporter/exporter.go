// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awskinesisexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awskinesisexporter"

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awskinesisexporter/internal/batch"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awskinesisexporter/internal/producer"
)

// kinesisExporter implements an OpenTelemetry trace exporter that exports all spans to AWS Kinesis
type kinesisExporter struct {
	producer producer.Batcher
	batcher  batch.Encoder
}

// options is used to override the default shipped behavior
// to allow for testing correct setup of components
type options struct {
	NewKinesisClient func(conf aws.Config, opts ...func(*kinesis.Options)) *kinesis.Client
}

func createExporter(ctx context.Context, c component.Config, log *zap.Logger, opts ...func(opt *options)) (*kinesisExporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// start validates that the Kinesis stream is available.
func (e kinesisExporter) start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// ConsumeTraces receives a span batch and exports it to AWS Kinesis
func (e kinesisExporter) consumeTraces(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func (e kinesisExporter) consumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (e kinesisExporter) consumeLogs(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}
