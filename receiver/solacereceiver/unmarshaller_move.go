// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package solacereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/solacereceiver"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/solacereceiver/internal/metadata"
	move_v1 "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/solacereceiver/internal/model/move/v1"
)

type brokerTraceMoveUnmarshallerV1 struct {
	logger           *zap.Logger
	telemetryBuilder *metadata.TelemetryBuilder
	metricAttrs      attribute.Set // other Otel attributes (to add to the metrics)
}

// unmarshal implements tracesUnmarshaller.unmarshal
func (u *brokerTraceMoveUnmarshallerV1) unmarshal(message *inboundMessage) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

// unmarshalToSpanData will consume an solaceMessage and unmarshal it into a SpanData.
// Returns an error if one occurred.
func (*brokerTraceMoveUnmarshallerV1) unmarshalToSpanData(message *inboundMessage) (*move_v1.SpanData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// populateTraces will create a new Span from the given traces and map the given SpanData to the span.
// This will set all required fields such as name version, trace and span ID, parent span ID (if applicable),
// timestamps, errors and states.
func (u *brokerTraceMoveUnmarshallerV1) populateTraces(spanData *move_v1.SpanData, traces ptrace.Traces) {
	_ = "STUB: not implemented"
	// Append new resource span and map any attributes
	return
}

// Create a new span

// map the tracing data for the span

// map the basic span data

func (*brokerTraceMoveUnmarshallerV1) mapResourceSpanAttributes(spanData *move_v1.SpanData, attrMap pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func (*brokerTraceMoveUnmarshallerV1) mapMoveSpanTracingInfo(spanData *move_v1.SpanData, span ptrace.Span) {
	_ = "STUB: not implemented"
	// hard coded to internal span
	// SPAN_KIND_CONSUMER == 1
	return
}

// map trace ID

// map span ID

// conditional parent-span-id

// timestamps

func (u *brokerTraceMoveUnmarshallerV1) mapClientSpanData(moveSpan *move_v1.SpanData, span ptrace.Span) {
	_ = "STUB: not implemented"
	return
}

// Delete Info reasons

// map the replication group ID for the move span

// source queue partition number

// destination queue partition number

// set source endpoint information
// don't fatal out when we receive invalid endpoint name, instead just log and increment stats

// set destination endpoint information
// don't fatal out when we receive invalid endpoint name, instead just log and increment stats

// do not fatal out when we don't have a valid move reason name
// instead just log and increment stats

// caused by expired ttl on message

// caused by consumer N(ack)ing with Rejected outcome

// caused by max redelivery reached/exceeded
