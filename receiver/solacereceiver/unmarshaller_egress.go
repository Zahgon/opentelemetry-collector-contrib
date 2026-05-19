// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package solacereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/solacereceiver"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/solacereceiver/internal/metadata"
	egress_v1 "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/solacereceiver/internal/model/egress/v1"
)

type brokerTraceEgressUnmarshallerV1 struct {
	logger           *zap.Logger
	telemetryBuilder *metadata.TelemetryBuilder
	metricAttrs      attribute.Set // other Otel attributes (to add to the metrics)
}

// unmarshal implements tracesUnmarshaller.unmarshal
func (u *brokerTraceEgressUnmarshallerV1) unmarshal(message *inboundMessage) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

// unmarshalToSpanData will consume an solaceMessage and unmarshal it into a SpanData.
// Returns an error if one occurred.
func (*brokerTraceEgressUnmarshallerV1) unmarshalToSpanData(message *inboundMessage) (*egress_v1.SpanData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// populateTraces will create a new Span from the given traces and map the given SpanData to the span.
// This will set all required fields such as name version, trace and span ID, parent span ID (if applicable),
// timestamps, errors and states.
func (u *brokerTraceEgressUnmarshallerV1) populateTraces(spanData *egress_v1.SpanData, traces ptrace.Traces) {
	_ = "STUB: not implemented"
	// Append new resource span and map any attributes
	return
}

func (*brokerTraceEgressUnmarshallerV1) mapResourceSpanAttributes(spanData *egress_v1.SpanData, attrMap pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func (u *brokerTraceEgressUnmarshallerV1) mapEgressSpan(spanData *egress_v1.SpanData_EgressSpan, clientSpans ptrace.SpanSlice) {
	_ = "STUB: not implemented"
	// at least a support Egress span is found
	return
}

// We only map for KNOWN span types. Current known list:[ SendSpan, DeleteSpan ]
// we drop any other unknown/unsupported span types

// map Egress Send span attributes

// map Egress Delete span attributes

// unknown span type, drop the span

// map any transaction events found

// malformed/incomplete egress span received, drop the span

func (*brokerTraceEgressUnmarshallerV1) mapEgressSpanCommon(spanData *egress_v1.SpanData_EgressSpan, clientSpan ptrace.Span) {
	_ = "STUB: not implemented"
	return
}

// conditional parent-span-id

// timestamps

// status

func (u *brokerTraceEgressUnmarshallerV1) mapSendSpan(sendSpan *egress_v1.SpanData_SendSpan, span ptrace.Span) {
	_ = "STUB: not implemented"
	return
}

// hard coded to producer span

// we don't fatal out when we don't have a valid kind, instead just log and increment stats

// include the partition number, if available

func (u *brokerTraceEgressUnmarshallerV1) mapDeleteSpan(deleteSpan *egress_v1.SpanData_DeleteSpan, span ptrace.Span) {
	_ = "STUB: not implemented"
	return
}

// Delete Info reasons

// hard coded to internal span

// include the partition number, if available

// Don't fatal out when we don't have a valid Endpoint name, instead just log and increment stats

// do not fatal out when we don't have a valid delete reason name
// instead just log and increment stats

// caused by expired ttl on message

// caused by consumer N(ack)ing with Rejected outcome

// caused by max redelivery reached/exceeded

// caused by exceeded hop count

// caused by destination unable to match any ingress selector rule

// caused by admin action

// mapDeleteSpanAdminActionInfo will map the delete admin action information
func (u *brokerTraceEgressUnmarshallerV1) mapDeleteSpanAdminActionInfo(adminActionInfo *egress_v1.SpanData_AdminActionInfo, attrMap pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

// Supported Admin Interface names

// the authenticated userId that performed the delete action

// Do not fatal out when there isn't a valid delete admin action session type, instead just log and increment stats

// from Cli

// get cli local session information

// set the admin interface name as "cli_terminal"

// session number for the cli connection that made the delete request

// get cli remote session information

// set the admin interface name as "cli_ssl"

// the peer IP address

// from SEMP

// set the admin interface name as "semp"

// maps a transaction event. We cannot reuse the code in receive unmarshaller since
// the protobuf model is different and the return type for things like type and initiator would not work in an interface
func (u *brokerTraceEgressUnmarshallerV1) mapTransactionEvent(transactionEvent *egress_v1.SpanData_TransactionEvent, clientEvent ptrace.SpanEvent) {
	_ = "STUB: not implemented"
	// map the transaction type to a name
	return
}

// Set the name to the unknown transaction event type to ensure forward compat.

// map initiator enums to expected initiator strings

// conditionally set the error description if one occurred, otherwise omit

// map the transaction type/id

// format xxxxxxxx-yyyyyyyy-zzzzzzzz where x is FormatID (hex rep of int32), y is BranchQualifier and z is GlobalID, hex encoded.

func isAnonymousQueue(name string) bool {
	_ = "STUB: not implemented"
	// all anonymous queues start with the prefix #P2P/QTMP
	return false
}

func isAnonymousTopicEndpoint(name string) bool {
	_ = "STUB: not implemented"
	// all anonymous topic endpoints are made up of hex strings of length 32
	return false
}

// []byte casting is more efficient in this loop
// check if we are outside 0-9 AND outside a-f
