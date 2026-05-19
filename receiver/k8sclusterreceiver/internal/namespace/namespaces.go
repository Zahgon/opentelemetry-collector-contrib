// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package namespace // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/namespace"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	corev1 "k8s.io/api/core/v1"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/metadata"
)

const (
	// Keys for namespace metadata and entity attributes.
	k8sNamespaceCreationTime = "k8s.namespace.creation_timestamp"
	k8sNamespacePhase        = "k8s.namespace.phase"
)

func RecordMetrics(mb *metadata.MetricsBuilder, ns *corev1.Namespace, ts pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

var namespacePhaseValues = map[corev1.NamespacePhase]int32{
	corev1.NamespaceActive:      1,
	corev1.NamespaceTerminating: 0,
	// If phase is blank for some reason, send as -1 for unknown.
	corev1.NamespacePhase(""): -1,
}

func GetMetadata(ns *corev1.Namespace) map[experimentalmetricmetadata.ResourceID]*metadata.KubernetesMetadata {
	_ = "STUB: not implemented"
	return nil
}
