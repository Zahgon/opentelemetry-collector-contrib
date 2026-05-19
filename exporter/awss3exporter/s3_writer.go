// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awss3exporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awss3exporter"

import (
	"context"

	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awss3exporter/internal/upload"
)

func newUploadManager(
	ctx context.Context,
	conf *Config,
	logger *zap.Logger,
	metadata string,
	format string,
	isCompressed bool,
) (upload.Manager, error) {
	_ = "STUB: not implemented"
	return *new(upload.Manager), nil
}
