// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package aggregator // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/signaltometricsconnector/internal/aggregator"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/connector/signaltometricsconnector/internal/model"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

// FilterAttrsFunc is a lazy function that produces a filtered attribute map.
// It is only called when the aggregator encounters a new unique attribute
// set and needs to create a datapoint. This avoids allocating a pcommon.Map
// for every signal item.
type FilterAttrsFunc func() (pcommon.Map, error)

// Aggregator provides a single interface to update all metrics
// datastructures. The required datastructure is selected using
// the metric definition.
type Aggregator[K any] struct {
	result pmetric.Metrics
	// smLookup maps resourceID against scope metrics since the aggregator
	// always produces a single scope.
	smLookup    map[[16]byte]pmetric.ScopeMetrics
	valueCounts map[model.MetricKey]map[[16]byte]map[[16]byte]*valueCountDP
	sums        map[model.MetricKey]map[[16]byte]map[[16]byte]*sumDP
	gauges      map[model.MetricKey]map[[16]byte]map[[16]byte]*gaugeDP
	timestamp   time.Time
	errorMode   ottl.ErrorMode
	logger      *zap.Logger
}

// NewAggregator creates a new instance of aggregator.
func NewAggregator[K any](metrics pmetric.Metrics, errorMode ottl.ErrorMode, logger *zap.Logger) *Aggregator[K] {
	_ = "STUB: not implemented"
	return nil
}

func (a *Aggregator[K]) Aggregate(
	ctx context.Context,
	tCtx K,
	md model.MetricDef[K],
	resAttrs pcommon.Map,
	attrID [16]byte,
	filterAttrs FilterAttrsFunc,
	defaultCount int64,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Gracefully skip missing keys in ExtractGrokPatterns

// handleError handles errors based on the configured ErrorMode.
// It returns nil for ignore/silent modes and returns the error for propagate mode.
func (a *Aggregator[K]) handleError(err error) error { _ = "STUB: not implemented"; return nil }

// Finalize finalizes the aggregations performed by the aggregator so far into
// the pmetric.Metrics used to create this instance of the aggregator. Finalize
// should be called once per aggregator instance and the aggregator instance
// should not be used after Finalize is called.
func (a *Aggregator[K]) Finalize(mds []model.MetricDef[K]) { _ = "STUB: not implemented"; return }

// If there are two metric defined with the same key required by metricKey
// then they will be aggregated within the same metric and produced
// together. Deleting the key ensures this while preventing duplicates.

func (a *Aggregator[K]) aggregateInt(
	md model.MetricDef[K],
	resAttrs pcommon.Map,
	attrID [16]byte,
	filterAttrs FilterAttrsFunc,
	v int64,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *Aggregator[K]) aggregateDouble(
	md model.MetricDef[K],
	resAttrs pcommon.Map,
	attrID [16]byte,
	filterAttrs FilterAttrsFunc,
	v float64,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *Aggregator[K]) aggregateGauge(
	md model.MetricDef[K],
	resAttrs pcommon.Map,
	attrID [16]byte,
	filterAttrs FilterAttrsFunc,
	v any,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *Aggregator[K]) aggregateValueCount(
	md model.MetricDef[K],
	resAttrs pcommon.Map,
	attrID [16]byte,
	filterAttrs FilterAttrsFunc,
	value float64, count int64,
) error {
	_ = "STUB: not implemented"

	// Nothing to record as count is zero
	return nil
}

func (a *Aggregator[K]) getResourceID(resourceAttrs pcommon.Map) [16]byte {
	_ = "STUB: not implemented"
	return nil
}

// getValueCount evaluates OTTL to get count and value respectively. Count is
// optional and defaults to the default count if the OTTL statement for count
// is missing. Value is required and returns an error if OTTL statement for
// value is missing.
func getValueCount[K any](
	ctx context.Context, tCtx K,
	valueExpr, countExpr *ottl.ValueExpression[K],
	defaultCount int64,
) (float64, int64, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func getIntFromOTTL[K any](
	ctx context.Context,
	tCtx K,
	s *ottl.ValueExpression[K],
) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func getDoubleFromOTTL[K any](
	ctx context.Context,
	tCtx K,
	s *ottl.ValueExpression[K],
) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
