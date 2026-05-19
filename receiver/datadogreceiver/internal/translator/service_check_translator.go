// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translator // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/datadogreceiver/internal/translator"

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

type ServiceCheck struct {
	Check     string                       `json:"check"`
	HostName  string                       `json:"host_name"`
	Status    datadogV1.ServiceCheckStatus `json:"status"`
	Timestamp int64                        `json:"timestamp,omitempty"`
	Tags      []string                     `json:"tags,omitempty"`
}

// More information on Datadog service checks: https://docs.datadoghq.com/api/latest/service-checks/
func (mt *MetricsTranslator) TranslateServices(services []ServiceCheck) pmetric.Metrics {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics)
}

// TODO(alexg): proper name

// OTel uses nanoseconds, while Datadog uses seconds

// TODO(alexg): Do this stream thing for service check metrics?
