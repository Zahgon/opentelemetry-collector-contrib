// Copyright The OpenTelemetry Authors
// Copyright (c) 2018 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package jaeger // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/jaeger/jaegerthriftcoverter"

import (
	"github.com/jaegertracing/jaeger-idl/proto-gen/api_v2"
	"github.com/jaegertracing/jaeger-idl/thrift-gen/sampling"
)

// ConvertSamplingResponseToDomain converts thrift sampling response to its proto representation.
func ConvertSamplingResponseToDomain(r *sampling.SamplingStrategyResponse) (*api_v2.SamplingStrategyResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertRateLimitingToDomain(s *sampling.RateLimitingSamplingStrategy) *api_v2.RateLimitingSamplingStrategy {
	_ = "STUB: not implemented"
	return nil
}

func convertProbabilisticToDomain(s *sampling.ProbabilisticSamplingStrategy) *api_v2.ProbabilisticSamplingStrategy {
	_ = "STUB: not implemented"
	return nil
}

func convertPerOperationToDomain(s *sampling.PerOperationSamplingStrategies) *api_v2.PerOperationSamplingStrategies {
	_ = "STUB: not implemented"
	return nil
}

func convertStrategyTypeToDomain(t sampling.SamplingStrategyType) (api_v2.SamplingStrategyType, error) {
	_ = "STUB: not implemented"
	return *new(api_v2.SamplingStrategyType), nil
}
