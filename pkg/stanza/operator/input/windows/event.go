// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package windows // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/windows"

// systemPropertiesRenderContext stores a custom rendering context to get only the event properties.
var (
	systemPropertiesRenderContext    = uintptr(0)
	systemPropertiesRenderContextErr error
)

func init() {
	// This is not expected to fail, however, collecting the error if a new failure mode appears.
	systemPropertiesRenderContext, systemPropertiesRenderContextErr = evtCreateRenderContext(0, nil, EvtRenderContextSystem)
}

// Event is an event stored in windows event log.
type Event struct {
	handle uintptr
}

// GetPublisherName will get the publisher name of the event.
func (e *Event) GetPublisherName(buffer *Buffer) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// utf16PtrToString converts Windows API LPWSTR (pointer to string) to go string
func utf16PtrToString(s *uint16) string { _ = "STUB: not implemented"; return "" }

// NewEvent will create a new event from an event handle.
func NewEvent(handle uintptr) Event { _ = "STUB: not implemented"; return *new(Event) }

// RenderSimple will render the event as a parsedEvent without formatted info.
func (e *Event) RenderSimple(buffer *Buffer) (parsedEvent, error) {
	_ = "STUB: not implemented"
	return *new(parsedEvent), nil
}

// RenderSimpleRaw will render the event XML but unmarshal only the fields
// needed when raw=true. Use this to avoid populating fields that will not be used.
func (e *Event) RenderSimpleRaw(buffer *Buffer) (parsedEvent, error) {
	_ = "STUB: not implemented"
	return *new(parsedEvent), nil
}

func (e *Event) renderSimpleEventXML(buffer *Buffer, unmarshal func([]byte) (parsedEvent, error)) (parsedEvent, error) {
	_ = "STUB: not implemented"
	return *new(parsedEvent), nil
}

// RenderDeep will render the event as a parsedEvent with all available formatted info.
func (e *Event) RenderDeep(buffer *Buffer, publisher Publisher) (parsedEvent, error) {
	_ = "STUB: not implemented"
	return *new(parsedEvent), nil
}

// RenderDeepRaw will render the event with formatted info but unmarshal only
// the fields needed when raw=true: timestamp, level, and the rendered level
// from RenderingInfo for accurate severity mapping.
func (e *Event) RenderDeepRaw(buffer *Buffer, publisher Publisher) (parsedEvent, error) {
	_ = "STUB: not implemented"
	return *new(parsedEvent), nil
}

func (e *Event) renderDeepEventXML(buffer *Buffer, publisher Publisher, unmarshal func([]byte) (parsedEvent, error)) (parsedEvent, error) {
	_ = "STUB: not implemented"
	return *new(parsedEvent), nil
}

// Close will close the event handle.
func (e *Event) Close() error { _ = "STUB: not implemented"; return nil }
