// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package attributesprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/attributesprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/attraction"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/expr"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlmetric"
)

type metricAttributesProcessor struct {
	logger   *zap.Logger
	attrProc *attraction.AttrProc
	skipExpr expr.BoolExpr[*ottlmetric.TransformContext]
}

// newMetricAttributesProcessor returns a processor that modifies attributes of a
// metric record. To construct the attributes processors, the use of the factory
// methods are required in order to validate the inputs.
func newMetricAttributesProcessor(logger *zap.Logger, attrProc *attraction.AttrProc, skipExpr expr.BoolExpr[*ottlmetric.TransformContext]) *metricAttributesProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (a *metricAttributesProcessor) processMetrics(ctx context.Context, md pmetric.Metrics) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// Attributes are provided for each log and trace, but not at the metric level
// Need to process attributes for every data point within a metric.
func (a *metricAttributesProcessor) processMetricAttributes(ctx context.Context, m pmetric.Metric) {
	_ = "STUB: not implemented"
	// This is a lot of repeated code, but since there is no single parent superclass
	// between metric data types, we can't use polymorphism.
	//exhaustive:enforce
	return
}
