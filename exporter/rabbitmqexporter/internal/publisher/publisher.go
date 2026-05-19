// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package publisher // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/rabbitmqexporter/internal/publisher"

import (
	"context"
	"time"

	"go.uber.org/zap"

	otelrabbitmq "github.com/open-telemetry/opentelemetry-collector-contrib/internal/rabbitmq"
)

type DialConfig struct {
	otelrabbitmq.DialConfig
	Durable                    bool
	PublishConfirmationTimeout time.Duration
}

type Message struct {
	Exchange   string
	RoutingKey string
	Body       []byte
}

func NewConnection(logger *zap.Logger, client otelrabbitmq.AmqpClient, config DialConfig) (Publisher, error) {
	_ = "STUB: not implemented"
	return *new(Publisher), nil
}

type Publisher interface {
	Publish(ctx context.Context, message Message) error
	Close() error
}

type publisher struct {
	logger     *zap.Logger
	client     otelrabbitmq.AmqpClient
	config     DialConfig
	connection otelrabbitmq.Connection
}

func (p *publisher) Publish(ctx context.Context, message Message) error {
	_ = "STUB: not implemented"
	return nil
}

// Create a new amqp channel for publishing messages and request that the broker confirms delivery.
// This could later be optimized to re-use channels which avoids repeated network calls to create and close them.
// Concurrency-control through something like a resource pool would be necessary since aqmp channels are not thread safe.

// Send the message

// Wait for async confirmation of the message

func (p *publisher) Close() error { _ = "STUB: not implemented"; return nil }
