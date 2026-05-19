// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package elasticsearchexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter"

import (
	"context"
	"io"
	"sync"
	"time"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-docappender/v2"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/metadata"
)

type bulkIndexer interface {
	// StartSession starts a new bulk indexing session.
	StartSession(context.Context) bulkIndexerSession

	// Close closes the bulk indexer, ending any in-progress
	// sessions and stopping any background processing.
	Close(ctx context.Context) error
}

type bulkIndexerSession interface {
	// Add adds a document to the bulk indexing session.
	Add(ctx context.Context, index, docID, pipeline string, document io.WriterTo, dynamicTemplates map[string]string, action string) error

	// End must be called on the session object once it is no longer
	// needed, in order to release any associated resources.
	//
	// Note that ending the session does _not_ implicitly flush
	// documents. Call Flush before calling End as needed.
	//
	// Calling other methods (including End) after End may panic.
	End()

	// Flush flushes any documents added to the bulk indexing session.
	Flush(context.Context) error
}

const defaultMaxRetries = 2

const (
	// errorHintKnownIssues is the hint message for errors that are documented in the Known issues section.
	errorHintKnownIssues = "check the \"Known issues\" section of Elasticsearch Exporter docs"
	// errorHintOTelMappingMode is the hint message for illegal_argument_exception
	// related to require_data_stream with incompatible Elasticsearch versions in OTel mode.
	errorHintOTelMappingMode = "OTel mapping mode requires Elasticsearch 8.12+ (see Known issues in README)"
	// errorHintECSMappingMode is the hint message for illegal_argument_exception
	// related to require_data_stream with incompatible Elasticsearch versions in ECS mode.
	errorHintECSMappingMode = "ECS mapping mode requires Elasticsearch 8.12+ (see Known issues in README)"
)

func newBulkIndexer(
	client elastictransport.Interface,
	config *Config,
	requireDataStream bool,
	tb *metadata.TelemetryBuilder,
	logger *zap.Logger,
	getErrorHintFunc func(index, errorType string) string,
) bulkIndexer {
	_ = "STUB: not implemented"
	return *new(bulkIndexer)
}

func bulkIndexerConfig(client elastictransport.Interface, config *Config, requireDataStream bool, logger *zap.Logger) docappender.BulkIndexerConfig {
	_ = "STUB: not implemented"
	return *new(docappender.BulkIndexerConfig)
}

func getQueryParamsFromEndpoint(config *Config, logger *zap.Logger) (queryParams map[string][]string) {
	_ = "STUB: not implemented"
	return nil
}

// we check the query params set on the first endpoint only
// this is enough to replicate to all requests

func bulkIndexerIncludeSourceOnError(includeSourceOnError *bool) docappender.Value {
	_ = "STUB: not implemented"
	return *new(docappender.Value)
}

func newSyncBulkIndexer(
	client elastictransport.Interface,
	config *Config,
	requireDataStream bool,
	tb *metadata.TelemetryBuilder,
	logger *zap.Logger,
	getErrorHintFunc func(index, errorType string) string,
) *syncBulkIndexer {
	_ = "STUB: not implemented"
	return nil
}

type syncBulkIndexer struct {
	config                 docappender.BulkIndexerConfig
	maxFlushBytes          int64
	flushTimeout           time.Duration
	retryConfig            RetrySettings
	metadataKeys           []string
	telemetryBuilder       *metadata.TelemetryBuilder
	logger                 *zap.Logger
	failedDocsInputLogger  *zap.Logger
	getErrorHintFunc       func(index, errorType string) string
	requireDataStream      bool
	suppressConflictErrors bool
}

// StartSession creates a new docappender.BulkIndexer, and wraps
// it with a syncBulkIndexerSession.
func (s *syncBulkIndexer) StartSession(ctx context.Context) bulkIndexerSession {
	_ = "STUB: not implemented"
	return *new(bulkIndexerSession)
}

// This should never happen in practice:
// NewBulkIndexer should only fail if the
// config is invalid, and we expect it to
// always be valid at this point.

// Compute the docs received attribute set once per session.
// Metadata keys are constant within a single request context,
// so recomputing them per document is wasteful.

// Close is a no-op.
func (*syncBulkIndexer) Close(context.Context) error { _ = "STUB: not implemented"; return nil }

