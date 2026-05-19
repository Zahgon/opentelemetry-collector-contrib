// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package marshaler // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/kafkaexporter/internal/marshaler"

import (
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/pprofile"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

var (
	_ LogsMarshaler     = pdataLogsMarshaler{}
	_ MetricsMarshaler  = pdataMetricsMarshaler{}
	_ TracesMarshaler   = pdataTracesMarshaler{}
	_ ProfilesMarshaler = pdataProfilesMarshaler{}
)

type pdataLogsMarshaler struct {
	marshaler plog.Marshaler
}

// NewPdataLogsMarshaler returns a new LogsMarshaler that marshals
// plog.Logs using the given plog.Marshaler. This can be used with
// the standard OTLP marshalers in the plog package, or with encoding
// extensions.
func NewPdataLogsMarshaler(m plog.Marshaler) LogsMarshaler {
	_ = "STUB: not implemented"
	return *new(LogsMarshaler)
}

func (p pdataLogsMarshaler) MarshalLogs(ld plog.Logs, yield func(key, value []byte)) error {
	_ = "STUB: not implemented"
	return nil
}

type pdataMetricsMarshaler struct {
	marshaler pmetric.Marshaler
}

// NewPdataMetricsMarshaler returns a new MetricsMarshaler that marshals
// pmetric.Metrics using the given pmetric.Marshaler. This can be used
// with the standard OTLP marshalers in the pmetric package, or with
// encoding extensions.
func NewPdataMetricsMarshaler(m pmetric.Marshaler) MetricsMarshaler {
	_ = "STUB: not implemented"
	return *new(MetricsMarshaler)
}

func (p pdataMetricsMarshaler) MarshalMetrics(ld pmetric.Metrics, yield func(key, value []byte)) error {
	_ = "STUB: not implemented"
	return nil
}

type pdataTracesMarshaler struct {
	marshaler ptrace.Marshaler
}

// NewPdataTracesMarshaler returns a new TracesMarshaler that marshals
// ptrace.Traces using the given ptrace.Marshaler. This can be used
// with the standard OTLP marshalers in the ptrace package, or with
// encoding extensions.
func NewPdataTracesMarshaler(m ptrace.Marshaler) TracesMarshaler {
	_ = "STUB: not implemented"
	return *new(TracesMarshaler)
}

func (p pdataTracesMarshaler) MarshalTraces(td ptrace.Traces, yield func(key, value []byte)) error {
	_ = "STUB: not implemented"
	return nil
}

type pdataProfilesMarshaler struct {
	marshaler pprofile.Marshaler
}

// NewPdataProfilesMarshaler returns a new ProfilesMarshaler that marshals
// pprofile.Profiles using the given pprofile.Marshaler. This can be used with
// the standard OTLP marshalers in the pprofile package, or with encoding
// extensions.
func NewPdataProfilesMarshaler(m pprofile.Marshaler) ProfilesMarshaler {
	_ = "STUB: not implemented"
	return *new(ProfilesMarshaler)
}

func (p pdataProfilesMarshaler) MarshalProfiles(ld pprofile.Profiles, yield func(key, value []byte)) error {
	_ = "STUB: not implemented"
	return nil
}
