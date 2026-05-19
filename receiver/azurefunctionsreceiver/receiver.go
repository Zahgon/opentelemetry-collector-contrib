// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azurefunctionsreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azurefunctionsreceiver"

import (
	"context"
	"net/http"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

type functionsReceiver struct {
	cfg      *Config
	settings receiver.Settings
	nextLogs consumer.Logs

	server     *http.Server
	shutdownWG sync.WaitGroup
}

func newFunctionsReceiver(cfg *Config, settings receiver.Settings, nextLogs consumer.Logs) receiver.Logs {
	_ = "STUB: not implemented"
	return *new(receiver.Logs)
}

func (r *functionsReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// registerTriggerRoutes attaches HTTP handlers for each configured trigger
func (r *functionsReceiver) registerTriggerRoutes(mux *http.ServeMux, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *functionsReceiver) registerEventHubRoutes(mux *http.ServeMux, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *functionsReceiver) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
