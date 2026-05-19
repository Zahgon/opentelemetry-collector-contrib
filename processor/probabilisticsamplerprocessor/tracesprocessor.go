// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package probabilisticsamplerprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/probabilisticsamplerprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/sampling"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/probabilisticsamplerprocessor/internal/metadata"
)

// samplingPriority has the semantic result of parsing the "sampling.priority"
// attribute per OpenTracing semantic conventions.
type samplingPriority int

const (
	// deferDecision means that the decision if a span will be "sampled" (ie.:
	// forwarded by the collector) is made by hashing the trace ID according
	// to the configured sampling rate.
	deferDecision samplingPriority = iota
	// mustSampleSpan indicates that the span had a "sampling.priority" attribute
	// greater than zero and it is going to be sampled, ie.: forwarded by the
	// collector.
	mustSampleSpan
	// doNotSampleSpan indicates that the span had a "sampling.priority" attribute
	// equal zero and it is NOT going to be sampled, ie.: it won't be forwarded
	// by the collector.
	doNotSampleSpan
)

type traceProcessor struct {
	sampler          dataSampler
	failClosed       bool
	logger           *zap.Logger
	telemetryBuilder *metadata.TelemetryBuilder
}

// tracestateCarrier conveys information about sampled spans between
// the call to parse incoming randomness/threshold and the call to
// decide.
type tracestateCarrier struct {
	span ptrace.Span
	sampling.W3CTraceState
}

var _ samplingCarrier = &tracestateCarrier{}

func newTracestateCarrier(s ptrace.Span) (samplingCarrier, error) {
	_ = "STUB: not implemented"
	return *new(samplingCarrier), nil
}

func (tc *tracestateCarrier) threshold() (sampling.Threshold, bool) {
	_ = "STUB: not implemented"
	return *new(sampling.Threshold), false
}

func (tc *tracestateCarrier) explicitRandomness() (randomnessNamer, bool) {
	_ = "STUB: not implemented"
	return *new(randomnessNamer), false
}

func (tc *tracestateCarrier) updateThreshold(th sampling.Threshold) error {
	_ = "STUB: not implemented"
	return nil
}

func (tc *tracestateCarrier) setExplicitRandomness(rnd randomnessNamer) {
	_ = "STUB: not implemented"
	return
}

func (tc *tracestateCarrier) clearThreshold() { _ = "STUB: not implemented"; return }

func (tc *tracestateCarrier) reserialize() error { _ = "STUB: not implemented"; return nil }

// newTracesProcessor returns a processor.TracesProcessor that will
// perform intermediate span sampling according to the given
// configuration.
func newTracesProcessor(ctx context.Context, set processor.Settings, cfg *Config, nextConsumer consumer.Traces) (processor.Traces, error) {
	_ = "STUB: not implemented"
	return *new(processor.Traces), nil
}

func (th *hashingSampler) randomnessFromSpan(s ptrace.Span) (randomnessNamer, samplingCarrier, error) {
	_ = "STUB: not implemented"
	return *new(randomnessNamer), *new(samplingCarrier), nil
}

// If the tracestate contains a proper R-value or T-value, we
// have to leave it alone.  The user should not be using this
// sampler mode if they are using specified forms of consistent
// sampling in OTel.

// When no sampling information is present, add a
// Randomness value.

func (*consistentTracestateCommon) randomnessFromSpan(s ptrace.Span) (randomnessNamer, samplingCarrier, error) {
	_ = "STUB: not implemented"
	return *new(randomnessNamer), *new(samplingCarrier), nil
}

// When the tracestate is OK and has r-value, use it.

func (*neverSampler) randomnessFromSpan(span ptrace.Span) (randomnessNamer, samplingCarrier, error) {
	_ = "STUB: not implemented"
	// We return a fake randomness value, since it will not be used.
	// This avoids a consistency check error for missing randomness.
	return *new(randomnessNamer), *new(samplingCarrier), nil
}

func (tp *traceProcessor) processTraces(ctx context.Context, td ptrace.Traces) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

// Filter out empty ScopeMetrics

// Filter out empty ResourceMetrics

func (*traceProcessor) priorityFunc(s ptrace.Span, rnd randomnessNamer, threshold sampling.Threshold) (randomnessNamer, sampling.Threshold) {
	_ = "STUB: not implemented"
	return *new(randomnessNamer), *new(sampling.Threshold)
}

// OpenTracing mentions this as a "hint". We take a stronger
// approach and do not sample the span since some may use it to
// remove specific spans from traces.

// override policy name

// override policy name

// Note that the logs processor has very different logic here,
// but that in tracing the priority can only force to never or
// always.

// parseSpanSamplingPriority checks if the span has the "sampling.priority" tag to
// decide if the span should be sampled or not. The usage of the tag follows the
// OpenTracing semantic tags:
// https://github.com/opentracing/specification/blob/main/semantic_conventions.md#span-tags-table
func parseSpanSamplingPriority(span ptrace.Span) samplingPriority {
	_ = "STUB: not implemented"
	return *new(samplingPriority)
}

// By default defer the decision.

// Try check for different types since there are various client libraries
// using different conventions regarding "sampling.priority". Besides the
// client libraries it is also possible that the type was lost in translation
// between different formats.
