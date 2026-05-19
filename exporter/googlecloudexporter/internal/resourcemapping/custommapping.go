// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package resourcemapping // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlecloudexporter/internal/resourcemapping"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	monitoredrespb "google.golang.org/genproto/googleapis/api/monitoredres"
)

var (
	mappingKey                   = "gcp.resource_type"
	monitoredResourceLabelPrefix = "gcp."
)

// CustomLoggingMonitoredResourceMapping allows mapping from OTel resources to
// Monitored Resources defined here:
// https://cloud.google.com/monitoring/api/resources
//
// The monitored resource type is extracted from the `gcp.resource_type`
// attribute. And the monitored resource labels are extracted from resource
// attributes with the prefix `gcp.<monitored resource type>.`.
func CustomLoggingMonitoredResourceMapping(r pcommon.Resource) *monitoredrespb.MonitoredResource {
	_ = "STUB: not implemented"
	return nil
}

// CustomMetricMonitoredResourceMapping allows mapping from OTel resources to
// Monitored Resources defined here:
// https://cloud.google.com/monitoring/api/resources
//
// The monitored resource type is extracted from the `gcp.resource_type`
// attribute. And the monitored resource labels are extracted from resource
// attributes with the prefix `gcp.<monitored resource type>.`.
func CustomMetricMonitoredResourceMapping(r pcommon.Resource) *monitoredrespb.MonitoredResource {
	_ = "STUB: not implemented"
	return nil
}

func customMonitoredResourceMapping(r pcommon.Resource, mmrFunc func(pcommon.Resource) *monitoredrespb.MonitoredResource) *monitoredrespb.MonitoredResource {
	_ = "STUB: not implemented"
	return nil
}

// Extract the monitored resource label by separating it from the prefix.
