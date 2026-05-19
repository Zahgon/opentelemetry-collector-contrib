// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package testutils // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/testutils"

import (
	"testing"

	"go.opentelemetry.io/collector/pdata/pmetric"
)

func AssertMetricInt(tb testing.TB, m pmetric.Metric, expectedMetric string, expectedType pmetric.MetricType, expectedValue any) {
	_ = "STUB: not implemented"
	return
}

func assertMetric(tb testing.TB, m pmetric.Metric, expectedMetric string, expectedType pmetric.MetricType) pmetric.NumberDataPointSlice {
	_ = "STUB: not implemented"
	return *new(pmetric.NumberDataPointSlice)
}

//exhaustive:enforce
