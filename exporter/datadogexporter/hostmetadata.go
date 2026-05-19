// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !aix

package datadogexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/datadog/hostmetadata"
	datadogconfig "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/datadog/config"
)

// newMetadataConfigfromConfig creates a new metadata pusher config from the main
func newMetadataConfigfromConfig(cfg *datadogconfig.Config) hostmetadata.PusherConfig {
	_ = "STUB: not implemented"
	return *new(hostmetadata.PusherConfig)
}
