// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package filterexpr // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filterexpr"

import (
	"sync"

	"github.com/expr-lang/expr/vm"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

var vmPool = sync.Pool{
	New: func() any {
		return &vm.VM{}
	},
}

type Matcher struct {
	program *vm.Program
}

type env struct {
	MetricName string
	MetricType string
	attributes pcommon.Map
}

func (e *env) HasLabel(key string) bool { _ = "STUB: not implemented"; return false }

func (e *env) Label(key string) string { _ = "STUB: not implemented"; return "" }

func NewMatcher(expression string) (*Matcher, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *Matcher) MatchMetric(metric pmetric.Metric) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

//exhaustive:enforce

func (m *Matcher) matchGauge(metricName string, gauge pmetric.Gauge, vm *vm.VM) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (m *Matcher) matchSum(metricName string, sum pmetric.Sum, vm *vm.VM) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (m *Matcher) matchHistogram(metricName string, histogram pmetric.Histogram, vm *vm.VM) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (m *Matcher) matchExponentialHistogram(metricName string, eh pmetric.ExponentialHistogram, vm *vm.VM) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (m *Matcher) matchSummary(metricName string, summary pmetric.Summary, vm *vm.VM) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (m *Matcher) matchEnv(metricName string, metricType pmetric.MetricType, attributes pcommon.Map, vm *vm.VM) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (m *Matcher) match(env env, vm *vm.VM) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
