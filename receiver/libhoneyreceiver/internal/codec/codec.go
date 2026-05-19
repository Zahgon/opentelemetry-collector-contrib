// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package codec provides encoding and decoding for the libhoney event format.
// It handles both JSON and MessagePack formats, supporting single events,
// batches, flat events, and structured events with header-based metadata.
package codec // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/libhoneyreceiver/internal/codec"

import (
	"net/http"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/libhoneyreceiver/internal/libhoneyevent"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/libhoneyreceiver/internal/response"
)

const (
	JSONContentType    = "application/json"
	MsgpackContentType = "application/msgpack"
)

var (
	JsEncoder = &jsonEncoder{}
	MpEncoder = &msgpackEncoder{}
)

// Encoder handles marshaling of libhoney response batches in the appropriate format
type Encoder interface {
	MarshalResponse([]response.ResponseInBatch) ([]byte, error)
	ContentType() string
}

type jsonEncoder struct{}

func (jsonEncoder) MarshalResponse(batchResponse []response.ResponseInBatch) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (jsonEncoder) ContentType() string { _ = "STUB: not implemented"; return "" }

type msgpackEncoder struct{}

func (msgpackEncoder) MarshalResponse(batchResponse []response.ResponseInBatch) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (msgpackEncoder) ContentType() string { _ = "STUB: not implemented"; return "" }

// GetEncoder returns the appropriate encoder for the given content type
func GetEncoder(contentType string) (Encoder, error) {
	_ = "STUB: not implemented"
	return *new(Encoder), nil
}

// DecodeEvents decodes libhoney events from the request body based on content type
func DecodeEvents(contentType string, body []byte, headers http.Header) ([]libhoneyevent.LibhoneyEvent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// decodeMsgpack decodes msgpack-encoded libhoney events
func decodeMsgpack(body []byte, headers http.Header) ([]libhoneyevent.LibhoneyEvent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// decodeJSON decodes JSON-encoded libhoney events
func decodeJSON(body []byte, headers http.Header) ([]libhoneyevent.LibhoneyEvent, error) {
	_ = "STUB: not implemented"
	// Check first non-whitespace character to determine structure
	return nil, nil
}

// isSingleMsgpackObject checks if the msgpack data represents a single object (map) vs an array
// msgpack format: 0x80-0x8f = fixmap, 0xde = map16, 0xdf = map32
//
//	0x90-0x9f = fixarray, 0xdc = array16, 0xdd = array32
func isSingleMsgpackObject(body []byte) bool { _ = "STUB: not implemented"; return false }

// Check if it's a map type

// decodeSingleMsgpackEvent decodes a single msgpack event
func decodeSingleMsgpackEvent(body []byte, headers http.Header) ([]libhoneyevent.LibhoneyEvent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Apply headers for single event (X-Honeycomb-Event-Time, X-Honeycomb-Samplerate)

// Wrap flat event if needed

// Construct LibhoneyEvent directly from the map

// decodeMsgpackArray decodes an array of msgpack events
func decodeMsgpackArray(body []byte) ([]libhoneyevent.LibhoneyEvent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// decodeSingleJSONEvent decodes a single JSON event
func decodeSingleJSONEvent(body []byte, headers http.Header) ([]libhoneyevent.LibhoneyEvent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Apply headers for single event (X-Honeycomb-Event-Time, X-Honeycomb-Samplerate)

// Wrap flat event if needed

// Now unmarshal the structured event

// decodeJSONArray decodes an array of JSON events
func decodeJSONArray(body []byte) ([]libhoneyevent.LibhoneyEvent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// applyHeadersToEvent applies X-Honeycomb-Event-Time and X-Honeycomb-Samplerate headers to the event
// Only applies if not already present in the body
func applyHeadersToEvent(rawEvent map[string]any, headers http.Header) {
	_ = "STUB: not implemented"
	return
}

// Convert string to number for compatibility

// Fallback to string

// wrapFlatEventIfNeeded wraps a flat event (one without a "data" field) into the structured format
func wrapFlatEventIfNeeded(rawEvent map[string]any) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

// Already has a data field, no wrapping needed

// Flat event format - wrap all fields into data

// Preserve time and samplerate at top level if they exist

// buildLibhoneyEventFromMap constructs a LibhoneyEvent from a raw map
func buildLibhoneyEventFromMap(rawEvent map[string]any) libhoneyevent.LibhoneyEvent {
	_ = "STUB: not implemented"
	return *new(libhoneyevent.LibhoneyEvent)
}

// default

// Extract samplerate

// Extract time

// Extract data

// ValidateEvents performs basic validation on decoded events
func ValidateEvents(events []libhoneyevent.LibhoneyEvent) error {
	_ = "STUB: not implemented"
	return nil
}
