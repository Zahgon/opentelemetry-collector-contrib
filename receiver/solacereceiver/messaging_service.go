// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package solacereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/solacereceiver"

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/Azure/go-amqp"
	"go.uber.org/zap"
)

// inboundMessage is an alias for amqp.Message
type inboundMessage = amqp.Message

// messagingService abstracts out the AMQP transport capabilities for unit testing
type messagingService interface {
	dial(ctx context.Context) error
	close(ctx context.Context)
	receiveMessage(ctx context.Context) (*inboundMessage, error)
	accept(ctx context.Context, msg *inboundMessage) error
	failed(ctx context.Context, msg *inboundMessage) error
}

// messagingServiceFactory is a factory to create new messagingService instances
type messagingServiceFactory func() messagingService

// newAMQPMessagingServiceFactory creates a new messagingServiceFactory backed by AMQP
func newAMQPMessagingServiceFactory(cfg *Config, logger *zap.Logger) (messagingServiceFactory, error) {
	_ = "STUB: not implemented"
	return *new(messagingServiceFactory), nil
}

// Use the default load config for TLS. Note that in the case where "insecure" is true and no
// ca file is provided, tlsConfig will be nil representing a plaintext connection.

// If the TLS config is nil, insecure is true and we should use amqp rather than amqps

type amqpConnectConfig struct {
	// connect config
	addr       string
	saslConfig amqp.SASLType
	tlsConfig  *tls.Config
}

type amqpReceiverConfig struct {
	queue       string
	maxUnacked  int32
	batchMaxAge time.Duration
}

type amqpMessagingService struct {
	// factory fields
	connectConfig  *amqpConnectConfig
	receiverConfig *amqpReceiverConfig
	logger         *zap.Logger

	// runtime fields
	client   *amqp.Conn
	session  *amqp.Session
	receiver *amqp.Receiver
}

// dialFunc is abstracted out into a variable in order for substitutions
var dialFunc = amqp.Dial

// telemetryLinkName will be used to create the single receiver link in order to standardize the connection.
// Mainly useful for testing to mock amqp frames.
const telemetryLinkName = "rx"

func (m *amqpMessagingService) dial(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (m *amqpMessagingService) close(ctx context.Context) { _ = "STUB: not implemented"; return }

func (m *amqpMessagingService) receiveMessage(ctx context.Context) (*inboundMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *amqpMessagingService) accept(ctx context.Context, msg *inboundMessage) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *amqpMessagingService) failed(ctx context.Context, msg *inboundMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// Allow for substitution in testing to assert correct data is passed to AMQP
// Due to the way that AMQP authentication is configured in Azure/amqp, we
// need to monkey substitute here since ConnSASL<auth> returns a function that
// acts on a private struct meaning we cannot meaningfully assert validity otherwise.
var (
	connSASLPlain    func(username, password string) amqp.SASLType                            = amqp.SASLTypePlain
	connSASLXOAUTH2  func(username, bearer string, maxFrameSizeOverride uint32) amqp.SASLType = amqp.SASLTypeXOAUTH2
	connSASLExternal func(resp string) amqp.SASLType                                          = amqp.SASLTypeExternal
)

// toAMQPAuthentication configures authentication in amqp.ConnOption slice
func toAMQPAuthentication(config *Config) (amqp.SASLType, error) {
	_ = "STUB: not implemented"
	return *new(amqp.SASLType), nil
}
