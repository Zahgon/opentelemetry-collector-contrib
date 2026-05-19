// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sentryexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sentryexporter"

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
)

var (
	errExporterShuttingDown = errors.New("sentry exporter is shutting down")
	errProjectBeingCreated  = errors.New("project is being created asynchronously, some events might be dropped")
)

const (
	maxProjects              = 1000 // only prevents auto-creation when total number of seen projects (existing + created) reaches this number
	projectCreationQueueSize = 1000
	timeout                  = 5 * time.Second
)

// newHTTPClient is used to override the http client factory. While running tests
// we need to mock the http transport, so that we don't open real sockets.
var newHTTPClient = func(ctx context.Context, cfg confighttp.ClientConfig, host component.Host, ts component.TelemetrySettings) (*http.Client, error) {
	var extensions map[component.ID]component.Component
	if host != nil {
		extensions = host.GetExtensions()
	}
	return cfg.ToClient(ctx, extensions, ts)
}

// projectCreationRequest represents a request to create a project
type projectCreationRequest struct {
	projectSlug string
	platform    string
}

// endpointState holds shared mutable state (client, caches) across signal exporters.
type endpointState struct {
	config            *Config
	telemetrySettings component.TelemetrySettings

	client       *http.Client
	sentryClient sentryAPIClient

	projectToEndpoint map[string]*otlpEndpoints
	projectMapMu      sync.RWMutex

	attributeKey    string
	projectMapping  map[string]string
	defaultTeamSlug string

	inflight singleflight.Group

	projectCreationQueue chan projectCreationRequest
	pendingCreations     sync.Map
	workerCtx            context.Context
	workerCancel         context.CancelFunc
	workerWg             sync.WaitGroup

	rateLimiter *rateLimiter

	startOnce    sync.Once
	startErr     error
	shutdownOnce sync.Once
	closing      atomic.Bool
}

// projectRoutingKey is a composite key for routing traces and logs by project slug and platform.
type projectRoutingKey struct {
	slug     string
	platform string
}

// sentryExporter is a per-signal wrapper that forwards to the shared endpointState.
type sentryExporter struct {
	logger *zap.Logger
	state  *endpointState
}

func newEndpointState(config *Config, set exporter.Settings) (*endpointState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newSentryExporter(state *endpointState, logger *zap.Logger) *sentryExporter {
	_ = "STUB: not implemented"
	return nil
}

// pushTraceData takes an incoming OpenTelemetry trace, and forwards it to the Sentry OTLP endpoint.
func (e *sentryExporter) pushTraceData(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *endpointState) pushTraceData(ctx context.Context, logger *zap.Logger, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

// routeTracesByProject splits traces by project and sends each batch to the appropriate endpoint
func (s *endpointState) routeTracesByProject(ctx context.Context, logger *zap.Logger, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

// sendTracesToEndpoint sends traces to a specific endpoint
func (s *endpointState) sendTracesToEndpoint(ctx context.Context, logger *zap.Logger, td ptrace.Traces, endpoint *otlpEndpoints) error {
	_ = "STUB: not implemented"
	return nil
}

// projectCreationWorker runs in the background to handle async project creation
func (s *endpointState) projectCreationWorker() { _ = "STUB: not implemented"; return }

// Start starts the exporter
func (s *endpointState) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *endpointState) preloadEndpointsFromOrgKeys(ctx context.Context, projects []projectInfo) {
	_ = "STUB: not implemented"
	return
}

// Shutdown stops the exporter
func (s *endpointState) Shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (e *sentryExporter) pushLogData(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *endpointState) pushLogData(ctx context.Context, logger *zap.Logger, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *endpointState) routeLogsByProject(ctx context.Context, logger *zap.Logger, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *endpointState) routeByProject(ctx context.Context, logger *zap.Logger, projectGroups map[projectRoutingKey]any, send func(context.Context, *zap.Logger, any, *otlpEndpoints) error) error {
	_ = "STUB: not implemented"
	return nil
}

func countSpansInResource(rs ptrace.ResourceSpans) int { _ = "STUB: not implemented"; return 0 }

func countLogsInResource(rl plog.ResourceLogs) int { _ = "STUB: not implemented"; return 0 }

func (s *endpointState) sendLogsToEndpoint(ctx context.Context, logger *zap.Logger, ld plog.Logs, endpoint *otlpEndpoints) error {
	_ = "STUB: not implemented"
	return nil
}

// sentryHTTPError represents an HTTP error from Sentry
type sentryHTTPError struct {
	statusCode int
	body       string
	headers    http.Header
}

func (e *sentryHTTPError) Error() string { _ = "STUB: not implemented"; return "" }

func (s *endpointState) sendOTLPData(ctx context.Context, logger *zap.Logger, dsn string, category dataCategory, endpoint string, data []byte, authHeader string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *endpointState) extractProjectSlug(attrs pcommon.Map) string {
	_ = "STUB: not implemented"
	return ""
}

func (*endpointState) extractPlatform(_ pcommon.Map) string {
	_ = "STUB: not implemented"
	// for correctly mapping `telemetry.sdk.language` to the Sentry platform we need to keep track of all supported platforms.
	// defaulting to `other` for now to avoid validation issues.
	return ""
}

func (s *endpointState) getOrCreateProjectEndpoint(ctx context.Context, logger *zap.Logger, projectSlug, platform string) (*otlpEndpoints, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
