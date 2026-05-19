// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package influxdbreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/influxdbreceiver"

import (
	"github.com/influxdata/influxdb-observability/common"
	"go.uber.org/zap"
)

type zapInfluxLogger struct {
	*zap.SugaredLogger
}

func newZapInfluxLogger(logger *zap.Logger) common.Logger {
	_ = "STUB: not implemented"
	return *new(common.Logger)
}

func (l zapInfluxLogger) Debug(msg string, kv ...any) { _ = "STUB: not implemented"; return }
