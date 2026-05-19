// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package remotetapextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/remotetapextension"

import (
	"context"
	"embed"
	"net/http"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

//go:embed http/*
var httpFS embed.FS

type remoteObserverExtension struct {
	config   *Config
	settings extension.Settings
	server   *http.Server
}

func (s *remoteObserverExtension) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *remoteObserverExtension) Shutdown(_ context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
