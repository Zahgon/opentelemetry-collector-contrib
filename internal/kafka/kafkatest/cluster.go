// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kafkatest // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/kafka/kafkatest"

import (
	"testing"

	"github.com/twmb/franz-go/pkg/kfake"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/kafka/configkafka"
)

// NewCluster returns a fake Kafka cluster and configkafka.ClientConfig
// with the default configuration, and brokers set to the cluster addresses.
func NewCluster(tb testing.TB, opts ...kfake.Opt) (*kfake.Cluster, configkafka.ClientConfig) {
	_ = "STUB: not implemented"
	return nil, *new(configkafka.ClientConfig)
}

// Cap the protocol version to one that uses topic names rather than topic IDs
// in Produce requests, so tests that inspect produced records by topic name
// continue to work against kfake.
