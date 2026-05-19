// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otlpencodingextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/otlpencodingextension"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/pprofile"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding"
)

const (
	otlpProto = "otlp_proto"
	otlpJSON  = "otlp_json"
)

var (
	_ encoding.TracesMarshalerExtension     = (*otlpExtension)(nil)
	_ encoding.TracesUnmarshalerExtension   = (*otlpExtension)(nil)
	_ encoding.LogsMarshalerExtension       = (*otlpExtension)(nil)
	_ encoding.LogsUnmarshalerExtension     = (*otlpExtension)(nil)
	_ encoding.MetricsMarshalerExtension    = (*otlpExtension)(nil)
	_ encoding.MetricsUnmarshalerExtension  = (*otlpExtension)(nil)
	_ encoding.ProfilesMarshalerExtension   = (*otlpExtension)(nil)
	_ encoding.ProfilesUnmarshalerExtension = (*otlpExtension)(nil)
)

type otlpExtension struct {
	config             *Config
	traceMarshaler     ptrace.Marshaler
	traceUnmarshaler   ptrace.Unmarshaler
	logMarshaler       plog.Marshaler
	logUnmarshaler     plog.Unmarshaler
	metricMarshaler    pmetric.Marshaler
	metricUnmarshaler  pmetric.Unmarshaler
	profileMarshaler   pprofile.Marshaler
	profileUnmarshaler pprofile.Unmarshaler
}

func newExtension(config *Config) (*otlpExtension, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ex *otlpExtension) UnmarshalTraces(buf []byte) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

func (ex *otlpExtension) MarshalTraces(traces ptrace.Traces) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ex *otlpExtension) UnmarshalMetrics(buf []byte) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

func (ex *otlpExtension) MarshalMetrics(metrics pmetric.Metrics) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ex *otlpExtension) UnmarshalLogs(buf []byte) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

func (ex *otlpExtension) MarshalLogs(logs plog.Logs) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalProfiles implements encoding.ProfilesUnmarshalerExtension.
func (ex *otlpExtension) UnmarshalProfiles(buf []byte) (pprofile.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(pprofile.Profiles), nil
}

// MarshalProfiles implements encoding.ProfilesMarshalerExtension.
func (ex *otlpExtension) MarshalProfiles(profiles pprofile.Profiles) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*otlpExtension) Start(context.Context, component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (*otlpExtension) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }
