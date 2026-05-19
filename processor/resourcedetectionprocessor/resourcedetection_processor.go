// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package resourcedetectionprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/pprofile"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
)

type resourceDetectionProcessor struct {
	provider           *internal.ResourceProvider
	override           bool
	httpClientSettings confighttp.ClientConfig
	refreshInterval    time.Duration
	telemetrySettings  component.TelemetrySettings
}

// Start is invoked during service startup.
func (rdp *resourceDetectionProcessor) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Perform initial resource detection

// Start periodic refresh if configured

// Shutdown is invoked during service shutdown.
func (rdp *resourceDetectionProcessor) Shutdown(_ context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// processTraces implements the ProcessTracesFunc type.
func (rdp *resourceDetectionProcessor) processTraces(ctx context.Context, td ptrace.Traces) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

// processMetrics implements the ProcessMetricsFunc type.
func (rdp *resourceDetectionProcessor) processMetrics(ctx context.Context, md pmetric.Metrics) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// processLogs implements the ProcessLogsFunc type.
func (rdp *resourceDetectionProcessor) processLogs(ctx context.Context, ld plog.Logs) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

// processProfiles implements the ProcessProfilesFunc type.
func (rdp *resourceDetectionProcessor) processProfiles(ctx context.Context, ld pprofile.Profiles) (pprofile.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(pprofile.Profiles), nil
}
