// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azureeventhubreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azureeventhubreceiver"

import (
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"
)

type rawLogsUnmarshaler struct {
	logger *zap.Logger
}

func newRawLogsUnmarshaler(logger *zap.Logger) eventLogsUnmarshaler {
	_ = "STUB: not implemented"
	return *new(eventLogsUnmarshaler)
}

func (rawLogsUnmarshaler) UnmarshalLogs(event *azureEvent) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}
