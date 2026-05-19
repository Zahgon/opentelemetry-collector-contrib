// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package geoipprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/geoipprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pmetric"
)

func (g *geoIPProcessor) processMetrics(ctx context.Context, ms pmetric.Metrics) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

func (g *geoIPProcessor) processMetricAttributes(ctx context.Context, m pmetric.Metric) error {
	_ = "STUB: not implemented"
	// This is a lot of repeated code, but since there is no single parent superclass
	// between metric data types, we can't use polymorphism.
	//exhaustive:enforce
	return nil
}
