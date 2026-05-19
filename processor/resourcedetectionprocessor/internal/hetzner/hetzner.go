// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package hetzner // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/hetzner"

import (
	"context"

	hcloudmeta "github.com/hetznercloud/hcloud-go/v2/hcloud/metadata"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/hetzner/internal/metadata"
)

const (
	// TypeStr is type of detector.
	TypeStr = "hetzner"
)

var _ internal.Detector = (*Detector)(nil)

// newHcloudClient is overridden in tests to point the client at a fake server.
var newHcloudClient = func() *hcloudmeta.Client {
	return hcloudmeta.NewClient()
}

// Detector is a Hetzner metadata detector.
type Detector struct {
	client *hcloudmeta.Client
	logger *zap.Logger
	rb     *metadata.ResourceBuilder
}

// NewDetector creates a new Hetzner metadata detector.
func NewDetector(p processor.Settings, dcfg internal.DetectorConfig) (internal.Detector, error) {
	_ = "STUB: not implemented"
	return *new(internal.Detector), nil
}

// Detect detects system metadata and returns a resource with the available ones.
func (d *Detector) Detect(_ context.Context) (pcommon.Resource, string, error) {
	_ = "STUB: not implemented"
	// Quick check: if not running in Hetzner Cloud, return empty.
	return *new(pcommon.Resource), "", nil
}
