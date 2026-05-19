// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package consumerretry // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/consumerretry"

import (
	"context"

	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"
)

type logsConsumer struct {
	consumer.Logs
	cfg    Config
	logger *zap.Logger
}

func NewLogs(config Config, logger *zap.Logger, next consumer.Logs) consumer.Logs {
	_ = "STUB: not implemented"
	return *new(consumer.Logs)
}

func (lc *logsConsumer) ConsumeLogs(ctx context.Context, logs plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

// Do not use NewExponentialBackOff since it calls Reset and the code here must
// call Reset after changing the InitialInterval (this saves an unnecessary call to Now).

// TODO: take delay from the error once it is available in the consumererror package.

// back-off, but get interrupted when shutting down or request is cancelled or timed out.
