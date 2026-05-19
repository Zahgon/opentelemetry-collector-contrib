// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azure // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/azure"

import (
	"context"
	"regexp"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/azure"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/azure/internal/metadata"
)

const (
	// TypeStr is type of detector.
	TypeStr   = "azure"
	tagPrefix = "azure.tag."
)

var _ internal.Detector = (*Detector)(nil)

// Detector is an Azure metadata detector
type Detector struct {
	provider      azure.Provider
	tagKeyRegexes []*regexp.Regexp
	logger        *zap.Logger
	rb            *metadata.ResourceBuilder
}

// NewDetector creates a new Azure metadata detector
func NewDetector(p processor.Settings, dcfg internal.DetectorConfig) (internal.Detector, error) {
	_ = "STUB: not implemented"
	return *new(internal.Detector), nil
}

// Detect detects system metadata and returns a resource with the available ones
func (d *Detector) Detect(ctx context.Context) (resource pcommon.Resource, schemaURL string, err error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource), "", nil
}

// return an empty Resource and no error

// Use osProfile.computerName for host.name, falling back to the VM name
// if computerName is empty (e.g., VMs created from specialized disks).

// Also save compute.Name in "azure.vm.name" as host.id (AttributeHostName) is
// used by system detector.

func matchAzureTags(azureTags []azure.ComputeTagsListMetadata, tagKeyRegexes []*regexp.Regexp) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func compileRegexes(cfg Config) ([]*regexp.Regexp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func regexArrayMatch(arr []*regexp.Regexp, val string) bool {
	_ = "STUB: not implemented"
	return false
}
