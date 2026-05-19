// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package elasticsearchexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter"

import (
	"time"
)

func generateIndexWithLogstashFormat(index string, conf *LogstashFormatSettings, t time.Time) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
