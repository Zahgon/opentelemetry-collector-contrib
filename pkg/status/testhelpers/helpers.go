// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package testhelpers // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/status/testhelpers"

import (
	"go.opentelemetry.io/collector/component/componentstatus"
	"go.opentelemetry.io/collector/pipeline"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/status"
)

// PipelineMetadata groups together component and instance IDs for a hypothetical pipeline used
// for testing purposes.
type PipelineMetadata struct {
	PipelineID  pipeline.ID
	ReceiverID  *componentstatus.InstanceID
	ProcessorID *componentstatus.InstanceID
	ExporterID  *componentstatus.InstanceID
}

// InstanceIDs returns a slice of instanceIDs for components within the hypothetical pipeline.
func (p *PipelineMetadata) InstanceIDs() []*componentstatus.InstanceID {
	_ = "STUB: not implemented"
	return nil
}

// NewPipelineMetadata returns a metadata for a hypothetical pipeline.
func NewPipelineMetadata(signal pipeline.Signal) *PipelineMetadata {
	_ = "STUB: not implemented"
	return nil
}

// SeedAggregator records a status event for each instanceID.
func SeedAggregator(
	agg *status.Aggregator,
	instanceIDs []*componentstatus.InstanceID,
	statuses ...componentstatus.Status,
) {
	_ = "STUB: not implemented"
	return
}
