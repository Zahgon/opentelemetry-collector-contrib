// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package correlation // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/correlation"

import (
	"context"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/apm/correlations"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/apm/tracetracker"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/timeutils"
)

// Tracker correlation
type Tracker struct {
	once         sync.Once
	log          *zap.Logger
	cfg          *Config
	params       exporter.Settings
	traceTracker *tracetracker.ActiveServiceTracker
	pTicker      timeutils.TTicker
	correlation  *correlationContext
	accessToken  configopaque.String
}

type correlationContext struct {
	correlations.CorrelationClient
	cancel context.CancelFunc
}

// NewTracker creates a new tracker instance for correlation.
func NewTracker(cfg *Config, accessToken configopaque.String, params exporter.Settings) *Tracker {
	_ = "STUB: not implemented"
	return nil
}

func newCorrelationClient(ctx context.Context, cfg *Config, accessToken configopaque.String, params exporter.Settings, host component.Host) (
	*correlationContext, error,
) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ProcessTraces processes the provided spans to correlate the services and environment observed
// to the resources (host, pods, etc.) emitting the spans.
func (cor *Tracker) ProcessTraces(ctx context.Context, traces ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

// Start correlation tracking.
func (cor *Tracker) Start(ctx context.Context, host component.Host) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Shutdown correlation tracking.
func (cor *Tracker) Shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }
