// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8sclient // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/k8s/k8sclient"

import (
	"sync"

	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

type PodClient interface {
	// Get the mapping between the namespace and the number of belonging pods
	NamespaceToRunningPodNum() map[string]int
}

type podClientOption func(*podClient)

func podSyncCheckerOption(checker initialSyncChecker) podClientOption {
	_ = "STUB: not implemented"
	return *new(podClientOption)
}

type podClient struct {
	stopChan chan struct{}
	store    *ObjStore

	stopped     bool
	syncChecker initialSyncChecker

	mu                          sync.RWMutex
	namespaceToRunningPodNumMap map[string]int
}

func (c *podClient) NamespaceToRunningPodNum() map[string]int {
	_ = "STUB: not implemented"
	return nil
}

func (c *podClient) refresh() { _ = "STUB: not implemented"; return }

func newPodClient(clientSet kubernetes.Interface, logger *zap.Logger, options ...podClientOption) *podClient {
	_ = "STUB: not implemented"
	return nil
}

// check the init sync for potential connection issue

func (c *podClient) shutdown() { _ = "STUB: not implemented"; return }

func transformFuncPod(obj any) (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func createPodListWatch(client kubernetes.Interface, ns string) cache.ListerWatcher {
	_ = "STUB: not implemented"
	return *new(cache.ListerWatcher)
}
