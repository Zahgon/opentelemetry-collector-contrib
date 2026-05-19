// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package transformprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottldatapoint"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottllog"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlmetric"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlprofile"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspan"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspanevent"
)

func DefaultLogFunctions() []ottl.Factory[*ottllog.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated: [v0.152.0] Use DefaultLogFunctions.
func DefaultLogFunctionsNew() []ottl.Factory[*ottllog.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func DefaultMetricFunctions() []ottl.Factory[*ottlmetric.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated: [v0.152.0] Use DefaultMetricFunctions.
func DefaultMetricFunctionsNew() []ottl.Factory[*ottlmetric.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func DefaultDataPointFunctions() []ottl.Factory[*ottldatapoint.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated: [v0.152.0] Use DefaultDataPointFunctions.
func DefaultDataPointFunctionsNew() []ottl.Factory[*ottldatapoint.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func DefaultSpanFunctions() []ottl.Factory[*ottlspan.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated: [v0.152.0] Use DefaultSpanFunctions.
func DefaultSpanFunctionsNew() []ottl.Factory[*ottlspan.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func DefaultSpanEventFunctions() []ottl.Factory[*ottlspanevent.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated: [v0.152.0] Use DefaultSpanEventFunctions.
func DefaultSpanEventFunctionsNew() []ottl.Factory[*ottlspanevent.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func DefaultProfileFunctions() []ottl.Factory[*ottlprofile.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated: [v0.152.0] Use DefaultProfileFunctions.
func DefaultProfileFunctionsNew() []ottl.Factory[*ottlprofile.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func defaultLogFunctionsMap() map[string]ottl.Factory[*ottllog.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func defaultMetricFunctionsMap() map[string]ottl.Factory[*ottlmetric.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func defaultDataPointFunctionsMap() map[string]ottl.Factory[*ottldatapoint.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func defaultSpanFunctionsMap() map[string]ottl.Factory[*ottlspan.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func defaultSpanEventFunctionsMap() map[string]ottl.Factory[*ottlspanevent.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func defaultProfileFunctionsMap() map[string]ottl.Factory[*ottlprofile.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func mergeFunctionsToMap[K any](functionMap map[string]ottl.Factory[K], functions []ottl.Factory[K]) map[string]ottl.Factory[K] {
	_ = "STUB: not implemented"
	return nil
}
