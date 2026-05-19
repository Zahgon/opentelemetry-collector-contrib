// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package customottl // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/signaltometricsconnector/internal/customottl"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottldatapoint"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottllog"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlprofile"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspan"
)

func SpanFuncs() map[string]ottl.Factory[*ottlspan.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func DatapointFuncs() map[string]ottl.Factory[*ottldatapoint.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func LogFuncs() map[string]ottl.Factory[*ottllog.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func ProfileFuncs() map[string]ottl.Factory[*ottlprofile.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func commonFuncs[K any]() map[string]ottl.Factory[K] { _ = "STUB: not implemented"; return nil }
