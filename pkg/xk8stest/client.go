// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package xk8stest // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/xk8stest"

import (
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
)

type K8sClient struct {
	DynamicClient   *dynamic.DynamicClient
	DiscoveryClient *discovery.DiscoveryClient
	Mapper          *restmapper.DeferredDiscoveryRESTMapper
	restConfig      *rest.Config
}

func NewK8sClient(kubeconfigPath string) (*K8sClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
