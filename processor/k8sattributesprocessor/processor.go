// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8sattributesprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/k8sattributesprocessor"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/pprofile"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sconfig"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/k8sattributesprocessor/internal/kube"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/k8sattributesprocessor/internal/metadata"
)

const (
	clientIPLabelName string = "ip"
)

type kubernetesprocessor struct {
	cfg                    component.Config
	options                []option
	telemetrySettings      component.TelemetrySettings
	telemetry              *metadata.TelemetryBuilder
	logger                 *zap.Logger
	apiConfig              k8sconfig.APIConfig
	kc                     kube.Client
	passthroughMode        bool
	rules                  kube.ExtractionRules
	filters                kube.Filters
	podAssociations        []kube.Association
	podIgnore              kube.Excludes
	waitForMetadata        bool
	waitForMetadataTimeout time.Duration
	watchSyncPeriod        time.Duration
}

func (kp *kubernetesprocessor) initKubeClient(set component.TelemetrySettings, kubeClient kube.ClientProvider) error {
	_ = "STUB: not implemented"
	return nil
}

func (kp *kubernetesprocessor) Start(_ context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// This might have been set by an option already

func (kp *kubernetesprocessor) Shutdown(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// processTraces process traces and add k8s metadata using resource IP or incoming IP as pod origin.
func (kp *kubernetesprocessor) processTraces(ctx context.Context, td ptrace.Traces) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

// processMetrics process metrics and add k8s metadata using resource IP, hostname or incoming IP as pod origin.
func (kp *kubernetesprocessor) processMetrics(ctx context.Context, md pmetric.Metrics) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// processLogs process logs and add k8s metadata using resource IP, hostname or incoming IP as pod origin.
func (kp *kubernetesprocessor) processLogs(ctx context.Context, ld plog.Logs) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

// processProfiles process profiles and add k8s metadata using resource IP, hostname or incoming IP as pod origin.
func (kp *kubernetesprocessor) processProfiles(ctx context.Context, pd pprofile.Profiles) (pprofile.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(pprofile.Profiles), nil
}

// processResource adds Pod metadata tags to resource based on pod association configuration.
// signal is the OTLP signal type ("traces", "metrics", "logs", "profiles") used for telemetry.
func (kp *kubernetesprocessor) processResource(ctx context.Context, resource pcommon.Resource, signal string) {
	_ = "STUB: not implemented"
	return
}

// Record failed pod association

// Record failed pod association when no identifier found

func setResourceAttribute(attributes pcommon.Map, key, val string) {
	_ = "STUB: not implemented"
	return
}

func getNamespace(pod *kube.Pod, resAttrs pcommon.Map) string { _ = "STUB: not implemented"; return "" }

func getNodeName(pod *kube.Pod, resAttrs pcommon.Map) string { _ = "STUB: not implemented"; return "" }

func getDeploymentUID(pod *kube.Pod, resAttrs pcommon.Map) string {
	_ = "STUB: not implemented"
	return ""
}

func getStatefulSetUID(pod *kube.Pod, resAttrs pcommon.Map) string {
	_ = "STUB: not implemented"
	return ""
}

func getDaemonSetUID(pod *kube.Pod, resAttrs pcommon.Map) string {
	_ = "STUB: not implemented"
	return ""
}

func getJobUID(pod *kube.Pod, resAttrs pcommon.Map) string { _ = "STUB: not implemented"; return "" }

// addContainerAttributes looks if pod has any container identifiers and adds additional container attributes
func (kp *kubernetesprocessor) addContainerAttributes(attrs pcommon.Map, pod *kube.Pod) {
	_ = "STUB: not implemented"
	return
}

// if there is only one container in the pod, we can fall back to that container

// attempt to get container ID from restart count

// take the highest runID (restart count) which represents the currently running container in most cases

func (kp *kubernetesprocessor) getAttributesForPodsNamespace(namespace string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (kp *kubernetesprocessor) getAttributesForPodsNode(nodeName string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (kp *kubernetesprocessor) getAttributesForPodsDeployment(deploymentUID string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (kp *kubernetesprocessor) getAttributesForPodsStatefulSet(statefulsetUID string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (kp *kubernetesprocessor) getAttributesForPodsDaemonSet(daemonsetUID string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (kp *kubernetesprocessor) getAttributesForPodsJob(jobUID string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (kp *kubernetesprocessor) getUIDForPodsNode(nodeName string) string {
	_ = "STUB: not implemented"
	return ""
}

// intFromAttribute extracts int value from an attribute stored as string or int
func intFromAttribute(val pcommon.Value) (int, error) { _ = "STUB: not implemented"; return 0, nil }
