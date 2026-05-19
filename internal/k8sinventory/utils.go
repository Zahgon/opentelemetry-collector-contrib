// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8sinventory // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sinventory"

import (
	"time"

	corev1 "k8s.io/api/core/v1"
)

// GetEventTimestamp returns the EventTimestamp based on the populated k8s event timestamps.
// Priority: EventTime > LastTimestamp > FirstTimestamp.
func GetEventTimestamp(ev *corev1.Event) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}
