// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package dorisexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/dorisexporter"

import (
	"time"

	"go.uber.org/zap"
)

type progressReporter struct {
	name       string
	totalBytes int64
	totalRows  int64
	interval   time.Duration
	logger     *zap.Logger
}

func newProgressReporter(name string, interval int, logger *zap.Logger) *progressReporter {
	_ = "STUB: not implemented"
	return nil
}

func (reporter *progressReporter) incrTotalBytes(bytes int64) { _ = "STUB: not implemented"; return }

func (reporter *progressReporter) incrTotalRows(rows int64) { _ = "STUB: not implemented"; return }

func (reporter *progressReporter) report() { _ = "STUB: not implemented"; return }
