// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metadata // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/metadata"

import (
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"

	metadataPkg "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata"
)

// GetEntityEvents processes metadata updates and returns entity events that describe the metadata changes.
func GetEntityEvents(oldMetadata, newMetadata map[metadataPkg.ResourceID]*KubernetesMetadata, timestamp pcommon.Timestamp, reportingInterval time.Duration) metadataPkg.EntityEventsSlice {
	_ = "STUB: not implemented"
	return *new(metadataPkg.EntityEventsSlice)
}

// An object was present, but no longer is. Create a "delete" event.

// All "new" are current objects. Create "state" events. "old" state does not matter.
