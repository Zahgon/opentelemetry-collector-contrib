// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awscloudwatchmetricstreamsencodingextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/awscloudwatchmetricstreamsencodingextension"

import (
	"errors"
	"io"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding"
)

var errInvalidUvarint = errors.New("invalid OTLP message length: failed to decode varint")

type formatOpenTelemetry10Unmarshaler struct {
	buildInfo component.BuildInfo
}

var _ pmetric.Unmarshaler = (*formatOpenTelemetry10Unmarshaler)(nil)

func (f *formatOpenTelemetry10Unmarshaler) UnmarshalMetrics(record []byte) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	// Decode as a stream but flush all at once using flush options
	return *new(pmetric.Metrics), nil
}

//nolint:errorlint

// EOF indicates no metrics were found, return any metrics that's available

func (f *formatOpenTelemetry10Unmarshaler) NewMetricsDecoder(reader io.Reader, options ...encoding.DecoderOption) (encoding.MetricsDecoder, error) {
	_ = "STUB: not implemented"
	return *new(encoding.MetricsDecoder), nil
}

// reached EOF

// Read next toRead bytes to get the length of the next OTLP metric message

// Skip bytesRead

// Reuse buffer, grow only if needed

// Read the OTLP metric message

// unmarshal metric

// add scope name and build info version to the resource metrics
