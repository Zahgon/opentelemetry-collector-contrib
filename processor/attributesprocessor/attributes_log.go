// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package attributesprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/attributesprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/attraction"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/expr"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottllog"
)

type logAttributesProcessor struct {
	logger   *zap.Logger
	attrProc *attraction.AttrProc
	skipExpr expr.BoolExpr[*ottllog.TransformContext]
}

// newLogAttributesProcessor returns a processor that modifies attributes of a
// log record. To construct the attributes processors, the use of the factory
// methods are required in order to validate the inputs.
func newLogAttributesProcessor(logger *zap.Logger, attrProc *attraction.AttrProc, skipExpr expr.BoolExpr[*ottllog.TransformContext]) *logAttributesProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (a *logAttributesProcessor) processLogs(ctx context.Context, ld plog.Logs) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}
