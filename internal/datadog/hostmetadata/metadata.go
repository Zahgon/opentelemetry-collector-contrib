// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package hostmetadata is responsible for collecting host metadata from different providers
// such as EC2, ECS, AWS, etc and pushing it to Datadog.
package hostmetadata // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/datadog/hostmetadata"

import (
	"context"
	"net/http"

	"github.com/DataDog/datadog-agent/pkg/opentelemetry-mapping-go/inframetadata"
	"github.com/DataDog/datadog-agent/pkg/opentelemetry-mapping-go/inframetadata/payload"
	"github.com/DataDog/datadog-agent/pkg/opentelemetry-mapping-go/otlp/attributes"
	"github.com/DataDog/datadog-agent/pkg/opentelemetry-mapping-go/otlp/attributes/source"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/datadog/clientutil"
)

// metadataFromAttributes gets metadata info from attributes following
// OpenTelemetry semantic conventions
func metadataFromAttributes(attrs pcommon.Map, hostFromAttributesHandler attributes.HostFromAttributesHandler) payload.HostMetadata {
	_ = "STUB: not implemented"
	return *new(payload.HostMetadata)
}

// AWS EC2 resource metadata

func fillHostMetadata(params exporter.Settings, pcfg PusherConfig, p source.Provider, hm *payload.HostMetadata) {
	_ = "STUB: not implemented"
	// Could not get hostname from attributes
	return
}

// This information always gets filled in here
// since it does not come from OTEL conventions

// EC2 data was not set from attributes

// System data was not set from attributes

type pusher struct {
	params     exporter.Settings
	pcfg       PusherConfig
	retrier    *clientutil.Retrier
	httpClient *http.Client
}

func (p *pusher) pushMetadata(hm payload.HostMetadata) error { _ = "STUB: not implemented"; return nil }

// Set the content type to JSON and the content encoding to gzip

func (p *pusher) Push(_ context.Context, hm payload.HostMetadata) error {
	_ = "STUB: not implemented"
	return nil

	// if the hostname is empty, don't send metadata; we don't need it.
}

var _ inframetadata.Pusher = (*pusher)(nil)

// NewPusher creates a new inframetadata.Pusher that pushes metadata payloads
func NewPusher(params exporter.Settings, pcfg PusherConfig) inframetadata.Pusher {
	_ = "STUB: not implemented"
	return *new(inframetadata.Pusher)
}

// deepCopyHostMetadata creates a deep copy of the host metadata payload to avoid
// race conditions when the payload is shared with the reporter's gohai collector
// that may refresh its internal maps concurrently.
// If deep copying fails, it returns the original payload to ensure the operation
// can still proceed (though this should not happen in normal operation).
func deepCopyHostMetadata(hm payload.HostMetadata) payload.HostMetadata {
	_ = "STUB: not implemented"
	// Use JSON marshal/unmarshal to create a deep copy
	// This ensures all nested maps and slices are properly copied
	return *new(payload.HostMetadata)
}

// Return original payload if marshaling fails

// Return original payload if unmarshaling fails

// RunPusher to push host metadata payloads from the host where the Collector is running periodically to Datadog intake.
// This function is blocking and it is meant to be run on a goroutine.
func RunPusher(ctx context.Context, params exporter.Settings, pcfg PusherConfig, p source.Provider, attrs pcommon.Map, reporter *inframetadata.Reporter) {
	_ = "STUB: not implemented"
	// Push metadata every 30 minutes
	return
}

// Get host metadata from resources and fill missing info using our exporter.
// Currently we only retrieve it once but still send the same payload
// every 30 minutes for consistency with the Datadog Agent behavior.
//
// NOTE: The gohai payload contains maps that are shared with the reporter's
// internal gohai collector. We deep copy the payload before each
// ConsumeHostMetadata call to avoid race conditions when the reporter refreshes
// maps concurrently with JSON marshaling.

// Consume one first time - deep copy to avoid race condition

// Deep copy before each consumption to avoid race condition with reporter's gohai refresh
