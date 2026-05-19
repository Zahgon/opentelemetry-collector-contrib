// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sampling // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor/internal/sampling"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor/pkg/samplingpolicy"
)

type statusCodeFilter struct {
	logger      *zap.Logger
	statusCodes []ptrace.StatusCode
}

var _ samplingpolicy.Evaluator = (*statusCodeFilter)(nil)

// NewStatusCodeFilter creates a policy evaluator that samples all traces with
// a given status code.
func NewStatusCodeFilter(settings component.TelemetrySettings, statusCodeString []string) (samplingpolicy.Evaluator, error) {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Evaluator), nil
}

// Evaluate looks at the trace data and returns a corresponding SamplingDecision.
func (r *statusCodeFilter) Evaluate(_ context.Context, _ pcommon.TraceID, trace *samplingpolicy.TraceData) (samplingpolicy.Decision, error) {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Decision), nil
}

func (*statusCodeFilter) IsStateful() bool { _ = "STUB: not implemented"; return false }
