// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package stores // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver/internal/stores"

import (
	"context"
	"regexp"
	"sync"
	"time"

	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/k8s/k8sclient"
	awsmetrics "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/metrics"
)

const (
	refreshInterval    = 30 * time.Second
	measurementsExpiry = 10 * time.Minute
	podsExpiry         = 2 * time.Minute
	memoryKey          = "memory"
	cpuKey             = "cpu"
	splitRegexStr      = "\\.|-"
	kubeProxy          = "kube-proxy"
)

var re = regexp.MustCompile(splitRegexStr)

type cachedEntry struct {
	pod      corev1.Pod
	creation time.Time
}

type Owner struct {
	OwnerKind string `json:"owner_kind"`
	OwnerName string `json:"owner_name"`
}

type prevPodMeasurement struct {
	containersRestarts int
}

type prevContainerMeasurement struct {
	restarts int
}

type mapWithExpiry struct {
	*awsmetrics.MapWithExpiry
}

func (m *mapWithExpiry) Get(key string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (m *mapWithExpiry) Set(key string, content any) { _ = "STUB: not implemented"; return }

func newMapWithExpiry(ttl time.Duration) *mapWithExpiry { _ = "STUB: not implemented"; return nil }

type replicaSetInfoProvider interface {
	GetReplicaSetClient() k8sclient.ReplicaSetClient
}

type podClient interface {
	ListPods() ([]corev1.Pod, error)
}

type PodStore struct {
	cache            *mapWithExpiry
	prevMeasurements map[string]*mapWithExpiry // preMeasurements per each Type (Pod, Container, etc)
	podClient        podClient
	k8sClient        replicaSetInfoProvider
	lastRefreshed    time.Time
	nodeInfo         *nodeInfo
	prefFullPodName  bool
	logger           *zap.Logger
	sync.Mutex
	addFullPodNameMetricLabel bool
}

func NewPodStore(hostIP string, prefFullPodName, addFullPodNameMetricLabel bool, logger *zap.Logger) (*PodStore, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Try to detect kubelet permission issue here

func (p *PodStore) Shutdown() error { _ = "STUB: not implemented"; return nil }

func (p *PodStore) getPrevMeasurement(metricType, metricKey string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (p *PodStore) setPrevMeasurement(metricType, metricKey string, content any) {
	_ = "STUB: not implemented"
	return
}

// RefreshTick triggers refreshing of the pod store.
// It will be called at relatively short intervals (e.g. 1 second).
// We can't do refresh in regular interval because the Decorate(...) function will
// call refresh(...) on demand when the pod metadata for the given metrics is not in
// cache yet. This will make the refresh interval irregular.
func (p *PodStore) RefreshTick(ctx context.Context) { _ = "STUB: not implemented"; return }

func (p *PodStore) Decorate(ctx context.Context, metric CIMetric, kubernetesBlob map[string]any) bool {
	_ = "STUB: not implemented"
	return false
}

// If pod is still not found, insert a placeholder to avoid too many refresh

// If the entry is not a placeholder, decorate the pod

func (p *PodStore) getCachedEntry(podKey string) *cachedEntry {
	_ = "STUB: not implemented"
	return nil
}

func (p *PodStore) setCachedEntry(podKey string, entry *cachedEntry) {
	_ = "STUB: not implemented"
	return
}

func (p *PodStore) refresh(ctx context.Context, now time.Time) { _ = "STUB: not implemented"; return }

func (p *PodStore) refreshInternal(now time.Time, podList []corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

func (p *PodStore) decorateNode(metric CIMetric) { _ = "STUB: not implemented"; return }

func (p *PodStore) decorateCPU(metric CIMetric, pod *corev1.Pod) { _ = "STUB: not implemented"; return }

// add cpu limit and request for pod cpu

// set podReq to the sum of containerReq which has req

// only set podLimit when all the containers has limit

// add cpu limit and request for container

func (p *PodStore) decorateMem(metric CIMetric, pod *corev1.Pod) { _ = "STUB: not implemented"; return }

// add mem limit and request for pod mem

// set podReq to the sum of containerReq which has req

// only set podLimit when all the containers has limit

// add mem limit and request for container

func (p *PodStore) addStatus(metric CIMetric, pod *corev1.Pod) { _ = "STUB: not implemented"; return }

// It could be used to get limit/request(depend on the passed-in fn) per pod
// return the sum of ResourceSetting and a bool which indicate whether all container set Resource
func getResourceSettingForPod(pod *corev1.Pod, bound uint64, resource corev1.ResourceName, fn func(resource corev1.ResourceName, spec *corev1.Container) (uint64, bool)) (uint64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func getLimitForContainer(resource corev1.ResourceName, spec *corev1.Container) (uint64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// it doesn't make sense for the limits to be negative

func getRequestForContainer(resource corev1.ResourceName, spec *corev1.Container) (uint64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// it doesn't make sense for the requests to be negative

func addContainerID(pod *corev1.Pod, metric CIMetric, kubernetesBlob map[string]any, logger *zap.Logger) {
	_ = "STUB: not implemented"
	return
}

func addLabels(pod *corev1.Pod, kubernetesBlob map[string]any) { _ = "STUB: not implemented"; return }

func getJobNamePrefix(podName string) string { _ = "STUB: not implemented"; return "" }

func (p *PodStore) addPodOwnersAndPodName(metric CIMetric, pod *corev1.Pod, kubernetesBlob map[string]any) {
	_ = "STUB: not implemented"
	return
}

// if podName is not set according to a well-known controllers, then set it to its own name

func addContainerCount(metric CIMetric, pod *corev1.Pod) { _ = "STUB: not implemented"; return }
