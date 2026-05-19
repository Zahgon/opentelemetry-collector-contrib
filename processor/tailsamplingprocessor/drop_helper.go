// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package tailsamplingprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor"

import (
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor/pkg/samplingpolicy"
)

func getNewDropPolicy(settings component.TelemetrySettings, config *DropCfg, policyExtensions map[string]samplingpolicy.Extension) (samplingpolicy.Evaluator, error) {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Evaluator), nil
}

// Return instance of and sub-policy
func getDropSubPolicyEvaluator(settings component.TelemetrySettings, cfg *AndSubPolicyCfg, policyExtensions map[string]samplingpolicy.Extension) (samplingpolicy.Evaluator, error) {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Evaluator), nil
}
