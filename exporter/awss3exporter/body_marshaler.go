// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awss3exporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awss3exporter"

import (
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type bodyMarshaler struct{}

func (*bodyMarshaler) format() string { _ = "STUB: not implemented"; return "" }

func newbodyMarshaler() bodyMarshaler { _ = "STUB: not implemented"; return *new(bodyMarshaler) }

func (bodyMarshaler) MarshalLogs(ld plog.Logs) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s bodyMarshaler) MarshalTraces(_ ptrace.Traces) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s bodyMarshaler) MarshalMetrics(_ pmetric.Metrics) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
