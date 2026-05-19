// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package golden // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden"

import (
	"testing"

	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/pprofile"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

// ReadMetrics reads a pmetric.Metrics from the specified YAML or JSON file.
func ReadMetrics(filePath string) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// WriteMetrics writes a pmetric.Metrics to the specified file in YAML format.
func WriteMetrics(tb testing.TB, filePath string, metrics pmetric.Metrics, opts ...WriteMetricsOption) error {
	_ = "STUB: not implemented"
	return nil
}

// MarshalMetricsYAML marshals a pmetric.Metrics to YAML format.
func MarshalMetricsYAML(metrics pmetric.Metrics) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WriteMetricsToFile writes a pmetric.Metrics to the specified file in YAML format.
// Prefer using WriteMetrics in tests.
func WriteMetricsToFile(filePath string, metrics pmetric.Metrics, opts ...WriteMetricsOption) error {
	_ = "STUB: not implemented"
	return nil
}

// ReadLogs reads a plog.Logs from the specified YAML or JSON file.
func ReadLogs(filePath string) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

// WriteLogs writes a plog.Logs to the specified file in YAML format.
func WriteLogs(tb testing.TB, filePath string, logs plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteLogsToFile writes a plog.Logs to the specified file in YAML format.
// Prefer using WriteLogs in tests.
func WriteLogsToFile(filePath string, logs plog.Logs) error { _ = "STUB: not implemented"; return nil }

// ReadTraces reads a ptrace.Traces from the specified YAML or JSON file.
func ReadTraces(filePath string) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

// WriteTraces writes a ptrace.Traces to the specified file in YAML format.
func WriteTraces(tb testing.TB, filePath string, traces ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteTracesToFile writes a ptrace.Traces to the specified file
// Prefer using WriteTraces in tests.
func WriteTracesToFile(filePath string, traces ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

// ReadProfiles reads a pprofile.Profiles from the specified YAML or JSON file.
func ReadProfiles(filePath string) (pprofile.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(pprofile.Profiles), nil
}

// WriteProfiles writes a pprofile.Profiles to the specified file in YAML format.
func WriteProfiles(tb testing.TB, filePath string, profiles pprofile.Profiles) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteProfilesToFile writes a pprofile.Profiles to the specified file in YAML format.
// Prefer using WriteProfiles in tests.
func WriteProfilesToFile(filePath string, profiles pprofile.Profiles) error {
	_ = "STUB: not implemented"
	return nil
}
