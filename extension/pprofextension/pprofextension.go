// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pprofextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/pprofextension"

import (
	"context"
	"net/http"
	_ "net/http/pprof" // #nosec Needed to enable the performance profiler
	"os"
	"sync/atomic"

	"go.opentelemetry.io/collector/component"
)

var running = &atomic.Bool{}

type pprofExtension struct {
	config            Config
	file              *os.File
	server            http.Server
	stopCh            chan struct{}
	telemetrySettings component.TelemetrySettings
}

func (p *pprofExtension) Start(_ context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	// The runtime settings are global to the application, so while in principle it
	// is possible to have more than one instance, running multiple will mean that
	// the settings of the last started instance will prevail. In order to avoid
	// this issue we will allow the start of a single instance once per process
	// Summary: only a single instance can be running in the same process.
	return nil
}

// Take care that if any error happen when starting the active instance is cleaned.

// Start the listener here so we can have earlier failure if port is
// already in use.

// The listener ownership goes to the server.

func (p *pprofExtension) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

// ignore the error

func newServer(config Config, params component.TelemetrySettings) *pprofExtension {
	_ = "STUB: not implemented"
	return nil
}
