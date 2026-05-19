// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metadata // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/gcp/internal/metadata"

import (
	"github.com/GoogleCloudPlatform/opentelemetry-operations-go/detectors/gcp"
)

func (*ResourceBuilder) SetFromCallable(set func(string), detect func() (string, error)) error {
	_ = "STUB: not implemented"
	return nil
}

func (rb *ResourceBuilder) SetZoneAndRegion(detect func() (string, string, error)) error {
	_ = "STUB: not implemented"
	return nil
}

func (rb *ResourceBuilder) SetZoneOrRegion(detect func() (string, gcp.LocationType, error)) error {
	_ = "STUB: not implemented"
	return nil
}

func (rb *ResourceBuilder) SetManagedInstanceGroup(detect func() (gcp.ManagedInstanceGroup, error)) error {
	_ = "STUB: not implemented"
	return nil
}
