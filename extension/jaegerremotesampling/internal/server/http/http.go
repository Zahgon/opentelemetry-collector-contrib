// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package http // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/jaegerremotesampling/internal/server/http"

import (
	"context"
	"errors"
	"net/http"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/jaegerremotesampling/internal/source"
)

var errMissingStrategyStore = errors.New("the strategy store has not been provided")

var _ component.Component = (*SamplingHTTPServer)(nil)

type SamplingHTTPServer struct {
	telemetry     component.TelemetrySettings
	settings      confighttp.ServerConfig
	strategyStore source.Source

	mux        *http.ServeMux
	srv        *http.Server
	shutdownWG *sync.WaitGroup
}

func NewHTTP(telemetry component.TelemetrySettings, settings confighttp.ServerConfig, strategyStore source.Source) (*SamplingHTTPServer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SEE: https://www.jaegertracing.io/docs/1.41/apis/#remote-sampling-configuration-stable

func (h *SamplingHTTPServer) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *SamplingHTTPServer) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *SamplingHTTPServer) samplingStrategyHandler(rw http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
