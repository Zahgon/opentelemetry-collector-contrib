// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package checkpoint // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/checkpoint"

import (
	"context"

	"go.opentelemetry.io/collector/extension/xextension/storage"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/reader"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
)

const knownFilesKey = "knownFiles"

// Save syncs the most recent set of files to the database
// Uses protobuf encoding if the feature gate is enabled, otherwise uses JSON
func Save(ctx context.Context, persister operator.Persister, rmds []*reader.Metadata) error {
	_ = "STUB: not implemented"
	return nil
}

func SaveKey(ctx context.Context, persister operator.Persister, rmds []*reader.Metadata, key string, ops ...*storage.Operation) error {
	_ = "STUB: not implemented"
	// Use protobuf if feature gate is enabled
	return nil
}

// Otherwise use JSON (default)

// Encode the number of known files

// Encode each known file

// Load loads the most recent set of files from the database
// Tries protobuf first for backward compatibility, falls back to JSON if protobuf fails
func Load(ctx context.Context, persister operator.Persister, logger *zap.Logger) ([]*reader.Metadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func LoadKey(ctx context.Context, persister operator.Persister, key string, logger *zap.Logger) ([]*reader.Metadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Try protobuf first (for backward compatibility with existing protobuf checkpoints)
// This allows seamless migration even when the feature gate is disabled

// Fall back to JSON if protobuf fails

// Decode the number of entries

// Decode each of the known files

// Migrate readers that used FileAttributes.HeaderAttributes
// This block can be removed in a future release, tentatively v0.90.0

// This reader won't be used for anything other than metadata reference, so just wrap the metadata
