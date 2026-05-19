// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build linux

package cadvisor // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver/internal/cadvisor"

import (
	cInfo "github.com/google/cadvisor/info/v1"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver/internal/cadvisor/extractors"
)

const (
	// TODO: https://github.com/containerd/cri/issues/922#issuecomment-423729537 the container name can be empty on containerd
	infraContainerName = "POD"
	podNameLabel       = "io.kubernetes.pod.name"
	namespaceLabel     = "io.kubernetes.pod.namespace"
	podIDLabel         = "io.kubernetes.pod.uid"
	containerNameLabel = "io.kubernetes.container.name"
)

// podKey contains information of a pod extracted from containers owned by it.
// containers has label information for a pod and their cgroup path is one level deeper.
type podKey struct {
	cgroupPath   string
	podID        string
	containerIDs []string
	podName      string
	namespace    string
}

func processContainers(cInfos []*cInfo.ContainerInfo, mInfo extractors.CPUMemInfoProvider, containerOrchestrator string, logger *zap.Logger, metricExtractors []extractors.MetricExtractor) []*extractors.CAdvisorMetric {
	_ = "STUB: not implemented"
	return nil
}

// first iteration of container infos processes individual container info and
// gather pod related info for the second iteration which generate pod metric

// Save pod cgroup path we collected from containers under it.

// collect the container ids associated with a pod

// second iteration of container infos generates pod metrics based on the info
// gathered from the first iteration

// This happens when our cgroup path based pod detection logic is not working.

// processContainers get metrics for individual container and gather information for pod so we can look it up later.
func processContainer(info *cInfo.ContainerInfo, mInfo extractors.CPUMemInfoProvider, containerOrchestrator string, logger *zap.Logger, metricExtractors []extractors.MetricExtractor) ([]*extractors.CAdvisorMetric, *podKey, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Only a container has all these three labels set.

// NOTE: containerName can be empty for pause container on containerd
// https://github.com/containerd/cri/issues/922#issuecomment-423729537

// Pod's cgroup path is parent for a container.
// container name: /kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pod04d39715_075e_4c7c_b128_67f7897c05b7.slice/docker-57b3dabd69b94beb462244a0c15c244b509adad0940cdcc67ca079b8208ec1f2.scope
// pod name:       /kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pod04d39715_075e_4c7c_b128_67f7897c05b7.slice/

// For docker, pause container name is set to POD while containerd does not set it.
// See https://github.com/aws/amazon-cloudwatch-agent/issues/188

// NOTE: the pod here is only used by NetMetricExtractor,
// other pod info like CPU, Mem are dealt within in processPod.

// TODO(pvasir): wait for upstream fix https://github.com/google/cadvisor/issues/2785

func processPod(info *cInfo.ContainerInfo, mInfo extractors.CPUMemInfoProvider, podKeys map[string]podKey, logger *zap.Logger, metricExtractors []extractors.MetricExtractor) []*extractors.CAdvisorMetric {
	_ = "STUB: not implemented"
	return nil
}

// Check if it's a container running inside container, caller will drop the metric when return value is true.
// The validation is based on ContainerReference.Name, which is essentially cgroup path.
// The first version is from https://github.com/aws/amazon-cloudwatch-agent/commit/e8daa5f5926c5a5f38e0ceb746c141be463e11e4#diff-599185154c116b295172b56311729990d20672f6659500870997c018ce072100
// But the logic no longer works when docker is using systemd as cgroup driver, because a prefix like `kubepods` is attached to each segment.
// The new name pattern with systemd is
// - Guaranteed /kubepods.slice/kubepods-podc8f7bb69_65f2_4b61_ae5a_9b19ac47a239.slice/docker-523b624a86a2a74c2bedf586d8448c86887ef7858a8dec037d6559e5ad3fccb5.scope
// - Burstable /kubepods.slice/kubepods-besteffort.slice/kubepods-besteffort-podab0e310c_0bdb_48e8_ac87_81a701514645.slice/docker-caa8a5e51cd6610f8f0110b491e8187d23488b9635acccf0355a7975fd3ff158.scope
// - Docker in Docker /kubepods.slice/kubepods-burstable.slice/kubepods-burstable-podc9adcee4_c874_4dad_8bc8_accdbd67ac3a.slice/docker-e58cfbc8b67f6e1af458efdd31cb2a8abdbf9f95db64f4c852b701285a09d40e.scope/docker/fb651068cfbd4bf3d45fb092ec9451f8d1a36b3753687bbaa0a9920617eae5b9
// So we check the number of segments within the cgroup path to determine if it's a container running in container.
func isContainerInContainer(p string) bool { _ = "STUB: not implemented"; return false }

// Without nested container, the number of segments (regardless of cgroupfs/systemd) are either 3 or 4 (depends on QoS)
// /kubepods/pod_id/docker_id
// /kubepods/qos/pod_id/docker_id
// With nested container, the number of segments are either 5 or 6
// /kubepods/pod_id/docker_id/docker/docker_id
// /kubepods/qos/pod_id/docker_id/docker/docker_id
