// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awscloudwatchmetricstreamsencodingextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/awscloudwatchmetricstreamsencodingextension"

import (
	"context"
	"io"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding"
)

var (
	_ encoding.MetricsUnmarshalerExtension = (*encodingExtension)(nil)
	_ encoding.MetricsDecoderExtension     = (*encodingExtension)(nil)
)

type encodingExtension struct {
	unmarshaler pmetric.Unmarshaler
	format      string
}

func newExtension(cfg *Config, settings extension.Settings) (*encodingExtension, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Format will have been validated by Config.Validate,
// so we'll only get here if we haven't handled a valid
// format.

func (*encodingExtension) Start(_ context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (*encodingExtension) Shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (e *encodingExtension) UnmarshalMetrics(record []byte) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// NewMetricsDecoder returns a MetricsDecoder if the underlying unmarshaler supports streaming.
// Caller must perform any decompression before passing the reader to the decoder.
// Implementations must utilize derived buffered readers as is.
func (e *encodingExtension) NewMetricsDecoder(reader io.Reader, options ...encoding.DecoderOption) (encoding.MetricsDecoder, error) {
	_ = "STUB: not implemented"
	return *new(encoding.MetricsDecoder), nil
}

// metricsUnmarshal is an interface that's expected to be implemented by metrics format implementations.
type streamUnmarshal interface {
	pmetric.Unmarshaler
	NewMetricsDecoder(reader io.Reader, options ...encoding.DecoderOption) (encoding.MetricsDecoder, error)
}
