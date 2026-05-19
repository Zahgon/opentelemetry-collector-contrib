// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translator // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsxrayreceiver/internal/translator"

import (
	"go.opentelemetry.io/collector/pdata/ptrace"

	awsxray "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray"
)

const (
	validAWSNamespace    = "aws"
	validRemoteNamespace = "remote"
	validLocalNamespace  = "local"
)

func addNameAndNamespace(seg *awsxray.Segment, span ptrace.Span) error {
	_ = "STUB: not implemented"
	// https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/c615d2db351929b99e46f7b427f39c12afe15b54/exporter/awsxrayexporter/translator/segment.go#L160
	return nil
}

// `ClientIP` is an optional field, we only attempt to use it to set
// a more specific spanKind if it exists.

// The `ClientIP` is not nil, it implies that this segment is generated
// by a server serving an incoming request

// seg is a subsegment

// https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/c615d2db351929b99e46f7b427f39c12afe15b54/exporter/awsxrayexporter/translator/segment.go#L163
// https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/trace/api.md#spankind

// https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/c615d2db351929b99e46f7b427f39c12afe15b54/exporter/awsxrayexporter/translator/segment.go#L116

// no op
