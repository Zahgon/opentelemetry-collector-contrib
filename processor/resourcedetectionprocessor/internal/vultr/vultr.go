// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package vultr // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/vultr"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/vultr"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/vultr/internal/metadata"
)

const (
	// TypeStr is type of detector.
	TypeStr = "vultr"
)

// newVultrProvider is overridden in tests to point the provider at a fake server.
var newVultrProvider = vultr.NewProvider

// Ensure Detector implements internal.Detector.
var _ internal.Detector = (*Detector)(nil)

// Detector is a Vultr metadata detector.
type Detector struct {
	provider              vultr.Provider
	logger                *zap.Logger
	rb                    *metadata.ResourceBuilder
	failOnMissingMetadata bool
}

// NewDetector creates a new Vultr metadata detector.
func NewDetector(p processor.Settings, dcfg internal.DetectorConfig) (internal.Detector, error) {
	_ = "STUB: not implemented"
	return *new(internal.Detector), nil
}

// Detect queries the Vultr metadata service and returns a populated resource.
func (d *Detector) Detect(ctx context.Context) (pcommon.Resource, string, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource), "", nil
}

// Prefer the v2 UUID if present; fall back to the legacy instance ID.
