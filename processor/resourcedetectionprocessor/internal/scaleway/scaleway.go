// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package scaleway // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/scaleway"

import (
	"context"

	instance "github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/scaleway/internal/metadata"
)

const (
	// TypeStr is type of detector.
	TypeStr = "scaleway"
)

var _ internal.Detector = (*Detector)(nil)

// newScalewayClient is overridden in tests to point the client at a fake server.
var newScalewayClient = instance.NewMetadataAPI

// Detector is a Scaleway metadata detector.
type Detector struct {
	client *instance.MetadataAPI
	logger *zap.Logger
	rb     *metadata.ResourceBuilder
}

// NewDetector creates a new Scaleway metadata detector.
func NewDetector(p processor.Settings, dcfg internal.DetectorConfig) (internal.Detector, error) {
	_ = "STUB: not implemented"
	return *new(internal.Detector), nil
}

// Detect detects system metadata and returns a resource with the available ones.
func (d *Detector) Detect(_ context.Context) (pcommon.Resource, string, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource), "", nil
}

// Cloud provider and platform values will be "scaleway_cloud" and "scaleway_cloud_platform" from conventions when it's merged.
// d.rb.SetCloudProvider(conventions.CloudProviderScalewayCloud.Value.AsString())
// d.rb.SetCloudPlatform(conventions.CloudPlatformScalewayCloud.Value.AsString())

// zoneToRegion extracts the region name from a Scaleway zone like "fr-par-1" -> "fr-par".
func zoneToRegion(zone string) string { _ = "STUB: not implemented"; return "" }
