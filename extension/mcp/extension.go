// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package mcp // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/mcp"

import (
	"context"
	"net/http"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

var _ extension.Extension = (*mcpExtension)(nil)

type mcpExtension struct {
	cfg        *Config
	settings   component.TelemetrySettings
	server     *http.Server
	shutdownWG sync.WaitGroup
}

func newExtension(cfg *Config, telemetry component.TelemetrySettings) *mcpExtension {
	_ = "STUB: not implemented"
	return nil
}

func (mcpe *mcpExtension) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (mcpe *mcpExtension) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
