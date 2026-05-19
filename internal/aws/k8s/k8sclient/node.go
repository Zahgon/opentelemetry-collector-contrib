// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8sclient // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/k8s/k8sclient"

import (
	"sync"

	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

// This needs to be reviewed for newer versions of k8s.
var failedNodeConditions = map[v1.NodeConditionType]bool{
	v1.NodeMemoryPressure:     true,
	v1.NodeDiskPressure:       true,
	v1.NodePIDPressure:        true,
	v1.NodeNetworkUnavailable: true,
}

type NodeClient interface {
	// Get the number of failed nodes for current cluster
	ClusterFailedNodeCount() int
	// Get the number of nodes for current cluster
	ClusterNodeCount() int
}

type nodeClientOption func(*nodeClient)

func nodeSyncCheckerOption(checker initialSyncChecker) nodeClientOption {
	_ = "STUB: not implemented"
	return *new(nodeClientOption)
}

type nodeClient struct {
	stopChan chan struct{}
	store    *ObjStore

	stopped     bool
	syncChecker initialSyncChecker

	mu                     sync.RWMutex
	clusterFailedNodeCount int
	clusterNodeCount       int
}

func (c *nodeClient) ClusterFailedNodeCount() int { _ = "STUB: not implemented"; return 0 }

func (c *nodeClient) ClusterNodeCount() int { _ = "STUB: not implemented"; return 0 }

func (c *nodeClient) refresh() { _ = "STUB: not implemented"; return }

// match the failedNodeConditions type we care about

// if this is not false, i.e. true or unknown

func newNodeClient(clientSet kubernetes.Interface, logger *zap.Logger, options ...nodeClientOption) *nodeClient {
	_ = "STUB: not implemented"
	return nil
}

// check the init sync for potential connection issue

func (c *nodeClient) shutdown() { _ = "STUB: not implemented"; return }

func transformFuncNode(obj any) (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func createNodeListWatch(client kubernetes.Interface) cache.ListerWatcher {
	_ = "STUB: not implemented"
	return *new(cache.ListerWatcher)
}
