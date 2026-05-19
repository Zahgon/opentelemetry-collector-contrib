// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package signalfxexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter"

import (
	"context"
	"errors"
	"net/http"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/dimensions"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/hostmetadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/translation"
	metadata "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata"
)

var errNotStarted = errors.New("exporter has not started")

// TODO: Find a place for this to be shared.
type baseMetricsExporter struct {
	component.Component
	consumer.Metrics
}

// TODO: Find a place for this to be shared.
type baseLogsExporter struct {
	component.Component
	consumer.Logs
}

type signalfMetadataExporter struct {
	exporter.Metrics
	exporter *signalfxExporter
}

func (sme *signalfMetadataExporter) ConsumeMetadata(metadata []*metadata.MetadataUpdate) error {
	_ = "STUB: not implemented"
	return nil
}

type signalfxExporter struct {
	config                 *Config
	version                string
	logger                 *zap.Logger
	telemetrySettings      component.TelemetrySettings
	pushMetricsData        func(ctx context.Context, md pmetric.Metrics) (droppedTimeSeries int, err error)
	hostMetadataSyncer     *hostmetadata.Syncer
	converter              *translation.MetricsConverter
	dimClient              *dimensions.DimensionClient
	eventClient            *sfxEventClient
	entityEventTransformer *dimensions.EntityEventTransformer
}

// newSignalFxExporter returns a new SignalFx exporter.
func newSignalFxExporter(
	config *Config,
	createSettings exporter.Settings,
) (*signalfxExporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if SendOTLPHistograms is true, do not process histograms when converting to SFx

func (se *signalfxExporter) start(ctx context.Context, host component.Host) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (se *signalfxExporter) startDimensionClient(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Duration to wait between property updates.

func newGzipPool() sync.Pool { _ = "STUB: not implemented"; return *new(sync.Pool) }

func newEventExporter(config *Config, createSettings exporter.Settings) (*signalfxExporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (se *signalfxExporter) startLogs(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Initialize dimension client for entity event processing if entity events processing is enabled.

func (se *signalfxExporter) createClient(ctx context.Context, host component.Host) (*http.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (se *signalfxExporter) pushMetrics(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (se *signalfxExporter) pushLogs(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

// Process logs that represent entity events and skip regular event conversion

func (se *signalfxExporter) processEntityEvents(logs plog.LogRecordSlice) error {
	_ = "STUB: not implemented"
	return nil
}

// Failure to accept dimension likely means exceeded buffer. Reject all further
// dimension updates for this export cycle.

func (se *signalfxExporter) shutdown(_ context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (se *signalfxExporter) pushMetadata(metadata []*metadata.MetadataUpdate) error {
	_ = "STUB: not implemented"
	return nil
}

func isEntityEventScope(sl plog.ScopeLogs) bool { _ = "STUB: not implemented"; return false }

func buildHeaders(config *Config, version string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// Add any custom headers from the config. They will override the pre-defined
// ones above in case of conflict, but, not the content encoding one since
// the latter one is defined according to the payload.

// we want to control how headers are set, overriding user headers with our passthrough.
