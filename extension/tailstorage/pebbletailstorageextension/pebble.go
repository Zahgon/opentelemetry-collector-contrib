// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pebbletailstorageextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/tailstorage/pebbletailstorageextension"

import (
	"github.com/cockroachdb/pebble/v2"
	"go.uber.org/zap"
)

func newPebbleDB(storageDir string, logger *zap.Logger) (*pebble.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: support persistence across restarts as storage schema matures.
// Meanwhile, error if DB already exists to prevent users from relying on persistence.

func traceKeyComparer() *pebble.Comparer { _ = "STUB: not implemented"; return nil }

// Since trace ID is fixed sized, and trace ID separator is a valid trace ID byte itself,
// split on fixed length instead to avoid corruption.

// BUG: key len should always be >= 17
