// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8sapiserver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver/internal/k8sapiserver"

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sync"

	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	"k8s.io/client-go/tools/record"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/k8s/k8sclient"
)

const (
	lockName = "otel-container-insight-clusterleader"
)

// eventBroadcaster is adapted from record.EventBroadcaster
type eventBroadcaster interface {
	// StartRecordingToSink starts sending events received from this EventBroadcaster to the given
	// sink. The return value can be ignored or used to stop recording, if desired.
	StartRecordingToSink(sink record.EventSink) watch.Interface
	// StartLogging starts sending events received from this EventBroadcaster to the given logging
	// function. The return value can be ignored or used to stop recording, if desired.
	StartLogging(logf func(format string, args ...any)) watch.Interface
	// NewRecorder returns an EventRecorder that can be used to send events to this EventBroadcaster
	// with the event source set to the given event source.
	NewRecorder(scheme *runtime.Scheme, source v1.EventSource) record.EventRecorderLogger
}

type K8sClient interface {
	GetClientSet() kubernetes.Interface
	GetEpClient() k8sclient.EpClient
	GetNodeClient() k8sclient.NodeClient
	GetPodClient() k8sclient.PodClient
	ShutdownNodeClient()
	ShutdownPodClient()
}

// K8sAPIServer is a struct that produces metrics from kubernetes api server
type K8sAPIServer struct {
	nodeName            string // get the value from downward API
	logger              *zap.Logger
	clusterNameProvider clusterNameProvider
	cancel              context.CancelFunc

	mu      sync.Mutex
	leading bool

	k8sClient  K8sClient // *k8sclient.K8sClient
	epClient   k8sclient.EpClient
	nodeClient k8sclient.NodeClient
	podClient  k8sclient.PodClient

	// the following can be set to mocks in testing
	broadcaster eventBroadcaster
	// the close of isLeadingC indicates the leader election is done. This is used in testing
	isLeadingC chan bool
}

type clusterNameProvider interface {
	GetClusterName() string
}

type k8sAPIServerOption func(*K8sAPIServer)

// New creates a k8sApiServer which can generate cluster-level metrics
func New(clusterNameProvider clusterNameProvider, logger *zap.Logger, options ...k8sAPIServerOption) (*K8sAPIServer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetMetrics returns an array of metrics
func (k *K8sAPIServer) GetMetrics() []pmetric.Metrics { _ = "STUB: not implemented"; return nil }

// don't generate any metrics if the current collector is not the leader

// don't emit metrics if the cluster name is not detected

//nolint:gocritic //sprintfQuotedString for JSON

//nolint:gocritic //sprintfQuotedString for JSON

func (k *K8sAPIServer) init() error {
	var ctx context.Context
	ctx, k.cancel = context.WithCancel(context.Background())

	k.nodeName = os.Getenv("HOST_NAME")
	if k.nodeName == "" {
		return errors.New("environment variable HOST_NAME is not set in k8s deployment config")
	}

	lockNamespace := os.Getenv("K8S_NAMESPACE")
	if lockNamespace == "" {
		return errors.New("environment variable K8S_NAMESPACE is not set in k8s deployment config")
	}

	clientSet := k.k8sClient.GetClientSet()
	configMapInterface := clientSet.CoreV1().ConfigMaps(lockNamespace)
	if configMap, err := configMapInterface.Get(ctx, lockName, metav1.GetOptions{}); configMap == nil || err != nil {
		k.logger.Info(fmt.Sprintf("Cannot get the leader config map: %v, try to create the config map...", err))
		configMap, err = configMapInterface.Create(ctx,
			&v1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: lockNamespace,
					Name:      lockName,
				},
			}, metav1.CreateOptions{})
		k.logger.Info(fmt.Sprintf("configMap: %v, err: %v", configMap, err))
	}

	lock, err := resourcelock.New(
		resourcelock.LeasesResourceLock,
		lockNamespace, lockName,
		clientSet.CoreV1(),
		clientSet.CoordinationV1(),
		resourcelock.ResourceLockConfig{
			Identity:      k.nodeName,
			EventRecorder: k.createRecorder(lockName, lockNamespace),
		})
	if err != nil {
		k.logger.Warn("Failed to create resource lock", zap.Error(err))
		return err
	}

	go k.startLeaderElection(ctx, lock)

	return nil
}

// Shutdown stops the k8sApiServer
func (k *K8sAPIServer) Shutdown() error { _ = "STUB: not implemented"; return nil }

func (k *K8sAPIServer) startLeaderElection(ctx context.Context, lock resourcelock.Interface) {
	_ = "STUB: not implemented"
	return
}

// IMPORTANT: you MUST ensure that any code you have that
// is protected by the lease must terminate **before**
// you call cancel. Otherwise, you could have a background
// loop still running and another process could
// get elected before your background loop finished, violating
// the stated goal of the lease.

// we're notified when we start

// always retrieve clients in case previous ones shut down during leader switching

// this executes only in testing

// we can do cleanup here, or after the RunOrDie method returns

// node and pod are only used for cluster level metrics, endpoint is used for decorator too.

// when leader election ends, the channel ctx.Done() will be closed

func (k *K8sAPIServer) createRecorder(name, namespace string) record.EventRecorder {
	_ = "STUB: not implemented"
	return *new(record.EventRecorder)
}
