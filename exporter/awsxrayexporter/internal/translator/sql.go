// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translator // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter/internal/translator"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"

	awsxray "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray"
)

func makeSQL(span ptrace.Span, attributes map[string]pcommon.Value) (map[string]pcommon.Value, *awsxray.SQLData) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Either no DB attributes or this is not an SQL DB.

// Despite what the X-Ray documents say, having the DB connection string
// set as the URL value of the segment is not useful. So let's use the
// current span name instead

// Let's keep the original format for connection_string

func isSQL(system string) bool { _ = "STUB: not implemented"; return false }
