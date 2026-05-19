// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package tailsamplingprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor"

import (
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor/pkg/samplingpolicy"
)

func getNewCompositePolicy(settings component.TelemetrySettings, config *CompositeCfg, policyExtensions map[string]samplingpolicy.Extension) (samplingpolicy.Evaluator, error) {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Evaluator), nil
}

// Apply rate allocations to the sub-policies
func getRateAllocationMap(config *CompositeCfg) map[string]float64 {
	_ = "STUB: not implemented"
	return nil
}

// Default SPS determined by equally diving number of sub policies

// Return instance of composite sub-policy
func getCompositeSubPolicyEvaluator(settings component.TelemetrySettings, cfg *CompositeSubPolicyCfg, policyExtensions map[string]samplingpolicy.Extension) (samplingpolicy.Evaluator, error) {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Evaluator), nil
}
