// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package oraclecloud // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/oraclecloud"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/oraclecloud"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/oraclecloud/internal/metadata"
)

const (
	// TypeStr is type of detector.
	TypeStr = "oraclecloud"
)

var _ internal.Detector = (*Detector)(nil)

// Detector is an Oracle Cloud metadata detector
type Detector struct {
	provider oraclecloud.Provider
	logger   *zap.Logger
	rb       *metadata.ResourceBuilder
}

// NewDetector creates a new Oracle Cloud metadata detector
func NewDetector(p processor.Settings, dcfg internal.DetectorConfig) (internal.Detector, error) {
	_ = "STUB: not implemented"
	return *new(internal.Detector), nil
}

// Detect detects system metadata and returns a resource with the available ones
func (d *Detector) Detect(ctx context.Context) (resource pcommon.Resource, schemaURL string, err error) {
	_ = "STUB: not implemented"
	// 1. Fast probe for Oracle Cloud platform
	return *new(pcommon.Resource), "", nil
}

// 2. After positive probe, attempt to fetch metadata

// signal error
