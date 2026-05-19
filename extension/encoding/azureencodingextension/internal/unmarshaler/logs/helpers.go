// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package logs // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/azureencodingextension/internal/unmarshaler/logs"

import (
	"crypto/tls"
	"encoding/json"
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
)

const (
	// OpenTelemetry attribute name for TLS protocol original value,
	// in case if the value could not be parsed into
	// `tls.protocol.name` and `tls.protocol.version` attributes
	attributeTLSProtocolOriginal = "tls.protocol.original"
	// OpenTelemetry attribute name for Network protocol original value,
	// in case if the value could not be parsed into
	// `network.protocol.name` and `network.protocol.version` attributes
	attributeNetworkProtocolOriginal = "network.protocol.original"
)

var (
	// As tls.VersionSSL30 constant is deprecated, will use simple string here
	tlsVersionSSLv3 = "SSLv3"
	tlsVersionTLS10 = tls.VersionName(tls.VersionTLS10)
	tlsVersionTLS11 = tls.VersionName(tls.VersionTLS11)
	tlsVersionTLS12 = tls.VersionName(tls.VersionTLS12)
	tlsVersionTLS13 = tls.VersionName(tls.VersionTLS13)
	// HTTP protocol versions
	httpVersion09 = "HTTP/0.9"
	httpVersion10 = "HTTP/1.0"
	httpVersion11 = "HTTP/1.1"
	httpVersion20 = "HTTP/2.0"
	httpVersion30 = "HTTP/3.0"
)

// asSeverity converts the Azure log level to equivalent
// OpenTelemetry severity numbers. If the log level is not
// valid, then the 'Unspecified' value is returned.
// According to the documentation, the level Must be one of:
// `Informational`, `Warning`, `Error` or `Critical`.
// see https://learn.microsoft.com/en-us/azure/azure-monitor/platform/resource-logs-schema
func asSeverity(number json.Number) plog.SeverityNumber {
	_ = "STUB: not implemented"
	return *new(plog.SeverityNumber)
}

// attrPutTLSProtoIf tries to parse provided value as TLS security protocol version,
// for example, "TLS 1.2" will be parsed into tls.protocol.name = "TLS" and tls.protocol.version = "1.2"
// If the value is not recognized - will set original value into "tls.protocol.original" attribute
// Puts at most 2 attributes
func attrPutTLSProtoIf(attrs pcommon.Map, securityProtocol string) {
	_ = "STUB: not implemented"
	return
}

// Nothing to do here

// attrPutHTTPProtoIf tries to parse provided value as HTTP protocol version,
// for example, "HTTP/1.1" will be parsed into network.protocol.name = "http" and network.protocol.version = "1.1"
// If the value is not recognized - will set original value into "network.protocol.original" attribute
// Puts at most 2 attributes
func attrPutHTTPProtoIf(attrs pcommon.Map, httpProtocol string) { _ = "STUB: not implemented"; return }

// Nothing to do here

// Protocol values SHOULD be normalized to lowercase as per SemConv

// convertInvalidSingleQuotedJSON tries to convert invalid, single quoted, JSON
// into valid and parsable JSON by substituting `'` to `"`, taking in account
// potential escaped single quotes, e.g. `\'`
func convertInvalidSingleQuotedJSON(data []byte) []byte { _ = "STUB: not implemented"; return nil }

// Enter quoted string

// Mark that we are in quoted string

// Replace single quote to double quote

// Leave quoted string

// Mark that we left quoted string

// Write closing double quote instead single quote

// Unescape escaped single quote inside quoted string

// Sometimes Azure double escapes single quotes,
// we will keep it as is to keep escaping consistent
// in whole line

// Escape unescaped double quote

// All other chars are going to output as-is

// convertStringToJSONNumber is a special helper function to mitigate issue with
// invalid JSON in Azure Log Records
// In some cases Azure can put into field, which is expected to be number (int/real),
// invalid data like empty string ("") or even dash ("-")
// This misbehavior was detected at leas in ApplicationGatewayAccessLog
// In any case of non-number input - it will return json.Number("")
func convertStringToJSONNumber(s string) json.Number {
	_ = "STUB: not implemented"
	return *

	// Actually it's invalid json.Number as it's will return an error on
	// any Int64() or Float64() calls, but that's OK for our case
	new(json.Number)
}

// Scan input string for allowed chars that represents numbers

// parseUnixTimestamp parses a Unix timestamp string (seconds since epoch) and returns a time.Time
func parseUnixTimestamp(s string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// unmarshalStringOrObjectJSON handles Azure properties fields that can be either
// a JSON object or a stringified JSON string (where Azure wraps the object in quotes).
// Some Azure Event Hub sources emit `"properties": "{\"key\":\"value\"}"` instead of
// `"properties": {"key":"value"}`.
func unmarshalStringOrObjectJSON(data []byte, v any) error { _ = "STUB: not implemented"; return nil }
