// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azurefunctionsreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azurefunctionsreceiver"

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azurefunctionsreceiver/internal/transport"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azurefunctionsreceiver/internal/trigger"
)

// Protocol parses the Azure Functions invoke HTTP body into a ParsedRequest and
// writes success or failure responses. Different triggers may share the same
// protocol (e.g. invoke envelope + transport decode) or use a different one.
type Protocol interface {
	ParseRequest(bindingName string, body []byte) (trigger.ParsedRequest, error)
	Success(w http.ResponseWriter)
	Failure(w http.ResponseWriter, err error, body []byte)
}

// profile binds an Azure Functions binding name to a Protocol and Consumer.
// The generic HTTP handler uses it to: parse request with protocol, then consume with consumer.
type profile struct {
	binding  string
	protocol Protocol
	consumer trigger.Consumer
}

func newProfile(binding string, protocol Protocol, consumer trigger.Consumer) profile {
	_ = "STUB: not implemented"
	return *new(profile)
}

// MetadataExtractor turns raw invoke Metadata JSON into resource attributes.
// Trigger-specific (e.g. Event Hub, HTTP); pass nil for no metadata extraction.
type MetadataExtractor func(raw []byte) map[string]string

// invokeProtocol implements Protocol for the Azure Functions invoke envelope:
// parse body, decode Data[bindingName] via transport, optionally fill Metadata using an extractor.
type invokeProtocol struct {
	decoder   *transport.BinaryDecoder
	logger    *zap.Logger
	extractor MetadataExtractor
}

// newInvokeProtocol returns a Protocol that parses invoke requests and decodes the binding
// payload. If extractor is non-nil and invoke Metadata is present, it is used to fill ParsedRequest.Metadata.
func newInvokeProtocol(decoder *transport.BinaryDecoder, logger *zap.Logger, extractor MetadataExtractor) Protocol {
	_ = "STUB: not implemented"
	return *new(Protocol)
}

// ParseRequest implements Protocol.
func (p *invokeProtocol) ParseRequest(bindingName string, body []byte) (trigger.ParsedRequest, error) {
	_ = "STUB: not implemented"
	return *new(trigger.ParsedRequest), nil
}

// Success implements Protocol.
func (*invokeProtocol) Success(w http.ResponseWriter) { _ = "STUB: not implemented"; return }

// Failure implements Protocol.
func (p *invokeProtocol) Failure(w http.ResponseWriter, err error, body []byte) {
	_ = "STUB: not implemented"
	return
}

// createHandler returns an HTTP handler for the given profile. It reads the body,
// parses it with the protocol, consumes with the consumer, and responds with success or failure.
// The same flow is used for every trigger; no trigger-specific logic here.
func createHandler(p profile) http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }
