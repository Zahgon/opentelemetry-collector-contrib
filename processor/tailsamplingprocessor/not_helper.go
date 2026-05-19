// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package tailsamplingprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor"

import (
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor/pkg/samplingpolicy"
)

func getNewNotPolicy(settings component.TelemetrySettings, config *NotCfg, policyExtensions map[string]samplingpolicy.Extension) (samplingpolicy.Evaluator, error) {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Evaluator), nil
}

// Return instance of not sub-policy
func getNotSubPolicyEvaluator(settings component.TelemetrySettings, cfg *NotSubPolicyCfg, policyExtensions map[string]samplingpolicy.Extension) (samplingpolicy.Evaluator, error) {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Evaluator), nil
}
