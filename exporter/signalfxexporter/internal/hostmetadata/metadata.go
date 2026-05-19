// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package hostmetadata // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/hostmetadata"

import (
	"sync"

	"github.com/shirou/gopsutil/v4/common"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/dimensions"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk"
	metadata "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata"
)

// Syncer is a config structure for host metadata syncer.
type Syncer struct {
	envMap    common.EnvMap
	logger    *zap.Logger
	dimClient dimensions.MetadataUpdateClient
	once      sync.Once
}

// NewSyncer creates new instance of host metadata syncer.
func NewSyncer(logger *zap.Logger, dimClient dimensions.MetadataUpdateClient, envMap common.EnvMap) *Syncer {
	_ = "STUB: not implemented"
	return nil
}

func (s *Syncer) Sync(md pmetric.Metrics) {
	_ = "STUB: not implemented"
	// skip if already synced or if metrics data is empty
	return
}

func (s *Syncer) syncOnResource(res pcommon.Resource) {
	_ = "STUB: not implemented"
	// If resourcedetection processor is enabled, all the metrics should have resource attributes
	// that can be used to update host metadata.
	// Based of this assumption we check just one ResourceMetrics object,
	return
}

// if no attributes found, we assume that resourcedetection is not enabled or
// it doesn't set right attributes, and we do not retry.

// do not retry if scraping failed.

func (*Syncer) prepareMetadataUpdate(props map[string]string, hostID splunk.HostID) *metadata.MetadataUpdate {
	_ = "STUB: not implemented"
	return nil
}

func (s *Syncer) scrapeHostProperties() map[string]string { _ = "STUB: not implemented"; return nil }
