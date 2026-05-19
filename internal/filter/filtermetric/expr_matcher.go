// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package filtermetric // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filtermetric"

import (
	"context"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filterexpr"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlmetric"
)

type exprMatcher struct {
	matchers []*filterexpr.Matcher
}

func newExprMatcher(expressions []string) (*exprMatcher, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *exprMatcher) Eval(_ context.Context, tCtx *ottlmetric.TransformContext) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
