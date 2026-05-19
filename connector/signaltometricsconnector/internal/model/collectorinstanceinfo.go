// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package model // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/signaltometricsconnector/internal/model"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	conventions "go.opentelemetry.io/otel/semconv/v1.40.0"

	"github.com/open-telemetry/opentelemetry-collector-contrib/connector/signaltometricsconnector/internal/metadata"
)

var prefix = metadata.Type.String()

// prefixedServiceInstanceIDKey is the service.instance.id key with the connector prefix,
// pre-computed to avoid a string allocation on every Copy() call.
var prefixedServiceInstanceIDKey = keyWithPrefix(string(conventions.ServiceInstanceIDKey))

// CollectorInstanceInfo holds the attributes that could uniquely identify
// the current collector instance. These attributes are initialized from the
// telemetry settings. The CollectorInstanceInfo can copy these attributes,
// with a given prefix, to a provided map.
type CollectorInstanceInfo struct {
	size              int
	serviceInstanceID string
}

func NewCollectorInstanceInfo(
	set component.TelemetrySettings,
) CollectorInstanceInfo {
	_ = "STUB: not implemented"
	return *new(CollectorInstanceInfo)
}

// Size returns the max number of attributes that defines a collector's
// instance information. Can be used to presize the attributes.
func (info CollectorInstanceInfo) Size() int { _ = "STUB: not implemented"; return 0 }

func (info CollectorInstanceInfo) Copy(to pcommon.Map) { _ = "STUB: not implemented"; return }

func keyWithPrefix(key string) string { _ = "STUB: not implemented"; return "" }
