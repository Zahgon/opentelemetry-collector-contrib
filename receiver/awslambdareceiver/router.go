// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awslambdareceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awslambdareceiver"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding"
)

// routedEncoding pairs an S3Encoding with its pre-split path pattern parts.
type routedEncoding struct {
	enc          S3Encoding
	patternParts []string
}

// logsDecoderRouter routes S3 object keys to the appropriate LogsDecoderFactory
// based on path pattern matching with wildcards.
//
// Encodings must be pre-sorted by specificity (most specific first, catch-all last).
// Use S3Config.sortedEncodings() to obtain a correctly sorted slice.
type logsDecoderRouter struct {
	encodings []routedEncoding
	decoders  map[string]encoding.LogsDecoderFactory
}

// newLogsDecoderRouter creates a new S3 logs decoder router.
//   - encodings: pre-sorted S3Encoding slice.
//   - decoders: maps encoding.Name => LogsDecoderFactory for every entry (including raw-passthrough ones).
//
// Path patterns are split on "/" at construction time so that GetDecoder only
// splits the object key once per S3 event rather than once per pattern.
func newLogsDecoderRouter(
	encodings []S3Encoding,
	decoders map[string]encoding.LogsDecoderFactory,
) *logsDecoderRouter {
	_ = "STUB: not implemented"
	return nil
}

// GetDecoder returns the LogsDecoderFactory and encoding name for the given S3 object key.
// It splits the object key once, then iterates encodings in order (most specific first)
// and returns on the first match. Returns an error if no encoding matches.
func (r *logsDecoderRouter) GetDecoder(objectKey string) (encoding.LogsDecoderFactory, string, error) {
	_ = "STUB: not implemented"
	return *new(encoding.LogsDecoderFactory), "", nil
}
