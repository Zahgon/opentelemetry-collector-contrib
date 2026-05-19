// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package upcloud // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/upcloud"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/upcloud"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/upcloud/internal/metadata"
)

const (
	// TypeStr is type of detector.
	TypeStr = "upcloud"
)

// newUpcloudProvider is overridden in tests to point the provider at a fake server.
var newUpcloudProvider = upcloud.NewProvider

// Ensure Detector implements internal.Detector.
var _ internal.Detector = (*Detector)(nil)

// Detector is a Upcloud metadata detector.
type Detector struct {
	provider              upcloud.Provider
	logger                *zap.Logger
	rb                    *metadata.ResourceBuilder
	failOnMissingMetadata bool
}

// NewDetector creates a new Upcloud metadata detector.
func NewDetector(p processor.Settings, dcfg internal.DetectorConfig) (internal.Detector, error) {
	_ = "STUB: not implemented"
	return *new(internal.Detector), nil
}

// Detect queries the Upcloud metadata service and returns a populated resource.
func (d *Detector) Detect(ctx context.Context) (pcommon.Resource, string, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource), "", nil
}
