// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metricsgenerationprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricsgenerationprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

type metricsGenerationProcessor struct {
	rules  []internalRule
	logger *zap.Logger
}

type internalRule struct {
	name      string
	unit      string
	ruleType  string
	metric1   string
	metric2   string
	operation string
	scaleBy   float64
}

func newMetricsGenerationProcessor(rules []internalRule, logger *zap.Logger) *metricsGenerationProcessor {
	_ = "STUB: not implemented"
	return nil
}

// Start is invoked during service startup.
func (*metricsGenerationProcessor) Start(context.Context, component.Host) error {
	_ = "STUB: not implemented"

	// processMetrics implements the ProcessMetricsFunc type.
	return nil
}

func (mgp *metricsGenerationProcessor) processMetrics(_ context.Context, md pmetric.Metrics) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// Operation type is validated during config validation, but this adds extra validation as a safety net

// When matching metric attributes isn't required the value of the first data point of metric2 is
// used for all calculations. The resulting logic is the same as generating a new metric from
// a scalar.

// Shutdown is invoked during service shutdown.
func (*metricsGenerationProcessor) Shutdown(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
