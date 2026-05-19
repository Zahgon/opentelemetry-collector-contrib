// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translator // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/datadogreceiver/internal/translator"

import (
	"net/http"

	"github.com/DataDog/agent-payload/v5/gogen"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

const (
	TypeGauge string = "gauge"
	TypeRate  string = "rate"
	TypeCount string = "count"
)

type SeriesList struct {
	Series []datadogV1.Series `json:"series"`
}

func (*MetricsTranslator) HandleSeriesV2Payload(req *http.Request) (mp []*gogen.MetricPayload_MetricSeries, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// handle json messages if set, otherwise handle protobuf

func (mt *MetricsTranslator) TranslateSeriesV1(series SeriesList) pmetric.Metrics {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics)
}

// See https://docs.datadoghq.com/metrics/types/?tab=count#definition

// Type is unset/unspecified

// The Datadog API returns a slice of slices of points [][]*float64 which is a bit awkward to work with.
// It looks like this:
// points := [][]*float64{
// 	{&timestamp1, &value1},
// 	{&timestamp2, &value2},
// }
// We need to flatten this to a slice of *float64 to work with it. And we know that in that slice, the first
// element is the timestamp and the second is the value.

// The datapoint is missing a timestamp and/or value, so this point should be skipped

// OTel uses nanoseconds, while Datadog uses seconds

func (mt *MetricsTranslator) TranslateSeriesV2(series []*gogen.MetricPayload_MetricSeries) pmetric.Metrics {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics)
}

// The V2 payload stores the host name under in the Resources field

// TODO(jesus.vazquez) (Do this with string interning)

// Host has already been added as a resource attribute in parseSeriesProperties(), so avoid duplicating that attribute

// TODO: check if this is correct handling of SourceTypeName field

// See https://docs.datadoghq.com/metrics/types/?tab=count#definition

// TODO: verify that this is always the case

// Type is unset/unspecified

// OTel uses nanoseconds, while Datadog uses seconds
// TODO(jesus.vazquez) Review this copy

func getV2Resources(resources []*gogen.MetricPayload_Resource) map[string]string {
	_ = "STUB: not implemented"
	return nil
}
