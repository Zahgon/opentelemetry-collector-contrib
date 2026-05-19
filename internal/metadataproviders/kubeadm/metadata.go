// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kubeadm // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/kubeadm"

import (
	"context"

	"k8s.io/client-go/kubernetes"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sconfig"
)

// clusterConfiguration represents the relevant fields from kubeadm's ClusterConfiguration
type clusterConfiguration struct {
	ClusterName string `yaml:"clusterName"`
}

type Provider interface {
	// ClusterName returns the current K8S cluster name
	ClusterName(ctx context.Context) (string, error)
	// ClusterUID returns the current K8S cluster UID
	ClusterUID(ctx context.Context) (string, error)
}

type LocalCache struct {
	ClusterName string
	ClusterUID  string
}

type kubeadmProvider struct {
	kubeadmClient       kubernetes.Interface
	configMapName       string
	kubeSystemNamespace string
	cache               LocalCache
}

func NewProvider(configMapName, kubeSystemNamespace string, apiConf k8sconfig.APIConfig) (Provider, error) {
	_ = "STUB: not implemented"
	return *new(Provider), nil
}

func (k *kubeadmProvider) ClusterName(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// kubeadm stores cluster configuration in the ClusterConfiguration key as YAML

func (k *kubeadmProvider) ClusterUID(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
