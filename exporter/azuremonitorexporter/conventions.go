// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azuremonitorexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azuremonitorexporter"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

/*
	This file encapsulates the extraction logic for the various kinds attributes
*/

const (
	// TODO replace with convention.* values once/if available
	attributeOtelStatusCode                          string = "otel.status_code"
	attributeOtelStatusDescription                   string = "otel.status_description"
	attributeMicrosoftCustomEventName                string = "microsoft.custom_event.name"
	attributeApplicationInsightsEventMarkerAttribute string = "APPLICATION_INSIGHTS_EVENT_MARKER_ATTRIBUTE"
)

// clientAttributes is the set of known client attributes
type clientAttributes struct {
	// see https://github.com/open-telemetry/semantic-conventions/blob/v1.37.0/docs/registry/attributes/client.md
	ClientAddress string
	ClientPort    int64
}

// serverAttributes is the set of known server attributes
type serverAttributes struct {
	// see https://github.com/open-telemetry/semantic-conventions/blob/v1.37.0/docs/registry/attributes/server.md
	ServerAddress string
	ServerPort    int64
}

// networkAttributes is the set of known network attributes
type networkAttributes struct {
	// see https://github.com/open-telemetry/semantic-conventions/blob/v1.37.0/docs/registry/attributes/network.md
	NetworkLocalAddress    string
	NetworkLocalPort       int64
	NetworkPeerAddress     string
	NetworkPeerPort        int64
	NetworkProtocolName    string
	NetworkProtocolVersion string
	NetworkTransport       string
	NetworkType            string
}

// MapAttribute attempts to map a Span attribute to one of the known types
func (attrs *networkAttributes) MapAttribute(k string, v pcommon.Value) bool {
	_ = "STUB: not implemented"
	return false
}

// urlAttributes is the set of known attributes for URL
type urlAttributes struct {
	// common attributes
	// https://github.com/open-telemetry/semantic-conventions/blob/v1.37.0/docs/registry/attributes/url.md
	URLFragment string
	URLFull     string
	URLPath     string
	URLQuery    string
	URLScheme   string
}

type userAgentAttributes struct {
	UserAgentOriginal string
	UserAgentName     string
	UserAgentVersion  string
}

// httpAttributes is the set of known attributes for HTTP Spans
type httpAttributes struct {
	// common attributes
	// https://github.com/open-telemetry/semantic-conventions/blob/v1.37.0/docs/registry/attributes/http.md
	HTTPRequestHeaders        map[string][]string
	HTTPRequestMethod         string
	HTTPRequestMethodOriginal string
	HTTPRequestResendCount    int64
	HTTPResponseHeaders       map[string][]string
	HTTPResponseStatusCode    int64
	HTTPRoute                 string

	HTTPRequestBodySize  int64
	HTTPResponseBodySize int64

	URLAttributes       urlAttributes
	ClientAttributes    clientAttributes
	ServerAttributes    serverAttributes
	UserAgentAttributes userAgentAttributes

	// any net.*
	NetworkAttributes networkAttributes
}

func setHeader(headers *map[string][]string, key string, v pcommon.Value) {
	_ = "STUB: not implemented"
	return
}

// MapAttribute attempts to map a Span attribute to one of the known types
func (attrs *httpAttributes) MapAttribute(k string, v pcommon.Value) bool {
	_ = "STUB: not implemented"
	return false
}

// URL attributes

// Network/server/client address attributes (new, replacing http.host)

// User agent attributes

// rpcAttributes is the set of known attributes for RPC Spans
type rpcAttributes struct {
	RPCSystem         string
	RPCService        string
	RPCMethod         string
	RPCGRPCStatusCode int64

	ClientAttributes  clientAttributes
	ServerAttributes  serverAttributes
	NetworkAttributes networkAttributes
}

// MapAttribute attempts to map a Span attribute to one of the known types
func (attrs *rpcAttributes) MapAttribute(k string, v pcommon.Value) bool {
	_ = "STUB: not implemented"
	return false
}

// databaseAttributes is the set of known attributes for Database Spans
type databaseAttributes struct {
	DBCollectionName      string
	DBNamespace           string
	DBOperationBatchSize  int64
	DBOperationName       string
	DBQuerySummary        string
	DBQueryText           string
	DBResponseStatusCode  string
	DBStoredProcedureName string
	DBSystemName          string

	ClientAttributes clientAttributes
	ServerAttributes serverAttributes

	NetworkAttributes networkAttributes
}

// MapAttribute attempts to map a Span attribute to one of the known types
func (attrs *databaseAttributes) MapAttribute(k string, v pcommon.Value) bool {
	_ = "STUB: not implemented"
	return false
}

// messagingAttributes is the set of known attributes for Messaging Spans
type messagingAttributes struct {
	MessagingBatchMessageCount      int64
	MessagingClientID               string
	MessagingConsumerGroup          string
	MessagingDestinationAnonymous   bool
	MessagingDestination            string
	MessagingDestinationPartitionID string
	MessagingDestinationSubName     string
	MessagingDestinationTemplate    string
	MessagingDestinationTemporary   bool
	MessagingMessageBodySize        int64
	MessagingMessageConversationID  string
	MessagingMessageEnvelopeSize    int64
	MessagingMessageID              string
	MessagingOperation              string
	MessagingOperationType          string
	MessagingSystem                 string

	ClientAttributes clientAttributes
	ServerAttributes serverAttributes

	NetworkAttributes networkAttributes
}

// MapAttribute attempts to map a Span attribute to one of the known types
func (attrs *messagingAttributes) MapAttribute(k string, v pcommon.Value) bool {
	_ = "STUB: not implemented"
	return false
}

// exceptionAttributes is the set of known attributes for Exception events
type exceptionAttributes struct {
	ExceptionMessage    string
	ExceptionStackTrace string
	ExceptionType       string
}

// MapAttribute attempts to map a SpanEvent attribute to one of the known types
func (attrs *exceptionAttributes) MapAttribute(k string, v pcommon.Value) bool {
	_ = "STUB: not implemented"
	return false
}

// Tries to return the value of the attribute as an int64
func getAttributeValueAsInt(attributeValue pcommon.Value) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// try to cast the string values to int64