type syncBulkIndexerSession struct {
	s                *syncBulkIndexer
	bi               *docappender.BulkIndexer
	docsReceivedAttr metric.MeasurementOption
}

// Add adds an item to the sync bulk indexer session.
func (s *syncBulkIndexerSession) Add(ctx context.Context, index, docID, pipeline string, document io.WriterTo, dynamicTemplates map[string]string, action string) error {
	_ = "STUB: not implemented"
	return nil
}

// sending_queue operates on flush sizes based on pdata model whereas bulk
// indexers operate on ndjson. Force a flush if the ndjson size is too large.
// when the uncompressed length exceeds the configured max flush size.

// End is a no-op.
func (*syncBulkIndexerSession) End() {
	_ = "STUB: not implemented"
	// TODO acquire docappender.BulkIndexer from pool in StartSession, release here
	return
}

// Flush flushes documents added to the bulk indexer session.
func (s *syncBulkIndexerSession) Flush(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// No documents in buffer waiting for per-document retry, exit retry loop.

// BUG: This should never happen in practice.
// When retry is disabled / document level retry limit is reached,
// documents should go into FailedDocs instead of indexer buffer.

// TODO: use exporterhelper retry_sender

func flushBulkIndexer(
	ctx context.Context,
	bi *docappender.BulkIndexer,
	timeout time.Duration,
	retryOnStatus []int,
	tMetaKeys []string,
	tb *metadata.TelemetryBuilder,
	logger *zap.Logger,
	failedDocsInputLogger *zap.Logger,
	getErrorHintFunc func(index, errorType string) string,
	suppressConflictErrors bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Create context with attempt counter to track http requests

// append metadata attributes to error log fields

// For other types, convert to string

// Record a successful completed bulk request

// Collect telemetry

// Rejection of duplicates are either expected (Profiling indices)
// or globally suppressed by the user. Do not log them.

// Log failed docs

func retryableStatusCode(statusCode int, retryOnStatus []int) bool {
	_ = "STUB: not implemented"
	return false
}

func getAttributesFromMetadataKeys(ctx context.Context, keys []string) []attribute.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

func getErrorHint(mode MappingMode, index, errorType string) string {
	_ = "STUB: not implemented"
	return ""
}

//  no error hint for this mapping mode

func newFailedDocsInputLogger(logger *zap.Logger, config *Config) *zap.Logger {
	_ = "STUB: not implemented"
	return nil
}

type bulkIndexers struct {
	// wg tracks active sessions
	wg sync.WaitGroup

	// NOTE(axw) we have removed async bulk indexer and there should be
	// no reason for having one per mode or for different document types.
	// Instead, the caller can create separate sessions as needed, and we
	// can either have one for required_data_stream=true and one for false,
	// or callers can set this per document.

	modes                [NumMappingModes]bulkIndexer
	profilingEvents      bulkIndexer // For profiling-events-*
	profilingStackTraces bulkIndexer // For profiling-stacktraces
	profilingStackFrames bulkIndexer // For profiling-stackframes
	profilingExecutables bulkIndexer // For profiling-executables

	telemetryBuilder *metadata.TelemetryBuilder
}

func (b *bulkIndexers) start(
	ctx context.Context,
	cfg *Config,
	set exporter.Settings,
	host component.Host,
	allowedMappingModes map[string]MappingMode,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *bulkIndexers) shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

type wgTrackingBulkIndexer struct {
	bulkIndexer
	wg *sync.WaitGroup
}

func (w *wgTrackingBulkIndexer) StartSession(ctx context.Context) bulkIndexerSession {
	_ = "STUB: not implemented"
	return *new(bulkIndexerSession)
}

type wgTrackingBulkIndexerSession struct {
	bulkIndexerSession
	wg *sync.WaitGroup
}

func (w *wgTrackingBulkIndexerSession) End() { _ = "STUB: not implemented"; return }

type errBulkIndexerSession struct {
	err error
}

func (s errBulkIndexerSession) Add(context.Context, string, string, string, io.WriterTo, map[string]string, string) error {
	_ = "STUB: not implemented"
	return nil
}

func (errBulkIndexerSession) End() { _ = "STUB: not implemented"; return }

func (s errBulkIndexerSession) Flush(context.Context) error { _ = "STUB: not implemented"; return nil }

func withOutcome(outcome string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func statusToOutcome(statusCode int) string { _ = "STUB: not implemented"; return "" }
