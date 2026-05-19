// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package httpserver // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/datadogextension/internal/httpserver"

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/DataDog/datadog-agent/pkg/serializer/marshaler"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/datadogextension/internal/payload"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/datadog/agentcomponents"
)

var nowFunc = time.Now

// Server provides local metadata server functionality (display the otel_collector payload locally) as well as
// functions to serialize and export this metadata to Datadog backend.
type Server struct {
	// serverConfig is used to create the HTTP server via confighttp
	serverConfig *confighttp.ServerConfig
	// handler is the HTTP handler for the server
	handler http.Handler
	// listenClose is used to shut down the server
	listenClose func(ctx context.Context) error
	// logger is passed from the extension to allow logging
	logger *zap.Logger
	// telemetrySettings holds the telemetry settings for the server
	telemetrySettings component.TelemetrySettings
	// serializer is a datadog-agent component used to forward payloads to Datadog backend
	serializer agentcomponents.SerializerWithForwarder
	// config contains the httpserver configuration values
	config *Config

	// payload is the metadata to send to Datadog backend
	payload marshaler.JSONMarshaler

	// mu protects concurrent access to the serializer
	// Note: Only protects serializer operations, not field reads since they're set once during initialization
	mu sync.Mutex
}

// NewServer creates a new HTTP server instance.
// It should be called after NotifyConfig has received full configuration.
// TODO: support generic payloads
func NewServer(
	logger *zap.Logger,
	s agentcomponents.SerializerWithForwarder,
	config *Config,
	hostname string,
	uuid string,
	p payload.OtelCollector,
	telemetrySettings component.TelemetrySettings,
) *Server {
	_ = "STUB: not implemented"
	// Create payload but don't add timestamp, that will happen in SendPayload
	return nil
}

// store as interface

// Start starts the HTTP server and begins sending payloads periodically.
func (s *Server) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Start HTTP server

// Stop shuts down the HTTP server, pass a context to allow for cancellation.
func (s *Server) Stop(ctx context.Context) { _ = "STUB: not implemented"; return }

// Ensure channel is always closed

// SendPayload prepares and sends the fleet automation payloads using Server's handlerDeps
// TODO: support generic payloads
func (s *Server) SendPayload() (marshaler.JSONMarshaler, error) {
	_ = "STUB: not implemented"
	// Use datadog-agent serializer to send these payloads
	return *new(marshaler.JSONMarshaler), nil
}

// Clone the payload to avoid data races

// shallow copy is sufficient since fields are value types or slices (which are not mutated)

// HandleMetadata writes the metadata payloads to the response writer and sends them to the Datadog backend
func (s *Server) HandleMetadata(w http.ResponseWriter, _ *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Marshal the combined payload to JSON
// Note: fullPayload is already thread-safe since SendPayload returned a marshaler interface

// Write the JSON response
