// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package solacereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/solacereceiver"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/solacereceiver/internal/metadata"
	receive_v1 "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/solacereceiver/internal/model/receive/v1"
)

type brokerTraceReceiveUnmarshallerV1 struct {
	logger           *zap.Logger
	telemetryBuilder *metadata.TelemetryBuilder
	metricAttrs      attribute.Set // other Otel attributes (to add to the metrics)
}

// unmarshal implements tracesUnmarshaller.unmarshal
func (u *brokerTraceReceiveUnmarshallerV1) unmarshal(message *inboundMessage) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

// unmarshalToSpanData will consume an solaceMessage and unmarshal it into a SpanData.
// Returns an error if one occurred.
func (*brokerTraceReceiveUnmarshallerV1) unmarshalToSpanData(message *inboundMessage) (*receive_v1.SpanData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// populateTraces will create a new Span from the given traces and map the given SpanData to the span.
// This will set all required fields such as name version, trace and span ID, parent span ID (if applicable),
// timestamps, errors and states.
func (u *brokerTraceReceiveUnmarshallerV1) populateTraces(spanData *receive_v1.SpanData, traces ptrace.Traces) {
	_ = "STUB: not implemented"
	// Append new resource span and map any attributes
	return
}

// Create a new span

// map the basic span data

// map all span attributes

// map all events

func (*brokerTraceReceiveUnmarshallerV1) mapResourceSpanAttributes(spanData *receive_v1.SpanData, attrMap pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func (*brokerTraceReceiveUnmarshallerV1) mapClientSpanData(spanData *receive_v1.SpanData, clientSpan ptrace.Span) {
	_ = "STUB: not implemented"
	// Set client span name
	return
}

// SPAN_KIND_CONSUMER == 5

// map trace ID

// map span ID

// conditional parent-span-id

// timestamps

// status

// trace state

// mapAttributes takes a set of attributes from SpanData and maps them to ClientSpan.Attributes().
// Will also copy any user properties stored in the SpanData with a best effort approach.
func (u *brokerTraceReceiveUnmarshallerV1) mapClientSpanAttributes(spanData *receive_v1.SpanData, attrMap pcommon.Map) {
	_ = "STUB: not implemented"
	// receive operation
	return
}

// only message payload
// payload with metadata

// rgmid := u.rgmidToString(spanData.ReplicationGroupMessageId)

// The IPs are now optional meaning we will not include them if they are zero length

// mapEvents maps all events contained in SpanData to relevant events within clientSpan.Events()
func (u *brokerTraceReceiveUnmarshallerV1) mapEvents(spanData *receive_v1.SpanData, clientSpan ptrace.Span) {
	_ = "STUB: not implemented"
	// handle enqueue events
	return
}

// handle transaction events

// mapEnqueueEvent maps a SpanData_EnqueueEvent to a ClientSpan.Event
func (u *brokerTraceReceiveUnmarshallerV1) mapEnqueueEvent(enqueueEvent *receive_v1.SpanData_EnqueueEvent, clientSpanEvents ptrace.SpanEventSlice) {
	_ = "STUB: not implemented"
	return
}

// Final should be `<dest> enqueue`

// mapTransactionEvent maps a SpanData_TransactionEvent to a ClientSpan.Event
func (u *brokerTraceReceiveUnmarshallerV1) mapTransactionEvent(transactionEvent *receive_v1.SpanData_TransactionEvent, clientSpanEvents ptrace.SpanEventSlice) {
	_ = "STUB: not implemented"
	// map the transaction type to a name
	return
}

// Set the name to the unknown transaction event type to ensure forward compat.

// create a new client span event for this event

// map initiator enums to expected initiator strings

// conditionally set the error description if one occurred, otherwise omit

// map the transaction type/id

// format xxxxxxxx-yyyyyyyy-zzzzzzzz where x is FormatID (hex rep of int32), y is BranchQualifier and z is GlobalID, hex encoded.

func (u *brokerTraceReceiveUnmarshallerV1) rgmidToString(rgmid []byte) string {
	_ = "STUB: not implemented"
	// rgmid[0] is the version of the rgmid
	return ""
}

// may be cases where the rgmid is empty or nil, len(rgmid) will return 0 if nil

// format: rmid1:aaaaa-bbbbbbbbbbb-cccccccc-dddddddd

// unmarshalBaggage will unmarshal a baggage string
// See spec https://github.com/open-telemetry/opentelemetry-go/blob/v1.11.1/baggage/baggage.go
func (*brokerTraceReceiveUnmarshallerV1) unmarshalBaggage(toMap pcommon.Map, baggageString string) error {
	_ = "STUB: not implemented"
	return nil
}

// we got a valid baggage string, assume everything else is valid

// member.Properties copies, we should cache

// Re-encode the properties and save them as a parameter

// insertUserProperty will insert a user property value with the given key to an attribute if possible.
// Since AttributeMap only supports int64 integer types, uint64 data may be misrepresented.
func (u *brokerTraceReceiveUnmarshallerV1) insertUserProperty(toMap pcommon.Map, key string, value any) {
	_ = "STUB: not implemented"

	// userPropertiesPrefixAttrKey is the key used to prefix all user properties
	return
}
