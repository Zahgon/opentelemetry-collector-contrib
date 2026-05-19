// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !aix

package datadogextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/datadogextension"

import (
	"context"
	"sync"
	"time"

	"github.com/DataDog/datadog-agent/pkg/opentelemetry-mapping-go/otlp/attributes/source"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componentstatus"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/extension/extensioncapabilities"
	"go.opentelemetry.io/collector/service"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/datadogextension/internal/httpserver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/datadogextension/internal/payload"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/datadog/agentcomponents"
)

const (
	payloadSendingInterval = 5 * time.Minute
	payloadTTL             = payloadSendingInterval * 3
)

// uuidProvider defines an interface for generating UUIDs, allowing for mocking in tests.
type uuidProvider interface {
	NewString() string
}

// realUUIDProvider is the concrete implementation of uuidProvider that uses the google/uuid package.
type realUUIDProvider struct{}

// NewString returns a new UUID string using the real uuid package.
func (*realUUIDProvider) NewString() string { _ = "STUB: not implemented"; return "" }

type configs struct {
	collector *confmap.Conf
	extension *Config
	mutex     sync.RWMutex
}

type info struct {
	host               source.Source
	hostnameSource     string
	uuid               string
	build              component.BuildInfo
	modules            service.ModuleInfos
	resourceAttributes map[string]string
}

type payloadSender struct {
	ticker  *time.Ticker
	ctx     context.Context
	cancel  context.CancelFunc
	channel chan struct{}
	once    sync.Once
}

type datadogExtension struct {
	extension.Extension // Embed base Extension for common functionality.

	// struct to store extension and collector configs
	configs *configs

	// components to assist in payload sending/logging
	logger            *zap.Logger
	serializer        agentcomponents.SerializerWithForwarder
	telemetrySettings component.TelemetrySettings

	// struct to store extension info
	info *info

	// host stores the component host for use when creating the HTTP server
	host component.Host

	otelCollectorMetadata *payload.OtelCollector
	httpServer            *httpserver.Server

	// Fields for periodic payload sending
	payloadSender *payloadSender
}

var (
	_ extensioncapabilities.ConfigWatcher   = (*datadogExtension)(nil)
	_ extensioncapabilities.PipelineWatcher = (*datadogExtension)(nil)
	_ componentstatus.Watcher               = (*datadogExtension)(nil)
)

// NotifyConfig implements the extensioncapabilities.ConfigWatcher interface, which allows
// this extension to be notified of the Collector's effective configuration.
// This method is called during startup by the Collector's service after calling Start.
func (e *datadogExtension) NotifyConfig(_ context.Context, conf *confmap.Conf) error {
	_ = "STUB: not implemented"
	return nil
}

// Create the build info struct for the payload

// Get the full collector configuration as a flattened JSON string

// Prepare the base payload

// This is the same version from buildInfo; it is possible we could want to set a different version here in the future

// Populate resource attributes collected from TelemetrySettings.Resource

// Populate the full list of components available in the collector build

// Populate the list of components that are active in a pipeline

// TODO: Populate HealthStatus from the pkg/status.
// https://datadoghq.atlassian.net/browse/OTEL-2663
// For now, we leave it as an empty string.

// Store the created payload in the extension struct

// Create and start the HTTP server

// Start periodic payload sending (every 30 minutes)

// Ready implements the extensioncapabilities.PipelineWatcher interface.
func (*datadogExtension) Ready() error {
	_ = "STUB: not implemented"

	// NotReady implements the extensioncapabilities.PipelineWatcher interface.
	return nil
}

func (*datadogExtension) NotReady() error {
	_ = "STUB: not implemented"

	// ComponentStatusChanged implements the componentstatus.Watcher interface.
	return nil
}

func (*datadogExtension) ComponentStatusChanged(
	*componentstatus.InstanceID,
	*componentstatus.Event,
) {
	_ = "STUB: not implemented"

	// Start starts the extension via the component interface.
	return
}

func (e *datadogExtension) Start(_ context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	// Store host for later use when creating the HTTP server
	return nil
}

// Retrieve module information from the host

// Start the serializer if it's available

// Shutdown stops the extension via the component interface.
// It shuts down the HTTP server, stops forwarder, and passes signal on
// channel to end goroutine that sends the Datadog otel_collector payloads.
func (e *datadogExtension) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Stop periodic payload sending

// Stop the serializer if it's available

// sendLivenessMetric sends the otel.datadog_extension.running metric to indicate the extension is active
func (e *datadogExtension) sendLivenessMetric(_ context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Add hostname source tag to indicate whether hostname was configured or inferred

// Create the liveness metric serie directly in agent format

// Send using the serializer (forwarder handles retries internally)

// startPeriodicPayloadSending starts a goroutine that sends payloads on payloadSendingInterval
func (e *datadogExtension) startPeriodicPayloadSending() { _ = "STUB: not implemented"; return }

// Send initial liveness metric on startup

// Send liveness metric first

// Wait 1 minute before sending metadata payload

// Send metadata payload after delay

// Allow manual triggering of payload sending if needed
// Manual triggers should send immediately without delay

// Send liveness metric first

// Send metadata payload immediately on manual trigger

// stopPeriodicPayloadSending stops the periodic payload sending goroutine
func (e *datadogExtension) stopPeriodicPayloadSending() { _ = "STUB: not implemented"; return }

// GetSerializer returns the configured serializer with proxy settings from ClientConfig.
// This allows other components (like exporters) to use the same serializer instance
// with the same proxy configuration.
func (e *datadogExtension) GetSerializer() agentcomponents.SerializerWithForwarder {
	_ = "STUB: not implemented"
	return *new(agentcomponents.SerializerWithForwarder)
}

func newExtension(
	ctx context.Context,
	cfg *Config,
	set extension.Settings,
	hostProvider source.Provider,
	uuidProvider uuidProvider,
) (*datadogExtension, error) {
	_ = "STUB: not implemented"
	// Create configuration for agent components
	// Convert datadogextension.Config to datadogconfig.Config
	return nil, nil
}

// Create agent components with proxy configuration from ClientConfig

// Use ClientConfig proxy settings instead of environment variables

// logging_frequency required to be set to avoid "divide by zero" error

// Create agent components

// Collect resource attributes from TelemetrySettings.Resource
// Format: map[string]string

// Ensure os.type is always present; defer to any value already set by a resource detector

// configure payloadSender struct

// moduleInfos will be populated in Start()
