// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// sdktest performs partial comparison of [sdk.ResourceMetrics] to a [Spec].
package sdktest // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor/internal/testing/sdktest"

import (
	"github.com/google/go-cmp/cmp"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/metric"
	sdk "go.opentelemetry.io/otel/sdk/metric/metricdata"
)

type Option = cmp.Option

// Test the metrics returned by [metric.ManualReader.Collect] against the [Spec]
func Test(spec Spec, mr *metric.ManualReader, opts ...Option) error {
	_ = "STUB: not implemented"
	return nil
}

// IgnoreTime ignores [sdk.DataPoint.Time] and [sdk.DataPoint.StartTime],
// because those are changing per run and typically not of interest.
func IgnoreTime() Option { _ = "STUB: not implemented"; return *new(Option) }

// IgnoreTime ignores [sdk.Metrics.Unit] and [sdk.Metrics.Description],
// because those are usually static
func IgnoreMetadata() Option { _ = "STUB: not implemented"; return *new(Option) }

// IgnoreUnspec ignores any Metrics not present in the [Spec]
func IgnoreUnspec(spec Spec) Option { _ = "STUB: not implemented"; return *new(Option) }

// Sort [sdk.Metrics] by name and [sdk.DataPoint] by their [attribute.Set]
func Sort() Option { _ = "STUB: not implemented"; return *new(Option) }

func sort[N int64 | float64]() Option { _ = "STUB: not implemented"; return *new(Option) }

// DataPoint holds a [sdk.DataPoints] and its attributes as a plain map.
// See [Transform]
type DataPoint[N int64 | float64] struct {
	Attributes map[string]any
	sdk.DataPoint[N]
}

// Transform turns []sdk.DataPoint[N] into []DataPoint[N].
//
// Primarily done to have DataPoint.Attributes as a flat, diffable map instead
// of the hard to understand internal structure of [attribute.Set], which is
// being truncated by go-cmp before reaching the depth where attribute values
// appear.
//
// This must happen on the slice level, transforming the values is not
// sufficient because when entire DataPoints are added / removed, go-cmp does
// not apply transformers on the fields.
func Transform() Option { _ = "STUB: not implemented"; return *new(Option) }

func transform[N int64 | float64]() Option { _ = "STUB: not implemented"; return *new(Option) }

func attrMap(set attribute.Set) map[string]any { _ = "STUB: not implemented"; return nil }
