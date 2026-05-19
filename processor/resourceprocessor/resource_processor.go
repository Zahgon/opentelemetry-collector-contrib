// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package resourceprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourceprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/pprofile"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/attraction"
)

type resourceProcessor struct {
	logger   *zap.Logger
	attrProc *attraction.AttrProc
}

func (rp *resourceProcessor) processTraces(ctx context.Context, td ptrace.Traces) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

func (rp *resourceProcessor) processMetrics(ctx context.Context, md pmetric.Metrics) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

func (rp *resourceProcessor) processLogs(ctx context.Context, ld plog.Logs) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

func (rp *resourceProcessor) processProfiles(ctx context.Context, pd pprofile.Profiles) (pprofile.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(pprofile.Profiles), nil
}
