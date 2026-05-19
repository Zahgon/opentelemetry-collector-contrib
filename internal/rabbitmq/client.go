// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package rabbitmq // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/rabbitmq"

import (
	"context"
	"crypto/tls"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type AmqpClient interface {
	DialConfig(config DialConfig) (Connection, error)
}

type Connection interface {
	ReconnectIfUnhealthy() error
	IsClosed() bool
	Channel() (Channel, error)
	NotifyClose(receiver chan *amqp.Error) chan *amqp.Error
	Close() error
}

type Channel interface {
	Confirm(noWait bool) error
	PublishWithDeferredConfirmWithContext(ctx context.Context, exchange, key string, mandatory, immediate bool, msg amqp.Publishing) (DeferredConfirmation, error)
	IsClosed() bool
	Close() error
}

type DeferredConfirmation interface {
	Done() <-chan struct{}
	Acked() bool
}

type connectionHolder struct {
	url              string
	config           amqp.Config
	connection       *amqp.Connection
	logger           *zap.Logger
	connLock         *sync.Mutex
	connectionErrors chan *amqp.Error
}

type channelHolder struct {
	channel *amqp.Channel
}

type deferredConfirmationHolder struct {
	confirmation *amqp.DeferredConfirmation
}

type DialConfig struct {
	URL               string
	Vhost             string
	Auth              amqp.Authentication
	ConnectionTimeout time.Duration
	Heartbeat         time.Duration
	TLS               *tls.Config
	ConnectionName    string
}

func NewAmqpClient(logger *zap.Logger) AmqpClient {
	_ = "STUB: not implemented"
	return *new(AmqpClient)
}

type client struct {
	logger *zap.Logger
}

func (c *client) DialConfig(config DialConfig) (Connection, error) {
	_ = "STUB: not implemented"
	return *new(Connection), nil
}

func (c *connectionHolder) ReconnectIfUnhealthy() error { _ = "STUB: not implemented"; return nil }

func (c *connectionHolder) connect() error { _ = "STUB: not implemented"; return nil }

// Goal is to lazily restore the connection so this needs to be buffered to avoid blocking on asynchronous amqp errors.
// Also re-create this channel each time because apparently the amqp library can close it

func (c *connectionHolder) Close() error { _ = "STUB: not implemented"; return nil }

func (c *connectionHolder) isConnected() bool { _ = "STUB: not implemented"; return false }

func (c *connectionHolder) Channel() (Channel, error) {
	_ = "STUB: not implemented"
	return *new(Channel), nil
}

func (c *connectionHolder) IsClosed() bool { _ = "STUB: not implemented"; return false }

func (c *connectionHolder) NotifyClose(receiver chan *amqp.Error) chan *amqp.Error {
	_ = "STUB: not implemented"
	return nil
}

func (c *channelHolder) Confirm(noWait bool) error { _ = "STUB: not implemented"; return nil }

func (c *channelHolder) PublishWithDeferredConfirmWithContext(ctx context.Context, exchange, key string, mandatory, immediate bool, msg amqp.Publishing) (DeferredConfirmation, error) {
	_ = "STUB: not implemented"
	return *new(DeferredConfirmation), nil
}

func (c *channelHolder) IsClosed() bool { _ = "STUB: not implemented"; return false }

func (c *channelHolder) Close() error { _ = "STUB: not implemented"; return nil }

func (d *deferredConfirmationHolder) Done() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (d *deferredConfirmationHolder) Acked() bool { _ = "STUB: not implemented"; return false }
