// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package batch // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awskinesisexporter/internal/batch"

import (
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type unsupported struct{}

var (
	_ ptrace.Marshaler  = (*unsupported)(nil)
	_ pmetric.Marshaler = (*unsupported)(nil)
	_ plog.Marshaler    = (*unsupported)(nil)
)

func (unsupported) MarshalTraces(_ ptrace.Traces) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (unsupported) MarshalMetrics(_ pmetric.Metrics) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (unsupported) MarshalLogs(_ plog.Logs) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
