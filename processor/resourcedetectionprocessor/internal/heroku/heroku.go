// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package heroku // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/heroku"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/heroku/internal/metadata"
)

const (
	// TypeStr is type of detector.
	TypeStr = "heroku"
)

// NewDetector returns a detector which can detect resource attributes on Heroku
func NewDetector(set processor.Settings, dcfg internal.DetectorConfig) (internal.Detector, error) {
	_ = "STUB: not implemented"
	return *new(internal.Detector), nil
}

type detector struct {
	logger *zap.Logger
	rb     *metadata.ResourceBuilder
}

// Detect detects heroku metadata and returns a resource with the available ones
func (d *detector) Detect(_ context.Context) (resource pcommon.Resource, schemaURL string, err error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource), "", nil
}

// some heroku deployments will enable some of the metadata.
