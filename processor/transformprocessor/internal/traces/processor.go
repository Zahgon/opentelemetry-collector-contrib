// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package traces // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor/internal/traces"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspan"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspanevent"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor/internal/common"
)

type Processor struct {
	contexts []common.TracesConsumer
	logger   *zap.Logger
}

func NewProcessor(contextStatements []common.ContextStatements, errorMode ottl.ErrorMode, settings component.TelemetrySettings, spanFunctions map[string]ottl.Factory[*ottlspan.TransformContext], spanEventFunctions map[string]ottl.Factory[*ottlspanevent.TransformContext]) (*Processor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Processor) ProcessTraces(ctx context.Context, td ptrace.Traces) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}
