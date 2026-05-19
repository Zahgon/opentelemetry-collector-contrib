// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sdktest // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor/internal/testing/sdktest"

import (
	sdk "go.opentelemetry.io/otel/sdk/metric/metricdata"
)

// Metrics returns the [sdk.Metrics] defined by this [Spec]
func Metrics(spec Spec) []sdk.Metrics { _ = "STUB: not implemented"; return nil }

func (spec Metric) Into() sdk.Metrics { _ = "STUB: not implemented"; return *new(sdk.Metrics) }

// Flatten turns the nested [sdk.ResourceMetrics] structure into a flat
// [sdk.Metrics] slice. If a metric is present multiple time in different scopes
// / resources, the last occurrence is used.
func Flatten(rm sdk.ResourceMetrics) []sdk.Metrics { _ = "STUB: not implemented"; return nil }
