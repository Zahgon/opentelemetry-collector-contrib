// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package unmarshaler // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/azureencodingextension/internal/unmarshaler"

import (
	"encoding/json"

	"go.opentelemetry.io/collector/pdata/pcommon"
)

type RecordsBatchFormat int

// Supported wrapper formats of Azure Logs Records batch
const (
	FormatUnknown RecordsBatchFormat = iota
	FormatObjectRecords
	FormatJSONArray
	FormatNDJSON
)

// JSON Path expressions that matches specific wrapper format
const (
	// As exported to Azure Event Hub, e.g. `{"records": [ {...}, {...} ]}`
	JSONPathEventHubRecords = "$.records[*]"
	// As exported to Azure Blob Storage, e.g. `[ {...}, {...} ]`
	JSONPathBlobStorageRecords = "$[*]"
)

// Commonly used attributes non-SemConv attributes across all telemetry signals
const (
	// OpenTelemetry attribute name for Azure Resource Log Category,
	// mostly used in "logs" telemetry, but also can be present in "metrics" and "traces"
	AttributeAzureCategory = "azure.category"

	// OpenTelemetry attribute name for Azure Resource Operation Name
	AttributeAzureOperationName = "azure.operation.name"
)

const (
	// OpenTelemetry attribute name for Destination original address,
	// in case if the value could not be parsed into
	// `destination.address` and `destination.port` attributes
	attributeDestinationAddressOriginal = "destination.original_address"

	// OpenTelemetry attribute name for Client original address,
	// in case if the value could not be parsed into
	// `client.address` and `client.port` attributes
	attributeClientAddressOriginal = "client.original_address"

	// OpenTelemetry attribute name for Client original address,
	// in case if the value could not be parsed into
	// `server.address` and `server.port` attributes
	attributeServerAddressOriginal = "server.original_address"

	// OpenTelemetry attribute name for Network Peer original address,
	// in case if the value could not be parsed into
	// `network.peer.address` and `network.peer.port` attributes
	attributeNetworkPeerAddressOriginal = "network.peer.original_address"

	// OpenTelemetry attribute name for Network Local original address,
	// in case if the value could not be parsed into
	// `network.local.address` and `network.local.port` attributes
	attributeNetworkLocalAddressOriginal = "network.local.original_address"
)

const (
	originalSuffix = "_original"
	redactedStr    = "REDACTED"
	// When field value is set to "-" it's indicate that the data was unknown
	// or unavailable, or that the field was not applicable to this request
	unknownField = "-"
)

// DetectWrapperFormat tries to detect format based on provided bytes input
// At the moment we support only JSON Array, "records" and ND JSON formats,
// anything else is detected as unsupported format
func DetectWrapperFormat(input []byte) (RecordsBatchFormat, error) {
	_ = "STUB: not implemented"

	// Not an error, just empty input
	return *new(RecordsBatchFormat), nil
}

// That's seems to be JSON Array format, e.g. `[ {...}, {...} ]`

// We'll scan only first 100 bytes to avoid performance bottleneck here

// Make sure that it's a top-level field, i.e. before it we have only '{' not-whitespace character

// Make sure that "records" field is an array, i.e. next non-whitespace characters are ':' and '['

// Next character should be '['

// That's seems to be JSON object format with "records" field, e.g. `{"records": [ {...}, {...} ]}`

// Detect ND JSON

// Everything else - is unsupported
// Will include first bytes of input in error to simplify further debug

// AsTimestamp tries to parse a string with timestamp into OpenTelemetry
// using provided list of formats layouts.
// First is used ISO8601 parser - as it's the most common format for Azure telemetry data
// After that will be tested provided formats if any
// If the string cannot be parsed, it will return zero and the error.
func AsTimestamp(s string, formats ...string) (pcommon.Timestamp, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp), nil
}

// In most cases - Azure telemetry data has ISO8601 formatted timestamp
// So, to optimize performance we'll check it first here

// Try parsing with provided formats first

// AttrPutStrIf is a helper function to set a string attribute
// only if the value is not empty
func AttrPutStrIf(attrs pcommon.Map, attrKey, attrValue string) { _ = "STUB: not implemented"; return }

// AttrPutStrPtrIf is a helper function to set a string attribute
// only if the value exists and is not empty
func AttrPutStrPtrIf(attrs pcommon.Map, attrKey string, attrValue *string) {
	_ = "STUB: not implemented"
	return
}

// AttrPutIntNumberIf is a helper function to set an int64 attribute with defined key,
// trying to parse it from json.Number value
// If parsing failed - no attribute will be set
func AttrPutIntNumberIf(attrs pcommon.Map, attrKey string, attrValue json.Number) {
	_ = "STUB: not implemented"
	return
}

// AttrPutIntNumberPtrIf is a same function as AttrPutIntNumberIf but
// accepts a pointer to json.Number instead of value
func AttrPutIntNumberPtrIf(attrs pcommon.Map, attrKey string, attrValue *json.Number) {
	_ = "STUB: not implemented"
	return
}

// AttrPutIntNumberIf is a helper function to set an float64 attribute with defined key,
// trying to parse it from json.Number value
// If parsing failed - no attribute will be set
func AttrPutFloatNumberIf(attrs pcommon.Map, attrKey string, attrValue json.Number) {
	_ = "STUB: not implemented"
	return
}

// AttrPutFloatNumberPtrIf is a same function as AttrPutIntNumberPtrIf but
// accepts a pointer to json.Number instead of value
func AttrPutFloatNumberPtrIf(attrs pcommon.Map, attrKey string, attrValue *json.Number) {
	_ = "STUB: not implemented"
	return
}

// attrPutMap is a helper function to set a map attribute with defined key,
// trying to parse it from raw value
// If parsing failed - no attribute will be set
func AttrPutMapIf(attrs pcommon.Map, attrKey string, attrValue map[string]any) {
	_ = "STUB: not implemented"
	return
}

// Failed to parse - put string representation of the attrValue

// AttrPutURLParsed is a helper function that parses provided URI into set
// of OpenTelemetry SemConv attributes
// It will set `url.original` on any non-empty URI string and other SemConv
// attributes in case if provided URI can be parsed
// Puts maximum 8 attributes
func AttrPutURLParsed(attrs pcommon.Map, uri string) { _ = "STUB: not implemented"; return }

// Try parsing provided URI

// Put original URI only if parsing failed to avoid data duplication
// unstable SemConv

// Mask credentials according to SemConv specs

// Set url.full

// Trying to parse port

// We can safely ignore the parse error here because `url.Parse` will fail
// in case of incorrect port, so it's parsable here 99%

// unstable SemConv

// Set other valuable `url.*` attributes according to SemConv

// unstable SemConv

// AttrPutHostPortIf tries to parse provided `value` as "host:port" format and put result
// into respective `addrKey` for "host" and `portKey` for "port"
// If no port was detected in `value` - only `addrKey` will be set
// If value is not in "host:port" format or port is invalid - then
// `fallbackKey` attribute will be set with the original value
// Puts at most 2 attributes
func AttrPutHostPortIf(attrs pcommon.Map, addrKey, portKey, value string) {
	_ = "STUB: not implemented"

	// Nothing to do here
	return
}

// Do not store "unspecified" address for both IPv4 and IPv6,
// including cases when it contains "unspecified" port
// as it doesn't provide any valuable information

// Not a "host:port" format or only IPv6
// put only "*.address" attribute

// Determine name of fallback attribute key

// Try to parse host:port

// net.SplitHostPort actually does not validates if port is a valid number,
// so will try to convert it and return error on failure
