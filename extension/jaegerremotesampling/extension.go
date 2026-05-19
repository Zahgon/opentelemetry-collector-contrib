// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package jaegerremotesampling // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/jaegerremotesampling"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/jaegerremotesampling/internal/source"
)

var _ extension.Extension = (*jrsExtension)(nil)

type jrsExtension struct {
	cfg       *Config
	telemetry component.TelemetrySettings

	httpServer    component.Component
	grpcServer    component.Component
	samplingStore source.Source

	closers []func() error
}

func newExtension(cfg *Config, telemetry component.TelemetrySettings) *jrsExtension {
	_ = "STUB: not implemented"
	return nil
}

func (jrse *jrsExtension) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	// the config validation will take care of ensuring we have one and only one of the following about the
	// source of the sampling config:
	// - remote (gRPC)
	// - local file
	// we can then use a simplified logic here to assign the appropriate store
	return nil
}

// there's a Close function on the concrete type, which is not visible to us...
// how can we close it then?

// then we start our own server interfaces, starting with the HTTP one

// start our gRPC server interface

func (jrse *jrsExtension) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	// we probably don't want to break whenever an error occurs, we want to continue and close the other resources
	return nil
}
