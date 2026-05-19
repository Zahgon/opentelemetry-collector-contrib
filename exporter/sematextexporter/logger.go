// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sematextexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sematextexporter"

import (
	"github.com/influxdata/influxdb-observability/common"
	"go.uber.org/zap"
)

type zapSematextLogger struct {
	*zap.SugaredLogger
}

func newZapSematextLogger(logger *zap.Logger) common.Logger {
	_ = "STUB: not implemented"
	return *new(common.Logger)
}

func (l zapSematextLogger) Debug(msg string, kv ...any) { _ = "STUB: not implemented"; return }
