// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awss3exporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awss3exporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awss3exporter/internal/upload"
)

type s3Exporter struct {
	config     *Config
	signalType string
	uploader   upload.Manager
	logger     *zap.Logger
	marshaler  marshaler
}

func newS3Exporter(
	config *Config,
	signalType string,
	params exporter.Settings,
) *s3Exporter {
	_ = "STUB: not implemented"
	return nil
}

func (e *s3Exporter) getUploadOpts(res pcommon.Resource) *upload.UploadOptions {
	_ = "STUB: not implemented"
	return nil
}

func (e *s3Exporter) start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (*s3Exporter) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (e *s3Exporter) ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *s3Exporter) ConsumeLogs(ctx context.Context, logs plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *s3Exporter) ConsumeTraces(ctx context.Context, traces ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}
