// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package akamai // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/akamai"

import (
	"context"

	linodemeta "github.com/linode/go-metadata"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/akamai/internal/metadata"
)

const (
	// TypeStr is type of detector.
	TypeStr = "akamai"
)

var _ internal.Detector = (*Detector)(nil)

// newAkamaiClient is overridden in tests to point the client at a fake server.
var newAkamaiClient = func(ctx context.Context) (akamaiAPI, error) {
	return linodemeta.NewClient(ctx)
}

type akamaiAPI interface {
	GetInstance(ctx context.Context) (*linodemeta.InstanceData, error)
}

// Detector is a Akamai metadata detector.
type Detector struct {
	client akamaiAPI
	logger *zap.Logger
	rb     *metadata.ResourceBuilder
}

// NewDetector creates a new Akamai metadata detector.
func NewDetector(p processor.Settings, dcfg internal.DetectorConfig) (internal.Detector, error) {
	_ = "STUB: not implemented"
	return *new(internal.Detector), nil
}

// Detect detects system metadata and returns a resource with the available ones.
func (d *Detector) Detect(ctx context.Context) (pcommon.Resource, string, error) {
	_ = "STUB: not implemented"
	// Try to fetch instance metadata; if it fails we're not on Akamai (or metadata unreachable).
	return *new(pcommon.Resource), "", nil
}
