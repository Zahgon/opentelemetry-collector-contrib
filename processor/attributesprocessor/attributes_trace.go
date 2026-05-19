// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package attributesprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/attributesprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/attraction"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/expr"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspan"
)

type spanAttributesProcessor struct {
	logger   *zap.Logger
	attrProc *attraction.AttrProc
	skipExpr expr.BoolExpr[*ottlspan.TransformContext]
}

// newTracesProcessor returns a processor that modifies attributes of a span.
// To construct the attributes processors, the use of the factory methods are required
// in order to validate the inputs.
func newSpanAttributesProcessor(logger *zap.Logger, attrProc *attraction.AttrProc, skipExpr expr.BoolExpr[*ottlspan.TransformContext]) *spanAttributesProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (a *spanAttributesProcessor) processTraces(ctx context.Context, td ptrace.Traces) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}
