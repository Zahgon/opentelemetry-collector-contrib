// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datasetexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datasetexporter"

import (
	"context"

	"github.com/scalyr/dataset-go/pkg/api/add_events"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

const ServiceNameKey = "service.name"

func createTracesExporter(ctx context.Context, set exporter.Settings, config component.Config) (exporter.Traces, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Traces), nil
}

func buildEventFromSpan(
	bundle spanBundle,
	serverHost string,
	settings TracesSettings,
) *add_events.EventBundle {
	_ = "STUB: not implemented"
	return nil
}

// for now we care only small subset of attributes
// updateWithPrefixedValues(attrs, "resource_", "_", resource.Attributes().AsRaw(), 0)

// since attributes are overwriting existing keys, they have to be at the end

const (
	resourceName = "resource_name"
	resourceType = "resource_type"
)

type ResourceType string

const (
	Service = ResourceType("service")
	Process = ResourceType("process")
)

func updateResource(attrs, resource map[string]any) {
	_ = "STUB: not implemented"
	// first detect, whether there is key service.name
	// if it's there, we are done
	return
}

// if we were not able to find service name, lets mark it as process

// but still try to search for anything, that start on service
// if we found it, we will mark it as service

// when we find process.pid - lets use it as name

type spanBundle struct {
	span     ptrace.Span
	resource pcommon.Resource
	scope    pcommon.InstrumentationScope
}

func buildEventsFromTraces(ld ptrace.Traces, serverHost string, settings TracesSettings) []*add_events.EventBundle {
	_ = "STUB: not implemented"
	return nil
}

// convert spans into events

func (e *datasetExporter) consumeTraces(_ context.Context, ld ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}
