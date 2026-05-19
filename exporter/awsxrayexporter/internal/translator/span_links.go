// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translator // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter/internal/translator"

import (
	"go.opentelemetry.io/collector/pdata/ptrace"

	awsxray "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray"
)

func makeSpanLinks(links ptrace.SpanLinkSlice, skipTimestampValidation bool) ([]awsxray.SpanLinkData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
