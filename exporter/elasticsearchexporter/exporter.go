// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package elasticsearchexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/pprofile"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/datapoints"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/metadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/metricgroup"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/pool"
)

type elasticsearchExporter struct {
	set                 exporter.Settings
	config              *Config
	index               string
	logstashFormat      LogstashFormatSettings
	defaultMappingMode  MappingMode
	allowedMappingModes map[string]MappingMode
	bulkIndexers        bulkIndexers
	bufferPool          *pool.BufferPool

	documentEncoders         [NumMappingModes]documentEncoder
	documentRouters          [NumMappingModes]documentRouter
	spanEventDocumentRouters [NumMappingModes]documentRouter

	telemetryBuilder *metadata.TelemetryBuilder
}

func newExporter(cfg *Config, set exporter.Settings, index string) (*elasticsearchExporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *elasticsearchExporter) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *elasticsearchExporter) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *elasticsearchExporter) pushLogsData(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *elasticsearchExporter) pushLogRecord(
	ctx context.Context,
	router documentRouter,
	encoder documentEncoder,
	ec encodingContext,
	record plog.LogRecord,
	bulkIndexerSession bulkIndexerSession,
) error {
	_ = "STUB: not implemented"
	return nil
}

// not recycling after Add returns an error as we don't know if it's already recycled

type dataPointsGroup struct {
	resource          pcommon.Resource
	resourceSchemaURL string
	scope             pcommon.InstrumentationScope
	scopeSchemaURL    string
	dataPoints        []datapoints.DataPoint
}

func (p *dataPointsGroup) addDataPoint(dp datapoints.DataPoint) { _ = "STUB: not implemented"; return }

func (e *elasticsearchExporter) pushMetricsData(ctx context.Context, metrics pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

// Maintain a 2 layer map to avoid storing lots of copies of index strings

// log instead of returning these so that upstream does not retry

// not recycling after Add returns an error as we don't know if it's already recycled

func (e *elasticsearchExporter) pushTraceData(
	ctx context.Context,
	td ptrace.Traces,
) error {
	_ = "STUB: not implemented"
	// Get the partioner key from the context
	// Decode the key to get the info
	return nil
}

func (e *elasticsearchExporter) pushTraceRecord(
	ctx context.Context,
	router documentRouter,
	encoder documentEncoder,
	ec encodingContext,
	span ptrace.Span,
	bulkIndexerSession bulkIndexerSession,
) error {
	_ = "STUB: not implemented"
	return nil
}

// not recycling after Add returns an error as we don't know if it's already recycled

func (e *elasticsearchExporter) pushSpanEvent(
	ctx context.Context,
	router documentRouter,
	encoder documentEncoder,
	ec encodingContext,
	span ptrace.Span,
	spanEvent ptrace.SpanEvent,
	bulkIndexerSession bulkIndexerSession,
) error {
	_ = "STUB: not implemented"
	return nil
}

// not recycling after Add returns an error as we don't know if it's already recycled

// controlAttrs holds the values of control-channel attributes the orchestrator
// needs before encoding a doc: whether to skip indexing entirely (_noindex
// mapping hint), the dynamic document ID, and the dynamic ingest pipeline.
type controlAttrs struct {
	noindex  bool
	docID    string
	pipeline string
}

// extractControlAttrs walks attrs once, collecting every control-channel value
// the push functions would otherwise read via separate pcommon.Map.Get calls.
// captureDocID and capturePipeline mirror the existing config gates
// (Logs/TracesDynamicID.Enabled, LogsDynamicPipeline.Enabled): when false, the
// corresponding value is left empty even if the attribute is present.
func extractControlAttrs(attrs pcommon.Map, captureDocID, capturePipeline bool) controlAttrs {
	_ = "STUB: not implemented"
	return *new(controlAttrs)
}

// MappingHintsAttrKey is always a candidate

// If _noindex is specified, nothing else matters.

func (e *elasticsearchExporter) pushProfilesData(ctx context.Context, pd pprofile.Profiles) error {
	_ = "STUB: not implemented"
	// TODO add support for routing profiles to different data_stream.namespaces?
	return nil
}

// scopeMappingModeSessions is used to create the default session according to
// the specified mapping mode.

func (*elasticsearchExporter) pushProfileRecord(
	ctx context.Context,
	encoder documentEncoder,
	ec encodingContext,
	dic pprofile.ProfilesDictionary,
	profile pprofile.Profile,
	defaultSession, eventsSession, stackTracesSession, stackFramesSession, executablesSession bulkIndexerSession,
) error {
	_ = "STUB: not implemented"
	return nil
}

// These regular indices have a low write-frequency and can share the executablesSession.

// mappingModeSessions holds mapping-mode specific bulk indexer sessions.
type mappingModeSessions struct {
	indexers *[NumMappingModes]bulkIndexer
	sessions [NumMappingModes]bulkIndexerSession
	sessionList
}

// StartSession starts a new session for the given mapping mode if one has
// not yet been started, otherwise it returns the existing session.
//
// Note: this is not safe for concurrent use. It is expected to be used
// within a single Consume* call.
func (s *mappingModeSessions) StartSession(ctx context.Context, mappingMode MappingMode) bulkIndexerSession {
	_ = "STUB: not implemented"
	return *new(bulkIndexerSession)
}

// sessionList holds a list of bulkIndexerSession instances.
//
// This provides Flush and End methods that flush/end all sessions in the list.
type sessionList []bulkIndexerSession

// Flush concurrently flushes all sessions.
func (sessions *sessionList) Flush(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// End ends all sessions.
func (sessions *sessionList) End() { _ = "STUB: not implemented"; return }

func (e *elasticsearchExporter) getRequestMappingMode(ctx context.Context) (MappingMode, error) {
	_ = "STUB: not implemented"
	return *new(MappingMode), nil
}

func (e *elasticsearchExporter) getScopeMappingMode(
	scope pcommon.InstrumentationScope, defaultMode MappingMode,
) (MappingMode, error) {
	_ = "STUB: not implemented"
	return *new(MappingMode), nil
}

func (e *elasticsearchExporter) parseMappingMode(s string) (MappingMode, error) {
	_ = "STUB: not implemented"
	return *new(MappingMode), nil
}

func newDataPointHasher(mode MappingMode) metricgroup.DataPointHasher {
	_ = "STUB: not implemented"
	return *new(metricgroup.DataPointHasher)
}

// Defaults to ECS for backward compatibility
