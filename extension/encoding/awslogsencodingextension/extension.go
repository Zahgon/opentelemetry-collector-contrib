// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awslogsencodingextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/awslogsencodingextension"

import (
	"context"
	"io"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/extension/extensioncapabilities"
	"go.opentelemetry.io/collector/featuregate"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/awslogsencodingextension/internal/constants"
	awsunmarshaler "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/awslogsencodingextension/internal/unmarshaler"
	subscriptionfilter "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/awslogsencodingextension/internal/unmarshaler/subscription-filter"
)

const (
	gzipEncoding    = "gzip"
	bytesEncoding   = "bytes"
	parquetEncoding = "parquet"
)

var (
	_ encoding.LogsUnmarshalerExtension = (*encodingExtension)(nil)
	_ encoding.LogsDecoderExtension     = (*encodingExtension)(nil)
	_ extensioncapabilities.Dependent   = (*encodingExtension)(nil)
)

var (
	vpcFlowStartISO8601FormatFeatureGate    *featuregate.Gate
	cloudTrailUserIdentityPrefixFeatureGate *featuregate.Gate
)

func init() {
	vpcFlowStartISO8601FormatFeatureGate = featuregate.GlobalRegistry().MustRegister(
		constants.VPCFlowStartISO8601FormatID,
		featuregate.StageAlpha,
		featuregate.WithRegisterDescription("When enabled, aws.vpc.flow.start field will be formatted as ISO-8601 string instead of seconds since epoch integer."),
		featuregate.WithRegisterReferenceURL("https://github.com/open-telemetry/opentelemetry-collector-contrib/pull/43390"),
	)

	cloudTrailUserIdentityPrefixFeatureGate = featuregate.GlobalRegistry().MustRegister(
		constants.CloudTrailEnableUserIdentityPrefixID,
		featuregate.StageAlpha,
		featuregate.WithRegisterDescription("When enabled, CloudTrail log userIdentity attributes will use 'aws.user_identity' prefix. This helps to preserve the attribute origin."),
		featuregate.WithRegisterReferenceURL("https://github.com/open-telemetry/opentelemetry-collector-contrib/pull/45459"))
}

type encodingExtension struct {
	cfg *Config

	unmarshaler             awsunmarshaler.AWSUnmarshaler
	subscriptionFilter      *subscriptionfilter.SubscriptionFilterUnmarshaler
	format                  string
	gzipPool                sync.Pool
	logger                  *zap.Logger
	warnGzipDeprecationOnce sync.Once
	selfID                  component.ID
}

func newExtension(cfg *Config, settings extension.Settings) (*encodingExtension, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Format will have been validated by Config.Validate,
// so we'll only get here if we haven't handled a valid
// format.

// Dependencies declares the inner encoding extensions referenced by the
// CloudWatch routing config so the framework starts them before this one.
func (e *encodingExtension) Dependencies() []component.ID { _ = "STUB: not implemented"; return nil }

func (e *encodingExtension) Start(_ context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (*encodingExtension) Shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (e *encodingExtension) UnmarshalLogs(buf []byte) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

// NewLogsDecoder returns a LogsDecoder if the underlying unmarshaler supports streaming.
// Caller must perform any decompression before passing the reader to the decoder.
// Implementations must utilize derived buffered readers as is.
func (e *encodingExtension) NewLogsDecoder(reader io.Reader, options ...encoding.DecoderOption) (encoding.LogsDecoder, error) {
	_ = "STUB: not implemented"
	return *new(encoding.LogsDecoder), nil
}

func (e *encodingExtension) getGzipReader(buf []byte) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

// isGzipData checks if the buffer contains gzip-compressed data by examining magic bytes
func isGzipData(buf []byte) bool { _ = "STUB: not implemented"; return false }

// getReaderForData returns the appropriate reader and encoding type based on data format
func (e *encodingExtension) getReaderForData(buf []byte) (string, io.Reader, error) {
	_ = "STUB: not implemented"
	return "", *new(io.Reader), nil
}

func (e *encodingExtension) getReaderFromFormat(buf []byte) (string, io.Reader, error) {
	_ = "STUB: not implemented"
	return "", *new(io.Reader), nil
}

// should not be possible

// should not be possible
