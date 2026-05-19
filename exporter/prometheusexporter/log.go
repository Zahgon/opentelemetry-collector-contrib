// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheusexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusexporter"

import (
	"go.uber.org/zap"
)

type promLogger struct {
	realLog *zap.Logger
}

func newPromLogger(zapLog *zap.Logger) *promLogger { _ = "STUB: not implemented"; return nil }

func (l *promLogger) Println(v ...any) { _ = "STUB: not implemented"; return }
