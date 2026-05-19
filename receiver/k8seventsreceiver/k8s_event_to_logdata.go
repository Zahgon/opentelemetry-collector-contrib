// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8seventsreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8seventsreceiver"

import (
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
)

const (
	// Number of log attributes to add to the plog.LogRecordSlice.
	totalLogAttributes = 9

	// Number of resource attributes to add to the plog.ResourceLogs.
	totalResourceAttributes = 7
)

// By default k8s event has only two types of events (Normal, Warning), here are we allowing other types as well.
// For more info: https://github.com/kubernetes/api/blob/release-1.34/events/v1/types_swagger_doc_generated.go#L42
var severityMap = map[string]plog.SeverityNumber{
	"normal":   plog.SeverityNumberInfo,
	"warning":  plog.SeverityNumberWarn,
	"error":    plog.SeverityNumberError,
	"critical": plog.SeverityNumberFatal,
}

// k8sEventToLogRecord converts Kubernetes event to plog.LogRecordSlice and adds the resource attributes.
func k8sEventToLogData(logger *zap.Logger, ev *corev1.Event, version string) plog.Logs {
	_ = "STUB: not implemented"
	return *new(plog.Logs)
}

// Attributes related to the object causing the event.

// The Message field contains description about the event,
// which is best suited for the "Body" of the LogRecordSlice.

// Set the "SeverityNumber" and "SeverityText" if a known type of
// severity is found.

// "Count" field of k8s event will be '0' in case it is
// not present in the collected event from k8s.
