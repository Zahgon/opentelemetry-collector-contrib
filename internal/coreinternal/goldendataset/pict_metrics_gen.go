// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package goldendataset // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/goldendataset"

import (
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// GenerateMetrics takes the filename of a PICT-generated file, walks through all of the rows in the PICT
// file and for each row, generates a MetricData object, collecting them and returning them to the caller.
func GenerateMetrics(metricPairsFile string) ([]pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func pictToCfg(inputs PICTMetricInputs) MetricsCfg {
	_ = "STUB: not implemented"
	return *new(MetricsCfg)
}
