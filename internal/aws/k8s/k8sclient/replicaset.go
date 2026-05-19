// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8sclient // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/k8s/k8sclient"

import (
	"sync"
	"time"

	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

const (
	deployment = "Deployment"
)

type ReplicaSetClient interface {
	// Get the mapping between replica set and deployment
	ReplicaSetToDeployment() map[string]string
}

type noOpReplicaSetClient struct{}

func (*noOpReplicaSetClient) ReplicaSetToDeployment() map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (*noOpReplicaSetClient) shutdown() { _ = "STUB: not implemented"; return }

type replicaSetClientOption func(*replicaSetClient)

func replicaSetSyncCheckerOption(checker initialSyncChecker) replicaSetClientOption {
	_ = "STUB: not implemented"
	return *new(replicaSetClientOption)
}

type replicaSetClient struct {
	stopChan chan struct{}
	store    *ObjStore

	stopped     bool
	syncChecker initialSyncChecker

	mu                        sync.RWMutex
	cachedReplicaSetMap       map[string]time.Time
	replicaSetToDeploymentMap map[string]string
}

func (c *replicaSetClient) ReplicaSetToDeployment() map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (c *replicaSetClient) refresh() { _ = "STUB: not implemented"; return }

func newReplicaSetClient(clientSet kubernetes.Interface, logger *zap.Logger, options ...replicaSetClientOption) (*replicaSetClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check the init sync for potential connection issue

func (c *replicaSetClient) shutdown() { _ = "STUB: not implemented"; return }

func transformFuncReplicaSet(obj any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func createReplicaSetListWatch(client kubernetes.Interface, ns string) cache.ListerWatcher {
	_ = "STUB: not implemented"
	return *new(cache.ListerWatcher)
}
