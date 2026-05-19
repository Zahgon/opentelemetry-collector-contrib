// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package alertmanagerexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/alertmanagerexporter"

import (
	"context"
	"net/http"

	"github.com/prometheus/common/model"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type alertmanagerExporter struct {
	config            *Config
	client            *http.Client
	tracesMarshaler   ptrace.Marshaler
	settings          component.TelemetrySettings
	endpoint          string
	generatorURL      string
	defaultSeverity   string
	severityAttribute string
	apiVersion        string
}

type alertmanagerEvent struct {
	spanEvent ptrace.SpanEvent
	traceID   string
	spanID    string
	severity  string
}

func (s *alertmanagerExporter) convertEventSliceToArray(eventSlice ptrace.SpanEventSlice, traceID pcommon.TraceID, spanID pcommon.SpanID) []*alertmanagerEvent {
	_ = "STUB: not implemented"
	return nil
}

func (s *alertmanagerExporter) extractEvents(td ptrace.Traces) []*alertmanagerEvent {
	_ = "STUB: not implemented"
	// Stitch parent trace ID and span ID
	return nil
}

func createAnnotations(event *alertmanagerEvent) model.LabelSet {
	_ = "STUB: not implemented"
	return *new(model.LabelSet)
}

func (s *alertmanagerExporter) createLabels(event *alertmanagerEvent) model.LabelSet {
	_ = "STUB: not implemented"
	return *new(model.LabelSet)
}

func (s *alertmanagerExporter) convertEventsToAlertPayload(events []*alertmanagerEvent) []model.Alert {
	_ = "STUB: not implemented"
	return nil
}

func (s *alertmanagerExporter) postAlert(ctx context.Context, payload []model.Alert) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *alertmanagerExporter) pushTraces(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *alertmanagerExporter) start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *alertmanagerExporter) shutdown(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func newAlertManagerExporter(cfg *Config, set component.TelemetrySettings) *alertmanagerExporter {
	_ = "STUB: not implemented"
	return nil
}

func newTracesExporter(ctx context.Context, cfg component.Config, set exporter.Settings) (exporter.Traces, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Traces), nil
}
