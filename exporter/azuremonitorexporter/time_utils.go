// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azuremonitorexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azuremonitorexporter"

// Contains code common to both trace and metrics exporters

import (
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
)

func toTime(t pcommon.Timestamp) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// Formats a Duration into the form DD.HH:MM:SS.MMMMMM
func formatDuration(d time.Duration) string { _ = "STUB: not implemented"; return "" }

var timeNow = time.Now
