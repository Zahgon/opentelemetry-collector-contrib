// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package classic // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/ibmcloud/classic"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"

	classicprovider "github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/ibmcloud/classic"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/ibmcloud/classic/internal/metadata"
)

const (
	// TypeStr is the detector type string.
	TypeStr = "ibmcloud_classic"
)

var _ internal.Detector = (*Detector)(nil)

// Detector queries the IBM Cloud Classic (SoftLayer) Resource Metadata Service
// and emits resource attributes.
type Detector struct {
	provider classicprovider.Provider
	logger   *zap.Logger
	rb       *metadata.ResourceBuilder
}

// NewDetector creates an IBM Cloud Classic detector.
func NewDetector(p processor.Settings, dcfg internal.DetectorConfig) (internal.Detector, error) {
	_ = "STUB: not implemented"
	return *new(internal.Detector), nil
}

// Detect detects IBM Cloud Classic instance metadata and returns a resource with the available attributes.
func (d *Detector) Detect(ctx context.Context) (pcommon.Resource, string, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource), "", nil
}

// TODO: Use semconv constant once CloudPlatformIBMCloudClassic is added.
