// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translator // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsxrayreceiver/internal/translator"

import (
	"go.opentelemetry.io/collector/pdata/ptrace"

	awsxray "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray"
)

func addHTTP(seg *awsxray.Segment, span ptrace.Span) { _ = "STUB: not implemented"; return }

// https://docs.aws.amazon.com/xray/latest/devguide/xray-api-segmentdocuments.html#api-segmentdocuments-http

// since the ClientIP is not nil, this means that this segment is generated
// by a server serving an incoming request

// in X-Ray exporter, the segment status is set:
// first via the span attribute, string(conventions.HTTPStatusCodeKey)
// then the span status. Since we are also setting the span attribute
// below, the span status code here will not be actually used
