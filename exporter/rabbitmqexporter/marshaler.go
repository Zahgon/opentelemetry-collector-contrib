// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package rabbitmqexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/rabbitmqexporter"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type marshaler struct {
	logsMarshaler    plog.Marshaler
	tracesMarshaler  ptrace.Marshaler
	metricsMarshaler pmetric.Marshaler
}

func newMarshaler(encoding *component.ID, host component.Host) (*marshaler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
