// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package batch // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awskinesisexporter/internal/batch"

import (
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

func partitionByTraceID(v any) string { _ = "STUB: not implemented"; return "" }

type jaegerEncoder struct {
	batchOptions []Option
}

var _ Encoder = (*jaegerEncoder)(nil)

func (je jaegerEncoder) Traces(td ptrace.Traces) (*Batch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (jaegerEncoder) Logs(plog.Logs) (*Batch, error) { _ = "STUB: not implemented"; return nil, nil }
func (jaegerEncoder) Metrics(pmetric.Metrics) (*Batch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
