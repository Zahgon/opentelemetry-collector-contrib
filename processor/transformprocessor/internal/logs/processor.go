// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package logs // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor/internal/logs"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottllog"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor/internal/common"
)

type Processor struct {
	contexts []common.LogsConsumer
	logger   *zap.Logger
	flatMode bool
}

func NewProcessor(contextStatements []common.ContextStatements, errorMode ottl.ErrorMode, flatMode bool, settings component.TelemetrySettings, logFunctions map[string]ottl.Factory[*ottllog.TransformContext]) (*Processor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Processor) ProcessLogs(ctx context.Context, ld plog.Logs) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}
