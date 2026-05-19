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
	"go.opentelemetry.io/collector/extension"
)

type factory struct {
	onceProvider   sync.Once
	sourceProvider source.Provider
	providerErr    error
}

func (f *factory) SourceProvider(set component.TelemetrySettings, configHostname string, timeout time.Duration) (source.Provider, error) {
	_ = "STUB: not implemented"
	return *new(source.Provider), nil
}

// NewFactory creates a factory for the Datadog extension.
func NewFactory() extension.Factory { _ = "STUB: not implemented"; return *new(extension.Factory) }

func (*factory) createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func (f *factory) create(ctx context.Context, set extension.Settings, cfg component.Config) (extension.Extension, error) {
	_ = "STUB: not implemented"
	return *new(extension.Extension), nil
}

// set timeout to 25 seconds to avoid default kube liveness probe of 10 seconds * 3 attempts

// Create the real UUID provider for the extension
