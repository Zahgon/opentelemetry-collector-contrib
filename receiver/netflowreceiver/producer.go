// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package netflowreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/netflowreceiver"

import (
	"github.com/netsampler/goflow2/v2/producer"
	"go.opentelemetry.io/collector/consumer"
	"go.uber.org/zap"
)

// otelLogsProducerWrapper is a wrapper around a producer.ProducerInterface that sends the messages to a log consumer
type otelLogsProducerWrapper struct {
	wrapped     producer.ProducerInterface
	logConsumer consumer.Logs
	logger      *zap.Logger
	sendRaw     bool
}

// Produce converts the message into a list log records and sends them to log consumer
func (o *otelLogsProducerWrapper) Produce(msg any, args *producer.ProduceArgs) ([]producer.ProducerMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// First we let the proto producer parse the message
// All the netflow protocol and structure is handled by the proto producer

// Create the otel log structure to hold our messages

// A single netflow packet can contain multiple flow messages

// Parse the message and add the attributes to the log record

func (o *otelLogsProducerWrapper) Close() { _ = "STUB: not implemented"; return }

func (o *otelLogsProducerWrapper) Commit(flowMessageSet []producer.ProducerMessage) {
	_ = "STUB: not implemented"
	return
}

func newOtelLogsProducer(wrapped producer.ProducerInterface, logConsumer consumer.Logs, logger *zap.Logger, sendRaw bool) producer.ProducerInterface {
	_ = "STUB: not implemented"
	return *new(producer.ProducerInterface)
}
