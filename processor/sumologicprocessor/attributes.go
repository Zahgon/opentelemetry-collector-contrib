// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sumologicprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/sumologicprocessor"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// This file contains some common functionalities for subprocessors that modify attributes (represented by pcommon.Map)

type attributesProcessor interface {
	processAttributes(pcommon.Map) error
}

func processMetricLevelAttributes(proc attributesProcessor, metric pmetric.Metric) error {
	_ = "STUB: not implemented"
	return nil
}

func mapToPcommonMap(m map[string]pcommon.Value) pcommon.Map {
	_ = "STUB: not implemented"
	return *new(pcommon.Map)
}

func mapToPcommonValue(m map[string]pcommon.Value) pcommon.Value {
	_ = "STUB: not implemented"
	return *new(pcommon.Value)
}
