// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kafkaclient // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/kafkaexporter/internal/kafkaclient"

import (
	"context"

	"github.com/twmb/franz-go/pkg/kgo"
)

// metadataToHeaders converts context metadata into a kgo.RecordHeader slice.
func metadataToHeaders(ctx context.Context, keys []string) []kgo.RecordHeader {
	_ = "STUB: not implemented"
	return nil
}
