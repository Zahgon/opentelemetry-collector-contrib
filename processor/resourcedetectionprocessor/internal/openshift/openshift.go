// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package openshift // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/openshift"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"

	ocp "github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/openshift"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/openshift/internal/metadata"
)

const (
	// TypeStr is type of detector.
	TypeStr = "openshift"
)

// NewDetector returns a detector which can detect resource attributes on OpenShift 4.
func NewDetector(set processor.Settings, dcfg internal.DetectorConfig) (internal.Detector, error) {
	_ = "STUB: not implemented"
	return *new(internal.Detector), nil
}

type detector struct {
	logger   *zap.Logger
	provider ocp.Provider
	rb       *metadata.ResourceBuilder
}

func (d *detector) Detect(ctx context.Context) (resource pcommon.Resource, schemaURL string, err error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource), "", nil
}

// return an empty Resource and no error

// TODO(frzifus): support conventions openshift and kubernetes cluster version.
// SEE: https://github.com/open-telemetry/opentelemetry-specification/issues/2913
