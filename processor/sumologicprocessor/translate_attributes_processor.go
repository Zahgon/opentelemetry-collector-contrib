// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sumologicprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/sumologicprocessor"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

// translateAttributesProcessor translates attribute names from OpenTelemetry to Sumo Logic convention
type translateAttributesProcessor struct {
	shouldTranslate bool
}

// attributeTranslations maps OpenTelemetry attribute names to Sumo Logic attribute names
var attributeTranslations = map[string]string{
	"cloud.account.id":        "AccountId",
	"cloud.availability_zone": "AvailabilityZone",
	"cloud.platform":          "aws_service",
	"cloud.region":            "Region",
	"host.id":                 "InstanceId",
	"host.name":               "host",
	"host.type":               "InstanceType",
	"k8s.cluster.name":        "Cluster",
	"k8s.container.name":      "container",
	"k8s.daemonset.name":      "daemonset",
	"k8s.deployment.name":     "deployment",
	"k8s.namespace.name":      "namespace",
	"k8s.node.name":           "node",
	"k8s.service.name":        "service",
	"k8s.pod.hostname":        "host",
	"k8s.pod.name":            "pod",
	"k8s.pod.uid":             "pod_id",
	"k8s.replicaset.name":     "replicaset",
	"k8s.statefulset.name":    "statefulset",
	"service.name":            "service",
	"log.file.path_resolved":  "_sourceName",
}

func newTranslateAttributesProcessor(shouldTranslate bool) *translateAttributesProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (proc *translateAttributesProcessor) processLogs(logs plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (proc *translateAttributesProcessor) processMetrics(metrics pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (*translateAttributesProcessor) processTraces(ptrace.Traces) error {
	_ = "STUB: not implemented"
	// No-op. Traces should not be translated.
	return nil
}

func (proc *translateAttributesProcessor) isEnabled() bool { _ = "STUB: not implemented"; return false }

func (*translateAttributesProcessor) ConfigPropertyName() string {
	_ = "STUB: not implemented"
	return ""
}

func translateAttributes(attributes pcommon.Map) { _ = "STUB: not implemented"; return }

// Only insert if it doesn't exist yet to prevent overwriting.
// We have to do it this way since the final return value is not
// ready yet to rely on .Insert() not overwriting.
