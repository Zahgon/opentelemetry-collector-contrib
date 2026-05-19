// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package customottl // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/signaltometricsconnector/internal/customottl"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspan"
)

func NewAdjustedCountFactory() ottl.Factory[*ottlspan.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func createAdjustedCountFunction(ottl.FunctionContext, ottl.Arguments) (ottl.ExprFunc[*ottlspan.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func adjustedCount() (ottl.ExprFunc[*ottlspan.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If otel trace state is missing, default to 1

// For non-probabilistic sampler OR always sampling threshold, default to 1
