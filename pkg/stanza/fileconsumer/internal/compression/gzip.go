// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package compression // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/compression"

import (
	"os"

	"go.uber.org/zap"
)

const gzipHeader = "\x1f\x8b" // RFC 1952 magic bytes

// IsGzipFile checks if a file is of gzip type by reading its header
func IsGzipFile(f *os.File, logger *zap.Logger) bool { _ = "STUB: not implemented"; return false }

// empty or too short file
