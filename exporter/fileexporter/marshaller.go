// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package fileexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/fileexporter"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/pprofile"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

// Marshaler configuration used for marshaling Protobuf
var tracesMarshalers = map[string]ptrace.Marshaler{
	formatTypeJSON:  &ptrace.JSONMarshaler{},
	formatTypeProto: &ptrace.ProtoMarshaler{},
}

var metricsMarshalers = map[string]pmetric.Marshaler{
	formatTypeJSON:  &pmetric.JSONMarshaler{},
	formatTypeProto: &pmetric.ProtoMarshaler{},
}

var logsMarshalers = map[string]plog.Marshaler{
	formatTypeJSON:  &plog.JSONMarshaler{},
	formatTypeProto: &plog.ProtoMarshaler{},
}

var profilesMarshalers = map[string]pprofile.Marshaler{
	formatTypeJSON:  &pprofile.JSONMarshaler{},
	formatTypeProto: &pprofile.ProtoMarshaler{},
}

type marshaller struct {
	tracesMarshaler   ptrace.Marshaler
	metricsMarshaler  pmetric.Marshaler
	logsMarshaler     plog.Marshaler
	profilesMarshaler pprofile.Marshaler

	compression string
	compressor  compressFunc

	formatType string
}

func newMarshaller(conf *Config, host component.Host) (*marshaller, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// cast with ok to avoid panics.

func (m *marshaller) marshalTraces(td ptrace.Traces) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *marshaller) marshalMetrics(md pmetric.Metrics) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *marshaller) marshalLogs(ld plog.Logs) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *marshaller) marshalProfiles(pd pprofile.Profiles) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
