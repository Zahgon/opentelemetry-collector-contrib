// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package filterottl // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filterottl"

import (
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottldatapoint"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottllog"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlmetric"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlprofile"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlresource"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlscope"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspan"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspanevent"
)

func StandardSpanFuncs() map[string]ottl.Factory[*ottlspan.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func StandardSpanEventFuncs() map[string]ottl.Factory[*ottlspanevent.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func StandardMetricFuncs() map[string]ottl.Factory[*ottlmetric.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func StandardDataPointFuncs() map[string]ottl.Factory[*ottldatapoint.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func StandardScopeFuncs() map[string]ottl.Factory[*ottlscope.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func StandardLogFuncs() map[string]ottl.Factory[*ottllog.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func StandardProfileFuncs() map[string]ottl.Factory[*ottlprofile.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func StandardResourceFuncs() map[string]ottl.Factory[*ottlresource.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

type hasAttributeOnDatapointArguments struct {
	Key         string
	ExpectedVal string
}

func newHasAttributeOnDatapointFactory() ottl.Factory[*ottlmetric.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func createHasAttributeOnDatapointFunction(_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[*ottlmetric.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func hasAttributeOnDatapoint(key, expectedVal string) (ottl.ExprFunc[*ottlmetric.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type hasAttributeKeyOnDatapointArguments struct {
	Key string
}

func newHasAttributeKeyOnDatapointFactory() ottl.Factory[*ottlmetric.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func createHasAttributeKeyOnDatapointFunction(_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[*ottlmetric.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func hasAttributeKeyOnDatapoint(key string) (ottl.ExprFunc[*ottlmetric.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func checkDataPoints(tCtx *ottlmetric.TransformContext, key string, expectedVal *string) (any, error) {
	_ = "STUB: not implemented"
	return *

	//exhaustive:enforce
	new(any), nil
}

func checkNumberDataPointSlice(dps pmetric.NumberDataPointSlice, key string, expectedVal *string) bool {
	_ = "STUB: not implemented"
	return false
}

func checkHistogramDataPointSlice(dps pmetric.HistogramDataPointSlice, key string, expectedVal *string) bool {
	_ = "STUB: not implemented"
	return false
}

func checkExponentialHistogramDataPointSlice(dps pmetric.ExponentialHistogramDataPointSlice, key string, expectedVal *string) bool {
	_ = "STUB: not implemented"
	return false
}

func checkSummaryDataPointSlice(dps pmetric.SummaryDataPointSlice, key string, expectedVal *string) bool {
	_ = "STUB: not implemented"
	return false
}
