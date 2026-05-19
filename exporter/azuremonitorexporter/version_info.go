// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azuremonitorexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azuremonitorexporter"

import (
	"sync"
)

var (
	once          sync.Once
	cachedVersion string
)

func getCollectorVersion() string { _ = "STUB: not implemented"; return "" }

// Extract the semantic version without metadata.
