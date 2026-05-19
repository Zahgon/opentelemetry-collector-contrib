// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package dimensions // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/dimensions"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"

	metadata "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata"
)

type EntityEventTransformer struct {
	defaultProperties map[string]string
}

func NewEntityEventTransformer(defaultProperties map[string]string) *EntityEventTransformer {
	_ = "STUB: not implemented"
	return nil
}

func (t *EntityEventTransformer) TransformEntityEvent(event metadata.EntityEvent) (*DimensionUpdate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *EntityEventTransformer) extractPropertiesAndTags(attrs pcommon.Map) (map[string]*string, map[string]bool) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractDimensionKeyValue(entityType string, entityID pcommon.Map) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}
