// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package traces // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor/internal/traces"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspan"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspanevent"
)

func SpanFunctions() map[string]ottl.Factory[*ottlspan.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func SpanEventFunctions() map[string]ottl.Factory[*ottlspanevent.TransformContext] {
	_ = "STUB: not implemented"
	// No trace-only functions yet.
	return nil
}
