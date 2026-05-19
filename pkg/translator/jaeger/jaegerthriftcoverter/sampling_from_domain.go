// Copyright The OpenTelemetry Authors
// Copyright (c) 2018 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package jaeger // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/jaeger/jaegerthriftcoverter"

import (
	"github.com/jaegertracing/jaeger-idl/proto-gen/api_v2"
	"github.com/jaegertracing/jaeger-idl/thrift-gen/sampling"
)

// ConvertSamplingResponseFromDomain converts proto sampling response to its thrift representation.
func ConvertSamplingResponseFromDomain(r *api_v2.SamplingStrategyResponse) (*sampling.SamplingStrategyResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertProbabilisticFromDomain(s *api_v2.ProbabilisticSamplingStrategy) *sampling.ProbabilisticSamplingStrategy {
	_ = "STUB: not implemented"
	return nil
}

func convertRateLimitingFromDomain(s *api_v2.RateLimitingSamplingStrategy) (*sampling.RateLimitingSamplingStrategy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertPerOperationFromDomain(s *api_v2.PerOperationSamplingStrategies) *sampling.PerOperationSamplingStrategies {
	_ = "STUB: not implemented"
	return nil
}

// Default to empty array so that json.Marshal returns [] instead of null (Issue #3891).

func convertOperationFromDomain(s *api_v2.OperationSamplingStrategy) *sampling.OperationSamplingStrategy {
	_ = "STUB: not implemented"
	return nil
}

func convertStrategyTypeFromDomain(s api_v2.SamplingStrategyType) (sampling.SamplingStrategyType, error) {
	_ = "STUB: not implemented"
	return *new(sampling.SamplingStrategyType), nil
}
