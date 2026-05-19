// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package deltatorateprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatorateprocessor"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

type deltaToRateProcessor struct {
	ConfiguredMetrics map[string]bool
	logger            *zap.Logger
}

func newDeltaToRateProcessor(config *Config, logger *zap.Logger) *deltaToRateProcessor {
	_ = "STUB: not implemented"
	return nil
}

// Start is invoked during service startup.
func (*deltaToRateProcessor) Start(context.Context, component.Host) error {
	_ = "STUB: not implemented"

	// processMetrics implements the ProcessMetricsFunc type.
	return nil
}

func (dtrp *deltaToRateProcessor) processMetrics(_ context.Context, md pmetric.Metrics) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// Setting the data type removed all the data points, so we must move them back to the metric.

// Append "/s" to nonempty units to reflect the per-second rate.
// There are a couple of edge cases that could be considered here:
// 1. The unit could already end with "/s" or a different denominator unit.
// 2. The unit could be a time unit which would result in a scalar.
//
// There are two relevant rules in the UCUM specification that apply here:
//  a. UCUM 2.2. §7.2, https://ucum.org/ucum#section-Syntax-Rules
//    > Terms are evaluated from left to right with the period and the solidus
//    > having the same operator precedence. Multiple division operators are
//    > allowed within one term."
//  b. UCUM 2 §2.2 https://ucum.org/ucum#section-Grammar-of-Units-and-Unit-Terms
//    > Programs that declare full conformance [...] must detect equivalence
//    > for different expressions with the same meaning
//
//  This means that resulting expressions such as "m/s/s" or "ms/s" are valid UCUM
//  units and there is no need to simplify them to be compliant.

// Shutdown is invoked during service shutdown.
func (*deltaToRateProcessor) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

func calculateRate(value float64, durationNanos time.Duration) float64 {
	_ = "STUB: not implemented"
	return 0
}
