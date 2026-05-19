// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pdatautil // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/pdatautil"

import (
	"github.com/lightstep/go-expohisto/structure"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// ExpoHistToExponentialDataPoint copies `lightstep/go-expohisto` structure.Histogram to
// pmetric.ExponentialHistogramDataPoint
func ExpoHistToExponentialDataPoint(agg *structure.Histogram[float64], dp pmetric.ExponentialHistogramDataPoint) {
	_ = "STUB: not implemented"
	return
}
