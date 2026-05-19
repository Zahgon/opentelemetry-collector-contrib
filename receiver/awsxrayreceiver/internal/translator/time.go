// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translator // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsxrayreceiver/internal/translator"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

func addStartTime(startTime *float64, span ptrace.Span) { _ = "STUB: not implemented"; return }

func addEndTime(endTime *float64, span ptrace.Span) { _ = "STUB: not implemented"; return }

func floatSecToNanoEpoch(epochSec *float64) pcommon.Timestamp {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp)
}
