// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package digitalocean // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/digitalocean"

import (
	"context"

	do "github.com/digitalocean/go-metadata"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/digitalocean/internal/metadata"
)

const (
	// TypeStr is type of detector.
	TypeStr = "digitalocean"
)

var _ internal.Detector = (*Detector)(nil)

// newDigitalOceanClient is overridden in tests to point the client at a fake server.
var newDigitalOceanClient = func() *do.Client {
	return do.NewClient()
}

// Detector is a DigitalOcean metadata detector.
type Detector struct {
	client *do.Client
	logger *zap.Logger
	rb     *metadata.ResourceBuilder
}

// NewDetector creates a new DigitalOcean metadata detector.
func NewDetector(p processor.Settings, dcfg internal.DetectorConfig) (internal.Detector, error) {
	_ = "STUB: not implemented"
	return *new(internal.Detector), nil
}

// Detect detects system metadata and returns a resource with the available ones.
func (d *Detector) Detect(_ context.Context) (pcommon.Resource, string, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource), "", nil
}
