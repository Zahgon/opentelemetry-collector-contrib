// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package hostmetadata // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/datadog/hostmetadata"

import (
	"time"

	"github.com/DataDog/datadog-agent/pkg/opentelemetry-mapping-go/otlp/attributes/source"
	"go.opentelemetry.io/collector/component"
)

// GetSourceProvider builds the hostname resolution chain.
func GetSourceProvider(set component.TelemetrySettings, configHostname string, timeout time.Duration) (source.Provider, error) {
	_ = "STUB: not implemented"
	return *new(source.Provider), nil
}
