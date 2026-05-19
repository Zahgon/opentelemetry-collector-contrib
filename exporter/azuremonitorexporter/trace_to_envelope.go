// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azuremonitorexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azuremonitorexporter"

// Contains code common to both trace and metrics exporters

import (
	"errors"
	"strings"

	"github.com/microsoft/ApplicationInsights-Go/appinsights/contracts"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

const (
	unknownSpanType   spanType = 0
	httpSpanType      spanType = 1
	rpcSpanType       spanType = 2
	databaseSpanType  spanType = 3
	messagingSpanType spanType = 4
	faasSpanType      spanType = 5

	exceptionSpanEventName string = "exception"
	msLinks                string = "_MS.links"
)

var (
	errUnexpectedAttributeValueType = errors.New("attribute value type is unexpected")
	errUnsupportedSpanType          = errors.New("unsupported Span type")
)

// Used to identify the type of a received Span
type spanType int8

type msLink struct {
	OperationID string `json:"operation_Id"`
	ID          string `json:"id"`
}

// Transforms a tuple of pcommon.Resource, pcommon.InstrumentationScope, ptrace.Span into one or more of AppInsights contracts.Envelope
// This is the only method that should be targeted in the unit tests
func spanToEnvelopes(
	resource pcommon.Resource,
	instrumentationScope pcommon.InstrumentationScope,
	span ptrace.Span,
	spanEventsEnabled bool,
	logger *zap.Logger,
) ([]*contracts.Envelope, error) {
	_ = "STUB: not implemented"
	return nil,

		// According to the SpanKind documentation, we can assume it to be INTERNAL
		// when we get UNSPECIFIED.
		nil
}

// For now, FaaS spans are unsupported

// First map the span itself

// Regardless of the detected Span type, if the SpanKind is Internal we need to set data.Type to InProc

// Record the raw Span status values as properties

// Sanitize the base data, the envelope and envelope tags

// Now add the span events. We always export exception events.

// skip non-exception events if configured

// Exceptions are a special case of span event.
// See https://opentelemetry.io/docs/reference/specification/trace/semantic_conventions/exceptions/#recording-an-exception

// Sanitize the base data, the envelope and envelope tags

func applyLinksToDataProperties(dataProperties map[string]string, spanLinkSlice ptrace.SpanLinkSlice, logger *zap.Logger) {
	_ = "STUB: not implemented"
	return
}

// Creates a new envelope with some basic tags populated
func newEnvelope(span ptrace.Span, time string) *contracts.Envelope {
	_ = "STUB: not implemented"
	return nil
}

// Maps Server/Consumer Span to AppInsights RequestData
func spanToRequestData(span ptrace.Span, incomingSpanType spanType) *contracts.RequestData {
	_ = "STUB: not implemented"
	// See https://github.com/microsoft/ApplicationInsights-Go/blob/master/appinsights/contracts/requestdata.go
	// Start with some reasonable default for server spans.
	return nil
}

// Maps Span to AppInsights RemoteDependencyData
func spanToRemoteDependencyData(span ptrace.Span, incomingSpanType spanType) *contracts.RemoteDependencyData {
	_ = "STUB: not implemented"
	// https://github.com/microsoft/ApplicationInsights-Go/blob/master/appinsights/contracts/remotedependencydata.go
	// Start with some reasonable default for dependent spans.
	return nil
}

// Maps SpanEvent to AppInsights ExceptionData
func spanEventToExceptionData(spanEvent ptrace.SpanEvent) *contracts.ExceptionData {
	_ = "STUB: not implemented"
	return nil
}

// Maps SpanEvent to AppInsights MessageData
func spanEventToMessageData(spanEvent ptrace.SpanEvent) *contracts.MessageData {
	_ = "STUB: not implemented"
	return nil
}

func getFormattedHTTPStatusValues(statusCode int64) (statusAsString string, success bool) {
	_ = "STUB: not implemented"
	// see https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/trace/semantic_conventions/http.md#status
	return "", false
}

// Maps HTTP Server Span to AppInsights RequestData
// https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/trace/semantic_conventions/http.md#semantic-conventions-for-http-spans
func fillRequestDataHTTP(span ptrace.Span, data *contracts.RequestData) {
	_ = "STUB: not implemented"
	return
}

// Construct data.Name
// The data.Name should be {HTTP METHOD} {HTTP SERVER ROUTE TEMPLATE}
// https://github.com/microsoft/ApplicationInsights-Home/blob/f1f9f619d74557c8db3dbde4b49c4193e10d8a81/EndpointSpecs/Schemas/Bond/RequestData.bond#L32

// Use httpRoute if available otherwise fallback to the span name
// https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/trace/semantic_conventions/http.md#name

/*
	To construct the value for data.Url we will use the following sets of attributes as defined by the otel spec
	Order of preference is:
	http.scheme, http.host, http.target
	http.scheme, http.server_name, net.host.port, http.target
	http.scheme, net.host.name, net.host.port, http.target
	http.url
*/

// Maps HTTP Client Span to AppInsights RemoteDependencyData
// https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/trace/semantic_conventions/http.md
func fillRemoteDependencyDataHTTP(span ptrace.Span, data *contracts.RemoteDependencyData) {
	_ = "STUB: not implemented"
	return
}

// Maps RPC Server Span to AppInsights RequestData
// https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/trace/semantic_conventions/rpc.md
func fillRequestDataRPC(span ptrace.Span, data *contracts.RequestData) {
	_ = "STUB: not implemented"
	return
}

// Prefix the name with the type of RPC

// Set the .Data property to .Name which contain the full RPC method

// Maps RPC Client Span to AppInsights RemoteDependencyData
// https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/trace/semantic_conventions/rpc.md
func fillRemoteDependencyDataRPC(span ptrace.Span, data *contracts.RemoteDependencyData) {
	_ = "STUB: not implemented"
	return
}

// Set the .Data property to .Name which contain the full RPC method

// Returns the RPC status code as a string
func getRPCStatusCodeAsString(rpcAttributes *rpcAttributes) (statusCodeAsString string) {
	_ = "STUB: not implemented"
	// Honor the attribute rpc.grpc.status_code if there
	return ""
}

// Maps Database Client Span to AppInsights RemoteDependencyData
// https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/trace/semantic_conventions/database.md
func fillRemoteDependencyDataDatabase(span ptrace.Span, data *contracts.RemoteDependencyData) {
	_ = "STUB: not implemented"
	return
}

// Maps Messaging Consumer/Server Span to AppInsights RequestData
// https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/trace/semantic_conventions/messaging.md
func fillRequestDataMessaging(span ptrace.Span, data *contracts.RequestData) {
	_ = "STUB: not implemented"
	return
}

// TODO Understand how to map attributes to RequestData fields

// Maps Messaging Producer/Client Span to AppInsights RemoteDependencyData
// https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/trace/semantic_conventions/messaging.md
func fillRemoteDependencyDataMessaging(span ptrace.Span, data *contracts.RemoteDependencyData) {
	_ = "STUB: not implemented"
	return
}

// TODO Understand how to map attributes to RemoteDependencyData fields

// Copies all attributes to either properties or measurements and passes the key/value to another mapping function
func copyAndMapAttributes(
	attributeMap pcommon.Map,
	properties map[string]string,
	mappingFunc func(k string, v pcommon.Value),
) {
	_ = "STUB: not implemented"
	return
}

// Copies all attributes to either properties or measurements without any kind of mapping to a known set of attributes
func copyAttributesWithoutMapping(
	attributeMap pcommon.Map,
	properties map[string]string,
) {
	_ = "STUB: not implemented"
	return
}

// Attribute extraction logic for HTTP Span attributes
func copyAndExtractHTTPAttributes(
	attributeMap pcommon.Map,
	properties map[string]string,
) *httpAttributes {
	_ = "STUB: not implemented"
	return nil
}

// Attribute extraction logic for RPC Span attributes
func copyAndExtractRPCAttributes(
	attributeMap pcommon.Map,
	properties map[string]string,
) *rpcAttributes {
	_ = "STUB: not implemented"
	return nil
}

// Attribute extraction logic for Database Span attributes
func copyAndExtractDatabaseAttributes(
	attributeMap pcommon.Map,
	properties map[string]string,
) *databaseAttributes {
	_ = "STUB: not implemented"
	return nil
}

// Attribute extraction logic for Messaging Span attributes
func copyAndExtractMessagingAttributes(
	attributeMap pcommon.Map,
	properties map[string]string,
) *messagingAttributes {
	_ = "STUB: not implemented"
	return nil
}

// Attribute extraction logic for Span event exception attributes
func copyAndExtractExceptionAttributes(
	attributeMap pcommon.Map,
	properties map[string]string,
) *exceptionAttributes {
	_ = "STUB: not implemented"
	return nil
}

func formatSpanDuration(span ptrace.Span) string { _ = "STUB: not implemented"; return "" }

// Maps incoming Span to a type defined in the specification
func mapIncomingSpanToType(attributeMap pcommon.Map) spanType {
	_ = "STUB: not implemented"
	// No attributes
	return *new(spanType)
}

// RPC

// HTTP

// Database

// Messaging

// https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/trace/api.md#set-status
func getDefaultFormattedSpanStatus(spanStatus ptrace.Status) (statusCodeAsString string, success bool) {
	_ = "STUB: not implemented"
	return "", false
}

func writeFormatedFromNetworkServerOrClient(networkAttributes *networkAttributes, addressName string, addressPort int64, sb *strings.Builder) {
	_ = "STUB: not implemented"
	// server.address or client.address
	return
}

func setAttributeValueAsProperty(
	key string,
	attributeValue pcommon.Value,
	properties map[string]string,
) {
	_ = "STUB: not implemented"
	return
}

func prefixIfNecessary(s, prefix string) string { _ = "STUB: not implemented"; return "" }

func sanitize(sanitizeFunc func() []string, logger *zap.Logger) { _ = "STUB: not implemented"; return }

func sanitizeWithCallback(sanitizeFunc func() []string, warningCallback func(string), logger *zap.Logger) {
	_ = "STUB: not implemented"
	return
}

// TODO error handling
