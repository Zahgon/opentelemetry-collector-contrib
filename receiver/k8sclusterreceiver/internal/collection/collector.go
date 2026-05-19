// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package collection // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/collection"

import (
	"time"

	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/metadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/service"
)

// TODO: Consider moving some of these constants to
// https://go.opentelemetry.io/collector/blob/main/model/semconv/opentelemetry.go.

// DataCollector emits metrics with CollectMetricData based on the Kubernetes API objects in the metadata store.
type DataCollector struct {
	settings                 receiver.Settings
	metadataStore            *metadata.Store
	nodeConditionsToReport   []string
	allocatableTypesToReport []string
	metricsBuilder           *metadata.MetricsBuilder
}

// NewDataCollector returns a DataCollector.
func NewDataCollector(set receiver.Settings, ms *metadata.Store,
	metricsBuilderConfig metadata.MetricsBuilderConfig, nodeConditionsToReport, allocatableTypesToReport []string,
) *DataCollector {
	_ = "STUB: not implemented"
	return nil
}

func (dc *DataCollector) CollectMetricData(currentTime time.Time) pmetric.Metrics {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics)
}

func (dc *DataCollector) calculateServiceEndpointCounts() map[string]service.EndpointCountsByKey {
	_ = "STUB: not implemented"
	return nil
}

// K8s Spec: ready == true or nil means endpoint can receive NEW connections

// K8s Spec: serving == true, or if nil "consumers should defer to the ready condition"
// https://kubernetes.io/docs/reference/kubernetes-api/service-resources/endpoint-slice-v1/#EndpointConditions

// K8s Spec: terminating == true means endpoint is draining
