// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package exceptionsconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/exceptionsconnector"

import (
	conventions "go.opentelemetry.io/otel/semconv/v1.40.0"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/pdatautil"
)

const (
	serviceNameKey         = string(conventions.ServiceNameKey)
	exceptionTypeKey       = string(conventions.ExceptionTypeKey)
	exceptionMessageKey    = string(conventions.ExceptionMessageKey)
	exceptionStacktraceKey = string(conventions.ExceptionStacktraceKey)
	// TODO(marctc): formalize these constants in the OpenTelemetry specification.
	spanKindKey   = "span.kind"   // OpenTelemetry non-standard constant.
	spanNameKey   = "span.name"   // OpenTelemetry non-standard constant.
	statusCodeKey = "status.code" // OpenTelemetry non-standard constant.
	eventNameExc  = "exception"   // OpenTelemetry non-standard constant.
)

func newDimensions(cfgDims []Dimension) []pdatautil.Dimension {
	_ = "STUB: not implemented"
	return nil
}
