// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheusremotewriteexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusremotewriteexporter"

import (
	"context"

	"go.uber.org/zap"
)

type ctxKey int

const (
	loggerCtxKey ctxKey = iota
)

func contextWithLogger(ctx context.Context, log *zap.Logger) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func loggerFromContext(ctx context.Context) (*zap.Logger, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
