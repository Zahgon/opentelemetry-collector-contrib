// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package checkpoint // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/checkpoint"

import (
	"context"

	"go.opentelemetry.io/collector/extension/xextension/storage"

	pb "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/checkpoint/proto"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/reader"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
)

// tryLoadProtobuf attempts to load checkpoint data using protobuf encoding
// Returns error if the data is not valid protobuf
func tryLoadProtobuf(encoded []byte) ([]*reader.Metadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func saveKeyProto(ctx context.Context, persister operator.Persister, rmds []*reader.Metadata, key string, ops ...*storage.Operation) error {
	_ = "STUB: not implemented"
	return nil
}

// metadataToPb converts reader.Metadata to protobuf Metadata
func metadataToPb(rmd *reader.Metadata) (*pb.Metadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Convert Fingerprint — Bytes() already returns a copy

// Convert FileAttributes - encode the entire map as JSON bytes

// Convert FlushState

// Only set LastDataChangeUnixNano if the time is not zero

// Convert TokenLenState

// pbToMetadata converts protobuf Metadata to reader.Metadata
func pbToMetadata(pbMeta *pb.Metadata) (*reader.Metadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Convert Fingerprint

// Convert FileAttributes - decode from JSON bytes

// Convert FlushState

// Convert TokenLenState
