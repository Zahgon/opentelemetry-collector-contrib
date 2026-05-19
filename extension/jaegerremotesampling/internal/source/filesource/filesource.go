// Copyright The OpenTelemetry Authors
// Copyright (c) 2018 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package filesource // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/jaegerremotesampling/internal/source/filesource"

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/jaegertracing/jaeger-idl/proto-gen/api_v2"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/jaegerremotesampling/internal/source"
)

// null represents "null" JSON value and
// it un-marshals to nil pointer.
var nullJSON = []byte("null")

type samplingProvider struct {
	logger *zap.Logger

	storedStrategies atomic.Value // holds *storedStrategies

	cancelFunc context.CancelFunc

	options Options

	wg sync.WaitGroup
}

type storedStrategies struct {
	defaultStrategy   *api_v2.SamplingStrategyResponse
	serviceStrategies map[string]*api_v2.SamplingStrategyResponse
}

type strategyLoader func() ([]byte, error)

// NewFileSource creates a strategy store that holds static sampling strategies.
func NewFileSource(options Options, logger *zap.Logger) (source.Source, error) {
	_ = "STUB: not implemented"
	return *new(source.Source), nil
}

// GetSamplingStrategy implements StrategyStore#GetSamplingStrategy.
func (h *samplingProvider) GetSamplingStrategy(_ context.Context, serviceName string) (*api_v2.SamplingStrategyResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Close stops updating the strategies
func (h *samplingProvider) Close() error { _ = "STUB: not implemented"; return nil }

func (h *samplingProvider) downloadSamplingStrategies(samplingURL string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isURL(str string) bool { _ = "STUB: not implemented"; return false }

func (h *samplingProvider) samplingStrategyLoader(strategiesFile string) strategyLoader {
	_ = "STUB: not implemented"
	return *new(strategyLoader)
}

func (h *samplingProvider) autoUpdateStrategies(ctx context.Context, interval time.Duration, loader strategyLoader) {
	_ = "STUB: not implemented"
	return
}

func (h *samplingProvider) reloadSamplingStrategy(loadFn strategyLoader, lastValue string) string {
	_ = "STUB: not implemented"
	return ""
}

func (h *samplingProvider) updateSamplingStrategy(dataBytes []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO good candidate for a global util function
func loadStrategies(loadFn strategyLoader) (*strategies, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *samplingProvider) parseStrategiesDeprecated(strategies *strategies) {
	_ = "STUB: not implemented"
	return
}

// Merge with the default operation strategies, because only merging with
// the default strategy has no effect on service strategies (the default strategy
// is not merged with and only used as a fallback).

// Service has no per-operation strategies, so just reference the default settings and change default samplingRate.

func (h *samplingProvider) parseStrategies(strategies *strategies) {
	_ = "STUB: not implemented"
	return
}

// Config for this service may not have per-operation strategies,
// but if the default strategy has them they should still apply.

// Default strategy doens't have them either, nothing to do.

// Service does not have its own per-operation rules, so copy (by value) from the default strategy.

// If the service's own default is probabilistic, then its sampling rate should take precedence.

// If the service did have its own per-operation strategies, then merge them with the default ones.

// mergePerOperationSamplingStrategies merges two operation strategies a and b, where a takes precedence over b.
func mergePerOperationSamplingStrategies(
	a, b []*api_v2.OperationSamplingStrategy,
) []*api_v2.OperationSamplingStrategy {
	_ = "STUB: not implemented"
	return nil
}

func (h *samplingProvider) parseServiceStrategies(strategy *serviceStrategy) *api_v2.SamplingStrategyResponse {
	_ = "STUB: not implemented"
	return nil
}

func (h *samplingProvider) parseOperationStrategy(
	strategy *operationStrategy,
	parent *api_v2.PerOperationSamplingStrategies,
) (s *api_v2.SamplingStrategyResponse, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// TODO OperationSamplingStrategy only supports probabilistic sampling

func (h *samplingProvider) parseStrategy(strategy *strategy) *api_v2.SamplingStrategyResponse {
	_ = "STUB: not implemented"
	return nil
}
