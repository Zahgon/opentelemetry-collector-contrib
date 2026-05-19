// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelserializer // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/serializer/otelserializer"

import (
	"bytes"

	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/datapoints"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/elasticsearch"
)

func (*Serializer) SerializeMetrics(resource pcommon.Resource, resourceSchemaURL string, scope pcommon.InstrumentationScope, scopeSchemaURL string, dataPoints []datapoints.DataPoint, validationErrors *[]error, idx elasticsearch.Index, buf *bytes.Buffer) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func serializeDataPoints(w *jsonWriter, dataPoints []datapoints.DataPoint, validationErrors *[]error, first bool) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// TODO here's potential for more optimization by directly serializing the value instead of allocating a pcommon.Value
//  the tradeoff is that this would imply a duplicated logic for the ECS mode

// TODO: support quantiles
// https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/34561

// DynamicTemplate returns the name of dynamic template that applies to the metric and data point,
// so that the field is indexed into Elasticsearch with the correct mapping. The name should correspond to a
// dynamic template that is defined in ES mapping, e.g.
// https://github.com/elastic/elasticsearch/blob/8.15/x-pack/plugin/core/template-resources/src/main/resources/metrics%40mappings.json

// workaround for https://github.com/elastic/elasticsearch/issues/99123
// should use a string field to benefit from run-length encoding
