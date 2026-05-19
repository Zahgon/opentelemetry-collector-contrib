// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sumologicprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/sumologicprocessor"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

// aggregateAttributesProcessor
type aggregateAttributesProcessor struct {
	aggregations []*aggregation
}

type aggregation struct {
	attribute string
	prefixes  []string
}

func newAggregateAttributesProcessor(config []AggregationPair) *aggregateAttributesProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (proc *aggregateAttributesProcessor) processLogs(logs plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (proc *aggregateAttributesProcessor) processMetrics(metrics pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (proc *aggregateAttributesProcessor) processTraces(traces ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func (proc *aggregateAttributesProcessor) isEnabled() bool { _ = "STUB: not implemented"; return false }

func (*aggregateAttributesProcessor) ConfigPropertyName() string {
	_ = "STUB: not implemented"
	return ""
}

func (proc *aggregateAttributesProcessor) processAttributes(attributes pcommon.Map) error {
	_ = "STUB: not implemented"
	return nil
}

// Create a new map. Unused keys will be added here,
// so we can check them against other prefixes.

// TODO: Potential name conflict to resolve, eg.:
// pod_* matches pod_foo
// pod2_* matches pod2_foo
// both will be renamed to foo
// ref: https://github.com/SumoLogic/sumologic-otel-collector/issues/1263

// Add a new attribute only if there's anything that should be put under it.
