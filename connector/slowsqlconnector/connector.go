// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package slowsqlconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/slowsqlconnector"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	conventions "go.opentelemetry.io/otel/semconv/v1.40.0"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/pdatautil"
)

const (
	serviceNameKey        = string(conventions.ServiceNameKey)
	dbSystemKey           = string(conventions.DBSystemNameKey)
	statementExecDuration = "db.client.operation.duration" // OpenTelemetry non-standard constant.
	spanKindKey           = "span.kind"                    // OpenTelemetry non-standard constant.
	spanNameKey           = "span.name"                    // OpenTelemetry non-standard constant.
	statusCodeKey         = "status.code"                  // OpenTelemetry non-standard constant.
	dbStatementKey        = "db.statement"                 // OpenTelemetry non-standard constant.
)

func newDimensions(cfgDims []Dimension) []pdatautil.Dimension {
	_ = "STUB: not implemented"
	return nil
}

func findAttributeValue(key string, attributes ...pcommon.Map) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}
