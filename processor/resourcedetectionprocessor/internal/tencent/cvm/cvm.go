// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package cvm // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/tencent/cvm"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"

	cvmprovider "github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/tencent/cvm"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/tencent/cvm/internal/metadata"
)

const (
	// TypeStr is the detector type id.
	TypeStr = "tencent_cvm"
)

var _ internal.Detector = (*Detector)(nil)

// Detector queries the Tencent Cloud CVM metadata service and emits resource attributes.
type Detector struct {
	logger                *zap.Logger
	rb                    *metadata.ResourceBuilder
	metadataProvider      cvmprovider.Provider
	failOnMissingMetadata bool
}

// NewDetector creates a Tencent Cloud CVM detector.
func NewDetector(set processor.Settings, dcfg internal.DetectorConfig) (internal.Detector, error) {
	_ = "STUB: not implemented"
	return *new(internal.Detector), nil
}

func (d *Detector) Detect(ctx context.Context) (resource pcommon.Resource, schemaURL string, err error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource), "", nil
}
