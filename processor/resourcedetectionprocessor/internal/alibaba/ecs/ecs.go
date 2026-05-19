// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ecs // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/alibaba/ecs"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"

	ecsprovider "github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/alibaba/ecs"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/alibaba/ecs/internal/metadata"
)

const (
	// TypeStr is the detector type id.
	TypeStr = "alibaba_ecs"
)

var _ internal.Detector = (*Detector)(nil)

// Detector queries the Alibaba Cloud ECS metadata service and emits resource attributes.
type Detector struct {
	logger                *zap.Logger
	rb                    *metadata.ResourceBuilder
	metadataProvider      ecsprovider.Provider
	failOnMissingMetadata bool
}

// NewDetector creates an Alibaba Cloud ECS detector.
func NewDetector(set processor.Settings, dcfg internal.DetectorConfig) (internal.Detector, error) {
	_ = "STUB: not implemented"
	return *new(internal.Detector), nil
}

func (d *Detector) Detect(ctx context.Context) (resource pcommon.Resource, schemaURL string, err error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource), "", nil
}
