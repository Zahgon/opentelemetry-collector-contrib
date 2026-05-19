// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kafkaexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/kafkaexporter"

import (
	"github.com/twmb/franz-go/pkg/kgo"
	"go.opentelemetry.io/collector/component"
)

// RecordPartitionerExtension is implemented by extensions that supply a custom Kafka record
// partitioner for use with the kafka exporter.
type RecordPartitionerExtension interface {
	component.Component

	GetPartitioner() kgo.Partitioner
}

func buildPartitionerOpt(cfg RecordPartitionerConfig, host component.Host) (kgo.Opt, error) {
	_ = "STUB: not implemented"
	return *new(kgo.Opt), nil
}

// in practice, this shouldn't happen.
// The config validation should catch the case where no partitioner is set.
