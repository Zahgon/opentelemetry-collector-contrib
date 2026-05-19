// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package consul // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/consul"

import (
	"context"

	"github.com/hashicorp/consul/api"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/consul"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/consul/internal/metadata"
)

const (
	// TypeStr is type of detector.
	TypeStr = "consul"
)

var _ internal.Detector = (*Detector)(nil)

// Detector is a system metadata detector
type Detector struct {
	provider consul.Provider
	logger   *zap.Logger
	rb       *metadata.ResourceBuilder
}

// buildConsulAPIConfig translates the detector Config into a hashicorp consul api.Config.
func buildConsulAPIConfig(userCfg Config) *api.Config { _ = "STUB: not implemented"; return nil }

func NewDetector(p processor.Settings, dcfg internal.DetectorConfig) (internal.Detector, error) {
	_ = "STUB: not implemented"
	return *new(internal.Detector), nil
}

// Detect detects system metadata and returns a resource with the available ones
func (d *Detector) Detect(ctx context.Context) (resource pcommon.Resource, schemaURL string, err error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource), "", nil
}
