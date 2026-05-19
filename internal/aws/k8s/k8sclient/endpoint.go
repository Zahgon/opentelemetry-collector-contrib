// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8sclient // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/k8s/k8sclient"

import (
	"sync"

	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

const (
	typePod = "Pod"
)

type Service struct {
	ServiceName string
	Namespace   string
}

func NewService(name, namespace string) Service { _ = "STUB: not implemented"; return *new(Service) }

type EpClient interface {
	// Get the mapping between pod key and the corresponding service names
	PodKeyToServiceNames() map[string][]string
	// Get the mapping between the service and the number of belonging pods
	ServiceToPodNum() map[Service]int
}

type epClientOption func(*epClient)

func epSyncCheckerOption(checker initialSyncChecker) epClientOption {
	_ = "STUB: not implemented"
	return *new(epClientOption)
}

type epClient struct {
	stopChan chan struct{}
	store    *ObjStore

	stopped bool

	syncChecker initialSyncChecker

	mu                      sync.RWMutex
	podKeyToServiceNamesMap map[string][]string
	serviceToPodNumMap      map[Service]int // only running pods will show behind endpoints
}

func (c *epClient) PodKeyToServiceNames() map[string][]string {
	_ = "STUB: not implemented"
	return nil
}

func (c *epClient) ServiceToPodNum() map[Service]int { _ = "STUB: not implemented"; return nil }

func (c *epClient) refresh() { _ = "STUB: not implemented"; return }

// pod key to service names

// each obj should be a uniq service.
// ignore the service which has 0 pods.

func newEpClient(clientSet kubernetes.Interface, logger *zap.Logger, options ...epClientOption) *epClient {
	_ = "STUB: not implemented"
	return nil
}

// check the init sync for potential connection issue

func (c *epClient) shutdown() { _ = "STUB: not implemented"; return }

func transformFuncEndpoint(obj any) (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

// EndpointSlice uses a label to reference the service

// Fallback to the EndpointSlice name if label is not present

// EndpointSlice has Endpoints field (not Subsets like old Endpoints)

// Check if endpoint is ready

func (*epClient) createEndpointListWatch(client kubernetes.Interface, ns string) cache.ListerWatcher {
	_ = "STUB: not implemented"
	return *new(cache.ListerWatcher)
}
