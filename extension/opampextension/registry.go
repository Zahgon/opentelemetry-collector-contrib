// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package opampextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/opampextension"

import (
	"container/list"
	"sync"

	"github.com/open-telemetry/opamp-go/protobufs"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/opampcustommessages"
)

// customCapabilityClient is a subset of OpAMP client containing only the methods needed for the customCapabilityRegistry.
type customCapabilityClient interface {
	SetCustomCapabilities(customCapabilities *protobufs.CustomCapabilities) error
	SendCustomMessage(message *protobufs.CustomMessage) (messageSendingChannel chan struct{}, err error)
}

type customCapabilityRegistry struct {
	mux                     *sync.Mutex
	capabilityToMsgChannels map[string]*list.List
	client                  customCapabilityClient
	logger                  *zap.Logger
}

var _ opampcustommessages.CustomCapabilityRegistry = (*customCapabilityRegistry)(nil)

func newCustomCapabilityRegistry(logger *zap.Logger, client customCapabilityClient) *customCapabilityRegistry {
	_ = "STUB: not implemented"
	return nil
}

// Register implements CustomCapabilityRegistry.Register
func (cr *customCapabilityRegistry) Register(capability string, opts ...opampcustommessages.CustomCapabilityRegisterOption) (opampcustommessages.CustomCapabilityHandler, error) {
	_ = "STUB: not implemented"
	return *new(opampcustommessages.CustomCapabilityHandler), nil
}

// ProcessMessage processes a custom message, asynchronously broadcasting it to all registered capability handlers for
// the messages capability.
func (cr customCapabilityRegistry) ProcessMessage(cm *protobufs.CustomMessage) {
	_ = "STUB: not implemented"
	return
}

// If the channel is full, we will skip sending the message to the receiver.
// We do this because we don't want a misbehaving component to be able to
// block the opamp extension, or block other components from receiving messages.

// removeCapabilityFunc returns a func that removes the custom capability with the given msg channel list element and sender,
// then recalculates and sets the list of custom capabilities on the OpAMP client.
func (cr *customCapabilityRegistry) removeCapabilityFunc(capability string, callbackElement *list.Element) func() {
	_ = "STUB: not implemented"
	return nil
}

// Since there are no more callbacks for this capability,
// this capability is no longer supported

// It's OK if we couldn't actually remove the capability, it just means we won't
// notify the server properly, and the server may send us messages that we have no associated callbacks for.

// capabilities gives the current set of custom capabilities with at least one
// callback registered.
func (cr *customCapabilityRegistry) capabilities() []string { _ = "STUB: not implemented"; return nil }

type customMessageHandler struct {
	// unregisteredMux protects unregistered, and makes sure that a message cannot be sent
	// on an unregistered capability.
	unregisteredMux *sync.Mutex

	capability               string
	opampClient              customCapabilityClient
	registry                 *customCapabilityRegistry
	sendChan                 <-chan *protobufs.CustomMessage
	unregisterCapabilityFunc func()

	unregistered bool
}

var _ opampcustommessages.CustomCapabilityHandler = (*customMessageHandler)(nil)

func newCustomMessageHandler(
	registry *customCapabilityRegistry,
	opampClient customCapabilityClient,
	capability string,
	sendChan <-chan *protobufs.CustomMessage,
	unregisterCapabilityFunc func(),
) *customMessageHandler {
	_ = "STUB: not implemented"
	return nil
}

// Message implements CustomCapabilityHandler.Message
func (c *customMessageHandler) Message() <-chan *protobufs.CustomMessage {
	_ = "STUB: not implemented"

	// SendMessage implements CustomCapabilityHandler.SendMessage
	return nil
}

func (c *customMessageHandler) SendMessage(messageType string, message []byte) (messageSendingChannel chan struct{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Unregister implements CustomCapabilityHandler.Unregister
func (c *customMessageHandler) Unregister() { _ = "STUB: not implemented"; return }
