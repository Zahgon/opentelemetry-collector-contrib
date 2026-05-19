// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kubeletutil // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver/internal/stores/kubeletutil"

import (
	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/kubelet"
)

type KubeletClient struct {
	KubeIP     string
	Port       string
	restClient kubelet.Client
}

func NewKubeletClient(kubeIP, port string, logger *zap.Logger) (*KubeletClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// use service account for authentication

func (k *KubeletClient) ListPods() ([]corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
