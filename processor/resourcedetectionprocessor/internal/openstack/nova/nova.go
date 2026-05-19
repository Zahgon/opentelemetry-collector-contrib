// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package nova // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/openstack/nova"

import (
	"context"
	"regexp"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"

	novaprovider "github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/openstack/nova"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/openstack/nova/internal/metadata"
)

const (
	// TypeStr is the detector type id.
	TypeStr = "nova"

	// LabelPrefix is the attribute prefix for Nova metadata keys (user/tenant-provided).
	LabelPrefix = "openstack.nova.meta."
)

var _ internal.Detector = (*Detector)(nil)

// Detector queries the OpenStack Nova metadata service and emits resource attributes.
type Detector struct {
	logger                *zap.Logger
	rb                    *metadata.ResourceBuilder
	metadataProvider      novaprovider.Provider
	labelRegexes          []*regexp.Regexp
	failOnMissingMetadata bool
}

// NewDetector creates a Nova detector.
func NewDetector(set processor.Settings, dcfg internal.DetectorConfig) (internal.Detector, error) {
	_ = "STUB: not implemented"
	return *new(internal.Detector), nil
}

func (d *Detector) Detect(ctx context.Context) (resource pcommon.Resource, schemaURL string, err error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource), "", nil
}

// Optional: EC2‑compatible instance type (don’t fail if missing)

// Optional: capture selected meta labels under openstack.meta.<key>

func compileRegexes(pats []string) ([]*regexp.Regexp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func regexArrayMatch(arr []*regexp.Regexp, s string) bool { _ = "STUB: not implemented"; return false }
